package encoding

import (
	"fmt"
	"reflect"
)

type CompactArrayItem interface {
	ToBytes() (data []byte, length uint64, err error)
}

type CompactArray struct {
	_length uint64
	_items  []CompactArrayItem
}

func (c CompactArray) Length() uint64 {
	return c._length
}

func (c CompactArray) Items() []CompactArrayItem {
	return c._items
}

func (c CompactArray) AddItem(item CompactArrayItem) CompactArray {
	// perform a type check if there are already items in the array
	if c._length != 0 && reflect.TypeOf(item) != reflect.TypeOf(c._items[0]) {
		panic(fmt.Errorf(
			"got %s wanted %s: %w",
			reflect.TypeOf(item),
			reflect.TypeOf(c._items[0]),
			ErrUnexpectedItemType,
		))
	}

	// add item to items
	c._items = append(
		c._items,
		item,
	)

	return c
}
