package form

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/pkg/errors"
	punqy "github.com/punqy/core"
	nethttp "net/http"
	"strings"
)

func GetError(dest interface{}) error {
	return Validate(dest)
}

func Validate(s interface{}) error {
	_, err := govalidator.ValidateStruct(s)
	return err
}

func NewValidationErrJsonResponse(error error) punqy.Response {
	switch error.(type) {
	case govalidator.Errors:
	default:
		return punqy.NewErrorJsonResponse(errors.New("error must be of type govalidator.errors"))
	}
	errs := error.(govalidator.Errors).Errors()
	var errorsList []govalidator.Error
	for _, e := range errs {
		makeFormError(e, &errorsList)
	}
	var errorsMap = make(map[string]interface{}, 0)
	for _, e := range errorsList {
		set(toPath(e), e.Err.Error(), &errorsMap)
	}
	return punqy.NewJsonResponse(errorsMap, nethttp.StatusBadRequest, error)
}

func makeFormError(err interface{}, errorsMap *[]govalidator.Error) {
	e, ok := err.(govalidator.Error)
	if ok {
		*errorsMap = append(*errorsMap, e)
	}
	errs, ok := err.(govalidator.Errors)
	if ok {
		for _, e := range errs {
			makeFormError(e, errorsMap)
		}
	}
}

func set(path []string, val interface{}, target *map[string]interface{}) {
	t := *target
	for i, key := range path {
		key = fmt.Sprintf("%s%s", strings.ToLower(string(key[0])), key[1:])
		if i == len(path)-1 {
			t[key] = val
			*target = t
			return
		}
		if _, ok := t[key]; !ok {
			nextTarget := make(map[string]interface{}, 0)
			set(path[1:], val, &nextTarget)
			t[key] = nextTarget
			*target = t
		} else {
			nextTargetPtr := t[key]
			nextTarget := nextTargetPtr.(map[string]interface{})
			set(path[1:], val, &nextTarget)
			t[key] = nextTarget
			*target = t
		}
		return
	}
}

func toPath(e govalidator.Error) []string {
	if len(e.Path) < 1 {
		return []string{e.Name}
	}
	return append(e.Path, e.Name)
}
