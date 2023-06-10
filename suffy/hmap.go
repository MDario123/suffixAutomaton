package suffy

type bucket struct {
	Key   rune
	Value int
}

// Hash table implementation
type hMap struct {
	buckets      []*bucket
	bucketAmount int32
}

// Len returns the len of the underlying slice
func (h *hMap) Len() int32 { return h.bucketAmount }

// Cap returns the cap of the underlying slice
func (h *hMap) Cap() int32 { return int32(cap(h.buckets)) }

// Resize resizes the underlying slice to the given size or does nothing if the given size is smaller than the amount of elements
func (h *hMap) Resize(size int32) {
	//
	if size < h.Len() {
		return
	}

	newMap := &hMap{make([]*bucket, size), h.bucketAmount}

	if h.Len() != 0 {
		for i := int32(0); i < h.Len(); i++ {
			if h.buckets[i] != nil {
				newMap.Insert(h.buckets[i].Key, h.buckets[i].Value)
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

	pos := key % h.Cap()
	element := &bucket{key, value}

	for ; h.buckets[pos] != nil; pos = (pos + 1) % h.Cap() {
		if h.buckets[pos].Key == element.Key {
			break
		} else if dist(element.Key, pos, h.Cap()) > dist(h.buckets[pos].Key, pos, h.Cap()) {
			element, h.buckets[pos] = h.buckets[pos], element
		}
	}

	h.buckets[pos] = element
	h.bucketAmount++

	return
}

func dist(key rune, pos int32, size int32) int32 {
	return (pos + size - key) % size
}

// Get the value of the element with the given key, and a bool to distinguish between value 0 and non-existing element
func (h *hMap) Get(key rune) (int, bool) {
	if h.Len() == 0 {
		return 0, false
	}
	pos := key % h.Cap()
	opos := pos

	for h.buckets[pos] != nil {
		if h.buckets[pos].Key == key {
			return h.buckets[pos].Value, true
		}
		pos = (pos + 1) % h.Cap()
		if pos == opos {
			break
		}
	}

	return 0, false
}

// Copy returns a deep copy of this map
func (h *hMap) Copy() hMap {
	newBuckets := make([]*bucket, h.Cap())
	copy(newBuckets, h.buckets)
	return hMap{newBuckets, h.bucketAmount}
}
