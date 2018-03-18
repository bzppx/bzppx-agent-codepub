package main

import (
	"bzppx-agent-codepub/containers"
	"net/rpc"
	"bzppx-agent-codepub/service"
	"crypto/tls"
	"os"
)

// rpc server start

func main()  {

	rpcRegister()
	rpcStartServer()
}

// rpc register
func rpcRegister()  {
	for _, ser := range service.RegisterServices {
		rpc.Register(ser)
	}
}

// start rpc server
func rpcStartServer()  {
	log := containers.Log
	cfg := containers.Cfg
	
	listenTcp := cfg.GetString("rpc.listen")
	token := cfg.GetString("access.token")
	keyFile := cfg.GetString("cert.key_file")
	crtFile := cfg.GetString("cert.crt_file")
	
	cert, err := tls.LoadX509KeyPair(crtFile, keyFile)
	if err != nil {
		log.Errorln(err.Error())
		os.Exit(1)
	}

	tlsConf := &tls.Config{
		Certificates:[]tls.Certificate{cert},
	}
	ln, err := tls.Listen("tcp", listenTcp, tlsConf)
	if err != nil {
		log.Errorln(err.Error())
		os.Exit(1)
	}
	defer ln.Close()

	log.Info("start listen tcp port"+listenTcp)

	for {
		c, err := ln.Accept()
		buf := make([]byte, 1024)
		n, err := c.Read(buf)
		if err != nil {
			log.Error(err.Error())
			break
		}
		clientToken := string(buf[:n])
		if clientToken != token {
			c.Write([]byte("token error"))
			continue
		}else {
			c.Write([]byte("success"))
		}

		go rpc.ServeConn(c)
	}
}