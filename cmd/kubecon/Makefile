# Makefile for building kubecon image
.PHONY: all build-bin build-image

TAG ?= $(shell git rev-parse --short HEAD | tr -d '\n')
SVC ?= localhost:5000/kubecon 
# Image name is targeted to a registry hosted in minikube,
# and port 5000 proxied to the registry pod.

IMAGE ?= $(SVC):$(TAG)
CGO_ENABLED=0

all: build-bin build-image

build-bin:
	go build --ldflags '-extldflags "-static"' -a -tags netgo

build-image: kubecon
	docker build --no-cache=true -t $(IMAGE) .

build: build-bin build-image

