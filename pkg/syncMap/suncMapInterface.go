package syncMap

type SyncMap interface {
	Load(key string) int64
	Increment(key string)
	Save(key string, v int64)
}
