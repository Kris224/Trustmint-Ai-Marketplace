package docker

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

type TrainOptions struct {
	DatasetPath string
	ConfigPath  string
	OutputPath  string
	ModelName   string
	Framework   string
	CPU         string
	Memory      string
	UseGPU      bool
}

// RunTrainingContainer spins up a Docker container for training
func RunTrainingContainer(opts TrainOptions) error {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}

	// ✅ Choose correct Docker image
	image := fmt.Sprintf("trustmint/train-%s:latest", opts.Framework)
	if opts.Framework == "pytorch" {
		image = "trustmint/train-pytorch:latest"
	} else if opts.Framework == "tensorflow" {
		image = "trustmint/train-tensorflow:latest"
	}

	fmt.Println("🚀 Starting training container with image:", image)

	// ✅ Ensure output directory exists
	if err := os.MkdirAll(opts.OutputPath, 0755); err != nil {
		return err
	}

	// ✅ Convert all paths to absolute
	datasetAbs, err := filepath.Abs(opts.DatasetPath)
	if err != nil {
		return fmt.Errorf("invalid dataset path: %v", err)
	}
	configAbs, err := filepath.Abs(opts.ConfigPath)
	if err != nil {
		return fmt.Errorf("invalid config path: %v", err)
	}
	outputAbs, err := filepath.Abs(opts.OutputPath)
	if err != nil {
		return fmt.Errorf("invalid output path: %v", err)
	}

	// ✅ Mount files properly
	hostConfig := &container.HostConfig{
		Binds: []string{
			fmt.Sprintf("%s:/workspace/data:ro", datasetAbs),
			fmt.Sprintf("%s:/workspace/%s:ro", configAbs, filepath.Base(configAbs)),
			fmt.Sprintf("%s:/workspace/output", outputAbs),
		},
	}

	// ✅ Handle old containers gracefully
	containerName := fmt.Sprintf("trustmint-train-%s", opts.ModelName)
	_ = cli.ContainerRemove(ctx, containerName, types.ContainerRemoveOptions{Force: true})

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: image,
		Cmd:   []string{"python", "/workspace/train.py", "--config", "/workspace/" + filepath.Base(configAbs), "--data", "/workspace/data", "--output", "/workspace/output"},
		Tty:   false,
	}, hostConfig, nil, nil, containerName)
	if err != nil {
		return err
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return err
	}

	// ✅ Wait for the container to complete
	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			return err
		}
	case <-statusCh:
	}

	// ✅ Capture logs
	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true, ShowStderr: true})
	if err != nil {
		return err
	}
	defer out.Close()

	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, out); err != nil {
		return err
	}
	fmt.Println("📜 Training logs:\n", buf.String())

	return nil
}
