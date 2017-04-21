package splitter

import (
	"github.com/fgimenez/validator/pkg/types"
)

type Splitter struct{}

func (p *Splitter) Split(options *types.Options, input []string) [][]string {
	var result [][]string
	var partial []string
	itemsPerBucket := 1
	if len(input) >= options.Executors {
		itemsPerBucket = len(input) / options.Executors
	}
	for i, item := range input {
		if i >= itemsPerBucket*options.Executors {
			result[i%options.Executors] = append(result[i%options.Executors], item)
			continue
		}
		partial = append(partial, item)
		if i == len(input)-1 || len(partial) == itemsPerBucket {
			result = append(result, partial)
			partial = []string{}
		}
	}
	return result
}
