#!/usr/bin/env bash

set -e # fail execution if execution of sourced files failed

my_dir="$(dirname "$0")"
source "$my_dir/envs.sh"

`get_env`

echo "Starting simplequeue project."
docker-compose up
#docker-compose up -d