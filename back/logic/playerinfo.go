package logic

import (
	"fmt"
	"math"
	"steve/back/data"
	"steve/entity/gamelog"
	"steve/external/robotclient"
	"time"

	"github.com/Sirupsen/logrus"
)

func updatePlayerInfo(detailInfo gamelog.TGameDetail) error {
	playerID := detailInfo.Playerid
	gameID := detailInfo.Gameid
	// 获取玩家游戏信息
	playerGame, err := data.GetTPlayerGame(gameID, playerID)
	playerGame.Playerid = int64(playerID)
	playerGame.Gameid = gameID
	if err != nil {
		return err
	}

	logrus.Debugf("获取玩家playerId:(%d),gameId:(%d)的游戏信息：(%v)", playerID, gameID, playerGame)

	//总局数+1
	playerGame.Totalbureau = playerGame.Totalbureau + 1

	// 该玩家此游戏的历史连胜记录key
	winStreamKey := fmt.Sprintf("win_stream:%v@%v", playerID, gameID)

	// 该玩家此游戏的历史连胜记录
	winStream, _ := data.GetPlayerMaxwinningstream(winStreamKey)

	logrus.Debugf("获取玩家playerId:(%d),gameId:(%d)的历史连胜记录：(%v)", playerID, gameID, winStream)

	// 赢家
	if detailInfo.Amount > 0 {
		//胜局+1
		playerGame.Winningburea = playerGame.Winningburea + 1
		//连胜+1
		winStream = winStream + 1
	} else {
		//输了，连胜终结
		winStream = 0
	}
	logrus.Debugf("获取玩家playerId:(%d),gameId:(%d)的新连胜记录：(%v)", playerID, gameID, winStream)

	// 连胜存入redis
	if err := data.SetPlayerMaxwinningstream(winStreamKey, winStream); err != nil {
		logrus.Errorf("failed set maxSream to redis,err:%v", err)
	}

	// 若连胜超过玩家的最高连胜记录，更新
	if winStream > playerGame.Maxwinningstream {
		playerGame.Maxwinningstream = winStream
	}
	// 若倍数超过玩家的最高赢牌倍数，更新
	if int(detailInfo.MaxTimes) > playerGame.Maxmultiple {
		playerGame.Maxmultiple = int(detailInfo.MaxTimes)
	}

	// 更新胜率
	playerGame.Winningrate = math.Trunc((float64(playerGame.Winningburea)/float64(playerGame.Totalbureau))*1e4+0.5) * 1e-4 * 100

	flag, _ := robotclient.IsRobotPlayer(uint64(playerGame.Playerid))
	if flag {
		// 更新机器人的胜率
		robotclient.UpdataRobotPlayerWinRate(uint64(playerGame.Playerid), int32(playerGame.Gameid), playerGame.Winningrate)
	}

	// 创建时间
	playerGame.Createtime = time.Now()

	if err := data.UpdateTPlayerGame(playerGame); err != nil {
		return err
	}
	if err := data.UpdatePlayerGameToredis(playerGame); err != nil {
		return err
	}
	return nil
}
