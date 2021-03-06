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
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"time"

	"./data"
	"./model"
	"github.com/alexedwards/scs"
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
	listen        = flag.String("listen", "127.0.0.1:8001", "listen address")
)

var sessionManager = scs.NewCookieManager("u46IpCV9y5Vlur8YvODJEhgOY8m9JVE4")
var usermanagementDriver data.UsermanagementDriver
var classificationDriver data.ClassificationDriver
var taxonomyBuilderDriver data.TaxonomyBuilderDriver
var paperReviewDriver data.PaperReviewDriver

func main() {
	sessionManager.Lifetime(time.Hour * 24) // session data expires after 24 hours
	sessionManager.Persist(true)            // session data persists after the browser has been closed by the user
	//sessionManager.Secure(true)
	flag.Parse()

	usermanagementDriver = data.InitUsermanagementDriver(*mysqlUser, *mysqlPassword, *mysqlServer)
	classificationDriver = data.InitClassificationDriver(*mysqlUser, *mysqlPassword, *mysqlServer)
	taxonomyBuilderDriver = data.InitTaxonomyBuilderDriver(*mysqlUser, *mysqlPassword, *mysqlServer)
	paperReviewDriver = data.InitPaperReviewDriver(*mysqlUser, *mysqlPassword, *mysqlServer)

	cors := tigertonic.NewCORSBuilder().AddAllowedOrigins(*listen)

	mux := tigertonic.NewTrieServeMux()

	// paper-review start
	mux.Handle("POST", "/research", cors.Build(tigertonic.Timed(tigertonic.Marshaled(postResearchHandler), "postResearchHandler", nil)))
	mux.Handle("GET", "/research/{id}", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getResearchHandler), "getResearchHandler", nil)))
	mux.Handle("GET", "/research", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getResearchHandler), "getAllResearchHandler", nil)))

	mux.Handle("POST", "/vote", cors.Build(tigertonic.Timed(tigertonic.Marshaled(postVoteHandler), "postVoteHandler", nil)))
	mux.Handle("GET", "/vote/{id}", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getVoteHandler), "getVoteHandler", nil)))
	mux.Handle("GET", "/votes/{researchID}", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getVotesHandler), "getResearchVotesHandler", nil)))
	mux.Handle("GET", "/votes", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getVotesHandler), "getAllVotesHandler", nil)))

	mux.Handle("GET", "/review/{researchID}/{mitarbeiterID}", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getReviewHandler), "getAllReviewsHandler", nil)))
	mux.Handle("GET", "/review/{researchID}/{mitarbeiterID}/{limit}", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getReviewHandler), "getNumReviewsHandler", nil)))
	mux.Handle("GET", "/review/stats/{researchID}/{mitarbeiterID}", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getReviewStatsHandler), "getReviewStatsPerMitarbeiterHandler", nil)))
	mux.Handle("GET", "/review/stats/{researchID}", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getReviewStatsHandler), "getReviewStatsHandler", nil)))

	mux.Handle("GET", "/approved/{researchID}/{threshold}", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getApprovedHandler), "getApprovedHandler", nil)))

	// mux.Handle("POST", "/mitarbeiter", cors.Build(tigertonic.Timed(tigertonic.Marshaled(postMitarbeiterHandler), "postMitarbeiterHandler", nil)))
	// mux.Handle("GET", "/mitarbeiter/{id}", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getMitarbeiterHandler), "getMitarbeiterHandler", nil)))
	// mux.Handle("GET", "/mitarbeiter", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getMitarbeiterHandler), "getAllMitarbeitersHandler", nil)))

	mux.Handle("GET", "/tag/{researchID}", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getTagsHandler), "getResearchTags", nil)))

	// paper-review end

	// taxonomyserver start

	mux.Handle("GET", "/taxonomy", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getTaxonomyHandler), "getTaxonomyHandler", nil)))
	mux.Handle("POST", "/getTaxonomyID", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getTaxonomyIDHandler), "getTaxonomyIDHandler", nil)))
	mux.Handle("POST", "/correlation", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getCorrelationHandler), "getCorrelationHandler", nil)))
	mux.Handle("POST", "/attributesPerDimension", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getAttributesPerDimensionHandler), "getAttributesPerDimensionHandler", nil)))
	mux.Handle("POST", "/leafAttributesPerDimension", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getLeafAttributesPerDimensionHandler), "getLeafAttributesPerDimensionHandler", nil)))
	mux.Handle("POST", "/attributeClusterPerDimension", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getAttributeClusterPerDimensionHandler), "getAttributeClusterPerDimensionHandler", nil)))
	mux.Handle("POST", "/allChildrenAttributes", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getAllChildrenAttributesHandler), "getAllChildrenAttributesHandler", nil)))
	mux.Handle("POST", "/allChildrenLeafAttributes", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getAllChildrenLeafAttributesHandler), "getAllChildrenLeafAttributesHandler", nil)))
	mux.Handle("POST", "/attributeRelations", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getAttributeRelationsHandler), "getAttributeRelationsHandler", nil)))
	mux.Handle("POST", "/interdimensionalRelations", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getInterdimensionalRelationsHandler), "getInterdimensionalRelationsHandler", nil)))
	mux.Handle("POST", "/intermediateAttributes", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getIntermediateAttributesHandler), "getIntermediateAttributesHandler", nil)))
	mux.Handle("POST", "/majorAttributes", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getMajorAttributesHandler), "getMajorAttributesHandler", nil)))
	mux.Handle("POST", "/citationsPerAttribute", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getCitationsPerAttributeHandler), "getCitationsPerAttributeHandler", nil)))
	mux.Handle("POST", "/citationsPerAttributeIncludingChildren", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getCitationsPerAttributeIncludingChildrenHandler), "getCitationsPerAttributeIncludingChildrenHandler", nil)))
	mux.Handle("POST", "/attribute", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getAttributesHandler), "getAttributesHandler", nil)))
	mux.Handle("POST", "/leafAttribute", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getLeafAttributesHandler), "getLeafAttributesHandler", nil)))
	mux.Handle("POST", "/dimension", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getDimensionsHandler), "getDimensionsHandler", nil)))
	mux.Handle("POST", "/citation", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getCitationsHandler), "getCitationsHandler", nil)))
	mux.Handle("POST", "/citationCount", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getCitationCountHandler), "getCitationCountHandler", nil)))
	mux.Handle("POST", "/citationCounts", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getCitationCountsHandler), "getCitationCountsHandler", nil)))
	mux.Handle("POST", "/citationCountsIncludingChildren", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getCitationCountsIncludingChildrenHandler), "getCitationCountsIncludingChildrenHandler", nil)))
	mux.Handle("GET", "/relationTypes", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getRelationTypesHandler), "getRelationTypesHandler", nil)))
	mux.Handle("POST", "/conceptCorrelation", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getConceptCorrelationsHandler), "getConceptCorrelationsHandler", nil)))
	mux.Handle("POST", "/conceptCorrelation3D", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getConceptCorrelations3DHandler), "getConceptCorrelations3DHandler", nil)))
	mux.Handle("POST", "/conceptCorrelationWithReferenceCounts", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getConceptCorrelationsWithReferenceCountsHandler), "getConceptCorrelationsWithReferenceCountsHandler", nil)))
	mux.Handle("POST", "/conceptCorrelationWithReferenceCounts3D", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getConceptCorrelationsWithReferenceCounts3DHandler), "getConceptCorrelationsWithReferenceCounts3DHandler", nil)))
	mux.Handle("POST", "/allConceptCorrelations", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getAllConceptCorrelationsHandler), "getAllConceptCorrelationsHandler", nil)))
	mux.Handle("POST", "/allConceptCorrelations3D", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getAllConceptCorrelations3DHandler), "getAllConceptCorrelations3DHandler", nil)))
	mux.Handle("POST", "/allConceptCorrelationsWithReferenceCounts", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getAllConceptCorrelationsWithReferenceCountsHandler), "getAllConceptCorrelationsWithReferenceCountsHandler", nil)))
	mux.Handle("POST", "/allConceptCorrelationsWithReferenceCounts3D", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getAllConceptCorrelationsWithReferenceCounts3DHandler), "getAllConceptCorrelationsWithReferenceCounts3DHandler", nil)))
	mux.Handle("POST", "/parentRelationsPerAttribute", cors.Build(tigertonic.Timed(tigertonic.Marshaled(GetParentRelationsPerAttributeHandler), "GetParentRelationsPerAttributeHandler", nil)))
	mux.Handle("POST", "/childRelationsPerAttribute", cors.Build(tigertonic.Timed(tigertonic.Marshaled(GetChildRelationsPerAttributeHandler), "GetChildRelationsPerAttributeHandler", nil)))
	mux.Handle("POST", "/sharedPapers", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getSharedPapersHandler), "getSharedPapersHandler", nil)))
	mux.Handle("POST", "/sharedPapers3D", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getSharedPapers3DHandler), "getSharedPapers3DHandler", nil)))
	mux.Handle("POST", "/sharedPapersIncludingChildren", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getSharedPapersIncludingChildrenHandler), "getSharedPapersIncludingChildrenHandler", nil)))
	mux.Handle("POST", "/sharedPapersIncludingChildren3D", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getSharedPapersIncludingChildren3DHandler), "getSharedPapersIncludingChildren3DHandler", nil)))
	mux.Handle("POST", "/attributeDetails", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getAttributeDetailsHandler), "getAttributeDetailsHandler", nil)))
	mux.Handle("POST", "/citationDetails", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getCitationDetailsHandler), "getCitationDetailsHandler", nil)))
	mux.Handle("POST", "/attributeCoverage", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getAttributeCoverageHandler), "getAttributeCoverageHandler", nil)))
	mux.Handle("POST", "/attributeCoverageWithOccurrenceCounts", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getAttributeCoverageWithOccurrenceCountsHandler), "getAttributeCoverageWithOccurrenceCountsHandler", nil)))
	mux.Handle("POST", "/attributeCoverageWithReferenceCounts", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getAttributeCoverageWithReferenceCountsHandler), "getAttributeCoverageWithReferenceCountsHandler", nil)))
	mux.Handle("POST", "/attributesByName", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getAttributesByNameHandler), "getAttributesByNameHandler", nil)))
	mux.Handle("POST", "/kMeans", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getKMeansHandler), "getKMeansHandler", nil)))

	// taxonomyserver end

	// user and access management:

	mux.HandleFunc("POST", "/login", func(w http.ResponseWriter, r *http.Request) {
		var loginRequest model.LoginRequest
		if r.Body == nil {
			http.Error(w, "Please send a request body", 400)
			return
		}
		err := json.NewDecoder(r.Body).Decode(&loginRequest)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		if loginRequest.Password == "" {
			session := sessionManager.Load(r)
			var email string
			email, err := session.GetString("email")
			if err != nil {
				email = ""
			}
			var adminInt int
			admin, err := session.GetString("admin")
			if err != nil {
				adminInt = 0
			}
			adminInt, err = strconv.Atoi(admin)
			if err != nil {
				adminInt = 0
			}
			var taxonomies string
			taxonomies, err = session.GetString("taxonomies")
			if err != nil {
				taxonomies = ""
			}
			userResult := model.LoginResult{Success: true}
			userResult.User = model.User{Email: email, Admin: adminInt, Taxonomies: taxonomies}
			output, err := json.Marshal(userResult)
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			w.Header().Set("content-type", "application/json")
			w.Write(output)
			return
		}
		// //driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
		loginResult, err := usermanagementDriver.Login(loginRequest.Email, loginRequest.Password)
		checkErr(err)
		if loginResult.Success {
			session := sessionManager.Load(r)
			err := session.PutString(w, "email", loginResult.User.Email)
			if err != nil {
				http.Error(w, err.Error(), 500)
			}
			err = session.PutString(w, "admin", strconv.Itoa(loginResult.User.Admin))
			if err != nil {
				http.Error(w, err.Error(), 500)
			}
			err = session.PutString(w, "taxonomies", loginResult.User.Taxonomies)
			if err != nil {
				http.Error(w, err.Error(), 500)
			}
		}
		output, err := json.Marshal(loginResult)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.Header().Set("content-type", "application/json")
		w.Write(output)
	})
	mux.HandleFunc("POST", "/logout", func(w http.ResponseWriter, r *http.Request) {
		session := sessionManager.Load(r)
		err := session.Destroy(w)
		result := model.Result{}
		if err != nil {
			result.Success = false
		} else {
			result.Success = true
		}
		output, err := json.Marshal(result)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.Header().Set("content-type", "application/json")
		w.Write(output)
	})
	mux.HandleFunc("POST", "/saveUser", func(w http.ResponseWriter, r *http.Request) {
		var loginRequest model.LoginRequest
		if r.Body == nil {
			http.Error(w, "Please send a request body", 400)
			return
		}
		err := json.NewDecoder(r.Body).Decode(&loginRequest)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		//driver = data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
		result, err := usermanagementDriver.SaveUser(loginRequest.Email, loginRequest.Password)
		checkErr(err)
		userResult := model.LoginResult{}
		if result.Success {
			userResult.Success = true
			userResult.User = model.User{Email: loginRequest.Email, Admin: 0}
			session := sessionManager.Load(r)
			session.Clear(w)
			session.RenewToken(w)
			err := session.PutString(w, "email", loginRequest.Email)
			if err != nil {
				http.Error(w, err.Error(), 500)
			}
			err = session.PutString(w, "admin", "0")
			if err != nil {
				http.Error(w, err.Error(), 500)
			}
		} else {
			userResult.Success = false
		}
		output, err := json.Marshal(userResult)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.Header().Set("content-type", "application/json")
		w.Write(output)
	})
	mux.HandleFunc("GET", "/user", func(w http.ResponseWriter, r *http.Request) {
		checkAdmin(w, r, getUserHandler)
	})
	mux.HandleFunc("GET", "/getUsers", func(w http.ResponseWriter, r *http.Request) {
		checkAdmin(w, r, getUsersHandler)
	})
	mux.HandleFunc("POST", "/query", func(w http.ResponseWriter, r *http.Request) {
		checkAdmin(w, r, queryHandler)
	})
	mux.HandleFunc("POST", "/createUser", func(w http.ResponseWriter, r *http.Request) {
		checkAdmin(w, r, createUserHandler)
	})
	mux.HandleFunc("POST", "/updateUser", func(w http.ResponseWriter, r *http.Request) {
		checkAdmin(w, r, updateUserHandler)
	})
	mux.HandleFunc("POST", "/deleteUser", func(w http.ResponseWriter, r *http.Request) {
		checkAdmin(w, r, deleteUserHandler)
	})
	mux.HandleFunc("POST", "/taxonomyPermissions", func(w http.ResponseWriter, r *http.Request) {
		checkAdmin(w, r, getTaxonomyPermissionsHandler)
	})
	mux.HandleFunc("POST", "/updateTaxonomyPermissions", func(w http.ResponseWriter, r *http.Request) {
		checkAdmin(w, r, getUpdateTaxonomyPermissionsHandler)
	})
	mux.HandleFunc("POST", "/addTaxonomy", func(w http.ResponseWriter, r *http.Request) {
		checkAdmin(w, r, getAddTaxonomyHandler)
	})
	mux.HandleFunc("POST", "/removeTaxonomy", func(w http.ResponseWriter, r *http.Request) {
		checkAdmin(w, r, getRemoveTaxonomyHandler)
	})
	mux.HandleFunc("POST", "/deleteCitation", func(w http.ResponseWriter, r *http.Request) {
		checkTaxonomyPermissions(w, r, deleteCitationHandler)
	})
	mux.HandleFunc("POST", "/savePositions", func(w http.ResponseWriter, r *http.Request) {
		checkTaxonomyPermissions(w, r, savePositionsHandler)
	})
	mux.HandleFunc("POST", "/saveMajorPositions", func(w http.ResponseWriter, r *http.Request) {
		checkTaxonomyPermissions(w, r, saveMajorPositionsHandler)
	})
	mux.HandleFunc("POST", "/save3DPositions", func(w http.ResponseWriter, r *http.Request) {
		checkTaxonomyPermissions(w, r, save3DPositionsHandler)
	})
	mux.HandleFunc("POST", "/saveMajor3DPositions", func(w http.ResponseWriter, r *http.Request) {
		checkTaxonomyPermissions(w, r, saveMajor3DPositionsHandler)
	})
	mux.HandleFunc("POST", "/saveEdgeBendPoints", func(w http.ResponseWriter, r *http.Request) {
		checkTaxonomyPermissions(w, r, saveEdgeBendPointsHandler)
	})
	mux.HandleFunc("POST", "/addAttribute", func(w http.ResponseWriter, r *http.Request) {
		checkTaxonomyPermissions(w, r, addAttributeHandler)
	})
	mux.HandleFunc("POST", "/addDimension", func(w http.ResponseWriter, r *http.Request) {
		checkTaxonomyPermissions(w, r, addDimensionHandler)
	})
	mux.HandleFunc("POST", "/changeDimension", func(w http.ResponseWriter, r *http.Request) {
		checkTaxonomyPermissions(w, r, changeDimensionHandler)
	})
	mux.HandleFunc("POST", "/removeAttribute", func(w http.ResponseWriter, r *http.Request) {
		checkTaxonomyPermissions(w, r, removeAttributeHandler)
	})
	mux.HandleFunc("POST", "/removeDimension", func(w http.ResponseWriter, r *http.Request) {
		checkTaxonomyPermissions(w, r, removeDimensionHandler)
	})
	mux.HandleFunc("POST", "/renameAttribute", func(w http.ResponseWriter, r *http.Request) {
		checkTaxonomyPermissions(w, r, renameAttributeHandler)
	})
	mux.HandleFunc("POST", "/updateSynonyms", func(w http.ResponseWriter, r *http.Request) {
		checkTaxonomyPermissions(w, r, updateSynonymsHandler)
	})
	mux.HandleFunc("POST", "/renameDimension", func(w http.ResponseWriter, r *http.Request) {
		checkTaxonomyPermissions(w, r, renameDimensionHandler)
	})
	mux.HandleFunc("POST", "/addTaxonomyRelation", func(w http.ResponseWriter, r *http.Request) {
		checkTaxonomyPermissions(w, r, addTaxonomyRelationHandler)
	})
	mux.HandleFunc("POST", "/removeTaxonomyRelation", func(w http.ResponseWriter, r *http.Request) {
		checkTaxonomyPermissions(w, r, removeTaxonomyRelationHandler)
	})
	mux.HandleFunc("POST", "/updateTaxonomyRelationType", func(w http.ResponseWriter, r *http.Request) {
		checkTaxonomyPermissions(w, r, updateTaxonomyRelationTypeHandler)
	})
	mux.HandleFunc("POST", "/updateTaxonomyRelationAnnotation", func(w http.ResponseWriter, r *http.Request) {
		checkTaxonomyPermissions(w, r, updateTaxonomyRelationAnnotationHandler)
	})
	mux.HandleFunc("POST", "/updateMajor", func(w http.ResponseWriter, r *http.Request) {
		checkTaxonomyPermissions(w, r, getUpdateMajorHandler)
	})
	mux.HandleFunc("POST", "/updateCitationMapping", func(w http.ResponseWriter, r *http.Request) {
		checkTaxonomyPermissions(w, r, getUpdateCitationMappingHandler)
	})
	mux.HandleFunc("POST", "/updateCitationMappings", func(w http.ResponseWriter, r *http.Request) {
		checkTaxonomyPermissions(w, r, getUpdateCitationMappingsHandler)
	})
	mux.HandleFunc("POST", "/updateCitationReferenceCounts", func(w http.ResponseWriter, r *http.Request) {
		checkTaxonomyPermissions(w, r, getUpdateCitationReferenceCountsHandler)
	})
	mux.HandleFunc("POST", "/mergeAttributes", func(w http.ResponseWriter, r *http.Request) {
		checkTaxonomyPermissions(w, r, getMergeAttributesHandler)
	})
	mux.HandleFunc("POST", "/forkAttribute", func(w http.ResponseWriter, r *http.Request) {
		checkTaxonomyPermissions(w, r, getForkAttributeHandler)
	})

	// make taxonomy specific?
	mux.HandleFunc("POST", "/getReviewList", func(w http.ResponseWriter, r *http.Request) {
		checkAdmin(w, r, getReviewListHandler)
	})
	mux.HandleFunc("POST", "/saveReviewMappings", func(w http.ResponseWriter, r *http.Request) {
		checkAdmin(w, r, saveReviewMappingsHandler)
	})
	mux.HandleFunc("POST", "/deleteArticleVotes", func(w http.ResponseWriter, r *http.Request) {
		checkAdmin(w, r, deleteArticleVotesHandler)
	})

	mux.HandleFunc("GET", "/review", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/review.html")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/landing", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/landing.html")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/approve", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/approve.html")
		fmt.Fprintf(w, "%s", p)
	})

	mux.HandleFunc("GET", "/error.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/error.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/stringComparison.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/stringComparison.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/tables.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/tables.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/users.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/users.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/userManagement.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/userManagement.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/exportHTML.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/exportHTML.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/loadData.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/loadData.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/fileUploader.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/fileUploader.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/scopus.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/scopus.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/scripts.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/scripts.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/style.css", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/css/style.css")
		w.Header().Add("Content-Type", "text/css")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/main.css", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/css/main.css")
		w.Header().Add("Content-Type", "text/css")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/editableTable.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/editableTable.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/pdf.min.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/libs/pdf.min.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/pdf.worker.min.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/libs/pdf.worker.min.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/lodash.core.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/libs/lodash.core.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/compare-strings.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/libs/compare-strings.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/levenshtein.min.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/libs/levenshtein.min.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/fuzzysort.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/libs/fuzzysort.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/zoomInIcon.png", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/icons/zoom_in_128.png")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/zoomOutIcon.png", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/icons/zoom_out_128.png")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/bluebird.min.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/libs/bluebird.min.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/FileSaver.min.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/libs/FileSaver.min.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/multiselect.min.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/libs/multiselect.min.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/jquery.min.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/libs/jquery.min.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/d3.min.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/libs/d3.min.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/d3.layout.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/libs/d3.layout.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/d3-hierarchy.v1.min.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/libs/d3-hierarchy.v1.min.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/d3-context-menu.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/libs/d3-context-menu.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/d3-context-menu.css", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/css/libs/d3-context-menu.css")
		w.Header().Add("Content-Type", "text/css")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/bootstrap.min.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/libs/bootstrap.min.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/bootstrap.min.css", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/css/libs/bootstrap.min.css")
		w.Header().Add("Content-Type", "text/css")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/bootstrap-waitingfor.min.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/libs/bootstrap-waitingfor.min.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/bootstrap-table.css", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/css/libs/bootstrap-table.css")
		w.Header().Add("Content-Type", "text/css")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/bootstrap-table.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/libs/bootstrap-table.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/selectize.min.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/libs/selectize.min.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/selectize.css", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/css/libs/selectize.css")
		w.Header().Add("Content-Type", "text/css")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/selectize.bootstrap3.css", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/css/libs/selectize.bootstrap3.css")
		w.Header().Add("Content-Type", "text/css")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/fonts/glyphicons-halflings-regular.woff", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/fonts/glyphicons-halflings-regular.woff")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/fonts/glyphicons-halflings-regular.woff2", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/fonts/glyphicons-halflings-regular.woff2")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/fonts/glyphicons-halflings-regular.ttf", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/fonts/glyphicons-halflings-regular.ttf")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/bootstrap-dialog.min.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/libs/bootstrap-dialog.min.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/bootstrap-dialog.min.css", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/css/libs/bootstrap-dialog.min.css")
		w.Header().Add("Content-Type", "text/css")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/d3.v4.min.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/libs/d3.v4.min.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/vis.min.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/libs/vis.min.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/vis.min.css", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/css/libs/vis.min.css")
		w.Header().Add("Content-Type", "text/css")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/cytoscape.min.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/cytoscape/cytoscape.min.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/konva.min.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/cytoscape/extensions/konva.min.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/cytoscape-undo-redo.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/cytoscape/extensions/cytoscape-undo-redo.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/cytoscape-clipboard.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/cytoscape/extensions/cytoscape-clipboard.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/cytoscape-graphml.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/cytoscape/extensions/cytoscape-graphml.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/cytoscape-view-utilities.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/cytoscape/extensions/cytoscape-view-utilities.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/cytoscape-context-menus.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/cytoscape/extensions/cytoscape-context-menus.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/cytoscape-context-menus.css", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/cytoscape/extensions/cytoscape-context-menus.css")
		w.Header().Add("Content-Type", "text/css")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/cytoscape-cxtmenu.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/cytoscape/extensions/cytoscape-cxtmenu.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/cytoscape-autopan-on-drag.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/cytoscape/extensions/cytoscape-autopan-on-drag.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/cytoscape-edge-bend-editing.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/cytoscape/extensions/cytoscape-edge-bend-editing.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/CytoscapeEdgeEditation.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/cytoscape/extensions/CytoscapeEdgeEditation.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/cytoscape-edgehandles.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/cytoscape/extensions/cytoscape-edgehandles.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/cytoscape-expand-collapse.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/cytoscape/extensions/cytoscape-expand-collapse.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/cytoscape.js-navigator.css", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/cytoscape/extensions/cytoscape.js-navigator.css")
		w.Header().Add("Content-Type", "text/css")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/cytoscape-navigator.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/cytoscape/extensions/cytoscape-navigator.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/cytoscape-node-resize.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/cytoscape/extensions/cytoscape-node-resize.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/cytoscape-noderesize.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/cytoscape/extensions/cytoscape-noderesize.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/cytoscape-panzoom.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/cytoscape/extensions/cytoscape-panzoom.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/cytoscape.js-panzoom.css", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/cytoscape/extensions/cytoscape.js-panzoom.css")
		w.Header().Add("Content-Type", "text/css")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/jquery.qtip.min.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/cytoscape/extensions/jquery.qtip.min.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/jquery.qtip.min.css", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/cytoscape/extensions/jquery.qtip.min.css")
		w.Header().Add("Content-Type", "text/css")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/cytoscape-qtip.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/cytoscape/extensions/cytoscape-qtip.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/cytoscape-snap-to-grid.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/cytoscape/extensions/cytoscape-snap-to-grid.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/d3-tip.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/libs/d3-tip.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/d3-tip.css", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/css/libs/d3-tip.css")
		w.Header().Add("Content-Type", "text/css")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/parse-bibtex.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/libs/parse-bibtex.js")
		fmt.Fprintf(w, "%s", p)
	})

	mux.HandleFunc("GET", "/", func(w http.ResponseWriter, r *http.Request) {
		// p := loadPage("frontend/taxonomy/index.html")
		p := loadPage("frontend/taxonomy/correlationMap/interactive/conceptCorrelations_two_dimensional.html")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/users", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/users/users.html")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/modals.html", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/users/modals.html")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/navbar.html", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/users/navbar.html")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/navbarAdmin.html", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/users/navbarAdmin.html")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/circlePacking", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/hierarchy/circlePacking.html")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/conceptCorrelationMatrix2D", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/correlationMap/interactive/conceptCorrelations_two_dimensional.html")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/conceptCorrelationMatrix3D", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/correlationMap/interactive/conceptCorrelations_three_dimensional.html")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/taxonomyRelations", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/cytoscape/taxonomyRelations.html")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/scopus", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/scopus/scopusAPI.html")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/pdf/{file}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("handling pdf")
		fp := filepath.Clean(r.URL.Path)
		var p []byte
		if strings.Contains(fp, "system") {
			p = loadPage("frontend/pdfs/system.pdf")
		} else if strings.Contains(fp, "attack") {
			p = loadPage("frontend/pdfs/attack.pdf")
		} else if strings.Contains(fp, "defense") {
			p = loadPage("frontend/pdfs/defense.pdf")
		} else if strings.Contains(fp, "relations") {
			p = loadPage("frontend/pdfs/view-relations.pdf")
		}
		fmt.Printf("Printing pdf %d", len(p))
		b := bytes.NewBuffer(p)
		// stream straight to client(browser)
		w.Header().Set("Content-type", "application/pdf")

		if _, err := b.WriteTo(w); err != nil { // <----- here!
			fmt.Fprintf(w, "%s", err)
		}
		fmt.Println("handling pdf end")
		w.Write([]byte("PDF Generated"))
	})
	mux.HandleFunc("GET", "/png/{file}", func(w http.ResponseWriter, r *http.Request) {
		fp := filepath.Clean(r.URL.Path)
		var p []byte
		if strings.Contains(fp, "system") {
			p = loadPage("frontend/pngs/system.png")
		} else if strings.Contains(fp, "attack") {
			p = loadPage("frontend/pngs/attack.png")
		} else if strings.Contains(fp, "defense") {
			p = loadPage("frontend/pngs/defense.png")
		} else if strings.Contains(fp, "relations") {
			p = loadPage("frontend/pngs/relations.png")
		}

		b := bytes.NewBuffer(p)
		// stream straight to client(browser)
		w.Header().Set("Content-type", "application/png")

		if _, err := b.WriteTo(w); err != nil { // <----- here!
			fmt.Fprintf(w, "%s", err)
		}
		w.Write([]byte("PNG Generated"))
	})

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

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	session := sessionManager.Load(r)
	var email string
	email, err := session.GetString("email")
	user := model.User{}
	if err == nil {
		// driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
		user, err = usermanagementDriver.GetUser(email)
		checkErr(err)
	}
	output, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}

// sends a SQL query to the database
func queryHandler(w http.ResponseWriter, r *http.Request) {
	var queryRequest model.QueryRequest
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&queryRequest)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	// driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	result, err := usermanagementDriver.QueryDB(queryRequest.Query)
	checkErr(err)
	output, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}

// get list of users and their permissions
func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	// driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	users, err := usermanagementDriver.GetUsers()
	checkErr(err)
	userResult := []model.User{}
	if err == nil {
		userResult = users
	}
	output, err := json.Marshal(userResult)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}

// add user
func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var userRequest model.CreateUserRequest
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	result, err := usermanagementDriver.CreateUser(userRequest.Email, userRequest.Password, userRequest.Admin)
	checkErr(err)
	output, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}

// update user permissions
func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	var userRequest model.UpdateUserRequest
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	result, err := usermanagementDriver.UpdateUser(userRequest.Email, userRequest.Admin)
	checkErr(err)
	output, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}

// delete user
func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	var userRequest model.UserRequest
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	result, err := usermanagementDriver.DeleteUser(userRequest.Email)
	checkErr(err)
	output, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}

func getTaxonomyHandler(u *url.URL, h http.Header, r *MyRequest) (int, http.Header, *MyResponse, error) {
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	taxonomies, err := taxonomyBuilderDriver.GetAllTaxonomies()
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{"0", len(taxonomies), taxonomies}, nil
}

func getTaxonomyIDHandler(u *url.URL, h http.Header, taxonomyIDRequest *model.TaxonomyIDRequest) (int, http.Header, *MyResponse, error) {
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	taxonomyID, err := taxonomyBuilderDriver.GetTaxonomyID(taxonomyIDRequest.Text)
	checkErr(err)
	return http.StatusOK, nil,
		&MyResponse{"0", len(taxonomyID), taxonomyID}, nil
}

// get taxonomy permissions
func getTaxonomyPermissionsHandler(w http.ResponseWriter, r *http.Request) {
	var userRequest model.UserRequest
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	result, err := taxonomyBuilderDriver.GetTaxonomyPermissions(userRequest.Email)
	checkErr(err)
	output, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}

// update taxonomy permissions
func getUpdateTaxonomyPermissionsHandler(w http.ResponseWriter, r *http.Request) {
	var taxonomyPermissionsRequest model.TaxonomyPermissionsRequest
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&taxonomyPermissionsRequest)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	result, err := taxonomyBuilderDriver.UpdateTaxonomyPermissions(taxonomyPermissionsRequest.Email, taxonomyPermissionsRequest.Permissions)
	checkErr(err)
	output, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}

func getAddTaxonomyHandler(w http.ResponseWriter, r *http.Request) {
	var addTaxonomyRequest model.AddTaxonomyRequest
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&addTaxonomyRequest)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	result, err := taxonomyBuilderDriver.AddTaxonomy(addTaxonomyRequest.Taxonomy, addTaxonomyRequest.Dimension)
	checkErr(err)
	output, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}

func getRemoveTaxonomyHandler(w http.ResponseWriter, r *http.Request) {
	var taxonomyRequest model.TaxonomyRequest
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&taxonomyRequest)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	result, err := taxonomyBuilderDriver.RemoveTaxonomy(taxonomyRequest.TaxonomyID)
	checkErr(err)
	output, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}

func getCorrelationHandler(u *url.URL, h http.Header, correlationRequest *model.CorrelationRequest) (int, http.Header, *MyResponse, error) {
	if len(correlationRequest.Attributes) == 0 {
		return http.StatusNotAcceptable, nil,
			&MyResponse{"0", 0, "I need some attributes to produce correlations"}, nil
	}
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	papers, err := classificationDriver.ExportCorrelations(
		correlationRequest.Attributes, correlationRequest.TaxonomyID)
	checkErr(err)
	return http.StatusOK, nil,
		&MyResponse{"0", len(papers), papers}, nil
}
func getAttributesPerDimensionHandler(u *url.URL, h http.Header, attributesPerDimensionRequest *model.AttributeRelationsRequest) (int, http.Header, *MyResponse, error) {
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	attributes, err := classificationDriver.GetAttributesPerDimension(
		attributesPerDimensionRequest.TaxonomyID, attributesPerDimensionRequest.Dimension)
	checkErr(err)
	return http.StatusOK, nil,
		&MyResponse{"0", len(attributes), attributes}, nil
}
func getLeafAttributesPerDimensionHandler(u *url.URL, h http.Header, attributesPerDimensionRequest *model.AttributeRelationsRequest) (int, http.Header, *MyResponse, error) {
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	attributes, err := classificationDriver.GetLeafAttributesPerDimension(
		attributesPerDimensionRequest.TaxonomyID, attributesPerDimensionRequest.Dimension)
	checkErr(err)
	return http.StatusOK, nil,
		&MyResponse{"0", len(attributes), attributes}, nil
}
func getAttributeClusterPerDimensionHandler(u *url.URL, h http.Header, attributeClusterPerDimensionRequest *model.AttributeRelationsRequest) (int, http.Header, *MyResponse, error) {
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	clusters, err := classificationDriver.GetAttributeClusterPerDimension(
		attributeClusterPerDimensionRequest.TaxonomyID, attributeClusterPerDimensionRequest.Dimension)
	checkErr(err)
	return http.StatusOK, nil,
		&MyResponse{"0", len(clusters), clusters}, nil
}
func getAllChildrenAttributesHandler(u *url.URL, h http.Header, allChildrenAttributesRequest *model.AllChildrenAttributesRequest) (int, http.Header, *MyResponse, error) {
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	attributes, err := classificationDriver.GetAllChildrenAttributes(
		allChildrenAttributesRequest.TaxonomyID, allChildrenAttributesRequest.Parent)
	checkErr(err)
	return http.StatusOK, nil,
		&MyResponse{"0", len(attributes), attributes}, nil
}
func getAllChildrenLeafAttributesHandler(u *url.URL, h http.Header, allChildrenAttributesRequest *model.AllChildrenAttributesRequest) (int, http.Header, *MyResponse, error) {
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	attributes, err := classificationDriver.GetAllChildrenLeafAttributes(
		allChildrenAttributesRequest.TaxonomyID, allChildrenAttributesRequest.Parent)
	checkErr(err)
	return http.StatusOK, nil,
		&MyResponse{"0", len(attributes), attributes}, nil
}
func getIntermediateAttributesHandler(u *url.URL, h http.Header, intermediateAttributesRequest *model.AttributeRelationsRequest) (int, http.Header, *MyResponse, error) {
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	intermediateAttributes, err := classificationDriver.GetIntermediateAttributes(
		intermediateAttributesRequest.TaxonomyID, intermediateAttributesRequest.MinValue, intermediateAttributesRequest.MaxValue)
	checkErr(err)
	return http.StatusOK, nil,
		&MyResponse{"0", len(intermediateAttributes), intermediateAttributes}, nil
}
func getMajorAttributesHandler(u *url.URL, h http.Header, majorAttributesRequest *model.TaxonomyRequest) (int, http.Header, *MyResponse, error) {
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	majorAttributes, err := classificationDriver.GetMajorAttributes(
		majorAttributesRequest.TaxonomyID)
	checkErr(err)
	return http.StatusOK, nil,
		&MyResponse{"0", len(majorAttributes), majorAttributes}, nil
}
func getAttributeRelationsHandler(u *url.URL, h http.Header, attributeRelationsRequest *model.AttributeRelationsRequest) (int, http.Header, *MyResponse, error) {
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	attributeRelations, err := classificationDriver.GetAttributeRelationsPerDimension(
		attributeRelationsRequest.TaxonomyID, attributeRelationsRequest.Dimension)
	checkErr(err)
	return http.StatusOK, nil,
		&MyResponse{"0", len(attributeRelations), attributeRelations}, nil
}
func getInterdimensionalRelationsHandler(u *url.URL, h http.Header, interdimensionalRelationsRequest *model.AttributeRelationsRequest) (int, http.Header, *MyResponse, error) {
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	attributeRelations, err := classificationDriver.GetInterdimensionalRelations(
		interdimensionalRelationsRequest.TaxonomyID)
	checkErr(err)
	return http.StatusOK, nil,
		&MyResponse{"0", len(attributeRelations), attributeRelations}, nil
}
func savePositionsHandler(w http.ResponseWriter, r *http.Request) {
	var savePositionsRequest model.SavePositionsRequest
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&savePositionsRequest)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	result, err := taxonomyBuilderDriver.SavePositions(
		savePositionsRequest.TaxonomyID, savePositionsRequest.Positions)
	checkErr(err)
	output, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}
func saveMajorPositionsHandler(w http.ResponseWriter, r *http.Request) {
	var savePositionsRequest model.SavePositionsRequest
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&savePositionsRequest)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	result, err := taxonomyBuilderDriver.SaveMajorPositions(
		savePositionsRequest.TaxonomyID, savePositionsRequest.Positions)
	checkErr(err)
	output, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}
func save3DPositionsHandler(w http.ResponseWriter, r *http.Request) {
	var savePositionsRequest model.SavePositionsRequest
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&savePositionsRequest)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	result, err := taxonomyBuilderDriver.Save3DPositions(
		savePositionsRequest.TaxonomyID, savePositionsRequest.Positions)
	checkErr(err)
	output, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}
func saveMajor3DPositionsHandler(w http.ResponseWriter, r *http.Request) {
	var savePositionsRequest model.SavePositionsRequest
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&savePositionsRequest)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	result, err := taxonomyBuilderDriver.SaveMajor3DPositions(
		savePositionsRequest.TaxonomyID, savePositionsRequest.Positions)
	checkErr(err)
	output, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}
func saveEdgeBendPointsHandler(w http.ResponseWriter, r *http.Request) {
	var saveEdgeBendPointsRequest model.SaveEdgeBendPointsRequest
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&saveEdgeBendPointsRequest)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	result, err := taxonomyBuilderDriver.SaveEdgeBendPoints(
		saveEdgeBendPointsRequest.TaxonomyID, saveEdgeBendPointsRequest.AttributeSrc, saveEdgeBendPointsRequest.AttributeDest, saveEdgeBendPointsRequest.EdgeBendPoints, saveEdgeBendPointsRequest.Dimension)
	checkErr(err)
	output, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}
func getCitationsPerAttributeHandler(u *url.URL, h http.Header, citationsPerAttributeRequest *model.CitationsPerAttributeRequest) (int, http.Header, *MyResponse, error) {
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	citations, err := classificationDriver.GetCitationsPerAttribute(citationsPerAttributeRequest.TaxonomyID, citationsPerAttributeRequest.Attribute)
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{"0", len(citations), citations}, nil
}
func getCitationsPerAttributeIncludingChildrenHandler(u *url.URL, h http.Header, citationsPerAttributeRequest *model.CitationsPerAttributeRequest) (int, http.Header, *MyResponse, error) {
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	citations, err := classificationDriver.GetCitationsPerAttributeIncludingChildren(citationsPerAttributeRequest.TaxonomyID, citationsPerAttributeRequest.Attribute)
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{"0", len(citations), citations}, nil
}
func addAttributeHandler(w http.ResponseWriter, r *http.Request) {
	var attributeRequest model.AttributeRequest
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&attributeRequest)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	attribute := model.Attribute{Text: attributeRequest.Text, X: attributeRequest.X, Y: attributeRequest.Y, XMajor: attributeRequest.XMajor, YMajor: attributeRequest.YMajor, Major: attributeRequest.Major, Dimension: attributeRequest.Dimension}
	result, err := taxonomyBuilderDriver.AddAttribute(attributeRequest.TaxonomyID, attribute)
	checkErr(err)
	output, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}
func addDimensionHandler(w http.ResponseWriter, r *http.Request) {
	var dimensionRequest model.DimensionRequest
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&dimensionRequest)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	result, err := taxonomyBuilderDriver.AddDimension(dimensionRequest.TaxonomyID, dimensionRequest.Text)
	checkErr(err)
	output, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}
func changeDimensionHandler(w http.ResponseWriter, r *http.Request) {
	var changeDimensionRequest model.ChangeDimensionRequest
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&changeDimensionRequest)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	result, err := taxonomyBuilderDriver.ChangeDimension(changeDimensionRequest.TaxonomyID, changeDimensionRequest.Attribute, changeDimensionRequest.Dimension)
	checkErr(err)
	output, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}
func deleteCitationHandler(w http.ResponseWriter, r *http.Request) {
	var citationRequest model.CitationRequest
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&citationRequest)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	citation := model.Paper{Citation: citationRequest.Citation}
	result, err := taxonomyBuilderDriver.DeleteCitation(citationRequest.TaxonomyID, citation)
	checkErr(err)
	output, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}
func removeAttributeHandler(w http.ResponseWriter, r *http.Request) {
	var attributeRequest model.AttributeRequest
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&attributeRequest)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	attribute := model.Attribute{Text: attributeRequest.Text}
	result, err := taxonomyBuilderDriver.RemoveAttribute(attributeRequest.TaxonomyID, attribute)
	checkErr(err)
	output, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}
func removeDimensionHandler(w http.ResponseWriter, r *http.Request) {
	var dimensionRequest model.AttributeRequest
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&dimensionRequest)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	dimension := model.Dimension{Text: dimensionRequest.Text}
	result, err := taxonomyBuilderDriver.RemoveDimension(dimensionRequest.TaxonomyID, dimension)
	checkErr(err)
	output, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}
func renameAttributeHandler(w http.ResponseWriter, r *http.Request) {
	var renameAttributeRequest model.RenameAttributeRequest
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&renameAttributeRequest)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	result, err := taxonomyBuilderDriver.RenameAttribute(renameAttributeRequest.TaxonomyID, renameAttributeRequest.PreviousName, renameAttributeRequest.NewName)
	checkErr(err)
	output, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}
func updateSynonymsHandler(w http.ResponseWriter, r *http.Request) {
	var updateSynonymsRequest model.UpdateSynonymsRequest
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&updateSynonymsRequest)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	result, err := taxonomyBuilderDriver.UpdateSynonyms(updateSynonymsRequest.TaxonomyID, updateSynonymsRequest.Attribute, updateSynonymsRequest.Synonyms)
	checkErr(err)
	output, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}
func renameDimensionHandler(w http.ResponseWriter, r *http.Request) {
	var renameDimensionRequest model.RenameAttributeRequest
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&renameDimensionRequest)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	result, err := taxonomyBuilderDriver.RenameDimension(renameDimensionRequest.TaxonomyID, renameDimensionRequest.PreviousName, renameDimensionRequest.NewName)
	checkErr(err)
	output, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}
func addTaxonomyRelationHandler(w http.ResponseWriter, r *http.Request) {
	var taxonomyRelationRequest model.AttributeRelationsRequest
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&taxonomyRelationRequest)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	relation := model.AttributeRelation{TaxonomyID: taxonomyRelationRequest.TaxonomyID, AttributeSrc: taxonomyRelationRequest.AttributeSrc, AttributeDest: taxonomyRelationRequest.AttributeDest, Dimension: taxonomyRelationRequest.Dimension, Relation: taxonomyRelationRequest.Text}
	result, err := taxonomyBuilderDriver.AddTaxonomyRelation(relation)
	checkErr(err)
	output, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}
func removeTaxonomyRelationHandler(w http.ResponseWriter, r *http.Request) {
	var taxonomyRelationRequest model.AttributeRelationsRequest
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&taxonomyRelationRequest)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	relation := model.AttributeRelation{TaxonomyID: taxonomyRelationRequest.TaxonomyID, AttributeSrc: taxonomyRelationRequest.AttributeSrc, AttributeDest: taxonomyRelationRequest.AttributeDest, Dimension: taxonomyRelationRequest.Dimension, Relation: taxonomyRelationRequest.Text}
	result, err := taxonomyBuilderDriver.RemoveTaxonomyRelation(relation)
	checkErr(err)
	output, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}
func updateTaxonomyRelationTypeHandler(w http.ResponseWriter, r *http.Request) {
	var taxonomyRelationRequest model.AttributeRelationsRequest
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&taxonomyRelationRequest)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	relation := model.AttributeRelation{TaxonomyID: taxonomyRelationRequest.TaxonomyID, AttributeSrc: taxonomyRelationRequest.AttributeSrc, AttributeDest: taxonomyRelationRequest.AttributeDest, Dimension: taxonomyRelationRequest.Dimension, Relation: taxonomyRelationRequest.Text}
	result, err := taxonomyBuilderDriver.UpdateTaxonomyRelationType(relation)
	checkErr(err)
	output, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}
func updateTaxonomyRelationAnnotationHandler(w http.ResponseWriter, r *http.Request) {
	var taxonomyRelationRequest model.AttributeRelationsRequest
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&taxonomyRelationRequest)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	relation := model.AttributeRelation{TaxonomyID: taxonomyRelationRequest.TaxonomyID, AttributeSrc: taxonomyRelationRequest.AttributeSrc, AttributeDest: taxonomyRelationRequest.AttributeDest, Dimension: taxonomyRelationRequest.Dimension, Annotation: taxonomyRelationRequest.Text}
	result, err := taxonomyBuilderDriver.UpdateTaxonomyRelationAnnotation(relation)
	checkErr(err)
	output, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}
func getAttributesHandler(u *url.URL, h http.Header, attributesRequest *model.TaxonomyRequest) (int, http.Header, *MyResponse, error) {
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	attributes, err := classificationDriver.GetAllAttributes(
		attributesRequest.TaxonomyID)
	checkErr(err)
	return http.StatusOK, nil,
		&MyResponse{"0", len(attributes), attributes}, nil
}
func getLeafAttributesHandler(u *url.URL, h http.Header, attributesRequest *model.TaxonomyRequest) (int, http.Header, *MyResponse, error) {
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	attributes, err := classificationDriver.GetLeafAttributes(
		attributesRequest.TaxonomyID)
	checkErr(err)
	return http.StatusOK, nil,
		&MyResponse{"0", len(attributes), attributes}, nil
}
func getDimensionsHandler(u *url.URL, h http.Header, dimensionRequest *model.TaxonomyRequest) (int, http.Header, *MyResponse, error) {
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	dimensions, err := classificationDriver.GetAllDimensions(
		dimensionRequest.TaxonomyID)
	checkErr(err)
	return http.StatusOK, nil,
		&MyResponse{"0", len(dimensions), dimensions}, nil
}
func getCitationsHandler(u *url.URL, h http.Header, taxonomyRequest *model.TaxonomyRequest) (int, http.Header, *MyResponse, error) {
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	citations, err := classificationDriver.GetAllCitations(taxonomyRequest.TaxonomyID)
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{"0", len(citations), citations}, nil
}
func getCitationCountHandler(u *url.URL, h http.Header, taxonomyRequest *model.TaxonomyRequest) (int, http.Header, *MyResponse, error) {
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	citationCount, err := classificationDriver.GetCitationCount(taxonomyRequest.TaxonomyID)
	checkErr(err)
	return http.StatusOK, nil,
		&MyResponse{"0", len(citationCount), citationCount}, nil
}
func getCitationCountsHandler(u *url.URL, h http.Header, taxonomyRequest *model.TaxonomyRequest) (int, http.Header, *MyResponse, error) {
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	citationCounts, err := classificationDriver.GetCitationCounts(
		taxonomyRequest.TaxonomyID)
	checkErr(err)
	return http.StatusOK, nil,
		&MyResponse{"0", len(citationCounts), citationCounts}, nil
}
func getUpdateCitationReferenceCountsHandler(w http.ResponseWriter, r *http.Request) {
	var updateReferenceCountsRequest model.UpdateReferenceCountsRequest
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&updateReferenceCountsRequest)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	result, err := taxonomyBuilderDriver.UpdateCitationReferenceCounts(updateReferenceCountsRequest.TaxonomyID, updateReferenceCountsRequest.ReferenceCounts)
	checkErr(err)
	output, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}
func getCitationCountsIncludingChildrenHandler(u *url.URL, h http.Header, taxonomyRequest *model.TaxonomyRequest) (int, http.Header, *MyResponse, error) {
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	citationCounts, err := classificationDriver.GetCitationCountsIncludingChildren(
		taxonomyRequest.TaxonomyID)
	checkErr(err)
	return http.StatusOK, nil,
		&MyResponse{"0", len(citationCounts), citationCounts}, nil
}
func getUpdateMajorHandler(w http.ResponseWriter, r *http.Request) {
	var updateMajorRequest model.UpdateMajorRequest
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&updateMajorRequest)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	result, err := taxonomyBuilderDriver.UpdateMajor(updateMajorRequest.TaxonomyID, updateMajorRequest.Text, updateMajorRequest.Major)
	checkErr(err)
	output, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}
func getUpdateCitationMappingHandler(w http.ResponseWriter, r *http.Request) {
	var updateCitationMappingRequest model.UpdateCitationMappingRequest
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&updateCitationMappingRequest)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	result, err := taxonomyBuilderDriver.UpdateCitationMapping(updateCitationMappingRequest.TaxonomyID, updateCitationMappingRequest.Attribute, updateCitationMappingRequest.Citations)
	checkErr(err)
	output, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}
func getUpdateCitationMappingsHandler(w http.ResponseWriter, r *http.Request) {
	var updateCitationMappingsRequest model.UpdateCitationMappingsRequest
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&updateCitationMappingsRequest)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	result, err := taxonomyBuilderDriver.UpdateCitationMappings(updateCitationMappingsRequest.TaxonomyID, updateCitationMappingsRequest.Mappings)
	checkErr(err)
	output, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}
func getRelationTypesHandler(u *url.URL, h http.Header, r *MyRequest) (int, http.Header, *MyResponse, error) {
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	relationTypes, err := classificationDriver.GetRelationTypes()
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{"0", len(relationTypes), relationTypes}, nil
}
func getConceptCorrelationsHandler(u *url.URL, h http.Header, taxonomyRequest *model.TaxonomyRequest) (int, http.Header, *MyResponse, error) {
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	conceptRelations, err := classificationDriver.GetConceptRelations(
		taxonomyRequest.TaxonomyID)
	checkErr(err)
	return http.StatusOK, nil,
		&MyResponse{"0", len(conceptRelations), conceptRelations}, nil
}
func getConceptCorrelations3DHandler(u *url.URL, h http.Header, taxonomyRequest *model.TaxonomyRequest) (int, http.Header, *MyResponse, error) {
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	conceptRelations, err := classificationDriver.GetConceptRelations3D(
		taxonomyRequest.TaxonomyID)
	checkErr(err)
	return http.StatusOK, nil,
		&MyResponse{"0", len(conceptRelations), conceptRelations}, nil
}
func getConceptCorrelationsWithReferenceCountsHandler(u *url.URL, h http.Header, taxonomyRequest *model.TaxonomyRequest) (int, http.Header, *MyResponse, error) {
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	conceptRelations, err := classificationDriver.GetConceptRelationsWithReferenceCounts(
		taxonomyRequest.TaxonomyID)
	checkErr(err)
	return http.StatusOK, nil,
		&MyResponse{"0", len(conceptRelations), conceptRelations}, nil
}
func getConceptCorrelationsWithReferenceCounts3DHandler(u *url.URL, h http.Header, taxonomyRequest *model.TaxonomyRequest) (int, http.Header, *MyResponse, error) {
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	conceptRelations, err := classificationDriver.GetConceptRelationsWithReferenceCounts3D(
		taxonomyRequest.TaxonomyID)
	checkErr(err)
	return http.StatusOK, nil,
		&MyResponse{"0", len(conceptRelations), conceptRelations}, nil
}
func getAllConceptCorrelationsHandler(u *url.URL, h http.Header, taxonomyRequest *model.TaxonomyRequest) (int, http.Header, *MyResponse, error) {
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	conceptRelations, err := classificationDriver.GetAllConceptRelations(
		taxonomyRequest.TaxonomyID)
	checkErr(err)
	return http.StatusOK, nil,
		&MyResponse{"0", len(conceptRelations), conceptRelations}, nil
}
func getAllConceptCorrelations3DHandler(u *url.URL, h http.Header, taxonomyRequest *model.TaxonomyRequest) (int, http.Header, *MyResponse, error) {
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	conceptRelations, err := classificationDriver.GetAllConceptRelations3D(
		taxonomyRequest.TaxonomyID)
	checkErr(err)
	return http.StatusOK, nil,
		&MyResponse{"0", len(conceptRelations), conceptRelations}, nil
}
func getAllConceptCorrelationsWithReferenceCountsHandler(u *url.URL, h http.Header, taxonomyRequest *model.TaxonomyRequest) (int, http.Header, *MyResponse, error) {
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	conceptRelations, err := classificationDriver.GetAllConceptRelationsWithReferenceCounts(
		taxonomyRequest.TaxonomyID)
	checkErr(err)
	return http.StatusOK, nil,
		&MyResponse{"0", len(conceptRelations), conceptRelations}, nil
}
func getAllConceptCorrelationsWithReferenceCounts3DHandler(u *url.URL, h http.Header, taxonomyRequest *model.TaxonomyRequest) (int, http.Header, *MyResponse, error) {
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	conceptRelations, err := classificationDriver.GetAllConceptRelationsWithReferenceCounts3D(
		taxonomyRequest.TaxonomyID)
	checkErr(err)
	return http.StatusOK, nil,
		&MyResponse{"0", len(conceptRelations), conceptRelations}, nil
}
func GetParentRelationsPerAttributeHandler(u *url.URL, h http.Header, attributeRelationsRequest *model.AttributeRelationsRequest) (int, http.Header, *MyResponse, error) {
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	relations, err := classificationDriver.GetParentRelationsPerAttribute(attributeRelationsRequest.TaxonomyID, attributeRelationsRequest.Text, attributeRelationsRequest.Dimension)
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{"0", len(relations), relations}, nil
}
func GetChildRelationsPerAttributeHandler(u *url.URL, h http.Header, attributeRelationsRequest *model.AttributeRelationsRequest) (int, http.Header, *MyResponse, error) {
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	relations, err := classificationDriver.GetChildRelationsPerAttribute(attributeRelationsRequest.TaxonomyID, attributeRelationsRequest.Text, attributeRelationsRequest.Dimension)
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{"0", len(relations), relations}, nil
}
func getSharedPapersHandler(u *url.URL, h http.Header, sharedPapersRequest *model.SharedPapersRequest) (int, http.Header, *MyResponse, error) {
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	papers, err := classificationDriver.GetSharedPapers(sharedPapersRequest.TaxonomyID, sharedPapersRequest.Text1, sharedPapersRequest.Text2)
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{"0", len(papers), papers}, nil
}
func getSharedPapers3DHandler(u *url.URL, h http.Header, sharedPapersRequest *model.SharedPapersRequest) (int, http.Header, *MyResponse, error) {
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	papers, err := classificationDriver.GetSharedPapers3D(sharedPapersRequest.TaxonomyID, sharedPapersRequest.Text1, sharedPapersRequest.Text2, sharedPapersRequest.Text3)
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{"0", len(papers), papers}, nil
}
func getSharedPapersIncludingChildrenHandler(u *url.URL, h http.Header, sharedPapersRequest *model.SharedPapersRequest) (int, http.Header, *MyResponse, error) {
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	papers, err := classificationDriver.GetSharedPapersIncludingChildren(sharedPapersRequest.TaxonomyID, sharedPapersRequest.Text1, sharedPapersRequest.Text2)
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{"0", len(papers), papers}, nil
}
func getSharedPapersIncludingChildren3DHandler(u *url.URL, h http.Header, sharedPapersRequest *model.SharedPapersRequest) (int, http.Header, *MyResponse, error) {
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	papers, err := classificationDriver.GetSharedPapersIncludingChildren3D(sharedPapersRequest.TaxonomyID, sharedPapersRequest.Text1, sharedPapersRequest.Text2, sharedPapersRequest.Text3)
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{"0", len(papers), papers}, nil
}
func getAttributeDetailsHandler(u *url.URL, h http.Header, attributeDetailsRequest *model.AttributeDetailsRequest) (int, http.Header, *MyResponse, error) {
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	attributeDetails, err := classificationDriver.GetAttributeDetails(attributeDetailsRequest.TaxonomyID, attributeDetailsRequest.Text)
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{"0", len(attributeDetails), attributeDetails}, nil
}
func getCitationDetailsHandler(u *url.URL, h http.Header, citationDetailsRequest *model.SharedPapersRequest) (int, http.Header, *MyResponse, error) {
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	citationDetails, err := classificationDriver.GetCitationDetails(citationDetailsRequest.TaxonomyID, citationDetailsRequest.Text1, citationDetailsRequest.Text2)
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{"0", len(citationDetails), citationDetails}, nil
}
func getMergeAttributesHandler(w http.ResponseWriter, r *http.Request) {
	var mergeAttributesRequest model.MergeAttributesRequest
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&mergeAttributesRequest)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	attribute1 := model.Attribute{Text: mergeAttributesRequest.Text1, Dimension: mergeAttributesRequest.Dimension1}
	attribute2 := model.Attribute{Text: mergeAttributesRequest.Text2, Dimension: mergeAttributesRequest.Dimension2}
	result, err := taxonomyBuilderDriver.MergeAttributes(mergeAttributesRequest.TaxonomyID, attribute1, attribute2)
	checkErr(err)
	output, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}
func getForkAttributeHandler(w http.ResponseWriter, r *http.Request) {
	var forkAttributeRequest model.ForkAttributeRequest
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&forkAttributeRequest)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	result, err := taxonomyBuilderDriver.ForkAttribute(forkAttributeRequest.TaxonomyID, forkAttributeRequest.Text, forkAttributeRequest.Dimension, forkAttributeRequest.Parents1, forkAttributeRequest.Parents2, forkAttributeRequest.Children1, forkAttributeRequest.Children2, forkAttributeRequest.Citations1, forkAttributeRequest.Citations2)
	checkErr(err)
	output, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}
func getAttributesByNameHandler(u *url.URL, h http.Header, attributesByNameRequest *model.AttributesByNameRequest) (int, http.Header, *MyResponse, error) {
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	attributes, err := classificationDriver.GetAttributesByName(
		attributesByNameRequest.TaxonomyID, attributesByNameRequest.Texts)
	checkErr(err)
	return http.StatusOK, nil,
		&MyResponse{"0", len(attributes), attributes}, nil
}
func getAttributeCoverageHandler(u *url.URL, h http.Header, taxonomyRequest *model.TaxonomyRequest) (int, http.Header, *MyResponse, error) {
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	attributeCoverage, err := classificationDriver.GetAttributeCoverage(
		taxonomyRequest.TaxonomyID)
	checkErr(err)
	return http.StatusOK, nil,
		&MyResponse{"0", len(attributeCoverage), attributeCoverage}, nil
}
func getAttributeCoverageWithOccurrenceCountsHandler(u *url.URL, h http.Header, taxonomyRequest *model.TaxonomyRequest) (int, http.Header, *MyResponse, error) {
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	attributeCoverage, err := classificationDriver.GetAttributeCoverageWithOcurrenceCounts(
		taxonomyRequest.TaxonomyID)
	checkErr(err)
	return http.StatusOK, nil,
		&MyResponse{"0", len(attributeCoverage), attributeCoverage}, nil
}
func getAttributeCoverageWithReferenceCountsHandler(u *url.URL, h http.Header, taxonomyRequest *model.TaxonomyRequest) (int, http.Header, *MyResponse, error) {
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	attributeCoverage, err := classificationDriver.GetAttributeCoverageWithReferenceCounts(
		taxonomyRequest.TaxonomyID)
	checkErr(err)
	return http.StatusOK, nil,
		&MyResponse{"0", len(attributeCoverage), attributeCoverage}, nil
}
func getKMeansHandler(u *url.URL, h http.Header, kMeansRequest *model.KMeansRequest) (int, http.Header, *MyResponse, error) {
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	clusters, err := classificationDriver.KMeans(
		kMeansRequest.TaxonomyID, kMeansRequest.N)
	checkErr(err)
	return http.StatusOK, nil,
		&MyResponse{"0", len(clusters), clusters}, nil
}

func loadPage(filename string) (body []byte) {
	// fmt.Println("handling loadpage")
	body, err := ioutil.ReadFile(filename)
	// fmt.Printf("%d", len(body))
	checkErr(err)
	return body
}

// paper-review

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
