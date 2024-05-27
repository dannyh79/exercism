#!/usr/bin/env bash

if (( $# == 0 || $# > 1 )); then
  echo "Usage: error_handling.sh <person>"
  exit 1
else
  echo "Hello, ${1}"
fi

exit 0
