# Auto Deployer

A CLI tool that automates the process of deploying applications to Kubernetes from GitHub repositories.

## Features

- Clones GitHub repositories
- Builds Docker images based on project type (Java/Node.js)
- Pushes images to Docker Hub
- Deploys applications to Kubernetes (Minikube)
- Interactive CLI prompts for configuration

## Prerequisites

- Go 1.21 or higher
- Docker installed and running
- Minikube installed and running
- kubectl configured to use Minikube
- Docker Hub account
- GitHub account (for private repositories)

### Setting up Minikube

1. Install Minikube:
   ```bash
   # For Windows (using Chocolatey)
   choco install minikube

   # For macOS (using Homebrew)
   brew install minikube

   # For Linux
   curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
   sudo install minikube-linux-amd64 /usr/local/bin/minikube
   ```

2. Start Minikube:
   ```bash
   minikube start
   ```

3. Enable required addons:
   ```bash
   minikube addons enable ingress
   ```

4. Verify installation:
   ```bash
   minikube status
   kubectl get nodes
   ```

### GitHub Token Setup (for private repositories)

1. Go to GitHub Settings > Developer settings > Personal access tokens
2. Generate a new token with 'repo' scope
3. Save the token securely - you'll need it when using the tool with private repositories

## Installation

```bash
# Clone the repository
git clone https://github.com/yourusername/auto-deployer

# Change to the project directory
cd auto-deployer

# Build the tool
go build -o auto-deployer

# Make it available globally (Linux/macOS)
sudo mv auto-deployer /usr/local/bin/
```

## Usage

Simply run the tool and follow the interactive prompts:

```bash
auto-deployer
```

The tool will ask for:
1. GitHub repository URL
2. Project type (Java/Node.js)
3. Docker Hub credentials
4. Image name
5. Application port

## Supported Project Types

### Java
- Expects a Maven project
- Builds using Maven
- Runs the resulting JAR file

### Node.js
- Expects a package.json
- Installs dependencies using npm
- Runs using npm start

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
