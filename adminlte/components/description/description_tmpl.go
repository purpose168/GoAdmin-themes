// 包 description 提供描述块组件的HTML模板
// 该组件用于在AdminLTE主题中显示带有百分比变化的数据描述块
// 常用于仪表盘、统计卡片等场景，展示关键指标及其变化趋势
package description

// List 定义了描述块组件的模板集合
// 键为模板标识符，值为对应的HTML模板字符串
//
// 模板说明：
//   - "description": 描述块模板，用于渲染带有百分比变化的数据展示块
//
// 模板变量：
//   - .Border: 边框样式（如 "right"、"left"、"bottom" 等）
//   - .Color: 百分比颜色标识（如 "green"、"red"、"yellow" 等）
//   - .Arrow: 箭头方向（如 "up"、"down"）
//   - .Percent: 百分比数值
//   - .Number: 主要数值
//   - .Title: 描述标题文本
//
// 注意事项：
//   - langHtml 函数用于支持多语言，会根据当前语言环境翻译文本
//   - 百分比和箭头方向通常配合使用，如增长用 "up"+"green"，下降用 "down"+"red"
//
// 使用示例：
//
//	template := description.List["description"]
//	data := map[string]interface{}{
//	    "Border":  "right",
//	    "Color":   "green",
//	    "Arrow":   "up",
//	    "Percent": "15",
//	    "Number":  "1,234",
//	    "Title":   "新用户",
//	}
//	tmpl.Execute(w, data)
var List = map[string]string{
	"description": `{{define "description"}}
<div class="description-block border-{{.Border}}">
    <span class="description-percentage text-{{.Color}}"><i class="fa fa-caret-{{.Arrow}}"></i>{{langHtml .Percent}}%</span>
    <h5 class="description-header">{{langHtml .Number}}</h5>
    <span class="description-text">{{langHtml .Title}}</span>
</div>
{{end}}`,
}
