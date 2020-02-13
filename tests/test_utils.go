package tests

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"time"
)

func Retry(interval time.Duration, count int, call func() bool) error {
	for i := 0; i < count; i++ {
		if call() {
			return nil
		}
		time.Sleep(interval)
	}
	return fmt.Errorf("Retry error: failed to satisfy the condition after %d times", count)
}

func ReadFixture(path string) (string, error) {
	fixturePath, err := filepath.Abs(path)
	if err != nil {
		return "", fmt.Errorf("failed to get fixture file path: %v", err)
	}
	data, err := ioutil.ReadFile(fixturePath)
	if err != nil {
		return "", fmt.Errorf("failed to open fixture file: %v", err)
	}
	return string(data), nil
}
