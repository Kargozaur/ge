package responses

type Token struct {
	AccessToken string `json:"access_token"`
	TokenType	string `json:"token_type"`
}

func NewToken(tokenStr string) Token {
	return Token{AccessToken: tokenStr, TokenType: "Bearer"}
}