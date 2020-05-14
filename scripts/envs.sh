#!/usr/bin/env bash

DOT_ENV_DIST_FILE="$PWD/.env.dist"
DOT_ENV_FILE="$PWD/.env"

if [[ ! -f "$DOT_ENV_FILE" ]]; then
    echo "(!) $DOT_ENV_FILE does not exist. Creating..." >&2

    if [[ ! -f "$DOT_ENV_DIST_FILE" ]]; then
        echo "(E) $DOT_ENV_DIST_FILE does not exist. Creating of $DOT_ENV_FILE failed." >&2
        exit 1
    fi

    cp "$DOT_ENV_DIST_FILE" "$DOT_ENV_FILE"

    echo "(!) $DOT_ENV_FILE is created!" >&2
fi

get_env() {
    local envs=$(cat $PWD/.env | grep -v '^#' | grep -o '^[^#]*')

    if [[ -z "$envs" ]]
    then
          echo "No enviroment valiables found in .env file" >&2
    else
          echo export $(echo "$envs" | xargs)
    fi
}