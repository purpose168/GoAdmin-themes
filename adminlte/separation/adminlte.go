// 包 separation 提供AdminLTE主题的分离版本实现
// 该包实现了AdminLTE主题的分离模式，将模板和资源进行分离管理
// 分离模式允许更灵活的主题配置和资源管理
package separation

import (
	"os"

	"github.com/purpose168/GoAdmin-themes/adminlte/resource"
	"github.com/purpose168/GoAdmin-themes/common"
	"github.com/purpose168/GoAdmin/modules/config"
	adminTemplate "github.com/purpose168/GoAdmin/template"
	"github.com/purpose168/GoAdmin/template/components"
	"github.com/purpose168/GoAdmin/template/types"
)

// Theme 主题结构体
// 定义了AdminLTE分离版本的主题结构
//
// 字段说明：
//   - ThemeName: 主题名称，用于标识当前主题
//   - Base: 基础组件，提供组件的基本功能
//   - BaseTheme: 基础主题，提供主题的通用功能
//
// 使用示例：
//
//	theme := separation.Get()
//	fmt.Println(theme.Name())
type Theme struct {
	ThemeName string
	components.Base
	*common.BaseTheme
}

// Adminlte AdminLTE分离版本的主题实例
// 该实例在包初始化时自动注册到模板系统中
//
// 字段说明：
//   - ThemeName: 主题名称，固定为 "adminlte_sep"
//   - Base: 基础组件配置，包含模板列表和分离模式标志
//   - BaseTheme: 基础主题配置，包含资源路径、模板列表和分离模式标志
//
// 注意事项：
//   - 该主题使用分离模式（Separation: true），将模板和资源分离管理
//   - 模板列表使用 common.SepTemplateList
//   - 资源路径使用 resource.AssetPaths
var Adminlte = Theme{
	ThemeName: "adminlte_sep",
	Base: components.Base{
		Attribute: types.Attribute{
			TemplateList: common.SepTemplateList,
			Separation:   true,
		},
	},
	BaseTheme: &common.BaseTheme{
		AssetPaths:   resource.AssetPaths,
		TemplateList: common.SepTemplateList,
		Separation:   true,
	},
}

// init 包初始化函数
// 在包被导入时自动执行，将AdminLTE分离版本主题注册到模板系统中
//
// 注意事项：
//   - 该函数在包导入时自动执行，无需手动调用
//   - 注册的主题名称为 "adminlte_sep"
//   - 注册后可以通过 adminTemplate.Get("adminlte_sep") 获取主题实例
func init() {
	adminTemplate.Add("adminlte_sep", &Adminlte)
}

// Get 获取AdminLTE分离版本的主题实例
//
// 返回值：
//   - *Theme: AdminLTE分离版本的主题实例指针
//
// 使用示例：
//
//	theme := separation.Get()
//	htmlContent := theme.GetContent()
func Get() *Theme {
	return &Adminlte
}

// Name 获取主题名称
//
// 返回值：
//   - string: 主题名称，固定为 "adminlte_sep"
//
// 使用示例：
//
//	theme := separation.Get()
//	name := theme.Name()
//	fmt.Println(name) // 输出: adminlte_sep
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
//	theme := separation.Get()
//	templates := theme.GetTmplList()
//	for name, content := range templates {
//	    fmt.Printf("模板: %s\n", name)
//	}
func (t *Theme) GetTmplList() map[string]string {
	return common.SepTemplateList
}

// GetAsset 获取指定路径的资源文件内容
// 从文件系统中读取资源文件并返回其内容
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
//	theme := separation.Get()
//	content, err := theme.GetAsset("/static/css/style.css")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(string(content))
//
// 注意事项：
//   - 资源文件路径是相对于配置的资源根目录的
//   - 资源根目录通过 config.GetAssetRootPath() 获取
func (t *Theme) GetAsset(path string) ([]byte, error) {
	return os.ReadFile(config.GetAssetRootPath() + path)
}

// GetAssetList 获取主题的所有资源文件列表
// 返回该主题包含的所有资源文件的路径列表
//
// 返回值：
//   - []string: 资源文件路径列表
//
// 使用示例：
//
//	theme := separation.Get()
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
