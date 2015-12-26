# goreader

The project is for querying the running status of linux system

Usage
----------------

Start the goreader server

    goreader install 
    goreader server

Query the status by the curl

    curl -k -d '["version","uptime","now"]' https://127.0.0.1:10443


If you want to add new status,please edit the file:

    goreader/internal/sys_info/sys_info_linux.go

