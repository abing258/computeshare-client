package agent

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewAgentService)
