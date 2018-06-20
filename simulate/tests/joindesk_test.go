package tests

import (
	"steve/client_pb/msgId"
	"steve/client_pb/room"
	"steve/simulate/global"
	"steve/simulate/interfaces"
	"steve/simulate/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApplyJoinDesk(t *testing.T) {
	createNtfExpectors := map[int]interfaces.MessageExpector{}
	gameStartNtfExpectors := map[int]interfaces.MessageExpector{}
	for i := 0; i < 4; i++ {
		// 登录用户
		accountID := global.AllocAccountID()
		accountName := utils.GenerateAccountName(accountID)
		player, err := utils.LoginPlayer(accountID, accountName)
		assert.Nil(t, err)
		assert.NotNil(t, player)
		client := player.GetClient()

		createNtfExpector, err := client.ExpectMessage(msgid.MsgID_ROOM_DESK_CREATED_NTF)
		assert.Nil(t, err)
		createNtfExpectors[i] = createNtfExpector

		gameStartNtfExpector, err := client.ExpectMessage(msgid.MsgID_ROOM_START_GAME_NTF)
		assert.Nil(t, err)
		gameStartNtfExpectors[i] = gameStartNtfExpector

		_, err = utils.ApplyJoinDesk(player)
		assert.Nil(t, err)
	}

	for _, e := range createNtfExpectors {
		ntf := &room.RoomDeskCreatedNtf{}
		assert.Nil(t, e.Recv(global.DefaultWaitMessageTime, ntf))
		assert.Equal(t, 4, len(ntf.GetPlayers()))
	}
	for _, e := range gameStartNtfExpectors {
		ntf := &room.RoomStartGameNtf{}
		assert.Nil(t, e.Recv(global.DefaultWaitMessageTime, ntf))
	}
}
