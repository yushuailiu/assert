package assert

import (
	"testing"
)

type structDemo struct {
	val int
}

func TestEqual(t *testing.T) {
	Equal(t, 1, 1)
	Equal(t, 1.0, 1.0)
	Equal(t, true, true)
	Equal(t, false, false)
	Equal(t, new(struct{}), new(struct{}))
	Equal(t, new(structDemo), new(structDemo))
	Equal(t, structDemo{}, new(structDemo))
	Equal(t, structDemo{val:1}, structDemo{val:1})
	Equal(t, &structDemo{val:1}, structDemo{val:1})
	Equal(t, &structDemo{val:1}, &structDemo{val:1})
	Equal(t, &structDemo{val:1}, structDemo{val:1})

	equalList := [][]interface{}{
		[]interface{}{1.0, float64(1.0)},
	}

	for _, items := range equalList {
		Equal(t, items[0], items[1])
	}
}

// test NotEqual
func TestNotEqual(t *testing.T) {
	NotEqual(t, 1, 2)
	NotEqual(t, 1, false)
	NotEqual(t, 1, true)
	NotEqual(t, 1, 1.0)
	NotEqual(t, 1.0, 1)

	var c1 complex64 = complex(1, 2)
	var c2 = complex(1, 2)
	NotEqual(t, c1, c2)

	notEqualList := [][]interface{}{
		[]interface{}{1, 2},
		[]interface{}{1, true},
		[]interface{}{1, nil},
		[]interface{}{1, 2.0},
		[]interface{}{1, false},
		[]interface{}{1, new(map[interface{}]interface{})},
		[]interface{}{1, new([]int)},
		[]interface{}{1, new(struct{})},
		[]interface{}{1, &struct {}{}},
		[]interface{}{1, &structDemo{val:1}},
		[]interface{}{1, 1.0},
		[]interface{}{1, uint(1)},
		[]interface{}{1, uint8(1)},
		[]interface{}{1, uint16(1)},
		[]interface{}{1, uint32(1)},
		[]interface{}{1, uint64(1)},
		[]interface{}{1, int8(1)},
		[]interface{}{1, int16(1)},
		[]interface{}{1, int32(1)},
		[]interface{}{1, int64(1)},
		[]interface{}{1.0, float32(1.0)},
	}

	for _, items := range notEqualList {
		NotEqual(t, items[0], items[1])
	}
}