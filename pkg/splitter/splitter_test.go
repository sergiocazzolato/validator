package splitter_test

import (
	"fmt"
	"testing"

	"github.com/fgimenez/validator/pkg/splitter"
	"github.com/fgimenez/validator/pkg/types"
)

func TestSplit(t *testing.T) {
	subject := &splitter.Splitter{}
	options := &types.Options{
		Executors: 4,
	}
	t.Run("empty input", func(t *testing.T) {
		input := []string{}
		result, err := subject.Split(options, input)
		if err != nil {
			t.Errorf("error not expected, got %v", err)
		}
		if len(result) != 0 {
			t.Errorf("expected empty result, got %v", result)
		}
	})
	t.Run("{Executors} input", func(t *testing.T) {
		input := []string{"line0", "line1", "line2", "line3"}
		result, err := subject.Split(options, input)
		if err != nil {
			t.Errorf("error not expected, got %v", err)
		}
		for i := 0; i < options.Executors; i++ {
			expected := []string{fmt.Sprintf("line%d", i)}
			if result[i][0] != expected[0] {
				t.Errorf("expected result %s, got %v", expected[0], result[i][0])
			}
		}
	})
}
