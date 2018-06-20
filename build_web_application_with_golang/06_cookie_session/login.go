package main

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var globalSessions *session.Manager

type Provider interface {
	SessionInit(sid string) (Session, error)
	SessionRead(sid string) (Session, error)
	SessionDestroy(sid string) error
	SessionGC(maxLifeTime int64)
}
type Session interface {
	Set(key, value interface{}) error
	Get(key interface{}) interface{}
	Delete(key interface{}) error
	SessionID() string
}
type Manager struct {
	cookieName  string
	lock        sync.Mutex
	provider    Provider
	maxLifeTime int64
}

// 这个文件还不能用，有点没理清

func init() {
	globalSessions, _ = NewManager("memory", "gosessionid", 3600)
}

var provides = make(map[string]Provider)

func Register(name string, provider Provider) {
	if provider == nil {
		panic("session: register provider is nil")
	}
	if _, dup := provides[name]; dup {
		panic("session: register called twice for provide " + name)
	}
	provides[name] = provider
}
func (manager *Manager) sessionId() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}
func NewManager(providerName, cookieName string, maxLifeTime int64) (*Manager, error) {
	provider, ok := provides[providerName]
	if !ok {
		return nil, fmt.Errorf("session: unknown provide %q (forgotten import?)", providerName)
	}
	return &Manager{provider: provider, cookieName: cookieName, maxLifeTime: maxLifeTime}, nil
}
func login(w http.ResponseWriter, r *http.Request) {
	sess, _ := beego.GlobalSessions.SessionStart(w, r)
	r.ParseForm()
	if r.Method == "GET" {
		cruntime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(cruntime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("login.gtpl")
		log.Println(t.Execute(w, token))
		w.Header().Set("Content-Type", "text/html")
		t.Execute(w, sess.Get("username"))
	} else {
		sess.Set("username", r.Form["username"])
		http.Redirect(w, r, "/", 302)
		token := r.Form.Get("token")
		if token != "" {
			// verify token
		} else {
			// errors.New("No token")
		}
		fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username")))
		fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
		template.HTMLEscape(w, []byte(r.Form.Get("username")))
	}
}

func main() {
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
