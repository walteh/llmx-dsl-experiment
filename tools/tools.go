package tools

//go:generate go run ./generate

import (
	_ "github.com/bufbuild/buf/cmd/buf"
	_ "github.com/go-task/task/v3/cmd/task"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/vektra/mockery/v2"
	_ "github.com/walteh/retab/cmd/retab"
	_ "gotest.tools/gotestsum"
)
