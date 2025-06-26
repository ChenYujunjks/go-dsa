package setmap

type entry struct {
	key   string
	value any
	next  *entry
}

type HashMap struct {
	buckets []*entry
	size    int
	cap     int
}

func NewHashMap(capacity int) *HashMap {
	if capacity <= 0 {
		capacity = 16
	}
	return &HashMap{
		buckets: make([]*entry, capacity),
		size:    0,
		cap:     capacity,
	}
}
