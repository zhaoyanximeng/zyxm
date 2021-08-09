/*
 Navicat Premium Data Transfer

 Source Server         : mysql57
 Source Server Type    : MySQL
 Source Server Version : 50721
 Source Host           : localhost:3307
 Source Schema         : p1

 Target Server Type    : MySQL
 Target Server Version : 50721
 File Encoding         : 65001

 Date: 07/12/2020 13:26:54
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for course_prices
-- ----------------------------
DROP TABLE IF EXISTS `course_prices`;
CREATE TABLE `course_prices`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `course_id` int(11) NOT NULL COMMENT '商品ID',
  `market_price` decimal(10, 2) NULL DEFAULT 0.00 COMMENT '市场价',
  `sale_price` decimal(10, 2) NULL DEFAULT 0.00 COMMENT '销售价',
  `setdate` datetime(0) NULL DEFAULT NULL COMMENT '价格设定时间',
  `comment` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  `iscurrent` tinyint(1) NULL DEFAULT '0' COMMENT '是否是当前价格',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
