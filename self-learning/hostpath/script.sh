#!/bin/bash

image_name=hostpath-demo
version=v1.0

docker build -t ${image_name}:${version} .
docker tag ${image_name}:${version} kaixin-registry:12345/${image_name}:${version}
docker push kaixin-registry:12345/${image_name}:${version}
