package main

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

type logInfo struct {
    id int `json:"id"`
    lid string `json:"lid"`
    logname string `json:logname"`
    game string `json:"game"`
    thread string `json:"thread"`
    loglevel string `json:"loglevel"`
    logtime string `json:"logtime"`
    loginfo string `json:"loginfo"`
}

func main() {
//   results :=  query()
   query()
//   fmt.Println(results)
}

//插入demo

//查询demo
// func query() map[string]string{
func query() {
    // db, err := sql.Open("mysql", "gameops:work@Mokylin@172.17.0.1/mysite?charset=utf8")
    db, err := sql.Open("mysql", "gameops:work@Mokylin@tcp(172.17.0.1:3306)/mysite?charset=utf8")
    checkErr(err)

    rows, err := db.Query("SELECT * FROM polls_datasave limit 10")
    checkErr(err)


    info := []logInfo{}
    //普通demo
    for rows.Next() {
    //  var id int
    //  var lid string
    //  var logname string
    //  var game string
    //  var thread string
    //  var loglevel string
    //  var logtime string
    //  var loginfo string
        inf := logInfo{}

        rows.Columns()
        // err = rows.Scan(&inf.id, &inf.lid, &logname, &game, &thread, &loglevel, &logtime, &loginfo)
        err = rows.Scan(&inf.id, &inf.lid, &inf.logname, &inf.game, &inf.thread, &inf.loglevel, &inf.logtime, &inf.loginfo)
        checkErr(err)
        info = append(info, inf)
    //  fmt.Println(id)
    //  fmt.Println(lid)
    //  fmt.Println(logname)
    //  fmt.Println(game)
    }
    fmt.Println(info)

    //字典类型
    //构造scanArgs、values两个数组，scanArgs的每个值指向values相应值的地址
  //info := []logInfo{}
  //columns, _ := rows.Columns()
  //scanArgs := make([]interface{}, len(columns))
  //values := make([]interface{}, len(columns))
  //for i := range values {
  //    scanArgs[i] = &values[i]
  //}

  //record := make(map[string]string)
  //for rows.Next() {
  //    //将行数据保存到record字典
  //    err = rows.Scan(scanArgs...)
  //    for i, col := range values {
  //        if col != nil {
  //            record[columns[i]] = string(col.([]byte))
  //        }
  //    }
////      fmt.Println(record)
  //   info = append(info, record)
  //}
////  return record
  //return info
}
//更新数据

//删除数据

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}
