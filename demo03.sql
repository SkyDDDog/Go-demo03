/*
SQLyog Ultimate v12.09 (64 bit)
MySQL - 8.0.27 : Database - demo03
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`demo03` /*!40100 DEFAULT CHARACTER SET utf8 */ /*!80016 DEFAULT ENCRYPTION='N' */;

USE `demo03`;

/*Table structure for table `note` */

DROP TABLE IF EXISTS `note`;

CREATE TABLE `note` (
  `id` varchar(255) NOT NULL COMMENT '主键',
  `del_flag` char(1) DEFAULT NULL COMMENT '逻辑删除',
  `create_date` datetime DEFAULT NULL COMMENT '创建时间',
  `update_date` datetime DEFAULT NULL COMMENT '更新时间',
  `remarks` varchar(255) DEFAULT NULL COMMENT '备注',
  `title` varchar(255) DEFAULT NULL COMMENT '标题',
  `content` varchar(255) DEFAULT NULL COMMENT '内容',
  `state` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '完成状态(0未完成)',
  `end_time` datetime DEFAULT NULL COMMENT '截止时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

/*Data for the table `note` */

insert  into `note`(`id`,`del_flag`,`create_date`,`update_date`,`remarks`,`title`,`content`,`state`,`end_time`) values ('1607321970309861376','0','2022-12-26 18:26:35','2022-12-26 18:26:35','','4444','内容','0','2022-12-27 02:26:34');
insert  into `note`(`id`,`del_flag`,`create_date`,`update_date`,`remarks`,`title`,`content`,`state`,`end_time`) values ('1607321983190568960','0','2022-12-26 18:26:38','2022-12-26 18:26:38','','55555','内容','0','2022-12-27 02:26:38');
insert  into `note`(`id`,`del_flag`,`create_date`,`update_date`,`remarks`,`title`,`content`,`state`,`end_time`) values ('1607321995068837888','0','2022-12-26 18:26:41','2022-12-26 18:26:41','','666','内容','0','2022-12-27 02:26:40');

/*Table structure for table `user` */

DROP TABLE IF EXISTS `user`;

CREATE TABLE `user` (
  `id` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '主键',
  `del_flag` char(1) DEFAULT NULL COMMENT '逻辑删除',
  `create_date` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `update_date` datetime(3) DEFAULT NULL COMMENT '更新时间',
  `remarks` varchar(255) DEFAULT NULL COMMENT '备注',
  `username` varchar(255) DEFAULT NULL COMMENT '用户名',
  `password` varchar(255) DEFAULT NULL COMMENT '密码',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

/*Data for the table `user` */

insert  into `user`(`id`,`del_flag`,`create_date`,`update_date`,`remarks`,`username`,`password`) values ('1607002749315059712','0','2022-12-25 21:18:06.818','2022-12-25 21:18:06.818','','test','$2a$10$IB8s9pnQMg15CaDkVIjnhO8F9tnbPv9mA6/JkJUqcXOxUejDWpCy.');

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
