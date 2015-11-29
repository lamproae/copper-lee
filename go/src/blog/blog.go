package main

import (
	"github.com/gorilla/sessions"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var err error
var tpl *template.Template
var store = sessions.NewCookieStore([]byte("something-very-secret"))

type Model struct {
	Files []string
}

var Data Model = getFilePaths()

func getFilePaths() Model {
	files := []string{}
	//files := make([]string, 0)

	filepath.Walk("./assets", func(path string, fi os.FileInfo, err error) error {
		if fi.IsDir() {
			return nil
		}

		path = strings.Replace(path, "\\", "/", -1)
		if strings.HasSuffix(path, ".go") {
			files = append(files, path)
		}

		return nil
	})

	return Model{Files: files}
}

func index(res http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(res, "index.gohtml", Data)
}

func login(res http.ResponseWriter, req *http.Request) {
	session, _ := store.Get(req, "session")
	if req.Method == "POST" {
		password := req.FormValue("password")
		if password == "secret" {
			session.Values["logged_in"] = true
		} else {
			http.Error(res, "invalid credentials", 401)
			return
		}

		session.Save(req, res)
		http.Redirect(res, req, "/", 302)
	}

	tpl.ExecuteTemplate(res, "login.gohtml", nil)
}

func logout(res http.ResponseWriter, req *http.Request) {
	session, _ := store.Get(req, "session")
	delete(session.Values, "logged_in")
	session.Save(req, res)
	http.Redirect(res, req, "/", 302)
}

func upload(res http.ResponseWriter, req *http.Request) {
	session, _ := store.Get(req, "session")
	if session.Values["logged_in"] == false ||
		session.Values["logged_in"] == nil {
		http.Redirect(res, req, "/login", 302)
	}

	if req.Method == "POST" {
		src, hdr, err := req.FormFile("my-file")
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
		defer src.Close()

		path := "./assets"
		dst, err := os.Create(path + hdr.Filename)
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
		defer dst.Close()

		io.Copy(dst, src)
		Data = getFilePaths()
		http.Redirect(res, req, "/", 302)
	}
	tpl.ExecuteTemplate(res, "upload-file.gohtml", nil)
}

func main() {
	tpl, err = tpl.ParseGlob("templates/*.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)
	http.HandleFunc("/upload", upload)
	http.HandleFunc("/logout", logout)
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())

	// to gernerate cert and key:
	// go run $(go env GROOT)/src/crypto/tsl/generate_cert.go --host=localhost
	http.ListenAndServeTLS("10.71.1.174:8080", "cert.pem", "key.pem", nil)
}
