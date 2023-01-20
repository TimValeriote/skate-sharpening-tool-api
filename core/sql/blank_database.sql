CREATE TABLE `users` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `first_name` varchar(255),
  `last_name` varchar(255),
  `email` varchar(255),
  `phone_number` varchar(255),
  `uuid` varchar(255),
  `isStaff` tinyint(1) DEFAULT 0 NOT NULL
);

CREATE TABLE `staff` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `user_id` int
);

CREATE TABLE `brands` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255),
  `short_name` varchar(255)
);

CREATE TABLE `skates` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `model_id` int,
  `brand_id` int,
  `fit_id` int,
  `size` int
);

CREATE TABLE `model` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255),
  `brand_id` int
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
  `phone_number` varchar(255)
);

CREATE TABLE `open_sharpenings` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `user_id` int,
  `user_skate_id` int,
  `store_id` int
);

CREATE TABLE `user_skates` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `user_id` int,
  `skate` int
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

ALTER TABLE `skates` ADD FOREIGN KEY (`fit_id`) REFERENCES `fits` (`id`);

ALTER TABLE `model` ADD FOREIGN KEY (`brand_id`) REFERENCES `brands` (`id`);

ALTER TABLE `open_sharpenings` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `open_sharpenings` ADD FOREIGN KEY (`user_skate_id`) REFERENCES `user_skates` (`id`);

ALTER TABLE `open_sharpenings` ADD FOREIGN KEY (`store_id`) REFERENCES `store` (`id`);

ALTER TABLE `user_skates` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `user_skates` ADD FOREIGN KEY (`skate`) REFERENCES `skates` (`id`);

ALTER TABLE `closed_sharpenings` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `closed_sharpenings` ADD FOREIGN KEY (`user_skate_id`) REFERENCES `user_skates` (`id`);

ALTER TABLE `closed_sharpenings` ADD FOREIGN KEY (`store_id`) REFERENCES `store` (`id`);
