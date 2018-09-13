USE beegoBBS;

DROP TABLE IF EXISTS `bbs_users`;
CREATE TABLE `bbs_users` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(100) DEFAULT NULL,
  `avatar` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `introduction` varchar(255) DEFAULT NULL,
  `remember_token` varchar(100) NOT NULL,
  `notification_count` int(11) DEFAULT NULL,
  `last_actived_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4

DROP TABLE IF EXISTS `bbs_categories`;
CREATE TABLE `bbs_categories` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '名称',
  `description` text COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '描述',
  `post_count` int(11) NOT NULL DEFAULT 0 COMMENT '帖子数',
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`),
  KEY `bbs_categories_name_index` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO `bbs_categories` VALUES ('1', '分享', '分享创造，分享发现', '0', null, null);
INSERT INTO `bbs_categories` VALUES ('2', '教程', '开发技巧、推荐扩展包等', '0', null, null);
INSERT INTO `bbs_categories` VALUES ('3', '问答', '请保持友善，互帮互助', '0', null, null);
INSERT INTO `bbs_categories` VALUES ('4', '公告', '站点公告', '0', null, null);


DROP TABLE IF EXISTS `bbs_topics`;
CREATE TABLE `bbs_topics` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `body` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `user_id` int(10) unsigned NOT NULL,
  `category_id` int(10) unsigned NOT NULL,
  `reply_count` int(10) unsigned NOT NULL DEFAULT 0,
  `view_count` int(10) unsigned NOT NULL DEFAULT 0,
  `last_reply_user_id` int(10) unsigned NOT NULL DEFAULT 0,
  `order` int(10) unsigned NOT NULL DEFAULT 0,
  `excerpt` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `slug` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `status` int(11) NOT NULL DEFAULT 1,
  PRIMARY KEY (`id`),
  KEY `bbs_topics_title_index` (`title`),
  KEY `bbs_topics_user_id_index` (`user_id`),
  KEY `bbs_topics_category_id_index` (`category_id`),
  CONSTRAINT `bbs_topics_user_id_foreign` FOREIGN KEY (`user_id`) REFERENCES `bbs_users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

DROP TABLE IF EXISTS `bbs_replies`;
CREATE TABLE `bbs_replies` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `topic_id` int(10) unsigned NOT NULL DEFAULT 0,
  `user_id` int(10) unsigned NOT NULL DEFAULT 0,
  `content` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `bbs_replies_topic_id_index` (`topic_id`),
  KEY `bbs_replies_user_id_index` (`user_id`),
  CONSTRAINT `bbs_replies_topic_id_foreign` FOREIGN KEY (`topic_id`) REFERENCES `bbs_topics` (`id`) ON DELETE CASCADE,
  CONSTRAINT `bbs_replies_user_id_foreign` FOREIGN KEY (`user_id`) REFERENCES `bbs_users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;



