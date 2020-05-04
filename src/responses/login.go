package responses

//Login godoc
//@Summary defines the login response model
type Login struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	Expires      int    `json:"expires"`
}
