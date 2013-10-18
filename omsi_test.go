package omsi

import (
	"fmt"
	"testing"
)

func Test_New(t *testing.T) {
	om := New()
	if om.Map == nil {
		t.Error("map failed to initialize after calling New.")
	}
}

func Test_Set(t *testing.T) {
	om := New()

	om.Set("cat", "funny")

	link := om.Map["cat"]

	if link.value != "funny" {
		t.Error("Set failed to set value")
	}
}

func Test_Get(t *testing.T) {
	om := New()

	om.Set("cat", "funny")
	value, exists := om.Get("cat")
	if exists == false {
		t.Error("Get failed to retrieve value")
	}

	if value != "funny" {
		t.Error("Get failed to get correct value")
	}
}

func Test_Delete(t *testing.T) {
	om := New()

	om.Set("cat", "funny")
	om.Set("dog", "silly")
	om.Set("parrot", "green")

	om.Delete("dog")
	dog, exists := om.Get("dog")
	if exists {
		t.Error("Deleted failed to remove value")
	}
	if dog != nil {
		msg := fmt.Sprintf("Deleted failed to remove value; expected nil, got %v", dog)
		t.Error(msg)
	}

	if om.startLink.next != om.endLink {
		t.Error("Delete failed to associate doubley linked list correctly")
	}

	if om.endLink.previous != om.startLink {
		t.Error("Delete failed to associate doubley linked list correctly")
	}
}

func Test_Pop(t *testing.T) {
	om := New()

	om.Set("cat", "funny")
	om.Set("dog", "silly")
	om.Set("parrot", "green")

	key, value, err := om.Pop()
	if err != nil {
		msg := fmt.Sprintf("Pop returned error: %s", err.Error())
		t.Error(msg)
	}
	if key != "parrot" {
		msg := fmt.Sprintf("Pop failed to yield correct key; expected 'parrot', got: %s", key)
		t.Error(msg)
	}

	if value != "green" {
		msg := fmt.Sprintf("Pop failed to yield correct key; expected 'green', got: %s", value)
		t.Error(msg)
	}
}

func Test_Keys(t *testing.T) {
	om := New()

	om.Set("cat", "funny")
	om.Set("dog", "silly")
	om.Set("parrot", "green")

	expected := [3]string{"cat", "dog", "parrot"}

	keys := om.Keys()

	if len(keys) != len(expected) {
		msg := fmt.Sprintf("Keys returned incorrect number of elements; expected %d got: %d", len(expected), len(keys))
		t.Error(msg)
	}

	for i, key := range keys {
		if key != expected[i] {
			t.Error("Keys returned out of order indexed slice of keys")
		}
	}
}

func Test_Values(t *testing.T) {
	om := New()

	om.Set("cat", "funny")
	om.Set("dog", "silly")
	om.Set("parrot", "green")

	expected := [3]string{"funny", "silly", "green"}

	values := om.Values()

	if len(values) != len(expected) {
		msg := fmt.Sprintf("Values returned incorrect number of elements; expected %d got: %d", len(expected), len(values))
		t.Error(msg)
	}

	for i, value := range values {
		if value != expected[i] {
			t.Error("Values returned out of order indexed slice of values")
		}
	}
}
