package yaml

import (
	. "github.com/stretchr/testify/assert"
	"path/filepath"
	"runtime"
	"testing"
)

func TestLoadYamlFromFile(t *testing.T) {
	_, fileName, _, _ := runtime.Caller(0)
	testFilePath := filepath.Join(filepath.Dir(fileName), "test", "test.yaml")
	result := make(map[string]interface{}, 1)
	err := LoadYamlFromFile(testFilePath, &result)
	if Nil(t, err) {
		EqualValues(t, result["hello"], "world")
	}
}
