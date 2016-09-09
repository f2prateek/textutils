package main

import (
	"fmt"
	"os"
)

type Handler interface {
	ServeCommand(args []string)
}

type HandlerFunc func(args []string)

func (h HandlerFunc) ServeCommand(args []string) {
	h(args)
}

func HandleFunc(path string, f HandlerFunc) {
	mux[path] = f
}

var mux map[string]Handler = make(map[string]Handler, 0)

func Serve() error {
	if h, ok := mux[os.Args[1]]; ok {
		h.ServeCommand(os.Args)
		return nil
	}

	return fmt.Errorf("unknown command: %s", os.Args)
}
