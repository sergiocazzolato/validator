package splitter

import "github.com/fgimenez/validator/pkg/types"

type Splitter struct{}

func (p *Splitter) Split(options *types.Options, input []string) ([][]string, error) {
	var result [][]string

	for _, item := range input {
		result = append(result, []string{item})
	}

	return result, nil
}
