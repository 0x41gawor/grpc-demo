package main

import (
	"context"
	"fmt"

	"github.com/0x41gawor/grpc-demo/invoicer"
)

type MyInvoicerServer struct {
	invoicer.UnimplementedInvoicerServer
}

func (s MyInvoicerServer) Create(c context.Context, r *invoicer.CreateRequest) (*invoicer.CreateResponse, error) {
	return &invoicer.CreateResponse{
		Pdf:  []byte(fmt.Sprintf("test %s", r.Amount)),
		Docx: []byte(fmt.Sprintf("test %s", r.From)),
	}, nil
}
