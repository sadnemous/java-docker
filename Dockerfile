FROM openjdk:21-jdk
EXPOSE 8080
COPY build/libs/docker-demo-0.0.1-SNAPSHOT-plain.jar docker-demo.jar
ENTRYPOINT ["java", "-jar", "/docker-demo.jar"]
