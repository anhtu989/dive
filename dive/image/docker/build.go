go
package docker

import (
	"os"
)

// buildImageFromCli builds a Docker image using the provided build arguments.
// It creates a temporary file to store the image ID and returns the image ID upon successful build.
// The temporary file is automatically removed after the function execution.
func buildImageFromCli(buildArgs []string) (string, error) {
	// Create a temporary file to store the image ID
	iidfile, err := os.CreateTemp("/tmp", "dive.*.iid")
	if err != nil {
		return "", err
	}
	defer os.Remove(iidfile.Name())

	// Append the --iidfile argument to the build arguments
	allArgs := append([]string{"--iidfile", iidfile.Name()}, buildArgs...)

	// Run the Docker build command with the combined arguments
	err = runDockerCmd("build", allArgs...)
	if err != nil {
		return "", err
	}

	// Read the image ID from the temporary file
	imageId, err := os.ReadFile(iidfile.Name())
	if err != nil {
		return "", err
	}

	// Return the image ID as a string
	return string(imageId), nil
}