package validator

import (
	"errors"
	"strconv"
)

type ValidateOption func(v *validator)

func init() {
	Register(vKeyPhone, func(v interface{}) error {
		phone, ok := v.(string)
		if !ok {
			return errors.New("invalid type")
		}
		if len(phone) != 11 {
			return errors.New("invalid length:" + strconv.Itoa(len(phone)))
		}
		return nil
	})

	Register(vKeyAddr, func(v interface{}) error {
		addr, ok := v.(string)
		if !ok {
			return errors.New("invalid type")
		}
		if len(addr) == 0 {
			return errors.New("empty value")
		}
		return nil
	})

	Register(vKeyName, func(v interface{}) error {
		name, ok := v.(string)
		if !ok {
			return errors.New("invalid type")
		}
		if len(name) == 0 {
			return errors.New("empty value")
		}
		return nil
	})

}

const (
	vKeyPhone string = "phone"
	vKeyAddr  string = "addr"
	vKeyName  string = "name"
)

func Phone(phone string) ValidateOption {
	return func(v *validator) {
		v.append(vKeyPhone, phone)
	}
}
func Address(addr string) ValidateOption {
	return func(v *validator) {
		v.append(vKeyAddr, addr)
	}
}
func UserName(name string) ValidateOption {
	return func(v *validator) {
		v.append(vKeyName, name)
	}
}

func WithKey(key string, value interface{}) ValidateOption {
	return func(v *validator) {
		v.append(key, value)
	}
}

func FastFail(fastFail bool) ValidateOption {
	return func(v *validator) {
		v.fastFail = fastFail
	}
}
