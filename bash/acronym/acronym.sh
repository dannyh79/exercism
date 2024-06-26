#!/usr/bin/env bash

to_acronym() {
  local phrase="$1"
  local cleaned_phrase=${phrase//[-|_|*]/\ }
  local acronym=""
  for word in $cleaned_phrase; do
    acronym+="${word:0:1}"
  done

  echo "$acronym" | tr '[:lower:]' '[:upper:]'
}

to_acronym "$@"
