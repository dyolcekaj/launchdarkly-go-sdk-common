package ldvalue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCopyArbitrarySlice(t *testing.T) {
	var sNil []string
	vm2 := CopyArbitraryValue(sNil)

	assert.Equal(t, NullType, vm2.Type())

	sEmpty := []string{}
	vm3 := CopyArbitraryValue(sEmpty)

	assert.Equal(t, ArrayType, vm3.Type())
	assert.Equal(t, []Value{}, vm3.arrayValue.data)

	sStr := []string{"a", "b", "c"}
	vm1 := CopyArbitraryValue(sStr)

	assert.Equal(t, ArrayType, vm1.Type())
	assert.Equal(t, []Value{String("a"), String("b"), String("c")}, vm1.arrayValue.data)

	sInt := []int{1, 2, 3}
	vm4 := CopyArbitraryValue(sInt)

	assert.Equal(t, ArrayType, vm4.Type())
	assert.Equal(t, []Value{Int(1), Int(2), Int(3)}, vm4.arrayValue.data)

	sUint := []uint{1, 2, 3}
	vm5 := CopyArbitraryValue(sUint)

	assert.Equal(t, ArrayType, vm5.Type())
	assert.Equal(t, []Value{Int(1), Int(2), Int(3)}, vm5.arrayValue.data)

	sFloat := []float64{1.1, 2.2, 3.3}
	vm6 := CopyArbitraryValue(sFloat)

	assert.Equal(t, ArrayType, vm6.Type())
	assert.Equal(t, []Value{Float64(1.1), Float64(2.2), Float64(3.3)}, vm6.arrayValue.data)

	sComplex := []complex128{1.1 + 2.2i, 3.3 + 4.4i, 5.5 + 6.6i}
	vm7 := CopyArbitraryValue(sComplex)
	assert.Equal(t, ArrayType, vm7.Type())
}

func TestCopyArbitraryMap(t *testing.T) {
	var mNil map[string]string
	vm2 := CopyArbitraryValue(mNil)

	assert.Equal(t, NullType, vm2.Type())

	mEmpty := map[string]string{}
	vm3 := CopyArbitraryValue(mEmpty)

	assert.Equal(t, ObjectType, vm3.Type())
	assert.Equal(t, map[string]Value{}, vm3.objectValue.data)

	mStr := map[string]string{"a": "1", "b": "2", "c": "3"}
	vm1 := CopyArbitraryValue(mStr)

	assert.Equal(t, ObjectType, vm1.Type())
	assert.Equal(t, map[string]Value{"a": String("1"), "b": String("2"), "c": String("3")}, vm1.objectValue.data)

	mInt := map[string]int{"a": 1, "b": 2, "c": 3}
	vm4 := CopyArbitraryValue(mInt)

	assert.Equal(t, ObjectType, vm4.Type())
	assert.Equal(t, map[string]Value{"a": Int(1), "b": Int(2), "c": Int(3)}, vm4.objectValue.data)

	mUint := map[string]uint{"a": 1, "b": 2, "c": 3}
	vm5 := CopyArbitraryValue(mUint)

	assert.Equal(t, ObjectType, vm5.Type())
	assert.Equal(t, map[string]Value{"a": Int(1), "b": Int(2), "c": Int(3)}, vm5.objectValue.data)

	mFloat := map[string]float64{"a": 1.1, "b": 2.2, "c": 3.3}
	vm6 := CopyArbitraryValue(mFloat)

	assert.Equal(t, ObjectType, vm6.Type())
	assert.Equal(t, map[string]Value{"a": Float64(1.1), "b": Float64(2.2), "c": Float64(3.3)}, vm6.objectValue.data)

	mComplex := map[string]complex128{"a": 1.1 + 2.2i, "b": 3.3 + 4.4i, "c": 5.5 + 6.6i}
	vm7 := CopyArbitraryValue(mComplex)

	assert.Equal(t, ObjectType, vm7.Type())

	mIntKeys := map[int]string{1: "a", 2: "b", 3: "c"}
	vm8 := CopyArbitraryValue(mIntKeys)

	assert.Equal(t, ObjectType, vm8.Type())
	assert.Equal(t, map[string]Value{"1": String("a"), "2": String("b"), "3": String("c")}, vm8.objectValue.data)
}
