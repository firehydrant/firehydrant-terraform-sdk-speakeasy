// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/shared"
	"net/http"
)

type UpdateCustomFieldDefinitionRequest struct {
	FieldID                               string                                       `pathParam:"style=simple,explode=false,name=field_id"`
	PatchV1CustomFieldsDefinitionsFieldID shared.PatchV1CustomFieldsDefinitionsFieldID `request:"mediaType=application/json"`
}

func (o *UpdateCustomFieldDefinitionRequest) GetFieldID() string {
	if o == nil {
		return ""
	}
	return o.FieldID
}

func (o *UpdateCustomFieldDefinitionRequest) GetPatchV1CustomFieldsDefinitionsFieldID() shared.PatchV1CustomFieldsDefinitionsFieldID {
	if o == nil {
		return shared.PatchV1CustomFieldsDefinitionsFieldID{}
	}
	return o.PatchV1CustomFieldsDefinitionsFieldID
}

type UpdateCustomFieldDefinitionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Update a single custom field definition
	OrganizationsCustomFieldDefinitionEntity *shared.OrganizationsCustomFieldDefinitionEntity
	// A collection of codes that generally means the end user got something wrong in making the request
	BadRequest *shared.BadRequest
	// A collection of codes that generally means the client was not authenticated correctly for the request they want to make
	Unauthorized *shared.Unauthorized
	// Status codes relating to the resource/entity they are requesting not being found or endpoints/routes not existing
	NotFound *shared.NotFound
	// Status codes relating to the client being rate limited by the server
	RateLimited *shared.RateLimited
	// A collection of status codes that generally mean the server failed in an unexpected way
	InternalServerError *shared.InternalServerError
	// Timeouts occurred with the request
	Timeout *shared.Timeout
}

func (o *UpdateCustomFieldDefinitionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateCustomFieldDefinitionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateCustomFieldDefinitionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateCustomFieldDefinitionResponse) GetOrganizationsCustomFieldDefinitionEntity() *shared.OrganizationsCustomFieldDefinitionEntity {
	if o == nil {
		return nil
	}
	return o.OrganizationsCustomFieldDefinitionEntity
}

func (o *UpdateCustomFieldDefinitionResponse) GetBadRequest() *shared.BadRequest {
	if o == nil {
		return nil
	}
	return o.BadRequest
}

func (o *UpdateCustomFieldDefinitionResponse) GetUnauthorized() *shared.Unauthorized {
	if o == nil {
		return nil
	}
	return o.Unauthorized
}

func (o *UpdateCustomFieldDefinitionResponse) GetNotFound() *shared.NotFound {
	if o == nil {
		return nil
	}
	return o.NotFound
}

func (o *UpdateCustomFieldDefinitionResponse) GetRateLimited() *shared.RateLimited {
	if o == nil {
		return nil
	}
	return o.RateLimited
}

func (o *UpdateCustomFieldDefinitionResponse) GetInternalServerError() *shared.InternalServerError {
	if o == nil {
		return nil
	}
	return o.InternalServerError
}

func (o *UpdateCustomFieldDefinitionResponse) GetTimeout() *shared.Timeout {
	if o == nil {
		return nil
	}
	return o.Timeout
}
