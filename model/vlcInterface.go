package model

import (
	"encoding/xml"
)

type Playlist struct {
	XMLName   xml.Name  `xml:"playlist"`
	Text      string    `xml:",chardata"`
	Xmlns     string    `xml:"xmlns,attr"`
	Vlc       string    `xml:"vlc,attr"`
	Version   string    `xml:"version,attr"`
	Title     string    `xml:"title"`
	TrackList TrackList `xml:"trackList"`
	Extension struct {
		Text        string `xml:",chardata"`
		Application string `xml:"application,attr"`
		Item        []struct {
			Text string `xml:",chardata"`
			Tid  string `xml:"tid,attr"`
		} `xml:"item"`
	} `xml:"extension"`
}

type TrackList struct {
	Text  string  `xml:",chardata"`
	Track []Track `xml:"track"`
}
type Track struct {
	Text       string `xml:",chardata"`
	Location   string `xml:"location"`
	Title      string `xml:"title"`
	Creator    string `xml:"creator"`
	Album      string `xml:"album"`
	Annotation string `xml:"annotation"`
	Duration   string `xml:"duration"`
	Extension  struct {
		Text        string `xml:",chardata"`
		Application string `xml:"application,attr"`
		ID          string `xml:"id"`
	} `xml:"extension"`
}

func NewPlaylist() Playlist {

	var pl Playlist
	pl.Xmlns = "http://xspf.org/ns/0/"
	pl.Title = "Playlist"
	pl.Version = "1"
	pl.Vlc = "http://www.videolan.org/vlc/playlist/ns/0/"
	pl.TrackList = TrackList{
		Track: []Track{},
	}
	return pl
}

func (pl *Playlist) AddTrack(aTrack Track) {
	pl.TrackList.Track = append(pl.TrackList.Track, aTrack)
}
