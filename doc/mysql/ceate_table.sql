CREATE TABLE `t_bases` (
  `id` bigint(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` DATETIME NOT NULL COMMENT '时间戳',
  `updated_at` DATETIME NOT NULL COMMENT '时间戳',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

CREATE TABLE `t_users` (
  `id` bigint(10) unsigned NOT NULL AUTO_INCREMENT,
  `account` varchar(128) NOT NULL UNIQUE,
  `password` varchar(128) NOT NULL,
  `nickname` varchar(128) NOT NULL UNIQUE,
  `created_at` DATETIME NOT NULL COMMENT '时间戳',
  `updated_at` DATETIME NOT NULL COMMENT '时间戳',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
