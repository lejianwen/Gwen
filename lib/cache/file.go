package cache

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

var Dir string = os.TempDir()

type FileCache struct {
}

func (c *FileCache) Get(key string, value interface{}) error {
	f := c.fileName(key)

	fileInfo, err := os.Stat(f)
	if err != nil {
		return err
	}
	difT := time.Now().Sub(fileInfo.ModTime())
	if difT >= 0 {
		os.Remove(f)
		return err
	}
	data, err := ioutil.ReadFile(f)
	if err != nil {
		return err
	}
	err1 := c.DecodeValue(string(data), value)
	return err1
}

func (c *FileCache) Set(key string, value interface{}, exp int) (bool, error) {
	f := c.fileName(key)
	str, err := c.EncodeValue(value)
	if err != nil {
		return false, err
	}
	// todo 给文件加锁，解锁， 懒得搞了
	//c.Lock()

	err = ioutil.WriteFile(f, ([]byte)(str), 0644)
	if err != nil {
		return false, err
	}
	if exp <= 0 {
		exp = MaxTimeOut
	}
	twoDaysFromNow := time.Now().Add(time.Duration(exp) * time.Second)
	err = os.Chtimes(f, twoDaysFromNow, twoDaysFromNow)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (c *FileCache) EncodeValue(value interface{}) (string, error) {
	js, err := json.Marshal(value)
	if err != nil {
		return "", err
	}
	return string(js), err
}
func (c *FileCache) DecodeValue(value string, rtv interface{}) error {
	err := json.Unmarshal(([]byte)(value), rtv)
	return err
}
func (c *FileCache) SetDir(path string) {
	Dir = path
}

func (c *FileCache) fileName(key string) string {
	f := Dir + string(os.PathSeparator) + fmt.Sprintf("%x", md5.Sum([]byte(key)))
	return f
}

func (c *FileCache) Gc() error {
	//检查文件过期时间，并删除
	return nil
}
