package internal

// DataStorage interface to storage key-value data
type DataStorage interface {
	Get(string) (string, bool)
	ToMap() (map[string]string, error)
	FromMap(map[string]string) error
	KeyExists(string) bool
	Commit() error
}
