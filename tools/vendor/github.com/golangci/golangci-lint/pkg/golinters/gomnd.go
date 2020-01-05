package golinters

import (
	"github.com/golangci/golangci-lint/pkg/golinters/goanalysis"
	mnd "github.com/tommy-muehle/go-mnd"
	"golang.org/x/tools/go/analysis"
)

func NewGoMND() *goanalysis.Linter {
	analyzers := []*analysis.Analyzer{
		mnd.Analyzer,
	}

	return goanalysis.NewLinter(
		"gomnd",
		"An analyzer to detect magic numbers.",
		analyzers,
		nil,
	).WithLoadMode(goanalysis.LoadModeSyntax)
}
