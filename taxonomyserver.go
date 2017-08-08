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

	//TODO: Remove cors
	// cors := tigertonic.NewCORSBuilder()//.AddAllowedOrigins("*").AddAllowedHeaders("Origin, X-Requested-With, Content-Type, Accept")

	mux := tigertonic.NewTrieServeMux()
	mux.Handle("POST", "/correlation", tigertonic.Timed(tigertonic.Marshaled(getCorrelationHandler), "getCorrelationHandler", nil))
	mux.Handle("GET", "/attribute", tigertonic.Timed(tigertonic.Marshaled(getAttributesHandler), "getAttributesHandler", nil))
	mux.Handle("GET", "/citation", tigertonic.Timed(tigertonic.Marshaled(getCitationsHandler), "getCitationsHandler", nil))
	// mux.Handle("GET","/",cors.Build(tigertonic.Timed(tigertonic.Marshaled(getIndexHandler), "getIndexHandler", nil)))
	mux.HandleFunc("GET", "/", func(w http.ResponseWriter, r *http.Request) {
    p := loadPage("frontend/taxonomy/index.html")
    fmt.Fprintf(w, "%s", p.Body)
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

func loadPage(filename string) *Page {
    body, err := ioutil.ReadFile(filename)
		checkErr(err)
    return &Page{Body: body}
}
