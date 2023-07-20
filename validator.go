package validator

import (
	"errors"
	"fmt"
	"strings"
)

type ValidateKey string
type ValidateOption func(v *validator)

func WithKey(key ValidateKey, value interface{}) ValidateOption {
	return func(v *validator) {
		v.append(key, value)
	}
}

func FastFail(fastFail bool) ValidateOption {
	return func(v *validator) {
		v.fastFail = fastFail
	}
}

var fns map[ValidateKey]ValidateFunc

func init() {
	fns = make(map[ValidateKey]ValidateFunc)
}

type ValidateFunc func(v interface{}) error

func Register(key ValidateKey, fn ValidateFunc) {
	if _, ok := fns[key]; ok {
		panic("dumplicate key:" + key + " in validator")
	}
	fns[key] = fn
}

type Validator interface {
	Validate() error
}

type validateItem struct {
	key   ValidateKey
	value interface{}
}
type validator struct {
	fastFail     bool
	validateList []*validateItem
	validateFunc map[ValidateKey]ValidateFunc
}

func New(opts ...ValidateOption) *validator {
	v := new(validator)
	for _, opt := range opts {
		opt(v)
	}
	return v
}
func (v *validator) append(key ValidateKey, value interface{}) {
	v.validateList = append(v.validateList, &validateItem{key: key, value: value})
}

func (v *validator) Validate() error {

	var errlist []string
	for idx, item := range v.validateList {

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
