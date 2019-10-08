-- phpMyAdmin SQL Dump
-- version 4.8.4
-- https://www.phpmyadmin.net/
--
-- 主机： 118.24.51.171
-- 生成日期： 2019-03-04 03:01:31
-- 服务器版本： 5.7.22
-- PHP 版本： 7.2.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- 数据库： `casbin`
--

-- --------------------------------------------------------

--
-- 表的结构 `admin_permissions`
--

CREATE TABLE `admin_permissions` (
  `id` int(10) UNSIGNED NOT NULL,
  `fid` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '菜单父ID',
  `icon` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '图标class',
  `url_path` varchar(255) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `display_name` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `description` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `is_menu` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否作为菜单显示,[1|0]',
  `sort` tinyint(4) NOT NULL DEFAULT '0' COMMENT '排序',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

--
-- 转存表中的数据 `admin_permissions`
--

INSERT INTO `admin_permissions` (`id`, `fid`, `icon`, `url_path`, `display_name`, `description`, `is_menu`, `sort`, `created_at`, `updated_at`) VALUES
(1, 0, '696', '#-1518089385', '管理员管理', 'edit', 1, 1, '2019-02-28 06:58:24', '2018-02-08 03:29:45'),
(2, 1, NULL, '/admin/rbac/permission_list', '权限管理', NULL, 1, 1, '2018-02-08 02:41:09', '2018-02-08 02:41:09'),
(5, 1, NULL, '/admin/rbac/role_list', '角色管理', NULL, 1, 3, '2018-02-09 08:02:04', '2018-02-09 00:02:04'),
(30, 1, '', '/admin/rbac/permission_del', '权限删除', '', 0, 0, '2019-02-28 12:51:58', '2019-02-28 22:44:11'),
(31, 2, NULL, '/admin/rbac/permission_edit', '权限编辑', '权限编辑', 0, 0, '2019-03-01 05:51:18', '2019-03-01 05:51:18'),
(32, 2, NULL, '/admin/rbac/permission_add', '权限添加', '权限添加', 0, 0, '2019-03-01 07:06:00', '2019-03-01 07:06:00'),
(35, 1, '', '/admin/rbac/category_list', '权限分类', '', 1, 0, '2019-02-28 23:27:47', '2019-02-28 23:27:47'),
(36, 1, '', '/admin/rbac/category_add', '权限分类-添加', '', 0, 0, '2019-02-28 23:46:41', '2019-02-28 23:46:41'),
(38, 1, '', '/admin/rbac/category_del', '权限分类-删除', '', 0, 0, '2019-02-28 23:56:53', '2019-02-28 23:56:53'),
(39, 1, '', '/admin/rbac/category_edit', '权限分类-编辑', '', 0, 0, '2019-02-28 23:57:30', '2019-02-28 23:57:30'),
(40, 1, '', '/admin/rbac/role_add', '角色-添加', '', 0, 0, '2019-03-01 00:48:05', '2019-03-01 00:48:05'),
(41, 1, '', '/admin/rbac/role_del', '角色-删除', '', 0, 0, '2019-03-01 02:35:12', '2019-03-01 02:35:12'),
(42, 1, '', '/admin/rbac/role_edit', '角色-编辑', '', 0, 0, '2019-03-01 02:36:02', '2019-03-01 02:36:02'),
(43, 1, '', '/admin/rbac/admin_list', '管理员列表', '', 1, 0, '2019-03-02 03:48:07', '2019-03-02 03:48:07'),
(44, 1, '', '/admin/rbac/admin_add', '管理员-添加', '', 0, 0, '2019-03-02 04:02:56', '2019-03-02 04:02:56'),
(45, 1, '', '/admin/rbac/admin_edit', '管理员-编辑', '', 0, 0, '2019-03-02 05:51:52', '2019-03-02 05:51:52'),
(46, 1, '', '/admin/rbac/admin_del', '管理员-删除', '', 0, 0, '2019-03-02 06:47:15', '2019-03-02 06:47:15');

-- --------------------------------------------------------

--
-- 表的结构 `admin_permission_role`
--

CREATE TABLE `admin_permission_role` (
  `permission_id` int(10) UNSIGNED NOT NULL,
  `role_id` int(10) UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

--
-- 转存表中的数据 `admin_permission_role`
--

INSERT INTO `admin_permission_role` (`permission_id`, `role_id`) VALUES
(1, 1),
(2, 1),
(3, 1),
(5, 1),
(11, 1),
(11, 2),
(11, 8),
(12, 1),
(12, 2),
(12, 8),
(29, 1),
(29, 2),
(29, 8),
(30, 1),
(35, 1),
(36, 1),
(38, 1),
(39, 1),
(40, 1),
(41, 1),
(42, 1);

-- --------------------------------------------------------

--
-- 表的结构 `admin_roles`
--

CREATE TABLE `admin_roles` (
  `id` int(10) UNSIGNED NOT NULL,
  `name` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `display_name` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `description` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

--
-- 转存表中的数据 `admin_roles`
--

INSERT INTO `admin_roles` (`id`, `name`, `display_name`, `description`, `created_at`, `updated_at`) VALUES
(1, 'admin', 'admin', 'all', '2019-03-01 12:02:33', '2019-03-01 12:02:35'),
(2, 'guest', 'guest', '游客', '2019-03-01 12:02:11', '2019-03-01 12:02:13'),
(8, '运营', '运营', '运营', '2019-03-01 11:53:41', '2019-03-01 11:53:41');

-- --------------------------------------------------------

--
-- 表的结构 `admin_role_user`
--

CREATE TABLE `admin_role_user` (
  `admin_user_id` int(10) UNSIGNED NOT NULL,
  `role_id` int(10) UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

--
-- 转存表中的数据 `admin_role_user`
--

INSERT INTO `admin_role_user` (`admin_user_id`, `role_id`) VALUES
(1, 1),
(3, 1);

-- --------------------------------------------------------

--
-- 表的结构 `admin_users`
--

CREATE TABLE `admin_users` (
  `id` int(10) UNSIGNED NOT NULL,
  `name` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `email` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `mobile_num` varchar(16) COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '手机号',
  `password` varchar(60) COLLATE utf8_unicode_ci NOT NULL,
  `is_super` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否超级管理员, 1-是， 0-否',
  `remember_token` varchar(100) COLLATE utf8_unicode_ci DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

--
-- 转存表中的数据 `admin_users`
--

INSERT INTO `admin_users` (`id`, `name`, `email`, `mobile_num`, `password`, `is_super`, `remember_token`, `created_at`, `updated_at`) VALUES
(1, '管理员', 'admin@admin.com', '15372406476', '$2a$04$Mx6zIsjjthPw4NirGd7Eeuh7002T50uN8LwxFLaH89BEwjAyiVmMa', 1, 'XYEfTEAqhW6ZOxHkmoTSTANy6ZrUojV5e2Kniw32LgRpROJ28qmNt8O3kRLY', '2019-02-27 13:36:20', '2019-03-02 06:34:23'),
(3, 'liugaoyun', 'liugy@soulgame.com', '18811597049', '$2a$04$RgUAcNU7Vrd6AQfntktl6OKz8gIwiKt64XN24.tLVccl8HdfEtMga', 1, '', '2019-03-02 05:29:13', '2019-03-02 14:59:02');

--
-- 转储表的索引
--

--
-- 表的索引 `admin_permissions`
--
ALTER TABLE `admin_permissions`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `admin_permission_role`
--
ALTER TABLE `admin_permission_role`
  ADD PRIMARY KEY (`permission_id`,`role_id`);

--
-- 表的索引 `admin_roles`
--
ALTER TABLE `admin_roles`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `roles_name_unique` (`name`);

--
-- 表的索引 `admin_role_user`
--
ALTER TABLE `admin_role_user`
  ADD PRIMARY KEY (`admin_user_id`,`role_id`);

--
-- 表的索引 `admin_users`
--
ALTER TABLE `admin_users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `admin_users_email_unique` (`email`);

--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `admin_permissions`
--
ALTER TABLE `admin_permissions`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=47;

--
-- 使用表AUTO_INCREMENT `admin_roles`
--
ALTER TABLE `admin_roles`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=10;

--
-- 使用表AUTO_INCREMENT `admin_users`
--
ALTER TABLE `admin_users`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
