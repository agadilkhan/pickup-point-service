package transport

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/agadilkhan/pickup-point-service/internal/auth/config"
	"io"
	"net/http"
	"time"
)

type UserTransport struct {
	cfg config.UserTransport
}

func NewUserTransport(cfg config.UserTransport) *UserTransport {
	return &UserTransport{
		cfg: cfg,
	}
}

type GetUserResponse struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Login     string `json:"login"`
	Password  string `json:"password"`
}

func (ut *UserTransport) GetUser(ctx context.Context, login string) (*GetUserResponse, error) {
	var response *GetUserResponse

	responseBody, err := ut.makeRequest(
		ctx,
		"GET",
		fmt.Sprintf("api/user/v1/user/%s", login),
		ut.cfg.ShutdownTimeout,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to makeRequest err: %w", err)
	}

	if err := json.Unmarshal(responseBody, response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response err: %w", err)
	}

	return response, nil
}

func (ut *UserTransport) makeRequest(
	ctx context.Context,
	httpMethod string,
	endPoint string,
	timeout time.Duration,
) (b []byte, err error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	requestUrl := ut.cfg.Host + endPoint

	req, err := http.NewRequestWithContext(ctx, httpMethod, requestUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to newRequestWithContext err: %w", err)
	}

	httpClient := &http.Client{}

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("client making http request:: %w", err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return body, nil
}
