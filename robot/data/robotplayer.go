package data

import (
	"steve/entity/cache"
	"steve/external/goldclient"
	"steve/server_pb/gold"
	"steve/server_pb/robot"

	"github.com/Sirupsen/logrus"
)

// RobotPlayer 机器人玩家
type RobotPlayer struct {
	PlayerID      uint64            `protobuf:"varint,1,opt,name=player_id,json=playerId" json:"player_id,omitempty"`
	Coin          uint64            `protobuf:"varint,4,opt,name=coin" json:"coin,omitempty"`
	State         uint64            `protobuf:"varint,5,opt,name=state" json:"state,omitempty"`
	GameIDWinRate map[uint64]uint64 `protobuf:"bytes,6,rep,name=game_id_win_rate,json=gameIdWinRate" json:"game_id_win_rate,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
}

//getRedisLeisureRobotPlayer 从redis 获取 空闲的RobotPlayer
func getRedisLeisureRobotPlayer(robotPlayerIDAll []uint64) ([]*RobotPlayer, []uint64) {
	robotsIDCoins := make([]*RobotPlayer, 0)
	lackRobotsID := make([]uint64, 0) // 没有存入redis的机器人
	for _, robotPlayerID := range robotPlayerIDAll {
		robotPlayerInfo, err := GetRobotFields(robotPlayerID, cache.GameState, RobotPlayerGameIDWinRate)
		if err != nil || len(robotPlayerInfo) != 2 {
			lackRobotsID = append(lackRobotsID, robotPlayerID)
			continue
		}
		robotPlayer := &RobotPlayer{}
		robotPlayer.State = InterToUint64(robotPlayerInfo[cache.GameState]) // 玩家状态
		if robotPlayer.State != uint64(robot.RobotPlayerState_RPS_IDIE) {   //是空闲状态
			continue
		}
		robotPlayer.PlayerID = robotPlayerID
		// 从金币服获取
		gold, err := goldclient.GetGold(robotPlayerID, int16(gold.GoldType_GOLD_COIN))
		if err != nil {
			logrus.WithError(err).Errorf("获取金币失败 robotPlayerID(%v)", robotPlayerIDAll)
			lackRobotsID = append(lackRobotsID, robotPlayerID)
			continue
		}
		robotPlayer.Coin = uint64(gold) // 金币
		if len(robotPlayerInfo[RobotPlayerGameIDWinRate].(string)) == 0 {
			logrus.WithError(err).Errorf("获取游戏对应的胜率失败 robotPlayerID(%v)", robotPlayerIDAll)
			lackRobotsID = append(lackRobotsID, robotPlayerID)
			continue
		}
		robotPlayer.GameIDWinRate = JSONToGameIDWinRate(robotPlayerInfo[RobotPlayerGameIDWinRate].(string)) // 游戏对应的胜率
		robotsIDCoins = append(robotsIDCoins, robotPlayer)
	}
	return robotsIDCoins, lackRobotsID
}

//getMysqlLeisureRobotPlayer 从mysql中获取空闲的玩家,并存入redis
func getMysqlLeisureRobotPlayer(robotsPlayers []*RobotPlayer, lackRobotsID []uint64) []*RobotPlayer {
	log := logrus.WithFields(logrus.Fields{"func_name": "getMysqlLeisureRobotPlayer"})
	failedIDErrMpa := make(map[uint64]error) //存入redis 失败 playerID
	for _, playerID := range lackRobotsID {
		robotPlayer := getMysqlRobotPropByPlayerID(playerID) // 从mysql获取 的一定是空闲的
		// 存入redis
		if err := SetRobotPlayerWatchs(playerID, FmtRobotPlayer(robotPlayer), RedisTimeOut); err != nil {
			failedIDErrMpa[playerID] = err
		}
		robotPlayer.PlayerID = playerID
		robotsPlayers = append(robotsPlayers, robotPlayer)
	}
	if len(failedIDErrMpa) > 0 {
		log = log.WithFields(logrus.Fields{"failedIDErrMpa": failedIDErrMpa})
	}
	log.Info("从mysql获取机器人,并存入redis")
	return robotsPlayers
}

//获取机器人需要的值
func getMysqlRobotFieldValuedAll(robotMap map[int64]*RobotPlayer) error {
	//gameid-winrate 游戏id对应的胜率
	robotsPGs, err := getMysqlRobotGameWinRateAll()
	if err != nil {
		return err
	}
	for _, robot := range robotsPGs {
		if rp := robotMap[robot.Playerid]; rp != nil {
			rp.GameIDWinRate[uint64(robot.Gameid)] = uint64(robot.Winningrate)
			robotMap[robot.Playerid] = rp
		} else {
			robotMap[robot.Playerid] = &RobotPlayer{
				PlayerID:      uint64(robot.Playerid),
				GameIDWinRate: map[uint64]uint64{uint64(robot.Gameid): uint64(robot.Winningrate)},
			}
		}
	}
	return err
}
