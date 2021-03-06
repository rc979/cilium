// Copyright 2016-2017 Authors of Cilium
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

package client

import (
	"github.com/cilium/cilium/api/v1/client/policy"
	"github.com/cilium/cilium/api/v1/models"
)

// PolicyPut inserts the `policyJSON`
func (c *Client) PolicyPut(policyJSON string) (string, error) {
	params := policy.NewPutPolicyParams().WithPolicy(&policyJSON)
	resp, err := c.Policy.PutPolicy(params)
	if err != nil {
		return "", err
	}
	return string(resp.Payload), nil
}

// PolicyGet returns policy rules
func (c *Client) PolicyGet(labels []string) (string, error) {
	params := policy.NewGetPolicyParams().WithLabels(labels)
	resp, err := c.Policy.GetPolicy(params)
	if err != nil {
		return "", err
	}
	return string(resp.Payload), nil
}

// PolicyDelete deletes policy rules
func (c *Client) PolicyDelete(labels []string) error {
	params := policy.NewDeletePolicyParams().WithLabels(labels)
	_, err := c.Policy.DeletePolicy(params)
	return err
}

// PolicyResolveGet resolves policy for a context with source and destination identity.
func (c *Client) PolicyResolveGet(context *models.IdentityContext) (*models.PolicyTraceResult, error) {
	params := policy.NewGetPolicyResolveParams().WithIdentityContext(context)
	resp, err := c.Policy.GetPolicyResolve(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}
