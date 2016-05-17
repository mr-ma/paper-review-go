package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/mr-ma/db"
	"github.com/mr-ma/review"
	"github.com/rcrowley/go-tigertonic"
)

type MyRequest struct {
	ID    string      `json:"id"`
	Stuff interface{} `json:"stuff"`
}
type MyResponse struct {
	ID       string      `json:"id"`
	Response interface{} `json:"response"`
}

// func myHandler(u *url.URL, h http.Header, r *MyRequest) (int, http.Header, *MyResponse, error) {
// 	return http.StatusOK, nil, &MyResponse{"ID", "STUFF"}, nil
// }
func postResearchHandler(u *url.URL, h http.Header, research *review.Research) (int, http.Header, *MyResponse, error) {
	driver := db.InitMySQLDriver()
	_, id, err := driver.InsertResearch(*research)
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{strconv.FormatInt(id, 10), "Research inserted"}, nil
}
func postVoteHandler(u *url.URL, h http.Header, vote *review.Vote) (int, http.Header, *MyResponse, error) {
	if vote.State < -1 || vote.State > 1 {
		return http.StatusNotAcceptable, nil, &MyResponse{"0", "Vote State value can strictly be set to [-1,1]"}, nil
	}
	if vote.Voter.ID <= 0 {
		return http.StatusNotAcceptable, nil, &MyResponse{"0", "Mitarbeiter id is missing"}, nil
	}
	if vote.AssociatedArticleID <= 0 {
		return http.StatusNotAcceptable, nil, &MyResponse{"0", "Article id is missing"}, nil
	}
	driver := db.InitMySQLDriver()
	_, id, err := driver.InsertVote(*vote)
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{strconv.FormatInt(id, 10), "Vote inserted"}, nil
}
func getResearchHandler(u *url.URL, h http.Header, r *MyRequest) (int, http.Header, *MyResponse, error) {
	fmt.Println("in getResearchHandler")
	driver := db.InitMySQLDriver()
	resp := MyResponse{}
	if i, _ := strconv.ParseInt(u.Query().Get("id"), 10, 64); i > 0 {
		fmt.Println(u.Query().Get("id"))
		research, err := driver.SelectResearchWithArticles(i)
		checkErr(err)
		resp.Response = research
	} else {
		//select all researches
		all, err := driver.SelectAllResearchWithArticles()
		checkErr(err)
		resp.Response = all
	}
	return http.StatusOK, nil, &resp, nil
}
func getVoteHandler(u *url.URL, h http.Header, r *MyRequest) (int, http.Header, *MyResponse, error) {
	driver := db.InitMySQLDriver()
	resp := MyResponse{}
	if i, _ := strconv.ParseInt(u.Query().Get("id"), 10, 64); i > 0 {
		fmt.Println(u.Query().Get("id"))
		research, err := driver.SelectVote(i)
		checkErr(err)
		resp.Response = research
	} else {
		//select all votes
		all, err := driver.SelectAllVotes()
		checkErr(err)
		resp.Response = all
	}
	return http.StatusOK, nil, &resp, nil
}
func getVotesHandler(u *url.URL, h http.Header, r *MyRequest) (int, http.Header, *MyResponse, error) {
	driver := db.InitMySQLDriver()
	resp := MyResponse{}
	if i, _ := strconv.ParseInt(u.Query().Get("researchID"), 10, 64); i > 0 {
		fmt.Println(u.Query().Get("researchID"))
		research, err := driver.SelectResearchVotes(i)
		fmt.Printf("HERE tryping %v\n", i)
		checkErr(err)
		resp.Response = research
	} else {
		//select all votes
		all, err := driver.SelectAllVotes()
		checkErr(err)
		resp.Response = all
	}
	return http.StatusOK, nil, &resp, nil
}
func getReviewHandler(u *url.URL, h http.Header, r *MyRequest) (int, http.Header, *MyResponse, error) {
	driver := db.InitMySQLDriver()
	researchID, err := strconv.ParseInt(u.Query().Get("researchID"), 10, 32)
	if err != nil {
		return http.StatusNotAcceptable, nil, &MyResponse{"0", "researchID is missing"}, nil
	}
	mitarbeiterID, err := strconv.ParseInt(u.Query().Get("mitarbeiterID"), 10, 64)
	if err != nil {
		return http.StatusNotAcceptable, nil, &MyResponse{"0", "mitarbeiterID is missing"}, nil
	}
	limit := 0
	if i, err := strconv.Atoi(u.Query().Get("limit")); err == nil {
		limit = i
	}

	if limit == 0 {
		a, _, err := driver.ReviewPapers(researchID, mitarbeiterID)
		if err != nil {
			return http.StatusNotAcceptable, nil, &MyResponse{"0", err.Error()}, nil
		}
		return http.StatusOK, nil, &MyResponse{"0", a}, nil
	}
	a, _, err := driver.ReviewNumPapers(researchID, mitarbeiterID, limit)
	if err != nil {
		return http.StatusNotAcceptable, nil, &MyResponse{"0", err.Error()}, nil
	}
	return http.StatusOK, nil, &MyResponse{"0", a}, nil

}

func postMitarbeiterHandler(u *url.URL, h http.Header, mitarbeiter *review.Mitarbeiter) (int, http.Header, *MyResponse, error) {
	driver := db.InitMySQLDriver()
	_, id, err := driver.InsertMitarbeiter(*mitarbeiter)
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{strconv.FormatInt(id, 10), "Mitarbeiter inserted"}, nil
}
func getMitarbeiterHandler(u *url.URL, h http.Header, r *MyRequest) (int, http.Header, *MyResponse, error) {
	driver := db.InitMySQLDriver()
	resp := MyResponse{}
	if i, _ := strconv.ParseInt(u.Query().Get("id"), 10, 64); i > 0 {
		fmt.Println(u.Query().Get("id"))
		research, err := driver.SelectMitarbeiter(i)
		checkErr(err)
		resp.Response = research
	} else {
		//select all researches
		all, err := driver.SelectAllMitarbeiters()
		checkErr(err)
		resp.Response = all
	}
	return http.StatusOK, nil, &resp, nil
}

func getTagsHandler(u *url.URL, h http.Header, r *MyRequest) (int, http.Header, *MyResponse, error) {
	driver := db.InitMySQLDriver()
	tags, err := driver.SelectAllTags()
	if err != nil {
		return http.StatusNotAcceptable, nil, &MyResponse{"0", err.Error()}, nil
	}
	return http.StatusOK, nil, &MyResponse{"0", tags}, nil
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Printf("Hello, world.\n")
	mux := tigertonic.NewTrieServeMux()
	mux.Handle("POST", "/research", tigertonic.Timed(tigertonic.Marshaled(postResearchHandler), "postResearchHandler", nil))
	mux.Handle("GET", "/research/{id}", tigertonic.Timed(tigertonic.Marshaled(getResearchHandler), "getResearchHandler", nil))
	mux.Handle("GET", "/research", tigertonic.Timed(tigertonic.Marshaled(getResearchHandler), "getAllResearchHandler", nil))

	mux.Handle("POST", "/vote", tigertonic.Timed(tigertonic.Marshaled(postVoteHandler), "postVoteHandler", nil))
	mux.Handle("GET", "/vote/{id}", tigertonic.Timed(tigertonic.Marshaled(getVoteHandler), "getVoteHandler", nil))
	mux.Handle("GET", "/votes/{researchID}", tigertonic.Timed(tigertonic.Marshaled(getVotesHandler), "getResearchVotesHandler", nil))
	mux.Handle("GET", "/votes", tigertonic.Timed(tigertonic.Marshaled(getVotesHandler), "getAllVotesHandler", nil))

	mux.Handle("GET", "/review/{researchID}/{mitarbeiterID}", tigertonic.Timed(tigertonic.Marshaled(getReviewHandler), "getAllReviewsHandler", nil))
	mux.Handle("GET", "/review/{researchID}/{mitarbeiterID}/{limit}", tigertonic.Timed(tigertonic.Marshaled(getReviewHandler), "getNumReviewsHandler", nil))

	mux.Handle("POST", "/mitarbeiter", tigertonic.Timed(tigertonic.Marshaled(postMitarbeiterHandler), "postMitarbeiterHandler", nil))
	mux.Handle("GET", "/mitarbeiter/{id}", tigertonic.Timed(tigertonic.Marshaled(getMitarbeiterHandler), "getMitarbeiterHandler", nil))
	mux.Handle("GET", "/mitarbeiter", tigertonic.Timed(tigertonic.Marshaled(getMitarbeiterHandler), "getAllMitarbeitersHandler", nil))

	mux.Handle("GET", "/tag", tigertonic.Timed(tigertonic.Marshaled(getTagsHandler), "getAllTags", nil))

	//mux.Handle("GET", "/", tigertonic.Timed(tigertonic.Marshaled(myHandler), "myHandler", nil))
	tigertonic.NewServer(":8000", tigertonic.Logged(mux, nil)).ListenAndServe()
}
