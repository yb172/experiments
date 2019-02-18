# Microservices example

This project shows how application that consists of multiple services that:

* Could be developed locally with convenience: no need in long setup (create database, download this, get those permissions, etc...)
* Is properly tested in an automated fashion so you're confident application works

Example application produces a sequence of words randomly generating either:

* Random word from internet (https://randomword.com/)
* Random word from predefined list
* Random number from 1 to 100

As of now there is no front-end, only `gen-gateway` service provides http api to trigger sequence creation.

## To run

First you would need some software to be installed:

* [Docker](https://docker.io)
* [Minikube](https://github.com/kubernetes/minikube) - should be installed and running
* [Skaffold](https://skaffold.dev)

To run project please run:

```bash
skaffold run --tail
```

Then in a separate terminal:

```bash
minikube service gen-gateway
```
