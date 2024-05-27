#!/usr/bin/env bash

(( $1 % 3 == 0 )) && output+="Pling"
(( $1 % 5 == 0 )) && output+="Plang"
(( $1 % 7 == 0 )) && output+="Plong"

echo "${output:-$1}"

exit 0
