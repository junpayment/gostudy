package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

// templは1つのテンプレートを表します。
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// ServeHTTPはHTTPリクエストを処理します。
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// パッケージ以下ディレクトリだと
	temp, _ := os.Getwd()
	dir := temp + "/templates"

	t.once.Do(func() {
		t.templ =
			template.Must(template.ParseFiles(filepath.Join(dir, t.filename)))
	})
	t.templ.Execute(w, nil)
}

func main() {
	r := newRoom()
	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/room", r)

	go r.run()

	// webサーバを開始します
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
