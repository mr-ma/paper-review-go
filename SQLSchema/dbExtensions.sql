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
ALTER TABLE attribute ADD COLUMN synonyms varchar(500) DEFAULT "";

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
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

ALTER TABLE taxonomy_relation ADD COLUMN id_dimension int(10) DEFAULT "0";

INSERT IGNORE dimension (id_dimension, text) VALUES (4, "Interdimensional view");

/* Interdimensional relations */
INSERT IGNORE INTO taxonomy_relation (id_taxonomy_relation, id_taxonomy, id_src_attribute, id_dest_attribute, id_relation, id_dimension) VALUES (232, 1, 32, 84, 2, 4), (233, 1, 32, 20, 2, 4), (234, 1, 32, 27, 2, 4), (235, 1, 32, 67, 1, 4), (236, 1, 57, 20, 2, 4), (237, 1, 62, 32, 2, 4), (238, 1, 93, 14, 1, 4), (239, 1, 5, 93, 2, 4), (240, 1, 84, 32, 2, 4), (241, 1, 80, 72, 2, 4), (242, 1, 27, 5, 2, 4), (243, 1, 27, 32, 2, 4), (244, 1, 73, 27, 2, 4), (245, 1, 84, 5, 2, 4);
INSERT IGNORE INTO taxonomy_relation_annotation (id_taxonomy, id_taxonomy_relation, annotation) VALUES (1, 232, "Impedes"), (1, 233, "Transforms"), (1, 234, "Mitigates or raises the bar"), (1, 236, "Affects"), (1, 237, "Strengthens"), (1, 239, "Contains"), (1, 240, "Identifies"), (1, 241, "Support"), (1, 242, "Tampers with"), (1, 243, "Tampers with"), (1, 244, "Executes"), (1, 245, "Identifies asset");

ALTER TABLE taxonomy_relation ADD COLUMN edgeBendPoints longtext;

ALTER TABLE paper ADD COLUMN id_taxonomy INT(11) UNSIGNED DEFAULT 1;
ALTER TABLE paper ADD COLUMN referenceCount int(20) DEFAULT "0";
ALTER TABLE paper ADD COLUMN author varchar(500) DEFAULT "";
ALTER TABLE paper ADD COLUMN keywords varchar(500) DEFAULT "";

ALTER TABLE paper DROP INDEX id_paper_UNIQUE;
ALTER TABLE paper DROP INDEX paper_id_paper;
ALTER TABLE paper ADD UNIQUE KEY paper_id_UNIQUE (id_taxonomy, id_paper);
ALTER TABLE paper ADD UNIQUE KEY paper_citation_UNIQUE (id_taxonomy, citation);

/* change primary key */
ALTER TABLE paper MODIFY id_paper int(11) UNSIGNED NOT NULL, DROP PRIMARY KEY, ADD PRIMARY KEY (id_taxonomy, id_paper);
ALTER TABLE paper ADD INDEX paper_id_paper (id_paper), MODIFY id_paper int(11) UNSIGNED NOT NULL AUTO_INCREMENT;
ALTER TABLE paper ADD CONSTRAINT paper_id_taxonomy_foreign FOREIGN KEY (id_taxonomy) REFERENCES taxonomy (id_taxonomy) ON DELETE CASCADE ON UPDATE NO ACTION;

ALTER TABLE mapping ADD COLUMN occurrenceCount int(20) DEFAULT "1";
ALTER TABLE mapping CHANGE COLUMN id_paper id_paper INT(11) UNSIGNED NOT NULL;
ALTER TABLE mapping CHANGE COLUMN id_attribute id_attribute INT(11) UNSIGNED NOT NULL;

ALTER TABLE attribute ADD COLUMN id_taxonomy INT(11) UNSIGNED DEFAULT 1;
ALTER TABLE attribute CHANGE COLUMN id_taxonomy id_taxonomy INT(11) UNSIGNED NOT NULL;
ALTER TABLE dimension ADD COLUMN id_taxonomy INT(11) UNSIGNED DEFAULT 1;
ALTER TABLE dimension CHANGE COLUMN id_taxonomy id_taxonomy INT(11) UNSIGNED NOT NULL;

ALTER TABLE attribute ADD UNIQUE KEY attribute_text_UNIQUE (id_taxonomy, text);
ALTER TABLE dimension ADD UNIQUE KEY dimension_text_UNIQUE (id_taxonomy, text);

/* change primary key */
ALTER TABLE mapping MODIFY id_mapping int(10) UNSIGNED NOT NULL, DROP PRIMARY KEY, ADD PRIMARY KEY (id_paper, id_attribute);
ALTER TABLE mapping ADD INDEX mapping_id_mapping (id_mapping), MODIFY id_mapping int(10) UNSIGNED NOT NULL AUTO_INCREMENT;

ALTER TABLE taxonomy_dimension ADD UNIQUE KEY taxonomy_dimension_id_attribute_UNIQUE (id_attribute);
ALTER TABLE taxonomy_relation ADD UNIQUE KEY taxonomy_relation_attributes_UNIQUE (id_taxonomy, id_src_attribute, id_dest_attribute, id_dimension);

/* foreign keys start */

SET FOREIGN_KEY_CHECKS = 0;

ALTER TABLE taxonomy CHANGE COLUMN id_taxonomy id_taxonomy INT(11) UNSIGNED NOT NULL AUTO_INCREMENT;

/* ALTER TABLE paper CHANGE COLUMN id_paper id_paper INT(11) UNSIGNED NOT NULL AUTO_INCREMENT; */
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

ALTER TABLE vote_tags ADD CONSTRAINT vote_tags_TagId_foreign FOREIGN KEY (Tag_Id) REFERENCES tags (TagId) ON DELETE CASCADE ON UPDATE NO ACTION;
ALTER TABLE vote_tags ADD CONSTRAINT vote_tags_VoteId_foreign FOREIGN KEY (VoteId) REFERENCES votes (VoteId) ON DELETE CASCADE ON UPDATE NO ACTION;

SET FOREIGN_KEY_CHECKS = 1;

DROP TABLE IF EXISTS user;
CREATE TABLE IF NOT EXISTS user (
  id int(11) unsigned NOT NULL AUTO_INCREMENT,
  email varchar(128) NOT NULL,
  name varchar(128) NOT NULL,
  password varchar(128) NOT NULL,
  taxonomies varchar(500) DEFAULT "",
  admin tinyint(1) DEFAULT "0",
  PRIMARY KEY (id),
  UNIQUE KEY user_id_UNIQUE (id),
  UNIQUE KEY user_email_UNIQUE (email)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

/* dummy user admin account */

INSERT INTO user (email, name, password, admin)
  VALUES
    ("mohsen.ahmadvand@tum.de", "Mohsen Ahmadvand", "123", 1),
    ("amjad.ibrahim@tum.de", "Amjad Ibrahim", "123", 1),
    ("felixhuber2@gmx.de", "Felix Huber", "123", 1);

/* foreign keys end */

UPDATE taxonomy_relation as relation SET id_dimension = (SELECT DISTINCT id_dimension from taxonomy_dimension WHERE taxonomy_dimension.id_attribute = relation.id_src_attribute) WHERE relation.id_dimension = 0;

/* store all children per attribute in a table */

DROP TABLE IF EXISTS allchildrenperattribute;
CREATE TABLE allchildrenperattribute (
  id_attribute int(11) unsigned NOT NULL,
  id_taxonomy int(11) unsigned NOT NULL,
  text varchar(50) NOT NULL,
  children longtext,
  PRIMARY KEY (id_attribute),
  UNIQUE KEY allchildrenperattribute_id_attribute_UNIQUE (id_attribute),
  CONSTRAINT allchildrenperattribute_id_attribute_foreign FOREIGN KEY (id_attribute) REFERENCES attribute (id_attribute) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

DROP PROCEDURE IF EXISTS insertallchildrenperattribute;
DELIMITER ;;

CREATE PROCEDURE insertallchildrenperattribute(taxonomyId INT(11))
BEGIN
  DECLARE cursor_id_attribute INT(11);
  DECLARE cursor_text VARCHAR(50);
  DECLARE done INT DEFAULT FALSE;
  DECLARE cursor_i CURSOR FOR SELECT id_attribute, text FROM attribute;
  DECLARE CONTINUE HANDLER FOR NOT FOUND SET done = TRUE;
  DELETE FROM allchildrenperattribute WHERE id_taxonomy = taxonomyId;
  OPEN cursor_i;
  read_loop: LOOP
    FETCH cursor_i INTO cursor_id_attribute, cursor_text;
    IF done THEN
      LEAVE read_loop;
    END IF;
    INSERT IGNORE INTO allchildrenperattribute(id_attribute, id_taxonomy, text, children) VALUES(cursor_id_attribute, taxonomyId, cursor_text, (SELECT (CASE WHEN b.children IS NULL THEN CAST(cursor_id_attribute AS CHAR(50)) ELSE CONCAT(CAST(cursor_id_attribute AS CHAR(50)), ",", b.children) END) AS children FROM (SELECT GROUP_CONCAT(lv SEPARATOR ',') AS children FROM (SELECT @pv:=(SELECT GROUP_CONCAT(DISTINCT relation1.id_src_attribute SEPARATOR ',') FROM taxonomy_relation AS relation1 WHERE relation1.id_taxonomy = taxonomyId AND relation1.id_relation > 2 AND FIND_IN_SET(relation1.id_dest_attribute, @pv)) AS lv FROM taxonomy_relation AS relation2 JOIN (SELECT @pv:=cursor_id_attribute) tmp ON (relation2.id_taxonomy = taxonomyId)) a) b));
  END LOOP;
  CLOSE cursor_i;
END;
;;

/* store all parents per attribute in a table */

DROP TABLE IF EXISTS allparentsperattribute;
CREATE TABLE allparentsperattribute (
  id_attribute int(11) unsigned NOT NULL,
  id_taxonomy int(11) unsigned NOT NULL,
  text varchar(50) NOT NULL,
  parents longtext,
  PRIMARY KEY (id_attribute),
  UNIQUE KEY allparentsperattribute_id_attribute_UNIQUE (id_attribute),
  CONSTRAINT allparentsperattribute_id_attribute_foreign FOREIGN KEY (id_attribute) REFERENCES attribute (id_attribute) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

DROP PROCEDURE IF EXISTS insertallparentsperattribute;
DELIMITER ;;

CREATE PROCEDURE insertallparentsperattribute(taxonomyId INT(11))
BEGIN
  DECLARE cursor_id_attribute INT(11);
  DECLARE cursor_text VARCHAR(50);
  DECLARE done INT DEFAULT FALSE;
  DECLARE cursor_i CURSOR FOR SELECT id_attribute, text FROM attribute;
  DECLARE CONTINUE HANDLER FOR NOT FOUND SET done = TRUE;
  DELETE FROM allparentsperattribute WHERE id_taxonomy = taxonomyId;
  OPEN cursor_i;
  read_loop: LOOP
    FETCH cursor_i INTO cursor_id_attribute, cursor_text;
    IF done THEN
      LEAVE read_loop;
    END IF;
    INSERT IGNORE INTO allparentsperattribute(id_attribute, id_taxonomy, text, parents) VALUES(cursor_id_attribute, taxonomyId, cursor_text, (SELECT (CASE WHEN b.parents IS NULL THEN "" ELSE b.parents END) AS parents FROM (SELECT GROUP_CONCAT(lv SEPARATOR ',') AS parents FROM (SELECT @pv:=(SELECT GROUP_CONCAT(DISTINCT parent.text SEPARATOR ',') FROM taxonomy_relation AS relation1 INNER JOIN attribute as parent ON (relation1.id_dest_attribute = parent.id_attribute AND parent.id_taxonomy = taxonomyId) WHERE relation1.id_taxonomy = taxonomyId AND relation1.id_relation > 2 AND FIND_IN_SET((SELECT DISTINCT text FROM attribute WHERE id_attribute = relation1.id_src_attribute AND id_taxonomy = taxonomyId), @pv)) AS lv FROM taxonomy_relation JOIN (SELECT @pv:=text FROM attribute WHERE id_attribute = cursor_id_attribute AND id_taxonomy = taxonomyId) tmp) a) b));
  END LOOP;
  CLOSE cursor_i;
END;
;;