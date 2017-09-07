package comm

import(
	"fmt"
	"github.com/op/go-logging"
	"os"
)

func NewLog(logname string) (*logging.Logger, error){
	log := logging.MustGetLogger(logname);

	format := logging.MustStringFormatter(
		`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}
		 %{color:reset} %{message}`,
	)

	logfile, err := os.OpenFile(logname, os.O_WRONLY, 0666);
	if err != nil {
		fmt.Println(err);
		return log , err;
	}

	backend1 := logging.NewLogBackend(logfile, "", 0);
	backend2 := logging.NewLogBackend(os.Stderr, "", 0);

	//backend1Formatter := logging.NewBackendFormatter(backend1, format);
	backend2Formatter := logging.NewBackendFormatter(backend2, format);

	backend1Leveled := logging.AddModuleLevel(backend1);
	backend1Leveled.SetLevel(logging.INFO, "");

	logging.SetBackend(backend1Leveled, backend2Formatter);

	log.Info("in func test log");
	return log, nil;
}  