package memstore

import "errors"

var NotFoundErr = errors.New("not found")

type MemStore[T interface{}] struct {
	list map[string]T
}

func New[T interface{}]() *MemStore[T] {
	var m MemStore[T]
	list := make(map[string]T)
	m.list = list
	return &m
}

func (m MemStore[T]) List() (map[string]T, error) {
	return m.list, nil
}

func (m MemStore[T]) Add(name string, val T) error {
	m.list[name] = val
	return nil
}

func (m MemStore[T]) Remove(name string) error {
	delete(m.list, name)
	return nil
}

func (m MemStore[T]) Get(name string) (T, error) {
	if val, ok := m.list[name]; ok {
		return val, nil
	}

	var val T
	return val, NotFoundErr
}

func (m MemStore[T]) Update(name string, new T) error {
	if _, ok := m.list[name]; ok {
		m.list[name] = new
		return nil
	}

	return NotFoundErr
}
