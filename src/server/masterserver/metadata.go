package masterserver

import (
	cm "GeekGFS/src/common"
	"strconv"
)

type MetaData struct {
	locations         []string            //
	files             map[string]File     // filepath to file
	chunkHandleToFile map[string]File     //
	locationDist      map[string][]string //
}

// SetLocations 设置服务器启动位置
func (md *MetaData) SetLocations(locations []string) {
	md.locations = locations
}

// CreateNewFile 创建新文件
func (md *MetaData) CreateNewFile(filePath *string, chunkHandle *string, code *cm.StatusCode) {
	// 如果已经创建该文件
	if _, ok := md.files[*filePath]; ok {
		code.Value = cm.FileExistsCode
		code.Exception = cm.FileExistsException + *filePath
		return
	}
	// 插入
	var f = new(File)
	f.SetFilePath(*filePath)
	md.files[*filePath] = *f
	// 创建chunk
	*chunkHandle = "1" + *chunkHandle
	prevChunkHandle := "-1"
	md.CreateNewChunk(filePath, &prevChunkHandle, chunkHandle, code)
}

// CreateNewChunk 创建新的chunk
func (md *MetaData) CreateNewChunk(filePath *string, prevChunkHandle *string, chunkHandle *string, code *cm.StatusCode) {
	// 如果该文件未被创建，就来创建 chunk 是不科学的
	if _, ok := md.files[*filePath]; !ok {
		code.Value = cm.FileNotExistsBeforeCreateChunkCode
		code.Exception = cm.FileNotExistsBeforeCreateChunkException + *filePath
		return
	}
	latestChunk := "-2"
	if *prevChunkHandle != strconv.Itoa(-1) {
		md.GetLatestChunk(filePath, &latestChunk)
	}
}

// GetLatestChunk 得到最新数据块
func (md *MetaData) GetLatestChunk(filePath *string, latestChunk *string) {
	if file, ok := md.files[*filePath]; ok {
		if len(file.GetChunks()) > 0 {
			file.GetChunks()
		}
	}
}
