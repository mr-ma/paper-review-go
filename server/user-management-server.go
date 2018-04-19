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
	listen        = flag.String("listen", "127.0.0.1:8002", "listen address")
)

var sessionManager *scs.Manager
var usermanagementDriver data.UsermanagementDriver

func main() {
	flag.Parse()

	usermanagementDriver = data.InitUsermanagementDriver(*mysqlUser, *mysqlPassword, *mysqlServer)
	dbRef, err := usermanagementDriver.OpenDB()
	if err == nil {
		sessionManager = scs.NewManager(mysqlstore.New(dbRef, 600000000000))
		sessionManager.Lifetime(time.Hour * 24) // session data expires after 24 hours
		sessionManager.Persist(true)            // session data persists after the browser has been closed by the user
		//sessionManager.Secure(true)
	}

	mux := tigertonic.NewTrieServeMux()

	// user and access management:

	mux.HandleFunc("POST", "/usermanagement/login", func(w http.ResponseWriter, r *http.Request) {
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
	mux.HandleFunc("POST", "/usermanagement/logout", func(w http.ResponseWriter, r *http.Request) {
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
	mux.HandleFunc("POST", "/usermanagement/saveUser", func(w http.ResponseWriter, r *http.Request) {
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
	mux.HandleFunc("GET", "/usermanagement/user", func(w http.ResponseWriter, r *http.Request) {
		checkAdmin(w, r, getUserHandler)
	})
	mux.HandleFunc("GET", "/usermanagement/getUsers", func(w http.ResponseWriter, r *http.Request) {
		checkAdmin(w, r, getUsersHandler)
	})
	mux.HandleFunc("POST", "/usermanagement/query", func(w http.ResponseWriter, r *http.Request) {
		checkAdmin(w, r, queryHandler)
	})
	mux.HandleFunc("POST", "/usermanagement/createUser", func(w http.ResponseWriter, r *http.Request) {
		checkAdmin(w, r, createUserHandler)
	})
	mux.HandleFunc("POST", "/usermanagement/updateUser", func(w http.ResponseWriter, r *http.Request) {
		checkAdmin(w, r, updateUserHandler)
	})
	mux.HandleFunc("POST", "/usermanagement/deleteUser", func(w http.ResponseWriter, r *http.Request) {
		checkAdmin(w, r, deleteUserHandler)
	})
	mux.HandleFunc("POST", "/usermanagement/taxonomyPermissions", func(w http.ResponseWriter, r *http.Request) {
		checkAdmin(w, r, getTaxonomyPermissionsHandler)
	})
	mux.HandleFunc("POST", "/usermanagement/updateTaxonomyPermissions", func(w http.ResponseWriter, r *http.Request) {
		checkAdmin(w, r, getUpdateTaxonomyPermissionsHandler)
	})

	mux.HandleFunc("GET", "/usermanagement/users", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("../frontend/taxonomy/users/users.html")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/usermanagement/modals.html", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("../frontend/taxonomy/users/modals.html")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/usermanagement/navbar.html", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("../frontend/taxonomy/users/navbar.html")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/usermanagement/navbarAdmin.html", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("../frontend/taxonomy/users/navbarAdmin.html")
		fmt.Fprintf(w, "%s", p)
	})

	mux.HandleFunc("GET", "/usermanagement/users.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("../frontend/src/js/users.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/usermanagement/userManagement.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("../frontend/src/js/userManagement.js")
		fmt.Fprintf(w, "%s", p)
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
	result, err := usermanagementDriver.GetTaxonomyPermissions(userRequest.Email)
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
	result, err := usermanagementDriver.UpdateTaxonomyPermissions(taxonomyPermissionsRequest.Email, taxonomyPermissionsRequest.Permissions)
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
