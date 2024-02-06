-- decks: table
CREATE TABLE `decks`
(
    `id`           bigint unsigned NOT NULL AUTO_INCREMENT,
    `created_at`   datetime(3) DEFAULT NULL COMMENT '创建时间',
    `updated_at`   datetime(3) DEFAULT NULL COMMENT '更新时间',
    `deleted_at`   datetime(3) DEFAULT NULL,
    `name`         longtext COLLATE utf8mb4_general_ci,
    `cardnum`      bigint                                  DEFAULT NULL,
    `learnnumber`  bigint                                  DEFAULT NULL,
    `introduction` longtext COLLATE utf8mb4_general_ci,
    `floderid`     bigint                                  DEFAULT NULL,
    `user_id`      varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- floder: table
CREATE TABLE `floder`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT,
    `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
    `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
    `deleted_at` datetime(3) DEFAULT NULL,
    `name`       longtext COLLATE utf8mb4_general_ci COMMENT '文件夹的名字叫啥',
    `user_id`    varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '属于哪个用户',
    `decknum`    bigint                                  DEFAULT NULL COMMENT '有几个牌组',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- knowledge: table
CREATE TABLE `knowledge`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT,
    `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
    `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
    `deleted_at` datetime(3) DEFAULT NULL,
    `font`       longtext COLLATE utf8mb4_general_ci,
    `originfont` longtext COLLATE utf8mb4_general_ci,
    `back`       longtext COLLATE utf8mb4_general_ci,
    `onlearning` tinyint(1) DEFAULT NULL,
    `typeof`     bigint                                  DEFAULT NULL,
    `deckid`     bigint unsigned DEFAULT NULL,
    `skilled` double DEFAULT NULL,
    `sectionid`  bigint                                  DEFAULT NULL,
    `user_id`    varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- record: table
CREATE TABLE `record`
(
    `id`            bigint unsigned NOT NULL AUTO_INCREMENT,
    `created_at`    datetime(3) DEFAULT NULL COMMENT '创建时间',
    `updated_at`    datetime(3) DEFAULT NULL COMMENT '更新时间',
    `deleted_at`    datetime(3) DEFAULT NULL,
    `knowledge_id`  bigint unsigned DEFAULT NULL COMMENT '这张卡片跟哪一个知识点有关',
    `Due`           datetime(3) DEFAULT NULL COMMENT '到期时间，也就是该复习的日子',
    `Stability` double DEFAULT NULL,
    `Difficulty` double DEFAULT NULL,
    `ElapsedDays`   bigint unsigned DEFAULT NULL,
    `ScheduledDays` bigint unsigned DEFAULT NULL,
    `Reps`          bigint unsigned DEFAULT NULL,
    `Lapses`        bigint unsigned DEFAULT NULL,
    `State`         tinyint                                 DEFAULT NULL,
    `On`            tinyint(1) DEFAULT NULL COMMENT '是否被使用',
    `lastop`        bigint                                  DEFAULT NULL COMMENT '上一次的选择',
    `LastReview`    datetime(3) DEFAULT NULL,
    `user_id`       varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '属于哪一个用户',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- section: table
CREATE TABLE `section`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT,
    `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
    `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
    `deleted_at` datetime(3) DEFAULT NULL,
    `deckid`     bigint                                  DEFAULT NULL,
    `name`       longtext COLLATE utf8mb4_general_ci,
    `user_id`    varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '属于哪一个用户',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- user: table
CREATE TABLE `user`
(
    `id`             bigint unsigned NOT NULL AUTO_INCREMENT,
    `created_at`     datetime(3) DEFAULT NULL COMMENT '创建时间',
    `updated_at`     datetime(3) DEFAULT NULL COMMENT '更新时间',
    `deleted_at`     datetime(3) DEFAULT NULL,
    `gender`         bigint                                  DEFAULT NULL COMMENT '性别',
    `age`            bigint                                  DEFAULT NULL,
    `password`       longtext COLLATE utf8mb4_general_ci,
    `head_url`       longtext COLLATE utf8mb4_general_ci COMMENT '头像',
    `email`          longtext COLLATE utf8mb4_general_ci,
    `username`       longtext COLLATE utf8mb4_general_ci COMMENT '登录的用户名',
    `introduction`   longtext COLLATE utf8mb4_general_ci,
    `learnnumoneday` bigint                                  DEFAULT NULL COMMENT '一天学多少',
    `user_id`        varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL,
    `nickname`       varchar(20) COLLATE utf8mb4_general_ci  DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

