package deploy

type NodeConfig struct {
	Endpoint    string
	KeyStoreDir string
	Address     string
	Password    string
}

type EnvConfig struct {
	SSHKey     string
	PrivateKey string
}
