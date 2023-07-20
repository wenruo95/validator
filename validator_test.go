package validator

import (
	"errors"
	"strconv"
	"strings"
	"testing"
	"time"
)

func doRegisterInit() {
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
	vKeyPhone ValidateKey = "phone"
	vKeyAddr  ValidateKey = "addr"
	vKeyName  ValidateKey = "name"
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

func TestHelloWorld(t *testing.T) {
	doRegisterInit()

	Register("email", func(v interface{}) error {
		email, ok := v.(string)
		if !ok {
			return errors.New("invalid type")
		}
		if len(email) == 0 {
			return errors.New("empty email")
		}
		if strings.Index(email, "@") == -1 {
			return errors.New("empty email")
		}
		return nil
	})
	time.Sleep(time.Second)

	var phone, addr, name, email string
	opts := []ValidateOption{
		Phone(phone), Address(addr), UserName(name), WithKey("email", email),
	}

	if err := New(opts...).Validate(); err != nil {
		t.Logf("show_all_error. error:%v", err.Error())
	}

	opts = append(opts, FastFail(true))
	if err := New(opts...).Validate(); err != nil {
		t.Logf("show_first_error. error:%v", err)
	}

}
