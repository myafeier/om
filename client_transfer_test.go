package om

import (
	"testing"
)

func TestTransferFromExtToExt(t *testing.T) {
	c, err := NewOmClient()
	if err != nil {
		t.Error(err)
	}
	t.Log(c.TransferFromExtToExt("2010", "3005"))
}

func TestTransferFromExtToOuter(t *testing.T) {
	c, err := NewOmClient()
	if err != nil {
		t.Error(err)
	}
	c.TransferFromExtToOuter("3005", "18987092886")
}
