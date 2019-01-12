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

## Speeding things up

Build process however is annoying: everytime gradle downloads itself and all dependencies
which could take a while (e.g. `time docker build . -t test-app` shows 1 minute), no docker
caching is leveraged.

It would be nice if we can separate dependencies downloading and actual build.

To do so we first need to add task to download deps to `build.gradle`:

```groovy
task getDeps(type: Copy) {
    from sourceSets.test.runtimeClasspath
    into 'runtime/'
}
```

And then update Dockerfile to following:

```Docker
FROM openjdk:8-alpine AS builder
COPY gradle ./gradle
COPY build.gradle settings.gradle gradlew ./
RUN ./gradlew getDeps
COPY . .
RUN ./gradlew build

FROM openjdk:8-alpine
COPY --from=builder ./build/libs/java-multistage.jar app.jar
CMD java -jar app.jar
```

What that is doing it is first copying all gradle files, running task to download deps
and only then copies the rest and does the build.

That way once we've downloaded gradle and resolved dependencies they are cached.

Before change every docker build took ~1 min.
After the change first build takes about the same, but subsequent builds are taking ~16s
(if build files are not updated)
