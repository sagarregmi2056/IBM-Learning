package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

type Config struct {
	GitHubRepo     string
	GitHubUsername string
	GitHubToken    string
	DockerUsername string
	DockerPassword string
	ImageName      string
	ProjectType    string
	Port           string
}

var version = "1.0.0"

func main() {
	var rootCmd = &cobra.Command{
		Version: version,
		Use:     "auto-deployer",
		Short:   "Automated deployment tool for Kubernetes",
		Long: `🚀 Kubernetes Auto Deployment Tool
Developed by: Sagar Regmi
GitHub: https://github.com/sagarregmi2056

This tool automates the process of cloning code, building Docker images, and deploying to Kubernetes.

This tool will help you:
1. Clone a GitHub repository (public or private)
2. Build Docker images based on project type (Java/Node.js)
3. Push images to Docker Hub
4. Deploy to Kubernetes (Minikube)

Prerequisites:
- Docker Desktop running
- Minikube installed and running
- kubectl configured
- Docker Hub account`,
		Run: run,
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run(cmd *cobra.Command, args []string) {
	config := &Config{}

	// Clear screen and show welcome message
	fmt.Print("\033[H\033[2J") // Clear screen
	fmt.Println("\n🚀 Kubernetes Auto Deployment Tool")
	fmt.Println("Developed by: Sagar Regmi")
	fmt.Println("GitHub: https://github.com/sagarregmi2056")
	fmt.Println("\n=============================================")
	fmt.Println("Starting deployment process...")
	fmt.Println("=============================================")

	// GitHub Repository
	prompt := promptui.Prompt{
		Label:    "GitHub Repository URL",
		Validate: validateGitURL,
	}
	result, err := prompt.Run()
	if err != nil {
		if err == promptui.ErrInterrupt {
			fmt.Println("\nOperation cancelled by user")
			os.Exit(0)
		}
		log.Fatalf("Prompt failed: %v\n", err)
	}
	config.GitHubRepo = result

	// GitHub Username
	prompt = promptui.Prompt{
		Label: "GitHub Username (for private repos, press enter if public repo)",
	}
	result, err = prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed: %v\n", err)
	}
	config.GitHubUsername = result

	// GitHub Token (if username was provided)
	if config.GitHubUsername != "" {
		prompt = promptui.Prompt{
			Label: "GitHub Personal Access Token",
			Mask:  '*',
		}
		result, err = prompt.Run()
		if err != nil {
			log.Fatalf("Prompt failed: %v\n", err)
		}
		config.GitHubToken = result
	}

	// Check Minikube status
	if err := checkPrerequisites(); err != nil {
		log.Fatalf("Prerequisites check failed: %v\n", err)
	}

	// Project Type
	typeSelect := promptui.Select{
		Label: "Select Project Type",
		Items: []string{"java", "nodejs"},
	}
	_, result, err = typeSelect.Run()
	if err != nil {
		log.Fatalf("Prompt failed: %v\n", err)
	}
	config.ProjectType = result

	// Docker Hub Username
	prompt = promptui.Prompt{
		Label: "Docker Hub Username",
	}
	result, err = prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed: %v\n", err)
	}
	config.DockerUsername = result

	// Docker Hub Password
	prompt = promptui.Prompt{
		Label: "Docker Hub Password",
		Mask:  '*',
	}
	result, err = prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed: %v\n", err)
	}
	config.DockerPassword = result

	// Image Name
	prompt = promptui.Prompt{
		Label: "Docker Image Name",
	}
	result, err = prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed: %v\n", err)
	}
	config.ImageName = result

	// Port
	prompt = promptui.Prompt{
		Label:    "Application Port",
		Default:  "8080",
		Validate: validatePort,
	}
	result, err = prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed: %v\n", err)
	}
	config.Port = result

	// Start deployment process
	if err := deployApplication(config); err != nil {
		log.Fatalf("Deployment failed: %v\n", err)
	}
}

func validateGitURL(input string) error {
	if len(input) < 10 {
		return fmt.Errorf("URL too short")
	}
	if !strings.HasPrefix(input, "https://") {
		return fmt.Errorf("URL must start with https://")
	}
	if !strings.Contains(input, "github.com/") {
		return fmt.Errorf("must be a github repository URL")
	}
	if strings.Count(input, "/") < 4 {
		return fmt.Errorf("invalid repository url format. expected: https://github.com/username/repository")
	}
	return nil
}

func validatePort(input string) error {
	if len(input) < 2 || len(input) > 5 {
		return fmt.Errorf("port must be between 2 and 5 characters")
	}
	return nil
}

func checkPrerequisites() error {
	// Check if Docker is running
	dockerCmd := exec.Command("docker", "info")
	if err := dockerCmd.Run(); err != nil {
		return fmt.Errorf("docker is not running or not installed: %v", err)
	}

	// Check if Minikube is installed
	minikubeCmd := exec.Command("minikube", "status")
	output, err := minikubeCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("minikube is not running or not installed: %v\noutput: %s", err, output)
	}

	// Check if kubectl is installed
	kubectlCmd := exec.Command("kubectl", "version", "--client")
	if err := kubectlCmd.Run(); err != nil {
		return fmt.Errorf("kubectl is not installed: %v", err)
	}

	return nil
}
