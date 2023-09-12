package service

type YoutubePlayerService interface {
	GetLists() []string
	GetListById(id string) string
	DownloadVideo(id string) string
}

type YoutubePlayerManager struct {
	Youtube YoutubePlayerService
}

func (ytpm *YoutubePlayerManager) GetVideos() []string {
	return ytpm.Youtube.GetLists()
}

func (ytpm *YoutubePlayerManager) InitiateDownload(id string) string {
	return ytpm.Youtube.DownloadVideo(id)
}

func (ytpm *YoutubePlayerManager) GetVideoByID(id string) string {
	return ytpm.Youtube.GetListById(id)
}
