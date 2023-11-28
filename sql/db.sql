CREATE TABLE IF NOT EXISTS `tb_user_info` (
    `uid` int unsigned NOT NULL AUTO_INCREMENT,
    `username` varchar(20) NOT NULL COMMENT '用户名',
    `status` tinyint unsigned NOT NULL DEFAULT 0 COMMENT '用户账号状态',
    `ctime` timestamp NOT NULL DEFAULT '2023-11-25 14:38:09',
    PRIMARY KEY (`uid`),
    UNIQUE KEY (`username`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT '用户信息表';

CREATE TABLE IF NOT EXISTS `tb_role` (
    `rid` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'roleId',
    `uid` int unsigned NOT NULL COMMENT '用户UID',
    `sex` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '性别，0:女 1男',
    `name` varchar(100) COMMENT '玩家名',
    `login_time` TIMESTAMP NULL DEFAULT NULL COMMENT '登录时间',
    `logout_time` TIMESTAMP NULL DEFAULT NULL COMMENT '登出时间',
    `created_time` timestamp NOT NULL DEFAULT '2023-11-25 14:38:09',
    UNIQUE KEY (`uid`),
    PRIMARY KEY (`rid`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT '玩家表';
