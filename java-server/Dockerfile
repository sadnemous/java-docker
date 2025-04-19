# syntax=docker/dockerfile:1
FROM openjdk:17-jdk-slim
COPY build/libs/docker-demo-0.0.1-SNAPSHOT.jar app.jar
ENTRYPOINT ["java", "-jar", "/app.jar"]