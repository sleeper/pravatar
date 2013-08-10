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
  Key string
  Content string
}

func (s* mockStore) Get(hash string) (io.Reader, error) {
  var content string = ""

  if hash == "foobarbaz" {
    content = "this is an image"
  }
  return strings.NewReader(content), nil
}


func (s* mockStore) Put(name string, content []byte) (error) {
  s.Key = name
  s.Content = string(content)
  return nil
}

func newMockStore() *mockStore {
  return &mockStore{Key: "", Content: ""}
}

func TestGetAvatar(t *testing.T) {
  var store = newMockStore()
  var pravatar = NewPravatar("", "3333", store)

  var req, err = http.NewRequest("GET", "/avatar/foobarbaz", nil)
  if err != nil {
    t.Errorf("Cannot create request: %s", err)
  }

  rw := httptest.NewRecorder()
  rw.Body = new(bytes.Buffer)

  pravatar.Router.ServeHTTP(rw, req)

  if rw.Code != 200 {
    t.Errorf("GET /avatar/foobarbaz: code = %d, want %d", rw.Code, 200)
  }

  if rw.Body.String() != "this is an image" {
    t.Errorf("GET /avatar/foobarbaz: body = %q, want \"this is an image\"", rw.Body.String())
  }
}

// FIXME: Test Put without body returns an error

func TestPutAvatar(t *testing.T) {
  var store = newMockStore()
  var pravatar = NewPravatar("", "3333", store)

  var req, err = http.NewRequest("POST", "/avatar/foobarbaz", strings.NewReader("this is an image"))
  if err != nil {
    t.Errorf("Cannot create request: %s", err)
  }

  rw := httptest.NewRecorder()
  rw.Body = new(bytes.Buffer)

  pravatar.Router.ServeHTTP(rw, req)

  if rw.Code != 200 {
    t.Errorf("POST /avatar/foobarbaz: code = %d, want %d", rw.Code, 200)
  }

  if store.Key != "foobarbaz" {
    t.Errorf("POST /avatar/foobarbaz: key = %q, want \"foobarbaz\"", store.Key)
  }

  if store.Content != "this is an image" {
    t.Errorf("POST /avatar/foobarbaz: last stored = %q, want \"this is an image\"", store.Content)
  }
}

