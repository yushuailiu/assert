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
		[]interface{}{[]int{1,2}, []int{1,2}},
		[]interface{}{map[interface{}]interface{}{"a":1,"b":2,"c":3},map[interface{}]interface{}{"b":2,"a":1,"c":3}},
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
		[]interface{}{[]int{1,2}, []int{2,1}},
		[]interface{}{map[interface{}]interface{}{"a":1,"b":2,"c":3},map[interface{}]interface{}{"b":2,"a":1}},
	}

	for _, items := range notEqualList {
		NotEqual(t, items[0], items[1])
	}
}

func TestMapHasKey(t *testing.T) {

	b := make(map[interface{}]int)
	b[&structDemo{val:1}] = 22
	c := &structDemo{val:2}
	f := &structDemo{val:3}

	s := structDemo{val:1}

	var a interface{}
	a = 1

	ch1 := make(chan int)

	mapHasKeyList := [][]interface{}{
		[]interface{}{map[interface{}]interface{}{nil:"ff"}, nil},
		[]interface{}{map[interface{}]interface{}{1:"ff"},1},
		[]interface{}{map[interface{}]interface{}{true:"ff"},true},
		[]interface{}{map[interface{}]interface{}{"ffs":"ff"},"ffs"},
		[]interface{}{map[interface{}]interface{}{"":"ff"},""},
		[]interface{}{map[interface{}]interface{}{ch1:"ff"},ch1},
		[]interface{}{map[interface{}]interface{}{[1]int{1}:"ff"},[1]int{1}},
		[]interface{}{map[interface{}]interface{}{[1]chan int{ch1}:"ff"},[1]chan int{ch1}},
		[]interface{}{map[interface{}]interface{}{[1]bool{true}:"ff"},[1]bool{true}},
		[]interface{}{map[interface{}]interface{}{[1]string{""}:"ff"},[1]string{""}},
		[]interface{}{map[interface{}]interface{}{[1]*structDemo{f}:"ff"},[1]*structDemo{f}},
		[]interface{}{map[interface{}]interface{}{[1]interface{}{a}:"ff"},[1]interface{}{a}},
		[]interface{}{map[interface{}]interface{}{[1]structDemo{s}:"ff"},[1]structDemo{s}},
		[]interface{}{map[interface{}]interface{}{[1]int{1}:"ff"},[1]int{1}},
		[]interface{}{map[interface{}]interface{}{"":"ff"},""},
		[]interface{}{map[interface{}]interface{}{a:"ff"},1},
		[]interface{}{map[interface{}]interface{}{a:"ff"},a},
		[]interface{}{map[interface{}]interface{}{struct {}{}:"ff"}, struct {}{}},
		[]interface{}{map[interface{}]interface{}{&struct {}{}:"ff"}, &struct {}{}}, // why ??
		[]interface{}{map[interface{}]interface{}{1.0:"ff"},1.0},
		[]interface{}{map[interface{}]interface{}{uint(1):"ff"},uint(1)},
		[]interface{}{map[interface{}]interface{}{uint8(1):"ff"},uint8(1)},
		[]interface{}{map[interface{}]interface{}{uint16(1):"ff"},uint16(1)},
		[]interface{}{map[interface{}]interface{}{uint32(1):"ff"},uint32(1)},
		[]interface{}{map[interface{}]interface{}{int(1):"ff"},int(1)},
		[]interface{}{map[interface{}]interface{}{int8(1):"ff"},int8(1)},
		[]interface{}{map[interface{}]interface{}{int16(1):"ff"},int16(1)},
		[]interface{}{map[interface{}]interface{}{int32(1):"ff"},int32(1)},
		[]interface{}{map[interface{}]interface{}{int64(1):"ff"},int64(1)},
		[]interface{}{map[interface{}]interface{}{f:"ff"}, f},
		[]interface{}{map[interface{}]interface{}{c:"ff"}, c},
	}

	for _, items := range mapHasKeyList {
		MapHasKey(t, items[0].(map[interface{}]interface{}), items[1])
	}
}