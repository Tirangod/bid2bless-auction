CREATE TABLE `auctions` (
  `id` INTEGER PRIMARY KEY AUTOINCREMENT,
  `name` VARCHAR(255) NOT NULL,
  `photo_path` VARCHAR(255) NOT NULL,
  `description` VARCHAR(255) NOT NULL,
  `start_at` TIMESTAMP NOT NULL,
  `end_at` TIMESTAMP NOT NULL,
  `initial_bet` INTEGER NOT NULL,
  `bet` INTEGER NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `user_id` INTEGER NOT NULL,
  FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
);

CREATE TABLE `users` (
  `id` INTEGER PRIMARY KEY AUTOINCREMENT,
  `email` VARCHAR(255) UNIQUE,
  `name` VARCHAR(255) NOT NULL,
  `hashed_password` VARCHAR(255) NOT NULL,
  `balance` INTEGER NOT NULL,
  `inuse_balance` INTEGER NOT NULL
);

CREATE TABLE `bet_history` (
  `id` INTEGER PRIMARY KEY AUTOINCREMENT,
  `auction_id` INTEGER NOT NULL,
  `updated_bet` INTEGER NOT NULL,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `user_id` INTEGER NOT NULL,
  FOREIGN KEY (`auction_id`) REFERENCES `auctions` (`id`),
  FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
);

CREATE TABLE `messages` (
  `id` INTEGER PRIMARY KEY AUTOINCREMENT,
  `auction_id` INTEGER NOT NULL,
  `sent_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `user_id` INTEGER NOT NULL,
  FOREIGN KEY (`auction_id`) REFERENCES `auctions` (`id`),
  FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
);

CREATE TABLE `categories` (
  `name` VARCHAR(255) PRIMARY KEY
);

CREATE TABLE `auctions_categories` (
  `connection_id` INTEGER PRIMARY KEY AUTOINCREMENT,
  `auction_id` INTEGER NOT NULL,
  `category_id` VARCHAR(255) NOT NULL,
  FOREIGN KEY (`auction_id`) REFERENCES `auctions` (`id`),
  FOREIGN KEY (`category_id`) REFERENCES `categories` (`name`)
);
