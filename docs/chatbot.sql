CREATE TABLE IF NOT EXISTS `chat_records_0` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '聊天记录id',
    `type` TINYINT(2) UNSIGNED NOT NULL DEFAULT '0' COMMENT '聊天主体：0 机器人,1 用户',
    `user_id` INT(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '用户id',
    `bot_provider` TINYINT(2) UNSIGNED NOT NULL DEFAULT '0' COMMENT '机器人模型提供者',
    `bot_model` CHAR(50) NOT NULL DEFAULT '' COMMENT '模型名称',
    `msg` VARCHAR(5000) NOT NULL DEFAULT '' COMMENT '聊天信息',
    `created_at` BIGINT NOT NULL DEFAULT '0' COMMENT '发布时间',

    PRIMARY KEY (`id`),
    KEY `uid_ctime` (`user_id`,`created_at`),
    KEY `ctime` (`created_at`)
)
COMMENT='聊天记录表'
DEFAULT CHARSET=utf8mb4
ENGINE=InnoDB
AUTO_INCREMENT=10000
;
