package http

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrCannotCreateRequest = errors.New("cannot create request")
	ErrCannotDoRequest     = errors.New("cannot do request")
)

func Get(url string, ctx context.Context) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrCannotCreateRequest, err)
	}

	req.Header.Add("Authorization", "Bearer "+ctx.Value("token").(string))

	res, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrCannotDoRequest, err)
	}

	return res, nil
}
