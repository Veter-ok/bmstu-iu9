#!/bin/bash

OUTPUT_FILE="./output.txt"
ERROR_FILE="./error.txt"

program="$1"
interval="$2"

start=$(date +%s)
$program > $OUTPUT_FILE 2> $ERROR_FILE
end=$(date +%s)
time=$(($end - $start))
waiting=$(($(($interval * 60)) - $time))

while true; do
    start=$(date +%s)
    $program >> $OUTPUT_FILE 2>> $ERROR_FILE
    end=$(date +%s)
    waiting=$(($(($interval * 60)) - $time))
    if [ $waiting -gt 0 ]; then
        sleep $waiting
    fi 
done