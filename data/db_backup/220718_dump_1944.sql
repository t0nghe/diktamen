-- MySQL dump 10.13  Distrib 8.0.22, for osx10.15 (x86_64)
--
-- Host: localhost    Database: diktdev
-- ------------------------------------------------------
-- Server version	8.0.29

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `article`
--

DROP TABLE IF EXISTS `article`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `article` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `title` varchar(256) DEFAULT NULL,
  `sent_count` int DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `article`
--

LOCK TABLES `article` WRITE;
/*!40000 ALTER TABLE `article` DISABLE KEYS */;
INSERT INTO `article` VALUES (5,'Vad gör livet meningsfullt? ',14),(6,'Zelenskyj är en man som kvinnor längtar efter',17);
/*!40000 ALTER TABLE `article` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `dummy_message`
--

DROP TABLE IF EXISTS `dummy_message`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `dummy_message` (
  `number` int DEFAULT NULL,
  `success` tinyint(1) DEFAULT NULL,
  `message` varchar(128) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `dummy_message`
--

LOCK TABLES `dummy_message` WRITE;
/*!40000 ALTER TABLE `dummy_message` DISABLE KEYS */;
INSERT INTO `dummy_message` VALUES (81,1,'Hello, soluppgång'),(87,0,'But what do you mean by that?'),(NULL,1,'hello');
/*!40000 ALTER TABLE `dummy_message` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sent`
--

DROP TABLE IF EXISTS `sent`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sent` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `text` varchar(400) DEFAULT NULL,
  `media_uri` varchar(256) DEFAULT NULL,
  `article_id` bigint DEFAULT NULL,
  `index_in_article` int DEFAULT NULL,
  `word_count` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `article_id` (`article_id`),
  CONSTRAINT `sent_ibfk_1` FOREIGN KEY (`article_id`) REFERENCES `article` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=32 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sent`
--

LOCK TABLES `sent` WRITE;
/*!40000 ALTER TABLE `sent` DISABLE KEYS */;
INSERT INTO `sent` VALUES (1,'Familjen ger livet dess värde.','http://d1zg52ope8o24c.cloudfront.net/meningsfullt/meningsfullt_01.mp3',5,1,6),(2,'Och inom familjen värderar vi barnen långt högre än en eventuell partner.','http://d1zg52ope8o24c.cloudfront.net/meningsfullt/meningsfullt_02.mp3',5,2,13),(3,'Men vi uppskattar familj och partner olika beroende på faktorer som exempelvis ålder, kön och var i världen vi bor. ','http://d1zg52ope8o24c.cloudfront.net/meningsfullt/meningsfullt_03.mp3',5,3,22),(4,'Den framgår av en ny internationell studie av befolkningarna i 17 relativt rika länder.','http://d1zg52ope8o24c.cloudfront.net/meningsfullt/meningsfullt_04.mp3',5,4,15),(5,'Familj, men även arbete, välstånd och hälsa. Det är vad som ger oss mening i livet ...','http://d1zg52ope8o24c.cloudfront.net/meningsfullt/meningsfullt_05.mp3',5,5,20),(6,'Svaren kanske inte förvånar, men frågan har nu analyserats i en stor enkätundersökning av Pew Research Center i USA.','http://d1zg52ope8o24c.cloudfront.net/meningsfullt/meningsfullt_06.mp3',5,6,21),(7,'Under våren 2021 intervjuades nästan 19 000 vuxna personer i 17 relativt rika länder i fyra världsdelar.','http://d1zg52ope8o24c.cloudfront.net/meningsfullt/meningsfullt_07.mp3',5,7,18),(8,'Sverige var ett av länderna i studien.','http://d1zg52ope8o24c.cloudfront.net/meningsfullt/meningsfullt_08.mp3',5,8,8),(9,'Huvudfrågan var öppen, och löd:','http://d1zg52ope8o24c.cloudfront.net/meningsfullt/meningsfullt_09.mp3',5,9,7),(10,'Vi är intresserade av att undersöka vad det betyder att leva ett meningsfullt liv.','http://d1zg52ope8o24c.cloudfront.net/meningsfullt/meningsfullt_10.mp3',5,10,15),(11,'Vilka aspekter av ditt nuvarande liv finner du meningsfulla och tillfredsställande?','http://d1zg52ope8o24c.cloudfront.net/meningsfullt/meningsfullt_11.mp3',5,11,12),(12,'Det återkommande och vanligaste svaret var familjen.','http://d1zg52ope8o24c.cloudfront.net/meningsfullt/meningsfullt_12.mp3',5,12,8),(13,'I 14 av de 17 länderna kom familjen allra först.','http://d1zg52ope8o24c.cloudfront.net/meningsfullt/meningsfullt_13.mp3',5,13,11),(14,'Men i de två östasiatiska länderna Sydkorea och Taiwan placerades familjen först på tredje plats.','http://d1zg52ope8o24c.cloudfront.net/meningsfullt/meningsfullt_14.mp3',5,14,16),(15,'Det är 75 mil mellan mitt skrivbord i Malmö och närmaste gränsstad till Ukraina, Medyka, en liten polsk by med några tusen invånare.','http://d1zg52ope8o24c.cloudfront.net/zelenskyj/zelenskyj_01.mp3',6,1,26),(16,'På andra sidan gränsen pågår ett krig, inte ett sådant som utspelar sig i spalter och debattprogram, utan ett sådant som dödar människor och utplånar städer. ','http://d1zg52ope8o24c.cloudfront.net/zelenskyj/zelenskyj_02.mp3',6,2,29),(17,'Mot bakgrund av detta ter sig samtalet i den svenska offentligheten som vi har vant oss vid det som bisarrt.','http://d1zg52ope8o24c.cloudfront.net/zelenskyj/zelenskyj_03.mp3',6,3,21),(18,'I åratal har ett oförsonligt kulturkrig rasat, med offer till både höger och vänster,','http://d1zg52ope8o24c.cloudfront.net/zelenskyj/zelenskyj_04.mp3',6,4,16),(19,'båda sidor lika övertygade om sin förträfflighet och om att just deras världsbild är korrekt.','http://d1zg52ope8o24c.cloudfront.net/zelenskyj/zelenskyj_05.mp3',6,5,16),(20,'Det har framställts  som en kamp mellan mörker och ljus, och människor har, ofta på tämligen dunkla grunder, dubbats till goda eller dömts som onda.','http://d1zg52ope8o24c.cloudfront.net/zelenskyj/zelenskyj_06.mp3',6,6,30),(21,'Hjältar, martyrer och profeter – som Greta Thunberg och Cissi Wallin, Katerina Janouch, Hanif Bali och Joakim Lamotte, allt efter tycke och smak – har kommit och gått.','http://d1zg52ope8o24c.cloudfront.net/zelenskyj/zelenskyj_07.mp3',6,7,33),(22,'Varenda människa som någon gång publicerat en åsikt har nogsamt sorterats in, oavsett om man vill vara en del av detta fiktiva krig eller inte.','http://d1zg52ope8o24c.cloudfront.net/zelenskyj/zelenskyj_08.mp3',6,8,27),(23,'Diskussioner om vilka ord som är tillåtna och otillåtna, uppförandekoder för standup-branschen,','http://d1zg52ope8o24c.cloudfront.net/zelenskyj/zelenskyj_09.mp3',6,9,14),(24,'vem som är ond respektive god och vilka pronomen som är lämpliga ter sig nu som väldigt obsoleta. ','http://d1zg52ope8o24c.cloudfront.net/zelenskyj/zelenskyj_10.mp3',6,10,19),(25,'Världen är inte bara en berättelse. Den är.','http://d1zg52ope8o24c.cloudfront.net/zelenskyj/zelenskyj_11.mp3',6,11,10),(26,'Människan är inte blott en konstruktion. Hon är en kropp som känner, lider och dör.','http://d1zg52ope8o24c.cloudfront.net/zelenskyj/zelenskyj_12.mp3',6,12,18),(27,'Krig förs inte bara på Twitter, utan på gator och torg, vapen är inte bara ord, utan gevär och bomber.','http://d1zg52ope8o24c.cloudfront.net/zelenskyj/zelenskyj_13.mp3',6,13,24),(28,'Trots dessa uppenbara påpekanden ser jag hur det pågående kriget i Ukraina fiktionaliseras.','http://d1zg52ope8o24c.cloudfront.net/zelenskyj/zelenskyj_14.mp3',6,14,14),(29,'Vi har två huvudpersoner, den hjältemodige president Zelenskyj och den galne diktatorn Vladimir Putin.','http://d1zg52ope8o24c.cloudfront.net/zelenskyj/zelenskyj_15.mp3',6,15,16),(30,'Sociala medier fylls av hyllningar till Zelenskyj och nidbilder av Putin,','http://d1zg52ope8o24c.cloudfront.net/zelenskyj/zelenskyj_16.mp3',6,16,12),(31,'reportage om den osannolike presidenten, från komiker till statschef skrivs i rasande fart, liksom analyser som avfärdar Putin som ”sjuk”, ”vansinnig” och ”galen”.','http://d1zg52ope8o24c.cloudfront.net/zelenskyj/zelenskyj_17.mp3',6,17,33);
/*!40000 ALTER TABLE `sent` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(32) NOT NULL,
  `email` varchar(32) DEFAULT NULL,
  `passwordHash` binary(60) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_article`
--

DROP TABLE IF EXISTS `user_article`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_article` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` int DEFAULT NULL,
  `article_id` bigint DEFAULT NULL,
  `finished_sent_index` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  KEY `article_id` (`article_id`),
  CONSTRAINT `user_article_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`),
  CONSTRAINT `user_article_ibfk_2` FOREIGN KEY (`article_id`) REFERENCES `article` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_article`
--

LOCK TABLES `user_article` WRITE;
/*!40000 ALTER TABLE `user_article` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_article` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_dictation`
--

DROP TABLE IF EXISTS `user_dictation`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_dictation` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` int DEFAULT NULL,
  `sent_id` bigint DEFAULT NULL,
  `repetition_count` int DEFAULT NULL,
  `repetition_time` datetime DEFAULT NULL,
  `repetition_score` float DEFAULT NULL,
  `next_repetition_time` datetime DEFAULT NULL,
  `perfection_count` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  KEY `sent_id` (`sent_id`),
  CONSTRAINT `user_dictation_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`),
  CONSTRAINT `user_dictation_ibfk_2` FOREIGN KEY (`sent_id`) REFERENCES `sent` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_dictation`
--

LOCK TABLES `user_dictation` WRITE;
/*!40000 ALTER TABLE `user_dictation` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_dictation` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_dictation_word`
--

DROP TABLE IF EXISTS `user_dictation_word`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_dictation_word` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` int DEFAULT NULL,
  `user_dictation_id` bigint DEFAULT NULL,
  `sent_id` bigint DEFAULT NULL,
  `word_id` bigint DEFAULT NULL,
  `index_in_sent` int DEFAULT NULL,
  `user_input_wordform` varchar(50) DEFAULT NULL,
  `user_input_score` float DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  KEY `user_dictation_id` (`user_dictation_id`),
  KEY `sent_id` (`sent_id`),
  CONSTRAINT `user_dictation_word_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`),
  CONSTRAINT `user_dictation_word_ibfk_2` FOREIGN KEY (`user_dictation_id`) REFERENCES `user_dictation` (`id`),
  CONSTRAINT `user_dictation_word_ibfk_3` FOREIGN KEY (`sent_id`) REFERENCES `sent` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_dictation_word`
--

LOCK TABLES `user_dictation_word` WRITE;
/*!40000 ALTER TABLE `user_dictation_word` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_dictation_word` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `word`
--

DROP TABLE IF EXISTS `word`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `word` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `wordform` varchar(50) DEFAULT NULL,
  `length` int DEFAULT NULL,
  `sent_id` bigint DEFAULT NULL,
  `index_in_sent` int DEFAULT NULL,
  `is_cloze` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `sent_id` (`sent_id`),
  CONSTRAINT `word_ibfk_1` FOREIGN KEY (`sent_id`) REFERENCES `sent` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=551 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `word`
--

LOCK TABLES `word` WRITE;
/*!40000 ALTER TABLE `word` DISABLE KEYS */;
INSERT INTO `word` VALUES (1,'Familjen',8,1,1,1),(2,'ger',3,1,2,1),(3,'livet',5,1,3,1),(4,'dess',4,1,4,1),(5,'värde',5,1,5,1),(6,'.',1,1,6,0),(7,'Och',3,2,1,1),(8,'inom',4,2,2,1),(9,'familjen',8,2,3,1),(10,'värderar',8,2,4,1),(11,'vi',2,2,5,1),(12,'barnen',6,2,6,1),(13,'långt',5,2,7,1),(14,'högre',5,2,8,1),(15,'än',2,2,9,1),(16,'en',2,2,10,1),(17,'eventuell',9,2,11,1),(18,'partner',7,2,12,1),(19,'.',1,2,13,0),(20,'Men',3,3,1,1),(21,'vi',2,3,2,1),(22,'uppskattar',10,3,3,1),(23,'familj',6,3,4,1),(24,'och',3,3,5,1),(25,'partner',7,3,6,1),(26,'olika',5,3,7,1),(27,'beroende',8,3,8,1),(28,'på',2,3,9,1),(29,'faktorer',8,3,10,1),(30,'som',3,3,11,1),(31,'exempelvis',10,3,12,1),(32,'ålder',5,3,13,1),(33,',',1,3,14,0),(34,'kön',3,3,15,1),(35,'och',3,3,16,1),(36,'var',3,3,17,1),(37,'i',1,3,18,1),(38,'världen',7,3,19,1),(39,'vi',2,3,20,1),(40,'bor',3,3,21,1),(41,'.',1,3,22,0),(42,'Den',3,4,1,1),(43,'framgår',7,4,2,1),(44,'av',2,4,3,1),(45,'en',2,4,4,1),(46,'ny',2,4,5,1),(47,'internationell',14,4,6,1),(48,'studie',6,4,7,1),(49,'av',2,4,8,1),(50,'befolkningarna',14,4,9,1),(51,'i',1,4,10,1),(52,'17',2,4,11,1),(53,'relativt',8,4,12,1),(54,'rika',4,4,13,1),(55,'länder',6,4,14,1),(56,'.',1,4,15,0),(57,'Familj',6,5,1,1),(58,',',1,5,2,0),(59,'men',3,5,3,1),(60,'även',4,5,4,1),(61,'arbete',6,5,5,1),(62,',',1,5,6,0),(63,'välstånd',8,5,7,1),(64,'och',3,5,8,1),(65,'hälsa',5,5,9,1),(66,'.',1,5,10,0),(67,'Det',3,5,11,1),(68,'är',2,5,12,1),(69,'vad',3,5,13,1),(70,'som',3,5,14,1),(71,'ger',3,5,15,1),(72,'oss',3,5,16,1),(73,'mening',6,5,17,1),(74,'i',1,5,18,1),(75,'livet',5,5,19,1),(76,'...',3,5,20,0),(77,'Svaren',6,6,1,1),(78,'kanske',6,6,2,1),(79,'inte',4,6,3,1),(80,'förvånar',8,6,4,1),(81,',',1,6,5,0),(82,'men',3,6,6,1),(83,'frågan',6,6,7,1),(84,'har',3,6,8,1),(85,'nu',2,6,9,1),(86,'analyserats',11,6,10,1),(87,'i',1,6,11,1),(88,'en',2,6,12,1),(89,'stor',4,6,13,1),(90,'enkätundersökning',17,6,14,1),(91,'av',2,6,15,1),(92,'Pew',3,6,16,0),(93,'Research',8,6,17,0),(94,'Center',6,6,18,0),(95,'i',1,6,19,1),(96,'USA',3,6,20,0),(97,'.',1,6,21,0),(98,'Under',5,7,1,1),(99,'våren',5,7,2,1),(100,'2021',4,7,3,1),(101,'intervjuades',12,7,4,1),(102,'nästan',6,7,5,1),(103,'19',2,7,6,1),(104,'000',3,7,7,1),(105,'vuxna',5,7,8,1),(106,'personer',8,7,9,1),(107,'i',1,7,10,1),(108,'17',2,7,11,1),(109,'relativt',8,7,12,1),(110,'rika',4,7,13,1),(111,'länder',6,7,14,1),(112,'i',1,7,15,1),(113,'fyra',4,7,16,1),(114,'världsdelar',11,7,17,1),(115,'.',1,7,18,0),(116,'Sverige',7,8,1,0),(117,'var',3,8,2,1),(118,'ett',3,8,3,1),(119,'av',2,8,4,1),(120,'länderna',8,8,5,1),(121,'i',1,8,6,1),(122,'studien',7,8,7,1),(123,'.',1,8,8,0),(124,'Huvudfrågan',11,9,1,1),(125,'var',3,9,2,1),(126,'öppen',5,9,3,1),(127,',',1,9,4,0),(128,'och',3,9,5,1),(129,'löd',3,9,6,1),(130,':',1,9,7,0),(131,'Vi',2,10,1,1),(132,'är',2,10,2,1),(133,'intresserade',12,10,3,1),(134,'av',2,10,4,1),(135,'att',3,10,5,1),(136,'undersöka',9,10,6,1),(137,'vad',3,10,7,1),(138,'det',3,10,8,1),(139,'betyder',7,10,9,1),(140,'att',3,10,10,1),(141,'leva',4,10,11,1),(142,'ett',3,10,12,1),(143,'meningsfullt',12,10,13,1),(144,'liv',3,10,14,1),(145,'.',1,10,15,0),(146,'Vilka',5,11,1,1),(147,'aspekter',8,11,2,1),(148,'av',2,11,3,1),(149,'ditt',4,11,4,1),(150,'nuvarande',9,11,5,1),(151,'liv',3,11,6,1),(152,'finner',6,11,7,1),(153,'du',2,11,8,1),(154,'meningsfulla',12,11,9,1),(155,'och',3,11,10,1),(156,'tillfredsställande',18,11,11,1),(157,'?',1,11,12,0),(158,'Det',3,12,1,1),(159,'återkommande',12,12,2,1),(160,'och',3,12,3,1),(161,'vanligaste',10,12,4,1),(162,'svaret',6,12,5,1),(163,'var',3,12,6,1),(164,'familjen',8,12,7,1),(165,'.',1,12,8,0),(166,'I',1,13,1,1),(167,'14',2,13,2,1),(168,'av',2,13,3,1),(169,'de',2,13,4,1),(170,'17',2,13,5,1),(171,'länderna',8,13,6,1),(172,'kom',3,13,7,1),(173,'familjen',8,13,8,1),(174,'allra',5,13,9,1),(175,'först',5,13,10,1),(176,'.',1,13,11,0),(177,'Men',3,14,1,1),(178,'i',1,14,2,1),(179,'de',2,14,3,1),(180,'två',3,14,4,1),(181,'östasiatiska',12,14,5,1),(182,'länderna',8,14,6,1),(183,'Sydkorea',8,14,7,0),(184,'och',3,14,8,1),(185,'Taiwan',6,14,9,0),(186,'placerades',10,14,10,1),(187,'familjen',8,14,11,1),(188,'först',5,14,12,1),(189,'på',2,14,13,1),(190,'tredje',6,14,14,1),(191,'plats',5,14,15,1),(192,'.',1,14,16,0),(193,'Det',3,15,1,1),(194,'är',2,15,2,1),(195,'75',2,15,3,1),(196,'mil',3,15,4,1),(197,'mellan',6,15,5,1),(198,'mitt',4,15,6,1),(199,'skrivbord',9,15,7,1),(200,'i',1,15,8,1),(201,'Malmö',5,15,9,0),(202,'och',3,15,10,1),(203,'närmaste',8,15,11,1),(204,'gränsstad',9,15,12,1),(205,'till',4,15,13,1),(206,'Ukraina',7,15,14,0),(207,',',1,15,15,0),(208,'Medyka',6,15,16,1),(209,',',1,15,17,0),(210,'en',2,15,18,1),(211,'liten',5,15,19,1),(212,'polsk',5,15,20,1),(213,'by',2,15,21,1),(214,'med',3,15,22,1),(215,'några',5,15,23,1),(216,'tusen',5,15,24,1),(217,'invånare',8,15,25,1),(218,'.',1,15,26,0),(219,'På',2,16,1,1),(220,'andra',5,16,2,1),(221,'sidan',5,16,3,1),(222,'gränsen',7,16,4,1),(223,'pågår',5,16,5,1),(224,'ett',3,16,6,1),(225,'krig',4,16,7,1),(226,',',1,16,8,0),(227,'inte',4,16,9,1),(228,'ett',3,16,10,1),(229,'sådant',6,16,11,1),(230,'som',3,16,12,1),(231,'utspelar',8,16,13,1),(232,'sig',3,16,14,1),(233,'i',1,16,15,1),(234,'spalter',7,16,16,1),(235,'och',3,16,17,1),(236,'debattprogram',13,16,18,1),(237,',',1,16,19,0),(238,'utan',4,16,20,1),(239,'ett',3,16,21,1),(240,'sådant',6,16,22,1),(241,'som',3,16,23,1),(242,'dödar',5,16,24,1),(243,'människor',9,16,25,1),(244,'och',3,16,26,1),(245,'utplånar',8,16,27,1),(246,'städer',6,16,28,1),(247,'.',1,16,29,0),(248,'Mot',3,17,1,1),(249,'bakgrund',8,17,2,1),(250,'av',2,17,3,1),(251,'detta',5,17,4,1),(252,'ter',3,17,5,1),(253,'sig',3,17,6,1),(254,'samtalet',8,17,7,1),(255,'i',1,17,8,1),(256,'den',3,17,9,1),(257,'svenska',7,17,10,1),(258,'offentligheten',14,17,11,1),(259,'som',3,17,12,1),(260,'vi',2,17,13,1),(261,'har',3,17,14,1),(262,'vant',4,17,15,1),(263,'oss',3,17,16,1),(264,'vid',3,17,17,1),(265,'det',3,17,18,1),(266,'som',3,17,19,1),(267,'bisarrt',7,17,20,1),(268,'.',1,17,21,0),(269,'I',1,18,1,1),(270,'åratal',6,18,2,1),(271,'har',3,18,3,1),(272,'ett',3,18,4,1),(273,'oförsonligt',11,18,5,1),(274,'kulturkrig',10,18,6,1),(275,'rasat',5,18,7,1),(276,',',1,18,8,0),(277,'med',3,18,9,1),(278,'offer',5,18,10,1),(279,'till',4,18,11,1),(280,'både',4,18,12,1),(281,'höger',5,18,13,1),(282,'och',3,18,14,1),(283,'vänster',7,18,15,1),(284,',',1,18,16,0),(285,'båda',4,19,1,1),(286,'sidor',5,19,2,1),(287,'lika',4,19,3,1),(288,'övertygade',10,19,4,1),(289,'om',2,19,5,1),(290,'sin',3,19,6,1),(291,'förträfflighet',14,19,7,1),(292,'och',3,19,8,1),(293,'om',2,19,9,1),(294,'att',3,19,10,1),(295,'just',4,19,11,1),(296,'deras',5,19,12,1),(297,'världsbild',10,19,13,1),(298,'är',2,19,14,1),(299,'korrekt',7,19,15,1),(300,'.',1,19,16,0),(301,'Det',3,20,1,1),(302,'har',3,20,2,1),(303,'framställts',11,20,3,1),(304,' ',1,20,4,1),(305,'som',3,20,5,1),(306,'en',2,20,6,1),(307,'kamp',4,20,7,1),(308,'mellan',6,20,8,1),(309,'mörker',6,20,9,1),(310,'och',3,20,10,1),(311,'ljus',4,20,11,1),(312,',',1,20,12,0),(313,'och',3,20,13,1),(314,'människor',9,20,14,1),(315,'har',3,20,15,1),(316,',',1,20,16,0),(317,'ofta',4,20,17,1),(318,'på',2,20,18,1),(319,'tämligen',8,20,19,1),(320,'dunkla',6,20,20,1),(321,'grunder',7,20,21,1),(322,',',1,20,22,0),(323,'dubbats',7,20,23,1),(324,'till',4,20,24,1),(325,'goda',4,20,25,1),(326,'eller',5,20,26,1),(327,'dömts',5,20,27,1),(328,'som',3,20,28,1),(329,'onda',4,20,29,1),(330,'.',1,20,30,0),(331,'Hjältar',7,21,1,1),(332,',',1,21,2,0),(333,'martyrer',8,21,3,1),(334,'och',3,21,4,1),(335,'profeter',8,21,5,1),(336,'–',1,21,6,0),(337,'som',3,21,7,1),(338,'Greta',5,21,8,0),(339,'Thunberg',8,21,9,0),(340,'och',3,21,10,1),(341,'Cissi',5,21,11,0),(342,'Wallin',6,21,12,0),(343,',',1,21,13,0),(344,'Katerina',8,21,14,0),(345,'Janouch',7,21,15,0),(346,',',1,21,16,0),(347,'Hanif',5,21,17,0),(348,'Bali',4,21,18,0),(349,'och',3,21,19,1),(350,'Joakim',6,21,20,0),(351,'Lamotte',7,21,21,0),(352,',',1,21,22,0),(353,'allt',4,21,23,1),(354,'efter',5,21,24,1),(355,'tycke',5,21,25,1),(356,'och',3,21,26,1),(357,'smak',4,21,27,1),(358,'–',1,21,28,0),(359,'har',3,21,29,1),(360,'kommit',6,21,30,1),(361,'och',3,21,31,1),(362,'gått',4,21,32,1),(363,'.',1,21,33,0),(364,'Varenda',7,22,1,1),(365,'människa',8,22,2,1),(366,'som',3,22,3,1),(367,'någon',5,22,4,1),(368,'gång',4,22,5,1),(369,'publicerat',10,22,6,1),(370,'en',2,22,7,1),(371,'åsikt',5,22,8,1),(372,'har',3,22,9,1),(373,'nogsamt',7,22,10,1),(374,'sorterats',9,22,11,1),(375,'in',2,22,12,1),(376,',',1,22,13,0),(377,'oavsett',7,22,14,1),(378,'om',2,22,15,1),(379,'man',3,22,16,1),(380,'vill',4,22,17,1),(381,'vara',4,22,18,1),(382,'en',2,22,19,1),(383,'del',3,22,20,1),(384,'av',2,22,21,1),(385,'detta',5,22,22,1),(386,'fiktiva',7,22,23,1),(387,'krig',4,22,24,1),(388,'eller',5,22,25,1),(389,'inte',4,22,26,1),(390,'.',1,22,27,0),(391,'Diskussioner',12,23,1,1),(392,'om',2,23,2,1),(393,'vilka',5,23,3,1),(394,'ord',3,23,4,1),(395,'som',3,23,5,1),(396,'är',2,23,6,1),(397,'tillåtna',8,23,7,1),(398,'och',3,23,8,1),(399,'otillåtna',9,23,9,1),(400,',',1,23,10,0),(401,'uppförandekoder',15,23,11,1),(402,'för',3,23,12,1),(403,'standup-branschen',17,23,13,1),(404,',',1,23,14,0),(405,'vem',3,24,1,1),(406,'som',3,24,2,1),(407,'är',2,24,3,1),(408,'ond',3,24,4,1),(409,'respektive',10,24,5,1),(410,'god',3,24,6,1),(411,'och',3,24,7,1),(412,'vilka',5,24,8,1),(413,'pronomen',8,24,9,1),(414,'som',3,24,10,1),(415,'är',2,24,11,1),(416,'lämpliga',8,24,12,1),(417,'ter',3,24,13,1),(418,'sig',3,24,14,1),(419,'nu',2,24,15,1),(420,'som',3,24,16,1),(421,'väldigt',7,24,17,1),(422,'obsoleta',8,24,18,1),(423,'.',1,24,19,0),(424,'Världen',7,25,1,1),(425,'är',2,25,2,1),(426,'inte',4,25,3,1),(427,'bara',4,25,4,1),(428,'en',2,25,5,1),(429,'berättelse',10,25,6,1),(430,'.',1,25,7,0),(431,'Den',3,25,8,1),(432,'är',2,25,9,1),(433,'.',1,25,10,0),(434,'Människan',9,26,1,1),(435,'är',2,26,2,1),(436,'inte',4,26,3,1),(437,'blott',5,26,4,1),(438,'en',2,26,5,1),(439,'konstruktion',12,26,6,1),(440,'.',1,26,7,0),(441,'Hon',3,26,8,1),(442,'är',2,26,9,1),(443,'en',2,26,10,1),(444,'kropp',5,26,11,1),(445,'som',3,26,12,1),(446,'känner',6,26,13,1),(447,',',1,26,14,0),(448,'lider',5,26,15,1),(449,'och',3,26,16,1),(450,'dör',3,26,17,1),(451,'.',1,26,18,0),(452,'Krig',4,27,1,1),(453,'förs',4,27,2,1),(454,'inte',4,27,3,1),(455,'bara',4,27,4,1),(456,'på',2,27,5,1),(457,'Twitter',7,27,6,1),(458,',',1,27,7,0),(459,'utan',4,27,8,1),(460,'på',2,27,9,1),(461,'gator',5,27,10,1),(462,'och',3,27,11,1),(463,'torg',4,27,12,1),(464,',',1,27,13,0),(465,'vapen',5,27,14,1),(466,'är',2,27,15,1),(467,'inte',4,27,16,1),(468,'bara',4,27,17,1),(469,'ord',3,27,18,1),(470,',',1,27,19,0),(471,'utan',4,27,20,1),(472,'gevär',5,27,21,1),(473,'och',3,27,22,1),(474,'bomber',6,27,23,1),(475,'.',1,27,24,0),(476,'Trots',5,28,1,1),(477,'dessa',5,28,2,1),(478,'uppenbara',9,28,3,1),(479,'påpekanden',10,28,4,1),(480,'ser',3,28,5,1),(481,'jag',3,28,6,1),(482,'hur',3,28,7,1),(483,'det',3,28,8,1),(484,'pågående',8,28,9,1),(485,'kriget',6,28,10,1),(486,'i',1,28,11,1),(487,'Ukraina',7,28,12,0),(488,'fiktionaliseras',15,28,13,1),(489,'.',1,28,14,0),(490,'Vi',2,29,1,1),(491,'har',3,29,2,1),(492,'två',3,29,3,1),(493,'huvudpersoner',13,29,4,1),(494,',',1,29,5,0),(495,'den',3,29,6,1),(496,'hjältemodige',12,29,7,1),(497,'president',9,29,8,1),(498,'Zelenskyj',9,29,9,0),(499,'och',3,29,10,1),(500,'den',3,29,11,1),(501,'galne',5,29,12,1),(502,'diktatorn',9,29,13,1),(503,'Vladimir',8,29,14,0),(504,'Putin',5,29,15,0),(505,'.',1,29,16,0),(506,'Sociala',7,30,1,1),(507,'medier',6,30,2,1),(508,'fylls',5,30,3,1),(509,'av',2,30,4,1),(510,'hyllningar',10,30,5,1),(511,'till',4,30,6,1),(512,'Zelenskyj',9,30,7,0),(513,'och',3,30,8,1),(514,'nidbilder',9,30,9,1),(515,'av',2,30,10,1),(516,'Putin',5,30,11,0),(517,',',1,30,12,0),(518,'reportage',9,31,1,1),(519,'om',2,31,2,1),(520,'den',3,31,3,1),(521,'osannolike',10,31,4,1),(522,'presidenten',11,31,5,1),(523,',',1,31,6,0),(524,'från',4,31,7,1),(525,'komiker',7,31,8,1),(526,'till',4,31,9,1),(527,'statschef',9,31,10,1),(528,'skrivs',6,31,11,1),(529,'i',1,31,12,1),(530,'rasande',7,31,13,1),(531,'fart',4,31,14,1),(532,',',1,31,15,0),(533,'liksom',6,31,16,1),(534,'analyser',8,31,17,1),(535,'som',3,31,18,1),(536,'avfärdar',8,31,19,1),(537,'Putin',5,31,20,0),(538,'som',3,31,21,1),(539,'”',1,31,22,0),(540,'sjuk',4,31,23,1),(541,'”',1,31,24,0),(542,',',1,31,25,0),(543,'”',1,31,26,0),(544,'vansinnig',9,31,27,1),(545,'”',1,31,28,0),(546,'och',3,31,29,1),(547,'”',1,31,30,0),(548,'galen',5,31,31,1),(549,'”',1,31,32,0),(550,'.',1,31,33,0);
/*!40000 ALTER TABLE `word` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-07-18 19:52:04
