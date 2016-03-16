package front

import (
    "net/http"
//    "strings"
    "fmt"
    "mime"
    "path/filepath"
    "log"
)

func httpServer(addr string) {
    initTmpls()
    http.HandleFunc("/ping", ping)
    http.HandleFunc("/static/", static)
    http.HandleFunc("/", index)
    http.HandleFunc("/index", index)

    fmt.Println("master http start!", addr)
    log.Fatal(http.ListenAndServe(addr, nil))
}

func ping(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("pong"))
}


func static(w http.ResponseWriter, r *http.Request) {
    urlpath := r.URL.Path
    data, err := Asset("html" + urlpath)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    ctype := mime.TypeByExtension(filepath.Ext(urlpath))
    w.Header().Set("Content-Type", ctype)
    w.Write(data)
}

func index(w http.ResponseWriter, r *http.Request) {
    err := templates.ExecuteTemplate(w, "index.html", "")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
