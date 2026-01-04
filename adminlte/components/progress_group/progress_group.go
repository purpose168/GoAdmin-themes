// 包 progress_group 提供进度条组组件的实现
// 该组件用于在AdminLTE主题中显示带有标题、数值和进度条的进度条组
// 支持自定义标题、分子、分母、百分比和颜色，并支持十六进制颜色值
package progress_group

import (
	"html/template"
	"strings"

	adminTemplate "github.com/purpose168/GoAdmin/template"
)

// ProgressGroup 进度条组组件结构体
// 继承自 BaseComponent，用于在页面中渲染带有标题、数值和进度条的展示块
//
// 字段说明：
//   - BaseComponent: 基础组件，提供组件的基本功能
//   - Title: 进度条标题文本，使用 template.HTML 类型以支持HTML内容
//   - Molecular: 分子（当前值），显示在进度条左侧，使用 int 类型
//   - Denominator: 分母（总值），显示在进度条右侧，使用 int 类型
//   - Color: 颜色标识，可以是预定义的颜色类（如 "success"、"danger"）或十六进制颜色值
//   - IsHexColor: 是否使用十六进制颜色值（布尔值），由 SetColor 方法自动设置
//   - Percent: 进度百分比（0-100），使用 int 类型
//
// 使用示例：
//
//	group := progress_group.New().
//	    SetTitle("任务完成度").
//	    SetMolecular(75).
//	    SetDenominator(100).
//	    SetPercent(75).
//	    SetColor("success")
type ProgressGroup struct {
	*adminTemplate.BaseComponent

	Title       template.HTML
	Molecular   int
	Denominator int
	Color       template.HTML
	IsHexColor  bool
	Percent     int
}

// New 创建一个新的进度条组组件实例
//
// 返回值：
//   - ProgressGroup: 初始化后的进度条组组件，包含默认的模板名称和HTML内容
//
// 使用示例：
//
//	group := progress_group.New()
func New() ProgressGroup {
	return ProgressGroup{
		BaseComponent: &adminTemplate.BaseComponent{
			Name:     "progress-group",
			HTMLData: List["progress-group"],
		},
	}
}

// SetTitle 设置进度条组的标题
//
// 参数：
//   - value: 标题文本，支持HTML内容
//
// 返回值：
//   - ProgressGroup: 返回设置标题后的组件实例，支持链式调用
//
// 使用示例：
//
//	group := progress_group.New().SetTitle("任务完成度")
func (p ProgressGroup) SetTitle(value template.HTML) ProgressGroup {
	p.Title = value
	return p
}

// SetColor 设置进度条组的颜色
//
// 参数：
//   - value: 颜色标识，可以是预定义的颜色类（如 "success"、"danger"、"warning" 等）或十六进制颜色值
//     如果包含 "#" 字符，会自动将 IsHexColor 设置为 true
//
// 返回值：
//   - ProgressGroup: 返回设置颜色后的组件实例，支持链式调用
//
// 使用示例：
//
//	group := progress_group.New().SetColor("success")
//	group := progress_group.New().SetColor(template.HTML("#3c8dbc"))
func (p ProgressGroup) SetColor(value template.HTML) ProgressGroup {
	p.Color = value
	if strings.Contains(string(value), "#") {
		p.IsHexColor = true
	}
	return p
}

// SetPercent 设置进度条组的百分比
//
// 参数：
//   - value: 进度百分比，应为 0-100 之间的整数
//
// 返回值：
//   - ProgressGroup: 返回设置百分比后的组件实例，支持链式调用
//
// 使用示例：
//
//	group := progress_group.New().SetPercent(75)
func (p ProgressGroup) SetPercent(value int) ProgressGroup {
	p.Percent = value
	return p
}

// SetDenominator 设置进度条组的分母（总值）
//
// 参数：
//   - value: 分母值，显示在进度条右侧
//
// 返回值：
//   - ProgressGroup: 返回设置分母后的组件实例，支持链式调用
//
// 使用示例：
//
//	group := progress_group.New().SetDenominator(100)
func (p ProgressGroup) SetDenominator(value int) ProgressGroup {
	p.Denominator = value
	return p
}

// SetMolecular 设置进度条组的分子（当前值）
//
// 参数：
//   - value: 分子值，显示在进度条左侧
//
// 返回值：
//   - ProgressGroup: 返回设置分子后的组件实例，支持链式调用
//
// 使用示例：
//
//	group := progress_group.New().SetMolecular(75)
func (p ProgressGroup) SetMolecular(value int) ProgressGroup {
	p.Molecular = value
	return p
}

// GetContent 获取进度条组组件的HTML内容
// 该方法将组件数据与模板结合，生成最终的HTML代码
//
// 返回值：
//   - template.HTML: 渲染后的HTML内容
//
// 使用示例：
//
//	group := progress_group.New().
//	    SetTitle("任务完成度").
//	    SetMolecular(75).
//	    SetDenominator(100).
//	    SetPercent(75).
//	    SetColor("success")
//	htmlContent := group.GetContent()
func (p ProgressGroup) GetContent() template.HTML { return p.GetContentWithData(p) }
