package checkvarname_test

import (
	"testing"

	"github.com/Makoto87/checkvarname"
	"github.com/gostaticanalysis/testutil"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	// テストデータ、静的解析、テストデータのパッケージ名
	analysistest.Run(t, testdata, checkvarname.Analyzer, "a")
}
