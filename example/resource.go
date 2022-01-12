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
	metaroute "github.com/aeraki-mesh/meta-protocol-control-plane-api/meta_protocol_proxy/config/route/v1alpha"
	httproute "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
)

var metaRoute = metaroute.RouteConfiguration{
	Name: "test",
	Routes: []*metaroute.Route{
		{
			Name: "test",
			Match: &metaroute.RouteMatch{
				Metadata: []*httproute.HeaderMatcher{
					{
						Name: "interface",
						HeaderMatchSpecifier: &httproute.HeaderMatcher_ExactMatch{
							ExactMatch: "org.apache.dubbo.samples.basic.api.DemoService",
						},
					},
					{
						Name: "method",
						HeaderMatchSpecifier: &httproute.HeaderMatcher_ExactMatch{
							ExactMatch: "sayHello",
						},
					},
				},
			},
			Route: &metaroute.RouteAction{
				ClusterSpecifier: &metaroute.RouteAction_Cluster{
					Cluster: "outbound|20880||org.apache.dubbo.samples.basic.api.demoservice",
				},
			},
		},
	},
}

func makeRoute() *httproute.RouteConfiguration {
	return metaProtocolRoute2HttpRoute(metaRoute)
}

//We use Envoy RDS(HTTP RouteConfiguration) to transmit Meta Protocol Configuration from the RDS server to the Proxy
func metaProtocolRoute2HttpRoute(metaRoute metaroute.RouteConfiguration) *httproute.RouteConfiguration {
	httpRoute := &httproute.RouteConfiguration{
		Name: metaRoute.Name,
		VirtualHosts: []*httproute.VirtualHost{
			{
				Name:    "dummy",
				Domains: []string{"*"},
			},
		},
	}

	for _, route := range metaRoute.Routes {
		routeMatch := &httproute.RouteMatch{
			PathSpecifier: &httproute.RouteMatch_Prefix{
				Prefix: "/",
			},
		}
		for _, metadata := range route.Match.Metadata {
			routeMatch.Headers = append(routeMatch.Headers, &httproute.HeaderMatcher{
				Name:                 metadata.Name,
				HeaderMatchSpecifier: metadata.HeaderMatchSpecifier,
			})
		}

		var routeAction *httproute.RouteAction
		if route.Route.GetWeightedClusters() != nil {
			routeAction = &httproute.RouteAction{
				ClusterSpecifier: &httproute.RouteAction_WeightedClusters{
					WeightedClusters: route.Route.GetWeightedClusters(),
				},
			}
		} else {
			routeAction = &httproute.RouteAction{
				ClusterSpecifier: &httproute.RouteAction_Cluster{
					Cluster: route.Route.GetCluster(),
				},
			}
		}

		httpRoute.VirtualHosts[0].Routes = append(httpRoute.VirtualHosts[0].Routes, &httproute.Route{
			Name:  route.Name,
			Match: routeMatch,
			Action: &httproute.Route_Route{
				Route: routeAction,
			},
		})
	}

	return httpRoute
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
