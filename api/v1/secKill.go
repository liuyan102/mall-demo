package v1

import (
	"fmt"
	"mall-demo/internal/dto"
	"mall-demo/internal/service"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

// AddSecKillProduct 添加秒杀商品
func AddSecKillProduct(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	files := form.File["file"]
	var secKillProductRequest dto.SecKillProductRequest
	err2 := ctx.ShouldBind(&secKillProductRequest)
	if err2 != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	var secKillService service.SecKillService
	response := secKillService.AddSecKillProduct(ctx, secKillProductRequest, files)
	ctx.JSON(http.StatusOK, response)
}

// SecKillWithoutLock 无锁秒杀商品
func SecKillWithoutLock(ctx *gin.Context) {
	var secKillRequest dto.SecKill
	err := ctx.ShouldBind(&secKillRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	var secKillService service.SecKillService
	secKillNum := 50
	var wg sync.WaitGroup
	wg.Add(secKillNum)
	for i := 0; i < secKillNum; i++ {
		go func(num int) {
			response := secKillService.SecKillWithoutLock(ctx, secKillRequest)
			fmt.Println(num, response)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("secKill Done!")
	ctx.JSON(http.StatusOK, "secKill Done!")
}

// SecKillWithMutexLock 互斥锁秒杀商品
func SecKillWithMutexLock(ctx *gin.Context) {
	var secKillRequest dto.SecKill
	err := ctx.ShouldBind(&secKillRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	var secKillService service.SecKillService
	secKillNum := 50
	var wg sync.WaitGroup
	wg.Add(secKillNum)
	var lock sync.Mutex
	for i := 0; i < secKillNum; i++ {
		go func(num int) {
			lock.Lock()
			response := secKillService.SecKillWithoutLock(ctx, secKillRequest)
			lock.Unlock()
			fmt.Println(num, response)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("secKill Done!")
	ctx.JSON(http.StatusOK, "secKill Done!")
}

// SecKillWithXLock 排他锁秒杀商品
func SecKillWithXLock(ctx *gin.Context) {
	var secKillRequest dto.SecKill
	err := ctx.ShouldBind(&secKillRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	var secKillService service.SecKillService
	secKillNum := 100000
	var wg sync.WaitGroup
	wg.Add(secKillNum)
	for i := 0; i < secKillNum; i++ {
		go func(num int) {
			response := secKillService.SecKillWithXLock(ctx, secKillRequest)
			fmt.Println(num, response)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("secKill Done!")
	ctx.JSON(http.StatusOK, "secKill Done!")
}

// SecKillWithRedis redis秒杀商品
func SecKillWithRedis(ctx *gin.Context) {
	var secKillRequest dto.SecKill
	err := ctx.ShouldBind(&secKillRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	var secKillService service.SecKillService
	secKillNum := 50
	var wg sync.WaitGroup
	wg.Add(secKillNum)
	for i := 0; i < secKillNum; i++ {
		go func(num int) {
			response := secKillService.SecKillWithRedis(ctx, secKillRequest, num)
			fmt.Println(num, response)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("secKill Done!")
	ctx.JSON(http.StatusOK, "secKill Done!")
}
