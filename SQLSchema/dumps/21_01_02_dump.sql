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
  `id_taxonomy` int(11) unsigned NOT NULL,
  `text` varchar(50) NOT NULL,
  `children` longtext,
  PRIMARY KEY (`id_attribute`,`id_taxonomy`),
  UNIQUE KEY `allChildrenPerAttribute_id_attribute_UNIQUE` (`id_attribute`,`id_taxonomy`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `allchildrenperattribute`
--

LOCK TABLES `allchildrenperattribute` WRITE;
/*!40000 ALTER TABLE `allchildrenperattribute` DISABLE KEYS */;
INSERT INTO `allchildrenperattribute` VALUES (1,1,'Integrity Protection Assets','1'),(2,1,'Behavior','2'),(3,1,'Data','3'),(4,1,'Data and behavior','4'),(5,1,'Representation','5,6,7,8,9,10,11,13,12'),(6,1,'Static','6'),(7,1,'In memory','7,9,10'),(8,1,'In execution','8,11,13,12'),(9,1,'Code invariants','9'),(10,1,'Data invariants','10'),(11,1,'Trace','11,12'),(12,1,'Timed trace','12'),(13,1,'HW counters','13'),(14,1,'Granularity','14,15,17,18,19,95'),(15,1,'Instructions','15'),(16,1,'BB','16'),(17,1,'Function','17'),(18,1,'Slice','18'),(19,1,'Application','19'),(20,1,'Lifecycle activity','20,21,22,23,24,25,94'),(21,1,'Pre-compile','21'),(22,1,'Compile','22'),(23,1,'Post-compile','23'),(24,1,'Load','24'),(25,1,'Run','25'),(26,1,'Not root','26'),(27,1,'Attack','27,28,29,30,31,78'),(28,1,'Binary ','28'),(29,1,'Process memory','29'),(30,1,'Runtime data','30'),(31,1,'Control flow','31'),(32,1,'Measure','32,33,34,35,38,41,44,50,57,62,67,36,37,39,42,43,45,46,47,48,49,51,52,53,54,55,58,59,60,61,63,64,65,69,70,71,89,90,91,92'),(33,1,'Local','33'),(34,1,'Remote','34'),(35,1,'Monitor','35,36,37'),(36,1,'State inspection','36'),(37,1,'Introspection','37'),(38,1,'Response','38,39,92'),(39,1,'Proactive','39'),(40,1,'Postmortem','40'),(41,1,'Transformation','41,42,43'),(42,1,'Manual','42'),(43,1,'Automatic','43'),(44,1,'Check','44,45,46,47,48,49'),(45,1,'Checksum','45'),(46,1,'Signature','46'),(47,1,'Equation eval','47'),(48,1,'Majority vote','48'),(49,1,'Access control','49'),(50,1,'Hardening','50,51,52,53,54,55,91'),(51,1,'Cyclic checks','51'),(52,1,'Mutation','52'),(53,1,'Code concealment','53'),(54,1,'Cloning','54'),(55,1,'Layered interpretation','55'),(56,1,'Block chain','56'),(57,1,'Overhead','57,58,59,60,61'),(58,1,'Fair','58'),(59,1,'Medium','59'),(60,1,'High','60'),(61,1,'N/A','61'),(62,1,'Trust anchor','62,63,64,65,71,89'),(63,1,'TPM','63'),(64,1,'SGX','64'),(65,1,'Other','65'),(66,1,'None','66'),(67,1,'Protection level','67,69,70,90'),(68,1,'Internal','68'),(69,1,'External','69'),(70,1,'Hypervisor','70'),(71,1,'Software','71'),(72,1,'Reverse engineering','72,27,73,80,84,26,28,29,30,31,78,79,81,82,83,85,86,87,88'),(73,1,'Attacker','73,26'),(78,1,'Call interposition','78'),(79,1,'Disassembler','79'),(80,1,'Tools','80,79,81,82,83'),(81,1,'Debugger','81'),(82,1,'Tracer','82'),(83,1,'Emulator','83'),(84,1,'Discovery','84,85,86,87,88'),(85,1,'Pattern matching','85'),(86,1,'Taint analysis','86'),(87,1,'Graph-based analysis','87'),(88,1,'Symbolic execution','88'),(89,1,'Dongle','89'),(90,1,'Self-check','90'),(91,1,'Hash chain','91'),(92,1,'Reactive','92'),(93,1,'Asset','93,2,3,4'),(94,1,'Link','94'),(95,1,'Basic block','95');
/*!40000 ALTER TABLE `allchildrenperattribute` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `allparentsperattribute`
--

DROP TABLE IF EXISTS `allparentsperattribute`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `allparentsperattribute` (
  `id_attribute` int(11) unsigned NOT NULL,
  `id_taxonomy` int(11) unsigned NOT NULL,
  `text` varchar(50) NOT NULL,
  `parents` longtext,
  PRIMARY KEY (`id_attribute`,`id_taxonomy`),
  UNIQUE KEY `allParentsPerAttribute_id_attribute_UNIQUE` (`id_attribute`,`id_taxonomy`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `allparentsperattribute`
--

LOCK TABLES `allparentsperattribute` WRITE;
/*!40000 ALTER TABLE `allparentsperattribute` DISABLE KEYS */;
INSERT INTO `allparentsperattribute` VALUES (1,1,'Integrity Protection Assets',''),(2,1,'Behavior','Asset'),(3,1,'Data','Asset'),(4,1,'Data and behavior','Asset'),(5,1,'Representation',''),(6,1,'Static','Representation'),(7,1,'In memory','Representation'),(8,1,'In execution','Representation'),(9,1,'Code invariants','In memory,Representation'),(10,1,'Data invariants','In memory,Representation'),(11,1,'Trace','In execution,Representation'),(12,1,'Timed trace','Trace,In execution,Representation'),(13,1,'HW counters','In execution,Representation'),(14,1,'Granularity',''),(15,1,'Instructions','Granularity'),(16,1,'BB',''),(17,1,'Function','Granularity'),(18,1,'Slice','Granularity'),(19,1,'Application','Granularity'),(20,1,'Lifecycle activity',''),(21,1,'Pre-compile','Lifecycle activity'),(22,1,'Compile','Lifecycle activity'),(23,1,'Post-compile','Lifecycle activity'),(24,1,'Load','Lifecycle activity'),(25,1,'Run','Lifecycle activity'),(26,1,'Not root','Attacker,Reverse engineering'),(27,1,'Attack','Reverse engineering'),(28,1,'Binary ','Attack,Reverse engineering'),(29,1,'Process memory','Attack,Reverse engineering'),(30,1,'Runtime data','Attack,Reverse engineering'),(31,1,'Control flow','Attack,Reverse engineering'),(32,1,'Measure',''),(33,1,'Local','Measure'),(34,1,'Remote','Measure'),(35,1,'Monitor','Measure'),(36,1,'State inspection','Monitor,Measure'),(37,1,'Introspection','Monitor,Measure'),(38,1,'Response','Measure'),(39,1,'Proactive','Response,Measure'),(40,1,'Postmortem',''),(41,1,'Transformation','Measure'),(42,1,'Manual','Transformation,Measure'),(43,1,'Automatic','Transformation,Measure'),(44,1,'Check','Measure'),(45,1,'Checksum','Check,Measure'),(46,1,'Signature','Check,Measure'),(47,1,'Equation eval','Check,Measure'),(48,1,'Majority vote','Check,Measure'),(49,1,'Access control','Check,Measure'),(50,1,'Hardening','Measure'),(51,1,'Cyclic checks','Hardening,Measure'),(52,1,'Mutation','Hardening,Measure'),(53,1,'Code concealment','Hardening,Measure'),(54,1,'Cloning','Hardening,Measure'),(55,1,'Layered interpretation','Hardening,Measure'),(56,1,'Block chain',''),(57,1,'Overhead','Measure'),(58,1,'Fair','Overhead,Measure'),(59,1,'Medium','Overhead,Measure'),(60,1,'High','Overhead,Measure'),(61,1,'N/A','Overhead,Measure'),(62,1,'Trust anchor','Measure'),(63,1,'TPM','Trust anchor,Measure'),(64,1,'SGX','Trust anchor,Measure'),(65,1,'Other','Trust anchor,Measure'),(66,1,'None',''),(67,1,'Protection level','Measure'),(68,1,'Internal',''),(69,1,'External','Protection level,Measure'),(70,1,'Hypervisor','Protection level,Measure'),(71,1,'Software','Trust anchor,Measure'),(72,1,'Reverse engineering',''),(73,1,'Attacker','Reverse engineering'),(78,1,'Call interposition','Attack,Reverse engineering'),(79,1,'Disassembler','Tools,Reverse engineering'),(80,1,'Tools','Reverse engineering'),(81,1,'Debugger','Tools,Reverse engineering'),(82,1,'Tracer','Tools,Reverse engineering'),(83,1,'Emulator','Tools,Reverse engineering'),(84,1,'Discovery','Reverse engineering'),(85,1,'Pattern matching','Discovery,Reverse engineering'),(86,1,'Taint analysis','Discovery,Reverse engineering'),(87,1,'Graph-based analysis','Discovery,Reverse engineering'),(88,1,'Symbolic execution','Discovery,Reverse engineering'),(89,1,'Dongle','Trust anchor,Measure'),(90,1,'Self-check','Protection level,Measure'),(91,1,'Hash chain','Hardening,Measure'),(92,1,'Reactive','Response,Measure'),(93,1,'Asset',''),(94,1,'Link','Lifecycle activity'),(95,1,'Basic block','Granularity');
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
  `id_taxonomy` int(11) unsigned NOT NULL,
  PRIMARY KEY (`id_attribute`),
  UNIQUE KEY `id_attribute_UNIQUE` (`id_attribute`),
  UNIQUE KEY `attribute_text_UNIQUE` (`text`,`id_taxonomy`),
  KEY `attribute_id_taxonomy_foreign` (`id_taxonomy`),
  CONSTRAINT `attribute_id_taxonomy_foreign` FOREIGN KEY (`id_taxonomy`) REFERENCES `taxonomy` (`id_taxonomy`) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=108 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `attribute`
--

LOCK TABLES `attribute` WRITE;
/*!40000 ALTER TABLE `attribute` DISABLE KEYS */;
INSERT INTO `attribute` VALUES (1,'Integrity Protection Assets','1539.6977208359122','828.406523846749',NULL,NULL,'2982.356969925021','-2285.019961769441','-864.2180305371212',NULL,NULL,NULL,0,1),(2,'Behavior','362.03679880998527','70.61697033346249',NULL,NULL,'-4105.754594713023','2970.805490853027','520.6240466093806',NULL,NULL,NULL,0,1),(3,'Data','360.50599492453625','156.00928187629017',NULL,NULL,'-5002.920658104225','2416.7437462541925','-667.9458699193781',NULL,NULL,NULL,0,1),(4,'Data and behavior','363.56267226780255','233.77468381901463',NULL,NULL,'-3116.814622838809','2150.293659766716','-739.4750876543186',NULL,NULL,NULL,0,1),(5,'Representation','328.44175759223543','432.6782494091842','-706.6576875701488','-609.3344062416174','-5631.8223863130315','-700.1480012803709','232.8933266727454','-2297.307274575899','1861.2451999882087','251.40799472895628',1,1),(6,'Static','1020.112697296182','169.14810291408673',NULL,NULL,'-753.8117777226684','2672.900536068831','55.5562648093229',NULL,NULL,NULL,0,1),(7,'In memory','1035.9341191918215','319.5410718472651',NULL,NULL,'-859.8475232068192','1419.0143417630459','1049.9932319804325',NULL,NULL,NULL,0,1),(8,'In execution','1036.422207754128','515.8949411042306',NULL,NULL,'-525.1480844836708','-997.8604194722125','-515.3976828985001',NULL,NULL,NULL,0,1),(9,'Code invariants','1473.5597140112222','232.10834584752325',NULL,NULL,'206.51897236227046','1964.5856119335288','1212.8380148609385',NULL,NULL,NULL,0,1),(10,'Data invariants','1475.0087028973157','391.5997213615837',NULL,NULL,'1584.2972884243895','1107.9452114293008','596.1960385203179',NULL,NULL,NULL,0,1),(11,'Trace','1476.9101394936774','659.8362915899118',NULL,NULL,'1414.656335785149','-1776.66513018276','-191.61968505004643',NULL,NULL,NULL,0,1),(12,'Timed trace','1652.9485817795407','661.3853305327655',NULL,NULL,'2706.2310058466305','-505.4745310046228','-577.2241671516888',NULL,NULL,NULL,0,1),(13,'HW counters','1477.121572065918','520.1881886758251',NULL,NULL,'1194.9742695396158','65.74160947273646','-137.41344005045494',NULL,NULL,NULL,0,1),(14,'Granularity','327.83669869646576','674.855512732843','-686.6596563047685','-281.3360728952638','-4849.209350403896','-1872.1017791288107','20.015773720053858','2750.4431883226216','-2948.6211902334258','-502.6910289016656',1,1),(15,'Instructions','593.8761270678542','600.541071847265',NULL,NULL,'-3677.00212371179','-2059.3275936855316','-1231.9318958901586',NULL,NULL,NULL,0,1),(16,'BB','1696.8048303141118','828.2885674277604',NULL,NULL,'2143.4276500237497','-2464.144864036762','-1036.6394584036368',NULL,NULL,NULL,0,1),(17,'Function','576.3228150877201','672.3462308565527',NULL,NULL,'-3825.8952008628225','-2536.6896428436626','-495.4057692218644',NULL,NULL,NULL,0,1),(18,'Slice','560.9285313738394','752.36643808501',NULL,NULL,'-3435.1692734008243','-2081.4646016536462','1249.2503410645315',NULL,NULL,NULL,0,1),(19,'Application','590.1813644885203','825.9356842563209',NULL,NULL,'-3659.547946100021','-2773.817578816789','732.0912311726131',NULL,NULL,NULL,0,1),(20,'Lifecycle activity','499.43222543214654','296.8923115428277','-690.6616875388991','-767.670541636393','-4688.243193111236','110.89828167666965','-422.7101612981819','3088.7134093484487','338.8041040030753','-411.428475846131',1,1),(21,'Pre-compile','727.983737113648','68.90956051470602',NULL,NULL,'-2098.4675273781886','1903.2747623767602','708.8565788744929',NULL,NULL,NULL,0,1),(22,'Compile','708.0208646188082','143.38072880030973',NULL,NULL,'-2060.668061605252','1333.5185979024836','700.4689784693802',NULL,NULL,NULL,0,1),(23,'Post-compile','733.842286514448','211.20478025735315',NULL,NULL,'-2417.367759420388','1072.6497859901137','-40.15143208390509',NULL,NULL,NULL,0,1),(24,'Load','693.7813361339005','366.75480528573263',NULL,NULL,'-2148.2653425899725','-308.27053657036','59.54550985321248',NULL,NULL,NULL,0,1),(25,'Run','695.4851303057279','438.1278956856295',NULL,NULL,'-2078.8620483112054','-1015.7537643835454','82.76219647055268',NULL,NULL,NULL,0,1),(26,'Not root','474.7266028589015','454.7732747885879','','',NULL,NULL,NULL,NULL,NULL,NULL,0,1),(27,'Attack','959.3360043758585','256.6619190499381','-323.9978125170894','-482.6650312627768','-560.3840244102373','102.1552480434342','-956.7462882957757','-1043.6553421224908','633.6478862886546','4.547473508864641e-13',1,1),(28,'Binary ','813.3219621018261','142.03790265880497',NULL,NULL,'-669.670563198074','778.7013671716411','-1338.7442185818304',NULL,NULL,NULL,0,1),(29,'Process memory','959.7840594430213','141.03790265880497',NULL,NULL,'-94.03937291157544','-1336.7099182805853','1261.323300421252',NULL,NULL,NULL,0,1),(30,'Runtime data','1103.0008419920632','139.4152452627283',NULL,NULL,'-1286.5345978331447','-12.458919208386305','-1284.887427267443',NULL,NULL,NULL,0,1),(31,'Control flow','1307.8590227685681','140.11370797641482',NULL,NULL,'613.0758806397226','625.6623474191132','-1220.5401846641005',NULL,NULL,NULL,0,1),(32,'Measure','847.3942993388207','425.0985353391625','-987.9939687971188','-766.6627500305989','-1540.1203017455578','164.35563430566253','519.8406371279184','-2885.666024652709','-362.649687924208','366.3875195434952',1,1),(33,'Local','1083.8113585010585','349.5760338996515',NULL,NULL,'414.01249022013945','650.3924349711187','217.48400016653432',NULL,NULL,NULL,0,1),(34,'Remote','1084.2995836339228','424.4551551031491',NULL,NULL,'491.74632812015005','97.3780742973924','-123.8359011677826',NULL,NULL,NULL,0,1),(35,'Monitor','422.4089459471961','352.089115445454',NULL,NULL,'-3569.5793861554353','1840.7840486857215','360.30867064656013',NULL,NULL,NULL,0,1),(36,'State inspection','139.41025262586263','152.3557294513971',NULL,NULL,'-5341.030340064095','2669.721222390146','-724.4818029428425',NULL,NULL,NULL,0,1),(37,'Introspection','136.15191089189352','352.71535517567395',NULL,NULL,'-4970.62799052557','1395.0657686414263','-235.261091445313',NULL,NULL,NULL,0,1),(38,'Response','426.9031877055812','232.79764502657298',NULL,NULL,'-2593.6039171763614','2544.8218260668036','-759.650458124624',NULL,NULL,NULL,0,1),(39,'Proactive','304.9316064436389','131.33039620209587',NULL,NULL,'-3372.016016595666','3430.1408603523755','-464.66544817688964',NULL,NULL,NULL,0,1),(40,'Postmortem','1796.8697013097888','637.6056002594439',NULL,NULL,'-5498.445586366219','-3168.741296759451','-466.7710346363615',NULL,NULL,NULL,0,1),(41,'Transformation','604.4856117755306','227.08131819475398',NULL,NULL,'-674.7676646465982','2349.977089886205','-139.82051733421395',NULL,NULL,NULL,0,1),(42,'Manual','553.6373636104923','121.95101950454278',NULL,NULL,'-921.5490625752307','3398.8951229358577','-479.87542564422074',NULL,NULL,NULL,0,1),(43,'Automatic','671.6729465958057','121.38822513286442',NULL,NULL,'127.94846090778958','3219.9382164220856','20.700303685751805',NULL,NULL,NULL,0,1),(44,'Check','847.6077972507001','241.9154365192295',NULL,NULL,'1837.275949199416','2425.3707306235287','-191.4718377402669',NULL,NULL,NULL,0,1),(45,'Checksum','796.7595490856617','119.67969555174551',NULL,NULL,'1281.3745416615216','2879.8359194492364','723.2800076185517',NULL,NULL,NULL,0,1),(46,'Signature','913.2121912332129','118.9711659706266',NULL,NULL,'2246.4861157601076','2810.5915954973075','942.1178163301893',NULL,NULL,NULL,0,1),(47,'Equation eval','1063.916010867478','118.97116597062654',NULL,NULL,'4753.967841269285','3722.180629086103','-665.6491212581645',NULL,NULL,NULL,0,1),(48,'Majority vote','1243.5593911034905','118.97116597062654',NULL,NULL,'5182.609118082028','2736.231141284869','551.5771540960122',NULL,NULL,NULL,0,1),(49,'Access control','1415.9515938527904','119.49366741013745',NULL,NULL,'6003.849502008718','2344.8187915623075','899.7533214767686',NULL,NULL,NULL,0,1),(50,'Hardening','1079.1773981587537','260.30455196468347',NULL,NULL,'3255.87157917825','840.7435575967074','88.59670625233503',NULL,NULL,NULL,0,1),(51,'Cyclic checks','1579.377556140925','264.23018137981813',NULL,NULL,'5319.912403964804','1727.0256279508856','111.736167113791',NULL,NULL,NULL,0,1),(52,'Mutation','1584.594457320991','342.1185238230755',NULL,NULL,'4730.551295434692','949.487537634357','746.8857972408391',NULL,NULL,NULL,0,1),(53,'Code concealment','1594.0826824538556','432.01779149265667',NULL,NULL,'5356.577880976774','409.08695466782','279.84508938293175',NULL,NULL,NULL,0,1),(54,'Cloning','1597.591054052804','522.4798535339163',NULL,NULL,'5418.041195227526','-113.16968086050889','220.2466009992163',NULL,NULL,NULL,0,1),(55,'Layered interpretation','1621.5637433394631','599.0876507846166',NULL,NULL,'5193.809046312347','-656.3386097542594','748.8403578656307',NULL,NULL,NULL,0,1),(56,'Block chain','1796.4704534387163','700.421241952652',NULL,NULL,'-6603.287282993331','-3590.335405994925','-415.31175748802025',NULL,NULL,NULL,0,1),(57,'Overhead','417.70302972341','425.33411832447575','-989.3283021226397','-621.3320208435873','-3111.5808170396704','-114.66231598591185','866.8210719505937','164.35378747602084','1134.055950173481','-753.9175926735629',1,1),(58,'Fair','232.52692883344275','671.5330491564459',NULL,NULL,'-5469.649462165422','-1106.5081701160102','688.9308524708213',NULL,NULL,NULL,0,1),(59,'Medium','320.36666614110965','672.6617658754263',NULL,NULL,'-4962.526245694544','-1110.5939508533318','750.2335288606791',NULL,NULL,NULL,0,1),(60,'High','414.8140466819174','672.4469216878636',NULL,NULL,'-4466.411357115115','-1226.125726118128','787.2822971600735',NULL,NULL,NULL,0,1),(61,'N/A','487.1906142580392','672.0880462774242',NULL,NULL,'-3918.0005729833483','-1244.1277199837466','842.8111976052214',NULL,NULL,NULL,0,1),(62,'Trust anchor','848.4620620412592','544.1450028790218','-982.6596771379395','-383.9977500175782','-1852.7437633603195','-1111.1540361743937','-97.31951956107605','-1496.7187517790915','-1076.584406805348','1016.42713987056',1,1),(63,'TPM','934.5379379587403','669.5063325898627',NULL,NULL,'-1184.9044012773659','-1789.3336295065346','572.4938645051229',NULL,NULL,NULL,0,1),(64,'SGX','1005.3626363895075','670.5063325898627',NULL,NULL,'-760.1823260827139','-1683.7777345675374','-319.8366899296602',NULL,NULL,NULL,0,1),(65,'Other','1077.3129235636318','673.3605973804222',NULL,NULL,'-294.31717100375323','-1603.1064742755818','528.323654467245',NULL,NULL,NULL,0,1),(66,'None','1710.7202051212796','700.8820213167838',NULL,NULL,'-8858.024127630688','-3467.2103127353244','-1518.0462141048988',NULL,NULL,NULL,0,1),(67,'Protection level','1295.8615880947477','548.2101522241273','-994.6595208891601','-281.33211459285496','996.9534648630378','-1705.7452487059247','512.4409064068868','-2534.5247518544247','256.0799143751864','-203.61227561712337',1,1),(68,'Internal','1637.7375791730985','699.6296638716772',NULL,NULL,'-7748.085581592561','-3476.806033232715','-1364.4036208165571',NULL,NULL,NULL,0,1),(69,'External','1368.6705318924207','700.6840866444072',NULL,NULL,'2529.0218352225365','-3007.2567771414238','285.6644084838954',NULL,NULL,NULL,0,1),(70,'Hypervisor','1228.9431625238847','700.7404915861159',NULL,NULL,'1583.1531063878915','-2645.529174665351','703.9929349935978',NULL,NULL,NULL,0,1),(71,'Software','729.4976450265727','669.7900057580437',NULL,NULL,'-2797.358495533419','-2152.495606294804','-225.68108123829825',NULL,NULL,NULL,0,1),(72,'Reverse engineering','475.2492122594372','254.81217069010142','-322.6662604198401','-776.0002187482914','-150.57370347755636','-715.1167132090086','-15.635084340648746','-197.80873206507522','-917.9333817485949','326.9834909860764',1,1),(73,'Attacker','475.12476237623423','368.1895132940248','-326.9977708507478','-283.333645830892','1144.1859704610501','240.24312276877822','-194.02308367976838','1024.0980845461536','-88.9569907091659','-1366.4392732671088',1,1),(78,'Call interposition','1150.347996774703','256.8145864563514',NULL,NULL,'-614.6855742024392','69.78717798899197','172.16880717031245',NULL,NULL,NULL,0,1),(79,'Disassembler','119.12950183558092','125.83901894908689',NULL,NULL,'-1995.7698478480065','-1269.2143260974635','-850.0941495725112',NULL,NULL,NULL,0,1),(80,'Tools','201.9306181258629','254.34923135684323','-325.33287500358057','-623.9985312614746','-1164.2709035330786','-836.8747820341652','-862.5478927977981','-1342.1191295012839','-114.96057988172356','-15.715835268157775',1,1),(81,'Debugger','241.45788738087992','128.4616763451635',NULL,NULL,'-898.1664589407969','-1951.0450139459258','-746.8683106090873',NULL,NULL,NULL,0,1),(82,'Tracer','340.07987666203144','127.12817980870784',NULL,NULL,'-1551.0856659947635','-478.5743355733096','-1242.1773531332879',NULL,NULL,NULL,0,1),(83,'Emulator','201.89365167201254','367.6398348095756',NULL,NULL,'-1150.7311957570096','-1400.1755468659949','217.24970035778017',NULL,NULL,NULL,0,1),(84,'Discovery','728.1518947222971','328.72709096376013','-320.002354148275','-889.3320416767579','134.6578654717242','-290.1059399368047','911.8235960463293','-107.54154854705472','-1454.515798633701','553.3090234561244',1,1),(85,'Pattern matching','728.5292554630064','453.4778093963763',NULL,NULL,'1037.20466085537','-142.0301624518263','1259.2820528498744',NULL,NULL,NULL,0,1),(86,'Taint analysis','919.9380630792848','451.93938973457784',NULL,NULL,'647.4909578714954','-571.3963813086439','1710.4840235215438',NULL,NULL,NULL,0,1),(87,'Graph-based analysis','1111.657682390608','453.1489477500765',NULL,NULL,'696.823432627275','-1526.425663903857','1189.8850144020796',NULL,NULL,NULL,0,1),(88,'Symbolic execution','1334.992734564465','451.1965639504301',NULL,NULL,'662.2672080618793','-1189.9103668925495','887.5497529647132',NULL,NULL,NULL,0,1),(89,'Dongle','848.9502871741241','670.352800129722',NULL,NULL,'-1758.2619761800556','-1832.5148037505676','517.4632417765356',NULL,NULL,NULL,0,1),(90,'Self-check','1498.733584648006','699.368986962789',NULL,NULL,'2254.929586909267','-2091.1030565912292','640.2494531206704',NULL,NULL,NULL,0,1),(91,'Hash chain','1575.609501237264','186.10961673256747',NULL,NULL,'5428.823081049158','-1392.7422049872685','622.1884914602401',NULL,NULL,NULL,0,1),(92,'Reactive','426.9542072101243','128.19895170523972',NULL,NULL,'-2156.9317741090754','3388.5632998016813','29.848118994063952',NULL,NULL,NULL,0,1),(93,'Asset','99.22342242131162','154.6749624574304','-686.6576250706369','-890.6693853954267','-4808.952804641889','1030.5020433257075','5.0788317295755405','3129.4424613034453','-1133.9728517827305','-488.9199597681988',1,1),(94,'Link','693.5984849922603','296.93075382868955',NULL,NULL,'-2652.1505220442427','412.8804880622755','131.98179481604484',NULL,NULL,NULL,0,1),(95,'Basic block','589.3599425928796','523.6192711996905',NULL,NULL,'-5634.276242050955','-2183.7509211360775','698.4873866131952',NULL,NULL,NULL,0,1);
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
  `id_taxonomy` int(11) unsigned NOT NULL,
  PRIMARY KEY (`id_dimension`),
  UNIQUE KEY `id_dimension_UNIQUE` (`id_dimension`),
  UNIQUE KEY `dimension_text_UNIQUE` (`text`,`id_taxonomy`),
  KEY `dimension_id_taxonomy_foreign` (`id_taxonomy`),
  CONSTRAINT `dimension_id_taxonomy_foreign` FOREIGN KEY (`id_taxonomy`) REFERENCES `taxonomy` (`id_taxonomy`) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `dimension`
--

LOCK TABLES `dimension` WRITE;
/*!40000 ALTER TABLE `dimension` DISABLE KEYS */;
INSERT INTO `dimension` VALUES (1,'System view','','','','',1),(2,'Attack view','','','','',1),(3,'Defense view','','','','',1),(4,'Interdimensional view','','','','',1);
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
  `id_paper` int(11) unsigned NOT NULL,
  `id_attribute` int(11) unsigned NOT NULL,
  `occurrenceCount` int(20) DEFAULT '1',
  PRIMARY KEY (`id_paper`,`id_attribute`),
  UNIQUE KEY `id_mapping_UNIQUE` (`id_mapping`),
  KEY `mapping_id_attribute_foreign` (`id_attribute`),
  KEY `mapping_id_mapping` (`id_mapping`),
  CONSTRAINT `mapping_id_attribute_foreign` FOREIGN KEY (`id_attribute`) REFERENCES `attribute` (`id_attribute`) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=678 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `mapping`
--

LOCK TABLES `mapping` WRITE;
/*!40000 ALTER TABLE `mapping` DISABLE KEYS */;
INSERT INTO `mapping` VALUES (1,1,2,1),(87,1,11,1),(103,1,15,1),(166,1,23,1),(199,1,26,1),(255,1,31,1),(265,1,33,1),(314,1,36,1),(360,1,39,1),(441,1,43,1),(531,1,58,1),(588,1,66,1),(628,1,68,1),(2,2,2,1),(62,2,9,1),(107,2,16,1),(167,2,23,1),(225,2,29,1),(266,2,33,1),(338,2,37,1),(361,2,39,1),(409,2,42,1),(458,2,45,1),(478,2,46,1),(503,2,53,1),(560,2,61,1),(589,2,66,1),(629,2,68,1),(3,3,2,1),(63,3,9,1),(108,3,16,1),(168,3,23,1),(201,3,26,1),(226,3,29,1),(268,3,33,1),(339,3,37,1),(363,3,39,1),(442,3,43,1),(459,3,45,1),(495,3,51,1),(505,3,53,1),(547,3,59,1),(591,3,66,1),(631,3,68,1),(4,4,2,1),(41,4,3,1),(51,4,4,1),(64,4,9,1),(137,4,19,1),(169,4,23,1),(227,4,29,1),(300,4,34,1),(340,4,37,1),(364,4,39,1),(443,4,43,1),(491,4,49,1),(511,4,55,1),(532,4,58,1),(585,4,64,1),(651,4,69,1),(5,5,2,1),(89,5,11,1),(105,5,15,1),(257,5,31,1),(269,5,33,1),(316,5,36,1),(365,5,39,1),(444,5,43,1),(561,5,61,1),(592,5,66,1),(652,5,69,1),(6,6,2,1),(42,6,3,1),(52,6,4,1),(65,6,9,1),(116,6,17,1),(170,6,23,1),(202,6,26,1),(228,6,29,1),(270,6,33,1),(317,6,36,1),(366,6,39,1),(445,6,43,1),(492,6,49,1),(512,6,55,1),(522,6,56,1),(553,6,60,1),(586,6,65,1),(671,6,70,1),(7,7,2,1),(54,7,6,1),(138,7,19,1),(171,7,23,1),(205,7,26,1),(218,7,28,1),(273,7,33,1),(341,7,37,1),(369,7,39,1),(412,7,42,1),(479,7,46,1),(564,7,61,1),(595,7,66,1),(653,7,69,1),(8,8,2,1),(66,8,9,1),(109,8,16,1),(172,8,23,1),(229,8,29,1),(274,8,33,1),(342,8,37,1),(370,8,39,1),(413,8,42,1),(460,8,45,1),(496,8,51,1),(554,8,60,1),(596,8,66,1),(634,8,68,1),(9,9,2,1),(90,9,11,1),(130,9,18,1),(163,9,22,1),(249,9,30,1),(258,9,31,1),(275,9,33,1),(320,9,36,1),(371,9,39,1),(414,9,42,1),(565,9,61,1),(597,9,66,1),(635,9,68,1),(10,10,2,1),(139,10,19,1),(173,10,23,1),(206,10,26,1),(276,10,33,1),(343,10,37,1),(372,10,39,1),(415,10,42,1),(493,10,49,1),(513,10,55,1),(533,10,58,1),(598,10,66,1),(672,10,70,1),(11,11,2,1),(55,11,6,1),(117,11,17,1),(197,11,25,1),(219,11,28,1),(301,11,34,1),(373,11,39,1),(447,11,43,1),(500,11,52,1),(534,11,58,1),(599,11,66,1),(654,11,69,1),(12,12,2,1),(56,12,6,1),(110,12,16,1),(220,12,28,1),(277,12,33,1),(344,12,37,1),(374,12,39,1),(416,12,42,1),(485,12,48,1),(566,12,61,1),(600,12,66,1),(636,12,68,1),(13,13,2,1),(57,13,6,1),(140,13,19,1),(174,13,23,1),(207,13,26,1),(221,13,28,1),(302,13,34,1),(345,13,37,1),(375,13,39,1),(417,13,42,1),(461,13,45,1),(567,13,61,1),(601,13,66,1),(655,13,69,1),(14,14,2,1),(45,14,3,1),(53,14,4,1),(67,14,9,1),(141,14,19,1),(156,14,21,1),(230,14,29,1),(278,14,33,1),(346,14,37,1),(376,14,39,1),(418,14,42,1),(462,14,45,1),(514,14,55,1),(523,14,56,1),(535,14,58,1),(579,14,63,1),(673,14,70,1),(15,15,2,1),(118,15,17,1),(175,15,23,1),(279,15,33,1),(347,15,37,1),(377,15,39,1),(419,15,42,1),(515,15,55,1),(548,15,59,1),(602,15,66,1),(674,15,70,1),(16,16,2,1),(68,16,9,1),(111,16,16,1),(177,16,23,1),(231,16,29,1),(280,16,33,1),(348,16,37,1),(379,16,39,1),(421,16,42,1),(497,16,51,1),(516,16,55,1),(536,16,58,1),(604,16,66,1),(637,16,68,1),(17,17,2,1),(69,17,9,1),(112,17,16,1),(178,17,23,1),(196,17,24,1),(232,17,29,1),(281,17,33,1),(349,17,37,1),(380,17,39,1),(422,17,42,1),(463,17,45,1),(498,17,51,1),(501,17,52,1),(517,17,55,1),(537,17,58,1),(605,17,66,1),(638,17,68,1),(18,18,2,1),(70,18,9,1),(131,18,18,1),(179,18,23,1),(233,18,29,1),(282,18,33,1),(350,18,37,1),(381,18,39,1),(423,18,42,1),(464,18,45,1),(569,18,61,1),(606,18,66,1),(639,18,68,1),(19,19,2,1),(83,19,10,1),(120,19,17,1),(157,19,21,1),(251,19,30,1),(283,19,33,1),(322,19,36,1),(382,19,39,1),(424,19,42,1),(487,19,48,1),(499,19,51,1),(570,19,61,1),(607,19,66,1),(640,19,68,1),(20,20,2,1),(91,20,11,1),(113,20,16,1),(132,20,18,1),(180,20,23,1),(259,20,31,1),(284,20,33,1),(323,20,36,1),(383,20,39,1),(425,20,42,1),(509,20,54,1),(555,20,60,1),(608,20,66,1),(641,20,68,1),(21,21,2,1),(96,21,12,1),(142,21,19,1),(181,21,23,1),(234,21,29,1),(304,21,34,1),(324,21,36,1),(384,21,39,1),(448,21,43,1),(465,21,45,1),(556,21,60,1),(609,21,66,1),(657,21,69,1),(22,22,2,1),(97,22,12,1),(143,22,19,1),(182,22,23,1),(235,22,29,1),(305,22,34,1),(325,22,36,1),(385,22,39,1),(449,22,43,1),(466,22,45,1),(549,22,59,1),(610,22,66,1),(658,22,69,1),(23,23,2,1),(92,23,11,1),(144,23,19,1),(158,23,21,1),(260,23,31,1),(306,23,34,1),(326,23,36,1),(407,23,40,1),(450,23,43,1),(467,23,45,1),(480,23,46,1),(488,23,48,1),(524,23,56,1),(550,23,59,1),(611,23,66,1),(642,23,68,1),(24,24,2,1),(71,24,9,1),(133,24,18,1),(164,24,22,1),(236,24,29,1),(285,24,33,1),(351,24,37,1),(386,24,39,1),(451,24,43,1),(468,24,45,1),(518,24,55,1),(571,24,61,1),(612,24,66,1),(643,24,68,1),(25,25,2,1),(72,25,9,1),(145,25,19,1),(183,25,23,1),(237,25,29,1),(286,25,33,1),(352,25,37,1),(387,25,39,1),(426,25,42,1),(469,25,45,1),(525,25,56,1),(580,25,63,1),(659,25,69,1),(26,26,2,1),(58,26,6,1),(146,26,19,1),(184,26,23,1),(210,26,26,1),(222,26,28,1),(287,26,33,1),(353,26,37,1),(390,26,39,1),(427,26,42,1),(481,26,46,1),(572,26,61,1),(614,26,66,1),(662,26,69,1),(27,27,2,1),(73,27,9,1),(147,27,19,1),(185,27,23,1),(211,27,26,1),(238,27,29,1),(288,27,33,1),(354,27,37,1),(391,27,39,1),(454,27,43,1),(470,27,45,1),(482,27,46,1),(519,27,55,1),(573,27,61,1),(615,27,66,1),(663,27,69,1),(28,28,2,1),(93,28,11,1),(122,28,17,1),(161,28,21,1),(261,28,31,1),(289,28,33,1),(392,28,39,1),(428,28,42,1),(510,28,54,1),(574,28,61,1),(616,28,66,1),(644,28,68,1),(29,29,2,1),(74,29,9,1),(135,29,18,1),(198,29,25,1),(239,29,29,1),(290,29,33,1),(355,29,37,1),(393,29,39,1),(455,29,43,1),(502,29,52,1),(506,29,53,1),(551,29,59,1),(617,29,66,1),(645,29,68,1),(30,30,2,1),(101,30,13,1),(186,30,23,1),(212,30,26,1),(262,30,31,1),(291,30,33,1),(329,30,36,1),(394,30,39,1),(429,30,42,1),(484,30,47,1),(527,30,56,1),(539,30,58,1),(582,30,63,1),(646,30,68,1),(31,31,2,1),(98,31,12,1),(123,31,17,1),(187,31,23,1),(309,31,34,1),(330,31,36,1),(395,31,39,1),(430,31,42,1),(471,31,45,1),(575,31,61,1),(618,31,66,1),(664,31,69,1),(32,32,2,1),(75,32,9,1),(148,32,19,1),(188,32,23,1),(240,32,29,1),(292,32,33,1),(331,32,36,1),(396,32,39,1),(431,32,42,1),(472,32,45,1),(520,32,55,1),(528,32,56,1),(540,32,58,1),(587,32,65,1),(675,32,70,1),(33,33,2,1),(60,33,6,1),(124,33,17,1),(191,33,23,1),(215,33,26,1),(241,33,29,1),(311,33,34,1),(357,33,37,1),(399,33,39,1),(433,33,42,1),(530,33,56,1),(543,33,58,1),(584,33,63,1),(676,33,70,1),(34,34,2,1),(76,34,9,1),(125,34,17,1),(192,34,23,1),(242,34,29,1),(295,34,33,1),(358,34,37,1),(400,34,39,1),(434,34,42,1),(474,34,45,1),(507,34,53,1),(558,34,60,1),(621,34,66,1),(648,34,68,1),(35,35,2,1),(99,35,12,1),(126,35,17,1),(193,35,23,1),(312,35,34,1),(334,35,36,1),(401,35,39,1),(435,35,42,1),(475,35,45,1),(544,35,58,1),(622,35,66,1),(667,35,69,1),(36,36,2,1),(77,36,9,1),(100,36,12,1),(102,36,13,1),(127,36,17,1),(162,36,21,1),(243,36,29,1),(296,36,33,1),(335,36,36,1),(402,36,39,1),(436,36,42,1),(577,36,61,1),(623,36,66,1),(668,36,69,1),(37,37,2,1),(61,37,6,1),(136,37,18,1),(165,37,22,1),(217,37,26,1),(224,37,28,1),(297,37,33,1),(337,37,36,1),(404,37,39,1),(438,37,42,1),(559,37,60,1),(625,37,66,1),(649,37,68,1),(38,38,2,1),(78,38,9,1),(115,38,16,1),(194,38,23,1),(244,38,29,1),(298,38,33,1),(405,38,39,1),(439,38,42,1),(476,38,45,1),(508,38,53,1),(552,38,59,1),(626,38,66,1),(650,38,68,1),(39,39,2,1),(79,39,9,1),(152,39,19,1),(195,39,23,1),(245,39,29,1),(299,39,33,1),(359,39,37,1),(406,39,39,1),(440,39,42,1),(477,39,45,1),(521,39,55,1),(545,39,58,1),(627,39,66,1),(670,39,69,1),(677,39,70,1),(40,40,3,1),(88,40,11,1),(104,40,15,1),(128,40,18,1),(153,40,21,1),(200,40,26,1),(246,40,30,1),(256,40,31,1),(267,40,33,1),(315,40,36,1),(362,40,39,1),(410,40,42,1),(490,40,49,1),(504,40,53,1),(546,40,59,1),(590,40,66,1),(630,40,68,1),(43,41,3,1),(80,41,10,1),(129,41,18,1),(154,41,21,1),(203,41,26,1),(247,41,30,1),(271,41,33,1),(318,41,36,1),(367,41,39,1),(446,41,43,1),(562,41,61,1),(593,41,66,1),(632,41,68,1),(44,42,3,1),(81,42,10,1),(106,42,15,1),(155,42,21,1),(204,42,26,1),(248,42,30,1),(272,42,33,1),(319,42,36,1),(368,42,39,1),(411,42,42,1),(563,42,61,1),(594,42,66,1),(633,42,68,1),(46,43,3,1),(82,43,10,1),(119,43,17,1),(176,43,23,1),(208,43,26,1),(250,43,30,1),(303,43,34,1),(321,43,36,1),(378,43,39,1),(420,43,42,1),(486,43,48,1),(568,43,61,1),(603,43,66,1),(656,43,69,1),(47,44,3,1),(84,44,10,1),(121,44,17,1),(159,44,21,1),(252,44,30,1),(307,44,34,1),(327,44,36,1),(388,44,39,1),(452,44,43,1),(557,44,60,1),(613,44,66,1),(660,44,69,1),(48,45,3,1),(85,45,10,1),(134,45,18,1),(160,45,21,1),(209,45,26,1),(253,45,30,1),(308,45,34,1),(328,45,36,1),(389,45,39,1),(453,45,43,1),(483,45,47,1),(526,45,56,1),(538,45,58,1),(581,45,63,1),(661,45,69,1),(49,46,3,1),(59,46,6,1),(149,46,19,1),(223,46,28,1),(310,46,34,1),(356,46,37,1),(408,46,40,1),(432,46,42,1),(473,46,45,1),(529,46,56,1),(576,46,61,1),(583,46,63,1),(665,46,69,1),(50,47,3,1),(86,47,10,1),(151,47,19,1),(216,47,26,1),(254,47,30,1),(313,47,34,1),(336,47,36,1),(403,47,39,1),(437,47,42,1),(489,47,48,1),(578,47,61,1),(624,47,66,1),(669,47,69,1),(94,48,11,1),(114,48,16,1),(189,48,23,1),(213,48,26,1),(263,48,31,1),(293,48,33,1),(332,48,36,1),(397,48,39,1),(456,48,43,1),(541,48,58,1),(619,48,66,1),(647,48,68,1),(95,49,11,1),(150,49,19,1),(190,49,23,1),(214,49,26,1),(264,49,31,1),(294,49,33,1),(333,49,36,1),(398,49,39,1),(457,49,43,1),(494,49,49,1),(542,49,58,1),(620,49,66,1),(666,49,69,1);
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
INSERT INTO `paper` VALUES (1,'abadi2005control','{\"author\": \"Abadi, Mart{\\\\\'\\\\i}n and Budiu, Mihai and Erlingsson, Ulfar and Ligatti, Jay\", \"booktitle\": \"Proceedings of the 12th ACM conference on Computer and communications security\", \"title\": \"Control-flow integrity\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2005\", \"organization\": \"ACM\", \"ID\": \"abadi2005control\", \"pages\": \"340--353\"}',5),(2,'aucsmith1996tamper','{\"isbn\": \"3-540-61996-8\", \"title\": \"Tamper resistant software: An implementation\", \"journal\": \"Proceedings of the First International Workshop on Information Hiding\", \"author\": \"Aucsmith, David\", \"ENTRYTYPE\": \"article\", \"mendeley-groups\": \"Tamperproofing,Tamperproofing/Methods\", \"link\": \"http://link.springer.com/chapter/10.1007/3-540-61996-8{\\\\_}49\", \"year\": \"1996\", \"ID\": \"aucsmith1996tamper\", \"pages\": \"317--333\"}',161),(3,'banescu2017detecting','{\"doi\": \"10.1145/3029806.3029835\", \"isbn\": \"978-1-4503-4523-1/17/03\", \"title\": \"Detecting Patching of Executables without System Calls\", \"booktitle\": \"Proceedings of the Conference on Data and Application Security and Privacy\", \"author\": \"Banescu, Sebastian and Ahmadvand, Mohsen and Pretschner, Alexander and Shield, Robert and Hamilton, Chris\", \"ID\": \"banescu2017detecting\", \"year\": \"2017\", \"ENTRYTYPE\": \"inproceedings\"}',0),(4,'baumann2015shielding','{\"publisher\": \"ACM\", \"author\": \"Baumann, Andrew and Peinado, Marcus and Hunt, Galen\", \"journal\": \"ACM Transactions on Computer Systems (TOCS)\", \"title\": \"Shielding applications from an untrusted cloud with haven\", \"number\": \"3\", \"ENTRYTYPE\": \"article\", \"volume\": \"33\", \"year\": \"2015\", \"ID\": \"baumann2015shielding\", \"pages\": \"8\"}',0),(5,'Blietz2006','{\"doi\": \"10.1007/11787952_12\", \"isbn\": \"3540359982\", \"author\": \"Blietz, Brian and Tyagi, Akhilesh\", \"journal\": \"Lecture Notes in Computer Science (including subseries Lecture Notes in Artificial Intelligence and Lecture Notes in Bioinformatics)\", \"issn\": \"16113349\", \"ENTRYTYPE\": \"article\", \"volume\": \"3919 LNCS\", \"mendeley-groups\": \"Tamperproofing/Methods\", \"file\": \":Users/mohsen-tum/Documents/Mendeley Desktop/Blietz, Tyagi/Lecture Notes in Computer Science (including subseries Lecture Notes in Artificial Intelligence and Lecture Notes in Bioinformatics)/Blietz, Tyagi - 2006 - Software tamper resistance through dynamic program monitoring.pdf:pdf\", \"year\": \"2006\", \"title\": \"Software tamper resistance through dynamic program monitoring\", \"ID\": \"Blietz2006\", \"pages\": \"146--163\"}',5),(6,'brasser2015tytan','{\"author\": \"Brasser, Ferdinand and El Mahjoub, Brahim and Sadeghi, Ahmad-Reza and Wachsmann, Christian and Koeberl, Patrick\", \"booktitle\": \"2015 52nd ACM/EDAC/IEEE Design Automation Conference (DAC)\", \"title\": \"TyTAN: tiny trust anchor for tiny devices\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2015\", \"organization\": \"IEEE\", \"ID\": \"brasser2015tytan\", \"pages\": \"1--6\"}',0),(7,'catuogno2002format','{\"author\": \"Catuogno, Luigi and Visconti, Ivan\", \"booktitle\": \"International Conference on Security in Communication Networks\", \"title\": \"A format-independent architecture for run-time integrity checking of executable code\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2002\", \"organization\": \"Springer\", \"ID\": \"catuogno2002format\", \"pages\": \"219--233\"}',9),(8,'chang2001protecting','{\"author\": \"Chang, Hoi and Atallah, Mikhail J\", \"booktitle\": \"ACM Workshop on Digital Rights Management\", \"title\": \"Protecting software code by guards\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2001\", \"organization\": \"Springer\", \"ID\": \"chang2001protecting\", \"pages\": \"160--175\"}',93),(9,'chen2002oblivious','{\"author\": \"Chen, Yuqun and Venkatesan, Ramarathnam and Cary, Matthew and Pang, Ruoming and Sinha, Saurabh and Jakubowski, Mariusz H\", \"booktitle\": \"International Workshop on Information Hiding\", \"title\": \"Oblivious hashing: A stealthy software integrity verification primitive\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2002\", \"organization\": \"Springer\", \"ID\": \"chen2002oblivious\", \"pages\": \"400--414\"}',55),(10,'christodorescu2009cloud','{\"author\": \"Christodorescu, Mihai and Sailer, Reiner and Schales, Douglas Lee and Sgandurra, Daniele and Zamboni, Diego\", \"booktitle\": \"Proceedings of the 2009 ACM workshop on Cloud computing security\", \"title\": \"Cloud security is not (just) virtualization security: a short paper\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2009\", \"organization\": \"ACM\", \"ID\": \"christodorescu2009cloud\", \"pages\": \"97--102\"}',0),(11,'collberg2012distributed','{\"author\": \"Collberg, Christian and Martin, Sam and Myers, Jonathan and Nagra, Jasvir\", \"booktitle\": \"Proceedings of the 28th Annual Computer Security Applications Conference\", \"title\": \"Distributed application tamper detection via continuous software updates\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2012\", \"organization\": \"ACM\", \"ID\": \"collberg2012distributed\", \"pages\": \"319--328\"}',15),(12,'dedic2007graph','{\"author\": \"Dedi{\\\\\'c}, Nenad and Jakubowski, Mariusz and Venkatesan, Ramarathnam\", \"booktitle\": \"International Workshop on Information Hiding\", \"title\": \"A graph game model for software tamper protection\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2007\", \"organization\": \"Springer\", \"ID\": \"dedic2007graph\", \"pages\": \"80--95\"}',5),(13,'deswarte2004remote','{\"publisher\": \"Springer\", \"author\": \"Deswarte, Yves and Quisquater, Jean-Jacques and Sa{\\\\\\\"\\\\i}dane, Ayda\", \"booktitle\": \"Integrity and internal control in information systems VI\", \"title\": \"Remote integrity checking\", \"ENTRYTYPE\": \"incollection\", \"year\": \"2004\", \"ID\": \"deswarte2004remote\", \"pages\": \"1--11\"}',34),(14,'dewan2008hypervisor','{\"author\": \"Dewan, Prashant and Durham, David and Khosravi, Hormuzd and Long, Men and Nagabhushan, Gayathri\", \"booktitle\": \"Proceedings of the 2008 Spring simulation multiconference\", \"title\": \"A hypervisor-based system for protecting software runtime memory and persistent storage\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2008\", \"organization\": \"Society for Computer Simulation International\", \"ID\": \"dewan2008hypervisor\", \"pages\": \"828--835\"}',22),(15,'Gan2015using','{\"doi\": \"10.1109/SPRO.2015.12\", \"title\": \"Using Virtual Machine Protections to Enhance Whitebox Cryptography\", \"booktitle\": \"Software Protection (SPRO), 2015 IEEE/ACM 1st International Workshop on\", \"author\": \"J. Gan and R. Kok and P. Kohli and Y. Ding and B. Mah\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2015\", \"ID\": \"Gan2015using\", \"pages\": \"17-23\"}',0),(16,'Ghosh2010secure','{\"isbn\": \"364216434X\", \"author\": \"Ghosh, Sudeep and Hiser, Jason D. and Davidson, Jack W.\", \"journal\": \"Lecture Notes in Computer Science (including subseries Lecture Notes in Artificial Intelligence and Lecture Notes in Bioinformatics)\", \"issn\": \"03029743\", \"ENTRYTYPE\": \"article\", \"volume\": \"6387 LNCS\", \"mendeley-groups\": \"Tamperproofing\", \"year\": \"2010\", \"title\": \"A secure and robust approach to software tamper resistance\", \"ID\": \"Ghosh2010secure\", \"pages\": \"33--47\"}',11),(17,'ghosh2013software','{\"author\": \"Ghosh, Sudeep and Hiser, Jason and Davidson, Jack W\", \"booktitle\": \"Proceedings of the 2nd ACM SIGPLAN Program Protection and Reverse Engineering Workshop\", \"title\": \"Software protection for dynamically-generated code\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2013\", \"organization\": \"ACM\", \"ID\": \"ghosh2013software\", \"pages\": \"1\"}',6),(18,'Horne2002','{\"doi\": \"10.1007/3-540-47870-1_9\", \"isbn\": \"978-3-540-43677-5\", \"author\": \"Horne, Bill and Matheson, Lesley and Sheehan, Casey and Tarjan, Robert\", \"ENTRYTYPE\": \"article\", \"abstract\": \"We describe a software self-checking mechanism designed to improve the tamper resistance of large programs. The mechanism consists of a number of testers that redundantly test for changes in the executable code as it is running and report modifications. The mechanism is built to be compatible with copy-specific static watermarking and other tamper-resistance techniques. The mechanism includes several innovations to make it stealthy and more robust.\", \"title\": \"Dynamic Self-Checking Techniques for Improved Tamper Resistance\", \"pages\": \"141--159\", \"mendeley-groups\": \"Tamperproofing/Methods,Tamperproofing\", \"link\": \"http://citeseerx.ist.psu.edu/viewdoc/summary?doi=10.1.1.13.3308\", \"file\": \":Users/mohsen-tum/Documents/Mendeley Desktop/Horne et al/Security and Privacy in Digital Rights Management/Horne et al. - 2002 - Dynamic Self-Checking Techniques for Improved Tamper Resistance.pdf:pdf\", \"year\": \"2002\", \"ID\": \"Horne2002\", \"annote\": \"They add testers in the post compilation process.\\nLinear checks no circular\\nTo avoid complexity, a block is checked only by one block\\nA 32bit space is added outside basic blocks as corrector that tries to fix the hash values in patch process. The patch process is part of sofware watermarking after-installation process\\nDid not quite get it where do they store hashes? They say we store them but not clear where?!\\nNo inidication of how Address space layout randomization is respected.\", \"journal\": \"Security and Privacy in Digital Rights Management\"}',0),(19,'ibrahim2016stins4cs','{\"author\": \"Ibrahim, Amjad and Banescu, Sebastian\", \"booktitle\": \"Proceedings of the 2016 ACM Workshop on Software PROtection\", \"title\": \"StIns4CS: A State Inspection Tool for C\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2016\", \"organization\": \"ACM\", \"ID\": \"ibrahim2016stins4cs\", \"pages\": \"61--71\"}',0),(20,'jacob2007towards','{\"author\": \"Jacob, Matthias and Jakubowski, Mariusz H and Venkatesan, Ramarathnam\", \"booktitle\": \"Proceedings of the 9th workshop on Multimedia \\\\& security\", \"title\": \"Towards integral binary execution: Implementing oblivious hashing using overlapped instruction encodings\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2007\", \"organization\": \"ACM\", \"ID\": \"jacob2007towards\", \"pages\": \"129--140\"}',25),(21,'jakobsson2010retroactive','{\"numpages\": \"13\", \"publisher\": \"USENIX Association\", \"title\": \"Retroactive Detection of Malware with Applications to Mobile Platforms\", \"series\": \"HotSec\'10\", \"booktitle\": \"Proceedings of the 5th USENIX Conference on Hot Topics in Security\", \"author\": \"Jakobsson, Markus and Johansson, Karl-Anders\", \"ENTRYTYPE\": \"inproceedings\", \"location\": \"Washinton, DC\", \"year\": \"2010\", \"ID\": \"jakobsson2010retroactive\", \"pages\": \"1--13\", \"address\": \"Berkeley, CA, USA\"}',0),(22,'jakobsson2011practical','{\"author\": \"Jakobsson, Markus and Johansson, Karl-Anders\", \"booktitle\": \"Lightweight Security \\\\& Privacy: Devices, Protocols and Applications (LightSec), 2011 Workshop on\", \"title\": \"Practical and secure software-based attestation\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2011\", \"organization\": \"IEEE\", \"ID\": \"jakobsson2011practical\", \"pages\": \"1--9\"}',13),(23,'jin2003forensic','{\"author\": \"Jin, Hongxia and Lotspiech, Jeffery\", \"booktitle\": \"Software Reliability Engineering, 2003. ISSRE 2003. 14th International Symposium on\", \"title\": \"Forensic analysis for tamper resistant software\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2003\", \"organization\": \"IEEE\", \"ID\": \"jin2003forensic\", \"pages\": \"133--142\"}',9),(24,'junod2015obfuscator','{\"author\": \"Junod, Pascal and Rinaldini, Julien and Wehrli, Johan and Michielin, Julie\", \"booktitle\": \"Proceedings of the 1st International Workshop on Software Protection\", \"title\": \"Obfuscator-LLVM: software protection for the masses\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2015\", \"organization\": \"IEEE Press\", \"ID\": \"junod2015obfuscator\", \"pages\": \"3--9\"}',0),(25,'kanstren2015architecture','{\"author\": \"Kanstr{\\\\\'e}n, Teemu and Lehtonen, Sami and Savola, Reijo and Kukkohovi, Hilkka and H{\\\\\\\"a}t{\\\\\\\"o}nen, Kimmo\", \"booktitle\": \"Cloud Engineering (IC2E), 2015 IEEE International Conference on\", \"title\": \"Architecture for high confidence cloud security monitoring\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2015\", \"organization\": \"IEEE\", \"ID\": \"kanstren2015architecture\", \"pages\": \"195--200\"}',3),(26,'kim1994experiences','{\"ID\": \"kim1994experiences\", \"author\": \"Kim, Gene H and Spafford, Eugene H\", \"year\": \"1994\", \"ENTRYTYPE\": \"article\", \"title\": \"Experiences with tripwire: Using integrity checkers for intrusion detection\"}',0),(27,'kimball2012emulation','{\"publisher\": \"Google Patents\", \"author\": \"Kimball, William B and Baldwin, Rusty O\", \"title\": \"Emulation-based software protection\", \"month\": \"oct~9\", \"note\": \"US Patent 8,285,987\", \"year\": \"2012\", \"ID\": \"kimball2012emulation\", \"ENTRYTYPE\": \"misc\"}',0),(28,'kulkarni2014new','{\"author\": \"Kulkarni, Aniket and Metta, Ravindra\", \"booktitle\": \"Service Oriented System Engineering (SOSE), 2014 IEEE 8th International Symposium on\", \"title\": \"A New Code Obfuscation Scheme for Software Protection\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2014\", \"organization\": \"IEEE\", \"ID\": \"kulkarni2014new\", \"pages\": \"409--414\"}',0),(29,'madou2005software','{\"author\": \"Madou, Matias and Anckaert, Bertrand and Moseley, Patrick and Debray, Saumya and De Sutter, Bjorn and De Bosschere, Koen\", \"booktitle\": \"International Workshop on Information Security Applications\", \"title\": \"Software protection through dynamic code mutation\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2005\", \"organization\": \"Springer\", \"ID\": \"madou2005software\", \"pages\": \"194--206\"}',31),(30,'Malone2011','{\"doi\": \"10.1145/2046582.2046596\", \"isbn\": \"9781450310017\", \"keyword\": \"hardware performance counters,integrity\", \"author\": \"Malone, Corey and Zahran, Mohamed and Karri, Ramesh\", \"journal\": \"Proceedings of the sixth ACM workshop on Scalable trusted computing - STC \'11\", \"issn\": \"15437221\", \"ENTRYTYPE\": \"article\", \"mendeley-groups\": \"Tamperproofing/Methods\", \"link\": \"http://www.scopus.com/inward/record.url?eid=2-s2.0-80755143408{\\\\&}partnerID=40{\\\\&}md5=ad5db1f8e5c0131a2a17f457ba1b0497$\\\\backslash$nhttp://dl.acm.org/citation.cfm?doid=2046582.2046596\", \"file\": \":Users/mohsen-tum/Documents/Mendeley Desktop/Malone, Zahran, Karri/Proceedings of the sixth ACM workshop on Scalable trusted computing - STC \'11/Malone, Zahran, Karri - 2011 - Are Hardware Performance Counters a Cost Effective Way for Integrity Checking of Programs.pdf:pdf\", \"year\": \"2011\", \"title\": \"Are Hardware Performance Counters a Cost Effective Way for Integrity Checking of Programs\", \"ID\": \"Malone2011\", \"pages\": \"71\"}',0),(31,'Martignoni2010conquer','{\"doi\": \"10.1007/978-3-642-14215-4_2\", \"title\": \"Conqueror: Tamper-proof code execution on legacy systems\", \"journal\": \"Lecture Notes in Computer Science (including subseries Lecture Notes in Artificial Intelligence and Lecture Notes in Bioinformatics)\", \"author\": \"Martignoni, Lorenzo and Paleari, Roberto and Bruschi, Danilo\", \"ENTRYTYPE\": \"article\", \"volume\": \"6201 LNCS\", \"year\": \"2010\", \"ID\": \"Martignoni2010conquer\", \"pages\": \"21--40\"}',12),(32,'morgan2015design','{\"author\": \"Morgan, Beno{\\\\^\\\\i}t and Alata, Eric and Nicomette, Vincent and Ka{\\\\^a}niche, Mohamed and Averlant, Guillaume\", \"booktitle\": \"Dependable Computing (PRDC), 2015 IEEE 21st Pacific Rim International Symposium on\", \"title\": \"Design and implementation of a hardware assisted security architecture for software integrity monitoring\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2015\", \"organization\": \"IEEE\", \"ID\": \"morgan2015design\", \"pages\": \"189--198\"}',0),(33,'park2015tgvisor','{\"author\": \"Park, Sungjin and Yoon, Jae Nam and Kang, Cheoloh and Kim, Kyong Hoon and Han, Taisook\", \"booktitle\": \"Mobile Cloud Computing, Services, and Engineering (MobileCloud), 2015 3rd IEEE International Conference on\", \"title\": \"TGVisor: A Tiny Hypervisor-Based Trusted Geolocation Framework for Mobile Cloud Clients\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2015\", \"organization\": \"IEEE\", \"ID\": \"park2015tgvisor\", \"pages\": \"99--108\"}',0),(34,'Protsenko2015dynamic','{\"doi\": \"10.1109/ARES.2015.98\", \"keyword\": \"Android (operating system);computer crime;cryptography;mobile computing;reverse engineering;Android apps;application piracy;dynamic code loading;dynamic obfuscation techniques;dynamic re-encryption;dynamic self-protection;mobile devices;native code;proprietary mobile software;reverse engineering;tamperproofing;Androids;Encryption;Humanoid robots;Loading;Runtime;Software protection;Android;Software Protection\", \"title\": \"Dynamic Self-Protection and Tamperproofing for Android Apps Using Native Code\", \"booktitle\": \"Availability, Reliability and Security (ARES), 2015 10th International Conference on\", \"author\": \"M. Protsenko and S. Kreuter and T. M\\u00fcller\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2015\", \"ID\": \"Protsenko2015dynamic\", \"pages\": \"129-138\"}',0),(35,'Seshadri2005pioneer','{\"doi\": \"10.1145/1095809.1095812\", \"isbn\": \"1-59593-079-5\", \"keyword\": \"dynamic root of trust,rootkit detection,self-check-summing code,software-based code attestation,verifiable code execution\", \"author\": \"Seshadri, Arvind and Luk, Mark and Shi, Elaine and Perrig, Adrian and van Doorn, Leendert and Khosla, Pradeep\", \"journal\": \"ACM SIGOPS Operating Systems Review\", \"issn\": \"01635980\", \"ID\": \"Seshadri2005pioneer\", \"mendeley-groups\": \"Tamperproofing\", \"link\": \"http://dl.acm.org/citation.cfm?id=1095809.1095812\", \"year\": \"2005\", \"title\": \"Pioneer: Verifying Code Integrity and Enforcing Untampered Code Execution on Legacy Systems\", \"ENTRYTYPE\": \"article\"}',2),(36,'Spinellis2000','{\"doi\": \"10.1145/353323.353383\", \"isbn\": \"1094-9224\", \"author\": \"Spinellis, Diomidis\", \"ENTRYTYPE\": \"article\", \"abstract\": \"The integrity verification of a device\'s controlling software is an important aspect of many emerging information appliances. We propose the use of reflection, whereby the software is able to examine its own operation, in conjunction with cryptographic hashes as a basis for developing a suitable software verification protocol. For more demanding applications meta-reflective techniques can be used to thwart attacks based on device emulation strategies. We demonstrate how our approach can be used to increase the security of mobile phones, devices for the delivery of digital content, and smartcards.\", \"issn\": \"10949224\", \"number\": \"1\", \"pages\": \"51--62\", \"volume\": \"3\", \"mendeley-groups\": \"Tamperproofing/Methods\", \"file\": \":Users/mohsen-tum/Documents/Mendeley Desktop/Spinellis/ACM Transactions on Information and System Security/Spinellis - 2000 - Reflection as a mechanism for software integrity verification.pdf:pdf\", \"year\": \"2000\", \"title\": \"Reflection as a mechanism for software integrity verification\", \"ID\": \"Spinellis2000\", \"annote\": \"In this approach a software integrity is verified with the help of an external (trusted) entity. Here, the program state is retrieved using reflection, a protocol is proposed to verify the state, and suggested to augment the scheme with CPU perfor.mance counter, before and after the verification call loops.\\nOne obvious attack is to keep an untouched version of the application in the memory next to the tampered with version. Then redirect all hash computations to the good version. The authors, suggest memory expanion and timing as possible countermeasures.\", \"journal\": \"ACM Transactions on Information and System Security\"}',0),(37,'teixeira2015siot','{\"author\": \"Teixeira, Fernando A and Machado, Gustavo V and Pereira, Fernando MQ and Wong, Hao Chi and Nogueira, Jos{\\\\\'e} and Oliveira, Leonardo B\", \"booktitle\": \"Proceedings of the 14th International Conference on Information Processing in Sensor Networks\", \"title\": \"SIoT: securing the internet of things through distributed system analysis\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2015\", \"organization\": \"ACM\", \"ID\": \"teixeira2015siot\", \"pages\": \"310--321\"}',0),(38,'Wang2005Tamper','{\"isbn\": \"8242866627\", \"keyword\": \"integrity checking,multi-blocking encryption,software piracy,tamper resistant\", \"title\": \"Tamper Resistant Software Through Dynamic Integrity Checking\", \"journal\": \"Proc. Symp. on Cyptography and Information Security (SCIS 05)\", \"author\": \"Wang, Ping and Kang, Seok-kyu and Kim, Kwangjo\", \"ID\": \"Wang2005Tamper\", \"year\": \"2005\", \"ENTRYTYPE\": \"article\"}',0),(39,'yao2014cryptvmi','{\"author\": \"Yao, Fangzhou and Sprabery, Read and Campbell, Roy H\", \"booktitle\": \"Proceedings of the 2nd international workshop on Security in cloud computing\", \"title\": \"CryptVMI: a flexible and encrypted virtual machine introspection system in the cloud\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2014\", \"organization\": \"ACM\", \"ID\": \"yao2014cryptvmi\", \"pages\": \"11--18\"}',0),(40,'banescu2015software','{\"author\": \"Banescu, Sebastian and Pretschner, Alexander and Battr{\\\\\'e}, Dominic and Cazzulani, St{\\\\\'e}fano and Shield, Robert and Thompson, Greg\", \"booktitle\": \"Proceedings of the 5th ACM Conference on Data and Application Security and Privacy\", \"title\": \"Software-based protection against changeware\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2015\", \"organization\": \"ACM\", \"ID\": \"banescu2015software\", \"pages\": \"231--242\"}',4),(41,'Carbone2009','{\"isbn\": \"9781605583525\", \"author\": \"Carbone, Martim and Cui, Weidong and Peinado, Marcus and Lu, Long and Lee, Wenke\", \"journal\": \"Analysis\", \"title\": \"Mapping Kernel Objects to Enable Systematic Integrity Checking\", \"ENTRYTYPE\": \"article\", \"mendeley-groups\": \"Tamperproofing/Methods\", \"file\": \":Users/mohsen-tum/Documents/Mendeley Desktop/Carbone et al/Analysis/Carbone et al. - 2009 - Mapping Kernel Objects to Enable Systematic Integrity Checking.pdf:pdf\", \"year\": \"2009\", \"ID\": \"Carbone2009\", \"pages\": \"555--565\"}',0),(42,'Castro2006','{\"isbn\": \"1-931971-47-1\", \"author\": \"Castro, Miguel and Costa, Manuel and Harris, Tim\", \"ENTRYTYPE\": \"article\", \"abstract\": \"Software attacks often subvert the intended data-flow in a vulnerable program. For example, attackers exploit buffer overflows and format string vulnerabilities to write data to unintended locations. We present a simple technique that prevents these attacks by enforcing data-flow integrity. It computes a data-flow graph using static analysis, and it instruments the program to ensure that the flow of data at runtime is allowed by the data-flow graph. We describe an efficient implementation of data-flow integrity enforcement that uses static analysis to reduce instrumentation overhead. This implementation can be used in practice to detect a broad class of attacks and errors because it can be applied automatically to C and C++ programs without modifications, it does not have false positives, and it has low overhead.\", \"title\": \"Securing software by enforcing data-flow integrity\", \"pages\": \"147--160\", \"mendeley-groups\": \"Tamperproofing/Methods\", \"link\": \"http://dl.acm.org/citation.cfm?id=1298455.1298470$\\\\backslash$nhttp://www.usenix.org/event/osdi06/tech/full{\\\\_}papers/castro/castro{\\\\_}html/\", \"file\": \":Users/mohsen-tum/Documents/Mendeley Desktop/Castro, Costa, Harris/Proceedings of the 7th symposium on Operating systems design and implementation/Castro, Costa, Harris - 2006 - Securing software by enforcing data-flow integrity.pdf:pdf\", \"year\": \"2006\", \"ID\": \"Castro2006\", \"journal\": \"Proceedings of the 7th symposium on Operating systems design and implementation\"}',0),(43,'gao2015integrity','{\"doi\": \"10.1109/ICAC.2015.34\", \"keyword\": \"Big Data;cloud computing;data integrity;data privacy;Big Data processing;cloud computing technology;dynamic redundancy computation;integrity protection solution;reputation based redundancy computation;Conferences;MapReduce;cloud computing;integrity protection\", \"title\": \"Integrity Protection for Big Data Processing with Dynamic Redundancy Computation\", \"booktitle\": \"Autonomic Computing (ICAC), 2015 IEEE International Conference on\", \"author\": \"Z. Gao and N. Desalvo and P. D. Khoa and S. H. Kim and L. Xu and W. W. Ro and R. M. Verma and W. Shi\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2015\", \"ID\": \"gao2015integrity\", \"pages\": \"159-160\"}',0),(44,'karapanos2016verena','{\"author\": \"Karapanos, Nikolaos and Filios, Alexandros and Popa, Raluca Ada and Capkun, Srdjan\", \"booktitle\": \"Proceedings of the 37th IEEE Symposium on Security and Privacy (IEEE S\\\\&P)\", \"title\": \"Verena: End-to-end integrity protection for web applications\", \"ID\": \"karapanos2016verena\", \"year\": \"2016\", \"ENTRYTYPE\": \"inproceedings\"}',0),(45,'Kil2009','{\"isbn\": \"9781424444212\", \"keyword\": \"dynamic attestation,integrity,remote attestation,runtime,system security,trusted computing\", \"author\": \"Kil, Chongkyung\", \"journal\": \"IEEE/IFIP International Conference on Dependable Systems {\\\\&} Networks\", \"title\": \"Remote Attestation to Dynamic System Properties: Towards Providing Complete System Integrity Evidence\", \"ENTRYTYPE\": \"article\", \"mendeley-groups\": \"Tamperproofing\", \"file\": \":Users/mohsen-tum/Documents/Mendeley Desktop/Kil/IEEEIFIP International Conference on Dependable Systems {\\\\&} Networks/Kil - 2009 - Remote Attestation to Dynamic System Properties Towards Providing Complete System Integrity Evidence.pdf:pdf\", \"year\": \"2009\", \"ID\": \"Kil2009\", \"pages\": \"115--124\"}',0),(46,'neisse2011implementing','{\"author\": \"Neisse, Ricardo and Holling, Dominik and Pretschner, Alexander\", \"booktitle\": \"Proceedings of the 2011 11th IEEE/ACM International Symposium on Cluster, Cloud and Grid Computing\", \"title\": \"Implementing trust in cloud infrastructures\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2011\", \"organization\": \"IEEE Computer Society\", \"ID\": \"neisse2011implementing\", \"pages\": \"524--533\"}',37),(47,'sun2015security','{\"author\": \"Sun, Yuqiong and Nanda, Susanta and Jaeger, Trent\", \"booktitle\": \"2015 IEEE 7th International Conference on Cloud Computing Technology and Science (CloudCom)\", \"title\": \"Security-as-a-Service for Microservices-Based Cloud Applications\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2015\", \"organization\": \"IEEE\", \"ID\": \"sun2015security\", \"pages\": \"50--57\"}',0),(48,'pappas2012smashing','{\"author\": \"Pappas, Vasilis and Polychronakis, Michalis and Keromytis, Angelos D\", \"booktitle\": \"2012 IEEE Symposium on Security and Privacy\", \"title\": \"Smashing the gadgets: Hindering return-oriented programming using in-place code randomization\", \"ENTRYTYPE\": \"inproceedings\", \"year\": \"2012\", \"organization\": \"IEEE\", \"ID\": \"pappas2012smashing\", \"pages\": \"601--615\"}',122),(49,'pappas2013transparent','{\"author\": \"Pappas, Vasilis and Polychronakis, Michalis and Keromytis, Angelos D\", \"booktitle\": \"Presented as part of the 22nd USENIX Security Symposium (USENIX Security 13)\", \"title\": \"Transparent ROP exploit mitigation using indirect branch tracing\", \"pages\": \"447--462\", \"year\": \"2013\", \"ID\": \"pappas2013transparent\", \"ENTRYTYPE\": \"inproceedings\"}',0);
/*!40000 ALTER TABLE `paper` ENABLE KEYS */;
UNLOCK TABLES;

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
  `id_taxonomy` int(11) unsigned NOT NULL AUTO_INCREMENT,
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
) ENGINE=InnoDB AUTO_INCREMENT=95 DEFAULT CHARSET=utf8;
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
) ENGINE=InnoDB AUTO_INCREMENT=248 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `taxonomy_relation`
--

LOCK TABLES `taxonomy_relation` WRITE;
/*!40000 ALTER TABLE `taxonomy_relation` DISABLE KEYS */;
INSERT INTO `taxonomy_relation` VALUES (56,1,2,93,3,1,'{\'x\':245.99999999999997,\'y\':71.99999999999996};{\'x\':246.99999999999997,\'y\':155}'),(57,1,3,93,3,1,'{\'x\':232.36470867292397,\'y\':155.35488918195506}'),(58,1,4,93,3,1,'{\'x\':246.99999999999997,\'y\':234};{\'x\':246.99999999999997,\'y\':155.99999999999994}'),(59,1,5,93,1,1,'{\'x\':99.99999999999991,\'y\':431.9999999999999}'),(60,1,21,20,3,1,'{\'x\':647.1794253666435,\'y\':69.82057517234267};{\'x\':650.6666891188631,\'y\':296.3333098646116}'),(61,1,22,20,3,1,'{\'x\':648.2100417611323,\'y\':143.79604123777884};{\'x\':649.4742480270277,\'y\':296.5104000328528}'),(62,1,23,20,3,1,'{\'x\':648.3866880746442,\'y\':211.6541379857174};{\'x\':649.3213879220136,\'y\':296.6063372240162}'),(63,1,24,20,3,1,'{\'x\':651.1102947476124,\'y\':366.7756595403927};{\'x\':650.602317217686,\'y\':297.7944639029661}'),(64,1,25,20,3,1,'{\'x\':651.0420523241444,\'y\':435.7642012738366};{\'x\':649.9017992903628,\'y\':297.80051601778246}'),(65,1,94,20,3,1,'{\'x\':631.5153552122035,\'y\':296.9184622110439}'),(66,1,6,21,1,1,'{\'x\':837.0000000000001,\'y\':168.00000000000003};{\'x\':835,\'y\':70.00000000000003}'),(67,1,6,22,1,1,'{\'x\':836.0000000000002,\'y\':169};{\'x\':836.0000000000002,\'y\':144}'),(68,1,6,94,1,1,'{\'x\':836.9999999999999,\'y\':169.0000000000001};{\'x\':837.9999999999997,\'y\':298.00000000000017}'),(69,1,6,5,5,1,'{\'x\':1019.9999999999998,\'y\':231.00000000000006};{\'x\':926,\'y\':231};{\'x\':929.9999999999997,\'y\':487};{\'x\':483.99999999999994,\'y\':482};{\'x\':480.9999999999998,\'y\':435}'),(70,1,7,5,5,1,'{\'x\':926.9999999999998,\'y\':319};{\'x\':929,\'y\':486};{\'x\':482.99999999999994,\'y\':482};{\'x\':482.0000000000001,\'y\':434.99999999999994}'),(71,1,8,5,5,1,'{\'x\':929,\'y\':514.9999999999999};{\'x\':929.0000000000002,\'y\':487};{\'x\':481.9999999999999,\'y\':481};{\'x\':480.9999999999999,\'y\':434}'),(72,1,9,7,3,1,'{\'x\':1474,\'y\':318}'),(73,1,10,7,3,1,'{\'x\':1473.7949380792272,\'y\':317.9784915175459}'),(74,1,13,8,3,1,'{\'x\':1259.2718899100228,\'y\':518.0659196311467}'),(75,1,11,8,3,1,'{\'x\':1361,\'y\':659};{\'x\':1360,\'y\':518.9999999999998}'),(76,1,12,11,3,1,'{\'x\':1549.9293606366093,\'y\':660.4788195267564}'),(77,1,7,24,1,1,'{\'x\':1036,\'y\':366.99999999999994}'),(78,1,7,25,1,1,'{\'x\':1036,\'y\':443.0000000000001}'),(79,1,8,24,1,1,'{\'x\':1035,\'y\':368.00000000000006}'),(80,1,8,25,1,1,'{\'x\':1035,\'y\':444};{\'x\':886.7188837054866,\'y\':441.4353930845009}'),(81,1,5,14,1,1,'{\'x\':328.1392281443622,\'y\':553.7668810710022}'),(82,1,15,14,3,1,'{\'x\':506.9999999999999,\'y\':603};{\'x\':506.99999999999994,\'y\':673.0000000000001}'),(83,1,95,14,3,1,'{\'x\':505.99999999999994,\'y\':525};{\'x\':506.99999999999994,\'y\':672.9999999999999}'),(84,1,17,14,3,1,'{\'x\':459.5797568920929,\'y\':673.5251347098841}'),(85,1,18,14,3,1,'{\'x\':507.9999999999999,\'y\':752};{\'x\':506.99999999999994,\'y\':674}'),(86,1,19,14,3,1,'{\'x\':509,\'y\':825};{\'x\':507,\'y\':673}'),(87,1,79,80,4,2,'{\'x\':119.21985912716738,\'y\':191.88190867508243};{\'x\':202.21985912716733,\'y\':190.88190867508223}'),(88,1,81,80,4,2,'{\'x\':241.21985912716747,\'y\':190.88190867508288};{\'x\':202.21985912716724,\'y\':190.881908675083}'),(89,1,82,80,4,2,'{\'x\':338.84847248584714,\'y\':191.5349713095123};{\'x\':203.21985912716718,\'y\':191.8819086750825}'),(90,1,83,80,4,2,'{\'x\':201.9121348988953,\'y\':310.9945330830352}'),(91,1,80,72,5,2,'{\'x\':303.5899151926502,\'y\':254.52141900969116}'),(92,1,73,72,5,2,'{\'x\':475.1869873177485,\'y\':311.500841992063}'),(93,1,84,72,5,2,'{\'x\':474.97515501915564,\'y\':327.65226755612076}'),(94,1,27,72,5,2,'{\'x\':749.7926083176478,\'y\':255.86123090994874}'),(95,1,28,27,3,2,'{\'x\':814.84152913225,\'y\':196.7518137564864};{\'x\':959.8415291322498,\'y\':196.75181375648648}'),(96,1,29,27,3,2,'{\'x\':959.5600319095079,\'y\':198.8499108543131}'),(97,1,30,27,3,2,'{\'x\':1103.8484724858467,\'y\':198.53497130951226};{\'x\':959.8415291322493,\'y\':197.7518137564864}'),(98,1,88,84,3,2,'{\'x\':1336.8484724858472,\'y\':396.53497130951206};{\'x\':728.0668472804629,\'y\':394.8950055421265}'),(99,1,31,27,3,2,'{\'x\':1307.8415291322503,\'y\':197.75181375648594};{\'x\':959.8415291322494,\'y\':198.75181375648614}'),(100,1,26,73,3,2,'{\'x\':474.9256826175398,\'y\':411.481394041331}'),(101,1,85,84,3,2,'{\'x\':728.34057509258,\'y\':391.10245018006844}'),(102,1,86,84,3,2,'{\'x\':919.8484724858474,\'y\':395.5349713095123};{\'x\':729.0743617269079,\'y\':394.50376596600955}'),(103,1,87,84,3,2,'{\'x\':1111.8484724858467,\'y\':396.534971309512};{\'x\':728.6180005479822,\'y\':393.9837999791198}'),(104,1,78,27,3,2,'{\'x\':1024.842000575281,\'y\':256.71427508485874}'),(105,1,57,32,5,3,'{\'x\':635.0486645311152,\'y\':425.21495617920175}'),(106,1,62,32,5,3,'{\'x\':847.9281806900183,\'y\':484.62176910908596}'),(107,1,67,32,5,3,'{\'x\':1293,\'y\':500.9999999999998};{\'x\':849,\'y\':501}'),(108,1,35,32,5,3,'{\'x\':557.9999999999999,\'y\':353};{\'x\':558,\'y\':424.99999999999983}'),(109,1,38,32,5,3,'{\'x\':425.9999999999999,\'y\':301.99999999999994};{\'x\':606,\'y\':303};{\'x\':847.9999999999999,\'y\':305}'),(110,1,41,32,5,3,'{\'x\':605,\'y\':302.0000000000002};{\'x\':848,\'y\':305.0000000000001}'),(111,1,44,32,5,3,'{\'x\':847.5010482948479,\'y\':333.5069859296838}'),(112,1,50,32,5,3,'{\'x\':1080,\'y\':304};{\'x\':848.0000000000002,\'y\':304.99999999999955}'),(113,1,33,32,3,3,'{\'x\':1012,\'y\':350};{\'x\':1013,\'y\':424}'),(114,1,34,32,3,3,'{\'x\':968.2652772623532,\'y\':424.7702775776905}'),(115,1,58,57,3,3,'{\'x\':233.99999999999994,\'y\':616};{\'x\':264.9999999999999,\'y\':616.0000000000001};{\'x\':415,\'y\':616.0000000000002}'),(116,1,59,57,3,3,'{\'x\':320.99999999999994,\'y\':616.0000000000001};{\'x\':414.99999999999994,\'y\':616.0000000000001}'),(117,1,60,57,3,3,'{\'x\':416.2585382026585,\'y\':548.8905200061671}'),(118,1,61,57,3,3,'{\'x\':486.00000000000006,\'y\':617.0000000000006};{\'x\':414.9999999999999,\'y\':617}'),(119,1,71,62,3,3,'{\'x\':728.9999999999997,\'y\':620.0000000000002};{\'x\':848.9999999999998,\'y\':621}'),(120,1,89,62,3,3,'{\'x\':848.7061746076743,\'y\':607.2489015044602}'),(121,1,63,62,3,3,'{\'x\':934,\'y\':622};{\'x\':848.0000000000006,\'y\':621.0000000000002}'),(122,1,64,62,3,3,'{\'x\':1005.0000000000001,\'y\':623};{\'x\':1005.9485620377977,\'y\':623.0121610517666};{\'x\':849.0000000000002,\'y\':621.0000000000003}'),(123,1,65,62,3,3,'{\'x\':1078,\'y\':624};{\'x\':849.0000000000005,\'y\':620.9999999999997}'),(124,1,70,67,3,3,'{\'x\':1228,\'y\':643.0000000000005};{\'x\':1297.9999999999993,\'y\':643.9999999999995}'),(125,1,69,67,3,3,'{\'x\':1369.0712099571522,\'y\':642.984697000612};{\'x\':1370.056723117731,\'y\':642.9706182411755};{\'x\':1298.000000000001,\'y\':643.9999999999995}'),(126,1,90,67,3,3,'{\'x\':1499,\'y\':643};{\'x\':1298.0000000000002,\'y\':644.0000000000002}'),(127,1,36,35,3,3,'{\'x\':138.99999999999997,\'y\':232};{\'x\':236,\'y\':231.99999999999997};{\'x\':236.00000000000003,\'y\':353}'),(128,1,37,35,3,3,'{\'x\':306.759866540906,\'y\':352.3421190031607}'),(129,1,39,38,3,3,'{\'x\':305.9999999999998,\'y\':178.9999999999999};{\'x\':427.00000000000006,\'y\':179}'),(130,1,92,38,3,3,'{\'x\':426.92869745792456,\'y\':180.49829836590638}'),(131,1,42,41,3,3,'{\'x\':604,\'y\':122.00000000000021}'),(132,1,43,41,3,3,'{\'x\':602.9999999999999,\'y\':122.0000000000002}'),(133,1,45,44,3,3,'{\'x\':795.6061590568141,\'y\':175.2083477327829};{\'x\':848.469182439591,\'y\':175.9429986885101}'),(134,1,46,44,3,3,'{\'x\':914,\'y\':175};{\'x\':847.9999999999998,\'y\':175.00000000000108}'),(135,1,47,44,3,3,'{\'x\':1064,\'y\':173.99999999999983};{\'x\':849.0794455437408,\'y\':176.13048301003045}'),(136,1,48,44,3,3,'{\'x\':1242.0000000000002,\'y\':171.99999999999994};{\'x\':849.0000000000002,\'y\':176.99999999999994}'),(137,1,49,44,3,3,'{\'x\':1417,\'y\':170.00000000000006};{\'x\':846.8473473305277,\'y\':177.10812131452283}'),(138,1,51,50,3,3,'{\'x\':1318.6167783100661,\'y\':262.1837002593493}'),(139,1,52,50,3,3,'{\'x\':1467,\'y\':340.9999999999999};{\'x\':1466.0000000000005,\'y\':264}'),(140,1,53,50,3,3,'{\'x\':1469.0000000000002,\'y\':430.99999999999994};{\'x\':1464.9999999999993,\'y\':263.99999999999983}'),(141,1,54,50,3,3,'{\'x\':1470.9999999999998,\'y\':521.0000000000001};{\'x\':1466,\'y\':263.9999999999998}'),(142,1,55,50,3,3,'{\'x\':1473,\'y\':598.9999999999998};{\'x\':1465.0000000000002,\'y\':264}'),(143,1,91,50,3,3,'{\'x\':1464,\'y\':186};{\'x\':1465.0000000000002,\'y\':264}'),(232,1,32,84,2,4,'{\'x\':-988.2359632913441,\'y\':-951.8103241540371};{\'x\':-319.4815098371267,\'y\':-953.0186736286499}'),(233,1,32,20,2,4,'{\'x\':-854.6013207506861,\'y\':-767.1148771605798}'),(234,1,32,27,2,4,'{\'x\':-988.3387121979307,\'y\':-671.6765206578543};{\'x\':-974.8746142326793,\'y\':-671.6633472599729};{\'x\':-881.6030760219256,\'y\':-671.6327409718069};{\'x\':-880.3387121979304,\'y\':-533.0098539911878};{\'x\':-323.00537886459705,\'y\':-530.3431873245208}'),(235,1,32,67,1,4,'{\'x\':-1112.6030760219257,\'y\':-766.6327409718037};{\'x\':-1109.4158110522417,\'y\':-280.0665701535794}'),(236,1,57,20,2,4,'{\'x\':-923.4018092467566,\'y\':-621.137974980062};{\'x\':-922.6030760219257,\'y\':-697.6327409718073};{\'x\':-779.605755831528,\'y\':-698.2517769899473};{\'x\':-691.6030760219255,\'y\':-697.6327409718069}'),(237,1,62,32,2,4,'{\'x\':-1064.0866850676534,\'y\':-384.47223943593764};{\'x\':-1063.6030760219257,\'y\':-766.6327409718072}'),(238,1,93,14,1,4,'{\'x\':-808.3387121979307,\'y\':-891.6765206622291};{\'x\':-808.6389747768841,\'y\':-281.22937748887347}'),(239,1,5,93,2,4,'{\'x\':-706.6030760219257,\'y\':-666.6327409718074};{\'x\':-545.6030760219257,\'y\':-665.6327409718065};{\'x\':-547.4698322288425,\'y\':-875.642814249965};{\'x\':-547.6030760219257,\'y\':-890.6327409718065}'),(240,1,84,32,2,4,'{\'x\':-319.94685346915963,\'y\':-837.0064609626666};{\'x\':-987.5115317947335,\'y\':-834.3519942169112}'),(241,1,80,72,2,4,'{\'x\':-181.6030760219259,\'y\':-623.6327409718075};{\'x\':-181.6030760219258,\'y\':-776.6327409718083}'),(242,1,27,5,2,4,'{\'x\':-369.47753487234206,\'y\':-482.1635275661221};{\'x\':-706.6030760219259,\'y\':-483.6327409718072}'),(243,1,27,32,2,4,'{\'x\':-403.0053788645975,\'y\':-482.3431873245204};{\'x\':-403.00537886459716,\'y\':-214.34318732452073};{\'x\':-1137.672045531264,\'y\':-214.34318732452084};{\'x\':-1137.6720455312643,\'y\':-767.6765206578542}'),(244,1,73,27,2,4,'{\'x\':-258.60307602192586,\'y\':-284.63274097180545};{\'x\':-258.60899150214834,\'y\':-483.8567189443867}'),(245,1,84,5,2,4,'{\'x\':-485.60307602192574,\'y\':-890.632740971807};{\'x\':-485.9305355073303,\'y\':-696.4536368161712};{\'x\':-485.09227442745885,\'y\':-608.4362234296373}'),(246,1,23,23,1,1,NULL),(247,1,23,6,1,1,'{\'x\':838.0000000000002,\'y\':211.00000000000006};{\'x\':836.9999999999999,\'y\':169.00000000000003}');
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
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user` (
  `email` varchar(128) NOT NULL,
  `password` varchar(128) NOT NULL,
  `admin` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`email`),
  UNIQUE KEY `email_UNIQUE` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES ('222@df.de','222',0),('felixhuber245@gmx.de','123',0),('felixhuber2465@gmx.de','123',0),('felixhuber2@gmx.de','123',1),('felixhuber4@gmx.de','123',0),('felixhuber5@gmx.de','123',0),('felixhuber62@gmx.de','123',0),('felixhuber6@gmx.de','123',0),('felixhuber7@gmx.de','123',0),('test3434@gmx.de','234',0),('test@gmx.de','123',0);
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping routines for database 'classification'
--
/*!50003 DROP PROCEDURE IF EXISTS `insertallchildrenperattribute` */;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8 */ ;
/*!50003 SET character_set_results = utf8 */ ;
/*!50003 SET collation_connection  = utf8_general_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'STRICT_TRANS_TABLES,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
CREATE DEFINER=`root`@`localhost` PROCEDURE `insertallchildrenperattribute`(taxonomyId INT(11))
BEGIN
  DECLARE cursor_id_attribute INT(11);
  DECLARE cursor_text VARCHAR(50);
  DECLARE done INT DEFAULT FALSE;
  DECLARE cursor_i CURSOR FOR SELECT id_attribute, text FROM attribute;
  DECLARE CONTINUE HANDLER FOR NOT FOUND SET done = TRUE;
  DELETE FROM allChildrenPerAttribute WHERE id_taxonomy = taxonomyId;
  OPEN cursor_i;
  read_loop: LOOP
    FETCH cursor_i INTO cursor_id_attribute, cursor_text;
    IF done THEN
      LEAVE read_loop;
    END IF;
    INSERT INTO allChildrenPerAttribute(id_attribute, id_taxonomy, text, children) VALUES(cursor_id_attribute, taxonomyId, cursor_text, (SELECT (CASE WHEN b.children IS NULL THEN CAST(cursor_id_attribute AS CHAR(50)) ELSE CONCAT(CAST(cursor_id_attribute AS CHAR(50)), ",", b.children) END) AS children FROM (SELECT GROUP_CONCAT(lv SEPARATOR ',') AS children FROM (SELECT @pv:=(SELECT GROUP_CONCAT(DISTINCT relation1.id_src_attribute SEPARATOR ',') FROM taxonomy_relation AS relation1 WHERE relation1.id_taxonomy = taxonomyId AND relation1.id_relation > 2 AND FIND_IN_SET(relation1.id_dest_attribute, @pv)) AS lv FROM taxonomy_relation AS relation2 JOIN (SELECT @pv:=cursor_id_attribute) tmp ON (relation2.id_taxonomy = taxonomyId)) a) b));
  END LOOP;
  CLOSE cursor_i;
END ;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;
/*!50003 DROP PROCEDURE IF EXISTS `insertallparentsperattribute` */;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8 */ ;
/*!50003 SET character_set_results = utf8 */ ;
/*!50003 SET collation_connection  = utf8_general_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'STRICT_TRANS_TABLES,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
CREATE DEFINER=`root`@`localhost` PROCEDURE `insertallparentsperattribute`(taxonomyId INT(11))
BEGIN
  DECLARE cursor_id_attribute INT(11);
  DECLARE cursor_text VARCHAR(50);
  DECLARE done INT DEFAULT FALSE;
  DECLARE cursor_i CURSOR FOR SELECT id_attribute, text FROM attribute;
  DECLARE CONTINUE HANDLER FOR NOT FOUND SET done = TRUE;
  DELETE FROM allParentsPerAttribute WHERE id_taxonomy = taxonomyId;
  OPEN cursor_i;
  read_loop: LOOP
    FETCH cursor_i INTO cursor_id_attribute, cursor_text;
    IF done THEN
      LEAVE read_loop;
    END IF;
    INSERT INTO allParentsPerAttribute(id_attribute, id_taxonomy, text, parents) VALUES(cursor_id_attribute, taxonomyId, cursor_text, (SELECT (CASE WHEN b.parents IS NULL THEN "" ELSE b.parents END) AS parents FROM (SELECT GROUP_CONCAT(lv SEPARATOR ',') AS parents FROM (SELECT @pv:=(SELECT GROUP_CONCAT(DISTINCT parent.text SEPARATOR ',') FROM taxonomy_relation AS relation1 INNER JOIN attribute as parent ON (relation1.id_dest_attribute = parent.id_attribute AND parent.id_taxonomy = taxonomyId) WHERE relation1.id_taxonomy = 1 AND relation1.id_relation > 2 AND FIND_IN_SET((SELECT DISTINCT text FROM attribute WHERE id_attribute = relation1.id_src_attribute AND id_taxonomy = taxonomyId), @pv)) AS lv FROM taxonomy_relation JOIN (SELECT @pv:=text FROM attribute WHERE id_attribute = cursor_id_attribute) tmp) a) b));
  END LOOP;
  CLOSE cursor_i;
END ;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;

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
/*!50001 VIEW `paper_merged_attributes` AS select 1 AS `id_taxonomy`,1 AS `id_paper`,1 AS `citation`,1 AS `bib`,1 AS `atts`,1 AS `leaf_atts` */;
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

-- Dump completed on 2018-01-21 20:02:41
