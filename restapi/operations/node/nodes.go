package node

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/nyasuto/goswagger/models"
)

func toPtr(s string) *string {
	return &s
}
func Search() middleware.Responder {
	payload := &models.Node{
		Name:   toPtr("nyan"),
		Status: "hungry",
	}

	return NewGetNodesOK().WithPayload(payload)
}
