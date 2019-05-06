package main

import rpcx "gnet/rpcx/client"

func main() {
	rpcx.SynchronousCall()
	rpcx.AsynchronousCall()
	//rpcx.PrintWrr()
	//rpcx.PrintWrrNgx()
}
