// 包 infobox 提供信息框组件的实现
// 该组件用于在AdminLTE主题中显示带有图标、文本和数值的信息框
// 支持自定义图标、文本、数值、内容、颜色，并支持十六进制颜色值和SVG图标
package infobox

import (
	"html/template"
	"strings"

	adminTemplate "github.com/purpose168/GoAdmin/template"
)

// InfoBox 信息框组件结构体
// 继承自 BaseComponent，用于在页面中渲染带有图标和数值的信息展示块
//
// 字段说明：
//   - BaseComponent: 基础组件，提供组件的基本功能
//   - Icon: 图标标识，可以是Font Awesome图标类（如 "fa-user"）或SVG HTML代码
//   - Text: 信息框的文本描述，使用 template.HTML 类型以支持HTML内容
//   - Number: 信息框的主要数值，使用 template.HTML 类型以支持HTML内容
//   - Content: 信息框的额外内容，使用 template.HTML 类型以支持HTML内容
//   - Color: 颜色标识，可以是预定义的颜色类（如 "aqua"、"green"）或十六进制颜色值
//   - IsHexColor: 是否使用十六进制颜色值（布尔值），由 SetColor 方法自动设置
//   - IsSvg: 是否使用SVG图标（布尔值），由 SetIcon 方法自动设置
//
// 使用示例：
//
//	box := infobox.New().
//	    SetIcon("fa-envelope-o").
//	    SetText("消息").
//	    SetNumber("1,410").
//	    SetColor("aqua")
type InfoBox struct {
	*adminTemplate.BaseComponent

	Icon       template.HTML
	Text       template.HTML
	Number     template.HTML
	Content    template.HTML
	Color      template.HTML
	IsHexColor bool
	IsSvg      bool
}

// New 创建一个新的信息框组件实例
//
// 返回值：
//   - InfoBox: 初始化后的信息框组件，包含默认的模板名称和HTML内容
//
// 使用示例：
//
//	box := infobox.New()
func New() InfoBox {
	return InfoBox{
		BaseComponent: &adminTemplate.BaseComponent{
			Name:     "infobox",
			HTMLData: List["infobox"],
		},
	}
}

// SetIcon 设置信息框的图标
//
// 参数：
//   - value: 图标标识，可以是Font Awesome图标类（如 "fa-user"、"fa-shopping-cart"）或SVG HTML代码
//     如果包含 "svg" 字符串，会自动将 IsSvg 设置为 true
//
// 返回值：
//   - InfoBox: 返回设置图标后的组件实例，支持链式调用
//
// 使用示例：
//
//	box := infobox.New().SetIcon("fa-envelope-o")
//	box := infobox.New().SetIcon(template.HTML("<svg>...</svg>"))
func (i InfoBox) SetIcon(value template.HTML) InfoBox {
	i.Icon = value
	if strings.Contains(string(value), "svg") {
		i.IsSvg = true
	}
	return i
}

// SetText 设置信息框的文本描述
//
// 参数：
//   - value: 文本描述，支持HTML内容
//
// 返回值：
//   - InfoBox: 返回设置文本后的组件实例，支持链式调用
//
// 使用示例：
//
//	box := infobox.New().SetText("消息")
func (i InfoBox) SetText(value template.HTML) InfoBox {
	i.Text = value
	return i
}

// SetNumber 设置信息框的主要数值
//
// 参数：
//   - value: 主要数值，支持HTML内容（如 "1,410" 或带有格式的HTML）
//
// 返回值：
//   - InfoBox: 返回设置数值后的组件实例，支持链式调用
//
// 使用示例：
//
//	box := infobox.New().SetNumber("1,410")
func (i InfoBox) SetNumber(value template.HTML) InfoBox {
	i.Number = value
	return i
}

// SetContent 设置信息框的额外内容
//
// 参数：
//   - value: 额外内容，支持HTML内容（如进度条、描述文字等）
//
// 返回值：
//   - InfoBox: 返回设置内容后的组件实例，支持链式调用
//
// 使用示例：
//
//	box := infobox.New().SetContent(template.HTML("<div class='progress'>...</div>"))
func (i InfoBox) SetContent(value template.HTML) InfoBox {
	i.Content = value
	return i
}

// SetColor 设置信息框的颜色
//
// 参数：
//   - value: 颜色标识，可以是预定义的颜色类（如 "aqua"、"green"、"red" 等）或十六进制颜色值
//     如果包含 "#" 字符，会自动将 IsHexColor 设置为 true
//
// 返回值：
//   - InfoBox: 返回设置颜色后的组件实例，支持链式调用
//
// 使用示例：
//
//	box := infobox.New().SetColor("aqua")
//	box := infobox.New().SetColor(template.HTML("#3c8dbc"))
func (i InfoBox) SetColor(value template.HTML) InfoBox {
	i.Color = value
	if strings.Contains(string(value), "#") {
		i.IsHexColor = true
	}
	return i
}

// GetContent 获取信息框组件的HTML内容
// 该方法将组件数据与模板结合，生成最终的HTML代码
//
// 返回值：
//   - template.HTML: 渲染后的HTML内容
//
// 使用示例：
//
//	box := infobox.New().
//	    SetIcon("fa-envelope-o").
//	    SetText("消息").
//	    SetNumber("1,410").
//	    SetColor("aqua")
//	htmlContent := box.GetContent()
func (i InfoBox) GetContent() template.HTML { return i.GetContentWithData(i) }
