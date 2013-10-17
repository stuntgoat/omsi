package om

import (
	"fmt"
	"errors"
)

// Link holds the information for a doubly linked list
// which holds the key as a string, the
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

func New() *OrderedMap {
	om := OrderedMap{
		Map: make(map[string]*Link),
	}
	return &om
}

// Set sets the `key` to the `value`.
func (om *OrderedMap) Set(key string, value interface{}) {
	retrievedLink, ok := om.Map[key]

	// add non existent key to map
	if !ok {
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
		return
	}
	retrievedLink.value = value
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

// Pop removes and returns the last value,  added to the OrderedMap
func (om *OrderedMap) Pop() (key string, value interface{}, err error) {
	if om.endLink == nil {
		return "", "", errors.New("OrderedMap is empty")
	}
	key, value = om.endLink.key, om.endLink.value
	om.Delete(key)
	return key, value, nil
}

// func (om *OrderedMap) SliceValues(start, stop, skip int) []string

// func (om *OrderedMap) SliceKeys(start, stop, skip int) []string

// func (om *OrderedMap) Keys()

// func (om *OrderedMap) Values()

// func (om *OrderedMap) Items() (key, value string)
