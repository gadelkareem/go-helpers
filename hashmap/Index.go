package hashmap

import (
	"sync"
)

type Index struct {
	sync.RWMutex
	records map[string]interface{}
	path    string
}

func NewIndex(records map[string]interface{}) *Index {
	return &Index{records: records}
}

func (i *Index) Add(id string, v interface{}) {
	i.Lock()
	defer i.Unlock()
	i.records[id] = v
	return
}

func (i *Index) Remove(id string) {
	i.Lock()
	defer i.Unlock()
	delete(i.records, id)
}

func (i *Index) Exists(id string) bool {
	i.Lock()
	defer i.Unlock()
	_, exists := i.records[id]
	return exists
}
