-- MySQL dump 10.13  Distrib 5.7.17, for Win64 (x86_64)
--
-- Host: 127.0.0.1    Database: classification
-- ------------------------------------------------------
-- Server version	5.7.20-log

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
-- Table structure for table `allchildrenperattribute`
--

DROP TABLE IF EXISTS `allchildrenperattribute`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `allchildrenperattribute` (
  `id_attribute` int(10) unsigned NOT NULL,
  `text` varchar(50) NOT NULL,
  `children` longtext,
  PRIMARY KEY (`id_attribute`),
  UNIQUE KEY `allChildrenPerAttribute_id_attribute_UNIQUE` (`id_attribute`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `allchildrenperattribute`
--

LOCK TABLES `allchildrenperattribute` WRITE;
/*!40000 ALTER TABLE `allchildrenperattribute` DISABLE KEYS */;
INSERT INTO `allchildrenperattribute` VALUES (1,'Integrity Protection Assets','1'),(2,'Behavior','2'),(3,'Data','3'),(4,'Data and behavior','4'),(5,'Representation','5,6,7,8,9,10,11,13,12'),(6,'Static','6'),(7,'In memory','7,9,10'),(8,'In execution','8,11,13,12'),(9,'Code invariants','9'),(10,'Data invariants','10'),(11,'Trace','11,12'),(12,'Timed trace','12'),(13,'HW counters','13'),(14,'Granularity','14,15,17,18,19,95'),(15,'Instructions','15'),(16,'BB','16'),(17,'Function','17'),(18,'Slice','18'),(19,'Application','19'),(20,'Lifecycle activity','20,21,22,23,24,25,94'),(21,'Pre-compile','21'),(22,'Compile','22'),(23,'Post-compile','23'),(24,'Load','24'),(25,'Run','25'),(26,'Not root','26'),(27,'Attack','27,28,29,30,31,78'),(28,'Binary ','28'),(29,'Process memory','29'),(30,'Runtime data','30'),(31,'Control flow','31'),(32,'Measure','32,33,34,35,38,41,44,50,57,62,67,36,37,39,42,43,45,46,47,48,49,51,52,53,54,55,58,59,60,61,63,64,65,69,70,71,89,90,91,92'),(33,'Local','33'),(34,'Remote','34'),(35,'Monitor','35,36,37'),(36,'State inspection','36'),(37,'Introspection','37'),(38,'Response','38,39,92'),(39,'Proactive','39'),(40,'Postmortem','40'),(41,'Transformation','41,42,43'),(42,'Manual','42'),(43,'Automatic','43'),(44,'Check','44,45,46,47,48,49'),(45,'Checksum','45'),(46,'Signature','46'),(47,'Equation eval','47'),(48,'Majority vote','48'),(49,'Access control','49'),(50,'Hardening','50,51,52,53,54,55,91'),(51,'Cyclic checks','51'),(52,'Mutation','52'),(53,'Code concealment','53'),(54,'Cloning','54'),(55,'Layered interpretation','55'),(56,'Block chain','56'),(57,'Overhead','57,58,59,60,61'),(58,'Fair','58'),(59,'Medium','59'),(60,'High','60'),(61,'N/A','61'),(62,'Trust anchor','62,63,64,65,71,89'),(63,'TPM','63'),(64,'SGX','64'),(65,'Other','65'),(66,'None','66'),(67,'Protection level','67,69,70,90'),(68,'Internal','68'),(69,'External','69'),(70,'Hypervisor','70'),(71,'Software','71'),(72,'Reverse engineering','72,27,73,80,84,26,28,29,30,31,78,79,81,82,83,85,86,87,88'),(73,'Attacker','73,26'),(78,'Call interposition','78'),(79,'Disassembler','79'),(80,'Tools','80,79,81,82,83'),(81,'Debugger','81'),(82,'Tracer','82'),(83,'Emulator','83'),(84,'Discovery','84,85,86,87,88'),(85,'Pattern matching','85'),(86,'Taint analysis','86'),(87,'Graph-based analysis','87'),(88,'Symbolic execution','88'),(89,'Dongle','89'),(90,'Self-check','90'),(91,'Hash chain','91'),(92,'Reactive','92'),(93,'Asset','93,2,3,4'),(94,'Link','94'),(95,'Basic block','95');
/*!40000 ALTER TABLE `allchildrenperattribute` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `allparentsperattribute`
--

DROP TABLE IF EXISTS `allparentsperattribute`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `allparentsperattribute` (
  `id_attribute` int(10) unsigned NOT NULL,
  `text` varchar(50) NOT NULL,
  `parents` longtext,
  PRIMARY KEY (`id_attribute`),
  UNIQUE KEY `allParentsPerAttribute_id_attribute_UNIQUE` (`id_attribute`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `allparentsperattribute`
--

LOCK TABLES `allparentsperattribute` WRITE;
/*!40000 ALTER TABLE `allparentsperattribute` DISABLE KEYS */;
INSERT INTO `allparentsperattribute` VALUES (1,'Integrity Protection Assets',''),(2,'Behavior','Asset'),(3,'Data','Asset'),(4,'Data and behavior','Asset'),(5,'Representation',''),(6,'Static','Representation'),(7,'In memory','Representation'),(8,'In execution','Representation'),(9,'Code invariants','In memory,Representation'),(10,'Data invariants','In memory,Representation'),(11,'Trace','In execution,Representation'),(12,'Timed trace','Trace,In execution,Representation'),(13,'HW counters','In execution,Representation'),(14,'Granularity',''),(15,'Instructions','Granularity'),(16,'BB',''),(17,'Function','Granularity'),(18,'Slice','Granularity'),(19,'Application','Granularity'),(20,'Lifecycle activity',''),(21,'Pre-compile','Lifecycle activity'),(22,'Compile','Lifecycle activity'),(23,'Post-compile','Lifecycle activity'),(24,'Load','Lifecycle activity'),(25,'Run','Lifecycle activity'),(26,'Not root','Attacker,Reverse engineering'),(27,'Attack','Reverse engineering'),(28,'Binary ','Attack,Reverse engineering'),(29,'Process memory','Attack,Reverse engineering'),(30,'Runtime data','Attack,Reverse engineering'),(31,'Control flow','Attack,Reverse engineering'),(32,'Measure',''),(33,'Local','Measure'),(34,'Remote','Measure'),(35,'Monitor','Measure'),(36,'State inspection','Monitor,Measure'),(37,'Introspection','Monitor,Measure'),(38,'Response','Measure'),(39,'Proactive','Response,Measure'),(40,'Postmortem',''),(41,'Transformation','Measure'),(42,'Manual','Transformation,Measure'),(43,'Automatic','Transformation,Measure'),(44,'Check','Measure'),(45,'Checksum','Check,Measure'),(46,'Signature','Check,Measure'),(47,'Equation eval','Check,Measure'),(48,'Majority vote','Check,Measure'),(49,'Access control','Check,Measure'),(50,'Hardening','Measure'),(51,'Cyclic checks','Hardening,Measure'),(52,'Mutation','Hardening,Measure'),(53,'Code concealment','Hardening,Measure'),(54,'Cloning','Hardening,Measure'),(55,'Layered interpretation','Hardening,Measure'),(56,'Block chain',''),(57,'Overhead','Measure'),(58,'Fair','Overhead,Measure'),(59,'Medium','Overhead,Measure'),(60,'High','Overhead,Measure'),(61,'N/A','Overhead,Measure'),(62,'Trust anchor','Measure'),(63,'TPM','Trust anchor,Measure'),(64,'SGX','Trust anchor,Measure'),(65,'Other','Trust anchor,Measure'),(66,'None',''),(67,'Protection level','Measure'),(68,'Internal',''),(69,'External','Protection level,Measure'),(70,'Hypervisor','Protection level,Measure'),(71,'Software','Trust anchor,Measure'),(72,'Reverse engineering',''),(73,'Attacker','Reverse engineering'),(78,'Call interposition','Attack,Reverse engineering'),(79,'Disassembler','Tools,Reverse engineering'),(80,'Tools','Reverse engineering'),(81,'Debugger','Tools,Reverse engineering'),(82,'Tracer','Tools,Reverse engineering'),(83,'Emulator','Tools,Reverse engineering'),(84,'Discovery','Reverse engineering'),(85,'Pattern matching','Discovery,Reverse engineering'),(86,'Taint analysis','Discovery,Reverse engineering'),(87,'Graph-based analysis','Discovery,Reverse engineering'),(88,'Symbolic execution','Discovery,Reverse engineering'),(89,'Dongle','Trust anchor,Measure'),(90,'Self-check','Protection level,Measure'),(91,'Hash chain','Hardening,Measure'),(92,'Reactive','Response,Measure'),(93,'Asset',''),(94,'Link','Lifecycle activity'),(95,'Basic block','Granularity');
/*!40000 ALTER TABLE `allparentsperattribute` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `attribute`
--

DROP TABLE IF EXISTS `attribute`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `attribute` (
  `id_attribute` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `text` varchar(50) NOT NULL,
  `x` varchar(50) DEFAULT '',
  `y` varchar(50) DEFAULT '',
  `xMajor` varchar(50) DEFAULT '',
  `yMajor` varchar(50) DEFAULT '',
  `x3D` varchar(50) DEFAULT '',
  `y3D` varchar(50) DEFAULT '',
  `z3D` varchar(50) DEFAULT '',
  `xMajor3D` varchar(50) DEFAULT '',
  `yMajor3D` varchar(50) DEFAULT '',
  `zMajor3D` varchar(50) DEFAULT '',
  `major` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`id_attribute`),
  UNIQUE KEY `id_attribute_UNIQUE` (`id_attribute`),
  UNIQUE KEY `attribute_text_UNIQUE` (`text`)
) ENGINE=InnoDB AUTO_INCREMENT=96 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `attribute`
--

LOCK TABLES `attribute` WRITE;
/*!40000 ALTER TABLE `attribute` DISABLE KEYS */;
INSERT INTO `attribute` VALUES (1,'Integrity Protection Assets','','','','','','','','','','',0),(2,'Behavior','','','','','','','','','','',0),(3,'Data','','','','','','','','','','',0),(4,'Data and behavior','','','','','','','','','','',0),(5,'Representation','','','','','','','','','','',1),(6,'Static','','','','','','','','','','',0),(7,'In memory','','','','','','','','','','',0),(8,'In execution','','','','','','','','','','',0),(9,'Code invariants','','','','','','','','','','',0),(10,'Data invariants','','','','','','','','','','',0),(11,'Trace','','','','','','','','','','',0),(12,'Timed trace','','','','','','','','','','',0),(13,'HW counters','','','','','','','','','','',0),(14,'Granularity','','','','','','','','','','',1),(15,'Instructions','','','','','','','','','','',0),(16,'BB','','','','','','','','','','',0),(17,'Function','','','','','','','','','','',0),(18,'Slice','','','','','','','','','','',0),(19,'Application','','','','','','','','','','',0),(20,'Lifecycle activity','','','','','','','','','','',1),(21,'Pre-compile','','','','','','','','','','',0),(22,'Compile','','','','','','','','','','',0),(23,'Post-compile','','','','','','','','','','',0),(24,'Load','','','','','','','','','','',0),(25,'Run','','','','','','','','','','',0),(26,'Not root','','','','','','','','','','',0),(27,'Attack','','','','','','','','','','',1),(28,'Binary ','','','','','','','','','','',0),(29,'Process memory','','','','','','','','','','',0),(30,'Runtime data','','','','','','','','','','',0),(31,'Control flow','','','','','','','','','','',0),(32,'Measure','','','','','','','','','','',1),(33,'Local','','','','','','','','','','',0),(34,'Remote','','','','','','','','','','',0),(35,'Monitor','','','','','','','','','','',0),(36,'State inspection','','','','','','','','','','',0),(37,'Introspection','','','','','','','','','','',0),(38,'Response','','','','','','','','','','',0),(39,'Proactive','','','','','','','','','','',0),(40,'Postmortem','','','','','','','','','','',0),(41,'Transformation','','','','','','','','','','',0),(42,'Manual','','','','','','','','','','',0),(43,'Automatic','','','','','','','','','','',0),(44,'Check','','','','','','','','','','',0),(45,'Checksum','','','','','','','','','','',0),(46,'Signature','','','','','','','','','','',0),(47,'Equation eval','','','','','','','','','','',0),(48,'Majority vote','','','','','','','','','','',0),(49,'Access control','','','','','','','','','','',0),(50,'Hardening','','','','','','','','','','',0),(51,'Cyclic checks','','','','','','','','','','',0),(52,'Mutation','','','','','','','','','','',0),(53,'Code concealment','','','','','','','','','','',0),(54,'Cloning','','','','','','','','','','',0),(55,'Layered interpretation','','','','','','','','','','',0),(56,'Block chain','','','','','','','','','','',0),(57,'Overhead','','','','','','','','','','',1),(58,'Fair','','','','','','','','','','',0),(59,'Medium','','','','','','','','','','',0),(60,'High','','','','','','','','','','',0),(61,'N/A','','','','','','','','','','',0),(62,'Trust anchor','','','','','','','','','','',1),(63,'TPM','','','','','','','','','','',0),(64,'SGX','','','','','','','','','','',0),(65,'Other','','','','','','','','','','',0),(66,'None','','','','','','','','','','',0),(67,'Protection level','','','','','','','','','','',1),(68,'Internal','','','','','','','','','','',0),(69,'External','','','','','','','','','','',0),(70,'Hypervisor','','','','','','','','','','',0),(71,'Software','','','','','','','','','','',0),(72,'Reverse engineering','','','','','','','','','','',1),(73,'Attacker','','','','','','','','','','',1),(78,'Call interposition','','','','','','','','','','',0),(79,'Disassembler','','','','','','','','','','',0),(80,'Tools','','','','','','','','','','',1),(81,'Debugger','','','','','','','','','','',0),(82,'Tracer','','','','','','','','','','',0),(83,'Emulator','','','','','','','','','','',0),(84,'Discovery','','','','','','','','','','',1),(85,'Pattern matching','','','','','','','','','','',0),(86,'Taint analysis','','','','','','','','','','',0),(87,'Graph-based analysis','','','','','','','','','','',0),(88,'Symbolic execution','','','','','','','','','','',0),(89,'Dongle','','','','','','','','','','',0),(90,'Self-check','','','','','','','','','','',0),(91,'Hash chain','','','','','','','','','','',0),(92,'Reactive','','','','','','','','','','',0),(93,'Asset','','','','','','','','','','',1),(94,'Link','','','','','','','','','','',0),(95,'Basic block','','','','','','','','','','',0);
/*!40000 ALTER TABLE `attribute` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `dimension`
--

DROP TABLE IF EXISTS `dimension`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `dimension` (
  `id_dimension` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `text` varchar(50) NOT NULL,
  `x` varchar(50) DEFAULT '',
  `y` varchar(50) DEFAULT '',
  `xMajor` varchar(50) DEFAULT '',
  `yMajor` varchar(50) DEFAULT '',
  PRIMARY KEY (`id_dimension`),
  UNIQUE KEY `id_dimension_UNIQUE` (`id_dimension`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `dimension`
--

LOCK TABLES `dimension` WRITE;
/*!40000 ALTER TABLE `dimension` DISABLE KEYS */;
INSERT INTO `dimension` VALUES (1,'System view','','','',''),(2,'Attack view','','','',''),(3,'Defense view','','','',''),(4,'Interdimensional view','','','','');
/*!40000 ALTER TABLE `dimension` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `mapping`
--

DROP TABLE IF EXISTS `mapping`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `mapping` (
  `id_mapping` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `id_paper` int(11) NOT NULL,
  `id_attribute` int(11) unsigned NOT NULL,
  `referenceCount` int(20) DEFAULT '0',
  PRIMARY KEY (`id_mapping`),
  UNIQUE KEY `id_mapping_UNIQUE` (`id_mapping`),
  UNIQUE KEY `mapping_id_paper_id_attribute_UNIQUE` (`id_paper`,`id_attribute`),
  KEY `mapping_id_attribute_foreign` (`id_attribute`),
  CONSTRAINT `mapping_id_attribute_foreign` FOREIGN KEY (`id_attribute`) REFERENCES `attribute` (`id_attribute`) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=678 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `mapping`
--

LOCK TABLES `mapping` WRITE;
/*!40000 ALTER TABLE `mapping` DISABLE KEYS */;
INSERT INTO `mapping` VALUES (1,1,2,0),(2,2,2,0),(3,3,2,0),(4,4,2,0),(5,5,2,0),(6,6,2,0),(7,7,2,0),(8,8,2,0),(9,9,2,0),(10,10,2,0),(11,11,2,0),(12,12,2,0),(13,13,2,0),(14,14,2,0),(15,15,2,0),(16,16,2,0),(17,17,2,0),(18,18,2,0),(19,19,2,0),(20,20,2,0),(21,21,2,0),(22,22,2,0),(23,23,2,0),(24,24,2,0),(25,25,2,0),(26,26,2,0),(27,27,2,0),(28,28,2,0),(29,29,2,0),(30,30,2,0),(31,31,2,0),(32,32,2,0),(33,33,2,0),(34,34,2,0),(35,35,2,0),(36,36,2,0),(37,37,2,0),(38,38,2,0),(39,39,2,0),(40,40,3,0),(41,4,3,0),(42,6,3,0),(43,41,3,0),(44,42,3,0),(45,14,3,0),(46,43,3,0),(47,44,3,0),(48,45,3,0),(49,46,3,0),(50,47,3,0),(51,4,4,0),(52,6,4,0),(53,14,4,0),(54,7,6,0),(55,11,6,0),(56,12,6,0),(57,13,6,0),(58,26,6,0),(59,46,6,0),(60,33,6,0),(61,37,6,0),(62,2,9,0),(63,3,9,0),(64,4,9,0),(65,6,9,0),(66,8,9,0),(67,14,9,0),(68,16,9,0),(69,17,9,0),(70,18,9,0),(71,24,9,0),(72,25,9,0),(73,27,9,0),(74,29,9,0),(75,32,9,0),(76,34,9,0),(77,36,9,0),(78,38,9,0),(79,39,9,0),(80,41,10,0),(81,42,10,0),(82,43,10,0),(83,19,10,0),(84,44,10,0),(85,45,10,0),(86,47,10,0),(87,1,11,0),(88,40,11,0),(89,5,11,0),(90,9,11,0),(91,20,11,0),(92,23,11,0),(93,28,11,0),(94,48,11,0),(95,49,11,0),(96,21,12,0),(97,22,12,0),(98,31,12,0),(99,35,12,0),(100,36,12,0),(101,30,13,0),(102,36,13,0),(103,1,15,0),(104,40,15,0),(105,5,15,0),(106,42,15,0),(107,2,16,0),(108,3,16,0),(109,8,16,0),(110,12,16,0),(111,16,16,0),(112,17,16,0),(113,20,16,0),(114,48,16,0),(115,38,16,0),(116,6,17,0),(117,11,17,0),(118,15,17,0),(119,43,17,0),(120,19,17,0),(121,44,17,0),(122,28,17,0),(123,31,17,0),(124,33,17,0),(125,34,17,0),(126,35,17,0),(127,36,17,0),(128,40,18,0),(129,41,18,0),(130,9,18,0),(131,18,18,0),(132,20,18,0),(133,24,18,0),(134,45,18,0),(135,29,18,0),(136,37,18,0),(137,4,19,0),(138,7,19,0),(139,10,19,0),(140,13,19,0),(141,14,19,0),(142,21,19,0),(143,22,19,0),(144,23,19,0),(145,25,19,0),(146,26,19,0),(147,27,19,0),(148,32,19,0),(149,46,19,0),(150,49,19,0),(151,47,19,0),(152,39,19,0),(153,40,21,0),(154,41,21,0),(155,42,21,0),(156,14,21,0),(157,19,21,0),(158,23,21,0),(159,44,21,0),(160,45,21,0),(161,28,21,0),(162,36,21,0),(163,9,22,0),(164,24,22,0),(165,37,22,0),(166,1,23,0),(167,2,23,0),(168,3,23,0),(169,4,23,0),(170,6,23,0),(171,7,23,0),(172,8,23,0),(173,10,23,0),(174,13,23,0),(175,15,23,0),(176,43,23,0),(177,16,23,0),(178,17,23,0),(179,18,23,0),(180,20,23,0),(181,21,23,0),(182,22,23,0),(183,25,23,0),(184,26,23,0),(185,27,23,0),(186,30,23,0),(187,31,23,0),(188,32,23,0),(189,48,23,0),(190,49,23,0),(191,33,23,0),(192,34,23,0),(193,35,23,0),(194,38,23,0),(195,39,23,0),(196,17,24,0),(197,11,25,0),(198,29,25,0),(199,1,26,0),(200,40,26,0),(201,3,26,0),(202,6,26,0),(203,41,26,0),(204,42,26,0),(205,7,26,0),(206,10,26,0),(207,13,26,0),(208,43,26,0),(209,45,26,0),(210,26,26,0),(211,27,26,0),(212,30,26,0),(213,48,26,0),(214,49,26,0),(215,33,26,0),(216,47,26,0),(217,37,26,0),(218,7,28,0),(219,11,28,0),(220,12,28,0),(221,13,28,0),(222,26,28,0),(223,46,28,0),(224,37,28,0),(225,2,29,0),(226,3,29,0),(227,4,29,0),(228,6,29,0),(229,8,29,0),(230,14,29,0),(231,16,29,0),(232,17,29,0),(233,18,29,0),(234,21,29,0),(235,22,29,0),(236,24,29,0),(237,25,29,0),(238,27,29,0),(239,29,29,0),(240,32,29,0),(241,33,29,0),(242,34,29,0),(243,36,29,0),(244,38,29,0),(245,39,29,0),(246,40,30,0),(247,41,30,0),(248,42,30,0),(249,9,30,0),(250,43,30,0),(251,19,30,0),(252,44,30,0),(253,45,30,0),(254,47,30,0),(255,1,31,0),(256,40,31,0),(257,5,31,0),(258,9,31,0),(259,20,31,0),(260,23,31,0),(261,28,31,0),(262,30,31,0),(263,48,31,0),(264,49,31,0),(265,1,33,0),(266,2,33,0),(267,40,33,0),(268,3,33,0),(269,5,33,0),(270,6,33,0),(271,41,33,0),(272,42,33,0),(273,7,33,0),(274,8,33,0),(275,9,33,0),(276,10,33,0),(277,12,33,0),(278,14,33,0),(279,15,33,0),(280,16,33,0),(281,17,33,0),(282,18,33,0),(283,19,33,0),(284,20,33,0),(285,24,33,0),(286,25,33,0),(287,26,33,0),(288,27,33,0),(289,28,33,0),(290,29,33,0),(291,30,33,0),(292,32,33,0),(293,48,33,0),(294,49,33,0),(295,34,33,0),(296,36,33,0),(297,37,33,0),(298,38,33,0),(299,39,33,0),(300,4,34,0),(301,11,34,0),(302,13,34,0),(303,43,34,0),(304,21,34,0),(305,22,34,0),(306,23,34,0),(307,44,34,0),(308,45,34,0),(309,31,34,0),(310,46,34,0),(311,33,34,0),(312,35,34,0),(313,47,34,0),(314,1,36,0),(315,40,36,0),(316,5,36,0),(317,6,36,0),(318,41,36,0),(319,42,36,0),(320,9,36,0),(321,43,36,0),(322,19,36,0),(323,20,36,0),(324,21,36,0),(325,22,36,0),(326,23,36,0),(327,44,36,0),(328,45,36,0),(329,30,36,0),(330,31,36,0),(331,32,36,0),(332,48,36,0),(333,49,36,0),(334,35,36,0),(335,36,36,0),(336,47,36,0),(337,37,36,0),(338,2,37,0),(339,3,37,0),(340,4,37,0),(341,7,37,0),(342,8,37,0),(343,10,37,0),(344,12,37,0),(345,13,37,0),(346,14,37,0),(347,15,37,0),(348,16,37,0),(349,17,37,0),(350,18,37,0),(351,24,37,0),(352,25,37,0),(353,26,37,0),(354,27,37,0),(355,29,37,0),(356,46,37,0),(357,33,37,0),(358,34,37,0),(359,39,37,0),(360,1,39,0),(361,2,39,0),(362,40,39,0),(363,3,39,0),(364,4,39,0),(365,5,39,0),(366,6,39,0),(367,41,39,0),(368,42,39,0),(369,7,39,0),(370,8,39,0),(371,9,39,0),(372,10,39,0),(373,11,39,0),(374,12,39,0),(375,13,39,0),(376,14,39,0),(377,15,39,0),(378,43,39,0),(379,16,39,0),(380,17,39,0),(381,18,39,0),(382,19,39,0),(383,20,39,0),(384,21,39,0),(385,22,39,0),(386,24,39,0),(387,25,39,0),(388,44,39,0),(389,45,39,0),(390,26,39,0),(391,27,39,0),(392,28,39,0),(393,29,39,0),(394,30,39,0),(395,31,39,0),(396,32,39,0),(397,48,39,0),(398,49,39,0),(399,33,39,0),(400,34,39,0),(401,35,39,0),(402,36,39,0),(403,47,39,0),(404,37,39,0),(405,38,39,0),(406,39,39,0),(407,23,40,0),(408,46,40,0),(409,2,42,0),(410,40,42,0),(411,42,42,0),(412,7,42,0),(413,8,42,0),(414,9,42,0),(415,10,42,0),(416,12,42,0),(417,13,42,0),(418,14,42,0),(419,15,42,0),(420,43,42,0),(421,16,42,0),(422,17,42,0),(423,18,42,0),(424,19,42,0),(425,20,42,0),(426,25,42,0),(427,26,42,0),(428,28,42,0),(429,30,42,0),(430,31,42,0),(431,32,42,0),(432,46,42,0),(433,33,42,0),(434,34,42,0),(435,35,42,0),(436,36,42,0),(437,47,42,0),(438,37,42,0),(439,38,42,0),(440,39,42,0),(441,1,43,0),(442,3,43,0),(443,4,43,0),(444,5,43,0),(445,6,43,0),(446,41,43,0),(447,11,43,0),(448,21,43,0),(449,22,43,0),(450,23,43,0),(451,24,43,0),(452,44,43,0),(453,45,43,0),(454,27,43,0),(455,29,43,0),(456,48,43,0),(457,49,43,0),(458,2,45,0),(459,3,45,0),(460,8,45,0),(461,13,45,0),(462,14,45,0),(463,17,45,0),(464,18,45,0),(465,21,45,0),(466,22,45,0),(467,23,45,0),(468,24,45,0),(469,25,45,0),(470,27,45,0),(471,31,45,0),(472,32,45,0),(473,46,45,0),(474,34,45,0),(475,35,45,0),(476,38,45,0),(477,39,45,0),(478,2,46,0),(479,7,46,0),(480,23,46,0),(481,26,46,0),(482,27,46,0),(483,45,47,0),(484,30,47,0),(485,12,48,0),(486,43,48,0),(487,19,48,0),(488,23,48,0),(489,47,48,0),(490,40,49,0),(491,4,49,0),(492,6,49,0),(493,10,49,0),(494,49,49,0),(495,3,51,0),(496,8,51,0),(497,16,51,0),(498,17,51,0),(499,19,51,0),(500,11,52,0),(501,17,52,0),(502,29,52,0),(503,2,53,0),(504,40,53,0),(505,3,53,0),(506,29,53,0),(507,34,53,0),(508,38,53,0),(509,20,54,0),(510,28,54,0),(511,4,55,0),(512,6,55,0),(513,10,55,0),(514,14,55,0),(515,15,55,0),(516,16,55,0),(517,17,55,0),(518,24,55,0),(519,27,55,0),(520,32,55,0),(521,39,55,0),(522,6,56,0),(523,14,56,0),(524,23,56,0),(525,25,56,0),(526,45,56,0),(527,30,56,0),(528,32,56,0),(529,46,56,0),(530,33,56,0),(531,1,58,0),(532,4,58,0),(533,10,58,0),(534,11,58,0),(535,14,58,0),(536,16,58,0),(537,17,58,0),(538,45,58,0),(539,30,58,0),(540,32,58,0),(541,48,58,0),(542,49,58,0),(543,33,58,0),(544,35,58,0),(545,39,58,0),(546,40,59,0),(547,3,59,0),(548,15,59,0),(549,22,59,0),(550,23,59,0),(551,29,59,0),(552,38,59,0),(553,6,60,0),(554,8,60,0),(555,20,60,0),(556,21,60,0),(557,44,60,0),(558,34,60,0),(559,37,60,0),(560,2,61,0),(561,5,61,0),(562,41,61,0),(563,42,61,0),(564,7,61,0),(565,9,61,0),(566,12,61,0),(567,13,61,0),(568,43,61,0),(569,18,61,0),(570,19,61,0),(571,24,61,0),(572,26,61,0),(573,27,61,0),(574,28,61,0),(575,31,61,0),(576,46,61,0),(577,36,61,0),(578,47,61,0),(579,14,63,0),(580,25,63,0),(581,45,63,0),(582,30,63,0),(583,46,63,0),(584,33,63,0),(585,4,64,0),(586,6,65,0),(587,32,65,0),(588,1,66,0),(589,2,66,0),(590,40,66,0),(591,3,66,0),(592,5,66,0),(593,41,66,0),(594,42,66,0),(595,7,66,0),(596,8,66,0),(597,9,66,0),(598,10,66,0),(599,11,66,0),(600,12,66,0),(601,13,66,0),(602,15,66,0),(603,43,66,0),(604,16,66,0),(605,17,66,0),(606,18,66,0),(607,19,66,0),(608,20,66,0),(609,21,66,0),(610,22,66,0),(611,23,66,0),(612,24,66,0),(613,44,66,0),(614,26,66,0),(615,27,66,0),(616,28,66,0),(617,29,66,0),(618,31,66,0),(619,48,66,0),(620,49,66,0),(621,34,66,0),(622,35,66,0),(623,36,66,0),(624,47,66,0),(625,37,66,0),(626,38,66,0),(627,39,66,0),(628,1,68,0),(629,2,68,0),(630,40,68,0),(631,3,68,0),(632,41,68,0),(633,42,68,0),(634,8,68,0),(635,9,68,0),(636,12,68,0),(637,16,68,0),(638,17,68,0),(639,18,68,0),(640,19,68,0),(641,20,68,0),(642,23,68,0),(643,24,68,0),(644,28,68,0),(645,29,68,0),(646,30,68,0),(647,48,68,0),(648,34,68,0),(649,37,68,0),(650,38,68,0),(651,4,69,0),(652,5,69,0),(653,7,69,0),(654,11,69,0),(655,13,69,0),(656,43,69,0),(657,21,69,0),(658,22,69,0),(659,25,69,0),(660,44,69,0),(661,45,69,0),(662,26,69,0),(663,27,69,0),(664,31,69,0),(665,46,69,0),(666,49,69,0),(667,35,69,0),(668,36,69,0),(669,47,69,0),(670,39,69,0),(671,6,70,0),(672,10,70,0),(673,14,70,0),(674,15,70,0),(675,32,70,0),(676,33,70,0),(677,39,70,0);
/*!40000 ALTER TABLE `mapping` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `paper`
--

DROP TABLE IF EXISTS `paper`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `paper` (
  `id_paper` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `citation` varchar(100) NOT NULL,
  `bib` mediumtext,
  `referenceCount` int(20) DEFAULT '0',
  PRIMARY KEY (`id_paper`),
  UNIQUE KEY `id_paper_UNIQUE` (`id_paper`)
) ENGINE=InnoDB AUTO_INCREMENT=50 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `paper`
--

LOCK TABLES `paper` WRITE;
/*!40000 ALTER TABLE `paper` DISABLE KEYS */;
INSERT INTO `paper` VALUES (1,'abadi2005control','{\"author\": \"Abadi, Mart{\\\\\'\\\\i}n and Budiu, Mihai and Erlingsson, Ulfar and Ligatti, Jay\", \"booktitle\": \"Proceedings of the 12th ACM conference on Computer and communications security\", \"title\": \"Control-flow integrity\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2005\", \"organization\": \"ACM\", \"ID\": \"abadi2005control\", \"pages\": \"340--353\"}',1),(2,'aucsmith1996tamper','{\"isbn\": \"3-540-61996-8\", \"title\": \"Tamper resistant software: An implementation\", \"journal\": \"Proceedings of the First International Workshop on Information Hiding\", \"author\": \"Aucsmith, David\", \"ENTRYTYPE\": \"article\", \"mendeley-groups\": \"Tamperproofing,Tamperproofing/Methods\", \"link\": \"http://link.springer.com/chapter/10.1007/3-540-61996-8{\\\\_}49\", \"year\": \"1996\", \"ID\": \"aucsmith1996tamper\", \"pages\": \"317--333\"}',2),(3,'banescu2017detecting','{\"doi\": \"10.1145/3029806.3029835\", \"isbn\": \"978-1-4503-4523-1/17/03\", \"title\": \"Detecting Patching of Executables without System Calls\", \"booktitle\": \"Proceedings of the Conference on Data and Application Security and Privacy\", \"author\": \"Banescu, Sebastian and Ahmadvand, Mohsen and Pretschner, Alexander and Shield, Robert and Hamilton, Chris\", \"ID\": \"banescu2017detecting\", \"year\": \"2017\", \"ENTRYTYPE\": \"inproceedings\"}',3),(4,'baumann2015shielding','{\"publisher\": \"ACM\", \"author\": \"Baumann, Andrew and Peinado, Marcus and Hunt, Galen\", \"journal\": \"ACM Transactions on Computer Systems (TOCS)\", \"title\": \"Shielding applications from an untrusted cloud with haven\", \"number\": \"3\", \"ENTRYTYPE\": \"article\", \"volume\": \"33\", \"year\": \"2015\", \"ID\": \"baumann2015shielding\", \"pages\": \"8\"}',4),(5,'Blietz2006','{\"doi\": \"10.1007/11787952_12\", \"isbn\": \"3540359982\", \"author\": \"Blietz, Brian and Tyagi, Akhilesh\", \"journal\": \"Lecture Notes in Computer Science (including subseries Lecture Notes in Artificial Intelligence and Lecture Notes in Bioinformatics)\", \"issn\": \"16113349\", \"ENTRYTYPE\": \"article\", \"volume\": \"3919 LNCS\", \"mendeley-groups\": \"Tamperproofing/Methods\", \"file\": \":Users/mohsen-tum/Documents/Mendeley Desktop/Blietz, Tyagi/Lecture Notes in Computer Science (including subseries Lecture Notes in Artificial Intelligence and Lecture Notes in Bioinformatics)/Blietz, Tyagi - 2006 - Software tamper resistance through dynamic program monitoring.pdf:pdf\", \"year\": \"2006\", \"title\": \"Software tamper resistance through dynamic program monitoring\", \"ID\": \"Blietz2006\", \"pages\": \"146--163\"}',5),(6,'brasser2015tytan','{\"author\": \"Brasser, Ferdinand and El Mahjoub, Brahim and Sadeghi, Ahmad-Reza and Wachsmann, Christian and Koeberl, Patrick\", \"booktitle\": \"2015 52nd ACM/EDAC/IEEE Design Automation Conference (DAC)\", \"title\": \"TyTAN: tiny trust anchor for tiny devices\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2015\", \"organization\": \"IEEE\", \"ID\": \"brasser2015tytan\", \"pages\": \"1--6\"}',6),(7,'catuogno2002format','{\"author\": \"Catuogno, Luigi and Visconti, Ivan\", \"booktitle\": \"International Conference on Security in Communication Networks\", \"title\": \"A format-independent architecture for run-time integrity checking of executable code\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2002\", \"organization\": \"Springer\", \"ID\": \"catuogno2002format\", \"pages\": \"219--233\"}',7),(8,'chang2001protecting','{\"author\": \"Chang, Hoi and Atallah, Mikhail J\", \"booktitle\": \"ACM Workshop on Digital Rights Management\", \"title\": \"Protecting software code by guards\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2001\", \"organization\": \"Springer\", \"ID\": \"chang2001protecting\", \"pages\": \"160--175\"}',8),(9,'chen2002oblivious','{\"author\": \"Chen, Yuqun and Venkatesan, Ramarathnam and Cary, Matthew and Pang, Ruoming and Sinha, Saurabh and Jakubowski, Mariusz H\", \"booktitle\": \"International Workshop on Information Hiding\", \"title\": \"Oblivious hashing: A stealthy software integrity verification primitive\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2002\", \"organization\": \"Springer\", \"ID\": \"chen2002oblivious\", \"pages\": \"400--414\"}',9),(10,'christodorescu2009cloud','{\"author\": \"Christodorescu, Mihai and Sailer, Reiner and Schales, Douglas Lee and Sgandurra, Daniele and Zamboni, Diego\", \"booktitle\": \"Proceedings of the 2009 ACM workshop on Cloud computing security\", \"title\": \"Cloud security is not (just) virtualization security: a short paper\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2009\", \"organization\": \"ACM\", \"ID\": \"christodorescu2009cloud\", \"pages\": \"97--102\"}',10),(11,'collberg2012distributed','{\"author\": \"Collberg, Christian and Martin, Sam and Myers, Jonathan and Nagra, Jasvir\", \"booktitle\": \"Proceedings of the 28th Annual Computer Security Applications Conference\", \"title\": \"Distributed application tamper detection via continuous software updates\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2012\", \"organization\": \"ACM\", \"ID\": \"collberg2012distributed\", \"pages\": \"319--328\"}',11),(12,'dedic2007graph','{\"author\": \"Dedi{\\\\\'c}, Nenad and Jakubowski, Mariusz and Venkatesan, Ramarathnam\", \"booktitle\": \"International Workshop on Information Hiding\", \"title\": \"A graph game model for software tamper protection\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2007\", \"organization\": \"Springer\", \"ID\": \"dedic2007graph\", \"pages\": \"80--95\"}',12),(13,'deswarte2004remote','{\"publisher\": \"Springer\", \"author\": \"Deswarte, Yves and Quisquater, Jean-Jacques and Sa{\\\\\\\"\\\\i}dane, Ayda\", \"booktitle\": \"Integrity and internal control in information systems VI\", \"title\": \"Remote integrity checking\", \"ENTRYTYPE\": \"incollection\", \"year\": \"2004\", \"ID\": \"deswarte2004remote\", \"pages\": \"1--11\"}',13),(14,'dewan2008hypervisor','{\"author\": \"Dewan, Prashant and Durham, David and Khosravi, Hormuzd and Long, Men and Nagabhushan, Gayathri\", \"booktitle\": \"Proceedings of the 2008 Spring simulation multiconference\", \"title\": \"A hypervisor-based system for protecting software runtime memory and persistent storage\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2008\", \"organization\": \"Society for Computer Simulation International\", \"ID\": \"dewan2008hypervisor\", \"pages\": \"828--835\"}',14),(15,'Gan2015using','{\"doi\": \"10.1109/SPRO.2015.12\", \"title\": \"Using Virtual Machine Protections to Enhance Whitebox Cryptography\", \"booktitle\": \"Software Protection (SPRO), 2015 IEEE/ACM 1st International Workshop on\", \"author\": \"J. Gan and R. Kok and P. Kohli and Y. Ding and B. Mah\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2015\", \"ID\": \"Gan2015using\", \"pages\": \"17-23\"}',15),(16,'Ghosh2010secure','{\"isbn\": \"364216434X\", \"author\": \"Ghosh, Sudeep and Hiser, Jason D. and Davidson, Jack W.\", \"journal\": \"Lecture Notes in Computer Science (including subseries Lecture Notes in Artificial Intelligence and Lecture Notes in Bioinformatics)\", \"issn\": \"03029743\", \"ENTRYTYPE\": \"article\", \"volume\": \"6387 LNCS\", \"mendeley-groups\": \"Tamperproofing\", \"year\": \"2010\", \"title\": \"A secure and robust approach to software tamper resistance\", \"ID\": \"Ghosh2010secure\", \"pages\": \"33--47\"}',16),(17,'ghosh2013software','{\"author\": \"Ghosh, Sudeep and Hiser, Jason and Davidson, Jack W\", \"booktitle\": \"Proceedings of the 2nd ACM SIGPLAN Program Protection and Reverse Engineering Workshop\", \"title\": \"Software protection for dynamically-generated code\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2013\", \"organization\": \"ACM\", \"ID\": \"ghosh2013software\", \"pages\": \"1\"}',17),(18,'Horne2002','{\"doi\": \"10.1007/3-540-47870-1_9\", \"isbn\": \"978-3-540-43677-5\", \"author\": \"Horne, Bill and Matheson, Lesley and Sheehan, Casey and Tarjan, Robert\", \"ENTRYTYPE\": \"article\", \"abstract\": \"We describe a software self-checking mechanism designed to improve the tamper resistance of large programs. The mechanism consists of a number of testers that redundantly test for changes in the executable code as it is running and report modifications. The mechanism is built to be compatible with copy-specific static watermarking and other tamper-resistance techniques. The mechanism includes several innovations to make it stealthy and more robust.\", \"title\": \"Dynamic Self-Checking Techniques for Improved Tamper Resistance\", \"pages\": \"141--159\", \"mendeley-groups\": \"Tamperproofing/Methods,Tamperproofing\", \"link\": \"http://citeseerx.ist.psu.edu/viewdoc/summary?doi=10.1.1.13.3308\", \"file\": \":Users/mohsen-tum/Documents/Mendeley Desktop/Horne et al/Security and Privacy in Digital Rights Management/Horne et al. - 2002 - Dynamic Self-Checking Techniques for Improved Tamper Resistance.pdf:pdf\", \"year\": \"2002\", \"ID\": \"Horne2002\", \"annote\": \"They add testers in the post compilation process.\\nLinear checks no circular\\nTo avoid complexity, a block is checked only by one block\\nA 32bit space is added outside basic blocks as corrector that tries to fix the hash values in patch process. The patch process is part of sofware watermarking after-installation process\\nDid not quite get it where do they store hashes? They say we store them but not clear where?!\\nNo inidication of how Address space layout randomization is respected.\", \"journal\": \"Security and Privacy in Digital Rights Management\"}',18),(19,'ibrahim2016stins4cs','{\"author\": \"Ibrahim, Amjad and Banescu, Sebastian\", \"booktitle\": \"Proceedings of the 2016 ACM Workshop on Software PROtection\", \"title\": \"StIns4CS: A State Inspection Tool for C\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2016\", \"organization\": \"ACM\", \"ID\": \"ibrahim2016stins4cs\", \"pages\": \"61--71\"}',19),(20,'jacob2007towards','{\"author\": \"Jacob, Matthias and Jakubowski, Mariusz H and Venkatesan, Ramarathnam\", \"booktitle\": \"Proceedings of the 9th workshop on Multimedia \\\\& security\", \"title\": \"Towards integral binary execution: Implementing oblivious hashing using overlapped instruction encodings\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2007\", \"organization\": \"ACM\", \"ID\": \"jacob2007towards\", \"pages\": \"129--140\"}',20),(21,'jakobsson2010retroactive','{\"numpages\": \"13\", \"publisher\": \"USENIX Association\", \"title\": \"Retroactive Detection of Malware with Applications to Mobile Platforms\", \"series\": \"HotSec\'10\", \"booktitle\": \"Proceedings of the 5th USENIX Conference on Hot Topics in Security\", \"author\": \"Jakobsson, Markus and Johansson, Karl-Anders\", \"ENTRYTYPE\": \"inproceedings\", \"location\": \"Washinton, DC\", \"year\": \"2010\", \"ID\": \"jakobsson2010retroactive\", \"pages\": \"1--13\", \"address\": \"Berkeley, CA, USA\"}',21),(22,'jakobsson2011practical','{\"author\": \"Jakobsson, Markus and Johansson, Karl-Anders\", \"booktitle\": \"Lightweight Security \\\\& Privacy: Devices, Protocols and Applications (LightSec), 2011 Workshop on\", \"title\": \"Practical and secure software-based attestation\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2011\", \"organization\": \"IEEE\", \"ID\": \"jakobsson2011practical\", \"pages\": \"1--9\"}',22),(23,'jin2003forensic','{\"author\": \"Jin, Hongxia and Lotspiech, Jeffery\", \"booktitle\": \"Software Reliability Engineering, 2003. ISSRE 2003. 14th International Symposium on\", \"title\": \"Forensic analysis for tamper resistant software\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2003\", \"organization\": \"IEEE\", \"ID\": \"jin2003forensic\", \"pages\": \"133--142\"}',23),(24,'junod2015obfuscator','{\"author\": \"Junod, Pascal and Rinaldini, Julien and Wehrli, Johan and Michielin, Julie\", \"booktitle\": \"Proceedings of the 1st International Workshop on Software Protection\", \"title\": \"Obfuscator-LLVM: software protection for the masses\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2015\", \"organization\": \"IEEE Press\", \"ID\": \"junod2015obfuscator\", \"pages\": \"3--9\"}',24),(25,'kanstren2015architecture','{\"author\": \"Kanstr{\\\\\'e}n, Teemu and Lehtonen, Sami and Savola, Reijo and Kukkohovi, Hilkka and H{\\\\\\\"a}t{\\\\\\\"o}nen, Kimmo\", \"booktitle\": \"Cloud Engineering (IC2E), 2015 IEEE International Conference on\", \"title\": \"Architecture for high confidence cloud security monitoring\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2015\", \"organization\": \"IEEE\", \"ID\": \"kanstren2015architecture\", \"pages\": \"195--200\"}',25),(26,'kim1994experiences','{\"ID\": \"kim1994experiences\", \"author\": \"Kim, Gene H and Spafford, Eugene H\", \"year\": \"1994\", \"ENTRYTYPE\": \"article\", \"title\": \"Experiences with tripwire: Using integrity checkers for intrusion detection\"}',26),(27,'kimball2012emulation','{\"publisher\": \"Google Patents\", \"author\": \"Kimball, William B and Baldwin, Rusty O\", \"title\": \"Emulation-based software protection\", \"month\": \"oct~9\", \"note\": \"US Patent 8,285,987\", \"year\": \"2012\", \"ID\": \"kimball2012emulation\", \"ENTRYTYPE\": \"misc\"}',27),(28,'kulkarni2014new','{\"author\": \"Kulkarni, Aniket and Metta, Ravindra\", \"booktitle\": \"Service Oriented System Engineering (SOSE), 2014 IEEE 8th International Symposium on\", \"title\": \"A New Code Obfuscation Scheme for Software Protection\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2014\", \"organization\": \"IEEE\", \"ID\": \"kulkarni2014new\", \"pages\": \"409--414\"}',28),(29,'madou2005software','{\"author\": \"Madou, Matias and Anckaert, Bertrand and Moseley, Patrick and Debray, Saumya and De Sutter, Bjorn and De Bosschere, Koen\", \"booktitle\": \"International Workshop on Information Security Applications\", \"title\": \"Software protection through dynamic code mutation\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2005\", \"organization\": \"Springer\", \"ID\": \"madou2005software\", \"pages\": \"194--206\"}',29),(30,'Malone2011','{\"doi\": \"10.1145/2046582.2046596\", \"isbn\": \"9781450310017\", \"keyword\": \"hardware performance counters,integrity\", \"author\": \"Malone, Corey and Zahran, Mohamed and Karri, Ramesh\", \"journal\": \"Proceedings of the sixth ACM workshop on Scalable trusted computing - STC \'11\", \"issn\": \"15437221\", \"ENTRYTYPE\": \"article\", \"mendeley-groups\": \"Tamperproofing/Methods\", \"link\": \"http://www.scopus.com/inward/record.url?eid=2-s2.0-80755143408{\\\\&}partnerID=40{\\\\&}md5=ad5db1f8e5c0131a2a17f457ba1b0497$\\\\backslash$nhttp://dl.acm.org/citation.cfm?doid=2046582.2046596\", \"file\": \":Users/mohsen-tum/Documents/Mendeley Desktop/Malone, Zahran, Karri/Proceedings of the sixth ACM workshop on Scalable trusted computing - STC \'11/Malone, Zahran, Karri - 2011 - Are Hardware Performance Counters a Cost Effective Way for Integrity Checking of Programs.pdf:pdf\", \"year\": \"2011\", \"title\": \"Are Hardware Performance Counters a Cost Effective Way for Integrity Checking of Programs\", \"ID\": \"Malone2011\", \"pages\": \"71\"}',30),(31,'Martignoni2010conquer','{\"doi\": \"10.1007/978-3-642-14215-4_2\", \"title\": \"Conqueror: Tamper-proof code execution on legacy systems\", \"journal\": \"Lecture Notes in Computer Science (including subseries Lecture Notes in Artificial Intelligence and Lecture Notes in Bioinformatics)\", \"author\": \"Martignoni, Lorenzo and Paleari, Roberto and Bruschi, Danilo\", \"ENTRYTYPE\": \"article\", \"volume\": \"6201 LNCS\", \"year\": \"2010\", \"ID\": \"Martignoni2010conquer\", \"pages\": \"21--40\"}',31),(32,'morgan2015design','{\"author\": \"Morgan, Beno{\\\\^\\\\i}t and Alata, Eric and Nicomette, Vincent and Ka{\\\\^a}niche, Mohamed and Averlant, Guillaume\", \"booktitle\": \"Dependable Computing (PRDC), 2015 IEEE 21st Pacific Rim International Symposium on\", \"title\": \"Design and implementation of a hardware assisted security architecture for software integrity monitoring\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2015\", \"organization\": \"IEEE\", \"ID\": \"morgan2015design\", \"pages\": \"189--198\"}',32),(33,'park2015tgvisor','{\"author\": \"Park, Sungjin and Yoon, Jae Nam and Kang, Cheoloh and Kim, Kyong Hoon and Han, Taisook\", \"booktitle\": \"Mobile Cloud Computing, Services, and Engineering (MobileCloud), 2015 3rd IEEE International Conference on\", \"title\": \"TGVisor: A Tiny Hypervisor-Based Trusted Geolocation Framework for Mobile Cloud Clients\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2015\", \"organization\": \"IEEE\", \"ID\": \"park2015tgvisor\", \"pages\": \"99--108\"}',33),(34,'Protsenko2015dynamic','{\"doi\": \"10.1109/ARES.2015.98\", \"keyword\": \"Android (operating system);computer crime;cryptography;mobile computing;reverse engineering;Android apps;application piracy;dynamic code loading;dynamic obfuscation techniques;dynamic re-encryption;dynamic self-protection;mobile devices;native code;proprietary mobile software;reverse engineering;tamperproofing;Androids;Encryption;Humanoid robots;Loading;Runtime;Software protection;Android;Software Protection\", \"title\": \"Dynamic Self-Protection and Tamperproofing for Android Apps Using Native Code\", \"booktitle\": \"Availability, Reliability and Security (ARES), 2015 10th International Conference on\", \"author\": \"M. Protsenko and S. Kreuter and T. M\\u00fcller\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2015\", \"ID\": \"Protsenko2015dynamic\", \"pages\": \"129-138\"}',34),(35,'Seshadri2005pioneer','{\"doi\": \"10.1145/1095809.1095812\", \"isbn\": \"1-59593-079-5\", \"keyword\": \"dynamic root of trust,rootkit detection,self-check-summing code,software-based code attestation,verifiable code execution\", \"author\": \"Seshadri, Arvind and Luk, Mark and Shi, Elaine and Perrig, Adrian and van Doorn, Leendert and Khosla, Pradeep\", \"journal\": \"ACM SIGOPS Operating Systems Review\", \"issn\": \"01635980\", \"ID\": \"Seshadri2005pioneer\", \"mendeley-groups\": \"Tamperproofing\", \"link\": \"http://dl.acm.org/citation.cfm?id=1095809.1095812\", \"year\": \"2005\", \"title\": \"Pioneer: Verifying Code Integrity and Enforcing Untampered Code Execution on Legacy Systems\", \"ENTRYTYPE\": \"article\"}',35),(36,'Spinellis2000','{\"doi\": \"10.1145/353323.353383\", \"isbn\": \"1094-9224\", \"author\": \"Spinellis, Diomidis\", \"ENTRYTYPE\": \"article\", \"abstract\": \"The integrity verification of a device\'s controlling software is an important aspect of many emerging information appliances. We propose the use of reflection, whereby the software is able to examine its own operation, in conjunction with cryptographic hashes as a basis for developing a suitable software verification protocol. For more demanding applications meta-reflective techniques can be used to thwart attacks based on device emulation strategies. We demonstrate how our approach can be used to increase the security of mobile phones, devices for the delivery of digital content, and smartcards.\", \"issn\": \"10949224\", \"number\": \"1\", \"pages\": \"51--62\", \"volume\": \"3\", \"mendeley-groups\": \"Tamperproofing/Methods\", \"file\": \":Users/mohsen-tum/Documents/Mendeley Desktop/Spinellis/ACM Transactions on Information and System Security/Spinellis - 2000 - Reflection as a mechanism for software integrity verification.pdf:pdf\", \"year\": \"2000\", \"title\": \"Reflection as a mechanism for software integrity verification\", \"ID\": \"Spinellis2000\", \"annote\": \"In this approach a software integrity is verified with the help of an external (trusted) entity. Here, the program state is retrieved using reflection, a protocol is proposed to verify the state, and suggested to augment the scheme with CPU perfor.mance counter, before and after the verification call loops.\\nOne obvious attack is to keep an untouched version of the application in the memory next to the tampered with version. Then redirect all hash computations to the good version. The authors, suggest memory expanion and timing as possible countermeasures.\", \"journal\": \"ACM Transactions on Information and System Security\"}',36),(37,'teixeira2015siot','{\"author\": \"Teixeira, Fernando A and Machado, Gustavo V and Pereira, Fernando MQ and Wong, Hao Chi and Nogueira, Jos{\\\\\'e} and Oliveira, Leonardo B\", \"booktitle\": \"Proceedings of the 14th International Conference on Information Processing in Sensor Networks\", \"title\": \"SIoT: securing the internet of things through distributed system analysis\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2015\", \"organization\": \"ACM\", \"ID\": \"teixeira2015siot\", \"pages\": \"310--321\"}',37),(38,'Wang2005Tamper','{\"isbn\": \"8242866627\", \"keyword\": \"integrity checking,multi-blocking encryption,software piracy,tamper resistant\", \"title\": \"Tamper Resistant Software Through Dynamic Integrity Checking\", \"journal\": \"Proc. Symp. on Cyptography and Information Security (SCIS 05)\", \"author\": \"Wang, Ping and Kang, Seok-kyu and Kim, Kwangjo\", \"ID\": \"Wang2005Tamper\", \"year\": \"2005\", \"ENTRYTYPE\": \"article\"}',38),(39,'yao2014cryptvmi','{\"author\": \"Yao, Fangzhou and Sprabery, Read and Campbell, Roy H\", \"booktitle\": \"Proceedings of the 2nd international workshop on Security in cloud computing\", \"title\": \"CryptVMI: a flexible and encrypted virtual machine introspection system in the cloud\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2014\", \"organization\": \"ACM\", \"ID\": \"yao2014cryptvmi\", \"pages\": \"11--18\"}',39),(40,'banescu2015software','{\"author\": \"Banescu, Sebastian and Pretschner, Alexander and Battr{\\\\\'e}, Dominic and Cazzulani, St{\\\\\'e}fano and Shield, Robert and Thompson, Greg\", \"booktitle\": \"Proceedings of the 5th ACM Conference on Data and Application Security and Privacy\", \"title\": \"Software-based protection against changeware\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2015\", \"organization\": \"ACM\", \"ID\": \"banescu2015software\", \"pages\": \"231--242\"}',40),(41,'Carbone2009','{\"isbn\": \"9781605583525\", \"author\": \"Carbone, Martim and Cui, Weidong and Peinado, Marcus and Lu, Long and Lee, Wenke\", \"journal\": \"Analysis\", \"title\": \"Mapping Kernel Objects to Enable Systematic Integrity Checking\", \"ENTRYTYPE\": \"article\", \"mendeley-groups\": \"Tamperproofing/Methods\", \"file\": \":Users/mohsen-tum/Documents/Mendeley Desktop/Carbone et al/Analysis/Carbone et al. - 2009 - Mapping Kernel Objects to Enable Systematic Integrity Checking.pdf:pdf\", \"year\": \"2009\", \"ID\": \"Carbone2009\", \"pages\": \"555--565\"}',41),(42,'Castro2006','{\"isbn\": \"1-931971-47-1\", \"author\": \"Castro, Miguel and Costa, Manuel and Harris, Tim\", \"ENTRYTYPE\": \"article\", \"abstract\": \"Software attacks often subvert the intended data-flow in a vulnerable program. For example, attackers exploit buffer overflows and format string vulnerabilities to write data to unintended locations. We present a simple technique that prevents these attacks by enforcing data-flow integrity. It computes a data-flow graph using static analysis, and it instruments the program to ensure that the flow of data at runtime is allowed by the data-flow graph. We describe an efficient implementation of data-flow integrity enforcement that uses static analysis to reduce instrumentation overhead. This implementation can be used in practice to detect a broad class of attacks and errors because it can be applied automatically to C and C++ programs without modifications, it does not have false positives, and it has low overhead.\", \"title\": \"Securing software by enforcing data-flow integrity\", \"pages\": \"147--160\", \"mendeley-groups\": \"Tamperproofing/Methods\", \"link\": \"http://dl.acm.org/citation.cfm?id=1298455.1298470$\\\\backslash$nhttp://www.usenix.org/event/osdi06/tech/full{\\\\_}papers/castro/castro{\\\\_}html/\", \"file\": \":Users/mohsen-tum/Documents/Mendeley Desktop/Castro, Costa, Harris/Proceedings of the 7th symposium on Operating systems design and implementation/Castro, Costa, Harris - 2006 - Securing software by enforcing data-flow integrity.pdf:pdf\", \"year\": \"2006\", \"ID\": \"Castro2006\", \"journal\": \"Proceedings of the 7th symposium on Operating systems design and implementation\"}',42),(43,'gao2015integrity','{\"doi\": \"10.1109/ICAC.2015.34\", \"keyword\": \"Big Data;cloud computing;data integrity;data privacy;Big Data processing;cloud computing technology;dynamic redundancy computation;integrity protection solution;reputation based redundancy computation;Conferences;MapReduce;cloud computing;integrity protection\", \"title\": \"Integrity Protection for Big Data Processing with Dynamic Redundancy Computation\", \"booktitle\": \"Autonomic Computing (ICAC), 2015 IEEE International Conference on\", \"author\": \"Z. Gao and N. Desalvo and P. D. Khoa and S. H. Kim and L. Xu and W. W. Ro and R. M. Verma and W. Shi\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2015\", \"ID\": \"gao2015integrity\", \"pages\": \"159-160\"}',43),(44,'karapanos2016verena','{\"author\": \"Karapanos, Nikolaos and Filios, Alexandros and Popa, Raluca Ada and Capkun, Srdjan\", \"booktitle\": \"Proceedings of the 37th IEEE Symposium on Security and Privacy (IEEE S\\\\&P)\", \"title\": \"Verena: End-to-end integrity protection for web applications\", \"ID\": \"karapanos2016verena\", \"year\": \"2016\", \"ENTRYTYPE\": \"inproceedings\"}',44),(45,'Kil2009','{\"isbn\": \"9781424444212\", \"keyword\": \"dynamic attestation,integrity,remote attestation,runtime,system security,trusted computing\", \"author\": \"Kil, Chongkyung\", \"journal\": \"IEEE/IFIP International Conference on Dependable Systems {\\\\&} Networks\", \"title\": \"Remote Attestation to Dynamic System Properties: Towards Providing Complete System Integrity Evidence\", \"ENTRYTYPE\": \"article\", \"mendeley-groups\": \"Tamperproofing\", \"file\": \":Users/mohsen-tum/Documents/Mendeley Desktop/Kil/IEEEIFIP International Conference on Dependable Systems {\\\\&} Networks/Kil - 2009 - Remote Attestation to Dynamic System Properties Towards Providing Complete System Integrity Evidence.pdf:pdf\", \"year\": \"2009\", \"ID\": \"Kil2009\", \"pages\": \"115--124\"}',45),(46,'neisse2011implementing','{\"author\": \"Neisse, Ricardo and Holling, Dominik and Pretschner, Alexander\", \"booktitle\": \"Proceedings of the 2011 11th IEEE/ACM International Symposium on Cluster, Cloud and Grid Computing\", \"title\": \"Implementing trust in cloud infrastructures\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2011\", \"organization\": \"IEEE Computer Society\", \"ID\": \"neisse2011implementing\", \"pages\": \"524--533\"}',46),(47,'sun2015security','{\"author\": \"Sun, Yuqiong and Nanda, Susanta and Jaeger, Trent\", \"booktitle\": \"2015 IEEE 7th International Conference on Cloud Computing Technology and Science (CloudCom)\", \"title\": \"Security-as-a-Service for Microservices-Based Cloud Applications\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2015\", \"organization\": \"IEEE\", \"ID\": \"sun2015security\", \"pages\": \"50--57\"}',47),(48,'pappas2012smashing','{\"author\": \"Pappas, Vasilis and Polychronakis, Michalis and Keromytis, Angelos D\", \"booktitle\": \"2012 IEEE Symposium on Security and Privacy\", \"title\": \"Smashing the gadgets: Hindering return-oriented programming using in-place code randomization\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2012\", \"organization\": \"IEEE\", \"ID\": \"pappas2012smashing\", \"pages\": \"601--615\"}',48),(49,'pappas2013transparent','{\"author\": \"Pappas, Vasilis and Polychronakis, Michalis and Keromytis, Angelos D\", \"booktitle\": \"Presented as part of the 22nd USENIX Security Symposium (USENIX Security 13)\", \"title\": \"Transparent ROP exploit mitigation using indirect branch tracing\", \"pages\": \"447--462\", \"year\": \"2013\", \"ID\": \"pappas2013transparent\", \"ENTRYTYPE\": \"inproceedings\"}',49);
/*!40000 ALTER TABLE `paper` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Temporary view structure for view `paper_attribute`
--

DROP TABLE IF EXISTS `paper_attribute`;
/*!50001 DROP VIEW IF EXISTS `paper_attribute`*/;
SET @saved_cs_client     = @@character_set_client;
SET character_set_client = utf8;
/*!50001 CREATE VIEW `paper_attribute` AS SELECT 
 1 AS `id_taxonomy`,
 1 AS `id_paper`,
 1 AS `citation`,
 1 AS `bib`,
 1 AS `id_attribute`,
 1 AS `text_attribute`,
 1 AS `leaf_attribute`*/;
SET character_set_client = @saved_cs_client;

--
-- Temporary view structure for view `paper_merged_attributes`
--

DROP TABLE IF EXISTS `paper_merged_attributes`;
/*!50001 DROP VIEW IF EXISTS `paper_merged_attributes`*/;
SET @saved_cs_client     = @@character_set_client;
SET character_set_client = utf8;
/*!50001 CREATE VIEW `paper_merged_attributes` AS SELECT 
 1 AS `id_taxonomy`,
 1 AS `id_paper`,
 1 AS `citation`,
 1 AS `bib`,
 1 AS `atts`,
 1 AS `leaf_atts`*/;
SET character_set_client = @saved_cs_client;

--
-- Table structure for table `relation`
--

DROP TABLE IF EXISTS `relation`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `relation` (
  `id_relation` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `text` varchar(50) NOT NULL,
  `comment` varchar(500) DEFAULT NULL,
  PRIMARY KEY (`id_relation`),
  UNIQUE KEY `id_relation_UNIQUE` (`id_relation`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `relation`
--

LOCK TABLES `relation` WRITE;
/*!40000 ALTER TABLE `relation` DISABLE KEYS */;
INSERT INTO `relation` VALUES (1,'Depends','simple dependency'),(2,'DependsDirected','directed dependency'),(3,'InstanceOf','inheritance'),(4,'MemberOf','aggregation'),(5,'PartOf','composition');
/*!40000 ALTER TABLE `relation` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `taxonomy`
--

DROP TABLE IF EXISTS `taxonomy`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `taxonomy` (
  `id_taxonomy` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `text` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id_taxonomy`),
  UNIQUE KEY `id_taxonomy_UNIQUE` (`id_taxonomy`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `taxonomy`
--

LOCK TABLES `taxonomy` WRITE;
/*!40000 ALTER TABLE `taxonomy` DISABLE KEYS */;
INSERT INTO `taxonomy` VALUES (1,'Integrity protection');
/*!40000 ALTER TABLE `taxonomy` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `taxonomy_dimension`
--

DROP TABLE IF EXISTS `taxonomy_dimension`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `taxonomy_dimension` (
  `id_taxonomy_dimension` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `id_taxonomy` int(11) NOT NULL,
  `id_attribute` int(11) unsigned NOT NULL,
  `id_dimension` int(11) unsigned NOT NULL,
  PRIMARY KEY (`id_taxonomy_dimension`),
  UNIQUE KEY `id_taxonomy_dimension_UNIQUE` (`id_taxonomy_dimension`),
  UNIQUE KEY `taxonomy_dimension_id_attribute_UNIQUE` (`id_taxonomy`,`id_attribute`),
  KEY `taxonomy_dimension_id_attribute_foreign` (`id_attribute`),
  KEY `taxonomy_dimension_id_dimension_foreign` (`id_dimension`),
  CONSTRAINT `taxonomy_dimension_id_attribute_foreign` FOREIGN KEY (`id_attribute`) REFERENCES `attribute` (`id_attribute`) ON DELETE CASCADE ON UPDATE NO ACTION,
  CONSTRAINT `taxonomy_dimension_id_dimension_foreign` FOREIGN KEY (`id_dimension`) REFERENCES `dimension` (`id_dimension`) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=92 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `taxonomy_dimension`
--

LOCK TABLES `taxonomy_dimension` WRITE;
/*!40000 ALTER TABLE `taxonomy_dimension` DISABLE KEYS */;
INSERT INTO `taxonomy_dimension` VALUES (1,1,2,1),(2,1,3,1),(3,1,4,1),(4,1,6,1),(5,1,7,1),(6,1,8,1),(7,1,9,1),(8,1,10,1),(9,1,11,1),(10,1,12,1),(11,1,13,1),(12,1,15,1),(13,1,16,1),(14,1,17,1),(15,1,18,1),(16,1,19,1),(17,1,21,1),(18,1,22,1),(19,1,23,1),(20,1,24,1),(21,1,25,1),(22,1,28,2),(23,1,29,2),(24,1,30,2),(25,1,31,2),(26,1,33,3),(27,1,34,3),(28,1,36,3),(29,1,37,3),(30,1,39,3),(31,1,40,3),(32,1,42,3),(33,1,43,3),(34,1,45,3),(35,1,46,3),(36,1,47,3),(37,1,48,3),(38,1,49,3),(39,1,51,3),(40,1,52,3),(41,1,53,3),(42,1,54,3),(43,1,55,3),(44,1,56,3),(45,1,58,3),(46,1,59,3),(47,1,60,3),(48,1,61,3),(49,1,63,3),(50,1,64,3),(51,1,65,3),(52,1,66,3),(53,1,68,3),(54,1,69,3),(55,1,70,3),(56,1,5,1),(57,1,71,3),(58,1,72,2),(59,1,73,2),(60,1,78,2),(61,1,79,2),(62,1,80,2),(63,1,81,2),(64,1,82,2),(65,1,83,2),(66,1,84,2),(67,1,85,2),(68,1,86,2),(69,1,87,2),(70,1,88,2),(71,1,89,3),(72,1,90,3),(73,1,91,3),(74,1,92,3),(75,1,93,1),(76,1,94,1),(77,1,95,1),(78,1,1,1),(79,1,14,1),(80,1,20,1),(81,1,26,2),(82,1,27,2),(83,1,32,3),(84,1,35,3),(85,1,38,3),(86,1,41,3),(87,1,44,3),(88,1,50,3),(89,1,57,3),(90,1,62,3),(91,1,67,3);
/*!40000 ALTER TABLE `taxonomy_dimension` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `taxonomy_relation`
--

DROP TABLE IF EXISTS `taxonomy_relation`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `taxonomy_relation` (
  `id_taxonomy_relation` int(11) NOT NULL AUTO_INCREMENT,
  `id_taxonomy` int(11) NOT NULL,
  `id_src_attribute` int(11) unsigned NOT NULL,
  `id_dest_attribute` int(11) unsigned NOT NULL,
  `id_relation` int(11) NOT NULL,
  `id_dimension` int(10) DEFAULT '0',
  `edgeBendPoints` longtext,
  PRIMARY KEY (`id_taxonomy_relation`),
  UNIQUE KEY `id_taxonomy_relation_UNIQUE` (`id_taxonomy_relation`),
  UNIQUE KEY `taxonomy_relation_attributes_UNIQUE` (`id_taxonomy`,`id_src_attribute`,`id_dest_attribute`,`id_dimension`),
  KEY `taxonomy_relation_id_src_attribute_foreign` (`id_src_attribute`),
  KEY `taxonomy_relation_id_dest_attribute_foreign` (`id_dest_attribute`),
  CONSTRAINT `taxonomy_relation_id_dest_attribute_foreign` FOREIGN KEY (`id_dest_attribute`) REFERENCES `attribute` (`id_attribute`) ON DELETE CASCADE ON UPDATE NO ACTION,
  CONSTRAINT `taxonomy_relation_id_src_attribute_foreign` FOREIGN KEY (`id_src_attribute`) REFERENCES `attribute` (`id_attribute`) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=246 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `taxonomy_relation`
--

LOCK TABLES `taxonomy_relation` WRITE;
/*!40000 ALTER TABLE `taxonomy_relation` DISABLE KEYS */;
INSERT INTO `taxonomy_relation` VALUES (56,1,2,93,3,1,NULL),(57,1,3,93,3,1,NULL),(58,1,4,93,3,1,NULL),(59,1,5,93,1,1,NULL),(60,1,21,20,3,1,NULL),(61,1,22,20,3,1,NULL),(62,1,23,20,3,1,NULL),(63,1,24,20,3,1,NULL),(64,1,25,20,3,1,NULL),(65,1,94,20,3,1,NULL),(66,1,6,21,1,1,NULL),(67,1,6,22,1,1,NULL),(68,1,6,94,1,1,NULL),(69,1,6,5,5,1,NULL),(70,1,7,5,5,1,NULL),(71,1,8,5,5,1,NULL),(72,1,9,7,3,1,NULL),(73,1,10,7,3,1,NULL),(74,1,13,8,3,1,NULL),(75,1,11,8,3,1,NULL),(76,1,12,11,3,1,NULL),(77,1,7,24,1,1,NULL),(78,1,7,25,1,1,NULL),(79,1,8,24,1,1,NULL),(80,1,8,25,1,1,NULL),(81,1,5,14,1,1,NULL),(82,1,15,14,3,1,NULL),(83,1,95,14,3,1,NULL),(84,1,17,14,3,1,NULL),(85,1,18,14,3,1,NULL),(86,1,19,14,3,1,NULL),(87,1,79,80,4,2,NULL),(88,1,81,80,4,2,NULL),(89,1,82,80,4,2,NULL),(90,1,83,80,4,2,NULL),(91,1,80,72,5,2,NULL),(92,1,73,72,5,2,NULL),(93,1,84,72,5,2,NULL),(94,1,27,72,5,2,NULL),(95,1,28,27,3,2,NULL),(96,1,29,27,3,2,NULL),(97,1,30,27,3,2,NULL),(98,1,88,84,3,2,NULL),(99,1,31,27,3,2,NULL),(100,1,26,73,3,2,NULL),(101,1,85,84,3,2,NULL),(102,1,86,84,3,2,NULL),(103,1,87,84,3,2,NULL),(104,1,78,27,3,2,NULL),(105,1,57,32,5,3,NULL),(106,1,62,32,5,3,NULL),(107,1,67,32,5,3,NULL),(108,1,35,32,5,3,NULL),(109,1,38,32,5,3,NULL),(110,1,41,32,5,3,NULL),(111,1,44,32,5,3,NULL),(112,1,50,32,5,3,NULL),(113,1,33,32,3,3,NULL),(114,1,34,32,3,3,NULL),(115,1,58,57,3,3,NULL),(116,1,59,57,3,3,NULL),(117,1,60,57,3,3,NULL),(118,1,61,57,3,3,NULL),(119,1,71,62,3,3,NULL),(120,1,89,62,3,3,NULL),(121,1,63,62,3,3,NULL),(122,1,64,62,3,3,NULL),(123,1,65,62,3,3,NULL),(124,1,70,67,3,3,NULL),(125,1,69,67,3,3,NULL),(126,1,90,67,3,3,NULL),(127,1,36,35,3,3,NULL),(128,1,37,35,3,3,NULL),(129,1,39,38,3,3,NULL),(130,1,92,38,3,3,NULL),(131,1,42,41,3,3,NULL),(132,1,43,41,3,3,NULL),(133,1,45,44,3,3,NULL),(134,1,46,44,3,3,NULL),(135,1,47,44,3,3,NULL),(136,1,48,44,3,3,NULL),(137,1,49,44,3,3,NULL),(138,1,51,50,3,3,NULL),(139,1,52,50,3,3,NULL),(140,1,53,50,3,3,NULL),(141,1,54,50,3,3,NULL),(142,1,55,50,3,3,NULL),(143,1,91,50,3,3,NULL),(232,1,32,84,2,4,NULL),(233,1,32,20,2,4,NULL),(234,1,32,27,2,4,NULL),(235,1,32,67,1,4,NULL),(236,1,57,20,2,4,NULL),(237,1,62,32,2,4,NULL),(238,1,93,14,1,4,NULL),(239,1,5,93,2,4,NULL),(240,1,84,32,2,4,NULL),(241,1,80,72,2,4,NULL),(242,1,27,5,2,4,NULL),(243,1,27,32,2,4,NULL),(244,1,73,27,2,4,NULL),(245,1,84,5,2,4,NULL);
/*!40000 ALTER TABLE `taxonomy_relation` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `taxonomy_relation_annotation`
--

DROP TABLE IF EXISTS `taxonomy_relation_annotation`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `taxonomy_relation_annotation` (
  `id_taxonomy` int(11) NOT NULL,
  `id_taxonomy_relation` int(11) NOT NULL,
  `annotation` longtext,
  PRIMARY KEY (`id_taxonomy_relation`),
  UNIQUE KEY `id_taxonomy_relation_annotation_UNIQUE` (`id_taxonomy_relation`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `taxonomy_relation_annotation`
--

LOCK TABLES `taxonomy_relation_annotation` WRITE;
/*!40000 ALTER TABLE `taxonomy_relation_annotation` DISABLE KEYS */;
INSERT INTO `taxonomy_relation_annotation` VALUES (1,232,'Impedes'),(1,233,'Transforms'),(1,234,'Mitigates or raises the bar'),(1,236,'Affects'),(1,237,'Strengthens'),(1,239,'Contains'),(1,240,'Identifies'),(1,241,'Support'),(1,242,'Tampers with'),(1,243,'Tampers with'),(1,244,'Executes'),(1,245,'Identifies asset');
/*!40000 ALTER TABLE `taxonomy_relation_annotation` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Final view structure for view `paper_attribute`
--

/*!50001 DROP VIEW IF EXISTS `paper_attribute`*/;
/*!50001 SET @saved_cs_client          = @@character_set_client */;
/*!50001 SET @saved_cs_results         = @@character_set_results */;
/*!50001 SET @saved_col_connection     = @@collation_connection */;
/*!50001 SET character_set_client      = utf8 */;
/*!50001 SET character_set_results     = utf8 */;
/*!50001 SET collation_connection      = utf8_general_ci */;
/*!50001 CREATE ALGORITHM=UNDEFINED */
/*!50013 DEFINER=`root`@`localhost` SQL SECURITY DEFINER */
/*!50001 VIEW `paper_attribute` AS select `taxonomy`.`id_taxonomy` AS `id_taxonomy`,`paper`.`id_paper` AS `id_paper`,`paper`.`citation` AS `citation`,`paper`.`bib` AS `bib`,`mapping`.`id_attribute` AS `id_attribute`,`GETATTRIBUTENAME`(`mapping`.`id_attribute`) AS `text_attribute`,`GETATTRIBUTENAME`(`mapping`.`id_attribute`) AS `leaf_attribute` from ((`paper` join `mapping` on((`paper`.`id_paper` = `mapping`.`id_paper`))) join `taxonomy`) union select `rel1`.`id_taxonomy` AS `id_taxonomy`,`paper`.`id_paper` AS `id_paper`,`paper`.`citation` AS `citation`,`paper`.`bib` AS `bib`,`rel1`.`id_src_attribute` AS `id_attribute`,`GETATTRIBUTENAME`(`rel1`.`id_src_attribute`) AS `text_attribute`,NULL AS `leaf_attribute` from ((`paper` join `mapping` on((`paper`.`id_paper` = `mapping`.`id_paper`))) left join `taxonomy_relation` `rel1` on((`rel1`.`id_dest_attribute` = `mapping`.`id_attribute`))) where (`rel1`.`id_src_attribute` is not null) union select `rel2`.`id_taxonomy` AS `id_taxonomy`,`paper`.`id_paper` AS `id_paper`,`paper`.`citation` AS `citation`,`paper`.`bib` AS `bib`,`rel2`.`id_src_attribute` AS `id_attribute`,`GETATTRIBUTENAME`(`rel2`.`id_src_attribute`) AS `text_attribute`,NULL AS `leaf_attribute` from (((`paper` join `mapping` on((`paper`.`id_paper` = `mapping`.`id_paper`))) left join `taxonomy_relation` `rel1` on((`rel1`.`id_dest_attribute` = `mapping`.`id_attribute`))) left join `taxonomy_relation` `rel2` on((`rel2`.`id_dest_attribute` = `rel1`.`id_src_attribute`))) where (`rel2`.`id_src_attribute` is not null) union select `rel3`.`id_taxonomy` AS `id_taxonomy`,`paper`.`id_paper` AS `id_paper`,`paper`.`citation` AS `citation`,`paper`.`bib` AS `bib`,`rel3`.`id_src_attribute` AS `id_attribute`,`GETATTRIBUTENAME`(`rel3`.`id_src_attribute`) AS `text_attribute`,NULL AS `leaf_attribute` from ((((`paper` join `mapping` on((`paper`.`id_paper` = `mapping`.`id_paper`))) left join `taxonomy_relation` `rel1` on((`rel1`.`id_dest_attribute` = `mapping`.`id_attribute`))) left join `taxonomy_relation` `rel2` on((`rel2`.`id_dest_attribute` = `rel1`.`id_src_attribute`))) left join `taxonomy_relation` `rel3` on((`rel3`.`id_dest_attribute` = `rel2`.`id_src_attribute`))) where (`rel3`.`id_src_attribute` is not null) */;
/*!50001 SET character_set_client      = @saved_cs_client */;
/*!50001 SET character_set_results     = @saved_cs_results */;
/*!50001 SET collation_connection      = @saved_col_connection */;

--
-- Final view structure for view `paper_merged_attributes`
--

/*!50001 DROP VIEW IF EXISTS `paper_merged_attributes`*/;
/*!50001 SET @saved_cs_client          = @@character_set_client */;
/*!50001 SET @saved_cs_results         = @@character_set_results */;
/*!50001 SET @saved_col_connection     = @@collation_connection */;
/*!50001 SET character_set_client      = utf8 */;
/*!50001 SET character_set_results     = utf8 */;
/*!50001 SET collation_connection      = utf8_general_ci */;
/*!50001 CREATE ALGORITHM=UNDEFINED */
/*!50013 DEFINER=`root`@`localhost` SQL SECURITY DEFINER */
/*!50001 VIEW `paper_merged_attributes` AS select distinct `paper_attribute`.`id_taxonomy` AS `id_taxonomy`,`paper_attribute`.`id_paper` AS `id_paper`,`paper_attribute`.`citation` AS `citation`,`paper_attribute`.`bib` AS `bib`,`att_table`.`atts` AS `atts`,`att_table`.`leaf_atts` AS `leaf_atts` from (`classification`.`paper_attribute` join (select `a`.`id_paper` AS `id_paper`,group_concat(concat(`a`.`text_attribute`) separator ',') AS `atts`,group_concat(concat(`a`.`leaf_attribute`) separator ',') AS `leaf_atts` from (select `paper_attribute`.`id_paper` AS `id_paper`,`paper_attribute`.`id_attribute` AS `id_attribute`,`paper_attribute`.`text_attribute` AS `text_attribute`,`paper_attribute`.`leaf_attribute` AS `leaf_attribute` from `classification`.`paper_attribute` order by `paper_attribute`.`id_attribute`) `a` group by `a`.`id_paper`) `att_table` on((`att_table`.`id_paper` = `paper_attribute`.`id_paper`))) */;
/*!50001 SET character_set_client      = @saved_cs_client */;
/*!50001 SET character_set_results     = @saved_cs_results */;
/*!50001 SET collation_connection      = @saved_col_connection */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2017-12-28 21:09:17
