package core

import (
	"steve/client_pb/msgid"
	"steve/hall/charge"
	"steve/hall/user"
	"steve/structs/exchanger"

	"github.com/Sirupsen/logrus"
)

func registerHandles(e exchanger.Exchanger) error {

	panicRegister := func(msgID msgid.MsgID, h interface{}) {
		if err := e.RegisterHandle(uint32(msgID), h); err != nil {
			logrus.WithField("msg_id", msgID).Panic(err)
		}
	}
	panicRegister(msgid.MsgID_HALL_GET_PLAYER_INFO_REQ, user.HandleGetPlayerInfoReq)
	panicRegister(msgid.MsgID_HALL_GET_PLAYER_STATE_REQ, user.HandleGetPlayerStateReq)
	panicRegister(msgid.MsgID_HALL_GET_GAME_LIST_INFO_REQ, user.HandleGetGameInfoReq)
	panicRegister(msgid.MsgID_GET_CHARGE_INFO_REQ, charge.HandleGetChargeInfoReq)
	panicRegister(msgid.MsgID_CHARGE_REQ, charge.HandleChargeReq)

	panicRegister(msgid.MsgID_HALL_UPDATE_PLAYER_INFO_REQ, user.HandleUpdatePlayerInoReq)
	panicRegister(msgid.MsgID_HALL_REAL_NAME_REQ, user.HandleRealNameReq)
	panicRegister(msgid.MsgID_HALL_GET_PLAYER_GAME_INFO_REQ, user.HandleGetPlayerGameInfoReq)

	panicRegister(msgid.MsgID_AUTH_CODE_REQ, user.HandleSendAuthCodeReq)
	panicRegister(msgid.MsgID_CHECK_AUTH_CODE_REQ, user.HandleCheckAuthCodeReq)
	panicRegister(msgid.MsgID_GET_BIND_PHONE_REWARD_REQ, user.HandleGetBindphoneRewardInfoReq)
	panicRegister(msgid.MsgID_BIND_PHONE_REQ, user.HandleBindPhoneReq)
	panicRegister(msgid.MsgID_CHANGE_PHONE_REQ, user.HandleChangePhoneReq)

	panicRegister(msgid.MsgID_CHANGE_PASSWORD_REQ, user.HandleChangePasswordReq)
	panicRegister(msgid.MsgID_RESET_PASSWORD_REQ, user.HandleResetPasswordReq)
	panicRegister(msgid.MsgID_CHECK_PASSWORD_REQ, user.HandleCheckPasswordReq)
	return nil
}
