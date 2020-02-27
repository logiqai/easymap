package easymap

import "fmt"

type EasyMap map[string]interface{}
type EasyList []interface{}
type NotFound error

func ( e EasyMap) Get(key string) (map[string]interface{}) {
	return e.lookup(key)
}

func (e EasyList) lookup(key string) (map[string]interface{}) {
	retval := map[string]interface{}{}
	for i, ival := range e {
		switch ival.(type) {
		case map[string]interface{}:
			vm, _ := (ival).(map[string]interface{})
			ret := (EasyMap)(vm).lookup(key);
			for k1,v1 := range ret {
					retval[fmt.Sprintf("%d",i)+"."+k1] = v1
					}
		}
	}
	return retval
}

func (e EasyMap) lookup(key string) (map[string]interface{}) {
	retval := EasyMap{}

	for k, v := range e {
		if k == key {
			retval[key]=v
		}else {
			switch v.(type) {
			case map[string]interface{}:
				vm, _ := (v).(map[string]interface{})
				ret := (EasyMap)(vm).lookup(key)
					for k1,v1 := range ret {
						retval[k+"."+k1] = v1
					}

			case []interface{}:
				vl := v.([]interface{})
				ret := (EasyList)(vl).lookup(key)
					for k1,v1 := range ret {
						retval[k+"."+k1] = v1
					}
			}
		}
	}
	return retval
}
