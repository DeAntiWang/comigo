(function(e){function t(t){for(var a,i,r=t[0],l=t[1],c=t[2],g=0,h=[];g<r.length;g++)i=r[g],Object.prototype.hasOwnProperty.call(n,i)&&n[i]&&h.push(n[i][0]),n[i]=0;for(a in l)Object.prototype.hasOwnProperty.call(l,a)&&(e[a]=l[a]);u&&u(t);while(h.length)h.shift()();return s.push.apply(s,c||[]),o()}function o(){for(var e,t=0;t<s.length;t++){for(var o=s[t],a=!0,r=1;r<o.length;r++){var l=o[r];0!==n[l]&&(a=!1)}a&&(s.splice(t--,1),e=i(i.s=o[0]))}return e}var a={},n={app:0},s=[];function i(t){if(a[t])return a[t].exports;var o=a[t]={i:t,l:!1,exports:{}};return e[t].call(o.exports,o,o.exports,i),o.l=!0,o.exports}i.m=e,i.c=a,i.d=function(e,t,o){i.o(e,t)||Object.defineProperty(e,t,{enumerable:!0,get:o})},i.r=function(e){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(e,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(e,"__esModule",{value:!0})},i.t=function(e,t){if(1&t&&(e=i(e)),8&t)return e;if(4&t&&"object"===typeof e&&e&&e.__esModule)return e;var o=Object.create(null);if(i.r(o),Object.defineProperty(o,"default",{enumerable:!0,value:e}),2&t&&"string"!=typeof e)for(var a in e)i.d(o,a,function(t){return e[t]}.bind(null,a));return o},i.n=function(e){var t=e&&e.__esModule?function(){return e["default"]}:function(){return e};return i.d(t,"a",t),t},i.o=function(e,t){return Object.prototype.hasOwnProperty.call(e,t)},i.p="";var r=window["webpackJsonp"]=window["webpackJsonp"]||[],l=r.push.bind(r);r.push=t,r=r.slice();for(var c=0;c<r.length;c++)t(r[c]);var u=l;s.push([0,"chunk-vendors"]),o()})({0:function(e,t,o){e.exports=o("56d7")},"034f":function(e,t,o){"use strict";o("85ec")},"0a90":function(e,t,o){},"3bc4":function(e,t,o){"use strict";o("d1e9")},"4e71":function(e,t,o){"use strict";o("d7e9")},"56d7":function(e,t,o){"use strict";o.r(t);o("4de4"),o("96cf");var a=o("1da1"),n=(o("e260"),o("e6cf"),o("cca6"),o("a79d"),o("2b0e")),s=function(){var e=this,t=e.$createElement,o=e._self._c||t;return o("div",{attrs:{id:"app"}},[this.$store.state.defaultSetting?o("div",["scroll"===this.$store.state.defaultSetting.template?o("ScrollTemplate"):e._e(),"sketch"===this.$store.state.defaultSetting.template?o("SketchTemplate"):e._e(),"single"===this.$store.state.defaultSetting.template?o("SinglePageTemplate"):e._e(),"double"===this.$store.state.defaultSetting.template?o("DoublePageTemplate"):e._e()],1):o("p",[e._v("loading.....")])])},i=[],r=(o("d3b7"),o("ddb0"),o("bc3a")),l=o.n(r),c=function(){var e=this,t=e.$createElement,o=e._self._c||t;return o("div",{attrs:{id:"ScrollPage"}},[o("Header",[o("h2",[this.$store.state.book.IsFolder?e._e():o("a",{attrs:{href:"raw/"+this.$store.state.book.name}},[e._v(e._s(this.$store.state.book.name)+"【Download】")]),this.$store.state.book.IsFolder?o("a",{attrs:{href:"raw/"+this.$store.state.book.name}},[e._v(e._s(this.$store.state.book.name))]):e._e()]),o("h4",[e._v("总页数："+e._s(this.$store.state.book.all_page_num))])]),e._l(this.$store.state.book.pages,(function(t,a){return o("div",{key:t.url,staticClass:"manga"},[o("img",{directives:[{name:"lazy",rawName:"v-lazy",value:t.url,expression:"page.url"}],key:a,class:e._f("check_image")(t.image_type,t.url),attrs:{H:t.height,W:t.width}}),e.showPageNum?o("p",[e._v(e._s(a+1)+"/"+e._s(e.AllPageNum))]):e._e()])})),o("p"),o("v-btn",{directives:[{name:"scroll",rawName:"v-scroll",value:e.onScroll,expression:"onScroll"},{name:"show",rawName:"v-show",value:e.btnFlag,expression:"btnFlag"}],attrs:{fab:"",color:"#bbcbff",bottom:"",right:""},on:{click:e.toTop}},[e._v("▲")]),e._t("default")],2)},u=[],g=(o("25f0"),o("93d3")),h={components:{Header:g["a"]},data:function(){return{page_mode:"multi",btnFlag:!1,showPageNum:!1,duration:300,offset:0,easing:"easeInOutCubic",AllPageNum:this.$store.state.book.all_page_num,message:{user_uuid:"",server_status:"",now_book_uuid:"",read_percent:0,msg:""}}},mounted:function(){this.initPage()},destroyed:function(){},methods:{initPage:function(){this.$cookies.keys()},getBook:function(){return this.$store.state.book},getNumber:function(e){this.page=e,console.log(e)},onScroll:function(e){if("undefined"!==typeof window){var t=window.pageYOffset||e.target.scrollTop||0;this.btnFlag=t>20}},toTop:function(){this.$vuetify.goTo(0)},initWebSocket:function(){this.$socket.onopen=this.websocketonopen,this.$socket.onerror=this.websocketonerror,this.$socket.onmessage=this.websocketonmessage,this.$socket.onclose=this.websocketclose,this.hint="连接建立"},websocketonopen:function(e){this.hint="连接成功",console.log("连接建立",e)},websocketonerror:function(e){this.hint="连接出错",this.initWebSocket(),console.log("Connection Error !!!",e)},websocketonmessage:function(e){console.log(e),this.msgList.push(JSON.parse(e.data)),this.hint="接收消息"},onChangeBook:function(e,t){this.message.now_book_uuid=t,this.message.msg="ChangeBook",this.$socket.send(JSON.stringify(this.message)),this.getBook()},websocketsend:function(e){this.$socket.send(JSON.stringify(this.message)),console.log(this.$socket.readyState,e)},websocketclose:function(e){var t=this;this.hint="连接断开",console.log("断开连接",e);var o=e.code,a=e.reason,n=e.wasClean;console.log(o,a,n);var s=setInterval((function(){t.$socket.onopen(),0==e.target.readyState&&clearInterval(s)}),3e3)}},filters:{check_image:function(e,t){if(e=e.toString(),"SinglePage"==e||"DoublePage"==e)return e;function o(e){var t=new Image;if(t.src=e,t.complete)return t.width<t.height?"SinglePage":"DoublePage";t.onload=function(){return t.onload=null,t.width<t.height?"SinglePage":"DoublePage"}}return""==e&&console.log("图片信息为空，开始本地JS分析"+t),e=o(t),e}}},p=h,m=(o("9803"),o("2877")),d=o("6544"),_=o.n(d),k=o("8336"),f=o("269a"),b=o.n(f),w=o("f977"),v=Object(m["a"])(p,c,u,!1,null,null,null),P=v.exports;_()(v,{VBtn:k["a"]}),b()(v,{Scroll:w["a"]});var $=o("d47c"),y=function(){var e=this,t=e.$createElement,o=e._self._c||t;return o("div",{attrs:{id:"SinglePageTemplate"}},[e.showHeader?o("Header",[o("h2",[this.$store.state.book.IsFolder?e._e():o("a",{attrs:{href:"raw/"+this.$store.state.book.name}},[e._v(e._s(this.$store.state.book.name)+"【Download】")]),this.$store.state.book.IsFolder?o("a",{attrs:{href:"raw/"+this.$store.state.book.name}},[e._v(e._s(this.$store.state.book.name))]):e._e()])]):e._e(),o("div",{staticClass:"single_page_main"},[e.now_page<=this.$store.state.book.all_page_num&&e.now_page>=1?o("img",{attrs:{"lazy-src":"/resources/favicon.ico",src:this.$store.state.book.pages[e.now_page-1].url},on:{click:function(t){return e.addPage(1)}}}):e._e(),o("img")]),o("v-pagination",{attrs:{length:this.$store.state.book.all_page_num,"total-visible":10},on:{input:e.toPage},model:{value:e.now_page,callback:function(t){e.now_page=t},expression:"now_page"}}),e._t("default")],2)},S=[],T={components:{Header:g["a"]},data:function(){return{now_page:1,showHeader:!1,showPagination:!0,alert:!1,easing:"easeInOutCubic",book:null,bookshelf:null,defaultSetting:null}},mounted:function(){this.book=this.$store.state.book,this.bookshelf=this.$store.state.bookshelf,this.defaultSetting=this.$store.state.defaultSetting,window.addEventListener("keyup",this.handleKeyup)},destroyed:function(){window.removeEventListener("keyup",this.handleKeyup)},methods:{initPage:function(){},addPage:function(e){this.now_page+e<this.book.all_page_num&&this.now_page+e>=1&&(this.now_page=this.now_page+e)},toPage:function(e){e<=this.book.all_page_num&&e>=1&&(this.now_page=e)},handleKeyup:function(e){var t=e||window.event||arguments.callee.caller.arguments[0];if(t)switch(t.key){case"ArrowUp":case"PageUp":case"ArrowLeft":this.addPage(-1);break;case"Space":case"ArrowDown":case"PageDown":case"ArrowRight":this.addPage(1);break;case"Home":this.toPage(1);break;case"End":this.toPage(this.book.all_page_num);break;case"Ctrl":break}},handleScroll:function(){document.body.scrollTop||document.documentElement.scrollTop}}},x=T,O=(o("9f5f"),o("891e")),A=Object(m["a"])(x,y,S,!1,null,null,null),C=A.exports;_()(A,{VPagination:O["a"]});var D=function(){var e=this,t=e.$createElement,o=e._self._c||t;return o("div",{attrs:{id:"DoublePageTemplate"}},[e.showHeader?o("Header",[o("h2",[this.$store.state.book.IsFolder?e._e():o("a",{attrs:{href:"raw/"+this.$store.state.book.name}},[e._v(e._s(this.$store.state.book.name)+"【Download】现在时刻："+e._s(e.currentTime))]),this.$store.state.book.IsFolder?o("a",{attrs:{href:"raw/"+this.$store.state.book.name}},[e._v(e._s(this.$store.state.book.name)+"现在时刻："+e._s(e.currentTime))]):e._e()])]):e._e(),o("div",{staticClass:"double_page_main"},[e.page_mark<this.$store.state.book.all_page_num?o("img",{attrs:{id:"image1","lazy-src":"/resources/favicon.ico",src:this.$store.state.book.pages[e.page_mark].url},on:{click:e.nextPageClick}}):e._e(),o("img"),e.page_mark-1>=0&&e.page_mark<this.$store.state.book.all_page_num&&"SinglePage"==this.$store.state.book.pages[e.page_mark].image_type&&"SinglePage"==this.$store.state.book.pages[e.page_mark-1].image_type?o("img",{attrs:{id:"image2","lazy-src":"/resources/favicon.ico",src:this.$store.state.book.pages[e.page_mark-1].url},on:{click:e.previousPageClick}}):e._e(),o("img")]),e.showPagination?o("v-pagination",{attrs:{length:this.$store.state.book.all_page_num-1,"total-visible":15},on:{input:e.toPage},model:{value:e.page_mark,callback:function(t){e.page_mark=t},expression:"page_mark"}}):e._e(),e._t("default")],2)},H=[],N=(o("b0c0"),{components:{Header:g["a"]},data:function(){return{localbook:{name:this.$store.state.book.name,all_image_num:this.$store.state.book.all_page_num,images:this.$store.state.book.pages,pages:null,all_page_num:0},bookshelf:null,defaultSetting:null,page_mark:0,showHeader:!1,showPagination:!0,AllPageNum:this.$store.state.book.all_page_num-1,time_cont:0,alert:!1,easing:"easeInOutCubic",timer:"",currentTime:new Date}},methods:{initLocalBook:function(){},initPageMark:function(){this.$store.state.book.all_page_num<2?this.page_mark=0:"SinglePage"==this.$store.state.book.pages[0].image_type&&"SinglePage"==this.$store.state.book.pages[1].image_type?this.page_mark=1:this.page_mark=0},toPage:function(e){(e>this.$store.state.book.all_page_num||e<0)&&console.log("page_mark error",e),this.page_mark=e,console.log(e)},nextPageClick:function(){this.page_mark>=this.AllPageNum||this.page_mark<0?console.log(this.page_mark):this.page_mark!=this.AllPageNum-1?this.page_mark!=this.AllPageNum-2?"SinglePage"==this.$store.state.book.pages[this.page_mark+1].image_type&&"SinglePage"==this.$store.state.book.pages[this.page_mark+2].image_type?(this.page_mark=this.page_mark+2,console.log(this.page_mark)):(this.page_mark=this.page_mark+1,console.log(this.page_mark)):"SinglePage"==this.$store.state.book.pages[this.AllPageNum-1].image_type&&"SinglePage"==this.$store.state.book.pages[this.AllPageNum-2].image_type?(this.page_mark=this.page_mark+2,console.log(this.page_mark)):(this.page_mark=this.page_mark+1,console.log(this.page_mark)):this.page_mark=this.page_mark+1},previousPageClick:function(){if(this.page_mark<=0)console.log(this.page_mark);else{if(1==this.page_mark)return this.page_mark=this.page_mark-1,void console.log(this.page_mark);if(2==this.page_mark)return this.page_mark=this.page_mark-1,void console.log(this.page_mark);if(this.page_mark>=this.AllPageNum)return this.page_mark=this.AllPageNum-1,void console.log(this.page_mark);if("SinglePage"==this.$store.state.book.pages[this.page_mark-2].image_type&&"SinglePage"==this.$store.state.book.pages[this.page_mark-2-1].image_type)return this.page_mark=this.page_mark-2,void console.log(this.page_mark);this.page_mark=this.page_mark-1,console.log(this.page_mark)}},nextPage:function(){this.page_mark>this.$store.state.book.all_page_num?console.log(this.page_mark):this.page_mark!=this.$store.state.book.all_page_num?this.nextPageClick():console.log(this.page_mark)},previousPage:function(){this.page_mark>this.$store.state.book.all_page_num?console.log(this.page_mark):this.page_mark!=this.$store.state.book.all_page_num?this.previousPageClick():this.page_mark=this.page_mark-1},handleKeyup:function(e){var t=e||window.event||arguments.callee.caller.arguments[0];if(t)switch(t.key){case"PageUp":case"ArrowUp":case"ArrowLeft":this.previousPage();break;case"Space":case"ArrowDown":case"PageDown":case"ArrowRight":this.nextPage();break;case"Home":this.toPage(1);break;case"End":this.toPage(this.$store.state.book.all_page_num-1);break;case"Ctrl":break}},handleScroll:function(){document.body.scrollTop||document.documentElement.scrollTop}},created:function(){var e=this;this.timer=setInterval((function(){var t=new Date,o=t.getFullYear(),a=t.getMonth()+1,n=t.getDate();a>=1&&a<=9&&(a="0"+a),n>=0&&n<=9&&(n="0"+n);var s=o+" 年 "+a+" 月 "+n+" 日 ",i=t.getHours();i>=0&&i<=9&&(i="0"+i);var r=t.getMinutes();r>=0&&r<=9&&(r="0"+r);var l=t.getSeconds();l>=0&&l<=9&&(l="0"+l),e.currentTime=s+" "+i+":"+r+":"+l}),1e3)},mounted:function(){this.time_cont=0,this.$cookies.keys(),this.initPageMark(),window.addEventListener("keyup",this.handleKeyup)},destroyed:function(){window.removeEventListener("keyup",this.handleKeyup),this.timer&&clearInterval(this.timer)}}),j=N,E=(o("810a"),Object(m["a"])(j,D,H,!1,null,null,null)),I=E.exports;_()(E,{VPagination:O["a"]});var F={name:"app",components:{ScrollTemplate:P,SketchTemplate:$["default"],SinglePageTemplate:C,DoublePageTemplate:I},data:function(){return{book:null,bookshelf:{},defaultSetting:{},now_page:1,duration:300,offset:0,easing:"easeInOutCubic",message:{user_uuid:"",server_status:"",now_book_uuid:"",read_percent:0,msg:""}}},mounted:function(){this.initPage(),this.$cookies.keys()},destroyed:function(){},methods:{initPage:function(){var e=this;l.a.get("/book.json").then((function(t){return e.$store.state.book=t.data})).finally(this.book=this.$store.book),l.a.get("/setting.json").then((function(t){return e.$store.state.defaultSetting=t.data})).finally(this.defaultSetting=this.$store.defaultSetting),l.a.get("/bookshelf.json").then((function(t){return e.$store.state.bookshelf=t.data})).finally(this.bookshelf=this.$store.bookshelf)},getNumber:function(e){this.page=e,console.log(e)}}},L=F,M=(o("034f"),Object(m["a"])(L,s,i,!1,null,null,null)),B=M.exports,K=o("8c4f");n["a"].use(K["a"]);var V=[{path:"/",name:"Scrool",component:P},{path:"/sketch",name:"Sketch",component:function(){return Promise.resolve().then(o.bind(null,"d47c"))}}],z=new K["a"]({routes:V}),J=z,U=o("caf9"),W=o("b408"),R=o.n(W),Y=o("f309");n["a"].use(Y["a"]);var q=new Y["a"]({}),G=o("2b27"),Q=o.n(G),X=o("2f62"),Z=void 0;n["a"].use(R.a,"ws://"+document.location.host+"/ws",{reconnection:!0,reconnectionAttempts:500,reconnectionDelay:1e3}),n["a"].config.productionTip=!1,n["a"].use(U["a"],{preLoad:4.5,attempt:10}),n["a"].use(Q.a),n["a"].$cookies.config("30d"),n["a"].use(X["a"]);var ee=new X["a"].Store({state:{count:0,todos:[{id:1,text:"...",done:!0},{id:2,text:"...",done:!1}],now_page:1,book:{name:"loading",page_num:1,pages:[{height:500,width:449,url:"/resources/favicon.ico",class:"Vertical"}]},bookshelf:{},defaultSetting:{default_page_template:"scroll",sketch_count_seconds:90},message:{user_uuid:"",server_status:"",now_book_uuid:"",read_percent:0,msg:""}},getters:{doneTodos:function(e){return e.todos.filter((function(e){return e.done}))},now_page:function(e){return e.now_page},book:function(e){return console.log(Z.state.book),e.book},bookshelf:function(e){return e.bookshelf},defaultSetting:function(e){return e.defaultSetting},message:function(e){return e.message}},mutations:{increment:function(e){e.count++},syncBookDate:function(e,t){e.book=t.msg,console.log(e.book),console.log("syncBookDate run")}},actions:{incrementAction:function(e){e.commit("increment")},getMessageAction:function(e){return Object(a["a"])(regeneratorRuntime.mark((function t(){var o,a;return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:return t.next=2,l.a.get("/book.json").then((function(e){return e.data}),(function(){return""}));case 2:o=t.sent,a={message:o},e.commit("syncBookDate",a);case 5:case"end":return t.stop()}}),t)})))()}}});new n["a"]({router:J,vuetify:q,store:ee,render:function(e){return e(B)}}).$mount("#app")},6528:function(e,t,o){},"810a":function(e,t,o){"use strict";o("0a90")},"85ec":function(e,t,o){},"93d3":function(e,t,o){"use strict";var a=function(){var e=this,t=e.$createElement,o=e._self._c||t;return o("header",{staticClass:"header"},[e._t("default")],2)},n=[],s={name:"Header",data:function(){return{mybook:this.book,rangeValue:500}},methods:{changeEvent:function(){console.log("change to:"+this.rangeValue)}}},i=s,r=(o("4e71"),o("2877")),l=Object(r["a"])(i,a,n,!1,null,"c3c1e8ec",null);t["a"]=l.exports},9803:function(e,t,o){"use strict";o("6528")},"9f5f":function(e,t,o){"use strict";o("f2b5")},d1e9:function(e,t,o){},d47c:function(e,t,o){"use strict";o.r(t);var a=function(){var e=this,t=e.$createElement,o=e._self._c||t;return o("div",{attrs:{id:"SketchPage"}},[e.showHeader?o("Header",[o("h2",[o("a",{attrs:{href:"raw/"+this.$store.state.book.name}},[e._v(e._s(this.$store.state.book.name)+"现在时刻："+e._s(e.currentTime))])])]):e._e(),o("div",{staticClass:"sketch_main"},[o("div",{attrs:{id:"SketchHint"}},[o("p",[e._v(" "+e._s(this.$store.state.defaultSetting.sketch_count_seconds)+"秒翻页,"+e._s(e.getNowCount())+"⏳ ")])]),o("img",{attrs:{"lazy-src":"/resources/favicon.ico",src:this.$store.state.book.pages[e.now_page-1].url},on:{click:function(t){return e.addPage(1)}}}),o("img"),o("div",{attrs:{id:"SketchHint"}},[o("p",[e._v("🕒"+e._s(e.currentTime))])])]),e.showPagination?o("v-pagination",{attrs:{length:this.$store.state.book.all_page_num,"total-visible":10},on:{input:e.toPage},model:{value:e.now_page,callback:function(t){e.now_page=t},expression:"now_page"}}):e._e(),e._t("default")],2)},n=[],s=(o("d3b7"),o("ddb0"),o("93d3")),i={components:{Header:s["a"]},data:function(){return{time_cont:1,WaitSeconds:this.$store.state.defaultSetting.sketch_count_seconds,book:null,bookshelf:null,defaultSetting:null,showHeader:!1,showPagination:!0,now_page:1,alert:!1,easing:"easeInOutCubic",timer:"",currentTime:null}},created:function(){var e=this;this.timer=setInterval((function(){var t=new Date,o=t.getFullYear(),a=t.getMonth()+1,n=t.getDate();a>=1&&a<=9&&(a="0"+a),n>=0&&n<=9&&(n="0"+n);var s=o+"年"+a+"月"+n+"日",i=t.getHours();i>=0&&i<=9&&(i="0"+i);var r=t.getMinutes();r>=0&&r<=9&&(r="0"+r);var l=t.getSeconds();l>=0&&l<=9&&(l="0"+l),e.currentTime=i+":"+r+":"+l,console.log(s+"time_cont："+e.time_cont),e.time_cont<e.WaitSeconds?e.time_cont++:(e.time_cont=1,console.log("时间到，翻页："+e.currentTime+"秒"),e.now_page<e.$store.state.book.all_page_num?e.now_page+=1:e.now_page=1)}),1e3)},mounted:function(){this.time_cont=0,this.$cookies.keys(),window.addEventListener("keyup",this.handleKeyup)},destroyed:function(){window.removeEventListener("keyup",this.handleKeyup),this.timer&&clearInterval(this.timer)},methods:{initPage:function(){},getWaitSeconds:function(){return this.$store.state.defaultSetting.sketch_count_seconds},getNowCount:function(){var e=this.time_cont;return e>=0&&e<=9&&(e="0"+e),e},addPage:function(e){this.now_page+e<this.$store.state.book.all_page_num&&this.now_page+e>=1&&(this.now_page=this.now_page+e)},toPage:function(e){this.now_page=e,console.log(e)},handleKeyup:function(e){var t=e||window.event||arguments.callee.caller.arguments[0];if(t)switch(t.key){case"PageUp":case"ArrowUp":case"ArrowLeft":this.addPage(-1);break;case"Space":case"ArrowDown":case"PageDown":case"ArrowRight":this.addPage(1);break;case"Home":this.toPage(1);break;case"End":this.toPage(this.$store.state.book.all_page_num-1);break;case"Ctrl":break}},handleScroll:function(){document.body.scrollTop||document.documentElement.scrollTop}}},r=i,l=(o("3bc4"),o("2877")),c=o("6544"),u=o.n(c),g=o("891e"),h=Object(l["a"])(r,a,n,!1,null,null,null);t["default"]=h.exports;u()(h,{VPagination:g["a"]})},d7e9:function(e,t,o){},f2b5:function(e,t,o){}});
//# sourceMappingURL=app.ae701700.js.map