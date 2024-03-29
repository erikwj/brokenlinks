package internal_test

import (
	"testing"

	"github.com/erikwj/brokenlinks/internal"
)

var ext = ".md"
var regexs = internal.ExtDocRegex(ext)

func TestValidateWebLine(t *testing.T) {
	line := "[GitHub](http://github.com) (and some extra text) [Gitlab](http://gitlab.com) "
	lineNum := 1
	filePath := "/path/to/file.md"

	// Test your validateLine function here
	// with the given line, lineNum, and filePath variables as input
	err := internal.ValidateLine(line, lineNum, filePath, regexs, false)

	// Assert the expected result
	if err != nil {
		t.Errorf("Expected validateWebLine to pass, but it failed")
	}
}

func TestValidateWebLineRst(t *testing.T) {
	line := "Table definitions may be constructed either from scratch (check out `the syntax <https://nightlies.apache.org/flink/flink-docs-release-1.17/docs/dev/table/sql/create/#create-table>`_)"
	lineNum := 1
	filePath := "/path/to/file.md"
	r := internal.ExtDocRegex(".rst")
	// Test your validateLine function here
	// with the given line, lineNum, and filePath variables as input
	err := internal.ValidateLine(line, lineNum, filePath, r, false)

	// Assert the expected result
	if err != nil {
		t.Errorf("Expected validateWebLine to pass, but it failed")
	}
}

func TestValidateFileLine(t *testing.T) {
	line := "[glossary](../testfiles/glossary.md)"
	lineNum := 1
	filePath := "./"

	// Test your validateLine function here
	// with the given line, lineNum, and filePath variables as input
	err := internal.ValidateLine(line, lineNum, filePath, regexs, false)

	// Assert the expected result
	if err != nil {
		t.Errorf("Expected validateLine to pass, but it failed")
	}
}

func TestValidateLongFileLine(t *testing.T) {
	line := "alpha beta gamma [bla](../testfiles/subdir/bla.md) or a [correct](../testfiles/correct.md), ... [glossary](../testfiles/glossary.md) and [corrupt](../testfiles/corrupt"
	lineNum := 1
	filePath := "./"

	// Test your validateLine function here
	// with the given line, lineNum, and filePath variables as input
	err := internal.ValidateLine(line, lineNum, filePath, regexs, false)

	// Assert the expected result
	if err != nil {
		t.Errorf("Expected validateLine to pass, but it failed")
	}
}

func TestValidateImgLine(t *testing.T) {
	line := "![](../testfiles/img/btn.png)"
	lineNum := 1
	filePath := "./"

	// Test your validateLine function here
	// with the given line, lineNum, and filePath variables as input
	err := internal.ValidateLine(line, lineNum, filePath, regexs, false)

	// Assert the expected result
	if err != nil {
		t.Errorf("Expected validateLine to pass, but it failed")
	}
}

func TestValidateImgLineRst(t *testing.T) {
	line := " asldfa fasdf lsd ::image ../testfiles/img/btn.png"
	lineNum := 1
	filePath := "./"
	r := internal.ExtDocRegex(".rst")
	// Test your validateLine function here
	// with the given line, lineNum, and filePath variables as input
	err := internal.ValidateLine(line, lineNum, filePath, r, false)

	// Assert the expected result
	if err != nil {
		t.Errorf("Expected validateLine to pass, but it failed")
	}
}

// test for failure of ValidateLine with broken file link
func TestValidateFileLineFail(t *testing.T) {
	line := "[GitHub](broken.md)"
	lineNum := 1
	filePath := "/path/to/file.md"

	// Test your validateLine function here
	err := internal.ValidateLine(line, lineNum, filePath, regexs, false)

	// Assert that the function fails
	if err == nil {
		t.Errorf("Expected validateLine to fail, but it succeeded")
	}
}

// test for failure of ValidateLine with broken image link
func TestValidateImageLineFail(t *testing.T) {
	line := "![](broken.png)"
	lineNum := 1
	filePath := "/path/to/file.md"

	// Test your validateLine function here
	err := internal.ValidateLine(line, lineNum, filePath, regexs, false)

	// Assert that the function fails
	if err == nil {
		t.Errorf("Expected validateLine to fail, but it succeeded")
	}
}

// test for failure of ValidateLine with broken image link
func TestValidateImageLineFailRst(t *testing.T) {
	line := " asldfa fasdf lsd ::image ../testfiles/img/broken.png"
	lineNum := 1
	filePath := "/path/to/file.md"
	r := internal.ExtDocRegex(".rst")

	// Test your validateLine function here
	err := internal.ValidateLine(line, lineNum, filePath, r, false)

	// Assert that the function fails
	if err == nil {
		t.Errorf("Expected validateLine to fail, but it succeeded")
	}
}
