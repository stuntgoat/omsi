package omsi

import (
	"fmt"
	"errors"
)

// Link holds the information for a doubly linked list
// which also holds the key.
type Link struct {
	key string
	value interface{}

	previous *Link
	next *Link
}

type OrderedMap struct {
	// maps a key to a value
	Map map[string]*Link

	startLink *Link
	endLink *Link
}

// New returns an OrderedMap with memory allocated for the Map
func New() *OrderedMap {
	om := OrderedMap{
		Map: make(map[string]*Link),
	}
	return &om
}

// Set sets the `key` to the `value`.
func (om *OrderedMap) Set(key string, value interface{}) {
	retrievedLink, exists := om.Map[key]

	if exists {
		retrievedLink.value = value
		return
	}

	// add non existent key to map
	link := &Link{
		key: key,
		value: value,
	}

	// first item
	if len(om.Map) == 0 {
		om.startLink = link
		om.endLink = link
	} else {
		// get the most recent 'endLink'
		lastEnd := om.endLink

		// set this link's 'previous' link as most recent 'endLink'
		link.previous = lastEnd
		// the link.next is defaulted to nil

		// set this new link as the 'next' link for
		// the om.endLink
		lastEnd.next = link

		// set this new link as the end
		om.endLink = link
	}
	om.Map[key] = link
}

// Get returns the value of the link found by the key
func (om *OrderedMap) Get(key string) (interface{}, bool) {
	retrievedLink, exists := om.Map[key]
	if !exists {
		return "", false
	}

	return retrievedLink.value, true
}

// Delete removes a key/value pair if it exists. If the key does not exist
// the operation is a no-op.
func (om *OrderedMap) Delete(key string) {
	retrievedLink, exists := om.Map[key]
	if !exists {
		return
	}

	previous := retrievedLink.previous
	next := retrievedLink.next

	// if we are deleting the beginning node
	if retrievedLink == om.startLink {

		// set the next node to the startLink
		om.startLink = next
	}

	// if the retrieved link was not the first item
	if previous != nil {
		previous.next = next
	}

	if retrievedLink == om.endLink {
		om.endLink = previous
	}

	// if the retrieved link was not the end link
	if next != nil {
		next.previous = previous
	}

	delete(om.Map, key)
}

// Pop removes and returns the last key, value, added to the OrderedMap, or
// an error.
func (om *OrderedMap) Pop() (key string, value interface{}, err error) {
	if om.endLink == nil {
		return key, value, errors.New("OrderedMap is empty")
	}
	key, value = om.endLink.key, om.endLink.value
	om.Delete(key)
	return key, value, nil
}

// Keys returns a slice of keys in order in which they were added
func (om *OrderedMap) Keys() (keys []string) {
	keys = make([]string, 0)

	var link = om.startLink
	for {
		if link == nil {
			break
		}
		keys = append(keys, link.key)
		link = link.next
	}
	return
}

// Values returns a slice of values in order in which they were added
func (om *OrderedMap) Values() (values []interface{}) {
	values = make([]interface{}, 0)

	var link = om.startLink
	for {
		if link == nil {
			break
		}
		values = append(values, link.value)
		link = link.next
	}
	return
}
