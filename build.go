package aria2go

/*
// Linux Build Tags
// ----------------
#cgo linux CXXFLAGS: -I${SRCDIR}/library -std=c++11
#cgo linux LDFLAGS: -L${SRCDIR}/library/linux -l:libaria2.a -l:libcares.a -l:libssh2.a -lxml2 -l:libsqlite3.a -l:libz.a -lssl -lcrypto -ldl

// Windows Build Tags
// ----------------
#cgo windows CXXFLAGS: -I${SRCDIR}/library -std=c++11
#cgo windows LDFLAGS: -L${SRCDIR}/library/win -l:libaria2.lib -l:libcares.lib -l:libexpat.lib -l:libgmp.lib -l:libsqlite3.lib -l:libssh2.lib -l:libz.lib -lws2_32 -lbcrypt -liphlpapi -lcrypt32 -lsecur32

// Darwin Build Tags
// ----------------ß
#cgo darwin CXXFLAGS: -I${SRCDIR}/library -std=c++11
#cgo darwin LDFLAGS:  -L${SRCDIR}/library/darwin ${SRCDIR}/library/darwin/libaria2.a ${SRCDIR}/library/darwin/libcares.a ${SRCDIR}/library/darwin/libcrypto.a ${SRCDIR}/library/darwin/libgmp.a ${SRCDIR}/library/darwin/libssh2.a ${SRCDIR}/library/darwin/libxml2.a ${SRCDIR}/library/darwin/libsqlite3.a ${SRCDIR}/library/darwin/libz.a ${SRCDIR}/library/darwin/libssl.a -ldl -framework CoreFoundation -framework Security -liconv

*/
import "C"
