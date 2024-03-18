package farewell

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func GreetInSpanish(ctx context.Context, name string) (string, error) {
	greeting, err := callService("get-spanish-greeting", name)
	return greeting, err
}

func FarewellInSpanish(ctx context.Context, name string) (string, error) {
	greeting, err := callService("get-spanish-farewell", name)
	return greeting, err
}

// utility function for making calls to the microservices
func callService(stem string, name string) (string, error) {
	base := "http://localhost:9999/" + stem + "?name=%s"
	url := fmt.Sprintf(base, url.QueryEscape(name))

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	translation := string(body)

	status := resp.StatusCode
	if status >= 400 {
		message := fmt.Sprintf("HTTP Error %d: %s", status, translation)
		return "", errors.New(message)
	}

	return translation, nil
}
