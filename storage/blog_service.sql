/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50736
 Source Host           : localhost:3306
 Source Schema         : blog_service

 Target Server Type    : MySQL
 Target Server Version : 50736
 File Encoding         : 65001

 Date: 17/03/2022 14:38:33
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for blog_article
-- ----------------------------
DROP TABLE IF EXISTS `blog_article`;
CREATE TABLE `blog_article`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '文章标题',
  `desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '文章简述',
  `cover_image_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '封面图片地址',
  `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '文章内容',
  `created_on` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '创建时间',
  `created_by` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '创建人',
  `modified_on` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '修改时间',
  `modified_by` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '修改人',
  `deleted_on` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '删除时间',
  `is_del` tinyint(3) UNSIGNED NULL DEFAULT 0 COMMENT '是否删除 0 为未删除、1 为已删除',
  `state` tinyint(3) UNSIGNED NULL DEFAULT 1 COMMENT '状态 0 为禁用、1 为启用',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '文章管理' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for blog_article_tag
-- ----------------------------
DROP TABLE IF EXISTS `blog_article_tag`;
CREATE TABLE `blog_article_tag`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `article_id` int(11) NOT NULL COMMENT '文章 ID',
  `tag_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '标签 ID',
  `created_on` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '创建时间',
  `created_by` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '创建人',
  `modified_on` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '修改时间',
  `modified_by` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '修改人',
  `deleted_on` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '删除时间',
  `is_del` tinyint(3) UNSIGNED NULL DEFAULT 0 COMMENT '是否删除 0 为未删除、1 为已删除',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '文章标签关联' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for blog_auth
-- ----------------------------
DROP TABLE IF EXISTS `blog_auth`;
CREATE TABLE `blog_auth`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `app_key` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT 'Key',
  `app_secret` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT 'Secret',
  `created_on` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '创建时间',
  `created_by` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '创建人',
  `modified_on` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '修改时间',
  `modified_by` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '修改人',
  `deleted_on` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '删除时间',
  `is_del` tinyint(3) UNSIGNED NULL DEFAULT 0 COMMENT '是否删除 0 为未删除、1 为已删除',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '认证管理' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of blog_auth
-- ----------------------------
INSERT INTO `blog_auth` VALUES (1, 'eddycjy', 'go-programming-tour-book', 0, 'eddycjy', 0, '', 0, 0);

-- ----------------------------
-- Table structure for blog_message
-- ----------------------------
DROP TABLE IF EXISTS `blog_message`;
CREATE TABLE `blog_message`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `tag_id` int(10) UNSIGNED NOT NULL COMMENT '标签ID',
  `title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '文章标题',
  `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '文章内容',
  `created_on` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `created_by` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '创建人',
  `modified_on` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '修改时间',
  `modified_by` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '修改人',
  `deleted_on` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '删除时间',
  `is_del` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除 0 为未删除、1 为已删除',
  `state` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '状态 0 未通知、1 已通知',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '文章管理' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of blog_message
-- ----------------------------
INSERT INTO `blog_message` VALUES (1, 1, 'GoLang', '测试', 1645705040, '白月光', 1645705040, '', 0, 0, 1);
INSERT INTO `blog_message` VALUES (2, 1, 'GoLang', '测试', 1645705048, '白月光', 1645705048, '', 0, 0, 1);
INSERT INTO `blog_message` VALUES (3, 1, 'GoLang', '测试', 1645760153, '白月光', 1645783590, '', 0, 0, 1);
INSERT INTO `blog_message` VALUES (4, 1, 'GoLang', '测试', 1645760162, '白月光', 1645783590, '', 0, 0, 1);
INSERT INTO `blog_message` VALUES (5, 1, 'GoLang', '测试', 1645760164, '白月光', 1645783590, '', 0, 0, 1);
INSERT INTO `blog_message` VALUES (6, 1, 'GoLang', '测试', 1645760300, '白月光', 1645784480, '', 0, 0, 1);
INSERT INTO `blog_message` VALUES (7, 1, 'GoLang', '测试', 1645760365, '白月光', 1645784481, '', 0, 0, 1);
INSERT INTO `blog_message` VALUES (8, 1, 'GoLang', '测试', 1646223581, '白月光', 1646223585, '', 0, 0, 1);

-- ----------------------------
-- Table structure for blog_message_tag
-- ----------------------------
DROP TABLE IF EXISTS `blog_message_tag`;
CREATE TABLE `blog_message_tag`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `title` varchar(90) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '信息标题',
  `state` tinyint(3) NOT NULL DEFAULT 0 COMMENT '0 未设置 1邮件 2 钉钉 3 短信',
  `created_on` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '创建时间',
  `created_by` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '创建人',
  `modified_on` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '修改时间',
  `modified_by` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '修改人',
  `deleted_on` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '删除时间',
  `is_del` tinyint(3) UNSIGNED NULL DEFAULT 0 COMMENT '是否删除 0 为未删除、1 为已删除',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '信息标签表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of blog_message_tag
-- ----------------------------
INSERT INTO `blog_message_tag` VALUES (1, 'GoLang', 1, 1645705040, '白月光', 1645705040, '', 0, 0);

-- ----------------------------
-- Table structure for blog_system_user
-- ----------------------------
DROP TABLE IF EXISTS `blog_system_user`;
CREATE TABLE `blog_system_user`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `username` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '密码',
  `salt` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '盐值',
  `created_on` int(11) NULL DEFAULT NULL COMMENT '创建时间',
  `created_by` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '创建人',
  `modified_on` int(11) NULL DEFAULT NULL COMMENT '修改时间',
  `modified_by` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '修改人',
  `deleted_on` int(10) NULL DEFAULT NULL COMMENT '删除时间',
  `is_del` tinyint(4) NULL DEFAULT NULL COMMENT '状态 0 为禁用、1 为启用',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of blog_system_user
-- ----------------------------
INSERT INTO `blog_system_user` VALUES (1, 'admin', '$2a$04$rZamaBkfKycDephbJOKIoOCNMIivSkFF6cmwoV1IdafF6dABuP5li', NULL, 1646273831, '白月光', 1646273831, '', 0, 1);
INSERT INTO `blog_system_user` VALUES (2, '白月光', '$2a$04$rZamaBkfKycDephbJOKIoOCNMIivSkFF6cmwoV1IdafF6dABuP5li', NULL, 1646273831, '白月光', 1646273831, '', 0, 1);

-- ----------------------------
-- Table structure for blog_tag
-- ----------------------------
DROP TABLE IF EXISTS `blog_tag`;
CREATE TABLE `blog_tag`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '标签名称',
  `created_on` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '创建时间',
  `created_by` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '创建人',
  `modified_on` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '修改时间',
  `modified_by` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '修改人',
  `deleted_on` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '删除时间',
  `is_del` tinyint(3) UNSIGNED NULL DEFAULT 0 COMMENT '是否删除 0 为未删除、1 为已删除',
  `state` tinyint(3) UNSIGNED NULL DEFAULT 1 COMMENT '状态 0 为禁用、1 为启用',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '标签管理' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of blog_tag
-- ----------------------------
INSERT INTO `blog_tag` VALUES (1, 'Goland', 1637551360, '白月光', 1644317888, 'eddycjy', 1644300298, 0, 0);
INSERT INTO `blog_tag` VALUES (2, 'GoLang', 1637551637, '白月光', 1637551637, '', 0, 0, 1);
INSERT INTO `blog_tag` VALUES (3, 'Go', 1637552026, '白月光', 1637552026, '', 0, 0, 1);
INSERT INTO `blog_tag` VALUES (4, 'GoLang', 1637552044, '白月光', 1637552044, '', 0, 0, 1);
INSERT INTO `blog_tag` VALUES (5, 'GoLang', 1637565079, '白月光', 1637565079, '', 0, 0, 1);
INSERT INTO `blog_tag` VALUES (6, 'GoLang', 1637565106, '白月光', 1637565106, '', 0, 0, 1);
INSERT INTO `blog_tag` VALUES (7, 'Go2', 1644298554, 'eddycjy', 1644298554, '', 0, 0, 1);

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule`  (
  `p_type` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '规则类型',
  `v0` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '角色ID',
  `v1` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'api路径',
  `v2` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'api访问方法',
  `v3` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `v4` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `v5` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '权限规则表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/casbin', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/casbin/list', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/hello', 'GET', '', '', '');

SET FOREIGN_KEY_CHECKS = 1;
