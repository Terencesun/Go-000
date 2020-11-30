package service

import (
	dao "week02/dao"
)

func Service() (ret []byte, err error) {
	err = dao.FindSomething()
	if err != nil {
		return nil, err
	}
	return []byte("hello"), nil
}
