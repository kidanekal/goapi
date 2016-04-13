package logger

import (
	"log/syslog"
	"os"
	"os/exec"
	"strings"

	"github.com/inconshreveable/log15"
	"github.com/kidanekal/goapi/constants"
	"github.com/onsi/ginkgo"
)

var cliLog = log15.New("goapi", constants.Version)
var hostLog = log15.New("goapi", constants.Version)

func init() {
	out, err := exec.Command("uname").Output()
	if err == nil && strings.TrimSpace(string(out)) == "Linux" {
		initHostLog()
	}

	initRootLog()
}

func initRootLog() {
	if os.Getenv("TEST") == "true" {
		logFmt := log15.TerminalFormat()

		handler := log15.StreamHandler(ginkgo.GinkgoWriter, logFmt)
		handler = log15.LvlFilterHandler(log15.LvlError, handler)
		handler = log15.CallerStackHandler("%+v", handler)

		log15.Root().SetHandler(handler)
	} else {
		// the default logging format is best for logging to stdout
	}
}

func initHostLog() {
	logFmt := log15.LogfmtFormat()

	writer, err := syslog.Dial("udp", "localhost:514", syslog.LOG_LOCAL5|syslog.LOG_INFO, constants.ApplicationName)
	if err != nil {
		panic(err)
	}

	stackHandler := log15.StreamHandler(writer, logFmt)
	stackHandler = log15.CallerStackHandler("%+v", stackHandler)
	// put filter last because it will be run first
	stackHandler = log15.FilterHandler(func(r *log15.Record) bool {
		return r.Lvl <= log15.LvlWarn
	}, stackHandler)

	infoHandler := log15.StreamHandler(writer, logFmt)
	infoHandler = log15.FilterHandler(func(r *log15.Record) bool {
		return r.Lvl == log15.LvlInfo
	}, infoHandler)

	hostLog.SetHandler(log15.MultiHandler(stackHandler, infoHandler))
}

func CLI(kvps ...interface{}) log15.Logger {
	return cliLog.New(kvps...)
}

func Host(kvps ...interface{}) log15.Logger {
	return hostLog.New(kvps...)
}
