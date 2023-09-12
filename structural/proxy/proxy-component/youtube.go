package proxycomponent

import (
	"fmt"

	"golang.org/x/exp/slices"
)

type Youtube struct {
	SearchKey string
}

func (youtube *Youtube) GetLists() []string {
	fmt.Println("Getting the youtube playlists")
	return []string{
		"Balle Balle",
		"Shava shava",
		"Nahi samjhoge toh code nahi kar paoge",
	}
}

func (youtube *Youtube) GetListById(id string) string {
	fmt.Println("Get list by id called with id", id)
	sliceStrings := []string{
		"Balle Balle",
		"Shava Shava",
		"Nahi samjhoge toh code nahi kar paoge",
	}
	foundIndex := slices.IndexFunc[[]string, string](sliceStrings, func(s string) bool {
		return s == id
	})
	if foundIndex == -1 {
		return "nothing found"
	}
	return sliceStrings[foundIndex]
}

func (youtube *Youtube) DownloadVideo(id string) string {
	return fmt.Sprintf("The file is getting downloaded for id %s", id)
}
