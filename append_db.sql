CREATE TABLE `env` (
                         `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
                         `name` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                         PRIMARY KEY (`id`)
);

CREATE TABLE `parameters` (
                                        `id` INT unsigned NOT NULL AUTO_INCREMENT,
                                        `param_key` VARCHAR(45) NULL,
                                        `param_value` VARCHAR(45) NULL,
                                        `param_type` VARCHAR(45) NULL,
                                        `env` VARCHAR(45) NULL,
                                        `domain` VARCHAR(45) NULL,
                                        `update_user` VARCHAR(45) NULL,
                                        `last_update` DATETIME NULL,
                                        `create_at` DATETIME NULL,
                                        PRIMARY KEY (`id`));


