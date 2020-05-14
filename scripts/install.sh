#!/usr/bin/env bash

set -e # fail execution if execution of sourced files failed

my_dir="$(dirname "$0")"
source "$my_dir/envs.sh"

`get_env`

echo "Installing fabric test setup."
curl -sSL https://bit.ly/2ysbOFE | bash -s
echo "Fabric test setup is installed."