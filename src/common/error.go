package common

import "errors"

var FileExistsCode = -1
var FileExistsException = errors.New("ERROR: File exists").Error()

var FileNotExistsBeforeCreateChunkCode = -2
var FileNotExistsBeforeCreateChunkException = errors.New("ERROR: File not exists before create chunk").Error()
