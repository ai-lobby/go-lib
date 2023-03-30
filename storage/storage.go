package storage

type Storager interface {
	Add(string, any) error
	Get(string, int) (any, error)
	Update(string, int, any) error
	Delete(string, int) error
}
