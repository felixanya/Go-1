package logic

import (
	"github.com/Sirupsen/logrus"
	"math/rand"
	"steve/idserver/data"
	"sync"
	"fmt"
	"time"
	"github.com/spf13/viper"
)

/*
  功能： 账号ID生成器： 生成playerId, showId,并且保证id不会出现重复的现象。showId支持随机生成和过滤靓号等.
  作者： SkyWang
  日期： 2018-8-15

*/

var muLock *sync.Mutex // 锁

var keepSum = uint64(100000)

func Init() error {

	hw := viper.GetInt64("hw")
	if hw >= 10000 && hw < 100000000{
		keepSum = uint64(hw)
	}

	muLock = new(sync.Mutex)
	go runLogicTask()
	return nil
}

func runMakeId() {

	// 目前号码总量达到一半，就不生成号码
	can, _ := data.GetCanUseSumFromDB()
	if can > keepSum/2 {
		return
	}

	for {
		for t := 0; t < 100; t++ {
			makeNewShowId()
			time.Sleep(time.Millisecond * 20)
		}
		time.Sleep(time.Minute)

		can, _ := data.GetCanUseSumFromDB()
		if can >= keepSum {
			// 可用号码得到设置水位，就不再生成新的号码
			break
		}
	}
}

// 运行逻辑任务
func runLogicTask() {
	rand.Seed(int64(time.Now().UnixNano()))
	for {
		runMakeId()
		time.Sleep(time.Minute*5)
	}
}

// 自动生成新的ShowID
func makeNewShowId() {
	d := 0
	sum := 100
	uids := make([]string, 0, sum)
	//num := ""
	num := make([]byte, 10, 10)
	for k := 0; k < sum; k++ {
		strNum := ""
		for i := 0; i < 10; i++ {
			if i == 0 {
				d = rand.Intn(9) + 1
			} else {
				d = rand.Intn(10)
			}
			num[i] = byte(d)
			strNum += fmt.Sprintf("%d", d)
			//num += fmt.Sprintf("%d", d)
		}
		// 过滤靓号
		if isGoodId(num) {
			continue
		}
		logrus.Debugln(strNum)
		uids = append(uids, strNum)

	}
	data.InsertShowId(uids)

}

// 过滤靓号
func isGoodId(num []byte) bool {
	/*
	   符合以下规则的号码为靓号
	   1. 全部数字都相同的号码 - 如：1111111111
	   2. 全部号码由两个数字组成⽽而且分开在两边 - 如：1111133333或1999999999
	   3. 前⾯面数字随机，尾数相同的号码，且尾数相同>=3位 - 如：737485999，6272222222
	   4. 前⾯面数字随机，尾数为连续数字的号码，且尾部连续>=6位 - 如：8274123456，1983456789
	   5. 前⾯面数字随机，尾数为ABABAB的形式，且>=6位
	   / 7 8
	   - 如：283412121212，129090909090
	*/

	if isGoodId1(num) {
		return true
	}
	if isGoodId2(num) {
		return true
	}
	if isGoodId3(num) {
		return true
	}
	if isGoodId41(num) {
		return true
	}
	if isGoodId42(num) {
		return true
	}
	if isGoodId5(num) {
		return true
	}

	return false
}

func isGoodId5(num []byte) bool {
	iSame := 1
	for i := len(num)-1; i > 1; i-- {
		if num[i] != num[i-2]{
			break
		}
		iSame++
	}
	if iSame >= 6 {
		return true
	}
	return false
}

func isGoodId41(num []byte) bool {
	iSame := 1
	for i := len(num)-1; i > 0; i-- {
		if num[i] != num[i-1] + 1 {
			break
		}
		iSame++
	}
	if iSame >= 6 {
		return true
	}
	return false
}

func isGoodId42(num []byte) bool {
	iSame := 1
	for i := len(num)-1; i > 0; i-- {
		if num[i] != num[i-1] - 1 {
			break
		}
		iSame++
	}
	if iSame >= 6 {
		return true
	}
	return false
}


func isGoodId3(num []byte) bool {
	iSame := 1
	for i := len(num)-1; i > 0; i-- {
		if num[i] != num[i-1] {
			break
		}
		iSame++
	}
	if iSame >= 3 {
		return true
	}
	return false
}


func isGoodId1(num []byte) bool {
	bGood := true
	for i := 0; i < len(num)-1; i++ {
		if num[i] != num[i+1] {
			bGood = false
			break
		}
	}
	if bGood {
		return true
	}
	return false
}

func isGoodId2(num []byte) bool {
	bGood := true
	first := byte(99)

	for i := 0; i < len(num)-1; i++ {
		if first == 99 {
			if num[i] != num[i+1] {
				first = num[i]
			}
		} else {
			if num[i] != num[i+1] {
				bGood = false
				break
			}
		}
	}
	if bGood {
		return true
	}
	return false
}

// 生成一个新的playerid 和 showid
func NewPlayerShowId() (uint64, uint64, error) {
	muLock.Lock()
	defer muLock.Unlock()

	playerId, err := data.NewPlayerIdFromDB()
	if err != nil {
		return 0, 0, err
	}
	showId, err := data.NewShowIdFromDB()
	if err != nil {
		return 0, 0, err
	}

	return playerId, showId, nil
}

// 生成一个新的playerid
func NewPlayerId() (uint64, error) {
	muLock.Lock()
	defer muLock.Unlock()
	return data.NewPlayerIdFromDB()
}

// 生成一个新的showid
func NewShowId() (uint64, error) {
	muLock.Lock()
	defer muLock.Unlock()
	return data.NewShowIdFromDB()
}
