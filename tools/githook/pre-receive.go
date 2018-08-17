package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

type CommitType string

const (
	BUG   CommitType = "BUG"
	DOC   CommitType = "DOC"
	TEST  CommitType = "TEST"
	FUNC  CommitType = "FUNC"
	MERGE CommitType = "MERGE"
)

const CommitMessagePattern = `^(\w*)\:.*`

// 是否开启严格模式，严格模式下将校验所有的提交信息格式(多 commit 下)
const strictMode = false

var commitMsgReg = regexp.MustCompile(CommitMessagePattern)

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)
	param := strings.Fields(string(input))

	// allow branch/tag delete
	if param[1] == "0000000000000000000000000000000000000000" {
		os.Exit(0)
	}

	odlCommitID, commitID := param[0], param[1]

	commitInfos := getCommitMsg(odlCommitID, commitID)

	for _, commitInfo := range commitInfos {
		commitTypes := commitMsgReg.FindAllStringSubmatch(commitInfo.msg, -1)

		if len(commitTypes) != 1 {
			checkFailed(commitInfo, odlCommitID, commitID)
		} else {
			switch commitTypes[0][1] {
			case string(BUG):
			case string(TEST):
			case string(DOC):
			case string(FUNC):
			case string(MERGE):
			default:
				checkFailed(commitInfo, odlCommitID, commitID)
			}
		}
		if !strictMode {
			os.Exit(0)
		}
	}

}

type commitInfo struct {
	hash string
	msg  string
}

func getCommitMsg(odlCommitID, commitID string) []commitInfo {
	getCommitMsgCmd := exec.Command("git", "log", odlCommitID+".."+commitID, "--pretty=format:%H %s")
	getCommitMsgCmd.Stdin = os.Stdin
	getCommitMsgCmd.Stderr = os.Stderr
	b, err := getCommitMsgCmd.Output()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	commitInfos := []commitInfo{}
	commits := strings.Split(string(b), "\n")
	for _, commit := range commits {
		infos := strings.Split(string(commit), " ")
		commitInfos = append(commitInfos, commitInfo{
			hash: infos[0],
			msg:  infos[1],
		})
	}

	return commitInfos
}

func checkFailed(commit commitInfo, oldCommitID, commitID string) {
	fmt.Fprintf(os.Stderr, "提交信息检测不通过，提交ID: %s， 提交信息：%s\n", commit.hash, commit.msg)
	fmt.Fprintf(os.Stderr, `正确提交示例：
	BUG: BUG:190245432-修复设置不了玩家服务端地址bug
	MERGE: 发布 develop 中的登录，充值等功能到 release
	FUNC: 登录相关流程实现
	DOC: 修改部署文档
	TEST: 增加登录功能测试用例
	
	请使用 'git log --oneline %s..%s' 查看对应的提交信息
	然后使用 'git rebase -i' 修改对应的提交信息
`, oldCommitID, commitID)
	os.Exit(1)
}
