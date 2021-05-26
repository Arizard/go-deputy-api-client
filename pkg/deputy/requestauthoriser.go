package deputy

import (
	"fmt"
	"net/http"
)

// RequestAuthoriser takes a *http.Request and makes it authorised.
// For example, this could add the Authorization header.
type RequestAuthoriser func(r *http.Request)

func NewBearerTokenRequestAuthoriser(token string) RequestAuthoriser {
	return func(req *http.Request) {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	}
}
