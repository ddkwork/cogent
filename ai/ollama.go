package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"cogentcore.org/core/base/errors"
	"cogentcore.org/core/base/fsx"
	"github.com/aandrew-me/tgpt/v2/structs"
	http "github.com/bogdanfinn/fhttp"
	tls_client "github.com/bogdanfinn/tls-client"
	"github.com/bogdanfinn/tls-client/profiles"
)

type Response struct {
	ID      string `json:"id"`
	Choices []struct {
		Delta struct {
			Content string `json:"content"`
		} `json:"delta"`
	} `json:"choices"`
}

func NewRequest(input string, params structs.Params, prevMessages string) (r *http.Response, err error) {
	client, err := NewClient()
	if errors.Log(err) != nil {
		return
	}

	model := "mistral"
	if params.ApiModel != "" {
		model = params.ApiModel
	}

	temperature := "0.5"
	if params.Temperature != "" {
		temperature = params.Temperature
	}

	topP := "0.5"
	if params.Top_p != "" {
		topP = params.Top_p
	}

	safeInput, err := json.Marshal(input)
	if errors.Log(err) != nil {
		return
	}

	data := strings.NewReader(fmt.Sprintf(`{
		"frequency_penalty": 0,
		"messages": [
			%v
			{
				"content": %v,
				"role": "user"
			}
		],
		"model": "%v",
		"presence_penalty": 0,
		"stream": true,
		"temperature": %v,
		"top_p": %v
	}
	`, prevMessages, string(safeInput), model, temperature, topP))

	req, err := http.NewRequest("POST", "http://localhost:11434/v1/chat/completions", data)
	if errors.Log(err) != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+params.ApiKey)

	return client.Do(req)
}

func HandleToken(respBody string) (string, error) {
	//https://521github.com/ollama/ollama/blob/main/openai/op#262
	//_, err = w.ResponseWriter.Write([]byte(fmt.Sprintf("data: %s\n\n", d)))
	//_, err = w.ResponseWriter.Write([]byte("data: [DONE]\n\n"))
	if strings.Contains(respBody, "data: [DONE]") {
		return "", nil
	}

	obj := "{}"
	if len(respBody) > 1 {
		split := strings.Split(respBody, "data: ")
		if len(split) > 1 {
			obj = split[1]
		} else {
			obj = split[0]
		}
	}

	var d Response
	if err := json.Unmarshal([]byte(obj), &d); err != nil {
		return "", err
	}

	if d.Choices != nil {
		token := d.Choices[0].Delta.Content
		return token, nil
	}
	return "", nil
}

func NewClient() (tls_client.HttpClient, error) {
	jar := tls_client.NewCookieJar()
	options := []tls_client.HttpClientOption{
		tls_client.WithTimeoutSeconds(1200),
		tls_client.WithClientProfile(profiles.Firefox_110),
		tls_client.WithNotFollowRedirects(),
		tls_client.WithCookieJar(jar),
		// tls_client.WithInsecureSkipVerify(),
	}

	//todo clean logic
	proxyAddress := os.Getenv("HTTP_PROXY")
	if proxyAddress == "" {
		proxyAddress = os.Getenv("http_proxy")
	}
	if proxyAddress != "" {
		if strings.HasPrefix(proxyAddress, "http://") || strings.HasPrefix(proxyAddress, "socks5://") {
			proxyOption := tls_client.WithProxyUrl(proxyAddress)
			options = append(options, proxyOption)
		}
	} else {
		if ok := errors.Log1(fsx.FileExists("proxy.txt")); ok {
			proxyConfig, err := os.ReadFile("proxy.txt")
			if err != nil {
				return nil, err
			}
			proxyAddress := strings.TrimSpace(string(proxyConfig))
			if proxyAddress != "" {
				if strings.HasPrefix(proxyAddress, "http://") || strings.HasPrefix(proxyAddress, "socks5://") {
					proxyOption := tls_client.WithProxyUrl(proxyAddress)
					options = append(options, proxyOption)
				}
			}
		}
	}
	return tls_client.NewHttpClient(tls_client.NewNoopLogger(), options...)
}
