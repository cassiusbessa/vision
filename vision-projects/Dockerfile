# Etapa 1: Construir o aplicativo
FROM jelastic/maven:3.9.4-openjdk-22.ea-b17 AS build
WORKDIR /usr/src/app

# Copiar pom.xml e o código-fonte
COPY pom.xml .
COPY src ./src

# Construir o aplicativo
RUN mvn package -DskipTests

# Etapa 2: Criar a imagem final
FROM openjdk:22-slim
WORKDIR /app

# Copiar o JAR construído da etapa de build
COPY --from=build /usr/src/app/target/vision-project-1.0-SNAPSHOT.jar vision-project.jar

# Definir o comando para executar a aplicação
ENTRYPOINT ["java", "-jar", "vision-project.jar"]

# Expor a porta padrão usada pela aplicação Spring Boot
EXPOSE 8080
