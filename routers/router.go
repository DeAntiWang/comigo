package routers

import (
	"embed"
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/yumenaka/comi/routers/token"
	"html/template"
	"io"
	"io/fs"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sanity-io/litter"

	"github.com/yumenaka/comi/book"
	"github.com/yumenaka/comi/common"
	"github.com/yumenaka/comi/locale"
	"github.com/yumenaka/comi/plugin"
	"github.com/yumenaka/comi/routers/handler"
	"github.com/yumenaka/comi/routers/websocket"
	"github.com/yumenaka/comi/tools"
)

// TemplateString 模板文件
//
//go:embed static/index.html
var TemplateString string

//go:embed  static
var staticFS embed.FS

//go:embed  static/assets
var staticAssetFS embed.FS

//go:embed  static/images
var staticImageFS embed.FS

// gin-jwt相关 https://github.com/appleboy/gin-jwt

// StartWebServer 启动web服务
func StartWebServer() {
	//设置 gin
	gin.SetMode(gin.ReleaseMode)
	//// 创建带有默认中间件的路由: 日志与恢复中间件
	engine := gin.Default()
	////Recovery 中间件会恢复(recovers) 任何恐慌(panics) 如果存在恐慌，中间件将会写入500 gin.Default()已经默认启用了这个中间件
	//engine.Use(gin.Recovery())
	////Logger() 以默认配置创建日志中间件，将所有请求信息按指定格式打印到标准输出。 gin.Default()已经默认启用了这个中间件
	//engine.Use(gin.Logger())

	//1、setStaticFiles
	setStaticFiles(engine)
	//2、setWebAPI
	setWebAPI(engine)
	//3、setPort
	setPort()
	//4、setWebpServer
	//setWebpServer(engine)
	//5、setFrpClient
	setFrpClient()
	//6、printCMDMessage
	printCMDMessage()
	//7、StartGinEngine 监听并启动web服务
	StartGinEngine(engine)
}

// 1、设置静态文件
func setStaticFiles(engine *gin.Engine) {
	//使用自定义的模板引擎，命名为"template-data"，为了与VUE兼容，把左右分隔符自定义为 [[ ]]
	tmpl := template.Must(template.New("template-data").Delims("[[", "]]").Parse(TemplateString))
	//使用模板
	engine.SetHTMLTemplate(tmpl)
	if common.Config.LogToFile {
		// 关闭 log 打印的字体颜色。输出到文件不需要颜色
		//gin.DisableConsoleColor()
		// 中间件，输出 log 到文件
		engine.Use(tools.LoggerToFile(common.Config.LogFilePath, common.Config.LogFileName))
		//禁止控制台输出
		gin.DefaultWriter = io.Discard
	}
	//自定义分隔符，避免与vue.js冲突
	engine.Delims("[[", "]]")
	//https://stackoverflow.com/questions/66248258/serve-embedded-filesystem-from-root-path-of-url
	assetsEmbedFS, err := fs.Sub(staticAssetFS, "static/assets")
	if err != nil {
		fmt.Println(err)
	}
	engine.StaticFS("/assets/", http.FS(assetsEmbedFS))
	imagesEmbedFS, errStaticImageFS := fs.Sub(staticImageFS, "static/images")
	if errStaticImageFS != nil {
		fmt.Println(errStaticImageFS)
	}
	engine.StaticFS("/images/", http.FS(imagesEmbedFS))
	//单独一张静态图片
	//singleStaticFiles(engine, "/favicon.ico", "static/images/favicon.ico", "image/x-icon")
	engine.GET("/favicon.ico", func(c *gin.Context) {
		file, _ := staticFS.ReadFile("static/images/favicon.ico")
		c.Data(
			http.StatusOK,
			"image/x-icon",
			file,
		)
	})
	//解析模板到HTML
	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "template-data", gin.H{
			"title": "Comigo 漫画阅读器 " + common.Version, //页面标题
		})
	})
}

// 简单的路由组: api,方便管理部分相同的URL
var api *gin.RouterGroup

// 2、设置获取书籍信息、图片文件的 API
func setWebAPI(engine *gin.Engine) {
	////TODO：处理登陆 https://www.chaindesk.cn/witbook/19/329
	////TODO：实现第三方认证，可参考 https://darjun.github.io/2021/07/26/godailylib/goth/

	////简单http认证
	//enableAuth := common.Config.UserName != "" && common.Config.Password != ""
	//if enableAuth {
	//	//api = engine.Group("/api", gin.BasicAuth(gin.Accounts{
	//	//	common.Config.UserName: common.Config.Password,
	//	//}))
	//}
	api = engine.Group("/api")

	// jwt登录、认证、鉴权etc
	// https://github.com/appleboy/gin-jwt
	// https://juejin.cn/post/7042520107976753165

	//使用gin+jwt实现身份验证功能： https://blog.firerain.me/article/18
	// gin-jwt 的使用简单，配置下中间件,然后在需要验证的api中用此中间件就行了

	//中间件配置。创建中间件的结构体 jwt.GinJWTMiddleware。
	ginJWTMiddleware := &jwt.GinJWTMiddleware{
		Realm:            "test zone",                                             //标识
		SigningAlgorithm: "HS256",                                                 //加密算法
		Key:              []byte(common.Config.UserName + common.Config.Password), //JWT服务端密钥，需要确保别人不知道
		Timeout:          time.Hour,                                               //jwt过期时间
		MaxRefresh:       time.Hour,                                               //刷新的时候，最大能延长多少时间
		IdentityKey:      "id",                                                    //指定cookie的id
		Authenticator:    token.Authenticator,                                     //认证器：在登录接口中使用的验证方法，返回验证成功后的用户对象。
		Authorizator:     token.Authorizator,                                      //授权者：登录后验证传入的 token 方法，可在此处写权限验证逻辑
		//验证失败后的函数调用，可用于自定义返回的 JSON 格式之类
		Unauthorized: func(c *gin.Context, code int, message string) {
			fmt.Println(code)
			fmt.Println(message)
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		//定义登录成功后用户名储存以及传递用户名到 Authorizator
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return claims["username"]
		},
		//添加额外业务相关的信息
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if user, ok := data.(token.User); ok {
				return jwt.MapClaims{"username": user.Username}
			}
			return jwt.MapClaims{}
		},
		// 指定从哪里获取token 其格式为："<source>:<name>" 如有多个，用逗号隔开，可用的值：
		//    "header:<name>"
		//    "query:<name>"
		//    "cookie:<name>"
		// 默认值 header:Authorization
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		//Header 中 token 的头部字段，默认值 Bearer
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	}
	// 创建 jwt 中间件
	jwtMiddleware, _ := jwt.New(ginJWTMiddleware)

	api = engine.Group("/api")

	// 登录 api ，直接用 jwtMiddleware 中的 `LoginHandler`
	//这个函数中，会执行上面设置的token.Authenticator来验证用户权限，如果通过就会返回token。
	api.POST("/login", jwtMiddleware.LoginHandler)
	//退出登录，会将用户的cookie中的token删除。
	api.POST("/logout", jwtMiddleware.LogoutHandler)
	// 刷新 token ，延长token的有效期
	api.GET("/refresh_token", jwtMiddleware.RefreshHandler)

	//简单的表单处理
	api.POST("/form", func(c *gin.Context) {
		t := c.DefaultPostForm("template", "scroll") //可设置默认值
		username := c.PostForm("username")
		password := c.PostForm("password")
		//bookList := c.PostFormMap("book_list")
		//bookList := c.QueryArray("book_list")
		//bookList := c.PostFormArray("book_list")
		if username == common.Config.UserName && password == common.Config.Password {
			c.String(http.StatusOK, fmt.Sprintf("template is %s, username is %s, password is %s,Config is %v", t, username, password, common.Config))
		} else {
			fmt.Println(fmt.Sprintf("template is %s, username is %s, password is %s,Config is %v", t, common.Config.UserName, common.Config.Password))
			c.String(http.StatusOK, " username or password is wrong.")
		}
	})

	// 在需要验证的api中用jwt中间件
	//通过URL字符串参数获取特定文件
	if common.Config.Debug {
		api.GET("/getfile", ginJWTMiddleware.MiddlewareFunc(), handler.GetFileHandler)
	} else {
		api.GET("/getfile", handler.GetFileHandler)
	}

	//文件上传
	api.POST("/upload", handler.UploadHandler)
	//web端需要的服务器状态，包括标题、机器状态等
	api.GET("/getstatus", handler.ServerStatusHandler)
	//获取书架信息，不包含每页信息
	api.GET("/getlist", handler.GetBookListHandler)
	//通过URL字符串参数查询书籍信息
	api.GET("/getbook", handler.GetBookHandler)

	////通过URL字符串参数PDF文件里的图片，效率太低，注释掉
	//api.GET("/get_pdf_image", handler.GetPdfImageHandler)
	//通过链接下载示例配置
	api.GET("/config.toml", handler.GetConfigHandler)
	//通过链接下载示例配置
	api.GET("/comigo.reg", handler.GetRegFIleHandler)
	//通过链接下载示例配置
	api.GET("/qrcode.png", handler.GetQrcodeHandler)
	//301重定向跳转示例
	api.GET("/redirect", func(c *gin.Context) {
		//支持内部和外部的重定向
		c.Redirect(http.StatusMovedPermanently, "https://www.google.com/")
	})
	//初始化websocket
	websocket.WsDebug = &common.Config.Debug
	api.GET("/ws", websocket.WsHandler)
	SetDownloadLink()
}

// 3、选择服务端口
func setPort() {
	//检测端口
	if !tools.CheckPort(common.Config.Port) {
		//获取一个空闲可用的系统端口号
		port, err := tools.GetFreePort()
		if err != nil {
			fmt.Println(err)
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			if common.Config.Port+2000 > 65535 {
				common.Config.Port = common.Config.Port + r.Intn(1024)
			} else {
				common.Config.Port = 30000 + r.Intn(20000)
			}
		} else {
			common.Config.Port = port
		}
		fmt.Println(locale.GetString("port_busy") + strconv.Itoa(common.Config.Port))
	}
}

// 5、setFrpClient
func setFrpClient() {
	//frp服务
	if !common.Config.EnableFrpcServer {
		return
	}
	if common.Config.FrpConfig.RemotePort <= 0 || common.Config.FrpConfig.RemotePort > 65535 {
		common.Config.FrpConfig.RemotePort = common.Config.Port
	}
	if common.Config.FrpConfig.RandomRemotePort {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		common.Config.FrpConfig.RemotePort = 50000 + r.Intn(10000)
	}
	frpcError := plugin.StartFrpC(common.Config.CachePath)
	if frpcError != nil {
		fmt.Println(locale.GetString("frpc_server_error"), frpcError.Error())
	} else {
		fmt.Println(locale.GetString("frpc_server_start"))
	}

}

// 6、printCMDMessage
func printCMDMessage() {
	//cmd打印链接二维码.如果只有一本书，就直接打开那本书.
	etcStr := ""
	//只有一本书的时候，URL需要附加的参数
	if book.GetBooksNumber() == 1 {
		bookList, err := book.GetAllBookInfoList("name")
		if err != nil {
			fmt.Println(err)
		}
		if len(bookList.BookInfos) == 1 {
			etcStr = "/#/scroll/" + bookList.BookInfos[0].BookID
		}
		if common.Config.DefaultMode != "" {
			etcStr = "/#/" + strings.ToLower(common.Config.DefaultMode) + "/" + bookList.BookInfos[0].BookID
		}
	}
	enableTls := common.Config.CertFile != "" && common.Config.KeyFile != ""
	tools.PrintAllReaderURL(common.Config.Port, common.Config.OpenBrowser, common.Config.EnableFrpcServer, common.Config.PrintAllIP, common.Config.Host, common.Config.FrpConfig.ServerAddr, common.Config.FrpConfig.RemotePort, common.Config.DisableLAN, enableTls, etcStr)
	//打印配置，调试用
	if common.Config.Debug {
		litter.Dump(common.Config)
	}
	fmt.Println(locale.GetString("ctrl_c_hint"))
}

// StartGinEngine 7、启动网页服务
func StartGinEngine(engine *gin.Engine) {
	//是否对外服务
	webHost := ":"
	if common.Config.DisableLAN {
		webHost = "localhost:"
	}
	//是否启用TLS
	enableTls := common.Config.CertFile != "" && common.Config.KeyFile != ""
	common.Srv = &http.Server{
		Addr:    webHost + strconv.Itoa(common.Config.Port),
		Handler: engine,
	}

	//在 goroutine 中初始化服务器，这样它就不会阻塞下面的正常关闭处理
	go func() {
		// 监听并启动服务(TLS)
		if enableTls {
			if err := common.Srv.ListenAndServeTLS(common.Config.CertFile, common.Config.KeyFile); err != nil && err != http.ErrServerClosed {
				time.Sleep(3 * time.Second)
				log.Fatalf("listen: %s\n", err)
			}
		}
		if !enableTls {
			// 监听并启动服务(HTTP)
			if err := common.Srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				time.Sleep(3 * time.Second)
				log.Fatalf("listen: %s\n", err)
			}
		}
	}()
}

////4、setWebpServer TODO：新的webp模式：https://docs.webp.sh/usage/remote-backend/
//func setWebpServer(engine *gin.Engine) {
//	//webp反向代理
//	if common.Config.EnableWebpServer {
//		webpError := common.StartWebPServer(common.CachePath+"/webp_config.json", common.ReadingBook.ExtractPath, common.CachePath+"/webp", common.Config.Port+1)
//		if webpError != nil {
//			fmt.Println(locale.GetString("webp_server_error"), webpError.Error())
//			//engine.Static("/cache", common.CachePath)
//
//		} else {
//			fmt.Println(locale.GetString("webp_server_start"))
//			engine.Use(reverse_proxy.ReverseProxyHandle("/cache", reverse_proxy.ReverseProxyOptions{
//				TargetHost:  "http://localhost",
//				TargetPort:  strconv.Itoa(common.Config.Port + 1),
//				RewritePath: "/cache",
//			}))
//		}
//	} else {
//		if common.ReadingBook.IsDir {
//			engine.Static("/cache/"+common.ReadingBook.BookID, common.ReadingBook.GetFilePath())
//		} else {
//			engine.Static("/cache", common.CachePath)
//		}
//	}
//}

//// 静态文件服务 单独设定某个文件
//func singleStaticFiles(engine *gin.Engine, fileUrl string, filePath string, contentType string) {
//	engine.GET(fileUrl, func(c *gin.Context) {
//		file, _ := staticFS.ReadFile(filePath)
//		c.Data(
//			http.StatusOK,
//			contentType,
//			file,
//		)
//	})
//}

//// getFileApi正常运作，虚拟文件系统实现方式
//func set-archiverFileSystem(engine *gin.Engine) {
////使用虚拟文件系统，设置服务路径（每本书都设置一遍）
////参考了: https://bitfieldconsulting.com/golang/filesystems
//for _, book := range common.BookList {
//	if book.NonUTF8Zip {
//		continue
//	}
//	ext := path.Ext(book.GetFilePath())
//	if (ext == ".zip" || ext == ".epub" || ext == ".cbz") && !book.NonUTF8Zip {
//		fsys, zipErr := zip.OpenReader(book.GetFilePath())
//		if zipErr != nil {
//			fmt.Println(zipErr)
//		}
//		httpFS := http.FS(fsys)
//		if book.IsDir {
//			engine.Static("/cache/"+book.BookID, book.GetFilePath())
//		} else {
//			engine.StaticFS("/cache/"+book.BookID, httpFS)
//		}
//	} else {
//		// 通过archiver/v4，建立虚拟FS。非UTF zip文件有编码问题
//		fsys, err := archiver.FileSystem(book.GetFilePath())
//		httpFS := http.FS(fsys)
//		if err != nil {
//			fmt.Println(err)
//		}
//		if book.IsDir {
//			engine.Static("/cache/"+book.BookID, book.GetFilePath())
//		} else {
//			engine.StaticFS("/cache/"+book.BookID, httpFS)
//		}
//	}
//}
//}
