package serializers

import (
	"encoding/xml"
)

type Xml struct {
}

func (x Xml) Serialize(i any) string {
	var (
		buf []byte
		err error
	)

	if buf, err = xml.Marshal(i); err != nil {
		panic(err)
	}

	return string(buf)
}

func (x Xml) Unserialize(s string, i any) error {
	if err := xml.Unmarshal([]byte(s), i); err != nil {
		return err
	}
	return nil
}
