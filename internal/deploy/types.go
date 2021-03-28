package deploy

type NodeConfig struct {
	Endpoint    string
	KeyStoreDir string
	Address     string
	Password    string
	SshKey      string
	IP          string
}

type EnvConfig struct {
	SSHKey       string
	ValidatorKey string
}
