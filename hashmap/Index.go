package hashmap

import (
	"sync"
)

type Index struct {
	sync.RWMutex
	Records map[string]interface{}
	path    string
}

func NewIndex(records map[string]interface{}) *Index {
	return &Index{Records: records}
}

func (i *Index) Add(id string, v interface{}) {
	i.Lock()
	defer i.Unlock()
	i.Records[id] = v
	return
}

func (i *Index) Remove(id string) {
	i.Lock()
	defer i.Unlock()
	delete(i.Records, id)
}

func (i *Index) Exists(id string) bool {
	i.Lock()
	defer i.Unlock()
	_, exists := i.Records[id]
	return exists
}

func (i *Index) Get(id string) interface{} {
	i.Lock()
	defer i.Unlock()
	r := i.Records[id]
	return r
}

func (i *Index) AddMultiple(m map[string]interface{}) {
	i.Lock()
	defer i.Unlock()
	for k, r := range m {
		i.Records[k] = r
	}
}

func (i *Index) Map(fn func(record interface{}) bool) {
	i.Lock()
	defer i.Unlock()
	for _, r := range i.Records {
		if fn(r) {
			return
		}
	}
}

func (i *Index) Length() int {
	i.Lock()
	defer i.Unlock()
	return len(i.Records)
}

func (i *Index) Copy() map[string]interface{} {
	i.Lock()
	defer i.Unlock()
	c := make(map[string]interface{})
	for k, v := range i.Records {
		c[k] = v
	}
	return c
}
