package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

func deployApplication(config *Config) error {
	// Create temporary directory for cloning
	tempDir, err := os.MkdirTemp("", "auto-deployer-*")
	if err != nil {
		return fmt.Errorf("failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Clone repository
	fmt.Println("Cloning repository...")
	if err := cloneRepo(config, tempDir); err != nil {
		return fmt.Errorf("failed to clone repository: %v", err)
	}

	// Build and push Docker image
	fmt.Println("Building Docker image...")
	if err := buildAndPushImage(config, tempDir); err != nil {
		return fmt.Errorf("failed to build/push Docker image: %v", err)
	}

	// Deploy to Kubernetes
	fmt.Println("Deploying to Kubernetes...")
	if err := deployToKubernetes(config); err != nil {
		return fmt.Errorf("failed to deploy to Kubernetes: %v", err)
	}

	fmt.Println("Deployment completed successfully!")
	return nil
}

func cloneRepo(config *Config, dir string) error {
	cloneOpts := &git.CloneOptions{
		URL:      config.GitHubRepo,
		Progress: os.Stdout,
	}

	// Add authentication for private repositories
	if config.GitHubUsername != "" && config.GitHubToken != "" {
		cloneOpts.Auth = &http.BasicAuth{
			Username: config.GitHubUsername,
			Password: config.GitHubToken,
		}
	}

	_, err := git.PlainClone(dir, false, cloneOpts)
	return err
}

func buildAndPushImage(config *Config, dir string) error {
	// Docker login
	loginCmd := exec.Command("docker", "login", "-u", config.DockerUsername, "--password-stdin")
	loginCmd.Stdin = strings.NewReader(config.DockerPassword)
	if err := loginCmd.Run(); err != nil {
		return fmt.Errorf("docker login failed: %v", err)
	}

	// Build image
	buildCmd := exec.Command("docker", "build", "-t",
		fmt.Sprintf("%s/%s", config.DockerUsername, config.ImageName), ".")
	buildCmd.Dir = dir
	buildCmd.Stdout = os.Stdout
	buildCmd.Stderr = os.Stderr
	if err := buildCmd.Run(); err != nil {
		return fmt.Errorf("docker build failed: %v", err)
	}

	// Push image
	pushCmd := exec.Command("docker", "push",
		fmt.Sprintf("%s/%s", config.DockerUsername, config.ImageName))
	pushCmd.Stdout = os.Stdout
	pushCmd.Stderr = os.Stderr
	return pushCmd.Run()
}

func deployToKubernetes(config *Config) error {
	// Create deployment YAML
	deploymentYAML := fmt.Sprintf(`
apiVersion: apps/v1
kind: Deployment
metadata:
  name: %s
spec:
  replicas: 1
  selector:
    matchLabels:
      app: %s
  template:
    metadata:
      labels:
        app: %s
    spec:
      containers:
      - name: %s
        image: %s/%s
        ports:
        - containerPort: %s
---
apiVersion: v1
kind: Service
metadata:
  name: %s
spec:
  type: NodePort
  ports:
  - port: %s
    targetPort: %s
  selector:
    app: %s
`, config.ImageName, config.ImageName, config.ImageName,
		config.ImageName, config.DockerUsername, config.ImageName,
		config.Port, config.ImageName, config.Port, config.Port, config.ImageName)

	// Write deployment YAML
	yamlFile := filepath.Join(os.TempDir(), "deployment.yaml")
	if err := os.WriteFile(yamlFile, []byte(deploymentYAML), 0644); err != nil {
		return err
	}

	// Apply deployment
	applyCmd := exec.Command("kubectl", "apply", "-f", yamlFile)
	applyCmd.Stdout = os.Stdout
	applyCmd.Stderr = os.Stderr
	return applyCmd.Run()
}
