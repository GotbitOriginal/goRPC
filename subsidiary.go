package gorpc

import (
	"errors"
	"math/big"
)

func floatify(intString string, precision int) (float64, error) {
	if len(intString) <= precision {
		toAdd := precision - len(intString) + 1
		for i := 0; i < toAdd; i++ {
			intString = "0" + intString
		}
	}
	index := len(intString) - precision
	q := intString[:index] + "." + intString[index:]
	f, ok := new(big.Float).SetString(q)
	if !ok {
		return 0, errors.New("!!floatify(): Could not parse this resulting string to float: " + q)
	}
	res, _ := f.Float64()
	return res, nil
}
