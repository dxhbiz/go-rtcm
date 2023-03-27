package rtcm

import (
	"fmt"
	"os"
	"testing"
)

var (
	msgs = []int{
		1001, 1002, 1003, 1004, 1005, 1006, 1007, 1008, 1009,
		1010, 1011, 1012, 1013, 1019,
		1020, 1029,
		1030, 1031, 1032, 1033,
		1042, 1044, 1045,
		1057, 1058, 1059,
		1060, 1063, 1064, 1065, 1066,
		1071, 1072, 1073, 1074, 1075, 1076, 1077,
		1081, 1082, 1083, 1084, 1085, 1086, 1087,
		1091, 1092, 1093, 1094, 1095, 1096, 1097,
		1104, 1107,
		1111, 1112, 1113, 1114, 1115, 1116, 1117,
		1121, 1122, 1123, 1124, 1125, 1126, 1127,
		1230,
	}
)

func TestCdecode(t *testing.T) {
	for _, msgType := range msgs {
		msgType := msgType

		file := fmt.Sprintf("%d_frame.bin", msgType)
		data, err := os.ReadFile("./data/" + file)
		if err != nil {
			t.Error(err)
		}

		rtcm, err := Cdecode(data)
		if err != nil {
			t.Error(err)
		}

		if rtcm.crc != 0 {
			t.Fatalf("%s crc error", file)
		}

		if int(rtcm.mtype) != msgType {
			t.Fatalf("expected to get %d, but got %d", msgType, rtcm.mtype)
		}
	}
}

func TestDecode(t *testing.T) {
	for _, msgType := range msgs {
		msgType := msgType

		file := fmt.Sprintf("%d_frame.bin", msgType)
		data, err := os.ReadFile("./data/" + file)
		if err != nil {
			t.Error(err)
		}

		rtcm, err := Decode(data)
		if err != nil {
			t.Error(err)
		}

		if rtcm.Crc != 0 {
			t.Fatalf("%s crc error", file)
		}

		if int(rtcm.Type) != msgType {
			t.Fatalf("expected to get %d, but got %d", msgType, rtcm.Type)
		}
	}
}

func TestCheckCrc(t *testing.T) {
	for _, msgType := range msgs {
		msgType := msgType

		file := fmt.Sprintf("%d_frame.bin", msgType)
		data, err := os.ReadFile("./data/" + file)
		if err != nil {
			t.Error(err)
		}

		crc := CheckCrc(data)

		if crc != true {
			t.Fatalf("expected to get %T, but got %T", true, crc)
		}
	}
}
