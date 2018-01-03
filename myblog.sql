/*
Navicat MySQL Data Transfer

Source Server         : MyTest
Source Server Version : 50717
Source Host           : localhost:3306
Source Database       : myblog

Target Server Type    : MYSQL
Target Server Version : 50717
File Encoding         : 65001

Date: 2018-01-03 14:05:52
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `banner`
-- ----------------------------
DROP TABLE IF EXISTS `banner`;
CREATE TABLE `banner` (
  `ID` bigint(20) NOT NULL,
  `IMGURL` varchar(50) DEFAULT NULL,
  `TITLE` varchar(500) DEFAULT NULL,
  `SUBTITLE` varchar(500) DEFAULT NULL,
  `URL` varchar(500) DEFAULT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of banner
-- ----------------------------
INSERT INTO `banner` VALUES ('1', './static/images/banner1.jpg', 'My Blog Based on Golang and Bootstrap1', 'My Blog Based on Golang and BootstrapMy Blog Based on Golang and BootstrapMy Blog Based on Golang and Bootstrap', 'http://www.baidu.com');
INSERT INTO `banner` VALUES ('2', './static/images/banner2.jpg', 'My Blog Based on Golang and Bootstrap2', 'My Blog Based on Golang and Bootstrap2My Blog Based on Golang and Bootstrap2My Blog Based on Golang and Bootstrap2', 'http://www.baidu.com');
INSERT INTO `banner` VALUES ('3', './static/images/banner3.jpg', 'My Blog Based on Golang and Bootstrap3', 'My Blog Based on Golang and Bootstrap3My Blog Based on Golang and Bootstrap3My Blog Based on Golang and Bootstrap3', 'http://www.baidu.com');
INSERT INTO `banner` VALUES ('4', 'static/upload/1514251430405.jpg', 'My Blog Based on Golang and Bootstrap4', 'My Blog Based on Golang and Bootstrap4My Blog Based on Golang and Bootstrap4My Blog Based on Golang and Bootstrap4', 'http://www.baidu.com');

-- ----------------------------
-- Table structure for `blog`
-- ----------------------------
DROP TABLE IF EXISTS `blog`;
CREATE TABLE `blog` (
  `ID` bigint(20) NOT NULL AUTO_INCREMENT,
  `AUTH` varchar(50) DEFAULT NULL,
  `TITLE` varchar(100) DEFAULT NULL,
  `KEYWORDS` varchar(100) DEFAULT NULL,
  `CATALOGID` varchar(20) DEFAULT NULL,
  `INTRODUCTION` longtext,
  `IMGURL` varchar(150) DEFAULT NULL,
  `CONTENT` longtext,
  `LASTUPDATE` datetime DEFAULT NULL,
  `TYPE` varchar(20) DEFAULT NULL,
  `STATUS` varchar(20) DEFAULT NULL,
  `VIEWS` varchar(20) DEFAULT NULL,
  `CREATETIME` datetime DEFAULT NULL,
  `SUBJECT` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=31 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of blog
-- ----------------------------
INSERT INTO `blog` VALUES ('13', null, '111111', '11111', '1', '11111', 'static/upload/16019175309d91516fe3e5e39eb171cf.png', '111111', '2017-12-27 14:41:37', 'original', 'private', '0', '2017-12-27 14:41:37', '11111');
INSERT INTO `blog` VALUES ('14', null, '111111', '1111', '1', '11111', 'static/upload/150752846099.jpg', '11111', '2017-12-27 14:41:51', 'original', '1', '0', '2017-12-27 14:41:51', '1111');
INSERT INTO `blog` VALUES ('15', null, '111111', '1111', '1', '11111', 'static/upload/150752846099.jpg', '11111', '2017-12-27 14:43:34', 'original', '1', '0', '2017-12-27 14:43:34', '1111');
INSERT INTO `blog` VALUES ('16', null, '111111', '1111', '1', '11111', 'static/upload/150752846099.jpg', '11111', '2017-12-27 14:43:37', 'original', '1', '0', '2017-12-27 14:43:37', '1111');
INSERT INTO `blog` VALUES ('17', null, '222222', '222', '1', '222222', 'static/upload/1514170960819.jpg', '2222', '2017-12-27 14:44:23', 'original', 'private', '0', '2017-12-27 14:44:23', '22222');
INSERT INTO `blog` VALUES ('18', null, '3333333333', '33333', '1', '333333333333', 'static/upload/1514169731237.jpg', '3333333333333', '2017-12-27 14:47:12', 'original', 'private', '0', '2017-12-27 14:47:12', '3333333333');
INSERT INTO `blog` VALUES ('19', null, '55555555555555', '555555555', '1', '555555555555', 'static/upload/1514251430405.jpg', '555555555555555555', '2017-12-27 14:52:34', 'original', 'private', '0', '2017-12-27 14:52:34', '555555555');
INSERT INTO `blog` VALUES ('20', null, '1112344', '11111', '1', '1111111', 'static/upload/150752846099.jpg', '111111111', '2017-12-27 14:55:26', 'original', 'private', '0', '2017-12-27 14:55:26', '11111');
INSERT INTO `blog` VALUES ('21', null, '55677888', '123', '2', 'qweqweqwe', 'static/upload/1514251458556.jpg', 'wqeqweqew', '2017-12-27 15:01:08', 'original', 'private', '0', '2017-12-27 15:01:08', '123123');
INSERT INTO `blog` VALUES ('22', null, '5678790---', 'dsfsdf', 'work', 'sdfsfsf', 'static/upload/150752846099.jpg', 'sdfsdfsdf', '2017-12-28 15:12:48', 'original', 'private', '', '2017-12-28 15:12:48', 'sdfsfds');
INSERT INTO `blog` VALUES ('23', null, 'werwer', 'werwe', 'work', 'werwerw', 'static/upload/150752846099.jpg', 'werwerw', '2017-12-28 15:14:08', 'original', 'private', '', '2017-12-28 15:14:08', 'werw');
INSERT INTO `blog` VALUES ('24', null, 'ewrwer', 'werw', 'work', 'werwerwer', 'static/upload/1507528466749.jpg', 'werwerw', '2017-12-28 15:14:20', 'original', 'private', '', '2017-12-28 15:14:20', 'werwer');
INSERT INTO `blog` VALUES ('25', null, 'test', 'kkkkkkkkkkkk', '', 'aaaaaaaaaa', 'static/upload/1514251449370.jpg', 'aaaaaaaaaawwqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq', '2017-12-29 15:45:07', 'original', '', '', null, 'kkkkkkkkkkkkkkkk');
INSERT INTO `blog` VALUES ('26', null, 'test', 'test', '', '1111111111111111', 'static/upload/16019175309d91516fe3e5e39eb171cf.png', '1111111111111111wwwwwwwwwwwwwwwwww<img src=\"http://localhost:8080/static/upload/16019175309d91516fe3e5e39eb171cf.png\" class=\"kv-preview-data file-preview-image\" title=\"\" alt=\"\" style=\"width:auto;height:160px;\">', '2017-12-29 15:46:34', 'original', '', '', null, 'test');
INSERT INTO `blog` VALUES ('27', '34534543', '123142354543', '345345', '', '3453534534543', 'static/upload/150752846099.jpg', '<p>3453534</p><p>34535345</p>', '2018-01-02 09:27:01', 'original', '', '', '2018-01-02 09:27:01', '345345');
INSERT INTO `blog` VALUES ('28', 'test10', 'test10', 'qwqe', '', 'qweqweqweqe', 'static/upload/1514169731237.jpg', 'qweqewqeqweqweqw', '2018-01-02 09:32:28', 'original', '', '', '2018-01-02 09:32:28', 'qweqwe');
INSERT INTO `blog` VALUES ('29', '12312313', 'test1000', '123213123', 'work', '123131312', 'static/upload/1514251046475.jpg', '12313123213', '2018-01-02 09:34:51', 'original', 'public', '', '2018-01-02 09:34:51', '12312313');
INSERT INTO `blog` VALUES ('30', 'asdasdasd', 'qwetwerterty', 'asdad', 'work', 'asdadasdsad', 'static/upload/1514251458556.jpg', 'asdadasdasd', '2018-01-02 09:37:01', 'original', 'public', '', '2018-01-02 09:37:01', 'asdasd');

-- ----------------------------
-- Table structure for `user`
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `ID` bigint(20) NOT NULL AUTO_INCREMENT,
  `USERNAME` varchar(50) DEFAULT NULL,
  `PASSWORD` varchar(50) DEFAULT NULL,
  `SEX` bigint(10) DEFAULT NULL,
  `AGE` bigint(10) DEFAULT NULL,
  `EMAIL` varchar(20) DEFAULT NULL,
  `IMGURL` varchar(50) DEFAULT NULL,
  `STATUS` bigint(10) DEFAULT NULL,
  `PERSONALKEY` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES ('1', 'yanfeilong', '123456', '1', '111', '11111@email.com', '11111', '1', '123456');
