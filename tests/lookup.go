package tests

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
)

// ErrNotFound is returned when the key didn't match
var ErrNotFound = errors.New("unable to find the key")

// LookupStringI returns the value found at the given path
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
