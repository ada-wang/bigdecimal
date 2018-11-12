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
}
