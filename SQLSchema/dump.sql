-- MySQL dump 10.13  Distrib 5.7.12, for osx10.10 (x86_64)
--
-- Host: localhost    Database: classification
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
-- Table structure for table `attribute`
--

DROP TABLE IF EXISTS `attribute`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `attribute` (
  `id_attribute` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `text` varchar(50) NOT NULL,
  PRIMARY KEY (`id_attribute`),
  UNIQUE KEY `id_attribute_UNIQUE` (`id_attribute`)
) ENGINE=InnoDB AUTO_INCREMENT=71 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `attribute`
--

LOCK TABLES `attribute` WRITE;
/*!40000 ALTER TABLE `attribute` DISABLE KEYS */;
INSERT INTO `attribute` VALUES (1,'Integrity Protection Assets'),(2,'Behavior'),(3,'Data'),(4,'Data and behavior'),(5,'Representation'),(6,'Static'),(7,'In memory'),(8,'In execution'),(9,'Code invariants'),(10,'Data invariants'),(11,'Trace'),(12,'Timed trace'),(13,'HW counters'),(14,'Granularity'),(15,'Instructions'),(16,'BB'),(17,'Function'),(18,'Slice'),(19,'Application'),(20,'Lifecycle activity'),(21,'Pre-compile'),(22,'Compile'),(23,'Post-compile'),(24,'Load'),(25,'Run'),(26,'Not root'),(27,'Attack'),(28,'Binary '),(29,'Process memory'),(30,'Runtime data'),(31,'Control flow'),(32,'Measure'),(33,'Local'),(34,'Remote'),(35,'Monitor'),(36,'State inspection'),(37,'Introspection'),(38,'Response'),(39,'Proactive'),(40,'Postmortem'),(41,'Transformation'),(42,'Manual'),(43,'Automatic'),(44,'Check'),(45,'Checksum'),(46,'Signature'),(47,'Equation eval'),(48,'Majority vote'),(49,'Access control'),(50,'Hardening'),(51,'Cyclic checks'),(52,'Mutation'),(53,'Code concealment'),(54,'Cloning'),(55,'Layered interpretation'),(56,'Block chain'),(57,'Overhead'),(58,'Fair'),(59,'Medium'),(60,'High'),(61,'N/A'),(62,'Trust anchor'),(63,'TPM'),(64,'SGX'),(65,'Other'),(66,'None'),(67,'Protection level'),(68,'Internal'),(69,'External'),(70,'Hypervisor');
/*!40000 ALTER TABLE `attribute` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `dimension`
--

DROP TABLE IF EXISTS `dimension`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `dimension` (
  `id_dimension` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `text` varchar(50) NOT NULL,
  PRIMARY KEY (`id_dimension`),
  UNIQUE KEY `id_dimension_UNIQUE` (`id_dimension`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `dimension`
--

LOCK TABLES `dimension` WRITE;
/*!40000 ALTER TABLE `dimension` DISABLE KEYS */;
INSERT INTO `dimension` VALUES (1,'System view'),(2,'Attack view'),(3,'Defense view');
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
  `id_attribute` int(11) NOT NULL,
  PRIMARY KEY (`id_mapping`),
  UNIQUE KEY `id_mapping_UNIQUE` (`id_mapping`)
) ENGINE=InnoDB AUTO_INCREMENT=678 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `mapping`
--

LOCK TABLES `mapping` WRITE;
/*!40000 ALTER TABLE `mapping` DISABLE KEYS */;
INSERT INTO `mapping` VALUES (1,1,2),(2,2,2),(3,3,2),(4,4,2),(5,5,2),(6,6,2),(7,7,2),(8,8,2),(9,9,2),(10,10,2),(11,11,2),(12,12,2),(13,13,2),(14,14,2),(15,15,2),(16,16,2),(17,17,2),(18,18,2),(19,19,2),(20,20,2),(21,21,2),(22,22,2),(23,23,2),(24,24,2),(25,25,2),(26,26,2),(27,27,2),(28,28,2),(29,29,2),(30,30,2),(31,31,2),(32,32,2),(33,33,2),(34,34,2),(35,35,2),(36,36,2),(37,37,2),(38,38,2),(39,39,2),(40,40,3),(41,4,3),(42,6,3),(43,41,3),(44,42,3),(45,14,3),(46,43,3),(47,44,3),(48,45,3),(49,46,3),(50,47,3),(51,4,4),(52,6,4),(53,14,4),(54,7,6),(55,11,6),(56,12,6),(57,13,6),(58,26,6),(59,46,6),(60,33,6),(61,37,6),(62,2,9),(63,3,9),(64,4,9),(65,6,9),(66,8,9),(67,14,9),(68,16,9),(69,17,9),(70,18,9),(71,24,9),(72,25,9),(73,27,9),(74,29,9),(75,32,9),(76,34,9),(77,36,9),(78,38,9),(79,39,9),(80,41,10),(81,42,10),(82,43,10),(83,19,10),(84,44,10),(85,45,10),(86,47,10),(87,1,11),(88,40,11),(89,5,11),(90,9,11),(91,20,11),(92,23,11),(93,28,11),(94,48,11),(95,49,11),(96,21,12),(97,22,12),(98,31,12),(99,35,12),(100,36,12),(101,30,13),(102,36,13),(103,1,15),(104,40,15),(105,5,15),(106,42,15),(107,2,16),(108,3,16),(109,8,16),(110,12,16),(111,16,16),(112,17,16),(113,20,16),(114,48,16),(115,38,16),(116,6,17),(117,11,17),(118,15,17),(119,43,17),(120,19,17),(121,44,17),(122,28,17),(123,31,17),(124,33,17),(125,34,17),(126,35,17),(127,36,17),(128,40,18),(129,41,18),(130,9,18),(131,18,18),(132,20,18),(133,24,18),(134,45,18),(135,29,18),(136,37,18),(137,4,19),(138,7,19),(139,10,19),(140,13,19),(141,14,19),(142,21,19),(143,22,19),(144,23,19),(145,25,19),(146,26,19),(147,27,19),(148,32,19),(149,46,19),(150,49,19),(151,47,19),(152,39,19),(153,40,21),(154,41,21),(155,42,21),(156,14,21),(157,19,21),(158,23,21),(159,44,21),(160,45,21),(161,28,21),(162,36,21),(163,9,22),(164,24,22),(165,37,22),(166,1,23),(167,2,23),(168,3,23),(169,4,23),(170,6,23),(171,7,23),(172,8,23),(173,10,23),(174,13,23),(175,15,23),(176,43,23),(177,16,23),(178,17,23),(179,18,23),(180,20,23),(181,21,23),(182,22,23),(183,25,23),(184,26,23),(185,27,23),(186,30,23),(187,31,23),(188,32,23),(189,48,23),(190,49,23),(191,33,23),(192,34,23),(193,35,23),(194,38,23),(195,39,23),(196,17,24),(197,11,25),(198,29,25),(199,1,26),(200,40,26),(201,3,26),(202,6,26),(203,41,26),(204,42,26),(205,7,26),(206,10,26),(207,13,26),(208,43,26),(209,45,26),(210,26,26),(211,27,26),(212,30,26),(213,48,26),(214,49,26),(215,33,26),(216,47,26),(217,37,26),(218,7,28),(219,11,28),(220,12,28),(221,13,28),(222,26,28),(223,46,28),(224,37,28),(225,2,29),(226,3,29),(227,4,29),(228,6,29),(229,8,29),(230,14,29),(231,16,29),(232,17,29),(233,18,29),(234,21,29),(235,22,29),(236,24,29),(237,25,29),(238,27,29),(239,29,29),(240,32,29),(241,33,29),(242,34,29),(243,36,29),(244,38,29),(245,39,29),(246,40,30),(247,41,30),(248,42,30),(249,9,30),(250,43,30),(251,19,30),(252,44,30),(253,45,30),(254,47,30),(255,1,31),(256,40,31),(257,5,31),(258,9,31),(259,20,31),(260,23,31),(261,28,31),(262,30,31),(263,48,31),(264,49,31),(265,1,33),(266,2,33),(267,40,33),(268,3,33),(269,5,33),(270,6,33),(271,41,33),(272,42,33),(273,7,33),(274,8,33),(275,9,33),(276,10,33),(277,12,33),(278,14,33),(279,15,33),(280,16,33),(281,17,33),(282,18,33),(283,19,33),(284,20,33),(285,24,33),(286,25,33),(287,26,33),(288,27,33),(289,28,33),(290,29,33),(291,30,33),(292,32,33),(293,48,33),(294,49,33),(295,34,33),(296,36,33),(297,37,33),(298,38,33),(299,39,33),(300,4,34),(301,11,34),(302,13,34),(303,43,34),(304,21,34),(305,22,34),(306,23,34),(307,44,34),(308,45,34),(309,31,34),(310,46,34),(311,33,34),(312,35,34),(313,47,34),(314,1,36),(315,40,36),(316,5,36),(317,6,36),(318,41,36),(319,42,36),(320,9,36),(321,43,36),(322,19,36),(323,20,36),(324,21,36),(325,22,36),(326,23,36),(327,44,36),(328,45,36),(329,30,36),(330,31,36),(331,32,36),(332,48,36),(333,49,36),(334,35,36),(335,36,36),(336,47,36),(337,37,36),(338,2,37),(339,3,37),(340,4,37),(341,7,37),(342,8,37),(343,10,37),(344,12,37),(345,13,37),(346,14,37),(347,15,37),(348,16,37),(349,17,37),(350,18,37),(351,24,37),(352,25,37),(353,26,37),(354,27,37),(355,29,37),(356,46,37),(357,33,37),(358,34,37),(359,39,37),(360,1,39),(361,2,39),(362,40,39),(363,3,39),(364,4,39),(365,5,39),(366,6,39),(367,41,39),(368,42,39),(369,7,39),(370,8,39),(371,9,39),(372,10,39),(373,11,39),(374,12,39),(375,13,39),(376,14,39),(377,15,39),(378,43,39),(379,16,39),(380,17,39),(381,18,39),(382,19,39),(383,20,39),(384,21,39),(385,22,39),(386,24,39),(387,25,39),(388,44,39),(389,45,39),(390,26,39),(391,27,39),(392,28,39),(393,29,39),(394,30,39),(395,31,39),(396,32,39),(397,48,39),(398,49,39),(399,33,39),(400,34,39),(401,35,39),(402,36,39),(403,47,39),(404,37,39),(405,38,39),(406,39,39),(407,23,40),(408,46,40),(409,2,42),(410,40,42),(411,42,42),(412,7,42),(413,8,42),(414,9,42),(415,10,42),(416,12,42),(417,13,42),(418,14,42),(419,15,42),(420,43,42),(421,16,42),(422,17,42),(423,18,42),(424,19,42),(425,20,42),(426,25,42),(427,26,42),(428,28,42),(429,30,42),(430,31,42),(431,32,42),(432,46,42),(433,33,42),(434,34,42),(435,35,42),(436,36,42),(437,47,42),(438,37,42),(439,38,42),(440,39,42),(441,1,43),(442,3,43),(443,4,43),(444,5,43),(445,6,43),(446,41,43),(447,11,43),(448,21,43),(449,22,43),(450,23,43),(451,24,43),(452,44,43),(453,45,43),(454,27,43),(455,29,43),(456,48,43),(457,49,43),(458,2,45),(459,3,45),(460,8,45),(461,13,45),(462,14,45),(463,17,45),(464,18,45),(465,21,45),(466,22,45),(467,23,45),(468,24,45),(469,25,45),(470,27,45),(471,31,45),(472,32,45),(473,46,45),(474,34,45),(475,35,45),(476,38,45),(477,39,45),(478,2,46),(479,7,46),(480,23,46),(481,26,46),(482,27,46),(483,45,47),(484,30,47),(485,12,48),(486,43,48),(487,19,48),(488,23,48),(489,47,48),(490,40,49),(491,4,49),(492,6,49),(493,10,49),(494,49,49),(495,3,51),(496,8,51),(497,16,51),(498,17,51),(499,19,51),(500,11,52),(501,17,52),(502,29,52),(503,2,53),(504,40,53),(505,3,53),(506,29,53),(507,34,53),(508,38,53),(509,20,54),(510,28,54),(511,4,55),(512,6,55),(513,10,55),(514,14,55),(515,15,55),(516,16,55),(517,17,55),(518,24,55),(519,27,55),(520,32,55),(521,39,55),(522,6,56),(523,14,56),(524,23,56),(525,25,56),(526,45,56),(527,30,56),(528,32,56),(529,46,56),(530,33,56),(531,1,58),(532,4,58),(533,10,58),(534,11,58),(535,14,58),(536,16,58),(537,17,58),(538,45,58),(539,30,58),(540,32,58),(541,48,58),(542,49,58),(543,33,58),(544,35,58),(545,39,58),(546,40,59),(547,3,59),(548,15,59),(549,22,59),(550,23,59),(551,29,59),(552,38,59),(553,6,60),(554,8,60),(555,20,60),(556,21,60),(557,44,60),(558,34,60),(559,37,60),(560,2,61),(561,5,61),(562,41,61),(563,42,61),(564,7,61),(565,9,61),(566,12,61),(567,13,61),(568,43,61),(569,18,61),(570,19,61),(571,24,61),(572,26,61),(573,27,61),(574,28,61),(575,31,61),(576,46,61),(577,36,61),(578,47,61),(579,14,63),(580,25,63),(581,45,63),(582,30,63),(583,46,63),(584,33,63),(585,4,64),(586,6,65),(587,32,65),(588,1,66),(589,2,66),(590,40,66),(591,3,66),(592,5,66),(593,41,66),(594,42,66),(595,7,66),(596,8,66),(597,9,66),(598,10,66),(599,11,66),(600,12,66),(601,13,66),(602,15,66),(603,43,66),(604,16,66),(605,17,66),(606,18,66),(607,19,66),(608,20,66),(609,21,66),(610,22,66),(611,23,66),(612,24,66),(613,44,66),(614,26,66),(615,27,66),(616,28,66),(617,29,66),(618,31,66),(619,48,66),(620,49,66),(621,34,66),(622,35,66),(623,36,66),(624,47,66),(625,37,66),(626,38,66),(627,39,66),(628,1,68),(629,2,68),(630,40,68),(631,3,68),(632,41,68),(633,42,68),(634,8,68),(635,9,68),(636,12,68),(637,16,68),(638,17,68),(639,18,68),(640,19,68),(641,20,68),(642,23,68),(643,24,68),(644,28,68),(645,29,68),(646,30,68),(647,48,68),(648,34,68),(649,37,68),(650,38,68),(651,4,69),(652,5,69),(653,7,69),(654,11,69),(655,13,69),(656,43,69),(657,21,69),(658,22,69),(659,25,69),(660,44,69),(661,45,69),(662,26,69),(663,27,69),(664,31,69),(665,46,69),(666,49,69),(667,35,69),(668,36,69),(669,47,69),(670,39,69),(671,6,70),(672,10,70),(673,14,70),(674,15,70),(675,32,70),(676,33,70),(677,39,70);
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
  PRIMARY KEY (`id_paper`),
  UNIQUE KEY `id_paper_UNIQUE` (`id_paper`)
) ENGINE=InnoDB AUTO_INCREMENT=50 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `paper`
--

LOCK TABLES `paper` WRITE;
/*!40000 ALTER TABLE `paper` DISABLE KEYS */;
INSERT INTO `paper` VALUES (1,'abadi2005control','{\"author\": \"Abadi, Mart{\\\\\'\\\\i}n and Budiu, Mihai and Erlingsson, Ulfar and Ligatti, Jay\", \"booktitle\": \"Proceedings of the 12th ACM conference on Computer and communications security\", \"title\": \"Control-flow integrity\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2005\", \"organization\": \"ACM\", \"ID\": \"abadi2005control\", \"pages\": \"340--353\"}'),(2,'aucsmith1996tamper','{\"isbn\": \"3-540-61996-8\", \"title\": \"Tamper resistant software: An implementation\", \"journal\": \"Proceedings of the First International Workshop on Information Hiding\", \"author\": \"Aucsmith, David\", \"ENTRYTYPE\": \"article\", \"mendeley-groups\": \"Tamperproofing,Tamperproofing/Methods\", \"link\": \"http://link.springer.com/chapter/10.1007/3-540-61996-8{\\\\_}49\", \"year\": \"1996\", \"ID\": \"aucsmith1996tamper\", \"pages\": \"317--333\"}'),(3,'banescu2017detecting','{\"doi\": \"10.1145/3029806.3029835\", \"isbn\": \"978-1-4503-4523-1/17/03\", \"title\": \"Detecting Patching of Executables without System Calls\", \"booktitle\": \"Proceedings of the Conference on Data and Application Security and Privacy\", \"author\": \"Banescu, Sebastian and Ahmadvand, Mohsen and Pretschner, Alexander and Shield, Robert and Hamilton, Chris\", \"ID\": \"banescu2017detecting\", \"year\": \"2017\", \"ENTRYTYPE\": \"inproceedings\"}'),(4,'baumann2015shielding','{\"publisher\": \"ACM\", \"author\": \"Baumann, Andrew and Peinado, Marcus and Hunt, Galen\", \"journal\": \"ACM Transactions on Computer Systems (TOCS)\", \"title\": \"Shielding applications from an untrusted cloud with haven\", \"number\": \"3\", \"ENTRYTYPE\": \"article\", \"volume\": \"33\", \"year\": \"2015\", \"ID\": \"baumann2015shielding\", \"pages\": \"8\"}'),(5,'Blietz2006','{\"doi\": \"10.1007/11787952_12\", \"isbn\": \"3540359982\", \"author\": \"Blietz, Brian and Tyagi, Akhilesh\", \"journal\": \"Lecture Notes in Computer Science (including subseries Lecture Notes in Artificial Intelligence and Lecture Notes in Bioinformatics)\", \"issn\": \"16113349\", \"ENTRYTYPE\": \"article\", \"volume\": \"3919 LNCS\", \"mendeley-groups\": \"Tamperproofing/Methods\", \"file\": \":Users/mohsen-tum/Documents/Mendeley Desktop/Blietz, Tyagi/Lecture Notes in Computer Science (including subseries Lecture Notes in Artificial Intelligence and Lecture Notes in Bioinformatics)/Blietz, Tyagi - 2006 - Software tamper resistance through dynamic program monitoring.pdf:pdf\", \"year\": \"2006\", \"title\": \"Software tamper resistance through dynamic program monitoring\", \"ID\": \"Blietz2006\", \"pages\": \"146--163\"}'),(6,'brasser2015tytan','{\"author\": \"Brasser, Ferdinand and El Mahjoub, Brahim and Sadeghi, Ahmad-Reza and Wachsmann, Christian and Koeberl, Patrick\", \"booktitle\": \"2015 52nd ACM/EDAC/IEEE Design Automation Conference (DAC)\", \"title\": \"TyTAN: tiny trust anchor for tiny devices\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2015\", \"organization\": \"IEEE\", \"ID\": \"brasser2015tytan\", \"pages\": \"1--6\"}'),(7,'catuogno2002format','{\"author\": \"Catuogno, Luigi and Visconti, Ivan\", \"booktitle\": \"International Conference on Security in Communication Networks\", \"title\": \"A format-independent architecture for run-time integrity checking of executable code\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2002\", \"organization\": \"Springer\", \"ID\": \"catuogno2002format\", \"pages\": \"219--233\"}'),(8,'chang2001protecting','{\"author\": \"Chang, Hoi and Atallah, Mikhail J\", \"booktitle\": \"ACM Workshop on Digital Rights Management\", \"title\": \"Protecting software code by guards\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2001\", \"organization\": \"Springer\", \"ID\": \"chang2001protecting\", \"pages\": \"160--175\"}'),(9,'chen2002oblivious','{\"author\": \"Chen, Yuqun and Venkatesan, Ramarathnam and Cary, Matthew and Pang, Ruoming and Sinha, Saurabh and Jakubowski, Mariusz H\", \"booktitle\": \"International Workshop on Information Hiding\", \"title\": \"Oblivious hashing: A stealthy software integrity verification primitive\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2002\", \"organization\": \"Springer\", \"ID\": \"chen2002oblivious\", \"pages\": \"400--414\"}'),(10,'christodorescu2009cloud','{\"author\": \"Christodorescu, Mihai and Sailer, Reiner and Schales, Douglas Lee and Sgandurra, Daniele and Zamboni, Diego\", \"booktitle\": \"Proceedings of the 2009 ACM workshop on Cloud computing security\", \"title\": \"Cloud security is not (just) virtualization security: a short paper\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2009\", \"organization\": \"ACM\", \"ID\": \"christodorescu2009cloud\", \"pages\": \"97--102\"}'),(11,'collberg2012distributed','{\"author\": \"Collberg, Christian and Martin, Sam and Myers, Jonathan and Nagra, Jasvir\", \"booktitle\": \"Proceedings of the 28th Annual Computer Security Applications Conference\", \"title\": \"Distributed application tamper detection via continuous software updates\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2012\", \"organization\": \"ACM\", \"ID\": \"collberg2012distributed\", \"pages\": \"319--328\"}'),(12,'dedic2007graph','{\"author\": \"Dedi{\\\\\'c}, Nenad and Jakubowski, Mariusz and Venkatesan, Ramarathnam\", \"booktitle\": \"International Workshop on Information Hiding\", \"title\": \"A graph game model for software tamper protection\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2007\", \"organization\": \"Springer\", \"ID\": \"dedic2007graph\", \"pages\": \"80--95\"}'),(13,'deswarte2004remote','{\"publisher\": \"Springer\", \"author\": \"Deswarte, Yves and Quisquater, Jean-Jacques and Sa{\\\\\\\"\\\\i}dane, Ayda\", \"booktitle\": \"Integrity and internal control in information systems VI\", \"title\": \"Remote integrity checking\", \"ENTRYTYPE\": \"incollection\", \"year\": \"2004\", \"ID\": \"deswarte2004remote\", \"pages\": \"1--11\"}'),(14,'dewan2008hypervisor','{\"author\": \"Dewan, Prashant and Durham, David and Khosravi, Hormuzd and Long, Men and Nagabhushan, Gayathri\", \"booktitle\": \"Proceedings of the 2008 Spring simulation multiconference\", \"title\": \"A hypervisor-based system for protecting software runtime memory and persistent storage\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2008\", \"organization\": \"Society for Computer Simulation International\", \"ID\": \"dewan2008hypervisor\", \"pages\": \"828--835\"}'),(15,'Gan2015using','{\"doi\": \"10.1109/SPRO.2015.12\", \"title\": \"Using Virtual Machine Protections to Enhance Whitebox Cryptography\", \"booktitle\": \"Software Protection (SPRO), 2015 IEEE/ACM 1st International Workshop on\", \"author\": \"J. Gan and R. Kok and P. Kohli and Y. Ding and B. Mah\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2015\", \"ID\": \"Gan2015using\", \"pages\": \"17-23\"}'),(16,'Ghosh2010secure','{\"isbn\": \"364216434X\", \"author\": \"Ghosh, Sudeep and Hiser, Jason D. and Davidson, Jack W.\", \"journal\": \"Lecture Notes in Computer Science (including subseries Lecture Notes in Artificial Intelligence and Lecture Notes in Bioinformatics)\", \"issn\": \"03029743\", \"ENTRYTYPE\": \"article\", \"volume\": \"6387 LNCS\", \"mendeley-groups\": \"Tamperproofing\", \"year\": \"2010\", \"title\": \"A secure and robust approach to software tamper resistance\", \"ID\": \"Ghosh2010secure\", \"pages\": \"33--47\"}'),(17,'ghosh2013software','{\"author\": \"Ghosh, Sudeep and Hiser, Jason and Davidson, Jack W\", \"booktitle\": \"Proceedings of the 2nd ACM SIGPLAN Program Protection and Reverse Engineering Workshop\", \"title\": \"Software protection for dynamically-generated code\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2013\", \"organization\": \"ACM\", \"ID\": \"ghosh2013software\", \"pages\": \"1\"}'),(18,'Horne2002','{\"doi\": \"10.1007/3-540-47870-1_9\", \"isbn\": \"978-3-540-43677-5\", \"author\": \"Horne, Bill and Matheson, Lesley and Sheehan, Casey and Tarjan, Robert\", \"ENTRYTYPE\": \"article\", \"abstract\": \"We describe a software self-checking mechanism designed to improve the tamper resistance of large programs. The mechanism consists of a number of testers that redundantly test for changes in the executable code as it is running and report modifications. The mechanism is built to be compatible with copy-specific static watermarking and other tamper-resistance techniques. The mechanism includes several innovations to make it stealthy and more robust.\", \"title\": \"Dynamic Self-Checking Techniques for Improved Tamper Resistance\", \"pages\": \"141--159\", \"mendeley-groups\": \"Tamperproofing/Methods,Tamperproofing\", \"link\": \"http://citeseerx.ist.psu.edu/viewdoc/summary?doi=10.1.1.13.3308\", \"file\": \":Users/mohsen-tum/Documents/Mendeley Desktop/Horne et al/Security and Privacy in Digital Rights Management/Horne et al. - 2002 - Dynamic Self-Checking Techniques for Improved Tamper Resistance.pdf:pdf\", \"year\": \"2002\", \"ID\": \"Horne2002\", \"annote\": \"They add testers in the post compilation process.\\nLinear checks no circular\\nTo avoid complexity, a block is checked only by one block\\nA 32bit space is added outside basic blocks as corrector that tries to fix the hash values in patch process. The patch process is part of sofware watermarking after-installation process\\nDid not quite get it where do they store hashes? They say we store them but not clear where?!\\nNo inidication of how Address space layout randomization is respected.\", \"journal\": \"Security and Privacy in Digital Rights Management\"}'),(19,'ibrahim2016stins4cs','{\"author\": \"Ibrahim, Amjad and Banescu, Sebastian\", \"booktitle\": \"Proceedings of the 2016 ACM Workshop on Software PROtection\", \"title\": \"StIns4CS: A State Inspection Tool for C\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2016\", \"organization\": \"ACM\", \"ID\": \"ibrahim2016stins4cs\", \"pages\": \"61--71\"}'),(20,'jacob2007towards','{\"author\": \"Jacob, Matthias and Jakubowski, Mariusz H and Venkatesan, Ramarathnam\", \"booktitle\": \"Proceedings of the 9th workshop on Multimedia \\\\& security\", \"title\": \"Towards integral binary execution: Implementing oblivious hashing using overlapped instruction encodings\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2007\", \"organization\": \"ACM\", \"ID\": \"jacob2007towards\", \"pages\": \"129--140\"}'),(21,'jakobsson2010retroactive','{\"numpages\": \"13\", \"publisher\": \"USENIX Association\", \"title\": \"Retroactive Detection of Malware with Applications to Mobile Platforms\", \"series\": \"HotSec\'10\", \"booktitle\": \"Proceedings of the 5th USENIX Conference on Hot Topics in Security\", \"author\": \"Jakobsson, Markus and Johansson, Karl-Anders\", \"ENTRYTYPE\": \"inproceedings\", \"location\": \"Washinton, DC\", \"year\": \"2010\", \"ID\": \"jakobsson2010retroactive\", \"pages\": \"1--13\", \"address\": \"Berkeley, CA, USA\"}'),(22,'jakobsson2011practical','{\"author\": \"Jakobsson, Markus and Johansson, Karl-Anders\", \"booktitle\": \"Lightweight Security \\\\& Privacy: Devices, Protocols and Applications (LightSec), 2011 Workshop on\", \"title\": \"Practical and secure software-based attestation\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2011\", \"organization\": \"IEEE\", \"ID\": \"jakobsson2011practical\", \"pages\": \"1--9\"}'),(23,'jin2003forensic','{\"author\": \"Jin, Hongxia and Lotspiech, Jeffery\", \"booktitle\": \"Software Reliability Engineering, 2003. ISSRE 2003. 14th International Symposium on\", \"title\": \"Forensic analysis for tamper resistant software\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2003\", \"organization\": \"IEEE\", \"ID\": \"jin2003forensic\", \"pages\": \"133--142\"}'),(24,'junod2015obfuscator','{\"author\": \"Junod, Pascal and Rinaldini, Julien and Wehrli, Johan and Michielin, Julie\", \"booktitle\": \"Proceedings of the 1st International Workshop on Software Protection\", \"title\": \"Obfuscator-LLVM: software protection for the masses\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2015\", \"organization\": \"IEEE Press\", \"ID\": \"junod2015obfuscator\", \"pages\": \"3--9\"}'),(25,'kanstren2015architecture','{\"author\": \"Kanstr{\\\\\'e}n, Teemu and Lehtonen, Sami and Savola, Reijo and Kukkohovi, Hilkka and H{\\\\\\\"a}t{\\\\\\\"o}nen, Kimmo\", \"booktitle\": \"Cloud Engineering (IC2E), 2015 IEEE International Conference on\", \"title\": \"Architecture for high confidence cloud security monitoring\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2015\", \"organization\": \"IEEE\", \"ID\": \"kanstren2015architecture\", \"pages\": \"195--200\"}'),(26,'kim1994experiences','{\"ID\": \"kim1994experiences\", \"author\": \"Kim, Gene H and Spafford, Eugene H\", \"year\": \"1994\", \"ENTRYTYPE\": \"article\", \"title\": \"Experiences with tripwire: Using integrity checkers for intrusion detection\"}'),(27,'kimball2012emulation','{\"publisher\": \"Google Patents\", \"author\": \"Kimball, William B and Baldwin, Rusty O\", \"title\": \"Emulation-based software protection\", \"month\": \"oct~9\", \"note\": \"US Patent 8,285,987\", \"year\": \"2012\", \"ID\": \"kimball2012emulation\", \"ENTRYTYPE\": \"misc\"}'),(28,'kulkarni2014new','{\"author\": \"Kulkarni, Aniket and Metta, Ravindra\", \"booktitle\": \"Service Oriented System Engineering (SOSE), 2014 IEEE 8th International Symposium on\", \"title\": \"A New Code Obfuscation Scheme for Software Protection\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2014\", \"organization\": \"IEEE\", \"ID\": \"kulkarni2014new\", \"pages\": \"409--414\"}'),(29,'madou2005software','{\"author\": \"Madou, Matias and Anckaert, Bertrand and Moseley, Patrick and Debray, Saumya and De Sutter, Bjorn and De Bosschere, Koen\", \"booktitle\": \"International Workshop on Information Security Applications\", \"title\": \"Software protection through dynamic code mutation\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2005\", \"organization\": \"Springer\", \"ID\": \"madou2005software\", \"pages\": \"194--206\"}'),(30,'Malone2011','{\"doi\": \"10.1145/2046582.2046596\", \"isbn\": \"9781450310017\", \"keyword\": \"hardware performance counters,integrity\", \"author\": \"Malone, Corey and Zahran, Mohamed and Karri, Ramesh\", \"journal\": \"Proceedings of the sixth ACM workshop on Scalable trusted computing - STC \'11\", \"issn\": \"15437221\", \"ENTRYTYPE\": \"article\", \"mendeley-groups\": \"Tamperproofing/Methods\", \"link\": \"http://www.scopus.com/inward/record.url?eid=2-s2.0-80755143408{\\\\&}partnerID=40{\\\\&}md5=ad5db1f8e5c0131a2a17f457ba1b0497$\\\\backslash$nhttp://dl.acm.org/citation.cfm?doid=2046582.2046596\", \"file\": \":Users/mohsen-tum/Documents/Mendeley Desktop/Malone, Zahran, Karri/Proceedings of the sixth ACM workshop on Scalable trusted computing - STC \'11/Malone, Zahran, Karri - 2011 - Are Hardware Performance Counters a Cost Effective Way for Integrity Checking of Programs.pdf:pdf\", \"year\": \"2011\", \"title\": \"Are Hardware Performance Counters a Cost Effective Way for Integrity Checking of Programs\", \"ID\": \"Malone2011\", \"pages\": \"71\"}'),(31,'Martignoni2010conquer','{\"doi\": \"10.1007/978-3-642-14215-4_2\", \"title\": \"Conqueror: Tamper-proof code execution on legacy systems\", \"journal\": \"Lecture Notes in Computer Science (including subseries Lecture Notes in Artificial Intelligence and Lecture Notes in Bioinformatics)\", \"author\": \"Martignoni, Lorenzo and Paleari, Roberto and Bruschi, Danilo\", \"ENTRYTYPE\": \"article\", \"volume\": \"6201 LNCS\", \"year\": \"2010\", \"ID\": \"Martignoni2010conquer\", \"pages\": \"21--40\"}'),(32,'morgan2015design','{\"author\": \"Morgan, Beno{\\\\^\\\\i}t and Alata, Eric and Nicomette, Vincent and Ka{\\\\^a}niche, Mohamed and Averlant, Guillaume\", \"booktitle\": \"Dependable Computing (PRDC), 2015 IEEE 21st Pacific Rim International Symposium on\", \"title\": \"Design and implementation of a hardware assisted security architecture for software integrity monitoring\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2015\", \"organization\": \"IEEE\", \"ID\": \"morgan2015design\", \"pages\": \"189--198\"}'),(33,'park2015tgvisor','{\"author\": \"Park, Sungjin and Yoon, Jae Nam and Kang, Cheoloh and Kim, Kyong Hoon and Han, Taisook\", \"booktitle\": \"Mobile Cloud Computing, Services, and Engineering (MobileCloud), 2015 3rd IEEE International Conference on\", \"title\": \"TGVisor: A Tiny Hypervisor-Based Trusted Geolocation Framework for Mobile Cloud Clients\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2015\", \"organization\": \"IEEE\", \"ID\": \"park2015tgvisor\", \"pages\": \"99--108\"}'),(34,'Protsenko2015dynamic','{\"doi\": \"10.1109/ARES.2015.98\", \"keyword\": \"Android (operating system);computer crime;cryptography;mobile computing;reverse engineering;Android apps;application piracy;dynamic code loading;dynamic obfuscation techniques;dynamic re-encryption;dynamic self-protection;mobile devices;native code;proprietary mobile software;reverse engineering;tamperproofing;Androids;Encryption;Humanoid robots;Loading;Runtime;Software protection;Android;Software Protection\", \"title\": \"Dynamic Self-Protection and Tamperproofing for Android Apps Using Native Code\", \"booktitle\": \"Availability, Reliability and Security (ARES), 2015 10th International Conference on\", \"author\": \"M. Protsenko and S. Kreuter and T. M\\u00fcller\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2015\", \"ID\": \"Protsenko2015dynamic\", \"pages\": \"129-138\"}'),(35,'Seshadri2005pioneer','{\"doi\": \"10.1145/1095809.1095812\", \"isbn\": \"1-59593-079-5\", \"keyword\": \"dynamic root of trust,rootkit detection,self-check-summing code,software-based code attestation,verifiable code execution\", \"author\": \"Seshadri, Arvind and Luk, Mark and Shi, Elaine and Perrig, Adrian and van Doorn, Leendert and Khosla, Pradeep\", \"journal\": \"ACM SIGOPS Operating Systems Review\", \"issn\": \"01635980\", \"ID\": \"Seshadri2005pioneer\", \"mendeley-groups\": \"Tamperproofing\", \"link\": \"http://dl.acm.org/citation.cfm?id=1095809.1095812\", \"year\": \"2005\", \"title\": \"Pioneer: Verifying Code Integrity and Enforcing Untampered Code Execution on Legacy Systems\", \"ENTRYTYPE\": \"article\"}'),(36,'Spinellis2000','{\"doi\": \"10.1145/353323.353383\", \"isbn\": \"1094-9224\", \"author\": \"Spinellis, Diomidis\", \"ENTRYTYPE\": \"article\", \"abstract\": \"The integrity verification of a device\'s controlling software is an important aspect of many emerging information appliances. We propose the use of reflection, whereby the software is able to examine its own operation, in conjunction with cryptographic hashes as a basis for developing a suitable software verification protocol. For more demanding applications meta-reflective techniques can be used to thwart attacks based on device emulation strategies. We demonstrate how our approach can be used to increase the security of mobile phones, devices for the delivery of digital content, and smartcards.\", \"issn\": \"10949224\", \"number\": \"1\", \"pages\": \"51--62\", \"volume\": \"3\", \"mendeley-groups\": \"Tamperproofing/Methods\", \"file\": \":Users/mohsen-tum/Documents/Mendeley Desktop/Spinellis/ACM Transactions on Information and System Security/Spinellis - 2000 - Reflection as a mechanism for software integrity verification.pdf:pdf\", \"year\": \"2000\", \"title\": \"Reflection as a mechanism for software integrity verification\", \"ID\": \"Spinellis2000\", \"annote\": \"In this approach a software integrity is verified with the help of an external (trusted) entity. Here, the program state is retrieved using reflection, a protocol is proposed to verify the state, and suggested to augment the scheme with CPU perfor.mance counter, before and after the verification call loops.\\nOne obvious attack is to keep an untouched version of the application in the memory next to the tampered with version. Then redirect all hash computations to the good version. The authors, suggest memory expanion and timing as possible countermeasures.\", \"journal\": \"ACM Transactions on Information and System Security\"}'),(37,'teixeira2015siot','{\"author\": \"Teixeira, Fernando A and Machado, Gustavo V and Pereira, Fernando MQ and Wong, Hao Chi and Nogueira, Jos{\\\\\'e} and Oliveira, Leonardo B\", \"booktitle\": \"Proceedings of the 14th International Conference on Information Processing in Sensor Networks\", \"title\": \"SIoT: securing the internet of things through distributed system analysis\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2015\", \"organization\": \"ACM\", \"ID\": \"teixeira2015siot\", \"pages\": \"310--321\"}'),(38,'Wang2005Tamper','{\"isbn\": \"8242866627\", \"keyword\": \"integrity checking,multi-blocking encryption,software piracy,tamper resistant\", \"title\": \"Tamper Resistant Software Through Dynamic Integrity Checking\", \"journal\": \"Proc. Symp. on Cyptography and Information Security (SCIS 05)\", \"author\": \"Wang, Ping and Kang, Seok-kyu and Kim, Kwangjo\", \"ID\": \"Wang2005Tamper\", \"year\": \"2005\", \"ENTRYTYPE\": \"article\"}'),(39,'yao2014cryptvmi','{\"author\": \"Yao, Fangzhou and Sprabery, Read and Campbell, Roy H\", \"booktitle\": \"Proceedings of the 2nd international workshop on Security in cloud computing\", \"title\": \"CryptVMI: a flexible and encrypted virtual machine introspection system in the cloud\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2014\", \"organization\": \"ACM\", \"ID\": \"yao2014cryptvmi\", \"pages\": \"11--18\"}'),(40,'banescu2015software','{\"author\": \"Banescu, Sebastian and Pretschner, Alexander and Battr{\\\\\'e}, Dominic and Cazzulani, St{\\\\\'e}fano and Shield, Robert and Thompson, Greg\", \"booktitle\": \"Proceedings of the 5th ACM Conference on Data and Application Security and Privacy\", \"title\": \"Software-based protection against changeware\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2015\", \"organization\": \"ACM\", \"ID\": \"banescu2015software\", \"pages\": \"231--242\"}'),(41,'Carbone2009','{\"isbn\": \"9781605583525\", \"author\": \"Carbone, Martim and Cui, Weidong and Peinado, Marcus and Lu, Long and Lee, Wenke\", \"journal\": \"Analysis\", \"title\": \"Mapping Kernel Objects to Enable Systematic Integrity Checking\", \"ENTRYTYPE\": \"article\", \"mendeley-groups\": \"Tamperproofing/Methods\", \"file\": \":Users/mohsen-tum/Documents/Mendeley Desktop/Carbone et al/Analysis/Carbone et al. - 2009 - Mapping Kernel Objects to Enable Systematic Integrity Checking.pdf:pdf\", \"year\": \"2009\", \"ID\": \"Carbone2009\", \"pages\": \"555--565\"}'),(42,'Castro2006','{\"isbn\": \"1-931971-47-1\", \"author\": \"Castro, Miguel and Costa, Manuel and Harris, Tim\", \"ENTRYTYPE\": \"article\", \"abstract\": \"Software attacks often subvert the intended data-flow in a vulnerable program. For example, attackers exploit buffer overflows and format string vulnerabilities to write data to unintended locations. We present a simple technique that prevents these attacks by enforcing data-flow integrity. It computes a data-flow graph using static analysis, and it instruments the program to ensure that the flow of data at runtime is allowed by the data-flow graph. We describe an efficient implementation of data-flow integrity enforcement that uses static analysis to reduce instrumentation overhead. This implementation can be used in practice to detect a broad class of attacks and errors because it can be applied automatically to C and C++ programs without modifications, it does not have false positives, and it has low overhead.\", \"title\": \"Securing software by enforcing data-flow integrity\", \"pages\": \"147--160\", \"mendeley-groups\": \"Tamperproofing/Methods\", \"link\": \"http://dl.acm.org/citation.cfm?id=1298455.1298470$\\\\backslash$nhttp://www.usenix.org/event/osdi06/tech/full{\\\\_}papers/castro/castro{\\\\_}html/\", \"file\": \":Users/mohsen-tum/Documents/Mendeley Desktop/Castro, Costa, Harris/Proceedings of the 7th symposium on Operating systems design and implementation/Castro, Costa, Harris - 2006 - Securing software by enforcing data-flow integrity.pdf:pdf\", \"year\": \"2006\", \"ID\": \"Castro2006\", \"journal\": \"Proceedings of the 7th symposium on Operating systems design and implementation\"}'),(43,'gao2015integrity','{\"doi\": \"10.1109/ICAC.2015.34\", \"keyword\": \"Big Data;cloud computing;data integrity;data privacy;Big Data processing;cloud computing technology;dynamic redundancy computation;integrity protection solution;reputation based redundancy computation;Conferences;MapReduce;cloud computing;integrity protection\", \"title\": \"Integrity Protection for Big Data Processing with Dynamic Redundancy Computation\", \"booktitle\": \"Autonomic Computing (ICAC), 2015 IEEE International Conference on\", \"author\": \"Z. Gao and N. Desalvo and P. D. Khoa and S. H. Kim and L. Xu and W. W. Ro and R. M. Verma and W. Shi\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2015\", \"ID\": \"gao2015integrity\", \"pages\": \"159-160\"}'),(44,'karapanos2016verena','{\"author\": \"Karapanos, Nikolaos and Filios, Alexandros and Popa, Raluca Ada and Capkun, Srdjan\", \"booktitle\": \"Proceedings of the 37th IEEE Symposium on Security and Privacy (IEEE S\\\\&P)\", \"title\": \"Verena: End-to-end integrity protection for web applications\", \"ID\": \"karapanos2016verena\", \"year\": \"2016\", \"ENTRYTYPE\": \"inproceedings\"}'),(45,'Kil2009','{\"isbn\": \"9781424444212\", \"keyword\": \"dynamic attestation,integrity,remote attestation,runtime,system security,trusted computing\", \"author\": \"Kil, Chongkyung\", \"journal\": \"IEEE/IFIP International Conference on Dependable Systems {\\\\&} Networks\", \"title\": \"Remote Attestation to Dynamic System Properties: Towards Providing Complete System Integrity Evidence\", \"ENTRYTYPE\": \"article\", \"mendeley-groups\": \"Tamperproofing\", \"file\": \":Users/mohsen-tum/Documents/Mendeley Desktop/Kil/IEEEIFIP International Conference on Dependable Systems {\\\\&} Networks/Kil - 2009 - Remote Attestation to Dynamic System Properties Towards Providing Complete System Integrity Evidence.pdf:pdf\", \"year\": \"2009\", \"ID\": \"Kil2009\", \"pages\": \"115--124\"}'),(46,'neisse2011implementing','{\"author\": \"Neisse, Ricardo and Holling, Dominik and Pretschner, Alexander\", \"booktitle\": \"Proceedings of the 2011 11th IEEE/ACM International Symposium on Cluster, Cloud and Grid Computing\", \"title\": \"Implementing trust in cloud infrastructures\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2011\", \"organization\": \"IEEE Computer Society\", \"ID\": \"neisse2011implementing\", \"pages\": \"524--533\"}'),(47,'sun2015security','{\"author\": \"Sun, Yuqiong and Nanda, Susanta and Jaeger, Trent\", \"booktitle\": \"2015 IEEE 7th International Conference on Cloud Computing Technology and Science (CloudCom)\", \"title\": \"Security-as-a-Service for Microservices-Based Cloud Applications\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2015\", \"organization\": \"IEEE\", \"ID\": \"sun2015security\", \"pages\": \"50--57\"}'),(48,'pappas2012smashing','{\"author\": \"Pappas, Vasilis and Polychronakis, Michalis and Keromytis, Angelos D\", \"booktitle\": \"2012 IEEE Symposium on Security and Privacy\", \"title\": \"Smashing the gadgets: Hindering return-oriented programming using in-place code randomization\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2012\", \"organization\": \"IEEE\", \"ID\": \"pappas2012smashing\", \"pages\": \"601--615\"}'),(49,'pappas2013transparent','{\"author\": \"Pappas, Vasilis and Polychronakis, Michalis and Keromytis, Angelos D\", \"booktitle\": \"Presented as part of the 22nd USENIX Security Symposium (USENIX Security 13)\", \"title\": \"Transparent ROP exploit mitigation using indirect branch tracing\", \"pages\": \"447--462\", \"year\": \"2013\", \"ID\": \"pappas2013transparent\", \"ENTRYTYPE\": \"inproceedings\"}');
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
 1 AS `text_attribute`*/;
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
 1 AS `atts`*/;
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
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `relation`
--

LOCK TABLES `relation` WRITE;
/*!40000 ALTER TABLE `relation` DISABLE KEYS */;
INSERT INTO `relation` VALUES (1,'Depends','simple dependency');
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
  `id_attribute` int(11) NOT NULL,
  `id_dimension` int(11) NOT NULL,
  PRIMARY KEY (`id_taxonomy_dimension`),
  UNIQUE KEY `id_taxonomy_dimension_UNIQUE` (`id_taxonomy_dimension`)
) ENGINE=InnoDB AUTO_INCREMENT=56 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `taxonomy_dimension`
--

LOCK TABLES `taxonomy_dimension` WRITE;
/*!40000 ALTER TABLE `taxonomy_dimension` DISABLE KEYS */;
INSERT INTO `taxonomy_dimension` VALUES (1,1,2,1),(2,1,3,1),(3,1,4,1),(4,1,6,1),(5,1,7,1),(6,1,8,1),(7,1,9,1),(8,1,10,1),(9,1,11,1),(10,1,12,1),(11,1,13,1),(12,1,15,1),(13,1,16,1),(14,1,17,1),(15,1,18,1),(16,1,19,1),(17,1,21,1),(18,1,22,1),(19,1,23,1),(20,1,24,1),(21,1,25,1),(22,1,28,2),(23,1,29,2),(24,1,30,2),(25,1,31,2),(26,1,33,3),(27,1,34,3),(28,1,36,3),(29,1,37,3),(30,1,39,3),(31,1,40,3),(32,1,42,3),(33,1,43,3),(34,1,45,3),(35,1,46,3),(36,1,47,3),(37,1,48,3),(38,1,49,3),(39,1,51,3),(40,1,52,3),(41,1,53,3),(42,1,54,3),(43,1,55,3),(44,1,56,3),(45,1,58,3),(46,1,59,3),(47,1,60,3),(48,1,61,3),(49,1,63,3),(50,1,64,3),(51,1,65,3),(52,1,66,3),(53,1,68,3),(54,1,69,3),(55,1,70,3);
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
  `id_src_attribute` int(11) NOT NULL,
  `id_dest_attribute` int(11) NOT NULL,
  `id_relation` int(11) NOT NULL,
  PRIMARY KEY (`id_taxonomy_relation`),
  UNIQUE KEY `id_taxonomy_relation_UNIQUE` (`id_taxonomy_relation`)
) ENGINE=InnoDB AUTO_INCREMENT=56 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `taxonomy_relation`
--

LOCK TABLES `taxonomy_relation` WRITE;
/*!40000 ALTER TABLE `taxonomy_relation` DISABLE KEYS */;
INSERT INTO `taxonomy_relation` VALUES (1,1,1,2,1),(2,1,1,3,1),(3,1,1,4,1),(4,1,5,6,1),(5,1,5,7,1),(6,1,5,8,1),(7,1,7,9,1),(8,1,7,10,1),(9,1,8,11,1),(10,1,8,12,1),(11,1,8,13,1),(12,1,14,15,1),(13,1,14,16,1),(14,1,14,17,1),(15,1,14,18,1),(16,1,14,19,1),(17,1,20,21,1),(18,1,20,22,1),(19,1,20,23,1),(20,1,20,24,1),(21,1,20,25,1),(22,1,27,28,1),(23,1,27,29,1),(24,1,27,30,1),(25,1,27,31,1),(26,1,32,33,1),(27,1,32,34,1),(28,1,35,36,1),(29,1,35,37,1),(30,1,38,39,1),(31,1,38,40,1),(32,1,41,42,1),(33,1,41,43,1),(34,1,44,45,1),(35,1,44,46,1),(36,1,44,47,1),(37,1,44,48,1),(38,1,44,49,1),(39,1,50,51,1),(40,1,50,52,1),(41,1,50,53,1),(42,1,50,54,1),(43,1,50,55,1),(44,1,50,56,1),(45,1,57,58,1),(46,1,57,59,1),(47,1,57,60,1),(48,1,57,61,1),(49,1,62,63,1),(50,1,62,64,1),(51,1,62,65,1),(52,1,62,66,1),(53,1,67,68,1),(54,1,67,69,1),(55,1,67,70,1);
/*!40000 ALTER TABLE `taxonomy_relation` ENABLE KEYS */;
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
/*!50001 VIEW `paper_attribute` AS select `taxonomy`.`id_taxonomy` AS `id_taxonomy`,`paper`.`id_paper` AS `id_paper`,`paper`.`citation` AS `citation`,`paper`.`bib` AS `bib`,`mapping`.`id_attribute` AS `id_attribute`,`GETATTRIBUTENAME`(`mapping`.`id_attribute`) AS `text_attribute` from ((`paper` join `mapping` on((`paper`.`id_paper` = `mapping`.`id_paper`))) join `taxonomy`) union select `rel1`.`id_taxonomy` AS `id_taxonomy`,`paper`.`id_paper` AS `id_paper`,`paper`.`citation` AS `citation`,`paper`.`bib` AS `bib`,`rel1`.`id_src_attribute` AS `id_attribute`,`GETATTRIBUTENAME`(`rel1`.`id_src_attribute`) AS `text_attribute` from ((`paper` join `mapping` on((`paper`.`id_paper` = `mapping`.`id_paper`))) left join `taxonomy_relation` `rel1` on((`rel1`.`id_dest_attribute` = `mapping`.`id_attribute`))) where (`rel1`.`id_src_attribute` is not null) union select `rel2`.`id_taxonomy` AS `id_taxonomy`,`paper`.`id_paper` AS `id_paper`,`paper`.`citation` AS `citation`,`paper`.`bib` AS `bib`,`rel2`.`id_src_attribute` AS `id_attribute`,`GETATTRIBUTENAME`(`rel2`.`id_src_attribute`) AS `text_attribute` from (((`paper` join `mapping` on((`paper`.`id_paper` = `mapping`.`id_paper`))) left join `taxonomy_relation` `rel1` on((`rel1`.`id_dest_attribute` = `mapping`.`id_attribute`))) left join `taxonomy_relation` `rel2` on((`rel2`.`id_dest_attribute` = `rel1`.`id_src_attribute`))) where (`rel2`.`id_src_attribute` is not null) union select `rel3`.`id_taxonomy` AS `id_taxonomy`,`paper`.`id_paper` AS `id_paper`,`paper`.`citation` AS `citation`,`paper`.`bib` AS `bib`,`rel3`.`id_src_attribute` AS `id_attribute`,`GETATTRIBUTENAME`(`rel3`.`id_src_attribute`) AS `text_attribute` from ((((`paper` join `mapping` on((`paper`.`id_paper` = `mapping`.`id_paper`))) left join `taxonomy_relation` `rel1` on((`rel1`.`id_dest_attribute` = `mapping`.`id_attribute`))) left join `taxonomy_relation` `rel2` on((`rel2`.`id_dest_attribute` = `rel1`.`id_src_attribute`))) left join `taxonomy_relation` `rel3` on((`rel3`.`id_dest_attribute` = `rel2`.`id_src_attribute`))) where (`rel3`.`id_src_attribute` is not null) */;
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
/*!50001 VIEW `paper_merged_attributes` AS select distinct `paper_attribute`.`id_taxonomy` AS `id_taxonomy`,`paper_attribute`.`id_paper` AS `id_paper`,`paper_attribute`.`citation` AS `citation`,`paper_attribute`.`bib` AS `bib`,`att_table`.`atts` AS `atts` from (`classification`.`paper_attribute` join (select `a`.`id_paper` AS `id_paper`,group_concat(concat(`a`.`text_attribute`) separator ',') AS `atts` from (select `paper_attribute`.`id_paper` AS `id_paper`,`paper_attribute`.`id_attribute` AS `id_attribute`,`paper_attribute`.`text_attribute` AS `text_attribute` from `classification`.`paper_attribute` order by `paper_attribute`.`id_attribute`) `a` group by `a`.`id_paper`) `att_table` on((`att_table`.`id_paper` = `paper_attribute`.`id_paper`))) */;
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

-- Dump completed on 2017-08-01 17:01:14
