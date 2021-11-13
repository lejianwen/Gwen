package utils

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
)

func Md5(str string) string {
	t := md5.Sum(([]byte)(str))
	return fmt.Sprintf("%x", t)
}

func CopyStructByJson(src, dst interface{}) {
	str, _ := json.Marshal(src)
	json.Unmarshal(str, dst)
}

//CopyStructToMap 结构体转map
func CopyStructToMap(src interface{}) map[string]interface{} {
	var res = map[string]interface{}{}
	str, _ := json.Marshal(src)
	json.Unmarshal(str, &res)
	return res
}
