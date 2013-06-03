package diskstore

import "testing"

func TestNewStore(t *testing.T) {
	var store = NewStore("tests")

	if store.Dir != "tests" {
		t.Error("Directory not set")
		return
	}
}

func TestGet(t *testing.T) {
	var store = NewStore("tests")
	var file, err = store.Get("foo")
	if err != nil {
		t.Error("Cannot find foo")
	}

	var stat, e = file.Stat()
	if e != nil {
		t.Errorf("Error on stat %s", e)
	}

	if stat.Name() != "foo.jpg" {
		t.Errorf("File %s returned instead of %s", stat.Name(), "foo")
	}
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
