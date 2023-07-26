package http

import (
	"bytes"
	"context"
	"io"
	"net/http"
)

type Service interface {
	Get(ctx context.Context, urlString string, header map[string]string) ([]byte, error)
	Post(ctx context.Context, urlString string, header map[string]string, data []byte) ([]byte, error)
	GetWithOutParse(ctx context.Context, urlString string, header map[string]string) (io.ReadCloser, error)
}

type httpService struct{}

func NewHTTPService() Service {
	return &httpService{}
}

func (s *httpService) Get(ctx context.Context, urlString string, header map[string]string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, urlString, bytes.NewBuffer([]byte{}))
	if err != nil {
		return nil, err
	}
	for k, v := range header {
		req.Header.Set(k, v)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (s *httpService) Post(
	ctx context.Context,
	urlString string,
	header map[string]string,
	data []byte,
) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, urlString, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	for k, v := range header {
		req.Header.Set(k, v)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (s *httpService) GetWithOutParse(
	ctx context.Context,
	urlString string,
	header map[string]string,
) (io.ReadCloser, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, urlString, bytes.NewBuffer([]byte{}))
	if err != nil {
		return nil, err
	}
	for k, v := range header {
		req.Header.Set(k, v)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}
