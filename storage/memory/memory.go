package memory

import "github.com/ai-lobby/go-lib/storage"

type Memory[T any] struct {
	tables map[string]*Table[T]
}

func NewMemoryDB[T any]() *Memory[T] {
	return &Memory[T]{
		tables: make(map[string]*Table[T]),
	}
}

func (m *Memory[T]) NewTable(t string) {
	m.tables[t] = NewTable[T]()
}

func (m *Memory[T]) Add(t string, v T) error {
	return m.add(t, v)
}

func (m Memory[T]) Get(t string, i int) (T, error) {
	return m.get(t, i)
}

func (m *Memory[T]) Update(t string, i int, v T) error {
	return m.set(t, i, v)
}

func (m *Memory[T]) Delete(t string, i int) error {
	return m.delete(t, i)
}

func (m Memory[T]) getTable(name string) (*Table[T], error) {
	table, ok := m.tables[name]
	if !ok {
		return table, storage.ErrTableNotFound
	}
	return table, nil
}

func (m *Memory[T]) add(t string, v T) error {
	table, err := m.getTable(t)
	if err != nil {
		return err
	}

	table.Add(v)
	return nil
}

func (m *Memory[T]) set(t string, i int, v T) error {
	table, err := m.getTable(t)
	if err != nil {
		return err
	}

	table.Set(i, v)
	return nil
}

func (m Memory[T]) get(t string, i int) (T, error) {
	table, err := m.getTable(t)
	if err != nil {
		return *new(T), err
	}

	v, err := table.Get(i)
	if err != nil {
		return *new(T), err
	}

	return v, nil
}

func (m Memory[T]) delete(t string, i int) error {
	table, err := m.getTable(t)
	if err != nil {
		return err
	}

	table.Delete(i)
	return nil
}
