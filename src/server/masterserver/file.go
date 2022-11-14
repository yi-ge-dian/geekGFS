package masterserver

type File struct {
	filePath       string            // 文件路径（命名空间）
	chunks         map[string]*Chunk // key:chunkHandle value:Chunk
	chunkHandleSet []string          // 存储chunkHandle的集合
}

func (f *File) GetChunks() *map[string]*Chunk {
	return &f.chunks
}

func NewFile(filePath *string) *File {
	return &File{
		filePath: *filePath,
		chunks:   make(map[string]*Chunk, 0),
	}
}
