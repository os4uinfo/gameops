package front

import (
	"fmt"
	"io/ioutil"
)

// bindata_read reads the given file from disk. It returns
// an error on failure.
func bindata_read(path, name string) ([]byte, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset %s at %s: %v", name, path, err)
	}
	return buf, err
}


// html_static_css_bootstrap_theme_css reads file data from disk.
// It panics if something went wrong in the process.
func html_static_css_bootstrap_theme_css() ([]byte, error) {
	return bindata_read(
		"/work/devops/src/gameops/front/html/static/css/bootstrap-theme.css",
		"html/static/css/bootstrap-theme.css",
	)
}

// html_static_css_bootstrap_theme_css_map reads file data from disk.
// It panics if something went wrong in the process.
func html_static_css_bootstrap_theme_css_map() ([]byte, error) {
	return bindata_read(
		"/work/devops/src/gameops/front/html/static/css/bootstrap-theme.css.map",
		"html/static/css/bootstrap-theme.css.map",
	)
}

// html_static_css_bootstrap_theme_min_css reads file data from disk.
// It panics if something went wrong in the process.
func html_static_css_bootstrap_theme_min_css() ([]byte, error) {
	return bindata_read(
		"/work/devops/src/gameops/front/html/static/css/bootstrap-theme.min.css",
		"html/static/css/bootstrap-theme.min.css",
	)
}

// html_static_css_bootstrap_css reads file data from disk.
// It panics if something went wrong in the process.
func html_static_css_bootstrap_css() ([]byte, error) {
	return bindata_read(
		"/work/devops/src/gameops/front/html/static/css/bootstrap.css",
		"html/static/css/bootstrap.css",
	)
}

// html_static_css_bootstrap_css_map reads file data from disk.
// It panics if something went wrong in the process.
func html_static_css_bootstrap_css_map() ([]byte, error) {
	return bindata_read(
		"/work/devops/src/gameops/front/html/static/css/bootstrap.css.map",
		"html/static/css/bootstrap.css.map",
	)
}

// html_static_css_bootstrap_min_css reads file data from disk.
// It panics if something went wrong in the process.
func html_static_css_bootstrap_min_css() ([]byte, error) {
	return bindata_read(
		"/work/devops/src/gameops/front/html/static/css/bootstrap.min.css",
		"html/static/css/bootstrap.min.css",
	)
}

// html_static_fonts_glyphicons_halflings_regular_eot reads file data from disk.
// It panics if something went wrong in the process.
func html_static_fonts_glyphicons_halflings_regular_eot() ([]byte, error) {
	return bindata_read(
		"/work/devops/src/gameops/front/html/static/fonts/glyphicons-halflings-regular.eot",
		"html/static/fonts/glyphicons-halflings-regular.eot",
	)
}

// html_static_fonts_glyphicons_halflings_regular_svg reads file data from disk.
// It panics if something went wrong in the process.
func html_static_fonts_glyphicons_halflings_regular_svg() ([]byte, error) {
	return bindata_read(
		"/work/devops/src/gameops/front/html/static/fonts/glyphicons-halflings-regular.svg",
		"html/static/fonts/glyphicons-halflings-regular.svg",
	)
}

// html_static_fonts_glyphicons_halflings_regular_ttf reads file data from disk.
// It panics if something went wrong in the process.
func html_static_fonts_glyphicons_halflings_regular_ttf() ([]byte, error) {
	return bindata_read(
		"/work/devops/src/gameops/front/html/static/fonts/glyphicons-halflings-regular.ttf",
		"html/static/fonts/glyphicons-halflings-regular.ttf",
	)
}

// html_static_fonts_glyphicons_halflings_regular_woff reads file data from disk.
// It panics if something went wrong in the process.
func html_static_fonts_glyphicons_halflings_regular_woff() ([]byte, error) {
	return bindata_read(
		"/work/devops/src/gameops/front/html/static/fonts/glyphicons-halflings-regular.woff",
		"html/static/fonts/glyphicons-halflings-regular.woff",
	)
}

// html_static_fonts_glyphicons_halflings_regular_woff2 reads file data from disk.
// It panics if something went wrong in the process.
func html_static_fonts_glyphicons_halflings_regular_woff2() ([]byte, error) {
	return bindata_read(
		"/work/devops/src/gameops/front/html/static/fonts/glyphicons-halflings-regular.woff2",
		"html/static/fonts/glyphicons-halflings-regular.woff2",
	)
}

// html_static_js_bootstrap_js reads file data from disk.
// It panics if something went wrong in the process.
func html_static_js_bootstrap_js() ([]byte, error) {
	return bindata_read(
		"/work/devops/src/gameops/front/html/static/js/bootstrap.js",
		"html/static/js/bootstrap.js",
	)
}

// html_static_js_bootstrap_min_js reads file data from disk.
// It panics if something went wrong in the process.
func html_static_js_bootstrap_min_js() ([]byte, error) {
	return bindata_read(
		"/work/devops/src/gameops/front/html/static/js/bootstrap.min.js",
		"html/static/js/bootstrap.min.js",
	)
}

// html_static_js_npm_js reads file data from disk.
// It panics if something went wrong in the process.
func html_static_js_npm_js() ([]byte, error) {
	return bindata_read(
		"/work/devops/src/gameops/front/html/static/js/npm.js",
		"html/static/js/npm.js",
	)
}

// html_tmpl_index_html reads file data from disk.
// It panics if something went wrong in the process.
func html_tmpl_index_html() ([]byte, error) {
	return bindata_read(
		"/work/devops/src/gameops/front/html/tmpl/index.html",
		"html/tmpl/index.html",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	if f, ok := _bindata[name]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string] func() ([]byte, error) {
	"html/static/css/bootstrap-theme.css": html_static_css_bootstrap_theme_css,
	"html/static/css/bootstrap-theme.css.map": html_static_css_bootstrap_theme_css_map,
	"html/static/css/bootstrap-theme.min.css": html_static_css_bootstrap_theme_min_css,
	"html/static/css/bootstrap.css": html_static_css_bootstrap_css,
	"html/static/css/bootstrap.css.map": html_static_css_bootstrap_css_map,
	"html/static/css/bootstrap.min.css": html_static_css_bootstrap_min_css,
	"html/static/fonts/glyphicons-halflings-regular.eot": html_static_fonts_glyphicons_halflings_regular_eot,
	"html/static/fonts/glyphicons-halflings-regular.svg": html_static_fonts_glyphicons_halflings_regular_svg,
	"html/static/fonts/glyphicons-halflings-regular.ttf": html_static_fonts_glyphicons_halflings_regular_ttf,
	"html/static/fonts/glyphicons-halflings-regular.woff": html_static_fonts_glyphicons_halflings_regular_woff,
	"html/static/fonts/glyphicons-halflings-regular.woff2": html_static_fonts_glyphicons_halflings_regular_woff2,
	"html/static/js/bootstrap.js": html_static_js_bootstrap_js,
	"html/static/js/bootstrap.min.js": html_static_js_bootstrap_min_js,
	"html/static/js/npm.js": html_static_js_npm_js,
	"html/tmpl/index.html": html_tmpl_index_html,

}
