package main

import (
	"fmt"
	"time"

	"github.com/jessevdk/go-flags"

	"github.com/alph4numb3r/netsuite-common"
)

var opts struct {
	Ports []uint16 `short:"p" long:"port" description:"Port to scan. Can be provided multiple times to scan multiple ports" required:"true"`
	Concurrency int `short:"c" description:"Max amount of concurrent connections" default:"500"`
	Delay uint `short:"d" long:"delay" description:"Delay between pings in ms" default:"50"`
	Timeout uint `short:"t" long:"timeout" description:"Timeout of each ping in ms" default:"500"`
	Args struct {
		IP string
	} `positional-args:"yes" required:"yes"`
}

var initerr error

func init(){
	_, initerr = flags.Parse(&opts)
}

func main(){
	if initerr != nil {
		panic(initerr)
	}
	p := netsuite.NewPortSniffer(map[string]interface{}{
		"timeout":time.Duration(opts.Timeout) * time.Millisecond,
		"maxConc":opts.Concurrency,
		"delay":time.Duration(opts.Delay) * time.Millisecond})
	res,err := p.PortSniffArray(opts.Args.IP,opts.Ports)
	if err != nil {
		panic(err)
	}
	fmt.Println("Scanned address ",opts.Args.IP," at ports ",opts.Ports,". \n  Found open ports ",res)
}