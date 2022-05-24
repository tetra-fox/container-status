package main

import (
	"context"
	"encoding/json"
	"golang.org/x/exp/slices"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type responseContainer struct {
	ID          string   `json:"id"`
	PrimaryName string   `json:"name"`
	Names       []string `json:"names"`
	State       string   `json:"state"`
	Status      string   `json:"status"`
	Image       string   `json:"image"`
	ImageHash   string   `json:"image_hash"`
}

type response struct {
	Containers []responseContainer `json:"containers"`
	Time       int64               `json:"time"`
}

func main() {
	router := gin.Default()
	router.GET("/", listContainers)
	router.GET("/:names", listContainers)

	router.Run()
}

// Helpers
func getContainers() ([]responseContainer, error) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{All: true})
	if err != nil {
		return nil, err
	}

	output := make([]responseContainer, len(containers))
	for i, container := range containers {
		output[i] = responseContainer{
			ID:        container.ID,
			State:     container.State,
			Status:    container.Status,
			Image:     container.Image,
			ImageHash: container.ImageID[7:],
		}
		for _, name := range container.Names {
			output[i].Names = append(output[i].Names, strings.TrimPrefix(name, "/"))
		}
		output[i].PrimaryName = output[i].Names[0]
	}

	return output, nil
}

func containersToJson(containers []responseContainer) ([]byte, error) {
	return json.Marshal(response{Containers: containers, Time: time.Now().UnixMilli()})
}

// Routes
func listContainers(c *gin.Context) {
	containers, err := getContainers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if c.Param("names") != "" {
		names := strings.Split(c.Param("names"), ",")
		filteredContainers := make([]responseContainer, 0)
		for _, container := range containers {
			for _, name := range names {
				if slices.Contains(container.Names, name) {
					filteredContainers = append(filteredContainers, container)
				}
			}
		}
		containers = filteredContainers
	}

	responseJson, err := containersToJson(containers)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Data(http.StatusOK, gin.MIMEJSON, responseJson)
}
