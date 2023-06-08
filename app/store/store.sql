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
        `currency` varchar(32) 	NOT NULL DEFAULT '' COMMENT '币种，默认：CNY',
        `amount` int 	NOT NULL DEFAULT 0 COMMENT '预定金额，单位：分',
        `tran_time` int 	NOT NULL DEFAULT 0 COMMENT '交易时间戳',
        `pay_info` varchar(256) NOT NULL DEFAULT '' COMMENT '交易备注',
        `remark` varchar(256) NOT NULL DEFAULT '' COMMENT '备注',
        `bill_status` varchar(64) NOT NULL DEFAULT '' COMMENT '',
        `pay_status` varchar(64) NOT NULL DEFAULT '' COMMENT '',
        `notify_status` varchar(64) NOT NULL DEFAULT '' COMMENT '',
        `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
        `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间' ,
        PRIMARY KEY (`id`),
        KEY `idx_oid` (`out_trade_no`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4


CREATE TABLE `store_info` (
  `id` int(20) unsigned NOT NULL AUTO_INCREMENT,
  `store_id` varchar(32) NOT NULL DEFAULT '' COMMENT '商店id',
  `store_name` varchar(512) NOT NULL DEFAULT '' COMMENT '商店名场',
  `province` varchar(128) NOT NULL DEFAULT '' COMMENT '商店所在省份',
  `city` varchar(128) NOT NULL DEFAULT '' COMMENT '商店所在城市',
  `address` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '店铺详细地址',
  `phone_no` varchar(64) NOT NULL DEFAULT '' COMMENT '联系电话',
  `introduction` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '店铺介绍',
  `longitude` double(16,2) NOT NULL DEFAULT '0.00' COMMENT '经度',
  `latitude` double(16,2) NOT NULL DEFAULT '0.00' COMMENT '纬度',
  `home_pic_url` varchar(1024) NOT NULL DEFAULT '' COMMENT '主页图片地址',
  `detail_pic_url` text COMMENT '明细页图片地址,多个地址用英文逗号分隔',
  `seat_pic_url` varchar(1024) NOT NULL DEFAULT '' COMMENT '座位图片地址',
  `introduction_pic_url` varchar(1024) NOT NULL DEFAULT '' COMMENT '介绍图片地址',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_oid` (`store_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4

CREATE TABLE `price` (
  `id` int(20) unsigned NOT NULL AUTO_INCREMENT,
  `store_id` varchar(32) NOT NULL DEFAULT '' COMMENT '商店id',
  `store_name` varchar(512) NOT NULL DEFAULT '' COMMENT '商店名场',
  `seat_type` varchar(128) NOT NULL DEFAULT '' COMMENT '座位类型, normal/vip/double 等',
  `price_type` varchar(128) NOT NULL DEFAULT '' COMMENT '价格分类：minute：按分钟价格; day: 按天价格',
  `ori_price` int NOT NULL DEFAULT 0 COMMENT '原价，单位：分',
  `real_price` int NOT NULL DEFAULT 0 COMMENT '实际价格，单位：分',
  `status` varchar(32) NOT NULL DEFAULT '' COMMENT '账户状态, valid:有效, invalid:无效',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_oid` (`store_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4

insert into price set store_id='s_001', store_name='测试店铺1', seat_type='normal', price_type='minute', ori_price=300, real_price=240;
insert into price set store_id='s_001', store_name='测试店铺1', seat_type='normal', price_type='day', ori_price=1800, real_price=1440;
insert into price set store_id='s_001', store_name='测试店铺1', seat_type='normal', price_type='month', ori_price=45000, real_price=36000;

insert into price set store_id='s_001', store_name='测试店铺1', seat_type='vip', price_type='minute', ori_price=300, real_price=240;
insert into price set store_id='s_001', store_name='测试店铺1', seat_type='vip', price_type='day', ori_price=1800, real_price=1440;

CREATE TABLE `store_notice` (
     `id` int(20) unsigned NOT NULL AUTO_INCREMENT,
     `wifi_name` varchar(64) NOT NULL DEFAULT '' COMMENT 'wifi名称',
     `wifi_passwd` varchar(64) NOT NULL DEFAULT '' COMMENT 'wifi密码',
     `customer_phone` varchar(64) NOT NULL DEFAULT '' COMMENT '客服电话',
     `status` varchar(32) NOT NULL DEFAULT '' COMMENT '账户状态, valid:有效, invalid:无效',
     `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
     `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
     PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4

insert into store_notice set wifi_name = 'test_wifi', wifi_passwd = 'test_wifi_passwd', customer_phone = '13800138000', status='valid';

CREATE TABLE `store_question` (
    `id` int(20) unsigned NOT NULL AUTO_INCREMENT,
    `ask` varchar(2048) NOT NULL DEFAULT '' COMMENT '问题',
    `answer` text NOT NULL  COMMENT '答案',
    `status` varchar(32) NOT NULL DEFAULT '' COMMENT '账户状态, valid:有效, invalid:无效',
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4

insert into store_question set ask = '问题1：你知道今天星期几吗？', answer = '答复：今天应该是星期二吧', status = 'valid';
insert into store_question set ask = '问题2：你知道今天星期几吗？', answer = '答复：今天应该是星期二吧', status = 'valid';

CREATE TABLE `bulletin` (
    `id` int(64) unsigned NOT NULL AUTO_INCREMENT,
    `title` varchar(2048) NOT NULL DEFAULT '' COMMENT '公告title',
    -- `publish_time` varchar(20) NOT NULL DEFAULT '' COMMENT '公告发布时间，YYYY-MM-DD HH:MM:SS 格式',
    `publish_time_unix_time` int(64) NOT NULL DEFAULT '0' COMMENT '公告发布时间，unix_time 格式',
    `publish_detail` text NOT NULL COMMENT '公告详细内容',
    `status` varchar(32) NOT NULL DEFAULT '' COMMENT '账户状态, VALID:有效, INVALID:无效',
    `create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4

insert into bulletin set title='公告1：五一假期的公告', publish_time_unix_time='1683648000', publish_detail='五一假期的公告详细内容\n五一假期的公告详细内容\n', status='VALID';
insert into bulletin set title='公告1：六一假期的公告', publish_time_unix_time='1686326400', publish_detail='六一儿童节假期的公告详细内容\n六一儿童节假期的公告详细内容\n', status='VALID';