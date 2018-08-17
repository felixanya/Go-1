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

SET FOREIGN_KEY_CHECKS = 1;