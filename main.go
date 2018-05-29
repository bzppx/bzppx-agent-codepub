package main

import (
	"bzppx-agent-codepub/app"
	"net/rpc"
	"bzppx-agent-codepub/app/service"
	"crypto/tls"
	"os"
	"net"
	"bzppx-agent-codepub/container"
	"bzppx-agent-codepub/utils"
)

var (
	TokenError = []byte("token error")
	TokenSuccess = []byte("success")
)

func main()  {

	// register rpc service
	service.RegisterRpc()

	// task worker start
	go container.NewWorker().StartTask()

	// start rpc  server
	rpcStartServer()
}

// start rpc server
func rpcStartServer()  {

	listenAddr := app.Conf.GetString("rpc.listen")
	token := app.Conf.GetString("access.token")
	keyFile := app.Conf.GetString("cert.key_file")
	crtFile := app.Conf.GetString("cert.crt_file")

	// load cert key
	cert, err := tls.LoadX509KeyPair(crtFile, keyFile)
	if err != nil {
		app.Log.Errorf("agent tls config load error, %s", err.Error())
		os.Exit(1)
	}
	tlsConf := &tls.Config{
		Certificates:[]tls.Certificate{cert},
	}
	ln, err := tls.Listen("tcp", listenAddr, tlsConf)
	if err != nil {
		app.Log.Errorf("tls listen error, %s", err.Error())
		os.Exit(1)
	}
	defer ln.Close()

	app.Log.Infof("agent start listen %s", listenAddr)

	for {
		conn, err := ln.Accept()
		if err != nil {
			app.Log.Errorf("agent accept error, %s", err.Error())
			break
		}
		clientToken, err := utils.Codec.DecodePack(conn)
		if err != nil {
			app.Log.Errorf("conn read token error, %s", err.Error())
			conn.Close()
			continue
		}

		// read byte and encode pack
		var checkRes []byte
		if clientToken != token {
			checkRes, err = utils.Codec.EncodePack(TokenError)
			conn.Write(checkRes)
			conn.Close()
			continue
		}else {
			checkRes, err = utils.Codec.EncodePack(TokenSuccess)
			conn.Write(checkRes)
		}

		// rpc conn serve
		go func(c *net.Conn) {
			defer func() {
				e := recover()
				if e != nil {
					app.Log.Errorf("conn rpc crash, %v", e)
				}
			}()
			rpc.ServeConn(*c)
		}(&conn)
	}
}