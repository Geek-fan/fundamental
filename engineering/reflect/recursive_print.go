package reflect

import (
	"reflect"
	"strconv"
	"strings"
)

func hexLetter(l uint64, capital bool) byte {
	if l < 10 {
		return byte('0' + l)
	}

	if capital {
		return byte('A' + l - 10)
	}
	return byte('a' + l - 10)
}

func printPtr(p uintptr) string {
	bits := 11
	m := uint64(p)
	b := strings.Builder{}
	b.WriteString("0x")
	temp := make([]byte, bits)
	for i := 0; i < bits; i++ {
		temp[bits-1-i] = hexLetter(m%16, false)
		m /= 16
		if m == 0 {
			break
		}
	}

	for _, c := range temp {
		if c == 0 {
			b.WriteByte('0')
		} else {
			b.WriteByte(c)
		}
	}

	return b.String()
}

type structField struct {
	tag   string
	value string
}

func printStruct(v reflect.Value) string {
	var fields []*structField
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		var value string

		field := v.Field(i)
		switch field.Kind() {
		case reflect.Ptr:
			if field.IsNil() {
				value = "<nil>"
			} else {
				if t.Field(i).IsExported() {
					value = "&" + RecursivePrint(field.Elem().Interface())
				} else {
					value = printPtr(field.Pointer())
				}
			}
		case reflect.Int:
			value = strconv.FormatInt(field.Int(), 10)
		case reflect.String:
			value = field.String()
		case reflect.Struct:
			value = printStruct(field)
		}

		tag := t.Field(i).Name
		fields = append(fields, &structField{tag, value})
	}

	b := strings.Builder{}
	b.WriteString("{")
	for i, field := range fields {
		b.WriteString(field.tag + ":" + field.value)
		if i < len(fields)-1 {
			b.WriteByte(' ')
		}
	}
	b.WriteString("}")
	return b.String()
}

func printSlice(v reflect.Value) string {
	b := strings.Builder{}
	b.WriteByte('[')
	length := v.Len()
	for i := 0; i < length; i++ {
		b.WriteString(RecursivePrint(v.Index(i).Interface()))
		if i < length-1 {
			b.WriteByte(' ')
		}
	}
	b.WriteByte(']')
	return b.String()
}

func RecursivePrint(s interface{}) string {
	v := reflect.ValueOf(s)
	if s == nil {
		return "<nil>"
	}

	switch v.Kind() {
	case reflect.Struct:
		return printStruct(reflect.ValueOf(s))
	case reflect.Ptr:
		return "&" + RecursivePrint(v.Elem().Interface())
	case reflect.Slice, reflect.Array:
		return printSlice(v)
	case reflect.Int:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.String, reflect.Func, reflect.Chan:
		return v.String()
	default:
	}
	return "not supported"
}
