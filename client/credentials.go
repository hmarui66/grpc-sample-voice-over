package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc/credentials"
)

var sslCred credentials.TransportCredentials

type (
	PasswordCompositeCredentials struct {
		sslCred      credentials.TransportCredentials
		passwordCred credentials.PerRPCCredentials
	}
	PasswordCredentials struct {
		password string
	}
)

func newPasswordCompositeCredentials(
	password string,
	caFile string,
	serverHostOverride string,
) (
	credentials.Bundle,
	error,
) {
	if sslCred == nil {
		var err error
		sslCred, err = credentials.NewClientTLSFromFile(caFile, serverHostOverride)
		if err != nil {
			return nil, fmt.Errorf("failed to new ssl credentials: %w", err)
		}
	}
	return &PasswordCompositeCredentials{
		sslCred: sslCred,
		passwordCred: &PasswordCredentials{
			password: password,
		},
	}, nil
}

func (b *PasswordCompositeCredentials) TransportCredentials() credentials.TransportCredentials {
	return b.sslCred
}

func (b *PasswordCompositeCredentials) PerRPCCredentials() credentials.PerRPCCredentials {
	return b.passwordCred
}

func (*PasswordCompositeCredentials) NewWithMode(mode string) (credentials.Bundle, error) {
	panic("implement me")
}

func (p *PasswordCredentials) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"password": p.password,
	}, nil
}

func (*PasswordCredentials) RequireTransportSecurity() bool {
	return true
}
