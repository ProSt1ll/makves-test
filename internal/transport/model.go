package transport

type Store interface {
	Get(key string, val []string) ([]byte, error)
}

type Server struct {
	store Store
}
