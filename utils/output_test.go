package utils

import "testing"

func TestEndsWithBracketedNumber(t *testing.T) {
	tests := []struct {
		input    []byte
		expected bool
	}{
		// 正确匹配的测试用例
		{[]byte("example[123]"), true},
		{[]byte("example[0]"), true},
		{[]byte("[1]"), true},
		{[]byte("data[456789]"), true},

		// 错误匹配的测试用例
		{[]byte("example[]"), false},     // 空括号
		{[]byte("example123]"), false},   // 缺少 '['
		{[]byte("[123]example"), false},  // 括号不在末尾
		{[]byte("example[abc]"), false},  // 括号内非数字
		{[]byte("example[123"), false},   // 缺少 ']'
		{[]byte("example[1a3]"), false},  // 括号内有非数字字符
		{[]byte(""), false},              // 空输入
		{[]byte("[]"), false},            // 空括号
		{[]byte("example[ 123]"), false}, // 括号内有空格

		// 边界测试
		{[]byte("[0]"), true},          // 最短的正确匹配
		{[]byte("x[999999999]"), true}, // 边界的数字
	}

	for _, tt := range tests {
		t.Run(string(tt.input), func(t *testing.T) {
			result := endsWithBracketedNumber(tt.input)
			if result != tt.expected {
				t.Errorf("For input '%s', expected %v but got %v", tt.input, tt.expected, result)
			}
		})
	}
}
