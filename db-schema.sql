-- MySQL dump 10.13  Distrib 5.7.12, for osx10.10 (x86_64)
--
-- Host: localhost    Database: paper_review
-- ------------------------------------------------------
-- Server version	5.7.12

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
-- Table structure for table `Mitarbeiters`
--

DROP TABLE IF EXISTS `Mitarbeiters`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Mitarbeiters` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Pass_Hash` binary(32) NOT NULL,
  `Nme` varchar(100) NOT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=47 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Mitarbeiters`
--

LOCK TABLES `Mitarbeiters` WRITE;
/*!40000 ALTER TABLE `Mitarbeiters` DISABLE KEYS */;
INSERT INTO `Mitarbeiters` VALUES (1,'\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','Mohsen'),(2,'\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','test mitarbeiter'),(3,'\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','test mitarbeiter'),(4,'\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','test mitarbeiter'),(5,'\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','test mitarbeiter'),(6,'\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','test mitarbeiter'),(7,'\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','test mitarbeiter'),(8,'\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','test mitarbeiter'),(9,'\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','test mitarbeiter'),(10,'\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','test mitarbeiter'),(11,'\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','test mitarbeiter'),(12,'\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','test mitarbeiter'),(13,'\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','test mitarbeiter'),(14,'\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','test mitarbeiter'),(15,'\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','test mitarbeiter'),(16,'\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','test mitarbeiter'),(17,'\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','test mitarbeiter'),(18,'\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','test mitarbeiter'),(19,'\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','test mitarbeiter'),(20,'\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','test mitarbeiter'),(21,'\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','test mitarbeiter'),(22,'\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','test mitarbeiter'),(23,'\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','test mitarbeiter'),(24,'\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','test mitarbeiter'),(25,'\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','test mitarbeiter'),(26,'\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','test mitarbeiter'),(27,'\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','test mitarbeiter'),(28,'\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','test mitarbeiter'),(29,'\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','test mitarbeiter'),(30,'\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','test mitarbeiter'),(31,'\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','test mitarbeiter'),(32,'\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','test mitarbeiter'),(33,'\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','test mitarbeiter'),(34,'\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','test mitarbeiter'),(35,'\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','test mitarbeiter'),(36,'\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','test mitarbeiter'),(37,'\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','test mitarbeiter'),(38,'\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','test mitarbeiter'),(39,'\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','test mitarbeiter'),(40,'\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','test mitarbeiter'),(41,'\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','test mitarbeiter'),(42,'\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','test mitarbeiter'),(43,'\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','test mitarbeiter'),(44,'\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','test mitarbeiter'),(45,'\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','test mitarbeiter'),(46,'\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0','Jake');
/*!40000 ALTER TABLE `Mitarbeiters` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Research`
--

DROP TABLE IF EXISTS `Research`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Research` (
  `ResearchId` int(11) NOT NULL AUTO_INCREMENT,
  `Questions` varchar(1000) NOT NULL,
  `Review_Template` varchar(2000) DEFAULT NULL,
  PRIMARY KEY (`ResearchId`)
) ENGINE=InnoDB AUTO_INCREMENT=53 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Research`
--

LOCK TABLES `Research` WRITE;
/*!40000 ALTER TABLE `Research` DISABLE KEYS */;
INSERT INTO `Research` VALUES (1,'test1','wrtie whatever'),(2,'test1','wrtie whatever'),(3,'test1','wrtie whatever'),(4,'test1','wrtie whatever'),(5,'test1','wrtie whatever'),(6,'test1','wrtie whatever'),(7,'test1','wrtie whatever'),(8,'test1','wrtie whatever'),(9,'test1','wrtie whatever'),(10,'test1','wrtie whatever'),(11,'test1','wrtie whatever'),(12,'test1','wrtie whatever'),(13,'test1','wrtie whatever'),(14,'test1','wrtie whatever'),(15,'test1','wrtie whatever'),(16,'test1','wrtie whatever'),(17,'test1','wrtie whatever'),(18,'test1','wrtie whatever'),(19,'test1','wrtie whatever'),(20,'test1','wrtie whatever'),(21,'test1','wrtie whatever'),(22,'test1','wrtie whatever'),(23,'test1','wrtie whatever'),(24,'test1','wrtie whatever'),(25,'test1','wrtie whatever'),(26,'test1','wrtie whatever'),(27,'test1','wrtie whatever'),(28,'test1','wrtie whatever'),(29,'test1','wrtie whatever'),(30,'test1','wrtie whatever'),(31,'test1','wrtie whatever'),(32,'test1','wrtie whatever'),(33,'test1','wrtie whatever'),(34,'test1','wrtie whatever'),(35,'test1','wrtie whatever'),(36,'test1','wrtie whatever'),(37,'test1','wrtie whatever'),(38,'test1','wrtie whatever'),(39,'test1','wrtie whatever'),(40,'test1','wrtie whatever'),(41,'test1','wrtie whatever'),(42,'test1','wrtie whatever'),(43,'test1','wrtie whatever'),(44,'test1','wrtie whatever'),(45,'test1','wrtie whatever'),(46,'test1','wrtie whatever'),(47,'test1','wrtie whatever'),(48,'test1','wrtie whatever'),(49,'Question1: Hello\n Question2: Hi','rev template'),(50,'Question1: Hello\n Question2: Hi','rev template'),(51,'Question1: Hello\n Question2: Hi','rev template'),(52,'test1','wrtie whatever');
/*!40000 ALTER TABLE `Research` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Tags`
--

DROP TABLE IF EXISTS `Tags`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Tags` (
  `TagId` int(11) NOT NULL AUTO_INCREMENT,
  `Text` varchar(500) NOT NULL,
  PRIMARY KEY (`TagId`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Tags`
--

LOCK TABLES `Tags` WRITE;
/*!40000 ALTER TABLE `Tags` DISABLE KEYS */;
INSERT INTO `Tags` VALUES (1,'test1'),(2,'test2'),(3,'test1'),(4,'test2'),(5,'test1'),(6,'test2'),(7,'test1'),(8,'test2');
/*!40000 ALTER TABLE `Tags` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Vote_Tags`
--

DROP TABLE IF EXISTS `Vote_Tags`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Vote_Tags` (
  `Vote_Tags_Id` int(11) NOT NULL AUTO_INCREMENT,
  `Tag_Id` int(11) NOT NULL,
  `VoteId` int(11) NOT NULL,
  PRIMARY KEY (`Vote_Tags_Id`)
) ENGINE=InnoDB AUTO_INCREMENT=103 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Vote_Tags`
--

LOCK TABLES `Vote_Tags` WRITE;
/*!40000 ALTER TABLE `Vote_Tags` DISABLE KEYS */;
INSERT INTO `Vote_Tags` VALUES (1,1,5),(2,2,5),(3,3,6),(4,4,6),(5,5,7),(6,6,7),(7,7,8),(8,8,8),(9,1,11),(10,2,11),(11,1,12),(12,2,12),(13,1,13),(14,2,13),(15,1,14),(16,2,14),(17,1,15),(18,2,15),(19,1,16),(20,2,16),(21,1,17),(22,2,17),(23,1,18),(24,2,18),(25,1,19),(26,2,19),(27,1,20),(28,2,20),(29,1,21),(30,2,21),(31,1,22),(32,2,22),(33,1,23),(34,2,23),(35,1,24),(36,2,24),(37,1,25),(38,2,25),(39,1,26),(40,2,26),(41,1,27),(42,2,27),(43,1,28),(44,2,28),(45,1,29),(46,2,29),(47,1,30),(48,2,30),(49,1,31),(50,2,31),(51,1,32),(52,2,32),(53,1,33),(54,2,33),(55,1,34),(56,2,34),(57,1,35),(58,2,35),(59,1,36),(60,2,36),(61,1,37),(62,2,37),(63,1,38),(64,2,38),(65,1,39),(66,2,39),(67,1,40),(68,2,40),(69,1,41),(70,2,41),(71,1,42),(72,2,42),(73,1,43),(74,2,43),(75,1,45),(76,2,45),(77,1,46),(78,2,46),(79,1,47),(80,2,47),(81,1,48),(82,2,48),(83,1,49),(84,2,49),(85,1,50),(86,2,50),(87,1,51),(88,2,51),(89,1,52),(90,2,52),(91,1,53),(92,2,53),(93,1,54),(94,2,54),(95,1,55),(96,2,55),(97,1,56),(98,2,56),(99,1,57),(100,2,57),(101,1,58),(102,2,58);
/*!40000 ALTER TABLE `Vote_Tags` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `articles`
--

DROP TABLE IF EXISTS `articles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `articles` (
  `ArticleId` int(11) NOT NULL AUTO_INCREMENT,
  `Title` varchar(500) NOT NULL,
  `year` varchar(4) DEFAULT NULL,
  `cited_by` varchar(100) DEFAULT NULL,
  `Keywords` varchar(500) DEFAULT NULL,
  `Abstract` varchar(4000) DEFAULT NULL,
  `Journal` varchar(500) DEFAULT NULL,
  `ResearchId` int(11) NOT NULL,
  `Authors` varchar(1000) NOT NULL,
  PRIMARY KEY (`ArticleId`)
) ENGINE=InnoDB AUTO_INCREMENT=51 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `articles`
--

LOCK TABLES `articles` WRITE;
/*!40000 ALTER TABLE `articles` DISABLE KEYS */;
INSERT INTO `articles` VALUES (1,'title','','','','','',4,'author'),(2,'title','','','','','',5,'author'),(3,'title','','','','','',6,'author'),(4,'title','','','','','',7,'author'),(5,'title','','','','','',8,'author'),(6,'title','','','','','',9,'author'),(7,'title','','','','','',10,'author'),(8,'title','','','','','',11,'author'),(9,'title','','','','','',12,'author'),(10,'title','','','','','',13,'author'),(11,'title','','','','','',14,'author'),(12,'title','','','','','',15,'author'),(13,'title','year','cited_by','keywords','abstract','journal',4,'authors'),(14,'title','','','','','',16,'author'),(15,'title','','','','','',17,'author'),(16,'title','','','','','',18,'author'),(17,'title','','','','','',19,'author'),(18,'title','','','','','',20,'author'),(19,'title','','','','','',21,'author'),(20,'title','','','','','',22,'author'),(21,'title','','','','','',23,'author'),(22,'title','','','','','',24,'author'),(23,'title','','','','','',25,'author'),(24,'title','','','','','',26,'author'),(25,'title','','','','','',27,'author'),(26,'title','','','','','',28,'author'),(27,'title','','','','','',29,'author'),(28,'title','','','','','',30,'author'),(29,'title','','','','','',31,'author'),(30,'title','','','','','',32,'author'),(31,'title','','','','','',33,'author'),(32,'title','','','','','',34,'author'),(33,'title','','','','','',35,'author'),(34,'title','','','','','',36,'author'),(35,'title','','','','','',37,'author'),(36,'title','','','','','',38,'author'),(37,'title','','','','','',39,'author'),(38,'title','','','','','',40,'author'),(39,'title','','','','','',41,'author'),(40,'title','year','cited_by','keywords','abstract','journal',6,'authors'),(41,'title','','','','','',42,'author'),(42,'title','','','','','',43,'author'),(43,'title','','','','','',44,'author'),(44,'title','','','','','',45,'author'),(45,'title','','','','','',46,'author'),(46,'title','','','','','',47,'author'),(47,'title','','','','','',48,'author'),(48,'article 1','2016','','key1 key2 key3','not so fun so far','a good one',51,'a1 a2 a3'),(49,'article 2','2015','','key1 key2 key3','so fun so far','a moderate one',51,'a1 a2 a3'),(50,'title','','','','','',52,'author');
/*!40000 ALTER TABLE `articles` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `votes`
--

DROP TABLE IF EXISTS `votes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `votes` (
  `VoteId` int(11) NOT NULL AUTO_INCREMENT,
  `Vote_State` int(11) NOT NULL,
  `MitarbeiterId` int(11) NOT NULL,
  `ArticleId` int(11) NOT NULL,
  `Review` varchar(5000) DEFAULT NULL,
  PRIMARY KEY (`VoteId`)
) ENGINE=InnoDB AUTO_INCREMENT=59 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `votes`
--

LOCK TABLES `votes` WRITE;
/*!40000 ALTER TABLE `votes` DISABLE KEYS */;
INSERT INTO `votes` VALUES (1,1,1,1,''),(2,1,1,1,''),(3,1,1,1,''),(4,1,1,1,''),(5,1,1,1,''),(6,1,1,1,''),(7,1,1,1,''),(8,1,1,1,''),(9,1,1,1,''),(10,1,1,1,''),(11,1,1,1,''),(12,1,1,1,''),(13,1,1,1,''),(14,1,1,1,''),(15,1,1,1,''),(16,1,1,1,''),(17,1,1,1,''),(18,1,1,1,''),(19,1,1,1,''),(20,1,1,1,''),(21,1,1,1,''),(22,1,1,1,''),(23,1,1,1,''),(24,1,1,1,''),(25,1,1,1,''),(26,1,1,1,''),(27,1,1,1,''),(28,1,1,1,''),(29,1,1,1,''),(30,1,1,1,''),(31,1,1,1,''),(32,1,1,1,''),(33,1,1,1,''),(34,1,1,1,''),(35,1,1,1,''),(36,1,1,1,''),(37,1,1,1,''),(38,1,1,1,''),(39,1,1,1,''),(40,1,1,1,''),(41,1,1,1,''),(42,1,1,1,''),(43,1,1,1,''),(44,1,1,2,'rev'),(45,1,1,1,''),(46,1,1,1,''),(47,1,1,1,''),(48,1,1,1,''),(49,1,1,1,''),(50,1,1,1,''),(51,1,1,1,''),(52,1,1,1,''),(53,1,1,1,''),(54,1,1,1,''),(55,1,1,1,''),(56,1,1,1,''),(57,1,1,5,'a good one'),(58,1,1,1,'');
/*!40000 ALTER TABLE `votes` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2016-05-17 19:02:01
