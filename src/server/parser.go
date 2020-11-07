package server

import "net/http"


func GetHeadersFromRequest(request *http.Request) map[string]string {
	headersMap := make(map[string]string)

	for name, headers := range request.Header {
		for _, h := range headers {
			headersMap[name] = h
		}
	}

	return headersMap
}
