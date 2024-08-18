package goraii

import "os"

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
