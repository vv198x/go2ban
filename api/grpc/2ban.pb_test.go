package grpc

import (
	"google.golang.org/protobuf/proto"
	"testing"
)

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
