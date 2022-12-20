package task

import (
	"context"
	"io"
	"log"
	"os"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/docker/go-connections/nat"
	"github.com/google/uuid"
)

type State int

const (
	Pending State = iota
	Scheduled
	Completed
	Running
	Failed
)

type Task struct {
	ID            uuid.UUID
	ContainerID   string
	Name          string
	State         State
	Image         string
	Memory        int64
	Disk          int64
	Cpu           float64
	ExposedPorts  nat.PortSet
	PortBindings  map[string]string
	RestartPolicy string
	StartTime     time.Time
	FinishTime    time.Time
}

type TaskEvent struct {
	ID        uuid.UUID
	State     State
	Timestamp time.Time
	Task      Task
}

type Config struct {
	Name         string
	AttachStdin  bool
	AttachStdout bool
	AttachStderr bool
	Cmd          []string
	Image        string
	// ExposedPorts list of ports exposed
	ExposedPorts nat.PortSet
	Cpu          float64
	Memory       int64
	Disk         int64
	Env          []string
	// RestartPolicy for the container ["", "always", "unless-stopped", "on-failure"]
	RestartPolicy string
}

func NewConfig(t *Task) *Config {
	return &Config{
		Name:          t.Name,
		ExposedPorts:  t.ExposedPorts,
		Image:         t.Image,
		Cpu:           t.Cpu,
		Memory:        t.Memory,
		Disk:          t.Disk,
		RestartPolicy: t.RestartPolicy,
	}
}

type Docker struct {
	Client *client.Client
	Config Config
}

func NewDocker(c *Config) *Docker {
	dc, _ := client.NewClientWithOpts(client.FromEnv)
	return &Docker{
		Client: dc,
		Config: *c,
	}
}

type DockerResult struct {
	Error       error
	Action      string
	ContainerId string
	Result      string
}

func (docker *Docker) Run() DockerResult {
	ctx := context.Background()
	reader, err := docker.Client.ImagePull(ctx, docker.Config.Image, types.ImagePullOptions{})

	if err != nil {
		log.Printf("Error pulling image %v: %v\n", docker.Config.Image, err)
	}

	io.Copy(os.Stdout, reader)

	rp := container.RestartPolicy{
		Name: docker.Config.RestartPolicy,
	}

	r := container.Resources{
		Memory: docker.Config.Memory,
	}

	cc := container.Config{
		Image: docker.Config.Image,
		Env:   docker.Config.Env,
	}

	hc := container.HostConfig{
		RestartPolicy:   rp,
		Resources:       r,
		PublishAllPorts: true,
	}

	resp, err := docker.Client.ContainerCreate(
		ctx,
		&cc,
		&hc,
		nil,
		nil,
		docker.Config.Name,
	)

	if err != nil {
		log.Printf("Error creating container using image %s:%v\n", docker.Config.Image, err)
		return DockerResult{Error: err}
	}

	err = docker.Client.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})

	if err != nil {
		log.Printf("Error starting container %s:%v\n", resp.ID, err)
		return DockerResult{Error: err}
	}

	out, err := docker.Client.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true, ShowStderr: true})

	if err != nil {
		log.Printf("Error getting logs for container %s: %v\n", resp.ID, err)
		return DockerResult{Error: err}
	}

	stdcopy.StdCopy(os.Stdout, os.Stderr, out)

	return DockerResult{
		ContainerId: resp.ID,
		Action:      "start",
		Result:      "success",
	}
}

func (docker *Docker) Stop(containerID string) DockerResult {
	ctx := context.Background()
	log.Printf("Attempting to stop container %v", containerID)
	err := docker.Client.ContainerStop(ctx, containerID, nil)
	if err != nil {
		panic(err)
	}

	removeOptions := types.ContainerRemoveOptions{
		RemoveVolumes: true,
		RemoveLinks:   false,
		Force:         false,
	}

	err = docker.Client.ContainerRemove(ctx, containerID, removeOptions)

	if err != nil {
		panic(err)
	}

	return DockerResult{Action: "stop", Result: "success", Error: nil}
}
