package utils

import (
	"regexp"
	"strconv"
	"strings"
)

// NutritionData represents the parsed nutrition facts
type NutritionData map[string]interface{}

// DailyValue represents daily value percentages
type DailyValue struct {
	TotalFatPercent         int `json:"Total Fat %,omitempty"`
	TotalCarbohydratePercent int `json:"Total Carbohydrate %,omitempty"`
}

// nutrientPatterns defines regex patterns for common nutrients
var nutrientPatterns = map[string]string{
	"Calories":           `calories\s+(\d+)`,
	"Total Fat":          `total\s+fat\s+(\d+)g`,
	"Saturated Fat":      `saturated\s+fat\s+(\d+)g`,
	"Trans Fat":          `trans\s+fat\s+(\d+)g`,
	"Polyunsaturated Fat": `polyunsaturated\s+fat\s+(\d+\.?\d*)g`,
	"Monounsaturated Fat": `monounsaturated\s+fat\s+(\d+\.?\d*)g`,
	"Cholesterol":        `cholesterol\s+(\d+)mg`,
	"Sodium":             `sodium\s+(\d+)mg`,
	"Potassium":          `potassium\s+(\d+)mg`,
	"Total Carbohydrate": `total\s+carbohydrate\s+(\d+)g`,
	"Dietary Fiber":      `dietary\s+fiber\s+(\d+)g`,
	"Sugars":             `sugars\s+(\d+)g`,
	"Protein":            `protein\s+(\d+)g`,
	"Vitamin A":          `vitamin\s+a\s+(\d*)`,
	"Vitamin C":          `vitamin\s+c\s+(\d*)`,
	"Calcium":            `calcium\s+(\d+)mg`,
	"Iron":               `iron\s+(\d+)mg`,
	"Vitamin D":          `vitamin\s+d\s+(\d*)`,
	"Thiamin":            `thiamin\s+(\d*)`,
	"Riboflavin":         `riboflavin\s+(\d*)`,
	"Niacin":             `niacin\s+(\d*)`,
	"Vitamin B6":         `vitamin\s+b6\s+(\d*)`,
	"Folic Acid":         `folic\s+acid\s+(\d*)`,
	"Vitamin B12":        `vitamin\s+b12\s+(\d*)`,
	"Pantothenic Acid":   `pantothenic\s+acid\s+(\d*)`,
}

// CleanText removes unnecessary characters, fixes OCR errors, and standardizes format
func CleanText(text string) string {
	// Remove newlines
	text = strings.ReplaceAll(text, "\n", " ")
	// Remove double spaces
	text = strings.ReplaceAll(text, "  ", " ")
	// Convert to lowercase
	text = strings.ToLower(text)

	// Remove non-ASCII characters
	nonASCII := regexp.MustCompile(`[^\x00-\x7F]+`)
	text = nonASCII.ReplaceAllString(text, "")

	// Remove unwanted symbols, keep only alphanumeric, %, ., ,, (, ), and space
	unwantedSymbols := regexp.MustCompile(`[^a-z0-9%.,() ]+`)
	text = unwantedSymbols.ReplaceAllString(text, "")

	return text
}

// extractNutrientData extracts nutrient values using regex patterns
func extractNutrientData(text string) NutritionData {
	data := make(NutritionData)

	for nutrient, pattern := range nutrientPatterns {
		re := regexp.MustCompile(pattern)
		matches := re.FindStringSubmatch(text)
		if len(matches) > 1 && matches[1] != "" {
			data[nutrient] = matches[1]
		}
	}

	return data
}

// ParseNutritionFacts cleans text and extracts nutrition data
func ParseNutritionFacts(text string) NutritionData {
	// Clean the extracted text
	cleanedText := CleanText(text)

	// Extract nutrient data
	data := extractNutrientData(cleanedText)

	// Look for "Daily Value" percentages
	dailyValueRe := regexp.MustCompile(`(\d+)%`)
	matches := dailyValueRe.FindAllStringSubmatch(cleanedText, -1)

	if len(matches) >= 2 {
		fatPercent, _ := strconv.Atoi(matches[0][1])
		carbPercent, _ := strconv.Atoi(matches[1][1])

		data["Daily Value"] = map[string]int{
			"Total Fat %":          fatPercent,
			"Total Carbohydrate %": carbPercent,
		}
	}

	return data
}
