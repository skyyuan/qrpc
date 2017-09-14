引包

	"github.com/skyyuan/qrpc"
	
项目中

    ✗ go get github.com/skyyuan/qrpc
    ✗ godep restore
    ✗ godep save

初始化

	qrpc.InitGConn()
    defer qrpc.Close()
	
	
调用

	qrpc.Log("login", "aaa", "bbb", "info")
	
