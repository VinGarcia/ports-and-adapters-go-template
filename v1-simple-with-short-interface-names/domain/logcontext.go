package domain

// This file is part of the domain language, and it allows different packages
// to save values to the context in a way that our log provider implementation
// is able to read and log later on.
//
// In the domain package modules that have function should be kept to a minimum,
// and should only be put here if they really need to be used by several different
// packages.

import (
	"context"

	"github.com/vingarcia/ddd-go-template/v1-simple-with-short-interface-names/infra/maps"
)

// Declaring a unique private type for the ctx key
// guarantees that no key colision will ever happen:
type logCtxKeyType uint8

var logCtxKey logCtxKeyType

type ctxBody = map[string]interface{}

// CtxWithValues merges received values with log body currently stored
// on the input ctx.
func CtxWithValues(ctx context.Context, values ctxBody) context.Context {
	m, _ := ctx.Value(logCtxKey).(ctxBody)
	return context.WithValue(ctx, logCtxKey, mergeMaps(m, values))
}

// GetCtxValues extracts the ctxBody currently stored on the input ctx.
func GetCtxValues(ctx context.Context) ctxBody {
	m, _ := ctx.Value(logCtxKey).(ctxBody)
	if m == nil {
		m = ctxBody{}
	}
	m["request_id"] = GetRequestIDFromContext(ctx)
	return m
}

func mergeMaps(bodies ...ctxBody) ctxBody {
	body := ctxBody{}
	maps.Merge(&body, bodies...)
	return body
}