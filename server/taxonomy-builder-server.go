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
	listen        = flag.String("listen", "127.0.0.1:8003", "listen address")
)

var sessionManager *scs.Manager
var taxonomyBuilderDriver data.TaxonomyBuilderDriver

func main() {
	flag.Parse()

	taxonomyBuilderDriver = data.InitTaxonomyBuilderDriver(*mysqlUser, *mysqlPassword, *mysqlServer)
	dbRef, err := taxonomyBuilderDriver.OpenDB()
	if err == nil {
		sessionManager = scs.NewManager(mysqlstore.New(dbRef, 600000000000))
		sessionManager.Lifetime(time.Hour * 24) // session data expires after 24 hours
		sessionManager.Persist(true)            // session data persists after the browser has been closed by the user
		//sessionManager.Secure(true)
	}

	cors := tigertonic.NewCORSBuilder().AddAllowedOrigins(*listen)

	mux := tigertonic.NewTrieServeMux()

	mux.Handle("GET", "/taxonomybuilder/taxonomy", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getTaxonomyHandler), "getTaxonomyHandler", nil)))
	mux.Handle("POST", "/taxonomybuilder/getTaxonomyID", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getTaxonomyIDHandler), "getTaxonomyIDHandler", nil)))

	mux.HandleFunc("POST", "/taxonomybuilder/addTaxonomy", func(w http.ResponseWriter, r *http.Request) {
		checkAdmin(w, r, getAddTaxonomyHandler)
	})
	mux.HandleFunc("POST", "/taxonomybuilder/removeTaxonomy", func(w http.ResponseWriter, r *http.Request) {
		checkAdmin(w, r, getRemoveTaxonomyHandler)
	})
	mux.HandleFunc("POST", "/taxonomybuilder/deleteCitation", func(w http.ResponseWriter, r *http.Request) {
		checkTaxonomyPermissions(w, r, deleteCitationHandler)
	})
	mux.HandleFunc("POST", "/taxonomybuilder/savePositions", func(w http.ResponseWriter, r *http.Request) {
		checkTaxonomyPermissions(w, r, savePositionsHandler)
	})
	mux.HandleFunc("POST", "/taxonomybuilder/saveMajorPositions", func(w http.ResponseWriter, r *http.Request) {
		checkTaxonomyPermissions(w, r, saveMajorPositionsHandler)
	})
	mux.HandleFunc("POST", "/taxonomybuilder/save3DPositions", func(w http.ResponseWriter, r *http.Request) {
		checkTaxonomyPermissions(w, r, save3DPositionsHandler)
	})
	mux.HandleFunc("POST", "/taxonomybuilder/saveMajor3DPositions", func(w http.ResponseWriter, r *http.Request) {
		checkTaxonomyPermissions(w, r, saveMajor3DPositionsHandler)
	})
	mux.HandleFunc("POST", "/taxonomybuilder/saveEdgeBendPoints", func(w http.ResponseWriter, r *http.Request) {
		checkTaxonomyPermissions(w, r, saveEdgeBendPointsHandler)
	})
	mux.HandleFunc("POST", "/taxonomybuilder/addAttribute", func(w http.ResponseWriter, r *http.Request) {
		checkTaxonomyPermissions(w, r, addAttributeHandler)
	})
	mux.HandleFunc("POST", "/taxonomybuilder/addDimension", func(w http.ResponseWriter, r *http.Request) {
		checkTaxonomyPermissions(w, r, addDimensionHandler)
	})
	mux.HandleFunc("POST", "/taxonomybuilder/changeDimension", func(w http.ResponseWriter, r *http.Request) {
		checkTaxonomyPermissions(w, r, changeDimensionHandler)
	})
	mux.HandleFunc("POST", "/taxonomybuilder/removeAttribute", func(w http.ResponseWriter, r *http.Request) {
		checkTaxonomyPermissions(w, r, removeAttributeHandler)
	})
	mux.HandleFunc("POST", "/taxonomybuilder/removeDimension", func(w http.ResponseWriter, r *http.Request) {
		checkTaxonomyPermissions(w, r, removeDimensionHandler)
	})
	mux.HandleFunc("POST", "/taxonomybuilder/renameAttribute", func(w http.ResponseWriter, r *http.Request) {
		checkTaxonomyPermissions(w, r, renameAttributeHandler)
	})
	mux.HandleFunc("POST", "/taxonomybuilder/updateSynonyms", func(w http.ResponseWriter, r *http.Request) {
		checkTaxonomyPermissions(w, r, updateSynonymsHandler)
	})
	mux.HandleFunc("POST", "/taxonomybuilder/renameDimension", func(w http.ResponseWriter, r *http.Request) {
		checkTaxonomyPermissions(w, r, renameDimensionHandler)
	})
	mux.HandleFunc("POST", "/taxonomybuilder/addTaxonomyRelation", func(w http.ResponseWriter, r *http.Request) {
		checkTaxonomyPermissions(w, r, addTaxonomyRelationHandler)
	})
	mux.HandleFunc("POST", "/taxonomybuilder/removeTaxonomyRelation", func(w http.ResponseWriter, r *http.Request) {
		checkTaxonomyPermissions(w, r, removeTaxonomyRelationHandler)
	})
	mux.HandleFunc("POST", "/taxonomybuilder/updateTaxonomyRelationType", func(w http.ResponseWriter, r *http.Request) {
		checkTaxonomyPermissions(w, r, updateTaxonomyRelationTypeHandler)
	})
	mux.HandleFunc("POST", "/taxonomybuilder/updateTaxonomyRelationAnnotation", func(w http.ResponseWriter, r *http.Request) {
		checkTaxonomyPermissions(w, r, updateTaxonomyRelationAnnotationHandler)
	})
	mux.HandleFunc("POST", "/taxonomybuilder/updateMajor", func(w http.ResponseWriter, r *http.Request) {
		checkTaxonomyPermissions(w, r, getUpdateMajorHandler)
	})
	mux.HandleFunc("POST", "/taxonomybuilder/updateCitationReferenceCounts", func(w http.ResponseWriter, r *http.Request) {
		checkTaxonomyPermissions(w, r, getUpdateCitationReferenceCountsHandler)
	})
	mux.HandleFunc("POST", "/taxonomybuilder/mergeAttributes", func(w http.ResponseWriter, r *http.Request) {
		checkTaxonomyPermissions(w, r, getMergeAttributesHandler)
	})
	mux.HandleFunc("POST", "/taxonomybuilder/forkAttribute", func(w http.ResponseWriter, r *http.Request) {
		checkTaxonomyPermissions(w, r, getForkAttributeHandler)
	})

	// mux.HandleFunc("GET", "/taxonomybuilder/zoomInIcon.png", func(w http.ResponseWriter, r *http.Request) {
	// 	p := loadPage("../frontend/src/icons/zoom_in_128.png")
	// 	fmt.Fprintf(w, "%s", p)
	// })
	// mux.HandleFunc("GET", "/taxonomybuilder/zoomOutIcon.png", func(w http.ResponseWriter, r *http.Request) {
	// 	p := loadPage("../frontend/src/icons/zoom_out_128.png")
	// 	fmt.Fprintf(w, "%s", p)
	// })
	// mux.HandleFunc("GET", "/taxonomybuilder/cytoscape.min.js", func(w http.ResponseWriter, r *http.Request) {
	// 	p := loadPage("../frontend/taxonomy/cytoscape/cytoscape.min.js")
	// 	fmt.Fprintf(w, "%s", p)
	// })
	// mux.HandleFunc("GET", "/taxonomybuilder/konva.min.js", func(w http.ResponseWriter, r *http.Request) {
	// 	p := loadPage("../frontend/taxonomy/cytoscape/extensions/konva.min.js")
	// 	fmt.Fprintf(w, "%s", p)
	// })
	// mux.HandleFunc("GET", "/taxonomybuilder/cytoscape-undo-redo.js", func(w http.ResponseWriter, r *http.Request) {
	// 	p := loadPage("../frontend/taxonomy/cytoscape/extensions/cytoscape-undo-redo.js")
	// 	fmt.Fprintf(w, "%s", p)
	// })
	// mux.HandleFunc("GET", "/taxonomybuilder/cytoscape-clipboard.js", func(w http.ResponseWriter, r *http.Request) {
	// 	p := loadPage("../frontend/taxonomy/cytoscape/extensions/cytoscape-clipboard.js")
	// 	fmt.Fprintf(w, "%s", p)
	// })
	// mux.HandleFunc("GET", "/taxonomybuilder/cytoscape-graphml.js", func(w http.ResponseWriter, r *http.Request) {
	// 	p := loadPage("../frontend/taxonomy/cytoscape/extensions/cytoscape-graphml.js")
	// 	fmt.Fprintf(w, "%s", p)
	// })
	// mux.HandleFunc("GET", "/taxonomybuilder/cytoscape-view-utilities.js", func(w http.ResponseWriter, r *http.Request) {
	// 	p := loadPage("../frontend/taxonomy/cytoscape/extensions/cytoscape-view-utilities.js")
	// 	fmt.Fprintf(w, "%s", p)
	// })
	// mux.HandleFunc("GET", "/taxonomybuilder/cytoscape-context-menus.js", func(w http.ResponseWriter, r *http.Request) {
	// 	p := loadPage("../frontend/taxonomy/cytoscape/extensions/cytoscape-context-menus.js")
	// 	fmt.Fprintf(w, "%s", p)
	// })
	// mux.HandleFunc("GET", "/taxonomybuilder/cytoscape-context-menus.css", func(w http.ResponseWriter, r *http.Request) {
	// 	p := loadPage("../frontend/taxonomy/cytoscape/extensions/cytoscape-context-menus.css")
	// 	w.Header().Add("Content-Type", "text/css")
	// 	fmt.Fprintf(w, "%s", p)
	// })
	// mux.HandleFunc("GET", "/taxonomybuilder/cytoscape-cxtmenu.js", func(w http.ResponseWriter, r *http.Request) {
	// 	p := loadPage("../frontend/taxonomy/cytoscape/extensions/cytoscape-cxtmenu.js")
	// 	fmt.Fprintf(w, "%s", p)
	// })
	// mux.HandleFunc("GET", "/taxonomybuilder/cytoscape-autopan-on-drag.js", func(w http.ResponseWriter, r *http.Request) {
	// 	p := loadPage("../frontend/taxonomy/cytoscape/extensions/cytoscape-autopan-on-drag.js")
	// 	fmt.Fprintf(w, "%s", p)
	// })
	// mux.HandleFunc("GET", "/taxonomybuilder/cytoscape-edge-bend-editing.js", func(w http.ResponseWriter, r *http.Request) {
	// 	p := loadPage("../frontend/taxonomy/cytoscape/extensions/cytoscape-edge-bend-editing.js")
	// 	fmt.Fprintf(w, "%s", p)
	// })
	// mux.HandleFunc("GET", "/taxonomybuilder/CytoscapeEdgeEditation.js", func(w http.ResponseWriter, r *http.Request) {
	// 	p := loadPage("../frontend/taxonomy/cytoscape/extensions/CytoscapeEdgeEditation.js")
	// 	fmt.Fprintf(w, "%s", p)
	// })
	// mux.HandleFunc("GET", "/taxonomybuilder/cytoscape-edgehandles.js", func(w http.ResponseWriter, r *http.Request) {
	// 	p := loadPage("../frontend/taxonomy/cytoscape/extensions/cytoscape-edgehandles.js")
	// 	fmt.Fprintf(w, "%s", p)
	// })
	// mux.HandleFunc("GET", "/taxonomybuilder/cytoscape-expand-collapse.js", func(w http.ResponseWriter, r *http.Request) {
	// 	p := loadPage("../frontend/taxonomy/cytoscape/extensions/cytoscape-expand-collapse.js")
	// 	fmt.Fprintf(w, "%s", p)
	// })
	// mux.HandleFunc("GET", "/taxonomybuilder/cytoscape.js-navigator.css", func(w http.ResponseWriter, r *http.Request) {
	// 	p := loadPage("../frontend/taxonomy/cytoscape/extensions/cytoscape.js-navigator.css")
	// 	w.Header().Add("Content-Type", "text/css")
	// 	fmt.Fprintf(w, "%s", p)
	// })
	// mux.HandleFunc("GET", "/taxonomybuilder/cytoscape-navigator.js", func(w http.ResponseWriter, r *http.Request) {
	// 	p := loadPage("../frontend/taxonomy/cytoscape/extensions/cytoscape-navigator.js")
	// 	fmt.Fprintf(w, "%s", p)
	// })
	// mux.HandleFunc("GET", "/taxonomybuilder/cytoscape-node-resize.js", func(w http.ResponseWriter, r *http.Request) {
	// 	p := loadPage("../frontend/taxonomy/cytoscape/extensions/cytoscape-node-resize.js")
	// 	fmt.Fprintf(w, "%s", p)
	// })
	// mux.HandleFunc("GET", "/taxonomybuilder/cytoscape-noderesize.js", func(w http.ResponseWriter, r *http.Request) {
	// 	p := loadPage("../frontend/taxonomy/cytoscape/extensions/cytoscape-noderesize.js")
	// 	fmt.Fprintf(w, "%s", p)
	// })
	// mux.HandleFunc("GET", "/taxonomybuilder/cytoscape-panzoom.js", func(w http.ResponseWriter, r *http.Request) {
	// 	p := loadPage("../frontend/taxonomy/cytoscape/extensions/cytoscape-panzoom.js")
	// 	fmt.Fprintf(w, "%s", p)
	// })
	// mux.HandleFunc("GET", "/taxonomybuilder/cytoscape.js-panzoom.css", func(w http.ResponseWriter, r *http.Request) {
	// 	p := loadPage("../frontend/taxonomy/cytoscape/extensions/cytoscape.js-panzoom.css")
	// 	w.Header().Add("Content-Type", "text/css")
	// 	fmt.Fprintf(w, "%s", p)
	// })
	// mux.HandleFunc("GET", "/taxonomybuilder/jquery.qtip.min.js", func(w http.ResponseWriter, r *http.Request) {
	// 	p := loadPage("../frontend/taxonomy/cytoscape/extensions/jquery.qtip.min.js")
	// 	fmt.Fprintf(w, "%s", p)
	// })
	// mux.HandleFunc("GET", "/taxonomybuilder/jquery.qtip.min.css", func(w http.ResponseWriter, r *http.Request) {
	// 	p := loadPage("../frontend/taxonomy/cytoscape/extensions/jquery.qtip.min.css")
	// 	w.Header().Add("Content-Type", "text/css")
	// 	fmt.Fprintf(w, "%s", p)
	// })
	// mux.HandleFunc("GET", "/taxonomybuilder/cytoscape-qtip.js", func(w http.ResponseWriter, r *http.Request) {
	// 	p := loadPage("../frontend/taxonomy/cytoscape/extensions/cytoscape-qtip.js")
	// 	fmt.Fprintf(w, "%s", p)
	// })
	// mux.HandleFunc("GET", "/taxonomybuilder/cytoscape-snap-to-grid.js", func(w http.ResponseWriter, r *http.Request) {
	// 	p := loadPage("../frontend/taxonomy/cytoscape/extensions/cytoscape-snap-to-grid.js")
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

func loadPage(filename string) (body []byte) {
	// fmt.Println("handling loadpage")
	body, err := ioutil.ReadFile(filename)
	// fmt.Printf("%d", len(body))
	checkErr(err)
	return body
}
