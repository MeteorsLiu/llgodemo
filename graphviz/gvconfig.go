package graphviz

import _ "unsafe"
// llgo:link (*GVCT).ConfigPluginInstallFromLibrary C.gvconfig_plugin_install_from_library
func (recv_ *GVCT) ConfigPluginInstallFromLibrary(package_path *int8, library *PluginLibraryT) {
}
