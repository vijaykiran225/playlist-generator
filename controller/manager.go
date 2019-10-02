package controller

import (
	"encoding/xml"
	"fmt"
	"os"

	metadata "github.com/vijaykiran225/playlist-generator/parser"

	walker "github.com/vijaykiran225/playlist-generator/file-walker"
	"github.com/vijaykiran225/playlist-generator/model"
	builder "github.com/vijaykiran225/playlist-generator/playlist-builder"
)

func ManageApp(rootDir string) {
	listOfFiles := walker.FindMusicFiles(rootDir)
	playlist := model.NewPlaylist()
	for _, path := range listOfFiles {
		m, err := metadata.ExtractMetaData(path)
		if err != nil {
			fmt.Println("skipping", path, err)
			continue
		}
		track := builder.BuildTrack(m, path)
		playlist.AddTrack(track)
	}
	b, err := xml.Marshal(playlist)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	f, err := os.Create("playlist.xspf")
	if err != nil {
		fmt.Println("err while opening file", err)
		return
	}
	_, err = f.WriteString("<?xml version=\"1.0\" encoding=\"UTF-8\"?>")
	if err != nil {
		fmt.Println("error while writing file", err)
		return
	}
	_, err = f.Write(b)
	if err != nil {
		fmt.Println("error while writing file", err)
		return
	}
	fmt.Println("done ")

}
