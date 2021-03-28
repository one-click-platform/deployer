package resources

type EnvConfig struct {
	Status       string `json:"status"`
	SshKey       string `json:"ssh_key"`
	ValidatorKey string `json:"validator_key"`
	Passphrase   string `json:"passphrase"`
	IP           string `json:"ip"`
}
