/*
Navicat MySQL Data Transfer

Source Server         : 127.0.0.1
Source Server Version : 50726
Source Host           : localhost:3306
Source Database       : beego

Target Server Type    : MYSQL
Target Server Version : 50726
File Encoding         : 65001

Date: 2020-10-16 14:49:09
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for admin
-- ----------------------------
DROP TABLE IF EXISTS `admin`;
CREATE TABLE `admin` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_name` varchar(64) DEFAULT NULL,
  `password` varchar(64) DEFAULT NULL,
  `image` varchar(64) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of admin
-- ----------------------------

-- ----------------------------
-- Table structure for article
-- ----------------------------
DROP TABLE IF EXISTS `article`;
CREATE TABLE `article` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(10) DEFAULT NULL,
  `keys_id` varchar(64) DEFAULT NULL,
  `name` varchar(64) NOT NULL COMMENT '文章标题',
  `class_id` int(10) NOT NULL DEFAULT '1' COMMENT '分类id',
  `child_id` int(10) NOT NULL DEFAULT '0' COMMENT '二级分类',
  `content` text,
  `click` int(10) DEFAULT '10' COMMENT '点击量',
  `image` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`,`class_id`)
) ENGINE=MyISAM AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COMMENT='文章表';

-- ----------------------------
-- Records of article
-- ----------------------------
INSERT INTO `article` VALUES ('1', null, null, '在线工具', '1', '3', '<p style=\"margin-top: 5px; margin-bottom: 5px; text-align: center;\">工具</p><p style=\"margin-top: 5px; margin-bottom: 5px;\"><span style=\"text-decoration-line: underline;\">markdown&nbsp;</span></p><p style=\"text-align:center\"><img src=\"http://127.0.0.1:8080/static/uploads/file_edit2020-10-16/jpg_1602820180.jpg\" alt=\"\"/></p><p style=\"margin-bottom: 5px; text-align: left; margin-top: 5px;\"><span style=\"text-decoration: underline;\">markdown&nbsp;</span></p><p style=\"margin-bottom: 5px; text-align: left; margin-top: 5px;\"><span style=\"text-decoration: underline;\"></span></p><table><tbody><tr class=\"firstRow\"><td width=\"291\" valign=\"top\" style=\"word-break: break-all;\">属性</td><td width=\"291\" valign=\"top\" style=\"word-break: break-all;\">值</td><td width=\"291\" valign=\"top\"><br/></td><td width=\"291\" valign=\"top\"><br/></td><td width=\"291\" valign=\"top\"><br/></td></tr><tr><td width=\"291\" valign=\"top\"><br/></td><td width=\"291\" valign=\"top\"><br/></td><td width=\"291\" valign=\"top\"><br/></td><td width=\"291\" valign=\"top\"><br/></td><td width=\"291\" valign=\"top\"><br/></td></tr><tr><td width=\"291\" valign=\"top\"><br/></td><td width=\"291\" valign=\"top\"><br/></td><td width=\"291\" valign=\"top\"><br/></td><td width=\"291\" valign=\"top\"><br/></td><td width=\"291\" valign=\"top\"><br/></td></tr><tr><td width=\"291\" valign=\"top\"><br/></td><td width=\"291\" valign=\"top\"><br/></td><td width=\"291\" valign=\"top\"><br/></td><td width=\"291\" valign=\"top\"><br/></td><td width=\"291\" valign=\"top\"><br/></td></tr><tr><td width=\"291\" valign=\"top\"><br/></td><td width=\"291\" valign=\"top\"><br/></td><td width=\"291\" valign=\"top\"><br/></td><td width=\"291\" valign=\"top\"><br/></td><td width=\"291\" valign=\"top\"><br/></td></tr></tbody></table><p style=\"margin-bottom: 5px; text-align: left; margin-top: 5px;\"><span style=\"text-decoration: underline;\"></span><br/></p>', '100', '/static/uploads/file2020-10-16/jpg_1602820446.jpg', '2020-10-16 03:53:51');
INSERT INTO `article` VALUES ('2', null, null, 'php', '2', '5', '<ol class=\" list-paddingleft-2\" style=\"list-style-type: decimal;\"><li><p><span style=\"color: #F73131; font-family: Arial, sans-serif; font-size: 13px; background-color: #FFFFFF;\">PHP</span><span style=\"color: #333333; font-family: Arial, sans-serif; font-size: 13px; background-color: #FFFFFF;\">即“超文本预处理器”，是一种通用开源脚本语言。</span></p></li><li><p><span style=\"color: #333333; font-family: Arial, sans-serif; font-size: 13px; background-color: #FFFFFF;\"></span><span style=\"color: #F73131; font-family: Arial, sans-serif; font-size: 13px; background-color: #FFFFFF;\">PHP</span><span style=\"color: #333333; font-family: Arial, sans-serif; font-size: 13px; background-color: #FFFFFF;\">是在服务器端执行的脚本语言，与C语言类似，是常用的网站编程语言。</span></p></li><li><p><span style=\"color: #F73131; font-family: Arial, sans-serif; font-size: 13px; background-color: #FFFFFF;\">PHP</span><span style=\"color: #333333; font-family: Arial, sans-serif; font-size: 13px; background-color: #FFFFFF;\">独特的语法混合了C、Java、Perl以及&nbsp;</span><span style=\"color: #F73131; font-family: Arial, sans-serif; font-size: 13px; background-color: #FFFFFF;\">PHP</span><span style=\"color: #333333; font-family: Arial, sans-serif; font-size: 13px; background-color: #FFFFFF;\">&nbsp;自创的语法。利于学习，使用广泛，主要适用于Web开发领域。</span></p></li></ol><p><br/></p><p><iframe class=\"ueditor_baidumap\" src=\"http://127.0.0.1:8080/resources/ueditor/dialogs/map/show.html#center=118.079879,38.086122&zoom=7&width=530&height=340&markers=116.404,39.915&markerStyles=l,A\" frameborder=\"0\" width=\"534\" height=\"344\"></iframe></p>', '10', '/static/uploads/file2020-10-16/jpg_1602820591.jpg', '2020-10-16 03:58:49');

-- ----------------------------
-- Table structure for class
-- ----------------------------
DROP TABLE IF EXISTS `class`;
CREATE TABLE `class` (
  `class_id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `class_name` varchar(64) NOT NULL DEFAULT '' COMMENT '分类名称',
  `link` varchar(255) NOT NULL DEFAULT 'http://blog.tfuu.cn' COMMENT '链接',
  `class_path` varchar(64) DEFAULT NULL COMMENT '父级id',
  `sort` int(10) NOT NULL DEFAULT '0' COMMENT '排序',
  `pid` int(10) NOT NULL DEFAULT '0' COMMENT '0一级分类',
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`class_id`),
  KEY `pid` (`pid`)
) ENGINE=MyISAM AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COMMENT='分类表';

-- ----------------------------
-- Records of class
-- ----------------------------
INSERT INTO `class` VALUES ('1', 'Linux', 'http://blog.tfuu.cn', null, '0', '0', null);
INSERT INTO `class` VALUES ('2', 'WEB', 'http://blog.tfuu.cn', null, '0', '0', null);
INSERT INTO `class` VALUES ('3', '前端', 'http://blog.tfuu.cn', null, '0', '1', null);
INSERT INTO `class` VALUES ('4', 'GO', 'http://blog.tfuu.cn', null, '0', '2', null);
INSERT INTO `class` VALUES ('5', 'PHP', 'http://blog.tfuu.cn', null, '0', '2', null);
INSERT INTO `class` VALUES ('6', '工具', 'http://blog.tfuu.cn', null, '0', '0', null);
INSERT INTO `class` VALUES ('7', '关于', 'http://blog.tfuu.cn', null, '0', '0', null);
INSERT INTO `class` VALUES ('8', '二维码生成', 'http://blog.tfuu.cn', null, '0', '6', null);
INSERT INTO `class` VALUES ('9', '时间戳转换', 'http://blog.tfuu.cn', null, '0', '6', null);
INSERT INTO `class` VALUES ('10', '关于我们', 'http://blog.tfuu.cn', null, '0', '7', null);

-- ----------------------------
-- Table structure for config
-- ----------------------------
DROP TABLE IF EXISTS `config`;
CREATE TABLE `config` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `config_key` varchar(255) NOT NULL DEFAULT '' COMMENT '配置key',
  `config_value` varchar(255) NOT NULL DEFAULT '' COMMENT '配置值',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COMMENT='配置表';

-- ----------------------------
-- Records of config
-- ----------------------------
INSERT INTO `config` VALUES ('1', 'host', 'http://www.tfuu.cn');
INSERT INTO `config` VALUES ('2', 'home_banner', 'https://images.tfuu.cn/home_banner.jpg');
INSERT INTO `config` VALUES ('3', 'banner', 'https://images.tfuu.cn/banner.jpg');
INSERT INTO `config` VALUES ('4', 'title_1', '技匠精神');
INSERT INTO `config` VALUES ('5', 'des_1', '精益求精，注重细节，追求完美和极致，不惜花费时间精力，孜孜不倦，反复改进产品，把99%提高到99.99%。');

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `mobile` varchar(11) DEFAULT NULL,
  `password` varchar(64) DEFAULT NULL,
  `name` varchar(64) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES ('1', '4554', null, null, null);
