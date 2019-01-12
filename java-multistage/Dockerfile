FROM openjdk:8

COPY . .

RUN ./gradlew build

CMD java -jar ./build/libs/java-multistage.jar
