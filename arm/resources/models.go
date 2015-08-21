package resources

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator 0.11.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/azure/go-autorest/autorest"
	"github.com/azure/go-autorest/autorest/date"
)

type DeploymentMode string

const (
	Incremental DeploymentMode = "Incremental"
)

// Deployment dependency information.
type BasicDependency struct {
	Id           string `json:"id,omitempty"`
	ResourceType string `json:"resourceType,omitempty"`
	ResourceName string `json:"resourceName,omitempty"`
}

// Deployment dependency information.
type Dependency struct {
	DependsOn    []BasicDependency `json:"dependsOn,omitempty"`
	Id           string            `json:"id,omitempty"`
	ResourceType string            `json:"resourceType,omitempty"`
	ResourceName string            `json:"resourceName,omitempty"`
}

// Deployment operation parameters.
type Deployment struct {
	Properties struct {
		Template     map[string]string `json:"template,omitempty"`
		TemplateLink struct {
			Uri            string `json:"uri,omitempty"`
			ContentVersion string `json:"contentVersion,omitempty"`
		} `json:"templateLink,omitempty"`
		Parameters     map[string]string `json:"parameters,omitempty"`
		ParametersLink struct {
			Uri            string `json:"uri,omitempty"`
			ContentVersion string `json:"contentVersion,omitempty"`
		} `json:"parametersLink,omitempty"`
		Mode DeploymentMode `json:"mode,omitempty"`
	} `json:"properties,omitempty"`
}

// Deployment information.
type DeploymentExtended struct {
	autorest.Response `json:"-"`
	Id                string `json:"id,omitempty"`
	Name              string `json:"name,omitempty"`
	Properties        struct {
		ProvisioningState string            `json:"provisioningState,omitempty"`
		CorrelationId     string            `json:"correlationId,omitempty"`
		Timestamp         date.Time         `json:"timestamp,omitempty"`
		Outputs           map[string]string `json:"outputs,omitempty"`
		Providers         []Provider        `json:"providers,omitempty"`
		Dependencies      []Dependency      `json:"dependencies,omitempty"`
		Template          map[string]string `json:"template,omitempty"`
		TemplateLink      struct {
			Uri            string `json:"uri,omitempty"`
			ContentVersion string `json:"contentVersion,omitempty"`
		} `json:"templateLink,omitempty"`
		Parameters     map[string]string `json:"parameters,omitempty"`
		ParametersLink struct {
			Uri            string `json:"uri,omitempty"`
			ContentVersion string `json:"contentVersion,omitempty"`
		} `json:"parametersLink,omitempty"`
		Mode DeploymentMode `json:"mode,omitempty"`
	} `json:"properties,omitempty"`
}

// List of deployments.
type DeploymentListResult struct {
	autorest.Response `json:"-"`
	Value             []DeploymentExtended `json:"value,omitempty"`
	NextLink          string               `json:"nextLink,omitempty"`
}

// Deployment operation information.
type DeploymentOperation struct {
	autorest.Response `json:"-"`
	Id                string `json:"id,omitempty"`
	OperationId       string `json:"operationId,omitempty"`
	Properties        struct {
		ProvisioningState string            `json:"provisioningState,omitempty"`
		Timestamp         date.Time         `json:"timestamp,omitempty"`
		StatusCode        string            `json:"statusCode,omitempty"`
		StatusMessage     map[string]string `json:"statusMessage,omitempty"`
		TargetResource    struct {
			Id           string `json:"id,omitempty"`
			ResourceName string `json:"resourceName,omitempty"`
			ResourceType string `json:"resourceType,omitempty"`
		} `json:"targetResource,omitempty"`
	} `json:"properties,omitempty"`
}

// List of deployment operations.
type DeploymentOperationsListResult struct {
	autorest.Response `json:"-"`
	Value             []DeploymentOperation `json:"value,omitempty"`
	NextLink          string                `json:"nextLink,omitempty"`
}

// Information from validate template deployment response.
type DeploymentValidateResult struct {
	autorest.Response `json:"-"`
	Error             struct {
		Details []ResourceManagementError `json:"details,omitempty"`
		Code    string                    `json:"code,omitempty"`
		Message string                    `json:"message,omitempty"`
		Target  string                    `json:"target,omitempty"`
	} `json:"error,omitempty"`
	Properties struct {
		ProvisioningState string            `json:"provisioningState,omitempty"`
		CorrelationId     string            `json:"correlationId,omitempty"`
		Timestamp         date.Time         `json:"timestamp,omitempty"`
		Outputs           map[string]string `json:"outputs,omitempty"`
		Providers         []Provider        `json:"providers,omitempty"`
		Dependencies      []Dependency      `json:"dependencies,omitempty"`
		Template          map[string]string `json:"template,omitempty"`
		TemplateLink      struct {
			Uri            string `json:"uri,omitempty"`
			ContentVersion string `json:"contentVersion,omitempty"`
		} `json:"templateLink,omitempty"`
		Parameters     map[string]string `json:"parameters,omitempty"`
		ParametersLink struct {
			Uri            string `json:"uri,omitempty"`
			ContentVersion string `json:"contentVersion,omitempty"`
		} `json:"parametersLink,omitempty"`
		Mode DeploymentMode `json:"mode,omitempty"`
	} `json:"properties,omitempty"`
}

// Resource information.
type GenericResource struct {
	autorest.Response `json:"-"`
	Id                string            `json:"id,omitempty"`
	Name              string            `json:"name,omitempty"`
	Type              string            `json:"type,omitempty"`
	Location          string            `json:"location,omitempty"`
	Tags              map[string]string `json:"tags,omitempty"`
	Plan              struct {
		Name          string `json:"name,omitempty"`
		Publisher     string `json:"publisher,omitempty"`
		Product       string `json:"product,omitempty"`
		PromotionCode string `json:"promotionCode,omitempty"`
	} `json:"plan,omitempty"`
	Properties map[string]string `json:"properties,omitempty"`
}

// Resource provider information.
type Provider struct {
	autorest.Response `json:"-"`
	Id                string                 `json:"id,omitempty"`
	Namespace         string                 `json:"namespace,omitempty"`
	RegistrationState string                 `json:"registrationState,omitempty"`
	ResourceTypes     []ProviderResourceType `json:"resourceTypes,omitempty"`
}

// List of resource providers.
type ProviderListResult struct {
	autorest.Response `json:"-"`
	Value             []Provider `json:"value,omitempty"`
	NextLink          string     `json:"nextLink,omitempty"`
}

// Resource type managed by the resource provider.
type ProviderResourceType struct {
	ResourceType string            `json:"resourceType,omitempty"`
	Locations    []string          `json:"locations,omitempty"`
	ApiVersions  []string          `json:"apiVersions,omitempty"`
	Properties   map[string]string `json:"properties,omitempty"`
}

// Resource group information.
type ResourceGroup struct {
	autorest.Response `json:"-"`
	Id                string `json:"id,omitempty"`
	Name              string `json:"name,omitempty"`
	Properties        struct {
		ProvisioningState string `json:"provisioningState,omitempty"`
	} `json:"properties,omitempty"`
	Location string            `json:"location,omitempty"`
	Tags     map[string]string `json:"tags,omitempty"`
}

// List of resource groups.
type ResourceGroupListResult struct {
	autorest.Response `json:"-"`
	Value             []ResourceGroup `json:"value,omitempty"`
	NextLink          string          `json:"nextLink,omitempty"`
}

// List of resource groups.
type ResourceListResult struct {
	autorest.Response `json:"-"`
	Value             []GenericResource `json:"value,omitempty"`
	NextLink          string            `json:"nextLink,omitempty"`
}

type ResourceManagementError struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Target  string `json:"target,omitempty"`
}

// Resource provider operation information.
type ResourceProviderOperationDefinition struct {
	Name    string `json:"name,omitempty"`
	Display struct {
		Publisher   string `json:"publisher,omitempty"`
		Provider    string `json:"provider,omitempty"`
		Resource    string `json:"resource,omitempty"`
		Operation   string `json:"operation,omitempty"`
		Description string `json:"description,omitempty"`
	} `json:"display,omitempty"`
}

// List of resource provider operations.
type ResourceProviderOperationDetailListResult struct {
	autorest.Response `json:"-"`
	Value             []ResourceProviderOperationDefinition `json:"value,omitempty"`
}

// Parameters of move resources.
type ResourcesMoveInfo struct {
	Resources           []string `json:"resources,omitempty"`
	TargetResourceGroup string   `json:"targetResourceGroup,omitempty"`
}

// Tag details.
type TagDetails struct {
	autorest.Response `json:"-"`
	Id                string `json:"id,omitempty"`
	TagName           string `json:"tagName,omitempty"`
	Count             struct {
		Type  string `json:"type,omitempty"`
		Value string `json:"value,omitempty"`
	} `json:"count,omitempty"`
	Values []TagValue `json:"values,omitempty"`
}

// List of subscription tags.
type TagsListResult struct {
	autorest.Response `json:"-"`
	Value             []TagDetails `json:"value,omitempty"`
	NextLink          string       `json:"nextLink,omitempty"`
}

// Tag information.
type TagValue struct {
	autorest.Response `json:"-"`
	Id                string `json:"id,omitempty"`
	TagValueProperty  string `json:"tagValue,omitempty"`
	Count             struct {
		Type  string `json:"type,omitempty"`
		Value string `json:"value,omitempty"`
	} `json:"count,omitempty"`
}
