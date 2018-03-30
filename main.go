package main

import (
	"bzppx-agent-codepub/containers"
	"net/rpc"
	"bzppx-agent-codepub/service"
	"crypto/tls"
	"os"
	"net"
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
		log.Errorf("agent tls config load error, %s",err.Error())
		os.Exit(1)
	}

	tlsConf := &tls.Config{
		Certificates:[]tls.Certificate{cert},
	}
	ln, err := tls.Listen("tcp", listenTcp, tlsConf)
	if err != nil {
		log.Errorf("tls listen error, %s", err.Error())
		os.Exit(1)
	}
	defer ln.Close()

	log.Info("agent start listen tcp port"+listenTcp)

	for {
		c, err := ln.Accept()
		if err != nil {
			log.Errorf("agent accept error, %s", err.Error())
			break
		}
		buf := make([]byte, 1024)
		n, err := c.Read(buf)
		if err != nil {
			log.Errorf("conn read error, %s", err.Error())
			break
		}
		clientToken := string(buf[:n])
		if clientToken != token {
			c.Write([]byte("token error"))
			continue
		}else {
			c.Write([]byte("success"))
		}

		go func(c *net.Conn) {
			defer func() {
				e := recover()
				if e != nil {
					log.Errorf("conn rpc crash, %v", e)
				}
			}()
			rpc.ServeConn(*c)
		}(&c)

	}
}