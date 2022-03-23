package tests

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// ErrNotFound is returned when the key didn't match
var ErrNotFound = errors.New("Unable to find the key")

// LookupStringI returnes the value found at the given path
func LookupStringI(i interface{}, path string) (reflect.Value, error) {
	value := reflect.ValueOf(i)
	var err error

	if path == "" {
		return value, nil
	}
	for _, part := range strings.Split(path, ".") {
		key, index := parseIndex(part)
		value, err = lookupPartI(value, key)
		if err != nil {
			break
		}
		if index != -1 {
			if value.Kind() != reflect.Slice || value.Len() <= index {
				return reflect.Value{}, ErrNotFound
			}
			value = value.Index(index)
			if value.Kind() == reflect.Interface {
				value = value.Elem()
			}
		}
	}
	return value, err
}

func lookupPartI(value reflect.Value, key string) (reflect.Value, error) {
	switch value.Kind() {
	case reflect.Map:
		iter := value.MapRange()
		for iter.Next() {
			if strings.EqualFold(key, iter.Key().String()) {
				value = iter.Value()
				if value.IsNil() {
					return value, nil
				}
				switch value.Elem().Kind() {
				case reflect.Float32, reflect.Float64:
					valueString := fmt.Sprintf("%f", value.Elem().Float())
					if strings.Contains(valueString, ".") {
						return value.Elem(), nil
					}
					valueInt64 := int64(value.Elem().Float())
					return reflect.ValueOf(valueInt64), nil
				}
				return value.Elem(), nil
			}
		}
	}
	return reflect.Value{}, ErrNotFound
}

func parseIndex(part string) (string, int) {
	opening := strings.Index(part, "[")
	closing := strings.Index(part, "]")
	if opening != -1 && closing != -1 {
		indexContent := part[opening+1 : closing]
		index, err := strconv.Atoi(indexContent)
		if err == nil {
			return part[:opening], index
		}
	}
	return part, -1
}
