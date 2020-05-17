package main

import (
  "log"
  "net/http"
  "text/template"
  "path/filepath"
  "sync"
)

// templは1つのテンプレートを表す
type templateHandler struct {
  once  sync.Once
  filename string
  templ *template.Template
}

// HTTPリクエストを処理するtemplateHandler型のメソッド
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  // t.once.Doを定義しておくと、複数のゴルーチンがServeHTTPを実行してもここは一度しか実行されない
  t.once.Do(func() {
    t.templ =
      template.Must(template.ParseFiles(filepath.Join("templates",
        t.filename)))
  })
  t.templ.Execute(w, nil)
}


func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte(`
      <html>
        <head>
          <title>チャット</title>
        </head>
        <body>
          チャットしましょう！
        </body>
      </html>
      `))
  })

  // WEBサーバーを開始
  if err := http.ListenAndServe(":8080", nil); err != nil {
    log.Fatal("ListenAndServe:", err)
  }
}
