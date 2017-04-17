package runner_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/fgimenez/validator/pkg/runner"
	"github.com/fgimenez/validator/pkg/types"
)

type fakeCli struct{}

var cliReturn string
var cliCalls int
var cliError bool

func (fc *fakeCli) ExecCommand(cmd ...string) (string, error) {
	cliCalls++
	if cliError {
		return "", errors.New("cli error")
	}
	return cliReturn, nil
}

type fakeSplitter struct{}

var splitReturn [][]string
var splitCalls int
var splitError bool

func (fs *fakeSplitter) Split(options *types.Options, input []string) ([][]string, error) {
	splitCalls++
	if splitError {
		return nil, errors.New("split error")
	}
	return splitReturn, nil
}

type fakeTestflinger struct{}

var generateCfgReturn string
var generateCfgCalls int
var generateCfgError bool

func (ts *fakeTestflinger) GenerateCfg(options *types.Options, input []string) (string, error) {
	generateCfgCalls++
	if generateCfgError {
		return "", errors.New("generateCfg error")
	}
	return fmt.Sprintf("%s-%d", generateCfgReturn, generateCfgCalls), nil
}

func TestRunner(t *testing.T) {
	s := runner.New(&types.RunnerDependencies{
		Cli:         &fakeCli{},
		Splitter:    &fakeSplitter{},
		Testflinger: &fakeTestflinger{},
	})
	options := &types.Options{
		System:    "mysystem",
		Executors: 4,
	}

	cliReturn = "line1\nline2\nline3\nline4"
	splitReturn = [][]string{{"line1"}, {"line2"}, {"line3"}, {"line4"}}
	generateCfgReturn = "/tmp/output"

	t.Run("happy-path", func(t *testing.T) {
		output, err := s.Run(options)
		t.Run("cli is called", func(t *testing.T) {
			if cliCalls != 1 {
				t.Errorf("expected 1 call to cli, obtained %d", cliCalls)
			}
		})
		t.Run("split is called", func(t *testing.T) {
			if splitCalls != 1 {
				t.Errorf("expected 1 call to split, obtained %d", splitCalls)
			}
		})
		t.Run("generateCfg is called", func(t *testing.T) {
			if generateCfgCalls != len(splitReturn) {
				t.Errorf("expected %d call to generateCfg, obtained %d", len(splitReturn), generateCfgCalls)
			}
		})
		t.Run("output is received", func(t *testing.T) {
			if err != nil {
				t.Errorf("expected nil error, got %v", err)
			}
			for i := 0; i < len(splitReturn); i++ {
				expected := fmt.Sprintf("/tmp/output-%d", i+1)
				if output[i] != expected {
					t.Errorf("expected output %s, got %s", expected, output[i])
				}
			}
		})
	})
	t.Run("unhappy-path cli error", func(t *testing.T) {
		cliError = true
		defer func() { cliError = false }()
		output, err := s.Run(options)
		if output != nil {
			t.Errorf("expected nil output, got %v", output)
		}
		if err.Error() != "cli error" {
			t.Errorf("expected cli error, got %v", err)
		}
	})
	t.Run("unhappy-path split error", func(t *testing.T) {
		splitError = true
		defer func() { splitError = false }()
		output, err := s.Run(options)
		if output != nil {
			t.Errorf("expected nil output, got %v", output)
		}
		if err.Error() != "split error" {
			t.Errorf("expected split error, got %v", err)
		}
	})
	t.Run("unhappy-path genCfg error", func(t *testing.T) {
		generateCfgError = true
		defer func() { generateCfgError = false }()
		output, err := s.Run(options)
		if output != nil {
			t.Errorf("expected nil output, got %v", output)
		}
		if err.Error() != "generateCfg error" {
			t.Errorf("expected generateCfg error, got %v", err)
		}
	})
}
