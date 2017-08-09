package main

import (
	"net/http"
	"net/url"
	"fmt"
	"os"
	"syscall"
	"flag"
	"os/signal"
	"io/ioutil"
	"github.com/mr-ma/paper-review-go/data"
	"github.com/mr-ma/paper-review-go/model"
	"github.com/rcrowley/go-tigertonic"
	"path/filepath"
	"strings"
	"bytes"
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

var
(
	mysqlUser = flag.String("mysqluser", "foo", "a mysql user")
	mysqlPassword = flag.String("mysqlpass", "bar", "the mysql password")
	cert   = flag.String("cert", "", "certificate pathname")
	key    = flag.String("key", "", "private key pathname")
	config = flag.String("config", "", "pathname of JSON configuration file")
	listen = flag.String("listen", "127.0.0.1:8001", "listen address")
)
func main() {

	flag.Parse()

	cors := tigertonic.NewCORSBuilder().AddAllowedOrigins(listen)//.AddAllowedHeaders("Origin, X-Requested-With, Content-Type, Accept")

	mux := tigertonic.NewTrieServeMux()
	mux.Handle("POST", "/correlation", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getCorrelationHandler), "getCorrelationHandler", nil)))
	mux.Handle("GET", "/attribute", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getAttributesHandler), "getAttributesHandler", nil)))
	mux.Handle("GET", "/citation", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getCitationsHandler), "getCitationsHandler", nil)))
	// mux.Handle("GET","/",cors.Build(tigertonic.Timed(tigertonic.Marshaled(getIndexHandler), "getIndexHandler", nil)))
	mux.HandleFunc("GET", "/", func(w http.ResponseWriter, r *http.Request) {
    p := loadPage("frontend/taxonomy/index.html")
    fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/pdf/{file}", func(w http.ResponseWriter, r *http.Request) {
		fp := filepath.Clean(r.URL.Path)
		var p []byte
		if strings.Contains(fp,"system.pdf"){
			p = loadPage("frontend/pdfs/system.pdf")
		} else if strings.Contains(fp,"attack.pdf"){
			p = loadPage("frontend/pdfs/attack.pdf")
		} else if strings.Contains(fp,"defense.pdf"){
			p = loadPage("frontend/pdfs/defense.pdf")
		} else if strings.Contains(fp,"relations.pdf"){
			p = loadPage("frontend/pdfs/view-relations.pdf")
		}

		b := bytes.NewBuffer(p)
		// stream straight to client(browser)
		w.Header().Set("Content-type", "application/pdf")

		if _, err := b.WriteTo(w); err != nil { // <----- here!
				fmt.Fprintf(w, "%s", err)
		}
		w.Write([]byte("PDF Generated"))
	})
	mux.HandleFunc("GET", "/png/{file}", func(w http.ResponseWriter, r *http.Request) {
		fp := filepath.Clean(r.URL.Path)
		var p []byte
		if strings.Contains(fp,"system"){
			p = loadPage("frontend/pngs/system.png")
		} else if strings.Contains(fp,"attack"){
			p = loadPage("frontend/pngs/attack.png")
		} else if strings.Contains(fp,"defense"){
			p = loadPage("frontend/pngs/defense.png")
		} else if strings.Contains(fp,"relations"){
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

	// c := &Config{}
	// if err := tigertonic.Configure(*config, c); nil != err {
	// 	checkErr(err)
	// }

	server:=tigertonic.NewServer(*listen, tigertonic.Logged(mux, nil))
	go func(){
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

func getCorrelationHandler(u *url.URL, h http.Header, correlationRequest *model.CorrelationRequest) (int, http.Header, *MyResponse, error) {
	if len(correlationRequest.Attributes) == 0 {
		return http.StatusNotAcceptable, nil,
		&MyResponse{"0", 0,"I need some attributes to produce correlations"}, nil
	}
	driver := data.InitClassificationDriver(*mysqlUser,*mysqlPassword)
	papers, err := driver.ExportCorrelations(
		correlationRequest.Attributes,correlationRequest.TaxonomyID)
	checkErr(err)
	return http.StatusOK, nil,
	&MyResponse{"0",len(papers), papers}, nil
}
func getAttributesHandler(u *url.URL, h http.Header, r *MyRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser,*mysqlPassword)
	attributes, err := driver.GetAllAttributes()
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{"0",len(attributes), attributes}, nil
}
func getCitationsHandler(u *url.URL, h http.Header, r *MyRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser,*mysqlPassword)
	citations, err := driver.GetAllCitations()
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{"0",len(citations), citations}, nil
}

func loadPage(filename string) []byte {
    body, err := ioutil.ReadFile(filename)
		checkErr(err)
    return body
}
