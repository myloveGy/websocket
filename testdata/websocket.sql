SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;
DROP TABLE IF EXISTS `app`;
CREATE TABLE `app` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `app_id` varchar(32) NOT NULL COMMENT '应用ID',
  `app_secret` varchar(32) NOT NULL COMMENT '应用秘钥',
  `app_name` varchar(32) NOT NULL COMMENT '应用名称',
  `status` tinyint(1) NOT NULL COMMENT '状态 1 启用 4 停用',
  `created_at` datetime NOT NULL COMMENT '创建时间',
  `updated_at` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `unq_app_id` (`app_id`) COMMENT 'app_id必须唯一'
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='应用信息';

INSERT INTO `app` VALUES (1, '2020110306161001', '777c9ce7926298d2bf7930436d7dd03a', '主要应用', 1, '2020-11-03 18:16:31', '2020-11-03 18:16:34');

DROP TABLE IF EXISTS `message`;
CREATE TABLE `message` (
  `message_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `app_id` int(11) NOT NULL COMMENT '所属应用主键',
  `type` tinyint(1) NOT NULL DEFAULT '1' COMMENT '消息类型 1 临时 2 永久',
  `content` text NOT NULL COMMENT '信息内容',
  `created_at` datetime NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`message_id`),
  KEY `idx_user_id` (`type`),
  KEY `idx_create_at` (`created_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COMMENT='消息信息';
INSERT INTO `message` VALUES (1, 1, 1, '{\"username\":\"jinxing.liu\", \"age\": 28}', '2020-11-20 18:12:36');
INSERT INTO `message` VALUES (2, 1, 1, '{\"username\":\"jinxing.liu\", \"age\": 28}', '2020-11-20 18:12:59');
DROP TABLE IF EXISTS `message_read`;
CREATE TABLE `message_read` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `message_id` char(40) NOT NULL DEFAULT '',
  `app_id` int(11) NOT NULL COMMENT '所属应用主键',
  `user_id` varchar(40) NOT NULL DEFAULT '' COMMENT '接收者会话id',
  `group_id` varchar(100) NOT NULL DEFAULT '' COMMENT '分组id',
  `status` tinyint(11) NOT NULL DEFAULT '1' COMMENT '状态 1 未读 2 已读',
  `created_at` datetime NOT NULL COMMENT '创建时间',
  `updated_at` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `idx_message_id` (`message_id`),
  KEY `idx_session_id` (`user_id`),
  KEY `idx_group_id` (`group_id`),
  KEY `idx_create_at` (`created_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COMMENT='消息读取记录信息';
INSERT INTO `message_read` VALUES (1, '1', 1, '2', '', 1, '2020-11-20 18:12:37', '2020-11-20 18:12:37');
INSERT INTO `message_read` VALUES (2, '2', 1, '2', '', 1, '2020-11-20 18:12:59', '2020-11-20 18:12:59');
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `user_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `app_id` int(11) NOT NULL COMMENT '应用ID',
  `username` varchar(32) NOT NULL COMMENT '用户名称',
  `phone` char(11) NOT NULL COMMENT '手机号',
  `password` varchar(200) NOT NULL COMMENT '密码',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态 1 启用 2 停用',
  `access_token` varchar(100) NOT NULL DEFAULT '' COMMENT '登录access_token信息',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`user_id`),
  UNIQUE KEY `unq_phone` (`phone`) COMMENT '手机号唯一',
  UNIQUE KEY `unq_username` (`username`) COMMENT '用户名唯一'
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='用户信息';
INSERT INTO `user` VALUES (1, 1, 'jinxing.liu', '13020137932', 'v123456', 1, '0d91e6f2f98adb3b329e2ab96a527880', '2020-11-19 11:46:41', '2020-11-20 18:15:44');
SET FOREIGN_KEY_CHECKS = 1;
