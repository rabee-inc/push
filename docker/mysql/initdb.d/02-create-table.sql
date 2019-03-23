CREATE TABLE IF NOT EXISTS `push`.`tokens` (
  `id`         varchar(32)         NOT NULL,
  `user_id`    varchar(255)        NOT NULL,
  `platform`   varchar(20)         NOT NULL,
  `device_id`  varchar(255)        NOT NULL DEFAULT '',
  `token`      varchar(1024)       NOT NULL,
  `created_at` bigint(20) unsigned NOT NULL,
  `updated_at` bigint(20) unsigned NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY idx_user_id_and_platform_and_device_id (`user_id`, `platform`, `device_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
COMMENT 'プッシュ通知のトークン';

CREATE TABLE IF NOT EXISTS `push`.`reserves` (
  `id`          varchar(32)         NOT NULL,
  `title`       varchar(255)        NOT NULL,
  `body`        varchar(255)        NOT NULL,
  `data`        json                NOT NULL,
  `status` enum('ready', 'started', 'finished', 'revoked') NOT NULL DEFAULT 'ready',
  `reserved_at` bigint(20) unsigned NOT NULL,
  `created_at`  bigint(20) unsigned NOT NULL,
  `updated_at`  bigint(20) unsigned NOT NULL,
  PRIMARY KEY (`id`),
  KEY idx_status (`status`),
  KEY idx_reserved_at  (`reserved_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
COMMENT 'プッシュ通知の予約';

