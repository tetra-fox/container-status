package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type responseContainer struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

type response struct {
	Containers []responseContainer `json:"containers"`
	Time       int64 `json:"time"`
}

func main() {
	router := gin.Default()
	router.GET("/", listContainers)
	router.GET("/:name", getContainerByName)

	router.Run("localhost:80")
}

func getContainers() ([]byte, error) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err;
	}

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{All: true})
	if err != nil {
		return nil, err;
	}

	sanitizedContainers := make([]responseContainer, len(containers))
	for i, container := range containers {
		sanitizedContainers[i] = responseContainer{container.Names[0][1:], container.State}
	}

	responseJson, err := json.Marshal(response{sanitizedContainers, time.Now().UnixMilli()})
	if err != nil {
		return nil, err;
	}

	return responseJson, nil;
}

func listContainers(c *gin.Context) {
	responseJson, err := getContainers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Data(http.StatusOK, gin.MIMEJSON, responseJson)
}

func getContainerByName(c *gin.Context) {
	responseJson, err := getContainers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var containers response
	err = json.Unmarshal(responseJson, &containers)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for _, container := range containers.Containers {
		if container.Name == c.Param("name") {
			responseJson, err := json.Marshal(response{[]responseContainer{container}, time.Now().UnixMilli()})
			if err != nil {
				panic(err)
			}
			c.Data(http.StatusOK, gin.MIMEJSON, responseJson)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "container not found"})
}