package helpers

import "net/http"

func SetAuthHeader(req *http.Request, accessToken string) {
	req.Header.Set("Authorization", "Bearer "+accessToken)
}
