// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msgid.proto

/*
Package msgid is a generated protocol buffer package.

It is generated from these files:
	msgid.proto

It has these top-level messages:
*/
package msgid

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// MsgID 消息 ID
type MsgID int32

const (
	MsgID_LOGIN_AUTH_REQ                MsgID = 1
	MsgID_LOGIN_AUTH_RSP                MsgID = 2
	MsgID_GATE_AUTH_REQ                 MsgID = 4097
	MsgID_GATE_AUTH_RSP                 MsgID = 4098
	MsgID_GATE_HEART_BEAT_REQ           MsgID = 4099
	MsgID_GATE_HEART_BEAT_RSP           MsgID = 4100
	MsgID_GATE_ANOTHER_LOGIN_NTF        MsgID = 4101
	MsgID_GATE_TRANSMIT_HTTP_REQ        MsgID = 4113
	MsgID_GATE_TRANSMIT_HTTP_RSP        MsgID = 4114
	MsgID_MATCH_REQ                     MsgID = 8193
	MsgID_MATCH_RSP                     MsgID = 8194
	MsgID_MATCH_SUC_CREATE_DESK_NTF     MsgID = 8198
	MsgID_CANCEL_MATCH_REQ              MsgID = 8199
	MsgID_CANCEL_MATCH_RSP              MsgID = 8200
	MsgID_HALL_GET_PLAYER_INFO_REQ      MsgID = 12289
	MsgID_HALL_GET_PLAYER_INFO_RSP      MsgID = 12290
	MsgID_HALL_GET_PLAYER_STATE_REQ     MsgID = 12291
	MsgID_HALL_GET_PLAYER_STATE_RSP     MsgID = 12292
	MsgID_HALL_GET_GAME_LIST_INFO_REQ   MsgID = 12293
	MsgID_HALL_GET_GAME_LIST_INFO_RSP   MsgID = 12294
	MsgID_HALL_REAL_NAME_REQ            MsgID = 12295
	MsgID_HALL_REAL_NAME_RSP            MsgID = 12296
	MsgID_HALL_UPDATE_PLAYER_INFO_REQ   MsgID = 12297
	MsgID_HALL_UPDATE_PLAYER_INFO_RSP   MsgID = 12298
	MsgID_GET_CHARGE_INFO_REQ           MsgID = 12299
	MsgID_GET_CHARGE_INFO_RSP           MsgID = 12300
	MsgID_CHARGE_REQ                    MsgID = 12301
	MsgID_CHARGE_RSP                    MsgID = 12302
	MsgID_HALL_GET_PLAYER_GAME_INFO_REQ MsgID = 12303
	MsgID_HALL_GET_PLAYER_GAME_INFO_RSP MsgID = 12304
	MsgID_AUTH_CODE_REQ                 MsgID = 12305
	MsgID_AUTH_CODE_RSP                 MsgID = 12306
	MsgID_CHECK_AUTH_CODE_REQ           MsgID = 12307
	MsgID_CHECK_AUTH_CODE_RSP           MsgID = 12308
	MsgID_GET_BIND_PHONE_REWARD_REQ     MsgID = 12309
	MsgID_GET_BIND_PHONE_REWARD_RSP     MsgID = 12310
	MsgID_BIND_PHONE_REQ                MsgID = 12311
	MsgID_BIND_PHONE_RSP                MsgID = 12312
	MsgID_CHANGE_PHONE_REQ              MsgID = 12313
	MsgID_CHANGE_PHONE_RSP              MsgID = 12320
	MsgID_CHECK_PASSWORD_REQ            MsgID = 12321
	MsgID_CHECK_PASSWORD_RSP            MsgID = 12322
	MsgID_CHANGE_PASSWORD_REQ           MsgID = 12323
	MsgID_CHANGE_PASSWORD_RSP           MsgID = 12324
	MsgID_RESET_PASSWORD_REQ            MsgID = 12325
	MsgID_RESET_PASSWORD_RSP            MsgID = 12326
	MsgID_MONEY_CHANGE_NTF              MsgID = 12327
	// MSGSVR_BEGIN 0x4001 -->msgserver/msgserver.proto
	MsgID_MSGSVR_GET_HORSE_RACE_REQ    MsgID = 16385
	MsgID_MSGSVR_GET_HORSE_RACE_RSP    MsgID = 16386
	MsgID_MSGSVR_HORSE_RACE_UPDATE_NTF MsgID = 16387
	// MAILSVR_BEGIN 0x5001 -->mailserver/mailserver.proto
	MsgID_MAILSVR_GET_UNREAD_SUM_REQ  MsgID = 20481
	MsgID_MAILSVR_GET_UNREAD_SUM_RSP  MsgID = 20482
	MsgID_MAILSVR_GET_MAIL_LIST_REQ   MsgID = 20483
	MsgID_MAILSVR_GET_MAIL_LIST_RSP   MsgID = 20484
	MsgID_MAILSVR_GET_MAIL_DETAIL_REQ MsgID = 20485
	MsgID_MAILSVR_GET_MAIL_DETAIL_RSP MsgID = 20486
	MsgID_MAILSVR_SET_READ_TAG_REQ    MsgID = 20487
	MsgID_MAILSVR_SET_READ_TAG_RSP    MsgID = 20488
	MsgID_MAILSVR_DEL_MAIL_REQ        MsgID = 20489
	MsgID_MAILSVR_DEL_MAIL_RSP        MsgID = 20490
	MsgID_MAILSVR_AWARD_ATTACH_REQ    MsgID = 20491
	MsgID_MAILSVR_AWARD_ATTACH_RSP    MsgID = 20492
	MsgID_MAILSVR_GET_AD_REQ          MsgID = 20493
	MsgID_MAILSVR_GET_AD_RSP          MsgID = 20494
	// ALMS_BEGIN 0x6001 -->msgserver/msgserver.proto
	MsgID_ALMS_GET_GOLD_REQ          MsgID = 24577
	MsgID_ALMS_GET_GOLD_RSP          MsgID = 24578
	MsgID_ALMS_LOGIN_GOLD_CONFIG_NTF MsgID = 24579
	MsgID_PACKSACK_INFO_REQ          MsgID = 24580
	MsgID_PACKSACK_INFO_RSP          MsgID = 24581
	MsgID_PACKSACK_GOLD_REQ          MsgID = 24582
	MsgID_PACKSACK_GOLD_RSP          MsgID = 24583
	MsgID_PACKSACK_GET_GOLD_REQ      MsgID = 24584
	MsgID_PACKSACK_GET_GOLD_RSP      MsgID = 24585
	// ROOM_BEGIN 消息区间开始
	MsgID_ROOM_BEGIN MsgID = 65536
	// ROOM_BASE_BEGIN 房间逻辑消息区间开始
	MsgID_ROOM_BASE_BEGIN                 MsgID = 65537
	MsgID_ROOM_LOGIN_REQ                  MsgID = 65538
	MsgID_ROOM_LOGIN_RSP                  MsgID = 65539
	MsgID_ROOM_JOIN_DESK_REQ              MsgID = 65540
	MsgID_ROOM_JOIN_DESK_RSP              MsgID = 65541
	MsgID_ROOM_DESK_CREATED_NTF           MsgID = 65542
	MsgID_ROOM_DESK_QUIT_REQ              MsgID = 65543
	MsgID_ROOM_DESK_DISMISS_NTF           MsgID = 65544
	MsgID_ROOM_DESK_CONTINUE_REQ          MsgID = 65545
	MsgID_ROOM_DESK_CONTINUE_RSP          MsgID = 65546
	MsgID_ROOM_VISITOR_LOGIN_REQ          MsgID = 65547
	MsgID_ROOM_VISITOR_LOGIN_RSP          MsgID = 65548
	MsgID_ROOM_DESK_NEED_RESUME_REQ       MsgID = 65549
	MsgID_ROOM_DESK_NEED_RESUME_RSP       MsgID = 65550
	MsgID_ROOM_DESK_QUIT_RSP              MsgID = 65551
	MsgID_ROOM_PLAYER_LOCATION_REQ        MsgID = 65552
	MsgID_ROOM_PLAYER_LOCATION_RSP        MsgID = 65553
	MsgID_ROOM_DESK_QUIT_ENTER_NTF        MsgID = 65554
	MsgID_ROOM_CHANGE_PLAYERS_REQ         MsgID = 65556
	MsgID_ROOM_CHANGE_PLAYERS_RSP         MsgID = 65557
	MsgID_ROOM_PLAYER_GIVEUP_REQ          MsgID = 65792
	MsgID_ROOM_PLAYER_GIVEUP_RSP          MsgID = 65793
	MsgID_ROOM_PLAYER_GIVEUP_NTF          MsgID = 65794
	MsgID_ROOM_BROKER_PLAYER_CONTINUE_REQ MsgID = 65795
	MsgID_ROOM_BROKER_PLAYER_CONTINUE_RSP MsgID = 65796
	MsgID_ROOM_USE_PROP_REQ               MsgID = 65808
	MsgID_ROOM_USE_PROP_RSP               MsgID = 65809
	MsgID_ROOM_USE_PROP_NTF               MsgID = 65810
	// ROOM_BASE_END ROOM 房间逻辑消息区间结束
	MsgID_ROOM_BASE_END MsgID = 69631
	// ROOM_GAME_BEGIN ROOM 游戏逻辑消息区间开始
	MsgID_ROOM_GAME_BEGIN              MsgID = 69632
	MsgID_ROOM_START_GAME_NTF          MsgID = 69633
	MsgID_ROOM_XIPAI_NTF               MsgID = 69634
	MsgID_ROOM_FAPAI_NTF               MsgID = 69635
	MsgID_ROOM_HUANSANZHANG_NTF        MsgID = 69664
	MsgID_ROOM_HUANSANZHANG_REQ        MsgID = 69665
	MsgID_ROOM_HUANSANZHANG_RSP        MsgID = 69666
	MsgID_ROOM_HUANSANZHANG_FINISH_NTF MsgID = 69667
	MsgID_ROOM_DINGQUE_NTF             MsgID = 69696
	MsgID_ROOM_DINGQUE_REQ             MsgID = 69697
	MsgID_ROOM_DINGQUE_RSP             MsgID = 69698
	MsgID_ROOM_DINGQUE_FINISH_NTF      MsgID = 69699
	MsgID_ROOM_CHUPAIWENXUN_NTF        MsgID = 69728
	MsgID_ROOM_XINGPAI_ACTION_REQ      MsgID = 69744
	// ROOM_PENG_REQ = 0x11080;	// 请求碰 客户端->服务器
	MsgID_ROOM_PENG_RSP MsgID = 69761
	MsgID_ROOM_PENG_NTF MsgID = 69762
	// ROOM_GANG_REQ = 0x110A0;	// 请求杠 客户端->服务器
	MsgID_ROOM_GANG_RSP MsgID = 69793
	MsgID_ROOM_GANG_NTF MsgID = 69794
	// ROOM_HU_REQ = 0x110C0;	   // 胡请求 客户端->服务端
	MsgID_ROOM_HU_RSP MsgID = 69825
	MsgID_ROOM_HU_NTF MsgID = 69826
	// ROOM_QI_REQ = 0x110E0;	   // 请求弃 客户端->服务端
	MsgID_ROOM_QI_RSP               MsgID = 69857
	MsgID_ROOM_CHI_RSP              MsgID = 69872
	MsgID_ROOM_CHI_NTF              MsgID = 69873
	MsgID_ROOM_ZIXUN_NTF            MsgID = 69888
	MsgID_ROOM_CHUPAI_REQ           MsgID = 69889
	MsgID_ROOM_CHUPAI_NTF           MsgID = 69890
	MsgID_ROOM_MOPAI_NTF            MsgID = 69891
	MsgID_ROOM_BUHUA_NTF            MsgID = 69892
	MsgID_ROOM_WAIT_QIANGGANGHU_NTF MsgID = 69920
	MsgID_ROOM_TINGINFO_NTF         MsgID = 69921
	MsgID_ROOM_CARTOON_FINISH_REQ   MsgID = 69925
	MsgID_ROOM_GAMEOVER_NTF         MsgID = 69952
	MsgID_ROOM_RESUME_GAME_REQ      MsgID = 69969
	MsgID_ROOM_RESUME_GAME_RSP      MsgID = 69970
	// ROOM_GAME_END 游戏逻辑消息区间结束
	MsgID_ROOM_GAME_END MsgID = 73727
	// ROOM_SETTLE_BEGIN ROOM 游戏结算消息区间开始
	MsgID_ROOM_SETTLE_BEGIN   MsgID = 73728
	MsgID_ROOM_INSTANT_SETTLE MsgID = 73729
	MsgID_ROOM_ROUND_SETTLE   MsgID = 73730
	// ROOM_SETTLE_END 游戏结算消息区间结束
	MsgID_ROOM_SETTLE_END MsgID = 77823
	// 托管逻辑区间开始
	MsgID_ROOM_TUOGUAN_BEGIN      MsgID = 77824
	MsgID_ROOM_TUOGUAN_NTF        MsgID = 77825
	MsgID_ROOM_CANCEL_TUOGUAN_REQ MsgID = 77826
	MsgID_ROOM_TUOGUAN_REQ        MsgID = 77827
	// 托管逻辑区间结束
	MsgID_ROOM_TUOGUAN_END    MsgID = 78079
	MsgID_ROOM_AUTOHU_REQ     MsgID = 78080
	MsgID_ROOM_AUTOHU_RSP     MsgID = 78081
	MsgID_ROOM_COUNT_DOWN_NTF MsgID = 78096
	// 房间聊天
	MsgID_ROOM_CHAT_REQ                  MsgID = 81920
	MsgID_ROOM_CHAT_NTF                  MsgID = 81921
	MsgID_ROOM_DDZ_START_GAME_NTF        MsgID = 86016
	MsgID_ROOM_DDZ_DEAL_NTF              MsgID = 86017
	MsgID_ROOM_DDZ_GRAB_LORD_REQ         MsgID = 86018
	MsgID_ROOM_DDZ_GRAB_LORD_RSP         MsgID = 86019
	MsgID_ROOM_DDZ_GRAB_LORD_NTF         MsgID = 86020
	MsgID_ROOM_DDZ_LORD_NTF              MsgID = 86021
	MsgID_ROOM_DDZ_DOUBLE_REQ            MsgID = 86022
	MsgID_ROOM_DDZ_DOUBLE_RSP            MsgID = 86023
	MsgID_ROOM_DDZ_DOUBLE_NTF            MsgID = 86024
	MsgID_ROOM_DDZ_PLAY_CARD_REQ         MsgID = 86025
	MsgID_ROOM_DDZ_PLAY_CARD_RSP         MsgID = 86026
	MsgID_ROOM_DDZ_PLAY_CARD_NTF         MsgID = 86027
	MsgID_ROOM_DDZ_GAME_OVER_NTF         MsgID = 86028
	MsgID_ROOM_DDZ_RESUME_REQ            MsgID = 86032
	MsgID_ROOM_DDZ_RESUME_RSP            MsgID = 86033
	MsgID_MATCH_CONTINUE_REQ             MsgID = 87552
	MsgID_MATCH_CONTINUE_RSP             MsgID = 87553
	MsgID_MATCH_CONTINUE_DESK_DIMISS_NTF MsgID = 87554
	// ROOM END 房间消息区间结束
	MsgID_ROOM_END MsgID = 131071
)

var MsgID_name = map[int32]string{
	1:      "LOGIN_AUTH_REQ",
	2:      "LOGIN_AUTH_RSP",
	4097:   "GATE_AUTH_REQ",
	4098:   "GATE_AUTH_RSP",
	4099:   "GATE_HEART_BEAT_REQ",
	4100:   "GATE_HEART_BEAT_RSP",
	4101:   "GATE_ANOTHER_LOGIN_NTF",
	4113:   "GATE_TRANSMIT_HTTP_REQ",
	4114:   "GATE_TRANSMIT_HTTP_RSP",
	8193:   "MATCH_REQ",
	8194:   "MATCH_RSP",
	8198:   "MATCH_SUC_CREATE_DESK_NTF",
	8199:   "CANCEL_MATCH_REQ",
	8200:   "CANCEL_MATCH_RSP",
	12289:  "HALL_GET_PLAYER_INFO_REQ",
	12290:  "HALL_GET_PLAYER_INFO_RSP",
	12291:  "HALL_GET_PLAYER_STATE_REQ",
	12292:  "HALL_GET_PLAYER_STATE_RSP",
	12293:  "HALL_GET_GAME_LIST_INFO_REQ",
	12294:  "HALL_GET_GAME_LIST_INFO_RSP",
	12295:  "HALL_REAL_NAME_REQ",
	12296:  "HALL_REAL_NAME_RSP",
	12297:  "HALL_UPDATE_PLAYER_INFO_REQ",
	12298:  "HALL_UPDATE_PLAYER_INFO_RSP",
	12299:  "GET_CHARGE_INFO_REQ",
	12300:  "GET_CHARGE_INFO_RSP",
	12301:  "CHARGE_REQ",
	12302:  "CHARGE_RSP",
	12303:  "HALL_GET_PLAYER_GAME_INFO_REQ",
	12304:  "HALL_GET_PLAYER_GAME_INFO_RSP",
	12305:  "AUTH_CODE_REQ",
	12306:  "AUTH_CODE_RSP",
	12307:  "CHECK_AUTH_CODE_REQ",
	12308:  "CHECK_AUTH_CODE_RSP",
	12309:  "GET_BIND_PHONE_REWARD_REQ",
	12310:  "GET_BIND_PHONE_REWARD_RSP",
	12311:  "BIND_PHONE_REQ",
	12312:  "BIND_PHONE_RSP",
	12313:  "CHANGE_PHONE_REQ",
	12320:  "CHANGE_PHONE_RSP",
	12321:  "CHECK_PASSWORD_REQ",
	12322:  "CHECK_PASSWORD_RSP",
	12323:  "CHANGE_PASSWORD_REQ",
	12324:  "CHANGE_PASSWORD_RSP",
	12325:  "RESET_PASSWORD_REQ",
	12326:  "RESET_PASSWORD_RSP",
	12327:  "MONEY_CHANGE_NTF",
	16385:  "MSGSVR_GET_HORSE_RACE_REQ",
	16386:  "MSGSVR_GET_HORSE_RACE_RSP",
	16387:  "MSGSVR_HORSE_RACE_UPDATE_NTF",
	20481:  "MAILSVR_GET_UNREAD_SUM_REQ",
	20482:  "MAILSVR_GET_UNREAD_SUM_RSP",
	20483:  "MAILSVR_GET_MAIL_LIST_REQ",
	20484:  "MAILSVR_GET_MAIL_LIST_RSP",
	20485:  "MAILSVR_GET_MAIL_DETAIL_REQ",
	20486:  "MAILSVR_GET_MAIL_DETAIL_RSP",
	20487:  "MAILSVR_SET_READ_TAG_REQ",
	20488:  "MAILSVR_SET_READ_TAG_RSP",
	20489:  "MAILSVR_DEL_MAIL_REQ",
	20490:  "MAILSVR_DEL_MAIL_RSP",
	20491:  "MAILSVR_AWARD_ATTACH_REQ",
	20492:  "MAILSVR_AWARD_ATTACH_RSP",
	20493:  "MAILSVR_GET_AD_REQ",
	20494:  "MAILSVR_GET_AD_RSP",
	24577:  "ALMS_GET_GOLD_REQ",
	24578:  "ALMS_GET_GOLD_RSP",
	24579:  "ALMS_LOGIN_GOLD_CONFIG_NTF",
	24580:  "PACKSACK_INFO_REQ",
	24581:  "PACKSACK_INFO_RSP",
	24582:  "PACKSACK_GOLD_REQ",
	24583:  "PACKSACK_GOLD_RSP",
	24584:  "PACKSACK_GET_GOLD_REQ",
	24585:  "PACKSACK_GET_GOLD_RSP",
	65536:  "ROOM_BEGIN",
	65537:  "ROOM_BASE_BEGIN",
	65538:  "ROOM_LOGIN_REQ",
	65539:  "ROOM_LOGIN_RSP",
	65540:  "ROOM_JOIN_DESK_REQ",
	65541:  "ROOM_JOIN_DESK_RSP",
	65542:  "ROOM_DESK_CREATED_NTF",
	65543:  "ROOM_DESK_QUIT_REQ",
	65544:  "ROOM_DESK_DISMISS_NTF",
	65545:  "ROOM_DESK_CONTINUE_REQ",
	65546:  "ROOM_DESK_CONTINUE_RSP",
	65547:  "ROOM_VISITOR_LOGIN_REQ",
	65548:  "ROOM_VISITOR_LOGIN_RSP",
	65549:  "ROOM_DESK_NEED_RESUME_REQ",
	65550:  "ROOM_DESK_NEED_RESUME_RSP",
	65551:  "ROOM_DESK_QUIT_RSP",
	65552:  "ROOM_PLAYER_LOCATION_REQ",
	65553:  "ROOM_PLAYER_LOCATION_RSP",
	65554:  "ROOM_DESK_QUIT_ENTER_NTF",
	65556:  "ROOM_CHANGE_PLAYERS_REQ",
	65557:  "ROOM_CHANGE_PLAYERS_RSP",
	65792:  "ROOM_PLAYER_GIVEUP_REQ",
	65793:  "ROOM_PLAYER_GIVEUP_RSP",
	65794:  "ROOM_PLAYER_GIVEUP_NTF",
	65795:  "ROOM_BROKER_PLAYER_CONTINUE_REQ",
	65796:  "ROOM_BROKER_PLAYER_CONTINUE_RSP",
	65808:  "ROOM_USE_PROP_REQ",
	65809:  "ROOM_USE_PROP_RSP",
	65810:  "ROOM_USE_PROP_NTF",
	69631:  "ROOM_BASE_END",
	69632:  "ROOM_GAME_BEGIN",
	69633:  "ROOM_START_GAME_NTF",
	69634:  "ROOM_XIPAI_NTF",
	69635:  "ROOM_FAPAI_NTF",
	69664:  "ROOM_HUANSANZHANG_NTF",
	69665:  "ROOM_HUANSANZHANG_REQ",
	69666:  "ROOM_HUANSANZHANG_RSP",
	69667:  "ROOM_HUANSANZHANG_FINISH_NTF",
	69696:  "ROOM_DINGQUE_NTF",
	69697:  "ROOM_DINGQUE_REQ",
	69698:  "ROOM_DINGQUE_RSP",
	69699:  "ROOM_DINGQUE_FINISH_NTF",
	69728:  "ROOM_CHUPAIWENXUN_NTF",
	69744:  "ROOM_XINGPAI_ACTION_REQ",
	69761:  "ROOM_PENG_RSP",
	69762:  "ROOM_PENG_NTF",
	69793:  "ROOM_GANG_RSP",
	69794:  "ROOM_GANG_NTF",
	69825:  "ROOM_HU_RSP",
	69826:  "ROOM_HU_NTF",
	69857:  "ROOM_QI_RSP",
	69872:  "ROOM_CHI_RSP",
	69873:  "ROOM_CHI_NTF",
	69888:  "ROOM_ZIXUN_NTF",
	69889:  "ROOM_CHUPAI_REQ",
	69890:  "ROOM_CHUPAI_NTF",
	69891:  "ROOM_MOPAI_NTF",
	69892:  "ROOM_BUHUA_NTF",
	69920:  "ROOM_WAIT_QIANGGANGHU_NTF",
	69921:  "ROOM_TINGINFO_NTF",
	69925:  "ROOM_CARTOON_FINISH_REQ",
	69952:  "ROOM_GAMEOVER_NTF",
	69969:  "ROOM_RESUME_GAME_REQ",
	69970:  "ROOM_RESUME_GAME_RSP",
	73727:  "ROOM_GAME_END",
	73728:  "ROOM_SETTLE_BEGIN",
	73729:  "ROOM_INSTANT_SETTLE",
	73730:  "ROOM_ROUND_SETTLE",
	77823:  "ROOM_SETTLE_END",
	77824:  "ROOM_TUOGUAN_BEGIN",
	77825:  "ROOM_TUOGUAN_NTF",
	77826:  "ROOM_CANCEL_TUOGUAN_REQ",
	77827:  "ROOM_TUOGUAN_REQ",
	78079:  "ROOM_TUOGUAN_END",
	78080:  "ROOM_AUTOHU_REQ",
	78081:  "ROOM_AUTOHU_RSP",
	78096:  "ROOM_COUNT_DOWN_NTF",
	81920:  "ROOM_CHAT_REQ",
	81921:  "ROOM_CHAT_NTF",
	86016:  "ROOM_DDZ_START_GAME_NTF",
	86017:  "ROOM_DDZ_DEAL_NTF",
	86018:  "ROOM_DDZ_GRAB_LORD_REQ",
	86019:  "ROOM_DDZ_GRAB_LORD_RSP",
	86020:  "ROOM_DDZ_GRAB_LORD_NTF",
	86021:  "ROOM_DDZ_LORD_NTF",
	86022:  "ROOM_DDZ_DOUBLE_REQ",
	86023:  "ROOM_DDZ_DOUBLE_RSP",
	86024:  "ROOM_DDZ_DOUBLE_NTF",
	86025:  "ROOM_DDZ_PLAY_CARD_REQ",
	86026:  "ROOM_DDZ_PLAY_CARD_RSP",
	86027:  "ROOM_DDZ_PLAY_CARD_NTF",
	86028:  "ROOM_DDZ_GAME_OVER_NTF",
	86032:  "ROOM_DDZ_RESUME_REQ",
	86033:  "ROOM_DDZ_RESUME_RSP",
	87552:  "MATCH_CONTINUE_REQ",
	87553:  "MATCH_CONTINUE_RSP",
	87554:  "MATCH_CONTINUE_DESK_DIMISS_NTF",
	131071: "ROOM_END",
}
var MsgID_value = map[string]int32{
	"LOGIN_AUTH_REQ":                  1,
	"LOGIN_AUTH_RSP":                  2,
	"GATE_AUTH_REQ":                   4097,
	"GATE_AUTH_RSP":                   4098,
	"GATE_HEART_BEAT_REQ":             4099,
	"GATE_HEART_BEAT_RSP":             4100,
	"GATE_ANOTHER_LOGIN_NTF":          4101,
	"GATE_TRANSMIT_HTTP_REQ":          4113,
	"GATE_TRANSMIT_HTTP_RSP":          4114,
	"MATCH_REQ":                       8193,
	"MATCH_RSP":                       8194,
	"MATCH_SUC_CREATE_DESK_NTF":       8198,
	"CANCEL_MATCH_REQ":                8199,
	"CANCEL_MATCH_RSP":                8200,
	"HALL_GET_PLAYER_INFO_REQ":        12289,
	"HALL_GET_PLAYER_INFO_RSP":        12290,
	"HALL_GET_PLAYER_STATE_REQ":       12291,
	"HALL_GET_PLAYER_STATE_RSP":       12292,
	"HALL_GET_GAME_LIST_INFO_REQ":     12293,
	"HALL_GET_GAME_LIST_INFO_RSP":     12294,
	"HALL_REAL_NAME_REQ":              12295,
	"HALL_REAL_NAME_RSP":              12296,
	"HALL_UPDATE_PLAYER_INFO_REQ":     12297,
	"HALL_UPDATE_PLAYER_INFO_RSP":     12298,
	"GET_CHARGE_INFO_REQ":             12299,
	"GET_CHARGE_INFO_RSP":             12300,
	"CHARGE_REQ":                      12301,
	"CHARGE_RSP":                      12302,
	"HALL_GET_PLAYER_GAME_INFO_REQ":   12303,
	"HALL_GET_PLAYER_GAME_INFO_RSP":   12304,
	"AUTH_CODE_REQ":                   12305,
	"AUTH_CODE_RSP":                   12306,
	"CHECK_AUTH_CODE_REQ":             12307,
	"CHECK_AUTH_CODE_RSP":             12308,
	"GET_BIND_PHONE_REWARD_REQ":       12309,
	"GET_BIND_PHONE_REWARD_RSP":       12310,
	"BIND_PHONE_REQ":                  12311,
	"BIND_PHONE_RSP":                  12312,
	"CHANGE_PHONE_REQ":                12313,
	"CHANGE_PHONE_RSP":                12320,
	"CHECK_PASSWORD_REQ":              12321,
	"CHECK_PASSWORD_RSP":              12322,
	"CHANGE_PASSWORD_REQ":             12323,
	"CHANGE_PASSWORD_RSP":             12324,
	"RESET_PASSWORD_REQ":              12325,
	"RESET_PASSWORD_RSP":              12326,
	"MONEY_CHANGE_NTF":                12327,
	"MSGSVR_GET_HORSE_RACE_REQ":       16385,
	"MSGSVR_GET_HORSE_RACE_RSP":       16386,
	"MSGSVR_HORSE_RACE_UPDATE_NTF":    16387,
	"MAILSVR_GET_UNREAD_SUM_REQ":      20481,
	"MAILSVR_GET_UNREAD_SUM_RSP":      20482,
	"MAILSVR_GET_MAIL_LIST_REQ":       20483,
	"MAILSVR_GET_MAIL_LIST_RSP":       20484,
	"MAILSVR_GET_MAIL_DETAIL_REQ":     20485,
	"MAILSVR_GET_MAIL_DETAIL_RSP":     20486,
	"MAILSVR_SET_READ_TAG_REQ":        20487,
	"MAILSVR_SET_READ_TAG_RSP":        20488,
	"MAILSVR_DEL_MAIL_REQ":            20489,
	"MAILSVR_DEL_MAIL_RSP":            20490,
	"MAILSVR_AWARD_ATTACH_REQ":        20491,
	"MAILSVR_AWARD_ATTACH_RSP":        20492,
	"MAILSVR_GET_AD_REQ":              20493,
	"MAILSVR_GET_AD_RSP":              20494,
	"ALMS_GET_GOLD_REQ":               24577,
	"ALMS_GET_GOLD_RSP":               24578,
	"ALMS_LOGIN_GOLD_CONFIG_NTF":      24579,
	"PACKSACK_INFO_REQ":               24580,
	"PACKSACK_INFO_RSP":               24581,
	"PACKSACK_GOLD_REQ":               24582,
	"PACKSACK_GOLD_RSP":               24583,
	"PACKSACK_GET_GOLD_REQ":           24584,
	"PACKSACK_GET_GOLD_RSP":           24585,
	"ROOM_BEGIN":                      65536,
	"ROOM_BASE_BEGIN":                 65537,
	"ROOM_LOGIN_REQ":                  65538,
	"ROOM_LOGIN_RSP":                  65539,
	"ROOM_JOIN_DESK_REQ":              65540,
	"ROOM_JOIN_DESK_RSP":              65541,
	"ROOM_DESK_CREATED_NTF":           65542,
	"ROOM_DESK_QUIT_REQ":              65543,
	"ROOM_DESK_DISMISS_NTF":           65544,
	"ROOM_DESK_CONTINUE_REQ":          65545,
	"ROOM_DESK_CONTINUE_RSP":          65546,
	"ROOM_VISITOR_LOGIN_REQ":          65547,
	"ROOM_VISITOR_LOGIN_RSP":          65548,
	"ROOM_DESK_NEED_RESUME_REQ":       65549,
	"ROOM_DESK_NEED_RESUME_RSP":       65550,
	"ROOM_DESK_QUIT_RSP":              65551,
	"ROOM_PLAYER_LOCATION_REQ":        65552,
	"ROOM_PLAYER_LOCATION_RSP":        65553,
	"ROOM_DESK_QUIT_ENTER_NTF":        65554,
	"ROOM_CHANGE_PLAYERS_REQ":         65556,
	"ROOM_CHANGE_PLAYERS_RSP":         65557,
	"ROOM_PLAYER_GIVEUP_REQ":          65792,
	"ROOM_PLAYER_GIVEUP_RSP":          65793,
	"ROOM_PLAYER_GIVEUP_NTF":          65794,
	"ROOM_BROKER_PLAYER_CONTINUE_REQ": 65795,
	"ROOM_BROKER_PLAYER_CONTINUE_RSP": 65796,
	"ROOM_USE_PROP_REQ":               65808,
	"ROOM_USE_PROP_RSP":               65809,
	"ROOM_USE_PROP_NTF":               65810,
	"ROOM_BASE_END":                   69631,
	"ROOM_GAME_BEGIN":                 69632,
	"ROOM_START_GAME_NTF":             69633,
	"ROOM_XIPAI_NTF":                  69634,
	"ROOM_FAPAI_NTF":                  69635,
	"ROOM_HUANSANZHANG_NTF":           69664,
	"ROOM_HUANSANZHANG_REQ":           69665,
	"ROOM_HUANSANZHANG_RSP":           69666,
	"ROOM_HUANSANZHANG_FINISH_NTF":    69667,
	"ROOM_DINGQUE_NTF":                69696,
	"ROOM_DINGQUE_REQ":                69697,
	"ROOM_DINGQUE_RSP":                69698,
	"ROOM_DINGQUE_FINISH_NTF":         69699,
	"ROOM_CHUPAIWENXUN_NTF":           69728,
	"ROOM_XINGPAI_ACTION_REQ":         69744,
	"ROOM_PENG_RSP":                   69761,
	"ROOM_PENG_NTF":                   69762,
	"ROOM_GANG_RSP":                   69793,
	"ROOM_GANG_NTF":                   69794,
	"ROOM_HU_RSP":                     69825,
	"ROOM_HU_NTF":                     69826,
	"ROOM_QI_RSP":                     69857,
	"ROOM_CHI_RSP":                    69872,
	"ROOM_CHI_NTF":                    69873,
	"ROOM_ZIXUN_NTF":                  69888,
	"ROOM_CHUPAI_REQ":                 69889,
	"ROOM_CHUPAI_NTF":                 69890,
	"ROOM_MOPAI_NTF":                  69891,
	"ROOM_BUHUA_NTF":                  69892,
	"ROOM_WAIT_QIANGGANGHU_NTF":       69920,
	"ROOM_TINGINFO_NTF":               69921,
	"ROOM_CARTOON_FINISH_REQ":         69925,
	"ROOM_GAMEOVER_NTF":               69952,
	"ROOM_RESUME_GAME_REQ":            69969,
	"ROOM_RESUME_GAME_RSP":            69970,
	"ROOM_GAME_END":                   73727,
	"ROOM_SETTLE_BEGIN":               73728,
	"ROOM_INSTANT_SETTLE":             73729,
	"ROOM_ROUND_SETTLE":               73730,
	"ROOM_SETTLE_END":                 77823,
	"ROOM_TUOGUAN_BEGIN":              77824,
	"ROOM_TUOGUAN_NTF":                77825,
	"ROOM_CANCEL_TUOGUAN_REQ":         77826,
	"ROOM_TUOGUAN_REQ":                77827,
	"ROOM_TUOGUAN_END":                78079,
	"ROOM_AUTOHU_REQ":                 78080,
	"ROOM_AUTOHU_RSP":                 78081,
	"ROOM_COUNT_DOWN_NTF":             78096,
	"ROOM_CHAT_REQ":                   81920,
	"ROOM_CHAT_NTF":                   81921,
	"ROOM_DDZ_START_GAME_NTF":         86016,
	"ROOM_DDZ_DEAL_NTF":               86017,
	"ROOM_DDZ_GRAB_LORD_REQ":          86018,
	"ROOM_DDZ_GRAB_LORD_RSP":          86019,
	"ROOM_DDZ_GRAB_LORD_NTF":          86020,
	"ROOM_DDZ_LORD_NTF":               86021,
	"ROOM_DDZ_DOUBLE_REQ":             86022,
	"ROOM_DDZ_DOUBLE_RSP":             86023,
	"ROOM_DDZ_DOUBLE_NTF":             86024,
	"ROOM_DDZ_PLAY_CARD_REQ":          86025,
	"ROOM_DDZ_PLAY_CARD_RSP":          86026,
	"ROOM_DDZ_PLAY_CARD_NTF":          86027,
	"ROOM_DDZ_GAME_OVER_NTF":          86028,
	"ROOM_DDZ_RESUME_REQ":             86032,
	"ROOM_DDZ_RESUME_RSP":             86033,
	"MATCH_CONTINUE_REQ":              87552,
	"MATCH_CONTINUE_RSP":              87553,
	"MATCH_CONTINUE_DESK_DIMISS_NTF":  87554,
	"ROOM_END":                        131071,
}

func (x MsgID) Enum() *MsgID {
	p := new(MsgID)
	*p = x
	return p
}
func (x MsgID) String() string {
	return proto.EnumName(MsgID_name, int32(x))
}
func (x *MsgID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MsgID_value, data, "MsgID")
	if err != nil {
		return err
	}
	*x = MsgID(value)
	return nil
}
func (MsgID) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterEnum("msgid.MsgID", MsgID_name, MsgID_value)
}

func init() { proto.RegisterFile("msgid.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1616 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x98, 0xcd, 0x72, 0xdc, 0xc6,
	0x11, 0x80, 0x6b, 0x53, 0xd8, 0x4a, 0x32, 0x8e, 0xe9, 0xf1, 0xc8, 0xfa, 0xb3, 0x2c, 0x99, 0x71,
	0x92, 0x4b, 0x0e, 0xf1, 0x2b, 0x68, 0x08, 0x0c, 0x01, 0x84, 0xbb, 0x83, 0x21, 0x66, 0x40, 0xca,
	0xbc, 0x60, 0x2b, 0x89, 0xca, 0xe5, 0xaa, 0x24, 0x76, 0x45, 0xaa, 0x9c, 0xb1, 0x4b, 0x91, 0x12,
	0x43, 0x2b, 0xa6, 0x1c, 0x3b, 0x3f, 0xa7, 0x80, 0xeb, 0xfc, 0x5c, 0xf4, 0x0e, 0x59, 0x3b, 0x4f,
	0x90, 0xbc, 0x40, 0x64, 0x3d, 0x81, 0x73, 0x4f, 0x95, 0x53, 0xdd, 0x03, 0xec, 0x60, 0x41, 0x60,
	0xed, 0x93, 0xb4, 0xfd, 0xf5, 0xf4, 0x74, 0xf7, 0x74, 0x4f, 0x0f, 0x48, 0x5e, 0xf8, 0xc5, 0xbd,
	0xb7, 0xdf, 0xf9, 0xd9, 0x8f, 0xde, 0xfb, 0xd5, 0xbb, 0xf7, 0xdf, 0x65, 0x43, 0xfc, 0xf1, 0xc3,
	0xff, 0x7d, 0x8f, 0x0c, 0xc7, 0xf7, 0xde, 0x8e, 0x03, 0xc6, 0xc8, 0xc6, 0x28, 0x09, 0x63, 0x99,
	0xf3, 0xcc, 0x44, 0x79, 0x2a, 0x76, 0xe9, 0xa0, 0x2d, 0xd3, 0x8a, 0x7e, 0x83, 0x31, 0xf2, 0x62,
	0xc8, 0x8d, 0x70, 0x6a, 0xd3, 0xcd, 0x96, 0x4c, 0x2b, 0x3a, 0xdb, 0x64, 0xd7, 0xc8, 0x25, 0x94,
	0x45, 0x82, 0xa7, 0x26, 0xdf, 0x12, 0xdc, 0xa0, 0xf6, 0x61, 0x37, 0xd1, 0x8a, 0x3e, 0xd8, 0x64,
	0x37, 0xc8, 0x15, 0x6b, 0x47, 0x26, 0x26, 0x12, 0x69, 0x6e, 0x37, 0x97, 0x66, 0x9b, 0x1e, 0x39,
	0x68, 0x52, 0x2e, 0xf5, 0x38, 0x36, 0x79, 0x64, 0x8c, 0x42, 0x9b, 0x4f, 0x7a, 0xa1, 0x56, 0xf4,
	0xc3, 0x4d, 0xb6, 0x41, 0xbe, 0x3d, 0xe6, 0xc6, 0xaf, 0xdc, 0xbd, 0xdd, 0xf8, 0x0d, 0xae, 0xde,
	0x66, 0xb7, 0xc8, 0x75, 0xfb, 0x5b, 0x67, 0x7e, 0xee, 0xa7, 0x02, 0x0c, 0x05, 0x42, 0xef, 0xe0,
	0xce, 0xc7, 0xb7, 0xd9, 0x65, 0x42, 0x7d, 0x2e, 0x7d, 0x31, 0xca, 0x9d, 0x99, 0x87, 0x1d, 0x62,
	0xad, 0xe8, 0xa3, 0xdb, 0xec, 0x26, 0xb9, 0x16, 0xf1, 0xd1, 0x28, 0x0f, 0x85, 0xc9, 0xd5, 0x88,
	0xbf, 0x25, 0xd2, 0x3c, 0x96, 0xdb, 0x89, 0xdd, 0x7c, 0xd2, 0x8f, 0xc1, 0x97, 0x09, 0xf8, 0xd2,
	0xc6, 0xda, 0x80, 0x43, 0x98, 0xbc, 0x75, 0x1c, 0x52, 0x38, 0x61, 0x9b, 0xe4, 0xc6, 0x92, 0x87,
	0x7c, 0x2c, 0xf2, 0x51, 0xac, 0x8d, 0x73, 0xe0, 0x68, 0xbd, 0x86, 0x56, 0xf4, 0x78, 0xc2, 0xae,
	0x12, 0x86, 0x1a, 0xa9, 0xe0, 0xa3, 0x5c, 0x82, 0x0a, 0x46, 0xdc, 0x09, 0x20, 0x66, 0x67, 0x33,
	0x53, 0x01, 0xf8, 0xd2, 0x0e, 0xfb, 0x64, 0xbd, 0x86, 0x56, 0xf4, 0x37, 0x13, 0x2c, 0x0b, 0x61,
	0x72, 0x3f, 0xe2, 0x69, 0x28, 0xdc, 0xda, 0xd3, 0x6e, 0xa2, 0x15, 0x7d, 0x7f, 0xc2, 0x5e, 0x22,
	0xa4, 0x92, 0x82, 0xea, 0xe3, 0x15, 0x81, 0x56, 0xf4, 0xb7, 0x13, 0xf6, 0x06, 0xb9, 0xd9, 0xce,
	0x17, 0x06, 0xbd, 0xb4, 0xff, 0xc1, 0x57, 0xe9, 0x68, 0x45, 0xcf, 0x26, 0x50, 0xe2, 0x58, 0xdd,
	0x7e, 0x12, 0xd8, 0xcd, 0x9e, 0xb4, 0x65, 0x50, 0x6b, 0xe8, 0xab, 0x1f, 0x09, 0x7f, 0x27, 0x5f,
	0xd5, 0xfe, 0x5d, 0x37, 0xd1, 0x8a, 0x7e, 0x84, 0x67, 0x0a, 0x5b, 0x6f, 0xc5, 0x32, 0xc8, 0x55,
	0x94, 0x48, 0x58, 0xb2, 0xcf, 0xd3, 0x00, 0x57, 0x7e, 0xbc, 0x8e, 0x6b, 0x45, 0x7f, 0x3f, 0x61,
	0x97, 0xc8, 0xc6, 0x0a, 0xdb, 0xa5, 0x7f, 0xb8, 0x20, 0xd4, 0x8a, 0xfe, 0x71, 0x82, 0x25, 0x1b,
	0x71, 0x19, 0x8a, 0x86, 0xee, 0x9f, 0x3a, 0xc4, 0x5a, 0xd1, 0x12, 0x8f, 0xdb, 0x7a, 0xac, 0xb8,
	0xd6, 0xfb, 0x49, 0xe5, 0xd0, 0x79, 0x27, 0xd0, 0x8a, 0xce, 0xab, 0x18, 0xad, 0xa1, 0xe6, 0x92,
	0x4f, 0xba, 0x89, 0x56, 0xf4, 0x2f, 0x68, 0x2c, 0x15, 0x1a, 0x52, 0xdf, 0x5c, 0xf2, 0xd7, 0x4e,
	0xa0, 0x15, 0xfd, 0x1b, 0xba, 0x3b, 0x4e, 0xa4, 0x78, 0x2b, 0xaf, 0x2c, 0x42, 0x9b, 0xfe, 0x7d,
	0xc2, 0x5e, 0x27, 0xd7, 0xc7, 0x3a, 0xd4, 0x7b, 0x29, 0x1e, 0x64, 0x94, 0xa4, 0x5a, 0xe4, 0x29,
	0xf7, 0x6d, 0x94, 0xd3, 0x62, 0xb0, 0x46, 0x01, 0x7a, 0xaf, 0x18, 0xb0, 0x37, 0xc8, 0x6b, 0x95,
	0x42, 0x03, 0x56, 0x15, 0x0b, 0x9b, 0x1c, 0x16, 0x03, 0xb6, 0x49, 0x5e, 0x1d, 0xf3, 0x78, 0x54,
	0x5b, 0xc9, 0x64, 0x2a, 0x78, 0x90, 0xeb, 0x6c, 0x6c, 0xb7, 0x29, 0xd7, 0x6a, 0xc0, 0x3e, 0xa5,
	0x75, 0xa4, 0xa1, 0x01, 0xff, 0xb7, 0x5d, 0x88, 0x4d, 0xbe, 0x56, 0x01, 0xba, 0xbc, 0x1c, 0xb0,
	0xef, 0x92, 0x1b, 0x17, 0x14, 0x02, 0x61, 0xe0, 0x1f, 0x6c, 0xf3, 0xaf, 0x50, 0x81, 0x3e, 0x2f,
	0x07, 0xec, 0x16, 0xb9, 0x56, 0xab, 0x40, 0x9e, 0xd1, 0x4f, 0xc3, 0x43, 0xdb, 0xee, 0xeb, 0x38,
	0x74, 0x7d, 0x39, 0x60, 0xaf, 0x92, 0x57, 0x6a, 0x1e, 0xe0, 0x35, 0x58, 0x6d, 0x7f, 0xd2, 0xc7,
	0xa0, 0xd3, 0x57, 0xed, 0x72, 0x2c, 0x64, 0x6e, 0x0c, 0xaf, 0x2e, 0xd6, 0xd3, 0x75, 0x1c, 0xba,
	0xbe, 0x1c, 0xb0, 0x6b, 0x84, 0x35, 0x43, 0xe3, 0xb6, 0x64, 0x1e, 0x77, 0x13, 0xb8, 0x07, 0xca,
	0x01, 0xbb, 0x4a, 0x5e, 0xe6, 0xa3, 0xb1, 0xb6, 0xd7, 0x5e, 0x32, 0xb2, 0x4b, 0xa6, 0x8b, 0x2e,
	0x00, 0xa7, 0xb4, 0xc0, 0x73, 0x44, 0x60, 0xa7, 0x10, 0x22, 0x3f, 0x91, 0xdb, 0x71, 0x68, 0x6b,
	0xc1, 0x2e, 0x55, 0xdc, 0xdf, 0xd1, 0xdc, 0xdf, 0x71, 0x17, 0xca, 0x83, 0x4e, 0xa0, 0x15, 0x3d,
	0x6a, 0x81, 0xa5, 0x17, 0xc7, 0x9d, 0x40, 0x2b, 0xfa, 0x70, 0x31, 0x60, 0x37, 0xc8, 0x65, 0x07,
	0x9a, 0xbe, 0x3f, 0xea, 0x85, 0x5a, 0xd1, 0x93, 0xc5, 0x80, 0x51, 0x42, 0xd2, 0x24, 0x19, 0xe7,
	0x5b, 0x22, 0x8c, 0x25, 0x2d, 0x0a, 0x8f, 0x5d, 0x26, 0x2f, 0x59, 0x09, 0xd7, 0xa2, 0x12, 0x4f,
	0x0b, 0x8f, 0xbd, 0x42, 0x36, 0x50, 0x6c, 0x03, 0x05, 0xdb, 0xb3, 0x8b, 0x52, 0xad, 0xe8, 0x61,
	0xe1, 0x41, 0x82, 0x51, 0xfa, 0xe3, 0x24, 0x96, 0x76, 0x48, 0x62, 0xcc, 0xdd, 0x04, 0x82, 0x2e,
	0x3c, 0xf0, 0x12, 0x09, 0x0a, 0xed, 0x7c, 0x0d, 0xec, 0x6c, 0x6d, 0x2c, 0x43, 0xb8, 0x9b, 0xc5,
	0xb6, 0x09, 0x1e, 0xb6, 0x97, 0x05, 0xb1, 0x1e, 0xc7, 0x5a, 0xe3, 0xb2, 0x47, 0x85, 0xc7, 0x5e,
	0x23, 0x57, 0x1a, 0x36, 0x13, 0x69, 0x62, 0x99, 0xd9, 0x4e, 0x3f, 0xe9, 0xa7, 0x50, 0x7e, 0x0d,
	0xba, 0x17, 0xeb, 0xd8, 0x24, 0x69, 0x23, 0xee, 0xd3, 0x7e, 0x0a, 0xa5, 0x57, 0x78, 0xd0, 0x99,
	0xce, 0xb2, 0x14, 0x02, 0xce, 0x42, 0x67, 0xd5, 0x88, 0x7c, 0xbc, 0x56, 0x01, 0x0a, 0xb1, 0x3b,
	0x60, 0xad, 0xe8, 0x07, 0x85, 0x07, 0x65, 0x8f, 0xa4, 0x9a, 0x41, 0xa3, 0xc4, 0xe7, 0x26, 0x4e,
	0xac, 0x67, 0x67, 0xeb, 0xb8, 0x56, 0xf4, 0x49, 0x83, 0x3b, 0xcb, 0x42, 0x1a, 0x91, 0x62, 0xce,
	0x3e, 0x2c, 0x3c, 0x76, 0x93, 0x5c, 0x45, 0x5e, 0x5f, 0xc4, 0x68, 0x46, 0xa3, 0xf9, 0x8f, 0xd6,
	0x60, 0xad, 0xe8, 0xc7, 0x8d, 0xbc, 0xd4, 0x13, 0x32, 0xde, 0x13, 0x99, 0x7d, 0x7f, 0x15, 0xb3,
	0x5e, 0xaa, 0x15, 0x9d, 0xf6, 0x52, 0xf0, 0x6b, 0x36, 0xf3, 0xd8, 0x0f, 0xc8, 0xeb, 0xb6, 0x2c,
	0xd3, 0x64, 0x47, 0xa4, 0xb5, 0xd2, 0xca, 0xa1, 0x1e, 0x7e, 0x0d, 0x35, 0xb8, 0x1a, 0x67, 0x1e,
	0x74, 0x12, 0xaa, 0x65, 0x5a, 0xe4, 0x2a, 0x4d, 0xac, 0x8b, 0x67, 0x9d, 0x00, 0xf2, 0xd6, 0x05,
	0x30, 0x61, 0x33, 0x8f, 0x5d, 0x22, 0x2f, 0xba, 0x7e, 0x11, 0x32, 0xa0, 0x5f, 0xfe, 0xd9, 0x35,
	0x11, 0x3e, 0x11, 0xaa, 0xde, 0x2a, 0x3d, 0x76, 0x9d, 0x5c, 0x42, 0xb1, 0x36, 0xf0, 0xaa, 0x45,
	0x08, 0x66, 0xa6, 0xa5, 0xeb, 0xa4, 0x3b, 0xb1, 0xe2, 0xb1, 0x8d, 0xba, 0x21, 0xdd, 0xe6, 0xb5,
	0xf4, 0xb0, 0x74, 0x45, 0x1f, 0x65, 0x5c, 0x6a, 0x2e, 0x0f, 0xe0, 0x2c, 0x10, 0x96, 0x7d, 0x10,
	0x67, 0x72, 0x2f, 0x84, 0xb9, 0x5c, 0x7a, 0x30, 0xd9, 0x2e, 0xc2, 0xed, 0x58, 0xc6, 0x3a, 0x42,
	0xeb, 0x9f, 0x94, 0x1e, 0xbb, 0x42, 0xa8, 0x2d, 0x9f, 0x58, 0x86, 0xbb, 0x99, 0x75, 0x7f, 0xd1,
	0x21, 0x87, 0x0d, 0x3f, 0xed, 0x92, 0x6b, 0x45, 0x3f, 0x2b, 0x5d, 0x1d, 0xd5, 0xf2, 0xc6, 0x36,
	0xff, 0x6c, 0xf8, 0xe9, 0x47, 0x99, 0xe2, 0xf1, 0xbe, 0x90, 0x77, 0x32, 0xfb, 0xc6, 0x7f, 0xd6,
	0x58, 0x7b, 0x27, 0x96, 0x21, 0xa4, 0x85, 0xfb, 0xcb, 0x0e, 0xf8, 0xa2, 0x74, 0x07, 0xa2, 0x44,
	0x15, 0xdb, 0xf4, 0xbc, 0x25, 0xc4, 0xec, 0x36, 0x84, 0x61, 0x9d, 0x85, 0xf3, 0xb6, 0x10, 0x34,
	0xe7, 0xe7, 0x1e, 0x7b, 0x99, 0xbc, 0x50, 0xa5, 0x06, 0xf5, 0x3e, 0x5d, 0x15, 0x81, 0xd6, 0x67,
	0x0d, 0xd1, 0x6e, 0x8c, 0x5a, 0x9f, 0x9f, 0x7b, 0x8c, 0x91, 0xef, 0x54, 0x81, 0x58, 0xd9, 0x17,
	0x2d, 0x19, 0x2c, 0xfd, 0xef, 0xb9, 0x3b, 0xe8, 0x83, 0xb8, 0x8e, 0xb4, 0x98, 0xbb, 0x32, 0xb2,
	0x69, 0xb0, 0xd3, 0xe8, 0xa2, 0x18, 0xc3, 0x99, 0x3b, 0x1b, 0xe3, 0x64, 0x59, 0x2c, 0x0d, 0xe9,
	0x56, 0x16, 0x65, 0x1c, 0xa5, 0x0f, 0xe6, 0xee, 0x06, 0xda, 0xe7, 0xb1, 0xc9, 0x77, 0x63, 0x2e,
	0x43, 0x88, 0xb7, 0x8a, 0xa5, 0x9c, 0xbb, 0x7a, 0x37, 0xb1, 0x0c, 0x71, 0x3a, 0x01, 0x38, 0x9f,
	0x37, 0x6e, 0x00, 0x9e, 0x9a, 0x24, 0x91, 0xf5, 0xc9, 0xe1, 0x7b, 0xac, 0xb1, 0x0e, 0x8a, 0x3b,
	0xd9, 0xab, 0x2e, 0x96, 0xc5, 0xdc, 0x83, 0x59, 0x8f, 0xa0, 0xba, 0xe9, 0xc2, 0xfa, 0x93, 0xe1,
	0x5f, 0x7d, 0x4c, 0x2b, 0xfa, 0xef, 0x79, 0xf3, 0x3c, 0xc6, 0x55, 0x7f, 0xfd, 0xc3, 0xed, 0xa2,
	0x85, 0x31, 0xa3, 0x65, 0x87, 0x2d, 0x5c, 0x87, 0xc5, 0x52, 0x1b, 0x2e, 0x4d, 0xa5, 0x40, 0xa7,
	0x0b, 0xb7, 0x26, 0x4d, 0x32, 0x19, 0xd4, 0x60, 0xb6, 0x70, 0xe9, 0xac, 0x8c, 0xe1, 0x1e, 0xff,
	0x71, 0x77, 0xb0, 0xc9, 0x92, 0x30, 0xe3, 0xb2, 0xde, 0xe4, 0x99, 0x2b, 0xea, 0x9a, 0x60, 0x0f,
	0x3f, 0x6b, 0xa6, 0x06, 0xbf, 0xf8, 0x6a, 0x8c, 0xc3, 0xb2, 0x63, 0x19, 0xde, 0x59, 0x1d, 0x72,
	0x74, 0xe0, 0x73, 0xe7, 0x17, 0xcf, 0x4c, 0x02, 0x85, 0x07, 0xb7, 0xe8, 0xf3, 0x8b, 0x62, 0xa8,
	0xf0, 0xe7, 0x2e, 0x72, 0x3f, 0xc9, 0xa4, 0xc9, 0x83, 0x64, 0xdf, 0xfa, 0x75, 0xf6, 0xdc, 0xa5,
	0xd0, 0x8f, 0xaa, 0x0f, 0xec, 0xa2, 0x18, 0xae, 0x0a, 0x31, 0x82, 0x62, 0xe8, 0xda, 0x32, 0x38,
	0x68, 0x5f, 0x52, 0x45, 0x39, 0x5c, 0xa6, 0x10, 0x70, 0x80, 0xdf, 0x77, 0x78, 0x7b, 0x0d, 0xdd,
	0x2c, 0x0d, 0x0e, 0xf2, 0x30, 0xe5, 0x5b, 0xf9, 0xa8, 0x7e, 0xa3, 0xcf, 0xfa, 0x29, 0xbc, 0x16,
	0x7a, 0x29, 0x16, 0x6a, 0x6b, 0xcb, 0x25, 0x38, 0x2a, 0x87, 0xcb, 0x78, 0xd1, 0x97, 0x24, 0xdb,
	0x1a, 0xd9, 0x72, 0x3a, 0xee, 0x41, 0xf0, 0x52, 0xea, 0x46, 0xf8, 0x5a, 0x68, 0xf9, 0x01, 0x73,
	0x03, 0x2a, 0x3c, 0xa8, 0x9e, 0xaa, 0xbd, 0x14, 0x1f, 0xab, 0x7d, 0x14, 0x2c, 0x9f, 0xb6, 0x23,
	0x84, 0x7c, 0x2e, 0x1b, 0xe3, 0xfd, 0x96, 0x4b, 0x8d, 0x77, 0xc2, 0x59, 0x0f, 0x82, 0x79, 0x54,
	0x0e, 0xed, 0x23, 0xd6, 0xf8, 0xd1, 0xea, 0x08, 0x2c, 0x9e, 0x76, 0x12, 0x28, 0x91, 0xa7, 0x43,
	0xf6, 0x7d, 0x72, 0xab, 0x45, 0xaa, 0x67, 0xd3, 0xf2, 0xd5, 0x34, 0x7b, 0x3a, 0x64, 0x1b, 0xe4,
	0x5b, 0xb8, 0x29, 0x96, 0xe1, 0x97, 0xdf, 0xdc, 0xba, 0x7a, 0x70, 0xf9, 0xde, 0xfd, 0xbb, 0xbf,
	0xbe, 0xfb, 0xe6, 0x4f, 0x7f, 0xfe, 0xce, 0xdd, 0x5f, 0xde, 0xcf, 0xdf, 0xfb, 0xc9, 0x9b, 0xf8,
	0x77, 0xa1, 0xff, 0x07, 0x00, 0x00, 0xff, 0xff, 0xc3, 0x04, 0x3a, 0xef, 0x2c, 0x12, 0x00, 0x00,
}
