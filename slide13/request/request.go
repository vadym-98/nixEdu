package request

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type Validatable interface {
	Validate() []string
}

func ParseJsonRequest(r *http.Request, v Validatable) error {
	ctHeader := r.Header.Get("Content-Type")
	if ctHeader != "application/json" {
		return fmt.Errorf("unsupported Content-Type header. Use application/json")
	}

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(v)
	if err != nil {
		return fmt.Errorf("failed decode request body. err: %s", err)
	}

	return Validate(v)
}

func Validate(v Validatable) error {
	if errs := v.Validate(); len(errs) > 0 {
		return errors.New(strings.Join(errs, "; "))
	}

	return nil
}
