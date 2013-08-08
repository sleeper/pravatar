package diskstore

import (
//  "bytes"
  "io/ioutil"
  "testing" 
  "os"
)

func TestNewStore(t *testing.T) {
	var store = NewStore("tests")

	if store.Dir != "tests" {
		t.Error("Directory not set")
		return
	}
}

func TestGet(t *testing.T) {
	var store = NewStore("tests")
	var reader, err = store.Get("foo")

	if err != nil {
		t.Error("Cannot find foo")
	}

  var b []byte
  b, _ = ioutil.ReadAll(reader)
  var content string

  content = string(b)

  if content != "Not a real JPG\n\n" {
    t.Errorf("Error in file content: %s not equal to %s", content, "Not a real JPG\n\n")
  }
}

func TestPut(t *testing.T) {
	var store = NewStore("tests")
	var err = store.Put("tst.jpg", []byte("file content"))
	if err != nil {
		t.Error("Cannot create file tst.jpg")
	}

	var _, e = os.Open("tst.jpg")
	if e != nil {
    t.Errorf("%s", e)
	}

  // Remove created file
  os.Remove("tst.jpg")
}

func TestGetWithExtension(t *testing.T) {
	var store = NewStore("tests")
	var file, err = store.Get("foo.jpg")
	if err != nil {
		t.Error("Cannot find foo.jpg")
	}

	var stat, e = file.Stat()
	if e != nil {
		t.Errorf("Error on stat %s", e)
	}

	if stat.Name() != "foo.jpg" {
		t.Errorf("File %s returned instead of %s", stat.Name(), "foo")
	}
}
