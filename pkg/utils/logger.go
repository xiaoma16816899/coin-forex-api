package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	Gray   = "\033[37m"
	White  = "\033[97m"
)

var (
	none    = log.New(os.Stdout, Blue, log.Ldate|log.Ltime)
	verbose = log.New(os.Stdout, Blue, log.Ldate|log.Ltime)
	info    = log.New(os.Stdout, Blue, log.Ldate|log.Ltime)
	warn    = log.New(os.Stdout, Blue, log.Ldate|log.Ltime)
	debug   = log.New(os.Stdout, Blue, log.Ldate|log.Ltime|log.Lshortfile)
	fetal   = log.New(os.Stdout, Blue, log.Ldate|log.Ltime|log.Lshortfile)
)

func None(args ...any) {
	none.Output(2, fmt.Sprint(args...)+Reset)
}
func Verbose(args ...any) {
	verbose.Output(2, Cyan+"[VERBOSE] "+Reset+fmt.Sprint(args...))
}
func Info(args ...any) {
	info.Output(2, Green+"[Info] "+Reset+fmt.Sprint(args...))
}
func Warn(args ...any) {
	warn.Output(2, Yellow+"[Warn] "+Reset+fmt.Sprint(args...))
}
func Debug(args ...any) {
	debug.Output(2, Purple+"[Debug] "+Reset+fmt.Sprint(args...))
}
func Error(args ...any) {

}

// Print and call os.Exit(1)
func Fetal(args ...any) {
	fetal.Output(2, Red+"[Fetal] "+fmt.Sprint(args...)+Reset)
	os.Exit(1)
}

// Print object or json byte
func Log(v ...any) {
	if b, ok := v[0].([]byte); ok {
		dst := &bytes.Buffer{}
		if err := json.Indent(dst, b, "", "   "); err != nil {
			panic(err)
		}
		None(Reset + dst.String())
	} else {
		s, _ := json.MarshalIndent(v, "", "   ")
		None(Reset + string(s))
	}
}
