package mailtests

import (
	"steve/client_pb/mailserver"
	"steve/client_pb/msgid"
	"steve/simulate/global"
	"steve/simulate/interfaces"
	"steve/simulate/utils"
	"testing"

	"github.com/Sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

var player interfaces.ClientPlayer
var err error

func init() {
	player, err = utils.LoginNewPlayer()

}

func Test_GetAD(t *testing.T) {

	reqCmd := msgid.MsgID_MAILSVR_GET_AD_REQ
	rspCmd := msgid.MsgID_MAILSVR_GET_AD_RSP
	req := &mailserver.MailSvrGetADReq{}
	rsp := &mailserver.MailSvrGetADRsp{}

	assert.NotNil(t, player)

	player.AddExpectors(rspCmd)
	player.GetClient().SendPackage(utils.CreateMsgHead(reqCmd), req)
	expector := player.GetExpector(rspCmd)

	assert.Nil(t, expector.Recv(global.DefaultWaitMessageTime, rsp))
	assert.Zero(t, rsp.GetErrCode())

	logrus.Debugf("Test_GetAD win:%v", rsp)

}

func Test_GetUnReadMailSum(t *testing.T) {

	reqCmd := msgid.MsgID_MAILSVR_GET_UNREAD_SUM_REQ
	rspCmd := msgid.MsgID_MAILSVR_GET_UNREAD_SUM_RSP
	req := &mailserver.MailSvrGetUnReadSumReq{}
	rsp := &mailserver.MailSvrGetUnReadSumRsp{}

	assert.NotNil(t, player)

	player.AddExpectors(rspCmd)
	player.GetClient().SendPackage(utils.CreateMsgHead(reqCmd), req)
	expector := player.GetExpector(rspCmd)

	assert.Nil(t, expector.Recv(global.DefaultWaitMessageTime, rsp))
	assert.Zero(t, rsp.GetErrCode())

	logrus.Debugf("Test_GetUnReadMailSum win:%v", rsp)

}

var mailId uint64 = 0

func Test_GetMailList(t *testing.T) {
	getMailList(t)
}
func getMailList(t *testing.T) {

	reqCmd := msgid.MsgID_MAILSVR_GET_MAIL_LIST_REQ
	rspCmd := msgid.MsgID_MAILSVR_GET_MAIL_LIST_RSP
	req := &mailserver.MailSvrGetMailListReq{}
	rsp := &mailserver.MailSvrGetMailListRsp{}

	assert.NotNil(t, player)

	player.AddExpectors(rspCmd)
	player.GetClient().SendPackage(utils.CreateMsgHead(reqCmd), req)
	expector := player.GetExpector(rspCmd)

	assert.Nil(t, expector.Recv(global.DefaultWaitMessageTime, rsp))
	assert.Zero(t, rsp.GetErrCode())

	logrus.Debugf("Test_GetMailList win:%v", rsp)

	if len(rsp.MailList) > 0 {
		mailId = rsp.MailList[0].GetMailId()
		logrus.Debugf("Test_GetMailList mailId:%v", mailId)
		//getMailDetail(t,player, id)
	}

}

func Test_GetMailDetail(t *testing.T) {

	reqCmd := msgid.MsgID_MAILSVR_GET_MAIL_DETAIL_REQ
	rspCmd := msgid.MsgID_MAILSVR_GET_MAIL_DETAIL_RSP
	req := &mailserver.MailSvrGetMailDetailReq{}
	rsp := &mailserver.MailSvrGetMailDetailRsp{}
	req.MailId = &mailId

	assert.NotNil(t, player)

	player.AddExpectors(rspCmd)
	player.GetClient().SendPackage(utils.CreateMsgHead(reqCmd), req)
	expector := player.GetExpector(rspCmd)

	assert.Nil(t, expector.Recv(global.DefaultWaitMessageTime, rsp))
	assert.Zero(t, rsp.GetErrCode())

	logrus.Debugf("getMailDetail win:%v", rsp)

}

func Test_SetMailReadTag(t *testing.T) {

	reqCmd := msgid.MsgID_MAILSVR_SET_READ_TAG_REQ
	rspCmd := msgid.MsgID_MAILSVR_SET_READ_TAG_RSP
	req := &mailserver.MailSvrSetReadTagReq{}
	rsp := &mailserver.MailSvrSetReadTagRsp{}
	req.MailId = &mailId

	assert.NotNil(t, player)

	player.AddExpectors(rspCmd)
	player.GetClient().SendPackage(utils.CreateMsgHead(reqCmd), req)
	expector := player.GetExpector(rspCmd)

	assert.Nil(t, expector.Recv(global.DefaultWaitMessageTime, rsp))
	assert.Zero(t, rsp.GetErrCode())

	logrus.Debugf("Test_SetMailReadTag win:%v", rsp)

	reqCmd = msgid.MsgID_MAILSVR_AWARD_ATTACH_REQ
	rspCmd = msgid.MsgID_MAILSVR_AWARD_ATTACH_RSP
	req2 := &mailserver.MailSvrAwardAttachReq{}
	rsp2 := &mailserver.MailSvrAwardAttachRsp{}
	req2.MailId = &mailId
	player.AddExpectors(rspCmd)
	player.GetClient().SendPackage(utils.CreateMsgHead(reqCmd), req2)
	expector = player.GetExpector(rspCmd)

	assert.Nil(t, expector.Recv(global.DefaultWaitMessageTime, rsp2))
	assert.Zero(t, rsp2.GetErrCode())

	reqCmd = msgid.MsgID_MAILSVR_DEL_MAIL_REQ
	rspCmd = msgid.MsgID_MAILSVR_DEL_MAIL_RSP
	req3 := &mailserver.MailSvrDelMailReq{}
	rsp3 := &mailserver.MailSvrDelMailRsp{}
	req3.MailId = &mailId
	player.AddExpectors(rspCmd)
	player.GetClient().SendPackage(utils.CreateMsgHead(reqCmd), req3)
	expector = player.GetExpector(rspCmd)

	assert.Nil(t, expector.Recv(global.DefaultWaitMessageTime, rsp3))
	assert.Zero(t, rsp3.GetErrCode())

	getMailList(t)

}
