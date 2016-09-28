#!/bin/bash

for i in $(seq 1 3); do
    echo "Running batch $i..."
    ../snappy_benchmark/tests/util/benchmark.sh linode: 5 "./benchmark_${i}.out"
done
