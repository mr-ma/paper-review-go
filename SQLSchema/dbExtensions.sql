USE classification;

ALTER TABLE attribute ADD COLUMN x varchar(50) DEFAULT "";
ALTER TABLE attribute ADD COLUMN y varchar(50) DEFAULT "";
ALTER TABLE attribute ADD COLUMN xMajor varchar(50) DEFAULT "";
ALTER TABLE attribute ADD COLUMN yMajor varchar(50) DEFAULT "";
ALTER TABLE attribute ADD COLUMN x3D varchar(50) DEFAULT "";
ALTER TABLE attribute ADD COLUMN y3D varchar(50) DEFAULT "";
ALTER TABLE attribute ADD COLUMN z3D varchar(50) DEFAULT "";
ALTER TABLE attribute ADD COLUMN xMajor3D varchar(50) DEFAULT "";
ALTER TABLE attribute ADD COLUMN yMajor3D varchar(50) DEFAULT "";
ALTER TABLE attribute ADD COLUMN zMajor3D varchar(50) DEFAULT "";
ALTER TABLE dimension ADD COLUMN x varchar(50) DEFAULT "";
ALTER TABLE dimension ADD COLUMN y varchar(50) DEFAULT "";
ALTER TABLE dimension ADD COLUMN xMajor varchar(50) DEFAULT "";
ALTER TABLE dimension ADD COLUMN yMajor varchar(50) DEFAULT "";
ALTER TABLE attribute ADD COLUMN major tinyint(1) DEFAULT "0";

INSERT IGNORE INTO taxonomy_dimension (id_taxonomy, id_attribute, id_dimension) VALUES (1, 5, 1);

INSERT IGNORE INTO relation (id_relation, text, comment) VALUES (2, 'DependsDirected', 'directed dependency'),(3, 'InstanceOf', 'inheritance'),(4, 'MemberOf', 'aggregation'),(5, 'PartOf', 'composition');

/* attack view */
INSERT IGNORE INTO attribute (id_attribute,text) VALUES (80,'Tools'), (72,'Reverse engineering'), (73,'Attacker'), (84,'Discovery'), (79,'Disassembler'), (81,'Debugger'), (82,'Tracer'), (83,'Emulator'), (78,'Call interposition'), (85,'Pattern matching'), (86,'Taint analysis'), (87,'Graph-based analysis'), (88,'Symbolic execution');
/* defense view */
INSERT IGNORE INTO attribute (id_attribute,text) VALUES (71,'Software'), (89,'Dongle'), (90,'Self-check'), (91,'Hash chain'), (92,'Reactive');
/* new missing dimensions */
INSERT IGNORE taxonomy_dimension (id_taxonomy, id_attribute, id_dimension) VALUES (1, 71, 3), (1, 72, 2), (1, 73, 2), (1, 78, 2), (1, 79, 2), (1, 80, 2), (1, 81, 2), (1, 82, 2), (1, 83, 2), (1, 84, 2), (1, 85, 2), (1, 86, 2), (1, 87, 2), (1, 88, 2), (1, 89, 3), (1, 90, 3), (1, 91, 3), (1, 92, 3), (1, 93, 1), (1, 94, 1), (1, 95, 1);

/* system view */
INSERT IGNORE INTO attribute (id_attribute,text) VALUES (93,'Asset'), (94,'Link'), (95,'Basic block');

/* old missing dimensions */
INSERT IGNORE taxonomy_dimension (id_taxonomy, id_attribute, id_dimension) VALUES (1, 1, 1),(1,14,1),(1,20,1),(1,26,2),(1,27,2),(1,32,3),(1,35,3),(1,38,3),(1,41,3),(1,44,3),(1,50,3),(1,57,3),(1,62,3),(1,67,3);

DELETE FROM taxonomy_relation;

/* system relations */
INSERT IGNORE taxonomy_relation (id_taxonomy, id_src_attribute, id_dest_attribute, id_relation) VALUES (1, 2, 93, 3),(1, 3, 93, 3),(1, 4, 93, 3),(1, 5, 93, 1),(1, 21, 20, 3),(1, 22, 20, 3),(1, 23, 20, 3),(1, 24, 20, 3),(1, 25, 20, 3),(1, 94, 20, 3);
INSERT IGNORE taxonomy_relation (id_taxonomy, id_src_attribute, id_dest_attribute, id_relation) VALUES (1, 6, 21, 1),(1, 6, 22, 1),(1, 6, 94, 1),(1, 6, 5, 5),(1, 7, 5, 5),(1, 8, 5, 5),(1, 9, 7, 3),(1, 10, 7, 3),(1, 13, 8, 3),(1, 11, 8, 3),(1, 12, 11, 3),(1, 7, 24, 1),(1, 7, 25, 1),(1, 8, 24, 1),(1, 8, 25, 1);
INSERT IGNORE taxonomy_relation (id_taxonomy, id_src_attribute, id_dest_attribute, id_relation) VALUES (1, 5, 14, 1),(1, 15, 14, 3),(1, 95, 14, 3),(1, 17, 14, 3),(1, 18, 14, 3),(1, 19, 14, 3);

/* attack relations */
INSERT IGNORE taxonomy_relation (id_taxonomy, id_src_attribute, id_dest_attribute, id_relation) VALUES (1, 79, 80, 4),(1, 81, 80, 4),(1, 82, 80, 4),(1, 83, 80, 4),(1, 80, 72, 5),(1, 73, 72, 5),(1, 84, 72, 5),(1, 27, 72, 5);
INSERT IGNORE taxonomy_relation (id_taxonomy, id_src_attribute, id_dest_attribute, id_relation) VALUES (1, 28, 27, 3),(1, 29, 27, 3),(1, 30, 27, 3),(1, 88, 84, 3),(1, 31, 27, 3),(1, 26, 73, 3),(1, 85, 84, 3),(1, 86, 84, 3),(1, 87, 84, 3),(1, 78, 27, 3);

/* defense relations */
INSERT IGNORE taxonomy_relation (id_taxonomy, id_src_attribute, id_dest_attribute, id_relation) VALUES (1, 57, 32, 5),(1, 62, 32, 5),(1, 67, 32, 5),(1, 35, 32, 5),(1, 38, 32, 5),(1, 41, 32, 5),(1, 44, 32, 5),(1, 50, 32, 5),(1, 33, 32, 3),(1, 34, 32, 3);
INSERT IGNORE taxonomy_relation (id_taxonomy, id_src_attribute, id_dest_attribute, id_relation) VALUES (1, 58, 57, 3),(1, 59, 57, 3),(1, 60, 57, 3),(1, 61, 57, 3),(1, 71, 62, 3),(1, 89, 62, 3),(1, 63, 62, 3),(1, 64, 62, 3),(1, 65, 62, 3),(1, 70, 67, 3),(1, 69, 67, 3),(1, 90, 67, 3);
INSERT IGNORE taxonomy_relation (id_taxonomy, id_src_attribute, id_dest_attribute, id_relation) VALUES (1, 36, 35, 3),(1, 37, 35, 3),(1, 39, 38, 3),(1, 92, 38, 3),(1, 42, 41, 3),(1, 43, 41, 3),(1, 45, 44, 3),(1, 46, 44, 3),(1, 47, 44, 3),(1, 48, 44, 3),(1, 49, 44, 3),(1, 51, 50, 3),(1, 52, 50, 3),(1, 53, 50, 3),(1, 54, 50, 3),(1, 55, 50, 3),(1, 91, 50, 3);

/* major attributes (nodes) */
UPDATE attribute SET major = 1 WHERE text = "Tools";
UPDATE attribute SET major = 1 WHERE text = "Reverse engineering";
UPDATE attribute SET major = 1 WHERE text = "Attacker";
UPDATE attribute SET major = 1 WHERE text = "Discovery";
UPDATE attribute SET major = 1 WHERE text = "Attack";

UPDATE attribute SET major = 1 WHERE text = "Measure";
UPDATE attribute SET major = 1 WHERE text = "Overhead";
UPDATE attribute SET major = 1 WHERE text = "Trust anchor";
UPDATE attribute SET major = 1 WHERE text = "Protection level";

UPDATE attribute SET major = 1 WHERE text = "Asset";
UPDATE attribute SET major = 1 WHERE text = "Lifecycle activity";
UPDATE attribute SET major = 1 WHERE text = "Representation";
UPDATE attribute SET major = 1 WHERE text = "Granularity";

DROP TABLE IF EXISTS taxonomy_relation_annotation;
CREATE TABLE taxonomy_relation_annotation (
  id_taxonomy int(11) NOT NULL,
  id_taxonomy_relation int(11) NOT NULL,
  annotation longtext,
  PRIMARY KEY (id_taxonomy_relation),
  UNIQUE KEY id_taxonomy_relation_annotation_UNIQUE (id_taxonomy_relation)
) ENGINE=InnoDB AUTO_INCREMENT=71 DEFAULT CHARSET=utf8;

ALTER TABLE taxonomy_relation ADD COLUMN id_dimension int(10) DEFAULT "0";

INSERT IGNORE dimension (id_dimension, text) VALUES (4, "Interdimensional view");

/*INSERT IGNORE INTO taxonomy_dimension (id_taxonomy, id_attribute, id_dimension) VALUES (1, 32, 4), (1, 57, 4), (1, 62, 4), (1, 67, 4), (1, 93, 4), (1, 20, 4), (1, 5, 4), (1, 14, 4), (1, 84, 4), (1, 72, 4), (1, 80, 4), (1, 27, 4), (1, 73, 4);*/

/* Interdimensional relations */
INSERT IGNORE INTO taxonomy_relation (id_taxonomy_relation, id_taxonomy, id_src_attribute, id_dest_attribute, id_relation, id_dimension) VALUES (232, 1, 32, 84, 2, 4), (233, 1, 32, 20, 2, 4), (234, 1, 32, 27, 2, 4), (235, 1, 32, 67, 1, 4), (236, 1, 57, 20, 2, 4), (237, 1, 62, 32, 2, 4), (238, 1, 93, 14, 1, 4), (239, 1, 5, 93, 2, 4), (240, 1, 84, 32, 2, 4), (241, 1, 80, 72, 2, 4), (242, 1, 27, 5, 2, 4), (243, 1, 27, 32, 2, 4), (244, 1, 73, 27, 2, 4), (245, 1, 84, 5, 2, 4);
INSERT IGNORE INTO taxonomy_relation_annotation (id_taxonomy, id_taxonomy_relation, annotation) VALUES (1, 232, "Impedes"), (1, 233, "Transforms"), (1, 234, "Mitigates or raises the bar"), (1, 236, "Affects"), (1, 237, "Strengthens"), (1, 239, "Contains"), (1, 240, "Identifies"), (1, 241, "Support"), (1, 242, "Tampers with"), (1, 243, "Tampers with"), (1, 244, "Executes"), (1, 245, "Identifies asset");

ALTER TABLE taxonomy_relation ADD COLUMN edgeBendPoints longtext;

ALTER TABLE paper ADD COLUMN referenceCount int(20) DEFAULT "0";
ALTER TABLE paper ADD COLUMN author varchar(500) DEFAULT "";
ALTER TABLE paper ADD COLUMN keywords varchar(500) DEFAULT "";
ALTER TABLE mapping ADD COLUMN occurrenceCount int(20) DEFAULT "1";

ALTER TABLE mapping CHANGE COLUMN id_paper id_paper INT(11) UNSIGNED NOT NULL;
ALTER TABLE mapping CHANGE COLUMN id_attribute id_attribute INT(11) UNSIGNED NOT NULL;

ALTER TABLE attribute ADD COLUMN id_taxonomy INT(11) UNSIGNED DEFAULT 1;
ALTER TABLE attribute CHANGE COLUMN id_taxonomy id_taxonomy INT(11) UNSIGNED NOT NULL;
ALTER TABLE dimension ADD COLUMN id_taxonomy INT(11) UNSIGNED DEFAULT 1;
ALTER TABLE dimension CHANGE COLUMN id_taxonomy id_taxonomy INT(11) UNSIGNED NOT NULL;

ALTER TABLE attribute ADD UNIQUE KEY attribute_text_UNIQUE (text, id_taxonomy);
ALTER TABLE dimension ADD UNIQUE KEY dimension_text_UNIQUE (text, id_taxonomy);
/* ALTER TABLE mapping ADD UNIQUE KEY mapping_id_paper_id_attribute_UNIQUE (id_paper, id_attribute); */

/* change primary key */
ALTER TABLE mapping MODIFY id_mapping int(10) UNSIGNED NOT NULL, DROP PRIMARY KEY, ADD PRIMARY KEY (id_paper, id_attribute);
ALTER TABLE mapping ADD INDEX mapping_id_mapping (id_mapping), MODIFY id_mapping int(10) UNSIGNED NOT NULL AUTO_INCREMENT;

ALTER TABLE taxonomy_dimension ADD UNIQUE KEY taxonomy_dimension_id_attribute_UNIQUE (id_taxonomy, id_attribute);
ALTER TABLE taxonomy_relation ADD UNIQUE KEY taxonomy_relation_attributes_UNIQUE (id_taxonomy, id_src_attribute, id_dest_attribute, id_dimension);

/* foreign keys start */

SET FOREIGN_KEY_CHECKS = 0;

ALTER TABLE taxonomy CHANGE COLUMN id_taxonomy id_taxonomy INT(11) UNSIGNED NOT NULL AUTO_INCREMENT;

ALTER TABLE paper CHANGE COLUMN id_paper id_paper INT(11) UNSIGNED NOT NULL AUTO_INCREMENT;
ALTER TABLE paper CHANGE COLUMN citation citation varchar(500) NOT NULL;

ALTER TABLE attribute CHANGE COLUMN id_attribute id_attribute INT(11) UNSIGNED NOT NULL AUTO_INCREMENT;
ALTER TABLE attribute CHANGE COLUMN text text varchar(500) NOT NULL;
ALTER TABLE dimension CHANGE COLUMN id_dimension id_dimension INT(11) UNSIGNED NOT NULL AUTO_INCREMENT;

ALTER TABLE taxonomy_dimension CHANGE COLUMN id_attribute id_attribute INT(11) UNSIGNED NOT NULL;
ALTER TABLE taxonomy_dimension CHANGE COLUMN id_dimension id_dimension INT(11) UNSIGNED NOT NULL;
ALTER TABLE taxonomy_relation CHANGE COLUMN id_src_attribute id_src_attribute INT(11) UNSIGNED NOT NULL, CHANGE COLUMN id_dest_attribute id_dest_attribute INT(11) UNSIGNED NOT NULL;


ALTER TABLE taxonomy_dimension ADD CONSTRAINT taxonomy_dimension_id_attribute_foreign FOREIGN KEY (id_attribute) REFERENCES attribute (id_attribute) ON DELETE CASCADE ON UPDATE NO ACTION;
ALTER TABLE taxonomy_dimension ADD CONSTRAINT taxonomy_dimension_id_dimension_foreign FOREIGN KEY (id_dimension) REFERENCES dimension (id_dimension) ON DELETE CASCADE ON UPDATE NO ACTION;
ALTER TABLE taxonomy_relation ADD CONSTRAINT taxonomy_relation_id_src_attribute_foreign FOREIGN KEY (id_src_attribute) REFERENCES attribute (id_attribute) ON DELETE CASCADE ON UPDATE NO ACTION;
ALTER TABLE taxonomy_relation ADD CONSTRAINT taxonomy_relation_id_dest_attribute_foreign FOREIGN KEY (id_dest_attribute) REFERENCES attribute (id_attribute) ON DELETE CASCADE ON UPDATE NO ACTION;
ALTER TABLE mapping ADD CONSTRAINT mapping_id_attribute_foreign FOREIGN KEY (id_attribute) REFERENCES attribute (id_attribute) ON DELETE CASCADE ON UPDATE NO ACTION;
ALTER TABLE mapping ADD CONSTRAINT mapping_id_paper_foreign FOREIGN KEY (id_paper) REFERENCES paper (id_paper) ON DELETE CASCADE ON UPDATE NO ACTION;

ALTER TABLE attribute ADD CONSTRAINT attribute_id_taxonomy_foreign FOREIGN KEY (id_taxonomy) REFERENCES taxonomy (id_taxonomy) ON DELETE CASCADE ON UPDATE NO ACTION;
ALTER TABLE dimension ADD CONSTRAINT dimension_id_taxonomy_foreign FOREIGN KEY (id_taxonomy) REFERENCES taxonomy (id_taxonomy) ON DELETE CASCADE ON UPDATE NO ACTION;

SET FOREIGN_KEY_CHECKS = 1;

CREATE TABLE IF NOT EXISTS user (
  email varchar(128) NOT NULL,
  password varchar(128) NOT NULL,
  admin tinyint(1) DEFAULT "0",
  PRIMARY KEY (email),
  UNIQUE KEY email_UNIQUE (email)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO user (email, password, admin)
  VALUES
    ("felixhuber2@gmx.de", "123", 1);

/* foreign keys end */

/* dummy reference count data */
/* UPDATE paper SET referenceCount = id_paper; */

/* update taxonomy_relation.id_dimension */
UPDATE taxonomy_relation as relation SET id_dimension = (SELECT DISTINCT id_dimension from taxonomy_dimension WHERE taxonomy_dimension.id_attribute = relation.id_src_attribute) WHERE relation.id_dimension = 0;

/* store all children per attribute in a table */

DROP TABLE IF EXISTS allChildrenPerAttribute;
CREATE TABLE allChildrenPerAttribute (
  id_attribute int(10) unsigned NOT NULL,
  id_taxonomy int(11) unsigned NOT NULL,
  text varchar(50) NOT NULL,
  children longtext,
  PRIMARY KEY (id_attribute, id_taxonomy),
  UNIQUE KEY allChildrenPerAttribute_id_attribute_UNIQUE (id_attribute, id_taxonomy)
) ENGINE=InnoDB AUTO_INCREMENT=71 DEFAULT CHARSET=utf8;

/* procedure creation + call has to be run after every db import/recreation: */

DROP PROCEDURE IF EXISTS insertallchildrenperattribute;
DELIMITER ;;

CREATE PROCEDURE insertallchildrenperattribute(taxonomyId INT(11))
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
END;
;;

/* CALL insertallchildrenperattribute(1); */

DROP TABLE IF EXISTS allParentsPerAttribute;
CREATE TABLE allParentsPerAttribute (
  id_attribute int(11) unsigned NOT NULL,
  id_taxonomy int(11) unsigned NOT NULL,
  text varchar(50) NOT NULL,
  parents longtext,
  PRIMARY KEY (id_attribute, id_taxonomy),
  UNIQUE KEY allParentsPerAttribute_id_attribute_UNIQUE (id_attribute, id_taxonomy)
) ENGINE=InnoDB AUTO_INCREMENT=71 DEFAULT CHARSET=utf8;

DROP PROCEDURE IF EXISTS insertallparentsperattribute;
DELIMITER ;;

CREATE PROCEDURE insertallparentsperattribute(taxonomyId INT(11))
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
    INSERT INTO allParentsPerAttribute(id_attribute, id_taxonomy, text, parents) VALUES(cursor_id_attribute, taxonomyId, cursor_text, (SELECT (CASE WHEN b.parents IS NULL THEN "" ELSE b.parents END) AS parents FROM (SELECT GROUP_CONCAT(lv SEPARATOR ',') AS parents FROM (SELECT @pv:=(SELECT GROUP_CONCAT(DISTINCT parent.text SEPARATOR ',') FROM taxonomy_relation AS relation1 INNER JOIN attribute as parent ON (relation1.id_dest_attribute = parent.id_attribute AND parent.id_taxonomy = taxonomyId) WHERE relation1.id_taxonomy = 1 AND relation1.id_relation > 2 AND FIND_IN_SET((SELECT DISTINCT text FROM attribute WHERE id_attribute = relation1.id_src_attribute AND id_taxonomy = taxonomyId), @pv)) AS lv FROM taxonomy_relation JOIN (SELECT @pv:=text FROM attribute WHERE id_attribute = cursor_id_attribute AND id_taxonomy = taxonomyId) tmp) a) b));
  END LOOP;
  CLOSE cursor_i;
END;
;;

/* CALL insertallparentsperattribute(1); */

/* dummy taxonomy coordinates */
/*
SET FOREIGN_KEY_CHECKS=0;
REPLACE INTO attribute (id_attribute, text, x, y, xMajor, yMajor, major, x3D, y3D, z3D, xMajor3D, yMajor3D, zMajor3D) VALUES (1,'Integrity Protection Assets','1649.0310541692468','657.0731905134155',NULL,NULL,0,'2982.356969925021','-2285.019961769441','-864.2180305371212',NULL,NULL,NULL),(2,'Behavior','101.37013214331847','-33.71636299987085',NULL,NULL,0,'-4105.754594713023','2970.805490853027','520.6240466093806',NULL,NULL,NULL),(3,'Data','197.83932825786945','-13.324051457043254',NULL,NULL,0,'-5002.920658104225','2416.7437462541925','-667.9458699193781',NULL,NULL,NULL),(4,'Data and behavior','194.8960056011358','78.44135048568128',NULL,NULL,0,'-3116.814622838809','2150.293659766716','-739.4750876543186',NULL,NULL,NULL),(5,'Representation','44.108424258902005','290.011582742518','-1149.3243542368157','137.33226042504893',1,'-5631.8223863130315','-700.1480012803709','232.8933266727454','-2297.307274575899','1861.2451999882087','251.40799472895628'),(6,'Static','887.4460306295156','193.1481029140868',NULL,NULL,0,'-753.8117777226684','2672.900536068831','55.5562648093229',NULL,NULL,NULL),(7,'In memory','889.6007858584886','300.2077385139318',NULL,NULL,0,'-859.8475232068192','1419.0143417630459','1049.9932319804325',NULL,NULL,NULL),(8,'In execution','891.7555410874618','422.5616077708977',NULL,NULL,0,'-525.1480844836708','-997.8604194722125','-515.3976828985001',NULL,NULL,NULL),(9,'Code invariants','1204.8930473445566','121.77501251418998',NULL,NULL,0,'206.51897236227046','1964.5856119335288','1212.8380148609385',NULL,NULL,NULL),(10,'Data invariants','1222.3420362306504','277.2663880282506',NULL,NULL,0,'1584.2972884243895','1107.9452114293008','596.1960385203179',NULL,NULL,NULL),(11,'Trace','1218.2434728270118','556.502958256579',NULL,NULL,0,'1414.656335785149','-1776.66513018276','-191.61968505004643',NULL,NULL,NULL),(12,'Timed trace','1396.2819151128742','559.0519971994323',NULL,NULL,0,'2706.2310058466305','-505.4745310046228','-577.2241671516888',NULL,NULL,NULL),(13,'HW counters','1217.4549053992523','418.8548553424921',NULL,NULL,0,'1194.9742695396158','65.74160947273646','-137.41344005045494',NULL,NULL,NULL),(14,'Granularity','131.17003202979896','505.5221793995098','-897.3263229714349','350.66392710473633',1,'-4849.209350403896','-1872.1017791288107','20.015773720053858','2750.4431883226216','-2948.6211902334258','-502.6910289016656'),(15,'Instructions','439.20946040118747','411.20773851393204',NULL,NULL,0,'-3677.00212371179','-2059.3275936855316','-1231.9318958901586',NULL,NULL,NULL),(16,'BB','1659.8048303141127','727.2885674277603',NULL,NULL,0,'2143.4276500237497','-2464.144864036762','-1036.6394584036368',NULL,NULL,NULL),(17,'Function','433.32281508771996','574.346230856553',NULL,NULL,0,'-3825.8952008628225','-2536.6896428436626','-495.4057692218644',NULL,NULL,NULL),(18,'Slice','432.9285313738393','653.3664380850103',NULL,NULL,0,'-3435.1692734008243','-2081.4646016536462','1249.2503410645315',NULL,NULL,NULL),(19,'Application','440.1813644885202','734.9356842563209',NULL,NULL,0,'-3659.547946100021','-2773.817578816789','732.0912311726131',NULL,NULL,NULL),(20,'Lifecycle activity','187.43222543214645','186.89231154282785','-637.3283542055661','495.99612503027356',1,'-4688.243193111236','110.89828167666965','-422.7101612981819','3088.7134093484487','338.8041040030753','-411.428475846131'),(21,'Pre-compile','580.9837371136482','-13.090439485293977',NULL,NULL,0,'-2098.4675273781886','1903.2747623767602','708.8565788744929',NULL,NULL,NULL),(22,'Compile','585.687531285475','63.38072880030976',NULL,NULL,0,'-2060.668061605252','1333.5185979024836','700.4689784693802',NULL,NULL,NULL),(23,'Post-compile','587.8422865144481','132.20478025735312',NULL,NULL,0,'-2417.367759420388','1072.6497859901137','-40.15143208390509',NULL,NULL,NULL),(24,'Load','587.4480028005672','264.75480528573274',NULL,NULL,0,'-2148.2653425899725','-308.27053657036','59.54550985321248',NULL,NULL,NULL),(25,'Run','592.1517969723947','336.12789568562965',NULL,NULL,0,'-2078.8620483112054','-1015.7537643835454','82.76219647055268',NULL,NULL,NULL),(27,'Attack','940.3360043758586','193.66191904993806','-279.99781251708964','-209.3316979294433',1,'-560.3840244102373','102.1552480434342','-956.7462882957757','-1043.6553421224908','633.6478862886546','4.547473508864641e-13'),(28,'Binary ','724.3219621018262','83.03790265880495',NULL,NULL,0,'-669.670563198074','778.7013671716411','-1338.7442185818304',NULL,NULL,NULL),(29,'Process memory','906.7840594430213','83.03790265880497',NULL,NULL,0,'-94.03937291157544','-1336.7099182805853','1261.323300421252',NULL,NULL,NULL),(30,'Runtime data','1084.0008419920632','80.41524526272829',NULL,NULL,0,'-1286.5345978331447','-12.458919208386305','-1284.887427267443',NULL,NULL,NULL),(31,'Control flow','1242.8590227685681','138.11370797641482',NULL,NULL,0,'613.0758806397226','625.6623474191132','-1220.5401846641005',NULL,NULL,NULL),(32,'Measure','873.3942993388207','163.0985353391625','-771.9939687971188','-501.32941669726534',1,'-1540.1203017455578','164.35563430566253','519.8406371279184','-2885.666024652709','-362.649687924208','366.3875195434952'),(33,'Local','1110.8113585010585','131.57603389965158',NULL,NULL,0,'414.01249022013945','650.3924349711187','217.48400016653432',NULL,NULL,NULL),(34,'Remote','1118.2995836339228','209.45515510314914',NULL,NULL,0,'491.74632812015005','97.3780742973924','-123.8359011677826',NULL,NULL,NULL),(35,'Monitor','265.40894594719606','-14.91088455454599',NULL,NULL,0,'-3569.5793861554353','1840.7840486857215','360.30867064656013',NULL,NULL,NULL),(36,'State inspection','126.41025262586263','-94.64427054860292',NULL,NULL,0,'-5341.030340064095','2669.721222390146','-724.4818029428425',NULL,NULL,NULL),(37,'Introspection','109.15191089189352','43.71535517567398',NULL,NULL,0,'-4970.62799052557','1395.0657686414263','-235.261091445313',NULL,NULL,NULL),(38,'Response','452.9031877055812','-11.20235497342702',NULL,NULL,0,'-2593.6039171763614','2544.8218260668036','-759.650458124624',NULL,NULL,NULL),(39,'Proactive','342.9316064436389','-99.66960379790412',NULL,NULL,0,'-3372.016016595666','3430.1408603523755','-464.66544817688964',NULL,NULL,NULL),(40,'Postmortem','1854.8697013097888','296.6056002594439',NULL,NULL,0,'-5498.445586366219','-3168.741296759451','-466.7710346363615',NULL,NULL,NULL),(41,'Transformation','630.4856117755306','-16.918681805246013',NULL,NULL,0,'-674.7676646465982','2349.977089886205','-139.82051733421395',NULL,NULL,NULL),(42,'Manual','578.6373636104923','-117.0489804954572',NULL,NULL,0,'-921.5490625752307','3398.8951229358577','-479.87542564422074',NULL,NULL,NULL),(43,'Automatic','693.6729465958057','-122.6117748671356',NULL,NULL,0,'127.94846090778958','3219.9382164220856','20.700303685751805',NULL,NULL,NULL),(44,'Check','873.6077972507001','-2.0845634807704645',NULL,NULL,0,'1837.275949199416','2425.3707306235287','-191.4718377402669',NULL,NULL,NULL),(45,'Checksum','821.7595490856617','-126.32030444825449',NULL,NULL,0,'1281.3745416615216','2879.8359194492364','723.2800076185517',NULL,NULL,NULL),(46,'Signature','944.2121912332129','-130.0288340293734',NULL,NULL,0,'2246.4861157601076','2810.5915954973075','942.1178163301893',NULL,NULL,NULL),(47,'Equation eval','1088.916010867478','-130.02883402937346',NULL,NULL,0,'4753.967841269285','3722.180629086103','-665.6491212581645',NULL,NULL,NULL),(48,'Majority vote','1272.5593911034905','-130.02883402937346',NULL,NULL,0,'5182.609118082028','2736.231141284869','551.5771540960122',NULL,NULL,NULL),(49,'Access control','1433.9515938527904','-98.50633258986255',NULL,NULL,0,'6003.849502008718','2344.8187915623075','899.7533214767686',NULL,NULL,NULL),(50,'Hardening','1105.1773981587537','16.30455196468352',NULL,NULL,0,'3255.87157917825','840.7435575967074','88.59670625233503',NULL,NULL,NULL),(51,'Cyclic checks','1574.377556140925','-28.769818620181876',NULL,NULL,0,'5319.912403964804','1727.0256279508856','111.736167113791',NULL,NULL,NULL),(52,'Mutation','1572.594457320991','44.118523823075535',NULL,NULL,0,'4730.551295434692','949.487537634357','746.8857972408391',NULL,NULL,NULL),(53,'Code concealment','1580.0826824538556','109.0177914926567',NULL,NULL,0,'5356.577880976774','409.08695466782','279.84508938293175',NULL,NULL,NULL),(54,'Cloning','1574.591054052804','179.4798535339163',NULL,NULL,0,'5418.041195227526','-113.16968086050889','220.2466009992163',NULL,NULL,NULL),(55,'Layered interpretation','1592.5637433394631','248.0876507846166',NULL,NULL,0,'5193.809046312347','-656.3386097542594','748.8403578656307',NULL,NULL,NULL),(56,'Block chain','1858.4704534387163','371.42124195265205',NULL,NULL,0,'-6603.287282993331','-3590.335405994925','-415.31175748802025',NULL,NULL,NULL),(57,'Overhead','443.70302972341','181.33411832447578','-643.9949687893062','-167.99868751025386',1,'-3111.5808170396704','-114.66231598591185','866.8210719505937','164.35378747602084','1134.055950173481','-753.9175926735629'),(58,'Fair','242.52692883344275','352.5330491564459',NULL,NULL,0,'-5469.649462165422','-1106.5081701160102','688.9308524708213',NULL,NULL,NULL),(59,'Medium','322.36666614110965','353.66176587542634',NULL,NULL,0,'-4962.526245694544','-1110.5939508533318','750.2335288606791',NULL,NULL,NULL),(60,'High','399.8140466819174','355.44692168786355',NULL,NULL,0,'-4466.411357115115','-1226.125726118128','787.2822971600735',NULL,NULL,NULL),(61,'N/A','480.19061425803926','356.0880462774242',NULL,NULL,0,'-3918.0005729833483','-1244.1277199837466','842.8111976052214',NULL,NULL,NULL),(62,'Trust anchor','875.4620620412592','296.1450028790219','-894.6596771379391','-287.99775001757797',1,'-1852.7437633603195','-1111.1540361743937','-97.31951956107605','-1496.7187517790915','-1076.584406805348','1016.42713987056'),(63,'TPM','964.5379379587403','431.50633258986267',NULL,NULL,0,'-1184.9044012773659','-1789.3336295065346','572.4938645051229',NULL,NULL,NULL),(64,'SGX','1031.3626363895075','431.5063325898628',NULL,NULL,0,'-760.1823260827139','-1683.7777345675374','-319.8366899296602',NULL,NULL,NULL),(65,'Other','1109.3129235636318','433.3605973804221',NULL,NULL,0,'-294.31717100375323','-1603.1064742755818','528.323654467245',NULL,NULL,NULL),(66,'None','1859.7202051212796','450.88202131678383',NULL,NULL,0,'-8858.024127630688','-3467.2103127353244','-1518.0462141048988',NULL,NULL,NULL),(67,'Protection level','1307.8615880947477','346.2101522241273','-914.6595208891597','-155.9987812595215',1,'996.9534648630378','-1705.7452487059247','512.4409064068868','-2534.5247518544247','256.0799143751864','-203.61227561712337'),(68,'Internal','1848.7375791730985','231.62966387167714',NULL,NULL,0,'-7748.085581592561','-3476.806033232715','-1364.4036208165571',NULL,NULL,NULL),(69,'External','1384.6705318924207','472.6840866444073',NULL,NULL,0,'2529.0218352225365','-3007.2567771414238','285.6644084838954',NULL,NULL,NULL),(70,'Hypervisor','1254.9431625238847','456.740491586116',NULL,NULL,0,'1583.1531063878915','-2645.529174665351','703.9929349935978',NULL,NULL,NULL),(71,'Software','760.4976450265727','425.7900057580437',NULL,NULL,0,'-2797.358495533419','-2152.495606294804','-225.68108123829825',NULL,NULL,NULL),(72,'Reverse engineering','456.2492122594371','195.8121706901014','-51.99959375317334','27.999781251708782',1,'-150.57370347755636','-715.1167132090086','-15.635084340648746','-197.80873206507522','-917.9333817485949','326.9834909860764'),(73,'Attacker','455.1247623762343','304.1895132940248','-285.3311041840815','39.999687502441375',1,'1144.1859704610501','240.24312276877822','-194.02308367976838','1024.0980845461536','-88.9569907091659','-1366.4392732671088'),(78,'Call interposition','1246.347996774703','268.8145864563514',NULL,NULL,0,'-614.6855742024392','69.78717798899197','172.16880717031245',NULL,NULL,NULL),(79,'Disassembler','100.12950183558092','66.83901894908688',NULL,NULL,0,'-1995.7698478480065','-1269.2143260974635','-850.0941495725112',NULL,NULL,NULL),(80,'Tools','182.93061812586294','195.34923135684326','-58.666208336913826','-187.99853126147457',1,'-1164.2709035330786','-836.8747820341652','-862.5478927977981','-1342.1191295012839','-114.96057988172356','-15.715835268157775'),(81,'Debugger','222.45788738087995','69.46167634516351',NULL,NULL,0,'-898.1664589407969','-1951.0450139459258','-746.8683106090873',NULL,NULL,NULL),(82,'Tracer','339.07987666203144','105.12817980870786',NULL,NULL,0,'-1551.0856659947635','-478.5743355733096','-1242.1773531332879',NULL,NULL,NULL),(83,'Emulator','181.22698500534597','328.30650147624226',NULL,NULL,0,'-1150.7311957570096','-1400.1755468659949','217.24970035778017',NULL,NULL,NULL),(84,'Discovery','713.1518947222972','272.72709096376013','301.3309791850585','-165.33204167675774',1,'134.6578654717242','-290.1059399368047','911.8235960463293','-107.54154854705472','-1454.515798633701','553.3090234561244'),(85,'Pattern matching','676.5292554630063','398.4778093963763',NULL,NULL,0,'1037.20466085537','-142.0301624518263','1259.2820528498744',NULL,NULL,NULL),(86,'Taint analysis','852.9380630792848','401.93938973457784',NULL,NULL,0,'647.4909578714954','-571.3963813086439','1710.4840235215438',NULL,NULL,NULL),(87,'Graph-based analysis','1061.657682390608','399.1489477500765',NULL,NULL,0,'696.823432627275','-1526.425663903857','1189.8850144020796',NULL,NULL,NULL),(88,'Symbolic execution','1302.992734564465','342.1965639504301',NULL,NULL,0,'662.2672080618793','-1189.9103668925495','887.5497529647132',NULL,NULL,NULL),(89,'Dongle','882.9502871741241','431.3528001297221',NULL,NULL,0,'-1758.2619761800556','-1832.5148037505676','517.4632417765356',NULL,NULL,NULL),(90,'Self-check','1527.733584648006','468.36898696278894',NULL,NULL,0,'2254.929586909267','-2091.1030565912292','640.2494531206704',NULL,NULL,NULL),(91,'Hash chain','1560.609501237264','336.10961673256753',NULL,NULL,0,'5428.823081049158','-1392.7422049872685','622.1884914602401',NULL,NULL,NULL),(92,'Reactive','465.9542072101243','-113.80104829476028',NULL,NULL,0,'-2156.9317741090754','3388.5632998016813','29.848118994063952',NULL,NULL,NULL),(93,'Asset','-25.109910912021626','78.67496245743045','-1157.3242917373038','347.99728127124',1,'-4808.952804641889','1030.5020433257075','5.0788317295755405','3129.4424613034453','-1133.9728517827305','-488.9199597681988'),(94,'Link','586.265151658927','195.93075382868957',NULL,NULL,0,'-2652.1505220442427','412.8804880622755','131.98179481604484',NULL,NULL,NULL),(95,'Basic block','438.0266092595463','491.6192711996905',NULL,NULL,0,'-5634.276242050955','-2183.7509211360775','698.4873866131952',NULL,NULL,NULL),(97,'asdf','181.67928743243047','461.74826607962','','',0,NULL,NULL,NULL,NULL,NULL,NULL),(98,'23','213.05761709609442','181.12503608253778','','',0,NULL,NULL,NULL,NULL,NULL,NULL),(26,'Not root','455.7266028589016','393.7732747885879','','',0,NULL,NULL,NULL,NULL,NULL,NULL),(103,'ffffffeee','620.1434389341558','504.7649461474462','','',0,NULL,NULL,NULL,NULL,NULL,NULL),(104,'lulkek','1033.4172749509523','624.2618774270801','','',0,NULL,NULL,NULL,NULL,NULL,NULL);
SET FOREIGN_KEY_CHECKS=1;
*/