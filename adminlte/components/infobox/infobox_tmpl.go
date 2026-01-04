// 包 infobox 提供信息框组件的HTML模板
// 该组件用于在AdminLTE主题中显示带有图标、文本和数值的信息框
// 常用于仪表盘、统计卡片等场景，展示关键指标和状态信息
package infobox

// List 定义了信息框组件的模板集合
// 键为模板标识符，值为对应的HTML模板字符串
//
// 模板说明：
//   - "infobox": 信息框模板，用于渲染带有图标和数值的信息展示块
//
// 模板变量：
//   - .IsHexColor: 是否使用十六进制颜色值（布尔值）
//   - .Color: 颜色标识，可以是预定义的颜色类（如 "aqua"、"green"、"red" 等）或十六进制颜色值
//   - .IsSvg: 是否使用SVG图标（布尔值）
//   - .Icon: 图标标识，可以是Font Awesome图标类（如 "fa-user"、"fa-shopping-cart"）或SVG HTML代码
//   - .Text: 信息框的文本描述
//   - .Number: 信息框的主要数值
//   - .Content: 信息框的额外内容，支持HTML
//
// 注意事项：
//   - langHtml 函数用于支持多语言，会根据当前语言环境翻译文本
//   - 当 IsHexColor 为 true 时，Color 应为十六进制颜色值（如 "#3c8dbc"）
//   - 当 IsHexColor 为 false 时，Color 应为预定义的颜色类（如 "aqua"、"green"、"red" 等）
//   - 当 IsSvg 为 true 时，Icon 应为完整的SVG HTML代码
//   - 当 IsSvg 为 false 时，Icon 应为Font Awesome图标类（如 "fa-user"）
//
// 使用示例：
//
//	template := infobox.List["infobox"]
//	data := map[string]interface{}{
//	    "IsHexColor": false,
//	    "Color":      "aqua",
//	    "IsSvg":      false,
//	    "Icon":       "fa-envelope-o",
//	    "Text":       "消息",
//	    "Number":     "1,410",
//	    "Content":    template.HTML("<div>额外内容</div>"),
//	}
//	tmpl.Execute(w, data)
var List = map[string]string{
	"infobox": `{{define "infobox"}}
<div class="info-box">
    {{if .IsHexColor}}
        <span class="info-box-icon" style="background-color: {{.Color}} !important;">
    {{else}}
        <span class="info-box-icon bg-{{.Color}}">
    {{end}}
    {{if .IsSvg}}
        {{.Icon}}
    {{else}}
        <i class="fa {{.Icon}}"></i>
    {{end}}
    </span>
    <div class="info-box-content">
        <span class="info-box-text">{{langHtml .Text}}</span>
        <span class="info-box-number">{{langHtml .Number}}</span>
        {{langHtml .Content}}
    </div>
</div>
{{end}}`,
}
