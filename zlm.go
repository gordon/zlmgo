// Package zlmgo provides Go bindings to the Zen License Manager (ZLM).
package zlmgo

/*
#cgo CFLAGS: -I/usr/local/include
#cgo LDFLAGS: /usr/local/lib/libzlm.a

#include <stdlib.h>
#include <zlm.h>

char zlmgo_errbuf_array[ZLM_ERRBUF];
char* zlmgo_errbuf = zlmgo_errbuf_array; // hack around cgo warning
*/
import "C"

import (
	"errors"
	"runtime"
	"unsafe"
)

// License wraps a ZLM license object.
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

// Get wraps zlm_license_get(), see
// https://zenlicensemanager.com/documentation/#API
func (license *License) Get(product, version, argv0, path, licenseString string) error {
	var c_product, c_version, c_argv0, c_path, c_license_string *C.char
	// convert argument to C strings
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
	if licenseString != "" {
		c_license_string = C.CString(licenseString)
	}
	// call actual method
	if C.zlm_license_get(license.l, c_product, c_version, c_argv0, c_path, c_license_string, C.zlmgo_errbuf) != C.ZLM_OK {
		return errors.New(C.GoString(C.zlmgo_errbuf))
	}
	return nil
}

func (license *License) free() {
	C.zlm_license_free(license.l)
}

func (license *License) Product() string {
	return C.GoString(C.zlm_license_product(license.l))
}

func (license *License) Expiry() string {
	return C.GoString(C.zlm_license_expiry(license.l))
}

func (license *License) ExpiryDays() int {
	return int(C.zlm_license_expiry_days(license.l))
}

func (license *License) Customer() string {
	return C.GoString(C.zlm_license_customer(license.l))
}

func (license *License) Userdata() string {
	return C.GoString(C.zlm_license_userdata(license.l))
}

func (license *License) Next() error {
	if C.zlm_license_next(license.l, C.zlmgo_errbuf) != C.ZLM_OK {
		return errors.New(C.GoString(C.zlmgo_errbuf))
	}
	return nil
}

func Version() string {
	return C.GoString(C.zlm_version())
}

func HostidJSON() (string, error) {
	cs := C.zlm_hostid_json(C.zlmgo_errbuf)
	if cs == nil {
		return "", errors.New(C.GoString(C.zlmgo_errbuf))
	}
	hostid := C.GoString(cs)
	C.free(unsafe.Pointer(cs))
	return hostid, nil
}

func (license *License) CheckA() {
	C.zlm_license_check_a(license.l)
}

func (license *License) CheckB() {
	C.zlm_license_check_b(license.l)
}

func (license *License) CheckC() {
	C.zlm_license_check_c(license.l)
}

func (license *License) CheckD() {
	C.zlm_license_check_d(license.l)
}

func (license *License) CheckE() {
	C.zlm_license_check_e(license.l)
}

func (license *License) CheckF() {
	C.zlm_license_check_f(license.l)
}

func (license *License) CheckG() {
	C.zlm_license_check_g(license.l)
}

func (license *License) CheckH() {
	C.zlm_license_check_h(license.l)
}

func (license *License) CheckI() {
	C.zlm_license_check_i(license.l)
}

func (license *License) CheckJ() {
	C.zlm_license_check_j(license.l)
}

func (license *License) CheckK() {
	C.zlm_license_check_k(license.l)
}

func (license *License) CheckL() {
	C.zlm_license_check_l(license.l)
}

func (license *License) CheckM() {
	C.zlm_license_check_m(license.l)
}

func (license *License) CheckN() {
	C.zlm_license_check_n(license.l)
}

func (license *License) CheckO() {
	C.zlm_license_check_o(license.l)
}

func (license *License) CheckP() {
	C.zlm_license_check_p(license.l)
}

func (license *License) CheckQ() {
	C.zlm_license_check_q(license.l)
}

func (license *License) CheckR() {
	C.zlm_license_check_r(license.l)
}
