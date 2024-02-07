#!/bin/bash

echo "Stop container on docker engin"
docker stop server nginx-server

echo "Remove related images"
docker image rm -f server nginx-server

echo "browing to part project"
cd "/home/server"

echo "Build up image container"
docker-compose up -d -V --build
