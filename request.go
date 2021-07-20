package compound

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/fox-one/pkg/uuid"
	"github.com/go-resty/resty/v2"
)

const (
	// HeaderKeyRequestID request id header key
	headerKeyRequestID = "X-Request-Id"
)

type ContextKey int

const (
	ContextKeyRequestID ContextKey = iota
)

var restyClient *resty.Client = resty.New().
	SetHeader("Content-Type", "application/json").
	SetHeader("Charset", "utf-8").
	SetTimeout(10 * time.Second).
	SetPreRequestHook(func(c *resty.Client, r *http.Request) error {
		ctx := r.Context()
		if values := r.Header.Values(headerKeyRequestID); len(values) == 0 {
			r.Header.Set(headerKeyRequestID, requestIDFromContext(ctx))
		}

		return nil
	}).OnAfterResponse(func(c *resty.Client, r *resty.Response) error {
	if err := checkResponseRequestID(r); err != nil {
		return err
	}

	return nil
})

// Request new resty request
func request(ctx context.Context) *resty.Request {
	return restyClient.R().SetContext(ctx)
}

// WithRequestID context with request id
func withRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, ContextKeyRequestID, requestID)
}

func requestIDFromContext(ctx context.Context) string {
	if v, ok := ctx.Value(ContextKeyRequestID).(string); ok {
		return v
	}

	return uuid.New()
}

func checkResponseRequestID(r *resty.Response) error {
	expect := r.Request.Header.Get(headerKeyRequestID)
	got := r.Header().Get(headerKeyRequestID)
	if expect != "" && got != "" && expect != got {
		return fmt.Errorf("%s mismatch, expect %q but got %q", headerKeyRequestID, expect, got)
	}

	return nil
}

func parseResponse(r *resty.Response, obj interface{}) error {
	if !r.IsSuccess() {
		return fmt.Errorf(string(r.Body()))
	}

	if obj != nil {
		if e := json.Unmarshal(r.Body(), obj); e != nil {
			return e
		}
	}

	return nil
}
