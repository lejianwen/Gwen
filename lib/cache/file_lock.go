//+build !windows

package cache

func (c *FileCache) Lock() {
	//syscall.Flock(int(f.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)
}
