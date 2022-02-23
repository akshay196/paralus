// Code generated by go-swagger; DO NOT EDIT.

package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/RafaySystems/rcloud-base/components/common/api/def/clients/scheduler/models"
)

// ClusterCreateClusterReader is a Reader for the ClusterCreateCluster structure.
type ClusterCreateClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClusterCreateClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewClusterCreateClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewClusterCreateClusterCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewClusterCreateClusterForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewClusterCreateClusterNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewClusterCreateClusterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewClusterCreateClusterOK creates a ClusterCreateClusterOK with default headers values
func NewClusterCreateClusterOK() *ClusterCreateClusterOK {
	return &ClusterCreateClusterOK{}
}

/* ClusterCreateClusterOK describes a response with status code 200, with default header values.

A successful response.
*/
type ClusterCreateClusterOK struct {
	Payload *models.V3Cluster
}

func (o *ClusterCreateClusterOK) Error() string {
	return fmt.Sprintf("[POST /infra/v3/project/{metadata.project}/cluster][%d] clusterCreateClusterOK  %+v", 200, o.Payload)
}
func (o *ClusterCreateClusterOK) GetPayload() *models.V3Cluster {
	return o.Payload
}

func (o *ClusterCreateClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V3Cluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClusterCreateClusterCreated creates a ClusterCreateClusterCreated with default headers values
func NewClusterCreateClusterCreated() *ClusterCreateClusterCreated {
	return &ClusterCreateClusterCreated{}
}

/* ClusterCreateClusterCreated describes a response with status code 201, with default header values.

Returned when edge is created successfully.
*/
type ClusterCreateClusterCreated struct {
	Payload interface{}
}

func (o *ClusterCreateClusterCreated) Error() string {
	return fmt.Sprintf("[POST /infra/v3/project/{metadata.project}/cluster][%d] clusterCreateClusterCreated  %+v", 201, o.Payload)
}
func (o *ClusterCreateClusterCreated) GetPayload() interface{} {
	return o.Payload
}

func (o *ClusterCreateClusterCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClusterCreateClusterForbidden creates a ClusterCreateClusterForbidden with default headers values
func NewClusterCreateClusterForbidden() *ClusterCreateClusterForbidden {
	return &ClusterCreateClusterForbidden{}
}

/* ClusterCreateClusterForbidden describes a response with status code 403, with default header values.

Returned when the user does not have permission to access the resource.
*/
type ClusterCreateClusterForbidden struct {
	Payload interface{}
}

func (o *ClusterCreateClusterForbidden) Error() string {
	return fmt.Sprintf("[POST /infra/v3/project/{metadata.project}/cluster][%d] clusterCreateClusterForbidden  %+v", 403, o.Payload)
}
func (o *ClusterCreateClusterForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ClusterCreateClusterForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClusterCreateClusterNotFound creates a ClusterCreateClusterNotFound with default headers values
func NewClusterCreateClusterNotFound() *ClusterCreateClusterNotFound {
	return &ClusterCreateClusterNotFound{}
}

/* ClusterCreateClusterNotFound describes a response with status code 404, with default header values.

Returned when the resource does not exist.
*/
type ClusterCreateClusterNotFound struct {
	Payload string
}

func (o *ClusterCreateClusterNotFound) Error() string {
	return fmt.Sprintf("[POST /infra/v3/project/{metadata.project}/cluster][%d] clusterCreateClusterNotFound  %+v", 404, o.Payload)
}
func (o *ClusterCreateClusterNotFound) GetPayload() string {
	return o.Payload
}

func (o *ClusterCreateClusterNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClusterCreateClusterDefault creates a ClusterCreateClusterDefault with default headers values
func NewClusterCreateClusterDefault(code int) *ClusterCreateClusterDefault {
	return &ClusterCreateClusterDefault{
		_statusCode: code,
	}
}

/* ClusterCreateClusterDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ClusterCreateClusterDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// Code gets the status code for the cluster create cluster default response
func (o *ClusterCreateClusterDefault) Code() int {
	return o._statusCode
}

func (o *ClusterCreateClusterDefault) Error() string {
	return fmt.Sprintf("[POST /infra/v3/project/{metadata.project}/cluster][%d] Cluster_CreateCluster default  %+v", o._statusCode, o.Payload)
}
func (o *ClusterCreateClusterDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *ClusterCreateClusterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}