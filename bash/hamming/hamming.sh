#!/usr/bin/env bash

main() {
  if [[ $# -lt 2 ]]; then
    echo "Usage: hamming.sh <string1> <string2>"
    exit 2
  fi

  local strand1="$1"
  local strand2="$2"

  local length1=${#strand1}
  local length2=${#strand2}

  if [[ "$length1" != "$length2" ]]; then
    echo "strands must be of equal length"
    exit 1
  fi

  local diff=0
  for (( i = 0; i < length1; i++ )); do
    local letter1=${strand1:i:1}
    local letter2=${strand2:i:1}

    if [[ "$letter1" != "$letter2" ]]; then
      ((diff++))
    fi
  done

  echo $diff
}

main "$@"
