package main

import (
	"github.com/kardianos/service"
	"log"
	"os"
)

type program struct{}

func (p *program) Start(s service.Service) error {
	go p.run()
	return nil
}
func (p *program) Stop(s service.Service) error {
	return nil
}
func (p *program) run() {
	// 这里放置程序要执行的代码……
}

func main() {
	cfg := &service.Config{
		Name:        "defunct",
		DisplayName: "a simple defunct service",
		Description: "a long-running defunct process",
	}
	prg := &program{}
	s, err := service.New(prg, cfg)
	if err != nil {
		log.Fatal(err)
	}

	logger, err := s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}
	if len(os.Args) == 2 { //如果有命令则执行
		err = service.Control(s, os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
	} else { //否则说明是方法启动了
		err = s.Run()
		if err != nil {
			logger.Error(err)
		}
	}
	if err != nil {
		logger.Error(err)
	}
}
