// +build with-x11

package gtk

import (
	"unsafe"

	"github.com/radosroka/go-gtk/gdk"
)

func (v *Window) XID() int32 {
	return gdk.WindowFromUnsafe(unsafe.Pointer(v.GWidget.window)).GetNativeWindowID()
}
