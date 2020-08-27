/*
Navicat MySQL Data Transfer

Source Server         : 127.0.0.1
Source Server Version : 50726
Source Host           : localhost:3306
Source Database       : beego

Target Server Type    : MYSQL
Target Server Version : 50726
File Encoding         : 65001

Date: 2020-08-27 20:02:39
*/

SET FOREIGN_KEY_CHECKS=0;

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
  `content` text,
  `click` int(10) DEFAULT '10' COMMENT '点击量',
  `image` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`,`class_id`)
) ENGINE=MyISAM AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COMMENT='文章表';

-- ----------------------------
-- Records of article
-- ----------------------------
INSERT INTO `article` VALUES ('1', null, null, 'php', '1', '09 Jul 2020\r\nPHP 7.4.8 Released!\r\nThe PHP development team announces the immediate availability of PHP 7.4.8. This is a security release impacting the official Windows builds of PHP.\r\n\r\nFor windows users running an official build, this release contains a patched version of libcurl addressing CVE-2020-8169.\r\n\r\nFor all other consumers of PHP, this is a bug fix release.\r\n\r\nAll PHP 7.4 users are encouraged to upgrade to this version.\r\n\r\nFor source downloads of PHP 7.4.8 please visit our downloads page, Windows source and binaries can be found on windows.php.net/download/. The list of changes is recorded in the ChangeLog.', '10', 'https://www.php.net/images/logos/php-logo.svg', '2020-07-23 09:33:42');
INSERT INTO `article` VALUES ('2', null, null, 'go', '1', 'Go（又称Golang）是Google开发的一种静态强类型、编译型、并发型，并具有垃圾回收功能的编程语言。\r\n罗伯特·格瑞史莫（Robert Griesemer），罗勃·派克（Rob Pike）及肯·汤普逊（Ken Thompson）于2007年9月开始设计Go，稍后Ian Lance Taylor、Russ Cox加入项目。Go是基于Inferno操作系统所开发的。Go于2009年11月正式宣布推出，成为开放源代码项目，并在Linux及Mac OS X平台上进行了实现，后来追加了Windows系统下的实现。在2016年，Go被软件评价公司TIOBE 选为“TIOBE 2016 年最佳语言”。 目前，Go每半年发布一个二级版本（即从a.x升级到a.y）。 [1]', '10', 'http://c.biancheng.net/uploads/allimg/181214/1-1Q214150602U9.jpg', '2020-07-05 09:41:31');
INSERT INTO `article` VALUES ('3', null, null, 'java', '1', 'ava是一门面向对象编程语言，不仅吸收了C++语言的各种优点，还摒弃了C++里难以理解的多继承、指针等概念，因此Java语言具有功能强大和简单易用两个特征。Java语言作为静态面向对象编程语言的代表，极好地实现了面向对象理论，允许程序员以优雅的思维方式进行复杂的编程 [1]  。\r\nJava具有简单性、面向对象、分布式、健壮性、安全性、平台独立与可移植性、多线程、动态性等特点 [2]  。Java可以编写桌面应用程序、Web应用程序、分布式系统和嵌入式系统应用程序等 [3]  。', '10', 'https://bkimg.cdn.bcebos.com/pic/cefc1e178a82b901d3ed95f0748da9773812efb5?x-bce-process=image/resize,m_lfit,w_220,h_220,limit_1', '2020-07-24 09:41:00');
INSERT INTO `article` VALUES ('4', null, null, 'python', '1', 'Python是一种跨平台的计算机程序设计语言。 是一个高层次的结合了解释性、编译性、互动性和面向对象的脚本语言。最初被设计用于编写自动化脚本(shell)，随着版本的不断更新和语言新功能的添加，越多被用于独立的、大型项目的开发。', '10', 'https://bkimg.cdn.bcebos.com/pic/35a85edf8db1cb1316f5387ed254564e92584b3a?x-bce-process=image/resize,m_lfit,w_268,limit_1/format,f_jpg', '2020-07-31 09:42:11');
INSERT INTO `article` VALUES ('5', null, null, '程序员的自我修养', '1', 'Elasticsearch是一个基于Lucene的搜索服务器。它提供了一个分布式多用户能力的全文搜索引擎，基于RESTful web接口。Elasticsearch是用Java语言开发的，并作为Apache许可条款下的开放源码发布，是一种流行的企业级搜索引擎。Elasticsearch用于云计算中，能够达到实时搜索，稳定，可靠，快速，安装使用方便。官方客户端在Java、.NET（C#）、PHP、Python、Apache Groovy、Ruby和许多其他语言中都是可用的。根据DB-Engines的排名显示，Elasticsearch是最受欢迎的企业搜索引擎，其次是Apache Solr，也是基于Lucene。\r\n', '10', null, null);
INSERT INTO `article` VALUES ('6', null, null, 'elasticsearch', '1', 'Elasticsearch是一个基于Lucene的搜索服务器。它提供了一个分布式多用户能力的全文搜索引擎，基于RESTful web接口。Elasticsearch是用Java语言开发的，并作为Apache许可条款下的开放源码发布，是一种流行的企业级搜索引擎。Elasticsearch用于云计算中，能够达到实时搜索，稳定，可靠，快速，安装使用方便。官方客户端在Java、.NET（C#）、PHP、Python、Apache Groovy、Ruby和许多其他语言中都是可用的。根据DB-Engines的排名显示，Elasticsearch是最受欢迎的企业搜索引擎，其次是Apache Solr，也是基于Lucene。\r\n', '10', null, null);
INSERT INTO `article` VALUES ('7', null, null, 'mysql', '1', '慢日志', '10', null, null);

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
) ENGINE=MyISAM AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COMMENT='配置表';

-- ----------------------------
-- Records of config
-- ----------------------------
INSERT INTO `config` VALUES ('1', 'host', 'http://www.tfuu.cn');
INSERT INTO `config` VALUES ('2', 'home_banner', 'https://images.tfuu.cn/home_banner.jpg');
INSERT INTO `config` VALUES ('3', 'banner', 'https://images.tfuu.cn/banner.jpg');

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
