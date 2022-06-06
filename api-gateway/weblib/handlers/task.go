package handlers

import (
	"api-gateway/pkg/utils"
	"api-gateway/services"
	"context"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetTaskList(ginCtx *gin.Context) {
	var taskReq services.TaskRequest
	PanicIfTaskError(ginCtx.Bind(&taskReq))
	taskService := ginCtx.Keys["taskService"].(services.TaskService)
	//通过拿到当前访问的用户的id，拿到用户的备忘录
	claims, _ := utils.ParseToken(ginCtx.GetHeader("Authorization"))
	taskReq.Uid = uint64(claims.Id)
	//调用服务端的函数
	taskResp, err := taskService.GetTasksList(context.Background(), &taskReq)
	if err != nil {
		PanicIfTaskError(err)
	}
	ginCtx.JSON(200, gin.H{
		"task":  taskResp.TaskList,
		"count": taskResp.Count,
	})
}

func CreateTaskList(ginCtx *gin.Context) {
	var taskReq services.TaskRequest
	PanicIfTaskError(ginCtx.Bind(&taskReq))

	//通过拿到当前访问的用户的id，拿到用户的备忘录
	claims, _ := utils.ParseToken(ginCtx.GetHeader("Authorization"))
	taskReq.Uid = uint64(claims.Id)
	taskService := ginCtx.Keys["taskService"].(services.TaskService)
	//调用服务端的函数
	taskResp, err := taskService.CreateTask(context.Background(), &taskReq)
	if err != nil {
		PanicIfTaskError(err)
	}
	ginCtx.JSON(200, gin.H{
		"taskRes": taskResp.TaskDetail,
	})

}

func GetTaskDetail(ginCtx *gin.Context) {
	var taskReq services.TaskRequest
	PanicIfTaskError(ginCtx.Bind(&taskReq))
	productService := ginCtx.Keys["taskService"].(services.TaskService)
	//通过拿到当前访问的用户的id，拿到用户的备忘录
	claims, _ := utils.ParseToken(ginCtx.GetHeader("Authorization"))
	taskReq.Uid = uint64(claims.Id)
	//获取task_id
	id, _ := strconv.Atoi(ginCtx.Param("id"))
	taskReq.Id = uint64(id)
	//调用服务端的函数
	productRes, err := productService.GetTask(context.Background(), &taskReq)
	if err != nil {
		PanicIfTaskError(err)
	}
	ginCtx.JSON(200, gin.H{
		"productRes": productRes.TaskDetail,
	})
}

func UpdateTask(ginCtx *gin.Context) {
	var taskReq services.TaskRequest
	PanicIfTaskError(ginCtx.Bind(&taskReq))
	productService := ginCtx.Keys["taskService"].(services.TaskService)
	//通过拿到当前访问的用户的id，拿到用户的备忘录
	claims, _ := utils.ParseToken(ginCtx.GetHeader("Authorization"))
	taskReq.Uid = uint64(claims.Id)
	//获取task_id
	id, _ := strconv.Atoi(ginCtx.Param("id"))
	taskReq.Id = uint64(id)
	//调用服务端的函数
	productRes, err := productService.UpdateTask(context.Background(), &taskReq)
	if err != nil {
		PanicIfTaskError(err)
	}
	ginCtx.JSON(200, gin.H{
		"productRes": productRes.TaskDetail,
	})
}

func DeleteTask(ginCtx *gin.Context) {
	var taskReq services.TaskRequest
	PanicIfTaskError(ginCtx.Bind(&taskReq))
	productService := ginCtx.Keys["taskService"].(services.TaskService)
	//通过拿到当前访问的用户的id，拿到用户的备忘录
	claims, _ := utils.ParseToken(ginCtx.GetHeader("Authorization"))
	taskReq.Uid = uint64(claims.Id)
	//获取task_id
	id, _ := strconv.Atoi(ginCtx.Param("id"))
	taskReq.Id = uint64(id)
	//调用服务端的函数
	productRes, err := productService.DeleteTask(context.Background(), &taskReq)
	if err != nil {
		PanicIfTaskError(err)
	}
	ginCtx.JSON(200, gin.H{
		"productRes": productRes.TaskDetail,
	})
}
