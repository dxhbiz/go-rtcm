package rtcm

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func Test1001(t *testing.T) {
	expected := 1001
	file := fmt.Sprintf("%d_frame.bin", expected)
	data, err := ioutil.ReadFile("./data/" + file)
	if err != nil {
		t.Error(err)
	}

	rtcm, err := Decode(data)
	if err != nil {
		t.Error(err)
	}

	if rtcm.crc != 0 {
		t.Fatalf("%s crc error", file)
	}

	if int(rtcm.mtype) != expected {
		t.Fatalf("expected to get %d, but got %d", expected, rtcm.mtype)
	}
}

func Test1002(t *testing.T) {
	expected := 1002
	file := fmt.Sprintf("%d_frame.bin", expected)
	data, err := ioutil.ReadFile("./data/" + file)
	if err != nil {
		t.Error(err)
	}

	rtcm, err := Decode(data)
	if err != nil {
		t.Error(err)
	}

	if rtcm.crc != 0 {
		t.Fatalf("%s crc error", file)
	}

	if int(rtcm.mtype) != expected {
		t.Fatalf("expected to get %d, but got %d", expected, rtcm.mtype)
	}
}

func Test1003(t *testing.T) {
	expected := 1003
	file := fmt.Sprintf("%d_frame.bin", expected)
	data, err := ioutil.ReadFile("./data/" + file)
	if err != nil {
		t.Error(err)
	}

	rtcm, err := Decode(data)
	if err != nil {
		t.Error(err)
	}

	if rtcm.crc != 0 {
		t.Fatalf("%s crc error", file)
	}

	if int(rtcm.mtype) != expected {
		t.Fatalf("expected to get %d, but got %d", expected, rtcm.mtype)
	}
}

func Test1004(t *testing.T) {
	expected := 1004
	file := fmt.Sprintf("%d_frame.bin", expected)
	data, err := ioutil.ReadFile("./data/" + file)
	if err != nil {
		t.Error(err)
	}

	rtcm, err := Decode(data)
	if err != nil {
		t.Error(err)
	}

	if rtcm.crc != 0 {
		t.Fatalf("%s crc error", file)
	}

	if int(rtcm.mtype) != expected {
		t.Fatalf("expected to get %d, but got %d", expected, rtcm.mtype)
	}
}

func Test1005(t *testing.T) {
	expected := 1005
	file := fmt.Sprintf("%d_frame.bin", expected)
	data, err := ioutil.ReadFile("./data/" + file)
	if err != nil {
		t.Error(err)
	}

	rtcm, err := Decode(data)
	if err != nil {
		t.Error(err)
	}

	if rtcm.crc != 0 {
		t.Fatalf("%s crc error", file)
	}

	if int(rtcm.mtype) != expected {
		t.Fatalf("expected to get %d, but got %d", expected, rtcm.mtype)
	}
}

func Test1006(t *testing.T) {
	expected := 1006
	file := fmt.Sprintf("%d_frame.bin", expected)
	data, err := ioutil.ReadFile("./data/" + file)
	if err != nil {
		t.Error(err)
	}

	rtcm, err := Decode(data)
	if err != nil {
		t.Error(err)
	}

	if rtcm.crc != 0 {
		t.Fatalf("%s crc error", file)
	}

	if int(rtcm.mtype) != expected {
		t.Fatalf("expected to get %d, but got %d", expected, rtcm.mtype)
	}
}

func Test1007(t *testing.T) {
	expected := 1007
	file := fmt.Sprintf("%d_frame.bin", expected)
	data, err := ioutil.ReadFile("./data/" + file)
	if err != nil {
		t.Error(err)
	}

	rtcm, err := Decode(data)
	if err != nil {
		t.Error(err)
	}

	if rtcm.crc != 0 {
		t.Fatalf("%s crc error", file)
	}

	if int(rtcm.mtype) != expected {
		t.Fatalf("expected to get %d, but got %d", expected, rtcm.mtype)
	}
}

func Test1008(t *testing.T) {
	expected := 1008
	file := fmt.Sprintf("%d_frame.bin", expected)
	data, err := ioutil.ReadFile("./data/" + file)
	if err != nil {
		t.Error(err)
	}

	rtcm, err := Decode(data)
	if err != nil {
		t.Error(err)
	}

	if rtcm.crc != 0 {
		t.Fatalf("%s crc error", file)
	}

	if int(rtcm.mtype) != expected {
		t.Fatalf("expected to get %d, but got %d", expected, rtcm.mtype)
	}
}

func Test1009(t *testing.T) {
	expected := 1009
	file := fmt.Sprintf("%d_frame.bin", expected)
	data, err := ioutil.ReadFile("./data/" + file)
	if err != nil {
		t.Error(err)
	}

	rtcm, err := Decode(data)
	if err != nil {
		t.Error(err)
	}

	if rtcm.crc != 0 {
		t.Fatalf("%s crc error", file)
	}

	if int(rtcm.mtype) != expected {
		t.Fatalf("expected to get %d, but got %d", expected, rtcm.mtype)
	}
}

func Test1010(t *testing.T) {
	expected := 1009
	file := fmt.Sprintf("%d_frame.bin", expected)
	data, err := ioutil.ReadFile("./data/" + file)
	if err != nil {
		t.Error(err)
	}

	rtcm, err := Decode(data)
	if err != nil {
		t.Error(err)
	}

	if rtcm.crc != 0 {
		t.Fatalf("%s crc error", file)
	}

	if int(rtcm.mtype) != expected {
		t.Fatalf("expected to get %d, but got %d", expected, rtcm.mtype)
	}
}

func Test1011(t *testing.T) {
	expected := 1011
	file := fmt.Sprintf("%d_frame.bin", expected)
	data, err := ioutil.ReadFile("./data/" + file)
	if err != nil {
		t.Error(err)
	}

	rtcm, err := Decode(data)
	if err != nil {
		t.Error(err)
	}

	if rtcm.crc != 0 {
		t.Fatalf("%s crc error", file)
	}

	if int(rtcm.mtype) != expected {
		t.Fatalf("expected to get %d, but got %d", expected, rtcm.mtype)
	}
}

func Test1012(t *testing.T) {
	expected := 1012
	file := fmt.Sprintf("%d_frame.bin", expected)
	data, err := ioutil.ReadFile("./data/" + file)
	if err != nil {
		t.Error(err)
	}

	rtcm, err := Decode(data)
	if err != nil {
		t.Error(err)
	}

	if rtcm.crc != 0 {
		t.Fatalf("%s crc error", file)
	}

	if int(rtcm.mtype) != expected {
		t.Fatalf("expected to get %d, but got %d", expected, rtcm.mtype)
	}
}

func Test1013(t *testing.T) {
	expected := 1013
	file := fmt.Sprintf("%d_frame.bin", expected)
	data, err := ioutil.ReadFile("./data/" + file)
	if err != nil {
		t.Error(err)
	}

	rtcm, err := Decode(data)
	if err != nil {
		t.Error(err)
	}

	if rtcm.crc != 0 {
		t.Fatalf("%s crc error", file)
	}

	if int(rtcm.mtype) != expected {
		t.Fatalf("expected to get %d, but got %d", expected, rtcm.mtype)
	}
}

func Test1019(t *testing.T) {
	expected := 1019
	file := fmt.Sprintf("%d_frame.bin", expected)
	data, err := ioutil.ReadFile("./data/" + file)
	if err != nil {
		t.Error(err)
	}

	rtcm, err := Decode(data)
	if err != nil {
		t.Error(err)
	}

	if rtcm.crc != 0 {
		t.Fatalf("%s crc error", file)
	}

	if int(rtcm.mtype) != expected {
		t.Fatalf("expected to get %d, but got %d", expected, rtcm.mtype)
	}
}

func Test1020(t *testing.T) {
	expected := 1020
	file := fmt.Sprintf("%d_frame.bin", expected)
	data, err := ioutil.ReadFile("./data/" + file)
	if err != nil {
		t.Error(err)
	}

	rtcm, err := Decode(data)
	if err != nil {
		t.Error(err)
	}

	if rtcm.crc != 0 {
		t.Fatalf("%s crc error", file)
	}

	if int(rtcm.mtype) != expected {
		t.Fatalf("expected to get %d, but got %d", expected, rtcm.mtype)
	}
}

func Test1029(t *testing.T) {
	expected := 1029
	file := fmt.Sprintf("%d_frame.bin", expected)
	data, err := ioutil.ReadFile("./data/" + file)
	if err != nil {
		t.Error(err)
	}

	rtcm, err := Decode(data)
	if err != nil {
		t.Error(err)
	}

	if rtcm.crc != 0 {
		t.Fatalf("%s crc error", file)
	}

	if int(rtcm.mtype) != expected {
		t.Fatalf("expected to get %d, but got %d", expected, rtcm.mtype)
	}
}

func Test1030(t *testing.T) {
	expected := 1030
	file := fmt.Sprintf("%d_frame.bin", expected)
	data, err := ioutil.ReadFile("./data/" + file)
	if err != nil {
		t.Error(err)
	}

	rtcm, err := Decode(data)
	if err != nil {
		t.Error(err)
	}

	if rtcm.crc != 0 {
		t.Fatalf("%s crc error", file)
	}

	if int(rtcm.mtype) != expected {
		t.Fatalf("expected to get %d, but got %d", expected, rtcm.mtype)
	}
}

func Test1031(t *testing.T) {
	expected := 1031
	file := fmt.Sprintf("%d_frame.bin", expected)
	data, err := ioutil.ReadFile("./data/" + file)
	if err != nil {
		t.Error(err)
	}

	rtcm, err := Decode(data)
	if err != nil {
		t.Error(err)
	}

	if rtcm.crc != 0 {
		t.Fatalf("%s crc error", file)
	}

	if int(rtcm.mtype) != expected {
		t.Fatalf("expected to get %d, but got %d", expected, rtcm.mtype)
	}
}

func Test1032(t *testing.T) {
	expected := 1032
	file := fmt.Sprintf("%d_frame.bin", expected)
	data, err := ioutil.ReadFile("./data/" + file)
	if err != nil {
		t.Error(err)
	}

	rtcm, err := Decode(data)
	if err != nil {
		t.Error(err)
	}

	if rtcm.crc != 0 {
		t.Fatalf("%s crc error", file)
	}

	if int(rtcm.mtype) != expected {
		t.Fatalf("expected to get %d, but got %d", expected, rtcm.mtype)
	}
}

func Test1033(t *testing.T) {
	expected := 1033
	file := fmt.Sprintf("%d_frame.bin", expected)
	data, err := ioutil.ReadFile("./data/" + file)
	if err != nil {
		t.Error(err)
	}

	rtcm, err := Decode(data)
	if err != nil {
		t.Error(err)
	}

	if rtcm.crc != 0 {
		t.Fatalf("%s crc error", file)
	}

	if int(rtcm.mtype) != expected {
		t.Fatalf("expected to get %d, but got %d", expected, rtcm.mtype)
	}
}

func Test1042(t *testing.T) {
	expected := 1042
	file := fmt.Sprintf("%d_frame.bin", expected)
	data, err := ioutil.ReadFile("./data/" + file)
	if err != nil {
		t.Error(err)
	}

	rtcm, err := Decode(data)
	if err != nil {
		t.Error(err)
	}

	if rtcm.crc != 0 {
		t.Fatalf("%s crc error", file)
	}

	if int(rtcm.mtype) != expected {
		t.Fatalf("expected to get %d, but got %d", expected, rtcm.mtype)
	}
}

func Test1044(t *testing.T) {
	expected := 1044
	file := fmt.Sprintf("%d_frame.bin", expected)
	data, err := ioutil.ReadFile("./data/" + file)
	if err != nil {
		t.Error(err)
	}

	rtcm, err := Decode(data)
	if err != nil {
		t.Error(err)
	}

	if rtcm.crc != 0 {
		t.Fatalf("%s crc error", file)
	}

	if int(rtcm.mtype) != expected {
		t.Fatalf("expected to get %d, but got %d", expected, rtcm.mtype)
	}
}

func Test1045(t *testing.T) {
	expected := 1045
	file := fmt.Sprintf("%d_frame.bin", expected)
	data, err := ioutil.ReadFile("./data/" + file)
	if err != nil {
		t.Error(err)
	}

	rtcm, err := Decode(data)
	if err != nil {
		t.Error(err)
	}

	if rtcm.crc != 0 {
		t.Fatalf("%s crc error", file)
	}

	if int(rtcm.mtype) != expected {
		t.Fatalf("expected to get %d, but got %d", expected, rtcm.mtype)
	}
}

func Test1057(t *testing.T) {
	expected := 1057
	file := fmt.Sprintf("%d_frame.bin", expected)
	data, err := ioutil.ReadFile("./data/" + file)
	if err != nil {
		t.Error(err)
	}

	rtcm, err := Decode(data)
	if err != nil {
		t.Error(err)
	}

	if rtcm.crc != 0 {
		t.Fatalf("%s crc error", file)
	}

	if int(rtcm.mtype) != expected {
		t.Fatalf("expected to get %d, but got %d", expected, rtcm.mtype)
	}
}

func Test1058(t *testing.T) {
	expected := 1058
	file := fmt.Sprintf("%d_frame.bin", expected)
	data, err := ioutil.ReadFile("./data/" + file)
	if err != nil {
		t.Error(err)
	}

	rtcm, err := Decode(data)
	if err != nil {
		t.Error(err)
	}

	if rtcm.crc != 0 {
		t.Fatalf("%s crc error", file)
	}

	if int(rtcm.mtype) != expected {
		t.Fatalf("expected to get %d, but got %d", expected, rtcm.mtype)
	}
}

func Test1059(t *testing.T) {
	expected := 1059
	file := fmt.Sprintf("%d_frame.bin", expected)
	data, err := ioutil.ReadFile("./data/" + file)
	if err != nil {
		t.Error(err)
	}

	rtcm, err := Decode(data)
	if err != nil {
		t.Error(err)
	}

	if rtcm.crc != 0 {
		t.Fatalf("%s crc error", file)
	}

	if int(rtcm.mtype) != expected {
		t.Fatalf("expected to get %d, but got %d", expected, rtcm.mtype)
	}
}

func Test1060(t *testing.T) {
	expected := 1060
	file := fmt.Sprintf("%d_frame.bin", expected)
	data, err := ioutil.ReadFile("./data/" + file)
	if err != nil {
		t.Error(err)
	}

	rtcm, err := Decode(data)
	if err != nil {
		t.Error(err)
	}

	if rtcm.crc != 0 {
		t.Fatalf("%s crc error", file)
	}

	if int(rtcm.mtype) != expected {
		t.Fatalf("expected to get %d, but got %d", expected, rtcm.mtype)
	}
}

func Test1063(t *testing.T) {
	expected := 1063
	file := fmt.Sprintf("%d_frame.bin", expected)
	data, err := ioutil.ReadFile("./data/" + file)
	if err != nil {
		t.Error(err)
	}

	rtcm, err := Decode(data)
	if err != nil {
		t.Error(err)
	}

	if rtcm.crc != 0 {
		t.Fatalf("%s crc error", file)
	}

	if int(rtcm.mtype) != expected {
		t.Fatalf("expected to get %d, but got %d", expected, rtcm.mtype)
	}
}

func Test1064(t *testing.T) {
	expected := 1064
	file := fmt.Sprintf("%d_frame.bin", expected)
	data, err := ioutil.ReadFile("./data/" + file)
	if err != nil {
		t.Error(err)
	}

	rtcm, err := Decode(data)
	if err != nil {
		t.Error(err)
	}

	if rtcm.crc != 0 {
		t.Fatalf("%s crc error", file)
	}

	if int(rtcm.mtype) != expected {
		t.Fatalf("expected to get %d, but got %d", expected, rtcm.mtype)
	}
}

func Test1065(t *testing.T) {
	expected := 1065
	file := fmt.Sprintf("%d_frame.bin", expected)
	data, err := ioutil.ReadFile("./data/" + file)
	if err != nil {
		t.Error(err)
	}

	rtcm, err := Decode(data)
	if err != nil {
		t.Error(err)
	}

	if rtcm.crc != 0 {
		t.Fatalf("%s crc error", file)
	}

	if int(rtcm.mtype) != expected {
		t.Fatalf("expected to get %d, but got %d", expected, rtcm.mtype)
	}
}

func Test1066(t *testing.T) {
	expected := 1066
	file := fmt.Sprintf("%d_frame.bin", expected)
	data, err := ioutil.ReadFile("./data/" + file)
	if err != nil {
		t.Error(err)
	}

	rtcm, err := Decode(data)
	if err != nil {
		t.Error(err)
	}

	if rtcm.crc != 0 {
		t.Fatalf("%s crc error", file)
	}

	if int(rtcm.mtype) != expected {
		t.Fatalf("expected to get %d, but got %d", expected, rtcm.mtype)
	}
}

func Test1071(t *testing.T) {
	expected := 1071
	file := fmt.Sprintf("%d_frame.bin", expected)
	data, err := ioutil.ReadFile("./data/" + file)
	if err != nil {
		t.Error(err)
	}

	rtcm, err := Decode(data)
	if err != nil {
		t.Error(err)
	}

	if rtcm.crc != 0 {
		t.Fatalf("%s crc error", file)
	}

	if int(rtcm.mtype) != expected {
		t.Fatalf("expected to get %d, but got %d", expected, rtcm.mtype)
	}
}

func Test1072(t *testing.T) {
	expected := 1072
	file := fmt.Sprintf("%d_frame.bin", expected)
	data, err := ioutil.ReadFile("./data/" + file)
	if err != nil {
		t.Error(err)
	}

	rtcm, err := Decode(data)
	if err != nil {
		t.Error(err)
	}

	if rtcm.crc != 0 {
		t.Fatalf("%s crc error", file)
	}

	if int(rtcm.mtype) != expected {
		t.Fatalf("expected to get %d, but got %d", expected, rtcm.mtype)
	}
}

func Test1073(t *testing.T) {
	expected := 1073
	file := fmt.Sprintf("%d_frame.bin", expected)
	data, err := ioutil.ReadFile("./data/" + file)
	if err != nil {
		t.Error(err)
	}

	rtcm, err := Decode(data)
	if err != nil {
		t.Error(err)
	}

	if rtcm.crc != 0 {
		t.Fatalf("%s crc error", file)
	}

	if int(rtcm.mtype) != expected {
		t.Fatalf("expected to get %d, but got %d", expected, rtcm.mtype)
	}
}

func Test1074(t *testing.T) {
	expected := 1074
	file := fmt.Sprintf("%d_frame.bin", expected)
	data, err := ioutil.ReadFile("./data/" + file)
	if err != nil {
		t.Error(err)
	}

	rtcm, err := Decode(data)
	if err != nil {
		t.Error(err)
	}

	if rtcm.crc != 0 {
		t.Fatalf("%s crc error", file)
	}

	if int(rtcm.mtype) != expected {
		t.Fatalf("expected to get %d, but got %d", expected, rtcm.mtype)
	}
}

func Test1075(t *testing.T) {
	expected := 1075
	file := fmt.Sprintf("%d_frame.bin", expected)
	data, err := ioutil.ReadFile("./data/" + file)
	if err != nil {
		t.Error(err)
	}

	rtcm, err := Decode(data)
	if err != nil {
		t.Error(err)
	}

	if rtcm.crc != 0 {
		t.Fatalf("%s crc error", file)
	}

	if int(rtcm.mtype) != expected {
		t.Fatalf("expected to get %d, but got %d", expected, rtcm.mtype)
	}
}

func Test1076(t *testing.T) {
	expected := 1076
	file := fmt.Sprintf("%d_frame.bin", expected)
	data, err := ioutil.ReadFile("./data/" + file)
	if err != nil {
		t.Error(err)
	}

	rtcm, err := Decode(data)
	if err != nil {
		t.Error(err)
	}

	if rtcm.crc != 0 {
		t.Fatalf("%s crc error", file)
	}

	if int(rtcm.mtype) != expected {
		t.Fatalf("expected to get %d, but got %d", expected, rtcm.mtype)
	}
}

func Test1077(t *testing.T) {
	expected := 1077
	file := fmt.Sprintf("%d_frame.bin", expected)
	data, err := ioutil.ReadFile("./data/" + file)
	if err != nil {
		t.Error(err)
	}

	rtcm, err := Decode(data)
	if err != nil {
		t.Error(err)
	}

	if rtcm.crc != 0 {
		t.Fatalf("%s crc error", file)
	}

	if int(rtcm.mtype) != expected {
		t.Fatalf("expected to get %d, but got %d", expected, rtcm.mtype)
	}
}

func Test1081(t *testing.T) {
	expected := 1081
	file := fmt.Sprintf("%d_frame.bin", expected)
	data, err := ioutil.ReadFile("./data/" + file)
	if err != nil {
		t.Error(err)
	}

	rtcm, err := Decode(data)
	if err != nil {
		t.Error(err)
	}

	if rtcm.crc != 0 {
		t.Fatalf("%s crc error", file)
	}

	if int(rtcm.mtype) != expected {
		t.Fatalf("expected to get %d, but got %d", expected, rtcm.mtype)
	}
}

func Test1082(t *testing.T) {
	expected := 1082
	file := fmt.Sprintf("%d_frame.bin", expected)
	data, err := ioutil.ReadFile("./data/" + file)
	if err != nil {
		t.Error(err)
	}

	rtcm, err := Decode(data)
	if err != nil {
		t.Error(err)
	}

	if rtcm.crc != 0 {
		t.Fatalf("%s crc error", file)
	}

	if int(rtcm.mtype) != expected {
		t.Fatalf("expected to get %d, but got %d", expected, rtcm.mtype)
	}
}

func Test1083(t *testing.T) {
	expected := 1083
	file := fmt.Sprintf("%d_frame.bin", expected)
	data, err := ioutil.ReadFile("./data/" + file)
	if err != nil {
		t.Error(err)
	}

	rtcm, err := Decode(data)
	if err != nil {
		t.Error(err)
	}

	if rtcm.crc != 0 {
		t.Fatalf("%s crc error", file)
	}

	if int(rtcm.mtype) != expected {
		t.Fatalf("expected to get %d, but got %d", expected, rtcm.mtype)
	}
}

func Test1084(t *testing.T) {
	expected := 1084
	file := fmt.Sprintf("%d_frame.bin", expected)
	data, err := ioutil.ReadFile("./data/" + file)
	if err != nil {
		t.Error(err)
	}

	rtcm, err := Decode(data)
	if err != nil {
		t.Error(err)
	}

	if rtcm.crc != 0 {
		t.Fatalf("%s crc error", file)
	}

	if int(rtcm.mtype) != expected {
		t.Fatalf("expected to get %d, but got %d", expected, rtcm.mtype)
	}
}

func Test1085(t *testing.T) {
	expected := 1085
	file := fmt.Sprintf("%d_frame.bin", expected)
	data, err := ioutil.ReadFile("./data/" + file)
	if err != nil {
		t.Error(err)
	}

	rtcm, err := Decode(data)
	if err != nil {
		t.Error(err)
	}

	if rtcm.crc != 0 {
		t.Fatalf("%s crc error", file)
	}

	if int(rtcm.mtype) != expected {
		t.Fatalf("expected to get %d, but got %d", expected, rtcm.mtype)
	}
}

func Test1086(t *testing.T) {
	expected := 1086
	file := fmt.Sprintf("%d_frame.bin", expected)
	data, err := ioutil.ReadFile("./data/" + file)
	if err != nil {
		t.Error(err)
	}

	rtcm, err := Decode(data)
	if err != nil {
		t.Error(err)
	}

	if rtcm.crc != 0 {
		t.Fatalf("%s crc error", file)
	}

	if int(rtcm.mtype) != expected {
		t.Fatalf("expected to get %d, but got %d", expected, rtcm.mtype)
	}
}

func Test1087(t *testing.T) {
	expected := 1087
	file := fmt.Sprintf("%d_frame.bin", expected)
	data, err := ioutil.ReadFile("./data/" + file)
	if err != nil {
		t.Error(err)
	}

	rtcm, err := Decode(data)
	if err != nil {
		t.Error(err)
	}

	if rtcm.crc != 0 {
		t.Fatalf("%s crc error", file)
	}

	if int(rtcm.mtype) != expected {
		t.Fatalf("expected to get %d, but got %d", expected, rtcm.mtype)
	}
}

func Test1091(t *testing.T) {
	expected := 1091
	file := fmt.Sprintf("%d_frame.bin", expected)
	data, err := ioutil.ReadFile("./data/" + file)
	if err != nil {
		t.Error(err)
	}

	rtcm, err := Decode(data)
	if err != nil {
		t.Error(err)
	}

	if rtcm.crc != 0 {
		t.Fatalf("%s crc error", file)
	}

	if int(rtcm.mtype) != expected {
		t.Fatalf("expected to get %d, but got %d", expected, rtcm.mtype)
	}
}

func Test1092(t *testing.T) {
	expected := 1092
	file := fmt.Sprintf("%d_frame.bin", expected)
	data, err := ioutil.ReadFile("./data/" + file)
	if err != nil {
		t.Error(err)
	}

	rtcm, err := Decode(data)
	if err != nil {
		t.Error(err)
	}

	if rtcm.crc != 0 {
		t.Fatalf("%s crc error", file)
	}

	if int(rtcm.mtype) != expected {
		t.Fatalf("expected to get %d, but got %d", expected, rtcm.mtype)
	}
}

func Test1093(t *testing.T) {
	expected := 1093
	file := fmt.Sprintf("%d_frame.bin", expected)
	data, err := ioutil.ReadFile("./data/" + file)
	if err != nil {
		t.Error(err)
	}

	rtcm, err := Decode(data)
	if err != nil {
		t.Error(err)
	}

	if rtcm.crc != 0 {
		t.Fatalf("%s crc error", file)
	}

	if int(rtcm.mtype) != expected {
		t.Fatalf("expected to get %d, but got %d", expected, rtcm.mtype)
	}
}

func Test1094(t *testing.T) {
	expected := 1094
	file := fmt.Sprintf("%d_frame.bin", expected)
	data, err := ioutil.ReadFile("./data/" + file)
	if err != nil {
		t.Error(err)
	}

	rtcm, err := Decode(data)
	if err != nil {
		t.Error(err)
	}

	if rtcm.crc != 0 {
		t.Fatalf("%s crc error", file)
	}

	if int(rtcm.mtype) != expected {
		t.Fatalf("expected to get %d, but got %d", expected, rtcm.mtype)
	}
}

func Test1095(t *testing.T) {
	expected := 1095
	file := fmt.Sprintf("%d_frame.bin", expected)
	data, err := ioutil.ReadFile("./data/" + file)
	if err != nil {
		t.Error(err)
	}

	rtcm, err := Decode(data)
	if err != nil {
		t.Error(err)
	}

	if rtcm.crc != 0 {
		t.Fatalf("%s crc error", file)
	}

	if int(rtcm.mtype) != expected {
		t.Fatalf("expected to get %d, but got %d", expected, rtcm.mtype)
	}
}

func Test1096(t *testing.T) {
	expected := 1096
	file := fmt.Sprintf("%d_frame.bin", expected)
	data, err := ioutil.ReadFile("./data/" + file)
	if err != nil {
		t.Error(err)
	}

	rtcm, err := Decode(data)
	if err != nil {
		t.Error(err)
	}

	if rtcm.crc != 0 {
		t.Fatalf("%s crc error", file)
	}

	if int(rtcm.mtype) != expected {
		t.Fatalf("expected to get %d, but got %d", expected, rtcm.mtype)
	}
}

func Test1097(t *testing.T) {
	expected := 1097
	file := fmt.Sprintf("%d_frame.bin", expected)
	data, err := ioutil.ReadFile("./data/" + file)
	if err != nil {
		t.Error(err)
	}

	rtcm, err := Decode(data)
	if err != nil {
		t.Error(err)
	}

	if rtcm.crc != 0 {
		t.Fatalf("%s crc error", file)
	}

	if int(rtcm.mtype) != expected {
		t.Fatalf("expected to get %d, but got %d", expected, rtcm.mtype)
	}
}

func Test1104(t *testing.T) {
	expected := 1104
	file := fmt.Sprintf("%d_frame.bin", expected)
	data, err := ioutil.ReadFile("./data/" + file)
	if err != nil {
		t.Error(err)
	}

	rtcm, err := Decode(data)
	if err != nil {
		t.Error(err)
	}

	if rtcm.crc != 0 {
		t.Fatalf("%s crc error", file)
	}

	if int(rtcm.mtype) != expected {
		t.Fatalf("expected to get %d, but got %d", expected, rtcm.mtype)
	}
}

func Test1107(t *testing.T) {
	expected := 1107
	file := fmt.Sprintf("%d_frame.bin", expected)
	data, err := ioutil.ReadFile("./data/" + file)
	if err != nil {
		t.Error(err)
	}

	rtcm, err := Decode(data)
	if err != nil {
		t.Error(err)
	}

	if rtcm.crc != 0 {
		t.Fatalf("%s crc error", file)
	}

	if int(rtcm.mtype) != expected {
		t.Fatalf("expected to get %d, but got %d", expected, rtcm.mtype)
	}
}

func Test1111(t *testing.T) {
	expected := 1111
	file := fmt.Sprintf("%d_frame.bin", expected)
	data, err := ioutil.ReadFile("./data/" + file)
	if err != nil {
		t.Error(err)
	}

	rtcm, err := Decode(data)
	if err != nil {
		t.Error(err)
	}

	if rtcm.crc != 0 {
		t.Fatalf("%s crc error", file)
	}

	if int(rtcm.mtype) != expected {
		t.Fatalf("expected to get %d, but got %d", expected, rtcm.mtype)
	}
}

func Test1112(t *testing.T) {
	expected := 1112
	file := fmt.Sprintf("%d_frame.bin", expected)
	data, err := ioutil.ReadFile("./data/" + file)
	if err != nil {
		t.Error(err)
	}

	rtcm, err := Decode(data)
	if err != nil {
		t.Error(err)
	}

	if rtcm.crc != 0 {
		t.Fatalf("%s crc error", file)
	}

	if int(rtcm.mtype) != expected {
		t.Fatalf("expected to get %d, but got %d", expected, rtcm.mtype)
	}
}

func Test1113(t *testing.T) {
	expected := 1113
	file := fmt.Sprintf("%d_frame.bin", expected)
	data, err := ioutil.ReadFile("./data/" + file)
	if err != nil {
		t.Error(err)
	}

	rtcm, err := Decode(data)
	if err != nil {
		t.Error(err)
	}

	if rtcm.crc != 0 {
		t.Fatalf("%s crc error", file)
	}

	if int(rtcm.mtype) != expected {
		t.Fatalf("expected to get %d, but got %d", expected, rtcm.mtype)
	}
}

func Test1114(t *testing.T) {
	expected := 1114
	file := fmt.Sprintf("%d_frame.bin", expected)
	data, err := ioutil.ReadFile("./data/" + file)
	if err != nil {
		t.Error(err)
	}

	rtcm, err := Decode(data)
	if err != nil {
		t.Error(err)
	}

	if rtcm.crc != 0 {
		t.Fatalf("%s crc error", file)
	}

	if int(rtcm.mtype) != expected {
		t.Fatalf("expected to get %d, but got %d", expected, rtcm.mtype)
	}
}

func Test1115(t *testing.T) {
	expected := 1115
	file := fmt.Sprintf("%d_frame.bin", expected)
	data, err := ioutil.ReadFile("./data/" + file)
	if err != nil {
		t.Error(err)
	}

	rtcm, err := Decode(data)
	if err != nil {
		t.Error(err)
	}

	if rtcm.crc != 0 {
		t.Fatalf("%s crc error", file)
	}

	if int(rtcm.mtype) != expected {
		t.Fatalf("expected to get %d, but got %d", expected, rtcm.mtype)
	}
}

func Test1116(t *testing.T) {
	expected := 1116
	file := fmt.Sprintf("%d_frame.bin", expected)
	data, err := ioutil.ReadFile("./data/" + file)
	if err != nil {
		t.Error(err)
	}

	rtcm, err := Decode(data)
	if err != nil {
		t.Error(err)
	}

	if rtcm.crc != 0 {
		t.Fatalf("%s crc error", file)
	}

	if int(rtcm.mtype) != expected {
		t.Fatalf("expected to get %d, but got %d", expected, rtcm.mtype)
	}
}

func Test1117(t *testing.T) {
	expected := 1117
	file := fmt.Sprintf("%d_frame.bin", expected)
	data, err := ioutil.ReadFile("./data/" + file)
	if err != nil {
		t.Error(err)
	}

	rtcm, err := Decode(data)
	if err != nil {
		t.Error(err)
	}

	if rtcm.crc != 0 {
		t.Fatalf("%s crc error", file)
	}

	if int(rtcm.mtype) != expected {
		t.Fatalf("expected to get %d, but got %d", expected, rtcm.mtype)
	}
}

func Test1121(t *testing.T) {
	expected := 1121
	file := fmt.Sprintf("%d_frame.bin", expected)
	data, err := ioutil.ReadFile("./data/" + file)
	if err != nil {
		t.Error(err)
	}

	rtcm, err := Decode(data)
	if err != nil {
		t.Error(err)
	}

	if rtcm.crc != 0 {
		t.Fatalf("%s crc error", file)
	}

	if int(rtcm.mtype) != expected {
		t.Fatalf("expected to get %d, but got %d", expected, rtcm.mtype)
	}
}

func Test1122(t *testing.T) {
	expected := 1122
	file := fmt.Sprintf("%d_frame.bin", expected)
	data, err := ioutil.ReadFile("./data/" + file)
	if err != nil {
		t.Error(err)
	}

	rtcm, err := Decode(data)
	if err != nil {
		t.Error(err)
	}

	if rtcm.crc != 0 {
		t.Fatalf("%s crc error", file)
	}

	if int(rtcm.mtype) != expected {
		t.Fatalf("expected to get %d, but got %d", expected, rtcm.mtype)
	}
}

func Test1123(t *testing.T) {
	expected := 1123
	file := fmt.Sprintf("%d_frame.bin", expected)
	data, err := ioutil.ReadFile("./data/" + file)
	if err != nil {
		t.Error(err)
	}

	rtcm, err := Decode(data)
	if err != nil {
		t.Error(err)
	}

	if rtcm.crc != 0 {
		t.Fatalf("%s crc error", file)
	}

	if int(rtcm.mtype) != expected {
		t.Fatalf("expected to get %d, but got %d", expected, rtcm.mtype)
	}
}

func Test1124(t *testing.T) {
	expected := 1124
	file := fmt.Sprintf("%d_frame.bin", expected)
	data, err := ioutil.ReadFile("./data/" + file)
	if err != nil {
		t.Error(err)
	}

	rtcm, err := Decode(data)
	if err != nil {
		t.Error(err)
	}

	if rtcm.crc != 0 {
		t.Fatalf("%s crc error", file)
	}

	if int(rtcm.mtype) != expected {
		t.Fatalf("expected to get %d, but got %d", expected, rtcm.mtype)
	}
}

func Test1125(t *testing.T) {
	expected := 1125
	file := fmt.Sprintf("%d_frame.bin", expected)
	data, err := ioutil.ReadFile("./data/" + file)
	if err != nil {
		t.Error(err)
	}

	rtcm, err := Decode(data)
	if err != nil {
		t.Error(err)
	}

	if rtcm.crc != 0 {
		t.Fatalf("%s crc error", file)
	}

	if int(rtcm.mtype) != expected {
		t.Fatalf("expected to get %d, but got %d", expected, rtcm.mtype)
	}
}

func Test1126(t *testing.T) {
	expected := 1126
	file := fmt.Sprintf("%d_frame.bin", expected)
	data, err := ioutil.ReadFile("./data/" + file)
	if err != nil {
		t.Error(err)
	}

	rtcm, err := Decode(data)
	if err != nil {
		t.Error(err)
	}

	if rtcm.crc != 0 {
		t.Fatalf("%s crc error", file)
	}

	if int(rtcm.mtype) != expected {
		t.Fatalf("expected to get %d, but got %d", expected, rtcm.mtype)
	}
}

func Test1127(t *testing.T) {
	expected := 1127
	file := fmt.Sprintf("%d_frame.bin", expected)
	data, err := ioutil.ReadFile("./data/" + file)
	if err != nil {
		t.Error(err)
	}

	rtcm, err := Decode(data)
	if err != nil {
		t.Error(err)
	}

	if rtcm.crc != 0 {
		t.Fatalf("%s crc error", file)
	}

	if int(rtcm.mtype) != expected {
		t.Fatalf("expected to get %d, but got %d", expected, rtcm.mtype)
	}
}

func Test1230(t *testing.T) {
	expected := 1230
	file := fmt.Sprintf("%d_frame.bin", expected)
	data, err := ioutil.ReadFile("./data/" + file)
	if err != nil {
		t.Error(err)
	}

	rtcm, err := Decode(data)
	if err != nil {
		t.Error(err)
	}

	if rtcm.crc != 0 {
		t.Fatalf("%s crc error", file)
	}

	if int(rtcm.mtype) != expected {
		t.Fatalf("expected to get %d, but got %d", expected, rtcm.mtype)
	}
}
