package service

import (
	"mall-demo/internal/dao"
	"mall-demo/internal/dto"
	"mall-demo/internal/model"
	"mall-demo/internal/pkg/e"
	"mall-demo/internal/pkg/res"
	"mall-demo/internal/pkg/util"
	"mall-demo/internal/vo"
	"mime/multipart"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

type ProductService struct {
}

// CreateProduct 创建商品
func (*ProductService) CreateProduct(ctx *gin.Context, request dto.CreateProductRequest, files []*multipart.FileHeader) res.Response {
	user, exists := ctx.Get("user")
	if !exists {
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "登录过期",
		}
	}
	boss := user.(*model.User)
	// 以第一张图片作为封面图
	firstFile, err := files[0].Open()
	if err != nil {
		util.Loggers.Errorln(err)
		return res.Response{
			Code: e.InvalidParam,
			Data: nil,
			Msg:  "文件打开失败",
		}
	}
	productPath, err := util.UploadProduct(firstFile, boss.ID, request.Name)
	if err != nil {
		util.Loggers.Errorln(err)
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "商品图片上传失败",
		}
	}
	product := &model.Product{
		Name:          request.Name,
		CategoryID:    request.CategoryID,
		Title:         request.Title,
		Info:          request.Info,
		ImgPath:       productPath,
		Price:         request.Price,
		DiscountPrice: request.DiscountPrice,
		OnSale:        true,
		Num:           request.Num,
		BossID:        boss.ID,
		BossName:      boss.UserName,
		BossAvatar:    boss.Avatar,
	}
	var productDao dao.ProductDao
	err = productDao.CreateProduct(product)
	if err != nil {
		util.Loggers.Errorln(err)
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "商品创建失败",
		}
	}

	var productImgDao dao.ProductImgDao
	// 并发执行
	// 创建计数器
	var wg sync.WaitGroup
	// 添加任务数
	wg.Add(len(files))
	// 创建channel接收返回值
	resChan := make(chan error, 10)
	// 循环创建协程执行并发
	for index, file := range files {
		// 创建协程执行匿名函数
		go func(goIndex int, goFile *multipart.FileHeader) { //todo:创建协程数可以加以限制
			open, err2 := goFile.Open()
			if err2 != nil {
				resChan <- err2
				wg.Done()
				return
			}
			path, err3 := util.UploadProduct(open, boss.ID, product.Name+strconv.Itoa(goIndex))
			if err3 != nil {
				resChan <- err3
				wg.Done()
				return
			}
			productImg := &model.ProductImg{
				ProductID: product.ID,
				ImgPath:   path,
			}
			err4 := productImgDao.CreateProductImg(productImg)
			if err4 != nil {
				resChan <- err4
				wg.Done()
				return
			}
			// 协程执行完毕，计数器-1
			wg.Done()
		}(index, file)
	}
	// 等待所有的协程执行完毕
	wg.Wait()
	// 关闭channel
	close(resChan)
	// 接收channel里传回的信息
	for err = range resChan {
		util.Loggers.Errorln(err)
		if err != nil {
			return res.Response{
				Code: e.Error,
				Data: nil,
				Msg:  "商品图片上传失败",
			}
		}
	}
	util.Loggers.Infoln("create product success")
	return res.Response{
		Code: e.Success,
		Data: vo.BuildProductResponse(product),
		Msg:  "商品创建成功",
	}

}

// ListProduct 展示商品
func (*ProductService) ListProduct(ctx *gin.Context, request dto.ListProductRequest) res.Response {
	if request.PageSize == 0 {
		request.PageSize = 15
	}
	condition := make(map[string]interface{})
	if request.CategoryID != 0 {
		condition["category_id"] = request.CategoryID
	}
	var productDao dao.ProductDao
	total, err := productDao.CountProductByCondition(condition)
	if err != nil {
		util.Loggers.Errorln(err)
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "查询失败",
		}
	}

	products, err := productDao.ListProductByCondition(condition, request.BasePage)
	if err != nil {
		util.Loggers.Errorln(err)
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "查询失败",
		}
	}
	util.Loggers.Infoln("list product success")
	return res.Response{
		Code: e.Success,
		Data: res.BuildDataList(vo.BuildProductResponseList(products), total),
		Msg:  "查询成功",
	}

}

// SearchProduct 搜索商品
func (*ProductService) SearchProduct(ctx *gin.Context, request dto.SearchProductRequest) res.Response {
	if request.PageSize == 0 {
		request.PageSize = 15
	}
	var productDao dao.ProductDao
	products, err := productDao.SearchProduct(request.Info, request.BasePage)
	if err != nil {
		util.Loggers.Errorln(err)
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "查询失败",
		}
	}
	util.Loggers.Infoln("search product success")
	return res.Response{
		Code: e.Success,
		Data: res.BuildDataList(vo.BuildProductResponseList(products), int64(len(products))),
		Msg:  "查询成功",
	}
}

// ShowProduct 查看商品详细信息
func (*ProductService) ShowProduct(ctx *gin.Context) res.Response {
	id := ctx.Param("id")
	productID, _ := strconv.Atoi(id)
	var productDao dao.ProductDao
	product, err := productDao.GetProductByID(uint(productID))
	if err != nil {
		util.Loggers.Errorln(err)
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "查询失败",
		}
	}
	product.AddView(ctx.Request.Context()) // 增加点击量
	util.Loggers.Infoln("show product success")
	return res.Response{
		Code: e.Success,
		Data: vo.BuildProductResponse(product),
		Msg:  "查询成功",
	}
}

// ShowProductImg 查看商品图片
func (*ProductService) ShowProductImg(ctx *gin.Context) res.Response {
	id := ctx.Param("id")
	productID, _ := strconv.Atoi(id)
	var productImgDao dao.ProductImgDao
	productImgList, err := productImgDao.GetProductImgByID(uint(productID))
	if err != nil {
		util.Loggers.Errorln(err)
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "查询失败",
		}
	}
	util.Loggers.Infoln("show product Img success")
	return res.Response{
		Code: e.Success,
		Data: res.BuildDataList(vo.BuildProductImgListResponse(productImgList), int64(len(productImgList))),
		Msg:  "查询成功",
	}

}
