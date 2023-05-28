
-- usersテーブルに値を挿入
INSERT INTO `users` (`id`, `username`, `password`, `email`, `created_at`, `updated_at`, `is_deleted`)
VALUES
  (1, 'ユーザー1', 'パスワード1', 'user1@example.com', NOW(), NULL, 0),
  (2, 'ユーザー2', 'パスワード2', 'user2@example.com', NOW(), NULL, 0),
  (3, 'ユーザー3', 'パスワード3', 'user3@example.com', NOW(), NULL, 0);

-- calendarテーブルに値を挿入
INSERT INTO `calendar` (`id`, `user_id`, `date`, `day`, `month`, `year`, `holiday`)
VALUES
  (1, 1, '2023-05-28', 28, 5, 2023, NULL),
  (2, 1, '2023-05-29', 29, 5, 2023, NULL),
  (3, 2, '2023-05-28', 28, 5, 2023, NULL);

-- eventsテーブルに値を挿入
INSERT INTO `events` (`id`, `user_id`, `event`, `from_date`, `to_date`, `description`, `created_at`, `updated_at`, `is_deleted`)
VALUES
  (1, 1, 'イベント1', '2023-06-01', '2023-06-03', 'イベントの詳細1', NOW(), NULL, 0),
  (2, 1, 'イベント2', '2023-06-05', '2023-06-07', 'イベントの詳細2', NOW(), NULL, 0),
  (3, 2, 'イベント3', '2023-06-10', '2023-06-12', 'イベントの詳細3', NOW(), NULL, 0);

-- notificationsテーブルに値を挿入
INSERT INTO `notifications` (`id`, `user_id`, `event_id`, `created_at`, `updated_at`, `is_deleted`)
VALUES
  (1, 1, 1, NOW(), NULL, 0),
  (2, 1, 2, NOW(), NULL, 0),
  (3, 2, 3, NOW(), NULL, 0);

-- workSpacesテーブルに値を挿入
INSERT INTO `work_spaces` (`id`, `work_space`, `created_at`, `updated_at`, `is_deleted`)
VALUES
  (1, 'ワークスペース1', NOW(), NULL, 0),
  (2, 'ワークスペース2', NOW(), NULL, 0);

-- chatsテーブルに値を挿入
INSERT INTO `chats` (`id`, `user_id`, `work_space_id`, `message`, `created_at`, `updated_at`, `is_deleted`)
VALUES
  (1, 1, 1, 'メッセージ1', NOW(), NULL, 0),
  (2, 1, 1, 'メッセージ2', NOW(), NULL, 0);


