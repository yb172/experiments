FROM openjdk:8-alpine AS builder
COPY gradle ./gradle
COPY build.gradle settings.gradle gradlew ./
RUN ./gradlew getDeps
COPY . .
RUN ./gradlew build

FROM openjdk:8-alpine
COPY --from=builder ./build/libs/java-multistage.jar app.jar
CMD java -jar app.jar