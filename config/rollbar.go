package config

import (
	"github.com/rollbar/rollbar-go"
	"github.com/spf13/viper"
)

func InitRollbar() {
	rollbar.SetToken(viper.GetString("rollbar.token"))
	rollbar.SetEnvironment(viper.GetString("rollbar.environment")) // defaults to "development"
	rollbar.SetCodeVersion(viper.GetString("rollbar.codeVersion")) // optional Git hash/branch/tag (required for GitHub integration)
	// rollbar.SetServerHost("web.1")                       // optional override; defaults to hostname
	rollbar.SetServerRoot(viper.GetString("rollbar.serverRoot")) // path of project (required for GitHub integration and non-project stacktrace collapsing)

}
