# Dockerized java application with multi-stage build

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
test-app                             latest              e60c168bf5c7        2 minutes ago       798MB
```

Simple hello world takes 800MB