package storage

type Storage interface {
	Load(key string) int64
	Increment(key string)
	Save(key string, v int64)
	ReadFromFile(fileMap string) error
	WriteToFile(fileMap string) error
}
