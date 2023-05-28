CREATE TABLE `users` (
  `id` int PRIMARY KEY,
  `username` varchar(50),
  `password` varchar(255),
  `email` varchar(255),
  `created_at` timestamp NOT NULL DEFAULT "NOW()",
  `updated_at` timestamp,
  `is_deleted` int NOT NULL DEFAULT 0
);
