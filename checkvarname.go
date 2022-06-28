package checkvarname

import (
	"go/ast"
	"regexp"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "checkvarname is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "checkvarname",
	Doc:  doc,
	Run:  run, // 静的解析の処理
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

// pass: 構文解析や型チェック、依存するAnalyserの結果など
// skelton作成後、run関数のみ変更すれば良い
func run(pass *analysis.Pass) (any, error) {
	// inspect.Analyzerの解析結果
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.Ident)(nil),
	}

	// 深さ優先探索
	inspect.Preorder(nodeFilter, func(n ast.Node) {
		// 識別子の場合に処理を行う
		switch n := n.(type) {
		case *ast.Ident:
			if regexp.MustCompile(`_`).MatchString(n.Name) || regexp.MustCompile(`-`).MatchString(n.Name) {
				// 指定した位置でエラーの報告
				pass.Reportf(n.Pos(), "_ or - is used")
			}
		}
	})

	return nil, nil
}
