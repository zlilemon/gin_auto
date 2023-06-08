## 创建账户db
create database account;

## 创建 金额&账户点数 对应关系表
use account;
CREATE TABLE `charge_point_mapping` (
       `id` int(20) unsigned NOT NULL AUTO_INCREMENT,
       `relation_id` varchar(64) NOT NULL DEFAULT '' COMMENT '映射关联id',
       `relation_name` varchar(128) NOT NULL DEFAULT '' COMMENT '映射关联名称，eg：充100得108',
       `charge_amount` int NOT NULL DEFAULT '0' COMMENT '现金充值金额，单位：分',
       `account_point` int NOT NULL DEFAULT '0' COMMENT '账户充值点数，单位：分',
       `status` varchar(32) NOT NULL DEFAULT '' COMMENT '账户状态, valid:有效, invalid:无效',
       `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
       `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间' ,
       PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4

insert into account.charge_point_mapping set relation_id='r_001', relation_name='充100到账108',charge_amount=10000,account_point=10800, status='valid';
insert into account.charge_point_mapping set relation_id='r_002', relation_name='充200到账220',charge_amount=20000,account_point=22000, status='valid';
insert into account.charge_point_mapping set relation_id='r_003', relation_name='充500到账565',charge_amount=50000,account_point=56500, status='valid';
insert into account.charge_point_mapping set relation_id='r_004', relation_name='充0.01到账0.01',charge_amount=1,account_point=1, status='valid';

## 创建账户余额表
use account;
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

## 创建账户流水表
CREATE TABLE `account_water` (
   `id` bigint NOT NULL AUTO_INCREMENT,
   `static_date` varchar(10) NOT NULL DEFAULT '' COMMENT '交易日期，格式 YYYY-MM-DD',
   `openid` varchar(32) NOT NULL DEFAULT '' COMMENT '用户对应的openid',
   `out_trade_no` varchar(64) NOT NULL DEFAULT '' COMMENT '充值内部订单号',
   `channel_order_no` varchar(64) NOT NULL DEFAULT '' COMMENT '外部订单号',
   `tran_type` varchar(64) NOT NULL DEFAULT '' COMMENT '交易类型，charge: 充值, consume：消耗, refund：退款',
   `sub_tran_type` varchar(64) NOT NULL DEFAULT '' COMMENT '子交易类型，先预留',
   `currency` varchar(64) NOT NULL DEFAULT 'CNY' COMMENT '币种，默认 CNY',
   `amount` int 	NOT NULL DEFAULT 0 COMMENT '交易金额，单位：分',
   `tran_time` timestamp NULL COMMENT '交易时间',
   `tran_time_unix_time` int NULL COMMENT '交易时间戳格式',
   `remark` varchar(2048) NOT NULL DEFAULT '' COMMENT '备注信息，预留',
   `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
   `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间' ,
   PRIMARY KEY (`id`),
   KEY `idx_oid` (`static_date`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4

## 创建积分账户余额表
use account;
CREATE TABLE `jf_account_balance` (
           `id` int(20) unsigned NOT NULL AUTO_INCREMENT,
           `openid` varchar(32) NOT NULL DEFAULT '' COMMENT '用户对应的openid',
           `currency` varchar(64) NOT NULL DEFAULT 'CNY' COMMENT '币种，在积分侧暂时没有意义，只为了兼容现金账户结构',
           `amount` int 	NOT NULL DEFAULT 0 COMMENT '积分余额，单位：分',
           `status` varchar(32) NOT NULL DEFAULT '' COMMENT '账户状态, VALID:有效, INVALID:无效',
           `version_id` int NOT NULL DEFAULT 0 COMMENT '版本号，每次更新+1',
           `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
           `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间' ,
           PRIMARY KEY (`id`),
           KEY `idx_oid` (`openid`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4

## 创建积分账户流水表
CREATE TABLE `jf_account_water` (
         `id` bigint NOT NULL AUTO_INCREMENT,
         `static_date` varchar(10) NOT NULL DEFAULT '' COMMENT '交易日期，格式 YYYY-MM-DD',
         `openid` varchar(32) NOT NULL DEFAULT '' COMMENT '用户对应的openid',
         `out_trade_no` varchar(64) NOT NULL DEFAULT '' COMMENT '充值内部订单号',
         `channel_order_no` varchar(64) NOT NULL DEFAULT '' COMMENT '外部订单号',
         `tran_type` varchar(64) NOT NULL DEFAULT '' COMMENT '交易类型，charge: 充值, consume：消耗, refund：退款',
         `sub_tran_type` varchar(64) NOT NULL DEFAULT '' COMMENT '子交易类型，先预留',
         `currency` varchar(64) NOT NULL DEFAULT 'CNY' COMMENT '币种，在积分侧暂时没有意义，只为了兼容现金账户结构',
         `amount` int 	NOT NULL DEFAULT 0 COMMENT '积分金额，单位：分',
         `tran_time` timestamp NULL COMMENT '交易时间',
         `tran_time_unix_time` int(64) NULL COMMENT '交易时间戳格式',
         `remark` varchar(2048) NOT NULL DEFAULT '' COMMENT '备注信息，预留',
         `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
         `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间' ,
         PRIMARY KEY (`id`),
         KEY `idx_oid` (`static_date`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4