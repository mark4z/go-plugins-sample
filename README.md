```dockerfile
docker build -t go-plugins . && docker run -v /tmp/go-plugins/out:/go/src/out go-plugins 
```

使用步骤：

1.在镜像中打包插件（+版本号） 
2.配置文件中增删插件 
3.watch事件，与缓存插件列表对比，先卸载，后加载，过程中加锁



Go plugin有如下几个**限制**：

1.**插件实现和主应用程序都必须使用完全相同的 Go工具链版本构建，且不支持win**

在与运行环境一致的镜像打包plugin可能是最简单可靠的方式

2.**公共依赖项**

这意味着插件要么要在项目内打包，要么就得把引用的包独立出来。 而且未来主包更新，插件大概率需要重新编译

3.**不能卸载**，内存占用，只能重启解决