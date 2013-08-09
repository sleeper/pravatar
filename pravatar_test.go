package main

import (
  "bytes"
  "testing"
  "io"
  "net/http"
  "net/http/httptest"
  "strings"
)

type mockStore struct {
  Dir string
}

func (s* mockStore) Get(hash string) (io.Reader, error) {
  var content string = ""

  if hash == "foobarbaz" {
    content = "this is an image"
  }
  return strings.NewReader(content), nil
}


func (s* mockStore) Put(name string, content []byte) (error) {
  return nil
}

func newMockStore() *mockStore {
  return &mockStore{Dir: "dummy"}
}

func TestGetAvatar(t *testing.T) {
  var store = newMockStore()
  var pravatar = NewPravatar("", "3333", store)

  dummy := httptest.NewServer(pravatar.Router)
  defer dummy.Close()

  var req, err = http.NewRequest("GET", "/avatar/foobarbaz", nil)
  if err != nil {
    t.Errorf("Cannot create request: %s", err)
  }

  rw := httptest.NewRecorder()
  rw.Body = new(bytes.Buffer)

//  dummy.ServeHTTP(rw, req)
  pravatar.Router.ServeHTTP(rw, req)

  if rw.Code != 200 {
    t.Errorf("www.example.com/avatar/foobarbaz: code = %d, want %d", rw.Code, 200)
  }

  if rw.Body.String() != "this is an image" {
    t.Errorf("www.example.com/avatar/foobarbaz: body = %q, want \"this is an image\"", rw.Body.String())
  }

//  handler := pravatar.getAvatarHandler()
//
//  req, err := http.NewRequest("GET", "/avatar/foobarbaz", nil)
//  if err != nil {
//    t.Errorf("Cannot create request: %s", err)
//  }
//
//  w := httptest.NewRecorder()
//  handler(w, req)
//  fmt.Printf("FRED ==> %s", w.Body.String())
//
//  if w.Body.String() != "This is an image" {
//    t.Error("Not receiving right reader")
//  }
}

//func TestGet(t *testing.T) {
////  handler := func(w http.ResponseWriter, r *http.Request) {
////
////    w.Header().Set("Content-Type", request.contenttype)
////    io.WriteString(w, request.body)
////  }
////
//
//  var store = newMockStore()
//  var pravatar = NewPravatar("", "3333", store)
//
//  handler := pravatar.getAvatarHandler()
//
//  server := httptest.NewServer(http.HandlerFunc(handler))
//  defer server.Close()
//
//  resp, err := http.Get(server.URL)
//  if err != nil {
//    t.Fatalf("Get: %v", err)
//  }
//  checkBody(t, resp, twitterResponse)
//}
//
//func checkBody(t *testing.T, r *http.Response, body string) {
//  b, err := ioutil.ReadAll(r.Body)
//  if err != nil {
//    t.Error("reading reponse body: %v, want %q", err, body)
//  }
//  if g, w := string(b), body; g != w {
//    t.Errorf("request body mismatch: got %q, want %q", g, w)
//  }
//}
//
//// func main() {
//   handler := func(w http.ResponseWriter, r *http.Request) {
//     http.Error(w, "something failed", http.StatusInternalServerError)
//   }
// 
//   req, err := http.NewRequest("GET", "http://GETexample.com/foo", nil)
//   if err != nil {
//     log.Fatal(err)
//   }
// 
//   w := httptesttest.NewRecorder()
//   handler(w, req)
// 
//   fmt.Printf("%d - %s", w.Code, w.Body.String())
// }

