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
	route "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
)

func makeRoute() *route.RouteConfiguration {
	return &route.RouteConfiguration{
		Name: "test",
		VirtualHosts: []*route.VirtualHost{{
			Name: "meta-protocol-route",
			Domains: []string{"*"},
			Routes: []*route.Route{{
				Name: "default",
				Match: &route.RouteMatch{
					PathSpecifier: &route.RouteMatch_Prefix{
						Prefix: "/",
					},
					Headers: []*route.HeaderMatcher{
						{
							Name:"interface",
							HeaderMatchSpecifier: &route.HeaderMatcher_ExactMatch{
								ExactMatch:"org.apache.dubbo.samples.basic.api.DemoService",
							},
						},
						{
							Name:"method",
							HeaderMatchSpecifier: &route.HeaderMatcher_ExactMatch{
								ExactMatch:"sayHello",
							},
						},
					},
				},
				Action: &route.Route_Route{
					Route: &route.RouteAction{
						ClusterSpecifier: &route.RouteAction_Cluster{
							Cluster: "outbound|20880||org.apache.dubbo.samples.basic.api.demoservice",
						},
					},
				},
			}},
		}},
	}
}
func GenerateSnapshot() cache.Snapshot {
	return cache.NewSnapshot(
		"1",
		[]types.Resource{}, // endpoints
		[]types.Resource{},// clusters
		[]types.Resource{makeRoute()},
		[]types.Resource{}, //listeners
		[]types.Resource{}, // runtimes
		[]types.Resource{}, // secrets
	)
}
