// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "computingProvider service APIs": Storage Resource Client
//
// Command:
// $ goagen
// --design=github.com/ZJU-DistributedAI/ComputingProvider/design
// --out=$(GOPATH)/src/github.com/ZJU-DistributedAI/ComputingProvider
// --version=v1.3.1

package client

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

// AddStoragePath computes a request path to the add action of Storage.
func AddStoragePath() string {

	return fmt.Sprintf("/api/v0/storage")
}

// Upload file to IPFS using multipart post
func (c *Client) AddStorage(ctx context.Context, path string, payload *FilePayload, contentType string) (*http.Response, error) {
	req, err := c.NewAddStorageRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewAddStorageRequest create the request corresponding to the add action endpoint of the Storage resource.
func (c *Client) NewAddStorageRequest(ctx context.Context, path string, payload *FilePayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	w := multipart.NewWriter(&body)

	{
		_, file := filepath.Split(payload.File)
		fw, err := w.CreateFormFile("file", file)
		if err != nil {
			return nil, err
		}
		fh, err := os.Open(payload.File)
		if err != nil {
			return nil, err
		}
		defer fh.Close()
		if _, err := io.Copy(fw, fh); err != nil {
			return nil, err
		}
	}
	if err := w.Close(); err != nil {
		return nil, err
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", w.FormDataContentType())
	return req, nil
}

// CatStoragePath computes a request path to the cat action of Storage.
func CatStoragePath(address string) string {
	param0 := address

	return fmt.Sprintf("/api/v0/storage/%s", param0)
}

// Cat the file in IPFS at :address
func (c *Client) CatStorage(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewCatStorageRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCatStorageRequest create the request corresponding to the cat action endpoint of the Storage resource.
func (c *Client) NewCatStorageRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
