package reader

import "testing"

type cfg struct {
}

func (c cfg) GetListenIP() string {
	panic("implement me")
}

func (c cfg) GetListenPort() uint16 {
	panic("implement me")
}

func BenchmarkUDPReader(b *testing.B) {
	reader, _ := NewUDPReader(&cfg{})
	defer reader.Close()

	for i := 0; i < b.N; i++ {
		_, _ = reader.RecvBytes()
	}
}
