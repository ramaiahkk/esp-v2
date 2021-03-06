// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package platform

// Keys for all files
type RuntimeFile int

const (
	// Javascript files
	GrpcBookstore RuntimeFile = iota

	// Binaries
	Bootstrapper
	ConfigManager
	Echo
	Envoy
	GrpcEchoClient
	GrpcEchoServer
	GrpcInteropClient
	GrpcInteropServer
	GrpcInteropStressClient

	// Proto descriptors
	FakeGrpcEchoDescriptor
	FakeGrpcInteropDescriptor
	FakeGrpcBookstoreDescriptor

	// Other files
	ServerCert
	ServerKey
	ProxyCert
	ProxyKey
	LogMetrics
	Version
	AccessLog
	ServiceAccountFile
	TestRootCaCerts

	// Configurations from examples directory
	AuthServiceConfig
	AuthEnvoyConfig
	ScServiceConfig
	ScEnvoyConfig
	DrServiceConfig
	DrEnvoyConfig
	RmServiceConfig
	RmEnvoyConfig
	SbServiceConfig
	SbEnvoyConfig
	GrpcEchoServiceConfig
	GrpcEchoEnvoyConfig

	// Other configurations for testing
	FixedDrServiceConfig
	FakeServiceAccountFile
)

// go tests are not executed from the root fo the repository.
// So the file path will change based on which test is using the file.
// i.e. unit tests are more deeply nested that integration tests.
var fileMap = map[RuntimeFile]string{

	// Used by integration test library.
	Version: "../../VERSION",

	// Used by integration tests.
	Bootstrapper:                "../../bin/bootstrap",
	ConfigManager:               "../../bin/configmanager",
	Echo:                        "../../bin/echo/server",
	Envoy:                       "../../bin/envoy",
	GrpcEchoClient:              "../../bin/grpc_echo_client",
	GrpcEchoServer:              "../../bin/grpc_echo_server",
	GrpcInteropClient:           "../../bin/interop_client",
	GrpcInteropServer:           "../../bin/interop_server",
	GrpcInteropStressClient:     "../../bin/stress_test",
	GrpcBookstore:               "../endpoints/bookstore_grpc/grpc_server.js",
	FakeGrpcEchoDescriptor:      "../endpoints/grpc_echo/proto/api_descriptor.pb",
	FakeGrpcInteropDescriptor:   "../endpoints/grpc_interop/proto/api_descriptor.pb",
	FakeGrpcBookstoreDescriptor: "../endpoints/bookstore_grpc/proto/api_descriptor.pb",
	ServerCert:                  "../env/testdata/server.crt",
	ServerKey:                   "../env/testdata/server.key",
	ProxyCert:                   "../env/testdata/proxy.crt",
	ProxyKey:                    "../env/testdata/proxy.key",
	LogMetrics:                  "../env/testdata/logs_metrics.pb.txt",
	AccessLog:                   "../env/testdata/access_log.txt",

	// Used by static bootstrap unit tests.
	AuthServiceConfig:     "../../../../examples/auth/service_config_generated.json",
	AuthEnvoyConfig:       "../../../../examples/auth/envoy_config.json",
	ScServiceConfig:       "../../../../examples/service_control/service_config_generated.json",
	ScEnvoyConfig:         "../../../../examples/service_control/envoy_config.json",
	DrServiceConfig:       "../../../../examples/dynamic_routing/service_config_generated.json",
	DrEnvoyConfig:         "../../../../examples/dynamic_routing/envoy_config.json",
	RmServiceConfig:       "../../../../examples/testdata/route_match/service_config_generated.json",
	RmEnvoyConfig:         "../../../../examples/testdata/route_match/envoy_config.json",
	SbServiceConfig:       "../../../../examples/testdata/sidecar_backend/service_config_generated.json",
	SbEnvoyConfig:         "../../../../examples/testdata/sidecar_backend/envoy_config.json",
	GrpcEchoServiceConfig: "../../../../examples/grpc_dynamic_routing/service_config_generated.json",
	GrpcEchoEnvoyConfig:   "../../../../examples/grpc_dynamic_routing/envoy_config.json",

	// Used by other unit tests.
	TestRootCaCerts:        "../../../tests/env/testdata/roots.pem",
	FixedDrServiceConfig:   "../../../tests/env/testdata/service_config_for_fixed_dynamic_routing.json",
	FakeServiceAccountFile: "./service_account.json",
}

// Get the runtime file path for the specified file.
func GetFilePath(file RuntimeFile) string {
	return fileMap[file]
}
