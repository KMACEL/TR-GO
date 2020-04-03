package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	ginserver "github.com/go-oauth2/gin-server"
	"github.com/go-session/session"
	"gopkg.in/oauth2.v3/errors"
	"gopkg.in/oauth2.v3/generates"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/models"
	"gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"
)

func main() {
	manager := manage.NewDefaultManager()
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)

	// token store
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	// generate jwt access token
	manager.MapAccessGenerate(generates.NewJWTAccessGenerate([]byte("00000000"), jwt.SigningMethodHS512))

	clientStore := store.NewClientStore()
	clientStore.Set("222222", &models.Client{
		ID:     "222222",
		Secret: "22222222",
		Domain: "http://localhost:9094",
	})

	manager.MapClientStorage(clientStore)

	srv := ginserver.InitServer(manager)
	ginserver.SetAllowGetAccessRequest(true)
	ginserver.SetClientInfoHandler(server.ClientFormHandler)

	srv.SetPasswordAuthorizationHandler(func(username, password string) (userID string, err error) {
		if username == "test" && password == "test" {
			userID = "test"
		}
		return
	})

	srv.SetUserAuthorizationHandler(userAuthorizeHandler)

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})

	g := gin.Default()

	g.GET("/auth", authHandler)
	g.GET("/login", loginHandler)
	g.POST("/login", loginHandler)

	a := func(c *gin.Context) {
		store, err := session.Start(nil, c.Writer, c.Request)
		if err != nil {
			http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
			return
		}

		var form url.Values
		if v, ok := store.Get("ReturnUri"); ok {
			form = v.(url.Values)
		}
		c.Request.Form = form

		store.Delete("ReturnUri")
		store.Save()

		err = srv.HandleAuthorizeRequest(c.Writer, c.Request)
		if err != nil {
			fmt.Println("mert")
			http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		}
	}

	g.GET("/authorize", a)
	g.POST("/authorize", a)

	g.POST("/token", func(c *gin.Context) {
		err := srv.HandleTokenRequest(c.Writer, c.Request)
		if err != nil {
			http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		}
	})

	g.GET("/test", func(c *gin.Context) {
		token, err := srv.ValidationBearerToken(c.Request)
		if err != nil {
			http.Error(c.Writer, err.Error(), http.StatusBadRequest)
			return
		}

		data := map[string]interface{}{
			"expires_in": int64(token.GetAccessCreateAt().Add(token.GetAccessExpiresIn()).Sub(time.Now()).Seconds()),
			"client_id":  token.GetClientID(),
			"user_id":    token.GetUserID(),
		}
		e := json.NewEncoder(c.Writer)
		e.SetIndent("", "  ")
		e.Encode(data)
	})

	log.Println("Server is running at 9096 port.")
	log.Fatal(g.Run(":9096"))
}

func userAuthorizeHandler(w http.ResponseWriter, r *http.Request) (userID string, err error) {
	fmt.Println("test")
	store, err := session.Start(nil, w, r)
	if err != nil {
		return
	}

	uid, ok := store.Get("LoggedInUserID")
	if !ok {
		if r.Form == nil {
			r.ParseForm()
		}

		store.Set("ReturnUri", r.Form)
		store.Save()

		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusFound)
		return
	}

	userID = uid.(string)
	store.Delete("LoggedInUserID")
	store.Save()
	return
}

func userAuthorizeHandlerGIN(c *gin.Context) (userID string, err error) {
	store, err := session.Start(nil, c.Writer, c.Request)
	if err != nil {
		return
	}

	uid, ok := store.Get("LoggedInUserID")
	if !ok {
		if c.Request.Form == nil {
			c.Request.ParseForm()
		}

		store.Set("ReturnUri", c.Request.Form)
		store.Save()

		c.Writer.Header().Set("Location", "/login")
		c.Writer.WriteHeader(http.StatusFound)
		return
	}

	userID = uid.(string)
	store.Delete("LoggedInUserID")
	store.Save()
	return
}

func loginHandler(c *gin.Context) {
	store, err := session.Start(nil, c.Writer, c.Request)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	if c.Request.Method == "POST" {
		store.Set("LoggedInUserID", "000000")
		store.Save()

		c.Writer.Header().Set("Location", "/auth")
		c.Writer.WriteHeader(http.StatusFound)
		return
	}
	outputHTML(c.Writer, c.Request, "static/login.html")
}

func authHandler(c *gin.Context) {
	store, err := session.Start(nil, c.Writer, c.Request)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, ok := store.Get("LoggedInUserID"); !ok {
		c.Writer.Header().Set("Location", "/login")
		c.Writer.WriteHeader(http.StatusFound)
		return
	}

	outputHTML(c.Writer, c.Request, "static/auth.html")
}

func outputHTML(w http.ResponseWriter, req *http.Request, filename string) {
	file, err := os.Open(filename)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer file.Close()
	fi, _ := file.Stat()
	http.ServeContent(w, req, file.Name(), fi.ModTime(), file)
}
