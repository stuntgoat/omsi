package main

import (
	"fmt"
	"reflect"
)

type OrderedMap struct {
	// maps a key to a value
	Map map[interface{}]interface{}

	// maps a key to an index in the KeyOrderArray
	KeyOrderMap map[interface{}]int

	// the order of keys added to Map
	KeyOrderArray []interface{}

	// holds the data type for the key/values
	keyType interface{}
	valueType interface{}

	ValueFromKey *func(interface{})interface{}
}

// validateTypes sets the types if the Map is empty, otherwise
// insures that the key and value are of the correct types
// for operations.
func (om *OrderedMap) validateTypes(key, value interface{}) (err error) {
//	if len(om.Map) == 0 {
		om.keyType = reflect.TypeOf(key)
		om.valueType = reflect.TypeOf(value)
		om.Map = make(map[interface{}]interface{}, 0)
		om.KeyOrderArray = make([]interface{}, 0)
		om.KeyOrderMap = make(map[interface{}]int, 0)
	// } else {

	// }

	return nil
}

// Set sets the `key` to the `value`.
func (om *OrderedMap) Set(key, value interface{}) {
	err := om.validateTypes(key, value)
	if err != nil {
		panic(err.Error())
	}

	value, ok := om.Map[key]

	// add non-existent key
	if !ok {
		om.KeyOrderMap[key] = len(om.Map)
		om.KeyOrderArray = append(om.KeyOrderArray, key)
	}

	om.Map[key] = value
}

// Get retrieves the value from om.Map if it exists
func (om *OrderedMap) Get(key interface{}) (interface{}, bool) {
	val, exists := om.Map[key]
	return val, exists
}

func InterfaceToInt(i interface{}) int {
	defer recover(
		func()nil{}()
	)
	val, ok := i.(int)
	if ok {
		return val
	}
	panic("value not converted")
}

func main() {
	o := new(OrderedMap)

	o.Set("cat", 9)
	o.Set("dog", 10)


	val := o.ValueForKey(`cat`)
	fmt.Println("o.Get(`cat`)", val)
}
