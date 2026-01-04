// 包 chart_legend 提供图表图例组件的实现
// 该组件用于在AdminLTE主题中显示图表的图例信息
// 支持自定义图例数据，包括颜色和标签
package chart_legend

import (
	"html/template"

	adminTemplate "github.com/purpose168/GoAdmin/template"
)

// ChartLegend 图表图例组件结构体
// 继承自 BaseComponent，用于在页面中渲染图表图例
//
// 字段说明：
//   - BaseComponent: 基础组件，提供组件的基本功能
//   - Data: 图例数据数组，每个元素包含 "color" 和 "label" 两个字段
//
// 使用示例：
//
//	legend := chart_legend.New().
//	    SetData([]map[string]string{
//	        {"color": "danger", "label": "销售额"},
//	        {"color": "success", "label": "利润"},
//	    })
type ChartLegend struct {
	*adminTemplate.BaseComponent

	Data []map[string]string
}

// New 创建一个新的图表图例组件实例
//
// 返回值：
//   - ChartLegend: 初始化后的图表图例组件，包含默认的模板名称和HTML内容
//
// 使用示例：
//
//	legend := chart_legend.New()
func New() ChartLegend {
	return ChartLegend{
		BaseComponent: &adminTemplate.BaseComponent{
			Name:     "chart-legend",
			HTMLData: List["chart-legend"],
		},
	}
}

// SetData 设置图表图例的数据
//
// 参数：
//   - value: 图例数据数组，每个元素必须包含 "color" 和 "label" 字段
//   - "color": 颜色标识（如 "danger"、"success"、"info" 等）
//   - "label": 图例标签文本
//
// 返回值：
//   - ChartLegend: 返回设置数据后的组件实例，支持链式调用
//
// 使用示例：
//
//	legend := chart_legend.New().
//	    SetData([]map[string]string{
//	        {"color": "danger", "label": "销售额"},
//	        {"color": "success", "label": "利润"},
//	        {"color": "info", "label": "成本"},
//	    })
func (c ChartLegend) SetData(value []map[string]string) ChartLegend {
	c.Data = value
	return c
}

// GetContent 获取图表图例组件的HTML内容
// 该方法将组件数据与模板结合，生成最终的HTML代码
//
// 返回值：
//   - template.HTML: 渲染后的HTML内容
//
// 使用示例：
//
//	legend := chart_legend.New().
//	    SetData([]map[string]string{
//	        {"color": "danger", "label": "销售额"},
//	    })
//	htmlContent := legend.GetContent()
func (c ChartLegend) GetContent() template.HTML { return c.GetContentWithData(c) }
