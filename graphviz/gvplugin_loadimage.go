package graphviz

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)
/// @ingroup plugin_api
/// @{
// llgo:link (*UsershapeT).UsershapeFileAccess C.gvusershape_file_access
func (recv_ *UsershapeT) UsershapeFileAccess() c.Int {
	return 0
}
// llgo:link (*UsershapeT).UsershapeFileRelease C.gvusershape_file_release
func (recv_ *UsershapeT) UsershapeFileRelease() {
}
