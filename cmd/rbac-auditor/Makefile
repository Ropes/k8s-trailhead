.PHONY: all build-image build tag-image push-tag minikube-run registry-run registry-port-uplink registry-setup

MINIKUBEIP ?= $(shell minikube ip)

build-bin: k8s-go-threads
	go build -a -v

build-image:
	docker build -t k8s-go-threads:0.1 .

tag-image:
	docker tag k8s-go-threads:0.1 localhost:5000/k8s-go-threads01

push-tag:
	docker push localhost:5000/k8s-go-threads01

minikube-run:
	kubectl run --image=localhost:5000/k8s-go-threads01 threads

registry-run:
	docker run -d -p 5000:5000 --restart=always --name registry registry:2

registry-port-uplink: registry-run
	ssh -f -i ~/.minikube/machines/minikube/id_rsa docker@$(MINIKUBEIP) -R 5000:localhost:5000

registry-setup:  registry-port-uplink

build: build-bin build-image

