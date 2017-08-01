package data
import (
	//overriding MySqlDriver
	_ "github.com/go-sql-driver/mysql"
		"github.com/mr-ma/paper-review-go/model"
)

type ClassificationDriver interface {
  DriverCore
	ExportCorrelations([]model.Attribute, int64) ([]model.Paper, error)
	GetAllAttributes() ([]model.Attribute, error)
	GetAllCitations() ([]model.Paper, error)
	ExportCorrelationsCSV(filterAttributes []model.Attribute)
}


//InitMySQLDriver initialize a new my sql driver instance
func InitClassificationDriver() ClassificationDriver {
	return MySQLDriver{username: "root", pass: "P$m7d2", database: "classification"}
}

//ExportCorrelations export correlations with the given attributes
func (d MySQLDriver) ExportCorrelations(filterAttributes []model.Attribute,
	taxonomyId int64) (papers []model.Paper, err error) {
	db, err := d.OpenDB()
	checkErr(err)
	//prepare list of attribute ids for the where clause
	attributeStr := "%"
	for _, attribute := range filterAttributes {
		attributeStr+=attribute.Text+",%"
	}
	db, stmt, err := d.Query(`select id_paper, citation, bib,atts
		from paper_merged_attributes
where id_taxonomy=? and atts like ?;`)
	defer stmt.Close()
	defer db.Close()
	rows, err := stmt.Query(taxonomyId,attributeStr)
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
	db, stmt, err := d.Query(`select text
		from attribute`)
	defer stmt.Close()
	defer db.Close()
	rows, err := stmt.Query()
	checkErr(err)
	for rows.Next() {
		a := model.Attribute{}
		rows.Scan(&a.Text)
		attributes = append(attributes, a)
	}
	return attributes, err

	}

	func (d MySQLDriver) GetAllCitations() (papers []model.Paper,
		err error){
		db, err := d.OpenDB()
		checkErr(err)
		db, stmt, err := d.Query(`select citation,bib
			from paper`)
		defer stmt.Close()
		defer db.Close()
		rows, err := stmt.Query()
		checkErr(err)
		for rows.Next() {
			a := model.Paper{}
			rows.Scan(&a.Citation,&a.Bib)
			papers = append(papers, a)
		}
		return papers, err

		}
