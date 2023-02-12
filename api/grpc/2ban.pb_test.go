package grpc

import (
	"github.com/golang/protobuf/proto"
	"testing"
)

func TestIPStringRequest(t *testing.T) {
	req := &IPStringRequest{
		Ip: "192.168.0.1",
	}

	// Ensure the IPStringRequest message is properly initialized
	if req == nil {
		t.Errorf("Expected IPStringRequest message to be initialized, got nil")
	}

	// Ensure the IPStringRequest message has the expected IP address
	if req.GetIp() != "192.168.0.1" {
		t.Errorf("Expected IPStringRequest message to have IP address %s, got %s", "192.168.0.1", req.GetIp())
	}

	// Ensure the IPStringRequest message can be reset correctly
	req.Reset()
	if req.GetIp() != "" {
		t.Errorf("Expected IPStringRequest message to have IP address %s after reset, got %s", "", req.GetIp())
	}
}

func TestOKReply(t *testing.T) {
	// Test Reset function
	okReply := &OKReply{
		Ok: true,
	}
	okReply.Reset()
	if okReply.Ok != false {
		t.Errorf("Reset function for OKReply not working as expected, expected: false, got: %v", okReply.Ok)
	}

	// Test String function
	okReply.Ok = true
	expectedString := "ok:true"
	if okReply.String() != expectedString {
		t.Errorf("String function for OKReply not working as expected, expected: %s, got: %s", expectedString, okReply.String())
	}

	// Test ProtoReflect function
	okReply.Ok = true
	expectedMessage := &OKReply{
		Ok: true,
	}
	if !proto.Equal(okReply.ProtoReflect().Interface().(*OKReply), expectedMessage) {
		t.Errorf("ProtoReflect function for OKReply not working as expected, expected: %v, got: %v", expectedMessage, okReply.ProtoReflect().Interface().(*OKReply))
	}

	// Test GetOk function
	okReply.Ok = true
	if okReply.GetOk() != true {
		t.Errorf("GetOk function for OKReply not working as expected, expected: true, got: %v", okReply.GetOk())
	}
	okReply = nil
	if okReply.GetOk() != false {
		t.Errorf("GetOk function for OKReply not working as expected, expected: false, got: %v", okReply.GetOk())
	}
}
