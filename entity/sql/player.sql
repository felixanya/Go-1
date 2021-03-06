/*
 Navicat Premium Data Transfer

 Source Server         : 192.168.7.108
 Source Server Type    : MySQL
 Source Server Version : 50722
 Source Host           : 192.168.7.108:3306
 Source Schema         : player

 Target Server Type    : MySQL
 Target Server Version : 50722
 File Encoding         : 65001

 Date: 09/08/2018 11:36:41
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for t_hall_info
-- ----------------------------
DROP TABLE IF EXISTS `t_hall_info`;
CREATE TABLE `t_hall_info` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `playerID` bigint(20) NOT NULL COMMENT '玩家ID',
  `recharge` int(11) DEFAULT NULL COMMENT '总充值金额',
  `bust` int(11) DEFAULT NULL COMMENT '累计破产次数',
  `lastGame` int(11) DEFAULT NULL COMMENT '上次金币场玩法',
  `lastLevel` int(11) DEFAULT NULL COMMENT '上次金币场场次',
  `lastFriendsBureauNum` int(11) DEFAULT NULL COMMENT '上次朋友局房号',
  `lastFriendsBureauGame` int(11) DEFAULT NULL COMMENT '上次朋友局玩法',
  `lastGameStartTime` datetime DEFAULT NULL COMMENT '最后游戏开始时间',
  `winningRate` int(11) DEFAULT NULL COMMENT '胜率',
  `backpackID` bigint(20) DEFAULT NULL COMMENT '背包ID',
  `almsGotTimes` int(11) DEFAULT NULL COMMENT '救济已领取次数',
  `remark` varchar(256) DEFAULT NULL COMMENT '备注',
  `createTime` datetime DEFAULT NULL COMMENT '创建时间',
  `createBy` varchar(64) DEFAULT NULL COMMENT '创建人',
  `updateTime` datetime DEFAULT NULL COMMENT '更新时间',
  `updateBy` varchar(64) DEFAULT NULL COMMENT '更新人',
  PRIMARY KEY (`id`),
  UNIQUE KEY `t_hall_info_1_UN` (`playerID`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='大厅信息表';


-- ----------------------------
-- Table structure for t_player
-- ----------------------------
DROP TABLE IF EXISTS `t_player`;
CREATE TABLE `t_player` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `accountID` bigint(20) NOT NULL COMMENT '账户ID',
  `playerID` bigint(20) NOT NULL COMMENT '玩家ID',
  `showUID` bigint(20) NOT NULL COMMENT '显示ID',
  `type` int(11) NOT NULL DEFAULT '1' COMMENT '玩家类型1.普通玩家，2.机器人，3.QA\n2.\n3.',
  `channelID` int(11) DEFAULT NULL COMMENT '渠道ID',
  `nickname` varchar(64) DEFAULT NULL COMMENT '昵称',
  `gender` int(11) DEFAULT '1' COMMENT '性别',
  `avatar` varchar(256) DEFAULT NULL COMMENT '头像地址',
  `provinceID` int(11) DEFAULT NULL COMMENT '省ID',
  `cityID` int(11) DEFAULT NULL COMMENT '市ID',
  `name` varchar(64) DEFAULT NULL COMMENT '真实姓名',
  `phone` varchar(11) DEFAULT NULL COMMENT '手机号码',
  `idCard` varchar(20) DEFAULT NULL COMMENT '身份证',
  `isWhiteList` tinyint(1) DEFAULT '0' COMMENT '是否QA，默认否',
  `zipCode` int(11) DEFAULT NULL COMMENT '邮编',
  `shippingAddr` varchar(256) DEFAULT NULL COMMENT '收获地址',
  `status` int(11) DEFAULT '1' COMMENT '1可登录，2冻结，默认为1',
  `remark` varchar(256) DEFAULT NULL COMMENT '备注',
  `createTime` datetime DEFAULT NULL COMMENT '创建时间，通常也是注册时间',
  `createBy` varchar(64) DEFAULT NULL COMMENT '创建人',
  `updateTime` datetime DEFAULT NULL COMMENT '更新时间',
  `updateBy` varchar(64) DEFAULT NULL COMMENT '更新人',
  PRIMARY KEY (`id`),
  UNIQUE KEY `t_player_1_UN_playerID` (`playerID`),
  UNIQUE KEY `t_player_1_UN_showUID` (`showUID`),
  UNIQUE KEY `t_player_1_UN_ACCOUTID` (`accountID`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='玩家表';

-- ----------------------------
-- Table structure for t_player_currency
-- ----------------------------
DROP TABLE IF EXISTS `t_player_currency`;
CREATE TABLE `t_player_currency` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `playerID` bigint(20) NOT NULL COMMENT '玩家ID',
  `coins` int(11) DEFAULT NULL COMMENT '当前金币数',
  `ingots` int(11) DEFAULT NULL COMMENT '当前元宝数',
  `keyCards` int(11) DEFAULT NULL COMMENT '当前房卡数',
  `obtainIngots` int(11) DEFAULT NULL COMMENT '总获得元宝数',
  `obtainKeyCards` int(11) DEFAULT NULL COMMENT '总获得房卡数',
  `costIngots` int(11) DEFAULT NULL COMMENT '累计消耗元宝数',
  `costKeyCards` int(11) DEFAULT NULL COMMENT '累计消耗房卡数',
  `remark` varchar(256) DEFAULT NULL COMMENT '备注',
  `createTime` datetime DEFAULT NULL COMMENT '创建时间',
  `createBy` varchar(64) DEFAULT NULL COMMENT '创建人',
  `updateTime` datetime DEFAULT NULL COMMENT '更新时间',
  `updateBy` varchar(64) DEFAULT NULL COMMENT '更新人',
  PRIMARY KEY (`id`),
  UNIQUE KEY `t_player_currency_1_UN_playerID` (`playerID`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='玩家货币表';

-- ----------------------------
-- Table structure for t_player_game
-- ----------------------------
DROP TABLE IF EXISTS `t_player_game`;
CREATE TABLE `t_player_game` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `playerID` bigint(20) NOT NULL COMMENT '玩家ID',
  `gameID` int(11) DEFAULT NULL COMMENT '游戏ID',
  `gameName` varchar(64) DEFAULT NULL COMMENT '游戏名称',
  `winningRate` double DEFAULT NULL COMMENT '胜率，百分比表示，50%，只记录 50，精确到个位数',
  `winningBurea` int(11) DEFAULT NULL COMMENT '胜利局数',
  `totalBureau` int(11) DEFAULT NULL COMMENT '总局数',
  `maxWinningStream` int(11) DEFAULT NULL COMMENT '最高连胜',
  `maxMultiple` int(11) DEFAULT NULL COMMENT '最大倍数',
  `createTime` datetime DEFAULT NULL COMMENT '创建时间',
  `createBy` varchar(64) DEFAULT NULL COMMENT '创建人',
  `updateTime` datetime DEFAULT NULL COMMENT '更新时间',
  `updateBy` varchar(64) DEFAULT NULL COMMENT '更新人',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='玩家游戏信息表';

-- ----------------------------
-- Table structure for t_player_mail
-- ----------------------------
DROP TABLE IF EXISTS `t_player_mail`;
CREATE TABLE `t_player_mail` (
  `n_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '递增ID',
  `n_playerid` bigint(20) NOT NULL COMMENT '玩家ID',
  `n_mailID` bigint(20) NOT NULL COMMENT '邮件ID',
  `n_isRead` tinyint(1) DEFAULT NULL COMMENT '是否已读: 0=未读, 1=已读 ',
  `n_isGetAttach` tinyint(1) DEFAULT NULL COMMENT '是否已领取附件: 0=未领, 1=已领',
  `n_isDel` tinyint(1) DEFAULT '0' COMMENT '是否被用户删除: 0=未删除, 1=删除',
  `n_deleteTime` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`n_id`),
  UNIQUE KEY `t_player_mail_UN` (`n_playerid`,`n_mailID`),
  KEY `t_player_mail_n_playerid_IDX` (`n_playerid`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='玩家邮件表';

-- ----------------------------
-- Table structure for t_player_props
-- ----------------------------
DROP TABLE IF EXISTS `t_player_props`;
CREATE TABLE `t_player_props` (
  `playerID` bigint(20) NOT NULL COMMENT '玩家ID',
  `propID` bigint(20) NOT NULL COMMENT '道具ID',
  `count` bigint(20) NOT NULL COMMENT '道具数量',
  `createTime` datetime DEFAULT NULL COMMENT '创建时间',
  `createBy` varchar(100) DEFAULT NULL COMMENT '创建人',
  `updateTime` datetime DEFAULT NULL COMMENT '更新时间',
  `updateBy` varchar(100) DEFAULT NULL COMMENT '更新人',
  PRIMARY KEY (`playerID`,`propID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='玩家道具表';


-- ----------------------------
-- Table structure for t_show_id
-- ----------------------------
DROP TABLE IF EXISTS `t_show_id`;
CREATE TABLE `t_show_id` (
  `n_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '递增ID',
  `n_showid` bigint(20) NOT NULL COMMENT 'show id 值',
  `n_isUse` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否被使用',
  PRIMARY KEY (`n_id`),
  UNIQUE KEY `t_show_id_UN_showid` (`n_showid`),
  KEY `t_show_id_n_isUse_IDX` (`n_isUse`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='showid记录表';

-- ----------------------------
-- Table structure for t_player_id
-- ----------------------------
DROP TABLE IF EXISTS `t_player_id`;
CREATE TABLE `t_player_id` (
  `n_id` bigint(20) NOT NULL COMMENT '通用变量ID',
  `n_value` bigint(20) DEFAULT '0' COMMENT '变量值',
  `n_des` varchar(255) DEFAULT NULL COMMENT '变量描述',
  PRIMARY KEY (`n_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='playerid表';

SET FOREIGN_KEY_CHECKS = 1;


-- ----------------------------
-- Table structure for t_player_packsack
-- ----------------------------
DROP TABLE IF EXISTS `t_player_packsack`;
CREATE TABLE `t_player_packsack`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `playerID` bigint(20) NOT NULL COMMENT '玩家ID',
  `gold` int(11) NULL DEFAULT NULL COMMENT '背包金币数',
  `createTime` datetime(0) NULL DEFAULT NULL,
  `createBy` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `updateTime` datetime(0) NULL DEFAULT NULL,
  `updateBy` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 175 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;