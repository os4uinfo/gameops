package front
// package main

import (
    "fmt"
//    "errors"
    "github.com/ziutek/mymysql/mysql"
    _ "github.com/ziutek/mymysql/native" // Native engine
//    "reflect"
)

const (
    defaultDbUrl = "172.17.0.1:3306"
    defaultDbUser = "gameops"
    defaultDbPwd = "work@Mokylin"
)
var (
    a string
)

//func (dbName, dbUser, dbPwd string) error {
// func main() {
func getinfo() []mysql.Row {
    db := mysql.New("tcp", "", defaultDbUrl, defaultDbUser, defaultDbPwd, "mysite")
    err := db.Connect()
    if err != nil {
        // log.Printf("connect db 失败: %s\n", err)
        fmt.Print("connect db 失败: %s\n", err)
        fmt.Println(err)
    }
    defer db.Close()
    stmt, err := db.Prepare("select * from mysite.polls_datasave limit 10")
    if err != nil {
        fmt.Println(err)
    }
    rows, _, err := stmt.Exec()
//  for _, row := range rows {
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
//  }
//    fmt.Println(rows)
//    fmt.Println(res.Map("loglevel"))
//    fmt.Println(res.Map("game"))
//    fmt.Println("Type:", reflect.TypeOf(rows))
//    fmt.Println("Type:", reflect.TypeOf(res))
    return rows
//    a = "t"
//    return a
//    _, _ , err = stmt.Exec()
//    return nil
}

