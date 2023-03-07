package main

import (
	"log"

	"github.com/apache/apisix-go-plugin-runner/pkg/plugin"
	"github.com/apache/apisix-go-plugin-runner/pkg/runner"
	"github.com/daemon1024/jwt2header/plugins"
	"go.uber.org/zap/zapcore"
)

func main() {
	cfg := runner.RunnerConfig{}
	cfg.LogLevel = zapcore.DebugLevel
	err := plugin.RegisterPlugin(&plugins.JWT2HeaderPlugin{})
	if err != nil {
		log.Fatalf("failed to register plugin JWT2HeaderPlugin: %s", err)
	}
	runner.Run(cfg)
}
