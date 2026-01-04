// 包 progress_group 提供进度条组组件的HTML模板
// 该组件用于在AdminLTE主题中显示带有标题、数值和进度条的进度条组
// 常用于仪表盘、统计卡片等场景，展示任务完成进度、数据占比等信息
package progress_group

// List 定义了进度条组组件的模板集合
// 键为模板标识符，值为对应的HTML模板字符串
//
// 模板说明：
//   - "progress-group": 进度条组模板，用于渲染带有标题、数值和进度条的展示块
//
// 模板变量：
//   - .Title: 进度条标题文本
//   - .Molecular: 分子（当前值），显示在进度条左侧
//   - .Denominator: 分母（总值），显示在进度条右侧
//   - .Percent: 进度百分比（0-100）
//   - .Color: 颜色标识，可以是预定义的颜色类（如 "success"、"danger"、"warning" 等）或十六进制颜色值
//   - .IsHexColor: 是否使用十六进制颜色值（布尔值）
//
// 注意事项：
//   - langHtml 函数用于支持多语言，会根据当前语言环境翻译文本
//   - 当 IsHexColor 为 true 时，Color 应为十六进制颜色值（如 "#3c8dbc"）
//   - 当 IsHexColor 为 false 时，Color 应为预定义的颜色类（如 "success"、"danger"、"warning" 等）
//   - Percent 值应为 0-100 之间的数字
//   - 进度条使用 "sm"（小尺寸）样式
//
// 使用示例：
//
//	template := progress_group.List["progress-group"]
//	data := map[string]interface{}{
//	    "Title":       "任务完成度",
//	    "Molecular":   "75",
//	    "Denominator": "100",
//	    "Percent":     "75",
//	    "Color":       "success",
//	    "IsHexColor":  false,
//	}
//	tmpl.Execute(w, data)
var List = map[string]string{
	"progress-group": `{{define "progress-group"}}
    <div class="progress-group">
        <span class="progress-text">{{langHtml .Title}}</span>
        <span class="progress-number"><b>{{.Molecular}}</b>/{{.Denominator}}</span>

        <div class="progress sm">
            {{if .IsHexColor}}
                <div class="progress-bar" style="width: {{.Percent}}%;background-color: {{.Color}} !important;"></div>
            {{else}}
                <div class="progress-bar progress-bar-{{.Color}}" style="width: {{.Percent}}%"></div>
            {{end}}
        </div>
    </div>
{{end}}`,
}
