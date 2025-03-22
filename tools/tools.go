//go:build tools
package tools

import (
	_ "github.com/cosmos/cosmos-proto/cmd/protoc-gen-go-cosmos"
	_ "github.com/cosmos/gogoproto/protoc-gen-gocosmos"
	_ "github.com/golang/protobuf/protoc-gen-go"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2"
	_ "github.com/regen-network/cosmos-proto/protoc-gen-gocosmos"
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
)

// This file imports packages required for generating protobuf code and other tools.
// It is not included in the final binary and is only used during development.

// The build tag 'tools' ensures this file is excluded from regular builds
// and is only used when explicitly requested.

// Required tool versions:
// - Go: 1.19
// - CIRCL: v1.3.7
// - Ignite CLI: v0.27.1
// - CometBFT: Latest compatible with Cosmos SDK
// - Cosmos SDK: Compatible with Go 1.19

const (
	// These constants define the major tool versions used in the project
	GoVersion     = "1.19"
	CirclVersion  = "v1.3.7"
	IgniteVersion = "v0.27.1"
	_ "github.com/regen-network/cosmos-proto/protoc-gen-gocosmos"
	_ "github.com/cloudflare/circl"           // CIRCL v1.3.7 for quantum-safe cryptography
	_ "github.com/cometbft/cometbft"          // CometBFT for consensus
	_ "github.com/cosmos/cosmos-sdk"          // Cosmos SDK
	_ "github.com/cosmos/ibc-go/v7"           // IBC for cross-chain communication
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
)

// This file imports packages required for generating protobuf code and other tools.
// It is not included in the final binary and is only used during development.

// The build tag 'tools' ensures this file is excluded from regular builds
// and is only used when explicitly requested.

// Required tool versions:
// - Go: 1.19
// - CIRCL: v1.3.7
// - Ignite CLI: v0.27.1
// - CometBFT: Latest compatible with Cosmos SDK
// - Cosmos SDK: Compatible with Go 1.19

const (
	// These constants define the major tool versions used in the project
	GoVersion    = "1.19"
	CirclVersion = "v1.3.7"
	IgniteVersion = "v0.27.1"
)
