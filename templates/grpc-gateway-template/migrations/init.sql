DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
  `id` BIGINT(20) PRIMARY KEY AUTO_INCREMENT,
  `name` VARCHAR(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `req_id` VARCHAR(128) COMMENT 'uinique_req_id' NOT NULL,
  `retry_time` INT COMMENT 'how many retry time after success' NOT NULL,
  `status` ENUM('INIT', 'SUCCESS', 'FAIL', 'NEED_TO_LEAN_UP'),
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  INDEX idx_req_id (`req_id`)
);
