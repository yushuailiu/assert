package assert

import (
	"errors"
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

	MapHasKey(t, map[string]string{"m":"v","m1":"v1"}, "m")
	MapHasKey(t, map[int]string{1:"v",2:"v1"}, 1)
	MapHasKey(t, map[uint8]string{uint8(1):"v",uint8(2):"v1"}, uint8(1))
}

func TestMapHasValue(t *testing.T) {

	b := make(map[interface{}]int)
	b[&structDemo{val:1}] = 22
	c := &structDemo{val:2}
	f := &structDemo{val:3}

	s := structDemo{val:1}

	var a interface{}
	a = 1

	ch1 := make(chan int)

	mapHasValueList := [][]interface{}{
		[]interface{}{map[interface{}]interface{}{nil:"ff"}, "ff"},
		[]interface{}{map[interface{}]interface{}{1:true}, true},
		[]interface{}{map[interface{}]interface{}{true:1}, 1},
		[]interface{}{map[interface{}]interface{}{"ffs":""}, ""},
		[]interface{}{map[interface{}]interface{}{"":"ff"}, "ff"},
		[]interface{}{map[interface{}]interface{}{ch1:ch1}, ch1},
		[]interface{}{map[interface{}]interface{}{[1]int{1}:"ff"},"ff"},
		[]interface{}{map[interface{}]interface{}{[1]chan int{ch1}:ch1}, ch1},
		[]interface{}{map[interface{}]interface{}{[1]bool{true}:"ff"}, "ff"},
		[]interface{}{map[interface{}]interface{}{[1]string{""}:"ff"}, "ff"},
		[]interface{}{map[interface{}]interface{}{[1]*structDemo{f}:[1]*structDemo{f}},[1]*structDemo{f}},
		[]interface{}{map[interface{}]interface{}{[1]interface{}{a}:"ff"}, "ff"},
		[]interface{}{map[interface{}]interface{}{[1]structDemo{s}:"ff"}, "ff"},
		[]interface{}{map[interface{}]interface{}{[1]int{1}:"ff"}, "ff"},
		[]interface{}{map[interface{}]interface{}{"":"ff"},"ff"},
		[]interface{}{map[interface{}]interface{}{a:"ff"},"ff"},
		[]interface{}{map[interface{}]interface{}{a:"ff"},"ff"},
		[]interface{}{map[interface{}]interface{}{struct {}{}:"ff"}, "ff"},
		[]interface{}{map[interface{}]interface{}{&struct {}{}:"ff"}, "ff"}, // why ??
		[]interface{}{map[interface{}]interface{}{1.1:1.2}, 1.2},
		[]interface{}{map[interface{}]interface{}{uint(1):uint(1)}, uint(1)},
		[]interface{}{map[interface{}]interface{}{uint8(1):uint8(1)}, uint8(1)},
		[]interface{}{map[interface{}]interface{}{uint16(1):uint16(1)}, uint16(1)},
		[]interface{}{map[interface{}]interface{}{uint32(1):uint32(1)}, uint32(1)},
		[]interface{}{map[interface{}]interface{}{int(1):int(1)}, int(1)},
		[]interface{}{map[interface{}]interface{}{int8(1):int8(1)},int8(1)},
		[]interface{}{map[interface{}]interface{}{int16(1):int16(1)},int16(1)},
		[]interface{}{map[interface{}]interface{}{int32(1):int32(1)},int32(1)},
		[]interface{}{map[interface{}]interface{}{int64(1):int64(1)},int64(1)},
		[]interface{}{map[interface{}]interface{}{f:f}, f},
		[]interface{}{map[interface{}]interface{}{c:c}, c},
	}

	for _, items := range mapHasValueList {
		MapHasValue(t, items[0].(map[interface{}]interface{}), items[1])
	}

	MapHasValue(t, map[string]string{"m":"v","m1":"v1"}, "v1")
	MapHasValue(t, map[int]string{1:"v",2:"v1"}, "v")
	MapHasValue(t, map[uint8]string{uint8(1):"v",uint8(2):"v1"}, "v1")
}

func TestMapNotHasKey(t *testing.T) {

	b := make(map[interface{}]int)
	b[&structDemo{val:1}] = 22
	c := &structDemo{val:2}
	f := &structDemo{val:3}

	s := structDemo{val:1}

	var a interface{}
	a = 1

	ch1 := make(chan int)
	ch2 := make(chan int)

	mapHasKeyList := [][]interface{}{
		[]interface{}{map[interface{}]interface{}{nil:"ff"}, ""},
		[]interface{}{map[interface{}]interface{}{nil:"ff"}, 0},
		[]interface{}{map[interface{}]interface{}{nil:"ff"}, struct {}{}},
		[]interface{}{map[interface{}]interface{}{1:"ff"},1.0},
		[]interface{}{map[interface{}]interface{}{true:"ff"},false},
		[]interface{}{map[interface{}]interface{}{"ffs":"ff"},"ff"},
		[]interface{}{map[interface{}]interface{}{"":"ff"},"0"},
		[]interface{}{map[interface{}]interface{}{ch1:"ff"},nil},
		[]interface{}{map[interface{}]interface{}{ch1:"ff"},ch2},
		[]interface{}{map[interface{}]interface{}{[1]int{1}:"ff"},[1]int{2}},
		[]interface{}{map[interface{}]interface{}{[1]chan int{ch1}:"ff"},[1]chan int{ch2}},
		[]interface{}{map[interface{}]interface{}{[1]bool{true}:"ff"},[1]bool{false}},
		[]interface{}{map[interface{}]interface{}{[1]string{""}:"ff"},[1]interface{}{nil}},
		[]interface{}{map[interface{}]interface{}{[1]*structDemo{f}:"ff"},[1]*structDemo{c}},
		[]interface{}{map[interface{}]interface{}{[1]interface{}{a}:"ff"},[1]interface{}{2}},
		[]interface{}{map[interface{}]interface{}{[1]structDemo{s}:"ff"},[1]structDemo{*f}},
		[]interface{}{map[interface{}]interface{}{[1]int{1}:"ff"},[1]float32{1}},
		[]interface{}{map[interface{}]interface{}{"":"ff"},nil},
		[]interface{}{map[interface{}]interface{}{a:"ff"},uint(1)},
		[]interface{}{map[interface{}]interface{}{a:"ff"},uint8(1)},
		[]interface{}{map[interface{}]interface{}{a:"ff"},uint16(1)},
		[]interface{}{map[interface{}]interface{}{a:"ff"},uint32(1)},
		[]interface{}{map[interface{}]interface{}{a:"ff"},uint64(1)},
		[]interface{}{map[interface{}]interface{}{a:"ff"},int64(1)},
		[]interface{}{map[interface{}]interface{}{a:"ff"},int32(1)},
		[]interface{}{map[interface{}]interface{}{a:"ff"},int16(1)},
		[]interface{}{map[interface{}]interface{}{a:"ff"},int8(1)},
		[]interface{}{map[interface{}]interface{}{a:"ff"},float32(1)},
		[]interface{}{map[interface{}]interface{}{a:"ff"},float64(1)},
	}

	for _, items := range mapHasKeyList {
		MapNotHasKey(t, items[0].(map[interface{}]interface{}), items[1])
	}

	dd := make(map[string]string)
	MapNotHasKey(t, dd, "ff")
	MapNotHasKey(t, map[string]string{"m":"v","m1":"v1"}, "mss")

}

func TestMapHasKeysOnly(t *testing.T) {

	b := make(map[interface{}]int)
	b[&structDemo{val:1}] = 22
	c := &structDemo{val:2}
	f := &structDemo{val:3}

	s := structDemo{val:1}

	var a interface{}
	a = 1

	ch1 := make(chan int)

	mapHasKeyList := [][]interface{}{
		[]interface{}{map[interface{}]interface{}{nil:"ff"}, []interface{}{nil}},
		[]interface{}{map[interface{}]interface{}{1:"ff"}, []interface{}{1}},
		[]interface{}{map[interface{}]interface{}{true:"ff"}, []interface{}{true}},
		[]interface{}{map[interface{}]interface{}{"ffs":"ff"}, []interface{}{"ffs"}},
		[]interface{}{map[interface{}]interface{}{"":"ff"}, []interface{}{""}},
		[]interface{}{map[interface{}]interface{}{ch1:"ff"}, []interface{}{ch1}},
		[]interface{}{map[interface{}]interface{}{[1]int{1}:"ff"}, []interface{}{[1]int{1}}},
		[]interface{}{map[interface{}]interface{}{[1]chan int{ch1}:"ff"}, []interface{}{[1]chan int{ch1}}},
		[]interface{}{map[interface{}]interface{}{[1]bool{true}:"ff"}, []interface{}{[1]bool{true}}},
		[]interface{}{map[interface{}]interface{}{[1]string{""}:"ff"}, []interface{}{[1]string{""}}},
		[]interface{}{map[interface{}]interface{}{[1]*structDemo{f}:"ff"}, []interface{}{[1]*structDemo{f}}},
		[]interface{}{map[interface{}]interface{}{[1]interface{}{a}:"ff"}, []interface{}{[1]interface{}{a}}},
		[]interface{}{map[interface{}]interface{}{[1]structDemo{s}:"ff"}, []interface{}{[1]structDemo{s}}},
		[]interface{}{map[interface{}]interface{}{[1]int{1}:"ff"}, []interface{}{[1]int{1}}},
		[]interface{}{map[interface{}]interface{}{"":"ff"}, []interface{}{""}},
		[]interface{}{map[interface{}]interface{}{a:"ff"}, []interface{}{1}},
		[]interface{}{map[interface{}]interface{}{a:"ff"}, []interface{}{a}},
		[]interface{}{map[interface{}]interface{}{struct {}{}:"ff"}, []interface{}{struct {}{}}},
		[]interface{}{map[interface{}]interface{}{&struct {}{}:"ff"}, []interface{}{&struct {}{}}}, // why ??
		[]interface{}{map[interface{}]interface{}{1.0:"ff"}, []interface{}{1.0}},
		[]interface{}{map[interface{}]interface{}{uint(1):"ff"}, []interface{}{uint(1)}},
		[]interface{}{map[interface{}]interface{}{uint8(1):"ff"}, []interface{}{uint8(1)}},
		[]interface{}{map[interface{}]interface{}{uint16(1):"ff"}, []interface{}{uint16(1)}},
		[]interface{}{map[interface{}]interface{}{uint32(1):"ff"}, []interface{}{uint32(1)}},
		[]interface{}{map[interface{}]interface{}{int(1):"ff"}, []interface{}{1}},
		[]interface{}{map[interface{}]interface{}{int8(1):"ff"}, []interface{}{int8(1)}},
		[]interface{}{map[interface{}]interface{}{int16(1):"ff"}, []interface{}{int16(1)}},
		[]interface{}{map[interface{}]interface{}{int32(1):"ff"}, []interface{}{int32(1)}},
		[]interface{}{map[interface{}]interface{}{int64(1):"ff"}, []interface{}{int64(1)}},
		[]interface{}{map[interface{}]interface{}{f:"ff", uint8(1): "dd"}, []interface{}{f, uint8(1)}},
		[]interface{}{map[interface{}]interface{}{c:"ff", uint(1): "ss"}, []interface{}{c, uint(1)}},
	}

	for _, items := range mapHasKeyList {
		MapHasKeysOnly(t, items[0].(map[interface{}]interface{}), items[1].([]interface{}))
	}
}

func TestContains(t *testing.T) {

	a := &structDemo{val:1}
	ch1 := make(chan int)
	ch2 := make(chan int)
	containersList := [][]interface{}{
		{[]interface{}{1,2}, 2,},
		{[]interface{}{1,"f"}, "f",},
		{[]interface{}{"lj",""}, "",},
		{[]interface{}{structDemo{val:1},1}, structDemo{val:1},},
		{[]interface{}{structDemo{val:1}, structDemo{val:2}}, structDemo{val:2},},
		{[]interface{}{1}, 1,},
		{[]interface{}{true, 1}, true,},
		{[]interface{}{nil, false}, false,},
		{[]interface{}{nil, false}, false,},
		{[]interface{}{a, false}, a,},
		{[]interface{}{ch1, ch2}, ch1,},
		{[]interface{}{ch1, ch2}, ch2,},
		{[]interface{}{errors.New("ff"), ch2}, errors.New("ff"),},
		{[]interface{}{errors.New("ff"), structDemo{val:1}}, structDemo{val:1},},
	}

	for _, items := range containersList {
		Contains(t, items[0], items[1])
	}

	c := [2]int{2,3}
	Contains(t, c, 2)
}

func TestNotContains(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	containersList := [][]interface{}{
		{[]interface{}{1,2}, 3,},
		{[]interface{}{1,""}, "f",},
		{[]interface{}{"lj",0}, "",},
		{[]interface{}{structDemo{val:2},1}, structDemo{val:1},},
		{[]interface{}{structDemo{val:1}, structDemo{val:3}}, structDemo{val:2},},
		{[]interface{}{2}, 1,},
		{[]interface{}{false, 1}, true,},
		{[]interface{}{nil, true}, false,},
		{[]interface{}{nil, true}, false,},
		{[]interface{}{ch2}, ch1,},
		{[]interface{}{ch1}, ch2,},
		{[]interface{}{errors.New("ffd"), ch2}, errors.New("ff"),},
		{[]interface{}{errors.New("ff"), structDemo{val:2}}, structDemo{val:1},},
	}

	for _, items := range containersList {
		NotContains(t, items[0], items[1])
	}

	c := [2]int{2,3}
	NotContains(t, c, 4)
}

func TestStringContainsString(t *testing.T) {
	a := [][]string{
		{"testString", "test"},
		{"testString", "g"},
		{"testString", "ing"},
		{"testString", "Strin"},
		{"testString", "te"},
		{"testString", "t"},
		{"testString&", "&"},
	}

	for _, items := range a {
		StringContainsString(t, items[0], items[1])
	}
}

func TestStringContainsStringIgnoringCase(t *testing.T) {
	a := [][]string{
		{"testString", "string"},
		{"testString", "G"},
		{"testString", "Ing"},
		{"testString", "strin"},
		{"testString", "Te"},
		{"testString", "T"},
		{"testSTRING&", "str"},
		{"testSTRING&", "Str"},
	}

	for _, items := range a {
		StringContainsStringIgnoringCase(t, items[0], items[1])
	}
}

func TestStringNotContainsString(t *testing.T) {
	a := [][]string{
		{"testString", "yushuailiu"},
		{"TestString", "ss"},
	}

	for _, items := range a {
		StringNotContainsStringIgnoringCase(t, items[0], items[1])
	}
}

func TestStringNotContainsStringIgnoringCase(t *testing.T) {
	a := [][]string{
		{"testString", "STRINGsss"},
		{"TestString", "testStrings"},
	}

	for _, items := range a {
		StringNotContainsStringIgnoringCase(t, items[0], items[1])
	}
}