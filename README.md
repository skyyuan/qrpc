引包

	"http://git.qycam.com/xiaofei/qrpc/src/master/src/qrpc/client"

初始化

	client.InitGConn()
    defer client.Close()
	
	
调用

	client.Log("login", "aaa", "bbb", "info")
	
