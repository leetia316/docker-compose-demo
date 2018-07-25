#!/bin/bash
set -eo pipefail

port=$port

if ping="$(curl -f http://localhost:8080/api/ping)" && [ "$ping" = 'pong' ]; then
	exit 0
fi

exit 1