package utils

type PipelineCfg struct {
	Concourse   concourse   `json:"Concourse"`
	Docker      docker      `json:"Docker"`
	FSLocations FSLocations `json:"FSLocations"`
	MaintAPI    MaintAPI    `json:"MaintAPI"`
	AWS         AWS         `json:"AWS"`
	OBS         OBS         `json:"OBS"`
	PackageName string
	PipeName    string
	RepoName    string
	HoldAfterUP bool
	Version     string
}

type OBS struct {
	User string
	Pass string
}

type AWS struct {
	ID     string `json:"ID"`
	Secret string `json:"SECRET"`
	Region string `json:"REGION"`
}

type FSLocations struct {
	MkcaaspRoot          string `json:"MkcaaspRoot"`
	SkubaRoot            string `json:`
	MkcaaspContainerRoot string `json:"MkcaaspContainerRoot"`
	VilyaRoot            string
}

type concourse struct {
	Username    string
	Password    string
	Team        string
	InstanceUrl string
	PipeName    string
}

type MaintAPI struct {
	QATeam  string
	APIUser string
	APIPwd  string
}

type docker struct {
	RegistryUsername string
	RegistryPassword string
	CaaSP3TestImage  string
	CaaSP4TestImage  string
}

type Updates struct {
	IncidentNumber string
	ReleaseRequest string
	SRCRPMS        []string
	Products       string
	AffectsCaaSP   string
	Repository     string
}