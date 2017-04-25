package testflinger_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
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
	t.Run("config file for single line, multigroup input", func(t *testing.T) {
		input := [][]string{{"line0"}, {"line2"}, {"line3"}, {"line4"}}
		result := subject.GenerateCfg(options, input)
		for i, item := range input {
			defer os.Remove(result[i])
			file := fmt.Sprintf("file%d", i)
			t.Run(file+" is created", func(t *testing.T) {
				if _, err := os.Stat(result[i]); os.IsNotExist(err) {
					t.Errorf("%v is not a file", item)
				}
			})
			t.Run(file+" has the right content", func(t *testing.T) {
				content, _ := ioutil.ReadFile(result[i])
				expected := fmt.Sprintf(testflinger.FromTargetFmt, options.Channel, input[i][0])
				if string(content) != expected {
					t.Errorf("%s file content wrong, actual %s, expected %s", item, content, expected)
				}
			})
		}
	})
	t.Run("config file for multiline, multigroup input", func(t *testing.T) {
		input := [][]string{
			{"line0", "line1", "line2"},
			{"line0"},
			{"line0", "line1", "line2", "line3", "line4"},
			{"line0", "line1", "line2", "line3", "line4", "line5", "line6", "line7", "line8"}}
		result := subject.GenerateCfg(options, input)
		for i, item := range input {
			defer os.Remove(result[i])
			file := fmt.Sprintf("file%d", i)
			t.Run(file+" is created", func(t *testing.T) {
				if _, err := os.Stat(result[i]); os.IsNotExist(err) {
					t.Errorf("%v is not a file", item)
				}
			})
			t.Run(file+" has the right content", func(t *testing.T) {
				content, _ := ioutil.ReadFile(result[i])
				mergedLines := strings.Join(input[i], " ")
				expected := fmt.Sprintf(testflinger.FromTargetFmt, options.Channel, mergedLines)
				if string(content) != expected {
					t.Errorf("%s file content wrong, actual %s, expected %s", item, content, expected)
				}
			})
		}
	})
}
