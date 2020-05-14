package shopware

// MediaTranslation see:
// https://github.com/shopware/platform/blob/6.2/src/Core/Content/Media/Aggregate/MediaTranslation/MediaTranslationEntity.php
type MediaTranslation struct {
	TranslationEntity
	MediaID string       `json:"mediaId"`
	Title   string       `json:"title"`
	Alt     string       `json:"alt"`
	Media   *MediaEntity `json:"media"`
}

// GetMediaID returns the mediaId
func (m *MediaTranslation) GetMediaID() string {
	return m.MediaID
}

// GetTitle returns the title
func (m *MediaTranslation) GetTitle() string {
	return m.Title
}

// GetAlt returns the ALT
func (m *MediaTranslation) GetAlt() string {
	return m.Alt
}

// GetMedia returns the Media
func (m *MediaTranslation) GetMedia() *MediaEntity {
	return m.Media
}

// GetAPIAlias returns the ApiAlias of thumbnail
func (m *MediaTranslation) GetAPIAlias() string {
	return "media_translation"
}
