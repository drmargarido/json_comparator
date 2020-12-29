package comparator

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"reflect"
)

/*JSONObject refresent a json object format*/
type JSONObject map[string]interface{}

func areValuesSimilar(val interface{}, val2 interface{}) bool {
	if val == nil || val2 == nil {
		return val == val2
	}

	kind := reflect.TypeOf(val).Kind()
	kind2 := reflect.TypeOf(val2).Kind()
	if kind != kind2 {
		return false
	}

	switch kind {
	case reflect.String, reflect.Float64, reflect.Bool:
		return val == val2
	case reflect.Slice:
		v1 := val.([]interface{})
		v2 := val2.([]interface{})
		return areOrderdArraysSimilar(v1, v2)
	case reflect.Map:
		v1 := val.(map[string]interface{})
		v2 := val2.(map[string]interface{})
		return areObjectsSimilar(v1, v2)
	default:
		log.Println("Not handled type:", kind)
		return false
	}
}

func areOrderdArraysSimilar(arr []interface{}, arr2 []interface{}) bool {
	if len(arr) != len(arr2) {
		return false
	}

	for i, obj := range arr {
		if !areValuesSimilar(obj, arr[i]) {
			return false
		}
	}

	return true
}

/* Json objects may have the keys in different order */
func areObjectsSimilar(obj JSONObject, obj2 JSONObject) bool {
	if len(obj) != len(obj2) {
		return false
	}

	for key, val := range obj {
		val2, exists := obj2[key]
		if !exists || !areValuesSimilar(val, val2) {
			return false
		}
	}

	return true
}

/* Both array can be in different orders and are filled with json objects */
func areUnorderedArraysSimilar(arr1 []JSONObject, arr2 []JSONObject) bool {
	if len(arr1) != len(arr2) {
		return false
	}

  canIgnore := map[int]bool{}
NEXT_OBJECT:
	for _, obj1 := range arr1 {
		for i, obj2 := range arr2 {
		  if canIgnore[i] {
		    continue
		  }
			if areObjectsSimilar(obj1, obj2) {
				canIgnore[i] = true
				continue NEXT_OBJECT
			}
		}
		return false
	}

	return true
}

/*AreSimilar checks if two json files have the same data in any order*/
func AreSimilar(file1 string, file2 string) (bool, error) {
	f1, err := ioutil.ReadFile(file1)
	if err != nil {
		return false, err
	}

	f2, err := ioutil.ReadFile(file2)
	if err != nil {
		return false, err
	}

	var data1 []JSONObject
	err = json.Unmarshal(f1, &data1)
	if err != nil {
		return false, err
	}

	var data2 []JSONObject
	err = json.Unmarshal(f2, &data2)
	if err != nil {
		return false, err
	}

	areSimilar := areUnorderedArraysSimilar(data1, data2)
	return areSimilar, nil
}
