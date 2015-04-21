package utils

import (
	"errors"
	"fmt"
	"sync"
)

// Provides a safe component hash
type ComponentHash struct {
	currId int64
	data   map[int64]interface{}
	mu     sync.RWMutex
}

// Create a new ComponentHash
func NewComponentHash() *ComponentHash {
	v := ComponentHash{
		currId: 0,
		data:   make(map[int64]interface{}),
	}

	return &v
}

// Length.
func (c *ComponentHash) Length() int {
	return len(c.data)
}

// Adds a new element to the hash and return the id of the element
func (c *ComponentHash) Add(el interface{}) int64 {
	c.mu.Lock()
	c.currId++
	c.data[c.currId] = el
	c.mu.Unlock()

	return c.currId
}

// Gets an element
func (c *ComponentHash) Get(id int64) (interface{}, error) {
	c.mu.Lock()
	if err := c.checkElementExists(id); err != nil {
		return nil, err
	}
	el := c.data[id]
	c.mu.Unlock()
	return el, nil
}

// Drops an element by id
func (c *ComponentHash) Drop(id int64) error {
	c.mu.Lock()
	if err := c.checkElementExists(id); err != nil {
		return err
	}
	delete(c.data, id)
	c.mu.Unlock()
	return nil
}

// Checks if the given id is in the hash. Calls to this function
// should be guarded by a mutex.
func (c *ComponentHash) checkElementExists(id int64) error {
	if _, ok := c.data[id]; !ok {
		msg := fmt.Sprintf("id %d was not found in the hash", id)
		return errors.New(msg)
	}
	return nil
}
