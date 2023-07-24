package ociinstaller

type pluginInstallConfig struct {
	skipConfigFile bool
}

type PluginInstallOption = func(config *pluginInstallConfig)

func WithSkipConfigEnabled(skipConfigFile bool) PluginInstallOption {
	return func(o *pluginInstallConfig) {
		o.skipConfigFile = skipConfigFile
	}
}
