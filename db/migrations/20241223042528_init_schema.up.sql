
CREATE TABLE `users` (
    `username` VARCHAR(30) PRIMARY KEY,
    `password` VARCHAR(255),
    `role` VARCHAR(30),
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE `inventories` (
    `id` INTEGER PRIMARY KEY AUTOINCREMENT,
    `name` TEXT NOT NULL,
    `quantity` NUMBER NOT NULL,
    `category_id` NUMBER NOT NULL,
    `condition` INT NOT NULL DEFAULT 0,
    `status` NUMBER NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (`category_id`) REFERENCES `ref_categories` (`id`),
    FOREIGN KEY (`status`) REFERENCES `ref_status` (`id`)
);

CREATE TABLE `consumables` (
    `id` INTEGER PRIMARY KEY AUTOINCREMENT,
    `name` TEXT NOT NULL,
    `quantity` NUMBER NOT NULL,
    `category_id` NUMBER NOT NULL,
    `status` NUMBER NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (`category_id`) REFERENCES `ref_categories` (`id`),
    FOREIGN KEY (`status`) REFERENCES `ref_status` (`id`)
);

CREATE TABLE `ref_categories` (
    `id` INTEGER PRIMARY KEY AUTOINCREMENT,
    `name` VARCHAR(255) NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE `ref_status` (
    `id` INTEGER PRIMARY KEY AUTOINCREMENT,
    `description` TEXT NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
