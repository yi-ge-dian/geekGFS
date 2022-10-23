package masterserver

type MetaData struct {
	locations         []string
	files             map[string]File //filepath to file
	chunkHandleToFile map[string]File
	locationDist      map[string][]string
}

func (md *MetaData) SetLocations(locations []string) {
	md.locations = locations
}

func (md *MetaData) GetLatestChunk(filePath string, latestChunkHandle string) {

}
