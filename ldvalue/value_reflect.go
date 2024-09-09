package ldvalue

import "reflect"

func tryArbitraryCollectionCopy(value any) (Value, bool) {
	vof := reflect.ValueOf(value)

	switch vof.Kind() {
	case reflect.Slice:
		if vof.IsNil() {
			return Null(), true
		} else if vof.Len() == 0 {
			return Value{valueType: ArrayType, arrayValue: ValueArray{data: emptyArray}}, true
		}

		return copyArbitrarySlice(vof), true
	case reflect.Map:
		if vof.IsNil() {
			return Null(), true
		} else if vof.Len() == 0 {
			return Value{valueType: ObjectType, objectValue: ValueMap{data: emptyMap}}, true
		}

		// anything else fall back to JSON slow path
		if vof.Type().Key().Kind() != reflect.String {
			return Value{}, false
		}

		return copyArbitraryMap(vof), true
	default:
		return Value{}, false
	}
}

func copyArbitraryMap(mapValue reflect.Value) Value {
	switch mapValue.Type().Elem().Kind() {
	case reflect.String:
		return buildArbitraryMap(mapValue, func(value reflect.Value) Value {
			return String(value.String())
		})
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return buildArbitraryMap(mapValue, func(value reflect.Value) Value {
			return Float64(float64(value.Int()))
		})
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return buildArbitraryMap(mapValue, func(value reflect.Value) Value {
			return Float64(float64(value.Uint()))
		})
	case reflect.Float32, reflect.Float64:
		return buildArbitraryMap(mapValue, func(value reflect.Value) Value {
			return Float64(value.Float())
		})
	default:
		return buildArbitraryMap(mapValue, func(value reflect.Value) Value {
			return CopyArbitraryValue(value.Interface())
		})
	}
}

func copyArbitrarySlice(slice reflect.Value) Value {
	switch slice.Type().Elem().Kind() {
	case reflect.String:
		return buildArbitrarySlice(slice, func(value reflect.Value) Value {
			return String(value.String())
		})
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return buildArbitrarySlice(slice, func(value reflect.Value) Value {
			return Float64(float64(value.Int()))
		})
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return buildArbitrarySlice(slice, func(value reflect.Value) Value {
			return Float64(float64(value.Uint()))
		})
	case reflect.Float32, reflect.Float64:
		return buildArbitrarySlice(slice, func(value reflect.Value) Value {
			return Float64(value.Float())
		})
	default:
		return buildArbitrarySlice(slice, func(value reflect.Value) Value {
			return CopyArbitraryValue(value.Interface())
		})
	}
}

func buildArbitraryMap(mapValue reflect.Value, valueFn func(value reflect.Value) Value) Value {
	builder := ValueMapBuildWithCapacity(mapValue.Len())

	for _, key := range mapValue.MapKeys() {
		builder.Set(key.String(), valueFn(mapValue.MapIndex(key)))
	}

	return builder.Build().AsValue()
}

func buildArbitrarySlice(slice reflect.Value, valueFn func(value reflect.Value) Value) Value {
	builder := ValueArrayBuildWithCapacity(slice.Len())

	for i := 0; i < slice.Len(); i++ {
		builder.Add(valueFn(slice.Index(i)))
	}

	return builder.Build().AsValue()
}
