-- migrations/001_create_users.sql
-- Jalankan manual di MySQL jika tidak pakai AutoMigrate

CREATE TABLE IF NOT EXISTS `users` (
  `id`         BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `name`       VARCHAR(100)    NOT NULL,
  `email`      VARCHAR(100)    NOT NULL,
  `phone`      VARCHAR(20)     DEFAULT '',
  `created_at` DATETIME(3)     DEFAULT NULL,
  `updated_at` DATETIME(3)     DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_users_email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;