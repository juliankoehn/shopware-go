package shopware

import "fmt"

// MediaThumbnail see:
// https://github.com/shopware/platform/blob/6.2/src/Core/Content/Media/Aggregate/MediaThumbnail/MediaThumbnailEntity.php
type MediaThumbnail struct {
	Entity
	Width   int          `json:"width"`
	Height  int          `json:"height"`
	URL     string       `json:"url"`
	MediaID string       `json:"mediaId"`
	Media   *MediaEntity `json:"media"`
}

// GetWidth returns the width of thumbnail
func (m *MediaThumbnail) GetWidth() int {
	return m.Width
}

// GetHeight returns the height of thumbnail
func (m *MediaThumbnail) GetHeight() int {
	return m.Height
}

// GetURL returns the URL of thumbnail
func (m *MediaThumbnail) GetURL() string {
	return m.URL
}

// GetMediaID returns the MediaID of thumbnail
func (m *MediaThumbnail) GetMediaID() string {
	return m.MediaID
}

// GetMedia returns the Media of thumbnail
func (m *MediaThumbnail) GetMedia() *MediaEntity {
	return m.Media
}

// GetIdentifier returns the Identifier of thumbnail
func (m *MediaThumbnail) GetIdentifier() string {
	identifier := fmt.Sprintf("%dx%d", m.Width, m.Height)
	return identifier
}

// GetAPIAlias returns the ApiAlias of thumbnail
func (m *MediaThumbnail) GetAPIAlias() string {
	return "media_thumbnail"
}
