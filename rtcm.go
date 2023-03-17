package rtcm

//#cgo CFLAGS:
//
//#include "rtcm.h"
import "C"

// Decode decode the rtcm data
func Decode(data []byte) (rtcm C.rtcm_t, err error) {
	C.init_rtcm(&rtcm)
	for _, d := range data {
		C.input_rtcm3(&rtcm, C.uint8_t(d))
	}
	C.free_rtcm(&rtcm)
	return
}
