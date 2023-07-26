package handler

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	authrepository "mytesting/pkg/repository/auth_repository"
	clientrepository "mytesting/pkg/repository/client_repository"
	codesrepository "mytesting/pkg/repository/codes_repository"
	userrepository "mytesting/pkg/repository/user_repository"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	userToken     = "user_token"
	usernameToken = "username_token"
	secretToken   = "secret_token"
	allToken      = "all_token"
	jwtCookieName = "jwt_token"
)

type Handler struct {
	auth    authrepository.Auth
	clients clientrepository.Clients
	users   userrepository.UserRepotisory
	codes   codesrepository.Codes
}

func NewHandler(auth authrepository.Auth, clients clientrepository.Clients, users userrepository.UserRepotisory, codes codesrepository.Codes) Handler {
	return Handler{auth: auth, clients: clients, users: users, codes: codes}
}

func (h Handler) InitRouter() http.Handler {
	r := gin.Default()

	r.GET("/signup", h.GetRegisterForm)
	r.POST("/signup", h.Register)

	r.GET("/signin", h.GetLoginForm)
	r.POST("/signin", h.Login)

	oauthGroup := r.Group("/oauth")
	oauthGroup.Use(h.CheckUserToken)
	oauthGroup.GET("/", h.GenerateCodeForm)
	oauthGroup.POST("/", h.GenerateCode)
	oauthGroup.GET("/authorize", h.GenerateToken)

	api := r.Group("/api")
	api.Use(h.CheckUserToken)
	api.GET("/info", h.ApiGetInfo)
	api.GET("/secret", h.ApiGetSecret)

	userGroup := r.Group("/user")
	userGroup.Use(h.CheckUserToken)
	userGroup.GET("/", h.GetUserData)

	userGroup.GET("/client", h.registerNewClientForm)
	userGroup.POST("/client", h.registerNewClient)

	userGroup.GET("/client/:client", h.getClient)

	return r
}

func (h Handler) ApiGetInfo(ctx *gin.Context) {
	d, ok := ctx.Get("username")
	if !ok {
		http.Error(ctx.Writer, "WTF", http.StatusInternalServerError)
	}

	v, ok := d.(authrepository.AuthData)
	if !ok {
		http.Error(ctx.Writer, "WTF", http.StatusInternalServerError)
	}

	if v.TokenType == usernameToken || v.TokenType == allToken {
		ctx.JSON(http.StatusOK, struct{
			Username string
		}{
			Username: v.Username,
		})
	} else {
		http.Error(ctx.Writer, "Access denied", http.StatusInternalServerError)
		return
	}
}

func (h Handler) ApiGetSecret(ctx *gin.Context) {
	d, ok := ctx.Get("username")
	if !ok {
		http.Error(ctx.Writer, "WTF", http.StatusInternalServerError)
	}

	v, ok := d.(authrepository.AuthData)
	if !ok {
		http.Error(ctx.Writer, "WTF", http.StatusInternalServerError)
	}

	if v.TokenType == secretToken || v.TokenType == allToken {
		user, err := h.users.GetByName(v.Username)
		if err != nil {
			http.Error(ctx.Writer, err.Error(), http.StatusInternalServerError)
			return
		}

		ctx.JSON(http.StatusOK, struct{
			Secret string
		}{
			Secret: user.Secret,
		})
	} else {
		http.Error(ctx.Writer, "Access denied", http.StatusInternalServerError)
		return
	}
}

func (h Handler) GenerateToken(ctx *gin.Context) {
	client_id := ctx.Query("client_id")
	redirect_uri := ctx.Query("redirect_uri")
	client_secret := ctx.Query("client_secret")
	code := ctx.Query("code")

	b, err := h.auth.EncryptedMessage(code)
	if err != nil {
		http.Error(ctx.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	codeStruct := struct {
		Id    string
		Time  time.Time
		Scope string
	}{}

	json.Unmarshal([]byte(b), &codeStruct)

	id, err := strconv.Atoi(client_id)
	if err != nil {
		http.Error(ctx.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	name, secret, ok := h.clients.GetById(id)
	if !ok {
		http.Error(ctx.Writer, "Cant find id", http.StatusInternalServerError)
		return
	}

	if secret != client_secret {
		http.Error(ctx.Writer, "Incorrect secret", http.StatusInternalServerError)
		return
	}

	if codeStruct.Id != client_id {
		http.Error(ctx.Writer, "Incorrect id or secret", http.StatusInternalServerError)
		return
	}

	token := h.auth.NewToken(name, codeStruct.Scope)
	p, _ := url.Parse(redirect_uri)
	q := p.Query()

	p.RawQuery = q.Encode()

	ctx.JSON(200, token)
	ctx.Redirect(301, p.String())
}

// http://localhost:8080/oauth?client_id=103971&redirect_uri=http://localhost:8080
func (h Handler) GenerateCode(ctx *gin.Context) {
	client_id := ctx.Query("client_id")
	redirect_uri := ctx.Query("redirect_uri")
	scope := ctx.Query("scope")
	if scope == "" {
		scope = usernameToken
	}

	id, err := strconv.Atoi(client_id)
	if err != nil {
		http.Error(ctx.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	_, _, ok := h.clients.GetById(id)
	if !ok {
		http.Error(ctx.Writer, "cant find program", http.StatusInternalServerError)
		return
	}

	codeStruct := struct {
		Id    string
		Time  time.Time
		Scope string
	}{
		Id:    client_id,
		Time:  time.Now(),
		Scope: scope,
	}
	b, _ := json.Marshal(codeStruct)
	message, err := h.auth.NewCryptedMessage(b)

	p, _ := url.Parse(redirect_uri)
	q := p.Query()
	q.Add("code", message)
	p.RawQuery = q.Encode()

	ctx.Redirect(301, p.String())
}

func (h Handler) GenerateCodeForm(ctx *gin.Context) {
	scope := ctx.Query("scope")
	data := strings.Builder{}
	if scope == "" || scope == usernameToken {
		data.WriteString("Watch Username")
	} else if scope == secretToken {
		data.WriteString("Watch Secret")
	} else if scope == allToken {
		data.WriteString("Watch Username and Watch Secret")
	} else {
		http.Error(ctx.Writer, "unknow scope", http.StatusInternalServerError)
		return
	}
	fp := path.Join("./templates", "code.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(ctx.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.ExecuteTemplate(ctx.Writer, "code", data.String()); err != nil {
		http.Error(ctx.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h Handler) getClient(ctx *gin.Context) {
	d, ok := ctx.Get("username")
	if !ok {
		http.Error(ctx.Writer, "WTF", http.StatusInternalServerError)
	}

	v, ok := d.(authrepository.AuthData)
	if !ok {
		http.Error(ctx.Writer, "WTF", http.StatusInternalServerError)
	}
	s, _ := ctx.Params.Get("client")

	id, secret, ok := h.clients.GetByName(s)
	if !ok {
		http.Error(ctx.Writer, "cant find program by name "+s, http.StatusInternalServerError)
		return
	}

	user, err := h.users.GetByName(v.Username)
	if err != nil {
		http.Error(ctx.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, ok := user.Clients[s]; !ok {
		http.Error(ctx.Writer, "cant find app", http.StatusInternalServerError)
		return
	}

	ctx.String(200, "Name: %s, Id: %d, Secret %s", s, id, secret)
}

func (h Handler) registerNewClientForm(ctx *gin.Context) {
	fp := path.Join("./templates", "client.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(ctx.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.ExecuteTemplate(ctx.Writer, "client", nil); err != nil {
		http.Error(ctx.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h Handler) registerNewClient(ctx *gin.Context) {
	d, ok := ctx.Get("username")
	if !ok {
		http.Error(ctx.Writer, "WTF", http.StatusInternalServerError)
	}

	v, ok := d.(authrepository.AuthData)
	if !ok {
		http.Error(ctx.Writer, "WTF", http.StatusInternalServerError)
	}

	data := struct {
		Client string `form:"client"`
	}{}

	ctx.Bind(&data)

	err := h.clients.Create(data.Client)
	if err != nil {
		http.Error(ctx.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	h.users.AddClientToUser(v.Username, data.Client)

	ctx.String(http.StatusOK, "%s", data.Client)
}

func (h Handler) GetUserData(ctx *gin.Context) {
	d, ok := ctx.Get("username")
	if !ok {
		http.Error(ctx.Writer, "WTF", http.StatusInternalServerError)
	}

	v, ok := d.(authrepository.AuthData)
	if !ok {
		http.Error(ctx.Writer, "WTF", http.StatusInternalServerError)
	}

	u, err := h.users.GetByName(v.Username)
	if err != nil {
		http.Error(ctx.Writer, err.Error(), http.StatusInternalServerError)
	}

	data := struct {
		Username string
		Secret   string
		Programs map[string]struct{}
	}{
		Username: u.Username,
		Secret:   u.Secret,
		Programs: u.Clients,
	}

	ctx.JSONP(http.StatusOK, data)
}

func (h Handler) CheckUserToken(ctx *gin.Context) {
	s, err := ctx.Cookie(jwtCookieName)
	if err != nil {
		ctx.Redirect(301, "/signin")
	}

	if s == "" {
		s = ctx.GetHeader("Authorization")
	}

	data, err := h.auth.CheckToken(s, userToken)
	if err != nil {
		ctx.Redirect(301, "/signin")
	}

	ctx.Set("username", data)
	ctx.Next()
}

func (h Handler) GetLoginForm(ctx *gin.Context) {
	fp := path.Join("./templates", "login.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		log.Println(err)
		return
	}

	if err := tmpl.ExecuteTemplate(ctx.Writer, "login", nil); err != nil {
		http.Error(ctx.Writer, err.Error(), http.StatusInternalServerError)
	}
}

func (h Handler) Login(ctx *gin.Context) {
	data := struct {
		Login    string `form:"login"`
		Password string `form:"password"`
	}{}

	ctx.Bind(&data)
	user, err := h.users.GetUser(data.Login, data.Password)
	if err != nil {
		http.Error(ctx.Writer, err.Error(), http.StatusNotFound)
		return
	}

	token := h.auth.NewToken(user.Username, userToken)

	fmt.Println(token)
	ctx.SetCookie(jwtCookieName, token, int(time.Minute*30), "", "", true, false)
	ctx.Redirect(301, "/user")
}

func (h Handler) GetRegisterForm(ctx *gin.Context) {
	fp := path.Join("./templates", "register.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		log.Println(err)
		return
	}

	if err := tmpl.ExecuteTemplate(ctx.Writer, "registration", nil); err != nil {
		http.Error(ctx.Writer, err.Error(), http.StatusInternalServerError)
	}
}

func (h Handler) Register(ctx *gin.Context) {
	data := struct {
		Login    string `form:"login"`
		Password string `form:"password"`
		Secret   string `form:"secret"`
	}{}

	ctx.Bind(&data)

	err := h.users.CreateUser(data.Login, data.Password, data.Secret)
	if err != nil {
		http.Error(ctx.Writer, err.Error(), http.StatusNotFound)
		return
	}

	ctx.Redirect(301, "/signin")
}
