package tcp

import (
	"os"
	"fmt"
	"time"
)

func die(format string, v ...interface{}) {
	os.Stderr.WriteString(fmt.Sprintf(format+"\n", v...))
	os.Exit(1)
}

func formatTime(t time.Time) string {
	return t.Format("2006.01.02-15.04.05")
}

func connectionLogger(data chan []byte, conn_n int, local_info, remote_info string) {
	log_name := fmt.Sprintf("log-%s-%04d-%s-%s.log", formatTime(time.Now()),
		conn_n, local_info, remote_info)
	loggerLoop(data, log_name)
}

func binaryLogger(data chan []byte, conn_n int, peer string) {
	log_name := fmt.Sprintf("log-binary-%s-%04d-%s.log", formatTime(time.Now()),
		conn_n, peer)
	loggerLoop(data, log_name)
}

func loggerLoop(data chan []byte, log_name string) {
	f, err := os.Create(log_name)
	if err != nil {
		die("Unable to create file %s, %v\n", log_name, err)
	}
	defer f.Close()
	for {
		b := <-data
		if len(b) == 0 {
			break
		}
		f.Write(b)
		f.Sync()
	}
}
