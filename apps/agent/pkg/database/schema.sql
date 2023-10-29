-- Create "apis" table
CREATE TABLE `apis` (
  `id` varchar(256) NOT NULL,
  `name` varchar(256) NOT NULL,
  `workspace_id` varchar(256) NOT NULL,
  `ip_whitelist` varchar(512) NULL,
  `auth_type` enum('key', 'jwt') NULL,
  `key_auth_id` varchar(256) NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `key_auth_id_idx` (`key_auth_id`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;

-- Create "audit_log_changes" table
CREATE TABLE `audit_log_changes` (
  `audit_log_id` varchar(256) NOT NULL,
  `field` varchar(256) NOT NULL,
  `old` varchar(1024) NULL,
  `new` varchar(1024) NULL,
  PRIMARY KEY (`audit_log_id`, `field`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;

-- Create "audit_logs" table
CREATE TABLE `audit_logs` (
  `id` varchar(256) NOT NULL,
  `workspace_id` varchar(256) NOT NULL,
  `action` enum('created', 'updated', 'deleted') NOT NULL,
  `description` varchar(512) NOT NULL,
  `time` datetime(3) NOT NULL,
  `actor_type` enum('user', 'key') NOT NULL,
  `actor_id` varchar(256) NOT NULL,
  `resource_type` enum('key', 'api', 'workspace') NOT NULL,
  `resource_id` varchar(256) NOT NULL,
  `tags` json NULL,
  PRIMARY KEY (`id`),
  INDEX `actor_id_idx` (`actor_id`),
  INDEX `resource_id_idx` (`resource_id`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;

-- Create "key_auth" table
CREATE TABLE `key_auth` (
  `id` varchar(256) NOT NULL,
  `workspace_id` varchar(256) NOT NULL,
  PRIMARY KEY (`id`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;

-- Create "keys" table
CREATE TABLE `keys` (
  `id` varchar(256) NOT NULL,
  `hash` varchar(256) NOT NULL,
  `start` varchar(256) NOT NULL,
  `owner_id` varchar(256) NULL,
  `meta` text NULL,
  `created_at` datetime(3) NOT NULL,
  `expires` datetime(3) NULL,
  `ratelimit_type` text NULL,
  `ratelimit_limit` int NULL,
  `ratelimit_refill_rate` int NULL,
  `ratelimit_refill_interval` int NULL,
  `workspace_id` varchar(256) NOT NULL,
  `for_workspace_id` varchar(256) NULL,
  `name` varchar(256) NULL,
  `remaining_requests` int NULL,
  `key_auth_id` varchar(256) NOT NULL,
  `total_uses` bigint NULL DEFAULT 0,
  `deleted_at` datetime(3) NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `hash_idx` (`hash`),
  INDEX `key_auth_id_idx` (`key_auth_id`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;

-- Create "vercel_bindings" table
CREATE TABLE `vercel_bindings` (
  `id` varchar(256) NOT NULL,
  `integration_id` varchar(256) NOT NULL,
  `workspace_id` varchar(256) NOT NULL,
  `project_id` varchar(256) NOT NULL,
  `environment` enum('development', 'preview', 'production') NOT NULL,
  `resource_id` varchar(256) NOT NULL,
  `resource_type` enum('rootKey', 'apiId') NOT NULL,
  `vercel_env_id` varchar(256) NOT NULL,
  `created_at` datetime(3) NOT NULL,
  `updated_at` datetime(3) NOT NULL,
  `last_edited_by` varchar(256) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `project_environment_resource_type_idx` (`project_id`, `environment`, `resource_type`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;

-- Create "vercel_integrations" table
CREATE TABLE `vercel_integrations` (
  `id` varchar(256) NOT NULL,
  `workspace_id` varchar(256) NOT NULL,
  `team_id` varchar(256) NULL,
  `access_token` varchar(256) NOT NULL,
  PRIMARY KEY (`id`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;

-- Create "workspaces" table
CREATE TABLE `workspaces` (
  `id` varchar(256) NOT NULL,
  `name` varchar(256) NOT NULL,
  `slug` varchar(256) NULL,
  `tenant_id` varchar(256) NOT NULL,
  `stripe_customer_id` varchar(256) NULL,
  `stripe_subscription_id` varchar(256) NULL,
  `plan` enum('free', 'pro', 'enterprise') NULL DEFAULT "free",
  `quota_max_active_keys` int NULL,
  `usage_active_keys` int NULL,
  `quota_max_verifications` int NULL,
  `usage_verifications` int NULL,
  `last_usage_update` datetime(3) NULL,
  `billing_period_start` datetime(3) NULL,
  `billing_period_end` datetime(3) NULL,
  `trial_ends` datetime(3) NULL,
  `features` json NOT NULL,
  `beta_features` json NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `slug_idx` (`slug`),
  UNIQUE INDEX `tenant_id_idx` (`tenant_id`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;