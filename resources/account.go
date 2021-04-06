package resources

type AccountRequest struct {
	Email    string `json:"email"`
	Password []byte `json:"password"`
}

type AccountResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
}
