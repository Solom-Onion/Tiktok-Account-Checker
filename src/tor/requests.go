package tor

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/net/proxy"
)

func SendOnionPost(urlStr string, torProxy string, payload interface{}) (*http.Response, error) {
	dialer, err := proxy.SOCKS5("tcp", torProxy, nil, proxy.Direct)
	if err != nil {
		fmt.Println("Error creating dialer: ", err)
		return nil, err
	}

	httpTransport := &http.Transport{}
	httpClient := &http.Client{Transport: httpTransport, Timeout: time.Second * 10}

	httpTransport.Dial = dialer.Dial

	proxyUrl, err := url.Parse("socks5://" + torProxy)
	if err != nil {
		fmt.Println("Error parsing proxy URL: ", err)
		return nil, err
	}

	httpTransport.Proxy = http.ProxyURL(proxyUrl)

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error encoding payload: ", err)
		return nil, err
	}

	request, err := http.NewRequest("POST", urlStr, bytes.NewBuffer(payloadBytes))
	if err != nil {
		fmt.Println("Error creating request: ", err)
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")

	response, err := httpClient.Do(request)
	if err != nil {
		fmt.Println("Error sending request: ", err)
		return nil, err
	}

	return response, nil
}
