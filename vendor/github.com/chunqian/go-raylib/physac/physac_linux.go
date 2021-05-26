// +build linux,!arm,!arm64

package physac

/*
// #cgo linux LDFLAGS: -L${SRCDIR}/../lib/raylib/src -lraylib -lGL -lm -pthread -ldl -lrt -lX11
#cgo linux LDFLAGS: -lGL -lm -pthread -ldl -lrt -lX11
#cgo linux CFLAGS: -I../lib/raylib/src -DPHYSAC_IMPLEMENTATION -DPHYSAC_STATIC -DPHYSAC_NO_THREADS
*/
import "C"
