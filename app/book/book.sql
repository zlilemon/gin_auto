## 创建db
create database wxapp;

##
use wxapp;
CREATE TABLE `account_balance` (
                                   `id` int(20) unsigned NOT NULL AUTO_INCREMENT,
                                   `openid` varchar(32) NOT NULL DEFAULT '' COMMENT '用户对应的openid',
                                   `currency` varchar(64) NOT NULL DEFAULT 'CNY' COMMENT '币种，默认 CNY',
                                   `amount` int 	NOT NULL DEFAULT 0 COMMENT '账户余额，单位：分',
                                   `jf_amount` int NOT NULL DEFAULT 0 COMMENT '积分账户余额',
                                   `status` varchar(32) NOT NULL DEFAULT '' COMMENT '账户状态, VALID:有效, INVALID:无效',
                                   `version_id` int NOT NULL DEFAULT 0 COMMENT '版本号，每次更新+1',
                                   `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                   `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间' ,
                                   PRIMARY KEY (`id`),
                                   KEY `idx_oid` (`openid`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4


##
use store;
CREATE TABLE `order_detail` (
    `id` int(20) unsigned NOT NULL AUTO_INCREMENT,
    `out_trade_no` varchar(64) NOT NULL DEFAULT '' COMMENT '业务订单号',
    `channel_order_no` varchar(64) NOT NULL DEFAULT '' COMMENT '渠道订单号',
    `order_type` varchar(64) NOT NULL DEFAULT '' COMMENT '订单类型，accountSave：充值账户，bookConsume：预定位置支付',
    `pay_method` varchar(64) NOT NULL DEFAULT '' COMMENT '支付方式，微信支付：wechat',
    `openid` varchar(32) NOT NULL DEFAULT '' COMMENT '预定用户对应的openid',
    `store_id` varchar(32) NOT NULL DEFAULT '' COMMENT '店铺id',
    `store_name` varchar(512) NOT NULL DEFAULT '' COMMENT '店铺名场',
    `seat_id` varchar(128) NOT NULL DEFAULT '' COMMENT '座位id',
    `seat_name` varchar(512) NOT NULL DEFAULT '' COMMENT '座位名称',
    `seat_type` varchar(128) NOT NULL DEFAULT '' COMMENT '座位类型, normal/vip/double 等',
    `book_begin_date` varchar(10) NOT NULL DEFAULT '' COMMENT '预定开始日期, eg: 2022-10-01',
    `book_end_date` varchar(10) NOT NULL DEFAULT '' COMMENT '预定结束日期, eg: 2022-10-01',
    `book_begin_time` varchar(20) NOT NULL DEFAULT '' COMMENT '预定开始时间, eg: 10:00',
    `book_end_time` varchar(20) NOT NULL DEFAULT '' COMMENT '预定结束时间, eg: 11:15',
    `book_duration` varchar(64) NOT NULL DEFAULT '' COMMENT '预定时长，eg：4.5小时',
    `begin_unix_time` int(11)  NOT NULL DEFAULT 0 COMMENT '预定开始时间 - unix_time格式',
    `end_unix_time` int(11)  NOT NULL DEFAULT 0 COMMENT '预定结束时间 - unix_time格式',
    `currency` varchar(32) NOT NULL DEFAULT '' COMMENT '币种，默认：CNY',
    `amount` int(11) NOT NULL DEFAULT '0' COMMENT '预定金额，单位：分',
    `tran_time` int(11) NOT NULL DEFAULT '0' COMMENT '交易时间戳',
    `pay_info` varchar(256) NOT NULL DEFAULT '' COMMENT '交易备注',
    `remark` varchar(256) NOT NULL DEFAULT '' COMMENT '备注',
    `bill_status` varchar(64) NOT NULL DEFAULT '',
    `pay_status` varchar(64) NOT NULL DEFAULT '',
    `notify_status` varchar(64) NOT NULL DEFAULT '',
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY `idx_oid` (`out_trade_no`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb4

use store;
alter table order_detail add begin_unix_time int(11) NOT NULL DEFAULT 0 after book_duration;
alter table order_detail add end_unix_time int(11) NOT NULL DEFAULT 0 after begin_unix_time;

update order_detail set begin_unix_time=1678120200, end_unix_time=1678127400 where out_trade_no=1678091496 limit 1;