// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "wit": login Resource Client
//
// Command:
// $ goagen
// --design=github.com/fabric8-services/fabric8-wit/design
// --out=$(GOPATH)/src/github.com/fabric8-services/fabric8-wit
// --version=v1.3.0

package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// AuthorizeLoginPath computes a request path to the authorize action of login.
func AuthorizeLoginPath() string {

	return fmt.Sprintf("/api/login/authorize")
}

// Authorize with the WIT
func (c *Client) AuthorizeLogin(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewAuthorizeLoginRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewAuthorizeLoginRequest create the request corresponding to the authorize action endpoint of the login resource.
func (c *Client) NewAuthorizeLoginRequest(ctx context.Context, path string) (*http.Request, error) {
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

// GenerateLoginPath computes a request path to the generate action of login.
func GenerateLoginPath() string {

	return fmt.Sprintf("/api/login/generate")
}

// Generate a set of Tokens for different Auth levels. NOT FOR PRODUCTION. Only available if server is running in dev mode
func (c *Client) GenerateLogin(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewGenerateLoginRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewGenerateLoginRequest create the request corresponding to the generate action endpoint of the login resource.
func (c *Client) NewGenerateLoginRequest(ctx context.Context, path string) (*http.Request, error) {
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

// RefreshLoginPath computes a request path to the refresh action of login.
func RefreshLoginPath() string {

	return fmt.Sprintf("/api/login/refresh")
}

// Refresh access token
func (c *Client) RefreshLogin(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewRefreshLoginRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewRefreshLoginRequest create the request corresponding to the refresh action endpoint of the login resource.
func (c *Client) NewRefreshLoginRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
