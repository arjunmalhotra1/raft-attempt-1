package messaging

import (
	"encoding/binary"
	"io"
	"net"
)

func SendMessage(conn net.Conn, message []byte) error {
	// First, we need to send the length of the message as a 4-byte integer
	// Convert the message length to a 4-byte binary representation
	length := make([]byte, 4)
	binary.BigEndian.PutUint32(length, uint32(len(message)))

	// Write the length to the connection
	if _, err := conn.Write(length); err != nil {
		return err
	}

	// Then, write the message itself to the connection
	if _, err := conn.Write(message); err != nil {
		return err
	}

	return nil
}

func ReceiveMessage(conn net.Conn) ([]byte, error) {
	// First, we need to read the 4-byte length of the message
	lengthBytes := make([]byte, 4)
	if _, err := io.ReadFull(conn, lengthBytes); err != nil {
		return nil, err
	}

	// Convert the 4-byte length to an integer
	length := binary.BigEndian.Uint32(lengthBytes)

	// Read the message itself
	message := make([]byte, length)
	if _, err := io.ReadFull(conn, message); err != nil {
		return nil, err
	}

	return message, nil
}
