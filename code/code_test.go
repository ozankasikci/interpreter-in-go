package code

import "testing"

func TestMake(t *testing.T) {
	tests := []struct{
		op Opcode
		operands []int
		expected []byte
	} {
		{ OpConstant, []int{65534}, []byte{byte(OpConstant), 255, 254}},
	}
}
