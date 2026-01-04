// 包 smallbox 提供小框组件的HTML模板
// 该组件用于在AdminLTE主题中显示带有图标、标题、数值和链接的小框
// 常用于仪表盘、统计卡片等场景，展示关键指标和快捷入口
package smallbox

// List 定义了小框组件的模板集合
// 键为模板标识符，值为对应的HTML模板字符串
//
// 模板说明：
//   - "smallbox": 小框模板，用于渲染带有图标、标题、数值和链接的展示块
//
// 模板变量：
//   - .Color: 背景颜色标识（如 "aqua"、"green"、"red"、"yellow" 等）
//   - .Value: 主要数值，显示在小框顶部
//   - .Title: 标题文本，显示在数值下方
//   - .Icon: 图标标识，使用Font Awesome图标类（如 "fa-envelope-o"、"fa-user" 等）
//   - .Url: 链接地址，点击小框底部区域时跳转的URL
//
// 注意事项：
//   - langHtml 函数用于支持多语言，会根据当前语言环境翻译文本
//   - lang 函数用于翻译预定义的文本键（如 "more" 表示"更多信息"）
//   - .Color 使用预定义的颜色类（如 "aqua"、"green"、"red"、"yellow" 等）
//   - .Icon 应为Font Awesome图标类，格式为 "fa-xxx"
//   - 小框底部包含"更多信息"链接，点击后跳转到 .Url 指定的地址
//
// 使用示例：
//
//	template := smallbox.List["smallbox"]
//	data := map[string]interface{}{
//	    "Color": "aqua",
//	    "Value": "150",
//	    "Title": "新消息",
//	    "Icon":  "fa-envelope-o",
//	    "Url":   "/messages",
//	}
//	tmpl.Execute(w, data)
var List = map[string]string{
	"smallbox": `{{define "smallbox"}}
    <div class="small-box bg-{{.Color}}">
        <div class="inner">
            <h3>{{langHtml .Value}}</h3>
            <p>{{langHtml .Title}}</p>
        </div>
        <div class="icon">
            <i class="fa {{.Icon}}"></i>
        </div>
        <a href="{{.Url}}" class="small-box-footer">
            {{lang "more"}}
            <i class="fa fa-arrow-circle-right"></i>
        </a>
    </div>
{{end}}`,
}
