# Test scenarios

These scenarios verify that end user scenarios are working correctly. That means that these tests require all components of our system deployed and working.

Which leads to an interesting question: how such test could be performed?

## Problem

It is easy to write test that verifies service-scoped behavior: write unit test or BDD-style test. If service has external dependencies (other services or database) - they could be mocked.

However in this scenario we have multiple services talking to each other and we would like to test their interaction. Well, it would be more correct to say that we would like to test the feature which involves multiple services.

So our first requirement for test: **Services required for testing (and their deps) should be up and running**. Meaning we have to have real cluster and do real deployment for these tests.

## After & Before

One way to do that is to have "test" (dev/staging/beta/whatever) environment where the most recent version of software is deployed and then tests are executed against it. This is a good solution, but such pattern would cause some inconvenience: tests would be executed *after* PR has been merged, so if PR breaks the test something would have to be done to fix it. It is more convenient to have such tests executed for every PR *before* it is merged. And it is still possible with this approach, though requires some coordination to queue test requests (if two PRs are being opened simultaneously).

This would be the second requriement for our test: **Tests should execute before PR is merged**.

Additional bonus if tests would be possible to execute locally.

## Tools to do that

It looks like [prow](https://github.com/kubernetes/test-infra/tree/master/prow) could help with running jobs. And first test of such job would be to either:

* Synchronize with shared environment (i.e. services running in it are ones built from PR code)
* Create a new environment

Second point seems easier and there are multiple choices:

* Use [skaffold](http://skaffold.dev)
* Use docker-compose + [kompose](http://kompose.io)
* Use something custom
