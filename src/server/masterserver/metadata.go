package masterserver

import (
	cm "GeekGFS/src/common"
	"math/rand"
	"time"
)

type MetaData struct {
	locations         []string            //
	files             map[string]File     // key:filepath  value:File
	chunkHandleToFile map[string]File     // key:chunkHandle value:File
	locationDist      map[string][]string // key: ***
}

func (md *MetaData) GetFiles() *map[string]File {
	return &md.files
}

func NewMetaData(locations []string) *MetaData {
	return &MetaData{locations: locations}
}

func (md *MetaData) Init() {
	md.files = make(map[string]File, 0)
	md.chunkHandleToFile = make(map[string]File, 0)
	md.locationDist = make(map[string][]string, 0)
}

// ChooseChunkServerLocations 选择存放 chunkServer的位置
func (md *MetaData) ChooseChunkServerLocations(chunkServerLocations *[]string) {
	//logger := gologger.GetLogger(gologger.CONSOLE, gologger.ColoredLog)
	gfsConfig := cm.NewGFSConfig(cm.GFSChunkSize, cm.GFSChunkServerLocations, cm.GFSChunkServerRoot)
	total := len(gfsConfig.ChunkServerLocations())
	rand.Seed(time.Now().Unix())
	index := rand.Int()
	for i := 0; i < 3; i++ {
		//logger.Debug(gfsConfig.ChunkServerLocations[(index+i)%total])
		*chunkServerLocations = append(*chunkServerLocations, gfsConfig.ChunkServerLocations()[(index+i)%total])
	}
}

// GetLatestChunkHandle 得到最新的chunkHandle
func (md *MetaData) GetLatestChunkHandle(filePath *string) string {
	// 如果找到了该文件
	if file, ok := md.files[*filePath]; ok {
		if len(file.chunkHandleSet) > 0 {
			return file.chunkHandleSet[len(file.chunkHandleSet)-1]
		}
	}
	return "-2"
}

// CreateNewFile 创建新文件
func (md *MetaData) CreateNewFile(filePath *string, chunkHandle *string, statusCode *cm.StatusCode) {
	//logger := gologger.GetLogger(gologger.CONSOLE, gologger.ColoredLog)
	// 如果已经创建该文件
	md.Init()
	if _, ok := md.files[*filePath]; ok {
		statusCode.Value = cm.FileExistsValue
		statusCode.Exception = cm.FileExistsException + *filePath
		return
	}
	// 插入
	f := NewFile(filePath)
	f.Init()
	md.files[*filePath] = *f
	// 创建chunk
	*chunkHandle = "1" + *chunkHandle
	prevChunkHandle := "-1"
	md.CreateNewChunk(filePath, &prevChunkHandle, chunkHandle, statusCode)
}

// CreateNewChunk 创建新的chunk
func (md *MetaData) CreateNewChunk(filePath *string, prevChunkHandle *string, chunkHandle *string, code *cm.StatusCode) {
	//logger := gologger.GetLogger(gologger.CONSOLE, gologger.ColoredLog)
	file, ok := md.files[*filePath]
	// 如果该文件未被创建，就来创建 chunk 是不科学的
	if !ok {
		code.Value = cm.FileNotExistsBeforeCreateChunkValue
		code.Exception = cm.FileNotExistsBeforeCreateChunkException + *filePath
		return
	}
	// chunk不是新建的
	latestChunkHandle := ""
	if *prevChunkHandle != "-1" {
		latestChunkHandle = md.GetLatestChunkHandle(filePath)
		//logger.Debug(latestChunkHandle)
	}
	// chunk 不是新建的并且与该文件最新的 chunkHandle 对不上号
	if *prevChunkHandle != "-1" && latestChunkHandle != *prevChunkHandle {
		code.Value = cm.ChunkExistsValue
		code.Exception = cm.ChunkExistsException + *filePath + ":" + *chunkHandle
		return
	}
	// 创建chunk
	newChunk := new(Chunk)
	chunks := file.GetChunks()
	(*chunks)[*chunkHandle] = newChunk
	// 拿到新创建的chunk
	chunk := (*chunks)[*chunkHandle]
	// 选择chunk的位置
	var chunkServerLocations []string
	md.ChooseChunkServerLocations(&chunkServerLocations)
	for i := 0; i < len(chunkServerLocations); i++ {
		chunk.locations = append(chunk.locations, chunkServerLocations[i])
	}
	// 更新文件信息
	file.chunkHandleSet = append(file.chunkHandleSet, *chunkHandle)
	// 设置状态码
	code.Value = 0
	code.Exception = "SUCCESS: New Chunk Created"
}
