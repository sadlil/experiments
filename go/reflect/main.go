package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"reflect"
)

const data = `apiVersion: v1
kind: Pod
metadata:
  creationTimestamp: null
  labels:
    chart: '{{.Chart.Name}}-{{.Chart.Version}}'
    heritage: '{{.Release.Service}}'
    release: '{{.Release.Name}}'
  name: '{{.Release.Name}}-{{.Values.Name}}'
spec:
  containers:
  - image: '{{.Values.myfrontend.image}}:{{.Values.myfrontend.imageTag}}'
    imagePullPolicy: '{{.Values.myfrontend.imagePullPolicy}}'
    name: myfrontend
    resources: {}
    volumeMounts:
    - mountPath: /var/www/html
      name: mypd
  serviceAccountName: ""
  volumes: []
status: {}
serviceAccountName5: null
`

func main() {
	var mp map[string]interface{}
	yaml.Unmarshal([]byte(data), &mp)

	for k, v := range mp {
		omitEmpty(mp, k, v)
	}

	for k, v := range mp {
		fmt.Println(k, v)
	}
}

func omitEmpty(mp map[string]interface{}, k string, v interface{}) {
	if reflect.ValueOf(v).Kind() == reflect.Ptr {
		v = reflect.ValueOf(v).Elem()
	}

	switch v.(type) {
	case string:
		if v == "" {
			delete(mp, k)
		}
	default:
		if isEmptyValue(reflect.ValueOf(v)) {
			delete(mp, k)
		} else if !reflect.ValueOf(v).IsValid() {
			delete(mp, k)
		} else if reflect.ValueOf(v).Kind() == reflect.Map || reflect.ValueOf(v).Kind() == reflect.Struct {
			data, err := json.Marshal(reflect.ValueOf(v).Interface())
			if err == nil {
				var newMap map[string]interface{}
				json.Unmarshal(data, newMap)
			}
		}
	}
}

func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return false
}

//func isZero(item interface{}) bool {
//	v := reflect.ValueOf(item)
//	switch v.Kind() {
//	case reflect.Func:
//		return v.IsNil()
//	case reflect.Map, reflect.Slice:
//
//	case reflect.Array:
//		z := true
//		for i := 0; i < v.Len(); i++ {
//			z = z && isZero(v.Index(i))
//		}
//		return z
//	case reflect.Struct:
//		z := true
//		for i := 0; i < v.NumField(); i++ {
//			if v.Field(i).CanSet() {
//				z = z && isZero(v.Field(i))
//			}
//		}
//		return z
//	case reflect.Ptr:
//		return isZero(reflect.Indirect(v))
//	}
//	// Compare other types directly:
//	z := reflect.Zero(v.Type())
//	result := v.Interface() == z.Interface()
//
//	return result
//}
