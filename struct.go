/*
Copyright github.com/ada-wang    wanggang-info@ruc.edu.cn
All Rights Reserved.
SPDX-License-Identifier: Apache-2.0
*/

package bigdecimal

import (
	"errors"
	"math/big"
	"strconv"
)

// BigDecimal a struct for decimal Add/Sub
// support point
// support negative
type BigDecimal struct {
	val *big.Int
	dp  int // decimalPointPosition
}

// New returns *BigDecimal
func New(value interface{}) (*BigDecimal, error) {
	var bigDecimal BigDecimal
	bigDecimal.val = big.NewInt(0)
	bigDecimal.dp = 0
	if _, ok := value.(int); ok {
		valueStr := strconv.Itoa((value.(int)))
		return bigDecimal.FormatString(valueStr)
	} else if _, ok := value.(float64); ok {
		valueStr := strconv.FormatFloat(value.(float64), 'f', -1, 64)
		return bigDecimal.FormatString(valueStr)
	} else if _, ok := value.([]byte); ok {
		return bigDecimal.FormatString(string(value.([]byte)))
	} else if _, ok := value.(string); ok {
		return bigDecimal.FormatString((value.(string)))
	}
	return nil, errors.New("value type is not supported")
}
