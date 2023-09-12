package proxycomponent

type YoutubeProxy struct {
	YT *Youtube
}

func (yt_proxy *YoutubeProxy) GetLists() []string {
	return yt_proxy.YT.GetLists()
}

func (yt_proxy *YoutubeProxy) GetListById(id string) string {
	return yt_proxy.YT.GetListById(id)
}

func (yt_proxy *YoutubeProxy) DownloadVideo(id string) string {
	return yt_proxy.YT.DownloadVideo(id)
}
