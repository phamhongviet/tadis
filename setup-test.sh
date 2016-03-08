#!/bin/bash

docker run -d -p 56379:6379 --name tadis_redis_test redis:3-alpine
