package data

import (
	"fmt"
	"strconv"
	"strings"
	//overriding MySqlDriver
	//"github.com/go-sql-driver/mysql"
	"../model"
)

type TaxonomyBuilderDriver interface {
	DriverCore
	GetAllTaxonomies() ([]model.Taxonomy, error)
	GetTaxonomyID(string) ([]model.Taxonomy, error)
	GetTaxonomyPermissions(string) ([]model.Taxonomy, error)
	UpdateTaxonomyPermissions(string, string) (model.Result, error)
	AddTaxonomy(string, string) (model.Result, error)
	RemoveTaxonomy(int64) (model.Result, error)
	UpdateCitationReferenceCounts(int64, []model.ReferenceCount) (model.Result, error)
	UpdateMajor(int64, string, int) (model.Result, error)
	UpdateCitationMapping(int64, string, []model.Paper) (model.Result, error)
	UpdateCitationMappings(int64, []model.CitationMapping) (model.Result, error)
	SavePositions(int64, []model.Position) (model.Result, error)
	SaveMajorPositions(int64, []model.Position) (model.Result, error)
	Save3DPositions(int64, []model.Position) (model.Result, error)
	SaveMajor3DPositions(int64, []model.Position) (model.Result, error)
	SaveEdgeBendPoints(int64, string, string, string, string) (model.Result, error)
	RemoveTaxonomyRelationsForAttribute(int64, model.Attribute) (model.Result, error)
	AddAttribute(int64, model.Attribute) (model.Result, error)
	AddDimension(int64, string) (model.Result, error)
	ChangeDimension(int64, string, string) (model.Result, error)
	RenameAttribute(int64, string, string) (model.Result, error)
	UpdateSynonyms(int64, string, string) (model.Result, error)
	RenameDimension(int64, string, string) (model.Result, error)
	CheckIfRelationIsValid(model.AttributeRelation) bool
	AddTaxonomyRelation(model.AttributeRelation) (model.Result, error)
	DeleteCitation(int64, model.Paper) (model.Result, error)
	RemoveAttribute(int64, model.Attribute) (model.Result, error)
	RemoveDimension(int64, model.Dimension) (model.Result, error)
	RemoveTaxonomyRelation(model.AttributeRelation) (model.Result, error)
	UpdateTaxonomyRelationType(model.AttributeRelation) (model.Result, error)
	UpdateTaxonomyRelationAnnotation(model.AttributeRelation) (model.Result, error)
	UpdateChildRelationshipTable(int64) chan model.Result
	UpdateParentRelationshipTable(int64) chan model.Result
	UpdateRelationshipTables(int64) model.Result
	MergeAttributes(int64, model.Attribute, model.Attribute) (model.Result, error)
	ForkAttribute(int64, string, string, []model.AttributeRelation, []model.AttributeRelation, []model.AttributeRelation, []model.AttributeRelation, []model.Paper, []model.Paper) (model.Result, error)
}

//InitMySQLDriver initialize a new my sql driver instance
func InitTaxonomyBuilderDriver(user string, password string, server string) TaxonomyBuilderDriver {
	return MySQLDriver{username: user, pass: password, database: "classification", server: server}
}


func (d MySQLDriver) GetAllTaxonomies() (taxonomies []model.Taxonomy,
	err error) {
	dbRef, err := d.OpenDB()
	defer dbRef.Close()
	checkErr(err)
	db, stmt, err := d.Query(`select id_taxonomy, text from taxonomy;`)
	defer stmt.Close()
	defer db.Close()
	rows, err := stmt.Query()
	checkErr(err)
	for rows.Next() {
		a := model.Taxonomy{}
		rows.Scan(&a.ID, &a.Text)
		taxonomies = append(taxonomies, a)
	}
	defer rows.Close()
	return taxonomies, err
}

func (d MySQLDriver) GetTaxonomyID(text string) (taxonomies []model.Taxonomy,
	err error) {
	dbRef, err := d.OpenDB()
	defer dbRef.Close()
	checkErr(err)
	db, stmt, err := d.Query("select id_taxonomy from taxonomy where BINARY text = ?;")
	defer stmt.Close()
	defer db.Close()
	rows, err := stmt.Query(text)
	checkErr(err)
	for rows.Next() {
		a := model.Taxonomy{}
		rows.Scan(&a.ID)
		taxonomies = append(taxonomies, a)
	}
	defer rows.Close()
	return taxonomies, err
}

func (d MySQLDriver) GetTaxonomyPermissions(email string) (taxonomies []model.Taxonomy, err error) {
	dbRef, err := d.OpenDB()
	defer dbRef.Close()
	checkErr(err)
	db, stmt, err := d.Query("SELECT taxonomies FROM user WHERE email = ?;")
	defer stmt.Close()
	defer db.Close()
	rows, err := stmt.Query(email)
	checkErr(err)
	var taxonomyPermissions string
	taxonomyPermissions = ""
	for rows.Next() {
		rows.Scan(&taxonomyPermissions)
	}
	if taxonomyPermissions != "" {
		array := strings.Split(taxonomyPermissions, ",")
		for _, elem := range array {
			id, err := strconv.Atoi(elem)
			if err == nil {
				a := model.Taxonomy{ID: id}
				taxonomies = append(taxonomies, a)
			}
		}
	}
	defer rows.Close()
	return taxonomies, err
}

func (d MySQLDriver) UpdateTaxonomyPermissions(email string, permissions string) (result model.Result, err error) {
	dbRef, err := d.OpenDB()
	defer dbRef.Close()
	checkErr(err)
	dbRef.Exec("UPDATE user SET taxonomies = ? WHERE email = ?;", permissions, email)
	result.Success = true
	return result, err
}

func (d MySQLDriver) AddTaxonomy(taxonomy string, dimension string) (result model.Result, err error) {
	dbRef, err := d.OpenDB()
	defer dbRef.Close()
	checkErr(err)
	if !strings.Contains(dimension, " view") {
		dimension += " view"
	}
	db, stmt, err := d.Query("select count(id_taxonomy) as taxonomyCount from taxonomy where BINARY text = ?;")
	defer stmt.Close()
	defer db.Close()
	rows, err := stmt.Query(taxonomy)
	checkErr(err)
	var rowCount int
	rowCount = 0
	for rows.Next() {
		rows.Scan(&rowCount)
	}
	defer rows.Close()
	if rowCount > 0 {
		result.Success = false
		return result, err
	}
	dbRef.Exec("INSERT IGNORE INTO taxonomy (text) VALUES (?);", taxonomy)
	dbRef.Exec("INSERT IGNORE INTO dimension (id_taxonomy, text) VALUES ((SELECT DISTINCT id_taxonomy FROM taxonomy WHERE BINARY text = ?), \"Interdimensional view\");", taxonomy)
	dbRef.Exec("INSERT IGNORE INTO dimension (id_taxonomy, text) VALUES ((SELECT DISTINCT id_taxonomy FROM taxonomy WHERE BINARY text = ?), ?);", taxonomy, dimension)
	result.Success = true
	return result, err
}

func (d MySQLDriver) RemoveTaxonomy(taxonomyId int64) (result model.Result, err error) {
	dbRef, err := d.OpenDB()
	defer dbRef.Close()
	checkErr(err)
	taxonomyIdStr := strconv.Itoa(int(taxonomyId))
	dbRef.Exec("DELETE FROM taxonomy WHERE id_taxonomy = ?;", taxonomyIdStr)
	result.Success = true
	return result, err
}


func (d MySQLDriver) UpdateCitationReferenceCounts(taxonomyId int64, referenceCounts []model.ReferenceCount) (result model.Result, err error) {
	dbRef, err := d.OpenDB()
	defer dbRef.Close()
	checkErr(err)
	taxonomyIdStr := strconv.Itoa(int(taxonomyId))
	for _, elem := range referenceCounts {
		referenceCountStr := strconv.Itoa(elem.ReferenceCount)
		dbRef.Exec("update paper set referenceCount = ? where id_taxonomy = ? and BINARY citation = ?;", referenceCountStr, taxonomyIdStr, elem.Citation)
	}
	result.Success = true
	return result, err
}

func (d MySQLDriver) UpdateMajor(taxonomyId int64, text string, major int) (result model.Result, err error) {
	dbRef, err := d.OpenDB()
	defer dbRef.Close()
	checkErr(err)
	majorStr := strconv.Itoa(major)
	taxonomyIdStr := strconv.Itoa(int(taxonomyId))
	dbRef.Exec("update attribute set major = ? where BINARY text = ? and id_taxonomy = ?;", majorStr, text, taxonomyIdStr)
	result.Success = true
	return result, err
}

func (d MySQLDriver) UpdateCitationMapping(taxonomyId int64, attribute string, citations []model.Paper) (result model.Result, err error) {
	dbRef, err := d.OpenDB()
	defer dbRef.Close()
	checkErr(err)
	taxonomyIdStr := strconv.Itoa(int(taxonomyId))
	db, stmt, err := d.Query("SELECT DISTINCT id_attribute FROM attribute WHERE BINARY text = ? and id_taxonomy = ?;")
	defer stmt.Close()
	defer db.Close()
	rows, err := stmt.Query(attribute, taxonomyIdStr)
	checkErr(err)
	var attributeID int
	attributeID = -1
	for rows.Next() {
		rows.Scan(&attributeID)
	}
	defer rows.Close()
	if attributeID < 0 {
		result.Success = false
		return result, err
	}
	attributeIDStr := strconv.Itoa(attributeID)
	/*
		db2, stmt2, err2 := d.Query("SELECT COUNT(DISTINCT id_src_attribute) FROM taxonomy_relation WHERE id_dest_attribute = ? AND id_relation > 2;")
		defer stmt2.Close()
		defer db2.Close()
		rows2, err2 := stmt2.Query(attributeIDStr)
		checkErr(err2)
		var childrenCount int
		for rows2.Next() {
			rows2.Scan(&childrenCount)
		}
		defer rows2.Close()
		if childrenCount > 0 {
			result.Success = false
			return result, err2
		}
	*/
	citationTupleString := ""
	for _, elem := range citations {
		if citationTupleString != "" {
			citationTupleString += ","
		}
		citationTupleString += "(" + strconv.Itoa(elem.ID) + "," + attributeIDStr + ")"
	}
	dbRef.Exec("DELETE FROM mapping WHERE id_attribute = ? AND (id_paper, id_attribute) NOT IN ("+citationTupleString+");", attributeIDStr)
	if citationTupleString != "" {
		dbRef.Exec("INSERT IGNORE INTO mapping (id_paper, id_attribute) VALUES " + citationTupleString + ";")
	}
	result.Success = true
	return result, err
}

func (d MySQLDriver) UpdateCitationMappings(taxonomyId int64, mappings []model.CitationMapping) (result model.Result, err error) {
	dbRef, err := d.OpenDB()
	defer dbRef.Close()
	checkErr(err)
	taxonomyIdStr := strconv.Itoa(int(taxonomyId))
	savedPapers := []model.Paper{}
	for _, elem := range mappings {
		paperIDStr := strconv.Itoa(elem.PaperID)
		found := false
		for i := 0; i < len(savedPapers); i++ {
			if savedPapers[i].ID == elem.PaperID {
				found = true
				break
			}
		}
		paperID := elem.PaperID
		if !found {
			paperIDStr = strconv.Itoa(paperID)
			referenceCountStr := strconv.Itoa(elem.ReferenceCount)
			if paperID < 0 {
				db, stmt, err := d.Query("SELECT MAX(id_paper) AS maxID FROM paper;")
				rows, err := stmt.Query()
				checkErr(err)
				stmt.Close()
				db.Close()
				var maxID int
				maxID = 0
				for rows.Next() {
					rows.Scan(&maxID)
				}
				rows.Close()
				paperID = maxID + 1
				paperIDStr = strconv.Itoa(paperID)
				bibTex := ""
				if elem.Bib != "empty" {
					bibTex = elem.Bib
				}
				dbRef.Exec("INSERT IGNORE INTO paper (id_taxonomy, id_paper, citation, bib, referenceCount, author, keywords) VALUES (?, ?, ?, ?, ?, ?, ?);", taxonomyIdStr, paperIDStr, elem.Citation, bibTex, referenceCountStr, elem.Author, elem.Keywords)
			} else {
				if elem.Bib != "empty" {
					dbRef.Exec("UPDATE paper SET BINARY citation = ?, bib = ?, referenceCount = ?, author = ?, keywords = ? WHERE id_taxonomy = ? AND id_paper = ?;", elem.Citation, elem.Bib, referenceCountStr, elem.Author, elem.Keywords, taxonomyIdStr, paperIDStr)
				} else {
					dbRef.Exec("UPDATE paper SET BINARY citation = ?, referenceCount = ?, author = ?, keywords = ?) WHERE id_taxonomy = ? AND id_paper = ?;", elem.Citation, referenceCountStr, elem.Author, elem.Keywords, taxonomyIdStr, paperIDStr)
				}
			}
			a := model.Paper{ID: paperID}
			savedPapers = append(savedPapers, a)
		}
		if elem.OccurrenceCount <= 0 {
			dbRef.Exec("DELETE FROM mapping WHERE id_paper = (SELECT DISTINCT id_paper FROM paper WHERE id_taxonomy = ? AND BINARY citation = ?) AND id_attribute = (SELECT DISTINCT id_attribute FROM attribute WHERE BINARY text = ? AND id_taxonomy = ?);", taxonomyIdStr, elem.Citation, elem.Attribute, taxonomyIdStr)
		} else {
			occurrenceCountStr := strconv.Itoa(elem.OccurrenceCount)
			dbRef.Exec("REPLACE INTO mapping (id_paper, id_attribute, occurrenceCount) VALUES ((SELECT DISTINCT id_paper FROM paper WHERE id_taxonomy = ? AND BINARY citation = ?), (SELECT DISTINCT id_attribute FROM attribute WHERE BINARY text = ? AND id_taxonomy = ?), ?);", taxonomyIdStr, elem.Citation, elem.Attribute, taxonomyIdStr, occurrenceCountStr)
		}
	}
	d.UpdateRelationshipTables(taxonomyId)
	result.Success = true
	return result, err
}


func (d MySQLDriver) SavePositions(taxonomyId int64, positions []model.Position) (result model.Result, err error) {
	dbRef, err := d.OpenDB()
	defer dbRef.Close()
	checkErr(err)
	taxonomyIdStr := strconv.Itoa(int(taxonomyId))
	for _, elem := range positions {
		dbRef.Exec("update "+elem.Table+" set x = ?, y = ? where BINARY text = ? and id_taxonomy = ?;", elem.X, elem.Y, elem.ID, taxonomyIdStr)
	}
	result.Success = true
	return result, err
}

func (d MySQLDriver) SaveMajorPositions(taxonomyId int64, positions []model.Position) (result model.Result, err error) {
	dbRef, err := d.OpenDB()
	defer dbRef.Close()
	checkErr(err)
	taxonomyIdStr := strconv.Itoa(int(taxonomyId))
	for _, elem := range positions {
		dbRef.Exec("update "+elem.Table+" set xMajor = ?, yMajor = ? where BINARY text = ? and id_taxonomy = ?;", elem.X, elem.Y, elem.ID, taxonomyIdStr)
	}
	result.Success = true
	return result, err
}

func (d MySQLDriver) Save3DPositions(taxonomyId int64, positions []model.Position) (result model.Result, err error) {
	dbRef, err := d.OpenDB()
	defer dbRef.Close()
	checkErr(err)
	taxonomyIdStr := strconv.Itoa(int(taxonomyId))
	for _, elem := range positions {
		dbRef.Exec("update "+elem.Table+" set x3D = ?, y3D = ?, z3D = ? where BINARY text = ? and id_taxonomy = ?;", elem.X, elem.Y, elem.Z, elem.ID, taxonomyIdStr)
	}
	result.Success = true
	return result, err
}

func (d MySQLDriver) SaveMajor3DPositions(taxonomyId int64, positions []model.Position) (result model.Result, err error) {
	dbRef, err := d.OpenDB()
	defer dbRef.Close()
	checkErr(err)
	taxonomyIdStr := strconv.Itoa(int(taxonomyId))
	for _, elem := range positions {
		dbRef.Exec("update "+elem.Table+" set xMajor3D = ?, yMajor3D = ?, zMajor3D = ? where BINARY text = ? and id_taxonomy = ?;", elem.X, elem.Y, elem.Z, elem.ID, taxonomyIdStr)
	}
	result.Success = true
	return result, err
}

func (d MySQLDriver) SaveEdgeBendPoints(taxonomyId int64, attributeSrc string, attributeDest string, edgeBendPoints string, dimension string) (result model.Result, err error) {
	dbRef, err := d.OpenDB()
	defer dbRef.Close()
	checkErr(err)
	taxonomyIdStr := strconv.Itoa(int(taxonomyId))
	dbRef.Exec("update taxonomy_relation set edgeBendPoints = ? where id_taxonomy = ? and id_src_attribute = (select distinct id_attribute from attribute where BINARY text = ? and id_taxonomy = ?) and id_dest_attribute = (select distinct id_attribute from attribute where BINARY text = ? and id_taxonomy = ?) and id_dimension = (select distinct id_dimension from dimension where BINARY text = ? and id_taxonomy = ?);", edgeBendPoints, taxonomyIdStr, attributeSrc, taxonomyIdStr, attributeDest, taxonomyIdStr, dimension, taxonomyIdStr)
	result.Success = true
	return result, err
}

func (d MySQLDriver) RemoveTaxonomyRelationsForAttribute(taxonomyId int64, attribute model.Attribute) (result model.Result, err error) {
	dbRef, err := d.OpenDB()
	defer dbRef.Close()
	checkErr(err)
	taxonomyIdStr := strconv.Itoa(int(taxonomyId))
	attributeIDStr := strconv.Itoa(attribute.ID)
	dbRef.Exec("DELETE FROM taxonomy_relation_annotation WHERE id_taxonomy_relation IN (SELECT id_taxonomy_relation FROM taxonomy_relation WHERE id_taxonomy = ? AND (id_src_attribute = ? OR id_dest_attribute = ?));", taxonomyIdStr, attributeIDStr, attributeIDStr)
	dbRef.Exec("DELETE FROM taxonomy_relation WHERE id_taxonomy = ? AND (id_src_attribute = ? OR id_dest_attribute = ?);", taxonomyIdStr, attributeIDStr, attributeIDStr)
	go d.UpdateRelationshipTables(taxonomyId)
	fmt.Println("Returning RemoveTaxonomyRelationsForAttribute")
	result.Success = true
	return result, err
}

func (d MySQLDriver) UpdateChildRelationshipTable(taxonomyId int64) (resultChan chan model.Result) {
	dbRef, err := d.OpenDB()
	checkErr(err)
	taxonomyIdStr := strconv.Itoa(int(taxonomyId))
	future := make(chan model.Result)
	go func() {
		dbRef.Exec("CALL insertallchildrenperattribute(?);", taxonomyIdStr)
		defer dbRef.Close()
		future <- model.Result{Success: true}
	}()
	return future
}

func (d MySQLDriver) UpdateParentRelationshipTable(taxonomyId int64) (resultChan chan model.Result) {
	dbRef, err := d.OpenDB()
	checkErr(err)
	taxonomyIdStr := strconv.Itoa(int(taxonomyId))
	future := make(chan model.Result)
	go func() {
		dbRef.Exec("CALL insertallparentsperattribute(?);", taxonomyIdStr)
		defer dbRef.Close()
		future <- model.Result{Success: true}
	}()
	return future
}

func (d MySQLDriver) UpdateRelationshipTables(taxonomyId int64) (result model.Result) {
	fmt.Println("Starting UpdateRelationshipTables")
	resultFuture1 := d.UpdateParentRelationshipTable(taxonomyId)
	resultFuture2 := d.UpdateChildRelationshipTable(taxonomyId)
	result1 := <-resultFuture1
	result2 := <-resultFuture2
	result.Success = result1.Success && result2.Success
	fmt.Println("Returning UpdateRelationshipTables")
	return result
}

func (d MySQLDriver) AddAttribute(taxonomyId int64, attribute model.Attribute) (result model.Result, err error) {
	dbRef, err := d.OpenDB()
	defer dbRef.Close()
	checkErr(err)
	taxonomyIdStr := strconv.Itoa(int(taxonomyId))
	db, stmt, err := d.Query("SELECT count(id_attribute) FROM attribute WHERE BINARY text = ? AND id_taxonomy = ?;")
	defer stmt.Close()
	defer db.Close()
	rows, err := stmt.Query(attribute.Text, taxonomyIdStr)
	checkErr(err)
	var rowCount int
	for rows.Next() {
		rows.Scan(&rowCount)
	}
	defer rows.Close()
	if rowCount > 0 {
		result.Success = false
		return result, err
	}
	majorStr := strconv.Itoa(int(attribute.Major))
	dbRef.Exec("INSERT IGNORE INTO attribute (id_taxonomy, text, x, y, xMajor, yMajor, major) VALUES (?, ?, ?, ?, ?, ?, ?);", taxonomyIdStr, attribute.Text, attribute.X, attribute.Y, attribute.XMajor, attribute.YMajor, majorStr)
	dbRef.Exec("INSERT IGNORE INTO taxonomy_dimension (id_taxonomy, id_attribute, id_dimension) VALUES (?, (SELECT DISTINCT id_attribute FROM attribute WHERE BINARY text = ? AND id_taxonomy = ?), (SELECT DISTINCT id_dimension FROM dimension WHERE BINARY text = ? AND id_taxonomy = ?));", taxonomyIdStr, attribute.Text, taxonomyIdStr, attribute.Dimension, taxonomyIdStr)
	go d.UpdateRelationshipTables(taxonomyId)
	fmt.Println("Returning addAttribute")
	result.Success = true
	return result, err
}

func (d MySQLDriver) AddDimension(taxonomyId int64, dimension string) (result model.Result, err error) {
	dbRef, err := d.OpenDB()
	defer dbRef.Close()
	checkErr(err)
	if !strings.Contains(dimension, " view") {
		dimension += " view"
	}
	taxonomyIdStr := strconv.Itoa(int(taxonomyId))
	db, stmt, err := d.Query("SELECT count(id_dimension) FROM dimension WHERE BINARY text = ? AND id_taxonomy = ?;")
	defer stmt.Close()
	defer db.Close()
	rows, err := stmt.Query(dimension, taxonomyIdStr)
	checkErr(err)
	var rowCount int
	for rows.Next() {
		rows.Scan(&rowCount)
	}
	defer rows.Close()
	if rowCount > 0 {
		result.Success = false
		return result, err
	}
	dbRef.Exec("INSERT IGNORE INTO dimension (id_taxonomy, text) VALUES (?, ?);", taxonomyIdStr, dimension)
	result.Success = true
	return result, err
}

func (d MySQLDriver) ChangeDimension(taxonomyId int64, attribute string, dimension string) (result model.Result, err error) {
	dbRef, err := d.OpenDB()
	defer dbRef.Close()
	checkErr(err)
	taxonomyIdStr := strconv.Itoa(int(taxonomyId))
	dbRef.Exec("UPDATE taxonomy_dimension SET id_dimension = (SELECT DISTINCT id_dimension FROM dimension WHERE BINARY text = ? AND id_taxonomy = ?) WHERE id_taxonomy = ? AND id_attribute = (SELECT DISTINCT id_attribute FROM attribute WHERE BINARY text = ? AND id_taxonomy = ?);", dimension, taxonomyIdStr, taxonomyIdStr, attribute, taxonomyIdStr)
	dbRef.Exec("DELETE FROM taxonomy_relation WHERE id_taxonomy = ? AND (id_src_attribute = (SELECT DISTINCT id_attribute FROM attribute WHERE BINARY text = ? AND id_taxonomy = ?) OR id_dest_attribute = (SELECT DISTINCT id_attribute FROM attribute WHERE BINARY text = ? AND id_taxonomy = ?));", taxonomyIdStr, attribute, taxonomyIdStr, attribute, taxonomyIdStr)
	d.UpdateRelationshipTables(taxonomyId)
	result.Success = true
	return result, err
}

func (d MySQLDriver) RenameAttribute(taxonomyId int64, previousName string, newName string) (result model.Result, err error) {
	dbRef, err := d.OpenDB()
	defer dbRef.Close()
	checkErr(err)
	taxonomyIdStr := strconv.Itoa(int(taxonomyId))
	db, stmt, err := d.Query("SELECT count(id_attribute) FROM attribute WHERE BINARY text = ? AND id_taxonomy = ?;")
	defer stmt.Close()
	defer db.Close()
	rows, err := stmt.Query(newName, taxonomyIdStr)
	checkErr(err)
	var rowCount int
	for rows.Next() {
		rows.Scan(&rowCount)
	}
	defer rows.Close()
	if rowCount > 0 {
		result.Success = false
		return result, err
	}
	dbRef.Exec("UPDATE attribute SET BINARY text = ? WHERE BINARY text = ? AND id_taxonomy = ?;", newName, previousName, taxonomyIdStr)
	d.UpdateRelationshipTables(taxonomyId)
	result.Success = true
	return result, err
}

func (d MySQLDriver) UpdateSynonyms(taxonomyId int64, attribute string, synonyms string) (result model.Result, err error) {
	dbRef, err := d.OpenDB()
	defer dbRef.Close()
	checkErr(err)
	taxonomyIdStr := strconv.Itoa(int(taxonomyId))
	dbRef.Exec("UPDATE attribute SET synonyms = ? WHERE BINARY text = ? AND id_taxonomy = ?;", synonyms, attribute, taxonomyIdStr)
	result.Success = true
	return result, err
}

func (d MySQLDriver) RenameDimension(taxonomyId int64, previousName string, newName string) (result model.Result, err error) {
	dbRef, err := d.OpenDB()
	defer dbRef.Close()
	checkErr(err)
	taxonomyIdStr := strconv.Itoa(int(taxonomyId))
	db, stmt, err := d.Query("SELECT count(id_dimension) FROM dimension WHERE BINARY text = ? AND id_taxonomy = ?;")
	defer stmt.Close()
	defer db.Close()
	rows, err := stmt.Query(newName, taxonomyIdStr)
	checkErr(err)
	var rowCount int
	for rows.Next() {
		rows.Scan(&rowCount)
	}
	defer rows.Close()
	if rowCount > 0 {
		result.Success = false
		return result, err
	}
	dbRef.Exec("UPDATE dimension SET BINARY text = ? WHERE BINARY text = ? AND id_taxonomy = ?;", newName, previousName, taxonomyIdStr)
	result.Success = true
	return result, err
}

func (d MySQLDriver) MergeAttributes(taxonomyId int64, attribute1 model.Attribute, attribute2 model.Attribute) (result model.Result, err error) {
	dbRef, err := d.OpenDB()
	defer dbRef.Close()
	checkErr(err)
	taxonomyIdStr := strconv.Itoa(int(taxonomyId))
	db, stmt, err := d.Query("SELECT count(id_attribute) FROM attribute WHERE BINARY text = ? AND id_taxonomy = ?;")
	defer stmt.Close()
	defer db.Close()
	rows, err := stmt.Query(attribute1.Text+":"+attribute2.Text, taxonomyIdStr)
	checkErr(err)
	var rowCount int
	for rows.Next() {
		rows.Scan(&rowCount)
	}
	defer rows.Close()
	if rowCount > 0 {
		result.Success = false
		return result, err
	}
	dbRef.Exec("INSERT IGNORE INTO attribute (id_taxonomy, text, x, y, xMajor, yMajor, major) SELECT attribute.id_taxonomy, ? as newName, attribute.x, attribute.y, attribute.xMajor, attribute.yMajor, attribute.major FROM attribute WHERE attribute.text = ? AND attribute.id_taxonomy = ?;", attribute1.Text+":"+attribute2.Text, attribute1.Text, taxonomyIdStr)
	dbRef.Exec("INSERT IGNORE INTO taxonomy_dimension (id_taxonomy, id_attribute, id_dimension) VALUES (?, (SELECT DISTINCT id_attribute FROM attribute WHERE BINARY text = ? AND id_taxonomy = ?), (SELECT DISTINCT id_dimension FROM dimension WHERE BINARY text = ? AND id_taxonomy = ?));", taxonomyIdStr, attribute1.Text+":"+attribute2.Text, taxonomyIdStr, attribute1.Dimension, taxonomyIdStr)
	dbRef.Exec("INSERT IGNORE INTO taxonomy_relation (id_taxonomy, id_src_attribute, id_dest_attribute, id_relation, id_dimension) SELECT id_taxonomy, (SELECT DISTINCT id_attribute FROM attribute WHERE BINARY text = ? AND id_taxonomy = ?), id_dest_attribute, id_relation, id_dimension FROM taxonomy_relation WHERE id_src_attribute = (SELECT id_attribute FROM attribute WHERE BINARY text = ? AND id_taxonomy = ?);", attribute1.Text+":"+attribute2.Text, taxonomyIdStr, attribute1.Text, taxonomyIdStr)
	dbRef.Exec("INSERT IGNORE INTO taxonomy_relation (id_taxonomy, id_src_attribute, id_dest_attribute, id_relation, id_dimension) SELECT id_taxonomy, (SELECT DISTINCT id_attribute FROM attribute WHERE BINARY text = ? AND id_taxonomy = ?), id_dest_attribute, id_relation, id_dimension FROM taxonomy_relation WHERE id_src_attribute = (SELECT id_attribute FROM attribute WHERE BINARY text = ? AND id_taxonomy = ?);", attribute1.Text+":"+attribute2.Text, taxonomyIdStr, attribute2.Text, taxonomyIdStr)
	dbRef.Exec("INSERT IGNORE INTO taxonomy_relation (id_taxonomy, id_src_attribute, id_dest_attribute, id_relation, id_dimension) SELECT id_taxonomy, id_src_attribute, (SELECT DISTINCT id_attribute FROM attribute WHERE BINARY text = ? AND id_taxonomy = ?), id_relation, id_dimension FROM taxonomy_relation WHERE id_dest_attribute = (SELECT id_attribute FROM attribute WHERE BINARY text = ? AND id_taxonomy = ?);", attribute1.Text+":"+attribute2.Text, taxonomyIdStr, attribute1.Text, taxonomyIdStr)
	dbRef.Exec("INSERT IGNORE INTO taxonomy_relation (id_taxonomy, id_src_attribute, id_dest_attribute, id_relation, id_dimension) SELECT id_taxonomy, id_src_attribute, (SELECT DISTINCT id_attribute FROM attribute WHERE BINARY text = ? AND id_taxonomy = ?), id_relation, id_dimension FROM taxonomy_relation WHERE id_dest_attribute = (SELECT id_attribute FROM attribute WHERE BINARY text = ? AND id_taxonomy = ?);", attribute1.Text+":"+attribute2.Text, taxonomyIdStr, attribute2.Text, taxonomyIdStr)
	dbRef.Exec("INSERT IGNORE INTO mapping (id_paper, id_attribute) SELECT id_paper, (SELECT DISTINCT id_attribute FROM attribute WHERE BINARY text = ? AND id_taxonomy = ?) FROM mapping WHERE id_attribute = (SELECT id_attribute FROM attribute WHERE BINARY text = ? AND id_taxonomy = ?);", attribute1.Text+":"+attribute2.Text, taxonomyIdStr, attribute1.Text, taxonomyIdStr)
	dbRef.Exec("INSERT IGNORE INTO mapping (id_paper, id_attribute) SELECT id_paper, (SELECT DISTINCT id_attribute FROM attribute WHERE BINARY text = ? AND id_taxonomy = ?) FROM mapping WHERE id_attribute = (SELECT id_attribute FROM attribute WHERE BINARY text = ? AND id_taxonomy = ?);", attribute1.Text+":"+attribute2.Text, taxonomyIdStr, attribute2.Text, taxonomyIdStr)
	dbRef.Exec("DELETE FROM attribute WHERE id_taxonomy = ? AND (text = ? OR BINARY text = ?);", taxonomyIdStr, attribute1.Text, attribute2.Text)
	d.UpdateRelationshipTables(taxonomyId)
	result.Success = true
	return result, err
}

func (d MySQLDriver) ForkAttribute(taxonomyId int64, attribute string, dimension string, parents1 []model.AttributeRelation, parents2 []model.AttributeRelation, children1 []model.AttributeRelation, children2 []model.AttributeRelation, citations1 []model.Paper, citations2 []model.Paper) (result model.Result, err error) {
	dbRef, err := d.OpenDB()
	defer dbRef.Close()
	checkErr(err)
	taxonomyIdStr := strconv.Itoa(int(taxonomyId))
	counter := 2
	var input int
	var input2 int
	var newAttributeName string
	var newAttributeName2 string
	for ok := true; ok; ok = (input != 2) {
		newAttributeName = attribute + ":" + strconv.Itoa(counter)
		db, stmt, err := d.Query("SELECT count(id_attribute) FROM attribute WHERE BINARY text = ? AND id_taxonomy = ?;")
		rows, err := stmt.Query(newAttributeName, taxonomyIdStr)
		stmt.Close()
		db.Close()
		checkErr(err)
		var rowCount int
		for rows.Next() {
			rows.Scan(&rowCount)
		}
		rows.Close()
		if rowCount > 0 {
			input = 1
		} else {
			input = 2
		}
		counter++
	}
	for ok := true; ok; ok = (input2 != 2) {
		newAttributeName2 = attribute + ":" + strconv.Itoa(counter)
		db, stmt, err := d.Query("SELECT count(id_attribute) FROM attribute WHERE BINARY text = ? AND id_taxonomy = ?;")
		rows, err := stmt.Query(newAttributeName, taxonomyIdStr)
		stmt.Close()
		db.Close()
		checkErr(err)
		var rowCount int
		for rows.Next() {
			rows.Scan(&rowCount)
		}
		rows.Close()
		if rowCount > 0 {
			input2 = 1
		} else {
			input2 = 2
		}
		counter++
	}
	dbRef.Exec("INSERT IGNORE INTO attribute (id_taxonomy, text, x, y, xMajor, yMajor, major) SELECT attribute.id_taxonomy, ? as newName, attribute.x, attribute.y, attribute.xMajor, attribute.yMajor, attribute.major FROM attribute WHERE attribute.text = ? AND attribute.id_taxonomy = ?;", newAttributeName, attribute, taxonomyIdStr)
	dbRef.Exec("INSERT IGNORE INTO taxonomy_dimension (id_taxonomy, id_attribute, id_dimension) VALUES (?, (SELECT DISTINCT id_attribute FROM attribute WHERE BINARY text = ? AND id_taxonomy = ?), (SELECT DISTINCT id_dimension from dimension where BINARY text = ? AND id_taxonomy = ?));", taxonomyIdStr, newAttributeName, taxonomyIdStr, dimension, taxonomyIdStr)
	dbRef.Exec("INSERT IGNORE INTO attribute (id_taxonomy, text, x, y, xMajor, yMajor, major) SELECT attribute.id_taxonomy, ? as newName, attribute.x + 150, attribute.y, attribute.xMajor + 150, attribute.yMajor, attribute.major FROM attribute WHERE attribute.text = ? AND attribute.id_taxonomy = ?;", newAttributeName2, attribute, taxonomyIdStr)
	dbRef.Exec("INSERT IGNORE INTO taxonomy_dimension (id_taxonomy, id_attribute, id_dimension) VALUES (?, (SELECT DISTINCT id_attribute FROM attribute WHERE BINARY text = ? AND id_taxonomy = ?), (SELECT DISTINCT id_dimension from dimension where BINARY text = ? AND id_taxonomy = ?));", taxonomyIdStr, newAttributeName2, taxonomyIdStr, dimension, taxonomyIdStr)
	for _, parent := range parents1 {
		dbRef.Exec("INSERT IGNORE INTO taxonomy_relation (id_taxonomy, id_src_attribute, id_dest_attribute, id_relation, id_dimension) VALUES (?, (select distinct id_attribute from attribute where BINARY text = ? AND id_taxonomy = ?), (select distinct id_attribute from attribute where BINARY text = ? AND id_taxonomy = ?), (select distinct id_relation from relation where BINARY text = ?), (select distinct id_dimension from dimension where BINARY text = ? AND id_taxonomy = ?));", taxonomyIdStr, newAttributeName, taxonomyIdStr, parent.Text, taxonomyIdStr, parent.Relation, dimension, taxonomyIdStr)
	}
	for _, parent := range parents2 {
		dbRef.Exec("INSERT IGNORE INTO taxonomy_relation (id_taxonomy, id_src_attribute, id_dest_attribute, id_relation, id_dimension) VALUES (?, (select distinct id_attribute from attribute where BINARY text = ? AND id_taxonomy = ?), (select distinct id_attribute from attribute where BINARY text = ? AND id_taxonomy = ?), (select distinct id_relation from relation where BINARY text = ?), (select distinct id_dimension from dimension where BINARY text = ? AND id_taxonomy = ?));", taxonomyIdStr, newAttributeName2, taxonomyIdStr, parent.Text, taxonomyIdStr, parent.Relation, dimension, taxonomyIdStr)
	}
	for _, child := range children1 {
		dbRef.Exec("INSERT IGNORE INTO taxonomy_relation (id_taxonomy, id_dest_attribute, id_src_attribute, id_relation, id_dimension) VALUES (?, (select distinct id_attribute from attribute where BINARY text = ? AND id_taxonomy = ?), (select distinct id_attribute from attribute where BINARY text = ? AND id_taxonomy = ?), (select distinct id_relation from relation where BINARY text = ?), (select distinct id_dimension from dimension where BINARY text = ? AND id_taxonomy = ?));", taxonomyIdStr, newAttributeName, taxonomyIdStr, child.Text, taxonomyIdStr, child.Relation, dimension, taxonomyIdStr)
	}
	for _, child := range children2 {
		dbRef.Exec("INSERT IGNORE INTO taxonomy_relation (id_taxonomy, id_dest_attribute, id_src_attribute, id_relation, id_dimension) VALUES (?, (select distinct id_attribute from attribute where BINARY text = ? AND id_taxonomy = ?), (select distinct id_attribute from attribute where BINARY text = ? AND id_taxonomy = ?), (select distinct id_relation from relation where BINARY text = ?), (select distinct id_dimension from dimension where BINARY text = ? AND id_taxonomy = ?));", taxonomyIdStr, newAttributeName2, taxonomyIdStr, child.Text, taxonomyIdStr, child.Relation, dimension, taxonomyIdStr)
	}
	for _, citation := range citations1 {
		dbRef.Exec("INSERT IGNORE INTO mapping (id_paper, id_attribute) VALUES (?, (select distinct id_attribute from attribute where BINARY text = ? AND id_taxonomy = ?));", strconv.Itoa(int(citation.ID)), newAttributeName, taxonomyIdStr)
	}
	for _, citation := range citations2 {
		dbRef.Exec("INSERT IGNORE INTO mapping (id_paper, id_attribute) VALUES (?, (select distinct id_attribute from attribute where BINARY text = ? AND id_taxonomy = ?));", strconv.Itoa(int(citation.ID)), newAttributeName2, taxonomyIdStr)
	}
	dbRef.Exec("DELETE FROM attribute WHERE BINARY text = ? AND id_taxonomy = ?;", attribute, taxonomyIdStr)
	d.UpdateRelationshipTables(taxonomyId)
	result.Success = true
	return result, err
}

func (d MySQLDriver) CheckIfRelationIsValid(relation model.AttributeRelation) (result bool) {
	dbRef, err := d.OpenDB()
	defer dbRef.Close()
	checkErr(err)
	taxonomyIdStr := strconv.Itoa(int(relation.TaxonomyID))
	db, stmt, err := d.Query("SELECT id_relation FROM relation WHERE BINARY text = ?;")
	checkErr(err)
	defer stmt.Close()
	defer db.Close()
	rows, err := stmt.Query(relation.Relation)
	checkErr(err)
	var relationID int
	relationID = -1
	for rows.Next() {
		rows.Scan(&relationID)
	}
	defer rows.Close()
	if relationID < 0 {
		return false
	}
	if relationID < 3 {
		return true
	}
	db, stmt, err = d.Query("SELECT parents FROM allparentsperattribute WHERE BINARY text = ? AND id_taxonomy = ?;")
	checkErr(err)
	defer stmt.Close()
	defer db.Close()
	rows, err = stmt.Query(relation.AttributeDest, taxonomyIdStr)
	checkErr(err)
	var parents string
	parents = ""
	for rows.Next() {
		rows.Scan(&parents)
	}
	defer rows.Close()
	array := strings.Split(parents, ",")
	for _, elem := range array {
		if elem == relation.AttributeSrc {
			return false
		}
	}
	db, stmt, err = d.Query("SELECT parents FROM allparentsperattribute WHERE BINARY text = ? AND id_taxonomy = ?;")
	checkErr(err)
	defer stmt.Close()
	defer db.Close()
	rows, err = stmt.Query(relation.AttributeSrc, taxonomyIdStr)
	checkErr(err)
	parents = ""
	for rows.Next() {
		rows.Scan(&parents)
	}
	defer rows.Close()
	array = strings.Split(parents, ",")
	for _, elem := range array {
		if elem == relation.AttributeDest {
			return false
		}
	}
	return true
}

func (d MySQLDriver) AddTaxonomyRelation(relation model.AttributeRelation) (result model.Result, err error) {
	dbRef, err := d.OpenDB()
	defer dbRef.Close()
	checkErr(err)
	taxonomyIdStr := strconv.Itoa(int(relation.TaxonomyID))
	isValid := d.CheckIfRelationIsValid(relation)
	if !isValid {
		result.Success = false
		return result, err
	}
	dbRef.Exec("INSERT IGNORE INTO taxonomy_relation (id_taxonomy, id_src_attribute, id_dest_attribute, id_relation, id_dimension) VALUES (?, (SELECT id_attribute FROM attribute WHERE BINARY text = ? AND id_taxonomy = ?), (SELECT id_attribute FROM attribute WHERE BINARY text = ? AND id_taxonomy = ?), (SELECT id_relation FROM relation WHERE BINARY text = ?), (SELECT id_dimension FROM dimension WHERE BINARY text = ? AND id_taxonomy = ?));", taxonomyIdStr, relation.AttributeSrc, taxonomyIdStr, relation.AttributeDest, taxonomyIdStr, relation.Relation, relation.Dimension, taxonomyIdStr)
	go d.UpdateRelationshipTables(relation.TaxonomyID)
	fmt.Println("Returning AddTaxonomyRelation")
	result.Success = true
	return result, err
}

func (d MySQLDriver) RemoveTaxonomyRelation(relation model.AttributeRelation) (result model.Result, err error) {
	dbRef, err := d.OpenDB()
	defer dbRef.Close()
	checkErr(err)
	taxonomyIdStr := strconv.Itoa(int(relation.TaxonomyID))
	dbRef.Exec("DELETE FROM taxonomy_relation_annotation WHERE id_taxonomy_relation IN (SELECT id_taxonomy_relation FROM taxonomy_relation WHERE id_taxonomy = ? AND id_src_attribute = (SELECT id_attribute FROM attribute WHERE BINARY text = ? AND id_taxonomy = ?) AND id_dest_attribute = (SELECT id_attribute FROM attribute WHERE BINARY text = ? AND id_taxonomy = ?) AND id_relation = (SELECT id_relation FROM relation WHERE BINARY text = ?) AND id_dimension = (SELECT id_dimension FROM dimension WHERE BINARY text = ? AND id_taxonomy = ?);", taxonomyIdStr, relation.AttributeSrc, taxonomyIdStr, relation.AttributeDest, taxonomyIdStr, relation.Relation, relation.Dimension, taxonomyIdStr)
	dbRef.Exec("DELETE FROM taxonomy_relation WHERE id_taxonomy = ? AND id_src_attribute = (SELECT id_attribute FROM attribute WHERE BINARY text = ? AND id_taxonomy = ?) AND id_dest_attribute = (SELECT id_attribute FROM attribute WHERE BINARY text = ? AND id_taxonomy = ?) AND id_relation = (SELECT id_relation FROM relation WHERE BINARY text = ?) AND id_dimension = (SELECT id_dimension FROM dimension WHERE BINARY text = ? AND id_taxonomy = ?);", taxonomyIdStr, relation.AttributeSrc, taxonomyIdStr, relation.AttributeDest, taxonomyIdStr, relation.Relation, relation.Dimension, taxonomyIdStr)
	go d.UpdateRelationshipTables(relation.TaxonomyID)
	fmt.Println("Returning RemoveTaxonomyRelation")
	result.Success = true
	return result, err
}

func (d MySQLDriver) UpdateTaxonomyRelationType(relation model.AttributeRelation) (result model.Result, err error) {
	dbRef, err := d.OpenDB()
	defer dbRef.Close()
	checkErr(err)
	taxonomyIdStr := strconv.Itoa(int(relation.TaxonomyID))
	isValid := d.CheckIfRelationIsValid(relation)
	if !isValid {
		result.Success = false
		return result, err
	}
	dbRef.Exec("UPDATE taxonomy_relation SET id_relation = (SELECT id_relation FROM relation WHERE BINARY text = ?) WHERE id_taxonomy = ? AND id_src_attribute = (SELECT id_attribute FROM attribute WHERE BINARY text = ? AND id_taxonomy = ?) AND id_dest_attribute = (SELECT id_attribute FROM attribute WHERE BINARY text = ? AND id_taxonomy = ?);", relation.Relation, taxonomyIdStr, relation.AttributeSrc, taxonomyIdStr, relation.AttributeDest, taxonomyIdStr) //  AND id_dimension = (SELECT id_dimension FROM dimension WHERE BINARY text = \"" + relation.Dimension + "\" AND id_taxonomy = " + taxonomyIdStr + ")
	go func() {
		d.UpdateRelationshipTables(relation.TaxonomyID)
	}()
	result.Success = true
	return result, err
}

func (d MySQLDriver) UpdateTaxonomyRelationAnnotation(relation model.AttributeRelation) (result model.Result, err error) {
	dbRef, err := d.OpenDB()
	defer dbRef.Close()
	checkErr(err)
	taxonomyIdStr := strconv.Itoa(int(relation.TaxonomyID))
	dbRef.Exec("REPLACE INTO taxonomy_relation_annotation (id_taxonomy, id_taxonomy_relation, annotation) VALUES (?, (SELECT DISTINCT id_taxonomy_relation FROM taxonomy_relation WHERE id_taxonomy = ? AND id_src_attribute = (SELECT id_attribute FROM attribute WHERE BINARY text = ? AND id_taxonomy = ?) AND id_dest_attribute = (SELECT id_attribute FROM attribute WHERE BINARY text = ? AND id_taxonomy = ?)), ?);", taxonomyIdStr, taxonomyIdStr, relation.AttributeSrc, taxonomyIdStr, relation.AttributeDest, taxonomyIdStr, relation.Annotation) //  AND id_dimension = (SELECT id_dimension FROM dimension WHERE BINARY text = \"" + relation.Dimension + "\" AND id_taxonomy = " + taxonomyIdStr + ")
	result.Success = true
	return result, err
}

func (d MySQLDriver) DeleteCitation(taxonomyId int64, paper model.Paper) (result model.Result, err error) {
	dbRef, err := d.OpenDB()
	defer dbRef.Close()
	checkErr(err)
	taxonomyIdStr := strconv.Itoa(int(taxonomyId))
	dbRef.Exec("DELETE FROM paper WHERE id_taxonomy = ? AND BINARY citation = ?;", taxonomyIdStr, paper.Citation)
	result.Success = true
	return result, err
}

func (d MySQLDriver) RemoveAttribute(taxonomyId int64, attribute model.Attribute) (result model.Result, err error) {
	dbRef, err := d.OpenDB()
	defer dbRef.Close()
	checkErr(err)
	taxonomyIdStr := strconv.Itoa(int(taxonomyId))
	db, stmt, err := d.Query("SELECT id_attribute FROM attribute WHERE BINARY text = ? AND id_taxonomy = ?;")
	checkErr(err)
	defer stmt.Close()
	defer db.Close()
	rows, err := stmt.Query(attribute.Text, taxonomyIdStr)
	checkErr(err)
	var attributeID int
	attributeID = -1
	for rows.Next() {
		rows.Scan(&attributeID)
	}
	defer rows.Close()
	if attributeID < 0 {
		result.Success = false
		return result, err
	}
	attributeIDStr := strconv.Itoa(attributeID)
	dbRef.Exec("DELETE FROM attribute WHERE id_attribute = ?;", attributeIDStr)
	dbRef.Exec("DELETE FROM taxonomy_dimension WHERE id_attribute = ?;", attributeIDStr)
	dbRef.Exec("DELETE FROM mapping WHERE id_attribute = ?;", attributeIDStr)
	d.RemoveTaxonomyRelationsForAttribute(taxonomyId, attribute)
	result.Success = true
	return result, err
}

func (d MySQLDriver) RemoveDimension(taxonomyId int64, dimension model.Dimension) (result model.Result, err error) {
	dbRef, err := d.OpenDB()
	defer dbRef.Close()
	checkErr(err)
	taxonomyIdStr := strconv.Itoa(int(taxonomyId))
	db, stmt, err := d.Query("SELECT id_dimension FROM dimension WHERE BINARY text = ? AND id_taxonomy = ?;")
	checkErr(err)
	defer stmt.Close()
	defer db.Close()
	rows, err := stmt.Query(dimension.Text, taxonomyIdStr)
	checkErr(err)
	var dimensionID int
	dimensionID = -1
	for rows.Next() {
		rows.Scan(&dimensionID)
	}
	defer rows.Close()
	if dimensionID < 0 {
		result.Success = false
		return result, err
	}
	dimensionIDStr := strconv.Itoa(dimensionID)
	db2, stmt2, err2 := d.Query("SELECT count(id_attribute) FROM taxonomy_dimension WHERE id_dimension = ?;")
	defer stmt2.Close()
	defer db2.Close()
	rows2, err2 := stmt2.Query(dimensionIDStr)
	checkErr(err2)
	var rowCount int
	for rows2.Next() {
		rows2.Scan(&rowCount)
	}
	defer rows2.Close()
	if rowCount > 0 {
		result.Success = false
		return result, err2
	}
	dbRef.Exec("DELETE FROM dimension WHERE id_dimension = ?;", dimensionIDStr)
	result.Success = true
	return result, err
}