package ptr

import (
	"reflect"
	"testing"
)

func TestToPtr(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		value := "some string"
		reflect.DeepEqual(ToPtr(value), &value)
	})

	t.Run("int", func(t *testing.T) {
		value := int(10)
		reflect.DeepEqual(ToPtr(10), &value)
	})

	t.Run("int8", func(t *testing.T) {
		value := int8(10)
		reflect.DeepEqual(ToPtr(10), &value)
	})

	t.Run("int16", func(t *testing.T) {
		value := int16(10)
		reflect.DeepEqual(ToPtr(10), &value)
	})

	t.Run("int32", func(t *testing.T) {
		value := int32(10)
		reflect.DeepEqual(ToPtr(10), &value)
	})

	t.Run("int32", func(t *testing.T) {
		value := int32(10)
		reflect.DeepEqual(ToPtr(10), &value)
	})
}
