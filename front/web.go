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
    http.HandleFunc("/login", login)
    http.HandleFunc("/forms", forms)
    http.HandleFunc("/forms.html", forms)
    http.HandleFunc("/index", index)
    http.HandleFunc("/allgame", allgame)

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

func forms(w http.ResponseWriter, r *http.Request) {
    err := templates.ExecuteTemplate(w, "forms.html", "")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func allgame(w http.ResponseWriter, r *http.Request) {
//    fmt.Println(getinfo())
//    htmlServer := [10]int{23,34, 12, 34, 1, 13, 43, 1, 5, 66}
//    fmt.Println(htmlServer)
//    abc := map[string]string{
//       "a": "bc",
//       "b": "ac",
//    }
    // err := templates.ExecuteTemplate(w, "allgame.html", htmlServer)
    results := getinfo()
    for _, row := range results {
        val1 := row[1].([]byte)
        fmt.Println(val1)
        fmt.Println(row.Int(0))
        fmt.Println(row.Str(1))
        fmt.Println(row.Str(2))
        fmt.Println(row.Str(3))
        fmt.Println(row.Str(4))
        fmt.Println(row.Str(5))
        fmt.Println(row.Str(6))
        fmt.Println(row.Str(7))
    }

    err := templates.ExecuteTemplate(w, "allgame.html", results)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func login(w http.ResponseWriter, r *http.Request) {
    err := templates.ExecuteTemplate(w, "login.html", "")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
