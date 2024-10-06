// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.778
package pages

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"fmt"
	"github.com/yumenaka/comigo/entity"
	"github.com/yumenaka/comigo/htmx/state"
	"github.com/yumenaka/comigo/htmx/templates/components"
	"strconv"
	"strings"
)

// js代码,滚动到顶部,显示返回顶部按钮,获取鼠标位置,决定是否打开设置面板等
func ScrollScripts() templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_ScrollScripts_49dc`,
		Function: `function __templ_ScrollScripts_49dc(){//滚动到顶部
function scrollToTop(scrollDuration) {
    let scrollStep = -window.scrollY / (scrollDuration / 15),
        scrollInterval = setInterval(function () {
            if (window.scrollY !== 0) {
                window.scrollBy(0, scrollStep);
            }
            else clearInterval(scrollInterval);
        }, 15);
}
// Button ID为BackTopButton的元素，点击后滚动到顶部
document.getElementById("BackTopButton").addEventListener("click", function () {
    scrollToTop(500);
});

//滚动到一定位置显示返回顶部按钮
let scrollTopSave = 0
let scrollDownFlag = false
let step = 0
function onScroll() {
    let scrollTop = document.documentElement.scrollTop || document.body.scrollTop;
    scrollDownFlag = scrollTop > scrollTopSave;
    //防手抖,小于一定数值状态就不变 Math.abs()会导致报错
    step = scrollTopSave - scrollTop;
    // console.log("this.scrollDownFlag:",this.scrollDownFlag,"scrollTop:",scrollTop,"step:", step);
    scrollTopSave = scrollTop
    if (step < -10 || step > 10) {
        showBackTopFlag = ((scrollTop > 400) && !scrollDownFlag);
        if (showBackTopFlag) {
            document.getElementById("BackTopButton").style.display = "block";
        } else {
            document.getElementById("BackTopButton").style.display = "none";
        }
    }
}
window.addEventListener("scroll", onScroll);

//可见区域变化的时候改变页面状态
function onResize() {
    this.ScrollModeConfig.imageMaxWidth = window.innerWidth
    this.clientWidth = document.documentElement.clientWidth
    this.clientHeight = document.documentElement.clientHeight
    // var aspectRatio = window.innerWidth / window.innerHeight
    this.aspectRatio = this.clientWidth / this.clientHeight
    // 为了调试的时候方便,阈值是正方形
    if (this.aspectRatio > (19 / 19)) {
        this.ScrollModeConfig.isLandscapeMode = true
        this.ScrollModeConfig.isPortraitMode = false
    } else {
        this.ScrollModeConfig.isLandscapeMode = false
        this.ScrollModeConfig.isPortraitMode = true
    }
}
//文档视图调整大小时触发 resize 事件。 https://developer.mozilla.org/zh-CN/docs/Web/API/Window/resize_event
window.addEventListener("resize", this.onResize);

//获取鼠标位置,决定是否打开设置面板
function onMouseClick(e) {
    this.clickX = e.x //获取鼠标的X坐标（鼠标与屏幕左侧的距离,单位为px）
    this.clickY = e.y //获取鼠标的Y坐标（鼠标与屏幕顶部的距离,单位为px）
    //浏览器的视口,不包括工具栏和滚动条:
    let innerWidth = window.innerWidth
    let innerHeight = window.innerHeight
    //设置区域为正方形，边长按照宽或高里面，比较小的值决定
    const setArea = 0.15;
    // innerWidth >= innerHeight 的情况下
    let MinY = innerHeight * (0.5 - setArea);
    let MaxY = innerHeight * (0.5 + setArea);
    let MinX = innerWidth * 0.5 - (MaxY - MinY) * 0.5;
    let MaxX = innerWidth * 0.5 + (MaxY - MinY) * 0.5;
    if (innerWidth < innerHeight) {
        MinX = innerWidth * (0.5 - setArea);
        MaxX = innerWidth * (0.5 + setArea);
        MinY = innerHeight * 0.5 - (MaxX - MinX) * 0.5;
        MaxY = innerHeight * 0.5 + (MaxX - MinX) * 0.5;
    }
    //在设置区域
    let inSetArea = false
    if ((this.clickX > MinX && this.clickX < MaxX) && (this.clickY > MinY && this.clickY < MaxY)) {
        console.log("点中了设置区域！");
        inSetArea = true
    }
    if (inSetArea) {
        //获取ID为 OpenSettingButton的元素，然后模拟点击
		document.getElementById("OpenSettingButton").click();
    }
}
//获取鼠标位置,决定是否打开设置面板
function onMouseMove(e) {
    this.clickX = e.x //获取鼠标的X坐标（鼠标与屏幕左侧的距离,单位为px）
    this.clickY = e.y //获取鼠标的Y坐标（鼠标与屏幕顶部的距离,单位为px）
    //浏览器的视口,不包括工具栏和滚动条:
    let innerWidth = window.innerWidth
    let innerHeight = window.innerHeight
    //设置区域为正方形，边长按照宽或高里面，比较小的值决定
    const setArea = 0.15;
    // innerWidth >= innerHeight 的情况下
    let MinY = innerHeight * (0.5 - setArea);
    let MaxY = innerHeight * (0.5 + setArea);
    let MinX = innerWidth * 0.5 - (MaxY - MinY) * 0.5;
    let MaxX = innerWidth * 0.5 + (MaxY - MinY) * 0.5;
    if (innerWidth < innerHeight) {
        MinX = innerWidth * (0.5 - setArea);
        MaxX = innerWidth * (0.5 + setArea);
        MinY = innerHeight * 0.5 - (MaxX - MinX) * 0.5;
        MaxY = innerHeight * 0.5 + (MaxX - MinX) * 0.5;
    }
    //在设置区域
    let inSetArea = false
    if ((this.clickX > MinX && this.clickX < MaxX) && (this.clickY > MinY && this.clickY < MaxY)) {
        inSetArea = true
    }
    if (inSetArea) {
        //console.log("在设置区域！");
        e.currentTarget.style.cursor = 'url(/static/images/SettingsOutline.png), pointer';
    } else {
        e.currentTarget.style.cursor = '';
    }
}
//获取ID为 mouseMoveArea 的元素
let mouseMoveArea = document.getElementById("mouseMoveArea")
// 鼠标移动的时候触发移动事件
mouseMoveArea.addEventListener('mousemove', onMouseMove)
// 点击的时候触发点击事件
mouseMoveArea.addEventListener('click', onMouseClick)
// 触摸的时候也触发点击事件
mouseMoveArea.addEventListener('touchstart', onMouseClick)
}`,
		Call:       templ.SafeScript(`__templ_ScrollScripts_49dc`),
		CallInline: templ.SafeScriptInline(`__templ_ScrollScripts_49dc`),
	}
}

// ScrollMainArea 定义 BodyHTML
func ScrollMainArea(s *state.GlobalState, book *entity.Book) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"mouseMoveArea\" class=\"flex flex-col items-center justify-center flex-1 w-full max-w-full pt-2 bg-base-100 text-base-content\" :class=\"(theme.toString() ===&#39;light&#39;||theme.toString() ===&#39;dark&#39;||theme.toString() ===&#39;retro&#39;||theme.toString() ===&#39;lofi&#39;||theme.toString() ===&#39;nord&#39;) &amp;&amp; &#39;bg-base-300&#39;\"><!-- Alpine.js 动态CSS，只支持内联写法 --><!-- Alpine.js 的 v-if 需要用template包裹起来，原因参照： https://alpinejs.dev/directives/if --><!-- https://htmx.org/docs/#triggers --><!-- https://htmx.org/docs/#swapping --><!-- hx-get 用于获取图片的URL，hx-trigger 用于触发加载，hx-swap 用于替换元素，innerHTML默认值，将内容放在目标元素内 outerHTML用返回的内容替换整个目标元素  hx-target 用于指定目标元素 -->")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for key, image := range book.Pages.Images {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div x-data=\"{\n                            orientation: &#39;&#39;,\n                            updateOrientation() {\n                                this.orientation = (window.innerWidth / window.innerHeight &gt;= 1) ? &#39;landscape&#39; : &#39;portrait&#39;;\n                            }\n                        }\" class=\"flex flex-col justify-start w-full max-w-full m-0 rounded item-center\" :style=\"{ marginBottom: $store.scroll.marginBottomOnScrollMode + &#39;px&#39; }\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if strings.Contains(image.Url, ".html") && !strings.Contains(image.Url, "hidden.") {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div hx-get=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var2 string
				templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(image.Url)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/pages/scroll.templ`, Line: 168, Col: 28}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-trigger=\"load\" hx-swap=\"innerHTML\" class=\"w-full m-0\"></div>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			if !strings.Contains(image.Url, "hidden.") && !strings.Contains(image.Url, ".html") {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<img class=\"w-full\" src=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var3 string
				templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(image.Url)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/pages/scroll.templ`, Line: 173, Col: 21}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" x-data=\"{\n                                    isDoublePage: false,\n                                    imageWidth: &#39;&#39;\n                                }\" x-init=\"updateOrientation();\" @load=\"\n                            if ($event.target.naturalWidth &gt; $event.target.naturalHeight) {\n                                // 双页\n                                isDoublePage = true;\n                                $el.classList.add(&#39;double&#39;);\n                                $el.classList.remove(&#39;single&#39;);\n                            } else {\n                                // 单页\n                                isDoublePage = false;\n                                $el.classList.add(&#39;single&#39;);\n                                $el.classList.remove(&#39;double&#39;);\n                            }\" @resize.window=\"updateOrientation()\" :style=\"{ width: orientation === &#39;landscape&#39;?(Alpine.store(&#39;scroll&#39;).widthUseFixedValue? (isDoublePage ? Alpine.store(&#39;scroll&#39;).doublePageWidth_PX +&#39;px&#39;: Alpine.store(&#39;scroll&#39;).singlePageWidth_PX +&#39;px&#39;): (isDoublePage ? Alpine.store(&#39;scroll&#39;).doublePageWidth_Percent + &#39;%&#39;: Alpine.store(&#39;scroll&#39;).singlePageWidth_Percent + &#39;%&#39;)): &#39;100%&#39;, maxWidth: &#39;100%&#39;}\" alt=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var4 string
				templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.Itoa(key))
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/pages/scroll.templ`, Line: 193, Col: 29}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<template x-if=\"$store.scroll.showPageNum\"><div class=\"w-full mt-0 mb-2 text-lg font-semibold text-center page_hint \">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var5 string
			templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("%d / %d", key+1, book.BookInfo.PageCount))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/pages/scroll.templ`, Line: 197, Col: 135}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></template></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><button id=\"BackTopButton\" style=\"display: none\" class=\"fixed flex items-center justify-center w-10 h-10 text-white bg-blue-500 rounded-full shadow-lg bottom-4 right-4\"><svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" viewBox=\"0 0 512 512\"><path d=\"M256 48C141.13 48 48 141.13 48 256s93.13 208 208 208s208-93.13 208-208S370.87 48 256 48zm96 270.63l-96-96l-96 96L137.37 296L256 177.37L374.63 296z\" fill=\"currentColor\"></path></svg></button>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = ScrollScripts().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<style>\n    /* https://developer.mozilla.org/zh-CN/docs/Web/CSS/object-fit */\n    img {\n        margin: auto;\n        margin-bottom: v-bind(margin+'px');\n        box-shadow: 0px 6px 3px 0px rgba(0, 0, 0, 0.19);\n    }\n\n    .page_hint {\n        /* 文字颜色 */\n        color: #413d3d;\n        /* 文字阴影：https://www.w3school.com.cn/css/css3_shadows.asp*/\n        text-shadow: -1px 0 rgb(240, 229, 229), 0 1px rgb(253, 242, 242),\n            1px 0 rgb(206, 183, 183), 0 -1px rgb(196, 175, 175);\n    }\n\n    .LoadingImage {\n        width: 90vw;\n        max-width: 90vw;\n    }\n\n    .ErrorImage {\n        width: 90vw;\n        max-width: 90vw;\n    }\n\n    /* 横屏（显示区域）时的CSS样式,IE无效 */\n    @media screen and (min-aspect-ratio: 19/19) {\n        .SinglePageImage {\n            width: v-bind(sPWL);\n            max-width: 100%;\n        }\n\n        .DoublePageImage {\n            width: v-bind(dPWL);\n            max-width: 100%;\n        }\n    }\n\n    /* 竖屏(显示区域)CSS样式,IE无效 */\n    @media screen and (max-aspect-ratio: 19/19) {\n        .SinglePageImage {\n            width: v-bind(sPWP);\n            max-width: 100%;\n        }\n\n        .DoublePageImage {\n            /* width: 100%; */\n            width: v-bind(dPWP);\n            max-width: 100%;\n        }\n    }\n    </style>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

// 卷轴阅读，侧栏设置slot
func ScrollDrawerSlot() templ.Component {
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
		templ_7745c5c3_Var6 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var6 == nil {
			templ_7745c5c3_Var6 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!-- 滑动选择组件: 下拉模式下，漫画页面的底部间距（px）--><!-- 组件来自：https://flowbite.com/docs/forms/range/--><div class=\"relative w-full my-4 outline outline-offset-8 outline-dotted hover:outline outline-2\"><label for=\"marginBottom\" class=\"block w-full mb-2 text-sm font-medium\" x-text=\"i18next.t(&#39;marginBottomOnScrollMode&#39;) + $store.scroll.marginBottomOnScrollMode+&#39;px&#39;\"></label> <input id=\"marginBottom\" type=\"range\" min=\"0\" max=\"100\" x-model=\"$store.scroll.marginBottomOnScrollMode\" step=\"1\" class=\"w-full h-2 bg-gray-200 rounded-lg appearance-none cursor-pointer dark:bg-gray-700\"></div><!-- toggle开关：显示页数 --><label class=\"inline-flex items-center w-full my-2 my-4 cursor-pointer outline outline-offset-8 outline-dotted hover:outline outline-2\"><input type=\"checkbox\" x-model=\"$store.scroll.showPageNum\" class=\"sr-only peer\" checked><div class=\"relative w-11 h-6 bg-gray-200 rounded-full peer peer-focus:ring-4 peer-focus:ring-blue-300 dark:peer-focus:ring-blue-800 dark:bg-gray-700 peer-checked:after:translate-x-full rtl:peer-checked:after:-translate-x-full peer-checked:after:border-white after:content-[&#39;&#39;] after:absolute after:top-0.5 after:start-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all dark:border-gray-600 peer-checked:bg-blue-600\"></div><span class=\"text-sm font-medium ms-3\" x-text=\"i18next.t(&#39;showPageNum&#39;)\"></span></label><!-- toggle开关组件：横屏状态下,宽度单位是固定值还是百分比 --><!-- 组件来自：https://flowbite.com/docs/forms/toggle/ --><label class=\"inline-flex items-center w-full my-2 my-4 cursor-pointer outline outline-offset-8 outline-dotted hover:outline outline-2\"><input type=\"checkbox\" x-model:value=\"$store.scroll.widthUseFixedValue\" class=\"sr-only peer\" checked><div class=\"relative w-11 h-6 bg-gray-200 rounded-full peer peer-focus:ring-4 peer-focus:ring-blue-300 dark:peer-focus:ring-blue-800 dark:bg-gray-700 peer-checked:after:translate-x-full rtl:peer-checked:after:-translate-x-full peer-checked:after:border-white after:content-[&#39;&#39;] after:absolute after:top-0.5 after:start-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all dark:border-gray-600 peer-checked:bg-blue-600\"></div><span class=\"text-sm font-medium ms-3\" x-text=\"$store.scroll.widthUseFixedValue?i18next.t(&#39;widthUseFixedValue&#39;):i18next.t(&#39;widthUsePercent&#39;)\"></span></label><!-- 单页漫画宽度（横屏状态+宽度限制使用百分比） --><!-- 组件来自：https://flowbite.com/docs/forms/range/--><div x-show=\"!$store.scroll.widthUseFixedValue\" class=\"relative w-full my-4 outline outline-offset-8 outline-dotted hover:outline outline-2 \"><label for=\"singlePageWidth_Percent\" class=\"block w-full mb-2 text-sm font-medium\" x-text=\"i18next.t(&#39;singlePageWidth&#39;) + $store.scroll.singlePageWidth_Percent+&#39;%&#39;\"></label> <input id=\"singlePageWidth_Percent\" type=\"range\" min=\"10\" max=\"100\" x-model=\"$store.scroll.singlePageWidth_Percent\" step=\"1\" class=\"w-full h-2 bg-gray-200 rounded-lg appearance-none cursor-pointer dark:bg-gray-700\"></div><!-- 双页漫画宽度（横屏状态+宽度限制使用百分比） --><!-- 组件来自：https://flowbite.com/docs/forms/range/--><div x-show=\"!$store.scroll.widthUseFixedValue\" class=\"relative w-full my-4 outline outline-offset-8 outline-dotted hover:outline outline-2\"><label for=\"doublePageWidth_Percent\" class=\"block w-full mb-2 text-sm font-medium\" x-text=\"i18next.t(&#39;doublePageWidth&#39;) + $store.scroll.doublePageWidth_Percent+&#39;%&#39;\"></label> <input id=\"doublePageWidth_Percent\" type=\"range\" min=\"10\" max=\"100\" x-model=\"$store.scroll.doublePageWidth_Percent\" step=\"1\" class=\"w-full h-2 bg-gray-200 rounded-lg appearance-none cursor-pointer dark:bg-gray-700\"></div><!-- 单页漫画宽度（横屏状态+宽度限制使用固定值） --><!-- 组件来自：https://flowbite.com/docs/forms/range/--><div x-show=\"$store.scroll.widthUseFixedValue\" class=\"relative w-full my-4 outline outline-offset-8 outline-dotted hover:outline outline-2\"><label for=\"singlePageWidth_PX\" class=\"block w-full mb-2 text-sm font-medium\" x-text=\"i18next.t(&#39;singlePageWidth&#39;) + $store.scroll.singlePageWidth_PX+&#39; px&#39;\"></label> <input id=\"singlePageWidth_PX\" type=\"range\" min=\"100\" max=\"1600\" x-model=\"$store.scroll.singlePageWidth_PX\" step=\"20\" class=\"w-full h-2 bg-gray-200 rounded-lg appearance-none cursor-pointer dark:bg-gray-700\"></div><!-- 双页漫画宽度（横屏状态+宽度限制使用固定值） --><!-- 组件来自：https://flowbite.com/docs/forms/range/--><div x-show=\"$store.scroll.widthUseFixedValue\" class=\"relative w-full my-4 outline outline-offset-8 outline-dotted hover:outline outline-2\"><label for=\"doublePageWidth_PX\" class=\"block w-full mb-2 text-sm font-medium\" x-text=\"i18next.t(&#39;doublePageWidth&#39;) + $store.scroll.doublePageWidth_PX+&#39; px&#39;\"></label> <input id=\"doublePageWidth_PX\" type=\"range\" min=\"100\" max=\"1600\" x-model=\"$store.scroll.doublePageWidth_PX\" step=\"20\" class=\"w-full h-2 bg-gray-200 rounded-lg appearance-none cursor-pointer dark:bg-gray-700\"></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

// ScrollPage 定义 BodyHTML
func ScrollPage(s *state.GlobalState, book *entity.Book) templ.Component {
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
		templ_7745c5c3_Var7 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var7 == nil {
			templ_7745c5c3_Var7 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = components.Header(components.HeaderProps{
			Title:           components.GetPageTitle(book.BookInfo.BookID),
			ShowReturnIcon:  true,
			ReturnUrl:       getReturnUrl(book.BookInfo.BookID),
			SetDownLoadLink: false,
			InShelf:         false,
			DownLoadLink:    "",
			SetTheme:        true,
		}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = ScrollMainArea(s, book).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.Footer(s.Version).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.Drawer(s, ScrollDrawerSlot()).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.QRCode(s).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
