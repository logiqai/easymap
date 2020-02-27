package easymap

import (
	"encoding/json"
	"testing"
)

var (
	sampleData = `
{
	"attr1":"value1",
	"attr2_list":[
	{
		"key":"val1"
	},
	{
		"key":"val2"
	},
	{
		"key":"val3",
		"attr3":"foo"
	},
	{
		"key":"val4",
		"attr3":{
			"attr4":"foo"
		}
	}
	],
	"attr3":"bar",
	"attr4":"bar1"
}
`
)

func TestEasyMap_Get(t *testing.T) {
	m := EasyMap{}
	json.Unmarshal(([]byte)(sampleData), &m)

	t.Run("single_map_key", func(t*testing.T){
		attr1 := m.Get("attr1")
		if len(attr1) != 1 {
			t.Error("Expected 1 entry in response, found : ", len(attr1))
		}
		for k, v := range attr1 {
			if k != "attr1" {
				t.Error("Found unexpected key : ", k)
			}
			if k == "attr1" && v != "value1" {
				t.Error("Found unexpected value: ", v)
			}
		}
	})

	t.Run("ArrayMaps", func(t *testing.T){
		ids := m.Get("key")
		if len(ids) != 4 {
			t.Error("Expected 4 entries in response, found : ", len(ids))
		}
		for k, v := range ids {
			if k == "attr2_list.0.key" && v != "val1" {
				t.Error("Found unexpected value: ", v)
			}
			if k == "attr2_list.1.key" && v != "val2" {
				t.Error("Found unexpected value: ", v)
			}
			if k == "attr2_list.2.key" && v != "val3" {
				t.Error("Found unexpected value: ", v)
			}
			if k == "attr2_list.3.key" && v != "val4" {
				t.Error("Found unexpected value: ", v)
			}
			if k != "attr2_list.0.key" && k != "attr2_list.1.key" && k != "attr2_list.2.key" && k != "attr2_list.3.key" {
				t.Error("Found unexpected key : ", k)
			}
		}
	})

	t.Run("MultipleKeysNested", func(t*testing.T){
		addrs := m.Get("attr4")
		if len(addrs) != 2 {
			t.Error("Expected 6 entries in response, found : ", len(addrs))
		}
		for k, v := range addrs {
			if k == "attr4" && v != "bar1" {
				t.Error("Found unexpected value: ", v)
			}
			if k == "attr2_list.3.attr3.attr4" && v != "foo" {
				t.Error("Found unexpected value: ", v)
			}
			if k != "attr4" && k != "attr2_list.3.attr3.attr4" {
				t.Error("Found unexpected key : ", k)
			}
		}
	})
}
