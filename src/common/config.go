package common

import (
	"strconv"
)

// StatusCode 状态码
type StatusCode struct {
	value     int
	exception string
}

// GFSConfig 状态配置
type GFSConfig struct {
	chunkSize            int
	chunkServerLocations []string
	chunkServerRoot      string
}

func (c *GFSConfig) Start() {
	c.chunkSize = 64
	for i := 30002; i <= 30006; i++ {
		c.chunkServerLocations = append(c.chunkServerLocations, strconv.Itoa(i))
	}
	c.chunkServerRoot = "Root"
}

func (c *GFSConfig) GetChunkServerLocations() []string {
	return c.chunkServerLocations
}
