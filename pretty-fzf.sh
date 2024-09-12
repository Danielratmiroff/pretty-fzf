#!/bin/bash
result=$(go run ./main.go "$@" | tail -n 1)

if [[ $result == cd* ]]; then
    echo "Changing directory to: $result"
    eval "$result"
else
    echo "$result"
fi