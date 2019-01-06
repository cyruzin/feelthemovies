-- MySQL dump 10.13  Distrib 5.7.24, for Linux (x86_64)
--
-- Host: 127.0.0.1    Database: api_feelthemovies
-- ------------------------------------------------------
-- Server version	5.7.24

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `genre_recommendation`
--

DROP TABLE IF EXISTS `genre_recommendation`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `genre_recommendation` (
  `recommendation_id` int(10) unsigned NOT NULL,
  `genre_id` int(10) unsigned NOT NULL,
  KEY `genre_recommendation_recommendation_id_foreign` (`recommendation_id`),
  KEY `genre_recommendation_genre_id_foreign` (`genre_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `genre_recommendation`
--

LOCK TABLES `genre_recommendation` WRITE;
/*!40000 ALTER TABLE `genre_recommendation` DISABLE KEYS */;
INSERT INTO `genre_recommendation` VALUES (1,1),(1,2),(2,1);
/*!40000 ALTER TABLE `genre_recommendation` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `genres`
--

DROP TABLE IF EXISTS `genres`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `genres` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(30) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `genres_name_unique` (`name`)
) ENGINE=MyISAM AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `genres`
--

LOCK TABLES `genres` WRITE;
/*!40000 ALTER TABLE `genres` DISABLE KEYS */;
INSERT INTO `genres` VALUES (1,'Horror','2019-01-05 22:34:32','2019-01-05 22:34:32'),(2,'Music','2019-01-05 22:34:37','2019-01-06 02:07:57'),(4,'Comedy','2019-01-05 22:35:10','2019-01-05 22:35:10'),(5,'War','2019-01-05 23:18:47','2019-01-05 23:18:47'),(6,'Bio','2019-01-05 23:30:39','2019-01-05 23:30:39'),(7,'Bioa','2019-01-05 23:37:45','2019-01-05 23:37:45'),(8,'Bioas','2019-01-05 23:38:22','2019-01-05 23:38:22'),(9,'World War II','2019-01-06 00:09:45','2019-01-06 00:09:45');
/*!40000 ALTER TABLE `genres` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `keyword_recommendation`
--

DROP TABLE IF EXISTS `keyword_recommendation`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `keyword_recommendation` (
  `recommendation_id` int(10) unsigned NOT NULL,
  `keyword_id` int(10) unsigned NOT NULL,
  KEY `keyword_recommendation_recommendation_id_foreign` (`recommendation_id`),
  KEY `keyword_recommendation_keyword_id_foreign` (`keyword_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `keyword_recommendation`
--

LOCK TABLES `keyword_recommendation` WRITE;
/*!40000 ALTER TABLE `keyword_recommendation` DISABLE KEYS */;
INSERT INTO `keyword_recommendation` VALUES (1,1),(1,2),(2,5);
/*!40000 ALTER TABLE `keyword_recommendation` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `keywords`
--

DROP TABLE IF EXISTS `keywords`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `keywords` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(30) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `keywords_name_unique` (`name`)
) ENGINE=MyISAM AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `keywords`
--

LOCK TABLES `keywords` WRITE;
/*!40000 ALTER TABLE `keywords` DISABLE KEYS */;
INSERT INTO `keywords` VALUES (1,'War','2019-01-06 02:28:03','2019-01-06 02:28:03'),(2,'Action','2019-01-06 02:28:07','2019-01-06 02:28:07');
/*!40000 ALTER TABLE `keywords` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `migrations`
--

DROP TABLE IF EXISTS `migrations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `migrations` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `migration` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `batch` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `migrations`
--

LOCK TABLES `migrations` WRITE;
/*!40000 ALTER TABLE `migrations` DISABLE KEYS */;
INSERT INTO `migrations` VALUES (1,'2018_09_30_013707_create_users_table',1),(2,'2018_09_30_013722_create_recommendations_table',1),(3,'2018_09_30_014443_create_recommendation_items_table',1),(4,'2018_10_01_160531_create_genres_table',1),(5,'2018_10_01_160552_create_genre_recommendation_table',1),(6,'2018_10_01_170339_create_keywords_table',1),(7,'2018_10_01_170402_create_keyword_recommendation_table',1),(8,'2018_10_01_233309_create_sources_table',1),(9,'2018_10_01_233540_create_recommendation_item_source_table',1),(10,'2018_10_31_235414_add_media_type_to_recommendation_items_table',2);
/*!40000 ALTER TABLE `migrations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `recommendation_item_source`
--

DROP TABLE IF EXISTS `recommendation_item_source`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `recommendation_item_source` (
  `recommendation_item_id` int(10) unsigned NOT NULL,
  `source_id` int(10) unsigned NOT NULL,
  KEY `recommendation_item_source_recommendation_item_id_foreign` (`recommendation_item_id`),
  KEY `recommendation_item_source_source_id_foreign` (`source_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `recommendation_item_source`
--

LOCK TABLES `recommendation_item_source` WRITE;
/*!40000 ALTER TABLE `recommendation_item_source` DISABLE KEYS */;
INSERT INTO `recommendation_item_source` VALUES (1,1),(1,2),(2,1),(2,2),(3,1),(3,2),(4,1),(4,2),(5,1),(5,2);
/*!40000 ALTER TABLE `recommendation_item_source` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `recommendation_items`
--

DROP TABLE IF EXISTS `recommendation_items`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `recommendation_items` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `recommendation_id` int(10) unsigned NOT NULL,
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
  `media_type` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`id`),
  KEY `recommendation_items_recommendation_id_foreign` (`recommendation_id`)
) ENGINE=MyISAM AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `recommendation_items`
--

LOCK TABLES `recommendation_items` WRITE;
/*!40000 ALTER TABLE `recommendation_items` DISABLE KEYS */;
INSERT INTO `recommendation_items` VALUES (1,1,'John Wick: Chapter II',5232,'2017-12-24','I\'ll kill them all!','OqAmzALlnkasQml','pKqnkmAqmlasas','/iqnonAsnkas','Really great movie!','2019-01-06 02:28:50','2019-01-06 02:28:50','movie'),(2,1,'Aquaman',5232,'2017-12-24','I\'ll kill them all!','OqAmzALlnkasQml','pKqnkmAqmlasas','/iqnonAsnkas','Really great movie!','2019-01-06 02:28:58','2019-01-06 02:28:58','movie'),(3,1,'Wonder Woman',5232,'2017-12-24','I\'ll kill them all!','OqAmzALlnkasQml','pKqnkmAqmlasas','/iqnonAsnkas','Really great movie!','2019-01-06 02:29:04','2019-01-06 02:29:04','movie'),(4,1,'Man of Steel',5232,'2017-12-24','I\'ll kill them all!','OqAmzALlnkasQml','pKqnkmAqmlasas','/iqnonAsnkas','Really great movie!','2019-01-06 02:29:11','2019-01-06 02:29:11','movie'),(5,2,'Aquamarine',5232,'2017-12-24','I\'ll kill them all!','OqAmzALlnkasQml','pKqnkmAqmlasas','/iqnonAsnkas','Really great movie!','2019-01-06 02:36:17','2019-01-06 02:36:17','movie');
/*!40000 ALTER TABLE `recommendation_items` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `recommendations`
--

DROP TABLE IF EXISTS `recommendations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `recommendations` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(10) unsigned NOT NULL,
  `title` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `type` tinyint(4) NOT NULL,
  `body` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `poster` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `backdrop` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `status` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `recommendations_user_id_foreign` (`user_id`)
) ENGINE=MyISAM AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `recommendations`
--

LOCK TABLES `recommendations` WRITE;
/*!40000 ALTER TABLE `recommendations` DISABLE KEYS */;
INSERT INTO `recommendations` VALUES (1,1,'The Best War Movies',0,'<p>Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.</p>','IjqIzzMninQmok','pOiiUyTZmlLznq',0,'2019-01-06 02:28:23','2019-01-06 02:28:23'),(2,1,'The Best Romance Movies',2,'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.','pOiiUyTZmlLznq','IjqIzzMninQmok',1,'2019-01-06 02:28:37','2019-01-06 02:36:23');
/*!40000 ALTER TABLE `recommendations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sources`
--

DROP TABLE IF EXISTS `sources`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sources` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(30) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `sources_name_unique` (`name`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sources`
--

LOCK TABLES `sources` WRITE;
/*!40000 ALTER TABLE `sources` DISABLE KEYS */;
/*!40000 ALTER TABLE `sources` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `email` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `password` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `api_token` varchar(200) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `users_email_unique` (`email`),
  UNIQUE KEY `users_api_token_unique` (`api_token`)
) ENGINE=MyISAM AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'Cyro Dubeux','xorycx@gmail.com','$2a$10$wxM3yhUhTgjs2MDmB7KA0uzXzliyHjtb6htrQlEx6Kjw3iDLagUjW','aa0f3579-c1ed-4377-867d-e5a11c45f76f','2019-01-02 00:00:00','2019-01-06 22:08:23'),(2,'John Doe','johndoe@admin.com','$2a$10$3KT68mNEtKHo4tI6hdXOQO/AbbVkH3iax2yGjM1ZN84iPKH7Noeca','bc409c91-1e0c-41e9-8837-eb04329aae75','2019-01-06 21:49:00','2019-01-06 21:49:00');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2019-01-06 19:13:16
