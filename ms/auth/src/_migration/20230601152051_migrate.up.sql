CREATE TABLE `client_master` (
  `client_id` varchar(50) PRIMARY KEY,
  `hash_password` varchar(255) NOT NULL UNIQUE,
  `email` varchar(255) NOT NULL UNIQUE,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `is_deleted` int NOT NULL DEFAULT 0
);
CREATE INDEX idx_client_id ON client_master(client_id);

CREATE TABLE authorization_code (
  `client_id` varchar(50) PRIMARY KEY,
  `code` VARCHAR(255) NOT NULL UNIQUE,
  `client_secret` varchar(255) NOT NULL UNIQUE,
  `scope` varchar(50) NOT NULL,
  `grant_type` varchar(50) NOT NULL,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `is_deleted` int NOT NULL DEFAULT 0,
  FOREIGN KEY (client_id) REFERENCES client_master(client_id)
);
CREATE INDEX idx_client_id ON authorization_code(client_id);

CREATE TABLE `client_token` (
  `client_id` varchar(50) PRIMARY KEY,
  `access_token` varchar(255) NOT NULL UNIQUE,
  `refresh_token` varchar(255) NOT NULL UNIQUE,
  `expires_in` TIMESTAMP NOT NULL,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `is_deleted` int NOT NULL DEFAULT 0,
  FOREIGN KEY (client_id) REFERENCES client_master(client_id)
);
CREATE INDEX idx_client_id ON client_token(client_id);

