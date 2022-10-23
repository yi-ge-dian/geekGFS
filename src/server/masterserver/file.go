package masterserver

type File struct {
	filePath string
	chunks   map[string]Chunk
}

func (f *File) SetFilePath(filePath string) {
	f.filePath = filePath
}

func (f *File) GetChunks() map[string]Chunk {
	return f.chunks
}
