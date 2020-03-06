package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc/credentials"
)

var sslCred credentials.TransportCredentials

type (
	creds struct {
		transportCreds credentials.TransportCredentials
		perRPCCreds    credentials.PerRPCCredentials
	}
	tokenCreds struct {
		token string
	}
)

func newCredentials(
	token string,
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
	return &creds{
		transportCreds: sslCred,
		perRPCCreds: &tokenCreds{
			token: token,
		},
	}, nil
}

func (c *creds) TransportCredentials() credentials.TransportCredentials {
	return c.transportCreds
}

func (c *creds) PerRPCCredentials() credentials.PerRPCCredentials {
	return c.perRPCCreds
}

func (c *creds) NewWithMode(mode string) (credentials.Bundle, error) {
	return c, nil
}

func (t *tokenCreds) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": "bearer " + t.token,
	}, nil
}

func (*tokenCreds) RequireTransportSecurity() bool {
	return true
}
