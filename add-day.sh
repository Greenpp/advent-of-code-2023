#!/bin/bash

set -e -o pipefail

if [ -z $1 ]; then
    echo "Day number required"
    exit 1
fi

day_dir="day${1}"
mkdir $day_dir
cp -r templates/part $day_dir/part1
cp -r templates/part $day_dir/part2
