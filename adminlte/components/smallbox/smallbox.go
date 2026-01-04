// 包 smallbox 提供小框组件的实现
// 该组件用于在AdminLTE主题中显示带有图标、标题、数值和链接的小框
// 支持自定义标题、数值、URL、颜色和图标，并支持十六进制颜色值和SVG图标
package smallbox

import (
	"html/template"
	"strings"

	adminTemplate "github.com/purpose168/GoAdmin/template"
)

// SmallBox 小框组件结构体
// 继承自 BaseComponent，用于在页面中渲染带有图标、标题、数值和链接的展示块
//
// 字段说明：
//   - BaseComponent: 基础组件，提供组件的基本功能
//   - Title: 标题文本，显示在数值下方，使用 template.HTML 类型以支持HTML内容
//   - Value: 主要数值，显示在小框顶部，使用 template.HTML 类型以支持HTML内容
//   - Url: 链接地址，点击小框底部区域时跳转的URL
//   - Color: 背景颜色标识，可以是预定义的颜色类（如 "aqua"、"green"）或十六进制颜色值
//   - IsSvg: 是否使用SVG图标（布尔值），由 SetIcon 方法自动设置
//   - IsHexColor: 是否使用十六进制颜色值（布尔值），由 SetColor 方法自动设置
//   - Icon: 图标标识，可以是Font Awesome图标类（如 "fa-user"）或SVG HTML代码
//
// 使用示例：
//
//	box := smallbox.New().
//	    SetTitle("新消息").
//	    SetValue("150").
//	    SetUrl("/messages").
//	    SetColor("aqua").
//	    SetIcon("fa-envelope-o")
type SmallBox struct {
	*adminTemplate.BaseComponent

	Title      template.HTML
	Value      template.HTML
	Url        string
	Color      template.HTML
	IsSvg      bool
	IsHexColor bool
	Icon       template.HTML
}

// New 创建一个新的小框组件实例
//
// 返回值：
//   - SmallBox: 初始化后的小框组件，包含默认的模板名称和HTML内容
//
// 使用示例：
//
//	box := smallbox.New()
func New() SmallBox {
	return SmallBox{
		BaseComponent: &adminTemplate.BaseComponent{
			Name:     "smallbox",
			HTMLData: List["smallbox"],
		},
	}
}

// SetTitle 设置小框的标题
//
// 参数：
//   - value: 标题文本，支持HTML内容
//
// 返回值：
//   - SmallBox: 返回设置标题后的组件实例，支持链式调用
//
// 使用示例：
//
//	box := smallbox.New().SetTitle("新消息")
func (s SmallBox) SetTitle(value template.HTML) SmallBox {
	s.Title = value
	return s
}

// SetValue 设置小框的主要数值
//
// 参数：
//   - value: 主要数值，支持HTML内容（如 "150" 或带有格式的HTML）
//
// 返回值：
//   - SmallBox: 返回设置数值后的组件实例，支持链式调用
//
// 使用示例：
//
//	box := smallbox.New().SetValue("150")
func (s SmallBox) SetValue(value template.HTML) SmallBox {
	s.Value = value
	return s
}

// SetColor 设置小框的背景颜色
//
// 参数：
//   - value: 颜色标识，可以是预定义的颜色类（如 "aqua"、"green"、"red" 等）或十六进制颜色值
//     如果包含 "#" 字符，会自动将 IsHexColor 设置为 true
//
// 返回值：
//   - SmallBox: 返回设置颜色后的组件实例，支持链式调用
//
// 使用示例：
//
//	box := smallbox.New().SetColor("aqua")
//	box := smallbox.New().SetColor(template.HTML("#3c8dbc"))
func (s SmallBox) SetColor(value template.HTML) SmallBox {
	s.Color = value
	if strings.Contains(string(value), "#") {
		s.IsHexColor = true
	}
	return s
}

// SetIcon 设置小框的图标
//
// 参数：
//   - value: 图标标识，可以是Font Awesome图标类（如 "fa-user"、"fa-envelope-o"）或SVG HTML代码
//     如果包含 "svg" 字符串，会自动将 IsSvg 设置为 true
//
// 返回值：
//   - SmallBox: 返回设置图标后的组件实例，支持链式调用
//
// 使用示例：
//
//	box := smallbox.New().SetIcon("fa-envelope-o")
//	box := smallbox.New().SetIcon(template.HTML("<svg>...</svg>"))
func (s SmallBox) SetIcon(value template.HTML) SmallBox {
	s.Icon = value
	if strings.Contains(string(value), "svg") {
		s.IsSvg = true
	}
	return s
}

// SetUrl 设置小框的链接地址
//
// 参数：
//   - value: 链接地址，点击小框底部区域时跳转的URL
//
// 返回值：
//   - SmallBox: 返回设置URL后的组件实例，支持链式调用
//
// 使用示例：
//
//	box := smallbox.New().SetUrl("/messages")
func (s SmallBox) SetUrl(value string) SmallBox {
	s.Url = value
	return s
}

// GetContent 获取小框组件的HTML内容
// 该方法将组件数据与模板结合，生成最终的HTML代码
//
// 返回值：
//   - template.HTML: 渲染后的HTML内容
//
// 使用示例：
//
//	box := smallbox.New().
//	    SetTitle("新消息").
//	    SetValue("150").
//	    SetUrl("/messages").
//	    SetColor("aqua").
//	    SetIcon("fa-envelope-o")
//	htmlContent := box.GetContent()
func (s SmallBox) GetContent() template.HTML { return s.GetContentWithData(s) }
