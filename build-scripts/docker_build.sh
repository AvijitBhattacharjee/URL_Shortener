#!/bin/bash

export DOCKER_BUILDKIT=0

docker build -t myurl:1.0 .
docker run -d myurl:1.0

docker tag myurl:1.0 avbhatta/myurl:1.0
docker login

docker push avbhatta/myurl:1.0