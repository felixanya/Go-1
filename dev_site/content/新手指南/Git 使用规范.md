---
title: "Git 使用规范"
date: 2018-04-03T17:30:47+08:00
draft: false
---


## 概述

为了确保大家的之间的相互协作的顺利进行，在 git 代码库的使用上，大家务必遵循下列规范。

## 代码提交检查清单

* 每次提交代码以前先 git pull 最新代码
* 提交代码前确保自己的提交的代码是可以运行的
* 提交 .proto 文件前，确保提交的 proto 文件可以被编译成 .go 文件
* 使用 git add . 命令后，请用 git status 确认只有想要提交的文件在提交文件列表
* 提交代码前，请检查运行代码需要的依赖库，确认它也已经提交


## 规范列表

### 提交说明规范

* 提交说明不可为空，不可为 ”第一次提交“ 这种无意义的说明
* 提交类型分为如下几种：
    * 新功能，提交说明必须为："FUNC:新功能描述"， 如 ”FUNC:登录相关流程实现“
    * 修复BUG时：
        * 测试提交的BUG，提交说明必须为： "BUG:BUG ID-BUG 描述", 如 ”BUG:190245432-修复设置不了玩家服务端地址bug“
        * 非测试提交的BUG，提交说明必须为："BUG:描述"，比如：“BUG:脚本错误”
    * 合并代码， 提交说明必须为： "MERGE:合并描述"，如 "MERGE:发布 develop 中的登录，充值等功能到 release"， 或者 "MERGE: 机器人功能联调通过，合并至 develop"
    * 文档类，提交说明必须为："DOC:文档描述"， 如 "DOC:修改 Git 使用规范文档"
 
 * PS:区分大小写。

### 不应该提交的文件

* 不要提交编译后的二进制程序文件
* 不要提交编辑器相关的配置文件
* 不要提交只是自己个人使用的文件
* 使用 .gitignore 忽略掉那些不想要提交的文件

### 使用 git rbase 的时机

git rbase 是一个强大的工具，如果使用不当可能会带来灾难，请务必仔细阅读下面的文档。

https://git-scm.com/book/en/v2/Git-Branching-Rebasing

* 清除不需要的 merge commit

当 git pull 代码后，如果本地有未 push 到远程的 local commit， 这时 git pull 会产生一个 merge commit，对于这种 merge commit 建议使用 git rebase 来消除。

In general the way to get the best of both worlds is to rebase local changes you’ve made but haven’t shared yet before you push them in order to clean up your story, but never rebase anything you’ve pushed somewhere.


* push commit 到 remote 库以前，合并一下 commit

对于一个功能的多次 commit，在 push 到服务器以前，尽可能用 git rebase 合并这些 commit 成一个 commit，这能很大提高代码历史的可读性。



## Tip

* 当在 Linux 系统工作时，git status 命令显示的中文文件名会是编码形式，例如：

```
new file:   "content/\346\226\260\346\211\213\346\214\207\345\215\227/Git \344\275\277\347\224\250\350\247\204\350\214\203.md"
modified:   "content/\346\226\260\346\211\213\346\214\207\345\215\227/_index.md"
```
但我们希望的显示是这样：

```
new file:   content/新手指南/Git 使用规范.md
modified:   content/新手指南/_index.md
```

这时可以用下面的命令让 git 显示正常字符。

```shell

git config --global core.quotePath false

```


