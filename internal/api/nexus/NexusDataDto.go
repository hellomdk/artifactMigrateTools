package nexus

type NexusDataDto struct {
	Name   string `json:"name"`
	Result string `json:"result"`
}

type NexusData struct {
	Users     []NexusUsers     `json:"users"`
	Groups    []NexusGroups    `json:"groups"`
	Privs     []NexusPrivs     `json:"privs"`
	Selectors []NexusSelectors `json:"selectors"`
	Ldaps     []NexusLdaps     `json:"ldaps"`
	Repos     []NexusRepos     `json:"repos"`
}

type NexusUsers struct {
}
type NexusGroups struct {
}
type NexusPrivs struct {
}
type NexusSelectors struct {
}
type NexusLdaps struct {
}
type NexusRepos struct {
	Type   string      `json:"type"`
	Format string      `json:"format"`
	Name   string      `json:"name"`
	Url    string      `json:"url"`
	Config NexusConfig `json:"config"`
}

type NexusConfig struct {
	Name       string          `json:"name"`
	Recipe     string          `json:"recipe"`
	Online     bool            `json:"online"`
	Attributes NexusAttributes `json:"attributes"`
}

type NexusAttributes struct {
	Component  NexusComponent  `json:"maven"`
	Maven      NexusMaven      `json:"maven"`
	Proxy      NexusProxy      `json:"proxy"`
	Httpclient NexusHttpclient `json:"httpclient"`
	Storage    NexusStorage    `json:"storage"`
	Group      NexusGroup      `json:"group"`
}

type NexusGroup struct {
	MemberNames []string `json:"memberNames"`
}

type NexusComponent struct {
	ProprietaryComponents bool `json:"proprietaryComponents"`
}

type NexusMaven struct {
	VersionPolicy string `json:"versionPolicy"`
	LayoutPolicy  string `json:"layoutPolicy"`
}

type NexusProxy struct {
	RemoteUrl      string `json:"remoteUrl"`
	ContentMaxAge  int    `json:"contentMaxAge"`
	MetadataMaxAge string `json:"metadataMaxAge"`
}

type NexusHttpclient struct {
	Blocked        bool                `json:"blocked"`
	AutoBlock      bool                `json:"autoBlock"`
	Authentication NexusAuthentication `json:"authentication"`
}

type NexusAuthentication struct {
	Type     string `json:"type"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type NexusStorage struct {
	BlobStoreName               string `json:"blobStoreName"`
	StrictContentTypeValidation bool   `json:"strictContentTypeValidation"`
}
