k8s-trailhead
-------------


This project aims to be a starting point for Kubernetes Users who are curious about interacting with the Kubernetes API via the [client-go](https://github.com/kubernetes/client-go) library. Inspired by the lead of Kelsey Hightower's [Kubernetes the hard way](https://github.com/kelseyhightower/kubernetes-the-hard-way) and Jessie Frazelle's [Kubernetes Snowflake](https://github.com/jessfraz/k8s-snowflake) project.  

Interacting directly with the Kubernetes API enables programatic controls of resources and clean definition of procedures. 

This project was originally demoed at 2017 Kubecon Austin in the talk "Oregon Trial to Kubernetes".(Demo gods were gracious, nerves less so.. attempting attonement for a shaky talk). 

## Why

the `client-go` library can be rough to get up and running with. Dependencies are the primary problem, and structure layout has changed between versions.

`client-go` version `v5.x` has landed with strong support for Kubernetes 1.8, and hopefully is the beginning of smooth client upgrades going forward.

## Vendoring

This project chooses to use [dep](https://github.com/golang/dep) as the dependency manager. This does have some [issues](https://github.com/kubernetes/client-go/blob/master/INSTALL.md#dep-not-supported-yet), but it works for the mainline cases. Supposedly the issues arise when utilizing the RBAC types which pull dependencies which cannot be flattened for the vendor directory.

For the mainline features, starting with the `Gopkg.toml` should work, copy it into other projects as a starting point. Once you've started writing yourown project, `dep ensure` to get the required packages vendored. `dep` is working on the [issue](https://github.com/golang/dep/issues/1124), although it's a [transitive dependency](https://github.com/golang/dep/issues/1124#issuecomment-333457621) issue, which is very [complex](https://github.com/golang/dep/issues/1124#issuecomment-331346439).

The `Gopkg.toml` locks client-go to `v5.0.1` and Kubernetes libraries to `kubernetes-1.8.2`. If you encounter any issues, please report them here and I'll try to resolve them and provide feedback upstream to `dep` if it's additional information.

## Showcase

### Deployment Object Specification

[Deployments](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/) are a core part object in Kubernetes. Typically defined in YAML, but if the desire is to build a set of Deployments with parameterized differences; YAML templating can be tricky to validated and test without dropping out to bash. The [Operator pattern](https://coreos.com/blog/introducing-operators.html) pioneered by CoreOS takes advantage of the Go client library and attempts to solve stateful problems with minimal manual intervention/management.

### End to End Testing

This runs a full integration test of a nginx(or whatever image you like) Deployment being created, tested for operation, and then cleanup. No bash, pipes, kubectl, or YAML involved!

Requirements: 
* [minikube](https://github.com/kubernetes/minikube) up and running.
* `export TESTMINIKUBE=1` to enable integration test.

Execute Simple E2E Test:

`go test ./kubernetes/... -v -run TestSimpleDeployE2E`

