#!/bin/bash

set -e

if [ "$1" = '' ]; then
  set -- go run main.go
fi

if [[ "$1" == -* ]]; then
	set -- go run main.go "$@"
fi

if [ "$1" = 'dev' ]; then
  if [ "$2" = '--init' -o ! -d "_vendor" ]; then
    # Install go packages with Vektah/gin
    gom install github.com/Vektah/gin
  fi

  # Execute "gin" command with excluding un-watch directory paths
  # since wathing large number of files causes high CPU usage
  # https://github.com/codegangsta/gin/issues/53
	set -- gin -x _vendor -x docs
fi

echo "$@"
exec "$@"
