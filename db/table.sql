
-- CREATE DATABASE IF NOT EXISTS `todo` DEFAULT CHARSET utf8 COLLATE utf8_general_ci;
-- use `todo`;
DROP TABLE IF EXISTS `cfg_user`;
CREATE TABLE `cfg_user`(
    `user_id` INT(11) AUTO_INCREMENT,
    `user_name` VARCHAR(128) NOT NULL DEFAULT '',
    `pwd` VARCHAR(128) NOT NULL DEFAULT '',
    `modified` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `created` TIMESTAMP NOT NULL DEFAULT '0000-00-00',
    PRIMARY KEY(`user_id`)
)Engine='InnoDB' CHARSET='utf8';

DROP TABLE IF EXISTS `to_do_list`;
CREATE TABLE `to_do_list`(
    `rec_id` INT(11) AUTO_INCREMENT,
    `status` INT(11) NOT NULL DEFAULT '0' COMMENT '0 待完成  5 已放弃 10 已完成 ',
    `expect_end_date` DATE NOT NULL DEFAULT '0000-00-00' COMMENT '预计截止时间',
    `end_date` DATE NOT NULL DEFAULT '0000-00-00' COMMENT '实际截止时间',
     `to_do_things` VARCHAR(1024) NOT NULL DEFAULT '',
    `user_id` INT(11) NOT NULL DEFAULT '0',
    `modified` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `created` TIMESTAMP NOT NULL DEFAULT '0000-00-00',
    PRIMARY KEY(`rec_id`),
    KEY `do_user_id`(`user_id`)
)Engine='InnoDB' CHARSET='utf8';

-- 日志表
DROP TABLE IF EXISTS `to_do_log`;
CREATE TABLE  `to_do_log`(
    `rec_id` INT(11) AUTO_INCREMENT,
    `do_id` INT(11) NOT NULL DEFAULT '0',
    `operator` INT(11) NOT NULL DEFAULT '0' COMMENT '操作人',
    `remark` VARCHAR(1024) NOT NULL DEFAULT '',
    `modified` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `created` TIMESTAMP NOT NULL DEFAULT '0000-00-00',
    PRIMARY KEY(`rec_id`),
    KEY `do_id`(`do_id`)
)Engine='InnoDb' CHARSET='utf8';




