package keyboard

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func GetFloat() (float64, error) {
	reader :=bufio.NewReader(os.Stdin)
	value, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	// 去除换行符和前后空格
	value = strings.TrimSpace(value)
	res, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, err
	}
	return res, nil
}
