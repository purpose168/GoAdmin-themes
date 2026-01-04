# CLI命令工具名称，用于执行资源合并和打包等操作
CLI=adm
# 资源文件根目录路径
RESOURCE_PATH=./resource
# 资源文件中的assets目录路径，包含前端静态资源
ASSETS_PATH=./resource/assets
# 分离主题目录路径，用于生成独立的静态资源包
SEPARATION_PATH=./separation
# 公共资源目录路径，存放各主题共享的资源文件
COMMON_PATH=./../common

# 默认目标，执行完整的主题构建流程
all: build

# 构建主题 (build a theme)
# 执行完整的构建流程，包括：重建dist目录、合并资源、生成压缩包、打包资源、打包模板、清理临时文件、代码格式化
build: rebuild-dist combine build-separation-zip build-assets build-tmpl clean fmt

# 清理旧的dist文件夹，创建新的dist文件夹 (clean old dist folder, create new dist folder)
# 此目标负责准备构建环境，包括：
# 1. 删除旧的dist和src目录
# 2. 创建新的dist目录结构（包含js和css子目录）
# 3. 从公共资源目录复制assets和pages到src目录
# 4. 复制AdminLTE样式文件
# 5. 复制JS和PNG文件到dist目录
# 6. 复制字体和图片资源
# 7. 准备分离主题的public目录结构
rebuild-dist:
	# 删除旧的dist目录
	rm -rf $(ASSETS_PATH)/dist
	# 创建新的dist目录
	mkdir $(ASSETS_PATH)/dist
	# 创建dist/js目录，用于存放合并后的JavaScript文件
	mkdir $(ASSETS_PATH)/dist/js
	# 创建dist/css目录，用于存放合并后的CSS文件
	mkdir $(ASSETS_PATH)/dist/css
	# 删除旧的src目录
	rm -rf $(ASSETS_PATH)/src
	# 从公共资源目录复制assets到src目录
	cp -R $(COMMON_PATH)/assets $(ASSETS_PATH)/src
	# 复制AdminLTE主样式文件到合并目录
	cp $(RESOURCE_PATH)/adminlte/adminlte.css $(ASSETS_PATH)/src/css/combine/g_all.css
	# 复制所有JavaScript文件到dist/js目录
	cp $(ASSETS_PATH)/src/js/*.js $(ASSETS_PATH)/dist/js/
	# 复制所有PNG图片文件到dist/css目录
	cp $(ASSETS_PATH)/src/css/*.png $(ASSETS_PATH)/dist/css/
	# 删除旧的pages目录
	rm -rf $(RESOURCE_PATH)/pages
	# 从公共资源目录复制pages到resource目录
	cp -R $(COMMON_PATH)/pages $(RESOURCE_PATH)/pages

	# 复制CSS字体文件到dist/css目录
	cp -R $(ASSETS_PATH)/src/css/fonts $(ASSETS_PATH)/dist/css/
	# 复制图片资源到dist目录
	cp -R $(ASSETS_PATH)/src/img $(ASSETS_PATH)/dist/
	# 复制字体文件到dist目录
	cp -R $(ASSETS_PATH)/src/fonts $(ASSETS_PATH)/dist/

	# 删除旧的分离主题public目录
	rm -rf $(SEPARATION_PATH)/public
	# 创建新的分离主题public目录
	mkdir $(SEPARATION_PATH)/public

	# 复制资源文件到分离主题的public目录
	cp -R $(RESOURCE_PATH)/assets $(SEPARATION_PATH)/public/assets
	# 复制页面文件到分离主题的public目录
	cp -R $(RESOURCE_PATH)/pages $(SEPARATION_PATH)/public/pages
	# 删除分离主题中的vendor目录（不需要打包第三方依赖）
	rm -rf $(SEPARATION_PATH)/public/assets/vendor

# 合并并压缩前端资源文件 (combine frontend assets(js/css))
# 先删除系统无用文件，然后分别合并JS和CSS文件
combine: remove combine-js combine-css

# 合并JavaScript文件
# 将多个JS文件按功能模块合并并压缩，生成带hash的文件名以支持缓存控制
combine-js:
	# 合并所有通用JS文件，生成all.min.js（带hash）
	$(CLI) merge js --hash=true --src=$(ASSETS_PATH)/src/js/all/ --dist=$(ASSETS_PATH)/dist/js/all.min.js
	# 合并第二组通用JS文件，生成all_2.min.js（带hash）
	$(CLI) merge js --hash=true --src=$(ASSETS_PATH)/src/js/all_2/ --dist=$(ASSETS_PATH)/dist/js/all_2.min.js
	# 合并表单组件JS文件，生成form.min.js（带hash）
	$(CLI) merge js --hash=true --src=$(ASSETS_PATH)/src/js/components/form/ --dist=$(ASSETS_PATH)/dist/js/form.min.js
	# 合并树形组件JS文件，生成tree.min.js（带hash）
	$(CLI) merge js --hash=true --src=$(ASSETS_PATH)/src/js/components/tree/ --dist=$(ASSETS_PATH)/dist/js/tree.min.js
	# 合并树形视图组件JS文件，生成treeview.min.js（带hash）
	$(CLI) merge js --hash=true --src=$(ASSETS_PATH)/src/js/components/treeview/ --dist=$(ASSETS_PATH)/dist/js/treeview.min.js
	# 合并数据表格组件JS文件，生成datatable.min.js（带hash）
	$(CLI) merge js --hash=true --src=$(ASSETS_PATH)/src/js/components/datatable/ --dist=$(ASSETS_PATH)/dist/js/datatable.min.js
	# 复制所有生成的JS文件到分离主题目录
	cp $(ASSETS_PATH)/dist/js/* $(SEPARATION_PATH)/public/assets/dist/js/

# 合并CSS文件
# 将CSS文件合并并压缩，生成带hash的文件名以支持缓存控制
combine-css:
	# 合并所有CSS文件，生成带hash的压缩文件
	$(CLI) merge css --hash=true
	# 复制所有生成的CSS文件到分离主题目录
	cp $(ASSETS_PATH)/dist/css/*.css $(SEPARATION_PATH)/public/assets/dist/css/

# 打包资源文件为Go代码
# 将静态资源文件打包到Go代码中，方便在编译时嵌入
build-assets:
	$(CLI) bundle asset

# 将Go模板文件转换为Go文件 (turn golang template files into a go file)
# 将所有Go模板文件(.tmpl)合并为一个Go文件，便于编译和分发
build-tmpl:
	# 使用adm工具打包模板文件，指定项目名为adminlte
	$(CLI) bundle tpl -p=adminlte

# 清理临时文件 (clean)
# 删除构建过程中生成的临时文件和目录
clean:
	# 删除src目录下的所有文件
	rm -rf $(ASSETS_PATH)/src/*
	# 删除pages目录下的所有文件
	rm -rf $(RESOURCE_PATH)/pages/*

# 格式化Go代码 (go fmt)
# 对生成的Go文件进行代码格式化，确保代码风格统一
fmt:
	# 使用go fmt工具格式化当前目录及子目录下的所有Go文件
	GO111MODULE=off go fmt ./...

# 删除系统无用文件 (remove useless system files)
# 删除系统产生的无用文件（如Mac的.DS_Store），防止被一并合并到Go文件中
remove:
	# 查找并删除所有.DS_Store文件（Mac系统自动生成的文件）
	find ./ -name ".DS_Store" -depth -exec rm {} \;

# 构建分离主题资源文件压缩包 (build assets zip of separation theme)
# 将分离主题的public目录打包成zip文件，用于独立部署
build-separation-zip:
	# 删除旧的压缩包
	rm $(SEPARATION_PATH)/public.zip
	# 进入public目录，打包所有文件为public.zip（排除.DS_Store和__MACOSX）
	cd $(SEPARATION_PATH)/public && zip -r -q ./../public.zip . -x "*.DS_Store" -x "__MACOSX"

# 声明伪目标，避免与同名文件冲突
.PHONY: all build rebuild-dist combine combine-js combine-css build-assets build-tmpl clean fmt remove