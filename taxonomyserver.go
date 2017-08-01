package main

import (
	"net/http"
	"net/url"
"fmt"
"io/ioutil"
	"github.com/mr-ma/paper-review-go/data"
	"github.com/mr-ma/paper-review-go/model"
	"github.com/rcrowley/go-tigertonic"
)

//MyRequest standard request
type MyRequest struct {
	ID    string      `json:"id"`
	Stuff interface{} `json:"stuff"`
}

//MyResponse standard response
type MyResponse struct {
	ID       string      `json:"id"`
	Count int `json:"count"`
	Response interface{} `json:"response"`
}
type Page struct {
    Title string
    Body  []byte
}
// func myHandler(u *url.URL, h http.Header, r *MyRequest) (int, http.Header, *MyResponse, error) {
// 	return http.StatusOK, nil, &MyResponse{"ID", "STUFF"}, nil
// }
// func postResearchHandler(u *url.URL, h http.Header, research *model.Research) (int, http.Header, *MyResponse, error) {
// 	driver := data.InitMySQLDriver()
// 	_, id, err := driver.InsertResearch(*research)
// 	checkErr(err)
// 	return http.StatusOK, nil, &MyResponse{strconv.FormatInt(id, 10), "Research inserted"}, nil
// }
// func postVoteHandler(u *url.URL, h http.Header, vote *model.Vote) (int, http.Header, *MyResponse, error) {
// 	if vote.State < -1 || vote.State > 1 {
// 		return http.StatusNotAcceptable, nil, &MyResponse{"0", "Vote State value can strictly be set to [-1,1]"}, nil
// 	}
// 	if vote.Voter.ID <= 0 {
// 		return http.StatusNotAcceptable, nil, &MyResponse{"0", "Mitarbeiter id is missing"}, nil
// 	}
// 	if vote.AssociatedArticleID <= 0 {
// 		return http.StatusNotAcceptable, nil, &MyResponse{"0", "Article id is missing"}, nil
// 	}
// 	driver := data.InitMySQLDriver()
// 	_, id, err := driver.InsertVote(*vote)
// 	checkErr(err)
// 	return http.StatusOK, nil, &MyResponse{strconv.FormatInt(id, 10), "Vote inserted"}, nil
// }
// func getResearchHandler(u *url.URL, h http.Header, r *MyRequest) (int, http.Header, *MyResponse, error) {
// 	//	fmt.Println("in getResearchHandler")
// 	driver := data.InitMySQLDriver()
// 	resp := MyResponse{}
// 	if i, _ := strconv.ParseInt(u.Query().Get("id"), 10, 64); i > 0 {
// 		//	fmt.Println(u.Query().Get("id"))
// 		research, err := driver.SelectResearchWithArticles(i)
// 		checkErr(err)
// 		resp.Response = research
// 	} else {
// 		//select all researches
// 		all, err := driver.SelectAllResearch()
// 		checkErr(err)
// 		resp.Response = all
// 	}
// 	return http.StatusOK, nil, &resp, nil
// }
// func getVoteHandler(u *url.URL, h http.Header, r *MyRequest) (int, http.Header, *MyResponse, error) {
// 	driver := data.InitMySQLDriver()
// 	resp := MyResponse{}
// 	if i, _ := strconv.ParseInt(u.Query().Get("id"), 10, 64); i > 0 {
// 		//fmt.Println(u.Query().Get("id"))
// 		research, err := driver.SelectVote(i)
// 		checkErr(err)
// 		resp.Response = research
// 	} else {
// 		//select all votes
// 		all, err := driver.SelectAllVotes()
// 		checkErr(err)
// 		resp.Response = all
// 	}
// 	return http.StatusOK, nil, &resp, nil
// }
// func getVotesHandler(u *url.URL, h http.Header, r *MyRequest) (int, http.Header, *MyResponse, error) {
// 	driver := data.InitMySQLDriver()
// 	resp := MyResponse{}
// 	if i, _ := strconv.ParseInt(u.Query().Get("researchID"), 10, 64); i > 0 {
// 		//fmt.Println(u.Query().Get("researchID"))
// 		research, err := driver.SelectResearchVotes(i)
// 		//fmt.Printf("HERE tryping %v\n", i)
// 		checkErr(err)
// 		resp.Response = research
// 	} else {
// 		//select all votes
// 		all, err := driver.SelectAllVotes()
// 		checkErr(err)
// 		resp.Response = all
// 	}
// 	return http.StatusOK, nil, &resp, nil
// }
// func getReviewHandler(u *url.URL, h http.Header, r *MyRequest) (int, http.Header, *MyResponse, error) {
// 	driver := data.InitMySQLDriver()
// 	researchID, err := strconv.ParseInt(u.Query().Get("researchID"), 10, 32)
// 	if err != nil {
// 		return http.StatusNotAcceptable, nil, &MyResponse{"0", "researchID is missing"}, nil
// 	}
// 	mitarbeiterID, err := strconv.ParseInt(u.Query().Get("mitarbeiterID"), 10, 64)
// 	if err != nil {
// 		return http.StatusNotAcceptable, nil, &MyResponse{"0", "mitarbeiterID is missing"}, nil
// 	}
// 	limit := 0
// 	if i, err := strconv.Atoi(u.Query().Get("limit")); err == nil {
// 		limit = i
// 	}
//
// 	if limit == 0 {
// 		a, _, err := driver.ReviewPapers(researchID, mitarbeiterID)
// 		if err != nil {
// 			return http.StatusNotAcceptable, nil, &MyResponse{"0", err.Error()}, nil
// 		}
// 		return http.StatusOK, nil, &MyResponse{"0", a}, nil
// 	}
// 	a, _, err := driver.ReviewNumPapers(researchID, mitarbeiterID, limit)
// 	if err != nil {
// 		return http.StatusNotAcceptable, nil, &MyResponse{"0", err.Error()}, nil
// 	}
// 	return http.StatusOK, nil, &MyResponse{"0", a}, nil
//
// }
//
// func postMitarbeiterHandler(u *url.URL, h http.Header, mitarbeiter *model.Mitarbeiter) (int, http.Header, *MyResponse, error) {
// 	driver := data.InitMySQLDriver()
// 	_, id, err := driver.InsertMitarbeiter(*mitarbeiter)
// 	checkErr(err)
// 	return http.StatusOK, nil, &MyResponse{strconv.FormatInt(id, 10), "Mitarbeiter inserted"}, nil
// }
// func getMitarbeiterHandler(u *url.URL, h http.Header, r *MyRequest) (int, http.Header, *MyResponse, error) {
// 	driver := data.InitMySQLDriver()
// 	resp := MyResponse{}
// 	if i, _ := strconv.ParseInt(u.Query().Get("id"), 10, 64); i > 0 {
// 		//fmt.Println(u.Query().Get("id"))
// 		research, err := driver.SelectMitarbeiter(i)
// 		checkErr(err)
// 		resp.Response = research
// 	} else {
// 		//select all researches
// 		all, err := driver.SelectAllMitarbeiters()
// 		checkErr(err)
// 		resp.Response = all
// 	}
// 	return http.StatusOK, nil, &resp, nil
// }
//
// func getTagsHandler(u *url.URL, h http.Header, r *MyRequest) (int, http.Header, *MyResponse, error) {
// 	driver := data.InitMySQLDriver()
// 	if i, _ := strconv.ParseInt(u.Query().Get("researchID"), 10, 64); i > 0 {
// 		tags, err := driver.SelectAllTags(i)
// 		if err != nil {
//
// 		}
// 		return http.StatusOK, nil, &MyResponse{"0", tags}, nil
// 	}
// 	return http.StatusNotAcceptable, nil, &MyResponse{"0", "researchID is missing/malformed"}, nil
//
// }
//
// func getApprovedHandler(u *url.URL, h http.Header, r *MyRequest) (int, http.Header, *MyResponse, error) {
// 	driver := data.InitMySQLDriver()
// 	researchID, err := strconv.ParseInt(u.Query().Get("researchID"), 10, 64)
// 	if err != nil {
// 		return http.StatusNotAcceptable, nil, &MyResponse{"0", "researchID is missing"}, nil
// 	}
// 	threshold, err := strconv.Atoi(u.Query().Get("threshold"))
// 	if err != nil {
// 		return http.StatusNotAcceptable, nil, &MyResponse{"0", err.Error()}, nil
// 	}
// 	papers, err := driver.GetApprovedPapers(researchID, threshold)
// 	checkErr(err)
// 	return http.StatusOK, nil, &MyResponse{"0", papers}, nil
// }
//
// func getReviewStatsHandler(u *url.URL, h http.Header, r *MyRequest) (int, http.Header, *MyResponse, error) {
// 	driver := data.InitMySQLDriver()
// 	resp := MyResponse{}
// 	researchID, err := strconv.ParseInt(u.Query().Get("researchID"), 10, 64)
// 	if err != nil {
// 		return http.StatusNotAcceptable, nil, &MyResponse{"0", "researchID is missing"}, nil
// 	}
// 	if mitarbeiterID, _ := strconv.ParseInt(u.Query().Get("mitarbeiterID"), 10, 64); mitarbeiterID > 0 {
// 		stats, err := driver.GetResearchStatsPerMitarbeiter(researchID, mitarbeiterID)
// 		checkErr(err)
// 		resp.Response = stats
// 	} else {
// 		//select all researches
// 		all, err := driver.GetResearchStats(researchID)
// 		checkErr(err)
// 		resp.Response = all
// 	}
// 	return http.StatusOK, nil, &resp, nil
// }

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func getCorrelationHandler(u *url.URL, h http.Header, correlationRequest *model.CorrelationRequest) (int, http.Header, *MyResponse, error) {
	if len(correlationRequest.Attributes) == 0 {
		return http.StatusNotAcceptable, nil,
		&MyResponse{"0", 0,"I need some attributes to produce correlations"}, nil
	}
	driver := data.InitClassificationDriver()
	papers, err := driver.ExportCorrelations(
		correlationRequest.Attributes,correlationRequest.TaxonomyID)
	checkErr(err)
	return http.StatusOK, nil,
	&MyResponse{"0",len(papers), papers}, nil
}
func getAttributesHandler(u *url.URL, h http.Header, r *MyRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver()
	attributes, err := driver.GetAllAttributes()
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{"0",len(attributes), attributes}, nil
}
func getCitationsHandler(u *url.URL, h http.Header, r *MyRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver()
	citations, err := driver.GetAllCitations()
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{"0",len(citations), citations}, nil
}

func loadPage(filename string) *Page {
    body, err := ioutil.ReadFile(filename)
		checkErr(err)
    return &Page{Body: body}
}

// func getIndexHandler(w http.ResponseWriter, r *http.Request) {
//  	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
// 	// fmt.Printf((w,"<html><head></head><body><h1>Welcome Home!</h1><a href=\"/static/img/test.gif\">Show Image!</a></body></html>")
// }
func main() {
	//TODO: Remove cors
	cors := tigertonic.NewCORSBuilder().AddAllowedOrigins("*").AddAllowedHeaders("Origin, X-Requested-With, Content-Type, Accept")

	mux := tigertonic.NewTrieServeMux()
	mux.Handle("POST", "/correlation", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getCorrelationHandler), "getCorrelationHandler", nil)))
	mux.Handle("GET", "/attribute", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getAttributesHandler), "getAttributesHandler", nil)))
	mux.Handle("GET", "/citation", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getCitationsHandler), "getCitationsHandler", nil)))
	// mux.Handle("GET","/",cors.Build(tigertonic.Timed(tigertonic.Marshaled(getIndexHandler), "getIndexHandler", nil)))
	mux.HandleFunc("GET", "/", func(w http.ResponseWriter, r *http.Request) {
    p := loadPage("frontend/taxonomy/index.html")
    fmt.Fprintf(w, "%s", p.Body)
	})
	// mux.Handle("POST", "/research", cors.Build(tigertonic.Timed(tigertonic.Marshaled(postResearchHandler), "postResearchHandler", nil)))
	// mux.Handle("GET", "/research/{id}", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getResearchHandler), "getResearchHandler", nil)))
	// mux.Handle("GET", "/research", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getResearchHandler), "getAllResearchHandler", nil)))
	//
	// mux.Handle("POST", "/vote", cors.Build(tigertonic.Timed(tigertonic.Marshaled(postVoteHandler), "postVoteHandler", nil)))
	// mux.Handle("GET", "/vote/{id}", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getVoteHandler), "getVoteHandler", nil)))
	// mux.Handle("GET", "/votes/{researchID}", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getVotesHandler), "getResearchVotesHandler", nil)))
	// mux.Handle("GET", "/votes", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getVotesHandler), "getAllVotesHandler", nil)))
	//
	// mux.Handle("GET", "/review/{researchID}/{mitarbeiterID}", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getReviewHandler), "getAllReviewsHandler", nil)))
	// mux.Handle("GET", "/review/{researchID}/{mitarbeiterID}/{limit}", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getReviewHandler), "getNumReviewsHandler", nil)))
	// mux.Handle("GET", "/review/stats/{researchID}/{mitarbeiterID}", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getReviewStatsHandler), "getReviewStatsPerMitarbeiterHandler", nil)))
	// mux.Handle("GET", "/review/stats/{researchID}", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getReviewStatsHandler), "getReviewStatsHandler", nil)))
	//
	// mux.Handle("GET", "/approved/{researchID}/{threshold}", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getApprovedHandler), "getApprovedHandler", nil)))
	//
	// mux.Handle("POST", "/mitarbeiter", cors.Build(tigertonic.Timed(tigertonic.Marshaled(postMitarbeiterHandler), "postMitarbeiterHandler", nil)))
	// mux.Handle("GET", "/mitarbeiter/{id}", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getMitarbeiterHandler), "getMitarbeiterHandler", nil)))
	// mux.Handle("GET", "/mitarbeiter", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getMitarbeiterHandler), "getAllMitarbeitersHandler", nil)))
	//
	// mux.Handle("GET", "/tag/{researchID}", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getTagsHandler), "getResearchTags", nil)))

	//mux.Handle("GET", "/", tigertonic.Timed(tigertonic.Marshaled(myHandler), "myHandler", nil))
	tigertonic.NewServer(":8001", tigertonic.Logged(mux, nil)).ListenAndServe()
}
