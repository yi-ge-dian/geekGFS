package common

import "errors"

// StatusCode 状态码
type StatusCode struct {
	Value     int
	Exception string
}

var FileExistsValue = -1
var FileExistsException = errors.New("ERROR: File exists").Error()

var FileNotExistsBeforeCreateChunkValue = -2
var FileNotExistsBeforeCreateChunkException = errors.New("ERROR: File not exists before create chunk").Error()

var ChunkExistsValue = -3
var ChunkExistsException = errors.New("ERROR: File exists").Error()
