// /*
// Copyright The Kubernetes Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// */

// Code generated by client-gen. DO NOT EDIT.
package virtualnetworkclient

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armnetwork "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v4"

	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/utils"
)

type Client struct {
	*armnetwork.VirtualNetworksClient
	subscriptionID string
}

func New(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (Interface, error) {
	if options == nil {
		options = utils.GetDefaultOption()
	}

	client, err := armnetwork.NewVirtualNetworksClient(subscriptionID, credential, options)
	if err != nil {
		return nil, err
	}
	return &Client{client, subscriptionID}, nil
}

// Get gets the VirtualNetwork
func (client *Client) Get(ctx context.Context, resourceGroupName string, resourceName string, expand *string) (result *armnetwork.VirtualNetwork, rerr error) {
	var ops *armnetwork.VirtualNetworksClientGetOptions
	if expand != nil {
		ops = &armnetwork.VirtualNetworksClientGetOptions{Expand: expand}
	}
	ctx = utils.ContextWithClientName(ctx, "VirtualNetworksClient")
	ctx = utils.ContextWithRequestMethod(ctx, "Get")
	ctx = utils.ContextWithResourceGroupName(ctx, resourceGroupName)
	ctx = utils.ContextWithSubscriptionID(ctx, client.subscriptionID)
	resp, err := client.VirtualNetworksClient.Get(ctx, resourceGroupName, resourceName, ops)
	if err != nil {
		return nil, err
	}
	//handle statuscode
	return &resp.VirtualNetwork, nil
}

// CreateOrUpdate creates or updates a VirtualNetwork.
func (client *Client) CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, resource armnetwork.VirtualNetwork) (*armnetwork.VirtualNetwork, error) {
	ctx = utils.ContextWithClientName(ctx, "VirtualNetworksClient")
	ctx = utils.ContextWithRequestMethod(ctx, "CreateOrUpdate")
	ctx = utils.ContextWithResourceGroupName(ctx, resourceGroupName)
	ctx = utils.ContextWithSubscriptionID(ctx, client.subscriptionID)
	resp, err := utils.NewPollerWrapper(client.VirtualNetworksClient.BeginCreateOrUpdate(ctx, resourceGroupName, resourceName, resource, nil)).WaitforPollerResp(ctx)
	if err != nil {
		return nil, err
	}
	if resp != nil {
		return &resp.VirtualNetwork, nil
	}
	return nil, nil
}

// Delete deletes a VirtualNetwork by name.
func (client *Client) Delete(ctx context.Context, resourceGroupName string, resourceName string) error {
	ctx = utils.ContextWithClientName(ctx, "VirtualNetworksClient")
	ctx = utils.ContextWithRequestMethod(ctx, "Delete")
	ctx = utils.ContextWithResourceGroupName(ctx, resourceGroupName)
	ctx = utils.ContextWithSubscriptionID(ctx, client.subscriptionID)
	_, err := utils.NewPollerWrapper(client.BeginDelete(ctx, resourceGroupName, resourceName, nil)).WaitforPollerResp(ctx)
	return err
}

// List gets a list of VirtualNetwork in the resource group.
func (client *Client) List(ctx context.Context, resourceGroupName string) (result []*armnetwork.VirtualNetwork, rerr error) {
	ctx = utils.ContextWithClientName(ctx, "VirtualNetworksClient")
	ctx = utils.ContextWithRequestMethod(ctx, "List")
	ctx = utils.ContextWithResourceGroupName(ctx, resourceGroupName)
	ctx = utils.ContextWithSubscriptionID(ctx, client.subscriptionID)
	pager := client.VirtualNetworksClient.NewListPager(resourceGroupName, nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		result = append(result, nextResult.Value...)
	}
	return result, nil
}
