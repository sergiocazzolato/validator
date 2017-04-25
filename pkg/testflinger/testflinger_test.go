package testflinger_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/fgimenez/validator/pkg/testflinger"
	"github.com/fgimenez/validator/pkg/types"
)

func TestGenerateCfg(t *testing.T) {
	subject := &testflinger.Testflinger{}
	options := &types.Options{
		Channel: "mychannel",
	}
	t.Run("empty input", func(t *testing.T) {
		input := [][]string{}
		result := subject.GenerateCfg(options, input)
		if len(result) != 0 {
			t.Errorf("expected empty result, got %v", result)
		}
	})

	t.Run("config file for sigle line, single group input", func(t *testing.T) {
		input := [][]string{{"line0"}}
		result := subject.GenerateCfg(options, input)
		defer os.Remove(result[0])
		t.Run("is created", func(t *testing.T) {
			if _, err := os.Stat(result[0]); os.IsNotExist(err) {
				t.Errorf("%v is not a file", result[0])
			}
		})
		t.Run("has the right content", func(t *testing.T) {
			content, _ := ioutil.ReadFile(result[0])
			expected := fmt.Sprintf(testflinger.FromTargetFmt, options.Channel, "line0")
			if string(content) != expected {
				t.Errorf("%s file content wrong, actual %s, expected %s", result[0], content, expected)
			}
		})
	})
}
