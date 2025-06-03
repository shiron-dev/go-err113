package err113

import (
	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"
)

func init() {
	register.Plugin("shironerr113", New)
}

type plugin struct{}

func New(_ any) (register.LinterPlugin, error) {
	return &plugin{}, nil
}

func (p *plugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{
		NewAnalyzer(),
	}, nil
}

func (p *plugin) GetLoadMode() string {
	return register.LoadModeSyntax
}
