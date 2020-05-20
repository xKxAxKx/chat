package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"text/template"

	"github.com/xKxAxKx/chat/trace"
)

// templは1つのテンプレートを表す
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// HTTPリクエストを処理するtemplateHandler型のメソッド
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// t.once.Doを定義しておくと、複数のゴルーチンがServeHTTPを実行してもここは一度しか実行されない
	t.once.Do(func() {
		t.templ =
			template.Must(template.ParseFiles(filepath.Join("templates",
				t.filename)))
	})
	t.templ.Execute(w, r)
}

func main() {
	var addr = flag.String("addr", ":8080", "アプリケーションのアドレス")
	flag.Parse() // フラグを解釈する

	r := newRoom()
	r.tracer = trace.New(os.Stdout)

	http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.Handle("/room", r)

	// チャットルームを開始する
	go r.run()

	// WEBサーバーを開始
	log.Println("Webサーバーを開始します。ポート:", *addr)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
