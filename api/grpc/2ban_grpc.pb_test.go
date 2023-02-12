package grpc

import (
	"context"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"testing"
)

func TestIP2BanClient(t *testing.T) {
	assert := assert.New(t)

	// Create an instance of context.Context
	ctx := context.Background()

	// Create an instance of IPStringRequest
	in := &IPStringRequest{
		Ip: "127.0.0.1",
	}

	// Create a mock implementation of the IP2BanClient interface
	mockClient := &mockIP2BanClient{}

	// Call the IP method of the mock client
	out, err := mockClient.IP(ctx, in)

	// Check for errors
	assert.Nil(err)
	assert.NotNil(out)

	// Compare the returned value with the expected value
	expectedMessage := &OKReply{
		Ok: true,
	}
	assert.Equal(expectedMessage, out)
}

type mockIP2BanClient struct{}

func (c *mockIP2BanClient) IP(ctx context.Context, in *IPStringRequest, opts ...grpc.CallOption) (*OKReply, error) {
	return &OKReply{
		Ok: true,
	}, nil
}
