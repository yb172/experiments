# Testing tools

This directory contains tools that test application beyond scope of one service, e.g. tests that involve multiple services or requrie external dependency.

## Naming convention

There could be a significant confusion in terminology. Take "integration" tests for instance: integration tests kind of supposed to test integration between components but what is considered as "component"? Class? Module? Service? System?

To avoid this confusion let's pretend there are only two groups of tests:

* Tests that test only once deployable unit - such tests could be run locally and require only language build system (gradle / go / npm), doesn't require docker or external dependencies with state. These would be unit tests and BDD-style tests that stand up service, run DB migration scripts on clean DB, mock any other service, execute series of API calls and verify it looks good. Such tests would provide test coverage metric to see which parts of code are verified.
* Tests that need external dependencies or other services.

## Approaches (strategy)

Let's say we have our generator service which consist of `gateway` that depends on `word`, `number`, and `internets` services. There are two fundamentally different approaches to type 2 testing:

* Have special environment ("dev" / "test" / "beta" / something similar). Tests are executed as part of CI / CD pipeline, e.g. Jenkins builds docker image, pushes it, then waits until it would be deployed to the environment and then runs tests.
* Create whole setup from scratch, run tests, tear setup down

As you might know this is "pets vs cattle" approach to infrastructure and kubernetes is all about "cattle".

So there should be something that would create whole setup, run tests and delete setup. It could be part of CI / CD pipeline or it could be run periodically.

## Tools (tactics)

Here I would try to explore and come up with some solution for cattle-type wholesale testing. Tools that looks relevant:

* [Skaffold](https://skaffold.dev)
* [Prow](https://github.com/kubernetes/test-infra/tree/master/prow)

As of now there are services, they are dockerized and skaffold is configured to launch it on minikube. There is also linkerd configured.

### Local development

One thing that would be convenient for me for local development is to have some utility that generates load on a service so that I could see some stats in linkerd.

It would be also nice if there would be a way to interactively change the load.

And this is not difficult to do: the simplest version would be to have go program that periodically triggers our app endpoint and kube spec where we specify number of replicas to scale the thing. That is easy to do but that is not interactive.

So another approach is to have same go program which does the same but runs locally. Then it could easily be tuned interactively, literally with up and down arrow keys. The only question is how to get service endpoint.

Mix of two would be programs running in the kube listening to the pubsub topic and program that runs locally which publishes messages to this pubsub. Though question of figuring out endpoint is still relevant: this time it would be pubsub endpoint (assuming it also runs on kube cluster).

### Test case

Another thing we could have is to run test that asks endpoint and verifies result is not empty (see [#3](https://github.com/yb172/experiments/issues/3)). Ok, that would work for local development, but whole part of setting up / tearing down is missed.
