package data
import (
	"strconv"
	"fmt"
	//overriding MySqlDriver
	_ "../mysql"
		"../model"
)

type ClassificationDriver interface {
  DriverCore
	ExportCorrelations([]model.Attribute, int64) ([]model.Paper, error)
	GetAttributesPerDimension(int64, string) ([]model.Attribute, error)
	GetAllChildrenAttributes(int64, string) ([]model.Attribute, error)
	GetIntermediateAttributes(int64, int64, int64) ([]model.Attribute, error)
	GetMajorAttributes(int64) ([]model.Attribute, error)
	GetAttributeRelationsPerDimension(int64, string) ([]model.AttributeRelation, error)
	GetInterdimensionalRelations(int64) ([]model.AttributeRelation, error)
	GetCitationCount() ([]model.CitationCount, error)
	GetCitationCounts() ([]model.CitationCount, error)
	GetCitationCountsIncludingChildren() ([]model.CitationCount, error)
	GetRelationTypes() ([]model.RelationType, error)
	GetCitationsPerAttribute(string) ([]model.Paper, error)
	GetCitationsPerAttributeIncludingChildren(string) ([]model.Paper, error)
	SavePositions([]model.Position) (error)
	SaveMajorPositions([]model.Position) (error)
	Save3DPositions([]model.Position) (error)
	SaveMajor3DPositions([]model.Position) (error)
	SaveEdgeBendPoints(string, string, string) (model.Result, error)
	GetAllAttributes() ([]model.Attribute, error)
	GetAllDimensions() ([]model.Dimension, error)
	GetAllCitations() ([]model.Paper, error)
	GetConceptRelations() ([]model.ConceptCorrelation, error)
	GetAllConceptRelations() ([]model.ConceptCorrelation, error)
	GetSharedPapers(string, string) ([]model.Paper, error)
	GetSharedPapersIncludingChildren(string, string) ([]model.Paper, error)
	GetAttributeDetails(string) ([]model.Attribute, error)
	GetCitationDetails(string, string) ([]model.Paper, error)
	GetAttributeCoverage() ([]model.AttributeCoverage, error)
	ExportCorrelationsCSV(filterAttributes []model.Attribute)
	AddAttribute(model.Attribute) (model.Result, error)
	RenameAttribute(string, string) (model.Result, error)
	AddTaxonomyRelation(model.AttributeRelation) (model.Result, error)
	RemoveAttribute(model.Attribute) (model.Result, error)
	RemoveTaxonomyRelation(model.AttributeRelation) (model.Result, error)
	UpdateTaxonomyRelationType(model.AttributeRelation) (model.Result, error)
	UpdateTaxonomyRelationAnnotation(model.AttributeRelation) (model.Result, error)
	RemoveTaxonomyRelationsForAttribute(model.Attribute) (model.Result, error)
	MergeAttributes(model.Attribute, model.Attribute) (model.Result, error)
}


//InitMySQLDriver initialize a new my sql driver instance
func InitClassificationDriver(user string, password string) ClassificationDriver {
	return MySQLDriver{username: user, pass: password, database: "classification"}
}

//ExportCorrelations export correlations with the given attributes
func (d MySQLDriver) ExportCorrelations(filterAttributes []model.Attribute,
	taxonomyId int64) (papers []model.Paper, err error) {
	db, err := d.OpenDB()
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
	return papers, err
}

func (d MySQLDriver) ExportCorrelationsCSV(filterAttributes []model.Attribute){

}


func (d MySQLDriver) GetAllAttributes() (attributes []model.Attribute,
	err error){
	db, err := d.OpenDB()
	checkErr(err)
	db, stmt, err := d.Query(`select distinct attribute1.id_attribute as id1, attribute1.text as attr1, relation.parentID as parentID, relation.parentText as parentText
		from attribute as attribute1 left outer join (select distinct taxonomy_relation.id_src_attribute, attribute2.id_attribute as parentID, attribute2.text as parentText from taxonomy_relation inner join attribute as attribute2 on (taxonomy_relation.id_dest_attribute = attribute2.id_attribute and taxonomy_relation.id_relation > 2)) as relation on (attribute1.id_attribute = relation.id_src_attribute) order by attribute1.id_attribute;`)
	defer stmt.Close()
	defer db.Close()
	rows, err := stmt.Query()
	checkErr(err)
	for rows.Next() {
		a := model.Attribute{}
		rows.Scan(&a.ID,&a.Text,&a.ParentID,&a.ParentText)
		attributes = append(attributes, a)
	}
	return attributes, err

	}

func (d MySQLDriver) GetAllDimensions() (dimensions []model.Dimension,
	err error){
	db, err := d.OpenDB()
	checkErr(err)
	db, stmt, err := d.Query(`select id_dimension,text,xMajor,yMajor
		from dimension where text != "Interdimensional view" order by id_dimension`)
	defer stmt.Close()
	defer db.Close()
	rows, err := stmt.Query()
	checkErr(err)
	for rows.Next() {
		a := model.Dimension{}
		rows.Scan(&a.ID,&a.Text,&a.XMajor,&a.YMajor)
		dimensions = append(dimensions, a)
	}
	return dimensions, err

	}

	func (d MySQLDriver) GetAllCitations() (papers []model.Paper,
		err error){
		db, err := d.OpenDB()
		checkErr(err)
		db, stmt, err := d.Query(`select id_paper,citation,citation as text,bib,referenceCount
			from paper order by id_paper`)
		defer stmt.Close()
		defer db.Close()
		rows, err := stmt.Query()
		checkErr(err)
		for rows.Next() {
			a := model.Paper{}
			rows.Scan(&a.ID,&a.Citation,&a.Text,&a.Bib,&a.ReferenceCount)
			papers = append(papers, a)
		}
		return papers, err

		}

	func (d MySQLDriver) GetCitationCount() (citationCounts []model.CitationCount, err error){
		db, err := d.OpenDB()
		checkErr(err)
		db, stmt, err := d.Query("select count(distinct id_paper) as citationCount, max(referenceCount) as maxReferenceCount from paper;")
		defer stmt.Close()
		defer db.Close()
		rows, err := stmt.Query()
		checkErr(err)
		for rows.Next() {
			a := model.CitationCount{}
			rows.Scan(&a.CitationCount,&a.MaxReferenceCount)
			citationCounts = append(citationCounts, a)
		}
		return citationCounts, err
		}

	func (d MySQLDriver) GetCitationCounts() (citationCounts []model.CitationCount, err error){
		db, err := d.OpenDB()
		checkErr(err)
		db, stmt, err := d.Query("select distinct attribute.text, count(distinct mapping.id_paper) as citationCount from attribute left outer join mapping on (attribute.id_attribute = mapping.id_attribute) group by attribute.id_attribute;")
		defer stmt.Close()
		defer db.Close()
		rows, err := stmt.Query()
		checkErr(err)
		for rows.Next() {
			a := model.CitationCount{}
			rows.Scan(&a.Attribute,&a.CitationCount)
			citationCounts = append(citationCounts, a)
		}
		return citationCounts, err
		}

	func (d MySQLDriver) GetCitationCountsIncludingChildren() (citationCounts []model.CitationCount, err error){
		db, err := d.OpenDB()
		checkErr(err)
		db, stmt, err := d.Query("select distinct allchildrenperattribute.text, count(distinct mapping.id_paper) as citationCount from allchildrenperattribute left outer join mapping on (FIND_IN_SET(mapping.id_attribute, allchildrenperattribute.children)) group by allchildrenperattribute.text;")
		defer stmt.Close()
		defer db.Close()
		rows, err := stmt.Query()
		checkErr(err)
		for rows.Next() {
			a := model.CitationCount{}
			rows.Scan(&a.Attribute,&a.CitationCount)
			citationCounts = append(citationCounts, a)
		}
		return citationCounts, err
		}

	func (d MySQLDriver) GetRelationTypes() (relationTypes []model.RelationType, err error){
		db, err := d.OpenDB()
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
		return relationTypes, err
		}


	func (d MySQLDriver) GetCitationsPerAttribute(attribute string) (papers []model.Paper, err error){
		db, err := d.OpenDB()
		checkErr(err)
		db, stmt, err := d.Query("select distinct paper.id_paper, paper.citation, paper.bib, paper.referenceCount from paper inner join mapping on (paper.id_paper = mapping.id_paper) inner join attribute on (mapping.id_attribute = attribute.id_attribute and attribute.text = \"" + attribute +  "\") order by paper.id_paper;")
		defer stmt.Close()
		defer db.Close()
		rows, err := stmt.Query()
		checkErr(err)
		for rows.Next() {
			a := model.Paper{}
			rows.Scan(&a.ID,&a.Citation,&a.Bib,&a.ReferenceCount)
			papers = append(papers, a)
		}
		return papers, err
		}

	func (d MySQLDriver) GetCitationsPerAttributeIncludingChildren(attribute string) (papers []model.Paper, err error){
		db, err := d.OpenDB()
		checkErr(err)
		db, stmt, err := d.Query("select distinct paper.id_paper, paper.citation, paper.bib, paper.referenceCount from allchildrenperattribute inner join mapping on (allchildrenperattribute.text = \"" + attribute + "\" and FIND_IN_SET(mapping.id_attribute, allchildrenperattribute.children)) inner join paper on (mapping.id_paper = paper.id_paper) order by paper.id_paper;")
		defer stmt.Close()
		defer db.Close()
		rows, err := stmt.Query()
		checkErr(err)
		for rows.Next() {
			a := model.Paper{}
			rows.Scan(&a.ID,&a.Citation,&a.Bib,&a.ReferenceCount)
			papers = append(papers, a)
		}
		return papers, err
		}

	func (d MySQLDriver) GetConceptRelations() (conceptCorrelations []model.ConceptCorrelation, err error){
		db, err := d.OpenDB()
		checkErr(err)
		db, stmt, err := d.Query(`select distinct attribute1.text as Attribute1, attribute2.text as Attribute2, attribute1.text as Text1, attribute2.text as Text2, attribute1.id_attribute as ID1, attribute2.id_attribute as ID2, correlation.value from attribute as attribute1 inner join (select distinct mapping1.id_attribute as attr1, mapping2.id_attribute as attr2, count(distinct mapping1.id_paper) as value from mapping as mapping1 inner join mapping as mapping2 on (mapping1.id_paper = mapping2.id_paper and mapping1.id_attribute < mapping2.id_attribute) group by mapping1.id_attribute,mapping2.id_attribute) as correlation on (attribute1.id_attribute = correlation.attr1) inner join attribute as attribute2 on (attribute2.id_attribute = correlation.attr2);`)
		defer stmt.Close()
		defer db.Close()
		rows, err := stmt.Query()
		checkErr(err)
		for rows.Next() {
			a := model.ConceptCorrelation{}
			rows.Scan(&a.Attribute1,&a.Attribute2,&a.Text1,&a.Text2,&a.ID1,&a.ID2,&a.Value)
			conceptCorrelations = append(conceptCorrelations, a)
		}
		return conceptCorrelations, err
		}

	func (d MySQLDriver) GetAllConceptRelations() (conceptCorrelations []model.ConceptCorrelation, err error){
		db, err := d.OpenDB()
		checkErr(err)
		db, stmt, err := d.Query(`select distinct attribute1.text as Attribute1, attribute2.text as Attribute2, attribute1.text as Text1, attribute2.text as Text2, attribute1.id_attribute as ID1, attribute2.id_attribute as ID2, count(distinct mapping1.id_paper) as value from allchildrenperattribute as allchildrenperattribute1 inner join mapping as mapping1 on (FIND_IN_SET(mapping1.id_attribute, allchildrenperattribute1.children)) inner join mapping as mapping2 on (mapping1.id_paper = mapping2.id_paper) inner join allchildrenperattribute as allchildrenperattribute2 on (allchildrenperattribute1.id_attribute <= allchildrenperattribute2.id_attribute and FIND_IN_SET(mapping2.id_attribute, allchildrenperattribute2.children)) inner join attribute as attribute1 on (allchildrenperattribute1.id_attribute = attribute1.id_attribute) inner join attribute as attribute2 on (allchildrenperattribute2.id_attribute = attribute2.id_attribute) group by attribute1.id_attribute, attribute2.id_attribute;`)
		defer stmt.Close()
		defer db.Close()
		rows, err := stmt.Query()
		checkErr(err)
		for rows.Next() {
			a := model.ConceptCorrelation{}
			rows.Scan(&a.Attribute1,&a.Attribute2,&a.Text1,&a.Text2,&a.ID1,&a.ID2,&a.Value)
			conceptCorrelations = append(conceptCorrelations, a)
		}
		return conceptCorrelations, err
		}

	func (d MySQLDriver) GetSharedPapers(text1 string, text2 string) (papers []model.Paper, err error){
		db, err := d.OpenDB()
		checkErr(err)
		db, stmt, err := d.Query("select distinct paper.id_paper, paper.citation, paper.bib, paper.referenceCount from attribute as attribute1 inner join mapping as mapping1 on (attribute1.text = \"" + text1 + "\" and attribute1.id_attribute = mapping1.id_attribute) inner join mapping as mapping2 on (mapping1.id_paper = mapping2.id_paper) inner join attribute as attribute2 on (attribute2.text = \"" + text2 + "\" and attribute2.id_attribute = mapping2.id_attribute) inner join paper on (mapping1.id_paper = paper.id_paper);")
		defer stmt.Close()
		defer db.Close()
		rows, err := stmt.Query()
		checkErr(err)
		for rows.Next() {
			a := model.Paper{}
			rows.Scan(&a.ID,&a.Citation,&a.Bib,&a.ReferenceCount)
			papers = append(papers, a)
		}
		return papers, err
		}

	func (d MySQLDriver) GetSharedPapersIncludingChildren(text1 string, text2 string) (papers []model.Paper, err error){
		db, err := d.OpenDB()
		checkErr(err)
		db, stmt, err := d.Query("select distinct paper.id_paper, paper.citation, paper.bib, paper.referenceCount from allchildrenperattribute as allchildrenperattribute1 inner join mapping as mapping1 on (allchildrenperattribute1.text = \"" + text1 + "\" and FIND_IN_SET(mapping1.id_attribute, allchildrenperattribute1.children)) inner join mapping as mapping2 on (mapping1.id_paper = mapping2.id_paper) inner join allchildrenperattribute as allchildrenperattribute2 on (allchildrenperattribute2.text = \"" + text2 + "\" and FIND_IN_SET(mapping2.id_attribute, allchildrenperattribute2.children)) inner join paper on (mapping1.id_paper = paper.id_paper);")
		defer stmt.Close()
		defer db.Close()
		rows, err := stmt.Query()
		checkErr(err)
		for rows.Next() {
			a := model.Paper{}
			rows.Scan(&a.ID,&a.Citation,&a.Bib,&a.ReferenceCount)
			papers = append(papers, a)
		}
		return papers, err
		}

	func (d MySQLDriver) GetAttributeDetails(text string) (attributes []model.Attribute, err error){
		db, err := d.OpenDB()
		checkErr(err)
		db, stmt, err := d.Query("select distinct attribute1.id_attribute as id1, attribute1.text as attr1, relation.parentID as parentID, relation.parentText as parentText from attribute as attribute1 left outer join (select distinct taxonomy_relation.id_src_attribute, attribute2.id_attribute as parentID, attribute2.text as parentText from taxonomy_relation inner join attribute as attribute2 on (taxonomy_relation.id_dest_attribute = attribute2.id_attribute and taxonomy_relation.id_relation > 2)) as relation on (attribute1.id_attribute = relation.id_src_attribute) where attribute1.text = \"" + text + "\" order by attribute1.id_attribute;")
		defer stmt.Close()
		defer db.Close()
		rows, err := stmt.Query()
		checkErr(err)
		for rows.Next() {
			a := model.Attribute{}
			rows.Scan(&a.ID,&a.Text,&a.ParentID,&a.ParentText)
			attributes = append(attributes, a)
		}
		return attributes, err
		}

	func (d MySQLDriver) GetCitationDetails(text1 string, text2 string) (papers []model.Paper, err error){
		db, err := d.OpenDB()
		checkErr(err)
		db, stmt, err := d.Query("select distinct id_paper, citation, bib, referenceCount from paper where citation = \"" + text1 + "\" or citation = \"" + text2 + "\" order by id_paper;")
		defer stmt.Close()
		defer db.Close()
		rows, err := stmt.Query()
		checkErr(err)
		for rows.Next() {
			a := model.Paper{}
			rows.Scan(&a.ID,&a.Citation,&a.Bib,&a.ReferenceCount)
			papers = append(papers, a)
		}
		return papers, err
		}

	func (d MySQLDriver) GetAttributeCoverage() (attributeCoverage []model.AttributeCoverage, err error){
		db, err := d.OpenDB()
		checkErr(err)
		db, stmt, err := d.Query(`select distinct attribute.text as attributeName, paper.citation as paperName, attribute.text as text1, paper.citation as text2, attribute.id_attribute as attributeID, paper.id_paper as paperID, paper.referenceCount as value from attribute inner join mapping on (attribute.id_attribute = mapping.id_attribute) inner join paper on (mapping.id_paper = paper.id_paper);`)
		defer stmt.Close()
		defer db.Close()
		rows, err := stmt.Query()
		checkErr(err)
		for rows.Next() {
			a := model.AttributeCoverage{}
			rows.Scan(&a.AttributeName,&a.PaperName,&a.Text1,&a.Text2,&a.AttributeID,&a.PaperID,&a.Value)
			attributeCoverage = append(attributeCoverage, a)
		}
		return attributeCoverage, err
		}

	func (d MySQLDriver) GetAttributesPerDimension(taxonomyId int64, dimension string) (attributes []model.Attribute, err error){
		db, err := d.OpenDB()
		checkErr(err)
		taxonomyIdStr := strconv.Itoa(int(taxonomyId))
		db, stmt, err := d.Query("select distinct attribute.id_attribute, attribute.text, dimension.text as dimensionText, attribute.major, attribute.x, attribute.y, attribute.x3D, attribute.y3D, attribute.z3D from attribute inner join taxonomy_dimension on (attribute.id_attribute = taxonomy_dimension.id_attribute and taxonomy_dimension.id_taxonomy = " + taxonomyIdStr + ") inner join dimension on (taxonomy_dimension.id_dimension = dimension.id_dimension and dimension.text = \"" + dimension + "\") order by id_attribute;")
		defer stmt.Close()
		defer db.Close()
		rows, err := stmt.Query()
		checkErr(err)
		for rows.Next() {
			a := model.Attribute{}
			rows.Scan(&a.ID,&a.Text,&a.Dimension,&a.Major,&a.X,&a.Y,&a.X3D,&a.Y3D,&a.Z3D)
			attributes = append(attributes, a)
		}
		return attributes, err
		}

	func (d MySQLDriver) GetAllChildrenAttributes(taxonomyId int64, parent string) (attributes []model.Attribute, err error){
		db, err := d.OpenDB()
		checkErr(err)
		taxonomyIdStr := strconv.Itoa(int(taxonomyId))
		db, stmt, err := d.Query("SELECT DISTINCT attributeSrc.id_attribute, attributeSrc.text FROM attribute AS attributeSrc INNER JOIN (SELECT GROUP_CONCAT(lv SEPARATOR ',') AS children FROM (SELECT @pv:=(SELECT GROUP_CONCAT(DISTINCT relation1.id_src_attribute SEPARATOR ',') FROM taxonomy_relation AS relation1 WHERE relation1.id_taxonomy = " + taxonomyIdStr + " AND relation1.id_relation > 2 AND FIND_IN_SET(relation1.id_dest_attribute, @pv)) AS lv FROM taxonomy_relation JOIN (SELECT @pv:=attributeDest.id_attribute from attribute as attributeDest where attributeDest.text = \"" + parent + "\") tmp) a) AS tmpTable on (attributeSrc.text = \"" + parent + "\" OR FIND_IN_SET(attributeSrc.id_attribute, tmpTable.children));")
		defer stmt.Close()	
		defer db.Close()
		rows, err := stmt.Query()
		checkErr(err)
		for rows.Next() {
			a := model.Attribute{}
			rows.Scan(&a.ID,&a.Text)
			attributes = append(attributes, a)
		}
		return attributes, err
		}

	// TODO with input = number of children, levels
	func (d MySQLDriver) GetIntermediateAttributes(taxonomyId int64, minValue int64, maxValue int64) (attributes []model.Attribute, err error){
		db, err := d.OpenDB()
		checkErr(err)
		taxonomyIdStr := strconv.Itoa(int(taxonomyId))
		minValueStr := strconv.Itoa(int(minValue))
		maxValueStr := strconv.Itoa(int(maxValue))
		db, stmt, err := d.Query("select distinct attribute.id_attribute, attribute.text from attribute left join taxonomy_relation on (attribute.id_attribute = taxonomy_relation.id_dest_attribute and taxonomy_relation.id_relation > 2 and taxonomy_relation.id_taxonomy = " + taxonomyIdStr + ") group by attribute.id_attribute having count(distinct taxonomy_relation.id_taxonomy_relation) >= " + minValueStr + " and count(distinct taxonomy_relation.id_taxonomy_relation) <= " + maxValueStr + ";")
		defer stmt.Close()
		defer db.Close()
		rows, err := stmt.Query()
		checkErr(err)
		for rows.Next() {
			a := model.Attribute{}
			rows.Scan(&a.ID,&a.Text)
			attributes = append(attributes, a)
		}
		return attributes, err
		}

	func (d MySQLDriver) GetMajorAttributes(taxonomyId int64) (attributes []model.Attribute, err error){
		db, err := d.OpenDB()
		checkErr(err)
		taxonomyIdStr := strconv.Itoa(int(taxonomyId))
		db, stmt, err := d.Query("select distinct attribute.id_attribute, attribute.text, dimension.text as dimensionText, attribute.xMajor, attribute.yMajor, attribute.xMajor3D, attribute.yMajor3D, attribute.zMajor3D from attribute inner join taxonomy_dimension on (attribute.major = 1 and attribute.id_attribute = taxonomy_dimension.id_attribute and taxonomy_dimension.id_taxonomy = " + taxonomyIdStr + ") inner join dimension on (taxonomy_dimension.id_dimension = dimension.id_dimension) order by attribute.id_attribute;")
		defer stmt.Close()
		defer db.Close()
		rows, err := stmt.Query()
		checkErr(err)
		for rows.Next() {
			a := model.Attribute{}
			rows.Scan(&a.ID,&a.Text,&a.Dimension,&a.X,&a.Y,&a.XMajor3D,&a.YMajor3D,&a.ZMajor3D)
			a.Major = 1;
			attributes = append(attributes, a)
		}
		return attributes, err
		}

	func (d MySQLDriver) GetAttributeRelationsPerDimension(taxonomyId int64, dimension string) (attributeRelations []model.AttributeRelation, err error){
		db, err := d.OpenDB()
		checkErr(err)
		taxonomyIdStr := strconv.Itoa(int(taxonomyId))
		dimensionStr := dimension
		db, stmt, err := d.Query("select distinct attribute1.text as attributeSrc, attribute2.text as attributeDest, relation.text as relation, annotation.annotation, taxonomy_relation.edgeBendPoints from attribute as attribute1 inner join taxonomy_relation on (attribute1.id_attribute = taxonomy_relation.id_src_attribute and taxonomy_relation.id_taxonomy = " + taxonomyIdStr + ") inner join attribute as attribute2 on (taxonomy_relation.id_dest_attribute = attribute2.id_attribute) inner join relation on (taxonomy_relation.id_relation = relation.id_relation) inner join taxonomy_dimension as dimension1 on (attribute1.id_attribute = dimension1.id_attribute and dimension1.id_taxonomy = " + taxonomyIdStr + ") inner join taxonomy_dimension as dimension2 on (attribute2.id_attribute = dimension2.id_attribute and dimension2.id_taxonomy = " + taxonomyIdStr + ") inner join dimension as dim1 on (dimension1.id_dimension = dim1.id_dimension and dim1.text = \"" + dimensionStr + "\") inner join dimension as dim2 on (dimension2.id_dimension = dim2.id_dimension and dim2.text = \"" + dimensionStr + "\") left outer join taxonomy_relation_annotation as annotation on (taxonomy_relation.id_taxonomy_relation = annotation.id_taxonomy_relation);")
		defer stmt.Close()
		defer db.Close()
		rows, err := stmt.Query()
		checkErr(err)
		for rows.Next() {
			a := model.AttributeRelation{}
			rows.Scan(&a.AttributeSrc,&a.AttributeDest,&a.Relation,&a.Annotation,&a.EdgeBendPoints)
			attributeRelations = append(attributeRelations, a)
		}
		return attributeRelations, err
		}

	func (d MySQLDriver) GetInterdimensionalRelations(taxonomyId int64) (attributeRelations []model.AttributeRelation, err error){
		db, err := d.OpenDB()
		checkErr(err)
		taxonomyIdStr := strconv.Itoa(int(taxonomyId))
		db, stmt, err := d.Query("select distinct attribute1.text as attributeSrc, attribute2.text as attributeDest, relation.text as relation, annotation.annotation, taxonomy_relation.edgeBendPoints from attribute as attribute1 inner join taxonomy_relation on (taxonomy_relation.id_dimension = 4 and attribute1.id_attribute = taxonomy_relation.id_src_attribute and taxonomy_relation.id_taxonomy = " + taxonomyIdStr + ") inner join attribute as attribute2 on (taxonomy_relation.id_dest_attribute = attribute2.id_attribute) inner join relation on (taxonomy_relation.id_relation = relation.id_relation) left outer join taxonomy_relation_annotation as annotation on (taxonomy_relation.id_taxonomy_relation = annotation.id_taxonomy_relation);")
		defer stmt.Close()
		defer db.Close()
		rows, err := stmt.Query()
		checkErr(err)
		for rows.Next() {
			a := model.AttributeRelation{}
			rows.Scan(&a.AttributeSrc,&a.AttributeDest,&a.Relation,&a.Annotation,&a.EdgeBendPoints)
			attributeRelations = append(attributeRelations, a)
		}
		return attributeRelations, err
		}

	func (d MySQLDriver) SavePositions(positions []model.Position) (err error){
    	for _, elem := range positions {
	  		db, err := d.OpenDB()
			checkErr(err)
			db, stmt, err := d.Query("update " + elem.Table + " set x = \"" + elem.X + "\", y = \"" + elem.Y + "\" where text = \"" + elem.ID + "\";");
			defer stmt.Close()
			defer db.Close()
			rows, err := stmt.Query()
			defer rows.Close()
			checkErr(err)
		}
		return err
		}

	func (d MySQLDriver) SaveMajorPositions(positions []model.Position) (err error){
    	for _, elem := range positions {
	  		db, err := d.OpenDB()
			checkErr(err)
			db, stmt, err := d.Query("update " + elem.Table + " set xMajor = \"" + elem.X + "\", yMajor = \"" + elem.Y + "\" where text = \"" + elem.ID + "\";");
			defer stmt.Close()
			defer db.Close()
			rows, err := stmt.Query()
			defer rows.Close()
			checkErr(err)
		}
		return err
		}

	func (d MySQLDriver) Save3DPositions(positions []model.Position) (err error){
    	for _, elem := range positions {
	  		db, err := d.OpenDB()
			checkErr(err)
			db, stmt, err := d.Query("update " + elem.Table + " set x3D = \"" + elem.X + "\", y3D = \"" + elem.Y + "\", z3D = \"" + elem.Z + "\" where text = \"" + elem.ID + "\";");
			defer stmt.Close()
			defer db.Close()
			rows, err := stmt.Query()
			defer rows.Close()
			checkErr(err)
		}
		return err
		}

	func (d MySQLDriver) SaveMajor3DPositions(positions []model.Position) (err error){
    	for _, elem := range positions {
	  		db, err := d.OpenDB()
			checkErr(err)
			db, stmt, err := d.Query("update " + elem.Table + " set xMajor3D = \"" + elem.X + "\", yMajor3D = \"" + elem.Y + "\", zMajor3D = \"" + elem.Z + "\" where text = \"" + elem.ID + "\";");
			defer stmt.Close()
			defer db.Close()
			rows, err := stmt.Query()
			defer rows.Close()
			checkErr(err)
		}
		return err
		}

	func (d MySQLDriver) SaveEdgeBendPoints(attributeSrc string, attributeDest string, edgeBendPoints string) (result model.Result, err error){
	  	db, err := d.OpenDB()
		checkErr(err)
		fmt.Sprintf("edgepoints: " + edgeBendPoints)
		db, stmt, err := d.Query("update taxonomy_relation set edgeBendPoints = \"" + edgeBendPoints + "\" where id_src_attribute = (select distinct id_attribute from attribute where text = \"" + attributeSrc + "\") and id_dest_attribute = (select distinct id_attribute from attribute where text = \"" + attributeDest + "\");")
		defer stmt.Close()
		defer db.Close()
		stmt.Query()
		result.Success = true
		return result, err
		}

	func (d MySQLDriver) RemoveTaxonomyRelationsForAttribute(attribute model.Attribute) (result model.Result, err error){
		db, err := d.OpenDB()
		checkErr(err)
		attributeIDStr := strconv.Itoa(int(attribute.ID))
		db, stmt, err := d.Query("DELETE FROM taxonomy_relation_annotation WHERE id_taxonomy_relation IN (SELECT id_taxonomy_relation FROM taxonomy_relation WHERE id_src_attribute = " + attributeIDStr + " OR id_dest_attribute = " + attributeIDStr + ");")
		db2, stmt2, err2 := d.Query("DELETE FROM taxonomy_relation WHERE id_src_attribute = " + attributeIDStr + " OR id_dest_attribute = " + attributeIDStr + ";")
		checkErr(err)
		checkErr(err2)
		defer stmt.Close()
		defer stmt2.Close()
		defer db.Close()
		defer db2.Close()
		stmt.Query()
		stmt2.Query()
		result.Success = true
		return result, err
		}

	func (d MySQLDriver) AddAttribute(attribute model.Attribute) (result model.Result, err error){
		db, err := d.OpenDB()
		checkErr(err)
		majorStr := strconv.Itoa(int(attribute.Major))
		db, stmt, err := d.Query("INSERT INTO attribute (text, x, y, xMajor, yMajor, major) VALUES (\"" + attribute.Text + "\", \"" + attribute.X + "\", \"" + attribute.Y + "\", \"" + attribute.XMajor + "\", \"" + attribute.YMajor + "\", " + majorStr + ");")
		db2, stmt2, err2 := d.Query("INSERT INTO taxonomy_dimension (id_taxonomy, id_attribute, id_dimension) VALUES (1, (SELECT DISTINCT id_attribute FROM attribute WHERE text = \"" + attribute.Text + "\"), (SELECT DISTINCT id_dimension FROM dimension WHERE text = \"" + attribute.Dimension + "\"));")
		db3, stmt3, err3 := d.Query("DELETE FROM allChildrenPerAttribute;")
		db4, stmt4, err4 := d.Query("CALL insertAllChildrenPerAttribute();")
		checkErr(err)
		checkErr(err2)
		checkErr(err3)
		checkErr(err4)
		defer stmt.Close()
		defer db.Close()
		defer stmt2.Close()
		defer db2.Close()
		defer stmt3.Close()
		defer db3.Close()
		defer stmt4.Close()
		defer db4.Close()
		stmt.Query()
		stmt2.Query()
		stmt3.Query()
		stmt4.Query()
		result.Success = true
		return result, err
		}

	func (d MySQLDriver) RenameAttribute(previousName string, newName string) (result model.Result, err error){
		db, err := d.OpenDB()
		checkErr(err)
		db, stmt, err := d.Query("SELECT count(id_attribute) FROM attribute WHERE text = \"" + newName + "\";")
		defer stmt.Close()
		defer db.Close()
		rows, err := stmt.Query()
		checkErr(err)
		var rowCount int
		for rows.Next() {
			rows.Scan(&rowCount)
		}
		if rowCount > 0 {
			result.Success = false
			return result, err
		}
		db2, stmt2, err2 := d.Query("UPDATE attribute SET text = \"" + newName + "\" WHERE text = \"" + previousName + "\";")
		db3, stmt3, err3 := d.Query("DELETE FROM allChildrenPerAttribute;")
		db4, stmt4, err4 := d.Query("CALL insertAllChildrenPerAttribute();")
		checkErr(err2)
		checkErr(err3)
		checkErr(err4)
		defer stmt2.Close()
		defer db2.Close()
		defer stmt3.Close()
		defer db3.Close()
		defer stmt4.Close()
		defer db4.Close()
		stmt2.Query()
		stmt3.Query()
		stmt4.Query()
		result.Success = true
		return result, err
		}


	func (d MySQLDriver) MergeAttributes(attribute1 model.Attribute, attribute2 model.Attribute) (result model.Result, err error){
		db, err := d.OpenDB()
		checkErr(err)
		db, stmt, err := d.Query("INSERT INTO attribute (text, x, y, xMajor, yMajor, major) SELECT \"" + attribute1.Text + ":" + attribute2.Text + "\" as newName, attribute.x, attribute.y, attribute.xMajor, attribute.yMajor, attribute.major FROM attribute WHERE attribute.text = \"" + attribute1.Text + "\";")
		db2, stmt2, err2 := d.Query("INSERT INTO taxonomy_dimension (id_taxonomy, id_attribute, id_dimension) VALUES (1, (SELECT DISTINCT id_attribute FROM attribute WHERE text = \"" + attribute1.Text + ":" + attribute2.Text + "\"), (SELECT DISTINCT id_dimension FROM dimension WHERE text = \"" + attribute1.Dimension + "\"));")
		db3, stmt3, err3 := d.Query("INSERT INTO taxonomy_relation (id_taxonomy, id_src_attribute, id_dest_attribute, id_relation, id_dimension) SELECT id_taxonomy, (SELECT DISTINCT id_attribute FROM attribute WHERE text = \"" + attribute1.Text + ":" + attribute2.Text + "\"), id_dest_attribute, id_relation, id_dimension FROM taxonomy_relation WHERE id_src_attribute = (SELECT id_attribute FROM attribute WHERE text = \"" + attribute1.Text + "\");")
		db4, stmt4, err4 := d.Query("INSERT INTO taxonomy_relation (id_taxonomy, id_src_attribute, id_dest_attribute, id_relation, id_dimension) SELECT id_taxonomy, (SELECT DISTINCT id_attribute FROM attribute WHERE text = \"" + attribute1.Text + ":" + attribute2.Text + "\"), id_dest_attribute, id_relation, id_dimension FROM taxonomy_relation WHERE id_src_attribute = (SELECT id_attribute FROM attribute WHERE text = \"" + attribute2.Text + "\");")
		db5, stmt5, err5 := d.Query("INSERT INTO taxonomy_relation (id_taxonomy, id_src_attribute, id_dest_attribute, id_relation, id_dimension) SELECT id_taxonomy, id_src_attribute, (SELECT DISTINCT id_attribute FROM attribute WHERE text = \"" + attribute1.Text + ":" + attribute2.Text + "\"), id_relation, id_dimension FROM taxonomy_relation WHERE id_dest_attribute = (SELECT id_attribute FROM attribute WHERE text = \"" + attribute1.Text + "\");")
		db6, stmt6, err6 := d.Query("INSERT INTO taxonomy_relation (id_taxonomy, id_src_attribute, id_dest_attribute, id_relation, id_dimension) SELECT id_taxonomy, id_src_attribute, (SELECT DISTINCT id_attribute FROM attribute WHERE text = \"" + attribute1.Text + ":" + attribute2.Text + "\"), id_relation, id_dimension FROM taxonomy_relation WHERE id_dest_attribute = (SELECT id_attribute FROM attribute WHERE text = \"" + attribute2.Text + "\");")
		db7, stmt7, err7 := d.Query("INSERT INTO mapping (id_paper, id_attribute) SELECT id_paper, (SELECT DISTINCT id_attribute FROM attribute WHERE text = \"" + attribute1.Text + ":" + attribute2.Text + "\") FROM mapping WHERE id_attribute = (SELECT id_attribute FROM attribute WHERE text = \"" + attribute1.Text + "\");")
		db8, stmt8, err8 := d.Query("INSERT INTO mapping (id_paper, id_attribute) SELECT id_paper, (SELECT DISTINCT id_attribute FROM attribute WHERE text = \"" + attribute1.Text + ":" + attribute2.Text + "\") FROM mapping WHERE id_attribute = (SELECT id_attribute FROM attribute WHERE text = \"" + attribute2.Text + "\");")
		checkErr(err)
		checkErr(err2)
		checkErr(err3)
		checkErr(err4)
		checkErr(err5)
		checkErr(err6)
		checkErr(err7)
		checkErr(err8)
		db9, stmt9, err9 := d.Query("DELETE FROM attribute WHERE text = \"" + attribute1.Text + "\";")
		db10, stmt10, err10 := d.Query("DELETE FROM attribute WHERE text = \"" + attribute2.Text + "\";")
		checkErr(err9)
		checkErr(err10)
		defer stmt.Close()
		defer db.Close()
		defer stmt2.Close()
		defer db2.Close()
		defer stmt3.Close()
		defer db3.Close()
		defer stmt4.Close()
		defer db4.Close()
		defer stmt5.Close()
		defer db5.Close()
		defer stmt6.Close()
		defer db6.Close()
		defer stmt7.Close()
		defer db7.Close()
		defer stmt8.Close()
		defer db8.Close()
		defer stmt9.Close()
		defer db9.Close()
		defer stmt10.Close()
		defer db10.Close()
		stmt.Query()
		stmt2.Query()
		stmt3.Query()
		stmt4.Query()
		stmt5.Query()
		stmt6.Query()
		stmt7.Query()
		stmt8.Query()
		stmt9.Query()
		stmt10.Query()
		result.Success = true
		return result, err
		}

/*

	func (d MySQLDriver) ForkAttributes(attribute model.Attribute) (result model.Result, err error){
		db, err := d.OpenDB()
		checkErr(err)
		majorStr := strconv.Itoa(int(attribute.Major))
		db, stmt, err := d.Query("INSERT INTO attribute (text, x, y, xMajor, yMajor, major) VALUES (\"" + attribute.Text + "\", \"" + attribute.X + "\", \"" + attribute.Y + "\", \"" + attribute.XMajor + "\", \"" + attribute.YMajor + "\", " + majorStr + ");")
		db2, stmt2, err2 := d.Query("INSERT INTO taxonomy_dimension (id_taxonomy, id_attribute, id_dimension) VALUES (1, (SELECT DISTINCT id_attribute FROM attribute WHERE text = \"" + attribute.Text + "\"), (SELECT DISTINCT id_dimension FROM dimension WHERE text = \"" + attribute.Dimension + "\"));")
		checkErr(err)
		checkErr(err2)
		defer stmt.Close()
		defer db.Close()
		defer stmt2.Close()
		defer db2.Close()
		stmt.Query()
		stmt2.Query()
		result.Success = true
		return result, err
		}

*/

	func (d MySQLDriver) AddTaxonomyRelation(relation model.AttributeRelation) (result model.Result, err error){
		db, err := d.OpenDB()
		checkErr(err)
		taxonomyIDStr := strconv.Itoa(int(relation.TaxonomyID))
		db, stmt, err := d.Query("INSERT INTO taxonomy_relation (id_taxonomy, id_src_attribute, id_dest_attribute, id_relation, id_dimension) VALUES (" + taxonomyIDStr + ", (SELECT id_attribute FROM attribute WHERE text = \"" + relation.AttributeSrc + "\"), (SELECT id_attribute FROM attribute WHERE text = \"" + relation.AttributeDest + "\"), (SELECT id_relation FROM relation WHERE text = \"" + relation.Relation + "\"), (SELECT id_dimension FROM dimension WHERE text = \"" + relation.Dimension + "\"));")
		db2, stmt2, err2 := d.Query("DELETE FROM allChildrenPerAttribute;")
		db3, stmt3, err3 := d.Query("CALL insertAllChildrenPerAttribute();")
		checkErr(err)
		checkErr(err2)
		checkErr(err3)
		defer stmt.Close()
		defer db.Close()
		defer stmt2.Close()
		defer db2.Close()
		defer stmt3.Close()
		defer db3.Close()
		stmt.Query()
		stmt2.Query()
		stmt3.Query()
		result.Success = true
		return result, err
		}

	func (d MySQLDriver) RemoveTaxonomyRelation(relation model.AttributeRelation) (result model.Result, err error){
		db, err := d.OpenDB()
		checkErr(err)
		taxonomyIDStr := strconv.Itoa(int(relation.TaxonomyID))
		db, stmt, err := d.Query("DELETE FROM taxonomy_relation WHERE id_taxonomy = " + taxonomyIDStr + " AND id_src_attribute = (SELECT id_attribute FROM attribute WHERE text = \"" + relation.AttributeSrc + "\") AND id_dest_attribute = (SELECT id_attribute FROM attribute WHERE text = \"" + relation.AttributeDest + "\") AND id_relation = (SELECT id_relation FROM relation WHERE text = \"" + relation.Relation + "\") AND id_dimension = (SELECT id_dimension FROM dimension WHERE text = \"" + relation.Dimension + "\");")
		defer stmt.Close()
		defer db.Close()
		stmt.Query()
		result.Success = true
		return result, err
		}

	func (d MySQLDriver) UpdateTaxonomyRelationType(relation model.AttributeRelation) (result model.Result, err error){
		db, err := d.OpenDB()
		checkErr(err)
		taxonomyIDStr := strconv.Itoa(int(relation.TaxonomyID))
		db, stmt, err := d.Query("UPDATE taxonomy_relation SET id_relation = (SELECT id_relation FROM relation WHERE text = \"" + relation.Relation + "\") WHERE id_taxonomy = " + taxonomyIDStr + " AND id_src_attribute = (SELECT id_attribute FROM attribute WHERE text = \"" + relation.AttributeSrc + "\") AND id_dest_attribute = (SELECT id_attribute FROM attribute WHERE text = \"" + relation.AttributeDest + "\") AND id_dimension = (SELECT id_dimension FROM dimension WHERE text = \"" + relation.Dimension + "\");")
		defer stmt.Close()
		defer db.Close()
		stmt.Query()
		result.Success = true
		return result, err
		}

	func (d MySQLDriver) UpdateTaxonomyRelationAnnotation(relation model.AttributeRelation) (result model.Result, err error){
		db, err := d.OpenDB()
		checkErr(err)
		taxonomyIDStr := strconv.Itoa(int(relation.TaxonomyID))
		db, stmt, err := d.Query("REPLACE INTO taxonomy_relation_annotation (id_taxonomy, id_taxonomy_relation, annotation) VALUES (" + taxonomyIDStr + ", (SELECT DISTINCT id_taxonomy_relation FROM taxonomy_relation WHERE id_taxonomy = " + taxonomyIDStr + " AND id_src_attribute = (SELECT id_attribute FROM attribute WHERE text = \"" + relation.AttributeSrc + "\") AND id_dest_attribute = (SELECT id_attribute FROM attribute WHERE text = \"" + relation.AttributeDest + "\") AND id_dimension = (SELECT id_dimension FROM dimension WHERE text = \"" + relation.Dimension + "\")), \"" + relation.Annotation + "\");")
		defer stmt.Close()
		defer db.Close()
		stmt.Query()
		result.Success = true
		return result, err
		}

	func (d MySQLDriver) RemoveAttribute(attribute model.Attribute) (result model.Result, err error){
		db, err := d.OpenDB()
		checkErr(err)
		db, stmt, err := d.Query("SELECT id_attribute FROM attribute WHERE text = \"" + attribute.Text + "\";")
		checkErr(err)
		defer stmt.Close()
		defer db.Close()
		rows, err := stmt.Query()
		checkErr(err)
		for rows.Next() {
			rows.Scan(&attribute.ID)
		}
		attributeIDStr := strconv.Itoa(int(attribute.ID))
		db2, stmt2, err2 := d.Query("DELETE FROM attribute WHERE id_attribute = " + attributeIDStr + ";")
		db3, stmt3, err3 := d.Query("DELETE FROM taxonomy_dimension WHERE id_attribute = " + attributeIDStr + ";")
		db4, stmt4, err4 := d.Query("DELETE FROM mapping WHERE id_attribute = " + attributeIDStr + ";")
		d.RemoveTaxonomyRelationsForAttribute(attribute)
		checkErr(err2)
		checkErr(err3)
		checkErr(err4)
		defer stmt2.Close()
		defer stmt3.Close()
		defer stmt4.Close()
		defer db2.Close()
		defer db3.Close()
		defer db4.Close()
		stmt2.Query()
		stmt3.Query()
		stmt4.Query()
		result.Success = true
		return result, err
		}