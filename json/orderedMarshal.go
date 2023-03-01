package json

import (
	"bytes"
)

// MapSlice enables you to marshal a map that is ordered
type MapSlice []MapItem

type MapItem struct {
	Key, Value string
}

func (ms *MapSlice) Add(k, v string) *MapSlice {
	*ms = append(*ms, MapItem{
		Key:   k,
		Value: v,
	})
	return ms
}

func (ms *MapSlice) AddMap(m map[string]string) *MapSlice {
	for k, v := range m {
		*ms = append(*ms, MapItem{Key: k, Value: v})
	}
	return ms
}

func (ms *MapSlice) Marshal() []byte {
	buf := &bytes.Buffer{}
	buf.Write([]byte{'{'})
	for i, mi := range *ms {
		buf.WriteByte('"')
		buf.WriteString(mi.Key)
		buf.WriteString(`:"`)
		buf.WriteString(mi.Value)
		buf.WriteByte('"')
		if i < len(*ms)-1 {
			buf.Write([]byte{','})
		}
	}
	buf.Write([]byte{'}'})
	return buf.Bytes()
}
