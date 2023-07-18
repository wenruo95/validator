package validator

import (
	"errors"
	"strings"
	"testing"
	"time"
)

func TestHelloWorld(t *testing.T) {
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
		t.Logf("error:%v", err.Error())
	}

	opts = append(opts, FastFail(true))
	if err := New(opts...).Validate(); err != nil {
		t.Logf("with_fast_fail. error:%v", err)
	}

}
