package reader

import (
	"errors"
	"fmt"
	"gitlab.rdp.ru/tt/events-collector-queue/internal/reader"
	"io"
	"net"

	"github.com/sirupsen/logrus"
)

const (
	MaxPayloadLength = 10000
)

type UDPServer struct {
	cfg  reader.Config
	conn net.PacketConn
}

func NewUDPServer(cfg reader.Config) (*UDPServer, error) {
	srv := &UDPServer{
		cfg: cfg,
	}

	// если останется эта реализация, то srv.cfg.GetListenPort -> string
	conn, err := net.ListenPacket("udp", ":8081")
	if err != nil {
		return nil, fmt.Errorf("net.ListenPacket(...): %w", err)
	}

	srv.conn = conn

	return srv, nil
}

func (s *UDPServer) Close() error {
	defer logrus.Info("UDP server stops listening")

	return s.conn.Close()
}

func (s *UDPServer) RecvBytes() ([]byte, error) {
	payload := make([]byte, MaxPayloadLength)
	n, _, err := s.conn.ReadFrom(payload)

	if errors.Is(err, io.EOF) {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("s.conn.ReadFrom(...): %w", err)
	}

	return payload[:n], nil
}
