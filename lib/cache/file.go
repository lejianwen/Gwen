package cache

import (
	"crypto/md5"
	"fmt"
	"os"
	"sync"
	"time"
)

type FileCache struct {
	mu    sync.Mutex
	locks map[string]*sync.Mutex
	Dir   string
}

func (fc *FileCache) getLock(key string) *sync.Mutex {
	fc.mu.Lock()
	defer fc.mu.Unlock()
	if fc.locks == nil {
		fc.locks = make(map[string]*sync.Mutex)
	}
	if _, ok := fc.locks[key]; !ok {
		fc.locks[key] = new(sync.Mutex)
	}
	return fc.locks[key]
}

func (c *FileCache) Get(key string, value interface{}) error {
	f := c.fileName(key)
	fileInfo, err := os.Stat(f)
	if err != nil {
		//文件不存在
		return nil
	}
	difT := time.Now().Sub(fileInfo.ModTime())
	if difT >= 0 {
		os.Remove(f)
		return nil
	}
	data, err := os.ReadFile(f)
	if err != nil {
		return err
	}
	err1 := DecodeValue(string(data), value)
	return err1
}

func (c *FileCache) Set(key string, value interface{}, exp int) (bool, error) {
	f := c.fileName(key)

	lock := c.getLock(f)
	lock.Lock()
	defer lock.Unlock()

	str, err := EncodeValue(value)
	if err != nil {
		return false, err
	}

	err = os.WriteFile(f, ([]byte)(str), 0644)
	if err != nil {
		return false, err
	}
	if exp <= 0 {
		exp = MaxTimeOut
	}
	expFromNow := time.Now().Add(time.Duration(exp) * time.Second)
	err = os.Chtimes(f, expFromNow, expFromNow)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (c *FileCache) SetDir(path string) {
	c.Dir = path
}

func (c *FileCache) fileName(key string) string {
	f := c.Dir + string(os.PathSeparator) + fmt.Sprintf("%x", md5.Sum([]byte(key)))
	return f
}

func (c *FileCache) Gc() error {
	//检查文件过期时间，并删除
	return nil
}

func NewFileCache() *FileCache {
	return &FileCache{
		locks: make(map[string]*sync.Mutex),
		Dir:   os.TempDir(),
	}
}
