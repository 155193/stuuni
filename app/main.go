package main

import (
	adminModels "./admin/models"
	adminServices "./admin/services"
	"./db"
	"./security/models"
	"./security/services"
	"context"
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/graphql-go/handler"
)

//user := os.Getenv("DB_USER")

var (
	ServerPort  = os.Getenv("SERVER_PORT")
	StoragePath = os.Getenv("STORAGE_PATH")
)

const maxUploadSize = 10 * 1024 * 1024 // 2 mb

// function to create and initialize collections in mongo
func InitializeBd() {
	// creating indexes

}

/**
main function
*/
func main() {
	_, _, err := db.ConnectDB()
	if err != nil {
		//panic(err)
	}
	InitializeBd()
	// Optional. Switch the session to a monotonic behavior.
	//session.SetMode(mgo.Monotonic, true)

	schema := *GetSchema()

	h := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   true,
		Playground: true,
		//ResultCallbackFn: MiddResultCallback,
	})

	http.Handle("/graphql", basicAuth(h))

	http.HandleFunc("/auth/login", loginHandle)

	http.HandleFunc("/storage/upload", uploadFileHandler())

	fs := http.FileServer(http.Dir(StoragePath))
	http.Handle("/storage/files/", http.StripPrefix("/storage/files", fs))

	fmt.Printf("Now server is running on port %s\n", ServerPort)
	tServerPort := fmt.Sprintf(":%s", ServerPort)
	http.ListenAndServe(tServerPort, nil)
}

func MiddResultCallback(ctx context.Context, params *graphql.Params, result *graphql.Result, responseBody []byte) {
	//fmt.Println("MiddResultCallback---------!!!!!")
	//fmt.Println(params.RootObject)
	//fmt.Println(result.Data)
	//fmt.Println(result.HasErrors())
	//fmt.Println(responseBody)
}

//
func basicAuth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		(w).Header().Set("Access-Control-Allow-Origin", "*")
		(w).Header().Set("Access-Control-Allow-Credentials", "true")
		(w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		// Get agent from client
		agent := r.Header.Get("User-Agent")
		// Get consult from graphql
		opts, _ := services.GetConsultGraphql(r)
		// Get local IP
		var local = ""
		ifaces, _ := net.InterfaceAddrs()
		for _, a := range ifaces {
			if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					local = ipnet.IP.String()
				}
			}
		}
		// Get client IP
		remote, _ := services.GetRemoteIp(r)

		s := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
		idUser := ""
		idAudit := primitive.NewObjectID()
		if len(s) != 2 {
			if err := services.AddAuditEvent(models.AuditUserToBson(idAudit, remote, local, agent, idUser, false, "", opts.Query)); err != nil {
				fmt.Println(err)
			}
			//http.Error(w, "Not authorized", 401)
			//return
			//println("not token")
		} else {
			id, err := adminServices.DecodeTokenString(s[1])
			if err == nil {
				idUser = id
				if err := services.AddAuditEvent(models.AuditUserToBson(idAudit, remote, local, agent, idUser, true, s[1], opts.Query)); err != nil {
					fmt.Println(err)
				}
			} else {
				if err := services.AddAuditEvent(models.AuditUserToBson(idAudit, remote, local, agent, idUser, false, s[1], opts.Query)); err != nil {
					fmt.Println(err)
				}
				//http.Error(w, "Not authorized", 401)
				//return
			}
		}
		authContext := adminServices.UserTkn{idUser}
		var v = models.Values{
			map[string]interface{}{
				"currentUser":  authContext,
				"currentAudit": idAudit,
				"idUser":       idUser,
				"time":         time.Now(),
			},
		}
		ctx := context.WithValue(r.Context(), "current", v)
		h.ServeHTTP(w, r.WithContext(ctx)) // call original
	})
}

type LoginAuth struct {
	Nick     string `json:"nick"`
	Password string `json:"password"`
}

type LoginToken struct {
	Token string `json:"token"`
}

func loginHandle(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var loginData LoginAuth
	err := decoder.Decode(&loginData)
	if err != nil || loginData.Nick == "" || loginData.Password == "" {
		http.Error(w, "Not authorized", 400)
		return
	}
	log.Println("-------test login")
	log.Println(loginData.Nick)
	log.Println(loginData.Password)

	nick, password, err := adminModels.Login(loginData.Nick, loginData.Password)
	if err != nil {
		http.Error(w, "Not authorized", 401)
		return
	}
	user, err := adminServices.Login(nick, password)
	fmt.Println(user)
	fmt.Println(err)
	if err != nil {
		http.Error(w, "Not authorized", 401)
		return
	}
	token, _ := adminServices.CreateTokenString(user.Id.Hex())
	if token == "" {
		http.Error(w, "Not authorized", 402)
		return
	}

	t := LoginToken{
		Token: token,
	}

	tkn, err := json.Marshal(t)
	if err != nil {
		http.Error(w, "Error", 500)
		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.WriteHeader(http.StatusOK)
	i, err := w.Write(tkn)
	if err != nil {
		fmt.Println(i)
		http.Error(w, "Error", 500)
	}
}

//Helper function to import json from file to map
func importJSONDataFromFile(fileName string, result interface{}) (isOK bool) {
	isOK = true
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Print("Error:", err)
		isOK = false
	}
	err = json.Unmarshal(content, result)
	if err != nil {
		isOK = false
		fmt.Print("Error:", err)
	}
	return
}
