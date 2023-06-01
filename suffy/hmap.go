package suffy

type bucket struct {
	Key   rune
	Value int
}

// Hash table implementation
type hMap struct {
	Buckets []*bucket
}

// Len returns the len of the underlying slice
func (h *hMap) Len() int32 { return int32(len(h.Buckets)) }

// Cap returns the cap of the underlying slice
func (h *hMap) Cap() int32 { return int32(cap(h.Buckets)) }

// Resize resizes the underlying slice to the given size or does nothing if the given size is smaller than the amount of elements
func (h *hMap) Resize(size int32) {
	//
	if size < h.Len() {
		return
	}

	newMap := &hMap{make([]*bucket, size)}

	if h.Len() != 0 {
		for i := int32(0); i < h.Len(); i++ {
			if h.Buckets[i] != nil {
				newMap.Insert(h.Buckets[i].Key, h.Buckets[i].Value)
			}
		}
	}

	*h = *newMap

	return
}

// Insert inserts a new element in the map
func (h *hMap) Insert(key rune, value int) {
	if h.Cap() == h.Len() {
		h.Resize(h.Len()*2 + 1)
	}

	pos := key % h.Len()
	element := &bucket{key, value}

	for ; h.Buckets[pos] != nil; pos = (pos + 1) % h.Len() {
		if h.Buckets[pos].Key == element.Key {
			break
		} else if dist(element.Key, pos, h.Len()) > dist(h.Buckets[pos].Key, pos, h.Len()) {
			element, h.Buckets[pos] = h.Buckets[pos], element
		}
	}

	h.Buckets[pos] = element

	return
}

func dist(key rune, pos int32, size int32) int32 {
	return (pos + size - key) % size
}

func (h *hMap) Get(key rune) (int, bool) {
	if h.Len() == 0 {
		return 0, false
	}
	pos := key % h.Len()
	opos := pos

	for h.Buckets[pos] != nil {
		if h.Buckets[pos].Key == key {
			return h.Buckets[pos].Value, true
		}
		pos = (pos + 1) % h.Len()
		if pos == opos {
			break
		}
	}

	return 0, false
}

func (h *hMap) Copy() hMap {
	newBuckets := make([]*bucket, h.Len())
	copy(newBuckets, h.Buckets)
	return hMap{newBuckets}
}
