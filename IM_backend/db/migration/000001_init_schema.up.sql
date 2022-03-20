CREATE TABLE `users` (
  `id` bigint PRIMARY KEY AUTO_INCREMENT,
  `username` varchar(255) UNIQUE NOT NULL,
  `hashed_password` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT now()
);

CREATE TABLE `conversation` (
  `id` bigint PRIMARY KEY AUTO_INCREMENT,
  `owner` bigint NOT NULL,
  `user_id` bigint NOT NULL
);

CREATE TABLE `friend` (
  `id` bigint PRIMARY KEY AUTO_INCREMENT,
  `owner` bigint NOT NULL,
  `friend_id` bigint NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT now()
);

ALTER TABLE `conversation` ADD FOREIGN KEY (`owner`) REFERENCES `users` (`id`);

ALTER TABLE `conversation` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `friend` ADD FOREIGN KEY (`owner`) REFERENCES `users` (`id`);

ALTER TABLE `friend` ADD FOREIGN KEY (`friend_id`) REFERENCES `users` (`id`);

CREATE UNIQUE INDEX `conversation_index_0` ON `conversation` (`owner`, `user_id`);

CREATE UNIQUE INDEX `friend_index_1` ON `friend` (`owner`, `friend_id`);