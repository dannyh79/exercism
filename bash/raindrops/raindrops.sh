#!/usr/bin/env bash

output=$1

if (( $1 % 3 == 0 )); then
  output="Pling"
fi

if (( $1 % 5 == 0 && output == $1 )); then
  output="Plang"
elif (( $1 % 5 == 0 && output != $1 )); then
  output="${output}Plang"
fi

if (( $1 % 7 == 0 && output == $1 )); then
  output="Plong"
elif (( $1 % 7 == 0 && output != $1 )); then
  output="${output}Plong"
fi

echo $output

exit 0
