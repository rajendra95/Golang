package main

import (
	"testing"

)


func TestSum(t  *testing.T){

	tt:=[]struct{
		testcase string
		input []int
		sum int
	}{
		{" one to five ",[]int{1,2,3,4,5}, 15} ,
		{"one and minus one" ,[]int{1,-1}, 0},
		{"no numbers" , nil,0},
	}

	for _,tc:= range tt{
		t.Run(tc.testcase,func(t *testing.T) {
			s := sumOfnums(tc.input)
			if s != tc.sum {
				t.Fatalf("sum of %v should be %v got %v ", tc.input, tc.sum, s)
			}
		})
	}

}
func TestFoo(t  *testing.T) {
	t.Errorf("test foo always fails !!!")
}
