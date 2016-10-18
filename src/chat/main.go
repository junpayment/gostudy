package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

// temp1は1つのテンプレートを表します。
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// ServeHTTPはHTTPリクエストを処理します。
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	temp, _ := os.Getwd()
	dir := temp + "/src/chat/templates"
	t.once.Do(func() {
		t.templ =
			template.Must(template.ParseFiles(filepath.Join(dir, t.filename)))
	})
	t.templ.Execute(w, nil)
}

func main() {
	log.Println(os.Getwd())

	http.Handle("/", &templateHandler{filename: "chat.html"})

	// webサーバを開始します
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
