package clients

import (
	"context"
	"github.com/pkg/errors"
	"log"
	"net/http"
)

type SlackClient interface {
	SendMessageForUser(token string, ctx context.Context) (*ApiResult, error)
}

type slackClient struct {
	httpClient *http.Client
}

func (sc *slackClient) SendMessageForUser(token string, ctx context.Context) (*ApiResult, error) {

	r, err := http.NewRequest(http.MethodPost, "", nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	r = r.WithContext(ctx)
	defer func() {
		if err := r.Body.Close(); err != nil {
			log.Println("ERROR - Closing request body")
		}
	}()

	res, err := sc.httpClient.Do(r)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer func() {
		if err := res.Body.Close; err  != nil {
			log.Println("ERROR - Closing response body")
		}
	}()

	return &ApiResult{}, nil
}

func NewSlackClient(client *http.Client) SlackClient {
	return &slackClient{}
}