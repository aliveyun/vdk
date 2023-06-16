package format

import (
	"github.com/aliveyun/vdk/av/avutil"
	"github.com/aliveyun/vdk/format/aac"
	"github.com/aliveyun/vdk/format/flv"
	"github.com/aliveyun/vdk/format/mp4"
	"github.com/aliveyun/vdk/format/rtmp"
	"github.com/aliveyun/vdk/format/rtsp"
	"github.com/aliveyun/vdk/format/ts"
)

func RegisterAll() {
	avutil.DefaultHandlers.Add(mp4.Handler)
	avutil.DefaultHandlers.Add(ts.Handler)
	avutil.DefaultHandlers.Add(rtmp.Handler)
	avutil.DefaultHandlers.Add(rtsp.Handler)
	avutil.DefaultHandlers.Add(flv.Handler)
	avutil.DefaultHandlers.Add(aac.Handler)
}
