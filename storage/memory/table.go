package memory

import "github.com/ai-lobby/go-lib/storage"

type Table[T any] struct {
	records []T
}

func NewTable[T any]() *Table[T] {
	return &Table[T]{}
}

func (t *Table[T]) Add(v T) {
	t.records = append(t.records, v)
}

func (t Table[T]) Get(i int) (T, error) {
	if i >= len(t.records) {
		return *new(T), storage.ErrIndexNotFound
	}
	return t.records[i], nil
}

func (t *Table[T]) Set(i int, v T) error {
	if i >= len(t.records) {
		return storage.ErrIndexNotFound
	}
	t.records[i] = v
	return nil
}

func (t *Table[T]) Delete(i int) {
	t.records = append(t.records[:i], t.records[i+1:]...)
}
