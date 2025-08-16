#!/bin/bash

CONTEXT=k3d-another-kaixin-demo
NAMESPACE=default

kubectl config use-context $CONTEXT
kubectl config set-context --current --namespace=$NAMESPACE