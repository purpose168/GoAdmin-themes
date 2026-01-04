// 包 chart_legend 提供图表图例组件的HTML模板
// 该组件用于在AdminLTE主题中显示图表的图例信息
package chart_legend

// List 定义了图表图例组件的模板集合
// 键为模板标识符，值为对应的HTML模板字符串
//
// 模板说明：
//   - "chart-legend": 图表图例模板，用于渲染图例列表
//
// 模板变量：
//   - .Data: 图例数据数组，每个元素包含以下字段：
//   - "color": 颜色标识（如 "danger"、"success" 等）
//   - "label": 图例标签文本
//
// 使用示例：
//
//	template := chart_legend.List["chart-legend"]
//	data := []map[string]interface{}{
//	    {"color": "danger", "label": "销售额"},
//	    {"color": "success", "label": "利润"},
//	}
//	tmpl.Execute(w, map[string]interface{}{"Data": data})
var List = map[string]string{
	"chart-legend": `{{define "chart-legend"}}
<ul class="chart-legend clearfix">
    {{range $key, $data := .Data}}
        <li><i class="fa fa-circle-o text-{{index $data "color"}}"></i>{{index $data "label"}}</li>
    {{end}}
</ul>
{{end}}`,
}
