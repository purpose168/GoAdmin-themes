# GoAdmin 官方主题

- [adminlte](https://github.com/purpose168/GoAdmin-themes/tree/master/adminlte)
- [sword](https://github.com/purpose168/GoAdmin-themes/tree/master/sword)

[中文介绍](./README_CN.md)

## 如何使用

- 导入主题
- 在 GoAdmin 的全局配置中进行设置

```go

package main

import (
	...
	_ "github.com/purpose168/GoAdmin-themes/adminlte" // 导入 adminlte 主题
	...
)

func main()  {
	
	...
	
	cfg := config.Config{
    		...
    		
    		Theme: "adminlte", // 设置使用的主题为 adminlte
    		
    		...
    	}
	
	...
 
}

```

## 如何修改并使其生效

使用每个主题目录下的 Makefile。