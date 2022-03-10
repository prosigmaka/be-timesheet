package service

import (
	"be-timesheet/pkg/entity"
	"be-timesheet/pkg/repository"
)

type DataService interface {
	GetAllData() ([]entity.DataResponse, error)
	GetDataByID(ID int) (entity.Data, error)
	AddData(dataRequest entity.DataRequest) (entity.Data, error)
	UpdateData(ID int, dataRequest entity.DataRequest) (entity.Data, error)
	DeleteData(ID int) error
}

type dataService struct {
	dataRepo repository.DataRespository
}

func NewServiceData(dataRepo repository.DataRespository) *dataService {
	return &dataService{dataRepo}
}

func (s *dataService) GetAllData() ([]entity.DataResponse, error) {
	result, err := s.dataRepo.GetAllData()
	if err != nil {
		return nil, err
	}

	var dataList []entity.DataResponse

	if result != nil {
		for _, item := range result{
			data := entity.DataResponse{
				ID: 			item.ID,
				UserID: 		item.UserID,
				ProjectID: 		item.ProjectID,
				TimesheetID: 	item.TimesheetID,
				Employee_name: 	item.Employee_name,
				Project_name: 	item.Project_name,
				StartDate: 		item.StartDate,
				EndDate: 		item.EndDate,	
			}
			dataList = append(dataList, data)
		}
	}
	return dataList, nil
}

func (s *dataService) GetDataByID(ID int) (entity.Data, error) {
	data, err := s.dataRepo.GetDataByID(ID)
	return data, err
}


func (s *dataService) AddData(dataRequest entity.DataRequest) (entity.Data, error) {
	data := entity.Data{
		UserID: dataRequest.UserID,
		TimesheetID: dataRequest.TimesheetID,
		ProjectID: dataRequest.ProjectID,
		Employee_name: dataRequest.Employee_name,
		Project_name: dataRequest.Project_name,
		StartDate: dataRequest.StartDate,
		EndDate: dataRequest.EndDate,
	}
	newData, err := s.dataRepo.AddData(data)
	return newData, err
}

func (s *dataService) UpdateData(ID int, dataRequest entity.DataRequest) (entity.Data, error) {
	data, err := s.dataRepo.GetDataByID(ID)

	data.UserID = dataRequest.UserID
	data.ProjectID = dataRequest.ProjectID
	data.TimesheetID = dataRequest.TimesheetID
	data.Employee_name = dataRequest.Employee_name
	data.Project_name = dataRequest.Project_name
	data.StartDate = dataRequest.StartDate
	data.EndDate = dataRequest.EndDate

	updatedData, err := s.dataRepo.UpdateData(data)
	return updatedData, err

}

func (s *dataService) DeleteData(ID int) error {
	err := s.dataRepo.DeleteData(ID)
	if err != nil {
		return err
	}

	return nil
}