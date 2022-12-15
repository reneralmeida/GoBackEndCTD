
-- Delete user if already exists
DROP USER IF EXISTS 'user1'@'localhost';

-- Create user with all privileges
CREATE USER 'user1'@'localhost' IDENTIFIED BY 'secret_password';
GRANT ALL PRIVILEGES ON *.* TO 'user1'@'localhost';
-- 
CREATE DATABASE `checkpoint_db`;
USE `checkpoint_db`;
-- MySQL dump 10.13  Distrib 8.0.25, for Linux (x86_64)
--
-- Host: localhost    Database: checkpoint_db
-- ------------------------------------------------------
-- Server version	8.0.29-0ubuntu0.20.04.3

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
-- Table structure for table `dentists`
--

DROP TABLE IF EXISTS `dentists`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `dentists` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  `lastname` varchar(50) NOT NULL,
  `registration` int NOT NULL UNIQUE,
  `email` varchar(50) NOT NULL,
  
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `dentists`
--

LOCK TABLES `dentists` WRITE;
/*!40000 ALTER TABLE `dentists` DISABLE KEYS */;
INSERT INTO `dentists` VALUES (1, 'Joao', 'Batista', 10, 'johnthebaptist@gmail.com'),(2, 'Eleonor', 'Magnolia', 11, 'elemag@gmail.com'),(3, 'Homer', 'Simpson', 12, 'homesimpson@gmail.com'),(4, 'Dennis', 'Taylor', 13, 'd.taylor@gmail.com'),(5, 'Juan', 'Pablo', 14, 'juan.p@gmail.com');
/*!40000 ALTER TABLE `dentists` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

-- Table structure for table `patients`

CREATE TABLE `patients` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  `lastname` varchar(50) NOT NULL,
  `document` varchar(50) NOT NULL UNIQUE,
  `reg_date` DATETIME DEFAULT NULL,
  
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `patients`
--

LOCK TABLES `patients` WRITE;
/*!40000 ALTER TABLE `patients` DISABLE KEYS */;
INSERT INTO `patients` VALUES (1, 'Miriam', 'Nunes', '58134849684', STR_TO_DATE('01-01-2018 15:42:15', '%d-%m-%Y %H:%i:%s')),(2, 'Clodoaldo', 'Silva', '25138151752', STR_TO_DATE('10-02-2020 15:42:15', '%d-%m-%Y %H:%i:%s')),(3, 'Juarez', 'Alves', '21351384674', STR_TO_DATE('25-03-2021 10:23:35', '%d-%m-%Y %H:%i:%s')),(4, 'Leticia', 'Gomes', '13438484684', STR_TO_DATE('30-04-2022 16:26:46', '%d-%m-%Y %H:%i:%s')),(5, 'Jonas', 'Gomes', '13418548438', STR_TO_DATE('18-06-2019 20:25:36', '%d-%m-%Y %H:%i:%s'));
/*!40000 ALTER TABLE `patients` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;


-- Table structure for table `appointments`

CREATE TABLE appointments (
    id INT NOT NULL AUTO_INCREMENT,
    description VARCHAR(250) NOT NULL,
    date_and_time DATETIME NOT NULL,
    dentists_registration int NOT NULL,
    patients_id int NOT NULL,

    PRIMARY KEY (id)
) ENGINE = INNODB;

ALTER TABLE `appointments`
  ADD KEY `FK_dentists` (`dentists_registration`),
  ADD KEY `FK_patients` (`patients_id`);
  
ALTER TABLE `appointments`
  ADD CONSTRAINT `FK_dentists` FOREIGN KEY (`dentists_registration`) REFERENCES `dentists` (`registration`),
  ADD CONSTRAINT `FK_patients` FOREIGN KEY (`patients_id`) REFERENCES `patients` (`id`);

--
-- Dumping data for table `appointments`
--

LOCK TABLES `appointments` WRITE;
/*!40000 ALTER TABLE `appointments` DISABLE KEYS */;
INSERT INTO `appointments` VALUES (1, 'Rotina', STR_TO_DATE('20-01-2023 10:00:00', '%d-%m-%Y %H:%i:%s'), 10, 1),(2, 'Endoscopia', STR_TO_DATE('10-02-2023 10:00:00', '%d-%m-%Y %H:%i:%s'), 11, 2),(3, 'Raio-X', STR_TO_DATE('12-01-2023 08:00:00', '%d-%m-%Y %H:%i:%s'), 12, 3),(4, 'Ressonância', STR_TO_DATE('10-02-2023 12:00:00', '%d-%m-%Y %H:%i:%s'), 13, 4),(5, 'Eletrocardiograma', STR_TO_DATE('13-01-2023 14:00:00', '%d-%m-%Y %H:%i:%s'), 14, 5);
/*!40000 ALTER TABLE `appointments` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;


/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-05-12 11:30:53