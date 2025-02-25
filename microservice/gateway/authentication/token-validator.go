package authentication

import (
	"encoding/json"

	"aidanwoods.dev/go-paseto"
)

type PasetoChecker struct {
	parser    paseto.Parser
	publicKey paseto.V4AsymmetricPublicKey
}

const TokenPublicKey = "aAvU8SHCtqdwoD3rL5S6s0hF35BnoH4+2jGbtrmuLiM="

var payload Payload

var checker *PasetoChecker

func (p *AuthPlugin) VerifyToken(token string) (*Payload, error) {
	key, _ := paseto.NewV4AsymmetricPublicKeyFromHex(TokenPublicKey)
	checker = &PasetoChecker{
		parser:    paseto.NewParser(),
		publicKey: key,
	}

	parsedToken, err := checker.parser.ParseV4Public(key, token, nil)
	if err != nil {
		if err.Error() == "this token has expired" {
			return nil, err
		}
		return nil, err
	}

	parsed, err := parsedToken.GetString("payload")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(parsed), &payload)
	if err != nil {
		return nil, err
	}

	return &payload, nil
}
