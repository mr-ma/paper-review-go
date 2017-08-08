package data

import (
	"testing"

	"github.com/mr-ma/paper-review-go/model"
	"github.com/stretchr/testify/assert"
)

var db = InitMySQLDriver("root","P$m7d2")
func TestInitMySQLDriver(t *testing.T) {
	assert.NotNil(t, db)
}

func TestOpenDB(t *testing.T) {
	_, err := db.OpenDB()
	assert.Nil(t, err)
}
func TestQuery(t *testing.T) {
	//TODO: fix this
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
	re := model.Research{Questions: "test1", ReviewTemplate: "wrtie whatever", Title: "ola"}
	articles := []model.Article{model.Article{Title: "title", Authors: "author", File: "file", Source: "source"}}
	re.Articles = articles
	a, _, err := db.InsertResearch(re)
	assert.Nil(t, err)
	assert.True(t, a > 0)
}
func TestInsertVote(t *testing.T) {
	vote := model.Vote{State: model.YES, Voter: model.Mitarbeiter{ID: 1, Name: "Mohsen"},
		AssociatedArticleID: 1,
		Tags:                []model.Tag{model.Tag{ID: 1, Text: "test1", ResearchID: 6}, model.Tag{ID: 2, Text: "test2"}}}
	a, _, err := db.InsertVote(vote)
	assert.Nil(t, err)
	assert.True(t, a > 0)
}
func TestInsertMitarbeiter(t *testing.T) {
	m := model.Mitarbeiter{Name: "test mitarbeiter", PassHash: ""}
	a, _, err := db.InsertMitarbeiter(m)
	assert.Nil(t, err)
	assert.True(t, a > 0)
}
func TestSelectResearchWithArticles(t *testing.T) {
	re, err := db.SelectResearchWithArticles(4)
	assert.NotZero(t, re.ID)
	assert.Nil(t, err)
	assert.True(t, len(re.Articles) > 0)
	assert.NotEmpty(t, re.Questions)
}
func TestSelectAllResearchWithArticles(t *testing.T) {
	re, err := db.SelectAllResearchWithArticles()
	assert.NotZero(t, len(re))
	assert.Nil(t, err)
}
func TestSelectVote(t *testing.T) {
	re, err := db.SelectVote(44)
	assert.NotZero(t, re.ID)
	assert.Nil(t, err)
	assert.True(t, len(re.Tags) > 0)
	assert.NotZero(t, re.Voter.ID)
}
func TestSelectAllVotes(t *testing.T) {
	re, err := db.SelectAllVotes()
	assert.NotZero(t, len(re))
	assert.Nil(t, err)
}
func TestSelectResearchVotes(t *testing.T) {
	re, err := db.SelectResearchVotes(4)
	assert.NotZero(t, len(re))
	for _, v := range re {
		assert.NotZero(t, v.AssociatedArticleID)
	}
	assert.Nil(t, err)
}
func TestReviewPapers(t *testing.T) {
	a, _, err := db.ReviewPapers(6, 1)
	assert.NotZero(t, len(a))
	assert.True(t, len(a) == 2)
	assert.Nil(t, err)
}
func TestReviewNumPapers(t *testing.T) {
	a, _, err := db.ReviewNumPapers(6, 1, 1)
	assert.NotZero(t, len(a))
	//assert.NotEmpty(t, re.Questions)
	assert.True(t, len(a) == 1)
	assert.Nil(t, err)
}
func TestSelectMitarbeiter(t *testing.T) {
	m, err := db.SelectMitarbeiter(1)
	assert.Nil(t, err)
	assert.NotZero(t, m.ID)
}
func TestSelectAllMitarbeiters(t *testing.T) {
	m, err := db.SelectAllMitarbeiters()
	assert.Nil(t, err)
	assert.NotZero(t, len(m))
}
func TestSelectAllTags(t *testing.T) {
	tags, err := db.SelectAllTags(4)
	assert.Nil(t, err)
	assert.NotZero(t, len(tags))
}
func TestGetResearchStats(t *testing.T) {
	stats, err := db.GetResearchStats(4)
	assert.Nil(t, err)
	assert.NotZero(t, len(stats))
}
func TestGetResearchStatsPerMitarbeiter(t *testing.T) {
	stat, err := db.GetResearchStatsPerMitarbeiter(4, 1)
	assert.Nil(t, err)
	assert.NotZero(t, stat.MitarbeiterID)
}
func TestGetApprovedPapers(t *testing.T) {
	atricles, err := db.GetApprovedPapers(4, 2)
	assert.Nil(t, err)
	assert.NotZero(t, len(atricles))
}
