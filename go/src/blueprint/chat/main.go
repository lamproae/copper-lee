package main

import (
	"flag"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/facebook"
	"github.com/stretchr/gomniauth/providers/github"
	"github.com/stretchr/gomniauth/providers/google"
	"github.com/stretchr/objx"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"text/template"
	"trace"
)

type templateHandler struct {
	once     sync.Once
	filename string
	tpl      *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.tpl = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})

	data := map[string]interface{}{
		"Host": r.Host,
	}

	if authCookie, err := r.Cookie("auth"); err == nil {
		data["UserData"] = objx.MustFromBase64(authCookie.Value)
	}
	t.tpl.Execute(w, data)
}

func main() {
	var addr = flag.String("addr", ":8080", "The addr of the application.")
	flag.Parse()
	gomniauth.SetSecurityKey("gTqMM6jUo7ogy5YnpfRDtISj")
	gomniauth.WithProviders(
		facebook.New("key", "secret", "http://www.kkkmmu.com:8080/auth/callback/facebook"),
		github.New("c6bb96679d3284c907cf", "a46ef8c09a4b1f081c2d08b244ab1ba6dbce8311", "http://www.kkkmmu.com:8080/auth/callback/github"),
		google.New("871487773908-kncl76ts3iisb3i3ua9r1gs1gns52pa9.apps.googleusercontent.com", "gTqMM6jUo7ogy5YnpfRDtISj", "http://www.kkkmmu.com:8080/auth/callback/google"),
	)

	r := newRoom()
	r.tracer = trace.New(os.Stdout)
	http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/room", r)
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
	http.HandleFunc("/auth/", loginHandler)
	go r.run()
	log.Println("Starting web server on", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServer: ", err.Error())
	}
}
