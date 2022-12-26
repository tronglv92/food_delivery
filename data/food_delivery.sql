-- MySQL dump 10.13  Distrib 8.0.31, for Linux (x86_64)
--
-- Host: localhost    Database: food_delivery
-- ------------------------------------------------------
-- Server version	8.0.31

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Current Database: `food_delivery`
--

CREATE DATABASE /*!32312 IF NOT EXISTS*/ `food_delivery` /*!40100 DEFAULT CHARACTER SET utf8mb3 */ /*!80016 DEFAULT ENCRYPTION='N' */;

USE `food_delivery`;

--
-- Table structure for table `carts`
--

DROP TABLE IF EXISTS `carts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `carts` (
  `user_id` int NOT NULL,
  `food_id` int NOT NULL,
  `quantity` int NOT NULL,
  `status` int NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`,`food_id`),
  KEY `food_id` (`food_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `carts`
--

LOCK TABLES `carts` WRITE;
/*!40000 ALTER TABLE `carts` DISABLE KEYS */;
/*!40000 ALTER TABLE `carts` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `categories`
--

DROP TABLE IF EXISTS `categories`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `categories` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `description` text,
  `icon` json DEFAULT NULL,
  `status` int NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `categories`
--

LOCK TABLES `categories` WRITE;
/*!40000 ALTER TABLE `categories` DISABLE KEYS */;
/*!40000 ALTER TABLE `categories` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `cities`
--

DROP TABLE IF EXISTS `cities`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `cities` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(100) NOT NULL,
  `status` int NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=64 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `cities`
--

LOCK TABLES `cities` WRITE;
/*!40000 ALTER TABLE `cities` DISABLE KEYS */;
/*!40000 ALTER TABLE `cities` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `food_likes`
--

DROP TABLE IF EXISTS `food_likes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `food_likes` (
  `user_id` int NOT NULL,
  `food_id` int NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`,`food_id`),
  KEY `food_id` (`food_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `food_likes`
--

LOCK TABLES `food_likes` WRITE;
/*!40000 ALTER TABLE `food_likes` DISABLE KEYS */;
/*!40000 ALTER TABLE `food_likes` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `food_ratings`
--

DROP TABLE IF EXISTS `food_ratings`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `food_ratings` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `food_id` int NOT NULL,
  `point` float DEFAULT '0',
  `comment` text,
  `status` int NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `food_id` (`food_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `food_ratings`
--

LOCK TABLES `food_ratings` WRITE;
/*!40000 ALTER TABLE `food_ratings` DISABLE KEYS */;
/*!40000 ALTER TABLE `food_ratings` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `foods`
--

DROP TABLE IF EXISTS `foods`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `foods` (
  `id` int NOT NULL AUTO_INCREMENT,
  `restaurant_id` int NOT NULL,
  `category_id` int DEFAULT NULL,
  `name` varchar(255) NOT NULL,
  `description` text,
  `price` float NOT NULL,
  `images` json NOT NULL,
  `status` int NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `restaurant_id` (`restaurant_id`) USING BTREE,
  KEY `category_id` (`category_id`) USING BTREE,
  KEY `status` (`status`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `foods`
--

LOCK TABLES `foods` WRITE;
/*!40000 ALTER TABLE `foods` DISABLE KEYS */;
/*!40000 ALTER TABLE `foods` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `order_details`
--

DROP TABLE IF EXISTS `order_details`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `order_details` (
  `id` int NOT NULL AUTO_INCREMENT,
  `order_id` int NOT NULL,
  `food_origin` json DEFAULT NULL,
  `price` float NOT NULL,
  `quantity` int NOT NULL,
  `discount` float DEFAULT '0',
  `status` int NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `order_id` (`order_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `order_details`
--

LOCK TABLES `order_details` WRITE;
/*!40000 ALTER TABLE `order_details` DISABLE KEYS */;
/*!40000 ALTER TABLE `order_details` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `order_trackings`
--

DROP TABLE IF EXISTS `order_trackings`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `order_trackings` (
  `id` int NOT NULL AUTO_INCREMENT,
  `order_id` int NOT NULL,
  `state` enum('waiting_for_shipper','preparing','on_the_way','delivered','cancel') NOT NULL,
  `status` int NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `order_id` (`order_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `order_trackings`
--

LOCK TABLES `order_trackings` WRITE;
/*!40000 ALTER TABLE `order_trackings` DISABLE KEYS */;
/*!40000 ALTER TABLE `order_trackings` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `orders`
--

DROP TABLE IF EXISTS `orders`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `orders` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `total_price` float NOT NULL,
  `shipper_id` int DEFAULT NULL,
  `status` int NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`) USING BTREE,
  KEY `shipper_id` (`shipper_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `orders`
--

LOCK TABLES `orders` WRITE;
/*!40000 ALTER TABLE `orders` DISABLE KEYS */;
/*!40000 ALTER TABLE `orders` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `restaurant_foods`
--

DROP TABLE IF EXISTS `restaurant_foods`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `restaurant_foods` (
  `restaurant_id` int NOT NULL,
  `food_id` int NOT NULL,
  `status` int NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`restaurant_id`,`food_id`),
  KEY `food_id` (`food_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `restaurant_foods`
--

LOCK TABLES `restaurant_foods` WRITE;
/*!40000 ALTER TABLE `restaurant_foods` DISABLE KEYS */;
/*!40000 ALTER TABLE `restaurant_foods` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `restaurant_likes`
--

DROP TABLE IF EXISTS `restaurant_likes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `restaurant_likes` (
  `restaurant_id` int NOT NULL,
  `user_id` int NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`restaurant_id`,`user_id`),
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `restaurant_likes`
--

LOCK TABLES `restaurant_likes` WRITE;
/*!40000 ALTER TABLE `restaurant_likes` DISABLE KEYS */;
/*!40000 ALTER TABLE `restaurant_likes` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `restaurant_ratings`
--

DROP TABLE IF EXISTS `restaurant_ratings`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `restaurant_ratings` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `restaurant_id` int NOT NULL,
  `point` float NOT NULL DEFAULT '0',
  `comment` text,
  `status` int NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`) USING BTREE,
  KEY `restaurant_id` (`restaurant_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `restaurant_ratings`
--

LOCK TABLES `restaurant_ratings` WRITE;
/*!40000 ALTER TABLE `restaurant_ratings` DISABLE KEYS */;
/*!40000 ALTER TABLE `restaurant_ratings` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `restaurants`
--

DROP TABLE IF EXISTS `restaurants`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `restaurants` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int DEFAULT NULL,
  `name` varchar(50) NOT NULL,
  `addr` varchar(255) NOT NULL,
  `city_id` int DEFAULT NULL,
  `lat` double DEFAULT NULL,
  `lng` double DEFAULT NULL,
  `cover` json DEFAULT NULL,
  `logo` json DEFAULT NULL,
  `shipping_fee_per_km` double DEFAULT '0',
  `liked_count` int DEFAULT '0',
  `status` int NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `owner_id` (`user_id`) USING BTREE,
  KEY `city_id` (`city_id`) USING BTREE,
  KEY `status` (`status`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `restaurants`
--

LOCK TABLES `restaurants` WRITE;
/*!40000 ALTER TABLE `restaurants` DISABLE KEYS */;
INSERT INTO `restaurants` VALUES (8,4,'test1','some where',NULL,NULL,NULL,NULL,NULL,1,0,0,'2022-11-14 15:26:45','2022-11-16 08:42:26'),(9,5,'Test restaurant 1','somewhere',NULL,NULL,NULL,NULL,NULL,2,0,0,'2022-11-16 10:12:09','2022-11-16 08:46:55'),(10,6,'Test restaurant 2','somewhere',NULL,NULL,NULL,NULL,NULL,3,0,1,'2022-11-16 10:12:12','2022-11-16 08:46:55'),(11,3,'Test restaurant 3','somewhere',NULL,NULL,NULL,NULL,NULL,4,0,1,'2022-11-16 10:12:14','2022-11-16 03:14:32'),(12,3,'Test restaurant 4','somewhere',NULL,NULL,NULL,NULL,NULL,5,0,1,'2022-11-16 10:12:17','2022-11-16 03:14:32'),(13,3,'Test restaurant 5','somewhere',NULL,NULL,NULL,NULL,NULL,6,0,1,'2022-11-16 10:12:19','2022-11-16 03:14:32'),(14,3,'Cua biển quy nhơn','somewhere',NULL,NULL,NULL,NULL,NULL,7,0,1,'2022-11-16 10:12:43','2022-11-16 03:14:32'),(15,3,'beef and beer','somewhere',NULL,NULL,NULL,NULL,NULL,8,0,1,'2022-11-16 10:13:23','2022-11-16 03:14:32'),(16,3,'mi tom chua cay','somewhere',NULL,NULL,NULL,NULL,NULL,9,0,1,'2022-11-16 10:13:41','2022-11-16 03:14:32'),(17,3,'susshi japan','somewhere',NULL,NULL,NULL,NULL,NULL,10,0,1,'2022-11-16 10:13:50','2022-11-16 03:14:32'),(18,3,'susshi japan','somewhere',NULL,NULL,NULL,NULL,NULL,0,0,1,'2022-11-16 11:50:36','2022-11-16 11:50:36');
/*!40000 ALTER TABLE `restaurants` ENABLE KEYS */;
UNLOCK TABLES;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb4 */ ;
/*!50003 SET character_set_results = utf8mb4 */ ;
/*!50003 SET collation_connection  = utf8mb4_0900_ai_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
/*!50003 CREATE*/ /*!50017 DEFINER=`root`@`%`*/ /*!50003 TRIGGER `restaurants_after_insert` AFTER INSERT ON `restaurants` FOR EACH ROW INSERT INTO restaurants_journal
	SET
		action_type = 'create',
		id = NEW.id,
		action_time = now() */;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb4 */ ;
/*!50003 SET character_set_results = utf8mb4 */ ;
/*!50003 SET collation_connection  = utf8mb4_0900_ai_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
/*!50003 CREATE*/ /*!50017 DEFINER=`root`@`%`*/ /*!50003 TRIGGER `restaurants_after_update` AFTER UPDATE ON `restaurants` FOR EACH ROW BEGIN
IF (NEW.id = OLD.id)
THEN
	INSERT INTO restaurants_journal
		SET action_type = 'update',
			id = OLD.id,
			action_time = NOW();
ELSE
		
		INSERT INTO restaurants_journal
		SET action_type = 'delete',
			id = OLD.id,
			action_time = NOW();
		
		INSERT INTO restaurants_journal
		SET action_type = 'create',
			id = NEW.id,
			action_time = NOW();
	END IF;
END */;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb4 */ ;
/*!50003 SET character_set_results = utf8mb4 */ ;
/*!50003 SET collation_connection  = utf8mb4_0900_ai_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
/*!50003 CREATE*/ /*!50017 DEFINER=`root`@`%`*/ /*!50003 TRIGGER `restaurants_after_delete` AFTER DELETE ON `restaurants` FOR EACH ROW INSERT INTO restaurants_journal
	SET action_type = 'delete',
		id = OLD.id,
		action_time = now() */;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;

--
-- Table structure for table `restaurants_journal`
--

DROP TABLE IF EXISTS `restaurants_journal`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `restaurants_journal` (
  `journal_id` int NOT NULL AUTO_INCREMENT,
  `id` varchar(15) DEFAULT NULL,
  `action_type` enum('create','update','delete') DEFAULT NULL,
  `action_time` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`journal_id`)
) ENGINE=InnoDB AUTO_INCREMENT=29 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `restaurants_journal`
--

LOCK TABLES `restaurants_journal` WRITE;
/*!40000 ALTER TABLE `restaurants_journal` DISABLE KEYS */;
/*!40000 ALTER TABLE `restaurants_journal` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sessions`
--

DROP TABLE IF EXISTS `sessions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sessions` (
  `id` varchar(255) NOT NULL,
  `user_id` int NOT NULL,
  `refresh_token` longtext NOT NULL,
  `user_agent` varchar(255) NOT NULL,
  `client_ip` varchar(255) NOT NULL,
  `is_blocked` tinyint(1) NOT NULL DEFAULT '0',
  `expires_at` timestamp NOT NULL,
  `status` int NOT NULL DEFAULT '1',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sessions`
--

LOCK TABLES `sessions` WRITE;
/*!40000 ALTER TABLE `sessions` DISABLE KEYS */;
INSERT INTO `sessions` VALUES ('004c6036-eb8a-4116-a1cd-18f42e2ab12c',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiIwMDRjNjAzNi1lYjhhLTQxMTYtYTFjZC0xOGY0MmUyYWIxMmMifSwiZXhwIjoxNjY5NjkxNDExLCJqdGkiOiIxNjY5NjkxMjMxNzQ2MDQ0MDAwIiwiaWF0IjoxNjY5NjkxMjMxfQ.OxX5iITnI9TP8iu8T24_GwBeeXn24siELKOAX3-qFCk','PostmanRuntime/7.28.4','::1',0,'2022-11-29 10:10:12',1,'2022-11-29 10:07:12','2022-11-29 10:07:12'),('054211b3-7326-423e-a621-9703abbd160a',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiIwNTQyMTFiMy03MzI2LTQyM2UtYTYyMS05NzAzYWJiZDE2MGEifSwiZXhwIjoxNjY5NTM3ODYyLCJqdGkiOiIxNjY5NTM3NjgyMDczNTYwMDAwIiwiaWF0IjoxNjY5NTM3NjgyfQ.lE1x4Wj68fdyYGp5MnrYo09PLiH04I99z93dx4HQZWY','PostmanRuntime/7.28.4','::1',0,'2022-11-27 15:31:02',1,'2022-11-27 15:28:02','2022-11-27 15:28:02'),('05d53261-fdcd-42b5-a5eb-1f5d8141b79f',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiIwNWQ1MzI2MS1mZGNkLTQyYjUtYTVlYi0xZjVkODE0MWI3OWYifSwiZXhwIjoxNjY5NjkxNDE2LCJqdGkiOiIxNjY5NjkxMjM2MDY2NTcyMDAwIiwiaWF0IjoxNjY5NjkxMjM2fQ.6-3Gxll7m8zXUpMZNkvOWHrarE2Kjt-Kkm-AqGJj7OI','PostmanRuntime/7.28.4','::1',0,'2022-11-29 10:10:16',1,'2022-11-29 10:07:16','2022-11-29 10:07:16'),('063bf0d6-9887-48a8-88b1-43ff046e0d90',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiIwNjNiZjBkNi05ODg3LTQ4YTgtODhiMS00M2ZmMDQ2ZTBkOTAifSwiZXhwIjoxNjY5NjQxOTMxLCJqdGkiOiIxNjY5NjQxNzUxMDE0MDUxMDAwIiwiaWF0IjoxNjY5NjQxNzUxfQ.5z7qjhzDAWNEDew4duEhI982T9rgw1kVbuigRDJWXJc','PostmanRuntime/7.28.4','::1',0,'2022-11-28 20:25:31',1,'2022-11-28 20:22:31','2022-11-28 20:22:31'),('072eb9a0-cece-46d1-a1b5-bf653302cc68',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiIwNzJlYjlhMC1jZWNlLTQ2ZDEtYTFiNS1iZjY1MzMwMmNjNjgifSwiZXhwIjoxNjY5NzM0OTMwLCJqdGkiOiIxNjY5NzM0NzUwNjY4MTc0MDAwIiwiaWF0IjoxNjY5NzM0NzUwfQ.GPv0mnjgv5j38PS09nNsjmfE3ky-tQJm-VDmil8JVqs','PostmanRuntime/7.28.4','::1',0,'2022-11-29 22:15:31',1,'2022-11-29 22:12:31','2022-11-29 22:12:31'),('0ac21329-fbcf-4559-9a0c-850d63cdeb52',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiIwYWMyMTMyOS1mYmNmLTQ1NTktOWEwYy04NTBkNjNjZGViNTIifSwiZXhwIjoxNjY5NjkwMzUyLCJqdGkiOiIxNjY5NjkwMTcyOTU5NTQzMDAwIiwiaWF0IjoxNjY5NjkwMTcyfQ.hHlFalcc89VosYgU1FynEL3x2nXhJiDgTsWBRCpniTE','PostmanRuntime/7.28.4','::1',0,'2022-11-29 09:52:33',1,'2022-11-29 09:49:33','2022-11-29 09:49:33'),('0b1fcbed-7768-45ea-8e8e-6457cac3e934',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiIwYjFmY2JlZC03NzY4LTQ1ZWEtOGU4ZS02NDU3Y2FjM2U5MzQifSwiZXhwIjoxNjY5NzMxNDI2LCJqdGkiOiIxNjY5NzMxMjQ2NTI4ODU4MDAwIiwiaWF0IjoxNjY5NzMxMjQ2fQ.gYRN1wQke66poHsjvZ-i_nyvKTbc-_KE_5BqBZSZuPA','PostmanRuntime/7.28.4','::1',0,'2022-11-29 21:17:07',1,'2022-11-29 21:14:07','2022-11-29 21:14:07'),('0bd132d9-0415-4e1d-bdfb-6d16899a757d',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiIwYmQxMzJkOS0wNDE1LTRlMWQtYmRmYi02ZDE2ODk5YTc1N2QifSwiZXhwIjoxNjY5NjkxNDE1LCJqdGkiOiIxNjY5NjkxMjM1NTM4NjM4MDAwIiwiaWF0IjoxNjY5NjkxMjM1fQ.OtZgw8Fp_TDff9PjSYHxGiNyvJB6Y8igtVaqRuvX24w','PostmanRuntime/7.28.4','::1',0,'2022-11-29 10:10:16',1,'2022-11-29 10:07:16','2022-11-29 10:07:16'),('0e16743b-74f5-4c0c-b45c-6635f7ede3fd',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiIwZTE2NzQzYi03NGY1LTRjMGMtYjQ1Yy02NjM1ZjdlZGUzZmQifSwiZXhwIjoxNjY5NjkwMzU2LCJqdGkiOiIxNjY5NjkwMTc2MTIxMzczMDAwIiwiaWF0IjoxNjY5NjkwMTc2fQ.1-zDR5IVupQ6nIy6hRdLAM0120YclG8NhU6XwK9y9Gg','PostmanRuntime/7.28.4','::1',0,'2022-11-29 09:52:36',1,'2022-11-29 09:49:36','2022-11-29 09:49:36'),('0e50ad53-0794-4db4-afc0-ce7da819ddaa',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiIwZTUwYWQ1My0wNzk0LTRkYjQtYWZjMC1jZTdkYTgxOWRkYWEifSwiZXhwIjoxNjY5NjkxMTMyLCJqdGkiOiIxNjY5NjkwOTUyNDMyOTAwMDAwIiwiaWF0IjoxNjY5NjkwOTUyfQ.1qECCTnFh38EqaSAPZlkEF9xu8zxBDgy4w_snV2fpkQ','PostmanRuntime/7.28.4','::1',0,'2022-11-29 10:05:32',1,'2022-11-29 10:02:32','2022-11-29 10:02:32'),('0f0c89bc-e03a-40ec-886c-9ddfab37ca8d',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiIwZjBjODliYy1lMDNhLTQwZWMtODg2Yy05ZGRmYWIzN2NhOGQifSwiZXhwIjoxNjY5NjkwMzUxLCJqdGkiOiIxNjY5NjkwMTcxNDM5NzAxMDAwIiwiaWF0IjoxNjY5NjkwMTcxfQ.jQzPcWzugCYErAp0km987oxvBxEA60WgfnizRlWPssk','PostmanRuntime/7.28.4','::1',0,'2022-11-29 09:52:31',1,'2022-11-29 09:49:31','2022-11-29 09:49:31'),('103dde69-4a1d-4cbc-95e0-09109c88116b',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiIxMDNkZGU2OS00YTFkLTRjYmMtOTVlMC0wOTEwOWM4ODExNmIifSwiZXhwIjoxNjY5NjkwMzU5LCJqdGkiOiIxNjY5NjkwMTc5MjU4NTk0MDAwIiwiaWF0IjoxNjY5NjkwMTc5fQ.aJx2_TpU2msGShl3-C2-LlMUoYpXeOF7qD-Sm2mW7kQ','PostmanRuntime/7.28.4','::1',0,'2022-11-29 09:52:39',1,'2022-11-29 09:49:39','2022-11-29 09:49:39'),('132793f9-6b7c-4df2-b5cd-c063409753bb',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiIxMzI3OTNmOS02YjdjLTRkZjItYjVjZC1jMDYzNDA5NzUzYmIifSwiZXhwIjoxNjY5NjkxNDIzLCJqdGkiOiIxNjY5NjkxMjQzNzQzNDM4MDAwIiwiaWF0IjoxNjY5NjkxMjQzfQ.x7bQ_tBqLjgFuFGYVDd4z5nWQQuGYoMTlp-COrf-pJE','PostmanRuntime/7.28.4','::1',0,'2022-11-29 10:10:24',1,'2022-11-29 10:07:24','2022-11-29 10:07:24'),('147f82de-ce1b-4782-be29-6a30ea3ec75a',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiIxNDdmODJkZS1jZTFiLTQ3ODItYmUyOS02YTMwZWEzZWM3NWEifSwiZXhwIjoxNjY5NjQ0MjY5LCJqdGkiOiIxNjY5NjQ0MDg5MjY1MjE3MDAwIiwiaWF0IjoxNjY5NjQ0MDg5fQ.Fb35s7BoUk1sJWYtGfRlWw7kWgoEIUrf0IBHCckhneM','PostmanRuntime/7.28.4','::1',0,'2022-11-28 21:04:29',1,'2022-11-28 21:01:29','2022-11-28 21:01:29'),('1482132b-fedf-4836-9ba5-b0be93421838',8,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjgsInJvbGUiOiJ1c2VyIiwiaWQiOiIxNDgyMTMyYi1mZWRmLTQ4MzYtOWJhNS1iMGJlOTM0MjE4MzgifSwiZXhwIjoxNjcxOTYwNDA0LCJqdGkiOiIxNjcxOTYwMjI0NzU5MjczMDAwIiwiaWF0IjoxNjcxOTYwMjI0fQ.5Mf1w8WHK0o7PFrMwCD_n7BW6ibMS3ukgSKVDuArzPw','PostmanRuntime/7.28.4','::1',0,'2022-12-25 16:26:45',1,'2022-12-25 16:23:45','2022-12-25 16:23:45'),('14b11580-f6d0-4ce2-bda3-15dbc5437543',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiIxNGIxMTU4MC1mNmQwLTRjZTItYmRhMy0xNWRiYzU0Mzc1NDMifSwiZXhwIjoxNjY5NDc2ODMzLCJqdGkiOiIxNjY5NDc2NjUzMzQzNzk0MDAwIiwiaWF0IjoxNjY5NDc2NjUzfQ.4L6MF_hv0IHA1GUFa7o8K9eDPhULoAsylMKHURDoLp8','PostmanRuntime/7.28.4','::1',0,'2022-11-26 22:33:53',1,'2022-11-26 22:30:53','2022-11-26 22:30:53'),('14f2684c-0d57-4539-8b0c-4ecb154241f4',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiIxNGYyNjg0Yy0wZDU3LTQ1MzktOGIwYy00ZWNiMTU0MjQxZjQifSwiZXhwIjoxNjY5NjQ0MjcyLCJqdGkiOiIxNjY5NjQ0MDkyMjkwODk4MDAwIiwiaWF0IjoxNjY5NjQ0MDkyfQ.beaVnFmL-wsE8Nezx5e2r0jqqp83WWM-gptuAIU5fe0','PostmanRuntime/7.28.4','::1',0,'2022-11-28 21:04:32',1,'2022-11-28 21:01:32','2022-11-28 21:01:32'),('15c563a9-20b1-421d-a6d7-5a9f97bb3f25',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiIxNWM1NjNhOS0yMGIxLTQyMWQtYTZkNy01YTlmOTdiYjNmMjUifSwiZXhwIjoxNjY5NzMzMjQxLCJqdGkiOiIxNjY5NzMzMDYxMTI3NTA3MDAwIiwiaWF0IjoxNjY5NzMzMDYxfQ.drjqmdAY5rYI_edWNM1eRPxV6LPmZz6yWOKbN5RXLPQ','PostmanRuntime/7.28.4','::1',0,'2022-11-29 21:47:21',1,'2022-11-29 21:44:21','2022-11-29 21:44:21'),('16b06736-fee5-47da-bd87-cdc68f253716',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiIxNmIwNjczNi1mZWU1LTQ3ZGEtYmQ4Ny1jZGM2OGYyNTM3MTYifSwiZXhwIjoxNjY5NjQxOTI5LCJqdGkiOiIxNjY5NjQxNzQ5NDc3NzIwMDAwIiwiaWF0IjoxNjY5NjQxNzQ5fQ.mLqx2cCgIvro8DmUUY6WDMVgCIdj8gTnHd5PiTM3Sg0','PostmanRuntime/7.28.4','::1',0,'2022-11-28 20:25:29',1,'2022-11-28 20:22:29','2022-11-28 20:22:29'),('17e2dbbc-7ff4-45f3-a950-22824771d561',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiIxN2UyZGJiYy03ZmY0LTQ1ZjMtYTk1MC0yMjgyNDc3MWQ1NjEifSwiZXhwIjoxNjY5NjkxMTM3LCJqdGkiOiIxNjY5NjkwOTU3MTUwNDcwMDAwIiwiaWF0IjoxNjY5NjkwOTU3fQ.wJEQVeRNbzP_0a6bxDRNnETWeqfnHuvYQAVp75lZrsg','PostmanRuntime/7.28.4','::1',0,'2022-11-29 10:05:37',1,'2022-11-29 10:02:37','2022-11-29 10:02:37'),('190c2b8d-10fb-4c35-8b5a-a783729fe00c',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiIxOTBjMmI4ZC0xMGZiLTRjMzUtOGI1YS1hNzgzNzI5ZmUwMGMifSwiZXhwIjoxNjY5NjkxMTM2LCJqdGkiOiIxNjY5NjkwOTU2NDg4MjcyMDAwIiwiaWF0IjoxNjY5NjkwOTU2fQ.gnXdWILVLO4UyeSz_nR0E_cmaLmuKnGiHbb32yZVomE','PostmanRuntime/7.28.4','::1',0,'2022-11-29 10:05:36',1,'2022-11-29 10:02:36','2022-11-29 10:02:36'),('194df51a-ce07-470c-8a13-c77b2d7a33a4',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiIxOTRkZjUxYS1jZTA3LTQ3MGMtOGExMy1jNzdiMmQ3YTMzYTQifSwiZXhwIjoxNjY5NjkxNDMzLCJqdGkiOiIxNjY5NjkxMjUzMDQ2NzcyMDAwIiwiaWF0IjoxNjY5NjkxMjUzfQ._8rrsYvSyPAI9qWdaHSHQ9TiHziU5Df4jqcQPLYNh44','PostmanRuntime/7.28.4','::1',0,'2022-11-29 10:10:33',1,'2022-11-29 10:07:33','2022-11-29 10:07:33'),('1a9ef244-8025-4194-bb70-e05d37ac3548',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiIxYTllZjI0NC04MDI1LTQxOTQtYmI3MC1lMDVkMzdhYzM1NDgifSwiZXhwIjoxNjY5NzA1MDM1LCJqdGkiOiIxNjY5NzA0ODU1NTY5Mzg3MDAwIiwiaWF0IjoxNjY5NzA0ODU1fQ.aZ-crUSkaAE2UgZpVibnnbcYkmj8FtykcmGmr9_uZek','PostmanRuntime/7.28.4','::1',0,'2022-11-29 13:57:16',1,'2022-11-29 13:54:16','2022-11-29 13:54:16'),('1cf7eee8-0895-4acd-924c-fce47637b1f8',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiIxY2Y3ZWVlOC0wODk1LTRhY2QtOTI0Yy1mY2U0NzYzN2IxZjgifSwiZXhwIjoxNjY5NjQ3MTEwLCJqdGkiOiIxNjY5NjQ2OTMwNjkyNDg0MDAwIiwiaWF0IjoxNjY5NjQ2OTMwfQ.dGvYPgls_OJYk9RaC1euf4U2YJ69ovD9rPQit915PAM','PostmanRuntime/7.28.4','::1',0,'2022-11-28 21:51:51',1,'2022-11-28 21:48:51','2022-11-28 21:48:51'),('1f39f736-cecc-4423-8698-523e29447839',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiIxZjM5ZjczNi1jZWNjLTQ0MjMtODY5OC01MjNlMjk0NDc4MzkifSwiZXhwIjoxNjY5NTU3MjIxLCJqdGkiOiIxNjY5NTU3MDQxNTUzMjkwMDAwIiwiaWF0IjoxNjY5NTU3MDQxfQ.3WYKVTntkeKIJpBC2rU5R-0x3tfS7Ccb4XRvlPzzOR0','PostmanRuntime/7.28.4','::1',0,'2022-11-27 20:53:42',1,'2022-11-27 20:50:42','2022-11-27 20:50:42'),('201d66d8-acf0-4a58-81c9-cc34a19ab371',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiIyMDFkNjZkOC1hY2YwLTRhNTgtODFjOS1jYzM0YTE5YWIzNzEifSwiZXhwIjoxNjY5NTU0NjQwLCJqdGkiOiIxNjY5NTU0NDYwMzA0MDA3MDAwIiwiaWF0IjoxNjY5NTU0NDYwfQ.LuklqvCL3JxRhbGR_jQgwW8Uk9uqEuHVOtn-Co-zBGA','PostmanRuntime/7.28.4','::1',0,'2022-11-27 20:10:40',1,'2022-11-27 20:07:40','2022-11-27 20:07:40'),('2191723f-9654-41cf-9d9a-54e0591ab6c5',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiIyMTkxNzIzZi05NjU0LTQxY2YtOWQ5YS01NGUwNTkxYWI2YzUifSwiZXhwIjoxNjY5NzM1Mjk2LCJqdGkiOiIxNjY5NzM1MTE2NzIwMTkyMDAwIiwiaWF0IjoxNjY5NzM1MTE2fQ.E7fLhqwNYrqAF2vqoi-E6ibdeIPjAe22Cs84xzoICXk','PostmanRuntime/7.28.4','::1',0,'2022-11-29 22:21:37',1,'2022-11-29 22:18:37','2022-11-29 22:18:37'),('231a1357-887f-476c-b4b4-38808e4a7a00',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiIyMzFhMTM1Ny04ODdmLTQ3NmMtYjRiNC0zODgwOGU0YTdhMDAifSwiZXhwIjoxNjY5NzMxMzEzLCJqdGkiOiIxNjY5NzMxMTMzNTAxMzA4MDAwIiwiaWF0IjoxNjY5NzMxMTMzfQ.0GBZ-Amc8TZB5XMQRTiALpn2mzInEK1g0A4qe2f8Pl8','PostmanRuntime/7.28.4','::1',0,'2022-11-29 21:15:14',1,'2022-11-29 21:12:14','2022-11-29 21:12:14'),('2461957b-f749-480c-a56a-3b39e8c8039c',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiIyNDYxOTU3Yi1mNzQ5LTQ4MGMtYTU2YS0zYjM5ZThjODAzOWMifSwiZXhwIjoxNjY5NjMxNDcxLCJqdGkiOiIxNjY5NjMxMjkxMzY4Njc1MDAwIiwiaWF0IjoxNjY5NjMxMjkxfQ._lYU0nI2uaEdLmadZTDTd-0hhR013vb7aNhd3Dfg0AE','PostmanRuntime/7.28.4','::1',0,'2022-11-28 17:31:11',1,'2022-11-28 17:28:11','2022-11-28 17:28:11'),('256dee7c-ca47-46f7-8c7a-425e74bd2a98',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiIyNTZkZWU3Yy1jYTQ3LTQ2ZjctOGM3YS00MjVlNzRiZDJhOTgifSwiZXhwIjoxNjY5NTU3MjI2LCJqdGkiOiIxNjY5NTU3MDQ2NDUxMzI3MDAwIiwiaWF0IjoxNjY5NTU3MDQ2fQ.qx9C6ePQF7WZVBCkGva-hcsO0UEHqZ3oHv6dNAH2Yok','PostmanRuntime/7.28.4','::1',0,'2022-11-27 20:53:46',1,'2022-11-27 20:50:46','2022-11-27 20:50:46'),('27f1f7b0-c2e3-4f9e-8943-c5f45cc3545f',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiIyN2YxZjdiMC1jMmUzLTRmOWUtODk0My1jNWY0NWNjMzU0NWYifSwiZXhwIjoxNjY5NjMxNDY4LCJqdGkiOiIxNjY5NjMxMjg4NDA2NTgzMDAwIiwiaWF0IjoxNjY5NjMxMjg4fQ.ew42Rrdsyjh9knfDzEgMABYZ-wchy09cnxFL-wxPAGQ','PostmanRuntime/7.28.4','::1',0,'2022-11-28 17:31:08',1,'2022-11-28 17:28:08','2022-11-28 17:28:08'),('2bab6072-a6f1-4de0-99ae-19677b47b0fb',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiIyYmFiNjA3Mi1hNmYxLTRkZTAtOTlhZS0xOTY3N2I0N2IwZmIifSwiZXhwIjoxNjY5NjQ0MjcxLCJqdGkiOiIxNjY5NjQ0MDkxNTk2Njc4MDAwIiwiaWF0IjoxNjY5NjQ0MDkxfQ.C4KjG2f8ISFOw5z7onVxCbmPSTk6jvfFc5Emj7ASi4o','PostmanRuntime/7.28.4','::1',0,'2022-11-28 21:04:32',1,'2022-11-28 21:01:32','2022-11-28 21:01:32'),('2bfbffae-8224-496b-ab48-5d2a0fcac499',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiIyYmZiZmZhZS04MjI0LTQ5NmItYWI0OC01ZDJhMGZjYWM0OTkifSwiZXhwIjoxNjY5NjkxMTM3LCJqdGkiOiIxNjY5NjkwOTU3ODkxOTc3MDAwIiwiaWF0IjoxNjY5NjkwOTU3fQ.7z0BDuCg_eeBeQMbaIDwLnH8IGFN_eQ22nLLZNKr2GA','PostmanRuntime/7.28.4','::1',0,'2022-11-29 10:05:38',1,'2022-11-29 10:02:38','2022-11-29 10:02:38'),('2ccb8671-3542-400f-9d2d-608a4f28bfa6',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiIyY2NiODY3MS0zNTQyLTQwMGYtOWQyZC02MDhhNGYyOGJmYTYifSwiZXhwIjoxNjY5NzM0NDI4LCJqdGkiOiIxNjY5NzM0MjQ4Njk5MDc3MDAwIiwiaWF0IjoxNjY5NzM0MjQ4fQ.ouNWnVi3xtx0AjgS0Uvf5PY6iDqN9or5eWAW0PQmiXQ','PostmanRuntime/7.28.4','::1',0,'2022-11-29 22:07:09',1,'2022-11-29 22:04:09','2022-11-29 22:04:09'),('2cdc8b5f-014b-47fe-8e25-28adcdc13848',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiIyY2RjOGI1Zi0wMTRiLTQ3ZmUtOGUyNS0yOGFkY2RjMTM4NDgifSwiZXhwIjoxNjY5NjkxNDMxLCJqdGkiOiIxNjY5NjkxMjUxODA0NDQ2MDAwIiwiaWF0IjoxNjY5NjkxMjUxfQ.2RbIYPU8tw5O91RwF7iJwhQE3FIimCc7LEDP0nf0Dc4','PostmanRuntime/7.28.4','::1',0,'2022-11-29 10:10:32',1,'2022-11-29 10:07:32','2022-11-29 10:07:32'),('356056a9-693f-46ee-8a52-23ba61e16d2f',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiIzNTYwNTZhOS02OTNmLTQ2ZWUtOGE1Mi0yM2JhNjFlMTZkMmYifSwiZXhwIjoxNjY5NzM2NDAzLCJqdGkiOiIxNjY5NzM2MjIzMTg0NjEzMDAwIiwiaWF0IjoxNjY5NzM2MjIzfQ.7OWNljq9Di1K94VF5N36va7sxbz7vyk5dzts94_mkJw','PostmanRuntime/7.28.4','::1',0,'2022-11-29 22:40:03',1,'2022-11-29 22:37:03','2022-11-29 22:37:03'),('372f6c08-6a68-4a74-9ed9-dca0cf6243ab',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiIzNzJmNmMwOC02YTY4LTRhNzQtOWVkOS1kY2EwY2Y2MjQzYWIifSwiZXhwIjoxNjY5NTM4Mjg5LCJqdGkiOiIxNjY5NTM4MTA5ODYxNDE0MDAwIiwiaWF0IjoxNjY5NTM4MTA5fQ.yVZfl6KJq3o0SRWdYoi1AzVI9Auox2_zlBaZYEMJsA4','PostmanRuntime/7.28.4','::1',0,'2022-11-27 15:38:10',1,'2022-11-27 15:35:10','2022-11-27 15:35:10'),('389c1e54-ab2b-4bd3-a9bf-49b95737d496',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiIzODljMWU1NC1hYjJiLTRiZDMtYTliZi00OWI5NTczN2Q0OTYifSwiZXhwIjoxNjY5NjkwMzU1LCJqdGkiOiIxNjY5NjkwMTc1NTIzMzY5MDAwIiwiaWF0IjoxNjY5NjkwMTc1fQ.vl7ue059Ovpr9sld_2InmZH4I_zC8h7yC9VjH809Uc8','PostmanRuntime/7.28.4','::1',0,'2022-11-29 09:52:36',1,'2022-11-29 09:49:36','2022-11-29 09:49:36'),('3a90609f-9cf1-4183-abbe-1316da36c0a1',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiIzYTkwNjA5Zi05Y2YxLTQxODMtYWJiZS0xMzE2ZGEzNmMwYTEifSwiZXhwIjoxNjY5NjkwMzU3LCJqdGkiOiIxNjY5NjkwMTc3MjczMzM2MDAwIiwiaWF0IjoxNjY5NjkwMTc3fQ.TLURgVyT83nTXeCs3L0fLFcKZ_IsXJRnFr92g__5-BI','PostmanRuntime/7.28.4','::1',0,'2022-11-29 09:52:37',1,'2022-11-29 09:49:37','2022-11-29 09:49:37'),('3bb41464-8726-43c6-8357-9c0f5fc9b8f7',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiIzYmI0MTQ2NC04NzI2LTQzYzYtODM1Ny05YzBmNWZjOWI4ZjcifSwiZXhwIjoxNjY5NjQxOTMwLCJqdGkiOiIxNjY5NjQxNzUwMjYxNjMwMDAwIiwiaWF0IjoxNjY5NjQxNzUwfQ.RX8egtbhptYldiqN6XwvxQvpkJlxkt0zu_01pCYMlm4','PostmanRuntime/7.28.4','::1',0,'2022-11-28 20:25:30',1,'2022-11-28 20:22:30','2022-11-28 20:22:30'),('40beec6a-ac44-40ef-a1f9-035f93504eb7',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiI0MGJlZWM2YS1hYzQ0LTQwZWYtYTFmOS0wMzVmOTM1MDRlYjcifSwiZXhwIjoxNjY5NjkwMzYyLCJqdGkiOiIxNjY5NjkwMTgyOTcwODY5MDAwIiwiaWF0IjoxNjY5NjkwMTgyfQ.7LME2Rr7CxXFraIsaFjrkn5I7UrnokoEfKPt_GpJIto','PostmanRuntime/7.28.4','::1',0,'2022-11-29 09:52:43',1,'2022-11-29 09:49:43','2022-11-29 09:49:43'),('41272021-d8ab-4a13-9a39-efbb8a9a6088',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiI0MTI3MjAyMS1kOGFiLTRhMTMtOWEzOS1lZmJiOGE5YTYwODgifSwiZXhwIjoxNjY5NzAzODQ3LCJqdGkiOiIxNjY5NzAzNjY3NjM1MzA3MDAwIiwiaWF0IjoxNjY5NzAzNjY3fQ.5vmBhDOCYBkZITmrx78F-ZR1BXtrH9KJY-ndT-Czmf4','PostmanRuntime/7.28.4','::1',0,'2022-11-29 13:37:28',1,'2022-11-29 13:34:28','2022-11-29 13:34:28'),('424dfc2c-1f4c-4ded-825f-8ca2aea86e8b',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiI0MjRkZmMyYy0xZjRjLTRkZWQtODI1Zi04Y2EyYWVhODZlOGIifSwiZXhwIjoxNjY5NDc3Njc1LCJqdGkiOiIxNjY5NDc3NDk1MTg1OTUzMDAwIiwiaWF0IjoxNjY5NDc3NDk1fQ.uqeTNxLyKiQMBqCRQ-LDWI4wFzovyha_hKtXqLE50K8','PostmanRuntime/7.28.4','::1',0,'2022-11-26 22:47:55',1,'2022-11-26 22:44:55','2022-11-26 22:44:55'),('491412ee-8746-479e-9259-cf871ed80d2c',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiI0OTE0MTJlZS04NzQ2LTQ3OWUtOTI1OS1jZjg3MWVkODBkMmMifSwiZXhwIjoxNjY5NDc5NjMzLCJqdGkiOiIxNjY5NDc5NDUzNjU0ODQwMDAwIiwiaWF0IjoxNjY5NDc5NDUzfQ.c0t4rMI4Qnq9oXgI04ltf2HdM9gpU3Rsaf3Ag-0vh1o','PostmanRuntime/7.28.4','::1',0,'2022-11-26 23:20:34',1,'2022-11-26 23:17:34','2022-11-26 23:17:34'),('495a25dc-becc-44f5-99d2-c623f22ebf19',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiI0OTVhMjVkYy1iZWNjLTQ0ZjUtOTlkMi1jNjIzZjIyZWJmMTkifSwiZXhwIjoxNjY5NTM2OTk5LCJqdGkiOiIxNjY5NTM2ODE5MTcxMzE4MDAwIiwiaWF0IjoxNjY5NTM2ODE5fQ.atQRuz4pWD_kwD6s2DLtUowxv_WJaX2YTVaF7gbPXQo','PostmanRuntime/7.28.4','::1',0,'2022-11-27 15:16:39',1,'2022-11-27 15:13:39','2022-11-27 15:13:39'),('4a957d9b-8192-48b6-bc04-4a9856a311a5',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiI0YTk1N2Q5Yi04MTkyLTQ4YjYtYmMwNC00YTk4NTZhMzExYTUifSwiZXhwIjoxNjY5NjkxNDEzLCJqdGkiOiIxNjY5NjkxMjMzMzIyODAyMDAwIiwiaWF0IjoxNjY5NjkxMjMzfQ.tzJj2HLjnYSeWfufFfnDVhklyRhdNwc-MyOw66aFhks','PostmanRuntime/7.28.4','::1',0,'2022-11-29 10:10:13',1,'2022-11-29 10:07:13','2022-11-29 10:07:13'),('4c223e19-81f9-4ee0-a914-5f2d80d01149',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiI0YzIyM2UxOS04MWY5LTRlZTAtYTkxNC01ZjJkODBkMDExNDkifSwiZXhwIjoxNjY5NjkxNDI4LCJqdGkiOiIxNjY5NjkxMjQ4MzY5NTI0MDAwIiwiaWF0IjoxNjY5NjkxMjQ4fQ.BWCsGw2ZQVRfNc85c2MrkBegBWVivxq9wPRVPS5SNxM','PostmanRuntime/7.28.4','::1',0,'2022-11-29 10:10:28',1,'2022-11-29 10:07:28','2022-11-29 10:07:28'),('4cf246ab-55fb-4393-8c2c-f3cccb7f66c9',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiI0Y2YyNDZhYi01NWZiLTQzOTMtOGMyYy1mM2NjY2I3ZjY2YzkifSwiZXhwIjoxNjY5NjkxNDI5LCJqdGkiOiIxNjY5NjkxMjQ5NTMxMzM0MDAwIiwiaWF0IjoxNjY5NjkxMjQ5fQ.zarFuDkBwNF20UvmUkGMlaVtO0bJVNFYVTYGtD20dRY','PostmanRuntime/7.28.4','::1',0,'2022-11-29 10:10:30',1,'2022-11-29 10:07:30','2022-11-29 10:07:30'),('4db0417e-6551-403f-9bff-9e2b69c40b65',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiI0ZGIwNDE3ZS02NTUxLTQwM2YtOWJmZi05ZTJiNjljNDBiNjUifSwiZXhwIjoxNjY5NjkwMzYxLCJqdGkiOiIxNjY5NjkwMTgxNTMzNjcwMDAwIiwiaWF0IjoxNjY5NjkwMTgxfQ.ys-GDSM0t_UNtXVqmCL2DMoojJVSymTpIBpEur2osPU','PostmanRuntime/7.28.4','::1',0,'2022-11-29 09:52:42',1,'2022-11-29 09:49:42','2022-11-29 09:49:42'),('4ead8417-c14d-4b50-af06-ff6393588570',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiI0ZWFkODQxNy1jMTRkLTRiNTAtYWYwNi1mZjYzOTM1ODg1NzAifSwiZXhwIjoxNjY5NzA0OTcxLCJqdGkiOiIxNjY5NzA0NzkxMDUyMjgyMDAwIiwiaWF0IjoxNjY5NzA0NzkxfQ.ExT2y4-vZFxuEOHMnyNk-oZJX2gZ4gPJIST8G4A1yG0','PostmanRuntime/7.28.4','::1',0,'2022-11-29 13:56:11',1,'2022-11-29 13:53:11','2022-11-29 13:53:11'),('4f7be194-adbd-45ed-a179-be01b4808695',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiI0ZjdiZTE5NC1hZGJkLTQ1ZWQtYTE3OS1iZTAxYjQ4MDg2OTUifSwiZXhwIjoxNjY5NjkwMzYzLCJqdGkiOiIxNjY5NjkwMTgzNTA4MDYxMDAwIiwiaWF0IjoxNjY5NjkwMTgzfQ.oxE-j3xNGZvmMmDP3-wCvjH3Xta9vetZYwCiBJg4fA0','PostmanRuntime/7.28.4','::1',0,'2022-11-29 09:52:44',1,'2022-11-29 09:49:44','2022-11-29 09:49:44'),('52c45e73-0837-4f0b-93e5-7ba0215e0124',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiI1MmM0NWU3My0wODM3LTRmMGItOTNlNS03YmEwMjE1ZTAxMjQifSwiZXhwIjoxNjY5NjkwMzUzLCJqdGkiOiIxNjY5NjkwMTczNTU4MDkyMDAwIiwiaWF0IjoxNjY5NjkwMTczfQ.KTEpELs2ChxZ2ExHWWnpZoK1TT8Z8N2pPvdZzRYToWk','PostmanRuntime/7.28.4','::1',0,'2022-11-29 09:52:34',1,'2022-11-29 09:49:34','2022-11-29 09:49:34'),('55e9f8f4-0d7f-4fcf-a7b2-ce6051d796c5',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiI1NWU5ZjhmNC0wZDdmLTRmY2YtYTdiMi1jZTYwNTFkNzk2YzUifSwiZXhwIjoxNjY5NzAzODQ3LCJqdGkiOiIxNjY5NzAzNjY3MDI1Mzc5MDAwIiwiaWF0IjoxNjY5NzAzNjY3fQ.SRTeegjw7TSzo94sdIKh-dpU1aEV1LnlQGSSEnR7YUg','PostmanRuntime/7.28.4','::1',0,'2022-11-29 13:37:27',1,'2022-11-29 13:34:27','2022-11-29 13:34:27'),('56cec3b1-f3ca-4604-bb37-f290b09babe7',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiI1NmNlYzNiMS1mM2NhLTQ2MDQtYmIzNy1mMjkwYjA5YmFiZTcifSwiZXhwIjoxNjY5NTU0NTA1LCJqdGkiOiIxNjY5NTU0MzI1OTMxMDg3MDAwIiwiaWF0IjoxNjY5NTU0MzI1fQ.CqoqtXhWPlZCPqGHcW_ZfLcbuXt5CNTk-jcpAT-sNSg','PostmanRuntime/7.28.4','::1',0,'2022-11-27 20:08:26',1,'2022-11-27 20:05:26','2022-11-27 20:05:26'),('56ef62f4-b922-43df-a595-6b5fae5ce4cc',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiI1NmVmNjJmNC1iOTIyLTQzZGYtYTU5NS02YjVmYWU1Y2U0Y2MifSwiZXhwIjoxNjY5NjkwMzU0LCJqdGkiOiIxNjY5NjkwMTc0MTg3MjgyMDAwIiwiaWF0IjoxNjY5NjkwMTc0fQ.3OKD9-g6ANbafAHqOTKWZDfc1i-Bq-8f2g4OfRupbbo','PostmanRuntime/7.28.4','::1',0,'2022-11-29 09:52:34',1,'2022-11-29 09:49:34','2022-11-29 09:49:34'),('577fa849-bac7-4fef-84bf-d358bdf171e9',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiI1NzdmYTg0OS1iYWM3LTRmZWYtODRiZi1kMzU4YmRmMTcxZTkifSwiZXhwIjoxNjY5NjQ0Mjg5LCJqdGkiOiIxNjY5NjQ0MTA5NjIzOTQwMDAwIiwiaWF0IjoxNjY5NjQ0MTA5fQ.I34A32RIgyPIb5Og9fRDWaGYq_cPeiWVmjCi0Zbid4c','PostmanRuntime/7.28.4','::1',0,'2022-11-28 21:04:50',1,'2022-11-28 21:01:50','2022-11-28 21:01:50'),('583167ee-8907-485e-ab9a-3b66085ea731',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiI1ODMxNjdlZS04OTA3LTQ4NWUtYWI5YS0zYjY2MDg1ZWE3MzEifSwiZXhwIjoxNjY5NjkwMzYwLCJqdGkiOiIxNjY5NjkwMTgwMTYzMDMyMDAwIiwiaWF0IjoxNjY5NjkwMTgwfQ.wNT4C_F4wEYTQKQCqPITzvvrw0TSsFJbZEBOa1ZmDp4','PostmanRuntime/7.28.4','::1',0,'2022-11-29 09:52:40',1,'2022-11-29 09:49:40','2022-11-29 09:49:40'),('59f41529-3d83-4bef-8fcc-11a89904b5df',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiI1OWY0MTUyOS0zZDgzLTRiZWYtOGZjYy0xMWE4OTkwNGI1ZGYifSwiZXhwIjoxNjY5NjQ0MjcwLCJqdGkiOiIxNjY5NjQ0MDkwOTU0MzU1MDAwIiwiaWF0IjoxNjY5NjQ0MDkwfQ.9FdPnQ53lCy7z-trXNAKObvHSBb0jjvobqiVhtxt4MM','PostmanRuntime/7.28.4','::1',0,'2022-11-28 21:04:31',1,'2022-11-28 21:01:31','2022-11-28 21:01:31'),('5bea2c24-2177-4560-97f4-260fdc54a98f',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiI1YmVhMmMyNC0yMTc3LTQ1NjAtOTdmNC0yNjBmZGM1NGE5OGYifSwiZXhwIjoxNjY5NjkwMzUyLCJqdGkiOiIxNjY5NjkwMTcyMjcyNDE1MDAwIiwiaWF0IjoxNjY5NjkwMTcyfQ.3eTGVEven-ZwebSo3TvS_zc5hYKugmhVreg5xYVjxCo','PostmanRuntime/7.28.4','::1',0,'2022-11-29 09:52:32',1,'2022-11-29 09:49:32','2022-11-29 09:49:32'),('616c39a0-e9af-46e5-bf0e-a755545914dc',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiI2MTZjMzlhMC1lOWFmLTQ2ZTUtYmYwZS1hNzU1NTQ1OTE0ZGMifSwiZXhwIjoxNjY5NjkxNDI3LCJqdGkiOiIxNjY5NjkxMjQ3NTgxNjI0MDAwIiwiaWF0IjoxNjY5NjkxMjQ3fQ.5qUjdmiKcSUfIc9951S3UIqlAR7sQjw4n_NfnYIvmcY','PostmanRuntime/7.28.4','::1',0,'2022-11-29 10:10:28',1,'2022-11-29 10:07:28','2022-11-29 10:07:28'),('62d0fe65-7c0f-43a3-a462-36f10e945c25',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiI2MmQwZmU2NS03YzBmLTQzYTMtYTQ2Mi0zNmYxMGU5NDVjMjUifSwiZXhwIjoxNjY5NzM0NDgxLCJqdGkiOiIxNjY5NzM0MzAxNTAyMjUyMDAwIiwiaWF0IjoxNjY5NzM0MzAxfQ.LMeVU8pxJW51FGe-ShBUWUBKxurZvloFDWd09pXbYcM','PostmanRuntime/7.28.4','::1',0,'2022-11-29 22:08:02',1,'2022-11-29 22:05:02','2022-11-29 22:05:02'),('67d5249a-dc8b-4e8b-8123-22b6ba3ed0ea',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiI2N2Q1MjQ5YS1kYzhiLTRlOGItODEyMy0yMmI2YmEzZWQwZWEifSwiZXhwIjoxNjY5NjkwMzYwLCJqdGkiOiIxNjY5NjkwMTgwODQzMzY3MDAwIiwiaWF0IjoxNjY5NjkwMTgwfQ.MnUpawUJQEtOem709Q0pIz_L_UVDYk_o6zRVomoUbCI','PostmanRuntime/7.28.4','::1',0,'2022-11-29 09:52:41',1,'2022-11-29 09:49:41','2022-11-29 09:49:41'),('6ad9904f-a090-4694-8cf4-ad44f9788278',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiI2YWQ5OTA0Zi1hMDkwLTQ2OTQtOGNmNC1hZDQ0Zjk3ODgyNzgifSwiZXhwIjoxNjY5NjkxMTM1LCJqdGkiOiIxNjY5NjkwOTU1MTA2NDg0MDAwIiwiaWF0IjoxNjY5NjkwOTU1fQ.BtOo2M_hShbLUHzlY96ls201OxhJxYOoZb6KFSRuidg','PostmanRuntime/7.28.4','::1',0,'2022-11-29 10:05:35',1,'2022-11-29 10:02:35','2022-11-29 10:02:35'),('6bc71686-5377-4f1f-b0ee-9898fb7c113a',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiI2YmM3MTY4Ni01Mzc3LTRmMWYtYjBlZS05ODk4ZmI3YzExM2EifSwiZXhwIjoxNjY5NjkwMzUwLCJqdGkiOiIxNjY5NjkwMTcwNTYzMDExMDAwIiwiaWF0IjoxNjY5NjkwMTcwfQ.CpKyxWiatYAfNgPibst-GawMPIQYnHp8JgHJb2FU6ak','PostmanRuntime/7.28.4','::1',0,'2022-11-29 09:52:31',1,'2022-11-29 09:49:31','2022-11-29 09:49:31'),('6d12c8f4-6006-420c-b702-e83ee4947774',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiI2ZDEyYzhmNC02MDA2LTQyMGMtYjcwMi1lODNlZTQ5NDc3NzQifSwiZXhwIjoxNjY5NjkwMzU3LCJqdGkiOiIxNjY5NjkwMTc3ODk3NDE2MDAwIiwiaWF0IjoxNjY5NjkwMTc3fQ.3M4EyRgh1WwIjbNUQ3DMWXOj_lcsxWbJhshlmmr2qJ0','PostmanRuntime/7.28.4','::1',0,'2022-11-29 09:52:38',1,'2022-11-29 09:49:38','2022-11-29 09:49:38'),('6ef2945f-e8ed-44b1-988b-cb7e5d18cf9d',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiI2ZWYyOTQ1Zi1lOGVkLTQ0YjEtOTg4Yi1jYjdlNWQxOGNmOWQifSwiZXhwIjoxNjY5NjkwMzU2LCJqdGkiOiIxNjY5NjkwMTc2NjI5NTU2MDAwIiwiaWF0IjoxNjY5NjkwMTc2fQ.6_RuUsI8ZrwxnlmOorAPLiGfDfo5rRDsBIPrQY_6aJ8','PostmanRuntime/7.28.4','::1',0,'2022-11-29 09:52:37',1,'2022-11-29 09:49:37','2022-11-29 09:49:37'),('7463d598-f633-4c5d-9156-d9397a862ca7',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiI3NDYzZDU5OC1mNjMzLTRjNWQtOTE1Ni1kOTM5N2E4NjJjYTcifSwiZXhwIjoxNjY5NjkxNDEyLCJqdGkiOiIxNjY5NjkxMjMyNjE4NjIzMDAwIiwiaWF0IjoxNjY5NjkxMjMyfQ.QvRXv1K5qFFxCJyi4FuXDW-4yTfuzyHQEilF32yR4zI','PostmanRuntime/7.28.4','::1',0,'2022-11-29 10:10:13',1,'2022-11-29 10:07:13','2022-11-29 10:07:13'),('75e755a7-7e3f-469f-90f7-c572b3d04572',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiI3NWU3NTVhNy03ZTNmLTQ2OWYtOTBmNy1jNTcyYjNkMDQ1NzIifSwiZXhwIjoxNjY5NDU5OTg4LCJqdGkiOiIxNjY5NDU5ODA4NjAxNDcxMDAwIiwiaWF0IjoxNjY5NDU5ODA4fQ.vjRjCRlnzbDPoIcLxUIvUMpqQwVa8Ucqz2hCGH-iW9g','PostmanRuntime/7.28.4','::1',0,'2022-11-26 17:53:09',1,'2022-11-26 17:50:09','2022-11-26 17:50:09'),('76d89c84-2386-4aa7-8b64-403996f636a7',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiI3NmQ4OWM4NC0yMzg2LTRhYTctOGI2NC00MDM5OTZmNjM2YTcifSwiZXhwIjoxNjY5NzM0NDA5LCJqdGkiOiIxNjY5NzM0MjI5OTM2Mzk1MDAwIiwiaWF0IjoxNjY5NzM0MjI5fQ.GfnObomPjuME1U7hqKFkHsDMqbksw21ADJi-MpkSsLw','PostmanRuntime/7.28.4','::1',0,'2022-11-29 22:06:50',1,'2022-11-29 22:03:50','2022-11-29 22:03:50'),('785fb645-bf61-4136-8155-c1cdf176d906',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiI3ODVmYjY0NS1iZjYxLTQxMzYtODE1NS1jMWNkZjE3NmQ5MDYifSwiZXhwIjoxNjY5NjkxMTM0LCJqdGkiOiIxNjY5NjkwOTU0Mzc4NDk5MDAwIiwiaWF0IjoxNjY5NjkwOTU0fQ.EQ7XsATZ9q7x98_hRj7RGby1wM_kXYnKTYzwwV_tsO4','PostmanRuntime/7.28.4','::1',0,'2022-11-29 10:05:34',1,'2022-11-29 10:02:34','2022-11-29 10:02:34'),('83145ef6-be5a-419a-9353-f95d3aaf6542',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiI4MzE0NWVmNi1iZTVhLTQxOWEtOTM1My1mOTVkM2FhZjY1NDIifSwiZXhwIjoxNjY5NzAzODQ1LCJqdGkiOiIxNjY5NzAzNjY1NTAwMzgzMDAwIiwiaWF0IjoxNjY5NzAzNjY1fQ.abTOo4gmoISOr5bTSDoSnsM6Kprt6rjoEVMETXHthdA','PostmanRuntime/7.28.4','::1',0,'2022-11-29 13:37:26',1,'2022-11-29 13:34:26','2022-11-29 13:34:26'),('84aa4468-3599-47c2-b21d-76565a0bbae1',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiI4NGFhNDQ2OC0zNTk5LTQ3YzItYjIxZC03NjU2NWEwYmJhZTEifSwiZXhwIjoxNjY5NjkwMzU0LCJqdGkiOiIxNjY5NjkwMTc0ODgzNzE3MDAwIiwiaWF0IjoxNjY5NjkwMTc0fQ.JO35no1h8Pl3zsE1UTUHz0fj_qFiO4taLoMgiY5JTBU','PostmanRuntime/7.28.4','::1',0,'2022-11-29 09:52:35',1,'2022-11-29 09:49:35','2022-11-29 09:49:35'),('89704a72-1d7c-42e1-a38f-10facf06ef80',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiI4OTcwNGE3Mi0xZDdjLTQyZTEtYTM4Zi0xMGZhY2YwNmVmODAifSwiZXhwIjoxNjY5NjQ0Mjg4LCJqdGkiOiIxNjY5NjQ0MTA4NjI5MjQ0MDAwIiwiaWF0IjoxNjY5NjQ0MTA4fQ.NNfWaeYHWL6ffsEUtaEo0Rb5-imCZxGGAyFXhKqV8NE','PostmanRuntime/7.28.4','::1',0,'2022-11-28 21:04:49',1,'2022-11-28 21:01:49','2022-11-28 21:01:49'),('8ec398fa-c5b5-4c7f-83e1-807c9db18bd3',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiI4ZWMzOThmYS1jNWI1LTRjN2YtODNlMS04MDdjOWRiMThiZDMifSwiZXhwIjoxNjY5NzA0MDk0LCJqdGkiOiIxNjY5NzAzOTE0MzgyODkzMDAwIiwiaWF0IjoxNjY5NzAzOTE0fQ.Om4GQf8YgcgvXnRnnz_W_zTr2090trurllSM-iMp37I','PostmanRuntime/7.28.4','::1',0,'2022-11-29 13:41:34',1,'2022-11-29 13:38:34','2022-11-29 13:38:34'),('91c54169-81e2-47f2-8183-f54bc1227866',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiI5MWM1NDE2OS04MWUyLTQ3ZjItODE4My1mNTRiYzEyMjc4NjYifSwiZXhwIjoxNjY5MzU0NTU2LCJqdGkiOiIxNjY5MzU0Mzc2NTIyOTg3MDAwIiwiaWF0IjoxNjY5MzU0Mzc2fQ.Xi_0oV3TCKsKpAj4H64-kTVwUWPwhelh8v-id_1nUIE','PostmanRuntime/7.28.4','::1',0,'2022-11-25 12:35:57',1,'2022-11-25 12:32:57','2022-11-25 12:32:57'),('91f66c40-f5ef-44b4-930c-b4a6d0527f68',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiI5MWY2NmM0MC1mNWVmLTQ0YjQtOTMwYy1iNGE2ZDA1MjdmNjgifSwiZXhwIjoxNjY5NjkxMTMxLCJqdGkiOiIxNjY5NjkwOTUxMDg2MzUyMDAwIiwiaWF0IjoxNjY5NjkwOTUxfQ.kXnMwd01xQ7lRNWyW9zkxYg3WjZd3elP6_-ZR_tF8KE','PostmanRuntime/7.28.4','::1',0,'2022-11-29 10:05:31',1,'2022-11-29 10:02:31','2022-11-29 10:02:31'),('9371ac97-6a2b-471b-8cc0-38402a755633',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiI5MzcxYWM5Ny02YTJiLTQ3MWItOGNjMC0zODQwMmE3NTU2MzMifSwiZXhwIjoxNjY5NjkwMzYyLCJqdGkiOiIxNjY5NjkwMTgyMjI1NjI1MDAwIiwiaWF0IjoxNjY5NjkwMTgyfQ.jjma23uiuHc0LSZ0abIJ9sEJWmv56gd2_rEiSrDKxj8','PostmanRuntime/7.28.4','::1',0,'2022-11-29 09:52:42',1,'2022-11-29 09:49:42','2022-11-29 09:49:42'),('9a910cf9-cb6b-4fa1-a8ed-14d55cd1591a',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiI5YTkxMGNmOS1jYjZiLTRmYTEtYThlZC0xNGQ1NWNkMTU5MWEifSwiZXhwIjoxNjY5NjQ0Mjg3LCJqdGkiOiIxNjY5NjQ0MTA3NTYzMjE3MDAwIiwiaWF0IjoxNjY5NjQ0MTA3fQ.8AlvMVmRSaD3Oi4anYn_VBdME3fbgWmk4c8a4cMHuoA','PostmanRuntime/7.28.4','::1',0,'2022-11-28 21:04:48',1,'2022-11-28 21:01:48','2022-11-28 21:01:48'),('9a928bf5-0a31-4f13-b155-462fba980a8e',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiI5YTkyOGJmNS0wYTMxLTRmMTMtYjE1NS00NjJmYmE5ODBhOGUifSwiZXhwIjoxNjY5NjQ0MjcwLCJqdGkiOiIxNjY5NjQ0MDkwMjMyMzU0MDAwIiwiaWF0IjoxNjY5NjQ0MDkwfQ.tkQBC1IGff8VckKEgj2aaWSXirRacdsjVItseTdfc3s','PostmanRuntime/7.28.4','::1',0,'2022-11-28 21:04:30',1,'2022-11-28 21:01:30','2022-11-28 21:01:30'),('9f8ff3ea-9386-4c3b-8e9d-8ab28aa112fc',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiI5ZjhmZjNlYS05Mzg2LTRjM2ItOGU5ZC04YWIyOGFhMTEyZmMifSwiZXhwIjoxNjY5NjkxNDI1LCJqdGkiOiIxNjY5NjkxMjQ1NTc4ODE3MDAwIiwiaWF0IjoxNjY5NjkxMjQ1fQ.soG1c5EjeA1Fk8z8V3pgLDduOtxNV5uyfhZMYnZF83Y','PostmanRuntime/7.28.4','::1',0,'2022-11-29 10:10:26',1,'2022-11-29 10:07:26','2022-11-29 10:07:26'),('a1c3f006-bead-4e46-8ac5-27ef1d271ec4',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiJhMWMzZjAwNi1iZWFkLTRlNDYtOGFjNS0yN2VmMWQyNzFlYzQifSwiZXhwIjoxNjY5NjkwNDYxLCJqdGkiOiIxNjY5NjkwMjgxNjM3ODY2MDAwIiwiaWF0IjoxNjY5NjkwMjgxfQ.sFdnH-Xb8koS_aXCk_dk5OrIuVnKzJjHzEjsJcPIEWQ','PostmanRuntime/7.28.4','::1',0,'2022-11-29 09:54:22',1,'2022-11-29 09:51:22','2022-11-29 09:51:22'),('a22a16ba-bf3c-4c76-87bf-eff8b50664af',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiJhMjJhMTZiYS1iZjNjLTRjNzYtODdiZi1lZmY4YjUwNjY0YWYifSwiZXhwIjoxNjY5NjMxNDY5LCJqdGkiOiIxNjY5NjMxMjg5MjU1NjU1MDAwIiwiaWF0IjoxNjY5NjMxMjg5fQ.imPncBpFCkyKtfelYC-8MOSbSi3y1jAVolBMqcC7CMA','PostmanRuntime/7.28.4','::1',0,'2022-11-28 17:31:09',1,'2022-11-28 17:28:09','2022-11-28 17:28:09'),('a494b477-67fd-4f5b-a4c1-dc133780605c',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiJhNDk0YjQ3Ny02N2ZkLTRmNWItYTRjMS1kYzEzMzc4MDYwNWMifSwiZXhwIjoxNjY5NzMxNzQxLCJqdGkiOiIxNjY5NzMxNTYxNTQzMTM0MDAwIiwiaWF0IjoxNjY5NzMxNTYxfQ.gGhtW6xFohSpJ2TkVDV0nBhO4vn28dgXk1-qm01ElwU','PostmanRuntime/7.28.4','::1',0,'2022-11-29 21:22:22',1,'2022-11-29 21:19:22','2022-11-29 21:19:22'),('a51e2860-2d47-458a-8d5e-4f69ed13e488',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiJhNTFlMjg2MC0yZDQ3LTQ1OGEtOGQ1ZS00ZjY5ZWQxM2U0ODgifSwiZXhwIjoxNjY5NjMxNDY5LCJqdGkiOiIxNjY5NjMxMjg5OTU4OTQyMDAwIiwiaWF0IjoxNjY5NjMxMjg5fQ.uRNoBcQzB2goQ4O-ee0JfnPT-RUuApUDn6Uq4GTh11Q','PostmanRuntime/7.28.4','::1',0,'2022-11-28 17:31:10',1,'2022-11-28 17:28:10','2022-11-28 17:28:10'),('a5738472-f45e-4bd8-94fa-735be81ae013',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiJhNTczODQ3Mi1mNDVlLTRiZDgtOTRmYS03MzViZTgxYWUwMTMifSwiZXhwIjoxNjY5NjkxNDI0LCJqdGkiOiIxNjY5NjkxMjQ0NzI2MDgwMDAwIiwiaWF0IjoxNjY5NjkxMjQ0fQ.EVtXRcoBu9SzvXKDqzdYcXxFEhiZdskHMGfFSKD7MCU','PostmanRuntime/7.28.4','::1',0,'2022-11-29 10:10:25',1,'2022-11-29 10:07:25','2022-11-29 10:07:25'),('a59cca26-e013-4e66-82c0-6cec46284a74',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiJhNTljY2EyNi1lMDEzLTRlNjYtODJjMC02Y2VjNDYyODRhNzQifSwiZXhwIjoxNjY5NjkxNDMxLCJqdGkiOiIxNjY5NjkxMjUxMDc5NTMzMDAwIiwiaWF0IjoxNjY5NjkxMjUxfQ.DXpqp9P7S1RCv3MSKrFTgNrMIgoeGlUdcaJIzXDjQSw','PostmanRuntime/7.28.4','::1',0,'2022-11-29 10:10:31',1,'2022-11-29 10:07:31','2022-11-29 10:07:31'),('aab565d4-f4c9-45cf-bb1e-ddab30b632ee',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiJhYWI1NjVkNC1mNGM5LTQ1Y2YtYmIxZS1kZGFiMzBiNjMyZWUifSwiZXhwIjoxNjY5NjkxNDE0LCJqdGkiOiIxNjY5NjkxMjM0MTI4NDkzMDAwIiwiaWF0IjoxNjY5NjkxMjM0fQ.pnEjcezwnKF_Fy3iP5_fY1iqAvxJ8grHPGuWAdrL7fY','PostmanRuntime/7.28.4','::1',0,'2022-11-29 10:10:14',1,'2022-11-29 10:07:14','2022-11-29 10:07:14'),('ab5e1f03-2794-442e-86e7-a4b9264ef5f2',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiJhYjVlMWYwMy0yNzk0LTQ0MmUtODZlNy1hNGI5MjY0ZWY1ZjIifSwiZXhwIjoxNjY5NjkxNDMyLCJqdGkiOiIxNjY5NjkxMjUyNDUyNjg4MDAwIiwiaWF0IjoxNjY5NjkxMjUyfQ.qnAsNvqOiLuOC-vyrP429VnPpOjT1Efvl_zXXOZAMBE','PostmanRuntime/7.28.4','::1',0,'2022-11-29 10:10:32',1,'2022-11-29 10:07:32','2022-11-29 10:07:32'),('ac07d192-5c9f-4a3a-ace0-0e41e06eddad',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiJhYzA3ZDE5Mi01YzlmLTRhM2EtYWNlMC0wZTQxZTA2ZWRkYWQifSwiZXhwIjoxNjY5NzM1MTIxLCJqdGkiOiIxNjY5NzM0OTQxNjEwODk3MDAwIiwiaWF0IjoxNjY5NzM0OTQxfQ.iBfhGT9iaLSMe9UxNvgWz66BdOFbooH75wGcy55n6Ww','PostmanRuntime/7.28.4','::1',0,'2022-11-29 22:18:42',1,'2022-11-29 22:15:42','2022-11-29 22:15:42'),('adbd728a-bb73-4a17-ab0d-5756959043fe',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiJhZGJkNzI4YS1iYjczLTRhMTctYWIwZC01NzU2OTU5MDQzZmUifSwiZXhwIjoxNjY5NjMxNDY2LCJqdGkiOiIxNjY5NjMxMjg2ODY5MDI1MDAwIiwiaWF0IjoxNjY5NjMxMjg2fQ.4s_irjVmuHgV7LpennxCiD_1nqxd5wIBCtXPUj3LvBQ','PostmanRuntime/7.28.4','::1',0,'2022-11-28 17:31:07',1,'2022-11-28 17:28:07','2022-11-28 17:28:07'),('adc3ffaa-7a3c-4216-b4b1-7edc72a2f80e',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiJhZGMzZmZhYS03YTNjLTQyMTYtYjRiMS03ZWRjNzJhMmY4MGUifSwiZXhwIjoxNjY5NzA3ODk1LCJqdGkiOiIxNjY5NzA3NzE1MTQ4ODE5MDAwIiwiaWF0IjoxNjY5NzA3NzE1fQ.hNx4Ijzf8ssUoHE-CXDbT1uHKjeHCRKG5ktqfgv6T5Q','PostmanRuntime/7.28.4','::1',0,'2022-11-29 14:44:55',1,'2022-11-29 14:41:55','2022-11-29 14:41:55'),('b264548b-b621-4e8f-b184-45d33e6905c5',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiJiMjY0NTQ4Yi1iNjIxLTRlOGYtYjE4NC00NWQzM2U2OTA1YzUifSwiZXhwIjoxNjY5NjkwMzY0LCJqdGkiOiIxNjY5NjkwMTg0MDMwMjcxMDAwIiwiaWF0IjoxNjY5NjkwMTg0fQ.1aHNoN53EDkM5mOQrajo3i3i0YdQBmJ69sFVnTOdfrY','PostmanRuntime/7.28.4','::1',0,'2022-11-29 09:52:44',1,'2022-11-29 09:49:44','2022-11-29 09:49:44'),('b816479e-b071-40d3-94cb-28666dea84e0',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiJiODE2NDc5ZS1iMDcxLTQwZDMtOTRjYi0yODY2NmRlYTg0ZTAifSwiZXhwIjoxNjY5NzM0NDU0LCJqdGkiOiIxNjY5NzM0Mjc0NTk2NzQ0MDAwIiwiaWF0IjoxNjY5NzM0Mjc0fQ.gGAEdwIwvlRWrmJcZMjbrr6TfOPub3sxce0zc_xiQS4','PostmanRuntime/7.28.4','::1',0,'2022-11-29 22:07:35',1,'2022-11-29 22:04:35','2022-11-29 22:04:35'),('ba12ae42-79db-4f1e-b690-a02d0a8b89d2',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiJiYTEyYWU0Mi03OWRiLTRmMWUtYjY5MC1hMDJkMGE4Yjg5ZDIifSwiZXhwIjoxNjY5NjQ0Mjg3LCJqdGkiOiIxNjY5NjQ0MTA3MDIwMzUyMDAwIiwiaWF0IjoxNjY5NjQ0MTA3fQ.myiN7Be89XX0kaKLV93aV2LUWbKooex74RFxxKljBJY','PostmanRuntime/7.28.4','::1',0,'2022-11-28 21:04:47',1,'2022-11-28 21:01:47','2022-11-28 21:01:47'),('bb3dde22-d14a-43a3-9578-da3551da43d7',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiJiYjNkZGUyMi1kMTRhLTQzYTMtOTU3OC1kYTM1NTFkYTQzZDcifSwiZXhwIjoxNjY5NDc3NDA1LCJqdGkiOiIxNjY5NDc3MjI1MzA1MjUxMDAwIiwiaWF0IjoxNjY5NDc3MjI1fQ.11G5Edej5YkPzKnqrveIIoB-sOgng21v9DcS5kwkb8U','PostmanRuntime/7.28.4','::1',0,'2022-11-26 22:43:25',1,'2022-11-26 22:40:25','2022-11-26 22:40:25'),('c0566b2d-5206-40ff-8474-ad338f40f6b0',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiJjMDU2NmIyZC01MjA2LTQwZmYtODQ3NC1hZDMzOGY0MGY2YjAifSwiZXhwIjoxNjY5NjQ0Mjg4LCJqdGkiOiIxNjY5NjQ0MTA4MDc4NTA2MDAwIiwiaWF0IjoxNjY5NjQ0MTA4fQ.vjVzcG_fdcrQQiorDMxhJIVvLWnm5qGcheqzMyRm44g','PostmanRuntime/7.28.4','::1',0,'2022-11-28 21:04:48',1,'2022-11-28 21:01:48','2022-11-28 21:01:48'),('c0bd6768-bbb8-4c25-98a5-5a6de7d089b9',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiJjMGJkNjc2OC1iYmI4LTRjMjUtOThhNS01YTZkZTdkMDg5YjkifSwiZXhwIjoxNjY5NDc3MzM4LCJqdGkiOiIxNjY5NDc3MTU4ODI3MDc4MDAwIiwiaWF0IjoxNjY5NDc3MTU4fQ.Lzqg5HNRJx2lWsv62gBmDCUGU57ocj278jGvSaHXSdA','PostmanRuntime/7.28.4','::1',0,'2022-11-26 22:42:19',1,'2022-11-26 22:39:19','2022-11-26 22:39:19'),('c1612a0c-bf99-47ef-8d96-610427bf5730',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiJjMTYxMmEwYy1iZjk5LTQ3ZWYtOGQ5Ni02MTA0MjdiZjU3MzAifSwiZXhwIjoxNjY5NDc3OTA2LCJqdGkiOiIxNjY5NDc3NzI2MzA3NDgwMDAwIiwiaWF0IjoxNjY5NDc3NzI2fQ.XkemuTwSjP2Mk9kwKfEG7If7k_4-lu2wi60K2zMixEU','PostmanRuntime/7.28.4','::1',0,'2022-11-26 22:51:46',1,'2022-11-26 22:48:46','2022-11-26 22:48:46'),('c5ab5d81-d255-4e6e-90c5-5ef9c1278d17',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiJjNWFiNWQ4MS1kMjU1LTRlNmUtOTBjNS01ZWY5YzEyNzhkMTcifSwiZXhwIjoxNjY5NjQ0Mjg5LCJqdGkiOiIxNjY5NjQ0MTA5MjQ0NDA0MDAwIiwiaWF0IjoxNjY5NjQ0MTA5fQ.q083CWKPuyxS9BA0mF4HCesr-umwhmEHlFmK2730a3E','PostmanRuntime/7.28.4','::1',0,'2022-11-28 21:04:49',1,'2022-11-28 21:01:49','2022-11-28 21:01:49'),('c6945ee7-ab6b-4443-832a-3ced554f91d8',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiJjNjk0NWVlNy1hYjZiLTQ0NDMtODMyYS0zY2VkNTU0ZjkxZDgifSwiZXhwIjoxNjY5NTU0MjQyLCJqdGkiOiIxNjY5NTU0MDYyNDU0MjAzMDAwIiwiaWF0IjoxNjY5NTU0MDYyfQ.ET8cY3gMMyyrj_QfVAoI3-mHjNI0ce2vN-4HC1JYptA','PostmanRuntime/7.28.4','::1',0,'2022-11-27 20:04:02',1,'2022-11-27 20:01:02','2022-11-27 20:01:02'),('c776cdef-daa6-4b32-81f2-803798a4518c',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiJjNzc2Y2RlZi1kYWE2LTRiMzItODFmMi04MDM3OThhNDUxOGMifSwiZXhwIjoxNjY5NzAzMzk3LCJqdGkiOiIxNjY5NzAzMjE3ODkxNzE0MDAwIiwiaWF0IjoxNjY5NzAzMjE3fQ.s7WommYM4Zo1iDZlCUmzKMtWm5oCzYkKDoug4cug_Ls','PostmanRuntime/7.28.4','::1',0,'2022-11-29 13:29:58',1,'2022-11-29 13:26:58','2022-11-29 13:26:58'),('c8ab924d-03dc-450c-b43a-57c214ab4718',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiJjOGFiOTI0ZC0wM2RjLTQ1MGMtYjQzYS01N2MyMTRhYjQ3MTgifSwiZXhwIjoxNjY5NDc4NjAyLCJqdGkiOiIxNjY5NDc4NDIyNTQ2MTE0MDAwIiwiaWF0IjoxNjY5NDc4NDIyfQ.S50gObgi2IUyTpBCvrANrmZj4PKNEa4Im_-l8vPNAYU','PostmanRuntime/7.28.4','::1',0,'2022-11-26 23:03:23',1,'2022-11-26 23:00:23','2022-11-26 23:00:23'),('c8f0c088-d3a1-49a2-b4bd-bb61f47f4f57',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiJjOGYwYzA4OC1kM2ExLTQ5YTItYjRiZC1iYjYxZjQ3ZjRmNTcifSwiZXhwIjoxNjY5NzM1NTk0LCJqdGkiOiIxNjY5NzM1NDE0MDAwNDYxMDAwIiwiaWF0IjoxNjY5NzM1NDE0fQ.80fma5_I8kZS7SXVpOBoCdNI42vmH8uC7dxXLlFpweM','PostmanRuntime/7.28.4','::1',0,'2022-11-29 22:26:34',1,'2022-11-29 22:23:34','2022-11-29 22:23:34'),('cd43c7fc-aa7e-4648-9df4-c3d44eaf430c',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiJjZDQzYzdmYy1hYTdlLTQ2NDgtOWRmNC1jM2Q0NGVhZjQzMGMifSwiZXhwIjoxNjY5NTU3MjEyLCJqdGkiOiIxNjY5NTU3MDMyNDIxOTc0MDAwIiwiaWF0IjoxNjY5NTU3MDMyfQ.k13oQVYtMP7nsE9YkW5xUABzRrUm0npMiOnMEiaSQrQ','PostmanRuntime/7.28.4','::1',0,'2022-11-27 20:53:32',1,'2022-11-27 20:50:32','2022-11-27 20:50:32'),('ce2e456d-5de1-4bb2-a5db-24a8bf86b5dc',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiJjZTJlNDU2ZC01ZGUxLTRiYjItYTVkYi0yNGE4YmY4NmI1ZGMifSwiZXhwIjoxNjY5NjMxNDcwLCJqdGkiOiIxNjY5NjMxMjkwNzI2Nzg3MDAwIiwiaWF0IjoxNjY5NjMxMjkwfQ.DYO1MoIDSp6ogryG5agYeXpB3aksbskUeZQrW_7Gyko','PostmanRuntime/7.28.4','::1',0,'2022-11-28 17:31:11',1,'2022-11-28 17:28:11','2022-11-28 17:28:11'),('d071fc9c-b1c5-4527-9b72-509ef1b919e0',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiJkMDcxZmM5Yy1iMWM1LTQ1MjctOWI3Mi01MDllZjFiOTE5ZTAifSwiZXhwIjoxNjY5NzM1MTI3LCJqdGkiOiIxNjY5NzM0OTQ3ODc4NDAzMDAwIiwiaWF0IjoxNjY5NzM0OTQ3fQ.bgyGEYtNbi2DqKWq5Z2ByfILbc_o3RcNGvZ5GV8t3Bs','PostmanRuntime/7.28.4','::1',0,'2022-11-29 22:18:48',1,'2022-11-29 22:15:48','2022-11-29 22:15:48'),('d74d575e-3e89-452e-b3b2-6ad2a38b0f8c',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiJkNzRkNTc1ZS0zZTg5LTQ1MmUtYjNiMi02YWQyYTM4YjBmOGMifSwiZXhwIjoxNjY5NjkwMzU4LCJqdGkiOiIxNjY5NjkwMTc4NTExODg1MDAwIiwiaWF0IjoxNjY5NjkwMTc4fQ.y3D-MszcevSOrNRtiuy6XNdNZMFaB0_tA_rT_eMNh_I','PostmanRuntime/7.28.4','::1',0,'2022-11-29 09:52:39',1,'2022-11-29 09:49:39','2022-11-29 09:49:39'),('d85f9f30-a6ad-4340-9425-eeed5633af94',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiJkODVmOWYzMC1hNmFkLTQzNDAtOTQyNS1lZWVkNTYzM2FmOTQifSwiZXhwIjoxNjY5NzMxOTM4LCJqdGkiOiIxNjY5NzMxNzU4MDIyNDU5MDAwIiwiaWF0IjoxNjY5NzMxNzU4fQ.vDW8qY9TZlttl-wxjB5phkrT_wZdfSMb0RGqGgXaIeU','PostmanRuntime/7.28.4','::1',0,'2022-11-29 21:25:38',1,'2022-11-29 21:22:38','2022-11-29 21:22:38'),('de5d32fd-1001-434a-97f8-164c928eb9f6',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiJkZTVkMzJmZC0xMDAxLTQzNGEtOTdmOC0xNjRjOTI4ZWI5ZjYifSwiZXhwIjoxNjY5NDc3NDI4LCJqdGkiOiIxNjY5NDc3MjQ4NjAxOTI0MDAwIiwiaWF0IjoxNjY5NDc3MjQ4fQ.L1NFfjZzOCV4iM5u74KhGCoLUelX5Nvr7wPfi0QD38c','PostmanRuntime/7.28.4','::1',0,'2022-11-26 22:43:49',1,'2022-11-26 22:40:49','2022-11-26 22:40:49'),('e1d7962d-be21-4c41-be11-69e2ff4c567a',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiJlMWQ3OTYyZC1iZTIxLTRjNDEtYmUxMS02OWUyZmY0YzU2N2EifSwiZXhwIjoxNjY5NjkwMzQ4LCJqdGkiOiIxNjY5NjkwMTY4MjI0ODk4MDAwIiwiaWF0IjoxNjY5NjkwMTY4fQ.U4reHtB0qMCBwG6vVSA5GDcX-zCWI12mZdqnN5WAqGc','PostmanRuntime/7.28.4','::1',0,'2022-11-29 09:52:28',1,'2022-11-29 09:49:28','2022-11-29 09:49:28'),('e321a67d-a1f8-406f-a5ae-43ae201fdd5e',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiJlMzIxYTY3ZC1hMWY4LTQwNmYtYTVhZS00M2FlMjAxZmRkNWUifSwiZXhwIjoxNjY5NjkwMzQ5LCJqdGkiOiIxNjY5NjkwMTY5NjQxNjkwMDAwIiwiaWF0IjoxNjY5NjkwMTY5fQ.NYfmmn3ViHJsab0iWaeGa3ERX7pxGB3l1YKvR-kl6dY','PostmanRuntime/7.28.4','::1',0,'2022-11-29 09:52:30',1,'2022-11-29 09:49:30','2022-11-29 09:49:30'),('e9795d1d-f291-4085-8c26-c09c25412119',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiJlOTc5NWQxZC1mMjkxLTQwODUtOGMyNi1jMDljMjU0MTIxMTkifSwiZXhwIjoxNjY5NDc3NTM3LCJqdGkiOiIxNjY5NDc3MzU3OTU2ODI3MDAwIiwiaWF0IjoxNjY5NDc3MzU3fQ.Iw3Q-YUR0w4q8GoiJt-Y79jwVKB7IbEo_f4_KgAoeZ4','PostmanRuntime/7.28.4','::1',0,'2022-11-26 22:45:38',1,'2022-11-26 22:42:38','2022-11-26 22:42:38'),('e9e48d91-1368-4ceb-88be-bad4f1130df3',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiJlOWU0OGQ5MS0xMzY4LTRjZWItODhiZS1iYWQ0ZjExMzBkZjMifSwiZXhwIjoxNjY5NjkxMTMzLCJqdGkiOiIxNjY5NjkwOTUzNTkzNTQyMDAwIiwiaWF0IjoxNjY5NjkwOTUzfQ.IsWpOLnX4N_l9MT06Y4kzBELAX9zRTkHE6A5TUScDCk','PostmanRuntime/7.28.4','::1',0,'2022-11-29 10:05:34',1,'2022-11-29 10:02:34','2022-11-29 10:02:34'),('ed7a604a-9aeb-4a70-9bdc-5fe71ce012c8',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiJlZDdhNjA0YS05YWViLTRhNzAtOWJkYy01ZmU3MWNlMDEyYzgifSwiZXhwIjoxNjY5NjkxMTM4LCJqdGkiOiIxNjY5NjkwOTU4NDMzNjc0MDAwIiwiaWF0IjoxNjY5NjkwOTU4fQ.eTwTRRE0mpP_UD43k5E5Xxlb6eNFzMIQK6rK1qHwtDI','PostmanRuntime/7.28.4','::1',0,'2022-11-29 10:05:38',1,'2022-11-29 10:02:38','2022-11-29 10:02:38'),('edc687cd-56f5-4a44-9b62-d5d9896fea14',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiJlZGM2ODdjZC01NmY1LTRhNDQtOWI2Mi1kNWQ5ODk2ZmVhMTQifSwiZXhwIjoxNjY5NDc3NTcyLCJqdGkiOiIxNjY5NDc3MzkyMjk2MjM4MDAwIiwiaWF0IjoxNjY5NDc3MzkyfQ.6SGHBaAwtJqLTeuslJH3QfC5MbQwxsTnA7VvfG-6wx0','PostmanRuntime/7.28.4','::1',0,'2022-11-26 22:46:12',1,'2022-11-26 22:43:12','2022-11-26 22:43:12'),('f1686a3a-6243-435d-8637-a6e0ccc3ce83',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiJmMTY4NmEzYS02MjQzLTQzNWQtODYzNy1hNmUwY2NjM2NlODMifSwiZXhwIjoxNjY5NzE3ODcyLCJqdGkiOiIxNjY5NzE3NjkyNjYwMDk4MDAwIiwiaWF0IjoxNjY5NzE3NjkyfQ.dHicPr2I6n63g2mpLdivHUAeAJmMZDMEGn5ZWQFqjcw','PostmanRuntime/7.28.4','::1',0,'2022-11-29 17:31:13',1,'2022-11-29 17:28:13','2022-11-29 17:28:13'),('f2da75ae-3615-4f78-8de2-3fb064372171',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiJmMmRhNzVhZS0zNjE1LTRmNzgtOGRlMi0zZmIwNjQzNzIxNzEifSwiZXhwIjoxNjY5NjkxNDMwLCJqdGkiOiIxNjY5NjkxMjUwMTc0Njk3MDAwIiwiaWF0IjoxNjY5NjkxMjUwfQ.b2Aodun6j79M_nALWgaA5D_H9kipp_gqq23y3ZKTwIY','PostmanRuntime/7.28.4','::1',0,'2022-11-29 10:10:30',1,'2022-11-29 10:07:30','2022-11-29 10:07:30'),('f44383ee-8423-485d-9dcd-8d500fe2715e',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiJmNDQzODNlZS04NDIzLTQ4NWQtOWRjZC04ZDUwMGZlMjcxNWUifSwiZXhwIjoxNjY5NjQxOTE2LCJqdGkiOiIxNjY5NjQxNzM2OTMzMDU0MDAwIiwiaWF0IjoxNjY5NjQxNzM2fQ.edWJCksKDzEemejLbiIRsqFjHewx32lfGqFVn8vAVko','PostmanRuntime/7.28.4','::1',0,'2022-11-28 20:25:17',1,'2022-11-28 20:22:17','2022-11-28 20:22:17'),('f4d63010-e591-46e4-84eb-8a3b79bafb7c',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiJmNGQ2MzAxMC1lNTkxLTQ2ZTQtODRlYi04YTNiNzliYWZiN2MifSwiZXhwIjoxNjY5NjkxMTM1LCJqdGkiOiIxNjY5NjkwOTU1NzgxOTg0MDAwIiwiaWF0IjoxNjY5NjkwOTU1fQ.G5nT-9GVqkrhzDCq5QGmG8QDXMa5bwmF8MusfGtjmRw','PostmanRuntime/7.28.4','::1',0,'2022-11-29 10:05:36',1,'2022-11-29 10:02:36','2022-11-29 10:02:36'),('f5c46850-c19e-44a5-8fd5-b4eee9768be5',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiJmNWM0Njg1MC1jMTllLTQ0YTUtOGZkNS1iNGVlZTk3NjhiZTUifSwiZXhwIjoxNjY5NzAzODQ2LCJqdGkiOiIxNjY5NzAzNjY2MjgwMjk2MDAwIiwiaWF0IjoxNjY5NzAzNjY2fQ.vXxGhPRw_Z7ccVZfqnPGqPQ01cT0jfEET1VZl7ucBc8','PostmanRuntime/7.28.4','::1',0,'2022-11-29 13:37:26',1,'2022-11-29 13:34:26','2022-11-29 13:34:26'),('f89c073f-4dbe-40a4-a01c-a6e4d80806b4',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiJmODljMDczZi00ZGJlLTQwYTQtYTAxYy1hNmU0ZDgwODA2YjQifSwiZXhwIjoxNjY5NzMxODM0LCJqdGkiOiIxNjY5NzMxNjU0NDg4NjMwMDAwIiwiaWF0IjoxNjY5NzMxNjU0fQ.KiptvZs6kf41kYEynRI4e-vnB3tvFvr-HPNVwjqUcmM','PostmanRuntime/7.28.4','::1',0,'2022-11-29 21:23:55',1,'2022-11-29 21:20:55','2022-11-29 21:20:55'),('fe0ad08b-f932-49d9-9133-958d2b25dd06',3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjMsInJvbGUiOiJ1c2VyIiwiaWQiOiJmZTBhZDA4Yi1mOTMyLTQ5ZDktOTEzMy05NThkMmIyNWRkMDYifSwiZXhwIjoxNjY5NjkxNDE0LCJqdGkiOiIxNjY5NjkxMjM0OTkxMjQwMDAwIiwiaWF0IjoxNjY5NjkxMjM0fQ.z_CgESTAvBzXIZKRQRuMK5Fm7i0O8FdJqhZQZLjiRC4','PostmanRuntime/7.28.4','::1',0,'2022-11-29 10:10:15',1,'2022-11-29 10:07:15','2022-11-29 10:07:15');
/*!40000 ALTER TABLE `sessions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_addresses`
--

DROP TABLE IF EXISTS `user_addresses`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_addresses` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `city_id` int NOT NULL,
  `title` varchar(100) DEFAULT NULL,
  `icon` json DEFAULT NULL,
  `addr` varchar(255) NOT NULL,
  `lat` double DEFAULT NULL,
  `lng` double DEFAULT NULL,
  `status` int NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`) USING BTREE,
  KEY `city_id` (`city_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_addresses`
--

LOCK TABLES `user_addresses` WRITE;
/*!40000 ALTER TABLE `user_addresses` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_addresses` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_device_tokens`
--

DROP TABLE IF EXISTS `user_device_tokens`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_device_tokens` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int unsigned DEFAULT NULL,
  `is_production` tinyint(1) DEFAULT '0',
  `os` enum('ios','android','web') DEFAULT 'ios' COMMENT '1: iOS, 2: Android',
  `token` varchar(255) DEFAULT NULL,
  `status` smallint unsigned NOT NULL DEFAULT '1',
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `device_id` varchar(256) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`) USING BTREE,
  KEY `os` (`os`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_device_tokens`
--

LOCK TABLES `user_device_tokens` WRITE;
/*!40000 ALTER TABLE `user_device_tokens` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_device_tokens` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_tokens`
--

DROP TABLE IF EXISTS `user_tokens`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_tokens` (
  `id` varchar(255) NOT NULL,
  `user_id` int NOT NULL,
  `refresh_token` longtext NOT NULL,
  `access_token` longtext NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_tokens`
--

LOCK TABLES `user_tokens` WRITE;
/*!40000 ALTER TABLE `user_tokens` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_tokens` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `email` varchar(50) NOT NULL,
  `fb_id` varchar(50) DEFAULT NULL,
  `gg_id` varchar(50) DEFAULT NULL,
  `password` varchar(50) NOT NULL,
  `salt` varchar(50) DEFAULT NULL,
  `last_name` varchar(50) NOT NULL,
  `first_name` varchar(50) NOT NULL,
  `phone` varchar(20) DEFAULT NULL,
  `role` enum('user','admin','shipper') NOT NULL DEFAULT 'user',
  `avatar` json DEFAULT NULL,
  `status` int NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `apple_id` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (3,'test@gmail.com',NULL,NULL,'c702176f37591d33293ab64e33746039','WrMENbtAwNvxpNdqwLEmsXJVjbKCzuBaqrJYaopeBpIwmbodHf','test2','test1',NULL,'user',NULL,1,'2022-11-16 10:11:29',NULL,NULL),(4,'test2@gmail.com',NULL,NULL,'8f12b6ee01147a952ed03574ee250713','dEcoMiZbbeyISUGkMIumiYbEYWJwFBBDIQTdHemuKfDKNCvxCF','test2','test1',NULL,'user',NULL,1,'2022-11-16 10:11:33',NULL,NULL),(5,'test3@gmail.com',NULL,NULL,'ef7f7c602ae4e987077b87d8c59a7908','ZUvvqSPYGTxjLwgqGyQylvcnEwZITmZCxtaLUpYMyXSFmWVmFq','test2','test1',NULL,'user',NULL,1,'2022-11-16 10:11:36',NULL,NULL),(6,'test4@gmail.com',NULL,NULL,'bfa2844b086e0ba143b9a62c1bc10ba2','mJPfzdOYLVHkNHRmzNugaAvKqzJuYTqxgOkeDasVqBPsPHgScs','test2','test1',NULL,'user',NULL,1,'2022-11-16 10:11:38',NULL,NULL),(7,'test5@gmail.com',NULL,NULL,'a84839ca385b7468538442a1b43e1e82','pJvCYLISHGBiusQzVdYLGBBRpokPpxVwftHmYJuUuiAXmPsILS','test2','test1',NULL,'user',NULL,1,'2022-11-16 10:11:41',NULL,NULL),(8,'84bvppdxcx@privaterelay.appleid.com',NULL,NULL,'dda1b3dae2338588c5fd175735a8994c','uCasUJoosTOzBpyasHibLRiICcqdwZLlDZRkozLeAaMWgbZiio','test','test',NULL,'user',NULL,1,'2022-12-25 16:23:45','2022-12-25 16:23:45','001681.aad7411f8b09494799a4b4d735b48e05.0702');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping routines for database 'food_delivery'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-12-26  1:44:29
