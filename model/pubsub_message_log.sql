CREATE TABLE `pubsub_message_log` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `pubsub_id` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `data` text COLLATE utf8_unicode_ci NOT NULL,
  `attribute` text COLLATE utf8_unicode_ci,
  `error_process` text COLLATE utf8_unicode_ci,
  `topic` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `receive_time` timestamp NULL DEFAULT NULL,
  `publish_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `source` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `subscription_id` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9331 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;