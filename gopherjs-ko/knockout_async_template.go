package ko

import (
	"github.com/gopherjs/gopherjs/js"
)

type AsynchronousTemplateConfig struct {
	*js.Object
}

func AsyncTemplateConfig() *AsynchronousTemplateConfig {
	return &AsynchronousTemplateConfig{
		Object: ko().Get("externalTemplateEngine"),
	}
}

// Set the default path/dir to load templates asynchronously
func (a *AsynchronousTemplateConfig) SetDefaultPath(path string) {
	a.Set("defaultPath", path)
}

// Set the default template suffix
func (a *AsynchronousTemplateConfig) SetDefaultSuffix(suffix string) {
	a.Set("defaultSuffix", suffix)
}
