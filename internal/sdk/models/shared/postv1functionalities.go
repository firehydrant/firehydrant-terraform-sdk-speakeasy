// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type Services struct {
	// ID of a service
	ID string `json:"id"`
}

func (o *Services) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type ExternalResources struct {
	RemoteID string `json:"remote_id"`
	// The integration slug for the external resource. Can be one of: github, opsgenie, pager_duty, statuspage, victorops. Not required if the resource has already been imported.
	ConnectionType *string `json:"connection_type,omitempty"`
}

func (o *ExternalResources) GetRemoteID() string {
	if o == nil {
		return ""
	}
	return o.RemoteID
}

func (o *ExternalResources) GetConnectionType() *string {
	if o == nil {
		return nil
	}
	return o.ConnectionType
}

type Links struct {
	// Short name used to display and identify this link
	Name string `json:"name"`
	// URL
	HrefURL string `json:"href_url"`
	// An optional URL to an icon representing this link
	IconURL *string `json:"icon_url,omitempty"`
}

func (o *Links) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Links) GetHrefURL() string {
	if o == nil {
		return ""
	}
	return o.HrefURL
}

func (o *Links) GetIconURL() *string {
	if o == nil {
		return nil
	}
	return o.IconURL
}

// Owner - An object representing a Team that owns the service
type Owner struct {
	ID string `json:"id"`
}

func (o *Owner) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type Teams struct {
	ID string `json:"id"`
}

func (o *Teams) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

// PostV1Functionalities - Creates a functionality for the organization
type PostV1Functionalities struct {
	Name        string     `json:"name"`
	Description *string    `json:"description,omitempty"`
	Services    []Services `json:"services,omitempty"`
	// A hash of label keys and values
	Labels                map[string]string `json:"labels,omitempty"`
	AlertOnAdd            *bool             `json:"alert_on_add,omitempty"`
	AutoAddRespondingTeam *bool             `json:"auto_add_responding_team,omitempty"`
	// An array of external resources to attach to this service.
	ExternalResources []ExternalResources `json:"external_resources,omitempty"`
	// An array of links to associate with this service
	Links []Links `json:"links,omitempty"`
	// An object representing a Team that owns the service
	Owner *Owner `json:"owner,omitempty"`
	// An array of teams to attach to this service.
	Teams []Teams `json:"teams,omitempty"`
}

func (o *PostV1Functionalities) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *PostV1Functionalities) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *PostV1Functionalities) GetServices() []Services {
	if o == nil {
		return nil
	}
	return o.Services
}

func (o *PostV1Functionalities) GetLabels() map[string]string {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *PostV1Functionalities) GetAlertOnAdd() *bool {
	if o == nil {
		return nil
	}
	return o.AlertOnAdd
}

func (o *PostV1Functionalities) GetAutoAddRespondingTeam() *bool {
	if o == nil {
		return nil
	}
	return o.AutoAddRespondingTeam
}

func (o *PostV1Functionalities) GetExternalResources() []ExternalResources {
	if o == nil {
		return nil
	}
	return o.ExternalResources
}

func (o *PostV1Functionalities) GetLinks() []Links {
	if o == nil {
		return nil
	}
	return o.Links
}

func (o *PostV1Functionalities) GetOwner() *Owner {
	if o == nil {
		return nil
	}
	return o.Owner
}

func (o *PostV1Functionalities) GetTeams() []Teams {
	if o == nil {
		return nil
	}
	return o.Teams
}