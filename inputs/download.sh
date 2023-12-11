#!/usr/bin/env bash

if test ! "$ADVENTSESSION"; then
  echo "Set ADVENTSESSION variable to proceed (it's the 'session' cookie value)"
  exit 1
fi

if [ "$1" = '--all' ]; then
  for day in $(seq -f "%02g" 1 25)
  do
    dayn="$day"
    if test "${dayn:0:1}" = '0'; then
      dayn="${day:1}"
    fi

    echo "Downloading input for day $day..."
    if ! curl -s -f -H "Cookie: session=$ADVENTSESSION" -o "day$day.txt" "https://adventofcode.com/2023/day/$dayn/input"; then
      echo "...No more input or session invalid"
      exit 0
    fi
  done
  exit 0
fi

if test "$1"; then
  day="$(printf %02d "$1")"

  echo "Downloading input for day $1..."
  if ! curl -s -f -H "Cookie: session=$ADVENTSESSION" -o "day$day.txt" "https://adventofcode.com/2023/day/$1/input"; then
    echo "No input for $1"
    exit 1
  fi
else
  echo "Usage: $0 [--all] [n]"
  exit 1
fi
