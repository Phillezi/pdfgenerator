package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/jung-kurt/gofpdf"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func LoadConfig() bool {
	viper.SetConfigName(".pdfconfig.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		return false
	}
	return true
}

func Run(cmd *cobra.Command, args []string) {

	readConf := LoadConfig()
	var includePatterns, excludePatterns []string

	if readConf {
		includePatterns = viper.GetStringSlice("includePatterns")
		excludePatterns = viper.GetStringSlice("excludePatterns")

	} else {
		includePatterns, excludePatterns = parsePatterns(args)
	}

	fmt.Println("Incl: ", includePatterns)
	fmt.Println("Excl: ", excludePatterns)

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	for _, pattern := range includePatterns {
		matches, err := filepath.Glob(pattern)
		if err != nil {
			fmt.Printf("Error with glob pattern %s: %v\n", pattern, err)
			continue
		}

		for _, match := range matches {
			err := filepath.Walk(match, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if info.IsDir() {
					return nil
				}

				matchedExclude := matchPatterns(path, excludePatterns)

				if !matchedExclude {
					fmt.Println("Adding file to PDF: " + path)
					appendFileContentsToPDF(pdf, path)
				}
				return nil
			})

			if err != nil {
				fmt.Printf("Error walking path %s: %v\n", match, err)
			}
		}
	}

	outputFile := viper.GetString("outputFile")
	err := pdf.OutputFileAndClose(outputFile)
	if err != nil {
		fmt.Printf("Error saving PDF: %v\n", err)
		return
	}

	fmt.Printf("PDF generated successfully: %s\n", outputFile)
}

func parsePatterns(args []string) ([]string, []string) {
	var includePatterns, excludePatterns []string
	for _, arg := range args {
		if strings.HasPrefix(arg, "!") {
			excludePatterns = append(excludePatterns, arg[1:])
		} else {
			includePatterns = append(includePatterns, arg)
		}
	}
	return includePatterns, excludePatterns
}

func compilePattern(pattern string) *regexp.Regexp {
	compiled, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Printf("Error compiling regex pattern %s: %v\n", pattern, err)
		os.Exit(1)
	}
	return compiled
}

func matchPatterns(path string, patterns []string) bool {
	for _, pattern := range patterns {
		matched, err := filepath.Match(pattern, path)
		if err != nil {
			fmt.Printf("Error matching pattern %s: %v\n", pattern, err)
			continue
		}
		if matched {
			return true
		}
	}
	return false
}

func appendFileContentsToPDF(pdf *gofpdf.Fpdf, path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", path, err)
		return
	}

	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(40, 10, fmt.Sprintf("File: %s", path))
	pdf.Ln(12)

	pdf.SetFont("Courier", "", 10)
	pdf.MultiCell(190, 6, string(data), "", "", false)
	pdf.Ln(10)

	fmt.Printf("Appended content from: %s\n", path)
}
