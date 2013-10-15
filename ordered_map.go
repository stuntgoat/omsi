package main

import (
	"fmt"
)



type OrderedMap struct {
	// maps a key to a value
	Map map[string]string

	// maps a key to an index in the KeyOrderArray
	KeyOrderMap map[string]int

	// the order of keys added to Map
	KeyOrderArray []string
	count int
}

func New() *OrderedMap {
	om := OrderedMap{
		Map: make(map[string]string),
		KeyOrderMap: make(map[string]int),
		KeyOrderArray: make([]string, 0),
	}
	return &om
}

// Set sets the `key` to the `value`.
func (om *OrderedMap) Set(key, value string) {
	_, ok := om.Map[key]

	// add non-existent key
	if !ok {
		om.KeyOrderMap[key] = om.count
		om.KeyOrderArray = append(om.KeyOrderArray, key)
		om.count += 1
	}
	om.Map[key] = value

}

// Get retrieves the value from om.Map if it exists
func (om *OrderedMap) Get(key string) (string, bool) {
	val, exists := om.Map[key]
	return val, exists
}

// Delete removes a key/value from the OrderedMap. If it does not exist
// in the map it is a no-op.
func (om *OrderedMap) Delete(key string) {
	i, ok := om.KeyOrderMap[key]
	if !ok {
		return
	}
	om.KeyOrderArray[i] = ""

	delete(om.KeyOrderMap, key)

	delete(om.Map, key)
}

func main() {
	o := New()

	o.Set("cat", "kitty")
	o.Set("dog", "doggy")
	o.Set("seal", "nails")

	val, ok := o.Get("cat")
	fmt.Println("val", val)
	fmt.Println("ok", ok)

	val, ok = o.Get("dog")
	fmt.Println("val", val)
	fmt.Println("ok", ok)

	val, ok = o.Get("whale")
	fmt.Println("val", val)
	fmt.Println("ok", ok)

	o.Delete("cat")
	o.Delete("dog")

	o.Set("chillins", "beer")
	o.Set("oxygen", "ladies")
	o.Delete("seal")

	fmt.Printf("%+v\n", o)
}
