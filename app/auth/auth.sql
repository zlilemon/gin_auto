CREATE TABLE `access_token` (
        `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '唯一标识',
        `access_token` varchar(1024) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'access_token',
        `expires_in` int(11) NOT NULL COMMENT '过期时间，单位(s)',
        `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
        `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间' ,
        PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci