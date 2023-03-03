package transport

//Store интерфейс базы с методом получения данных
type Store interface {
	Get(key string, val []string) ([]byte, error)
}

//Server структура с обработчиками запросов
type Server struct {
	store Store
}
