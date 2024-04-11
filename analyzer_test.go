package linguachecker

import (
	"testing"

	// "github.com/tj/assert"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestLinguachecker(t *testing.T) {
	tdDir := analysistest.TestData()
	tCases := []struct {
		name           string
		include        string
		exclude        string
		testDataDir    string
		expectedErrors int
	}{
		{
			name:           "no include or exclude(includes all languages)",
			include:        "",
			exclude:        "",
			testDataDir:    tdDir + "/pack",
			expectedErrors: 3, // Update with the expected number of errors
		},
		{
			name:           "include only English",
			include:        "en",
			exclude:        "",
			testDataDir:    tdDir + "/pack2",
			expectedErrors: 1, // Update with the expected number of errors
		},
		{
			name:           "exclude Russian",
			include:        "",
			exclude:        "ru",
			testDataDir:    tdDir + "/pack2",
			expectedErrors: 1, // Update with the expected number of errors
		},
		// Add more test cases as needed
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(t *testing.T) {
			Analyzer.Flags.Set("include", tc.include)
			Analyzer.Flags.Set("exclude", tc.exclude)
			_ = analysistest.Run(t, tc.testDataDir, Analyzer)

			// // Count the number of errors reported
			// errorCount := 0
			// for _, r := range results {
			// 	errorCount += len(r.Diagnostics)
			// }

			// // Check if the number of errors matches the expected count
			// if errorCount != tc.expectedErrors {
			// 	t.Errorf("unexpected number of errors, got %d, want %d", errorCount, tc.expectedErrors)
			// }
		})
	}
}
