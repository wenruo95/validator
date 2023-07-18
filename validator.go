package validator

import (
	"errors"
	"fmt"
	"strings"
	"sync"
)

var once sync.Once
var fns map[string]CheckerFunc = make(map[string]CheckerFunc)

type CheckerFunc func(v interface{}) error

func Register(key string, fn CheckerFunc) {
	once.Do(func() {
		fns = make(map[string]CheckerFunc)
	})
	if _, ok := fns[key]; ok {
		panic("dumplicate key:" + key + " in validator")
	}
	fns[key] = fn
}

type Validator interface {
	Validate() error
}

type checkItem struct {
	key   string
	value interface{}
}
type validator struct {
	fastFail  bool
	checkList []*checkItem
	checkFunc map[string]CheckerFunc
}

func New(opts ...ValidateOption) *validator {
	v := new(validator)
	for _, opt := range opts {
		opt(v)
	}
	return v
}
func (v *validator) append(key string, value interface{}) {
	v.checkList = append(v.checkList, &checkItem{key: key, value: value})
}

func (v *validator) Validate() error {

	var errlist []string
	for idx, item := range v.checkList {

		fn, ok := fns[item.key]
		if !ok {
			return fmt.Errorf("idx:%v not find key:%v validate rule", idx, item.key)
		}

		if err := fn(item.value); err != nil {
			err = fmt.Errorf("%w, idx:%v key:%v value:%v", err, idx, item.key, item.value)

			if v.fastFail {
				return err
			}
			errlist = append(errlist, err.Error())
		}
	}

	if len(errlist) == 0 {
		return nil
	}
	return errors.New(strings.Join(errlist, " | "))
}
