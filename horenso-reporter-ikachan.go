package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/jessevdk/go-flags"
	"github.com/Songmu/horenso"
	"github.com/yuuki/go-horenso-reporter-helper"
)

const version = "0.0.1"

type opts struct {
	Host   string `short:"H" long:"host" required:"true" value-name:"hostname" description:"ikachan hostname"`
	Port   int `short:"p" long:"port" value-name:"port" default:"4979" description:"ikachan port"`
	Channel   string `short:"c" long:"channel" required:"true" value-name:"'#channel'" description:"destination channel"`
	MsgType   string `short:"t" long:"type" value-name:"msgtype" default:"notice" description:"message type notice/privmsg)"`
	ErrorOnly bool   `short:"e" long:"error-only" default:"false" description:"report only when error ocurrs"`
}

type ikachanOpts struct {
	Host	string
	Port	int
	Channel string
	Message string
	MsgType string
}

func run(o *opts) error {
	r, err := horensoreporter.GetReport()
	if err != nil {
		return fmt.Errorf("Failed to get report from horenso: %s", err)
	}
	if o.ErrorOnly && *r.ExitCode == 0 {
		return nil
	}

	message := formatMessage(r)

	err = postToIkachan(&ikachanOpts{
		Host: o.Host,
		Port: o.Port,
		Channel: o.Channel,
		MsgType: o.MsgType,
		Message: message,
	})
	if err != nil {
		return fmt.Errorf("Failed to post ikachan: %s", err)
	}

	return nil
}

func formatMessage(r *horenso.Report) string {
	msg := fmt.Sprintf("command:%s hostname:%s systime:%f usertime:%f output:%s result:%s", r.Command, r.Hostname, r.SystemTime, r.UserTime, r.Output, r.Result)
	if r.Tag != "" {
		return fmt.Sprintf("tag:%s %s", msg)
	}
	return msg
}

func postToIkachan(o *ikachanOpts) error {
	client := &http.Client{
		Timeout: time.Duration(10) * time.Second,
	}
	values := url.Values{"channel": {o.Channel}, "message": {o.Message}}
	_, err := client.PostForm(fmt.Sprintf("http://%s:%d/%s", o.Host, o.Port, o.MsgType), values)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	opt := &opts{}
	p := flags.NewParser(opt, flags.Default)
	p.Usage = "--host HOSTNAME --channel '#CHANNEL' [--port=PORT] [--type=MSGTYPE] [--error-only] \n\nVerion: " + version
	_, err := p.ParseArgs(os.Args)
	if err != nil {
		if ferr, ok := err.(*flags.Error); !ok || ferr.Type != flags.ErrHelp {
			p.WriteHelp(os.Stderr)
		}
		os.Exit(2)
	}

	if err := run(opt); err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(1)
	}

	os.Exit(0)
}
