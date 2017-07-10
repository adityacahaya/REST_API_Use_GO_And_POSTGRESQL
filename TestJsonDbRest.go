package main

import "encoding/json"
import "fmt"
import "database/sql"
import _"github.com/lib/pq"
import "log"
import "net/http"
//import "github.com/gorilla/mux"
//import "io/ioutil"

const (
  DB_USER     = "postgres"
  DB_PASSWORD = "123456"
  DB_NAME     = "DataMahasiswa"
)

type Mahasiswa struct {
  ID      int
  Nama    string
  Jurusan string
}

/* fungsi-fungsi REST API ---------------------------------------------------*/
func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
  db := connectToDb()
  mahasiswa := readFromDb(db)
  defer db.Close()

  j,_ := json.Marshal(mahasiswa)
  w.WriteHeader(http.StatusOK)
  w.Write(j)
}
/* ------------------------------------------------------------------------- */

func main() {
  http.HandleFunc("/Mahasiswa", GetPeopleEndpoint)
  log.Fatal(http.ListenAndServe(":12345", nil))
  /*
  resp, err := http.Get("http://localhost:12334/Mahasiswa")
  checkErr(err)
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)

  dataMahasiswa := unmarshalData(body)
	fmt.Printf("%+v", dataMahasiswa, "\n")


  fmt.Println("Bentuk Data ke JSON")
	b, err := json.Marshal(mahasiswa)
	checkErr(err)
  os.Stdout.Write(b)
  fmt.Println("\n")

  fmt.Println("Dari JSON bentuk ke Struct")
  var dataMahasiswa []Mahasiswa
	err := json.Unmarshal(b, &dataMahasiswa)
	checkErr(err)
	fmt.Printf("%+v", dataMahasiswa)
  */
}

/* fungsi-fungsi Database ---------------------------------------------------*/
func connectToDb()(*sql.DB){
  fmt.Println("Initialize DB")
  dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
      DB_USER, DB_PASSWORD, DB_NAME)
  db, err := sql.Open("postgres", dbinfo)
  checkErr(err)
  return db
}

func readFromDb(db *sql.DB)([]Mahasiswa){
  var mahasiswa []Mahasiswa

  fmt.Println("Read Data From DB")
  rows, err := db.Query("SELECT * FROM userinfo")
  checkErr(err)

  for rows.Next() {
      var uid int
      var nama string
      var jurusan string
      err = rows.Scan(&uid, &nama, &jurusan)
      checkErr(err)
      mahasiswa = append(mahasiswa, Mahasiswa{ID: uid, Nama: nama,
        Jurusan: jurusan})
  }
  fmt.Println("Read Data From DB Succes\n")

  return mahasiswa
}
/* ------------------------------------------------------------------------- */

func checkErr(err error) {
  if err != nil {
      panic(err)
  }
}
