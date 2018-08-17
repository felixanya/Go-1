/*
 Navicat Premium Data Transfer

 Source Server         : 192.168.7.108
 Source Server Type    : MySQL
 Source Server Version : 50722
 Source Host           : 192.168.7.108:3306
 Source Schema         : config

 Target Server Type    : MySQL
 Target Server Version : 50722
 File Encoding         : 65001

 Date: 09/08/2018 11:35:02
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for t_common_config
-- ----------------------------
DROP TABLE IF EXISTS `t_common_config`;
CREATE TABLE `t_common_config` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `key` varchar(128) NOT NULL COMMENT 'config key',
  `subkey` varchar(128) NOT NULL COMMENT 'config sub key',
  `value` text COMMENT 'config context, json format',
  PRIMARY KEY (`id`),
  UNIQUE KEY `key` (`key`,`subkey`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='通用配置表';

-- ----------------------------
-- Table structure for t_game_config
-- ----------------------------
DROP TABLE IF EXISTS `t_game_config`;
CREATE TABLE `t_game_config` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `gameID` int(11) DEFAULT NULL COMMENT '游戏ID',
  `name` varchar(128) DEFAULT NULL COMMENT '游戏名称',
  `type` int(11) DEFAULT NULL COMMENT '游戏类型',
  `minPeople` int(11) DEFAULT NULL COMMENT '允许最少人数',
  `maxPeople` int(11) DEFAULT NULL COMMENT '允许最多人数',
  `playform` int(11) DEFAULT NULL COMMENT '平台,1:安卓;2:ios',
  `countryID` int(11) DEFAULT NULL COMMENT '国区（默认中国）',
  `provinceID` int(11) DEFAULT NULL COMMENT '省ID',
  `cityID` int(11) DEFAULT NULL COMMENT '市ID',
  `channelID` int(11) DEFAULT NULL COMMENT '渠道ID',
  `createTime` datetime DEFAULT NULL COMMENT '创建时间',
  `createBy` varchar(64) DEFAULT NULL COMMENT '创建人',
  `updateTime` datetime DEFAULT NULL COMMENT '更新时间',
  `updateBy` varchar(64) DEFAULT NULL COMMENT '更新人',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='游戏配置表';

-- ----------------------------
-- Table structure for t_game_level_config
-- ----------------------------
DROP TABLE IF EXISTS `t_game_level_config`;
CREATE TABLE `t_game_level_config`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `gameID` int(11) NULL DEFAULT NULL,
  `levelID` int(11) NULL DEFAULT NULL,
  `name` varchar(256) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `fee` int(11) NULL DEFAULT NULL COMMENT '费用',
  `baseScores` int(11) NULL DEFAULT NULL,
  `lowScores` int(11) NULL DEFAULT NULL,
  `highScores` int(11) NULL DEFAULT NULL,
  `realOnlinePeople` int(11) NULL DEFAULT NULL COMMENT '实时在线人数',
  `showOnlinePeople` int(11) NULL DEFAULT NULL COMMENT '显示在线人数',
  `status` int(11) NULL DEFAULT NULL,
  `tag` int(11) NULL DEFAULT NULL COMMENT '标签：1.热门；2.New',
  `isAlms` int(11) NULL DEFAULT NULL COMMENT '是否为救济金场，0：关闭，1：开启',
  `remark` varchar(256) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `createTime` datetime(0) NULL DEFAULT NULL,
  `createBy` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `updateTime` datetime(0) NULL DEFAULT NULL,
  `updateBy` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '游戏场次配置表' ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;

-- ----------------------------
-- Table structure for t_horse_race
-- ----------------------------
DROP TABLE IF EXISTS `t_horse_race`;
CREATE TABLE `t_horse_race` (
  `n_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '数据递增ID',
  `n_channel` bigint(20) NOT NULL COMMENT '渠道ID',
  `n_prov` bigint(20) DEFAULT NULL COMMENT '省包ID',
  `n_city` bigint(20) DEFAULT NULL COMMENT '城市ID',
  `n_bUse` tinyint(1) DEFAULT '1' COMMENT '是否启用',
  `n_bUseParent` tinyint(1) DEFAULT '1' COMMENT '是否启用上级配置',
  `n_horseData` text COMMENT 'json格式的跑马灯配置，具体格式参考相关说明文件',
  PRIMARY KEY (`n_id`),
  KEY `t_horse_race_n_channel_IDX` (`n_channel`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='跑马灯表';

-- ----------------------------
-- Table structure for t_mail
-- ----------------------------
DROP TABLE IF EXISTS `t_mail`;
CREATE TABLE `t_mail` (
  `n_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '递增ID',
  `n_title` varchar(150) DEFAULT NULL COMMENT '邮件标题',
  `n_detail` text COMMENT '邮件内容',
  `n_attach` varchar(256) DEFAULT NULL COMMENT '邮件附件：json格式 ',
  `n_dest` text COMMENT '发送对象:json格式',
  `n_state` int(11) NOT NULL COMMENT '邮件状态：未发送=0＞审核中=1＞已审核=2＞发送中=3＞发送结束=4＞已拒绝=5＞已撤回=6＞已失效=7 ',
  `n_starttime` datetime DEFAULT NULL COMMENT '发送开始时间: 2018-08-08 12:00:00',
  `n_endtime` datetime DEFAULT NULL COMMENT '发送截至时间: 2018-08-18 12:00:00',
  `n_deltime` datetime DEFAULT NULL COMMENT '邮件删除时间: 2018-09-18 12:00:00',
  `n_createTime` datetime DEFAULT NULL COMMENT '创建时间: 2018-08-08 12:00:00',
  `n_createBy` varchar(64) DEFAULT NULL COMMENT '创建人',
  `n_updateTime` datetime DEFAULT NULL COMMENT '最后更新时间: 2018-08-08 12:00:00',
  `n_updateBy` varchar(64) DEFAULT NULL COMMENT '最后更新人',
  `n_isUseEndTime` tinyint(1) DEFAULT '1' COMMENT '是否启用截至时间',
  `n_isUseDelTime` tinyint(1) DEFAULT '1' COMMENT '是否启用删除时间',
  PRIMARY KEY (`n_id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='系统消息表，邮件表'

SET FOREIGN_KEY_CHECKS = 1;