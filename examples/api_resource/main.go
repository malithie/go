/*
 * Copyright (c) 2025, WSO2 LLC. (http://www.wso2.com).
 *
 * WSO2 LLC. licenses this file to you under the Apache License,
 * Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package main

import (
	"context"
	"log"
	"time"

	"github.com/asgardeo/go/pkg/config"
	"github.com/asgardeo/go/pkg/sdk"
)

func main() {

	// Initialize the client configurations.
	cfg := config.DefaultClientConfig().
		WithBaseURL("https://api.asgardeo.io/t/<tenant-domain>").
		WithTimeout(10*time.Second).
		WithClientCredentials(
			"client_id",
			"client_secret",
		)

	// Create a client with the given configurations.
	client, err := sdk.New(cfg)
	if err != nil {
		log.Fatalf("Failed to create SDK client: %v", err)
	}

	// Use the client with token authentication.
	ctx := context.Background()

	// List API resources.
	apiResources, err := client.APIResource.List(ctx, nil)
	if err != nil {
		log.Printf("Error listing API Resources: %v", err)
	} else {
		log.Printf("Found %d API Resources.\n", len(*apiResources.APIResources))
	}

	// Get a specific API resource by ID.
	apiResource, err := client.APIResource.Get(ctx, "api_resource_uuid")
	if err != nil {
		log.Printf("Error getting API Resource: %v", err)
	} else {
		log.Printf("Found API Resource: %s\n", apiResource.Name)
	}

	// Get API Resources by name.
	apiResourcesByName, err := client.APIResource.GetByName(ctx, "api_resource_name")
	if err != nil {
		log.Printf("Error getting API Resources by name: %v", err)
	} else {
		log.Printf("Found %d API Resources by name.\n", len(*apiResourcesByName))
	}

	// Get API Resource By Identifier.
	apiResourcesByIdentifier, err := client.APIResource.GetByIdentifier(ctx, "api_resource_identifier")
	if err != nil {
		log.Printf("Error getting API Resource by identifier: %v", err)
	} else {
		log.Printf("Found API Resource by identifier: %s\n", apiResourcesByIdentifier.Name)
	}
}
