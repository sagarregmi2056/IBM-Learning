package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func createDockerfile(projectType, dir string) error {
	var dockerfile string

	switch projectType {
	case "java":
		dockerfile = `FROM maven:latest AS build
WORKDIR /app
COPY . .
RUN mvn clean package -DskipTests

FROM openjdk:11-jre-slim
COPY --from=build /app/target/*.jar app.jar
EXPOSE ${PORT}
ENTRYPOINT ["java","-jar","/app.jar"]`

	case "nodejs":
		dockerfile = `FROM node:latest
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
EXPOSE ${PORT}
CMD ["npm", "start"]`

	default:
		return fmt.Errorf("unsupported project type: %s", projectType)
	}

	return os.WriteFile(filepath.Join(dir, "Dockerfile"), []byte(dockerfile), 0644)
}
