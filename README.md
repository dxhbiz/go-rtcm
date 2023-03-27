# <center>go-rtcm</center>
rtcm decoding tool based on cgo, use [RTKLIB](https://github.com/tomojitakasu/RTKLIB/tree/rtklib_2.4.3/src) source code for core decoding

### Usage
```go
go get github.com/dxhbiz/go-rtcm
```
example.go
```go
package main

import (
	"fmt"

	"github.com/dxhbiz/go-rtcm"
)

func main() {
	rtcmBuf := []byte{211, 0, 19, 62, 208, 0, 3, 181, 150, 68, 101, 184, 134, 55, 187, 179, 82, 183, 116, 225, 229, 121, 76, 148, 146}

	check := rtcm.CheckCrc(rtcmBuf)
	if !check {
		panic("crc error")
	}

	grtcm, err := rtcm.Decode(rtcmBuf)
	if err != nil {
		panic(err)
	}

	fmt.Printf("rtcm type:%d, position x:%f, y:%f, z:%f\n", grtcm.Type, grtcm.Sta.Pos[0], grtcm.Sta.Pos[1], grtcm.Sta.Pos[2])
}

```

