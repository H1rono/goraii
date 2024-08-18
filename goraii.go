package goraii

import (
	"os"
	"sync"
)

type ResourceManager[T any] interface {
	Enter() T
	Exit(T)
}

func WithResourceManager[T any, R ResourceManager[T]](r R) func(yield func(T) bool) {
	return func(yield func(T) bool) {
		v := r.Enter()
		defer r.Exit(v)
		yield(v)
	}
}

type defaultFileManager struct {
	name string
}

func (f *defaultFileManager) Enter() *os.File {
	file, err := os.Open(f.name)
	if err != nil {
		panic(err)
	}
	return file
}

func (f *defaultFileManager) Exit(file *os.File) {
	err := file.Close()
	if err != nil {
		panic(err)
	}
}

func OpenFile(name string) func(func(*os.File) bool) {
	return WithResourceManager(&defaultFileManager{name})
}

type mutexManager struct {
	mu *sync.Mutex
}

func (m *mutexManager) Enter() struct{} {
	m.mu.Lock()
	return struct{}{}
}

func (m *mutexManager) Exit(struct{}) {
	m.mu.Unlock()
}

func MutexLockGuard(m *sync.Mutex) func(yield func(struct{}) bool) {
	return WithResourceManager(&mutexManager{m})
}
