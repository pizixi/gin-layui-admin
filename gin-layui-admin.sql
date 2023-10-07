/*
 Navicat Premium Data Transfer

 Source Server         : 192.168.1.37test
 Source Server Type    : MySQL
 Source Server Version : 50726
 Source Host           : 192.168.1.37:3306
 Source Schema         : gin-layui-admin

 Target Server Type    : MySQL
 Target Server Version : 50726
 File Encoding         : 65001

 Date: 07/10/2023 18:48:30
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for kgo_agent
-- ----------------------------
DROP TABLE IF EXISTS `kgo_agent`;
CREATE TABLE `kgo_agent`  (
  `aid` int(10) UNSIGNED NOT NULL COMMENT '平台id(对应conf的plat)',
  `flag` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '平台标示, eg.: dalan',
  `name` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '平台名称, eg.: 大蓝',
  `lang` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '平台语言',
  `miniapp` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否是小程序（1=小程序）',
  `audit_version` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '当前提审版本',
  `last_audit` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '最后提审通过的版本',
  `vpsid` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'cdn所在vps',
  `source` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT 'cdn回源域名',
  `domain` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT 'cdn正式域名',
  `create_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `status` tinyint(1) UNSIGNED NOT NULL DEFAULT 1 COMMENT '状态1-正常，0-删除',
  PRIMARY KEY (`aid`) USING BTREE,
  UNIQUE INDEX `flag`(`flag`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '平台列表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of kgo_agent
-- ----------------------------

-- ----------------------------
-- Table structure for kgo_cdn
-- ----------------------------
DROP TABLE IF EXISTS `kgo_cdn`;
CREATE TABLE `kgo_cdn`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '流水号id',
  `aid` int(10) UNSIGNED NOT NULL COMMENT '平台id',
  `version` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT 'cdn版本, eg.: cn_andou_20190101',
  `create_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `status` tinyint(1) UNSIGNED NOT NULL DEFAULT 1 COMMENT '0=未安装,1=已安装',
  `install_log` text CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT 'CDN安装日志',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `version`(`version`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = 'CDN列表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of kgo_cdn
-- ----------------------------

-- ----------------------------
-- Table structure for kgo_cross
-- ----------------------------
DROP TABLE IF EXISTS `kgo_cross`;
CREATE TABLE `kgo_cross`  (
  `id` int(10) UNSIGNED NOT NULL COMMENT '流水号id',
  `plat` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '平台集合id(目前认为是aid，后期有需要再改)',
  `version` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '当前版本',
  `vpsid` int(10) UNSIGNED NOT NULL COMMENT '所在的vps',
  `port` smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '跨服服端口',
  `db_port` smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '游服mysql端口',
  `create_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `status` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态 0=未安装 1=已安装 2=运行中 3=已暂停',
  `procs` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '进程信息',
  `install_log` text CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '安装日志',
  `hoted` text CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '已安装的补丁',
  `hot_log` text CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '安装补丁日志',
  `start_log` text CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '启服日志',
  `stop_log` text CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '关服日志',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '跨服列表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of kgo_cross
-- ----------------------------

-- ----------------------------
-- Table structure for kgo_game
-- ----------------------------
DROP TABLE IF EXISTS `kgo_game`;
CREATE TABLE `kgo_game`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '流水号id',
  `aid` int(10) UNSIGNED NOT NULL COMMENT '平台id',
  `sid` int(10) UNSIGNED NOT NULL COMMENT '服务器id（大于10000表示外测服）',
  `serial` int(8) UNSIGNED NOT NULL COMMENT '游服唯一id，对应conf的serial',
  `gid` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '合服后的组id, 0=未被合服',
  `mid` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '数据源服务器ID(0=无导入)',
  `version` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '当前服务器版本',
  `name` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '服务器名称(例如：大鹏展翅)',
  `vpsid` int(10) UNSIGNED NOT NULL COMMENT '游服所在的vps',
  `create_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `install_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '安装时间',
  `update_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
  `port` smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '游服端口',
  `db_port` smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '游服mysql端口',
  `db_share` smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '共享数据库的id',
  `open_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '开服时间',
  `merge_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '合服时间',
  `is_tls` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否走tls',
  `domain` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '单服域名(eg.: s1-andou-jmxy.kgogame.com)',
  `procs` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '进程信息',
  `status` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态 0=未安装 1=已安装 2=运行中 3=已暂停',
  `nginx_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'nginx转发服务器id',
  `ws` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '游服websocket入口地址（支持域名和ip:port）',
  `single` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '游服single入口地址',
  `mode` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '客户端显示模式 1=新服 2=火爆 3=维护',
  `install_log` text CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '安装日志',
  `hoted` text CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '已安装的补丁',
  `hot_log` text CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '安装补丁日志',
  `start_log` text CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '启服日志',
  `stop_log` text CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '关服日志',
  `cid` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '跨服组id',
  `join_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '加入跨服的时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `server`(`aid`, `sid`) USING BTREE,
  UNIQUE INDEX `serial`(`serial`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '服务器列表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of kgo_game
-- ----------------------------

-- ----------------------------
-- Table structure for kgo_nginx
-- ----------------------------
DROP TABLE IF EXISTS `kgo_nginx`;
CREATE TABLE `kgo_nginx`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '流水号id',
  `aid` int(10) UNSIGNED NOT NULL COMMENT '平台id',
  `sid` int(10) UNSIGNED NOT NULL COMMENT '起始sid',
  `vpsid` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'nginx所在的vps',
  `domain` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '转发域名',
  `ws` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT 'websocket 转发域名',
  `single` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT 'single 转发域名',
  `create_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '添加时间',
  `status` tinyint(1) UNSIGNED NOT NULL DEFAULT 1 COMMENT '状态1-正常，0-删除',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx`(`aid`, `sid`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '平台nginx转发服务器' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of kgo_nginx
-- ----------------------------

-- ----------------------------
-- Table structure for kgo_vps
-- ----------------------------
DROP TABLE IF EXISTS `kgo_vps`;
CREATE TABLE `kgo_vps`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '流水号id',
  `ip` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT 'ip地址',
  `type` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '类型（1=game,2=cdn,4=nginx,8=center）',
  `create_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '注册时间',
  `domain` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '被控端域名',
  `detail` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '备注',
  `status` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态0=离线，1=在线',
  `version` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT 'oam-ctl 版本',
  `vps_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '当前时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `ip`(`ip`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '云服务器列表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of kgo_vps
-- ----------------------------

-- ----------------------------
-- Table structure for menu
-- ----------------------------
DROP TABLE IF EXISTS `menu`;
CREATE TABLE `menu`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `pid` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '上级ID，0为顶级',
  `auth_name` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '权限名称',
  `auth_url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT 'URL地址',
  `sort` int(11) UNSIGNED NOT NULL DEFAULT 999 COMMENT '排序，越小越前',
  `icon` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `is_show` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否显示，0-隐藏，1-显示',
  `auth_bit` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '菜单权限',
  `user_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '操作者ID',
  `create_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建者ID',
  `update_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '修改者ID',
  `status` tinyint(1) UNSIGNED NOT NULL DEFAULT 1 COMMENT '状态，1-正常，0-删除',
  `create_time` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `update_time` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `auth_url`(`auth_url`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 17 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '菜单列表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of menu
-- ----------------------------
INSERT INTO `menu` VALUES (1, 0, '所有权限', '', 1, '', 1, 7, 1, 1, 1, 1, 1505620970, 0);
INSERT INTO `menu` VALUES (2, 1, '权限管理', '/', 999, 'fa-id-card', 1, 1, 1, 1, 1, 1, 1505622360, 0);
INSERT INTO `menu` VALUES (3, 2, '用户管理', '/home/user', 1, 'fa-user-o', 1, 1, 0, 0, 1, 1, 1528385411, 0);
INSERT INTO `menu` VALUES (4, 2, '角色管理', '/home/role', 2, 'fa-user-circle-o', 1, 1, 0, 1, 1, 1, 1505621852, 0);
INSERT INTO `menu` VALUES (5, 2, '菜单管理', '/home/menu', 3, 'fa-list', 1, 1, 1, 1, 1, 1, 1505621986, 0);
INSERT INTO `menu` VALUES (6, 1, '运维管理', '/oam', 1, 'fa-tasks', 1, 0, 1, 1, 1, 1, 0, 0);
INSERT INTO `menu` VALUES (7, 6, 'vps列表', '/home/vps', 1, 'fa-cloud', 1, 0, 1, 1, 1, 1, 0, 1696667381);
INSERT INTO `menu` VALUES (8, 1, '个人中心', '/personal', 999, 'fa-user-circle-o', 1, 7, 1, 1, 1, 1, 1547000410, 0);
INSERT INTO `menu` VALUES (9, 6, '游服管理', '/home/game', 3, 'fa-server', 1, 0, 1, 1, 1, 1, 1546920455, 0);
INSERT INTO `menu` VALUES (10, 6, '合服管理', '/home/merge', 999, 'fa-compress', 0, 0, 1, 1, 1, 1, 1546920606, 0);
INSERT INTO `menu` VALUES (11, 6, 'cdn管理', '/home/cdn', 4, 'fa-file-code-o', 1, 0, 1, 1, 1, 1, 1546921447, 0);
INSERT INTO `menu` VALUES (12, 8, '资料修改', '/home/personal', 1, 'fa-edit', 1, 7, 1, 1, 1, 1, 1547000565, 1696675325);
INSERT INTO `menu` VALUES (13, 6, '平台列表', '/home/agent', 2, 'fa-map-signs', 1, 0, 1, 1, 1, 1, 1547724451, 0);
INSERT INTO `menu` VALUES (14, 6, '版本管理', '/home/version', 5, 'fa-file-archive-o', 1, 0, 1, 1, 1, 1, 1547785510, 0);
INSERT INTO `menu` VALUES (15, 6, 'nginx管理', '/home/nginx', 6, 'fa-sitemap', 1, 0, 1, 1, 1, 1, 1547785510, 0);
INSERT INTO `menu` VALUES (16, 1, '基本管理7', '/set7', 1, 'fa-tasks', 1, 4, 0, 1, 1, 0, 1696669069, 1696672348);

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `role_name` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '0' COMMENT '角色名称',
  `detail` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '0' COMMENT '备注',
  `create_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建者ID',
  `update_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '修改这ID',
  `status` tinyint(1) UNSIGNED NOT NULL DEFAULT 1 COMMENT '状态1-正常，0-删除',
  `create_time` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '添加时间',
  `update_time` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `role_name`(`role_name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '角色表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of role
-- ----------------------------
INSERT INTO `role` VALUES (1, '超级管理员', '超级管理员，具有所有权限', 1, 1, 1, 0, 1696666716);
INSERT INTO `role` VALUES (2, '普通管理员', '普通管理员，无菜单管理权限', 1, 1, 1, 0, 1696666747);
INSERT INTO `role` VALUES (3, '测试', '测试角色1', 1, 1, 1, 1696664214, 1696675302);

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户id',
  `role_id` int(11) UNSIGNED NOT NULL COMMENT '角色id',
  `login_name` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '用户名',
  `real_name` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '0' COMMENT '真实姓名',
  `password` char(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '密码',
  `phone` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '0' COMMENT '手机号码',
  `email` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '邮箱',
  `salt` char(10) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '密码盐',
  `last_login` int(11) NOT NULL DEFAULT 0 COMMENT '最后登录时间',
  `last_ip` char(15) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '最后登录IP',
  `status` tinyint(4) NOT NULL DEFAULT 0 COMMENT '状态，1-正常 0禁用',
  `create_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建者ID',
  `update_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '修改者ID',
  `create_time` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `update_time` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_user_name`(`login_name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '用户表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, 1, 'admin', 'admin', '97aae6915f853246abf22773112d2fb4', '13866668888', 'xxx@qq.com', '58QM', 1696674269, '127.0.0.1', 1, 0, 0, 0, 1548146771);
INSERT INTO `user` VALUES (2, 2, 'admin2', 'admin2', '97aae6915f853246abf22773112d2fb4', '13855556666', 'gjs3ob_t@streetsinus.com', '58QM', 1696599396, '127.0.0.1', 1, 1, 1, 1696598348, 1696598348);
INSERT INTO `user` VALUES (3, 2, 'admin3', 'admin3', '26a72db35baca65333d1e314ab70411f', '15395040869', 'cpktbf54906@chaco.net', 'uY7c', 1696663962, '127.0.0.1', 0, 1, 3, 1696649699, 1696664313);
INSERT INTO `user` VALUES (4, 3, 'admin4', 'admin4', '55ba69f874cf544b2fa46f4ac9412c54', '13178687958', 'cpktbf@chaco.net', 'pHRV', 1696667404, '127.0.0.1', 0, 1, 1, 1696664346, 1696674031);
INSERT INTO `user` VALUES (5, 1, 'admin5', 'admin5', '2fbf13521c2fce3aa5b42ef2b052cc4f', '15888888888', 'pizixi@qq.com', 'ujlR', 1696674238, '127.0.0.1', 1, 1, 1, 1696674224, 1696675277);

SET FOREIGN_KEY_CHECKS = 1;
