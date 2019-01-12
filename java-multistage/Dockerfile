FROM openjdk:8-alpine AS builder
COPY . .
RUN ./gradlew build

FROM openjdk:8-alpine
COPY --from=builder ./build/libs/java-multistage.jar app.jar
CMD java -jar app.jar