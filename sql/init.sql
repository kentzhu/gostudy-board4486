/*
 Navicat Premium Data Transfer

 Source Server         : 爱碰-开发环境
 Source Server Type    : MySQL
 Source Server Version : 50623
 Source Host           : 192.168.10.10:3306
 Source Schema         : golang-board4486

 Target Server Type    : MySQL
 Target Server Version : 50623
 File Encoding         : 65001

 Date: 09/03/2021 01:22:43
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admin
-- ----------------------------
DROP TABLE IF EXISTS `admin`;
CREATE TABLE `admin`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `ctime` timestamp(0) NULL DEFAULT CURRENT_TIMESTAMP(0),
  `etime` timestamp(0) NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0),
  `username` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `password` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `username`(`username`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of admin
-- ----------------------------
INSERT INTO `admin` VALUES (1, '2021-03-09 01:00:29', '2021-03-09 01:00:34', 'admin', 'admin987654321');

-- ----------------------------
-- Table structure for message
-- ----------------------------
DROP TABLE IF EXISTS `message`;
CREATE TABLE `message`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `ctime` timestamp(0) NULL DEFAULT CURRENT_TIMESTAMP(0),
  `etime` timestamp(0) NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0),
  `title` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `author` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 17 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of message
-- ----------------------------
INSERT INTO `message` VALUES (1, '2021-03-08 13:57:16', '2021-03-08 13:57:58', 'first message', 'this is my first message', 'Kent');
INSERT INTO `message` VALUES (2, '2021-03-08 15:09:22', '2021-03-08 15:10:07', 'second message', 'this is the second message', 'Li');
INSERT INTO `message` VALUES (3, '2021-03-08 15:40:17', '2021-03-08 15:40:17', 'haha', 'hahah just test', 'postman');
INSERT INTO `message` VALUES (7, '2021-03-08 18:07:11', '2021-03-08 18:07:11', 'Your', 'KKKKKK', 'Man');
INSERT INTO `message` VALUES (13, '2021-03-08 22:12:40', '2021-03-08 22:12:40', '阿斯蒂芬斯蒂芬', '阿斯顿发发呆发送到发顺丰', '打发斯蒂芬');
INSERT INTO `message` VALUES (16, '2021-03-09 01:16:17', '2021-03-09 01:16:17', '真的啊', '绝了呀', '真的吗');

SET FOREIGN_KEY_CHECKS = 1;
