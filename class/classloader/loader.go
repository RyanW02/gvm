package classloader

import "github.com/RyanW02/gvm/class"

type ClassLoader struct {
	Classes []class.Class
}

func NewClassLoader() *ClassLoader {
	return &ClassLoader{
	}
}
