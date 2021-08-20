package bincode_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/forbole/soljuno/solana/bincode"
)

type testStruct struct {
	Bytes []uint8
}

func TestDecode(t *testing.T) {
	testCases := []struct {
		name     string
		bz       []byte
		expected interface{}
	}{
		{
			name:     "decode bool",
			bz:       []byte{1},
			expected: true,
		},
		{
			name:     "decode int8",
			bz:       []byte{1},
			expected: int8(1),
		},
		{
			name:     "decode int16",
			bz:       []byte{1, 0},
			expected: int16(1),
		},
		{
			name:     "decode int32",
			bz:       []byte{1, 0, 0, 0},
			expected: int32(1),
		},
		{
			name:     "decode int64",
			bz:       []byte{1, 0, 0, 0, 0, 0, 0, 0},
			expected: int64(1),
		},
		{
			name:     "decode uint8",
			bz:       []byte{1},
			expected: uint(1),
		},
		{
			name:     "decode uint16",
			bz:       []byte{1, 0},
			expected: uint16(1),
		},
		{
			name:     "decode uint64",
			bz:       []byte{1, 0, 0, 0, 0, 0, 0, 0},
			expected: uint64(1),
		},
		{
			name:     "decode array",
			bz:       []byte{1, 2},
			expected: [2]uint8{1, 2},
		},
		{
			name:     "decode slice",
			bz:       []byte{2, 0, 0, 0, 0, 0, 0, 0, 1, 2},
			expected: []uint8{1, 2},
		},
		{
			name:     "decode string",
			bz:       []byte{2, 0, 0, 0, 0, 0, 0, 0, 49, 50},
			expected: "12",
		},
		{
			name:     "decode struct",
			bz:       []byte{2, 0, 0, 0, 0, 0, 0, 0, 1, 2},
			expected: testStruct{Bytes: []uint8{1, 2}},
		},
		{
			name:     "decode ptr",
			bz:       []byte{1, 2, 0, 0, 0, 0, 0, 0, 0, 1, 2},
			expected: &testStruct{Bytes: []uint8{1, 2}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testFunc := func(check interface{}) {
				require.Equal(t, tc.expected, check)
			}
			testDecodeWithType(tc.expected, tc.bz, testFunc)
		})
	}
}

func testDecodeWithType(data interface{}, bz []byte, fun func(check interface{})) {
	switch data.(type) {
	case bool:
		var v bool
		bincode.Decode(bz, &v)
		fun(v)
	case int8:
		var v int8
		bincode.Decode(bz, &v)
		fun(v)
	case int16:
		var v int16
		bincode.Decode(bz, &v)
		fun(v)
	case int32:
		var v int32
		bincode.Decode(bz, &v)
		fun(v)
	case int64:
		var v int64
		bincode.Decode(bz, &v)
		fun(v)
	case uint8:
		var v uint8
		bincode.Decode(bz, &v)
		fun(v)
	case uint16:
		var v uint16
		bincode.Decode(bz, &v)
		fun(v)
	case uint32:
		var v uint32
		bincode.Decode(bz, &v)
		fun(v)
	case uint64:
		var v uint64
		bincode.Decode(bz, &v)
		fun(v)
	case [2]uint8:
		var v [2]uint8
		bincode.Decode(bz, &v)
		fun(v)
	case []uint8:
		var v []uint8
		bincode.Decode(bz, &v)
		fun(v)
	case string:
		var v string
		bincode.Decode(bz, &v)
		fun(v)
	case testStruct:
		var v testStruct
		bincode.Decode(bz, &v)
		fun(v)
	case *testStruct:
		var v *testStruct
		bincode.Decode(bz, &v)
		fun(v)
	}
}
