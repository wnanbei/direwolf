package direwolf

import (
	"testing"
)

func TestHeaders(t *testing.T) {
	headers := NewHeaders(
		"key1", "value1",
		"key2", "value2",
	)
	if headers.Get("key1") != "value1" {
		t.Fatal("Headers.Get() failed.")
	}
	if headers.Get("key3") != "" {
		t.Fatal("Headers.Get() failed.")
	}
}

func TestStringSliceMap(t *testing.T) {
	params := NewParams(
		"key1", "value1",
		"key2", "value2",
	)
	if params.Get("key1") != "value1" {
		t.Fatal("params.Get() failed.")
	}
	if params.Get("key3") != "" {
		t.Fatal("params.Get() failed.")
	}

	if params.URLEncode() != "key1=value1&key2=value2" {
		t.Fatal("params.URLEncode() failed.")
	}

	params.Add("key1", "value3")
	if params.Get("key1", 1) != "value3" {
		t.Fatal("params.Add() failed.")
	}

	params.Set("key1", "value4")
	if params.Get("key1") != "value4" {
		t.Fatal("params.Set() failed.")
	}

	params.Del("key2")
	if params.Get("key2") != "" {
		t.Fatal("params.Del() failed.")
	}

	postform := NewPostForm(
		"key1", "value1",
		"key2", "value2",
	)
	if postform.Get("key1") != "value1" {
		t.Fatal("postform.Get() failed.")
	}
}
