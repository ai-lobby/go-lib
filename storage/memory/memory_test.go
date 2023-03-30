package memory

import (
	"testing"
)

func TestMemoryDB(t *testing.T) {
	type tableSchema struct {
		name string
		age  int
	}

	db := NewMemoryDB[tableSchema]()
	db.NewTable("tableA")
	db.Add("tableA", tableSchema{name: "abc", age: 12})

	item, err := db.Get("tableA", 0)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", item)
}
