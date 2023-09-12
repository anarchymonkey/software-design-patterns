package main

import (
	"fmt"

	proxycomponent "github.com/anarchymonkey/design-patterns/structural/proxy/proxy-component"
	"github.com/anarchymonkey/design-patterns/structural/proxy/service"
)

func main() {
	fmt.Println("Main")

	youtube := proxycomponent.Youtube{}

	youtubeProxy := proxycomponent.YoutubeProxy{
		YT: &youtube,
	}

	youtubeManager := service.YoutubePlayerManager{
		Youtube: &youtubeProxy,
	}

	youtubeManager.GetVideos()
	fmt.Println("GetVideoByID response", youtubeManager.GetVideoByID("Balle Balle"))
	fmt.Println("GetVideoByID response", youtubeManager.GetVideoByID("Balle Balle, Shava Shava"))

	fmt.Println(youtubeManager.Youtube.DownloadVideo("Balle Balle"))
}
