package file

import (
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestReadFileLine(t *testing.T) {
	t.Run("read line one by one", func(t *testing.T) {
		_, file, _, _ := runtime.Caller(0)
		baseDir := filepath.Dir(file)
		target := filepath.Join(baseDir, "test", "test.text")
		result := make(map[string]int, 10)
		err := ReadFileLine(target, func(oneLine string) (isQuit bool) {
			if strings.Contains(oneLine, "one") {
				result["one"] += 1
			}

			if strings.Contains(oneLine, "aaa") {
				result["aaa"] += 1
			}
			return false
		})
		if assert.Nil(t, err) {
			assert.Equal(t, 1, result["one"])
			assert.Equal(t, 0, result["second"])
		}
	})
	t.Run("skip case", func(t *testing.T) {
		_, file, _, _ := runtime.Caller(0)
		baseDir := filepath.Dir(file)
		target := filepath.Join(baseDir, "test", "test.text")
		result := make(map[string]int, 10)
		err := ReadFileLine(target, func(oneLine string) (isQuit bool) {
			if strings.Contains(oneLine, "one") {
				result["one"] += 1
			}

			if strings.Contains(oneLine, "aaa") {
				result["aaa"] += 1
				return true
			}
			return false
		})
		if assert.Nil(t, err) {
			assert.Equal(t, 0, result["one"])
			assert.Equal(t, 1, result["aaa"])
		}
	})
}
