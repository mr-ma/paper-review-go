package data
import (
	"strconv"
	"strings"
	"../gabs"
	"sort"
	"fmt"
	"math/rand"
	//overriding MySqlDriver
	_ "../mysql"
		"../model"
)

type ClassificationDriver interface {
  DriverCore
  	Login(string, string) (model.LoginResult, error)
  	SaveUser(string, string) (model.Result, error)
  	CreateUser(string, string, int) (model.Result, error)
  	UpdateUser(string, int) (model.Result, error)
  	DeleteUser(string) (model.Result, error)
  	GetUser(string) (model.User, error)
  	GetUsers() ([]model.User, error)
	ExportCorrelations([]model.Attribute, int64) ([]model.Paper, error)
	GetAllTaxonomies() ([]model.Taxonomy, error)
	GetTaxonomyID(string) ([]model.Taxonomy, error)
	GetTaxonomyPermissions(string) ([]model.Taxonomy, error)
	UpdateTaxonomyPermissions(string, string) (model.Result, error)
	AddTaxonomy(string, string) (model.Result, error)
	RemoveTaxonomy(int64) (model.Result, error)
	GetAttributesPerDimension(int64, string) ([]model.Attribute, error)
	GetLeafAttributesPerDimension(int64, string) ([]model.Attribute, error)
	GetAttributeChildren(string, string, model.AttributeCluster, model.AttributeCluster) (string)
	GetAttributeClusterPerDimension(int64, string) (string, error)
	GetAllChildrenAttributes(int64, string) ([]model.Attribute, error)
	GetAllChildrenLeafAttributes(int64, string) ([]model.Attribute, error)
	GetIntermediateAttributes(int64, int64, int64) ([]model.Attribute, error)
	GetMajorAttributes(int64) ([]model.Attribute, error)
	GetAttributeRelationsPerDimension(int64, string) ([]model.AttributeRelation, error)
	GetInterdimensionalRelations(int64) ([]model.AttributeRelation, error)
	GetCitationCount(int64) ([]model.CitationCount, error)
	GetCitationCounts(int64) ([]model.CitationCount, error)
	GetCitationCountsIncludingChildren(int64) ([]model.CitationCount, error)
	UpdateCitationReferenceCounts(int64, []model.ReferenceCount) (model.Result, error)
	UpdateMajor(int64, string, int) (model.Result, error)
	UpdateCitationMapping(int64, string, []model.Paper) (model.Result, error)
	UpdateCitationMappings(int64, []model.CitationMapping) (model.Result, error)
	GetRelationTypes() ([]model.RelationType, error)
	GetCitationsPerAttribute(int64, string) ([]model.Paper, error)
	GetCitationsPerAttributeIncludingChildren(int64, string) ([]model.Paper, error)
	SavePositions(int64, []model.Position) (model.Result, error)
	SaveMajorPositions(int64, []model.Position) (model.Result, error)
	Save3DPositions(int64, []model.Position) (model.Result, error)
	SaveMajor3DPositions(int64, []model.Position) (model.Result, error)
	SaveEdgeBendPoints(int64, string, string, string, string) (model.Result, error)
	GetAllAttributes(int64) ([]model.Attribute, error)
	GetLeafAttributes(int64) ([]model.Attribute, error)
	GetAllDimensions(int64) ([]model.Dimension, error)
	GetAllCitations(int64) ([]model.Paper, error)
	GetConceptRelations(int64) ([]model.ConceptCorrelation, error)
	GetConceptRelations3D(int64) ([]model.ConceptCorrelation, error)
	GetConceptRelationsWithReferenceCounts(int64) ([]model.ConceptCorrelation, error)
	GetConceptRelationsWithReferenceCounts3D(int64) ([]model.ConceptCorrelation, error)
	GetAllConceptRelations(int64) ([]model.ConceptCorrelation, error)
	GetAllConceptRelations3D(int64) ([]model.ConceptCorrelation, error)
	GetAllConceptRelationsWithReferenceCounts(int64) ([]model.ConceptCorrelation, error)
	GetAllConceptRelationsWithReferenceCounts3D(int64) ([]model.ConceptCorrelation, error)
	GetParentRelationsPerAttribute(int64, string, string) ([]model.AttributeRelation, error)
	GetChildRelationsPerAttribute(int64, string, string) ([]model.AttributeRelation, error)
	GetSharedPapers(int64, string, string) ([]model.Paper, error)
	GetSharedPapers3D(int64, string, string, string) ([]model.Paper, error)
	GetSharedPapersIncludingChildren(int64, string, string) ([]model.Paper, error)
	GetSharedPapersIncludingChildren3D(int64, string, string, string) ([]model.Paper, error)
	GetAttributeDetails(int64, string) ([]model.Attribute, error)
	GetCitationDetails(int64, string, string) ([]model.Paper, error)
	GetAttributeCoverage(int64) ([]model.AttributeCoverage, error)
	GetAttributeCoverageWithOcurrenceCounts(int64) ([]model.AttributeCoverage, error)
	GetAttributeCoverageWithReferenceCounts(int64) ([]model.AttributeCoverage, error)
	ExportCorrelationsCSV(filterAttributes []model.Attribute)
	AddAttribute(int64, model.Attribute) (model.Result, error)
	AddDimension(int64, string) (model.Result, error)
	ChangeDimension(int64, string, string) (model.Result, error)
	RenameAttribute(int64, string, string) (model.Result, error)
	UpdateSynonyms(int64, string, string) (model.Result, error)
	RenameDimension(int64, string, string) (model.Result, error)
	AddTaxonomyRelation(model.AttributeRelation) (model.Result, error)
	DeleteCitation(int64, model.Paper) (model.Result, error)
	RemoveAttribute(int64, model.Attribute) (model.Result, error)
	RemoveDimension(int64, model.Dimension) (model.Result, error)
	RemoveTaxonomyRelation(model.AttributeRelation) (model.Result, error)
	UpdateTaxonomyRelationType(model.AttributeRelation) (model.Result, error)
	UpdateTaxonomyRelationAnnotation(model.AttributeRelation) (model.Result, error)
	RemoveTaxonomyRelationsForAttribute(int64, model.Attribute) (model.Result, error)
	UpdateChildRelationshipTable(int64) (chan model.Result)
	UpdateParentRelationshipTable(int64) (chan model.Result)
	UpdateRelationshipTables(int64) (model.Result)
	MergeAttributes(int64, model.Attribute, model.Attribute) (model.Result, error)
	ForkAttribute(int64, string, string, []model.AttributeRelation, []model.AttributeRelation, []model.AttributeRelation, []model.AttributeRelation, []model.Paper, []model.Paper) (model.Result, error)
	SaveReviewMappings(int64, []model.Attribute, []model.ReviewMapping) (model.Result, error)
	KMeans(int64, int) ([]model.Cluster, error)
}


//InitMySQLDriver initialize a new my sql driver instance
func InitClassificationDriver(user string, password string) ClassificationDriver {
	return MySQLDriver{username: user, pass: password, database: "classification"}
}

func (d MySQLDriver) Login(email string, password string) (result model.LoginResult, err error){
	dbRef, err := d.OpenDB()
	defer dbRef.Close()
	checkErr(err)
	db, stmt, err := d.Query("SELECT email, taxonomies, admin FROM user WHERE email = ? AND password = ?;")
	defer stmt.Close()
	defer db.Close()
	rows, err := stmt.Query(email, password)
	checkErr(err)
	if rows.Next() {
		a := model.User{}
		rows.Scan(&a.Email,&a.Taxonomies,&a.Admin)
		result.Success = true
		result.User = a
	} else {
		result.Success = false
	}
	defer rows.Close()
	return result, err
	}

func (d MySQLDriver) SaveUser(email string, password string) (result model.Result, err error){
	dbRef, err := d.OpenDB()
	defer dbRef.Close()
	checkErr(err)
	db, stmt, err := d.Query("SELECT COUNT(email) as userCount FROM user WHERE email = ?;")
	defer stmt.Close()
	defer db.Close()
	rows, err := stmt.Query(email)
	checkErr(err)
	var count int
	for rows.Next() {
		rows.Scan(&count)
	}
	defer rows.Close()
	if count > 0 {
		result.Success = false
		return result, err
	}
	dbRef.Exec("INSERT IGNORE INTO user (email, name, password) VALUES (?, ?, ?);", email, email, password)
	result.Success = true
	return result, err
	}

func (d MySQLDriver) CreateUser(email string, password string, admin int) (result model.Result, err error){
	dbRef, err := d.OpenDB()
	defer dbRef.Close()
	checkErr(err)
	db, stmt, err := d.Query("SELECT COUNT(email) as userCount FROM user WHERE email = ?;")
	defer stmt.Close()
	defer db.Close()
	rows, err := stmt.Query(email)
	checkErr(err)
	var count int
	for rows.Next() {
		rows.Scan(&count)
	}
	defer rows.Close()
	if count > 0 {
		result.Success = false
		return result, err
	}
	adminStr := strconv.Itoa(admin)
	dbRef.Exec("INSERT IGNORE INTO user (email, name, password, admin) VALUES (?, ?, ?, ?);", email, email, password, adminStr)
	result.Success = true
	return result, err
	}

func (d MySQLDriver) UpdateUser(email string, admin int) (result model.Result, err error){
	dbRef, err := d.OpenDB()
	defer dbRef.Close()
	checkErr(err)
	adminStr := strconv.Itoa(admin)
	dbRef.Exec("UPDATE user SET admin = ? WHERE email = ?;", adminStr, email)
	result.Success = true
	return result, err
	}

func (d MySQLDriver) DeleteUser(email string) (result model.Result, err error){
	dbRef, err := d.OpenDB()
	defer dbRef.Close()
	checkErr(err)
	dbRef.Exec("DELETE FROM user WHERE email = ?;", email)
	result.Success = true
	return result, err
	}

func (d MySQLDriver) GetUser(email string) (user model.User, err error){
	dbRef, err := d.OpenDB()
	defer dbRef.Close()
	checkErr(err)
	db, stmt, err := d.Query("SELECT id, name FROM user WHERE email = ?;")
	defer stmt.Close()
	defer db.Close()
	rows, err := stmt.Query(email)
	checkErr(err)
	if rows.Next() {
		a := model.User{}
		rows.Scan(&a.ID,&a.Name)
		user = a
	}
	defer rows.Close()
	user.Email = email
	return user, err
	}

func (d MySQLDriver) GetUsers() (users []model.User, err error){
	dbRef, err := d.OpenDB()
	defer dbRef.Close()
	checkErr(err)
	db, stmt, err := d.Query("SELECT email, admin FROM user;")
	defer stmt.Close()
	defer db.Close()
	rows, err := stmt.Query()
	checkErr(err)
	for rows.Next() {
		a := model.User{}
		rows.Scan(&a.Email,&a.Admin)
		users = append(users, a)
	}
	defer rows.Close()
	return users, err
	}

//ExportCorrelations export correlations with the given attributes
func (d MySQLDriver) ExportCorrelations(filterAttributes []model.Attribute,
	taxonomyId int64) (papers []model.Paper, err error) {
	dbRef, err := d.OpenDB()
	defer dbRef.Close()
	checkErr(err)
	//prepare list of attribute ids for the where clause
	queryStr := ""
	parameters := []interface{}{taxonomyId}
	for _, attribute := range filterAttributes {
		queryStr+=" and atts REGEXP ?"
		parameters = append(parameters,attribute.Text)
	}
	queryStr+=";"
	queryStr = `select id_paper, citation, bib,leaf_atts
		from paper_merged_attributes
		where id_taxonomy=?`+queryStr
	fmt.Println(queryStr)
	fmt.Println(parameters...)
	fmt.Println(len(parameters))
	db, stmt, err := d.Query(queryStr)
	defer stmt.Close()
	defer db.Close()
	rows, err := stmt.Query(parameters...)
	checkErr(err)
	for rows.Next() {
		a := model.Paper{}
		rows.Scan(&a.ID, &a.Citation, &a.Bib,&a.StrAttributes)
		papers = append(papers, a)
	}
	defer rows.Close()
	return papers, err
}

func (d MySQLDriver) ExportCorrelationsCSV(filterAttributes []model.Attribute){

}

func (d MySQLDriver) GetAllTaxonomies() (taxonomies []model.Taxonomy,
	err error){
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
		rows.Scan(&a.ID,&a.Text)
		taxonomies = append(taxonomies, a)
	}
	defer rows.Close()
	return taxonomies, err
	}

func (d MySQLDriver) GetTaxonomyID(text string) (taxonomies []model.Taxonomy,
	err error){
	dbRef, err := d.OpenDB()
	defer dbRef.Close()
	checkErr(err)
	db, stmt, err := d.Query("select id_taxonomy from taxonomy where text = ?;")
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

func (d MySQLDriver) GetTaxonomyPermissions(email string) (taxonomies []model.Taxonomy, err error){
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

func (d MySQLDriver) UpdateTaxonomyPermissions(email string, permissions string) (result model.Result, err error){
	dbRef, err := d.OpenDB()
	defer dbRef.Close()
	checkErr(err)
	dbRef.Exec("UPDATE user SET taxonomies = ? WHERE email = ?;", permissions, email)
	result.Success = true
	return result, err
	}

func (d MySQLDriver) AddTaxonomy(taxonomy string, dimension string) (result model.Result, err error){
	dbRef, err := d.OpenDB()
	defer dbRef.Close()
	checkErr(err)
	db, stmt, err := d.Query("select count(id_taxonomy) as taxonomyCount from taxonomy where text = ?;")
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
	dbRef.Exec("INSERT IGNORE INTO taxonomy (text) VALUES (?);", taxonomy);
	dbRef.Exec("INSERT IGNORE INTO dimension (id_taxonomy, text) VALUES ((SELECT DISTINCT id_taxonomy FROM taxonomy WHERE text = ?), \"Interdimensional view\");", taxonomy)
	dbRef.Exec("INSERT IGNORE INTO dimension (id_taxonomy, text) VALUES ((SELECT DISTINCT id_taxonomy FROM taxonomy WHERE text = ?), ?);", taxonomy, dimension)
	result.Success = true
	return result, err
	}

func (d MySQLDriver) RemoveTaxonomy(taxonomyId int64) (result model.Result, err error){
	dbRef, err := d.OpenDB()
	defer dbRef.Close()
	checkErr(err)
	taxonomyIdStr := strconv.Itoa(int(taxonomyId))
	dbRef.Exec("DELETE FROM taxonomy WHERE id_taxonomy = ?;", taxonomyIdStr)
	result.Success = true
	return result, err
	}

func (d MySQLDriver) GetAllAttributes(taxonomyId int64) (attributes []model.Attribute,
	err error){
	dbRef, err := d.OpenDB()
	defer dbRef.Close()
	checkErr(err)
	taxonomyIdStr := strconv.Itoa(int(taxonomyId))
	db, stmt, err := d.Query(`select distinct attribute.id_attribute as id, attribute.text, allparentsperattribute.parents as parentText, dimension.text as dimensionText, attribute.synonyms
		from attribute inner join allparentsperattribute on (attribute.id_attribute = allparentsperattribute.id_attribute and attribute.id_taxonomy = ?) left outer join taxonomy_dimension on (attribute.id_attribute = taxonomy_dimension.id_attribute) left outer join dimension on (taxonomy_dimension.id_dimension = dimension.id_dimension) order by attribute.id_attribute;`)
	defer stmt.Close()
	defer db.Close()
	rows, err := stmt.Query(taxonomyIdStr)
	checkErr(err)
	for rows.Next() {
		a := model.Attribute{}
		rows.Scan(&a.ID,&a.Text,&a.ParentText,&a.Dimension,&a.Synonyms)
		attributes = append(attributes, a)
	}
	defer rows.Close()
	return attributes, err
	}

func (d MySQLDriver) GetLeafAttributes(taxonomyId int64) (attributes []model.Attribute,
	err error){
	dbRef, err := d.OpenDB()
	defer dbRef.Close()
	checkErr(err)
	taxonomyIdStr := strconv.Itoa(int(taxonomyId))
	db, stmt, err := d.Query(`select distinct attr.id_attribute as id1, attr.text as attr1, allparentsperattribute.parents as parentText, dimension.text as dimensionText, attr.synonyms
		from (select distinct attribute.id_attribute, attribute.text, attribute.synonyms from attribute left outer join taxonomy_relation on (attribute.id_attribute = taxonomy_relation.id_dest_attribute and taxonomy_relation.id_relation > 2) where attribute.id_taxonomy = ? and taxonomy_relation.id_taxonomy_relation IS NULL) as attr inner join allparentsperattribute on (attr.id_attribute = allparentsperattribute.id_attribute) left outer join taxonomy_dimension on (attr.id_attribute = taxonomy_dimension.id_attribute) left outer join dimension on (taxonomy_dimension.id_dimension = dimension.id_dimension) order by attr.id_attribute;`)
	defer stmt.Close()
	defer db.Close()
	rows, err := stmt.Query(taxonomyIdStr)
	checkErr(err)
	for rows.Next() {
		a := model.Attribute{}
		rows.Scan(&a.ID,&a.Text,&a.ParentText,&a.Dimension,&a.Synonyms)
		attributes = append(attributes, a)
	}
	defer rows.Close()
	return attributes, err
	}
func (d MySQLDriver) GetAllDimensions(taxonomyId int64) (dimensions []model.Dimension,
	err error){
	dbRef, err := d.OpenDB()
	defer dbRef.Close()
	checkErr(err)
	taxonomyIdStr := strconv.Itoa(int(taxonomyId))
	db, stmt, err := d.Query(`select id_dimension,text,xMajor,yMajor
		from dimension where id_taxonomy = ? order by id_dimension`)
	defer stmt.Close()
	defer db.Close()
	rows, err := stmt.Query(taxonomyIdStr)
	checkErr(err)
	for rows.Next() {
		a := model.Dimension{}
		rows.Scan(&a.ID,&a.Text,&a.XMajor,&a.YMajor)
		dimensions = append(dimensions, a)
	}
	defer rows.Close()
	return dimensions, err
	}

	func (d MySQLDriver) GetAllCitations(taxonomyId int64) (papers []model.Paper,
		err error){
		dbRef, err := d.OpenDB()
		defer dbRef.Close()
		checkErr(err)
		taxonomyIdStr := strconv.Itoa(int(taxonomyId))
		db, stmt, err := d.Query(`select id_paper,citation,citation as text,referenceCount,bib
			from paper where id_taxonomy = ? order by id_paper`)
		defer stmt.Close()
		defer db.Close()
		rows, err := stmt.Query(taxonomyIdStr)
		checkErr(err)
		for rows.Next() {
			a := model.Paper{}
			rows.Scan(&a.ID,&a.Citation,&a.Text,&a.ReferenceCount,&a.Bib)
			papers = append(papers, a)
		}
		defer rows.Close()
		return papers, err
		}

	func (d MySQLDriver) GetCitationCount(taxonomyId int64) (citationCounts []model.CitationCount, err error){
		dbRef, err := d.OpenDB()
		defer dbRef.Close()
		checkErr(err)
		taxonomyIdStr := strconv.Itoa(int(taxonomyId))
		db, stmt, err := d.Query("select count(distinct id_paper) as citationCount, max(referenceCount) as maxReferenceCount, sum(referenceCount) as referenceCountSum from paper where id_taxonomy = ?;")
		defer stmt.Close()
		defer db.Close()
		rows, err := stmt.Query(taxonomyIdStr)
		checkErr(err)
		for rows.Next() {
			a := model.CitationCount{}
			rows.Scan(&a.CitationCount,&a.MaxReferenceCount,&a.ReferenceCountSum)
			citationCounts = append(citationCounts, a)
		}
		defer rows.Close()
		return citationCounts, err
		}

	func (d MySQLDriver) GetCitationCounts(taxonomyId int64) (citationCounts []model.CitationCount, err error){
		dbRef, err := d.OpenDB()
		defer dbRef.Close()
		checkErr(err)
		taxonomyIdStr := strconv.Itoa(int(taxonomyId))
		db, stmt, err := d.Query("select distinct attribute.text, count(distinct tmp.id_paper) as citationCount, sum(case when tmp.referenceCount is not null then tmp.referenceCount else 0 end) as referenceCountSum from attribute left outer join (select distinct mapping.id_attribute, mapping.id_paper, paper.referenceCount from mapping inner join paper on (mapping.id_paper = paper.id_paper and paper.id_taxonomy = ?)) as tmp on (attribute.id_attribute = tmp.id_attribute) where attribute.id_taxonomy = ? group by attribute.id_attribute;")
		defer stmt.Close()
		defer db.Close()
		rows, err := stmt.Query(taxonomyIdStr, taxonomyIdStr)
		checkErr(err)
		for rows.Next() {
			a := model.CitationCount{}
			rows.Scan(&a.Attribute,&a.CitationCount,&a.ReferenceCountSum)
			citationCounts = append(citationCounts, a)
		}
		defer rows.Close()
		return citationCounts, err
		}

	func (d MySQLDriver) GetCitationCountsIncludingChildren(taxonomyId int64) (citationCounts []model.CitationCount, err error){
		dbRef, err := d.OpenDB()
		defer dbRef.Close()
		checkErr(err)
		taxonomyIdStr := strconv.Itoa(int(taxonomyId))
		db, stmt, err := d.Query("select distinct allchildrenperattribute.text, count(distinct tmp.id_paper) as citationCount, sum(case when tmp.referenceCount is not null then tmp.referenceCount else 0 end) from allchildrenperattribute left outer join (select distinct mapping.id_attribute, mapping.id_paper, paper.referenceCount from mapping inner join paper on (mapping.id_paper = paper.id_paper and paper.id_taxonomy = ?)) as tmp on (FIND_IN_SET(tmp.id_attribute, allchildrenperattribute.children)) where allchildrenperattribute.id_taxonomy = ? group by allchildrenperattribute.text;")
		defer stmt.Close()
		defer db.Close()
		rows, err := stmt.Query(taxonomyIdStr, taxonomyIdStr)
		checkErr(err)
		for rows.Next() {
			a := model.CitationCount{}
			rows.Scan(&a.Attribute,&a.CitationCount,&a.ReferenceCountSum)
			citationCounts = append(citationCounts, a)
		}
		defer rows.Close()
		return citationCounts, err
		}

	func (d MySQLDriver) UpdateCitationReferenceCounts(taxonomyId int64, referenceCounts []model.ReferenceCount) (result model.Result, err error){
	  	dbRef, err := d.OpenDB()
		defer dbRef.Close()
		checkErr(err)
		taxonomyIdStr := strconv.Itoa(int(taxonomyId))
    	for _, elem := range referenceCounts {
    		referenceCountStr := strconv.Itoa(elem.ReferenceCount)
			dbRef.Exec("update paper set referenceCount = ? where id_taxonomy = ? and citation = ?;", referenceCountStr, taxonomyIdStr, elem.Citation);
		}
		result.Success = true
		return result, err
		}

	func (d MySQLDriver) UpdateMajor(taxonomyId int64, text string, major int) (result model.Result, err error){
	  	dbRef, err := d.OpenDB()
		defer dbRef.Close()
		checkErr(err)
		majorStr := strconv.Itoa(major)
		taxonomyIdStr := strconv.Itoa(int(taxonomyId))
    	dbRef.Exec("update attribute set major = ? where text = ? and id_taxonomy = ?;", majorStr, text, taxonomyIdStr)
		result.Success = true
		return result, err
		}

	func (d MySQLDriver) UpdateCitationMapping(taxonomyId int64, attribute string, citations []model.Paper) (result model.Result, err error){
	  	dbRef, err := d.OpenDB()
		defer dbRef.Close()
		checkErr(err)
		taxonomyIdStr := strconv.Itoa(int(taxonomyId))
		db, stmt, err := d.Query("SELECT DISTINCT id_attribute FROM attribute WHERE text = ? and id_taxonomy = ?;")
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
		citationTupleString := ""
		for _, elem := range citations {
			if citationTupleString != "" {
				citationTupleString += ","
			}
			citationTupleString += "(" + strconv.Itoa(elem.ID) + "," + attributeIDStr + ")"
	    }
	    dbRef.Exec("DELETE FROM mapping WHERE id_attribute = ? AND (id_paper, id_attribute) NOT IN (?);", attributeIDStr, citationTupleString)
	    dbRef.Exec("INSERT IGNORE INTO mapping (id_paper, id_attribute) VALUES ?;", citationTupleString);
		result.Success = true
		return result, err
		}

	func (d MySQLDriver) UpdateCitationMappings(taxonomyId int64, mappings []model.CitationMapping) (result model.Result, err error){
	  	dbRef, err := d.OpenDB()
		defer dbRef.Close()
		checkErr(err)
		taxonomyIdStr := strconv.Itoa(int(taxonomyId))
		savedPapers := []model.Paper{}
		for _, elem := range mappings {
			paperIDStr := strconv.Itoa(elem.PaperID)
			found := false
		    for i := 0; i < len(savedPapers); i++ {
		        if (savedPapers[i].ID == elem.PaperID) {
		        	found = true
		        	break
		        }
		    }
		    paperID := elem.PaperID
		   	if (!found) {
				paperIDStr = strconv.Itoa(paperID)
				referenceCountStr := strconv.Itoa(elem.ReferenceCount)
				if paperID < 0 {
					db, stmt, err := d.Query("SELECT MAX(id_paper) AS maxID FROM paper WHERE id_taxonomy = ?;")
					rows, err := stmt.Query(taxonomyIdStr)
					checkErr(err)
					stmt.Close()
					db.Close()
					var maxID int
					maxID = 0
					for rows.Next() {
						rows.Scan(&maxID)
					}
					rows.Close()
					paperID = maxID+1
					paperIDStr = strconv.Itoa(paperID)
					bibTex := ""
					if elem.Bib != "empty" {
						bibTex = elem.Bib
					}
					dbRef.Exec("INSERT IGNORE INTO paper (id_taxonomy, id_paper, citation, bib, referenceCount, author, keywords) VALUES (?, ?, ?, ?, ?, ?, ?);", taxonomyIdStr, paperIDStr, elem.Citation, bibTex, referenceCountStr, elem.Author, elem.Keywords)
				} else {
					if elem.Bib != "empty" {
						dbRef.Exec("UPDATE paper SET citation = ?, bib = ?, referenceCount = ?, author = ?, keywords = ? WHERE id_taxonomy = ? AND id_paper = ?;", elem.Citation, elem.Bib, referenceCountStr, elem.Author, elem.Keywords, taxonomyIdStr, paperIDStr)
					} else {
						dbRef.Exec("UPDATE paper SET citation = ?, referenceCount = ?, author = ?, keywords = ?) WHERE id_taxonomy = ? AND id_paper = ?;", elem.Citation, referenceCountStr, elem.Author, elem.Keywords, taxonomyIdStr, paperIDStr)
					}
				}
				a := model.Paper{ID: paperID}
				savedPapers = append(savedPapers, a)
			}
			if elem.OccurrenceCount <= 0 {
				dbRef.Exec("DELETE FROM mapping WHERE id_paper = ? AND id_attribute = (SELECT DISTINCT id_attribute FROM attribute WHERE text = ? AND id_taxonomy = ?);", paperIDStr, elem.Attribute, taxonomyIdStr)
			} else {
				occurrenceCountStr := strconv.Itoa(elem.OccurrenceCount)
				dbRef.Exec("REPLACE INTO mapping (id_paper, id_attribute, occurrenceCount) VALUES (?, (SELECT DISTINCT id_attribute FROM attribute WHERE text = ? AND id_taxonomy = ?), ?);", paperIDStr, elem.Attribute, taxonomyIdStr, occurrenceCountStr)
			}
		}
		result.Success = true
		return result, err
		}

	func (d MySQLDriver) GetRelationTypes() (relationTypes []model.RelationType, err error){
		dbRef, err := d.OpenDB()
		defer dbRef.Close()
		checkErr(err)
		db, stmt, err := d.Query("select id_relation as id, text, comment from relation order by id_relation;")
		defer stmt.Close()
		defer db.Close()
		rows, err := stmt.Query()
		checkErr(err)
		for rows.Next() {
			a := model.RelationType{}
			rows.Scan(&a.ID,&a.Text,&a.Comment)
			relationTypes = append(relationTypes, a)
		}
		defer rows.Close()
		return relationTypes, err
		}


	func (d MySQLDriver) GetCitationsPerAttribute(taxonomyId int64, attribute string) (papers []model.Paper, err error){
		dbRef, err := d.OpenDB()
		defer dbRef.Close()
		checkErr(err)
		taxonomyIdStr := strconv.Itoa(int(taxonomyId))
		db, stmt, err := d.Query("select distinct paper.id_paper, paper.citation, paper.bib, paper.referenceCount, mapping.occurrenceCount from paper inner join mapping on (paper.id_paper = mapping.id_paper and paper.id_taxonomy = ?) inner join attribute on (mapping.id_attribute = attribute.id_attribute and attribute.text = ? and attribute.id_taxonomy = ?) order by paper.id_paper;")
		defer stmt.Close()
		defer db.Close()
		rows, err := stmt.Query(taxonomyIdStr, attribute, taxonomyIdStr)
		checkErr(err)
		for rows.Next() {
			a := model.Paper{}
			rows.Scan(&a.ID,&a.Citation,&a.Bib,&a.ReferenceCount,&a.OccurrenceCount)
			papers = append(papers, a)
		}
		defer rows.Close()
		return papers, err
		}

	func (d MySQLDriver) GetCitationsPerAttributeIncludingChildren(taxonomyId int64, attribute string) (papers []model.Paper, err error){
		dbRef, err := d.OpenDB()
		defer dbRef.Close()
		checkErr(err)
		taxonomyIdStr := strconv.Itoa(int(taxonomyId))
		db, stmt, err := d.Query("select distinct paper.id_paper, paper.citation, paper.bib, paper.referenceCount, sum(mapping.occurrenceCount) as occurrenceCount from allchildrenperattribute inner join mapping on (allchildrenperattribute.text = ? and allchildrenperattribute.id_taxonomy = ? and FIND_IN_SET(mapping.id_attribute, allchildrenperattribute.children)) inner join paper on (mapping.id_paper = paper.id_paper and paper.id_taxonomy = ?) group by allchildrenperattribute.id_attribute, mapping.id_paper order by paper.id_paper;")
		defer stmt.Close()
		defer db.Close()
		rows, err := stmt.Query(attribute, taxonomyIdStr, taxonomyIdStr)
		checkErr(err)
		for rows.Next() {
			a := model.Paper{}
			rows.Scan(&a.ID,&a.Citation,&a.Bib,&a.ReferenceCount,&a.OccurrenceCount)
			papers = append(papers, a)
		}
		defer rows.Close()
		return papers, err
		}

	func (d MySQLDriver) GetConceptRelations(taxonomyId int64) (conceptCorrelations []model.ConceptCorrelation, err error){
		dbRef, err := d.OpenDB()
		defer dbRef.Close()
		checkErr(err)
		taxonomyIdStr := strconv.Itoa(int(taxonomyId))
		db, stmt, err := d.Query(`select distinct attribute1.text as Attribute1, attribute2.text as Attribute2, attribute1.text as Text1, attribute2.text as Text2, attribute1.id_attribute as ID1, attribute2.id_attribute as ID2, correlation.value from attribute as attribute1 inner join (select distinct mapping1.id_attribute as attr1, mapping2.id_attribute as attr2, count(distinct mapping1.id_paper) as value from mapping as mapping1 inner join mapping as mapping2 on (mapping1.id_paper = mapping2.id_paper and mapping1.id_attribute <= mapping2.id_attribute) group by mapping1.id_attribute,mapping2.id_attribute) as correlation on (attribute1.id_taxonomy = ? and attribute1.id_attribute = correlation.attr1) inner join attribute as attribute2 on (attribute2.id_taxonomy = ? and attribute2.id_attribute = correlation.attr2);`)
		defer stmt.Close()
		defer db.Close()
		rows, err := stmt.Query(taxonomyIdStr, taxonomyIdStr)
		checkErr(err)
		for rows.Next() {
			a := model.ConceptCorrelation{}
			rows.Scan(&a.Attribute1,&a.Attribute2,&a.Text1,&a.Text2,&a.ID1,&a.ID2,&a.Value)
			conceptCorrelations = append(conceptCorrelations, a)
		}
		defer rows.Close()
		return conceptCorrelations, err
		}

	func (d MySQLDriver) GetConceptRelations3D(taxonomyId int64) (conceptCorrelations []model.ConceptCorrelation, err error){
		dbRef, err := d.OpenDB()
		defer dbRef.Close()
		checkErr(err)
		taxonomyIdStr := strconv.Itoa(int(taxonomyId))
		db, stmt, err := d.Query(`select distinct attribute1.text as Attribute1, attribute2.text as Attribute2, attribute3.text as Attribute3, attribute1.text as Text1, attribute2.text as Text2, attribute3.text as Text3, attribute1.id_attribute as ID1, attribute2.id_attribute as ID2, attribute3.id_attribute as ID3, correlation.value from attribute as attribute1 inner join (select distinct mapping1.id_attribute as attr1, mapping2.id_attribute as attr2, mapping3.id_attribute as attr3, count(distinct mapping1.id_paper) as value from mapping as mapping1 inner join mapping as mapping2 on (mapping1.id_paper = mapping2.id_paper and mapping1.id_attribute < mapping2.id_attribute) inner join mapping as mapping3 on (mapping1.id_paper = mapping3.id_paper and mapping2.id_attribute < mapping3.id_attribute) group by mapping1.id_attribute,mapping2.id_attribute,mapping3.id_attribute) as correlation on (attribute1.id_taxonomy = ? and attribute1.id_attribute = correlation.attr1) inner join attribute as attribute2 on (attribute2.id_taxonomy = ? and attribute2.id_attribute = correlation.attr2) inner join attribute as attribute3 on (attribute3.id_taxonomy = ? and attribute3.id_attribute = correlation.attr3);`)
		defer stmt.Close()
		defer db.Close()
		rows, err := stmt.Query(taxonomyIdStr, taxonomyIdStr, taxonomyIdStr)
		checkErr(err)
		for rows.Next() {
			a := model.ConceptCorrelation{}
			rows.Scan(&a.Attribute1,&a.Attribute2,&a.Attribute3,&a.Text1,&a.Text2,&a.Text3,&a.ID1,&a.ID2,&a.ID3,&a.Value)
			conceptCorrelations = append(conceptCorrelations, a)
		}
		defer rows.Close()
		return conceptCorrelations, err
		}

	func (d MySQLDriver) GetConceptRelationsWithReferenceCounts(taxonomyId int64) (conceptCorrelations []model.ConceptCorrelation, err error){
		dbRef, err := d.OpenDB()
		defer dbRef.Close()
		checkErr(err)
		taxonomyIdStr := strconv.Itoa(int(taxonomyId))
		db, stmt, err := d.Query(`select distinct attribute1.text as Attribute1, attribute2.text as Attribute2, attribute1.text as Text1, attribute2.text as Text2, attribute1.id_attribute as ID1, attribute2.id_attribute as ID2, SUM(paper.referenceCount) as value from attribute as attribute1 inner join (select distinct mapping1.id_attribute as attr1, mapping2.id_attribute as attr2, mapping1.id_paper from mapping as mapping1 inner join mapping as mapping2 on (mapping1.id_paper = mapping2.id_paper and mapping1.id_attribute <= mapping2.id_attribute)) as correlation on (attribute1.id_taxonomy = ? and attribute1.id_attribute = correlation.attr1) inner join attribute as attribute2 on (attribute2.id_taxonomy = ? and attribute2.id_attribute = correlation.attr2) inner join paper on (correlation.id_paper = paper.id_paper and paper.id_taxonomy = ?) group by attribute1.id_attribute, attribute2.id_attribute;`)
		defer stmt.Close()
		defer db.Close()
		rows, err := stmt.Query(taxonomyIdStr, taxonomyIdStr, taxonomyIdStr)
		checkErr(err)
		for rows.Next() {
			a := model.ConceptCorrelation{}
			rows.Scan(&a.Attribute1,&a.Attribute2,&a.Text1,&a.Text2,&a.ID1,&a.ID2,&a.Value)
			conceptCorrelations = append(conceptCorrelations, a)
		}
		defer rows.Close()
		return conceptCorrelations, err
		}

	func (d MySQLDriver) GetConceptRelationsWithReferenceCounts3D(taxonomyId int64) (conceptCorrelations []model.ConceptCorrelation, err error){
		dbRef, err := d.OpenDB()
		defer dbRef.Close()
		checkErr(err)
		taxonomyIdStr := strconv.Itoa(int(taxonomyId))
		db, stmt, err := d.Query(`select distinct attribute1.text as Attribute1, attribute2.text as Attribute2, attribute3.text as Attribute3, attribute1.text as Text1, attribute2.text as Text2, attribute3.text as Text3, attribute1.id_attribute as ID1, attribute2.id_attribute as ID2, attribute3.id_attribute as ID3, SUM(paper.referenceCount) as value from attribute as attribute1 inner join (select distinct mapping1.id_attribute as attr1, mapping2.id_attribute as attr2, mapping3.id_attribute as attr3, mapping1.id_paper from mapping as mapping1 inner join mapping as mapping2 on (mapping1.id_paper = mapping2.id_paper) inner join mapping as mapping3 on (mapping1.id_paper = mapping3.id_paper)) as correlation on (attribute1.id_taxonomy = ? and attribute1.id_attribute = correlation.attr1) inner join attribute as attribute2 on (attribute2.id_taxonomy = ? and attribute2.id_attribute = correlation.attr2) inner join attribute as attribute3 on (attribute3.id_taxonomy = ? and attribute3.id_attribute = correlation.attr3) inner join paper on (correlation.id_paper = paper.id_paper and paper.id_taxonomy = ?) group by attribute1.id_attribute, attribute2.id_attribute, attribute3.id_attribute;`)
		defer stmt.Close()
		defer db.Close()
		rows, err := stmt.Query(taxonomyIdStr, taxonomyIdStr, taxonomyIdStr, taxonomyIdStr)
		checkErr(err)
		for rows.Next() {
			a := model.ConceptCorrelation{}
			rows.Scan(&a.Attribute1,&a.Attribute2,&a.Attribute3,&a.Text1,&a.Text2,&a.Text3,&a.ID1,&a.ID2,&a.ID3,&a.Value)
			conceptCorrelations = append(conceptCorrelations, a)
		}
		defer rows.Close()
		return conceptCorrelations, err
		}

	func (d MySQLDriver) GetAllConceptRelations(taxonomyId int64) (conceptCorrelations []model.ConceptCorrelation, err error){
		dbRef, err := d.OpenDB()
		defer dbRef.Close()
		checkErr(err)
		taxonomyIdStr := strconv.Itoa(int(taxonomyId))
		db, stmt, err := d.Query(`select distinct attribute1.text as Attribute1, attribute2.text as Attribute2, attribute1.text as Text1, attribute2.text as Text2, attribute1.id_attribute as ID1, attribute2.id_attribute as ID2, count(distinct mapping1.id_paper) as value from allchildrenperattribute as allchildrenperattribute1 inner join mapping as mapping1 on (allchildrenperattribute1.id_taxonomy = ? and FIND_IN_SET(mapping1.id_attribute, allchildrenperattribute1.children)) inner join mapping as mapping2 on (mapping1.id_paper = mapping2.id_paper) inner join allchildrenperattribute as allchildrenperattribute2 on (allchildrenperattribute2.id_taxonomy = ? and allchildrenperattribute1.id_attribute <= allchildrenperattribute2.id_attribute and FIND_IN_SET(mapping2.id_attribute, allchildrenperattribute2.children)) inner join attribute as attribute1 on (allchildrenperattribute1.id_attribute = attribute1.id_attribute) inner join attribute as attribute2 on (allchildrenperattribute2.id_attribute = attribute2.id_attribute) group by attribute1.id_attribute, attribute2.id_attribute;`)
		defer stmt.Close()
		defer db.Close()
		rows, err := stmt.Query(taxonomyIdStr, taxonomyIdStr)
		checkErr(err)
		for rows.Next() {
			a := model.ConceptCorrelation{}
			rows.Scan(&a.Attribute1,&a.Attribute2,&a.Text1,&a.Text2,&a.ID1,&a.ID2,&a.Value)
			conceptCorrelations = append(conceptCorrelations, a)
		}
		defer rows.Close()
		return conceptCorrelations, err
		}

	func (d MySQLDriver) GetAllConceptRelations3D(taxonomyId int64) (conceptCorrelations []model.ConceptCorrelation, err error){
		dbRef, err := d.OpenDB()
		defer dbRef.Close()
		checkErr(err)
		taxonomyIdStr := strconv.Itoa(int(taxonomyId))
		db, stmt, err := d.Query(`select distinct attribute1.text as Attribute1, attribute2.text as Attribute2, attribute3.text as Attribute3, attribute1.text as Text1, attribute2.text as Text2, attribute3.text as Text3, attribute1.id_attribute as ID1, attribute2.id_attribute as ID2, attribute3.id_attribute as ID3, count(distinct mapping1.id_paper) as value from allchildrenperattribute as allchildrenperattribute1 inner join mapping as mapping1 on (allchildrenperattribute1.id_taxonomy = ? and FIND_IN_SET(mapping1.id_attribute, allchildrenperattribute1.children)) inner join mapping as mapping2 on (mapping1.id_paper = mapping2.id_paper) inner join allchildrenperattribute as allchildrenperattribute2 on (allchildrenperattribute2.id_taxonomy = ? and FIND_IN_SET(mapping2.id_attribute, allchildrenperattribute2.children)) inner join mapping as mapping3 on (mapping1.id_paper = mapping3.id_paper) inner join allChildrenPerAttribute as allchildrenperattribute3 on (allchildrenperattribute3.id_taxonomy = ? and FIND_IN_SET(mapping3.id_attribute, allchildrenperattribute3.children)) inner join attribute as attribute1 on (allchildrenperattribute1.id_attribute = attribute1.id_attribute) inner join attribute as attribute2 on (allchildrenperattribute2.id_attribute = attribute2.id_attribute) inner join attribute as attribute3 on (allchildrenperattribute3.id_attribute = attribute3.id_attribute) group by attribute1.id_attribute, attribute2.id_attribute, attribute3.id_attribute;`)
		defer stmt.Close()
		defer db.Close()
		rows, err := stmt.Query(taxonomyIdStr, taxonomyIdStr, taxonomyIdStr)
		checkErr(err)
		for rows.Next() {
			a := model.ConceptCorrelation{}
			rows.Scan(&a.Attribute1,&a.Attribute2,&a.Attribute3,&a.Text1,&a.Text2,&a.Text3,&a.ID1,&a.ID2,&a.ID3,&a.Value)
			conceptCorrelations = append(conceptCorrelations, a)
		}
		defer rows.Close()
		return conceptCorrelations, err
		}

	func (d MySQLDriver) GetAllConceptRelationsWithReferenceCounts(taxonomyId int64) (conceptCorrelations []model.ConceptCorrelation, err error){
		dbRef, err := d.OpenDB()
		defer dbRef.Close()
		checkErr(err)
		taxonomyIdStr := strconv.Itoa(int(taxonomyId))
		db, stmt, err := d.Query(`select distinct correlation.Attribute1, correlation.Attribute2, correlation.Text1, correlation.Text2, correlation.ID1, correlation.ID2, SUM(paper.referenceCount) as value from (select distinct attribute1.text as Attribute1, attribute2.text as Attribute2, attribute1.text as Text1, attribute2.text as Text2, attribute1.id_attribute as ID1, attribute2.id_attribute as ID2, mapping1.id_paper from allchildrenperattribute as allchildrenperattribute1 inner join mapping as mapping1 on (allchildrenperattribute1.id_taxonomy = ? and FIND_IN_SET(mapping1.id_attribute, allchildrenperattribute1.children)) inner join mapping as mapping2 on (mapping1.id_paper = mapping2.id_paper) inner join allchildrenperattribute as allchildrenperattribute2 on (allchildrenperattribute2.id_taxonomy = ? and allchildrenperattribute1.id_attribute <= allchildrenperattribute2.id_attribute and FIND_IN_SET(mapping2.id_attribute, allchildrenperattribute2.children)) inner join attribute as attribute1 on (allchildrenperattribute1.id_attribute = attribute1.id_attribute) inner join attribute as attribute2 on (allchildrenperattribute2.id_attribute = attribute2.id_attribute)) as correlation inner join paper on (correlation.id_paper = paper.id_paper and paper.id_taxonomy = ?) group by correlation.ID1, correlation.ID2;`)
		defer stmt.Close()
		defer db.Close()
		rows, err := stmt.Query(taxonomyIdStr, taxonomyIdStr, taxonomyIdStr)
		checkErr(err)
		for rows.Next() {
			a := model.ConceptCorrelation{}
			rows.Scan(&a.Attribute1,&a.Attribute2,&a.Text1,&a.Text2,&a.ID1,&a.ID2,&a.Value)
			conceptCorrelations = append(conceptCorrelations, a)
		}
		defer rows.Close()
		return conceptCorrelations, err
		}

	func (d MySQLDriver) GetAllConceptRelationsWithReferenceCounts3D(taxonomyId int64) (conceptCorrelations []model.ConceptCorrelation, err error){
		dbRef, err := d.OpenDB()
		defer dbRef.Close()
		checkErr(err)
		taxonomyIdStr := strconv.Itoa(int(taxonomyId))
		db, stmt, err := d.Query(`select distinct correlation.Attribute1, correlation.Attribute2, correlation.Attribute3, correlation.Text1, correlation.Text2, correlation.Text3, correlation.ID1, correlation.ID2, correlation.ID3, SUM(paper.referenceCount) as value from (select distinct attribute1.text as Attribute1, attribute2.text as Attribute2, attribute3.text as Attribute3, attribute1.text as Text1, attribute2.text as Text2, attribute3.text as Text3, attribute1.id_attribute as ID1, attribute2.id_attribute as ID2, attribute3.id_attribute as ID3, mapping1.id_paper from allchildrenperattribute as allchildrenperattribute1 inner join mapping as mapping1 on (allchildrenperattribute1.id_taxonomy = ? and FIND_IN_SET(mapping1.id_attribute, allchildrenperattribute1.children)) inner join mapping as mapping2 on (mapping1.id_paper = mapping2.id_paper) inner join allchildrenperattribute as allchildrenperattribute2 on (allchildrenperattribute2.id_taxonomy = ? and FIND_IN_SET(mapping2.id_attribute, allchildrenperattribute2.children) inner join mapping as mapping3 on (mapping1.id_paper = mapping3.id_paper) inner join allparentsperattribute as allchildrenperattribute3 on (allchildrenperattribute3.id_taxonomy = ? and FIND_IN_SET(mapping3.id_attribute, allchildrenperattribute3.children)) inner join attribute as attribute1 on (allchildrenperattribute1.id_attribute = attribute1.id_attribute) inner join attribute as attribute2 on (allchildrenperattribute2.id_attribute = attribute2.id_attribute) inner join attribute as attribute3 on (allchildrenperattribute3.id_attribute = attribute3.id_attribute)) as correlation inner join paper on (correlation.id_paper = paper.id_paper and paper.id_taxonomy = ?) group by correlation.ID1, correlation.ID2, correlation.ID3;`)
		defer stmt.Close()
		defer db.Close()
		rows, err := stmt.Query(taxonomyIdStr, taxonomyIdStr, taxonomyIdStr, taxonomyIdStr)
		checkErr(err)
		for rows.Next() {
			a := model.ConceptCorrelation{}
			rows.Scan(&a.Attribute1,&a.Attribute2,&a.Attribute3,&a.Text1,&a.Text2,&a.Text3,&a.ID1,&a.ID2,&a.ID3,&a.Value)
			conceptCorrelations = append(conceptCorrelations, a)
		}
		defer rows.Close()
		return conceptCorrelations, err
		}

	func (d MySQLDriver) GetSharedPapers(taxonomyId int64, text1 string, text2 string) (papers []model.Paper, err error){
		dbRef, err := d.OpenDB()
		defer dbRef.Close()
		checkErr(err)
		taxonomyIdStr := strconv.Itoa(int(taxonomyId))
		db, stmt, err := d.Query("select distinct paper.id_paper, paper.citation, paper.bib, paper.referenceCount from attribute as attribute1 inner join mapping as mapping1 on (attribute1.text = ? and attribute1.id_taxonomy = ? and attribute1.id_attribute = mapping1.id_attribute) inner join mapping as mapping2 on (mapping1.id_paper = mapping2.id_paper) inner join attribute as attribute2 on (attribute2.text = ? and attribute2.id_taxonomy = ? and attribute2.id_attribute = mapping2.id_attribute) inner join paper on (mapping1.id_paper = paper.id_paper and paper.id_taxonomy = ?);")
		defer stmt.Close()
		defer db.Close()
		rows, err := stmt.Query(text1, taxonomyIdStr, text2, taxonomyIdStr, taxonomyIdStr)
		checkErr(err)
		for rows.Next() {
			a := model.Paper{}
			rows.Scan(&a.ID,&a.Citation,&a.Bib,&a.ReferenceCount)
			papers = append(papers, a)
		}
		defer rows.Close()
		return papers, err
		}

	func (d MySQLDriver) GetSharedPapers3D(taxonomyId int64, text1 string, text2 string, text3 string) (papers []model.Paper, err error){
		dbRef, err := d.OpenDB()
		defer dbRef.Close()
		checkErr(err)
		taxonomyIdStr := strconv.Itoa(int(taxonomyId))
		db, stmt, err := d.Query("select distinct paper.id_paper, paper.citation, paper.bib, paper.referenceCount from attribute as attribute1 inner join mapping as mapping1 on (attribute1.text = ? and attribute1.id_taxonomy = ? and attribute1.id_attribute = mapping1.id_attribute) inner join mapping as mapping2 on (mapping1.id_paper = mapping2.id_paper) inner join attribute as attribute2 on (attribute2.text = ? and attribute2.id_taxonomy = ? and attribute2.id_attribute = mapping2.id_attribute) inner join mapping as mapping3 on (mapping1.id_paper = mapping3.id_paper) inner join attribute as attribute3 on (attribute3.text = ? and attribute3.id_taxonomy = ? and attribute3.id_attribute = mapping3.id_attribute) inner join paper on (mapping1.id_paper = paper.id_paper and paper.id_taxonomy = ?);")
		defer stmt.Close()
		defer db.Close()
		rows, err := stmt.Query(text1, taxonomyIdStr, text2, taxonomyIdStr, text3, taxonomyIdStr, taxonomyIdStr)
		checkErr(err)
		for rows.Next() {
			a := model.Paper{}
			rows.Scan(&a.ID,&a.Citation,&a.Bib,&a.ReferenceCount)
			papers = append(papers, a)
		}
		defer rows.Close()
		return papers, err
		}

	func (d MySQLDriver) GetSharedPapersIncludingChildren(taxonomyId int64, text1 string, text2 string) (papers []model.Paper, err error){
		dbRef, err := d.OpenDB()
		defer dbRef.Close()
		checkErr(err)
		taxonomyIdStr := strconv.Itoa(int(taxonomyId))
		db, stmt, err := d.Query("select distinct paper.id_paper, paper.citation, paper.bib, paper.referenceCount from allchildrenperattribute as allchildrenperattribute1 inner join mapping as mapping1 on (allchildrenperattribute1.text = ? and allchildrenperattribute1.id_taxonomy = ? and FIND_IN_SET(mapping1.id_attribute, allchildrenperattribute1.children)) inner join mapping as mapping2 on (mapping1.id_paper = mapping2.id_paper) inner join allchildrenperattribute as allchildrenperattribute2 on (allchildrenperattribute2.text = ? and allchildrenperattribute2.id_taxonomy = ? and FIND_IN_SET(mapping2.id_attribute, allchildrenperattribute2.children)) inner join paper on (mapping1.id_paper = paper.id_paper and paper.id_taxonomy = ?);")
		defer stmt.Close()
		defer db.Close()
		rows, err := stmt.Query(text1, taxonomyIdStr, text2, taxonomyIdStr, taxonomyIdStr)
		checkErr(err)
		for rows.Next() {
			a := model.Paper{}
			rows.Scan(&a.ID,&a.Citation,&a.Bib,&a.ReferenceCount)
			papers = append(papers, a)
		}
		defer rows.Close()
		return papers, err
		}

	func (d MySQLDriver) GetSharedPapersIncludingChildren3D(taxonomyId int64, text1 string, text2 string, text3 string) (papers []model.Paper, err error){
		dbRef, err := d.OpenDB()
		defer dbRef.Close()
		checkErr(err)
		taxonomyIdStr := strconv.Itoa(int(taxonomyId))
		db, stmt, err := d.Query("select distinct paper.id_paper, paper.citation, paper.bib, paper.referenceCount from allchildrenperattribute as allchildrenperattribute1 inner join mapping as mapping1 on (allchildrenperattribute1.text = ? and allchildrenperattribute1.id_taxonomy = ? and FIND_IN_SET(mapping1.id_attribute, allchildrenperattribute1.children)) inner join mapping as mapping2 on (mapping1.id_paper = mapping2.id_paper) inner join allchildrenperattribute as allchildrenperattribute2 on (allchildrenperattribute2.text = ? and allchildrenperattribute2.id_taxonomy = ? and FIND_IN_SET(mapping2.id_attribute, allchildrenperattribute2.children)) inner join mapping as mapping3 on (mapping1.id_paper = mapping3.id_paper) inner join allchildrenperattribute as allchildrenperattribute3 on (allchildrenperattribute3.text = ? and allchildrenperattribute3.id_taxonomy = ? and FIND_IN_SET(mapping3.id_attribute, allchildrenperattribute3.children)) inner join paper on (mapping1.id_paper = paper.id_paper and paper.id_taxonomy = ?);")
		defer stmt.Close()
		defer db.Close()
		rows, err := stmt.Query(text1, taxonomyIdStr, text2, taxonomyIdStr, text3, taxonomyIdStr, taxonomyIdStr)
		checkErr(err)
		for rows.Next() {
			a := model.Paper{}
			rows.Scan(&a.ID,&a.Citation,&a.Bib,&a.ReferenceCount)
			papers = append(papers, a)
		}
		defer rows.Close()
		return papers, err
		}

	func (d MySQLDriver) GetAttributeDetails(taxonomyId int64, text string) (attributes []model.Attribute, err error){
		dbRef, err := d.OpenDB()
		defer dbRef.Close()
		checkErr(err)
		taxonomyIdStr := strconv.Itoa(int(taxonomyId))
		db, stmt, err := d.Query("select distinct attribute1.id_attribute as id1, attribute1.text as attr1, relation.parentID as parentID, relation.parentText as parentText, attribute1.synonyms from attribute as attribute1 left outer join (select distinct taxonomy_relation.id_src_attribute, attribute2.id_attribute as parentID, attribute2.text as parentText from taxonomy_relation inner join attribute as attribute2 on (attribute2.id_taxonomy = ? and taxonomy_relation.id_dest_attribute = attribute2.id_attribute and taxonomy_relation.id_relation > 2)) as relation on (attribute1.id_attribute = relation.id_src_attribute) where attribute1.text = ? and attribute1.id_taxonomy = ? order by attribute1.id_attribute;")
		defer stmt.Close()
		defer db.Close()
		rows, err := stmt.Query(taxonomyIdStr, text, taxonomyIdStr)
		checkErr(err)
		for rows.Next() {
			a := model.Attribute{}
			rows.Scan(&a.ID,&a.Text,&a.ParentID,&a.ParentText,&a.Synonyms)
			attributes = append(attributes, a)
		}
		defer rows.Close()
		return attributes, err
		}

	func (d MySQLDriver) GetCitationDetails(taxonomyId int64, text1 string, text2 string) (papers []model.Paper, err error){
		dbRef, err := d.OpenDB()
		defer dbRef.Close()
		checkErr(err)
		taxonomyIdStr := strconv.Itoa(int(taxonomyId))
		db, stmt, err := d.Query("select distinct id_paper, citation, bib, referenceCount from paper where id_taxonomy = ? and (citation = ? or citation = ?) order by id_paper;")
		defer stmt.Close()
		defer db.Close()
		rows, err := stmt.Query(taxonomyIdStr, text1, text2)
		checkErr(err)
		for rows.Next() {
			a := model.Paper{}
			rows.Scan(&a.ID,&a.Citation,&a.Bib,&a.ReferenceCount)
			papers = append(papers, a)
		}
		defer rows.Close()
		return papers, err
		}

	func (d MySQLDriver) GetAttributeCoverage(taxonomyId int64) (attributeCoverage []model.AttributeCoverage, err error){
		dbRef, err := d.OpenDB()
		defer dbRef.Close()
		checkErr(err)
		taxonomyIdStr := strconv.Itoa(int(taxonomyId))
		db, stmt, err := d.Query(`select distinct attribute.text as attributeName, paper.citation as paperName, attribute.text as text1, paper.citation as text2, attribute.id_attribute as attributeID, paper.id_paper as paperID, 1 as value from attribute inner join mapping on (attribute.id_taxonomy = ? and attribute.id_attribute = mapping.id_attribute) inner join paper on (mapping.id_paper = paper.id_paper and paper.id_taxonomy = ?);`)
		defer stmt.Close()
		defer db.Close()
		rows, err := stmt.Query(taxonomyIdStr, taxonomyIdStr)
		checkErr(err)
		for rows.Next() {
			a := model.AttributeCoverage{}
			rows.Scan(&a.AttributeName,&a.PaperName,&a.Text1,&a.Text2,&a.AttributeID,&a.PaperID,&a.Value)
			attributeCoverage = append(attributeCoverage, a)
		}
		defer rows.Close()
		return attributeCoverage, err
		}

	func (d MySQLDriver) GetAttributeCoverageWithOcurrenceCounts(taxonomyId int64) (attributeCoverage []model.AttributeCoverage, err error){
		dbRef, err := d.OpenDB()
		defer dbRef.Close()
		checkErr(err)
		taxonomyIdStr := strconv.Itoa(int(taxonomyId))
		// mapping.occurrenceCount as value (for occurrence counts instead of 1's)
		db, stmt, err := d.Query(`select distinct attribute.text as attributeName, paper.citation as paperName, attribute.text as text1, paper.citation as text2, attribute.id_attribute as attributeID, paper.id_paper as paperID, 1 as value from attribute inner join mapping on (attribute.id_taxonomy = ? and attribute.id_attribute = mapping.id_attribute) inner join paper on (mapping.id_paper = paper.id_paper and paper.id_taxonomy = ?);`)
		defer stmt.Close()
		defer db.Close()
		rows, err := stmt.Query(taxonomyIdStr, taxonomyIdStr)
		checkErr(err)
		for rows.Next() {
			a := model.AttributeCoverage{}
			rows.Scan(&a.AttributeName,&a.PaperName,&a.Text1,&a.Text2,&a.AttributeID,&a.PaperID,&a.Value)
			attributeCoverage = append(attributeCoverage, a)
		}
		defer rows.Close()
		return attributeCoverage, err
		}

	func (d MySQLDriver) GetAttributeCoverageWithReferenceCounts(taxonomyId int64) (attributeCoverage []model.AttributeCoverage, err error){
		dbRef, err := d.OpenDB()
		defer dbRef.Close()
		checkErr(err)
		taxonomyIdStr := strconv.Itoa(int(taxonomyId))
		db, stmt, err := d.Query(`select distinct attribute.text as attributeName, paper.citation as paperName, attribute.text as text1, paper.citation as text2, attribute.id_attribute as attributeID, paper.id_paper as paperID, paper.referenceCount as value from attribute inner join mapping on (attribute.id_taxonomy = ? and attribute.id_attribute = mapping.id_attribute) inner join paper on (mapping.id_paper = paper.id_paper and paper.id_taxonomy = ?);`)
		defer stmt.Close()
		defer db.Close()
		rows, err := stmt.Query(taxonomyIdStr, taxonomyIdStr)
		checkErr(err)
		for rows.Next() {
			a := model.AttributeCoverage{}
			rows.Scan(&a.AttributeName,&a.PaperName,&a.Text1,&a.Text2,&a.AttributeID,&a.PaperID,&a.Value)
			attributeCoverage = append(attributeCoverage, a)
		}
		defer rows.Close()
		return attributeCoverage, err
		}

	func (d MySQLDriver) GetParentRelationsPerAttribute(taxonomyId int64, attribute string, dimension string) (relations []model.AttributeRelation, err error){
		dbRef, err := d.OpenDB()
		defer dbRef.Close()
		checkErr(err)
		taxonomyIdStr := strconv.Itoa(int(taxonomyId))
		db, stmt, err := d.Query("select distinct attribute.text as attr, relation.text as relationText from taxonomy_relation inner join attribute on (taxonomy_relation.id_src_attribute = (select distinct id_attribute from attribute where text = ? and attribute.id_taxonomy = ?) and taxonomy_relation.id_dimension = (select distinct id_dimension from dimension where text = ? and dimension.id_taxonomy = ?) and taxonomy_relation.id_dest_attribute = attribute.id_attribute) inner join relation on (taxonomy_relation.id_relation = relation.id_relation);")
		defer stmt.Close()
		defer db.Close()
		rows, err := stmt.Query(attribute, taxonomyIdStr, dimension, taxonomyIdStr)
		checkErr(err)
		for rows.Next() {
			a := model.AttributeRelation{}
			rows.Scan(&a.Text,&a.Relation)
			relations = append(relations, a)
		}
		defer rows.Close()
		return relations, err
		}

	func (d MySQLDriver) GetChildRelationsPerAttribute(taxonomyId int64, attribute string, dimension string) (relations []model.AttributeRelation, err error){
		dbRef, err := d.OpenDB()
		defer dbRef.Close()
		checkErr(err)
		taxonomyIdStr := strconv.Itoa(int(taxonomyId))
		db, stmt, err := d.Query("select distinct attribute.text as attr, relation.text as relationText from taxonomy_relation inner join attribute on (taxonomy_relation.id_dest_attribute = (select distinct id_attribute from attribute where text = ? and attribute.id_taxonomy = ?) and taxonomy_relation.id_dimension = (select distinct id_dimension from dimension where text = ? and dimension.id_taxonomy = ?) and taxonomy_relation.id_src_attribute = attribute.id_attribute) inner join relation on (taxonomy_relation.id_relation = relation.id_relation);")
		defer stmt.Close()
		defer db.Close()
		rows, err := stmt.Query(attribute, taxonomyIdStr, dimension, taxonomyIdStr)
		checkErr(err)
		for rows.Next() {
			a := model.AttributeRelation{}
			rows.Scan(&a.Text,&a.Relation)
			relations = append(relations, a)
		}
		defer rows.Close()
		return relations, err
		}

	func (d MySQLDriver) GetAttributesPerDimension(taxonomyId int64, dimension string) (attributes []model.Attribute, err error){
		dbRef, err := d.OpenDB()
		defer dbRef.Close()
		checkErr(err)
		taxonomyIdStr := strconv.Itoa(int(taxonomyId))
		db, stmt, err := d.Query("select distinct attr.id_attribute, attr.text, allparentsperattribute.parents as parentText, attr.dimensionText, attr.synonyms, attr.major, attr.x, attr.y, attr.x3D, attr.y3D, attr.z3D from (select attribute.id_attribute, attribute.text, dimension.text as dimensionText, attribute.synonyms, attribute.major, attribute.x, attribute.y, attribute.x3D, attribute.y3D, attribute.z3D from attribute inner join taxonomy_dimension on (attribute.id_taxonomy = ? and attribute.id_attribute = taxonomy_dimension.id_attribute) inner join dimension on (taxonomy_dimension.id_dimension = dimension.id_dimension and dimension.text = ? and dimension.id_taxonomy = ?)) as attr inner join allparentsperattribute on (attr.id_attribute = allparentsperattribute.id_attribute) order by attr.id_attribute;")
		defer stmt.Close()
		defer db.Close()
		rows, err := stmt.Query(taxonomyIdStr, dimension, taxonomyIdStr)
		checkErr(err)
		for rows.Next() {
			a := model.Attribute{}
			rows.Scan(&a.ID,&a.Text,&a.ParentText,&a.Dimension,&a.Synonyms,&a.Major,&a.X,&a.Y,&a.X3D,&a.Y3D,&a.Z3D)
			attributes = append(attributes, a)
		}
		defer rows.Close()
		return attributes, err
		}

	func (d MySQLDriver) GetLeafAttributesPerDimension(taxonomyId int64, dimension string) (attributes []model.Attribute, err error){
		dbRef, err := d.OpenDB()
		defer dbRef.Close()
		checkErr(err)
		taxonomyIdStr := strconv.Itoa(int(taxonomyId))
		db, stmt, err := d.Query("select distinct attr.id_attribute, attr.text, allparentsperattribute.parents as parentText, attr.dimensionText, attr.synonyms, attr.major, attr.x, attr.y, attr.x3D, attr.y3D, attr.z3D from (select attribute.id_attribute, attribute.text, dimension.text as dimensionText, attribute.synonyms, attribute.major, attribute.x, attribute.y, attribute.x3D, attribute.y3D, attribute.z3D from attribute inner join taxonomy_dimension on (attribute.id_taxonomy = ? and attribute.id_attribute = taxonomy_dimension.id_attribute) inner join dimension on (taxonomy_dimension.id_dimension = dimension.id_dimension and dimension.text = ? and dimension.id_taxonomy = ?) left outer join taxonomy_relation on (attribute.id_attribute = taxonomy_relation.id_dest_attribute and taxonomy_relation.id_relation > 2) where taxonomy_relation.id_taxonomy_relation IS NULL) as attr inner join allparentsperattribute on (attr.id_attribute = allparentsperattribute.id_attribute) order by attr.id_attribute;")
		defer stmt.Close()
		defer db.Close()
		rows, err := stmt.Query(taxonomyIdStr, dimension, taxonomyIdStr)
		checkErr(err)
		for rows.Next() {
			a := model.Attribute{}
			rows.Scan(&a.ID,&a.Text,&a.ParentText,&a.Dimension,&a.Synonyms,&a.Major,&a.X,&a.Y,&a.X3D,&a.Y3D,&a.Z3D)
			attributes = append(attributes, a)
		}
		defer rows.Close()
		return attributes, err
		}

	func (d MySQLDriver) GetAttributeChildren(taxonomyIdStr string, dimension string, cluster model.AttributeCluster, parent model.AttributeCluster) (result string){
		dbRef, err := d.OpenDB()
		defer dbRef.Close()
		checkErr(err)
		destAttributeIdStr := strconv.Itoa(cluster.ID)
		db, stmt, err := d.Query("select distinct attribute.id_attribute, attribute.text from attribute inner join taxonomy_relation on (attribute.id_taxonomy = ? and taxonomy_relation.id_dest_attribute = ? and attribute.id_attribute = taxonomy_relation.id_src_attribute and taxonomy_relation.id_dimension = (select distinct id_dimension from dimension where text = ? and dimension.id_taxonomy = ?));")
		defer stmt.Close()
		defer db.Close()
		rows, err := stmt.Query(taxonomyIdStr, destAttributeIdStr, dimension, taxonomyIdStr)
		checkErr(err)
		clusterObjects := make([]model.AttributeCluster, 0)
		for rows.Next() {
			a := model.AttributeCluster{}
			rows.Scan(&a.ID,&a.Text)
			clusterObjects = append(clusterObjects, a)
		}
		rows.Close()
		dbRef.Close()
		jsonObj := gabs.New()
		jsonObj.Set(cluster.Text, "text")
		jsonObj.Set(parent.Text, "parent")
		jsonObj.Array("children")
		counter := 0
		for _, elem := range clusterObjects {
			jsonObj.ArrayAppend(d.GetAttributeChildren(taxonomyIdStr, dimension, elem, cluster), "children")
			counter++
		}
		if counter == 0 {
			counter = 1
		}
		jsonObj.Set(counter, "value")
		return strings.Replace(jsonObj.String(), "\\", "", -1)
		}

	func (d MySQLDriver) GetAttributeClusterPerDimension(taxonomyId int64, dimension string) (result string, err error){
		dbRef, err := d.OpenDB()
		defer dbRef.Close()
		checkErr(err)
		taxonomyIdStr := strconv.Itoa(int(taxonomyId))
		// Get root attributes
		db, stmt, err := d.Query("select distinct attr.id_attribute, attr.text from (select distinct attribute.id_attribute, attribute.text from attribute inner join taxonomy_dimension on (attribute.id_taxonomy = ? and attribute.id_attribute = taxonomy_dimension.id_attribute) inner join dimension on (taxonomy_dimension.id_dimension = dimension.id_dimension and dimension.text = ? and dimension.id_taxonomy = ?)) as attr left outer join taxonomy_relation on (attr.id_attribute = taxonomy_relation.id_src_attribute and taxonomy_relation.id_dimension = (select distinct id_dimension from dimension where text = ? and dimension.id_taxonomy = ?)) where taxonomy_relation.id_taxonomy_relation is null;")
		defer stmt.Close()
		defer db.Close()
		rows, err := stmt.Query(taxonomyIdStr, dimension, taxonomyIdStr, dimension, taxonomyIdStr)
		checkErr(err)
		clusters := make([]model.AttributeCluster, 0)
		for rows.Next() {
			a := model.AttributeCluster{}
			rows.Scan(&a.ID,&a.Text)
			clusters = append(clusters, a)
		}
		rows.Close()
		dbRef.Close()
		jsonObj := gabs.New()
		jsonObj.Set(dimension, "text")
		jsonObj.Set("", "parent")
		jsonObj.Array("children")
		root := model.AttributeCluster{ID: -1, Text: ""}
		counter := 0
		for _, elem := range clusters {
			jsonObj.ArrayAppend(d.GetAttributeChildren(taxonomyIdStr, dimension, elem, root), "children")
			counter++
		}
		if counter == 0 {
			counter = 1
		}
		jsonObj.Set(counter, "value")
		return strings.Replace(jsonObj.String(), "\\", "", -1), err
		}

	func (d MySQLDriver) GetAllChildrenAttributes(taxonomyId int64, parent string) (attributes []model.Attribute, err error){
		dbRef, err := d.OpenDB()
		defer dbRef.Close()
		checkErr(err)
		taxonomyIdStr := strconv.Itoa(int(taxonomyId))
		db, stmt, err := d.Query("SELECT DISTINCT attributeSrc.id_attribute, attributeSrc.text, allparentsperattribute.parents AS parentText, dimension.text as dimensionText FROM attribute AS attributeSrc INNER JOIN (SELECT GROUP_CONCAT(lv SEPARATOR ',') AS children FROM (SELECT @pv:=(SELECT GROUP_CONCAT(DISTINCT relation1.id_src_attribute SEPARATOR ',') FROM taxonomy_relation AS relation1 WHERE relation1.id_taxonomy = ? AND relation1.id_relation > 2 AND FIND_IN_SET(relation1.id_dest_attribute, @pv)) AS lv FROM taxonomy_relation JOIN (SELECT @pv:=attributeDest.id_attribute from attribute as attributeDest where attributeDest.text = ? and attributeDest.id_taxonomy = ?) tmp) a) AS tmpTable on ((attributeSrc.text = ? AND attributeSrc.id_taxonomy = ?) OR FIND_IN_SET(attributeSrc.id_attribute, tmpTable.children)) inner join allparentsperattribute on (allparentsperattribute.id_taxonomy = ? and attributeSrc.id_attribute = allparentsperattribute.id_attribute) left outer join taxonomy_dimension on (attributeSrc.id_attribute = taxonomy_dimension.id_attribute) left outer join dimension on (taxonomy_dimension.id_dimension = dimension.id_dimension and dimension.id_taxonomy = ?) order by attributeSrc.id_attribute;")
		defer stmt.Close()
		defer stmt.Close()	
		defer db.Close()
		rows, err := stmt.Query(taxonomyIdStr, parent, taxonomyIdStr, parent, taxonomyIdStr, taxonomyIdStr, taxonomyIdStr)
		checkErr(err)
		for rows.Next() {
			a := model.Attribute{}
			rows.Scan(&a.ID,&a.Text,&a.ParentText)
			attributes = append(attributes, a)
		}
		defer rows.Close()
		return attributes, err
		}

	func (d MySQLDriver) GetAllChildrenLeafAttributes(taxonomyId int64, parent string) (attributes []model.Attribute, err error){
		dbRef, err := d.OpenDB()
		defer dbRef.Close()
		checkErr(err)
		taxonomyIdStr := strconv.Itoa(int(taxonomyId))
		db, stmt, err := d.Query("SELECT DISTINCT attr.id_attribute, attr.text, allparentsperattribute.parents AS parentText, dimension.text AS dimensionText, attr.synonyms FROM (SELECT DISTINCT attributeSrc.id_attribute, attributeSrc.text, attributeSrc.synonyms FROM attribute AS attributeSrc INNER JOIN (SELECT GROUP_CONCAT(lv SEPARATOR ',') AS children FROM (SELECT @pv:=(SELECT GROUP_CONCAT(DISTINCT relation1.id_src_attribute SEPARATOR ',') FROM taxonomy_relation AS relation1 WHERE relation1.id_taxonomy = ? AND relation1.id_relation > 2 AND FIND_IN_SET(relation1.id_dest_attribute, @pv)) AS lv FROM taxonomy_relation JOIN (SELECT @pv:=attributeDest.id_attribute from attribute as attributeDest where attributeDest.text = ? and attributeDest.id_taxonomy = ?) tmp) a) AS tmpTable on ((attributeSrc.text = ? AND attributeSrc.id_taxonomy = ?) OR FIND_IN_SET(attributeSrc.id_attribute, tmpTable.children)) left outer join taxonomy_relation on (attributeSrc.id_attribute = taxonomy_relation.id_dest_attribute and taxonomy_relation.id_relation > 2) where taxonomy_relation.id_taxonomy_relation IS NULL) as attr inner join allparentsperattribute on (attr.id_attribute = allparentsperattribute.id_attribute) left outer join taxonomy_dimension on (attr.id_attribute = taxonomy_dimension.id_attribute) left outer join dimension on (taxonomy_dimension.id_dimension = dimension.id_dimension and dimension.id_taxonomy = ?) order by attr.id_attribute;")
		defer stmt.Close()
		defer stmt.Close()	
		defer db.Close()
		rows, err := stmt.Query(taxonomyIdStr, parent, taxonomyIdStr, parent, taxonomyIdStr, taxonomyIdStr)
		checkErr(err)
		for rows.Next() {
			a := model.Attribute{}
			rows.Scan(&a.ID,&a.Text,&a.ParentText,&a.Dimension,&a.Synonyms)
			attributes = append(attributes, a)
		}
		defer rows.Close()
		return attributes, err
		}

	func (d MySQLDriver) GetIntermediateAttributes(taxonomyId int64, minValue int64, maxValue int64) (attributes []model.Attribute, err error){
		dbRef, err := d.OpenDB()
		defer dbRef.Close()
		checkErr(err)
		taxonomyIdStr := strconv.Itoa(int(taxonomyId))
		minValueStr := strconv.Itoa(int(minValue))
		maxValueStr := strconv.Itoa(int(maxValue))
		db, stmt, err := d.Query("select distinct attribute.id_attribute, attribute.text from attribute left join taxonomy_relation on (attribute.id_attribute = taxonomy_relation.id_dest_attribute and taxonomy_relation.id_relation > 2) where attribute.id_taxonomy = ? group by attribute.id_attribute having count(distinct taxonomy_relation.id_taxonomy_relation) >= ? and count(distinct taxonomy_relation.id_taxonomy_relation) <= ?;")
		defer stmt.Close()
		defer db.Close()
		rows, err := stmt.Query(taxonomyIdStr, minValueStr, maxValueStr)
		checkErr(err)
		for rows.Next() {
			a := model.Attribute{}
			rows.Scan(&a.ID,&a.Text)
			attributes = append(attributes, a)
		}
		defer rows.Close()
		return attributes, err
		}

	func (d MySQLDriver) GetMajorAttributes(taxonomyId int64) (attributes []model.Attribute, err error){
		dbRef, err := d.OpenDB()
		defer dbRef.Close()
		checkErr(err)
		taxonomyIdStr := strconv.Itoa(int(taxonomyId))
		db, stmt, err := d.Query("select distinct attribute.id_attribute, attribute.text, allparentsperattribute.parents as parentText, dimension.text as dimensionText, attribute.synonyms, attribute.xMajor, attribute.yMajor, attribute.xMajor3D, attribute.yMajor3D, attribute.zMajor3D from attribute inner join allparentsperattribute on (attribute.id_taxonomy = ? and attribute.id_attribute = allparentsperattribute.id_attribute) inner join taxonomy_dimension on (attribute.major = 1 and attribute.id_attribute = taxonomy_dimension.id_attribute) inner join dimension on (taxonomy_dimension.id_dimension = dimension.id_dimension and dimension.id_taxonomy = ?) order by attribute.id_attribute;")
		defer stmt.Close()
		defer db.Close()
		rows, err := stmt.Query(taxonomyIdStr, taxonomyIdStr)
		checkErr(err)
		for rows.Next() {
			a := model.Attribute{}
			rows.Scan(&a.ID,&a.Text,&a.ParentText,&a.Dimension,&a.Synonyms,&a.X,&a.Y,&a.XMajor3D,&a.YMajor3D,&a.ZMajor3D)
			a.Major = 1;
			attributes = append(attributes, a)
		}
		defer rows.Close()
		return attributes, err
		}

	func (d MySQLDriver) GetAttributeRelationsPerDimension(taxonomyId int64, dimension string) (attributeRelations []model.AttributeRelation, err error){
		dbRef, err := d.OpenDB()
		defer dbRef.Close()
		checkErr(err)
		taxonomyIdStr := strconv.Itoa(int(taxonomyId))
		db, stmt, err := d.Query("select distinct attribute1.text as attributeSrc, attribute2.text as attributeDest, relation.text as relation, (case when taxonomy_relation.edgeBendPoints IS NOT NULL then taxonomy_relation.edgeBendPoints else \"\" end), annotation.annotation from attribute as attribute1 inner join taxonomy_relation on (attribute1.id_taxonomy = ? and taxonomy_relation.id_dimension = (select distinct id_dimension from dimension where text = ? and dimension.id_taxonomy = ?) and attribute1.id_attribute = taxonomy_relation.id_src_attribute) inner join attribute as attribute2 on (attribute2.id_taxonomy = ? and taxonomy_relation.id_dest_attribute = attribute2.id_attribute) inner join relation on (taxonomy_relation.id_relation = relation.id_relation) inner join taxonomy_dimension as dimension1 on (attribute1.id_attribute = dimension1.id_attribute) inner join taxonomy_dimension as dimension2 on (attribute2.id_attribute = dimension2.id_attribute) inner join dimension as dim1 on (dimension1.id_dimension = dim1.id_dimension and dim1.text = ? and dim1.id_taxonomy = ?) inner join dimension as dim2 on (dimension2.id_dimension = dim2.id_dimension and dim2.text = ? and dim2.id_taxonomy = ?) left outer join taxonomy_relation_annotation as annotation on (taxonomy_relation.id_taxonomy_relation = annotation.id_taxonomy_relation);")
		defer stmt.Close()
		defer db.Close()
		rows, err := stmt.Query(taxonomyIdStr, dimension, taxonomyIdStr, taxonomyIdStr, dimension, taxonomyIdStr, dimension, taxonomyIdStr)
		checkErr(err)
		for rows.Next() {
			a := model.AttributeRelation{}
			rows.Scan(&a.AttributeSrc,&a.AttributeDest,&a.Relation,&a.EdgeBendPoints,&a.Annotation)
			attributeRelations = append(attributeRelations, a)
		}
		defer rows.Close()
		return attributeRelations, err
		}

	func (d MySQLDriver) GetInterdimensionalRelations(taxonomyId int64) (attributeRelations []model.AttributeRelation, err error){
		dbRef, err := d.OpenDB()
		defer dbRef.Close()
		checkErr(err)
		taxonomyIdStr := strconv.Itoa(int(taxonomyId))
		db, stmt, err := d.Query("select distinct attribute1.text as attributeSrc, attribute2.text as attributeDest, relation.text as relation, (case when taxonomy_relation.edgeBendPoints IS NOT NULL then taxonomy_relation.edgeBendPoints else \"\" end), annotation.annotation from attribute as attribute1 inner join taxonomy_relation on (attribute1.id_taxonomy = ? and taxonomy_relation.id_dimension = (select distinct id_dimension from dimension where text = \"Interdimensional view\" and dimension.id_taxonomy = ?) and attribute1.id_attribute = taxonomy_relation.id_src_attribute) inner join attribute as attribute2 on (attribute2.id_taxonomy = ? and taxonomy_relation.id_dest_attribute = attribute2.id_attribute) inner join relation on (taxonomy_relation.id_relation = relation.id_relation) left outer join taxonomy_relation_annotation as annotation on (taxonomy_relation.id_taxonomy_relation = annotation.id_taxonomy_relation);")
		defer stmt.Close()
		defer db.Close()
		rows, err := stmt.Query(taxonomyIdStr, taxonomyIdStr, taxonomyIdStr)
		checkErr(err)
		for rows.Next() {
			a := model.AttributeRelation{}
			rows.Scan(&a.AttributeSrc,&a.AttributeDest,&a.Relation,&a.EdgeBendPoints,&a.Annotation)
			attributeRelations = append(attributeRelations, a)
		}
		defer rows.Close()
		return attributeRelations, err
		}

	func (d MySQLDriver) SavePositions(taxonomyId int64, positions []model.Position) (result model.Result, err error){
	  	dbRef, err := d.OpenDB()
		defer dbRef.Close()
		checkErr(err)
		taxonomyIdStr := strconv.Itoa(int(taxonomyId))
    	for _, elem := range positions {
			dbRef.Exec("update " + elem.Table + " set x = ?, y = ? where text = ? and id_taxonomy = ?;", elem.X, elem.Y, elem.ID, taxonomyIdStr)
		}
		result.Success = true
		return result, err
		}

	func (d MySQLDriver) SaveMajorPositions(taxonomyId int64, positions []model.Position) (result model.Result, err error){
	  	dbRef, err := d.OpenDB()
		defer dbRef.Close()
		checkErr(err)
		taxonomyIdStr := strconv.Itoa(int(taxonomyId))
    	for _, elem := range positions {
			dbRef.Exec("update " + elem.Table + " set xMajor = ?, yMajor = ? where text = ? and id_taxonomy = ?;", elem.X, elem.Y, elem.ID, taxonomyIdStr)
		}
		result.Success = true
		return result, err
		}

	func (d MySQLDriver) Save3DPositions(taxonomyId int64, positions []model.Position) (result model.Result, err error){
	  	dbRef, err := d.OpenDB()
		defer dbRef.Close()
		checkErr(err)
		taxonomyIdStr := strconv.Itoa(int(taxonomyId))
    	for _, elem := range positions {
			dbRef.Exec("update " + elem.Table + " set x3D = ?, y3D = ?, z3D = ? where text = ? and id_taxonomy = ?;", elem.X, elem.Y, elem.Z, elem.ID, taxonomyIdStr)
		}
		result.Success = true
		return result, err
		}

	func (d MySQLDriver) SaveMajor3DPositions(taxonomyId int64, positions []model.Position) (result model.Result, err error){
	  	dbRef, err := d.OpenDB()
		defer dbRef.Close()
		checkErr(err)
		taxonomyIdStr := strconv.Itoa(int(taxonomyId))
    	for _, elem := range positions {
			dbRef.Exec("update " + elem.Table + " set xMajor3D = ?, yMajor3D = ?, zMajor3D = ? where text = ? and id_taxonomy = ?;", elem.X, elem.Y, elem.Z, elem.ID, taxonomyIdStr)
		}
		result.Success = true
		return result, err
		}

	func (d MySQLDriver) SaveEdgeBendPoints(taxonomyId int64, attributeSrc string, attributeDest string, edgeBendPoints string, dimension string) (result model.Result, err error){
	  	dbRef, err := d.OpenDB()
		defer dbRef.Close()
		checkErr(err)
		taxonomyIdStr := strconv.Itoa(int(taxonomyId))
		dbRef.Exec("update taxonomy_relation set edgeBendPoints = ? where id_taxonomy = ? and id_src_attribute = (select distinct id_attribute from attribute where text = ? and id_taxonomy = ?) and id_dest_attribute = (select distinct id_attribute from attribute where text = ? and id_taxonomy = ?) and id_dimension = (select distinct id_dimension from dimension where text = ? and id_taxonomy = ?);", edgeBendPoints, taxonomyIdStr, attributeSrc, taxonomyIdStr, attributeDest, taxonomyIdStr, dimension, taxonomyIdStr)
		result.Success = true
		return result, err
		}

	func (d MySQLDriver) RemoveTaxonomyRelationsForAttribute(taxonomyId int64, attribute model.Attribute) (result model.Result, err error){
		dbRef, err := d.OpenDB()
		defer dbRef.Close()
		checkErr(err)
		taxonomyIdStr := strconv.Itoa(int(taxonomyId))
		attributeIDStr := strconv.Itoa(attribute.ID)
		dbRef.Exec("DELETE FROM taxonomy_relation_annotation WHERE id_taxonomy_relation IN (SELECT id_taxonomy_relation FROM taxonomy_relation WHERE id_taxonomy = ? AND (id_src_attribute = ? OR id_dest_attribute = ?));", taxonomyIdStr, attributeIDStr, attributeIDStr)
		dbRef.Exec("DELETE FROM taxonomy_relation WHERE id_taxonomy = ? AND (id_src_attribute = ? OR id_dest_attribute = ?);", taxonomyIdStr, attributeIDStr, attributeIDStr)
		d.UpdateRelationshipTables(taxonomyId)
		result.Success = true
		return result, err
		}

	func (d MySQLDriver) UpdateChildRelationshipTable(taxonomyId int64) (resultChan chan model.Result){
		dbRef, err := d.OpenDB()
		checkErr(err)
		taxonomyIdStr := strconv.Itoa(int(taxonomyId))
		future := make(chan model.Result)
		go func () {
			dbRef.Exec("CALL insertallchildrenperattribute(?);", taxonomyIdStr)
			defer dbRef.Close()
			future <- model.Result{Success: true}
		}()
		return future
		}

	func (d MySQLDriver) UpdateParentRelationshipTable(taxonomyId int64) (resultChan chan model.Result){
		dbRef, err := d.OpenDB()
		checkErr(err)
		taxonomyIdStr := strconv.Itoa(int(taxonomyId))
		future := make(chan model.Result)
		go func () {
			dbRef.Exec("CALL insertallparentsperattribute(?);", taxonomyIdStr)
			defer dbRef.Close()
			future <- model.Result{Success: true}
		}()
		return future
		}

	func (d MySQLDriver) UpdateRelationshipTables(taxonomyId int64) (result model.Result){
		resultFuture1 := d.UpdateParentRelationshipTable(taxonomyId)
		resultFuture2 := d.UpdateChildRelationshipTable(taxonomyId)
		result1 := <-resultFuture1
		result2 := <-resultFuture2
		result.Success = result1.Success && result2.Success
		return result
		}

	func (d MySQLDriver) AddAttribute(taxonomyId int64, attribute model.Attribute) (result model.Result, err error){
		dbRef, err := d.OpenDB()
		defer dbRef.Close()
		checkErr(err)
		taxonomyIdStr := strconv.Itoa(int(taxonomyId))
		db, stmt, err := d.Query("SELECT count(id_attribute) FROM attribute WHERE text = ? AND id_taxonomy = ?;")
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
		dbRef.Exec("INSERT IGNORE INTO taxonomy_dimension (id_taxonomy, id_attribute, id_dimension) VALUES (?, (SELECT DISTINCT id_attribute FROM attribute WHERE text = ? AND id_taxonomy = ?), (SELECT DISTINCT id_dimension FROM dimension WHERE text = ? AND id_taxonomy = ?));", taxonomyIdStr, attribute.Text, taxonomyIdStr, attribute.Dimension, taxonomyIdStr)
		d.UpdateRelationshipTables(taxonomyId)
		result.Success = true
		return result, err
		}

	func (d MySQLDriver) AddDimension(taxonomyId int64, dimension string) (result model.Result, err error){
		dbRef, err := d.OpenDB()
		defer dbRef.Close()
		checkErr(err)
		taxonomyIdStr := strconv.Itoa(int(taxonomyId))
		db, stmt, err := d.Query("SELECT count(id_dimension) FROM dimension WHERE text = ? AND id_taxonomy = ?;")
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

	func (d MySQLDriver) ChangeDimension(taxonomyId int64, attribute string, dimension string) (result model.Result, err error){
		dbRef, err := d.OpenDB()
		defer dbRef.Close()
		checkErr(err)
		taxonomyIdStr := strconv.Itoa(int(taxonomyId))
		dbRef.Exec("UPDATE taxonomy_dimension SET id_dimension = (SELECT DISTINCT id_dimension FROM dimension WHERE text = ?) WHERE id_taxonomy = ? AND id_attribute = (SELECT DISTINCT id_attribute FROM attribute WHERE text = ?);", dimension, taxonomyIdStr, attribute)
		dbRef.Exec("DELETE FROM taxonomy_relation WHERE id_taxonomy = ? AND (id_src_attribute = (SELECT DISTINCT id_attribute FROM attribute WHERE text = ?) OR id_dest_attribute = (SELECT DISTINCT id_attribute FROM attribute WHERE text = ?));", taxonomyIdStr, attribute, attribute)
		result.Success = true
		return result, err
		}

	func (d MySQLDriver) RenameAttribute(taxonomyId int64, previousName string, newName string) (result model.Result, err error){
		dbRef, err := d.OpenDB()
		defer dbRef.Close()
		checkErr(err)
		taxonomyIdStr := strconv.Itoa(int(taxonomyId))
		db, stmt, err := d.Query("SELECT count(id_attribute) FROM attribute WHERE text = ? AND id_taxonomy = ?;")
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
		dbRef.Exec("UPDATE attribute SET text = ? WHERE text = ? AND id_taxonomy = ?;", newName, previousName, taxonomyIdStr)
		d.UpdateRelationshipTables(taxonomyId)
		result.Success = true
		return result, err
		}

	func (d MySQLDriver) UpdateSynonyms(taxonomyId int64, attribute string, synonyms string) (result model.Result, err error){
		dbRef, err := d.OpenDB()
		defer dbRef.Close()
		checkErr(err)
		taxonomyIdStr := strconv.Itoa(int(taxonomyId))
		dbRef.Exec("UPDATE attribute SET synonyms = ? WHERE text = ? AND id_taxonomy = ?;", synonyms, attribute, taxonomyIdStr)
		result.Success = true
		return result, err
		}		

	func (d MySQLDriver) RenameDimension(taxonomyId int64, previousName string, newName string) (result model.Result, err error){
		dbRef, err := d.OpenDB()
		defer dbRef.Close()
		checkErr(err)
		taxonomyIdStr := strconv.Itoa(int(taxonomyId))
		db, stmt, err := d.Query("SELECT count(id_dimension) FROM dimension WHERE text = ? AND id_taxonomy = ?;")
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
		dbRef.Exec("UPDATE dimension SET text = ? WHERE text = ? AND id_taxonomy = ?;", newName, previousName, taxonomyIdStr)
		result.Success = true
		return result, err
		}

	func (d MySQLDriver) MergeAttributes(taxonomyId int64, attribute1 model.Attribute, attribute2 model.Attribute) (result model.Result, err error){
		dbRef, err := d.OpenDB()
		defer dbRef.Close()
		checkErr(err)
		taxonomyIdStr := strconv.Itoa(int(taxonomyId))
		db, stmt, err := d.Query("SELECT count(id_attribute) FROM attribute WHERE text = ? AND id_taxonomy = ?;")
		defer stmt.Close()
		defer db.Close()
		rows, err := stmt.Query(attribute1.Text + ":" + attribute2.Text, taxonomyIdStr)
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
		dbRef.Exec("INSERT IGNORE INTO attribute (id_taxonomy, text, x, y, xMajor, yMajor, major) SELECT attribute.id_taxonomy, ? as newName, attribute.x, attribute.y, attribute.xMajor, attribute.yMajor, attribute.major FROM attribute WHERE attribute.text = ? AND attribute.id_taxonomy = ?;", attribute1.Text + ":" + attribute2.Text, attribute1.Text, taxonomyIdStr)
		dbRef.Exec("INSERT IGNORE INTO taxonomy_dimension (id_taxonomy, id_attribute, id_dimension) VALUES (?, (SELECT DISTINCT id_attribute FROM attribute WHERE text = ? AND id_taxonomy = ?), (SELECT DISTINCT id_dimension FROM dimension WHERE text = ? AND id_taxonomy = ?));", taxonomyIdStr, attribute1.Text + ":" + attribute2.Text, taxonomyIdStr, attribute1.Dimension, taxonomyIdStr)
		dbRef.Exec("INSERT IGNORE INTO taxonomy_relation (id_taxonomy, id_src_attribute, id_dest_attribute, id_relation, id_dimension) SELECT id_taxonomy, (SELECT DISTINCT id_attribute FROM attribute WHERE text = ? AND id_taxonomy = ?), id_dest_attribute, id_relation, id_dimension FROM taxonomy_relation WHERE id_src_attribute = (SELECT id_attribute FROM attribute WHERE text = ? AND id_taxonomy = ?);", attribute1.Text + ":" + attribute2.Text, taxonomyIdStr, attribute1.Text, taxonomyIdStr)
		dbRef.Exec("INSERT IGNORE INTO taxonomy_relation (id_taxonomy, id_src_attribute, id_dest_attribute, id_relation, id_dimension) SELECT id_taxonomy, (SELECT DISTINCT id_attribute FROM attribute WHERE text = ? AND id_taxonomy = ?), id_dest_attribute, id_relation, id_dimension FROM taxonomy_relation WHERE id_src_attribute = (SELECT id_attribute FROM attribute WHERE text = ? AND id_taxonomy = ?);", attribute1.Text + ":" + attribute2.Text, taxonomyIdStr, attribute2.Text, taxonomyIdStr)
		dbRef.Exec("INSERT IGNORE INTO taxonomy_relation (id_taxonomy, id_src_attribute, id_dest_attribute, id_relation, id_dimension) SELECT id_taxonomy, id_src_attribute, (SELECT DISTINCT id_attribute FROM attribute WHERE text = ? AND id_taxonomy = ?), id_relation, id_dimension FROM taxonomy_relation WHERE id_dest_attribute = (SELECT id_attribute FROM attribute WHERE text = ? AND id_taxonomy = ?);", attribute1.Text + ":" + attribute2.Text, taxonomyIdStr, attribute1.Text, taxonomyIdStr)
		dbRef.Exec("INSERT IGNORE INTO taxonomy_relation (id_taxonomy, id_src_attribute, id_dest_attribute, id_relation, id_dimension) SELECT id_taxonomy, id_src_attribute, (SELECT DISTINCT id_attribute FROM attribute WHERE text = ? AND id_taxonomy = ?), id_relation, id_dimension FROM taxonomy_relation WHERE id_dest_attribute = (SELECT id_attribute FROM attribute WHERE text = ? AND id_taxonomy = ?);", attribute1.Text + ":" + attribute2.Text, taxonomyIdStr, attribute2.Text, taxonomyIdStr)
		dbRef.Exec("INSERT IGNORE INTO mapping (id_paper, id_attribute) SELECT id_paper, (SELECT DISTINCT id_attribute FROM attribute WHERE text = ? AND id_taxonomy = ?) FROM mapping WHERE id_attribute = (SELECT id_attribute FROM attribute WHERE text = ? AND id_taxonomy = ?);", attribute1.Text + ":" + attribute2.Text, taxonomyIdStr, attribute1.Text, taxonomyIdStr)
		dbRef.Exec("INSERT IGNORE INTO mapping (id_paper, id_attribute) SELECT id_paper, (SELECT DISTINCT id_attribute FROM attribute WHERE text = ? AND id_taxonomy = ?) FROM mapping WHERE id_attribute = (SELECT id_attribute FROM attribute WHERE text = ? AND id_taxonomy = ?);", attribute1.Text + ":" + attribute2.Text, taxonomyIdStr, attribute2.Text, taxonomyIdStr)
		dbRef.Exec("DELETE FROM attribute WHERE id_taxonomy = ? AND (text = ? OR text = ?);", taxonomyIdStr, attribute1.Text, attribute2.Text)
		d.UpdateRelationshipTables(taxonomyId)
		result.Success = true
		return result, err
		}

	func (d MySQLDriver) ForkAttribute(taxonomyId int64, attribute string, dimension string, parents1 []model.AttributeRelation, parents2 []model.AttributeRelation, children1 []model.AttributeRelation, children2 []model.AttributeRelation, citations1 []model.Paper, citations2 []model.Paper) (result model.Result, err error){
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
			db, stmt, err := d.Query("SELECT count(id_attribute) FROM attribute WHERE text = ? AND id_taxonomy = ?;")
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
			db, stmt, err := d.Query("SELECT count(id_attribute) FROM attribute WHERE text = ? AND id_taxonomy = ?;")
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
		dbRef.Exec("INSERT IGNORE INTO taxonomy_dimension (id_taxonomy, id_attribute, id_dimension) VALUES (?, (SELECT DISTINCT id_attribute FROM attribute WHERE text = ? AND id_taxonomy = ?), (SELECT DISTINCT id_dimension from dimension where text = ? AND id_taxonomy = ?));", taxonomyIdStr, newAttributeName, taxonomyIdStr, dimension, taxonomyIdStr)
		dbRef.Exec("INSERT IGNORE INTO attribute (id_taxonomy, text, x, y, xMajor, yMajor, major) SELECT attribute.id_taxonomy, ? as newName, attribute.x + 150, attribute.y, attribute.xMajor + 150, attribute.yMajor, attribute.major FROM attribute WHERE attribute.text = ? AND attribute.id_taxonomy = ?;", newAttributeName2, attribute, taxonomyIdStr)
		dbRef.Exec("INSERT IGNORE INTO taxonomy_dimension (id_taxonomy, id_attribute, id_dimension) VALUES (?, (SELECT DISTINCT id_attribute FROM attribute WHERE text = ? AND id_taxonomy = ?), (SELECT DISTINCT id_dimension from dimension where text = ? AND id_taxonomy = ?));", taxonomyIdStr, newAttributeName2, taxonomyIdStr, dimension, taxonomyIdStr)
	    for _, parent := range parents1 {
	        dbRef.Exec("INSERT IGNORE INTO taxonomy_relation (id_taxonomy, id_src_attribute, id_dest_attribute, id_relation, id_dimension) VALUES (?, (select distinct id_attribute from attribute where text = ? AND id_taxonomy = ?), (select distinct id_attribute from attribute where text = ? AND id_taxonomy = ?), (select distinct id_relation from relation where text = ?), (select distinct id_dimension from dimension where text = ? AND id_taxonomy = ?));", taxonomyIdStr, newAttributeName, taxonomyIdStr, parent.Text, taxonomyIdStr, parent.Relation, dimension, taxonomyIdStr)
	    }
	    for _, parent := range parents2 {
	        dbRef.Exec("INSERT IGNORE INTO taxonomy_relation (id_taxonomy, id_src_attribute, id_dest_attribute, id_relation, id_dimension) VALUES (?, (select distinct id_attribute from attribute where text = ? AND id_taxonomy = ?), (select distinct id_attribute from attribute where text = ? AND id_taxonomy = ?), (select distinct id_relation from relation where text = ?), (select distinct id_dimension from dimension where text = ? AND id_taxonomy = ?));", taxonomyIdStr, newAttributeName2, taxonomyIdStr, parent.Text, taxonomyIdStr, parent.Relation, dimension, taxonomyIdStr)
	    }
	    for _, child := range children1 {
	        dbRef.Exec("INSERT IGNORE INTO taxonomy_relation (id_taxonomy, id_dest_attribute, id_src_attribute, id_relation, id_dimension) VALUES (?, (select distinct id_attribute from attribute where text = ? AND id_taxonomy = ?), (select distinct id_attribute from attribute where text = ? AND id_taxonomy = ?), (select distinct id_relation from relation where text = ?), (select distinct id_dimension from dimension where text = ? AND id_taxonomy = ?));", taxonomyIdStr, newAttributeName, taxonomyIdStr, child.Text, taxonomyIdStr, child.Relation, dimension, taxonomyIdStr)
	    }
	    for _, child := range children2 {
	        dbRef.Exec("INSERT IGNORE INTO taxonomy_relation (id_taxonomy, id_dest_attribute, id_src_attribute, id_relation, id_dimension) VALUES (?, (select distinct id_attribute from attribute where text = ? AND id_taxonomy = ?), (select distinct id_attribute from attribute where text = ? AND id_taxonomy = ?), (select distinct id_relation from relation where text = ?), (select distinct id_dimension from dimension where text = ? AND id_taxonomy = ?));", taxonomyIdStr, newAttributeName2, taxonomyIdStr, child.Text, taxonomyIdStr, child.Relation, dimension, taxonomyIdStr)
	    }
	    for _, citation := range citations1 {
	        dbRef.Exec("INSERT IGNORE INTO mapping (id_paper, id_attribute) VALUES (?, (select distinct id_attribute from attribute where text = ? AND id_taxonomy = ?));", strconv.Itoa(int(citation.ID)), newAttributeName, taxonomyIdStr)
	    }
	    for _, citation := range citations2 {
	        dbRef.Exec("INSERT IGNORE INTO mapping (id_paper, id_attribute) VALUES (?, (select distinct id_attribute from attribute where text = ? AND id_taxonomy = ?));", strconv.Itoa(int(citation.ID)), newAttributeName2, taxonomyIdStr)
	    }
		dbRef.Exec("DELETE FROM attribute WHERE text = ? AND id_taxonomy = ?;", attribute, taxonomyIdStr)
		d.UpdateRelationshipTables(taxonomyId)
		result.Success = true
		return result, err
		}

	func (d MySQLDriver) AddTaxonomyRelation(relation model.AttributeRelation) (result model.Result, err error){
		dbRef, err := d.OpenDB()
		defer dbRef.Close()
		checkErr(err)
		taxonomyIdStr := strconv.Itoa(int(relation.TaxonomyID))
		dbRef.Exec("INSERT IGNORE INTO taxonomy_relation (id_taxonomy, id_src_attribute, id_dest_attribute, id_relation, id_dimension) VALUES (?, (SELECT id_attribute FROM attribute WHERE text = ? AND id_taxonomy = ?), (SELECT id_attribute FROM attribute WHERE text = ? AND id_taxonomy = ?), (SELECT id_relation FROM relation WHERE text = ?), (SELECT id_dimension FROM dimension WHERE text = ? AND id_taxonomy = ?));", taxonomyIdStr, relation.AttributeSrc, taxonomyIdStr, relation.AttributeDest, taxonomyIdStr, relation.Relation, relation.Dimension, taxonomyIdStr)
		d.UpdateRelationshipTables(relation.TaxonomyID)
		result.Success = true
		return result, err
		}

	func (d MySQLDriver) RemoveTaxonomyRelation(relation model.AttributeRelation) (result model.Result, err error){
		dbRef, err := d.OpenDB()
		defer dbRef.Close()
		checkErr(err)
		taxonomyIdStr := strconv.Itoa(int(relation.TaxonomyID))
		dbRef.Exec("DELETE FROM taxonomy_relation_annotation WHERE id_taxonomy_relation IN (SELECT id_taxonomy_relation FROM taxonomy_relation WHERE id_taxonomy = ? AND id_src_attribute = (SELECT id_attribute FROM attribute WHERE text = ? AND id_taxonomy = ?) AND id_dest_attribute = (SELECT id_attribute FROM attribute WHERE text = ? AND id_taxonomy = ?) AND id_relation = (SELECT id_relation FROM relation WHERE text = ?) AND id_dimension = (SELECT id_dimension FROM dimension WHERE text = ? AND id_taxonomy = ?);", taxonomyIdStr, relation.AttributeSrc, taxonomyIdStr, relation.AttributeDest, taxonomyIdStr, relation.Relation, relation.Dimension, taxonomyIdStr)
		dbRef.Exec("DELETE FROM taxonomy_relation WHERE id_taxonomy = ? AND id_src_attribute = (SELECT id_attribute FROM attribute WHERE text = ? AND id_taxonomy = ?) AND id_dest_attribute = (SELECT id_attribute FROM attribute WHERE text = ? AND id_taxonomy = ?) AND id_relation = (SELECT id_relation FROM relation WHERE text = ?) AND id_dimension = (SELECT id_dimension FROM dimension WHERE text = ? AND id_taxonomy = ?);", taxonomyIdStr, relation.AttributeSrc, taxonomyIdStr, relation.AttributeDest, taxonomyIdStr, relation.Relation, relation.Dimension, taxonomyIdStr)
		d.UpdateRelationshipTables(relation.TaxonomyID)
		result.Success = true
		return result, err
		}

	func (d MySQLDriver) UpdateTaxonomyRelationType(relation model.AttributeRelation) (result model.Result, err error){
		dbRef, err := d.OpenDB()
		defer dbRef.Close()
		checkErr(err)
		taxonomyIdStr := strconv.Itoa(int(relation.TaxonomyID))
		dbRef.Exec("UPDATE taxonomy_relation SET id_relation = (SELECT id_relation FROM relation WHERE text = ?) WHERE id_taxonomy = ? AND id_src_attribute = (SELECT id_attribute FROM attribute WHERE text = ? AND id_taxonomy = ?) AND id_dest_attribute = (SELECT id_attribute FROM attribute WHERE text = ? AND id_taxonomy = ?);", relation.Relation, taxonomyIdStr, relation.AttributeSrc, taxonomyIdStr, relation.AttributeDest, taxonomyIdStr) //  AND id_dimension = (SELECT id_dimension FROM dimension WHERE text = \"" + relation.Dimension + "\" AND id_taxonomy = " + taxonomyIdStr + ")
		go func () {
			d.UpdateRelationshipTables(relation.TaxonomyID)
		}()
		result.Success = true
		return result, err
		}

	func (d MySQLDriver) UpdateTaxonomyRelationAnnotation(relation model.AttributeRelation) (result model.Result, err error){
		dbRef, err := d.OpenDB()
		defer dbRef.Close()
		checkErr(err)
		taxonomyIdStr := strconv.Itoa(int(relation.TaxonomyID))
		dbRef.Exec("REPLACE INTO taxonomy_relation_annotation (id_taxonomy, id_taxonomy_relation, annotation) VALUES (?, (SELECT DISTINCT id_taxonomy_relation FROM taxonomy_relation WHERE id_taxonomy = ? AND id_src_attribute = (SELECT id_attribute FROM attribute WHERE text = ? AND id_taxonomy = ?) AND id_dest_attribute = (SELECT id_attribute FROM attribute WHERE text = ? AND id_taxonomy = ?)), ?);", taxonomyIdStr, taxonomyIdStr, relation.AttributeSrc, taxonomyIdStr, relation.AttributeDest, taxonomyIdStr, relation.Annotation) //  AND id_dimension = (SELECT id_dimension FROM dimension WHERE text = \"" + relation.Dimension + "\" AND id_taxonomy = " + taxonomyIdStr + ")
		result.Success = true
		return result, err
		}

	func (d MySQLDriver) DeleteCitation(taxonomyId int64, paper model.Paper) (result model.Result, err error){
		dbRef, err := d.OpenDB()
		defer dbRef.Close()
		checkErr(err)
		taxonomyIdStr := strconv.Itoa(int(taxonomyId))
		dbRef.Exec("DELETE FROM paper WHERE id_taxonomy = ? AND citation = ?;", taxonomyIdStr, paper.Citation)
		result.Success = true
		return result, err
		}

	func (d MySQLDriver) RemoveAttribute(taxonomyId int64, attribute model.Attribute) (result model.Result, err error){
		dbRef, err := d.OpenDB()
		defer dbRef.Close()
		checkErr(err)
		taxonomyIdStr := strconv.Itoa(int(taxonomyId))
		db, stmt, err := d.Query("SELECT id_attribute FROM attribute WHERE text = ? AND id_taxonomy = ?;")
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

	func (d MySQLDriver) RemoveDimension(taxonomyId int64, dimension model.Dimension) (result model.Result, err error){
		dbRef, err := d.OpenDB()
		defer dbRef.Close()
		checkErr(err)
		taxonomyIdStr := strconv.Itoa(int(taxonomyId))
		db, stmt, err := d.Query("SELECT id_dimension FROM dimension WHERE text = ? AND id_taxonomy = ?;")
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

	func (d MySQLDriver) SaveReviewMappings(taxonomyId int64, attributes []model.Attribute, mappings []model.ReviewMapping) (result model.Result, err error) {
		dbRef, err := d.OpenDB()
		defer dbRef.Close()
		checkErr(err)
		taxonomyIdStr := strconv.Itoa(int(taxonomyId))
		for _, elem := range mappings {
			db, stmt, err := d.Query("SELECT DISTINCT Title, year, Authors, Keywords, cited_by FROM articles WHERE ArticleId = ?;")
			checkErr(err)
			rows, err := stmt.Query(elem.ArticleID)
			checkErr(err)
			stmt.Close()
			db.Close()
			var article model.Article
			foundArticle := false
			for rows.Next() {
				article = model.Article{}
				rows.Scan(&article.Title,&article.Year,&article.Authors,&article.Keywords,&article.CitedBy)
				foundArticle = true
			}
			rows.Close()
			if (foundArticle) {
				dimension := ""
			    for i := 0; i < len(attributes); i++ {
			        if (attributes[i].Text == elem.Attribute) {
			        	dimension = attributes[i].Dimension
			        	break
			        }
			    }
			    if len(dimension) == 0 {
			    	dimension = "Interdimensional view"
			    }
			    major := 0
			    if dimension == "Interdimensional view" {
			    	major = 1
			    }
				db, stmt, err := d.Query("SELECT DISTINCT id_attribute FROM attribute WHERE id_taxonomy = ? AND text = ?;")
				rows, err := stmt.Query(taxonomyIdStr, elem.Attribute)
				checkErr(err)
				stmt.Close()
				db.Close()
				var attributeID int
				attributeID = -1
				for rows.Next() {
					rows.Scan(&attributeID)
				}
				rows.Close()
				if (attributeID < 0) {
					dbRef.Exec("INSERT IGNORE INTO attribute (id_taxonomy, text, major) VALUES (?, ?, ?);", taxonomyIdStr, elem.Attribute, major)
					dbRef.Exec("INSERT IGNORE INTO taxonomy_dimension (id_taxonomy, id_attribute, id_dimension) VALUES (?, (SELECT DISTINCT id_attribute FROM attribute WHERE id_taxonomy = ? AND text = ?), (SELECT DISTINCT id_dimension FROM dimension WHERE id_taxonomy = ? AND text = ?));", taxonomyIdStr, taxonomyIdStr, elem.Attribute, taxonomyIdStr, dimension)
				}
				db, stmt, err = d.Query("SELECT tmp.maxID, paper.id_paper FROM (SELECT MAX(id_paper) AS maxID FROM paper WHERE id_taxonomy = ?) AS tmp LEFT OUTER JOIN paper ON (paper.citation = ? and paper.id_taxonomy = ?);")
				rows, err = stmt.Query(taxonomyIdStr, article.Title, taxonomyIdStr)
				checkErr(err)
				stmt.Close()
				db.Close()
				var maxID int
				var paperID int
				maxID = 0
				paperID = -1
				for rows.Next() {
					rows.Scan(&maxID,&paperID)
				}
				rows.Close()
				var paperIDStr string
				if paperID < 0 {
					paperID = maxID+1
					paperIDStr = strconv.Itoa(paperID)
					dbRef.Exec("INSERT IGNORE INTO paper (id_taxonomy, id_paper, citation, author, keywords, referenceCount) VALUES (?, ?, ?, ?, ?, ?);", taxonomyIdStr, paperIDStr, article.Title, article.Authors, article.Keywords, article.CitedBy) // bibTex
				} else {
					paperIDStr = strconv.Itoa(paperID)
				}
				dbRef.Exec("INSERT IGNORE INTO mapping (id_paper, id_attribute) VALUES (?, (SELECT DISTINCT id_attribute FROM attribute WHERE id_taxonomy = ? AND text = ?));", paperIDStr, taxonomyIdStr, elem.Attribute)
			}
		}
		result.Success = true
		return result, err
	}

	func (d MySQLDriver) KMeans(taxonomyId int64, n int) (clusters []model.Cluster, err error){
		dbRef, err := d.OpenDB()
		defer dbRef.Close()
		checkErr(err)
		if n <= 0 {
			return clusters, err
		}
		taxonomyIdStr := strconv.Itoa(int(taxonomyId))
		db, stmt, err := d.Query("SELECT id_paper, citation FROM paper WHERE id_taxonomy = ?;")
		checkErr(err)
		rows, err := stmt.Query(taxonomyIdStr)
		checkErr(err)
		stmt.Close()
		db.Close()
		papers := []model.Paper{}
		for rows.Next() {
			a := model.Paper{}
			rows.Scan(&a.ID,&a.Citation)
			papers = append(papers, a)
		}
		rows.Close()
		// shuffle papers
	    for k := len(papers) - 1; k > 0; k-- {
	        l := rand.Intn(k + 1)
	        papers[k], papers[l] = papers[l], papers[k]
	    }
		i := int(len(papers) / n)
		for j := 0; j < n; j++ {
			a := model.Cluster{ID: j}
			a.Papers = []int{}
			a.Attributes = []string{}
			clusters = append(clusters, a)
		}
		index := 0
		counter := 0
		for _, elem := range papers {
			clusters[index].Papers = append(clusters[index].Papers, elem.ID)
			counter++
			if counter >= i && index < n-1 {
				index++
				counter = 0
			}
		}
		db, stmt, err = d.Query("SELECT id_attribute, text FROM attribute WHERE id_taxonomy = ?;")
		checkErr(err)
		defer stmt.Close()
		defer db.Close()
		rows, err = stmt.Query(taxonomyIdStr)
		checkErr(err)
		attributes := []model.Attribute{}
		for rows.Next() {
			a := model.Attribute{}
			rows.Scan(&a.ID,&a.Text)
			attributes = append(attributes, a)
		}
		defer rows.Close()
		for i, _ := range attributes {
			idStr := strconv.Itoa(attributes[i].ID)
			db, stmt, err = d.Query("SELECT id_paper FROM mapping WHERE id_attribute = ?;")
			checkErr(err)
			rows, err = stmt.Query(idStr)
			checkErr(err)
			stmt.Close()
			db.Close()
			attributes[i].PaperIDs = []int{}
			for rows.Next() {
				var paperID int
				rows.Scan(&paperID)
				attributes[i].PaperIDs = append(attributes[i].PaperIDs, paperID)
			}
			defer rows.Close()
		}
		// sort attributes
		sort.Slice(attributes, func(i, j int) bool {
		    return len(attributes[i].PaperIDs) < len(attributes[j].PaperIDs)
		})
		for _, elem := range attributes {
			max := 0
			clusterIndex := 0
			for clusterID, cluster := range clusters {
				count := 0
				for _, id := range cluster.Papers {
					for _, paperID := range elem.PaperIDs {
						if id == paperID {
							count++
							break
						}
					}
				}
				if count > max {
					max = count
					clusterIndex = clusterID
				}
			}
			clusters[clusterIndex].Attributes = append(clusters[clusterIndex].Attributes, elem.Text)
			for _, paperID := range elem.PaperIDs {
				found := false
				for _, clusterPaperID := range clusters[clusterIndex].Papers {
					if paperID == clusterPaperID {
						found = true
						break
					}
				}
				if !found {
					clusters[clusterIndex].Papers = append(clusters[clusterIndex].Papers, paperID)
				}
			}
		}
		return clusters, err
		}