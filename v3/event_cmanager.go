package vlc

/*
#cgo LDFLAGS: -lvlc
#include <vlc/vlc.h>
*/
import "C"

import (
	"unsafe"
)

//export eventDispatch
func eventDispatch(event C.clibvlc_event_t, userData unsafe.Pointer) {
	ctx, ok := inst.events.get(EventID(uintptr(userData)))
	if !ok {
		return
	}
	if ctx.callback == nil {
		return
	}

	ctx.callback(ctx.event, ctx.userData)
}
