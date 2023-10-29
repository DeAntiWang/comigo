package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yumenaka/comi/logger"
	"github.com/yumenaka/comi/types"
)

// HandlerGetBook 相关参数：
// id：书籍的ID，必须项目       							&id=2b17a130
// author：书籍的作者，未必存在									&author=佚名
// sort_page：按照自然文件名重新排序							&sort_page=true
// 示例 URL： http://127.0.0.1:1234/api/getbook?id=1215a&sort_by=name
// 示例 URL： http://127.0.0.1:1234/api/getbook?&author=Doe&name=book_name
func HandlerGetBook(c *gin.Context) {
	author := c.DefaultQuery("author", "")
	sortBy := c.DefaultQuery("sort_by", "default")
	id := c.DefaultQuery("id", "")
	if author != "" {
		bookList, err := types.GetBookByAuthor(author, sortBy)
		if err != nil {
			logger.Info(err)
		}
		c.PureJSON(http.StatusOK, bookList)
		return
	}
	if id != "" {
		b, err := types.GetBookByID(id, sortBy)
		if err != nil {
			logger.Info(err)
			c.PureJSON(http.StatusBadRequest, "id not found")
			return
		}
		c.PureJSON(http.StatusOK, b)
	}
}

// HandlerGetQuickJumpInfo 示例 URL： http://127.0.0.1:1234/api/getbook?id=1215a&sort_by=name
func HandlerGetQuickJumpInfo(c *gin.Context) {
	sortBy := c.DefaultQuery("sort_by", "default")
	id := c.DefaultQuery("id", "")
	if id == "" {
		c.PureJSON(http.StatusBadRequest, "id not set")
		return
	}
	//sortBy: 根据压缩包原始顺序、时间、文件名排序
	b, err := types.GetBookByID(id, sortBy)
	if err != nil {
		logger.Info(err)
		c.PureJSON(http.StatusBadRequest, "id not found")
		return
	}
	bookList, err := types.GetBookIDByParentFolder(b.ParentFolder, sortBy)
	if err != nil {
		logger.Info(err)
		c.PureJSON(http.StatusBadRequest, "ParentFolder, not found")
		return
	}
	idArray := make([]string, 0)
	for _, book := range bookList {
		idArray = append(idArray, book.BookID)
	}
	c.PureJSON(http.StatusOK, idArray)
}
