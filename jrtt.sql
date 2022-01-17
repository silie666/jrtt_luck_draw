/*
 Navicat MySQL Data Transfer

 Source Server         : myaliyun
 Source Server Type    : MySQL
 Source Server Version : 80025
 Source Host           : 47.106.136.31:3306
 Source Schema         : jrtt

 Target Server Type    : MySQL
 Target Server Version : 80025
 File Encoding         : 65001

 Date: 17/01/2022 09:18:05
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for jr_list
-- ----------------------------
DROP TABLE IF EXISTS `jr_list`;
CREATE TABLE `jr_list`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `search_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '',
  `query_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '',
  `search_result_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT 'api id',
  `is_winner` char(5) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '0' COMMENT '是否中奖',
  `lottery_time` datetime NULL DEFAULT NULL COMMENT '开奖时间',
  `participate_type` tinyint(1) NULL DEFAULT NULL COMMENT '参与方式 1-转发',
  `reward` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '奖品文本',
  `status` tinyint NULL DEFAULT NULL COMMENT '1-未开将  2-已开奖',
  `winner_type` tinyint NULL DEFAULT NULL COMMENT '参与范围 1-所有人  2-仅粉丝',
  `detail` mediumtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '详情',
  `luck_data` mediumtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '抽奖数据',
  `is_ok` tinyint(1) NULL DEFAULT 0 COMMENT '是否完成抽奖',
  `is_repost` tinyint(1) NULL DEFAULT 0 COMMENT '是否转发',
  `is_like` tinyint(1) NULL DEFAULT 0 COMMENT '是否点赞',
  `user_id` bigint NULL DEFAULT NULL COMMENT '博主id',
  `user_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '名称',
  `zhuanfa_uid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '完成抽奖的uid',
  `create_time` datetime NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 640 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for jr_log
-- ----------------------------
DROP TABLE IF EXISTS `jr_log`;
CREATE TABLE `jr_log`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `log` mediumtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `create_time` datetime NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 34 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for jr_user
-- ----------------------------
DROP TABLE IF EXISTS `jr_user`;
CREATE TABLE `jr_user`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `uid` bigint NULL DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '链接',
  `is_modify` tinyint(1) NULL DEFAULT 0 COMMENT '是否关注',
  `zhuanfa_uid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `token` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '用户加密token',
  `create_time` datetime NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 448 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
