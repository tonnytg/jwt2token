package jwt2token

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
)

// GoogleAccessTokenFromJSON get json file and convert to token, but you need start for CleanToken
func GoogleAccessTokenFromJSON(secretJson string, scope string) (string, error) {
	data, err := ioutil.ReadFile(secretJson)
	if err != nil {
		return "", err
	}

	conf, err := google.JWTConfigFromJSON(data, scope)
	if err != nil {
		return "", err
	}

	res := conf.TokenSource(oauth2.NoContext)
	token, err := res.Token()
	if err != nil {
		return "", err
	}

	return token.AccessToken, err
}

// CleanToken start here to get token and clean many dots .....
func CleanToken(fileJson string) (string, error) {

	token, _ := GoogleAccessTokenFromJSON(fileJson, "https://www.googleapis.com/auth/cloud-platform")
	count := 0
	newToken := ""
	for i := 0; i < len(token); i++ {
		if count < 173 {
			newToken = newToken + string(token[i])
			count++
		} else {
			if string(token[i]) != "." {
				newToken = newToken + string(token[i])
				count++
			}
		}
	}
	return newToken, nil
}

func JWT2token() {
	// create a JSON key of service key and save with name credential.json
	CleanToken("credential.json")
}
