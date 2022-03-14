package handler

import (
	"be-timesheet/pkg/entity"
	"be-timesheet/pkg/jwttoken"
	"be-timesheet/pkg/response"
	"be-timesheet/pkg/service"

	"fmt"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type addFileHandler struct {
	addFileService service.FileService
}

func NewAddFileHandler(addFileService service.FileService) *addFileHandler {
	return &addFileHandler{addFileService}
}

func (h *addFileHandler) GetAllFiles(c *gin.Context) {
	result, err := h.addFileService.GetAllFiles()
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	if result == nil {
		result = []entity.AdditionFileResponse{}
	}

	response.ResponseOKWithData(c, result)
}

func (h *addFileHandler) GetFileByID(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	result, err := h.addFileService.GetFileByID(id)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	response.ResponseOKWithData(c, result)
}

func (h *addFileHandler) AddNewFile(c *gin.Context) {
	idDocType, _ := strconv.Atoi(c.DefaultPostForm("id_doctype", "id_doctype"))

	tokenMetadata, err := jwttoken.ExtractTokenMetadata(c.Request)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	userId := tokenMetadata.UserID

	file, err := c.FormFile("file")
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	path := viper.GetString("Files.Path")
	filename := filepath.Base(file.Filename)
	if err := c.SaveUploadedFile(file, fmt.Sprintf("%s/%s", path, filename)); err != nil {
		response.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	var uploadFile = entity.AdditionFileResponse{
		IDDocType: idDocType,
		File:      filename,
		UserID:    int(userId),
	}

	result, err := h.addFileService.AddNewFile(&uploadFile)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
	}

	response.ResponseCreated(c, result)
	// var file entity.AdditionFileRequest

	// tokenMetadata, _ := jwttoken.ExtractTokenMetadata(c.Request)

	// err := c.ShouldBindJSON(&file)
	// if err != nil {
	// 	response.ResponseCustomError(c, err, http.StatusBadRequest)
	// 	return
	// } else {
	// 	userId := tokenMetadata.UserID
	// 	file.UserID = int(userId)

	// 	uploadFile, err := c.FormFile("file")
	// 	if err != nil {
	// 		response.ResponseError(c, err.Error(), http.StatusBadRequest)
	// 		return
	// 	}

	// 	path := viper.GetString("Files.Path")
	// 	filename := filepath.Base(uploadFile.Filename)
	// 	if err := c.SaveUploadedFile(uploadFile, fmt.Sprintf("%s/%s", path, filename)); err != nil {
	// 		response.ResponseError(c, err.Error(), http.StatusBadRequest)
	// 		return
	// 	}
	// 	file.File = filename

	// }

	// newFile, err := h.addFileService.AddNewFile(file)

	// if err != nil {
	// 	response.ResponseError(c, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// response.ResponseCreated(c, newFile)
}

func (h *addFileHandler) UpdateFile(c *gin.Context) {
	fileId, err := strconv.Atoi(c.Param("id_file"))
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	idDocType, _ := strconv.Atoi(c.DefaultPostForm("id_doctype", "id_doctype"))

	tokenMetadata, err := jwttoken.ExtractTokenMetadata(c.Request)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	userId := tokenMetadata.UserID

	file, err := c.FormFile("file")
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	filename := filepath.Base(file.Filename)
	if err := c.SaveUploadedFile(file, filename); err != nil {
		response.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	var updatedFile = entity.AdditionFileResponse{
		ID:        fileId,
		IDDocType: idDocType,
		File:      filename,
		UserID:    int(userId),
	}

	result, err := h.addFileService.UpdateFile(&updatedFile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, result)

	// var file entity.AdditionFileRequest

	// tokenMetadata, _ := jwttoken.ExtractTokenMetadata(c.Request)

	// err := c.ShouldBindJSON(&file)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error": err,
	// 	})
	// 	return
	// } else {
	// 	userId := tokenMetadata.UserID
	// 	file.UserID = int(userId)

	// 	updatedFile, err := c.FormFile("file")
	// 	if err != nil {
	// 		response.ResponseError(c, err.Error(), http.StatusBadRequest)
	// 		return
	// 	}

	// 	filename := filepath.Base(updatedFile.Filename)
	// 	if err := c.SaveUploadedFile(updatedFile, filename); err != nil {
	// 		response.ResponseError(c, err.Error(), http.StatusBadRequest)
	// 		return
	// 	}

	// 	file.File = filename
	// }

	// idStr := c.Param("id_file")
	// id, _ := strconv.Atoi(idStr)

	// result, err := h.addFileService.UpdateFile(id, file)

	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error": err,
	// 	})
	// 	return
	// }

	// c.JSON(http.StatusOK, gin.H{
	// 	"data": result,
	// })

}

func (h *addFileHandler) DeleteFile(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.addFileService.DeleteFile(id)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	response.ResponseOK(c, "Succesfully Deleted Timesheet")

}

// func (h *addFileHandler) DownloadFile(c *gin.Context) {
// 	var f entity.AdditionFileDownload
// 	if err := c.ShouldBindUri(&f); err != nil {
// 		response.ResponseError(c, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	m, cn, err := h.DownloadFile(f.Name)
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": err})
// 		return
// 	}
// 	c.Header("Content-Disposition", "attachment; filename"+f.Name)
// }
