package main

import "encoding/json"
import "fmt"
import "net/http"
import "io/ioutil"

type Mahasiswa struct {
  ID      int
  Nama    string
  Jurusan string
}

var mahasiswa []Mahasiswa

func main() {
  resp, err := http.Get("http://localhost:12345/Mahasiswa")
  if err != nil {
	   // handle error
   }
   defer resp.Body.Close()
   body, err := ioutil.ReadAll(resp.Body)
   err = json.Unmarshal(body, &mahasiswa)
   if err != nil {
     fmt.Println("error:", err)
   }
   fmt.Printf("%+v", mahasiswa)

}
