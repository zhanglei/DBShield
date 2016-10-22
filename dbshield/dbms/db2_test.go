package dbms_test

import (
	"errors"
	"net"
	"testing"

	"github.com/nim4/DBShield/dbshield/dbms"
	"github.com/nim4/mock"
)

var db2Count int

func db2DummyReader(c net.Conn) (buf []byte, err error) {
	sampleIO := [][]byte{
		{
			0x00, 0xd0, 0xd0, 0x41, 0x00, 0x01, 0x00, 0xca, 0x10, 0x41, 0x00, 0x82,
			0x11, 0x5e, 0x97, 0xa8, 0xa3, 0x88, 0x96, 0x95, 0x40, 0x40, 0x40, 0x40,
			0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0xf0, 0xf0,
			0xf0, 0xf0, 0xf7, 0xf8, 0xf8, 0xf1, 0xf0, 0xf0, 0xf0, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x60, 0xf0, 0xf0, 0xf0, 0xf1, 0xa4, 0xa2,
			0x85, 0x99, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40,
			0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40,
			0x40, 0x40, 0x40, 0x40, 0xe2, 0xc1, 0xd4, 0xd7, 0xd3, 0xc5, 0x40, 0x40,
			0xf0, 0xc4, 0xc2, 0xf2, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40,
			0x40, 0x40, 0x40, 0x40, 0x40, 0xf0, 0xf0, 0xf1, 0x00, 0x1c, 0x14, 0x04,
			0x14, 0x03, 0x00, 0x0a, 0x24, 0x07, 0x00, 0x0b, 0x14, 0x74, 0x00, 0x05,
			0x24, 0x0f, 0x00, 0x08, 0x14, 0x40, 0x00, 0x09, 0x1c, 0x08, 0x04, 0xb8,
			0x00, 0x13, 0x11, 0x47, 0xd8, 0xc4, 0xc2, 0xf2, 0x61, 0xd3, 0xc9, 0xd5,
			0xe4, 0xe7, 0xe7, 0xf8, 0xf6, 0xf6, 0xf4, 0x00, 0x09, 0x11, 0x6d, 0xd5,
			0x96, 0x92, 0x89, 0x81, 0x00, 0x0c, 0x11, 0x5a, 0xe2, 0xd8, 0xd3, 0xf1,
			0xf0, 0xf0, 0xf5, 0xf5, 0x00, 0x4a, 0xd0, 0x01, 0x00, 0x02, 0x00, 0x44,
			0x10, 0x6d, 0x00, 0x06, 0x11, 0xa2, 0x00, 0x09, 0x00, 0x16, 0x21, 0x10,
			0xe2, 0xc1, 0xd4, 0xd7, 0xd3, 0xc5, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40,
			0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x00, 0x24, 0x11, 0xdc, 0x23, 0xc4,
			0xde, 0x5b, 0xff, 0xf1, 0x70, 0xa2, 0xdf, 0xa8, 0x4c, 0x89, 0x2d, 0xf0,
			0xd7, 0xa3, 0x8b, 0x8a, 0x1c, 0x9b, 0x62, 0x33, 0x97, 0x55, 0xc6, 0xba,
			0x9b, 0xcb, 0x01, 0x4d, 0x7e, 0xf3,
		}, //Client
		{
			0x00, 0x86, 0xd0, 0x43, 0x00, 0x01, 0x00, 0x80, 0x14, 0x43, 0x00, 0x35,
			0x11, 0x5e, 0x84, 0x82, 0xf2, 0x89, 0x95, 0xa2, 0xa3, 0xf1, 0x84, 0x82,
			0xf2, 0x81, 0x87, 0x85, 0x95, 0xa3, 0xf0, 0xf0, 0xf0, 0xf0, 0xf2, 0xf4,
			0xc6, 0xc1, 0x6c, 0xc6, 0xc5, 0xc4, 0x6c, 0xe8, 0xf0, 0xf0, 0x40, 0x40,
			0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40,
			0x40, 0x40, 0xf1, 0x00, 0x1c, 0x14, 0x04, 0x14, 0x03, 0x00, 0x0a, 0x24,
			0x07, 0x00, 0x0a, 0x14, 0x74, 0x00, 0x05, 0x24, 0x0f, 0x00, 0x08, 0x14,
			0x40, 0x00, 0x09, 0x1c, 0x08, 0x00, 0x00, 0x00, 0x13, 0x11, 0x47, 0xd8,
			0xc4, 0xc2, 0xf2, 0x61, 0xd3, 0xc9, 0xd5, 0xe4, 0xe7, 0xe7, 0xf8, 0xf6,
			0xf6, 0xf4, 0x00, 0x0c, 0x11, 0x6d, 0x84, 0x82, 0xf2, 0x89, 0x95, 0xa2,
			0xa3, 0xf1, 0x00, 0x0c, 0x11, 0x5a, 0xe2, 0xd8, 0xd3, 0xf1, 0xf0, 0xf0,
			0xf5, 0xf5, 0x00, 0x17, 0xd0, 0x03, 0x00, 0x02, 0x00, 0x11, 0x14, 0xac,
			0x00, 0x08, 0x11, 0xa2, 0x00, 0x03, 0x00, 0x05, 0x00, 0x05, 0x11, 0xa4,
			0x01,
		},
		{
			0x00, 0x26, 0xd0, 0x41, 0x00, 0x01, 0x00, 0x20, 0x10, 0x6d, 0x00, 0x06,
			0x11, 0xa2, 0x00, 0x03, 0x00, 0x16, 0x21, 0x10, 0xe2, 0xc1, 0xd4, 0xd7,
			0xd3, 0xc5, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40,
			0x40, 0x40, 0x00, 0x3e, 0xd0, 0x41, 0x00, 0x02, 0x00, 0x38, 0x10, 0x6e,
			0x00, 0x06, 0x11, 0xa2, 0x00, 0x03, 0x00, 0x16, 0x21, 0x10, 0xe2, 0xc1,
			0xd4, 0xd7, 0xd3, 0xc5, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40,
			0x40, 0x40, 0x40, 0x40, 0x00, 0x0c, 0x11, 0xa1, 0x84, 0x82, 0xf2, 0x89,
			0x95, 0xa2, 0xa3, 0xf1, 0x00, 0x0c, 0x11, 0xa0, 0x84, 0x82, 0xf2, 0x89,
			0x95, 0xa2, 0xa3, 0xf1, 0x00, 0xbc, 0xd0, 0x01, 0x00, 0x03, 0x00, 0xb6,
			0x20, 0x01, 0x00, 0x06, 0x21, 0x0f, 0x24, 0x07, 0x00, 0x20, 0x21, 0x35,
			0xf1, 0xf2, 0xf7, 0x4b, 0xf0, 0x4b, 0xf0, 0x4b, 0xf1, 0x4b, 0xf3, 0xf6,
			0xf3, 0xf9, 0xf4, 0x4b, 0xf1, 0xf6, 0xf1, 0xf0, 0xf2, 0xf2, 0xf0, 0xf9,
			0xf0, 0xf8, 0xf2, 0xf8, 0x00, 0x16, 0x21, 0x10, 0xe2, 0xc1, 0xd4, 0xd7,
			0xd3, 0xc5, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40,
			0x40, 0x40, 0x00, 0x0c, 0x11, 0x2e, 0xe2, 0xd8, 0xd3, 0xf1, 0xf0, 0xf0,
			0xf5, 0xf5, 0x00, 0x0d, 0x00, 0x2f, 0xd8, 0xe3, 0xc4, 0xe2, 0xd8, 0xd3,
			0xe7, 0xf8, 0xf6, 0x00, 0x1c, 0x00, 0x35, 0x00, 0x06, 0x11, 0x9c, 0x04,
			0xb8, 0x00, 0x06, 0x11, 0x9d, 0x04, 0xb0, 0x00, 0x06, 0x11, 0x9e, 0x04,
			0xb8, 0x00, 0x06, 0x19, 0x13, 0x04, 0xb8, 0x00, 0x3c, 0x21, 0x04, 0x37,
			0xe2, 0xd8, 0xd3, 0xf1, 0xf0, 0xf0, 0xf5, 0xf5, 0xd3, 0x89, 0x95, 0xa4,
			0xa7, 0x61, 0xe7, 0xf8, 0xf6, 0xf6, 0xf4, 0x40, 0x40, 0x40, 0x40, 0x40,
			0x40, 0x40, 0x97, 0xa8, 0xa3, 0x88, 0x96, 0x95, 0x40, 0x40, 0x40, 0x40,
			0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x84, 0x82,
			0xf2, 0x89, 0x95, 0xa2, 0xa3, 0xf1, 0x00, 0x00, 0x05, 0x21, 0x3b, 0xf1,
		},
		{
			0x00, 0x10, 0xd0, 0x43, 0x00, 0x01, 0x00, 0x0a, 0x14, 0xac, 0x00, 0x06,
			0x11, 0xa2, 0x00, 0x03, 0x00, 0x15, 0xd0, 0x42, 0x00, 0x02, 0x00, 0x0f,
			0x12, 0x19, 0x00, 0x06, 0x11, 0x49, 0x00, 0x00, 0x00, 0x05, 0x11, 0xa4,
			0x00, 0x00, 0x5d, 0xd0, 0x52, 0x00, 0x03, 0x00, 0x57, 0x22, 0x01, 0x00,
			0x06, 0x11, 0x49, 0x00, 0x00, 0x00, 0x0c, 0x11, 0x2e, 0xe2, 0xd8, 0xd3,
			0xf1, 0xf0, 0xf0, 0xf5, 0xf5, 0x00, 0x0d, 0x00, 0x2f, 0xd8, 0xe3, 0xc4,
			0xe2, 0xd8, 0xd3, 0xe7, 0xf8, 0xf6, 0x00, 0x1c, 0x00, 0x35, 0x00, 0x06,
			0x11, 0x9c, 0x04, 0xb8, 0x00, 0x06, 0x11, 0x9d, 0x04, 0xb0, 0x00, 0x06,
			0x11, 0x9e, 0x04, 0xb8, 0x00, 0x06, 0x19, 0x13, 0x04, 0xb8, 0x00, 0x06,
			0x21, 0x03, 0x01, 0x8c, 0x00, 0x06, 0x21, 0x25, 0x24, 0x35, 0x00, 0x0c,
			0x11, 0xa0, 0xc4, 0xc2, 0xf2, 0xc9, 0xd5, 0xe2, 0xe3, 0xf1, 0x00, 0x91,
			0xd0, 0x03, 0x00, 0x03, 0x00, 0x8b, 0x24, 0x08, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x30, 0x30, 0x30, 0x30, 0x30, 0x53, 0x51, 0x4c, 0x31, 0x30, 0x30,
			0x35, 0x35, 0x00, 0x01, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
			0x20, 0x20, 0x00, 0x12, 0x53, 0x41, 0x4d, 0x50, 0x4c, 0x45, 0x20, 0x20,
			0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x00, 0x38,
			0x31, 0xff, 0x31, 0x32, 0x30, 0x38, 0xff, 0x44, 0x42, 0x32, 0x49, 0x4e,
			0x53, 0x54, 0x31, 0xff, 0x53, 0x41, 0x4d, 0x50, 0x4c, 0x45, 0xff, 0x51,
			0x44, 0x42, 0x32, 0x2f, 0x4c, 0x49, 0x4e, 0x55, 0x58, 0x58, 0x38, 0x36,
			0x36, 0x34, 0xff, 0x33, 0x39, 0x36, 0xff, 0x33, 0x39, 0x36, 0xff, 0x30,
			0xff, 0x31, 0x32, 0x30, 0x38, 0xff, 0x30, 0xff, 0x00, 0x00, 0xff,
		},
		{
			0x00, 0x12, 0xd0, 0x41, 0x00, 0x01, 0x00, 0x0c, 0x10, 0x41, 0x00, 0x08,
			0x14, 0x04, 0x14, 0xcc, 0x04, 0xb8, 0x00, 0x4e, 0xd0, 0x51, 0x00, 0x02,
			0x00, 0x48, 0x20, 0x14, 0x00, 0x44, 0x21, 0x13, 0x53, 0x41, 0x4d, 0x50,
			0x4c, 0x45, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
			0x20, 0x20, 0x44, 0x42, 0x32, 0x4c, 0x49, 0x43, 0x20, 0x20, 0x20, 0x20,
			0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x53, 0x59, 0x53, 0x4c,
			0x49, 0x43, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
			0x20, 0x20, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x00, 0x01,
			0x00, 0x35, 0xd0, 0x43, 0x00, 0x02, 0x00, 0x2f, 0x24, 0x14, 0x00, 0x00,
			0x00, 0x00, 0x25, 0x53, 0x45, 0x54, 0x20, 0x43, 0x55, 0x52, 0x52, 0x45,
			0x4e, 0x54, 0x20, 0x4c, 0x4f, 0x43, 0x41, 0x4c, 0x45, 0x20, 0x4c, 0x43,
			0x5f, 0x43, 0x54, 0x59, 0x50, 0x45, 0x20, 0x3d, 0x20, 0x27, 0x65, 0x6e,
			0x5f, 0x55, 0x53, 0x27, 0xff, 0x00, 0x6c, 0xd0, 0x51, 0x00, 0x03, 0x00,
			0x66, 0x20, 0x0b, 0x00, 0x44, 0x21, 0x13, 0x53, 0x41, 0x4d, 0x50, 0x4c,
			0x45, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
			0x20, 0x44, 0x42, 0x32, 0x4c, 0x49, 0x43, 0x20, 0x20, 0x20, 0x20, 0x20,
			0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x53, 0x59, 0x53, 0x4c, 0x49,
			0x43, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
			0x20, 0x50, 0x42, 0x35, 0x4a, 0x4c, 0x46, 0x4a, 0x77, 0x00, 0x01, 0x00,
			0x05, 0x21, 0x05, 0xf1, 0x00, 0x05, 0x21, 0x11, 0xf1, 0x00, 0x08, 0x21,
			0x14, 0x00, 0x00, 0xff, 0xff, 0x00, 0x06, 0x21, 0x41, 0xff, 0xff, 0x00,
			0x06, 0x21, 0x40, 0xff, 0xff, 0x02, 0xb4, 0xd0, 0x43, 0x00, 0x03, 0x02,
			0xae, 0x24, 0x12, 0x00, 0x22, 0x00, 0x10, 0x18, 0x76, 0xd0, 0x3f, 0x00,
			0x05, 0x3f, 0x00, 0x09, 0x3f, 0x00, 0x08, 0x3f, 0x00, 0x04, 0x3f, 0x00,
			0x01, 0x3f, 0x00, 0x01, 0x3f, 0x02, 0x58, 0x06, 0x71, 0xe4, 0xd0, 0x00,
			0x01, 0x02, 0x88, 0x14, 0x7a, 0x00, 0x00, 0x00, 0x05, 0x4e, 0x6f, 0x6b,
			0x69, 0x61, 0x00, 0x00, 0x09, 0x31, 0x32, 0x37, 0x2e, 0x30, 0x2e, 0x31,
			0x2e, 0x31, 0x00, 0x00, 0x08, 0x44, 0x42, 0x32, 0x49, 0x4e, 0x53, 0x54,
			0x31, 0x00, 0x00, 0x04, 0x31, 0x30, 0x30, 0x35, 0xff, 0xff, 0x00, 0x02,
			0x58, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x0a, 0xd0, 0x01, 0x00, 0x04, 0x00, 0x04, 0x20, 0x0e,
		},
		{
			0x00, 0x12, 0xd0, 0x43, 0x00, 0x01, 0x00, 0x0c, 0x14, 0x43, 0x00, 0x08,
			0x14, 0x04, 0x14, 0xcc, 0x04, 0xb8, 0x00, 0x0b, 0xd0, 0x43, 0x00, 0x02,
			0x00, 0x05, 0x24, 0x08, 0xff, 0x00, 0x79, 0xd0, 0x43, 0x00, 0x03, 0x00,
			0x73, 0x24, 0x08, 0x00, 0xdb, 0xfc, 0xff, 0xff, 0x35, 0x31, 0x30, 0x30,
			0x32, 0x53, 0x51, 0x4c, 0x52, 0x41, 0x31, 0x34, 0x44, 0x00, 0x6d, 0x00,
			0x1a, 0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x9c, 0xff, 0xff, 0xff, 0x00, 0x00, 0x00, 0x00, 0x20, 0x20,
			0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x00, 0x12, 0x53,
			0x41, 0x4d, 0x50, 0x4c, 0x45, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
			0x20, 0x20, 0x20, 0x20, 0x20, 0x00, 0x20, 0x44, 0x42, 0x32, 0x4c, 0x49,
			0x43, 0x2e, 0x53, 0x59, 0x53, 0x4c, 0x49, 0x43, 0x20, 0x30, 0x58, 0x35,
			0x30, 0x34, 0x32, 0x33, 0x35, 0x34, 0x41, 0x34, 0x43, 0x34, 0x36, 0x34,
			0x41, 0x37, 0x37, 0x00, 0x00, 0xff, 0x00, 0x2b, 0xd0, 0x52, 0x00, 0x04,
			0x00, 0x25, 0x22, 0x0c, 0x00, 0x06, 0x11, 0x49, 0x00, 0x04, 0x00, 0x05,
			0x21, 0x15, 0x01, 0x00, 0x16, 0x21, 0x10, 0x53, 0x41, 0x4d, 0x50, 0x4c,
			0x45, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
			0x20, 0x00, 0x0b, 0xd0, 0x03, 0x00, 0x04, 0x00, 0x05, 0x24, 0x08, 0xff,
		},
		{
			0x00, 0x0a, 0xd0, 0x01, 0x00, 0x01, 0x00, 0x04, 0x20, 0x0e,
		},
		{
			0x00, 0x2b, 0xd0, 0x52, 0x00, 0x01, 0x00, 0x25, 0x22, 0x0c, 0x00, 0x06,
			0x11, 0x49, 0x00, 0x04, 0x00, 0x05, 0x21, 0x15, 0x01, 0x00, 0x16, 0x21,
			0x10, 0x53, 0x41, 0x4d, 0x50, 0x4c, 0x45, 0x20, 0x20, 0x20, 0x20, 0x20,
			0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x00, 0x0b, 0xd0, 0x03, 0x00,
			0x01, 0x00, 0x05, 0x24, 0x08, 0xff,
		},
		{
			0x00, 0x4e, 0xd0, 0x51, 0x00, 0x01, 0x00, 0x48, 0x20, 0x14,
			0x00, 0x44, 0x21, 0x13, 0x53, 0x41, 0x4d, 0x50, 0x4c, 0x45, 0x20, 0x20,
			0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x4e, 0x55,
			0x4c, 0x4c, 0x49, 0x44, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
			0x20, 0x20, 0x20, 0x20, 0x53, 0x59, 0x53, 0x53, 0x48, 0x32, 0x30, 0x30,
			0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x01, 0x01,
			0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x00, 0x01, 0x00, 0x2d, 0xd0, 0x43,
			0x00, 0x01, 0x00, 0x27, 0x24, 0x14, 0x00, 0x00, 0x00, 0x00, 0x1d, 0x53,
			0x45, 0x54, 0x20, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x20, 0x57, 0x52,
			0x4b, 0x53, 0x54, 0x4e, 0x4e, 0x41, 0x4d, 0x45, 0x20, 0x27, 0x4e, 0x6f,
			0x6b, 0x69, 0x61, 0x27, 0xff, 0x00, 0x53, 0xd0, 0x51, 0x00, 0x02, 0x00,
			0x4d, 0x20, 0x0d, 0x00, 0x44, 0x21, 0x13, 0x53, 0x41, 0x4d, 0x50, 0x4c,
			0x45, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
			0x20, 0x4e, 0x55, 0x4c, 0x4c, 0x49, 0x44, 0x20, 0x20, 0x20, 0x20, 0x20,
			0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x53, 0x59, 0x53, 0x53, 0x48,
			0x32, 0x30, 0x30, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
			0x20, 0x53, 0x59, 0x53, 0x4c, 0x56, 0x4c, 0x30, 0x31, 0x00, 0x04, 0x00,
			0x05, 0x21, 0x16, 0xf1, 0x00, 0x1a, 0xd0, 0x53, 0x00, 0x02, 0x00, 0x14,
			0x24, 0x50, 0x00, 0x00, 0x00, 0x00, 0x0a, 0x57, 0x49, 0x54, 0x48, 0x20,
			0x48, 0x4f, 0x4c, 0x44, 0x20, 0xff, 0x00, 0x57, 0xd0, 0x43, 0x00, 0x02,
			0x00, 0x51, 0x24, 0x14, 0x00, 0x00, 0x00, 0x00, 0x47, 0x73, 0x65, 0x6c,
			0x65, 0x63, 0x74, 0x20, 0x50, 0x43, 0x54, 0x46, 0x52, 0x45, 0x45, 0x2c,
			0x20, 0x43, 0x41, 0x52, 0x44, 0x2c, 0x20, 0x43, 0x48, 0x49, 0x4c, 0x44,
			0x52, 0x45, 0x4e, 0x20, 0x66, 0x72, 0x6f, 0x6d, 0x20, 0x73, 0x79, 0x73,
			0x63, 0x61, 0x74, 0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x20, 0x77,
			0x68, 0x65, 0x72, 0x65, 0x20, 0x54, 0x41, 0x42, 0x4e, 0x41, 0x4d, 0x45,
			0x3d, 0x27, 0x53, 0x41, 0x4c, 0x45, 0x53, 0x27, 0xff, 0x00, 0x72, 0xd0,
			0x01, 0x00, 0x03, 0x00, 0x6c, 0x20, 0x0c, 0x00, 0x44, 0x21, 0x13, 0x53,
			0x41, 0x4d, 0x50, 0x4c, 0x45, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
			0x20, 0x20, 0x20, 0x20, 0x20, 0x4e, 0x55, 0x4c, 0x4c, 0x49, 0x44, 0x20,
			0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x53,
			0x59, 0x53, 0x53, 0x48, 0x32, 0x30, 0x30, 0x20, 0x20, 0x20, 0x20, 0x20,
			0x20, 0x20, 0x20, 0x20, 0x20, 0x53, 0x59, 0x53, 0x4c, 0x56, 0x4c, 0x30,
			0x31, 0x00, 0x04, 0x00, 0x08, 0x21, 0x14, 0x00, 0x00, 0xff, 0xff, 0x00,
			0x06, 0x21, 0x41, 0xff, 0xff, 0x00, 0x05, 0x21, 0x5d, 0x01, 0x00, 0x05,
			0x21, 0x4b, 0xf1, 0x00, 0x0c, 0x21, 0x36, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0xff, 0xff,
		},
		{
			0x00, 0x0b, 0xd0, 0x43, 0x00, 0x01, 0x00, 0x05, 0x24, 0x08, 0xff, 0x00,
			0xf3, 0xd0, 0x43, 0x00, 0x02, 0x00, 0xed, 0x24, 0x11, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x30, 0x30, 0x30, 0x30, 0x30, 0x53, 0x51, 0x4c, 0x31, 0x30,
			0x30, 0x35, 0x35, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x04, 0x00, 0x00, 0x00, 0x3b, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
			0x20, 0x20, 0x20, 0x00, 0x12, 0x53, 0x41, 0x4d, 0x50, 0x4c, 0x45, 0x20,
			0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x00,
			0x00, 0x00, 0x00, 0xff, 0xff, 0x03, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xf4, 0x01, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x07, 0x50, 0x43, 0x54, 0x46, 0x52, 0x45, 0x45, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0xff, 0xff, 0x00, 0x00,
			0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xec, 0x01,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x04, 0x43, 0x41, 0x52, 0x44, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0xff, 0xff, 0x00, 0x00,
			0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xf5, 0x01,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x08, 0x43, 0x48, 0x49, 0x4c, 0x44, 0x52, 0x45,
			0x4e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff,
			0xff, 0xff, 0x00, 0x37, 0xd0, 0x52, 0x00, 0x03, 0x00, 0x31, 0x22, 0x05,
			0x00, 0x06, 0x11, 0x49, 0x00, 0x00, 0x00, 0x06, 0x21, 0x02, 0x24, 0x17,
			0x00, 0x05, 0x21, 0x1f, 0xf1, 0x00, 0x05, 0x21, 0x50, 0x01, 0x00, 0x0c,
			0x21, 0x5b, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x05,
			0x21, 0x4b, 0xf1, 0x00, 0x06, 0x24, 0x60, 0x24, 0x42, 0x00, 0x25, 0xd0,
			0x53, 0x00, 0x03, 0x00, 0x1f, 0x24, 0x1a, 0x0c, 0x76, 0xd0, 0x04, 0x00,
			0x02, 0x16, 0x00, 0x08, 0x05, 0x00, 0x02, 0x09, 0x71, 0xe0, 0x54, 0x00,
			0x01, 0xd0, 0x00, 0x01, 0x06, 0x71, 0xf0, 0xe0, 0x00, 0x00, 0x00, 0x19,
			0xd0, 0x53, 0x00, 0x03, 0x00, 0x13, 0x24, 0x1b, 0xff, 0x00, 0xff, 0xff,
			0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x00, 0x00, 0x00,
			0x26, 0xd0, 0x52, 0x00, 0x03, 0x00, 0x20, 0x22, 0x0b, 0x00, 0x06, 0x11,
			0x49, 0x00, 0x04, 0x00, 0x16, 0x21, 0x10, 0x53, 0x41, 0x4d, 0x50, 0x4c,
			0x45, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
			0x20, 0x00, 0x59, 0xd0, 0x03, 0x00, 0x03, 0x00, 0x53, 0x24, 0x08, 0x00,
			0x64, 0x00, 0x00, 0x00, 0x30, 0x32, 0x30, 0x30, 0x30, 0x53, 0x51, 0x4c,
			0x52, 0x49, 0x30, 0x31, 0x46, 0x00, 0x01, 0x00, 0x04, 0x80, 0x01, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
			0x20, 0x20, 0x20, 0x20, 0x20, 0x00, 0x12, 0x53, 0x41, 0x4d, 0x50, 0x4c,
			0x45, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
			0x20, 0x00, 0x00, 0x00, 0x00, 0xff,
		},
		{
			0x00, 0x0a, 0xd0, 0x01, 0x00, 0x01, 0x00, 0x04, 0x20, 0x0e,
		},
		{
			0x00, 0x2b, 0xd0, 0x52, 0x00, 0x01, 0x00, 0x25, 0x22, 0x0c, 0x00, 0x06,
			0x11, 0x49, 0x00, 0x04, 0x00, 0x05, 0x21, 0x15, 0x01, 0x00, 0x16, 0x21,
			0x10, 0x53, 0x41, 0x4d, 0x50, 0x4c, 0x45, 0x20, 0x20, 0x20, 0x20, 0x20,
			0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x00, 0x0b, 0xd0, 0x03, 0x00,
			0x01, 0x00, 0x05, 0x24, 0x08, 0xff,
		},
		{
			0x00, 0x0a, 0xd0, 0x05, 0x00, 0x01, 0x00, 0x04, 0xc0, 0x04,
		},
	}
	if db2Count < len(sampleIO) {
		buf = sampleIO[db2Count]
		db2Count++
	} else {
		err = errors.New("EOF")
	}
	return
}

func TestDB2(t *testing.T) {
	d := new(dbms.DB2)
	port := d.DefaultPort()
	if d.DefaultPort() != 50000 {
		t.Error("Expected 50000, got ", port)
	}
	err := d.SetCertificate("", "")
	if err == nil {
		t.Error("Expected error")
	}
	d.SetReader(db2DummyReader)
	var s mock.ConnMock
	d.SetSockets(s, s)
	err = d.Handler()
	if err != nil {
		t.Error("Got error", err)
	}
	d.Close()
}