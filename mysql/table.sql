CREATE TABLE `azur_lane` (
                             `createdAt` datetime(3) DEFAULT NULL,  --
                             `updatedAt` datetime(3) DEFAULT NULL, --
                             `id` bigint NOT NULL AUTO_INCREMENT, --
                             `createdById` bigint DEFAULT NULL, --
                             `updatedById` bigint DEFAULT NULL, --
                             `code` varchar(255) DEFAULT NULL, --wiki编码
                             `name` varchar(255) DEFAULT NULL, --舰船名称
                             `camp` varchar(255) DEFAULT NULL, --舰船阵营
                             `ship_type` varchar(255) DEFAULT NULL, --舰船类型
                             `tech_point_get` bigint DEFAULT '0', --获得舰船加的科技点
                             `tech_point_star` bigint DEFAULT '0', --舰船突破满星加的科技点
                             `tech_point_lv120` bigint DEFAULT '0', --120级后加的科技点
                             `tech_point_total` bigint DEFAULT NULL, --最多可以加的科技点
                             `attribute_get_apply_ship` varchar(255) DEFAULT NULL, --获得舰船后 属性作用的舰船类型 eg:战巡/战列/航战
                             `attribute_lv120_apply_ship` varchar(255) DEFAULT NULL, --舰船120级后 属性作用的舰船类型 eg:战巡/战列/航战
                             `attribute_name_get` varchar(255) DEFAULT NULL, --获取舰船 解锁的属性类型 eg:耐久
                             `attribute_name_lv120` varchar(255) DEFAULT NULL, --120级后 解锁的属性类型 eg:命中
                             `attribute_get` bigint DEFAULT '0', --获得舰船属性加成值
                             `attribute_lv120` bigint DEFAULT '0', --舰船120级后获得属性值
                             `is_get_tech` varchar(255) DEFAULT NULL, --是否已获取科技点标识
                             PRIMARY KEY (`id`),
                             UNIQUE KEY `code` (`code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci