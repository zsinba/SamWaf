package api

import (
	"SamWaf/model/common/response"
	"SamWaf/model/request"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type WafWhiteIpApi struct {
}

func (w *WafWhiteIpApi) AddApi(c *gin.Context) {
	var req request.WafWhiteIpAddReq
	err := c.ShouldBind(&req)
	if err == nil {
		err = wafIpWhiteService.CheckIsExistApi(req)
		if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
			err = wafIpWhiteService.AddApi(req)
			if err == nil {

				response.OkWithMessage("添加成功", c)
			} else {

				response.FailWithMessage("添加失败", c)
			}
			return
		} else {
			response.FailWithMessage("当前网站的IP已经存在", c)
			return
		}

	} else {
		response.FailWithMessage("解析失败", c)
	}
}
func (w *WafWhiteIpApi) GetDetailApi(c *gin.Context) {
	var req request.WafWhiteIpDetailReq
	err := c.ShouldBind(&req)
	if err == nil {
		wafIpWhite := wafIpWhiteService.GetDetailApi(req)
		response.OkWithDetailed(wafIpWhite, "获取成功", c)
	} else {
		response.FailWithMessage("解析失败", c)
	}
}
func (w *WafWhiteIpApi) GetListApi(c *gin.Context) {
	var req request.WafWhiteIpSearchReq
	err := c.ShouldBind(&req)
	if err == nil {
		wafIpWhites, total, _ := wafIpWhiteService.GetListApi(req)
		response.OkWithDetailed(response.PageResult{
			List:      wafIpWhites,
			Total:     total,
			PageIndex: req.PageIndex,
			PageSize:  req.PageSize,
		}, "获取成功", c)
	} else {
		response.FailWithMessage("解析失败", c)
	}
}
func (w *WafWhiteIpApi) DelWhiteIpApi(c *gin.Context) {
	var req request.WafWhiteIpDelReq
	err := c.ShouldBind(&req)
	if err == nil {
		err = wafIpWhiteService.DelApi(req)
		if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
			response.FailWithMessage("请检测参数", c)
		} else if err != nil {
			response.FailWithMessage("发生错误", c)
		} else {
			response.FailWithMessage("删除成功", c)
		}

	} else {
		response.FailWithMessage("解析失败", c)
	}
}

func (w *WafWhiteIpApi) ModifyWhiteIpApi(c *gin.Context) {
	var req request.WafWhiteIpEditReq
	err := c.ShouldBind(&req)
	if err == nil {
		err = wafIpWhiteService.ModifyApi(req)
		if err != nil {
			response.FailWithMessage("编辑发生错误", c)
		} else {
			response.OkWithMessage("编辑成功", c)
		}

	} else {
		response.FailWithMessage("解析失败", c)
	}
}
