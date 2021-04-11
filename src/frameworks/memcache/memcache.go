package memcache

type (
	Memcacher interface {
		Increment(key string) (int, error)
	}
)
