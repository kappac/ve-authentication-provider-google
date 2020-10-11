package handler

import (
	"context"

	"github.com/kappac/restx-helpers/pkg/proto/api"
)

// Handler serves gRPC requests
type Handler struct{}

// ValidateToken handles ValidateToken RPC call
func (h *Handler) ValidateToken(ctx context.Context, req *api.RestXRequest, rsp *api.RestXResponse) error {
	return nil
}
