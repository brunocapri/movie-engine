// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func SearchBar() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<form method=\"get\" action=\"/search\" class=\"flex w-80 items-center justify-between bg-neutral-800 rounded-md p-2\"><img src=\"dist/mag.svg\"> <input id=\"query\" name=\"q\" type=\"text\" maxlength=\"45\" placeholder=\"Search for a movie plot...\" class=\"w-full bg-neutral-800 text-white placeholder-gray-400 pl-2 py-1 rounded-md focus:outline-none\"> <button type=\"submit\" id=\"form-button\" class=\"hidden bg-neutral-700 rounded-md p-1.5 ml-1\"><img src=\"dist/arrow.svg\"></button></form><script>\n    const input = document.getElementById(\"query\")\n    const button = document.getElementById(\"form-button\")\n    input.addEventListener('input', function (evt) {\n        if (this.value !== \"\") {\n          button.classList.remove(\"hidden\");\n          return\n        }\n        button.classList.add(\"hidden\");\n    })\n  </script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
