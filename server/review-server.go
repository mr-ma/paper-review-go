package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"../data"
	"../model"
	"github.com/alexedwards/scs"
	"github.com/alexedwards/scs/stores/mysqlstore"
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
	Count    int         `json:"count"`
	Response interface{} `json:"response"`
}
type Page struct {
	Title string
	Body  []byte
}

type fn func(http.ResponseWriter, *http.Request)

var (
	mysqlUser     = flag.String("mysqluser", "foo", "a mysql user")
	mysqlPassword = flag.String("mysqlpass", "bar", "the mysql password")
	mysqlServer   = flag.String("mysqlserver", "", "Mysql server address")
	cert          = flag.String("cert", "", "certificate pathname")
	key           = flag.String("key", "", "private key pathname")
	config        = flag.String("config", "", "pathname of JSON configuration file")
	listen        = flag.String("listen", "127.0.0.1:8004", "listen address")
)

var sessionManager *scs.Manager
var paperReviewDriver data.PaperReviewDriver

func main() {
	flag.Parse()

	paperReviewDriver = data.InitPaperReviewDriver(*mysqlUser, *mysqlPassword, *mysqlServer)
	dbRef, err := paperReviewDriver.OpenDB()
	if err == nil {
		sessionManager = scs.NewManager(mysqlstore.New(dbRef, 600000000000))
		sessionManager.Lifetime(time.Hour * 24) // session data expires after 24 hours
		sessionManager.Persist(true)            // session data persists after the browser has been closed by the user
		//sessionManager.Secure(true)
	}

	cors := tigertonic.NewCORSBuilder().AddAllowedOrigins(*listen)

	mux := tigertonic.NewTrieServeMux()

	mux.Handle("POST", "/paperreview/research", cors.Build(tigertonic.Timed(tigertonic.Marshaled(postResearchHandler), "postResearchHandler", nil)))
	mux.Handle("GET", "/paperreview/research/{id}", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getResearchHandler), "getResearchHandler", nil)))
	mux.Handle("GET", "/paperreview/research", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getResearchHandler), "getAllResearchHandler", nil)))

	mux.Handle("POST", "/paperreview/vote", cors.Build(tigertonic.Timed(tigertonic.Marshaled(postVoteHandler), "postVoteHandler", nil)))
	mux.Handle("GET", "/paperreview/vote/{id}", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getVoteHandler), "getVoteHandler", nil)))
	mux.Handle("GET", "/paperreview/votes/{researchID}", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getVotesHandler), "getResearchVotesHandler", nil)))
	mux.Handle("GET", "/paperreview/votes", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getVotesHandler), "getAllVotesHandler", nil)))

	mux.Handle("GET", "/paperreview/review/{researchID}/{mitarbeiterID}", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getReviewHandler), "getAllReviewsHandler", nil)))
	mux.Handle("GET", "/paperreview/review/{researchID}/{mitarbeiterID}/{limit}", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getReviewHandler), "getNumReviewsHandler", nil)))
	mux.Handle("GET", "/paperreview/review/stats/{researchID}/{mitarbeiterID}", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getReviewStatsHandler), "getReviewStatsPerMitarbeiterHandler", nil)))
	mux.Handle("GET", "/paperreview/review/stats/{researchID}", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getReviewStatsHandler), "getReviewStatsHandler", nil)))

	mux.Handle("GET", "/paperreview/approved/{researchID}/{threshold}", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getApprovedHandler), "getApprovedHandler", nil)))

	// mux.Handle("POST", "/paperreview/mitarbeiter", cors.Build(tigertonic.Timed(tigertonic.Marshaled(postMitarbeiterHandler), "postMitarbeiterHandler", nil)))
	// mux.Handle("GET", "/paperreview/mitarbeiter/{id}", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getMitarbeiterHandler), "getMitarbeiterHandler", nil)))
	// mux.Handle("GET", "/paperreview/mitarbeiter", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getMitarbeiterHandler), "getAllMitarbeitersHandler", nil)))

	mux.Handle("GET", "/paperreview/tag/{researchID}", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getTagsHandler), "getResearchTags", nil)))

	mux.HandleFunc("POST", "/paperreview/getReviewList", func(w http.ResponseWriter, r *http.Request) {
		checkAdmin(w, r, getReviewListHandler)
	})
	mux.HandleFunc("POST", "/paperreview/saveReviewMappings", func(w http.ResponseWriter, r *http.Request) {
		checkAdmin(w, r, saveReviewMappingsHandler)
	})
	mux.HandleFunc("POST", "/paperreview/deleteArticleVotes", func(w http.ResponseWriter, r *http.Request) {
		checkAdmin(w, r, deleteArticleVotesHandler)
	})

	// mux.HandleFunc("GET", "/paperreview/review", func(w http.ResponseWriter, r *http.Request) {
	// 	p := loadPage("../frontend/src/review.html")
	// 	fmt.Fprintf(w, "%s", p)
	// })
	// mux.HandleFunc("GET", "/paperreview/landing", func(w http.ResponseWriter, r *http.Request) {
	// 	p := loadPage("../frontend/src/landing.html")
	// 	fmt.Fprintf(w, "%s", p)
	// })
	// mux.HandleFunc("GET", "/paperreview/approve", func(w http.ResponseWriter, r *http.Request) {
	// 	p := loadPage("../frontend/src/approve.html")
	// 	fmt.Fprintf(w, "%s", p)
	// })

	server := tigertonic.NewServer(*listen, tigertonic.Logged(sessionManager.Use(mux), nil)) // context.ClearHandler(mux), to avoid memory leaks
	go func() {
		var err error
		if "" != *cert && "" != *key {
			err = server.ListenAndServeTLS(*cert, *key)
		} else {
			err = server.ListenAndServe()
		}
		if nil != err {
			checkErr(err)
		}
	}()
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	fmt.Println(<-ch)
	server.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func saveFile(w http.ResponseWriter, file multipart.File, handle *multipart.FileHeader) {
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return
	}

	err = ioutil.WriteFile("./files/"+handle.Filename, data, 0666)
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return
	}
	jsonResponse(w, http.StatusCreated, "File uploaded successfully!.")
}

func jsonResponse(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	fmt.Fprint(w, message)
}

// check if user is admin
func checkAdmin(w http.ResponseWriter, r *http.Request, callback fn) {
	session := sessionManager.Load(r)
	var admin int
	adminStr, err := session.GetString("admin")
	if err != nil {
		admin = 0
	} else {
		admin, err = strconv.Atoi(adminStr)
		if err != nil {
			admin = 0
		}
	}
	if admin == 1 {
		callback(w, r)
		return
	}
	result := []int{}
	output, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// check if user has edit permissions on a taxonomy with a specific ID
// admins can edit every taxonomy
func checkTaxonomyPermissions(w http.ResponseWriter, r *http.Request, callback fn) {
	session := sessionManager.Load(r)
	var admin int
	adminStr, err := session.GetString("admin")
	if err != nil {
		admin = 0
	} else {
		admin, err = strconv.Atoi(adminStr)
		if err != nil {
			admin = 0
		}
	}
	if admin == 1 {
		callback(w, r)
		return
	}
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	var objmap map[string]*json.RawMessage
	err = json.Unmarshal(body, &objmap)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	var taxonomyId64 int64
	err = json.Unmarshal(*objmap["taxonomy_id"], &taxonomyId64)
	if err == nil {
		taxonomyId := int(taxonomyId64)
		taxonomyIDs := []int{}
		taxonomyString, err := session.GetString("taxonomies")
		if err == nil && taxonomyString != "" {
			taxonomyArray := strings.Split(taxonomyString, ",")
			for _, elem := range taxonomyArray {
				id, err := strconv.Atoi(elem)
				if err == nil {
					taxonomyIDs = append(taxonomyIDs, id)
				}
			}
		}
		if len(taxonomyIDs) > 0 && contains(taxonomyIDs, taxonomyId) {
			r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
			callback(w, r)
			return
		}
	}
	result := []int{}
	output, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}

func postResearchHandler(u *url.URL, h http.Header, research *model.Research) (int, http.Header, *MyResponse, error) {
	//driver := data.InitPaperReviewDriver(*mysqlUser, *mysqlPassword)
	_, id, err := paperReviewDriver.InsertResearch(*research)
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{strconv.FormatInt(id, 10), 1, "Research inserted"}, nil
}
func postVoteHandler(u *url.URL, h http.Header, vote *model.Vote) (int, http.Header, *MyResponse, error) {
	if vote.State < -1 || vote.State > 1 {
		return http.StatusNotAcceptable, nil, &MyResponse{"0", 1, "Vote State value can strictly be set to [-1,1]"}, nil
	}
	if vote.Voter.ID <= 0 {
		return http.StatusNotAcceptable, nil, &MyResponse{"0", 1, "Mitarbeiter id is missing"}, nil
	}
	if vote.AssociatedArticleID <= 0 {
		return http.StatusNotAcceptable, nil, &MyResponse{"0", 1, "Article id is missing"}, nil
	}
	//driver := data.InitPaperReviewDriver(*mysqlUser, *mysqlPassword)
	_, id, err := paperReviewDriver.InsertVote(*vote)
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{strconv.FormatInt(id, 10), 1, "Vote inserted"}, nil
}
func getResearchHandler(u *url.URL, h http.Header, r *MyRequest) (int, http.Header, *MyResponse, error) {
	//	fmt.Println("in getResearchHandler")
	//driver := data.InitPaperReviewDriver(*mysqlUser, *mysqlPassword)
	resp := MyResponse{}
	if i, _ := strconv.ParseInt(u.Query().Get("id"), 10, 64); i > 0 {
		//	fmt.Println(u.Query().Get("id"))
		research, err := paperReviewDriver.SelectResearchWithArticles(i)
		checkErr(err)
		resp.Response = research
	} else {
		//select all researches
		all, err := paperReviewDriver.SelectAllResearch()
		checkErr(err)
		resp.Response = all
	}
	return http.StatusOK, nil, &resp, nil
}
func getVoteHandler(u *url.URL, h http.Header, r *MyRequest) (int, http.Header, *MyResponse, error) {
	//driver := data.InitPaperReviewDriver(*mysqlUser, *mysqlPassword)
	resp := MyResponse{}
	if i, _ := strconv.ParseInt(u.Query().Get("id"), 10, 64); i > 0 {
		//fmt.Println(u.Query().Get("id"))
		research, err := paperReviewDriver.SelectVote(i)
		checkErr(err)
		resp.Response = research
	} else {
		//select all votes
		all, err := paperReviewDriver.SelectAllVotes()
		checkErr(err)
		resp.Response = all
	}
	return http.StatusOK, nil, &resp, nil
}
func getVotesHandler(u *url.URL, h http.Header, r *MyRequest) (int, http.Header, *MyResponse, error) {
	//driver := data.InitPaperReviewDriver(*mysqlUser, *mysqlPassword)
	resp := MyResponse{}
	if i, _ := strconv.ParseInt(u.Query().Get("researchID"), 10, 64); i > 0 {
		//fmt.Println(u.Query().Get("researchID"))
		research, err := paperReviewDriver.SelectResearchVotes(i)
		//fmt.Printf("HERE tryping %v\n", i)
		checkErr(err)
		resp.Response = research
	} else {
		//select all votes
		all, err := paperReviewDriver.SelectAllVotes()
		checkErr(err)
		resp.Response = all
	}
	return http.StatusOK, nil, &resp, nil
}
func getReviewHandler(u *url.URL, h http.Header, r *MyRequest) (int, http.Header, *MyResponse, error) {
	//driver := data.InitPaperReviewDriver(*mysqlUser, *mysqlPassword)
	researchID, err := strconv.ParseInt(u.Query().Get("researchID"), 10, 32)
	if err != nil {
		return http.StatusNotAcceptable, nil, &MyResponse{"0", 1, "researchID is missing"}, nil
	}
	mitarbeiterID, err := strconv.ParseInt(u.Query().Get("mitarbeiterID"), 10, 64)
	if err != nil {
		return http.StatusNotAcceptable, nil, &MyResponse{"0", 1, "mitarbeiterID is missing"}, nil
	}
	limit := 0
	if i, err := strconv.Atoi(u.Query().Get("limit")); err == nil {
		limit = i
	}

	if limit == 0 {
		a, _, err := paperReviewDriver.ReviewPapers(researchID, mitarbeiterID)
		if err != nil {
			return http.StatusNotAcceptable, nil, &MyResponse{"0", 0, err.Error()}, nil
		}
		return http.StatusOK, nil, &MyResponse{"0", 1, a}, nil
	}
	a, _, err := paperReviewDriver.ReviewNumPapers(researchID, mitarbeiterID, limit)
	if err != nil {
		return http.StatusNotAcceptable, nil, &MyResponse{"0", 0, err.Error()}, nil
	}
	return http.StatusOK, nil, &MyResponse{"0", 1, a}, nil

}

/*
func postMitarbeiterHandler(u *url.URL, h http.Header, mitarbeiter *model.Mitarbeiter) (int, http.Header, *MyResponse, error) {
	//driver := data.InitPaperReviewDriver(*mysqlUser, *mysqlPassword)
	_, id, err := paperReviewDriver.InsertMitarbeiter(*mitarbeiter)
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{strconv.FormatInt(id, 10), 1, "Mitarbeiter inserted"}, nil
}

func getMitarbeiterHandler(u *url.URL, h http.Header, r *MyRequest) (int, http.Header, *MyResponse, error) {
	//driver := data.InitPaperReviewDriver(*mysqlUser, *mysqlPassword)
	resp := MyResponse{}
	if i, _ := strconv.ParseInt(u.Query().Get("id"), 10, 64); i > 0 {
		//fmt.Println(u.Query().Get("id"))
		research, err := paperReviewDriver.SelectMitarbeiter(i)
		checkErr(err)
		resp.Response = research
	} else {
		//select all researches
		all, err := paperReviewDriver.SelectAllMitarbeiters()
		checkErr(err)
		resp.Response = all
	}
	return http.StatusOK, nil, &resp, nil
}
*/

func getTagsHandler(u *url.URL, h http.Header, r *MyRequest) (int, http.Header, *MyResponse, error) {
	//driver := data.InitPaperReviewDriver(*mysqlUser, *mysqlPassword)
	if i, _ := strconv.ParseInt(u.Query().Get("researchID"), 10, 64); i > 0 {
		tags, err := paperReviewDriver.SelectAllTags(i)
		if err != nil {

		}
		return http.StatusOK, nil, &MyResponse{"0", 1, tags}, nil
	}
	return http.StatusNotAcceptable, nil, &MyResponse{"0", 1, "researchID is missing/malformed"}, nil

}

func getApprovedHandler(u *url.URL, h http.Header, r *MyRequest) (int, http.Header, *MyResponse, error) {
	//driver := data.InitPaperReviewDriver(*mysqlUser, *mysqlPassword)
	researchID, err := strconv.ParseInt(u.Query().Get("researchID"), 10, 64)
	if err != nil {
		return http.StatusNotAcceptable, nil, &MyResponse{"0", 1, "researchID is missing"}, nil
	}
	threshold, err := strconv.Atoi(u.Query().Get("threshold"))
	if err != nil {
		return http.StatusNotAcceptable, nil, &MyResponse{"0", 0, err.Error()}, nil
	}
	papers, err := paperReviewDriver.GetApprovedPapers(researchID, threshold)
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{"0", len(papers), papers}, nil
}

func getReviewStatsHandler(u *url.URL, h http.Header, r *MyRequest) (int, http.Header, *MyResponse, error) {
	//driver := data.InitPaperReviewDriver(*mysqlUser, *mysqlPassword)
	resp := MyResponse{}
	researchID, err := strconv.ParseInt(u.Query().Get("researchID"), 10, 64)
	if err != nil {
		return http.StatusNotAcceptable, nil, &MyResponse{"0", 1, "researchID is missing"}, nil
	}
	if mitarbeiterID, _ := strconv.ParseInt(u.Query().Get("mitarbeiterID"), 10, 64); mitarbeiterID > 0 {
		stats, err := paperReviewDriver.GetResearchStatsPerMitarbeiter(researchID, mitarbeiterID)
		checkErr(err)
		resp.Response = stats
	} else {
		//select all researches
		all, err := paperReviewDriver.GetResearchStats(researchID)
		checkErr(err)
		resp.Response = all
	}
	return http.StatusOK, nil, &resp, nil
}

func getReviewListHandler(w http.ResponseWriter, r *http.Request) {
	var reviewListRequest model.ReviewListRequest
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&reviewListRequest)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	//driver := data.InitPaperReviewDriver(*mysqlUser, *mysqlPassword)
	result, err := paperReviewDriver.GetApprovedPapersWithDetails(reviewListRequest.ResearchID, reviewListRequest.Threshold)
	checkErr(err)
	output, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}

func saveReviewMappingsHandler(w http.ResponseWriter, r *http.Request) {
	var saveReviewMappingsRequest model.SaveReviewMappingsRequest
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&saveReviewMappingsRequest)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	result, err := paperReviewDriver.SaveReviewMappings(saveReviewMappingsRequest.TaxonomyID, saveReviewMappingsRequest.Attributes, saveReviewMappingsRequest.Mappings)
	checkErr(err)
	output, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}

func deleteArticleVotesHandler(w http.ResponseWriter, r *http.Request) {
	var deleteArticleVotesRequest model.DeleteArticleVotesRequest
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&deleteArticleVotesRequest)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	//driver := data.InitPaperReviewDriver(*mysqlUser, *mysqlPassword)
	result, err := paperReviewDriver.DeleteArticleVotes(deleteArticleVotesRequest.Articles)
	checkErr(err)
	output, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}

func loadPage(filename string) (body []byte) {
	// fmt.Println("handling loadpage")
	body, err := ioutil.ReadFile(filename)
	// fmt.Printf("%d", len(body))
	checkErr(err)
	return body
}
