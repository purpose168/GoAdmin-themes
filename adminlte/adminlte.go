// 包 adminlte 提供AdminLTE主题的标准版本实现
// 该包实现了AdminLTE主题的标准模式，将模板和资源打包在二进制文件中
// 标准模式适合生产环境部署，提供完整的主题功能
package adminlte

import (
	"strings"

	"github.com/purpose168/GoAdmin-themes/adminlte/resource"
	"github.com/purpose168/GoAdmin-themes/common"
	adminTemplate "github.com/purpose168/GoAdmin/template"
	"github.com/purpose168/GoAdmin/template/components"
	"github.com/purpose168/GoAdmin/template/types"
)

// 颜色方案常量定义
// 这些常量用于设置AdminLTE主题的皮肤颜色方案
const (
	ColorschemeSkinBlack       = "skin-black"        // 黑色皮肤
	ColorschemeSkinBlackLight  = "skin-black-light"  // 浅黑色皮肤
	ColorschemeSkinBlue        = "skin-blue"         // 蓝色皮肤
	ColorschemeSkinBlueLight   = "skin-blue-light"   // 浅蓝色皮肤
	ColorschemeSkinGreen       = "skin-green"        // 绿色皮肤
	ColorschemeSkinGreenLight  = "skin-green-light"  // 浅绿色皮肤
	ColorschemeSkinPurple      = "skin-purple"       // 紫色皮肤
	ColorschemeSkinPurpleLight = "skin-purple-light" // 浅紫色皮肤
	ColorschemeSkinRed         = "skin-red"          // 红色皮肤
	ColorschemeSkinRedLight    = "skin-red-light"    // 浅红色皮肤
	ColorschemeSkinYellow      = "skin-yellow"       // 黄色皮肤
	ColorschemeSkinYellowLight = "skin-yellow-light" // 浅黄色皮肤
)

// Theme 主题结构体
// 定义了AdminLTE标准版本的主题结构
//
// 字段说明：
//   - ThemeName: 主题名称，用于标识当前主题
//   - Base: 基础组件，提供组件的基本功能
//   - BaseTheme: 基础主题，提供主题的通用功能
//
// 使用示例：
//
//	theme := adminlte.Get()
//	fmt.Println(theme.Name())
type Theme struct {
	ThemeName string
	components.Base
	*common.BaseTheme
}

// Adminlte AdminLTE标准版本的主题实例
// 该实例在包初始化时自动注册到模板系统中
//
// 字段说明：
//   - ThemeName: 主题名称，固定为 "adminlte"
//   - Base: 基础组件配置，包含模板列表
//   - BaseTheme: 基础主题配置，包含资源路径和模板列表
//
// 注意事项：
//   - 该主题使用标准模式，模板和资源打包在二进制文件中
//   - 模板列表使用 TemplateList
//   - 资源路径使用 resource.AssetPaths
//   - 资源通过 resource.AssetFS 访问（嵌入的文件系统）
var Adminlte = Theme{
	ThemeName: "adminlte",
	Base: components.Base{
		Attribute: types.Attribute{
			TemplateList: TemplateList,
		},
	},
	BaseTheme: &common.BaseTheme{
		AssetPaths:   resource.AssetPaths,
		TemplateList: TemplateList,
	},
}

// init 包初始化函数
// 在包被导入时自动执行，将AdminLTE标准版本主题注册到模板系统中
//
// 注意事项：
//   - 该函数在包导入时自动执行，无需手动调用
//   - 注册的主题名称为 "adminlte"
//   - 注册后可以通过 adminTemplate.Get("adminlte") 获取主题实例
func init() {
	adminTemplate.Add("adminlte", &Adminlte)
}

// Get 获取AdminLTE标准版本的主题实例
//
// 返回值：
//   - *Theme: AdminLTE标准版本的主题实例指针
//
// 使用示例：
//
//	theme := adminlte.Get()
//	htmlContent := theme.GetContent()
func Get() *Theme {
	return &Adminlte
}

// Name 获取主题名称
//
// 返回值：
//   - string: 主题名称，固定为 "adminlte"
//
// 使用示例：
//
//	theme := adminlte.Get()
//	name := theme.Name()
//	fmt.Println(name) // 输出: adminlte
func (t *Theme) Name() string {
	return t.ThemeName
}

// GetTmplList 获取主题的模板列表
// 返回该主题支持的所有模板及其内容
//
// 返回值：
//   - map[string]string: 模板列表，键为模板名称，值为模板内容
//
// 使用示例：
//
//	theme := adminlte.Get()
//	templates := theme.GetTmplList()
//	for name, content := range templates {
//	    fmt.Printf("模板: %s\n", name)
//	}
func (t *Theme) GetTmplList() map[string]string {
	return TemplateList
}

// GetAssetList 获取主题的所有资源文件列表
// 返回该主题包含的所有资源文件的路径列表
//
// 返回值：
//   - []string: 资源文件路径列表
//
// 使用示例：
//
//	theme := adminlte.Get()
//	assets := theme.GetAssetList()
//	for _, asset := range assets {
//	    fmt.Println(asset)
//	}
//
// 注意事项：
//   - 返回的资源文件列表是预定义的，来自 resource.AssetsList
//   - 资源文件包括CSS、JavaScript、图片等静态资源
func (t *Theme) GetAssetList() []string {
	return resource.AssetsList
}

// GetAsset 获取指定路径的资源文件内容
// 从嵌入的文件系统中读取资源文件并返回其内容
//
// 参数：
//   - path: 资源文件的相对路径
//
// 返回值：
//   - []byte: 资源文件的二进制内容
//   - error: 读取错误，如果读取成功则为 nil
//
// 使用示例：
//
//	theme := adminlte.Get()
//	content, err := theme.GetAsset("/assets/dist/css/style.css")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(string(content))
//
// 注意事项：
//   - 资源文件路径会自动处理，将 "/assets/dist" 替换为 "assets/dist"
//   - 资源从嵌入的文件系统 resource.AssetFS 中读取
//   - 标准模式下，资源文件打包在二进制文件中
func (t *Theme) GetAsset(path string) ([]byte, error) {
	path = strings.Replace(path, "/assets/dist", "assets/dist", -1)
	return resource.AssetFS.ReadFile(path)
}
