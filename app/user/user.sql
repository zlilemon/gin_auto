## 创建账户db
create database wxapp;

## 创建 金额&账户点数 对应关系表
use wxapp;
CREATE TABLE `clock_info` (
        `id` int(20) unsigned NOT NULL AUTO_INCREMENT,
        `openid` varchar(32) NOT NULL DEFAULT '' COMMENT '用户对应的openid',
        `out_trade_no` varchar(64) NOT NULL DEFAULT '' COMMENT '充值内部订单号',
        `static_date` varchar(10) NOT NULL DEFAULT '' COMMENT '打卡日期，YYYY-MM-DD',
        `duration_time` varchar(64) NOT NULL DEFAULT '' COMMENT '预定时长，eg：4.5小时',
        `status` varchar(32) NOT NULL DEFAULT '' COMMENT '账户状态, valid:有效, invalid:无效',
        `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
        `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间' ,
        PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4

CREATE TABLE `ranking_info` (
      `id` int(20) unsigned NOT NULL AUTO_INCREMENT,
      `batch_id` varchar(64) NOT NULL DEFAULT '' COMMENT '批次号，每备份一次更新一下批次号',
      `openid` varchar(32) NOT NULL DEFAULT '' COMMENT '用户对应的openid',
      `avatar_url` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '用户头像',
      `nick_name` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '用户昵称',
      `weekly_times` int NOT NULL DEFAULT '0' COMMENT '本周次数，本周从周一开始',
      `weekly_duration_time` int NOT NULL DEFAULT '0' COMMENT '本周预定时长，单位：小时，本周从周一开始',
      `monthly_times` int NOT NULL DEFAULT '0' COMMENT '本月次数，本月从1号开始',
      `monthly_duration_time` int NOT NULL DEFAULT '0' COMMENT '本月预定时长，单位：小时，本月从1号开始',
      `totally_times` int NOT NULL DEFAULT '0' COMMENT '累计次数',
      `totally_duration_time` int NOT NULL DEFAULT '0' COMMENT '累计预定时长，单位：小时',
      `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
      `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间' ,
      PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4