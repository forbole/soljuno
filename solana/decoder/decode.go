package decoder

import (
	"encoding/binary"
	"math"
	"reflect"
)

type decoder struct {
	order  binary.ByteOrder
	buf    []byte
	offset int // next read offset in data
}

func Deserialize(bz []byte, data interface{}) error {
	var d decoder
	v := reflect.ValueOf(data)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	d.init(bz)
	d.value(v)
	return nil
}

func (d *decoder) init(data []byte) *decoder {
	d.order = binary.LittleEndian
	d.buf = data
	d.offset = 0
	return d
}

func (d *decoder) value(v reflect.Value) {
	switch v.Kind() {
	case reflect.Bool:
		v.SetBool(d.bool())

	case reflect.Int8:
		v.SetInt(int64(d.int8()))
	case reflect.Int16:
		v.SetInt(int64(d.int16()))
	case reflect.Int32:
		v.SetInt(int64(d.int32()))
	case reflect.Int64:
		v.SetInt(d.int64())

	case reflect.Uint8:
		v.SetUint(uint64(d.uint8()))
	case reflect.Uint16:
		v.SetUint(uint64(d.uint16()))
	case reflect.Uint32:
		v.SetUint(uint64(d.uint32()))
	case reflect.Uint64:
		v.SetUint(d.uint64())

	case reflect.Float32:
		v.SetFloat(float64(math.Float32frombits(d.uint32())))
	case reflect.Float64:
		v.SetFloat(math.Float64frombits(d.uint64()))

	case reflect.Array:
		l := v.Len()
		for i := 0; i < l; i++ {
			d.value(v.Index(i))
		}

	case reflect.Slice:
		len := d.uint64()
		typ := v.Type()
		v.Set(reflect.MakeSlice(typ, int(len), int(len)))
		for i := 0; i < int(len); i++ {
			d.value(v.Index(i))
		}

	case reflect.Struct:
		t := v.Type()
		for i := 0; i < v.NumField(); i++ {
			if v := v.Field(i); v.CanSet() || t.Field(i).Name != "_" {
				d.value(v)
			}
		}

	// TODO support Ptr
	case reflect.Ptr:
		if d.bool() {
			//d.value(v.Elem())
		}
	}
}

func (d *decoder) bool() bool {
	x := d.buf[d.offset]
	d.offset++
	return x != 0
}

func (d *decoder) uint8() uint8 {
	x := d.buf[d.offset]
	d.offset++
	return x
}

func (d *decoder) uint16() uint16 {
	x := d.order.Uint16(d.buf[d.offset : d.offset+2])
	d.offset += 2
	return x
}

func (d *decoder) uint32() uint32 {
	x := d.order.Uint32(d.buf[d.offset : d.offset+4])
	d.offset += 4
	return x
}

func (d *decoder) uint64() uint64 {
	x := d.order.Uint64(d.buf[d.offset : d.offset+8])
	d.offset += 8
	return x
}

func (d *decoder) int8() int8 { return int8(d.uint8()) }

func (d *decoder) int16() int16 { return int16(d.uint16()) }

func (d *decoder) int32() int32 { return int32(d.uint32()) }

func (d *decoder) int64() int64 { return int64(d.uint64()) }
