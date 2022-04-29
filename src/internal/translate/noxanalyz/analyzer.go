package noxanalyz

import (
	"log"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const Doc = `applies different Nox-specific refactorings`

var Analyzer = &analysis.Analyzer{
	Name:             "noxanalyz",
	Doc:              Doc,
	Run:              run,
	RunDespiteErrors: true,
	Requires:         []*analysis.Analyzer{inspect.Analyzer},
	FactTypes:        []analysis.Fact{new(testFact)},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	// collect imports
	//inspect.Preorder([]ast.Node{(*ast.ImportSpec)(nil)})
	_ = inspect
	for _, f := range pass.Files {
		log.Printf("noxanalyz: %q", f.Name.Name)
	}
	return nil, nil
}

type testFact struct{}

func (*testFact) String() string { return "test fact" }
func (*testFact) AFact()         {}
