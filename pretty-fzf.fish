#!/usr/bin/env fish

set result (go run ./main.go $argv | tail -n 1)

if string match -q "cd *" $result
    echo "Changing directory to: $result"
    eval $result
else if string match -q "vim *" $result
    echo "Opening vie: $result"
    eval $result
else
    echo $result
end