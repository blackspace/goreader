# goreader

The project is for querying the running status of linux system by the https connecting.
And the client gets the data in the json format.

Usage
----------------
First as other golang project at the github site

    go get github.com/blackspace/goreader

Install the key.gem and cert.gem files in the ~/.goreader directory for the https server

    goreader install

Start the goreader server

    goreader server

Query the status by the curl

    $ curl -k -d '["actions"]' https://127.0.0.1:10443

    {"actions":[{"Path":"/uptime","Alias":"uptime","Descript":"the uptime"},{"Path":"/version","Alias":"version","Descript":"the version"},{"Path":"/now","Alias":"now","Descript":"the now"},{"Path":"/actions","Alias":"actions","Descript":"Get all status actionses"}]}

    $ curl -k -d '["uptime","now"]' https://127.0.0.1:10443

    {"now":1451206955,"uptime":3501.67}

If you want to add new status,please edit the file:

    goreader/internal/sys_info/sys_info_linux.go

