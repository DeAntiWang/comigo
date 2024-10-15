// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.778
package common

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "github.com/gin-gonic/gin"

// MainLayout 定义网页布局
func MainLayout(c *gin.Context, bodyContent templ.Component) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><meta http-equiv=\"X-UA-Compatible\" content=\"ie=edge\"><title>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(GetPageTitle(c.Param("id")))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/common/main_layout.templ`, Line: 13, Col: 39}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</title><meta name=\"keywords\" content=\"Comigo  Comic Manga Reader 在线漫画 阅读器\"><meta name=\"description\" content=\"Simple Manga Reader in Linux，Windows，Mac OS\"><!--TODO:PWA模式  <link rel=\"manifest\" href=\"/static/manifest.webmanifest\"/>  --><link rel=\"apple-touch-icon\" href=\"/static/apple-touch-icon.png\"><link rel=\"shortcut icon\" href=\"/static/favicon.ico\" type=\"image/x-icon\"><link rel=\"icon\" href=\"/static/favicon.png\" sizes=\"any\"><!--  font-sans：https://tailwindcss.com/docs/font-family --><link href=\"/static/styles.css\" rel=\"stylesheet\"><!--  让旧版浏览器也支持 原生的 image 延迟加载属性  https://github.com/mfranzke/loading-attribute-polyfill  --><link rel=\"stylesheet\" href=\"/static/script/loading-attribute-polyfill.css\"><script src=\"/static/script/loading-attribute-polyfill.umd.js\" async></script></head><!-- x-bind: Alpine.js的语法，声明全局主题 theme --><!-- $persist 可以存储原始值以及数组和对象。本地存储，默认的key是 _x_变量名 --><!-- ！！！当变量的类型发生变化时，必须手动清除 localStorage，否则相应数值将无法正确更新。！！！ --><!-- 详细用法参见： https://alpinejs.dev/plugins/persist -->")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.RenderScriptItems(ctx, templ_7745c5c3_Buffer, BodyScripts())
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<body x-data=\"{ theme: $persist(&#39;retro&#39;) }\" x-bind:data-theme=\"theme\" onload=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 templ.ComponentScript = BodyScripts()
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var3.Call)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"flex flex-col items-center justify-between w-full h-full min-h-screen p-0 m-0 font-sans\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = bodyContent.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</body><script src=\"/static/scripts.js\"></script></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

// BodyScripts JavaScript 代码。在页面加载时，绑定全屏按钮的点击事件。
func BodyScripts() templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_BodyScripts_6fa6`,
		Function: `function __templ_BodyScripts_6fa6(){var element = document.getElementById('FullScreenIcon');
    if (element) {
        element.addEventListener('click', function() {
        	if (!FullScreenHelper.supported()){
            	alert('Your browser does not support full screen mode.');
            	return;
            }
            if(!FullScreenHelper.state()){
                FullScreenHelper.request()
            }else{
                FullScreenHelper.exit()
            }
        });
    }
}`,
		Call:       templ.SafeScript(`__templ_BodyScripts_6fa6`),
		CallInline: templ.SafeScriptInline(`__templ_BodyScripts_6fa6`),
	}
}

var _ = templruntime.GeneratedTemplate
