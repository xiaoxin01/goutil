package strutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubBytesWithRuneCheck(t *testing.T) {
	s := []byte(`123456789壹贰叁肆伍陆柒捌玖éèǔ`)
	testCases := []struct {
		desc  string
		limit int
		want  []byte
	}{
		{
			desc:  "get 1 byte words",
			limit: 9,
			want:  []byte("123456789"),
		},
		{
			desc:  "get 1 and 3 bytes words",
			limit: 12,
			want:  []byte("123456789壹"),
		},
		{
			desc:  "get 1 and 3 bytes words skip 1 byte",
			limit: 13,
			want:  []byte("123456789壹"),
		},
		{
			desc:  "get 1, 2 and 3 bytes words",
			limit: 40,
			want:  []byte("123456789壹贰叁肆伍陆柒捌玖éè"),
		},
		{
			desc:  "get 1, 2 and 3 bytes words skip 1 byte",
			limit: 41,
			want:  []byte("123456789壹贰叁肆伍陆柒捌玖éè"),
		},
		{
			desc:  "get nil",
			limit: 0,
			want:  nil,
		},
		{
			desc:  "get words with limit overflow",
			limit: len(s) * 2,
			want:  []byte(`123456789壹贰叁肆伍陆柒捌玖éèǔ`),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := SubBytesWithRuneCheck(s, tC.limit)
			assert.Equal(t, tC.want, result)
		})
	}
}

func TestSubStringWithByteLimit(t *testing.T) {
	s := `123456789壹贰叁肆伍陆柒捌玖éèǔ`
	testCases := []struct {
		desc  string
		limit int
		want  string
	}{
		{
			desc:  "get 1 byte words",
			limit: 12,
			want:  "123456789壹",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := SubStringWithByteLimit(s, tC.limit)
			assert.Equal(t, tC.want, result)
		})
	}
}

func BenchmarkSubBytesWithRuneCheck(b *testing.B) {
	s := []byte(`123456789壹贰叁肆伍陆柒捌玖éèǔ`)
	for i := 0; i < b.N; i++ {
		SubBytesWithRuneCheck(s, 20)
	}
}

func BenchmarkSubStringWithByteLimit(b *testing.B) {
	s := `123456789壹贰叁肆伍陆柒捌玖éèǔ`
	for i := 0; i < b.N; i++ {
		SubStringWithByteLimit(s, 20)
	}
}
