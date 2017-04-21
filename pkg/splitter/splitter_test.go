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
		result := subject.Split(options, input)
		if len(result) != 0 {
			t.Errorf("expected empty result, got %v", result)
		}
	})
	t.Run("exactly {Executors} input", func(t *testing.T) {
		input := []string{"line0", "line1", "line2", "line3"}
		result := subject.Split(options, input)
		for i := 0; i < len(result); i++ {
			expected := []string{fmt.Sprintf("line%d", i)}
			if result[i][0] != expected[0] {
				t.Errorf("expected result %s, got %v", expected[0], result[i][0])
			}
		}
	})
	t.Run("greater than {Executors} input", func(t *testing.T) {
		input := []string{
			"line0", "line1", "line2", "line3", "line4",
			"line5", "line6", "line7", "line8", "line9",
			"line10", "line11", "line12", "line13", "line14",
			"line15", "line16", "line17", "line18", "line19",
		}
		result := subject.Split(options, input)
		offset := 0
		for i := 0; i < len(result); i++ {
			if len(result[i]) != len(input)/options.Executors {
				t.Errorf("expected results of length 4, got %d", len(result[i]))
			}
			for j := 0; j < len(result[i]); j++ {
				expected := fmt.Sprintf("line%d", offset+j)
				if result[i][j] != expected {
					t.Errorf("expected result %s, got %v", expected, result[i][j])
				}
			}
			offset += len(result[i])
		}
	})
	t.Run("less than {Executors} input", func(t *testing.T) {
		input := []string{
			"line0", "line1",
		}
		result := subject.Split(options, input)
		offset := 0
		for i := 0; i < len(result); i++ {
			if len(result[i]) != 1 {
				t.Errorf("expected results of length 1, got %d", len(result[i]))
			}
			expected := fmt.Sprintf("line%d", offset)
			if result[i][0] != expected {
				t.Errorf("expected result %s, got %v", expected, result[i][0])
			}
			offset++
		}
	})
	t.Run("uneven {Executors} multiple input", func(t *testing.T) {
		input := []string{
			"line0", "line1", "line2", "line3",
			"line4", "line5",
		}
		result := subject.Split(options, input)

		if len(result) != 4 {
			t.Errorf("expected total length of results of 4, got %d", len(result))
		}
		t.Log("result: ", result)
		if len(result[0]) != 2 {
			t.Errorf("expected results of length 2, got %d", len(result[0]))
		}
		if len(result[1]) != 2 {
			t.Errorf("expected results of length 2, got %d", len(result[1]))
		}
		if len(result[2]) != 1 {
			t.Errorf("expected results of length 1, got %d", len(result[2]))
		}
		if len(result[2]) != 1 {
			t.Errorf("expected results of length 1, got %d", len(result[2]))
		}
		if result[0][0] != "line0" || result[0][1] != "line4" {
			t.Errorf("expected results [line0, line4], got %v", result[0])
		}
	})
}
