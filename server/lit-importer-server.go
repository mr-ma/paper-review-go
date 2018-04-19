package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
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
	listen        = flag.String("listen", "127.0.0.1:8005", "listen address")
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

	mux := tigertonic.NewTrieServeMux()

	mux.HandleFunc("POST", "/litimporter/updateCitationMapping", func(w http.ResponseWriter, r *http.Request) {
		checkTaxonomyPermissions(w, r, getUpdateCitationMappingHandler)
	})
	mux.HandleFunc("POST", "/litimporter/updateCitationMappings", func(w http.ResponseWriter, r *http.Request) {
		checkTaxonomyPermissions(w, r, getUpdateCitationMappingsHandler)
	})

	// mux.HandleFunc("GET", "/litimporter/stringComparison.js", func(w http.ResponseWriter, r *http.Request) {
	// 	p := loadPage("../frontend/src/js/stringComparison.js")
	// 	fmt.Fprintf(w, "%s", p)
	// })
	// mux.HandleFunc("GET", "/litimporter/pdf.min.js", func(w http.ResponseWriter, r *http.Request) {
	// 	p := loadPage("../frontend/src/js/libs/pdf.min.js")
	// 	fmt.Fprintf(w, "%s", p)
	// })
	// mux.HandleFunc("GET", "/litimporter/pdf.worker.min.js", func(w http.ResponseWriter, r *http.Request) {
	// 	p := loadPage("../frontend/src/js/libs/pdf.worker.min.js")
	// 	fmt.Fprintf(w, "%s", p)
	// })
	// mux.HandleFunc("GET", "/litimporter/lodash.core.js", func(w http.ResponseWriter, r *http.Request) {
	// 	p := loadPage("../frontend/src/js/libs/lodash.core.js")
	// 	fmt.Fprintf(w, "%s", p)
	// })
	// mux.HandleFunc("GET", "/litimporter/fileUploader.js", func(w http.ResponseWriter, r *http.Request) {
	// 	p := loadPage("../frontend/src/js/fileUploader.js")
	// 	fmt.Fprintf(w, "%s", p)
	// })
	// mux.HandleFunc("GET", "/litimporter/scopus.js", func(w http.ResponseWriter, r *http.Request) {
	// 	p := loadPage("../frontend/src/js/scopus.js")
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

func loadPage(filename string) (body []byte) {
	// fmt.Println("handling loadpage")
	body, err := ioutil.ReadFile(filename)
	// fmt.Printf("%d", len(body))
	checkErr(err)
	return body
}
