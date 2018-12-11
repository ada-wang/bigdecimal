/*
Copyright github.com/ada-wang    wanggang-info@ruc.edu.cn
All Rights Reserved.
SPDX-License-Identifier: Apache-2.0
*/

package bigdecimal

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	var err error
	t1, _ := New(0)
	fmt.Println(t1.ValString(), t1.dp)
	t1, _ = New(-1)
	fmt.Println(t1.ValString(), t1.dp)
	t1, _ = New(1.0)
	fmt.Println(t1.ValString(), t1.dp)
	t1, err = New(-0.01)
	fmt.Println(err)
	fmt.Println(t1.ValString(), t1.dp)
	t1, err = New("-01.01")
	fmt.Println(err)
	fmt.Println(t1.ValString())
	t1, err = New("100.00")
	fmt.Println(err)
	fmt.Println(t1.ValString(), t1.val.String(), t1.dp)

	t1.IncreaseDigit(5)
	if t1.ValString() != "100.0000000" {
		fmt.Println(t1.ValString())
		t.FailNow()
	}

	t1, err = New("100")
	t2, err := New("100.00")
	t1.Add(t2)
	fmt.Println(t1.ValString(), t1.val.String(), t1.dp)

	if t1.ValString() != "200.00" || t1.dp != 2 {
		t.FailNow()
	}
	t1, err = New("99")
	t2, err = New("-100.00")
	fmt.Println(t2.ValString(), t2.val.String(), t2.dp)
	t2.Add(t1)
	fmt.Println(t2.ValString(), t2.val.String(), t2.dp)
	if t2.ValString() != "-1.00" || t1.dp != 2 {
		t.FailNow()
	}

	t1, err = New("99")
	fmt.Println(t1.ValString(), t1.val.String(), t1.dp)
	t1.IncreaseDigit(2)
	fmt.Println(t1.ValString(), t1.val.String(), t1.dp)

	t1, err = New("-99.99")
	fmt.Println(t1.ValString(), t1.val.String(), t1.dp)
	t1.IncreaseDigit(2)
	fmt.Println(t1.ValString(), t1.val.String(), t1.dp)

	t1, err = New("-9.99")
	fmt.Println(t1.ValString(), t1.val.String(), t1.dp)
	t1.IncreaseDigit(2)
	fmt.Println(t1.ValString(), t1.val.String(), t1.dp)

	t1, err = New("-0.0099")
	fmt.Println(t1.ValString(), t1.val.String(), t1.dp)
	t1.IncreaseDigit(2)
	fmt.Println(t1.ValString(), t1.val.String(), t1.dp)

	t1, err = New("-0.0099")
	t2, err = New("9900")
	t1.Sub(t2)
	fmt.Println(t1.ValString(), t1.val.String(), t1.dp)
	fmt.Println(t2.ValString(), t2.val.String(), t2.dp)

	t1, err = New("-0.0099")
	t2, err = New("9900")
	t2.Sub(t1)
	fmt.Println(t1.ValString(), t1.val.String(), t1.dp)
	fmt.Println(t2.ValString(), t2.val.String(), t2.dp)

	// test IncreaseDigit
	var t3 BigDecimal
	err = t3.IncreaseDigit(1)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		t.Error("t3.IncreaseDigit(1) t3 should New first")
	}

}
