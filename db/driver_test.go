package db

import (
	"testing"

	"github.com/mr-ma/review"
	"github.com/stretchr/testify/assert"
)

func TestInitMySQLDriver(t *testing.T) {
	db := InitMySQLDriver()
	assert.NotNil(t, db)
}

func TestOpenDB(t *testing.T) {
	db := InitMySQLDriver()
	_, err := db.OpenDB()
	assert.Nil(t, err)
}
func TestQuery(t *testing.T) {
	//TODO: fix this
	db := InitMySQLDriver()
	d, stmt, err := db.Query("select Id from mitarbeiters where Nme like ?")
	defer stmt.Close()
	defer d.Close()
	id := 0
	err = stmt.QueryRow("Mohsen").Scan(&id)
	assert.Nil(t, err)
	assert.NotZero(t, id)
}
func TestInsert(t *testing.T) {

}
func TestInsertArticle(t *testing.T) {

}
func TestInsertResearch(t *testing.T) {
	re := review.Research{Questions: "test1", ReviewTemplate: "wrtie whatever"}
	articles := []review.Article{review.Article{Title: "title", Authors: "author"}}
	re.Articles = articles
	d := InitMySQLDriver()
	a, _, err := d.InsertResearch(re)
	assert.Nil(t, err)
	assert.True(t, a > 0)
}
func TestInsertVote(t *testing.T) {
	vote := review.Vote{State: review.YES, Voter: review.Mitarbeiter{ID: 1, Name: "Mohsen"},
		AssociatedArticleID: 1,
		Tags:                []review.Tag{review.Tag{ID: 1, Text: "test1"}, review.Tag{ID: 2, Text: "test2"}}}
	d := InitMySQLDriver()
	a, _, err := d.InsertVote(vote)
	assert.Nil(t, err)
	assert.True(t, a > 0)
}
func TestInsertMitarbeiter(t *testing.T) {
	m := review.Mitarbeiter{Name: "test mitarbeiter", PassHash: ""}
	d := InitMySQLDriver()
	a, _, err := d.InsertMitarbeiter(m)
	assert.Nil(t, err)
	assert.True(t, a > 0)
}
func TestSelectResearchWithArticles(t *testing.T) {
	d := InitMySQLDriver()
	re, err := d.SelectResearchWithArticles(4)
	assert.NotZero(t, re.ID)
	assert.Nil(t, err)
	assert.True(t, len(re.Articles) > 0)
	assert.NotEmpty(t, re.Questions)
}
func TestSelectAllResearchWithArticles(t *testing.T) {
	d := InitMySQLDriver()
	re, err := d.SelectAllResearchWithArticles()
	assert.NotZero(t, len(re))
	assert.Nil(t, err)
}
func TestSelectVote(t *testing.T) {
	d := InitMySQLDriver()
	re, err := d.SelectVote(5)
	assert.NotZero(t, re.ID)
	assert.Nil(t, err)
	assert.True(t, len(re.Tags) > 0)
	assert.NotZero(t, re.Voter.ID)
}
func TestSelectAllVotes(t *testing.T) {
	d := InitMySQLDriver()
	re, err := d.SelectAllVotes()
	assert.NotZero(t, len(re))
	assert.Nil(t, err)
}
func TestSelectResearchVotes(t *testing.T) {
	d := InitMySQLDriver()
	re, err := d.SelectResearchVotes(4)
	assert.NotZero(t, len(re))
	for _, v := range re {
		assert.NotZero(t, v.AssociatedArticleID)
	}
	assert.Nil(t, err)
}
func TestReviewPapers(t *testing.T) {
	d := InitMySQLDriver()
	a, _, err := d.ReviewPapers(6, 1)
	assert.NotZero(t, len(a))
	assert.True(t, len(a) == 2)
	assert.Nil(t, err)
}
func TestReviewNumPapers(t *testing.T) {
	d := InitMySQLDriver()
	a, _, err := d.ReviewNumPapers(6, 1, 1)
	assert.NotZero(t, len(a))
	//assert.NotEmpty(t, re.Questions)
	assert.True(t, len(a) == 1)
	assert.Nil(t, err)
}
func TestSelectMitarbeiter(t *testing.T) {
	d := InitMySQLDriver()
	m, err := d.SelectMitarbeiter(1)
	assert.Nil(t, err)
	assert.NotZero(t, m.ID)
}
func TestSelectAllMitarbeiters(t *testing.T) {
	d := InitMySQLDriver()
	m, err := d.SelectAllMitarbeiters()
	assert.Nil(t, err)
	assert.NotZero(t, len(m))
}
func TestSelectAllTags(t *testing.T) {
	d := InitMySQLDriver()
	tags, err := d.SelectAllTags()
	assert.Nil(t, err)
	assert.NotZero(t, len(tags))
}
