# goreader

Linux系统运行状态查询

使用
----------------

启动服务器

    goreader server

查询

    curl -k -d '["version","uptime","now"]' https://127.0.0.1:10443

添加状态

    编辑goreader/internal/sys_info/sys_info_linux.go

