# Dockerized java application with multi-stage build

## Initial version

It is not difficult to create a docker image of java application.

Let's start by creating java application itself by running

```bash
gradle init --type java-application
```

(also I had to change repository to `mavenCentral` and added `jar` section in `gradle.build`)

Then we could create dockerfile which could be as simple as

```Docker
FROM openjdk:8

COPY . .

RUN ./gradlew build

CMD ./gradlew run
```

Which we could build using command

```bash
docker build . -t test-app
```

and then execute by running

```bash
docker run test-app
```

There is however a small problem:

```text
test-app            latest              e60c168bf5c7        2 minutes ago       798MB
```

Simple hello world takes 800MB

## Using alpine

First thing we could do - use alpine image: `FROM openjdk:8-alpine`. This simple change
helps to save ~500MB:

```text
test-app            latest              940fe5f97601        10 seconds ago      277MB
```

## Multistage build

However we can do even better! All that we really need is the jar that we're using in `CMD`
Docker's multi stage builds can help with that, Dockerfile would look like following:

```Docker
FROM openjdk:8-alpine AS builder
COPY . .
RUN ./gradlew build

FROM openjdk:8-alpine
COPY --from=builder ./build/libs/java-multistage.jar app.jar
CMD java -jar app.jar
```

Run `docker build . -t test-app` again and we have

```text
test-app            latest              e8a7aadfa0f2        11 seconds ago      103MB
```

8 times less than it was originally!
