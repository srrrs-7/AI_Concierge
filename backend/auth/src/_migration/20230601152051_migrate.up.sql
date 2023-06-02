CREATE TABLE `auth` (
  `id` int PRIMARY KEY,
  `user_id` varchar(50) NOT NULL,
  `password` varchar(50) NOT NULL,
  `mail` varchar(255) NOT NULL,
  `scope` varchar(50) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT NOW(),
  `updated_at` timestamp,
  `is_deleted` int NOT NULL DEFAULT 0
);