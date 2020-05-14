package shopware

// MediaFolder see:
// https://github.com/shopware/platform/blob/6.2/src/Core/Content/Media/Aggregate/MediaFolder/MediaFolderEntity.php
type MediaFolder struct {
	Entity
	Name                   string                    `json:"name"`
	ParentID               string                    `json:"parentId"`
	Parent                 *MediaFolder              `json:"parent"`
	ChildCount             int                       `json:"childCount"`
	Media                  []*MediaEntity            `json:"media"`
	ConfigurationID        string                    `json:"configurationId"`
	Configuration          *MediaFolderConfiguration `json:"configuration"`
	UseParentConfiguration bool                      `json:"useParentConfiguration"`
	Children               []*MediaFolder            `json:"children"`
	DefaultFolder          *MediaFolder              `json:"defaultFolder"`
	DefaultFolderID        string                    `json:"defaultFolderId"`
}

// GetName returns the name of the media fodler
func (m *MediaFolder) GetName() string {
	return m.Name
}

// GetParentID returns the ParentID
func (m *MediaFolder) GetParentID() string {
	return m.Name
}

// GetParent returns the Parent
func (m *MediaFolder) GetParent() *MediaFolder {
	return m.Parent
}

// GetChildCount returns the ChildCount
func (m *MediaFolder) GetChildCount() int {
	return m.ChildCount
}

// GetMedia returns the Media
func (m *MediaFolder) GetMedia() []*MediaEntity {
	return m.Media
}

// GetConfigurationID returns the ConfigurationID
func (m *MediaFolder) GetConfigurationID() string {
	return m.ConfigurationID
}

// GetConfiguration returns the Configuration
func (m *MediaFolder) GetConfiguration() *MediaFolderConfiguration {
	return m.Configuration
}

// GetUseParentConfiguration returns the UseParentConfiguration
func (m *MediaFolder) GetUseParentConfiguration() bool {
	return m.UseParentConfiguration
}

// GetChildren returns the Children
func (m *MediaFolder) GetChildren() []*MediaFolder {
	return m.Children
}

// GetDefaultFolder returns the DefaultFolder
func (m *MediaFolder) GetDefaultFolder() *MediaFolder {
	return m.DefaultFolder
}

// GetDefaultFolderID returns the DefaultFolderID
func (m *MediaFolder) GetDefaultFolderID() string {
	return m.DefaultFolderID
}

// GetAPIAlias returns the ApiAlias
func (m *MediaFolder) GetAPIAlias() string {
	return "media_folder"
}
