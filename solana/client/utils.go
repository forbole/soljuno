package client

import "reflect"

func isEmpty(obj interface{}) bool {
	return reflect.ValueOf(obj).IsZero()
}
