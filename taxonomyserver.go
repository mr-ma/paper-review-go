package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
    "mime/multipart"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"encoding/json"
	"strconv"
	"time"

	"./data"
	"./model"
	"./go-tigertonic"
	"./scs"
	"./pdf"
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
	cert          = flag.String("cert", "", "certificate pathname")
	key           = flag.String("key", "", "private key pathname")
	config        = flag.String("config", "", "pathname of JSON configuration file")
	listen        = flag.String("listen", "127.0.0.1:8002", "listen address")
)

//var store = sessions.NewCookieStore([]byte("test-secret-4353522"))
var sessionManager = scs.NewCookieManager("u46IpCV9y5Vlur8YvODJEhgOY8m9JVE4")

func main() {
/*
	store.Options = &sessions.Options{
	    Path:     "/",
	    MaxAge:   3600 * 8, // 8 hours
	}
*/
	sessionManager.Lifetime(time.Hour*24)
	sessionManager.Persist(true) // Persist the session after a user has closed their browser.
	//sessionManager.Secure(true) // Set the Secure flag on the session cookie.
	flag.Parse()

	cors := tigertonic.NewCORSBuilder().AddAllowedOrigins(*listen) //.AddAllowedHeaders("Origin, X-Requested-With, Content-Type, Accept")

	mux := tigertonic.NewTrieServeMux()
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
	mux.Handle("GET", "/citation", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getCitationsHandler), "getCitationsHandler", nil)))
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
	mux.Handle("POST", "/kMeans", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getKMeansHandler), "getKMeansHandler", nil)))

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
        	/*
        	session, err := store.Get(r, loginRequest.Email)
		    if err != nil {
		        http.Error(w, err.Error(), http.StatusInternalServerError)
		        return
		    }
		    email := session.Values["email"]
		    admin := session.Values["admin"]
		    emailStr, ok := email.(string)
		    if !ok {
		        http.Error(w, err.Error(), http.StatusInternalServerError)
		        return
		    }
		    adminInt, ok := admin.(int)
		    if !ok {
		        http.Error(w, err.Error(), http.StatusInternalServerError)
		        return
		    }
		    */
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
		driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
		loginResult, err := driver.Login(loginRequest.Email, loginRequest.Password)
		checkErr(err)
		if loginResult.Success {
			/*
			session, _ := store.Get(r, "session")
		    session.Values["email"] = loginResult.User.Email
		    session.Values["admin"] = loginResult.User.Admin
		    err := session.Save(r, w)
		    if err != nil {
		    	fmt.Println("Error saving session.")
		    	http.Error(w, err.Error(), 500)
				return
		    }
		    */
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
		driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
		result, err := driver.SaveUser(loginRequest.Email, loginRequest.Password)
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
	mux.HandleFunc("GET", "/getUsers", func(w http.ResponseWriter, r *http.Request) {
		checkAdmin(w, r, getUsersHandler)
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
		checkAdmin(w, r, deleteCitationHandler)
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
	mux.HandleFunc("POST", "/removeAttribute", func(w http.ResponseWriter, r *http.Request) {
		checkTaxonomyPermissions(w, r, removeAttributeHandler)
	})
	mux.HandleFunc("POST", "/removeDimension", func(w http.ResponseWriter, r *http.Request) {
		checkTaxonomyPermissions(w, r, removeDimensionHandler)
	})
	mux.HandleFunc("POST", "/renameAttribute", func(w http.ResponseWriter, r *http.Request) {
		checkTaxonomyPermissions(w, r, renameAttributeHandler)
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

	mux.HandleFunc("GET", "/error.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/error.js")
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
	mux.HandleFunc("GET", "/pdf.min.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/pdf.min.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/pdf.worker.min.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/pdf.worker.min.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/lodash.core.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/lodash.core.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/compare-strings.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/compare-strings.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/fuzzysort.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/fuzzysort.js")
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
	mux.HandleFunc("GET", "/zoomInIcon.png", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/icons/zoom_in_128.png")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/zoomOutIcon.png", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/icons/zoom_out_128.png")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/bluebird.min.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/bluebird.min.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/FileSaver.min.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/FileSaver.min.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/multiselect.min.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/multiselect.min.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/three.min.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/three/three.min.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/threex.dynamictexture.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/three/threex.dynamictexture.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/threex.dynamictext2dobject.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/three/threex.dynamictext2dobject.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/DragControls.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/three/DragControls.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/TrackballControls.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/three/TrackballControls.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/PointerLockControls.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/three/PointerLockControls.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/FlyControls.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/three/FlyControls.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/stats.min.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/stats.min.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/cytoscape.min.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/cytoscape/cytoscape.min.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/jquery.min.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/jquery.min.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/d3.min.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/d3.min.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/d3.layout.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/d3.layout.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/d3-hierarchy.v1.min.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/d3-hierarchy.v1.min.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/d3-context-menu.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/d3-context-menu.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/d3-context-menu.css", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/css/d3-context-menu.css")
		w.Header().Add("Content-Type", "text/css")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/runner.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/runner.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/bootstrap.min.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/bootstrap.min.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/bootstrap.min.css", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/css/bootstrap.min.css")
		w.Header().Add("Content-Type", "text/css")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/bootstrap-waitingfor.min.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/bootstrap-waitingfor.min.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/bootstrap-table.css", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/css/bootstrap-table.css")
		w.Header().Add("Content-Type", "text/css")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/bootstrap-table.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/bootstrap-table.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/selectize.min.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/selectize.min.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/selectize.css", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/css/selectize.css")
		w.Header().Add("Content-Type", "text/css")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/loginForm.css", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/css/loginForm.css")
		w.Header().Add("Content-Type", "text/css")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/editableTable.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/editableTable.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/selectize.bootstrap3.css", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/css/selectize.bootstrap3.css")
		w.Header().Add("Content-Type", "text/css")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/style.css", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/css/style.css")
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
		p := loadPage("frontend/src/js/bootstrap-dialog.min.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/bootstrap-dialog.min.css", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/css/bootstrap-dialog.min.css")
		w.Header().Add("Content-Type", "text/css")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/d3.v2.min.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/d3.v2.min.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/d3.v4.min.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/d3.v4.min.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/viz.min.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/viz.v1.1.2.min.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/vis.min.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/vis.min.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/vis.min.css", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/css/vis.min.css")
		w.Header().Add("Content-Type", "text/css")
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
	mux.HandleFunc("GET", "/konva.min.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/cytoscape/extensions/konva.min.js")
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
	mux.HandleFunc("GET", "/d3-tip.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/d3-tip.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/d3-tip.css", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/css/d3-tip.css")
		w.Header().Add("Content-Type", "text/css")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/cytoscape-snap-to-grid.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/cytoscape/extensions/cytoscape-snap-to-grid.js")
		fmt.Fprintf(w, "%s", p)
	})
	// mux.Handle("GET","/",cors.Build(tigertonic.Timed(tigertonic.Marshaled(getIndexHandler), "getIndexHandler", nil)))
	mux.HandleFunc("GET", "/", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/index.html")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/benchmark", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/Benchmark.html")
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
	mux.HandleFunc("GET", "/treemap", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/hierarchy/treemap.html")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/treemap2", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/hierarchy/treemap2.html")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/conceptCorrelationMatrix", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/correlationMap/interactive/conceptCorrelations.html")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/conceptCorrelationMatrix2D", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/correlationMap/interactive/conceptCorrelations_two_dimensional.html")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/conceptCorrelationMatrix2DCluster", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/correlationMap/interactive/conceptCorrelations_two_dimensional_cluster.html")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/conceptCorrelationMatrix3D", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/correlationMap/interactive/conceptCorrelations_three_dimensional.html")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/conceptCorrelationMatrix3D2", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/correlationMap/interactive/conceptCorrelations_three_dimensional2.html")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/attributeCoverageMatrix", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/correlationMap/interactive/attributeCoverage.html")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/3D", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/3D/hierarchy.html")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/chord", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/chordMap.html")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/horizontal", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/BPHorizontal.html")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/vertical", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/BPVertical.html")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/correlationMap", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/correlationMap/index.html")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/coverage", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/correlationMap/attributeCoverage.html")
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
	/*
	// https://astaxie.gitbooks.io/build-web-application-with-golang/en/04.5.html
	mux.HandleFunc("POST", "/upload", func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("file")
		if err != nil {
		    fmt.Println(err)
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./files/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
		    fmt.Println(err)
		}
		defer f.Close()
		io.Copy(f, file)
	})
	*/
	mux.HandleFunc("POST", "/upload", func(w http.ResponseWriter, r *http.Request) {
/*
		file, header, err := r.FormFile("file")

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "{'error': %s}", err)
			return
		}
		defer file.Close()

		out, err := os.Create("files/uploaded-" + header.Filename)
		checkErr(err)
		if err != nil {
			fmt.Fprintf(w, "[-] Unable to create the file for writing. Check your write access privilege.", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		defer out.Close()

		// write the content from POST to the file
		_, err = io.Copy(out, file)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
*/
		UploadFile(w, r)
	})

	// c := &Config{}
	// if err := tigertonic.Configure(*config, c); nil != err {
	// 	checkErr(err)
	// }

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

func parsePDF(path string) (string, error) {
	r, err := pdf.Open(path)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	buf.ReadFrom(r.GetPlainText())
	return buf.String(), nil
}

/*
func parsePDF(path string) (string, error) {
	res, err := docconv.ConvertPath(path)
	checkErr(err)
	if err != nil {
		return "", err
	}
	return res.Body, nil
}
*/

func deleteFile(path string) (error) {
	// delete file
	var err = os.Remove(path)
	return err
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {  
      var (  
           status int  
           err  error  
      )  
      defer func() {  
           if nil != err {  
                http.Error(w, err.Error(), status)  
           }  
      }()  
      if err = r.ParseMultipartForm(32 << 20); nil != err {  
           status = http.StatusInternalServerError  
           return  
      }
      paths := []string{}
      for _, fheaders := range r.MultipartForm.File {
           for _, hdr := range fheaders {
           		path := "./files/" + hdr.Filename
                var infile multipart.File  
                if infile, err = hdr.Open(); nil != err {  
                     status = http.StatusInternalServerError  
                     return  
                }  
                var outfile *os.File  
                if outfile, err = os.Create(path); nil != err {  
                     status = http.StatusInternalServerError  
                     return  
                }   
                var written int64  
                if written, err = io.Copy(outfile, infile); nil != err {  
                     status = http.StatusInternalServerError  
                     return  
                }
                fmt.Println("written: " + strconv.Itoa(int(written)))
				text, err := parsePDF(path)
				if err == nil {
					fmt.Println("parsed file: " + strconv.Itoa(len(text)))
				} else {
					fmt.Println("parse file Error")
				}
				defer infile.Close()
				defer outfile.Close()
				paths = append(paths, path)
           }
        }
        for _, elem := range paths {
			err = deleteFile(elem)
			if err != nil {
				fmt.Println("deleting file failed.")
			}
        }
        /*
        for _, fheaders := range r.MultipartForm.File {
            for _, hdr := range fheaders {
				path := "./files/" + hdr.Filename
				err = deleteFile(path)
				if err != nil {
					fmt.Println("deleting file failed.")
				}
			}
		}
		*/
 } 

func UploadFile(w http.ResponseWriter, r *http.Request) {
/*
    if r.Method != http.MethodPost {
        http.Redirect(w, r, "/", http.StatusSeeOther)
        return
    }

    file, handle, err := r.FormFile("file")
    if err != nil {
        fmt.Fprintf(w, "%v", err)
        return
    }
    defer file.Close()

    mimeType := handle.Header.Get("Content-Type")
    switch mimeType {
    case "application/pdf":
        saveFile(w, file, handle)
    default:
        jsonResponse(w, http.StatusBadRequest, "The format file is not valid.")
    }
*/
/*
 		err := r.ParseMultipartForm(100000)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//get a ref to the parsed multipart form
		m := r.MultipartForm

		//get the *fileheaders
		files := m.File["file"]
		for i, _ := range files {
			//for each fileheader, get a handle to the actual file
			file, err := files[i].Open()
			defer file.Close()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			path := "./files/" + files[i].Filename
			//create destination file making sure the path is writeable.
			dst, err := os.Create(path)
			defer dst.Close()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			//copy the uploaded file to the destination file
			if _, err := io.Copy(dst, file); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			text, err := parsePDF(path)
			if err == nil {
				fmt.Println("parsed file: " + strconv.Itoa(len(text)))
			}
		}
*/
		UploadHandler(w, r)
		jsonResponse(w, http.StatusCreated, "File uploaded successfully!.")
/*
		result := model.Result{Success: true}
		output, err := json.Marshal(result)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.Header().Set("content-type", "application/json")
		w.Write(output)
*/
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

func checkAdmin (w http.ResponseWriter, r *http.Request, callback fn) {
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

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	users, err := driver.GetUsers()
	checkErr(err)
	userResult := []model.User{}
	if err == nil {
		userResult = users
	}
	for _, el := range userResult {
		fmt.Println("user: " + el.Email)
	}
	output, err := json.Marshal(userResult)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}

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
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	result, err := driver.CreateUser(userRequest.Email, userRequest.Password, userRequest.Admin)
	checkErr(err)
	output, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}

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
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	result, err := driver.UpdateUser(userRequest.Email, userRequest.Admin)
	checkErr(err)
	output, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}

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
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	result, err := driver.DeleteUser(userRequest.Email)
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
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	taxonomies, err := driver.GetAllTaxonomies()
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{"0", len(taxonomies), taxonomies}, nil
}

func getTaxonomyIDHandler(u *url.URL, h http.Header, taxonomyIDRequest *model.TaxonomyIDRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	taxonomyID, err := driver.GetTaxonomyID(taxonomyIDRequest.Text)
	checkErr(err)
	return http.StatusOK, nil,
		&MyResponse{"0", len(taxonomyID), taxonomyID}, nil
}

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
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	result, err := driver.GetTaxonomyPermissions(userRequest.Email)
	checkErr(err)
	output, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}

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
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	result, err := driver.UpdateTaxonomyPermissions(taxonomyPermissionsRequest.Email, taxonomyPermissionsRequest.Permissions)
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
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	result, err := driver.AddTaxonomy(addTaxonomyRequest.Taxonomy, addTaxonomyRequest.Dimension)
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
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	result, err := driver.RemoveTaxonomy(taxonomyRequest.TaxonomyID)
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
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	papers, err := driver.ExportCorrelations(
		correlationRequest.Attributes, correlationRequest.TaxonomyID)
	checkErr(err)
	return http.StatusOK, nil,
		&MyResponse{"0", len(papers), papers}, nil
}
func getAttributesPerDimensionHandler(u *url.URL, h http.Header, attributesPerDimensionRequest *model.AttributeRelationsRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	attributes, err := driver.GetAttributesPerDimension(
		attributesPerDimensionRequest.TaxonomyID, attributesPerDimensionRequest.Dimension)
	checkErr(err)
	return http.StatusOK, nil,
		&MyResponse{"0", len(attributes), attributes}, nil
}
func getLeafAttributesPerDimensionHandler(u *url.URL, h http.Header, attributesPerDimensionRequest *model.AttributeRelationsRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	attributes, err := driver.GetLeafAttributesPerDimension(
		attributesPerDimensionRequest.TaxonomyID, attributesPerDimensionRequest.Dimension)
	checkErr(err)
	return http.StatusOK, nil,
		&MyResponse{"0", len(attributes), attributes}, nil
}
func getAttributeClusterPerDimensionHandler(u *url.URL, h http.Header, attributeClusterPerDimensionRequest *model.AttributeRelationsRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	clusters, err := driver.GetAttributeClusterPerDimension(
		attributeClusterPerDimensionRequest.TaxonomyID, attributeClusterPerDimensionRequest.Dimension)
	checkErr(err)
	return http.StatusOK, nil,
		&MyResponse{"0", len(clusters), clusters}, nil
}
func getAllChildrenAttributesHandler(u *url.URL, h http.Header, allChildrenAttributesRequest *model.AllChildrenAttributesRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	attributes, err := driver.GetAllChildrenAttributes(
		allChildrenAttributesRequest.TaxonomyID, allChildrenAttributesRequest.Parent)
	checkErr(err)
	return http.StatusOK, nil,
		&MyResponse{"0", len(attributes), attributes}, nil
}
func getAllChildrenLeafAttributesHandler(u *url.URL, h http.Header, allChildrenAttributesRequest *model.AllChildrenAttributesRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	attributes, err := driver.GetAllChildrenLeafAttributes(
		allChildrenAttributesRequest.TaxonomyID, allChildrenAttributesRequest.Parent)
	checkErr(err)
	return http.StatusOK, nil,
		&MyResponse{"0", len(attributes), attributes}, nil
}
func getIntermediateAttributesHandler(u *url.URL, h http.Header, intermediateAttributesRequest *model.AttributeRelationsRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	intermediateAttributes, err := driver.GetIntermediateAttributes(
		intermediateAttributesRequest.TaxonomyID, intermediateAttributesRequest.MinValue, intermediateAttributesRequest.MaxValue)
	checkErr(err)
	return http.StatusOK, nil,
		&MyResponse{"0", len(intermediateAttributes), intermediateAttributes}, nil
}
func getMajorAttributesHandler(u *url.URL, h http.Header, majorAttributesRequest *model.TaxonomyRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	majorAttributes, err := driver.GetMajorAttributes(
		majorAttributesRequest.TaxonomyID)
	checkErr(err)
	return http.StatusOK, nil,
		&MyResponse{"0", len(majorAttributes), majorAttributes}, nil
}
func getAttributeRelationsHandler(u *url.URL, h http.Header, attributeRelationsRequest *model.AttributeRelationsRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	attributeRelations, err := driver.GetAttributeRelationsPerDimension(
		attributeRelationsRequest.TaxonomyID, attributeRelationsRequest.Dimension)
	checkErr(err)
	return http.StatusOK, nil,
		&MyResponse{"0", len(attributeRelations), attributeRelations}, nil
}
func getInterdimensionalRelationsHandler(u *url.URL, h http.Header, interdimensionalRelationsRequest *model.AttributeRelationsRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	attributeRelations, err := driver.GetInterdimensionalRelations(
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
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	result, err := driver.SavePositions(
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
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	result, err := driver.SaveMajorPositions(
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
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	result, err := driver.Save3DPositions(
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
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	result, err := driver.SaveMajor3DPositions(
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
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	result, err := driver.SaveEdgeBendPoints(
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
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	citations, err := driver.GetCitationsPerAttribute(citationsPerAttributeRequest.TaxonomyID, citationsPerAttributeRequest.Attribute)
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{"0", len(citations), citations}, nil
}
func getCitationsPerAttributeIncludingChildrenHandler(u *url.URL, h http.Header, citationsPerAttributeRequest *model.CitationsPerAttributeRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	citations, err := driver.GetCitationsPerAttributeIncludingChildren(citationsPerAttributeRequest.TaxonomyID, citationsPerAttributeRequest.Attribute)
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
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	attribute := model.Attribute{Text: attributeRequest.Text, X: attributeRequest.X, Y: attributeRequest.Y, XMajor: attributeRequest.XMajor, YMajor: attributeRequest.YMajor, Major: attributeRequest.Major, Dimension: attributeRequest.Dimension}
	result, err := driver.AddAttribute(attributeRequest.TaxonomyID, attribute)
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
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	result, err := driver.AddDimension(dimensionRequest.TaxonomyID, dimensionRequest.Text)
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
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	citation := model.Paper{Citation: citationRequest.Citation}
	result, err := driver.DeleteCitation(citation)
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
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	attribute := model.Attribute{Text: attributeRequest.Text}
	result, err := driver.RemoveAttribute(attributeRequest.TaxonomyID, attribute)
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
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	dimension := model.Dimension{Text: dimensionRequest.Text}
	result, err := driver.RemoveDimension(dimensionRequest.TaxonomyID, dimension)
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
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	result, err := driver.RenameAttribute(renameAttributeRequest.TaxonomyID, renameAttributeRequest.PreviousName, renameAttributeRequest.NewName)
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
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	result, err := driver.RenameDimension(renameDimensionRequest.TaxonomyID, renameDimensionRequest.PreviousName, renameDimensionRequest.NewName)
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
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	relation := model.AttributeRelation{TaxonomyID: taxonomyRelationRequest.TaxonomyID, AttributeSrc: taxonomyRelationRequest.AttributeSrc, AttributeDest: taxonomyRelationRequest.AttributeDest, Dimension: taxonomyRelationRequest.Dimension, Relation: taxonomyRelationRequest.Text}
	result, err := driver.AddTaxonomyRelation(relation)
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
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	relation := model.AttributeRelation{TaxonomyID: taxonomyRelationRequest.TaxonomyID, AttributeSrc: taxonomyRelationRequest.AttributeSrc, AttributeDest: taxonomyRelationRequest.AttributeDest, Dimension: taxonomyRelationRequest.Dimension, Relation: taxonomyRelationRequest.Text}
	result, err := driver.RemoveTaxonomyRelation(relation)
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
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	relation := model.AttributeRelation{TaxonomyID: taxonomyRelationRequest.TaxonomyID, AttributeSrc: taxonomyRelationRequest.AttributeSrc, AttributeDest: taxonomyRelationRequest.AttributeDest, Dimension: taxonomyRelationRequest.Dimension, Relation: taxonomyRelationRequest.Text}
	result, err := driver.UpdateTaxonomyRelationType(relation)
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
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	relation := model.AttributeRelation{TaxonomyID: taxonomyRelationRequest.TaxonomyID, AttributeSrc: taxonomyRelationRequest.AttributeSrc, AttributeDest: taxonomyRelationRequest.AttributeDest, Dimension: taxonomyRelationRequest.Dimension, Annotation: taxonomyRelationRequest.Text}
	result, err := driver.UpdateTaxonomyRelationAnnotation(relation)
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
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	attributes, err := driver.GetAllAttributes(
		attributesRequest.TaxonomyID)
	checkErr(err)
	return http.StatusOK, nil,
		&MyResponse{"0", len(attributes), attributes}, nil
}
func getLeafAttributesHandler(u *url.URL, h http.Header, attributesRequest *model.TaxonomyRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	attributes, err := driver.GetLeafAttributes(
		attributesRequest.TaxonomyID)
	checkErr(err)
	return http.StatusOK, nil,
		&MyResponse{"0", len(attributes), attributes}, nil
}
func getDimensionsHandler(u *url.URL, h http.Header, dimensionRequest *model.TaxonomyRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	dimensions, err := driver.GetAllDimensions(
		dimensionRequest.TaxonomyID)
	checkErr(err)
	return http.StatusOK, nil,
		&MyResponse{"0", len(dimensions), dimensions}, nil
}
func getCitationsHandler(u *url.URL, h http.Header, r *MyRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	citations, err := driver.GetAllCitations()
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{"0", len(citations), citations}, nil
}
func getCitationCountHandler(u *url.URL, h http.Header, taxonomyRequest *model.TaxonomyRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	citationCount, err := driver.GetCitationCount()
	checkErr(err)
	return http.StatusOK, nil,
		&MyResponse{"0", len(citationCount), citationCount}, nil
}
func getCitationCountsHandler(u *url.URL, h http.Header, taxonomyRequest *model.TaxonomyRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	citationCounts, err := driver.GetCitationCounts(
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
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	result, err := driver.UpdateCitationReferenceCounts(updateReferenceCountsRequest.ReferenceCounts)
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
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	citationCounts, err := driver.GetCitationCountsIncludingChildren(
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
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	result, err := driver.UpdateMajor(updateMajorRequest.TaxonomyID, updateMajorRequest.Text, updateMajorRequest.Major)
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
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	result, err := driver.UpdateCitationMapping(updateCitationMappingRequest.TaxonomyID, updateCitationMappingRequest.Attribute, updateCitationMappingRequest.Citations)
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
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	result, err := driver.UpdateCitationMappings(updateCitationMappingsRequest.TaxonomyID, updateCitationMappingsRequest.Mappings)
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
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	relationTypes, err := driver.GetRelationTypes()
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{"0", len(relationTypes), relationTypes}, nil
}
func getConceptCorrelationsHandler(u *url.URL, h http.Header, taxonomyRequest *model.TaxonomyRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	conceptRelations, err := driver.GetConceptRelations(
		taxonomyRequest.TaxonomyID)
	checkErr(err)
	return http.StatusOK, nil,
		&MyResponse{"0", len(conceptRelations), conceptRelations}, nil
}
func getConceptCorrelations3DHandler(u *url.URL, h http.Header, taxonomyRequest *model.TaxonomyRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	conceptRelations, err := driver.GetConceptRelations3D(
		taxonomyRequest.TaxonomyID)
	checkErr(err)
	return http.StatusOK, nil,
		&MyResponse{"0", len(conceptRelations), conceptRelations}, nil
}
func getConceptCorrelationsWithReferenceCountsHandler(u *url.URL, h http.Header, taxonomyRequest *model.TaxonomyRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	conceptRelations, err := driver.GetConceptRelationsWithReferenceCounts(
		taxonomyRequest.TaxonomyID)
	checkErr(err)
	return http.StatusOK, nil,
		&MyResponse{"0", len(conceptRelations), conceptRelations}, nil
}
func getConceptCorrelationsWithReferenceCounts3DHandler(u *url.URL, h http.Header, taxonomyRequest *model.TaxonomyRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	conceptRelations, err := driver.GetConceptRelationsWithReferenceCounts3D(
		taxonomyRequest.TaxonomyID)
	checkErr(err)
	return http.StatusOK, nil,
		&MyResponse{"0", len(conceptRelations), conceptRelations}, nil
}
func getAllConceptCorrelationsHandler(u *url.URL, h http.Header, taxonomyRequest *model.TaxonomyRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	conceptRelations, err := driver.GetAllConceptRelations(
		taxonomyRequest.TaxonomyID)
	checkErr(err)
	return http.StatusOK, nil,
		&MyResponse{"0", len(conceptRelations), conceptRelations}, nil
}
func getAllConceptCorrelations3DHandler(u *url.URL, h http.Header, taxonomyRequest *model.TaxonomyRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	conceptRelations, err := driver.GetAllConceptRelations3D(
		taxonomyRequest.TaxonomyID)
	checkErr(err)
	return http.StatusOK, nil,
		&MyResponse{"0", len(conceptRelations), conceptRelations}, nil
}
func getAllConceptCorrelationsWithReferenceCountsHandler(u *url.URL, h http.Header, taxonomyRequest *model.TaxonomyRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	conceptRelations, err := driver.GetAllConceptRelationsWithReferenceCounts(
		taxonomyRequest.TaxonomyID)
	checkErr(err)
	return http.StatusOK, nil,
		&MyResponse{"0", len(conceptRelations), conceptRelations}, nil
}
func getAllConceptCorrelationsWithReferenceCounts3DHandler(u *url.URL, h http.Header, taxonomyRequest *model.TaxonomyRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	conceptRelations, err := driver.GetAllConceptRelationsWithReferenceCounts3D(
		taxonomyRequest.TaxonomyID)
	checkErr(err)
	return http.StatusOK, nil,
		&MyResponse{"0", len(conceptRelations), conceptRelations}, nil
}
func GetParentRelationsPerAttributeHandler(u *url.URL, h http.Header, attributeRelationsRequest *model.AttributeRelationsRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	relations, err := driver.GetParentRelationsPerAttribute(attributeRelationsRequest.TaxonomyID, attributeRelationsRequest.Text, attributeRelationsRequest.Dimension)
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{"0", len(relations), relations}, nil
}
func GetChildRelationsPerAttributeHandler(u *url.URL, h http.Header, attributeRelationsRequest *model.AttributeRelationsRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	relations, err := driver.GetChildRelationsPerAttribute(attributeRelationsRequest.TaxonomyID, attributeRelationsRequest.Text, attributeRelationsRequest.Dimension)
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{"0", len(relations), relations}, nil
}
func getSharedPapersHandler(u *url.URL, h http.Header, sharedPapersRequest *model.SharedPapersRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	papers, err := driver.GetSharedPapers(sharedPapersRequest.TaxonomyID, sharedPapersRequest.Text1, sharedPapersRequest.Text2)
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{"0", len(papers), papers}, nil
}
func getSharedPapers3DHandler(u *url.URL, h http.Header, sharedPapersRequest *model.SharedPapersRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	papers, err := driver.GetSharedPapers3D(sharedPapersRequest.TaxonomyID, sharedPapersRequest.Text1, sharedPapersRequest.Text2, sharedPapersRequest.Text3)
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{"0", len(papers), papers}, nil
}
func getSharedPapersIncludingChildrenHandler(u *url.URL, h http.Header, sharedPapersRequest *model.SharedPapersRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	papers, err := driver.GetSharedPapersIncludingChildren(sharedPapersRequest.TaxonomyID, sharedPapersRequest.Text1, sharedPapersRequest.Text2)
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{"0", len(papers), papers}, nil
}
func getSharedPapersIncludingChildren3DHandler(u *url.URL, h http.Header, sharedPapersRequest *model.SharedPapersRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	papers, err := driver.GetSharedPapersIncludingChildren3D(sharedPapersRequest.TaxonomyID, sharedPapersRequest.Text1, sharedPapersRequest.Text2, sharedPapersRequest.Text3)
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{"0", len(papers), papers}, nil
}
func getAttributeDetailsHandler(u *url.URL, h http.Header, attributeDetailsRequest *model.AttributeDetailsRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	attributeDetails, err := driver.GetAttributeDetails(attributeDetailsRequest.TaxonomyID, attributeDetailsRequest.Text)
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{"0", len(attributeDetails), attributeDetails}, nil
}
func getCitationDetailsHandler(u *url.URL, h http.Header, citationDetailsRequest *model.SharedPapersRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	citationDetails, err := driver.GetCitationDetails(citationDetailsRequest.Text1, citationDetailsRequest.Text2)
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
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	attribute1 := model.Attribute{Text: mergeAttributesRequest.Text1, Dimension: mergeAttributesRequest.Dimension1}
	attribute2 := model.Attribute{Text: mergeAttributesRequest.Text2, Dimension: mergeAttributesRequest.Dimension2}
	result, err := driver.MergeAttributes(mergeAttributesRequest.TaxonomyID, attribute1, attribute2)
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
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	result, err := driver.ForkAttribute(forkAttributeRequest.TaxonomyID, forkAttributeRequest.Text, forkAttributeRequest.Dimension, forkAttributeRequest.Parents1, forkAttributeRequest.Parents2, forkAttributeRequest.Children1, forkAttributeRequest.Children2, forkAttributeRequest.Citations1, forkAttributeRequest.Citations2)
	checkErr(err)
	output, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}
func getAttributeCoverageHandler(u *url.URL, h http.Header, taxonomyRequest *model.TaxonomyRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	attributeCoverage, err := driver.GetAttributeCoverage(
		taxonomyRequest.TaxonomyID)
	checkErr(err)
	return http.StatusOK, nil,
		&MyResponse{"0", len(attributeCoverage), attributeCoverage}, nil
}
func getAttributeCoverageWithOccurrenceCountsHandler(u *url.URL, h http.Header, taxonomyRequest *model.TaxonomyRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	attributeCoverage, err := driver.GetAttributeCoverageWithOcurrenceCounts(
		taxonomyRequest.TaxonomyID)
	checkErr(err)
	return http.StatusOK, nil,
		&MyResponse{"0", len(attributeCoverage), attributeCoverage}, nil
}
func getAttributeCoverageWithReferenceCountsHandler(u *url.URL, h http.Header, taxonomyRequest *model.TaxonomyRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	attributeCoverage, err := driver.GetAttributeCoverageWithReferenceCounts(
		taxonomyRequest.TaxonomyID)
	checkErr(err)
	return http.StatusOK, nil,
		&MyResponse{"0", len(attributeCoverage), attributeCoverage}, nil
}
func getKMeansHandler(u *url.URL, h http.Header, kMeansRequest *model.KMeansRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	clusters, err := driver.KMeans(
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
