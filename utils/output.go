package utils

import (
	"bytes"
	"fmt"
	"strings"
)

func ParseCmdOutput(output []byte, expectedCountOfFields int) ([][]string, error) {
	lines := bytes.FieldsFunc(output, func(c rune) bool { return c == '\n' || c == '\r' })

	var recordLines [][]string
	for i, line := range lines {
		recordLine := splitBySeparator(":", string(line))
		if len(recordLine) != expectedCountOfFields {
			return nil, fmt.Errorf(
				"line %d contains %d fields but should %d",
				i, len(recordLine), expectedCountOfFields,
			)
		}

		recordLines = append(recordLines, recordLine)
	}

	return recordLines, nil
}

func endsWithBracketedNumber(b []byte) bool {
	// 检查最短长度，至少 "[1]"
	if len(b) < 3 || b[len(b)-1] != ']' {
		return false
	}

	// 从倒数第二个字符开始向前查找 '['
	for i := len(b) - 2; i >= 0; i-- {
		if b[i] == '[' {
			// 确保括号内至少有一个数字
			return i+1 < len(b)-1
		}
		// 如果出现非数字字符（在遇到 '[' 之前），立即返回 false
		if b[i] < '0' || b[i] > '9' {
			return false
		}
	}
	return false
}

func ParseCmdsHaveFieldNameOutput(output []byte) []map[string][][]string {
	Group := bytes.Split(output, []byte("\n\n"))
	res := make([]map[string][][]string, len(Group))
	for i, Item := range Group {
		res[i] = ParseCmdHaveFieldNameOutput(Item)
	}
	return res
}

func ParseCmdHaveFieldNameOutput(Item []byte) map[string][][]string {
	res := make(map[string][][]string)
	lines := bytes.FieldsFunc(Item, func(c rune) bool { return c == '\n' || c == '\r' })
	for _, line := range lines {
		FirstDelimiter := bytes.Index(line, []byte(":"))
		if endsWithBracketedNumber(line[:FirstDelimiter]) {
			LBrackets := bytes.Index(line[:FirstDelimiter], []byte("["))
			val, oK := res[string(line[:LBrackets])]
			if !oK {
				res[string(line[:LBrackets])] = make([][]string, 0)
			}
			res[string(line[:LBrackets])] = append(val, splitBySeparator(":", string(line[FirstDelimiter+1:])))
		} else {
			res[string(line[:FirstDelimiter])] = make([][]string, 0)
			res[string(line[:FirstDelimiter])] = [][]string{splitBySeparator(":", string(line[FirstDelimiter+1:]))}
		}
	}
	return res
}

func splitBySeparator(separator, line string) []string {
	escape := `\`
	tempEscapedSeparator := "\x00"

	replacedEscape := strings.ReplaceAll(line, escape+separator, tempEscapedSeparator)
	records := strings.Split(replacedEscape, separator)

	for i, record := range records {
		records[i] = strings.ReplaceAll(record, tempEscapedSeparator, separator)
	}
	return records
}
