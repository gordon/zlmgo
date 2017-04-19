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
	var cProduct, cVersion, cArgv0, cPath, cLicenseString *C.char
	// convert argument to C strings
	if product != "" {
		cProduct = C.CString(product)
	}
	if version != "" {
		cVersion = C.CString(version)
	}
	if argv0 != "" {
		cArgv0 = C.CString(argv0)
	}
	if path != "" {
		cPath = C.CString(path)
	}
	if licenseString != "" {
		cLicenseString = C.CString(licenseString)
	}
	// call actual method
	if C.zlm_license_get(license.l, cProduct, cVersion, cArgv0, cPath, cLicenseString, C.zlmgo_errbuf) != C.ZLM_OK {
		return errors.New(C.GoString(C.zlmgo_errbuf))
	}
	return nil
}

func (license *License) free() {
	C.zlm_license_free(license.l)
}

// Product wraps zlm_license_product(), see
// https://zenlicensemanager.com/documentation/#API
func (license *License) Product() string {
	return C.GoString(C.zlm_license_product(license.l))
}

// Expiry wraps zlm_license_expiry(), see
// https://zenlicensemanager.com/documentation/#API
func (license *License) Expiry() string {
	return C.GoString(C.zlm_license_expiry(license.l))
}

// ExpiryDays wraps zlm_license_expiry_days(), see
// https://zenlicensemanager.com/documentation/#API
func (license *License) ExpiryDays() int {
	return int(C.zlm_license_expiry_days(license.l))
}

// Customer wraps zlm_license_customer(), see
// https://zenlicensemanager.com/documentation/#API
func (license *License) Customer() string {
	return C.GoString(C.zlm_license_customer(license.l))
}

// Userdata wraps zlm_license_userdata(), see
// https://zenlicensemanager.com/documentation/#API
func (license *License) Userdata() string {
	return C.GoString(C.zlm_license_userdata(license.l))
}

// UserdataUnescaped wraps zlm_license_userdata_unescaped(), see
// https://zenlicensemanager.com/documentation/#API
func (license *License) UserdataUnescaped() string {
	return C.GoString(C.zlm_license_userdata_unescaped(license.l))
}

// Next wraps zlm_license_next(), see
// https://zenlicensemanager.com/documentation/#API
func (license *License) Next() error {
	if C.zlm_license_next(license.l, C.zlmgo_errbuf) != C.ZLM_OK {
		return errors.New(C.GoString(C.zlmgo_errbuf))
	}
	return nil
}

// Version wraps zlm_version(), see
// https://zenlicensemanager.com/documentation/#API
func Version() string {
	return C.GoString(C.zlm_version())
}

// HostidJSON wraps zlm_hostid_json(), see
// https://zenlicensemanager.com/documentation/#API
func HostidJSON() (string, error) {
	cs := C.zlm_hostid_json(C.zlmgo_errbuf)
	if cs == nil {
		return "", errors.New(C.GoString(C.zlmgo_errbuf))
	}
	hostid := C.GoString(cs)
	C.zlm_free(unsafe.Pointer(cs))
	return hostid, nil
}

// CheckA wraps zlm_license_check_a(), see
// https://zenlicensemanager.com/documentation/#API
func (license *License) CheckA() {
	C.zlm_license_check_a(license.l)
}

// CheckB wraps zlm_license_check_b(), see
// https://zenlicensemanager.com/documentation/#API
func (license *License) CheckB() {
	C.zlm_license_check_b(license.l)
}

// CheckC wraps zlm_license_check_c(), see
// https://zenlicensemanager.com/documentation/#API
func (license *License) CheckC() {
	C.zlm_license_check_c(license.l)
}

// CheckD wraps zlm_license_check_d(), see
// https://zenlicensemanager.com/documentation/#API
func (license *License) CheckD() {
	C.zlm_license_check_d(license.l)
}

// CheckE wraps zlm_license_check_e(), see
// https://zenlicensemanager.com/documentation/#API
func (license *License) CheckE() {
	C.zlm_license_check_e(license.l)
}

// CheckF wraps zlm_license_check_f(), see
// https://zenlicensemanager.com/documentation/#API
func (license *License) CheckF() {
	C.zlm_license_check_f(license.l)
}

// CheckG wraps zlm_license_check_g(), see
// https://zenlicensemanager.com/documentation/#API
func (license *License) CheckG() {
	C.zlm_license_check_g(license.l)
}

// CheckH wraps zlm_license_check_h(), see
// https://zenlicensemanager.com/documentation/#API
func (license *License) CheckH() {
	C.zlm_license_check_h(license.l)
}

// CheckI wraps zlm_license_check_i(), see
// https://zenlicensemanager.com/documentation/#API
func (license *License) CheckI() {
	C.zlm_license_check_i(license.l)
}

// CheckJ wraps zlm_license_check_j(), see
// https://zenlicensemanager.com/documentation/#API
func (license *License) CheckJ() {
	C.zlm_license_check_j(license.l)
}

// CheckK wraps zlm_license_check_k(), see
// https://zenlicensemanager.com/documentation/#API
func (license *License) CheckK() {
	C.zlm_license_check_k(license.l)
}

// CheckL wraps zlm_license_check_l(), see
// https://zenlicensemanager.com/documentation/#API
func (license *License) CheckL() {
	C.zlm_license_check_l(license.l)
}

// CheckM wraps zlm_license_check_m(), see
// https://zenlicensemanager.com/documentation/#API
func (license *License) CheckM() {
	C.zlm_license_check_m(license.l)
}

// CheckN wraps zlm_license_check_n(), see
// https://zenlicensemanager.com/documentation/#API
func (license *License) CheckN() {
	C.zlm_license_check_n(license.l)
}

// CheckO wraps zlm_license_check_o(), see
// https://zenlicensemanager.com/documentation/#API
func (license *License) CheckO() {
	C.zlm_license_check_o(license.l)
}

// CheckP wraps zlm_license_check_p(), see
// https://zenlicensemanager.com/documentation/#API
func (license *License) CheckP() {
	C.zlm_license_check_p(license.l)
}

// CheckQ wraps zlm_license_check_q(), see
// https://zenlicensemanager.com/documentation/#API
func (license *License) CheckQ() {
	C.zlm_license_check_q(license.l)
}

// CheckR wraps zlm_license_check_r(), see
// https://zenlicensemanager.com/documentation/#API
func (license *License) CheckR() {
	C.zlm_license_check_r(license.l)
}
