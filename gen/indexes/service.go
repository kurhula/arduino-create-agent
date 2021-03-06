// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// indexes service
//
// Command:
// $ goa gen github.com/arduino/arduino-create-agent/design

package indexes

import (
	"context"

	indexesviews "github.com/arduino/arduino-create-agent/gen/indexes/views"
	"goa.design/goa"
)

// The indexes service manages the package_index files
type Service interface {
	// List implements list.
	List(context.Context) (res []string, err error)
	// Add implements add.
	Add(context.Context, *IndexPayload) (res *Operation, err error)
	// Remove implements remove.
	Remove(context.Context, *IndexPayload) (res *Operation, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "indexes"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [3]string{"list", "add", "remove"}

// IndexPayload is the payload type of the indexes service add method.
type IndexPayload struct {
	// The url of the index file
	URL string
}

// Operation is the result type of the indexes service add method.
type Operation struct {
	// The status of the operation
	Status string
}

// MakeInvalidURL builds a goa.ServiceError from an error.
func MakeInvalidURL(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "invalid_url",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// NewOperation initializes result type Operation from viewed result type
// Operation.
func NewOperation(vres *indexesviews.Operation) *Operation {
	var res *Operation
	switch vres.View {
	case "default", "":
		res = newOperation(vres.Projected)
	}
	return res
}

// NewViewedOperation initializes viewed result type Operation from result type
// Operation using the given view.
func NewViewedOperation(res *Operation, view string) *indexesviews.Operation {
	var vres *indexesviews.Operation
	switch view {
	case "default", "":
		p := newOperationView(res)
		vres = &indexesviews.Operation{p, "default"}
	}
	return vres
}

// newOperation converts projected type Operation to service type Operation.
func newOperation(vres *indexesviews.OperationView) *Operation {
	res := &Operation{}
	if vres.Status != nil {
		res.Status = *vres.Status
	}
	return res
}

// newOperationView projects result type Operation into projected type
// OperationView using the "default" view.
func newOperationView(res *Operation) *indexesviews.OperationView {
	vres := &indexesviews.OperationView{
		Status: &res.Status,
	}
	return vres
}
