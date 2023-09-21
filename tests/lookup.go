package tests

import (
	"errors"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

// ErrNotFound is returned when the key didn't match
var ErrNotFound = errors.New("unable to find the key")
var nestedIndexRe = regexp.MustCompile(`\]\[`)

// LookupStringI returns the value found at the given path
func LookupStringI(i interface{}, path string) (reflect.Value, error) {
	value := reflect.ValueOf(i)
	var err error

	if path == "" {
		return value, nil
	}
	for _, part := range strings.Split(path, ".") {
		key, _ := parseIndex(part)
		value, err = lookupPartI(value, key)
		if err != nil {
			break
		}
		for _, listIndex := range splitNestedSlices(part) {
			_, indexValue := parseIndex(listIndex)
			value = lookupIndexI(value, indexValue)
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

func lookupIndexI(value reflect.Value, index int) reflect.Value {
	if value.Kind() != reflect.Slice || value.Len() <= index || index < 0 {
		return value
	}
	value = value.Index(index)
	if value.Kind() == reflect.Interface {
		value = value.Elem()
	}
	return value
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

// This function splits a nested slice into bracketed indexes
// Ex, "list[0][1]" is split into ["list[0]", "[1]"]
// Strings without a nested slice are returned in a slice but unchanged
// Ex, "list[0]" is returned as ["list[0]"]
func splitNestedSlices(part string) []string {
	nestedIndices := nestedIndexRe.FindAllStringIndex(part, -1)
	if nestedIndices != nil {
		parts := []string{}
		lastIdx := 0
		for _, pair := range nestedIndices {
			parts = append(parts, part[lastIdx:pair[1]-1])
			lastIdx = pair[1] - 1
		}
		parts = append(parts, part[lastIdx:])
		return parts
	}
	return []string{part}
}
