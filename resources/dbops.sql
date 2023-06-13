create table sys_version(
    `id` int unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(20) NOT NULL COMMENT '版本号',
    PRIMARY KEY (`id`)
);
insert into sys_version values(1,'1.0.0');

CREATE TABLE `sys_user` (
                            `id` int(50) unsigned NOT NULL AUTO_INCREMENT,
                            `user_name` varchar(50) NOT NULL COMMENT '用户名',
                            `nick_name` varchar(150) DEFAULT NULL COMMENT '昵称',
                            `avatar` varchar(150) DEFAULT NULL COMMENT '头像',
                            `password` varchar(100) DEFAULT NULL COMMENT '密码',
                            `email` varchar(100) DEFAULT NULL COMMENT '邮箱',
                            `mobile` varchar(100) DEFAULT NULL COMMENT '手机号',
                            `create_time` bigint(50) DEFAULT NULL COMMENT '更新时间',
                            `del_status` tinyint(4) DEFAULT '0' COMMENT '是否删除 -1：已删除   0：正常',
                            PRIMARY KEY (`id`),
                            UNIQUE KEY `name` (`user_name`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='用户表';

-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO `sys_user` VALUES ('1', 'admin', '', '', '123456', '924931408@qq.com', '15751004355', '0', '0');