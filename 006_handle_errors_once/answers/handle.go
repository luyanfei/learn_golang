package handle

import (
	"encoding/json"
	"io"
	"log"
)

type Config struct {
	Address string
}

/*
You should only handle errors once. Handling an error means inspecting the error value, and making a single decision.
WriteAll函数和WriteConfig函数将同一个错误写入了日志，导致写入错误会在日志中出现两次。
处理错误意味着对错误的处理，写入日志的的行为应当只发生一次。请改写下面的代码，保证同一个错误只会被写入日志一次。
*/
func WriteAll(w io.Writer, buf []byte, logger *log.Logger) error {
	_, err := w.Write(buf)
	return err
}

func WriteConfig(w io.Writer, conf *Config, logger *log.Logger) error {
	buf, err := json.Marshal(conf)
	if err != nil {
		logger.Printf("could not marshal config: %v", err)
		return err
	}
	if err := WriteAll(w, buf, logger); err != nil {
		logger.Printf("could not write config: %v", err)
		return err
	}
	return nil
}
