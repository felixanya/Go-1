package logic

import (
	"errors"
	"fmt"
	"github.com/Sirupsen/logrus"
	"steve/client_pb/mailserver"
	"steve/entity/goods"
	"steve/external/hallclient"
	"steve/mailserver/data"
	"steve/mailserver/define"
	"steve/structs"
	"time"
	"sort"
)

/*
  功能： 邮件管理:
		1.获取未读邮件总数
		2.获取邮件列表
		3.获取指定邮件详情.
		4.删除邮件
		5.领取附件.
  作者： SkyWang
  日期： 2018-8-7
*/

// 是否是主节点
var isMasterNode = false

// 清理过期邮件开始点数
var clearBeginHour = 4

// 清理过期邮件结束点数
var clearEndHour = 6

// 邮件列表
var mailList map[uint64]*define.MailInfo

// 省包节点列表
var provSendList map[int64][]*define.MailInfo

var cityAdList map[int64]*define.ADJson    // 城市级别的AD列表
var provAdList map[int64]*define.ADJson    // 省份级别的AD列表
var channelAdList map[int64]*define.ADJson // 渠道级别的AD列表

//定义interface{},并实现sort.Interface接口的三个方法
type mailSlice  []*define.MailInfo

func (c mailSlice) Len() int {
	return len(c)
}
func (c mailSlice) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c mailSlice) Less(i, j int) bool {
	return c[i].StartTime > c[j].StartTime

}


func Init() error {
	testADJson()
	err := getDataFromDB()
	if err != nil {
		logrus.Errorf("getDataFromDB first err:%v", err)
		//return err
	} else {
		logrus.Debugf("getDataFromDB first win...")
	}

	//testJsonObject()
	// 启动跑马灯是否开始检测协程
	go runCheckMailChange()
	return nil
}

func testJsonObject() {
	dst := &define.SendDest{}
	dst.SendType = 0
	dst.Prov = 1
	dst.Channel = 0
	jsonDest, _ := data.MarshalSendDest(dst)
	logrus.Debugln(jsonDest)

	gs := &goods.Goods{}
	gs.GoodsType = 0
	gs.GoodsId = 1
	gs.GoodsNum = 100

	jsonGs, _ := data.MarshalAttachGoods([]*goods.Goods{gs, gs})
	logrus.Debugln(jsonGs)
}

// 获取广告位列表
func GetAD(uid uint64) ([]*mailserver.ADInfo, error) {

	// 获取玩家渠道ID
	channel, prov, city, ok := getUserInfo(uid)
	logrus.Debugf("getUserInfo: uid=%d,channel=%d,prov=%d,city=%d", uid, channel, prov, city)
	if !ok {
		return nil, errors.New("获取玩家失败")
	}

	if channel != 0 {
		ad, ok := channelAdList[channel]
		if ok && ad != nil {
			return ad.AdList, nil
		}
	}

	if city != 0 {
		// 先读取城市级别的AD
		ad, ok := cityAdList[city]
		if ok && ad != nil {
			return ad.AdList, nil
		}
	}

	if prov >= 0 {
		// 再读取省份级别的AD
		ad, ok := provAdList[prov]
		if ok && ad != nil {
			return ad.AdList, nil
		}
	}

	// 再读取省份通用的AD
	ad, ok := provAdList[0]
	if ok && ad != nil {
		return ad.AdList, nil
	}
	return nil, nil
}

// 获取未读消息总数
func GetGetUnReadSum(uid uint64) (int32, error) {
	// 获取玩家渠道ID
	channel, prov, _, ok := getUserInfo(uid)
	if !ok {
		return 0, errors.New("获取玩家失败")
	}

	// 从DB获取玩家的已读邮件列表
	readList, err := data.GetUserMailFromDB(uid)
	if err != nil {
		return 0, errors.New("从DB获取玩家已读邮件列表失败")
	}
	sum := int32(0)

	// 获取玩家所属省包的邮件
	if prov >= 0 {
		list := provSendList[prov]
		for _, mail := range list {

			// 检测是否符合省包和渠道ID
			isOk := checkMailProvChannel(uid, mail, channel, prov)
			if !isOk {
				continue
			}

			if one, ok := readList[mail.Id]; !ok || !one.IsRead {
				sum++
			}
		}
	}
	return sum, nil
}



	// 获取邮件消息列表
func GetMailList(uid uint64) ([]*mailserver.MailTitle, error) {
	// 获取玩家渠道ID
	channel, prov, _, ok := getUserInfo(uid)
	if !ok {
		return nil, errors.New("获取玩家失败")
	}
	if prov < 0 {
		return nil, errors.New("获取玩家省包ID < 0")
	}

	// 从DB获取玩家的已读邮件列表
	readList, err := data.GetUserMailFromDB(uid)
	if err != nil {
		return nil, errors.New("从DB获取玩家已读邮件列表失败")
	}

	titleList := make([]*mailserver.MailTitle, 0, 5)
	// 获取玩家所属省包的邮件

	format := "2006-01-02 15:04:05"
	list := provSendList[prov]
	for _, mail := range list {

		// 检测是否符合省包和渠道ID
		isOk := checkMailProvChannel(uid, mail, channel, prov)
		if !isOk {
			continue
		}
		title := new(mailserver.MailTitle)
		title.MailId = &mail.Id
		title.MailTitle = &mail.Title
		title.CreateTime = &mail.StartTime
		t, err := time.Parse(format, mail.StartTime)
		if err == nil {
			strTime := fmt.Sprintf("%02d-%02d", t.Month(), t.Day())
			title.CreateTime = &strTime
		}

		isRead := int32(0)
		one, ok := readList[mail.Id]
		if !ok || !one.IsRead {

		} else {
			isRead = 1
		}
		// 已经删除的，不返回
		if one != nil && one.IsDel {
			continue
		}
		title.IsRead = &isRead

		isHaveAttach := int32(0)
		if len(mail.Attach) > 0 {
			isHaveAttach = 1
		}
		if one != nil && one.IsGetAttach {
			isHaveAttach = 2
		}
		title.IsHaveAttach = &isHaveAttach

		titleList = append(titleList, title)
	}


	return titleList, nil
}

// 获取指定邮件详情
func GetMailDetail(uid uint64, mailId uint64) (*mailserver.MailDetail, error) {

	mail, ok := mailList[mailId]
	if !ok {
		return nil, errors.New("指定邮件不存在")
	}
	if mail.State != define.StateSended && mail.State != define.StateSending {
		return nil, errors.New("指定邮件状态错误")
	}
	isHaveAttach := int32(0)
	if len(mail.AttachGoods) > 0 {
		isHaveAttach = 1
	}

	// 从DB获取玩家的已读邮件列表
	one, _ := data.GetTheMailFromDB(uid, mailId)

	if one != nil && one.IsDel {
		return nil, errors.New("邮件已被用户删除")
	}

	isRead := int32(1)

	if one == nil {
		// 设置邮件=已读
		if isHaveAttach == 0 {
			isRead = 1
			data.SetEmailReadTagFromDB(uid, mailId, true, mail.DelTime, 1)

		} else {
			isRead = 0
			data.SetEmailReadTagFromDB(uid, mailId, true, mail.DelTime, 0)
		}
	} else {
		if one.IsGetAttach {
			isRead = 1
			if !one.IsRead {
				// 设置邮件=已读
				data.SetEmailReadTagFromDB(uid, mailId, false, mail.DelTime, 1)

			}
		} else {
			isRead = 0
			if one.IsRead {
				// 设置邮件=已读
				data.SetEmailReadTagFromDB(uid, mailId, false, mail.DelTime, 0)
			}
		}
	}

	detail := new(mailserver.MailDetail)
	detail.MailId = &mail.Id
	detail.MailTitle = &mail.Title
	detail.Content = &mail.Detail
	//detail.DelTime = &mail.DelTime

	strDel := ""
	format := "2006-01-02 15:04:05"
	tm, err := time.Parse(format, mail.DelTime)

	OneDay := int64(3600 * 24)
	if err == nil {
		sb := tm.Unix() - time.Now().Unix()
		if sb > 0 {
			d := sb / OneDay
			sb = sb % OneDay
			h := sb / 3600
			sb = sb % 3600
			m := sb / 60

			strDel = fmt.Sprintf("%02d天%02d时%02d分", d, h, m)
		}
	}
	detail.DelTime = &strDel

	t := mailserver.GoodsType_GOODSTYPE_PROPS

	for _, ach := range mail.AttachGoods {
		newGoods := new(mailserver.Goods)
		newGoods.GoodsType = &t
		newGoods.GoodsId = &ach.GoodsId
		newGoods.GoodsNum = &ach.GoodsNum
		detail.AttachGoods = newGoods
		break
	}


	detail.IsRead = &isRead


	if one != nil && one.IsGetAttach {
		isHaveAttach = 2
	}
	detail.IsHaveAttach = &isHaveAttach

	return detail, nil
}

// 标记邮件为已读请求
func SetReadTag(uid uint64, mailId uint64) error {
	mail, ok := mailList[mailId]
	if !ok {
		return errors.New("指定邮件不存在")
	}

	if mail.State != define.StateSended && mail.State != define.StateSending {
		return errors.New("指定邮件状态错误")
	}

	// 从DB获取玩家的已读邮件列表
	one, _ := data.GetTheMailFromDB(uid, mailId)

	if one != nil && one.IsDel {
		return errors.New("邮件已被用户删除")
	}

	if one == nil {
		// 设置邮件=已读
		data.SetEmailReadTagFromDB(uid, mailId, true, mail.DelTime, 1)
	} else {
		if !one.IsRead {
			// 设置邮件=已读
			data.SetEmailReadTagFromDB(uid, mailId, false, mail.DelTime, 1)
		}
	}

	return nil
}

// 删除邮件
func DelMail(uid uint64, mailId uint64) error {

	mail, ok := mailList[mailId]
	if !ok {
		return errors.New("指定邮件不存在")
	}
	// 从DB获取玩家的已读邮件列表
	one, _ := data.GetTheMailFromDB(uid, mailId)
	if one != nil && one.IsDel {
		return errors.New("邮件已被用户删除")
	}

	if one != nil {
		if one.IsDel {
			return errors.New("邮件已被用户删除")
		}

		if len(mail.AttachGoods) > 0 {
			if !one.IsGetAttach {
				return errors.New("附件未领取,无法删除")
			}
		}
	} else {
		// 设置邮件为删除状态
		return data.DelEmailFromDB(uid, mailId, true, mail.DelTime)
	}

	return data.DelEmailFromDB(uid, mailId, false, mail.DelTime)

}

// 领取附件奖励请求
func AwardAttach(uid uint64, mailId uint64) (*mailserver.Goods, error) {

	mail, ok := mailList[mailId]
	if !ok {
		return nil, errors.New("指定邮件不存在")
	}
	if len(mail.AttachGoods) == 0 {
		return nil, errors.New("此邮件无附件")
	}
	// 从DB获取玩家的已读邮件列表
	one, _ := data.GetTheMailFromDB(uid, mailId)

	if one == nil {
		return nil, errors.New("玩家邮件不存在")
	}
	if one.IsDel {
		return nil, errors.New("邮件已被玩家删除")
	}
	// 如果已领取，直接返回
	if one.IsGetAttach {
		return nil, errors.New("邮件已领取")
	}

	// 发放附件奖励

	// 标记为已领取
	data.SetAttachGettedDB(uid, mailId)

	attach := &mailserver.Goods{}
	attach.GoodsId = &mail.AttachGoods[0].GoodsId
	attach.GoodsNum = &mail.AttachGoods[0].GoodsNum

	gType := mailserver.GoodsType_GOODSTYPE_PROPS
	attach.GoodsType = &gType

	return attach, nil
}

// 启动邮件列表变化检测协程
func runCheckMailChange() error {

	// 1分钟更新一次邮件列表
	for {
		time.Sleep(time.Minute)

		// 判断当时是否是主节点
		isMasterNode = structs.GetGlobalExposer().ConsulReq.IsMasterNode()
		getDataFromDB()
	}
	return nil
}

func testADJson() {
	adList := make([]*define.ADJson, 0, 2)

	ad := new(define.ADJson)
	ad.Id = 1
	ad.IsUse = 1
	ad.Prov = 0
	ad.Channel = 0
	ad.Prov = 0

	for i := uint64(1); i <= 4; i++ {
		one := new(mailserver.ADInfo)
		one.AdId = &i
		param := ""
		one.AdParam = &param

		tick := int32(5)
		one.AdTick = &tick
		pic := "http:/www.qq.com/pic123.jpg"
		one.PicUrl = &pic
		gourl := "http:/www.qq.com"
		one.GoUrl = &gourl

		ad.AdList = append(ad.AdList, one)
	}

	adList = append(adList, ad)

	ad2 := new(define.ADJson)

	*ad2 = *ad
	ad2.Id = 2
	ad2.Channel = 1
	adList = append(adList, ad2)

	strJson, err := data.MarshalADs(adList)
	logrus.Debugf("MarshalADs: json=%v,err=%v", strJson, err)

}

// 从DB获取AD
func getADFromDB() error {
	list, err := data.GetADFromDB()
	if err != nil {
		logrus.Errorf("getADFromDB err:%v", err)
		return err
	}

	channelList := make(map[int64]*define.ADJson)
	cityList := make(map[int64]*define.ADJson)
	provList := make(map[int64]*define.ADJson)

	for _, ad := range list {
		if ad.Channel == 0 && ad.City == 0 {
			provList[ad.Prov] = ad
		} else if ad.City > 0 && ad.Channel == 0 {
			cityList[ad.City] = ad
		} else if ad.Channel > 0 {
			channelList[ad.Channel] = ad
		}
	}

	channelAdList = channelList
	cityAdList = cityList
	provAdList = provList

	return nil
}

// 从DB获取邮件列表
func getDataFromDB() error {

	getADFromDB()

	list, err := data.LoadMailListFromDB()
	if err != nil {
		logrus.Errorln("load email list from db err:", list)
		return err
	}
	logrus.Debugln("email list:", list)
	mailList = list
	// 检测邮件状态
	checkMailStatus(mailList)

	// 主节点每日半夜4-6点，清理过期邮件
	clearExpiredEmail()
	return err
}

// 主节点每日半夜4-6点，清理过期邮件
var thisDay = 0

func clearExpiredEmail() {
	if !isMasterNode {
		// 非主节点，直接返回
		return
	}

	now := time.Now()
	// 每日只执行1次
	if thisDay == now.YearDay() {
		return
	}
	if now.Hour() >= clearBeginHour && now.Hour() < clearEndHour {
		data.ClearExpiredEmailFromDB()
		data.ClearExpiredUserEmailFromDB()
		thisDay = now.YearDay()
	}
}

// 检测邮件状态是否变化
func checkMailStatus(mailList map[uint64]*define.MailInfo) error {

	curDate := time.Now().Format("2006-01-02 15:04:05")

	bUpdate := false

	for _, mail := range mailList {
		if mail.State == define.StateChecked {
			// 检测是否开始
			if curDate >= mail.StartTime {
				mail.State = define.StateSending
				bUpdate = true

				// 主节点负责保存邮件状态到DB
				if isMasterNode {
					// 保存邮件状态变化到DB
					data.SetEmailStateToDB(mail.Id, mail.State)
				}

			}
		} else if mail.State == define.StateSending {
			// 检测是否结束
			if mail.IsUseEndTime && curDate >= mail.EndTime {
				mail.State = define.StateSended
				bUpdate = true
				// 主节点负责保存邮件状态到DB
				if isMasterNode {
					// 保存邮件状态变化到DB
					data.SetEmailStateToDB(mail.Id, mail.State)
				}
			}
		} else if mail.State == define.StateSended {
			// 检测是否达到删除时间
			if mail.IsUseDelTime && curDate >= mail.DelTime {
				mail.State = define.StateDelete
				bUpdate = true
				// 主节点负责保存邮件状态到DB
				if isMasterNode {
					// 保存邮件状态变化到DB
					data.SetEmailStateToDB(mail.Id, mail.State)
				}
			}
		}
	}

	bUpdate = true
	// 更新发送列表provSendList
	if bUpdate {
		// 将发送中和发送截至的加入到指定列表中
		myList := make(map[int64]mailSlice)
		for _, mail := range mailList {
			if mail.State == define.StateSending || mail.State == define.StateSended {

				for _, dest := range mail.DestList {

					myList[dest.Prov] = append(myList[dest.Prov], mail)
				}
			}
		}

		// 排序邮件列表
		myListSort := make(map[int64][]*define.MailInfo)

		for k, slic := range myList {
			sort.Sort(slic)
			myListSort[k] = slic
		}

		provSendList = myListSort
	}

	return nil
}

// 调用hall接口获取用户信息
// 返回:渠道ID，省ID，城市ID
func getUserInfo(uid uint64) (int64, int64, int64, bool) {
	return 0, 1, 0, true
	info, err := hallclient.GetPlayerInfo(uid)
	if err != nil {
		return 0, 0, 0, false
	}
	if info == nil {
		return 0, 0, 0, false
	}

	return int64(info.ChannelId), int64(info.ProvinceId), int64(info.CityId), true
}

// 检测是否符合省包和渠道ID
func checkMailProvChannel(uid uint64, mail *define.MailInfo, channel int64, prov int64) bool {
	isOk := false
	for _, dest := range mail.DestList {
		if dest.Prov != 0 && prov != dest.Prov {
			continue
		}

		if dest.Channel != 0 && channel != dest.Channel {
			continue
		}
		if dest.SendType == define.SendAll {
			// 发送给所有人
			isOk = true
		} else {
			// 发送给指定玩家列表
			for _, id := range dest.PlayerList {
				if id == uid {
					isOk = true
					break
				}
			}
		}
		if isOk {
			break
		}

	}
	return isOk
}
