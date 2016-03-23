package front
//package main

import (
    "fmt"
//    "errors"
    "github.com/ziutek/mymysql/mysql"
    _ "github.com/ziutek/mymysql/native" // Native engine
//    "reflect"
)

type logInfo struct {
    Id int `json:"id"`
    Lid string `json:"lid"`
    Logname string `json:logname"`
    Game string `json:"game"`
    Thread string `json:"thread"`
    Loglevel string `json:"loglevel"`
    Logtime string `json:"logtime"`
    Loginfo string `json:"loginfo"`
}

const (
    defaultDbUrl = "172.17.0.1:3306"
    defaultDbUser = "gameops"
    defaultDbPwd = "work@Mokylin"
)

//func (dbName, dbUser, dbPwd string) error {
// func main() {
//     r := getInfo()
//     fmt.Println(r)
//}
//func getinfo() []mysql.Row {
func getInfo() []logInfo{
    db := mysql.New("tcp", "", defaultDbUrl, defaultDbUser, defaultDbPwd, "mysite")
    db.Register("set names utf8")
    err := db.Connect()
    if err != nil {
        // log.Printf("connect db 失败: %s\n", err)
        fmt.Print("connect db 失败: %s\n", err)
        fmt.Println(err)
    }
    defer db.Close()
//    stmt, err := db.Prepare("select * from mysite.polls_datasave limit 10")
//    stmt, err := db.Prepare("select * from polls_datasave where loglevel = 'ERROR' and id = 4682")
    stmt, err := db.Prepare("select * from polls_datasave where loglevel = 'ERROR' limit 100")
    if err != nil {
        fmt.Println(err)
    }
    info := []logInfo{}
    rows, _, err := stmt.Exec()
    for _, row := range rows {
        inf := logInfo{
           Id:          row.Int(0),
           Lid:         row.Str(1),
           Logname:     row.Str(2),
           Game:        row.Str(3),
           Thread:      row.Str(4),
           Loglevel:    row.Str(5),
           Logtime:     row.Str(6),
           Loginfo:     row.Str(7),
        }

        info = append(info, inf)
//      val1 := row[1].([]byte)
//      fmt.Println(val1)
//      fmt.Println(row.Int(0))
//      fmt.Println(row.Str(1))
//      fmt.Println(row.Str(2))
//      fmt.Println(row.Str(3))
//      fmt.Println(row.Str(4))
//      fmt.Println(row.Str(5))
//      fmt.Println(row.Str(6))
//      fmt.Println(row.Str(7))
    }
//    fmt.Println(rows)
//    fmt.Println(res.Map("loglevel"))
//    fmt.Println(res.Map("game"))
//    fmt.Println("Type:", reflect.TypeOf(rows))
//    fmt.Println("Type:", reflect.TypeOf(res))
    return info
//    a = "t"
//    return a
//    _, _ , err = stmt.Exec()
//    return nil
}

