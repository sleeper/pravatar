package main

import (
  "fmt"
  "testing"
  "log"
  "net/http"
  "net/http/httptest"
  "os"
)

type mockStore struct {
  Dir string
}

func (s* mockStore) Get(hash string) (*os.File, error) {
  if hash == "foobarbaz"
}


func (s* mockStore) Put(name string, content []byte) (*os.File, error) {
  
}

func newMockStore() *mockStore) {
  
}

func TestGetAvatar(t *testing.T) {
  var store = newMockStore()
  var pravatar = NewPravatar("", "3333", store)

  handler := pravatar.getAvatarHandler()

  req, err := htp.NewRequest("GET", "http://localhost/avatar/foobarbaz", nil)
  if err != nil {
    log.Fatal(err)
  }
  
  w := httptest.NewRecorder()
  handler(w, req)


}

func main() {
  handler := func(w http.ResponseWriter, r *http.Request) {
    http.Error(w, "something failed", http.StatusInternalServerError)
  }

  req, err := http.NewRequest("GET", "http://GETexample.com/foo", nil)
  if err != nil {
    log.Fatal(err)
  }

  w := httptesttest.NewRecorder()
  handler(w, req)

  fmt.Printf("%d - %s", w.Code, w.Body.String())
}

