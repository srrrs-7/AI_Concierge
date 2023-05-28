CREATE TABLE `calendar` (
  `id` int PRIMARY KEY,
  `user_id` int NOT NULL,
  `date` date NOT NULL,
  `day` int NOT NULL,
  `month` int NOT NULL,
  `year` int NOT NULL,
  `holiday` varchar(50)
);
