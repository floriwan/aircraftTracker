package http

import (
	"fmt"
	"io"
	"net/http"
)

func SendRequest(req *http.Request) (b []byte, err error) {

	//log.Printf("> %v : %s\n  %v", req.Method, req.URL, req.Header)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("got response status code %v %v", res.StatusCode, http.StatusText(res.StatusCode))
	}

	//log.Printf("response status %v %v\n", res.StatusCode, http.StatusText(res.StatusCode))
	defer res.Body.Close()
	b, _ = io.ReadAll(res.Body)
	return b, nil
}
