package common

import (
	"io"
	"strings"
)

func HandleMsg(reader io.Reader, callback func(msg string)error) error {
    data := ""
    buf := make([]byte, 4096)
    for {
        n, err := reader.Read(buf)
        if err != nil {
            return err
        }

        data += string(buf[:n])

        index := strings.Index(data, "\n")
        for index != -1 {
            line := data[:index]
            data = data[index+1:]
            if err := callback(line); err != nil {
                return err
            }
            index = strings.Index(data, "\n")
        }
    }
}
