// 包 description 提供描述块组件的实现
// 该组件用于在AdminLTE主题中显示带有百分比变化的数据描述块
// 支持自定义边框样式、数值、标题、箭头方向、颜色和百分比
package description

import (
	"html/template"

	adminTemplate "github.com/purpose168/GoAdmin/template"
)

// Description 描述块组件结构体
// 继承自 BaseComponent，用于在页面中渲染带有百分比变化的数据展示块
//
// 字段说明：
//   - BaseComponent: 基础组件，提供组件的基本功能
//   - Border: 边框样式（如 "right"、"left"、"bottom" 等）
//   - Number: 主要数值，使用 template.HTML 类型以支持HTML内容
//   - Title: 描述标题文本，使用 template.HTML 类型以支持HTML内容
//   - Arrow: 箭头方向（如 "up"、"down"）
//   - Color: 百分比颜色标识（如 "green"、"red"、"yellow" 等）
//   - Percent: 百分比数值，使用 template.HTML 类型以支持HTML内容
//
// 使用示例：
//
//	desc := description.New().
//	    SetBorder("right").
//	    SetNumber("1,234").
//	    SetTitle("新用户").
//	    SetArrow("up").
//	    SetColor("green").
//	    SetPercent("15")
type Description struct {
	*adminTemplate.BaseComponent

	Border  string
	Number  template.HTML
	Title   template.HTML
	Arrow   string
	Color   template.HTML
	Percent template.HTML
}

// New 创建一个新的描述块组件实例
//
// 返回值：
//   - Description: 初始化后的描述块组件，包含默认的模板名称和HTML内容
//
// 使用示例：
//
//	desc := description.New()
func New() Description {
	return Description{
		BaseComponent: &adminTemplate.BaseComponent{
			Name:     "description",
			HTMLData: List["description"],
		},
	}
}

// SetNumber 设置描述块的主要数值
//
// 参数：
//   - value: 主要数值，支持HTML内容（如 "1,234" 或带有格式的HTML）
//
// 返回值：
//   - Description: 返回设置数值后的组件实例，支持链式调用
//
// 使用示例：
//
//	desc := description.New().SetNumber("1,234")
func (c Description) SetNumber(value template.HTML) Description {
	c.Number = value
	return c
}

// SetTitle 设置描述块的标题文本
//
// 参数：
//   - value: 标题文本，支持HTML内容
//
// 返回值：
//   - Description: 返回设置标题后的组件实例，支持链式调用
//
// 使用示例：
//
//	desc := description.New().SetTitle("新用户")
func (c Description) SetTitle(value template.HTML) Description {
	c.Title = value
	return c
}

// SetArrow 设置描述块的箭头方向
//
// 参数：
//   - value: 箭头方向，可选值为 "up"（向上，表示增长）或 "down"（向下，表示下降）
//
// 返回值：
//   - Description: 返回设置箭头方向后的组件实例，支持链式调用
//
// 使用示例：
//
//	desc := description.New().SetArrow("up")
func (c Description) SetArrow(value string) Description {
	c.Arrow = value
	return c
}

// SetPercent 设置描述块的百分比数值
//
// 参数：
//   - value: 百分比数值，支持HTML内容（如 "15" 表示15%）
//
// 返回值：
//   - Description: 返回设置百分比后的组件实例，支持链式调用
//
// 使用示例：
//
//	desc := description.New().SetPercent("15")
func (c Description) SetPercent(value template.HTML) Description {
	c.Percent = value
	return c
}

// SetColor 设置描述块的百分比颜色
//
// 参数：
//   - value: 颜色标识，可选值为 "green"（绿色，表示增长）、"red"（红色，表示下降）、"yellow"（黄色）等
//
// 返回值：
//   - Description: 返回设置颜色后的组件实例，支持链式调用
//
// 使用示例：
//
//	desc := description.New().SetColor("green")
func (c Description) SetColor(value template.HTML) Description {
	c.Color = value
	return c
}

// SetBorder 设置描述块的边框样式
//
// 参数：
//   - value: 边框样式，可选值为 "right"（右边框）、"left"（左边框）、"bottom"（下边框）等
//
// 返回值：
//   - Description: 返回设置边框样式后的组件实例，支持链式调用
//
// 使用示例：
//
//	desc := description.New().SetBorder("right")
func (c Description) SetBorder(value string) Description {
	c.Border = value
	return c
}

// GetContent 获取描述块组件的HTML内容
// 该方法将组件数据与模板结合，生成最终的HTML代码
//
// 返回值：
//   - template.HTML: 渲染后的HTML内容
//
// 使用示例：
//
//	desc := description.New().
//	    SetBorder("right").
//	    SetNumber("1,234").
//	    SetTitle("新用户").
//	    SetArrow("up").
//	    SetColor("green").
//	    SetPercent("15")
//	htmlContent := desc.GetContent()
func (c Description) GetContent() template.HTML { return c.GetContentWithData(c) }
