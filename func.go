/*
Copyright github.com/ada-wang    wanggang-info@ruc.edu.cn
All Rights Reserved.
SPDX-License-Identifier: Apache-2.0
*/

package bigdecimal

import (
	"bytes"
	"errors"
	"strconv"
)

// FormatString returns *BigDecimal by parsing valueStr to Struct
func (t *BigDecimal) FormatString(valueStr string) (*BigDecimal, error) {
	neg, dp, noPointStr, err := scanner(valueStr)
	if err != nil {
		return nil, err
	}
	t.dp = dp
	t.val.SetString(noPointStr, 10)
	if neg == false && t.val.Sign() >= 0 {
		return t, nil
	} else if neg == true && t.val.Sign() < 0 {
		return t, nil
	}
	return nil, errors.New("negative parse error")
}

func scanner(valueStr string) (bool, int, string, error) {
	if valueStr == "" {
		return false, 0, "", errors.New("value is empty")
	}
	valueStrLength := len(valueStr)
	index := 0
	dashFound := false
	pointFound := false
	decimalPointPosition := 0
	var b bytes.Buffer
	b.Grow(valueStrLength)
	if valueStr[0] == '-' {
		dashFound = true
		index = 1
		b.WriteRune('-')
	}
	// after dash check, the next char should be number
	if valueStr[index] < '0' || valueStr[index] > '9' {
		return false, 0, "", errors.New("value is not decimal")
	}
	for ; index < valueStrLength; index++ {
		if valueStr[index] >= '0' && valueStr[index] <= '9' {
			b.WriteByte((valueStr[index]))
		} else if valueStr[index] == '.' {
			if pointFound == false {
				// just find point, do not write
				decimalPointPosition = valueStrLength - 1 - index
				pointFound = true
			} else {
				return false, 0, "", errors.New("found the second point, index is " + strconv.Itoa(index))
			}
		} else {
			return false, 0, "", errors.New("value is not decimal, index is " + strconv.Itoa(index))
		}
	}
	return dashFound, decimalPointPosition, b.String(), nil
}

// ValString returns string just like input
func (t *BigDecimal) ValString() string {
	switch t.val.Sign() {
	case 0:
		return "0"
	case -1:
		valStr := t.val.String()
		var b bytes.Buffer
		b.Grow(len(valStr) + t.dp)
		b.WriteRune('-')
		writeBuffer(&b, valStr[1:], t.dp)
		return b.String()
	case 1:
		valStr := t.val.String()
		var b bytes.Buffer
		b.Grow(len(valStr) + t.dp)
		writeBuffer(&b, valStr, t.dp)
		return b.String()
	}
	return ""
}

func writeBuffer(b *bytes.Buffer, absStr string, dp int) *bytes.Buffer {
	if len(absStr) >= dp+1 {
		// -1.001 -> 1001 -> len 4 == 3+1 -> -11.001 >= 3
		b.WriteString(absStr[0 : len(absStr)-dp])
		// check if '.' is the last one
		if dp != 0 {
			b.WriteRune('.')
			b.WriteString(absStr[len(absStr)-dp : len(absStr)])
		}
	} else {
		// -0.0011 -> 1 -> len 1 < 4+1
		b.WriteRune('0')
		//
		b.WriteRune('.')
		for index := 0; index < dp-len(absStr); index++ {
			b.WriteRune('0')
		}
		b.WriteString(absStr)
	}
	return b
}

// IncreaseDigit to modify decimalPointPosition
func (t *BigDecimal) IncreaseDigit(digitIncreaseBy int) error {
	if digitIncreaseBy <= 0 {
		return errors.New("digitIncreaseBy should greater than 0")
	}
	oldIntStr := t.val.String()
	var b bytes.Buffer
	b.Grow(len(oldIntStr) + digitIncreaseBy)
	b.WriteString(oldIntStr)
	for index := 0; index < digitIncreaseBy; index++ {
		b.WriteRune('0')
	}
	t.val.SetString(b.String(), 10)
	t.dp += digitIncreaseBy
	return nil
}
