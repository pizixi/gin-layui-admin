SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

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
INSERT INTO `menu` VALUES (8, 1, '个人中心', '/personal', 999, 'fa-user-circle-o', 1, 7, 1, 1, 1, 1, 1547000410, 0);
INSERT INTO `menu` VALUES (12, 8, '资料修改', '/home/personal', 1, 'fa-edit', 1, 7, 1, 1, 1, 1, 1547000565, 1696675325);

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
INSERT INTO `user` VALUES (4, 3, 'admin4', 'admin4', '55ba69f874cf544b2fa46f4ac9412c54', '13178687958', 'cpktbf@chaco.net', 'pHRV', 1696667404, '127.0.0.1', 0, 1, 1, 1696664346, 1696674031);

SET FOREIGN_KEY_CHECKS = 1;
