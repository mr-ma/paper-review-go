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

func main() {
	sessionManager.Lifetime(time.Hour * 24) // session data expires after 24 hours
	sessionManager.Persist(true)            // session data persists after the browser has been closed by the user
	//sessionManager.Secure(true)
	flag.Parse()

	cors := tigertonic.NewCORSBuilder().AddAllowedOrigins(*listen)

	mux := tigertonic.NewTrieServeMux()

	mux.HandleFunc("GET", "/visualization/error.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/error.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/visualization/tables.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/tables.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/visualization/exportHTML.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/exportHTML.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/visualization/scripts.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/scripts.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/visualization/style.css", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/css/style.css")
		w.Header().Add("Content-Type", "text/css")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/visualization/main.css", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/css/main.css")
		w.Header().Add("Content-Type", "text/css")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/visualization/editableTable.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/editableTable.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/visualization/bluebird.min.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/libs/bluebird.min.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/visualization/FileSaver.min.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/libs/FileSaver.min.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/visualization/multiselect.min.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/libs/multiselect.min.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/visualization/jquery.min.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/libs/jquery.min.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/visualization/d3.min.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/libs/d3.min.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/visualization/d3.layout.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/libs/d3.layout.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/visualization/d3-hierarchy.v1.min.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/libs/d3-hierarchy.v1.min.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/visualization/d3-context-menu.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/libs/d3-context-menu.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/visualization/d3-context-menu.css", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/css/libs/d3-context-menu.css")
		w.Header().Add("Content-Type", "text/css")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/visualization/bootstrap.min.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/libs/bootstrap.min.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/visualization/bootstrap.min.css", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/css/libs/bootstrap.min.css")
		w.Header().Add("Content-Type", "text/css")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/visualization/bootstrap-waitingfor.min.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/libs/bootstrap-waitingfor.min.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/visualization/bootstrap-table.css", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/css/libs/bootstrap-table.css")
		w.Header().Add("Content-Type", "text/css")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/visualization/bootstrap-table.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/libs/bootstrap-table.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/visualization/selectize.min.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/libs/selectize.min.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/visualization/selectize.css", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/css/libs/selectize.css")
		w.Header().Add("Content-Type", "text/css")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/visualization/selectize.bootstrap3.css", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/css/libs/selectize.bootstrap3.css")
		w.Header().Add("Content-Type", "text/css")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/visualization/fonts/glyphicons-halflings-regular.woff", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/fonts/glyphicons-halflings-regular.woff")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/visualization/fonts/glyphicons-halflings-regular.woff2", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/fonts/glyphicons-halflings-regular.woff2")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/visualization/fonts/glyphicons-halflings-regular.ttf", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/fonts/glyphicons-halflings-regular.ttf")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/visualization/bootstrap-dialog.min.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/libs/bootstrap-dialog.min.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/visualization/bootstrap-dialog.min.css", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/css/libs/bootstrap-dialog.min.css")
		w.Header().Add("Content-Type", "text/css")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/visualization/d3.v4.min.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/libs/d3.v4.min.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/visualization/vis.min.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/libs/vis.min.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/visualization/vis.min.css", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/css/libs/vis.min.css")
		w.Header().Add("Content-Type", "text/css")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/visualization/d3-tip.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/libs/d3-tip.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/visualization/d3-tip.css", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/css/libs/d3-tip.css")
		w.Header().Add("Content-Type", "text/css")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/visualization/parse-bibtex.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/libs/parse-bibtex.js")
		fmt.Fprintf(w, "%s", p)
	})

	mux.HandleFunc("GET", "/", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/cytoscape/taxonomyRelations.html")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/visualization/circlePacking", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/hierarchy/circlePacking.html")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/visualization/conceptCorrelationMatrix2D", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/correlationMap/interactive/conceptCorrelations_two_dimensional.html")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/visualization/conceptCorrelationMatrix3D", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/correlationMap/interactive/conceptCorrelations_three_dimensional.html")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/visualization/taxonomyRelations", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/cytoscape/taxonomyRelations.html")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/visualization/pdf/{file}", func(w http.ResponseWriter, r *http.Request) {
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
	mux.HandleFunc("GET", "/visualization/png/{file}", func(w http.ResponseWriter, r *http.Request) {
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

func loadPage(filename string) (body []byte) {
	// fmt.Println("handling loadpage")
	body, err := ioutil.ReadFile(filename)
	// fmt.Printf("%d", len(body))
	checkErr(err)
	return body
}