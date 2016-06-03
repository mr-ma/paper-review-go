CREATE DATABASE  IF NOT EXISTS `paper_review` /*!40100 DEFAULT CHARACTER SET utf8 */;
USE `paper_review`;
-- MySQL dump 10.13  Distrib 5.7.9, for Win64 (x86_64)
--
-- Host: localhost    Database: paper_review
-- ------------------------------------------------------
-- Server version	5.7.12-log

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
  `Keywords` varchar(2000) DEFAULT NULL,
  `Abstract` varchar(10000) DEFAULT NULL,
  `Journal` varchar(500) DEFAULT NULL,
  `ResearchId` int(11) NOT NULL,
  `Authors` varchar(1000) NOT NULL,
  `File` varchar(200) NOT NULL,
  `Source` varchar(50) NOT NULL,
  `Enabled` bit(1) NOT NULL DEFAULT b'0',
  PRIMARY KEY (`ArticleId`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Temporary view structure for view `articles_view`
--

DROP TABLE IF EXISTS `articles_view`;
/*!50001 DROP VIEW IF EXISTS `articles_view`*/;
SET @saved_cs_client     = @@character_set_client;
SET character_set_client = utf8;
/*!50001 CREATE VIEW `articles_view` AS SELECT
 1 AS `ArticleId`,
 1 AS `Title`,
 1 AS `year`,
 1 AS `cited_by`,
 1 AS `Keywords`,
 1 AS `Abstract`,
 1 AS `Journal`,
 1 AS `ResearchId`,
 1 AS `Authors`,
 1 AS `File`,
 1 AS `Source`*/;
SET character_set_client = @saved_cs_client;

--
-- Table structure for table `mitarbeiters`
--

DROP TABLE IF EXISTS `mitarbeiters`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `mitarbeiters` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Pass_Hash` binary(32) NOT NULL,
  `Nme` varchar(100) NOT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `research`
--

DROP TABLE IF EXISTS `research`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `research` (
  `ResearchId` int(11) NOT NULL AUTO_INCREMENT,
  `Questions` varchar(1000) NOT NULL,
  `Review_Template` varchar(2000) DEFAULT NULL,
  `Title` varchar(100) NOT NULL,
  PRIMARY KEY (`ResearchId`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `tags`
--

DROP TABLE IF EXISTS `tags`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tags` (
  `TagId` int(11) NOT NULL AUTO_INCREMENT,
  `Text` varchar(500) NOT NULL,
  `ResearchID` int(11) NOT NULL,
  PRIMARY KEY (`TagId`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Temporary view structure for view `unique_articles_view`
--

DROP TABLE IF EXISTS `unique_articles_view`;
/*!50001 DROP VIEW IF EXISTS `unique_articles_view`*/;
SET @saved_cs_client     = @@character_set_client;
SET character_set_client = utf8;
/*!50001 CREATE VIEW `unique_articles_view` AS SELECT
 1 AS `ArticleId`,
 1 AS `Title`,
 1 AS `year`,
 1 AS `cited_by`,
 1 AS `Keywords`,
 1 AS `Abstract`,
 1 AS `Journal`,
 1 AS `ResearchId`,
 1 AS `Authors`,
 1 AS `File`,
 1 AS `Source`*/;
SET character_set_client = @saved_cs_client;

--
-- Table structure for table `vote_tags`
--

DROP TABLE IF EXISTS `vote_tags`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `vote_tags` (
  `Vote_Tags_Id` int(11) NOT NULL AUTO_INCREMENT,
  `Tag_Id` int(11) NOT NULL,
  `VoteId` int(11) NOT NULL,
  PRIMARY KEY (`Vote_Tags_Id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

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
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Final view structure for view `articles_view`
--

/*!50001 DROP VIEW IF EXISTS `articles_view`*/;
/*!50001 SET @saved_cs_client          = @@character_set_client */;
/*!50001 SET @saved_cs_results         = @@character_set_results */;
/*!50001 SET @saved_col_connection     = @@collation_connection */;
/*!50001 SET character_set_client      = utf8 */;
/*!50001 SET character_set_results     = utf8 */;
/*!50001 SET collation_connection      = utf8_general_ci */;
/*!50001 CREATE ALGORITHM=UNDEFINED */
/*!50013 DEFINER=`root`@`localhost` SQL SECURITY DEFINER */
/*!50001 VIEW `articles_view` AS select `distict_articles`.`ArticleId` AS `ArticleId`,`distict_articles`.`Title` AS `Title`,`distict_articles`.`year` AS `year`,`distict_articles`.`cited_by` AS `cited_by`,`distict_articles`.`Keywords` AS `Keywords`,`distict_articles`.`Abstract` AS `Abstract`,`distict_articles`.`Journal` AS `Journal`,`distict_articles`.`ResearchId` AS `ResearchId`,`distict_articles`.`Authors` AS `Authors`,`distict_articles`.`File` AS `File`,`distict_articles`.`Source` AS `Source` from (select max(`a`.`ArticleId`) AS `ArticleId`,`a`.`Title` AS `Title`,max(cast(`a`.`year` as unsigned)) AS `year`,max(cast(`a`.`cited_by` as unsigned)) AS `cited_by`,max(`a`.`Keywords`) AS `Keywords`,max(`a`.`Abstract`) AS `Abstract`,max(`a`.`Journal`) AS `Journal`,`a`.`ResearchId` AS `ResearchId`,max(`a`.`Authors`) AS `Authors`,min(`a`.`File`) AS `File`,max(`a`.`Source`) AS `Source`,max(`a`.`Enabled`) AS `Enabled` from `paper_review`.`articles` `a` group by `a`.`Title`,`a`.`ResearchId`) `distict_articles` where (`distict_articles`.`Enabled` = 1) */;
/*!50001 SET character_set_client      = @saved_cs_client */;
/*!50001 SET character_set_results     = @saved_cs_results */;
/*!50001 SET collation_connection      = @saved_col_connection */;

--
-- Final view structure for view `unique_articles_view`
--

/*!50001 DROP VIEW IF EXISTS `unique_articles_view`*/;
/*!50001 SET @saved_cs_client          = @@character_set_client */;
/*!50001 SET @saved_cs_results         = @@character_set_results */;
/*!50001 SET @saved_col_connection     = @@collation_connection */;
/*!50001 SET character_set_client      = utf8 */;
/*!50001 SET character_set_results     = utf8 */;
/*!50001 SET collation_connection      = utf8_general_ci */;
/*!50001 CREATE ALGORITHM=UNDEFINED */
/*!50013 DEFINER=`root`@`localhost` SQL SECURITY DEFINER */
/*!50001 VIEW `unique_articles_view` AS select `distict_articles`.`ArticleId` AS `ArticleId`,`distict_articles`.`Title` AS `Title`,`distict_articles`.`year` AS `year`,`distict_articles`.`cited_by` AS `cited_by`,`distict_articles`.`Keywords` AS `Keywords`,`distict_articles`.`Abstract` AS `Abstract`,`distict_articles`.`Journal` AS `Journal`,`distict_articles`.`ResearchId` AS `ResearchId`,`distict_articles`.`Authors` AS `Authors`,`distict_articles`.`File` AS `File`,`distict_articles`.`Source` AS `Source` from (select max(`a`.`ArticleId`) AS `ArticleId`,`a`.`Title` AS `Title`,max(cast(`a`.`year` as unsigned)) AS `year`,max(cast(`a`.`cited_by` as unsigned)) AS `cited_by`,max(`a`.`Keywords`) AS `Keywords`,max(`a`.`Abstract`) AS `Abstract`,max(`a`.`Journal`) AS `Journal`,`a`.`ResearchId` AS `ResearchId`,max(`a`.`Authors`) AS `Authors`,min(`a`.`File`) AS `File`,max(`a`.`Source`) AS `Source`,max(`a`.`Enabled`) AS `Enabled` from `paper_review`.`articles` `a` group by `a`.`Title`,`a`.`ResearchId`) `distict_articles` */;
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

-- Dump completed on 2016-05-27 19:02:11
