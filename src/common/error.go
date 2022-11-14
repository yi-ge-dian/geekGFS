package common

import "errors"

// StatusCode 状态码
type StatusCode struct {
	Value     int
	Exception string
}

var FileExistsValue = -2
var FileExistsException = errors.New("ERROR: File exists").Error()

var FileNotExistsValue = -3
var FileNotExistsException = errors.New("ERROR: File not exists").Error()

var FileNotExistsBeforeCreateChunkValue = -4
var FileNotExistsBeforeCreateChunkException = errors.New("ERROR: File not exists before create chunk").Error()

var ChunkExistsValue = -5
var ChunkExistsException = errors.New("ERROR: Chunk exists").Error()

var FilePathNotExistsValue = -6
var FilePathNotExistsException = errors.New("ERROR: Filepath not exists").Error()
