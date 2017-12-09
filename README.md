k8s-trailhead
-------------


This project aims to be a starting point for Kubernetes Users who are curious about interacting with the Kubernetes API via the [client-go](https://github.com/kubernetes/client-go) library. Attempting to follow the lead of Kelsey Hightower's [Kubernetes the hard way](https://github.com/kelseyhightower/kubernetes-the-hard-way) and Jessie Frazelle's [Kubernetes Snowflake](https://github.com/jessfraz/k8s-snowflake) project.  


## Why

the `client-go` library can be rough to get up and running with. Dependencies are the primary problem, and structure layout has changed between versions.

`client-go` version `v5.x` has landed with strong support for Kubernetes 1.8, and hopefully is the beginning of smooth client upgrades going forward.

## Vendoring

This project chooses to use [dep](https://github.com/golang/dep) as the dependency manager. 
