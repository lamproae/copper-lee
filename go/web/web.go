package main

import (
	"fmt"
	"log"
	"net/http"
	//	"net/url"
	"crypto/md5"
	"io"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"
	//	"sync"
	//	"crypto/rand"
	"html/template"
	//	"encoding/base64"
	tt "text/template"
)

/*
var globalSessons *Manager
var providers = make(map[string]Provider)

type Manager struct {
	cookieName string
	lock sync.Mutex
	provider Provider
	maxlifetime int64
}

func NewManager(providerName, cookieName string, maxlifetime int64)(*Manager, error) {
	provider, ok := providers[providerName]
	if !ok {
		return nil, fmt.Errorf("session: unknown provider %q(forgotten import?)", providerName)
	}
	return &Manager{provider:provider, cookieName:cookieName, maxlifetime:maxlifetime}, nil
}

func (manager *Manager) SessionId() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}

	fmt.Println("in manager sessionid")
	return base64.URLEncoding.EncodeToString(b)
}

func (manager *Manager)SessionStart(w http.ResponseWriter, r *http.Request)(session Session) {
	fmt.Println("in Session start");
	manager.lock.Lock()
	fmt.Println("in Session start1");
	defer manager.lock.Unlock()
	cookie, err := r.Cookie(manager.cookieName)
	fmt.Println("in Session start1");
	if err != nil || cookie.Value == "" {
	fmt.Println("in Session start2");
		sid := manager.SessionId()
		session, _ = manager.provider.SessionInit(sid)
		cookie := http.Cookie{Name: manager.cookieName, Value:url.QueryEscape(sid), Path:"/", HttpOnly:true, MaxAge:int(manager.maxlifetime)}
		http.SetCookie(w, &cookie)
	} else {
	fmt.Println("in Session start3");
		sid, _ := url.QueryUnescape(cookie.Value)
		session, _ = manager.provider.SessionRead(sid)
	}

	fmt.Println("in Session start");
	return
}

type Provider interface {
	SessionInit(sid string) (Session, error)
	SessionRead(sid string) (Session, error)
	SessionDestroy(sid string) error
	SessionGC(maxlifetime int64)
}

type Session interface {
	Set(key, value interface{}) error
	Get(key interface{}) interface{}
	Delete(key interface{}) error
	SessionID() string
}

type TP struct {
	se map[string]TS
}

type TS struct {
	me string
}

func (p TS)Set(key , value  interface{}) (err error) {
	p.me = value.(string)
	return nil
}

func (p TS)Get(key interface{}) interface{} {
	return p.me
}

func (p TS)Delete(key interface{}) (err error){
	p.me  = "";
	return err
}

func (p TS)SessionID() string {
	return p.me
}

func (p *TP)SessionInit(sid string) (session Session, err error) {
	if v, ok :=  p.se[sid]; ok {
		return v, nil;
	} else {
		sess := TS{}
		sess.me = ""
		p.se[sid] = sess;
		return sess, nil
	}
}

func (p *TP)SessionRead(sid string) (session Session, err error) {
	if v, ok :=  p.se[sid]; ok {
		return v, nil;
	} else {
		sess := TS{}
		sess.me = ""
		p.se[sid] = sess;
		return sess, nil
	}
}

func (p *TP)SessionDestroy(sid string) error {
	if _, ok :=  p.se[sid]; ok {
		delete(p.se, sid)
	}
	return nil
}

func (p *TP)SessionGC(maxlifetime int64) {

}

func Register(name string, provider Provider) {
	if provider == nil {
		panic("Session: Register provide is nil")
	}

	if _, dup := providers[name]; dup {
		panic("Session: Register called twice for provide " + name)
	}

	providers[name] = provider
}

func init() {
	prov := &TP{}
	Register("memory", prov)
	globalSessons, _ = NewManager("memory", "gosessionid", 3600)
}
*/

func login(w http.ResponseWriter, r *http.Request) {
	//sess := globalSessons.SessionStart(w, r)
	r.ParseForm()
	fmt.Println("methd: ", r.Method)
	if r.Method == "GET" {
		fmt.Println(r.Form)
		fmt.Println("path", r.URL.Path)
		for k, v := range r.Form {
			fmt.Println("key:", k)
			fmt.Println("val:", strings.Join(v, ""))
		}

		curtime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(curtime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		expiration := time.Now().Add(time.Duration(3600 * time.Second))
		cookie := http.Cookie{Name: "liwei", Value: "goooooooooooogl", Expires: expiration}
		http.SetCookie(w, &cookie)
		cookie = http.Cookie{Name: "wxmm", Value: "zoooooooooooogl", Expires: expiration}
		http.SetCookie(w, &cookie)
		cookie = http.Cookie{Name: "yandz", Value: "zxxxxxxsxxxxxgl", Expires: expiration}
		http.SetCookie(w, &cookie)
		t, _ := template.ParseFiles("./html/login.html")
		//w.Header().Set("Content-Type", "text/html")
		t.Execute(w, token)
		//t.Execute(w, sess.Get("username"))
	} else {
		/*
			rcookie, _ := r.Cookie("liwei")
			fmt.Println(rcookie)

			for _, rcookie = range r.Cookies() {
				fmt.Println(rcookie);
			}
			fmt.Fprintf(w, "Your username is: %s\n", r.Form.Get("username"))
			fmt.Fprintf(w, "Your telephone is: %s\n", r.Form.Get("mobile"))
			fmt.Fprintf(w, "Your email is: %s\n", r.Form.Get("email"))
			fmt.Fprintf(w, "Your token is: %s\n", r.Form.Get("token"))
			fmt.Fprintf(w, "Your passord is: %s\n", r.Form.Get("password"))
			expiration := time.Now().Add(time.Duration(3600 * time.Second))
			cookie := http.Cookie{Name: "liwei", Value:"Gooooooooooooooooogle", Expires: expiration}
			http.SetCookie(w,&cookie)
			t, _ := template.ParseFiles("./html/welcome.html")
			t.Execute(w, nil)
		*/
		//sess.Set("username", r.Form.Get("username"))
		http.Redirect(w, r, "./html/welcome.html", 302)
	}
}

func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("methd: ", r.Method)

	if r.Method == "GET" {
		curtime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(curtime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("./html/upload.html")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		defer file.Close()
		/*
			fmt.Fprintf(w, "%v", handler.Header)
		*/
		f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(r.FormFile("uploadfile"))
			fmt.Println(err)
			return
		}

		defer f.Close()
		io.Copy(f, file)

		t, _ := template.ParseFiles("./html/upload.html")
		t.Execute(w, nil)
	}
}

func Jsload(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "GET" {
		//	w.Header().Set("Content-type", "text/html; charset=utf-8")
		w.Header().Set("Content-Type", "application/javascript")
		fmt.Println("path: %s\n", r.URL.Path)
		if m, _ := regexp.MatchString("^/js", r.URL.Path); m {
			file := path.Base(r.URL.Path)
			file = "./js/" + file
			fmt.Println(file)
			t, _ := tt.ParseFiles(file)
			t.Execute(w, nil)
		}
	}
}

func CSSload(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "GET" {
		fmt.Println("path: %s\n", r.URL.Path)
		if m, _ := regexp.MatchString("^/css", r.URL.Path); m {
			file := path.Base(r.URL.Path)
			fmt.Println(file)
			file = "./css/" + file
			fmt.Println(file)
			t, _ := template.ParseFiles(file)
			t.Execute(w, nil)
		}
	}
}

func HTMLload(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "GET" {
		fmt.Println("path: %s\n", r.URL.Path)
		if m, _ := regexp.MatchString("^/html", r.URL.Path); m {
			file := path.Base(r.URL.Path)
			fmt.Println(file)
			file = "./html/" + file
			fmt.Println(file)
			t, _ := template.ParseFiles(file)
			t.Execute(w, nil)
		}
	}
}

func download(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "GET" {
		fmt.Println("path: %s\n", r.URL.Path)
		if m, _ := regexp.MatchString("^/test", r.URL.Path); m {
			file := path.Base(r.URL.Path)
			fmt.Println(file)
			file = "./test/" + file
			fmt.Println(file)
			f, err := os.Open(file)
			if err != nil {
				http.Error(w, "file not found", 404)
				return
			}
			defer f.Close()

			io.Copy(w, f)
		}
	}
}

func main() {
	http.HandleFunc("/", login)
	http.HandleFunc("/test/", download)
	http.HandleFunc("/login", login)
	http.HandleFunc("/html/upload", upload)
	http.HandleFunc("/html/upload.html", upload)
	http.HandleFunc("/html/", HTMLload)
	http.HandleFunc("/js/", Jsload)
	http.HandleFunc("/css/", CSSload)

	//err := http.ListenAndServe("10.71.1.124:9090", nil)
	err := http.ListenAndServe("192.168.1.103:9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe err: ", err)
	}
}
