package deploy

type NodeConfig struct {
	Endpoint    string
	KeyStoreDir string
	Address     string
	Password    string
	SshKey      string
}

type EnvConfig struct {
	SSHKey       string
	ValidatorKey string
}
