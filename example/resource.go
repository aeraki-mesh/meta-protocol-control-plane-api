// Copyright 2020 Envoyproxy Authors
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.
package example

import (
	metaRoute "github.com/aeraki-framework/go-control-plane/api/v1alpha"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
)

func makeRoute() *metaRoute.RouteConfiguration {
	return &metaRoute.RouteConfiguration{
		Name: "test",
		Routes: []*metaRoute.Route{
			{
				Name: "default",
				Route: &metaRoute.RouteAction{
					ClusterSpecifier: &metaRoute.RouteAction_Cluster{
						Cluster: "outbound|20880||org.apache.dubbo.samples.basic.api.demoservice",
					},
				},
			},
		},
	}
}

func GenerateSnapshot() cache.Snapshot {
	return cache.NewSnapshot(
		"1",
		[]types.Resource{}, // endpoints
		[]types.Resource{}, // clusters
		[]types.Resource{makeRoute()},
		[]types.Resource{}, //listeners
		[]types.Resource{}, // runtimes
		[]types.Resource{}, // secrets
	)
}
