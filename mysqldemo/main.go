package main

import (
  "fmt"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

func connect() *sql.DB {
  db, err := sql.Open("mysql", string("root@tcp(127.0.0.1:3306)/researchdb"))
  if err != nil {
    panic(err)
  }
  err = db.Ping() // test connection
  if err != nil {
    panic(err.Error())
  }
  fmt.Println("connected")
  return db
}

func close(db *sql.DB){
  fmt.Println("close")
  db.Close()
}

func insertData(db *sql.DB){
  // prepare development
  stmt, err := db.Prepare("insert into sensordevice (deviceid, temperature, humidity, ight_intensity, pressure) value (?,?,?,?,?)")
  if err != nil {
    panic(err)
  }
  defer stmt.Close()

  fmt.Println("inserting data...")
  for i := 0; i < 10; i++ {
    _, err = stmt.Exec(2, 0.2*float64(i), 0.6*float64(i), 0.5*float64(i),
                          0.3*float64(i))
    if err != nil {
      panic(err)
    }
  }
  fmt.Println("done")
}

func testConnection() {
  // change database user and password
  db, err := sql.Open("mysql", string("root@tcp(127.0.0.1:3306)/researchdb"))
  if err != nil {
    panic(err)
  }
  err = db.Ping() // test connection
  if err != nil {
    panic(err.Error())
  }
  fmt.Println("connected")
  defer db.Close()
}

func main(){
  // testConnection()
  db := connect()
  insertData(db)
  close(db)
}
