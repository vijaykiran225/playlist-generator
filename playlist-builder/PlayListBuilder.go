package builder

import (
	"github.com/dhowden/tag"
	"github.com/vijaykiran225/playlist-generator/model"
)

func BuildTrack(m tag.Metadata, path string) model.Track {

	var t model.Track
	t.Title = m.Title()
	t.Album = m.Album()
	t.Location = "file://" + path
	t.Creator = m.Artist()
	return t
}
