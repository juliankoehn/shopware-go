package shopware

// MediaFolderConfiguration see:
// https://github.com/shopware/platform/blob/6.2/src/Core/Content/Media/Aggregate/MediaFolderConfiguration/MediaFolderConfigurationEntity.php
type MediaFolderConfiguration struct {
	Entity
	MediaFolders          []*MediaFolder        `json:"mediaFolders"`
	CreateThumbnails      bool                  `json:"createThumbnails"`
	KeepAspectRatio       bool                  `json:"keepAspectRatio"`
	ThumbnailQuality      int                   `json:"thumbnailQuality"`
	Private               bool                  `json:"private"`
	NoAssociation         bool                  `json:"noAssociation"`
	MediaThumbnailSizes   []*MediaThumbnailSize `json:"mediaThumbnailSizes"`
	MediaThumbnailSizesRo string                `json:"mediaThumbnailSizesRo"`
}
