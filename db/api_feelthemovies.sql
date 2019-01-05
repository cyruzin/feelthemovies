-- phpMyAdmin SQL Dump
-- version 4.8.3
-- https://www.phpmyadmin.net/
--
-- Host: mysql
-- Tempo de geração: 05/01/2019 às 17:17
-- Versão do servidor: 5.7.24
-- Versão do PHP: 7.2.8

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Banco de dados: `api_feelthemovies`
--

-- --------------------------------------------------------

--
-- Estrutura para tabela `genres`
--

CREATE TABLE `genres` (
  `id` int(10) UNSIGNED NOT NULL,
  `name` varchar(30) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Fazendo dump de dados para tabela `genres`
--

INSERT INTO `genres` (`id`, `name`, `created_at`, `updated_at`) VALUES
(1, 'Horror', '2019-01-03 00:20:12', '2019-01-03 00:20:12'),
(2, 'Thriller', '2019-01-03 00:20:17', '2019-01-03 00:20:17'),
(4, 'Comedy', '2019-01-03 02:41:28', '2019-01-03 02:41:28'),
(5, 'Drama', '2019-01-03 02:41:32', '2019-01-03 02:41:32'),
(6, 'Family', '2019-01-03 16:16:49', '2019-01-03 16:16:49'),
(7, 'Biography', '2019-01-03 16:17:06', '2019-01-03 16:17:42'),
(8, 'Ocean', '2019-01-03 16:17:52', '2019-01-03 16:17:52'),
(9, 'War', '2019-01-04 00:05:04', '2019-01-04 00:05:04'),
(10, 'Music', '2019-01-04 20:09:54', '2019-01-04 20:09:54'),
(11, 'Fantasy', '2019-01-04 20:13:49', '2019-01-04 20:13:49'),
(12, 'Western', '2019-01-04 20:15:22', '2019-01-04 20:15:22'),
(13, 'NewGenre', '2019-01-04 20:31:26', '2019-01-04 20:31:26'),
(17, 'Animation', '2019-01-05 01:29:11', '2019-01-05 01:29:11'),
(15, 'Action', '2019-01-04 22:54:01', '2019-01-05 01:28:52'),
(16, 'Adventure', '2019-01-04 22:54:04', '2019-01-04 22:55:32');

-- --------------------------------------------------------

--
-- Estrutura para tabela `genre_recommendation`
--

CREATE TABLE `genre_recommendation` (
  `recommendation_id` int(10) UNSIGNED NOT NULL,
  `genre_id` int(10) UNSIGNED NOT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Fazendo dump de dados para tabela `genre_recommendation`
--

INSERT INTO `genre_recommendation` (`recommendation_id`, `genre_id`) VALUES
(14, 1),
(14, 2),
(15, 1),
(15, 2),
(15, 4),
(15, 5),
(16, 1),
(16, 2),
(16, 4),
(16, 5),
(19, 5),
(21, 1),
(21, 5),
(22, 5),
(23, 5),
(24, 5),
(25, 5),
(27, 1),
(28, 1),
(35, 1),
(36, 1),
(37, 1),
(38, 4),
(38, 1),
(40, 1),
(39, 4),
(39, 1),
(41, 9),
(41, 1),
(42, 9),
(42, 1),
(45, 9),
(45, 1),
(46, 9),
(46, 1),
(47, 10),
(48, 10),
(49, 10),
(50, 10),
(51, 1),
(52, 12),
(53, 13),
(54, 16),
(54, 17),
(55, 13),
(57, 17),
(56, 1),
(58, 7),
(59, 5),
(60, 11),
(61, 16),
(61, 9),
(62, 17),
(63, 17),
(64, 1);

-- --------------------------------------------------------

--
-- Estrutura para tabela `keywords`
--

CREATE TABLE `keywords` (
  `id` int(10) UNSIGNED NOT NULL,
  `name` varchar(30) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Fazendo dump de dados para tabela `keywords`
--

INSERT INTO `keywords` (`id`, `name`, `created_at`, `updated_at`) VALUES
(5, 'Alligator', '2019-01-03 02:41:07', '2019-01-03 02:41:07'),
(4, 'Shark', '2019-01-03 02:13:29', '2019-01-03 02:13:29'),
(6, 'Blood', '2019-01-03 02:41:16', '2019-01-04 22:55:46'),
(7, 'Bloodbath', '2019-01-03 17:37:11', '2019-01-03 17:37:11'),
(10, 'Golden Retriever', '2019-01-05 01:28:16', '2019-01-05 01:28:16'),
(8, 'Fork Silver', '2019-01-04 23:09:14', '2019-01-05 01:27:43'),
(9, 'Martial Arts', '2019-01-04 23:18:35', '2019-01-04 23:18:41');

-- --------------------------------------------------------

--
-- Estrutura para tabela `keyword_recommendation`
--

CREATE TABLE `keyword_recommendation` (
  `recommendation_id` int(10) UNSIGNED NOT NULL,
  `keyword_id` int(10) UNSIGNED NOT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Fazendo dump de dados para tabela `keyword_recommendation`
--

INSERT INTO `keyword_recommendation` (`recommendation_id`, `keyword_id`) VALUES
(14, 1),
(14, 4),
(15, 6),
(16, 6),
(19, 4),
(21, 1),
(21, 4),
(22, 4),
(23, 4),
(24, 4),
(25, 4),
(26, 4),
(27, 4),
(28, 4),
(35, 1),
(36, 1),
(37, 1),
(40, 5),
(38, 5),
(39, 5),
(41, 1),
(42, 1),
(43, 1),
(44, 1),
(45, 1),
(46, 1),
(47, 1),
(48, 1),
(49, 1),
(50, 1),
(51, 1),
(52, 1),
(53, 1),
(54, 7),
(54, 6),
(55, 1),
(57, 9),
(56, 5),
(58, 9),
(59, 10),
(60, 9),
(61, 9),
(61, 5),
(62, 9),
(63, 7),
(64, 5);

-- --------------------------------------------------------

--
-- Estrutura para tabela `migrations`
--

CREATE TABLE `migrations` (
  `id` int(10) UNSIGNED NOT NULL,
  `migration` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `batch` int(11) NOT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Fazendo dump de dados para tabela `migrations`
--

INSERT INTO `migrations` (`id`, `migration`, `batch`) VALUES
(1, '2018_09_30_013707_create_users_table', 1),
(2, '2018_09_30_013722_create_recommendations_table', 1),
(3, '2018_09_30_014443_create_recommendation_items_table', 1),
(4, '2018_10_01_160531_create_genres_table', 1),
(5, '2018_10_01_160552_create_genre_recommendation_table', 1),
(6, '2018_10_01_170339_create_keywords_table', 1),
(7, '2018_10_01_170402_create_keyword_recommendation_table', 1),
(8, '2018_10_01_233309_create_sources_table', 1),
(9, '2018_10_01_233540_create_recommendation_item_source_table', 1),
(10, '2018_10_31_235414_add_media_type_to_recommendation_items_table', 2);

-- --------------------------------------------------------

--
-- Estrutura para tabela `recommendations`
--

CREATE TABLE `recommendations` (
  `id` int(10) UNSIGNED NOT NULL,
  `user_id` int(10) UNSIGNED NOT NULL,
  `title` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `type` tinyint(4) NOT NULL,
  `body` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `poster` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `backdrop` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `status` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Fazendo dump de dados para tabela `recommendations`
--

INSERT INTO `recommendations` (`id`, `user_id`, `title`, `type`, `body`, `poster`, `backdrop`, `status`, `created_at`, `updated_at`) VALUES
(1, 1, 'Go is a beast!', 0, 'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.', 'pOiiUyTZmlLznqaa', 'IjqIzzMninQmokaaa', 0, '2019-01-02 23:46:23', '2019-01-02 23:51:06'),
(4, 1, 'Java is another beast!', 0, 'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.', 'IjqIzzMninQmok', 'pOiiUyTZmlLznq', 0, '2019-01-03 02:11:03', '2019-01-03 02:11:03'),
(5, 1, 'Python is another beast!', 0, 'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.', 'IjqIzzMninQmok', 'pOiiUyTZmlLznq', 0, '2019-01-03 02:12:38', '2019-01-03 02:12:38'),
(6, 1, 'Python is another beast!', 0, 'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.', 'IjqIzzMninQmok', 'pOiiUyTZmlLznq', 0, '2019-01-03 02:15:29', '2019-01-03 02:15:29'),
(7, 1, 'Python is another beast!', 0, 'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.', 'IjqIzzMninQmok', 'pOiiUyTZmlLznq', 0, '2019-01-03 02:16:10', '2019-01-03 02:16:10'),
(8, 1, 'Python is another beast!', 0, 'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.', 'IjqIzzMninQmok', 'pOiiUyTZmlLznq', 0, '2019-01-03 02:16:29', '2019-01-03 02:16:29'),
(9, 1, 'Python is another beast!', 0, 'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.', 'IjqIzzMninQmok', 'pOiiUyTZmlLznq', 0, '2019-01-03 02:17:41', '2019-01-03 02:17:41'),
(10, 1, 'Python is another beast!', 0, 'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.', 'IjqIzzMninQmok', 'pOiiUyTZmlLznq', 0, '2019-01-03 02:19:14', '2019-01-03 02:19:14'),
(11, 1, 'Python is another beast!', 0, 'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.', 'IjqIzzMninQmok', 'pOiiUyTZmlLznq', 0, '2019-01-03 02:22:50', '2019-01-03 02:22:50'),
(12, 1, 'Python is another beast!', 0, 'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.', 'IjqIzzMninQmok', 'pOiiUyTZmlLznq', 0, '2019-01-03 02:24:45', '2019-01-03 02:24:45'),
(13, 1, 'Python is another beast!', 0, 'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.', 'IjqIzzMninQmok', 'pOiiUyTZmlLznq', 0, '2019-01-03 02:25:56', '2019-01-03 02:25:56'),
(14, 1, 'Python is another beast!', 0, 'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.', 'IjqIzzMninQmok', 'pOiiUyTZmlLznq', 0, '2019-01-03 02:27:33', '2019-01-03 02:27:33'),
(15, 1, 'Python is another beast!', 0, 'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.', 'IjqIzzMninQmok', 'pOiiUyTZmlLznq', 0, '2019-01-03 02:42:07', '2019-01-03 02:42:07'),
(16, 1, 'Python is another beast!', 0, 'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.', 'IjqIzzMninQmok', 'pOiiUyTZmlLznq', 0, '2019-01-03 02:49:38', '2019-01-03 02:49:38'),
(17, 1, 'Python is another beast!', 0, 'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.', 'IjqIzzMninQmok', 'pOiiUyTZmlLznq', 0, '2019-01-03 03:16:59', '2019-01-03 03:16:59'),
(18, 1, 'Python is another beast!', 0, 'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.', 'IjqIzzMninQmok', 'pOiiUyTZmlLznq', 0, '2019-01-03 03:18:11', '2019-01-03 03:18:11'),
(19, 1, 'Python is another beast!', 0, 'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.', 'IjqIzzMninQmok', 'pOiiUyTZmlLznq', 0, '2019-01-03 03:18:55', '2019-01-03 03:18:55'),
(20, 1, 'Python is another beast!', 0, 'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.', 'IjqIzzMninQmok', 'pOiiUyTZmlLznq', 0, '2019-01-03 03:19:19', '2019-01-03 03:19:19'),
(21, 1, 'Python is another beast!', 0, 'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.', 'pOiiUyTZmlLznq', 'IjqIzzMninQmok', 1, '2019-01-03 03:19:49', '2019-01-03 04:04:56'),
(22, 1, 'PHP is another beast!', 0, 'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.', 'IjqIzzMninQmok', 'pOiiUyTZmlLznq', 0, '2019-01-03 20:14:59', '2019-01-03 20:14:59'),
(23, 1, 'PHP is another beast!', 0, 'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.', 'IjqIzzMninQmok', 'pOiiUyTZmlLznq', 0, '2019-01-03 20:15:23', '2019-01-03 20:15:23'),
(24, 1, 'PHP is another beast!', 0, 'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.', 'IjqIzzMninQmok', 'pOiiUyTZmlLznq', 0, '2019-01-03 20:15:50', '2019-01-03 20:15:50'),
(25, 1, 'PHP is another beast!', 0, 'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.', 'IjqIzzMninQmok', 'pOiiUyTZmlLznq', 0, '2019-01-03 20:16:12', '2019-01-03 20:16:12'),
(26, 1, 'PHP is another beast!', 0, 'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.', 'IjqIzzMninQmok', 'pOiiUyTZmlLznq', 0, '2019-01-03 20:16:26', '2019-01-03 20:16:26'),
(27, 1, 'PHP is another beast!', 2, 'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.', 'IjqIzzMninQmok', 'pOiiUyTZmlLznq', 0, '2019-01-03 20:27:06', '2019-01-03 20:27:06'),
(28, 1, 'PHP is another beast!', 2, 'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.', 'IjqIzzMninQmok', 'pOiiUyTZmlLznq', 0, '2019-01-03 20:29:00', '2019-01-03 20:29:00'),
(29, 1, 'PHP is another beast!', 2, 'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.', 'IjqIzzMninQmok', 'pOiiUyTZmlLznq', 0, '2019-01-03 20:29:34', '2019-01-03 20:29:34'),
(30, 1, 'PHP is another beast!', 2, 'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.', 'IjqIzzMninQmok', 'pOiiUyTZmlLznq', 0, '2019-01-03 20:30:36', '2019-01-03 20:30:36'),
(31, 1, 'PHP is another beast!', 2, 'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.', 'IjqIzzMninQmok', 'pOiiUyTZmlLznq', 0, '2019-01-03 20:31:22', '2019-01-03 20:31:22'),
(32, 1, 'PHP is another beast!', 2, 'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.', 'IjqIzzMninQmok', 'pOiiUyTZmlLznq', 0, '2019-01-03 20:36:05', '2019-01-03 20:36:05'),
(33, 1, 'PHP is another beast!', 2, 'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.', 'IjqIzzMninQmok', 'pOiiUyTZmlLznq', 0, '2019-01-03 20:36:22', '2019-01-03 20:36:22'),
(34, 1, 'PHP is another beast!', 2, 'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.', 'IjqIzzMninQmok', 'pOiiUyTZmlLznq', 0, '2019-01-03 20:36:55', '2019-01-03 20:36:55'),
(35, 1, 'PHP is another beast!', 2, 'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.', 'IjqIzzMninQmok', 'pOiiUyTZmlLznq', 0, '2019-01-03 20:39:04', '2019-01-03 20:39:04'),
(61, 1, 'Teste do Insomnia 2', 0, '<p>Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.</p>', 'IjqIzzMninQmok', 'pOiiUyTZmlLznq', 0, '2019-01-05 01:59:20', '2019-01-05 01:59:20'),
(37, 1, 'PHP is another beast!', 2, 'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.', 'IjqIzzMninQmok', 'pOiiUyTZmlLznq', 0, '2019-01-03 20:46:37', '2019-01-03 20:46:37'),
(54, 1, 'Golang é foda!', 1, '<p>Golang &eacute; foda!</p>', '/5jsSUKwg8pnFqmuBfCA5aESDVI7.jpg', '/ee5D9trr1qJrNom0eksacWzCfuT.jpg', 0, '2019-01-04 23:20:25', '2019-01-05 01:31:04'),
(55, 1, 'Teste do Insomnia', 2, 'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.', 'IjqIzzMninQmok', 'pOiiUyTZmlLznq', 0, '2019-01-05 01:37:05', '2019-01-05 01:37:05'),
(56, 1, 'Teste', 1, 'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.', 'pOiiUyTZmlLznq', 'IjqIzzMninQmok', 1, '2019-01-05 01:37:30', '2019-01-05 01:37:46'),
(40, 1, 'Teste', 1, 'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.', 'pOiiUyTZmlLznq', 'IjqIzzMninQmok', 1, '2019-01-04 00:05:24', '2019-01-04 00:07:04'),
(62, 1, 'aa', 0, '<p>aa</p>', '/9XgVqYPY7gDuCpymmTwwpXA1IB5.jpg', '/chLSPHjb1dAZRKstpx17lnDbZU9.jpg', 0, '2019-01-05 02:11:10', '2019-01-05 02:11:10'),
(63, 1, 'aba', 0, '<p>asasas</p>', '/9XgVqYPY7gDuCpymmTwwpXA1IB5.jpg', '/chLSPHjb1dAZRKstpx17lnDbZU9.jpg', 1, '2019-01-05 02:24:19', '2019-01-05 02:36:48'),
(64, 1, 'Teste', 0, '<p>Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.</p>', 'pOiiUyTZmlLznq', 'IjqIzzMninQmok', 0, '2019-01-05 02:35:00', '2019-01-05 02:37:22'),
(57, 1, 'a', 0, '<p>a</p>', '/9XgVqYPY7gDuCpymmTwwpXA1IB5.jpg', '/chLSPHjb1dAZRKstpx17lnDbZU9.jpg', 0, '2019-01-05 01:48:34', '2019-01-05 01:48:34'),
(58, 1, 'bb', 0, '<p>bb</p>', '/5jsSUKwg8pnFqmuBfCA5aESDVI7.jpg', '/ee5D9trr1qJrNom0eksacWzCfuT.jpg', 0, '2019-01-05 01:49:14', '2019-01-05 01:49:14'),
(59, 1, 'afsds', 2, '<p>sdsd</p>', '/y6JABtgWMVYPx84Rvy7tROU5aNH.jpg', '/mgOZSS2FFIGtfVeac1buBw3Cx5w.jpg', 0, '2019-01-05 01:49:46', '2019-01-05 01:49:46'),
(60, 1, 'yyy', 1, '<p>yyy</p>', '/5WuiQt7pogeaa2uKJiyX2icytBx.jpg', '/9S0SL3gsWd3HaNuELJDHSSdNOPm.jpg', 0, '2019-01-05 01:50:09', '2019-01-05 01:50:09'),
(48, 1, 'Lord of the Rings', 2, 'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.', 'IjqIzzMninQmok', 'pOiiUyTZmlLznq', 0, '2019-01-04 20:11:25', '2019-01-04 20:11:25'),
(49, 1, 'Lord of the Rings II', 2, 'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.', 'IjqIzzMninQmok', 'pOiiUyTZmlLznq', 0, '2019-01-04 20:11:48', '2019-01-04 20:11:48'),
(51, 1, 'Rambo', 2, 'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.', 'IjqIzzMninQmok', 'pOiiUyTZmlLznq', 0, '2019-01-04 20:14:00', '2019-01-04 20:14:00'),
(52, 1, 'Huehue', 2, 'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.', 'IjqIzzMninQmok', 'pOiiUyTZmlLznq', 0, '2019-01-04 20:15:38', '2019-01-04 20:15:38'),
(53, 1, 'New Genre', 2, 'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.', 'IjqIzzMninQmok', 'pOiiUyTZmlLznq', 0, '2019-01-04 20:31:46', '2019-01-04 20:31:46');

-- --------------------------------------------------------

--
-- Estrutura para tabela `recommendation_items`
--

CREATE TABLE `recommendation_items` (
  `id` int(10) UNSIGNED NOT NULL,
  `recommendation_id` int(10) UNSIGNED NOT NULL,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `tmdb_id` int(11) NOT NULL,
  `year` date NOT NULL,
  `overview` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `poster` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `backdrop` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `trailer` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `commentary` text COLLATE utf8mb4_unicode_ci,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `media_type` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Fazendo dump de dados para tabela `recommendation_items`
--

INSERT INTO `recommendation_items` (`id`, `recommendation_id`, `name`, `tmdb_id`, `year`, `overview`, `poster`, `backdrop`, `trailer`, `commentary`, `created_at`, `updated_at`, `media_type`) VALUES
(1, 1, 'The Grinch', 500, '0000-00-00', 'He hates Christmas!', 'OqAmzALlnkasQml', 'pKqnkmAqmlasas', '/iqnonAsnkas', '<p>Really funny movie!!</p>', '2019-01-03 01:13:16', '2019-01-04 23:04:52', 'movie'),
(35, 54, 'Man of Steel', 49521, '2013-06-12', 'A young boy learns that he has extraordinary powers and is not of this earth. As a young man, he journeys to discover where he came from and what he was sent here to do. But the hero in him must emerge if he is to save the world from annihilation and become the symbol of hope for all mankind.', '/xWlaTLnD8NJMTT9PGOD9z5re1SL.jpg', '/jYLh4mdOqkt30i7LTFs3o02UcGF.jpg', 'KVu3gS7iJu4', '<p>Hue</p>', '2019-01-05 00:49:24', '2019-01-05 01:26:00', 'movie'),
(4, 1, 'John Wick: Chapter II', 5232, '2017-12-24', 'I\'ll kill them all!', 'OqAmzALlnkasQml', 'pKqnkmAqmlasas', '/iqnonAsnkas', 'Really great movie!', '2019-01-03 04:16:17', '2019-01-03 04:16:17', 'movie'),
(5, 1, 'John Wick: Chapter II', 5232, '2017-12-24', 'I\'ll kill them all!', 'OqAmzALlnkasQml', 'pKqnkmAqmlasas', '/iqnonAsnkas', 'Really great movie!', '2019-01-03 04:16:36', '2019-01-03 04:20:37', 'movie'),
(6, 1, 'John Wick: Chapter II', 5232, '2017-12-24', 'I\'ll kill them all!', 'OqAmzALlnkasQml', 'pKqnkmAqmlasas', '/iqnonAsnkas', 'Really great movie!', '2019-01-03 20:53:40', '2019-01-03 20:53:40', 'movie'),
(7, 1, 'John Wick: Chapter II', 5232, '2017-12-24', 'I\'ll kill them all!', 'OqAmzALlnkasQml', 'pKqnkmAqmlasas', '/iqnonAsnkas', 'Really great movie!', '2019-01-03 20:54:00', '2019-01-03 20:54:00', 'movie'),
(8, 1, 'John Wick: Chapter II', 5232, '2017-12-24', 'I\'ll kill them all!', 'OqAmzALlnkasQml', 'pKqnkmAqmlasas', '/iqnonAsnkas', 'Really great movie!', '2019-01-03 20:54:23', '2019-01-03 20:54:23', 'movie'),
(9, 1, 'John Wick: Chapter II', 5232, '2017-12-24', 'I\'ll kill them all!', 'OqAmzALlnkasQml', 'pKqnkmAqmlasas', '/iqnonAsnkas', 'Really great movie!', '2019-01-03 20:54:28', '2019-01-03 20:54:28', 'movie'),
(10, 1, 'John Wick: Chapter II', 5232, '2017-12-24', 'I\'ll kill them all!', 'OqAmzALlnkasQml', 'pKqnkmAqmlasas', '/', 'Really great movie!', '2019-01-03 20:54:31', '2019-01-03 20:54:31', 'movie'),
(11, 1, 'John Wick: Chapter II', 5232, '2017-12-24', 'I\'ll kill them all!', 'OqAmzALlnkasQml', 'pKqnkmAqmlasas', '', 'Really great movie!', '2019-01-03 20:54:34', '2019-01-03 20:54:34', 'movie'),
(12, 1, 'John Wick: Chapter III', 5232, '2017-12-24', 'I\'ll kill them all!', 'OqAmzALlnkasQml', 'pKqnkmAqmlasas', '/iqnonAsnkas', 'Really great movie!', '2019-01-03 20:54:43', '2019-01-03 20:56:57', 'movie'),
(13, 1, 'John Wick: Chapter II', 5232, '2017-12-24', 'I\'ll kill them all!', 'OqAmzALlnkasQml', 'pKqnkmAqmlasas', '/iqnonAsnkas', 'Really great movie!', '2019-01-04 03:29:07', '2019-01-04 03:29:07', 'movie'),
(14, 1, 'John Wick: Chapter II', 5232, '2017-12-24', 'I\'ll kill them all!', 'OqAmzALlnkasQml', 'pKqnkmAqmlasas', '/iqnonAsnkas', 'Really great movie!', '2019-01-04 03:29:20', '2019-01-04 03:29:20', 'movie'),
(15, 1, 'John Wick: Chapter II', 5232, '2017-12-24', 'I\'ll kill them all!', 'OqAmzALlnkasQml', 'pKqnkmAqmlasas', '/iqnonAsnkas', 'Really great movie!', '2019-01-04 03:29:27', '2019-01-04 03:29:27', 'movie'),
(16, 1, 'John Wick: Chapter III', 5232, '2017-12-24', 'I\'ll kill them all!', 'OqAmzALlnkasQml', 'pKqnkmAqmlasas', '/iqnonAsnkas', 'Really great movie!', '2019-01-04 03:29:30', '2019-01-04 06:03:02', 'movie'),
(17, 1, 'John Wick: Chapter II', 5232, '2017-12-24', 'I\'ll kill them all!', 'OqAmzALlnkasQml', 'pKqnkmAqmlasas', '/iqnonAsnkas', 'Really great movie!', '2019-01-04 03:31:25', '2019-01-04 03:31:25', 'movie'),
(18, 1, 'John Wick: Chapter II', 5232, '2017-12-24', 'I\'ll kill them all!', 'OqAmzALlnkasQml', 'pKqnkmAqmlasas', '/iqnonAsnkas', 'Really great movie!', '2019-01-04 03:31:36', '2019-01-04 03:31:36', 'movie'),
(19, 1, 'John Wick: Chapter II', 5232, '2017-12-24', 'I\'ll kill them all!', 'OqAmzALlnkasQml', 'pKqnkmAqmlasas', '/iqnonAsnkas', 'Really great movie!', '2019-01-04 03:31:59', '2019-01-04 03:31:59', 'movie'),
(20, 1, 'John Wick: Chapter II', 5232, '2017-12-24', 'I\'ll kill them all!', 'OqAmzALlnkasQml', 'pKqnkmAqmlasas', '/iqnonAsnkas', 'Really great movie!', '2019-01-04 03:40:06', '2019-01-04 03:40:06', 'movie'),
(21, 1, 'John Wick: Chapter II', 5232, '2017-12-24', 'I\'ll kill them all!', 'OqAmzALlnkasQml', 'pKqnkmAqmlasas', '/iqnonAsnkas', 'Really great movie!', '2019-01-04 03:40:45', '2019-01-04 03:40:45', 'movie'),
(22, 1, 'John Wick: Chapter II', 5232, '2017-12-24', 'I\'ll kill them all!', 'OqAmzALlnkasQml', 'pKqnkmAqmlasas', '/iqnonAsnkas', 'Really great movie!', '2019-01-04 03:42:01', '2019-01-04 03:42:01', 'movie'),
(23, 1, 'John Wick: Chapter II', 5232, '2017-12-24', 'I\'ll kill them all!', 'OqAmzALlnkasQml', 'pKqnkmAqmlasas', '/iqnonAsnkas', 'Really great movie!', '2019-01-04 03:44:23', '2019-01-04 03:44:23', 'movie'),
(24, 1, 'John Wick: Chapter II', 5232, '2017-12-24', 'I\'ll kill them all!', 'OqAmzALlnkasQml', 'pKqnkmAqmlasas', '/iqnonAsnkas', 'Really great movie!', '2019-01-04 03:45:24', '2019-01-04 03:45:24', 'movie'),
(25, 1, 'John Wick: Chapter II', 5232, '2017-12-24', 'I\'ll kill them all!', 'OqAmzALlnkasQml', 'pKqnkmAqmlasas', '/iqnonAsnkas', 'Really great movie!', '2019-01-04 03:47:06', '2019-01-04 03:47:06', 'movie'),
(26, 1, 'John Wick: Chapter II', 5232, '2017-12-24', 'I\'ll kill them all!', 'OqAmzALlnkasQml', 'pKqnkmAqmlasas', '/iqnonAsnkas', 'Really great movie!', '2019-01-04 03:47:35', '2019-01-04 03:47:35', 'movie'),
(27, 1, 'John Wick: Chapter II', 5232, '2017-12-24', 'I\'ll kill them all!', 'OqAmzALlnkasQml', 'pKqnkmAqmlasas', '/iqnonAsnkas', 'Really great movie!', '2019-01-04 03:52:41', '2019-01-04 03:52:41', 'movie'),
(28, 1, 'John Wick: Chapter II', 5232, '2017-12-24', 'I\'ll kill them all!', 'OqAmzALlnkasQml', 'pKqnkmAqmlasas', '/iqnonAsnkas', 'Really great movie!', '2019-01-04 03:52:52', '2019-01-04 03:52:52', 'movie'),
(29, 1, 'John Wick: Chapter II', 5232, '2017-12-24', 'I\'ll kill them all!', 'OqAmzALlnkasQml', 'pKqnkmAqmlasas', '/iqnonAsnkas', 'Really great movie!', '2019-01-04 03:56:47', '2019-01-04 03:56:47', 'movie'),
(30, 1, 'John Wick: Chapter II', 5232, '2017-12-24', 'I\'ll kill them all!', 'OqAmzALlnkasQml', 'pKqnkmAqmlasas', '/iqnonAsnkas', 'Really great movie!', '2019-01-04 04:05:35', '2019-01-04 04:05:35', 'movie'),
(31, 1, 'John Wick: Chapter II', 5232, '2017-12-24', 'I\'ll kill them all!', 'OqAmzALlnkasQml', 'pKqnkmAqmlasas', '/iqnonAsnkas', 'Really great movie!', '2019-01-04 04:07:01', '2019-01-04 04:07:01', 'movie'),
(32, 1, 'John Wick: Chapter II', 5232, '2017-12-24', 'I\'ll kill them all!', 'OqAmzALlnkasQml', 'pKqnkmAqmlasas', '/iqnonAsnkas', 'Really great movie!', '2019-01-04 04:07:20', '2019-01-04 04:07:20', 'movie'),
(33, 1, 'John Wick: Chapter II', 5232, '2017-12-24', 'I\'ll kill them all!', 'OqAmzALlnkasQml', 'pKqnkmAqmlasas', '/iqnonAsnkas', 'Really great movie!', '2019-01-04 04:08:04', '2019-01-04 04:08:04', 'movie'),
(34, 1, 'John Wick: Chapter II', 5232, '2017-12-24', 'I\'ll kill them all!', 'OqAmzALlnkasQml', 'pKqnkmAqmlasas', '/iqnonAsnkas', 'Really great movie!', '2019-01-04 04:09:01', '2019-01-04 04:09:01', 'movie'),
(36, 54, 'Wonder Woman', 297762, '2017-05-30', 'An Amazon princess comes to the world of Man in the grips of the First World War to confront the forces of evil and bring an end to human conflict.', '/imekS7f1OuHyUP2LAiTEM0zBzUz.jpg', '/6iUNJZymJBMXXriQyFZfLAKnjO6.jpg', '1Q8fG0TtVAY', '<p>Muito bom!</p>', '2019-01-05 01:10:27', '2019-01-05 01:25:29', 'movie'),
(37, 54, 'Wonder Woman', 297762, '2017-05-30', 'An Amazon princess comes to the world of Man in the grips of the First World War to confront the forces of evil and bring an end to human conflict.', '/imekS7f1OuHyUP2LAiTEM0zBzUz.jpg', '/6iUNJZymJBMXXriQyFZfLAKnjO6.jpg', '1Q8fG0TtVAY', '<p>aa</p>', '2019-01-05 01:26:28', '2019-01-05 01:26:28', 'movie');

-- --------------------------------------------------------

--
-- Estrutura para tabela `recommendation_item_source`
--

CREATE TABLE `recommendation_item_source` (
  `recommendation_item_id` int(10) UNSIGNED NOT NULL,
  `source_id` int(10) UNSIGNED NOT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Fazendo dump de dados para tabela `recommendation_item_source`
--

INSERT INTO `recommendation_item_source` (`recommendation_item_id`, `source_id`) VALUES
(6, 4),
(6, 1),
(5, 3),
(6, 5),
(7, 1),
(7, 4),
(7, 5),
(8, 1),
(8, 4),
(8, 5),
(9, 1),
(9, 4),
(9, 5),
(10, 1),
(10, 4),
(10, 5),
(11, 1),
(11, 4),
(11, 5),
(13, 4),
(13, 1),
(12, 3),
(13, 5),
(14, 1),
(14, 4),
(14, 5),
(35, 1),
(16, 6),
(16, 4),
(18, 1),
(18, 4),
(18, 5),
(26, 1),
(27, 1),
(29, 1),
(33, 1),
(33, 2),
(34, 1),
(34, 2),
(35, 6),
(36, 8),
(36, 4),
(37, 4);

-- --------------------------------------------------------

--
-- Estrutura para tabela `sources`
--

CREATE TABLE `sources` (
  `id` int(10) UNSIGNED NOT NULL,
  `name` varchar(30) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Fazendo dump de dados para tabela `sources`
--

INSERT INTO `sources` (`id`, `name`, `created_at`, `updated_at`) VALUES
(1, 'Netflix', '2019-01-02 23:55:55', '2019-01-02 23:55:55'),
(5, 'HBO GO', '2019-01-03 04:13:33', '2019-01-03 04:13:33'),
(4, 'Prime Video', '2019-01-03 03:26:20', '2019-01-03 03:26:20'),
(6, 'YouTube', '2019-01-03 17:39:16', '2019-01-03 17:39:36'),
(7, 'NET Now', '2019-01-04 22:42:19', '2019-01-04 22:42:19'),
(8, 'iTunes', '2019-01-04 22:44:59', '2019-01-04 22:55:59'),
(9, 'Cinemax', '2019-01-04 22:46:02', '2019-01-04 22:46:02');

-- --------------------------------------------------------

--
-- Estrutura para tabela `users`
--

CREATE TABLE `users` (
  `id` int(10) UNSIGNED NOT NULL,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `email` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `password` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `api_token` varchar(200) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Fazendo dump de dados para tabela `users`
--

INSERT INTO `users` (`id`, `name`, `email`, `password`, `api_token`, `created_at`, `updated_at`) VALUES
(1, 'Cyro Dubeux', 'xorycx@gmail.com', '$2a$10$9tRbJ5BECRwpplS7xMhcqOACC.2JYucBaBEkf./P4khvl3G3elnua', '54B19FD4-EB1B-76BE-7D21-B8F5F4F3A521', '2019-01-02 00:00:00', '2019-01-02 22:34:11'),
(3, 'Paulo Braga', 'paulo_braga@gmail.com', '$2a$10$MRb8T41mO8aCtZzasWLDhupndLEbQ3RNFVPaes0eiw5wKV/wJuy52', '13CA6AC5-92C3-737B-CDA6-5285B381DE3C', '2019-01-02 22:59:45', '2019-01-02 22:59:45'),
(4, 'Vinícius Enéas!', 'vini.eneas@gmail.com', '$2a$10$BRAboTAYjR.WTRPDFCnN9.JjiL3JpYCb7R5Bmn9aloPzLH2oyif/O', '9590BCDF-7C81-F4B0-8DFC-E2E2A3DB947D', '2019-01-02 23:21:42', '2019-01-05 01:29:31'),
(20, 'Administrator', 'admin2@admin.com', '$2a$10$7P5ePXGmXJ7VOxXDmrDR.eHv./FvSkNV6AspjbH/Pz8eEPtSATLri', 'FF5140C0-AC88-D006-3B08-84F3ED83A41C', '2019-01-04 23:02:29', '2019-01-04 23:02:29'),
(13, 'Jams', 'jamesssassaa_rodruiguez@gmail.com', '$2a$10$Mi2dENJrMphEGt5Ura1xNuJl6eadCekxtZXbBCiYp.jirv8cn3Qnm', '81D3591C-9CCF-FAF9-57BC-17793907DE1E', '2019-01-03 05:21:06', '2019-01-03 05:21:06'),
(21, 'Ronaldo Fenômeno', 'ronaldo__fenomeno@gmail.com', '$2a$10$GATTX6UjkCS7AHT.uc4oceS4bEbNxt6sPEB9p9wz5f11uWDJYfhoW', '10750C91-7374-9D95-2FC9-AFF0F14DB3E2', '2019-01-04 23:06:40', '2019-01-04 23:07:22'),
(22, 'Débora Melo', 'debif@hotmail.com', '$2a$10$YmZ8v6MJSh.HNfqxAVWdl.D67afoX42hmQZP3uR8CS6rc21Avx9fC', '90F5D29E-8A6E-CBC4-DAE5-F6DD019A1C45', '2019-01-04 23:19:03', '2019-01-04 23:19:22'),
(16, 'Cristiano Ronaldo', 'cristiano.ronaldo@gmail.com', '$2a$10$IC4lRAk60yE3sj1Rf5BJvunFpnNQEhmA8fuW.TKAlxGFntxw1lvqq', 'E9822A0C-B901-2E20-7F41-090F401D06B2', '2019-01-03 16:04:36', '2019-01-03 16:06:45'),
(19, 'Admin', 'admin@admin.com', '$2a$10$KHsiHimoTwAgPw6H63UvK.1niWLEYg/avcuqmOBz1/ts5XoAZ0l/K', '057B1B0B-F0A2-9E61-E89D-A8DF6F7B5112', '2019-01-03 21:07:07', '2019-01-03 21:07:07');

--
-- Índices de tabelas apagadas
--

--
-- Índices de tabela `genres`
--
ALTER TABLE `genres`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `genres_name_unique` (`name`);

--
-- Índices de tabela `genre_recommendation`
--
ALTER TABLE `genre_recommendation`
  ADD KEY `genre_recommendation_recommendation_id_foreign` (`recommendation_id`),
  ADD KEY `genre_recommendation_genre_id_foreign` (`genre_id`);

--
-- Índices de tabela `keywords`
--
ALTER TABLE `keywords`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `keywords_name_unique` (`name`);

--
-- Índices de tabela `keyword_recommendation`
--
ALTER TABLE `keyword_recommendation`
  ADD KEY `keyword_recommendation_recommendation_id_foreign` (`recommendation_id`),
  ADD KEY `keyword_recommendation_keyword_id_foreign` (`keyword_id`);

--
-- Índices de tabela `migrations`
--
ALTER TABLE `migrations`
  ADD PRIMARY KEY (`id`);

--
-- Índices de tabela `recommendations`
--
ALTER TABLE `recommendations`
  ADD PRIMARY KEY (`id`),
  ADD KEY `recommendations_user_id_foreign` (`user_id`);

--
-- Índices de tabela `recommendation_items`
--
ALTER TABLE `recommendation_items`
  ADD PRIMARY KEY (`id`),
  ADD KEY `recommendation_items_recommendation_id_foreign` (`recommendation_id`);

--
-- Índices de tabela `recommendation_item_source`
--
ALTER TABLE `recommendation_item_source`
  ADD KEY `recommendation_item_source_recommendation_item_id_foreign` (`recommendation_item_id`),
  ADD KEY `recommendation_item_source_source_id_foreign` (`source_id`);

--
-- Índices de tabela `sources`
--
ALTER TABLE `sources`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `sources_name_unique` (`name`);

--
-- Índices de tabela `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `users_email_unique` (`email`),
  ADD UNIQUE KEY `users_api_token_unique` (`api_token`);

--
-- AUTO_INCREMENT de tabelas apagadas
--

--
-- AUTO_INCREMENT de tabela `genres`
--
ALTER TABLE `genres`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=18;

--
-- AUTO_INCREMENT de tabela `keywords`
--
ALTER TABLE `keywords`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;

--
-- AUTO_INCREMENT de tabela `migrations`
--
ALTER TABLE `migrations`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;

--
-- AUTO_INCREMENT de tabela `recommendations`
--
ALTER TABLE `recommendations`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=65;

--
-- AUTO_INCREMENT de tabela `recommendation_items`
--
ALTER TABLE `recommendation_items`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=38;

--
-- AUTO_INCREMENT de tabela `sources`
--
ALTER TABLE `sources`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;

--
-- AUTO_INCREMENT de tabela `users`
--
ALTER TABLE `users`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=24;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
