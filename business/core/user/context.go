package user

import "context"

// ctxKey is a type for context keys
type ctxKey string

// ctxTokenKey is the context key for the token
const ctxTokenKey ctxKey = "userToken"

// ContextWithToken adds the token to the context
func ContextWithToken(ctx context.Context, token Token) context.Context {
	return context.WithValue(ctx, ctxTokenKey, token)
}

// GetContextToken returns the token from the context
func GetContextToken(ctx context.Context) Token {
	token, ok := ctx.Value(ctxTokenKey).(Token)
	if !ok {
		return Token{}
	}
	return token
}
