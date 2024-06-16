#!/bin/bash

minikube start --driver=docker
minikube status

kubectl create namespace test
cd helm/

helm install myurl . -n test 
kubectl get po -n test | grep myurl