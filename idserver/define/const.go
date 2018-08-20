package define
/*
 功能： 基础结构和常量定义
 作者： SkyWang
 日期： 2018-7-24

 */

import "fmt"

// 错误定义
var (
	ErrPropId = fmt.Errorf("prop id error")
	ErrNoUser = fmt.Errorf("no this user")
	ErrNoProp = fmt.Errorf("prop num < changenum")
	ErrLoadDB = fmt.Errorf("load from db failed")
	ErrSeqNo = fmt.Errorf("seq is same")
)
