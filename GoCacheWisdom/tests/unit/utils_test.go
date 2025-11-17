package unit

import (
	"github.com/GoServePro/GoDevExperience/GoCacheWisdom/internal"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 测试CalculateSize计算结果是否稳定、符合预期
func TestCalculateSize(t *testing.T) {
	// 测试用例：类型->输入值->实际合理的预期范围
	testCases := []struct {
		name    string
		value   interface{}
		minSize int64 // 最小预期字节数
		maxSize int64 // 最大预期字节数
	}{
		{
			name:    "空字符串",
			value:   "",
			minSize: 4, // 实际编码后约4字节（精简）
			maxSize: 10,
		},
		{
			name:    "普通字符串（张三）",
			value:   "张三",
			minSize: 10, // 实际编码后约10-15字节
			maxSize: 15,
		},
		{
			name:    "int类型（100）",
			value:   100,
			minSize: 5, // 实际编码后约5-8字节
			maxSize: 8,
		},
		{
			name: "结构体（User）",
			value: struct {
				ID   int
				Name string
			}{ID: 1001, Name: "李四"},
			minSize: 25,
			maxSize: 50,
		},
		{
			name:    "切片（[]int{1,2,3}）",
			value:   []int{1, 2, 3},
			minSize: 15, // 实际编码后约15-30字节
			maxSize: 30,
		},
		{
			name:    "nil指针（返回默认64字节）",
			value:   (*int)(nil),
			minSize: 64,
			maxSize: 64, // 严格等于默认值
		},
	}

	// 执行测试
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			size := internal.CalculateSize(tc.value)
			// 断言结果是否在预期范围内
			assert.GreaterOrEqual(t, size, tc.minSize, "计算结果偏小")
			assert.LessOrEqual(t, size, tc.maxSize, "计算结果偏大")
		})
	}
}
