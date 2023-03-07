package plugins

import (
	"encoding/json"
	"net/http"

	pkgHTTP "github.com/apache/apisix-go-plugin-runner/pkg/http"
	"github.com/golang-jwt/jwt/v5"
)

// Say is a demo to show how to return data directly instead of proxying
// it to the upstream.
type JWT2HeaderPlugin struct {
}

type JWT2HeaderConfig struct {
	TokenRequired bool `json:"token"`
}

func (b *JWT2HeaderPlugin) Name() string {
	return "jwt2header"
}

func (b *JWT2HeaderPlugin) ParseConf(in []byte) (interface{}, error) {
	conf := JWT2HeaderConfig{}
	err := json.Unmarshal(in, &conf)
	return conf, err
}

func (b *JWT2HeaderPlugin) RequestFilter(conf interface{}, w http.ResponseWriter, r pkgHTTP.Request) {
	config := conf.(JWT2HeaderConfig)
	auth := r.Header().Get("Authorization")
	if auth == "" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`No JWT Token found`))
		return
	}
	token, _ := jwt.Parse(string(auth), nil)

	if token.Valid && config.TokenRequired {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`Invalid JWT Token`))
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	for k, v := range claims {
		switch val := v.(type) {
		case string:
			r.Header().Set("APISIX-jwt-claim-"+k, val)
		default:
		}
	}

}

func (b *JWT2HeaderPlugin) ResponseFilter(conf interface{}, r pkgHTTP.Response) {
}
