package infrastructure

type StorageClient interface {
	Save(key string, doc string) (string, error)
}
