package resources

type EnvConfig struct {
	SshKey       string `json:"ssh_key"`
	ValidatorKey string `json:"validator_key"`
	Passphrase   string `json:"passphrase"`
}
