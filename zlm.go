// Package zlmgo provides Go bindings to the Zen License Manager (ZLM).
package zlmgo

/*
#cgo CFLAGS: -I/usr/local/include
#cgo LDFLAGS: /usr/local/lib/libzlm.a

#include <zlm.h>

char zlmgo_errbuf_array[ZLM_ERRBUF];
char* zlmgo_errbuf = zlmgo_errbuf_array; // hack around cgo warning
*/
import "C"

import (
	"errors"
	"runtime"
)

type License struct {
	l *C.ZlmLicense
}

// LicenseNew returns a new license object (panics if not enough memory is available).
func LicenseNew() *License {
	license := &License{C.zlm_license_new(C.zlmgo_errbuf)}
	if license.l == nil {
		panic(C.GoString(C.zlmgo_errbuf))
	}
	runtime.SetFinalizer(license, (*License).free)
	return license
}

func (license *License) Get(product, version, argv0, path, license_string string) error {
	var c_product, c_version, c_argv0, c_path, c_license_string *C.char
	if product != "" {
		c_product = C.CString(product)
	}
	if version != "" {
		c_version = C.CString(version)
	}
	if argv0 != "" {
		c_argv0 = C.CString(argv0)
	}
	if path != "" {
		c_path = C.CString(path)
	}
	if license_string != "" {
		c_license_string = C.CString(license_string)
	}
	if C.zlm_license_get(license.l, c_product, c_version, c_argv0, c_path, c_license_string, C.zlmgo_errbuf) !=
		C.ZLM_OK {
		return errors.New(C.GoString(C.zlmgo_errbuf))
	}
	return nil
}

func (license *License) free() {
	C.zlm_license_free(license.l)
}
