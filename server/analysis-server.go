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
	listen        = flag.String("listen", "127.0.0.1:8006", "listen address")
)

var sessionManager *scs.Manager
var classificationDriver data.ClassificationDriver

func main() {
	flag.Parse()

	classificationDriver = data.InitClassificationDriver(*mysqlUser, *mysqlPassword, *mysqlServer)
	dbRef, err := classificationDriver.OpenDB()
	if err == nil {
		sessionManager = scs.NewManager(mysqlstore.New(dbRef, 600000000000))
		sessionManager.Lifetime(time.Hour * 24) // session data expires after 24 hours
		sessionManager.Persist(true)            // session data persists after the browser has been closed by the user
		//sessionManager.Secure(true)
	}

	cors := tigertonic.NewCORSBuilder().AddAllowedOrigins(*listen)

	mux := tigertonic.NewTrieServeMux()

	mux.Handle("POST", "/analysis/correlation", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getCorrelationHandler), "getCorrelationHandler", nil)))
	mux.Handle("POST", "/analysis/attributesPerDimension", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getAttributesPerDimensionHandler), "getAttributesPerDimensionHandler", nil)))
	mux.Handle("POST", "/analysis/leafAttributesPerDimension", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getLeafAttributesPerDimensionHandler), "getLeafAttributesPerDimensionHandler", nil)))
	mux.Handle("POST", "/analysis/attributeClusterPerDimension", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getAttributeClusterPerDimensionHandler), "getAttributeClusterPerDimensionHandler", nil)))
	mux.Handle("POST", "/analysis/allChildrenAttributes", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getAllChildrenAttributesHandler), "getAllChildrenAttributesHandler", nil)))
	mux.Handle("POST", "/analysis/allChildrenLeafAttributes", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getAllChildrenLeafAttributesHandler), "getAllChildrenLeafAttributesHandler", nil)))
	mux.Handle("POST", "/analysis/attributeRelations", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getAttributeRelationsHandler), "getAttributeRelationsHandler", nil)))
	mux.Handle("POST", "/analysis/interdimensionalRelations", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getInterdimensionalRelationsHandler), "getInterdimensionalRelationsHandler", nil)))
	mux.Handle("POST", "/analysis/intermediateAttributes", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getIntermediateAttributesHandler), "getIntermediateAttributesHandler", nil)))
	mux.Handle("POST", "/analysis/majorAttributes", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getMajorAttributesHandler), "getMajorAttributesHandler", nil)))
	mux.Handle("POST", "/analysis/citationsPerAttribute", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getCitationsPerAttributeHandler), "getCitationsPerAttributeHandler", nil)))
	mux.Handle("POST", "/analysis/citationsPerAttributeIncludingChildren", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getCitationsPerAttributeIncludingChildrenHandler), "getCitationsPerAttributeIncludingChildrenHandler", nil)))
	mux.Handle("POST", "/analysis/attribute", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getAttributesHandler), "getAttributesHandler", nil)))
	mux.Handle("POST", "/analysis/leafAttribute", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getLeafAttributesHandler), "getLeafAttributesHandler", nil)))
	mux.Handle("POST", "/analysis/dimension", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getDimensionsHandler), "getDimensionsHandler", nil)))
	mux.Handle("POST", "/analysis/citation", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getCitationsHandler), "getCitationsHandler", nil)))
	mux.Handle("POST", "/analysis/citationCount", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getCitationCountHandler), "getCitationCountHandler", nil)))
	mux.Handle("POST", "/analysis/citationCounts", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getCitationCountsHandler), "getCitationCountsHandler", nil)))
	mux.Handle("POST", "/analysis/citationCountsIncludingChildren", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getCitationCountsIncludingChildrenHandler), "getCitationCountsIncludingChildrenHandler", nil)))
	mux.Handle("GET", "/analysis/relationTypes", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getRelationTypesHandler), "getRelationTypesHandler", nil)))
	mux.Handle("POST", "/analysis/conceptCorrelation", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getConceptCorrelationsHandler), "getConceptCorrelationsHandler", nil)))
	mux.Handle("POST", "/analysis/conceptCorrelation3D", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getConceptCorrelations3DHandler), "getConceptCorrelations3DHandler", nil)))
	mux.Handle("POST", "/analysis/conceptCorrelationWithReferenceCounts", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getConceptCorrelationsWithReferenceCountsHandler), "getConceptCorrelationsWithReferenceCountsHandler", nil)))
	mux.Handle("POST", "/analysis/conceptCorrelationWithReferenceCounts3D", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getConceptCorrelationsWithReferenceCounts3DHandler), "getConceptCorrelationsWithReferenceCounts3DHandler", nil)))
	mux.Handle("POST", "/analysis/allConceptCorrelations", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getAllConceptCorrelationsHandler), "getAllConceptCorrelationsHandler", nil)))
	mux.Handle("POST", "/analysis/allConceptCorrelations3D", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getAllConceptCorrelations3DHandler), "getAllConceptCorrelations3DHandler", nil)))
	mux.Handle("POST", "/analysis/allConceptCorrelationsWithReferenceCounts", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getAllConceptCorrelationsWithReferenceCountsHandler), "getAllConceptCorrelationsWithReferenceCountsHandler", nil)))
	mux.Handle("POST", "/analysis/allConceptCorrelationsWithReferenceCounts3D", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getAllConceptCorrelationsWithReferenceCounts3DHandler), "getAllConceptCorrelationsWithReferenceCounts3DHandler", nil)))
	mux.Handle("POST", "/analysis/parentRelationsPerAttribute", cors.Build(tigertonic.Timed(tigertonic.Marshaled(GetParentRelationsPerAttributeHandler), "GetParentRelationsPerAttributeHandler", nil)))
	mux.Handle("POST", "/analysis/childRelationsPerAttribute", cors.Build(tigertonic.Timed(tigertonic.Marshaled(GetChildRelationsPerAttributeHandler), "GetChildRelationsPerAttributeHandler", nil)))
	mux.Handle("POST", "/analysis/sharedPapers", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getSharedPapersHandler), "getSharedPapersHandler", nil)))
	mux.Handle("POST", "/analysis/sharedPapers3D", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getSharedPapers3DHandler), "getSharedPapers3DHandler", nil)))
	mux.Handle("POST", "/analysis/sharedPapersIncludingChildren", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getSharedPapersIncludingChildrenHandler), "getSharedPapersIncludingChildrenHandler", nil)))
	mux.Handle("POST", "/analysis/sharedPapersIncludingChildren3D", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getSharedPapersIncludingChildren3DHandler), "getSharedPapersIncludingChildren3DHandler", nil)))
	mux.Handle("POST", "/analysis/attributeDetails", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getAttributeDetailsHandler), "getAttributeDetailsHandler", nil)))
	mux.Handle("POST", "/analysis/citationDetails", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getCitationDetailsHandler), "getCitationDetailsHandler", nil)))
	mux.Handle("POST", "/analysis/attributeCoverage", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getAttributeCoverageHandler), "getAttributeCoverageHandler", nil)))
	mux.Handle("POST", "/analysis/attributeCoverageWithOccurrenceCounts", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getAttributeCoverageWithOccurrenceCountsHandler), "getAttributeCoverageWithOccurrenceCountsHandler", nil)))
	mux.Handle("POST", "/analysis/attributeCoverageWithReferenceCounts", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getAttributeCoverageWithReferenceCountsHandler), "getAttributeCoverageWithReferenceCountsHandler", nil)))
	mux.Handle("POST", "/analysis/attributesByName", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getAttributesByNameHandler), "getAttributesByNameHandler", nil)))
	mux.Handle("POST", "/analysis/kMeans", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getKMeansHandler), "getKMeansHandler", nil)))

	// mux.HandleFunc("GET", "/analysis/loadData.js", func(w http.ResponseWriter, r *http.Request) {
	// 	p := loadPage("../frontend/src/js/loadData.js")
	// 	fmt.Fprintf(w, "%s", p)
	// })
	//
	// mux.HandleFunc("GET", "/analysis/scopus", func(w http.ResponseWriter, r *http.Request) {
	// 	p := loadPage("../frontend/taxonomy/scopus/scopusAPI.html")
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
func getCitationCountsIncludingChildrenHandler(u *url.URL, h http.Header, taxonomyRequest *model.TaxonomyRequest) (int, http.Header, *MyResponse, error) {
	//driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	citationCounts, err := classificationDriver.GetCitationCountsIncludingChildren(
		taxonomyRequest.TaxonomyID)
	checkErr(err)
	return http.StatusOK, nil,
		&MyResponse{"0", len(citationCounts), citationCounts}, nil
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
