package datetime

import (
	"fmt"
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
	"strconv"
	"time"
)

type fnSubMins struct {
}

func init() {
	function.Register(&fnSubMins{})
}

func (s *fnSubMins) Name() string {
	return "subMins"
}

func (s *fnSubMins) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeDateTime, data.TypeFloat64}, false
}

func (s *fnSubMins) Eval(in ...interface{}) (interface{}, error) {
	t, err := coerce.ToDateTime(in[0])
	if err != nil {
		return nil, err
	}
	mins, err := coerce.ToFloat64(in[1])
	if err != nil {
		return nil, err
	}

	d, err := time.ParseDuration("-" + strconv.FormatFloat(mins, 'f', -1, 64) + "m")
	if err != nil {
		return nil, fmt.Errorf("Invalid minutes [%f]", mins)
	}
	return t.Add(d), nil

}
