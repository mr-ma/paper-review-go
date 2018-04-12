package data

import (
	"errors"
	"fmt"
	"strconv"
	//overriding MySqlDriver
	"../model"
	_ "github.com/go-sql-driver/mysql"
)

type PaperReviewDriver interface {
	DriverCore
	InsertArticle(article model.Article, researchID int64) (int64, error)
	InsertResearch(research model.Research) (int64, int64, error)
	InsertTag(tag model.Tag) (affected int64, id int64, err error)
	InsertVoteTags(tags []model.Tag, voteID int64) (affected int64, err error)
	InsertVote(vote model.Vote) (affected int64, id int64, err error)
	// InsertMitarbeiter(mitarbeiter model.Mitarbeiter) (affected int64, id int64, err error)
	// SelectMitarbeiter(string) (model.User, error)
	// SelectAllMitarbeiters() ([]model.User, error)
	SelectResearchWithArticles(id int64) (r model.Research, err error)
	SelectAllResearch() (r []model.Research, err error)
	SelectAllResearchWithArticles() (r []model.Research, err error)
	SelectVote(id int64) (r model.Vote, err error)
	SelectAllVotes() (r []model.Vote, err error)
	SelectResearchVotes(researchID int64) (r []model.Vote, err error)
	ReviewPapers(researchID int64, mitarbeiterID int64) (a []model.Article, r model.Research, err error)
	ReviewNumPapers(researchID int64, mitarbeiterID int64, limit int) (a []model.Article, r model.Research, err error)
	SelectAllTags(researchID int64) ([]model.Tag, error)
	GetResearchStatsPerMitarbeiter(researchID int64, mitarbeiterID int64) (s model.Stats, err error)
	GetResearchStats(researchID int64) (s []model.Stats, err error)
	GetApprovedPapers(researchID int64, threshold int) ([]model.Article, error)
	GetApprovedPapersWithDetails(researchID int64, threshold int) ([]model.ArticleMapping, error)
	DeleteArticleVotes([]model.Article) (model.Result, error)
}

//InitMySQLDriver initialize a new my sql driver instance
func InitPaperReviewDriver(user string, password string) PaperReviewDriver {
	return MySQLDriver{username: user, pass: password, database: "classification"}
}

//SelectResearchWithArticles a research with it's associated articles
func (d MySQLDriver) SelectResearchWithArticles(id int64) (r model.Research, err error) {
	db, err := d.OpenDB()
	defer db.Close()
	checkErr(err)
	db, stmt, err := d.Query(`Select Research.researchid,research.questions,research.Review_template, research.Title researchTitle,
		a.ArticleId, a.Title, a.year, a.cited_by, a.keywords, a.abstract, a.journal,
		a.authors,a.researchid
		from research left outer join articles a on research.researchid=a.researchid
		where research.researchid=?`)
	defer stmt.Close()
	defer db.Close()
	rows, err := stmt.Query(id)
	checkErr(err)
	for rows.Next() {
		a := model.Article{}
		rows.Scan(&r.ID, &r.Questions, &r.ReviewTemplate, &r.Title, &a.ID, &a.Title, &a.Year, &a.CitedBy, &a.Keywords, &a.Abstract, &a.Journal, &a.Authors, &a.AssociatedResearchId)
		if a.AssociatedResearchId != 0 {
			r.Articles = append(r.Articles, a)
		}
	}
	return r, err
}

//SelectAllResearchWithArticles a research with it's associated articles
func (d MySQLDriver) SelectAllResearchWithArticles() (r []model.Research, err error) {
	db, err := d.OpenDB()
	defer db.Close()
	checkErr(err)
	db, stmt, err := d.Query(`Select Research.researchid,research.questions,
		research.Review_template,research.Title researchTitle, a.ArticleId, a.Title, a.year,
		a.cited_by, a.keywords, a.abstract, a.journal, a.authors, a.researchid
		from research inner join articles a on research.researchid=a.researchid`)
	defer stmt.Close()
	defer db.Close()
	rows, err := stmt.Query()
	checkErr(err)
	m := make(map[int]model.Research)
	for rows.Next() {
		id := 0
		questions := ""
		template := ""
		title := ""
		a := model.Article{}
		rows.Scan(&id, &questions, &template, &title, &a.ID, &a.Title, &a.Year, &a.CitedBy, &a.Keywords, &a.Abstract, &a.Journal, &a.Authors, &a.AssociatedResearchId)
		research := m[id]
		research.ID = id
		research.Questions = questions
		research.ReviewTemplate = template
		research.Title = title
		research.Articles = append(research.Articles, a)
		m[id] = research
	}
	researcharray := []model.Research{}
	for _, value := range m {
		researcharray = append(researcharray, value)
	}
	return researcharray, err
}

//SelectAllResearch a research without articles
func (d MySQLDriver) SelectAllResearch() (r []model.Research, err error) {
	db, err := d.OpenDB()
	defer db.Close()
	checkErr(err)
	db, stmt, err := d.Query(`Select researchid,questions,Review_template,Title
		from research`)
	defer stmt.Close()
	defer db.Close()
	rows, err := stmt.Query()
	checkErr(err)
	m := make(map[int]model.Research)
	for rows.Next() {
		id := 0
		title := ""
		questions := ""
		template := ""
		rows.Scan(&id, &questions, &template, &title)
		research := m[id]
		research.ID = id
		research.Title = title
		research.Questions = questions
		research.ReviewTemplate = template
		m[id] = research
	}
	researcharray := []model.Research{}
	for _, value := range m {
		researcharray = append(researcharray, value)
	}
	return researcharray, err
}

//SelectVote picks an arbitrary vote
func (d MySQLDriver) SelectVote(id int64) (model.Vote, error) {
	v := model.Vote{}
	db, err := d.OpenDB()
	defer db.Close()
	checkErr(err)
	db, stmt, err := d.Query(`select Votes.VoteId, vote_State,Review,a.ArticleID,
		m.id MitarbeiterID,m.name,t.TagId as TagID,t.text as TagText,t.ResearchID
		from Votes inner join articles a
		on Votes.ArticleId = a.ArticleID inner join user m
		on Votes.MitarbeiterId = m.id left outer join Vote_Tags vt
		on Votes.VoteId =vt.VoteId left outer join tags t
		on vt.Tag_Id = t.TagId where votes.VoteId=?`)
	defer stmt.Close()
	defer db.Close()
	rows, err := stmt.Query(id)
	checkErr(err)
	defer rows.Close()

	for rows.Next() {
		a := model.Tag{}
		rows.Scan(&v.ID, &v.State, &v.Review, &v.AssociatedArticleID, &v.Voter.ID, &v.Voter.Name, &a.ID, &a.Text, &a.ResearchID)
		v.Tags = append(v.Tags, a)
	}
	if v.ID <= 0 {
		panic("NO ID")
	}
	return v, err
}

//SelectAllVotes selects all votes
func (d MySQLDriver) SelectAllVotes() (r []model.Vote, err error) {
	db, err := d.OpenDB()
	defer db.Close()
	checkErr(err)
	db, stmt, err := d.Query(`select Votes.VoteId, vote_State,Review,a.ArticleID,
m.id MitarbeiterID,m.name
,t.TagId as TagID,t.text as TagText,t.ResearchID
from Votes inner join articles a
on Votes.ArticleId = a.ArticleID
inner join user m
on Votes.MitarbeiterId = m.id
left outer join Vote_Tags vt on Votes.VoteId =vt.VoteId
left outer join tags t on vt.Tag_Id = t.TagId`)
	defer stmt.Close()
	defer db.Close()
	rows, err := stmt.Query()
	checkErr(err)
	m := make(map[int]model.Vote)
	for rows.Next() {
		id := 0
		State := model.UNSURE
		voteReview := ""
		articleID := 0
		mit := model.User{}
		a := model.Tag{}
		rows.Scan(&id, &State, &voteReview, &articleID, &mit.ID, &mit.Name, &a.ID, &a.Text, &a.ResearchID)
		vote := m[id]
		vote.ID = id
		vote.Voter = mit
		vote.Review = voteReview
		vote.State = State
		if a.ID > 0 {
			vote.Tags = append(vote.Tags, a)
		}
		m[id] = vote
	}
	votearray := []model.Vote{}
	for _, value := range m {
		votearray = append(votearray, value)
	}
	return votearray, err
}

//SelectResearchVotes selects all votes
func (d MySQLDriver) SelectResearchVotes(researchID int64) (r []model.Vote, err error) {
	db, err := d.OpenDB()
	defer db.Close()
	checkErr(err)
	db, stmt, err := d.Query(`select Votes.VoteId, vote_State,Review,a.ArticleID,
m.id MitarbeiterID,m.name
,t.TagId as TagID,t.text as TagText,t.ResearchID
from Votes inner join articles a
on Votes.ArticleId = a.ArticleID
inner join user m
on Votes.MitarbeiterId = m.id
left outer join Vote_Tags vt on Votes.VoteId =vt.VoteId
left outer join Tags t on vt.Tag_Id = t.TagId
where a.ResearchId=?`)
	defer stmt.Close()
	defer db.Close()
	rows, err := stmt.Query(researchID)
	checkErr(err)
	m := make(map[int]model.Vote)
	for rows.Next() {
		id := 0
		State := model.UNSURE
		voteReview := ""
		articleID := 0
		mit := model.User{}
		a := model.Tag{}
		rows.Scan(&id, &State, &voteReview, &articleID, &mit.ID, &mit.Name, &a.ID, &a.Text, &a.ResearchID)
		vote := m[id]
		vote.ID = id
		vote.Voter = mit
		vote.Review = voteReview
		vote.State = State
		vote.AssociatedArticleID = articleID
		if a.ID > 0 {
			vote.Tags = append(vote.Tags, a)
		}
		m[id] = vote
	}
	votearray := []model.Vote{}
	for _, value := range m {
		votearray = append(votearray, value)
	}
	return votearray, err
}
func (d MySQLDriver) SelectAllTags(researchID int64) (tags []model.Tag, err error) {
	db, err := d.OpenDB()
	defer db.Close()
	checkErr(err)
	db, stmt, err := d.Query(`select TagId,Text,ResearchID from tags where ResearchID=?`)
	defer stmt.Close()
	defer db.Close()
	rows, err := stmt.Query(researchID)
	checkErr(err)
	for rows.Next() {
		id := 0
		text := ""
		researchID := 0
		rows.Scan(&id, &text, &researchID)
		tags = append(tags, model.Tag{ID: int64(id), Text: text, ResearchID: researchID})
	}
	return tags, err
}

/*
//SelectMitarbeiter select
func (d MySQLDriver) SelectMitarbeiter(email string) (m model.User, err error) {
	db, err := d.OpenDB()
	defer db.Close()
	checkErr(err)
	db, stmt, err := d.Query(`select id,name from user where email=?`)
	defer stmt.Close()
	defer db.Close()
	rows, err := stmt.Query(email)
	checkErr(err)
	for rows.Next() {
		m = model.User{}
		rows.Scan(&m.ID, &m.Name)
	}
	return m, err
}

//SelectAllMitarbeiters all mitarbeiters
func (d MySQLDriver) SelectAllMitarbeiters() (marr []model.User, err error) {
	db, err := d.OpenDB()
	defer db.Close()
	checkErr(err)
	db, stmt, err := d.Query(`select id,name from user`)
	defer stmt.Close()
	defer db.Close()
	rows, err := stmt.Query()
	checkErr(err)
	for rows.Next() {
		id := 0
		name := ""
		rows.Scan(&id, &name)
		marr = append(marr, model.User{ID: id, Name: name})
	}
	return marr, err
}
*/

func (d MySQLDriver) ReviewPapers(researchID int64, mitarbeiterID int64) (articleArray []model.Article, r model.Research, err error) {
	db, err := d.OpenDB()
	defer db.Close()
	checkErr(err)
	db, stmt, err := d.Query(`select a.ArticleId,a.Title,a.year, a.cited_by,a.Keywords,
a.Abstract,a.Journal,a.Authors,a.ResearchId,
r.Questions,r.Review_Template

from articles a inner join research r on a.researchId =r.researchid and a.Enabled = 1
left outer join (select * from votes where MitarbeiterId =?) v on a.ArticleId = v.ArticleId
where v.MitarbeiterId is null and a.researchId=?`)
	defer stmt.Close()
	defer db.Close()
	rows, err := stmt.Query(mitarbeiterID, researchID)
	checkErr(err)

	for rows.Next() {
		a := model.Article{}
		rows.Scan(&a.ID, &a.Title, &a.Year, &a.CitedBy, &a.Keywords,
			&a.Abstract, &a.Journal, &a.Authors, &a.AssociatedResearchId,
			&r.Questions, &r.ReviewTemplate)
		r.ID = a.ID
		articleArray = append(articleArray, a)
	}
	return articleArray, r, err
}

//ReviewNumPapers review limited papers
func (d MySQLDriver) ReviewNumPapers(researchID int64, mitarbeiterID int64, limit int) (articleArray []model.Article, r model.Research, err error) {
	db, err := d.OpenDB()
	defer db.Close()
	checkErr(err)
	db, stmt, err := d.Query(`select a.ArticleId,a.Title,a.year, a.cited_by,a.Keywords,
	a.Abstract,a.Journal,a.Authors,a.ResearchId,
	r.Questions,r.Review_Template

	from articles a inner join research r on a.researchId =r.researchid and a.Enabled = 1
	left outer join (select * from votes where MitarbeiterId =?) v on a.ArticleId = v.ArticleId
	where v.MitarbeiterId is null and a.researchId=?
	limit ?`)
	defer stmt.Close()
	defer db.Close()
	rows, err := stmt.Query(mitarbeiterID, researchID, limit)
	checkErr(err)

	for rows.Next() {
		a := model.Article{}
		rows.Scan(&a.ID, &a.Title, &a.Year, &a.CitedBy, &a.Keywords, &a.Abstract, &a.Journal, &a.Authors, &a.AssociatedResearchId,
			&r.Questions, &r.ReviewTemplate)
		r.ID = a.ID
		articleArray = append(articleArray, a)
	}
	return articleArray, r, err
}

//InsertArticle insert publication
func (d MySQLDriver) InsertArticle(article model.Article, researchID int64) (int64, error) {
	affect, _, err := d.Insert("Articles", "Title=?,Authors=?,year=?,Cited_by=?,Keywords=?,Abstract=?,Journal=?,File=?,Source=?,ResearchId=?",
		article.Title, article.Authors, article.Year, article.CitedBy, article.Keywords,
		article.Abstract, article.Journal, article.File, article.Source, researchID)
	return affect, err
}

//InsertResearch insert overall research including articles
func (d MySQLDriver) InsertResearch(research model.Research) (int64, int64, error) {
	affect, id, err := d.Insert("Research", "Questions=?,Review_template=?,Title=?",
		research.Questions, research.ReviewTemplate, research.Title)
	for _, element := range research.Articles {
		a, _ := d.InsertArticle(element, id)
		affect += a
	}
	return affect, id, err
}

//InsertTag insert article tags
func (d MySQLDriver) InsertTag(tag model.Tag) (affected int64, id int64, err error) {
	affect, id, err := d.Insert("Tags", "Text=?, ResearchID=?", tag.Text, tag.ResearchID)
	return affect, id, err
}

//InsertVoteTags insert tags corresponding to a vote
func (d MySQLDriver) InsertVoteTags(tags []model.Tag, voteID int64) (affected int64, err error) {
	for _, tag := range tags {
		if tag.ID <= 0 {
			if tag.ResearchID == 0 {
				return 0, errors.New("Tag.ResearchID is missing")
			}
			//tag doesn't exist
			//first query by Text
			// TODO: figure this out
			// stmtOut, err := d.Query("select TagId from Tags where text like ?")
			// checkErr(err)
			// defer stmtOut.Close()
			// row := stmtOut.QueryRow(tag.Text)
			// checkErr(err)
			// err = row.Scan(&tag.ID)
			// checkErr(err)
			//not found insert tag
			_, tag.ID, _ = d.InsertTag(tag)
		}
		if tag.ID <= 0 {
			panic(errors.New("Can't get tag id"))
		}
		affected, _, err = d.Insert("vote_tags", "Tag_Id=?, VoteId=?", tag.ID, voteID)
	} //end for
	return affected, err
}

//InsertVote insert review vote
func (d MySQLDriver) InsertVote(vote model.Vote) (affected int64, id int64, err error) {
	if vote.Voter.ID <= 0 {
		err = errors.New("Voter id is missing")
	}
	checkErr(err)
	if vote.AssociatedArticleID <= 0 {
		err = errors.New("Article id is missing")
	}
	checkErr(err)

	affect, id, err := d.Insert("Votes", "Vote_State=?,MitarbeiterId=?,ArticleId=?,Review=?",
		vote.State, vote.Voter.ID, vote.AssociatedArticleID, vote.Review)
	//Insert Tags
	fmt.Println(vote.Tags)
	affect, err = d.InsertVoteTags(vote.Tags, id)

	checkErr(err)

	return affect, id, err
}

/*
//InsertMitarbeiter insert researcher
func (d MySQLDriver) InsertMitarbeiter(mitarbeiter model.Mitarbeiter) (affected int64, id int64, err error) {
	affect, id, err := d.Insert("Mitarbeiters", "Pass_Hash=?,Nme=?", mitarbeiter.PassHash, mitarbeiter.Name)
	return affect, id, err
}
*/

//GetResearchStatsPerMitarbeiter get statistics on reviewing process
func (d MySQLDriver) GetResearchStatsPerMitarbeiter(researchID int64, mitarbeiterID int64) (s model.Stats, err error) {
	db, err := d.OpenDB()
	defer db.Close()
	checkErr(err)
	db, stmt, err := d.Query(`select LEAST(count(votes.Vote_State),ar.CountArticles) votes,ar.CountArticles
from articles cross join user
left outer join votes on articles.ArticleId =votes.ArticleId and votes.MitarbeiterId = user.id
inner join (select ResearchId,count(*) CountArticles from articles where articles.Enabled = 1
group by ResearchId) ar on articles.ResearchId = ar.ResearchId and articles.Enabled = 1
group by articles.ResearchId, user.id
having articles.ResearchId = ? and user.id = ?`)
	defer stmt.Close()
	defer db.Close()
	rows, err := stmt.Query(researchID, mitarbeiterID)
	checkErr(err)
	allArticles := 0
	for rows.Next() {
		rows.Scan(&s.ReviewedArticles, &allArticles)
	}
	s.RemainingArticles = allArticles - s.ReviewedArticles
	s.MitarbeiterID = int(mitarbeiterID)
	return s, err
}

//GetResearchStats get statistics on reviewing process
func (d MySQLDriver) GetResearchStats(researchID int64) (s []model.Stats, err error) {
	db, err := d.OpenDB()
	defer db.Close()
	checkErr(err)
	db, stmt, err := d.Query(`select user.id, LEAST(count(votes.Vote_State),ar.CountArticles) votes,ar.CountArticles
from articles cross join user
left outer join votes on articles.ArticleId =votes.ArticleId and votes.MitarbeiterId = user.id
inner join (select ResearchId,count(*) CountArticles from articles where articles.Enabled = 1
group by ResearchId) ar on articles.ResearchId = ar.ResearchId and articles.Enabled = 1
group by articles.ResearchId, user.id
having articles.ResearchId = ? `)
	defer stmt.Close()
	defer db.Close()
	rows, err := stmt.Query(researchID)
	checkErr(err)

	for rows.Next() {
		allArticles := 0
		stats := model.Stats{}
		rows.Scan(&stats.MitarbeiterID, &stats.ReviewedArticles, &allArticles)
		stats.RemainingArticles = allArticles - stats.ReviewedArticles
		s = append(s, stats)
	}

	return s, err
}

//GetApprovedPapers get approved papers by threshold
func (d MySQLDriver) GetApprovedPapers(researchID int64, threshold int) (articles []model.Article, err error) {
	db, err := d.OpenDB()
	defer db.Close()
	checkErr(err)
	db, stmt, err := d.Query(`select a.ArticleId,a.Title,a.year,a.cited_by,
a.Keywords,a.Abstract,a.Journal,a.ResearchId,a.Authors
from articles a
inner join votes on a.ArticleId = votes.ArticleId and a.Enabled = 1
group by a.ArticleId,a.ResearchId
having a.ResearchId = ? and count(votes.Vote_State) > ?`)
	defer stmt.Close()
	defer db.Close()
	rows, err := stmt.Query(researchID, threshold)
	checkErr(err)

	for rows.Next() {
		a := model.Article{}
		rows.Scan(&a.ID, &a.Title, &a.Year, &a.CitedBy, &a.Keywords,
			&a.Abstract, &a.Journal, &a.AssociatedResearchId, &a.Authors)
		articles = append(articles, a)
	}

	return articles, err
}

//GetApprovedPapers get approved papers with details by threshold
func (d MySQLDriver) GetApprovedPapersWithDetails(researchID int64, threshold int) (articleMappings []model.ArticleMapping, err error) {
	db, err := d.OpenDB()
	defer db.Close()
	checkErr(err)
	db, stmt, err := d.Query(`select a.ArticleId,a.Title,a.year,a.cited_by,
a.Keywords,a.Abstract,a.Journal,a.ResearchId,a.Authors, GROUP_CONCAT(tmp.TagId SEPARATOR ',') as tagIDs, GROUP_CONCAT(tmp.Tag SEPARATOR ',') as tags, sum(case when votes.Vote_State = 1 then 1 else 0 end) as approvedCount, sum(case when votes.Vote_State = 0 then 1 else 0 end) as rejectedCount
from articles a
inner join votes on a.ArticleId = votes.ArticleId and a.Enabled = 1
left outer join (select vote_tags.VoteId as VoteId, GROUP_CONCAT(tags.TagId SEPARATOR ',') as TagId, GROUP_CONCAT(tags.Text SEPARATOR ',') as Tag from vote_tags inner join tags on (vote_tags.Tag_Id = tags.TagId) group by vote_tags.VoteId) as tmp on votes.VoteId = tmp.VoteId and votes.Vote_State = 1
group by a.ArticleId,a.ResearchId
having a.ResearchId = ? and sum(case when votes.Vote_State = 1 then 1 else 0 end) >= ?`)
	defer stmt.Close()
	defer db.Close()
	rows, err := stmt.Query(researchID, threshold)
	checkErr(err)

	for rows.Next() {
		a := model.ArticleMapping{}
		rows.Scan(&a.ID, &a.Title, &a.Year, &a.CitedBy, &a.Keywords,
			&a.Abstract, &a.Journal, &a.AssociatedResearchId, &a.Authors, &a.TagIds, &a.Tags, &a.ApprovedCount, &a.RejectedCount)
		articleMappings = append(articleMappings, a)
	}

	return articleMappings, err
}

func (d MySQLDriver) DeleteArticleVotes(articles []model.Article) (result model.Result, err error) {
	db, err := d.OpenDB()
	defer db.Close()
	checkErr(err)
	articleIDString := ""
	for _, elem := range articles {
		if articleIDString != "" {
			articleIDString += ","
		}
		articleIDString += strconv.Itoa(elem.ID)
	}
	db.Exec("DELETE FROM votes WHERE ArticleId IN (" + articleIDString + ");")
	result.Success = true
	return result, err
}
