CREATE TABLE `users`
(
  `id` int PRIMARY KEY,
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT NOW(),
  `updated_at` timestamp,
  `is_deleted` int NOT NULL DEFAULT 0
);

CREATE TABLE `calendar`
(
  `id` int PRIMARY KEY,
  `user_id` int NOT NULL,
  `date` date NOT NULL,
  `day` int NOT NULL,
  `month` int NOT NULL,
  `year` int NOT NULL,
  `holiday` varchar(50)
);

CREATE TABLE `events`
(
  `id` int PRIMARY KEY,
  `user_id` int NOT NULL,
  `event` varchar(100) NOT NULL,
  `from_date` date NOT NULL,
  `to_date` date NOT NULL,
  `description` text,
  `created_at` timestamp NOT NULL DEFAULT NOW(),
  `updated_at` timestamp,
  `is_deleted` int NOT NULL DEFAULT 0
);

CREATE TABLE `notifications`
(
  `id` int PRIMARY KEY,
  `user_id` int NOT NULL,
  `event_id` int,
  `created_at` timestamp NOT NULL DEFAULT NOW(),
  `updated_at` timestamp,
  `is_deleted` int NOT NULL DEFAULT 0
);

CREATE TABLE `work_spaces`
(
  `id` int PRIMARY KEY,
  `work_space` varchar(50) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT NOW(),
  `updated_at` timestamp,
  `is_deleted` int NOT NULL DEFAULT 0
);

CREATE TABLE `chats`
(
  `id` int PRIMARY KEY,
  `user_id` int NOT NULL,
  `work_space_id` int NOT NULL,
  `message` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT NOW(),
  `updated_at` timestamp,
  `is_deleted` int NOT NULL DEFAULT 0
);

ALTER TABLE `calendar` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `events` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `notifications` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `notifications` ADD FOREIGN KEY (`event_id`) REFERENCES `events` (`id`);

ALTER TABLE `chats` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `chats` ADD FOREIGN KEY (`work_space_id`) REFERENCES `work_spaces` (`id`);