package dao

import (
	"errors"
	"fmt"
	"time"
)

var ErrNoRoes = errors.New("no found.")

func FindSomething() error {
	fmt.Println("正在查询 -> 假查询，等待2s，抛出错误")
	time.Sleep(time.Second * 2)
	return ErrNoRoes
}
