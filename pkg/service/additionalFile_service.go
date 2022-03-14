package service

import (
	"be-timesheet/pkg/entity"
	"be-timesheet/pkg/repository"
)

type FileService interface {
	GetAllFiles() ([]entity.AdditionFileResponse, error)
	GetFileByID(ID int) (entity.AdditionalFile, error)
	AddNewFile(file *entity.AdditionFileResponse) (*entity.AdditionFileResponse, error)
	UpdateFile(fileResponse *entity.AdditionFileResponse) (*entity.AdditionFileResponse, error)
	DeleteFile(ID int) error
	// DownloadFile(FileName string) (string, []byte, error)
}

type fileService struct {
	fileRepo repository.FileRepository
}

func NewServiceAdditionalFile(fileRepo repository.FileRepository) *fileService {
	return &fileService{fileRepo}
}

func (s *fileService) GetAllFiles() ([]entity.AdditionFileResponse, error) {
	result, err := s.fileRepo.GetAllFiles()
	if err != nil {
		return nil, err
	}

	var fileList []entity.AdditionFileResponse

	if result != nil {
		for _, item := range result {
			file := entity.AdditionFileResponse{
				ID:        item.ID,
				IDDocType: item.IDDocType,
				File:      item.File,
				UserID:    item.UserID,
			}
			fileList = append(fileList, file)
		}
	}
	return fileList, nil
}

func (s *fileService) GetFileByID(ID int) (entity.AdditionalFile, error) {
	file, err := s.fileRepo.GetFileByID(ID)
	return file, err
}

func (s *fileService) AddNewFile(file *entity.AdditionFileResponse) (*entity.AdditionFileResponse, error) {
	var newFile = entity.AdditionalFile{
		IDDocType: file.IDDocType,
		File:      file.File,
		UserID:    file.UserID,
	}

	result, err := s.fileRepo.AddNewFile(&newFile)
	if err != nil {
		return nil, err
	}

	if result != nil {
		file = &entity.AdditionFileResponse{
			ID:        result.ID,
			IDDocType: result.IDDocType,
			File:      result.File,
			UserID:    result.UserID,
		}
	}

	return file, nil
}

func (s *fileService) UpdateFile(fileResponse *entity.AdditionFileResponse) (*entity.AdditionFileResponse, error) {
	var file = entity.AdditionalFile{
		ID:        fileResponse.ID,
		IDDocType: fileResponse.IDDocType,
		File:      fileResponse.File,
		UserID:    fileResponse.UserID,
	}

	_, err := s.fileRepo.UpdateFile(&file)

	if err != nil {
		return nil, err
	}

	return fileResponse, nil
}

func (s *fileService) DeleteFile(ID int) error {
	err := s.fileRepo.DeleteFile(ID)
	if err != nil {
		return err
	}

	return nil
}

// var path = viper.GetString("Files.Path")

// func (s *fileService) DownloadFile(fileName string) (string, []byte, error) {
// 	dst := fmt.Sprintf("%s %s", path, fileName)
// 	b, err := ioutil.ReadFile(dst)
// 	if err != nil {
// 		return "", nil, err
// 	}
// 	m := http.DetectContentType(b[:512])

// 	return m, b, nil
// }
