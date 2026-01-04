// 包 productlist 提供产品列表组件的HTML模板
// 该组件用于在AdminLTE主题中显示产品列表
// 常用于仪表盘、产品管理等场景，展示带有图片、标题、描述和标签的产品信息
package productlist

// List 定义了产品列表组件的模板集合
// 键为模板标识符，值为对应的HTML模板字符串
//
// 模板说明：
//   - "productlist": 产品列表模板，用于渲染带有图片、标题、描述和标签的产品列表
//
// 模板变量：
//   - .Data: 产品数据数组，每个元素包含以下字段：
//   - "img": 产品图片URL
//   - "title": 产品标题
//   - "has_tabel": 是否显示标签（字符串类型的布尔值，"true" 或 "false"）
//   - "labeltype": 标签类型（如 "success"、"danger"、"warning" 等）
//   - "label": 标签文本
//   - "description": 产品描述
//
// 注意事项：
//   - "has_tabel" 字段是字符串类型的布尔值，值为 "true" 时显示标签
//   - "labeltype" 使用Bootstrap的标签颜色类（如 "success"、"danger"、"warning" 等）
//   - 产品图片使用 "Product Image" 作为alt属性，可根据需要自定义
//
// 使用示例：
//
//	template := productlist.List["productlist"]
//	data := []map[string]interface{}{
//	    {
//	        "img":         "/static/img/product1.jpg",
//	        "title":       "高级笔记本电脑",
//	        "has_tabel":   "true",
//	        "labeltype":   "success",
//	        "label":       "热销",
//	        "description": "高性能处理器，16GB内存，512GB固态硬盘",
//	    },
//	    {
//	        "img":         "/static/img/product2.jpg",
//	        "title":       "无线鼠标",
//	        "has_tabel":   "false",
//	        "description": "人体工学设计，长续航",
//	    },
//	}
//	tmpl.Execute(w, map[string]interface{}{"Data": data})
var List = map[string]string{
	"productlist": `{{define "productlist"}}
<ul class="products-list product-list-in-box">
    {{range $key, $data := .Data}}
    <li class="item">
        <div class="product-img">
            <img src="{{index $data "img"}}" alt="Product Image">
        </div>
        <div class="product-info">
            <a href="javascript:void(0)" class="product-title">{{index $data "title"}}
                {{if eq (index $data "has_tabel") "true"}}
                    <span class="label label-{{index $data "labeltype"}} pull-right">{{index $data "label"}}</span>
                {{end}}
            </a>
            <span class="product-description">
                {{index $data "description"}}
            </span>
        </div>
    </li>
    {{end}}
</ul>
{{end}}`,
}
