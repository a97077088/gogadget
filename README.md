"# gogadget 基于gum-js的go绑定" 

go get github.com/a97077088/gogadget


可以运行在osx x86_64
可以运行在ios arm64

先下载frida源代码，编译一次
make gum-ios-thin 生成工具链后连接到工具链，注意修改ld_x

其他平台可以拓展ld_
