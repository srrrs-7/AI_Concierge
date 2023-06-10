INSERT INTO `auth` (`id`, `user_id`, `password`, `mail`, `scope`, `created_at`, `is_deleted`)
VALUES
  (1, 'john123', 'password123', 'john@example.com', 'read', NOW(), 0),
  (2, 'emma456', 'securepass789', 'emma@example.com', 'write', NOW(), 0),
  (3, 'mark789', 'strongpassword', 'mark@example.com', 'readwrite', NOW(), 0);
