package pages

import (
	"net/http"

	"github.com/angelofallars/htmx-go"
	"github.com/gin-gonic/gin"
	"github.com/yumenaka/comigo/entity"
	"github.com/yumenaka/comigo/htmx/state"
	"github.com/yumenaka/comigo/htmx/templates/components"
	"github.com/yumenaka/comigo/util/logger"
)

// ScrollHandler 阅读界面（先做卷轴模式）
func ScrollHandler(c *gin.Context) {
	bookID := c.Param("id")
	if bookID != "" {
		state.Global.RequestBookID = bookID
	}
	book, err := entity.GetBookByID(bookID, "default")
	if err != nil {
		logger.Infof("GetBookByID: %v", err)
	}
	// TODO: 如果没有找到书籍，返回 HTTP 404 错误信息，或建议跳转上传页面。
	state.Global.TopBooks, err = entity.TopOfShelfInfo("name")
	if err != nil {
		logger.Infof("TopOfShelfInfo: %v", err)
	}

	// 定义模板主体内容。
	scrollPage := ScrollPage(&state.Global, book)
	// 为首页定义模板布局。
	indexTemplate := components.MainLayout(
		c,
		scrollPage, // define body content
	)

	// 渲染索引页模板。
	if err := htmx.NewResponse().RenderTempl(c.Request.Context(), c.Writer, indexTemplate); err != nil {
		// 如果不是，返回 HTTP 500 错误。
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}
