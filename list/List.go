package list

import (
	"reflect"
)

var nil Type
type Type int

func RemoveByVal(list []reflect.Type, v reflect.Type)  {

}



func RemoveByIdx(list []string, idx int) []string {
	return list[:idx+copy(list[idx:], list[idx+1:])];
}