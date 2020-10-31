package context

import (
	stdcontext "context"
)

const (
	valueRequestID = "request_id"
	valueTenantID  = "tenant_id"
	valueUserID    = "user_id"
)

// New creates context with RequestID, TenantID, UserID set
func New(rid, tid, uid string) stdcontext.Context {
	ctx := stdcontext.WithValue(stdcontext.Background(), valueRequestID, rid)
	ctx = stdcontext.WithValue(ctx, valueTenantID, tid)
	ctx = stdcontext.WithValue(ctx, valueUserID, uid)

	return ctx
}

// GetRequestID returns RequestID of the Context
func GetRequestID(c stdcontext.Context) string {
	return c.Value(valueRequestID).(string)
}

// GetTenantID returns TenantID of the Context
func GetTenantID(c stdcontext.Context) string {
	return c.Value(valueTenantID).(string)
}

// GetUserID returns UserID of the Context
func GetUserID(c stdcontext.Context) string {
	return c.Value(valueUserID).(string)
}
