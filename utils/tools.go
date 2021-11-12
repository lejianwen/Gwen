package utils

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"reflect"
)

func Md5(str string) string {
	t := md5.Sum(([]byte)(str))
	return fmt.Sprintf("%x", t)
}

func CopyStruct(src, dst interface{}) {
	sval := reflect.ValueOf(src).Elem()
	dval := reflect.ValueOf(dst).Elem()

	for i := 0; i < sval.NumField(); i++ {
		value := sval.Field(i)
		f := sval.Type().Field(i)

		dvalue := dval.FieldByName(f.Name)
		if dvalue.IsValid() == false {
			continue
		}
		if value.Kind() == dvalue.Kind() {
			dvalue.Set(value)
		}
	}
}

func CopyStructByJson(src, dst interface{}) {
	str, _ := json.Marshal(src)
	json.Unmarshal(str, dst)
}
