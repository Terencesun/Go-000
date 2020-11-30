package dao

import (
	"github.com/pkg/errors"
	"fmt"
	"time"
)

var ErrNoRows = errors.New("no found.")

func FindOne() (err error) {
	return ErrNoRows
}

func FindSomething() error {
	fmt.Println("正在查询 -> 假查询，等待2s，抛出错误")
	err := FindOne()
	time.Sleep(time.Second * 2)
	if errors.Is(err, ErrNoRows) {
		return errors.New("NoFoundData")
	} else {
		return nil
	}
}
