package app

// TEMPLATE FILE

import (
	"html/template"
	"path/filepath"

	"github.com/cbonello/revel-csrf"
	"github.com/iassic/revel-modz/modules/grunt"
	"github.com/revel/revel"

	ctrl "github.com/iassic/revel-modz/sample/app/controllers"
)

// in your app initialization..
func init() {
	revel.TemplateFuncs["appendjs"] = func(renderArgs map[string]interface{}, key string, value interface{}) template.HTML {
		s := value.(string)
		js_code := template.JS(s)

		if renderArgs[key] == nil {
			renderArgs[key] = []interface{}{js_code}
		} else {
			renderArgs[key] = append(renderArgs[key].([]interface{}), js_code)
		}
		return template.HTML("")
	}

	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		revel.SessionFilter,           // Restore and write the session cookie.
		revel.FlashFilter,             // Restore and write the flash cookie.
		csrf.CSRFFilter,               // CSRF prevention.
		revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,              // Resolve the requested language
		HeaderFilter,                  // Add some security based headers
		revel.InterceptorFilter,       // Run interceptors around the action.
		revel.ActionInvoker,           // Invoke the action.
	}

	revel.OnAppStart(func() {
		appPath := revel.BasePath
		for _, gruntCompiler := range compilers {
			path := filepath.Join(appPath, gruntCompiler.Path)
			revel.INFO.Printf("Listening: %q\n", path)
			revel.MainWatcher.Listen(gruntCompiler, path)
		}
	})

	// add interceptors
	revel.InterceptMethod((*ctrl.DbController).Begin, revel.BEFORE)
	revel.InterceptMethod(ctrl.App.RenderArgsFill, revel.BEFORE)
	revel.InterceptMethod(ctrl.App.RecordPageRequest, revel.BEFORE)
	revel.InterceptMethod(ctrl.User.CheckLoggedIn, revel.BEFORE)
	revel.InterceptMethod(ctrl.Admin.CheckLoggedIn, revel.BEFORE)
	revel.InterceptMethod((*ctrl.DbController).Commit, revel.AFTER)
	revel.InterceptMethod((*ctrl.DbController).Rollback, revel.FINALLY)

	revel.OnAppStart(func() {
		ctrl.InitDB()
		if revel.RunMode == "dev" {
			ctrl.SetupDevDB()
		}
		if revel.RunMode == "prod" {
			ctrl.SetupTables()
		}
	})

}

var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
	// Add some common security headers
	c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
	c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
	c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")

	fc[0](c, fc[1:]) // Execute the next filter stage.
}

var compilers = []*grunt.GruntCompiler{
	grunt.NewCompiler("Gruntfile.js", "Gruntfile.js", "default"),

	grunt.NewCompiler("Foundation JS", "app/assets/js/foundation", "uglify:foundation_js"),
	grunt.NewCompiler("Foundation SASS", "app/assets/sass/foundation", "sass:foundation_css"),
	grunt.NewCompiler("Foundation SASS", "app/assets/sass/foundation_custom.scss", "sass:foundation_css"),

	grunt.NewCompilerOnStart("sample JS", "app/assets/js/sample", "uglify:sample_js"),
	grunt.NewCompiler("sample SASS", "app/assets/sass/sample", "sass:sample_css"),
	grunt.NewCompiler("sample SASS", "app/assets/sass/sample.scss", "sass:sample_css"),
}
