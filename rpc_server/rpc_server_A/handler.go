package main

import (
	"context"
	demo "github.com/211250125liu/rpc_server/kitex_gen/demo"
)

// Service_1Impl implements the last service interface defined in the IDL.
type Service_1Impl struct{}

// GetMessage implements the Service_1Impl interface.
func (s *Service_1Impl) GetMessage(ctx context.Context, req *demo.Request) (resp *demo.Response, err error) {
	// TODO: Your code here...

	resp = &demo.Response{Message: req.Message + "get it"}
	return
}
