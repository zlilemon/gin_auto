use store;
CREATE TABLE `device_mapping` (
        `id` int(20) unsigned NOT NULL AUTO_INCREMENT,
        `store_id` varchar(32) NOT NULL DEFAULT '' COMMENT '店铺id',
        `seat_id` varchar(512) NOT NULL DEFAULT '' COMMENT '座位id',
        `device_function_type` varchar(32) NOT NULL DEFAULT '' COMMENT '设备作用分类，entrance_door：入户门， switch：开关',
        `device_id` varchar(64) NOT NULL DEFAULT '' COMMENT '硬件设备在厂家侧对应的设备id',
        `device_name` varchar(64) NOT NULL DEFAULT '' COMMENT '硬件设备在厂家侧对应的设备名称',
        `device_category` varchar(64) NOT NULL DEFAULT '' COMMENT '硬件设备在厂家侧对应的设备分类，开关、电灯、led等',
        `device_brand` varchar(64) NOT NULL DEFAULT '' COMMENT '硬件设备品牌',
        `status` varchar(128) NOT NULL DEFAULT '' COMMENT '状态，VALID：生效， INVALID：失效',
        `remark` varchar(256) NOT NULL DEFAULT '' COMMENT '备注',
        `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
        `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间' ,
        PRIMARY KEY (`id`),
        KEY `idx_oid` (`store_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4

insert into device_mapping set store_id='s_001', seat_id='seat_002', device_function_type='entrance_door', device_id='', device_name='', device_category='', device_brand='tuya', status='VALID';
insert into device_mapping set store_id='s_001', seat_id='seat_002', device_function_type='switch', device_id='switch.zimi_zncz01_b007_switch', device_name='小米开关', device_category='switch', device_brand='xiaomi', status='VALID';