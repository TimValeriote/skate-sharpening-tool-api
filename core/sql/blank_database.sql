CREATE TABLE `users` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `first_name` varchar(255),
  `last_name` varchar(255),
  `email` varchar(255),
  `phone_number` varchar(255),
  `uuid` varchar(255),
  `is_staff` tinyint(1) NOT NULL DEFAULT 0
);

CREATE TABLE `staff` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `user_id` int
);

CREATE TABLE `brands` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255),
  `short_name` varchar(255),
  `is_skate` tinyint(1),
  `is_steel` tinyint(1),
  `is_holder` tinyint(1)
);

CREATE TABLE `skates` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `model_id` int,
  `brand_id` int
);

CREATE TABLE `model` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255),
  `alias` varchar(255)
);

CREATE TABLE `colour` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `colour` varchar(255)
);

CREATE TABLE `fits` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255)
);

CREATE TABLE `store` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255),
  `address` varchar(255),
  `city` varchar(255),
  `country` varchar(255),
  `phone_number` varchar(255),
  `store_number` int
);

CREATE TABLE `progress_trans` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `progress_text` varchar(255)
);

CREATE TABLE `open_sharpenings` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `user_id` int,
  `user_skate_id` int,
  `store_id` int,
  `progress_trans_id` int
);

CREATE TABLE `user_skates` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `user_id` int,
  `skate_id` int,
  `holder_brand_id` int,
  `holder_size` float,
  `skate_size` float,
  `lace_colour_id` int,
  `has_steel` tinyint(1),
  `steel_id` int,
  `has_guards` tinyint(1),
  `guard_colour_id` int,
  `fit_id` int,
  `preferred_radius` varchar(255)
);

CREATE TABLE `closed_sharpenings` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `user_id` int,
  `user_skate_id` int,
  `store_id` int
);

ALTER TABLE `staff` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `skates` ADD FOREIGN KEY (`model_id`) REFERENCES `model` (`id`);

ALTER TABLE `skates` ADD FOREIGN KEY (`brand_id`) REFERENCES `brands` (`id`);

ALTER TABLE `user_skates` ADD FOREIGN KEY (`fit_id`) REFERENCES `fits` (`id`);

ALTER TABLE `open_sharpenings` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `open_sharpenings` ADD FOREIGN KEY (`user_skate_id`) REFERENCES `user_skates` (`id`);

ALTER TABLE `open_sharpenings` ADD FOREIGN KEY (`store_id`) REFERENCES `store` (`id`);

ALTER TABLE `open_sharpenings` ADD FOREIGN KEY (`progress_trans_id`) REFERENCES `progress_trans` (`id`);

ALTER TABLE `user_skates` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `user_skates` ADD FOREIGN KEY (`skate_id`) REFERENCES `skates` (`id`);

ALTER TABLE `user_skates` ADD FOREIGN KEY (`steel_id`) REFERENCES `brands` (`id`);

ALTER TABLE `user_skates` ADD FOREIGN KEY (`guard_colour_id`) REFERENCES `colour` (`id`);

ALTER TABLE `user_skates` ADD FOREIGN KEY (`lace_colour_id`) REFERENCES `colour` (`id`);

ALTER TABLE `user_skates` ADD FOREIGN KEY (`holder_brand_id`) REFERENCES `brands` (`id`);

ALTER TABLE `closed_sharpenings` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `closed_sharpenings` ADD FOREIGN KEY (`user_skate_id`) REFERENCES `user_skates` (`id`);

ALTER TABLE `closed_sharpenings` ADD FOREIGN KEY (`store_id`) REFERENCES `store` (`id`);
