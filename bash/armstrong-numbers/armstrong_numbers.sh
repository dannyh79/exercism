#!/usr/bin/env bash

is_armstrong_number() {
  local number="$1"
  local length=${#number}
  local converted=0
  for ((i = 0; i < length; i++)); do
    local digit=${number:$i:1}
    converted=$(echo "$converted + $digit ^ $length" | bc)
  done

  if [[ "$number" -eq "$converted" ]]; then
    echo true
  else
    echo false
  fi
}

is_armstrong_number "$@"
