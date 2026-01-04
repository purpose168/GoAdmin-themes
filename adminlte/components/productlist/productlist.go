// 包 productlist 提供产品列表组件的实现
// 该组件用于在AdminLTE主题中显示产品列表
// 支持自定义产品数据，包括图片、标题、描述和标签
package productlist

import (
	"html/template"

	adminTemplate "github.com/purpose168/GoAdmin/template"
)

// ProductList 产品列表组件结构体
// 继承自 BaseComponent，用于在页面中渲染产品列表
//
// 字段说明：
//   - BaseComponent: 基础组件，提供组件的基本功能
//   - Data: 产品数据数组，每个元素包含以下字段：
//   - "img": 产品图片URL
//   - "title": 产品标题
//   - "has_tabel": 是否显示标签（字符串类型的布尔值，"true" 或 "false"）
//   - "labeltype": 标签类型（如 "success"、"danger"、"warning" 等）
//   - "label": 标签文本
//   - "description": 产品描述
//
// 使用示例：
//
//	list := productlist.New().
//	    SetData([]map[string]string{
//	        {
//	            "img":         "/static/img/product1.jpg",
//	            "title":       "高级笔记本电脑",
//	            "has_tabel":   "true",
//	            "labeltype":   "success",
//	            "label":       "热销",
//	            "description": "高性能处理器，16GB内存，512GB固态硬盘",
//	        },
//	        {
//	            "img":         "/static/img/product2.jpg",
//	            "title":       "无线鼠标",
//	            "has_tabel":   "false",
//	            "description": "人体工学设计，长续航",
//	        },
//	    })
type ProductList struct {
	*adminTemplate.BaseComponent

	Data []map[string]string
}

// New 创建一个新的产品列表组件实例
//
// 返回值：
//   - ProductList: 初始化后的产品列表组件，包含默认的模板名称和HTML内容
//
// 使用示例：
//
//	list := productlist.New()
func New() ProductList {
	return ProductList{
		BaseComponent: &adminTemplate.BaseComponent{
			Name:     "productlist",
			HTMLData: List["productlist"],
		},
	}
}

// SetData 设置产品列表的数据
//
// 参数：
//   - value: 产品数据数组，每个元素必须包含以下字段：
//   - "img": 产品图片URL
//   - "title": 产品标题
//   - "has_tabel": 是否显示标签（字符串类型的布尔值，"true" 或 "false"）
//   - "labeltype": 标签类型（如 "success"、"danger"、"warning" 等）
//   - "label": 标签文本
//   - "description": 产品描述
//
// 返回值：
//   - ProductList: 返回设置数据后的组件实例，支持链式调用
//
// 使用示例：
//
//	list := productlist.New().
//	    SetData([]map[string]string{
//	        {
//	            "img":         "/static/img/product1.jpg",
//	            "title":       "高级笔记本电脑",
//	            "has_tabel":   "true",
//	            "labeltype":   "success",
//	            "label":       "热销",
//	            "description": "高性能处理器，16GB内存，512GB固态硬盘",
//	        },
//	    })
func (p ProductList) SetData(value []map[string]string) ProductList {
	p.Data = value
	return p
}

// GetContent 获取产品列表组件的HTML内容
// 该方法将组件数据与模板结合，生成最终的HTML代码
//
// 返回值：
//   - template.HTML: 渲染后的HTML内容
//
// 使用示例：
//
//	list := productlist.New().
//	    SetData([]map[string]string{
//	        {
//	            "img":         "/static/img/product1.jpg",
//	            "title":       "高级笔记本电脑",
//	            "has_tabel":   "true",
//	            "labeltype":   "success",
//	            "label":       "热销",
//	            "description": "高性能处理器，16GB内存，512GB固态硬盘",
//	        },
//	    })
//	htmlContent := list.GetContent()
func (p ProductList) GetContent() template.HTML { return p.GetContentWithData(p) }
