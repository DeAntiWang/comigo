// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package flip

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"github.com/gin-gonic/gin"
	"github.com/yumenaka/comigo/entity"
	"github.com/yumenaka/comigo/htmx/embed_files"
	"github.com/yumenaka/comigo/htmx/state"
	"github.com/yumenaka/comigo/htmx/templates/common"
)

// FlipPage 定义 BodyHTML
func FlipPage(c *gin.Context, s *state.GlobalState, book *entity.Book, readingProgress *entity.ReadingProgress) templ.Component {
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
		templ_7745c5c3_Err = InsertData(book, s).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = common.Header(
			c,
			common.HeaderProps{
				Title:           common.GetPageTitle(book.BookInfo.BookID),
				ShowReturnIcon:  true,
				ReturnUrl:       common.GetReturnUrl(book.BookInfo.BookID),
				SetDownLoadLink: false,
				InShelf:         false,
				DownLoadLink:    "",
				SetTheme:        true,
				FlipMode:        true,
			}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = FlipMainArea(s, book).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = common.Drawer(s, FlipDrawerSlot()).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = common.QRCode(s).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func InsertData(bookData any, stateData any) templ.Component {
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
		templ_7745c5c3_Var2 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var2 == nil {
			templ_7745c5c3_Var2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templ.JSONScript("NowBook", bookData).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.JSONScript("GlobalState", stateData).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func FlipMainArea(s *state.GlobalState, book *entity.Book) templ.Component {
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
		templ_7745c5c3_Var3 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var3 == nil {
			templ_7745c5c3_Var3 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"mouseMoveArea\" class=\"grid-line flex flex-col items-center justify-center flex-1 w-full max-w-full bg-base-100 text-base-content\" :class=\"(theme.toString() ===&#39;light&#39;||theme.toString() ===&#39;dark&#39;||theme.toString() ===&#39;retro&#39;||theme.toString() ===&#39;lofi&#39;||theme.toString() ===&#39;nord&#39;) ? ($store.global.bgPattern !== &#39;none&#39;?$store.global.bgPattern+&#39; bg-base-300&#39;:&#39;bg-base-300&#39;):($store.global.bgPattern !== &#39;none&#39;?$store.global.bgPattern:&#39;&#39;)\"><div class=\"manga_area\" id=\"MangaMain\"><div class=\"manga_area_img_div\"><!-- 非自动拼合模式最简单,直接显示一张图 --><img id=\"NowImage\" class=\"max-w-full max-h-screen m-0 rounded shadow-lg\" src=\"/static/images/ball-triangle.svg\" alt=\"\"><!-- 简单拼合双页,不管单双页什么的 --><img id=\"NextImage\" x-show=\"$store.flip.doublePageMode &amp;&amp; $store.flipnowPageNum &lt; $store.flip.allPageNum\" src=\"/static/images/ball-triangle.svg\"></div></div></div><!-- 底部的阅读进度条 --><!-- https://flowbite.com/docs/forms/range/ --><!-- 宽度：w-5/6 https://www.tailwindcss.cn/docs/width 使用 w-{fraction} 或 w-full 将元素设置为基于百分比的宽度。 --><!-- 定位：https://www.tailwindcss.cn/docs/position  --><!-- 使用 fixed 来定位一个元素相对于浏览器窗视口的位置。偏移量是相对于视口计算的，且该元素将作为绝对定位的子元素的位置参考。 --><!-- 控制 flex 和 grid 项目如何沿着容器的主轴定位:https://www.tailwindcss.cn/docs/justify-content --><!-- Tailwind 的容器不会自动居中，也没有任何内置的水平方向的内边距。要使一个容器居中，使用 mx-auto 功能类： --><div id=\"steps-range_area\" onmouseover=\"showToolbar();\" onmouseout=\"hideToolbar()\" class=\"w-full px-2 overflow-hidden bg-gray-400 border border-blue-800 rounded toolbar h-14 opacity-80\" :class=\"Alpine.store(&#39;flip&#39;).autoHideToolbar? &#39;absolute fixed bottom-0&#39;:&#39;flex flex-col justify-center&#39;\"><label for=\"steps-range\" class=\"block m-0 text-sm font-medium text-center text-gray-900 dark:text-white\" x-text=\"$store.flip.nowPageNum+&#39;/&#39;+$store.flip.allPageNum\"></label> <input id=\"steps-range\" class=\"w-full h-2 mb-2 bg-yellow-800 rounded-lg appearance-none cursor-pointer dark:bg-gray-700\" type=\"range\" min=\"1\" :max=\"$store.flip.allPageNum\" x-model=\"$store.flip.nowPageNum\" onchange=\"setImageSrc()\" step=\"1\"></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.Raw("<style>"+embed_files.GetFileStr("static/flip.css")+"</style>").Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func FlipDrawerSlot() templ.Component {
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
		templ_7745c5c3_Var4 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var4 == nil {
			templ_7745c5c3_Var4 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!-- 阅读模式 --><!-- toggle组件来自： https://flowbite.com/docs/forms/toggle/ --><label class=\"inline-flex items-center w-full my-4 cursor-pointer outline outline-offset-8 outline-dotted hover:outline outline-2\"><input type=\"checkbox\" :value=\"$store.global.readMode === &#39;scroll&#39;\" x-on:click=\"$store.global.toggleReadMode()\" class=\"sr-only peer\"><div class=\"relative w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 dark:peer-focus:ring-blue-800 rounded-full peer dark:bg-gray-700 peer-checked:after:translate-x-full rtl:peer-checked:after:-translate-x-full peer-checked:after:border-white after:content-[&#39;&#39;] after:absolute after:top-[2px] after:start-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all dark:border-gray-600 peer-checked:bg-blue-600\"></div><span class=\"text-sm font-medium ms-3\" x-text=\"$store.global.readMode === &#39;scroll&#39;?i18next.t(&#39;scroll_mode&#39;):i18next.t(&#39;flip_mode&#39;)\">Toggle me</span></label><!-- 远程同步翻页 --><label class=\"inline-flex items-center w-full my-4 cursor-pointer outline outline-offset-8 outline-dotted hover:outline outline-2\"><input type=\"checkbox\" :value=\"$store.global.syncPageByWS\" x-on:click=\"$store.global.syncPageByWS =!$store.global.syncPageByWS\" class=\"sr-only peer\"><div class=\"relative w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 dark:peer-focus:ring-blue-800 rounded-full peer dark:bg-gray-700 peer-checked:after:translate-x-full rtl:peer-checked:after:-translate-x-full peer-checked:after:border-white after:content-[&#39;&#39;] after:absolute after:top-[2px] after:start-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all dark:border-gray-600 peer-checked:bg-blue-600\"></div><span class=\"text-sm font-medium ms-3\" x-text=\"i18next.t(&#39;SyncPage&#39;)\"></span></label><!-- 保存阅读进度 SaveReadingProgress --><label class=\"inline-flex items-center w-full my-4 cursor-pointer outline outline-offset-8 outline-dotted hover:outline outline-2\"><input type=\"checkbox\" :value=\"$store.flip.saveReadingProgress\" x-on:click=\"$store.flip.saveReadingProgress =!$store.flip.saveReadingProgress\" class=\"sr-only peer\"><div class=\"relative w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 dark:peer-focus:ring-blue-800 rounded-full peer dark:bg-gray-700 peer-checked:after:translate-x-full rtl:peer-checked:after:-translate-x-full peer-checked:after:border-white after:content-[&#39;&#39;] after:absolute after:top-[2px] after:start-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all dark:border-gray-600 peer-checked:bg-blue-600\"></div><span class=\"text-sm font-medium ms-3\" x-text=\"i18next.t(&#39;SavePageNum&#39;)\"></span></label><!-- 右开本（日漫模式）  --><label class=\"inline-flex items-center w-full my-4 cursor-pointer outline outline-offset-8 outline-dotted hover:outline outline-2\"><input type=\"checkbox\" :value=\"$store.flip.rightToLeft\" x-on:click=\"$store.flip.rightToLeft =!$store.flip.rightToLeft\" class=\"sr-only peer\"><div class=\"relative w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 dark:peer-focus:ring-blue-800 rounded-full peer dark:bg-gray-700 peer-checked:after:translate-x-full rtl:peer-checked:after:-translate-x-full peer-checked:after:border-white after:content-[&#39;&#39;] after:absolute after:top-[2px] after:start-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all dark:border-gray-600 peer-checked:bg-blue-600\"></div><span class=\"text-sm font-medium ms-3\" x-text=\"$store.flip.rightToLeft?i18next.t(&#39;LeftScreenToNext&#39;):i18next.t(&#39;RightScreenToNext&#39;)\"></span></label><!-- 单页模式  --><label class=\"inline-flex items-center w-full my-4 cursor-pointer outline outline-offset-8 outline-dotted hover:outline outline-2\"><input type=\"checkbox\" :value=\"$store.flip.doublePageMode\" x-on:click=\"$store.flip.doublePageMode =!$store.flip.doublePageMode\" class=\"sr-only peer\"><div class=\"relative w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 dark:peer-focus:ring-blue-800 rounded-full peer dark:bg-gray-700 peer-checked:after:translate-x-full rtl:peer-checked:after:-translate-x-full peer-checked:after:border-white after:content-[&#39;&#39;] after:absolute after:top-[2px] after:start-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all dark:border-gray-600 peer-checked:bg-blue-600\"></div><span class=\"text-sm font-medium ms-3\" x-text=\"$store.flip.doublePageMode?i18next.t(&#39;DoublePageMode&#39;):i18next.t(&#39;SinglePageMode&#39;)\"></span></label><!-- 自动隐藏工具栏  --><label class=\"inline-flex items-center w-full my-4 cursor-pointer outline outline-offset-8 outline-dotted hover:outline outline-2\"><input type=\"checkbox\" :value=\"$store.flip.autoHideToolbar\" x-on:click=\"$store.flip.autoHideToolbar =!$store.flip.autoHideToolbar;showToolbar();\" class=\"sr-only peer\"><div class=\"relative w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 dark:peer-focus:ring-blue-800 rounded-full peer dark:bg-gray-700 peer-checked:after:translate-x-full rtl:peer-checked:after:-translate-x-full peer-checked:after:border-white after:content-[&#39;&#39;] after:absolute after:top-[2px] after:start-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all dark:border-gray-600 peer-checked:bg-blue-600\"></div><span class=\"text-sm font-medium ms-3\" x-text=\"i18next.t(&#39;AutoHideToolbar&#39;)\"></span></label><!-- 显示页数  --><label class=\"inline-flex items-center w-full my-4 cursor-pointer outline outline-offset-8 outline-dotted hover:outline outline-2\"><input type=\"checkbox\" :value=\"$store.flip.showPageNum\" x-on:click=\"$store.flip.showPageNum =!$store.flip.showPageNum\" class=\"sr-only peer\"><div class=\"relative w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 dark:peer-focus:ring-blue-800 rounded-full peer dark:bg-gray-700 peer-checked:after:translate-x-full rtl:peer-checked:after:-translate-x-full peer-checked:after:border-white after:content-[&#39;&#39;] after:absolute after:top-[2px] after:start-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all dark:border-gray-600 peer-checked:bg-blue-600\"></div><span class=\"text-sm font-medium ms-3\" x-text=\"i18next.t(&#39;ShowPageNum&#39;)\"></span></label>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
