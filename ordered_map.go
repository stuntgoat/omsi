package main

import (
	"fmt"
)

// Link holds the information for a doubly linked list
// which holds the key as a string, the
type Link struct {
	value string
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
func (om *OrderedMap) Set(key, value string) {
	retrievedLink, ok := om.Map[key]

	// add non existent key to map
	if !ok {
		link := &Link{
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
func (om *OrderedMap) Get(key string) (string, bool) {
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

	if previous != nil {
		previous.next = next
	}

	if retrievedLink == om.endLink {
		om.endLink = previous
	}

	if next != nil {
		next.previous = previous
	}

	delete(om.Map, key)
}



func main() {
	o := New()

	o.Set("cat", "kitty")
	o.Set("dog", "doggy")
	o.Set("seal", "nails")

	fmt.Printf("%+v\n", o)

	val, _ := o.Get("cat")
	fmt.Println("val", val)
	val, _ = o.Get("dog")
	fmt.Println("val", val)
	val, _ = o.Get("seal")
	fmt.Println("val", val)

	o.Delete("seal")
	fmt.Printf("%+v\n", o)

	o.Delete("cat")
	fmt.Printf("%+v\n", o)

	o.Delete("dog")
	fmt.Printf("%+v\n", o)

	o.Set("cat", "kitty")
	o.Set("dog", "doggy")
	o.Set("seal", "nails")

	fmt.Printf("%+v\n", o)

	o.Delete("cat")
	fmt.Printf("%+v\n", o)

	o.Delete("seal")
	fmt.Printf("%+v\n", o)

	o.Delete("dog")
	fmt.Printf("%+v\n", o)

	o.Set("cat", "kitty")
	o.Set("dog", "doggy")
	o.Set("seal", "nails")

	o.Delete("dog")
	fmt.Printf("%+v\n", o)

	o.Delete("cat")
	fmt.Printf("%+v\n", o)

	o.Delete("seal")
	fmt.Printf("%+v\n", o)

}
