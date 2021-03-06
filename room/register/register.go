package registers

import (
	"steve/client_pb/msgid"
	"steve/structs/exchanger"
)

// RegisterHandlers 注册消息处理器
func RegisterHandlers(e exchanger.Exchanger) {
	registe := func(id msgid.MsgID, handler interface{}) {
		err := e.RegisterHandle(uint32(id), handler)
		if err != nil {
			panic(err)
		}
	}

	// registe(msgid.MsgID_ROOM_JOIN_DESK_REQ, deskapply.HandleRoomJoinDeskReq)          // 加入牌桌请求
	registe(msgid.MsgID_ROOM_DESK_QUIT_REQ, HandleRoomDeskQuitReq) // 退出牌桌请求
	// registe(msgid.MsgID_ROOM_DESK_CONTINUE_REQ, deskapply.HandleRoomContinueReq)      // 续局请求
	registe(msgid.MsgID_ROOM_CANCEL_TUOGUAN_REQ, HandleCancelTuoGuanReq) // 取消托管请求
	registe(msgid.MsgID_ROOM_RESUME_GAME_REQ, HandleResumeGameReq)       // 恢复对局请求
	registe(msgid.MsgID_ROOM_TUOGUAN_REQ, HandleTuoGuanReq)              // 托管/取消托管请求
	registe(msgid.MsgID_ROOM_AUTOHU_REQ, HandleAutoHuReq)                // 自动胡牌请求
	registe(msgid.MsgID_ROOM_CHAT_REQ, HandleRoomChatReq)                // 房间玩家聊天请求
	// registe(msgid.MsgID_ROOM_DESK_NEED_RESUME_REQ, deskapply.HandleRoomNeedResumeReq) // 是否需要恢复对局请求
	// registe(msgid.MsgID_ROOM_CHANGE_PLAYERS_REQ, deskapply.HandleRoomChangePlayerReq) // 换对手请求
	registe(msgid.MsgID_ROOM_USE_PROP_REQ, HandleUsePropReq) // 使用道具请求
	registe(msgid.MsgID_MATCH_CONTINUE_REQ, HandleContinueReq)
	registe(msgid.MsgID_ROOM_PLAYER_GIVEUP_REQ, HandlePlayerGameGiveUp)                  // 游戏内认输请求
	registe(msgid.MsgID_ROOM_BROKER_PLAYER_CONTINUE_REQ, HandleRoomBrokerPlayerContinue) // 破产玩家继续游戏请求

	RegisterRoomReqHandlers(e)
}
