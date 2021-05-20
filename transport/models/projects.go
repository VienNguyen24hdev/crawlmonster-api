package models

import "github.com/jasonbronson/crawlmonster-api/config"

func GetProjects(customerID string) (*Projects, error) {

	projects := Projects{}
	result := config.Cfg.GormDB.Where("customer_id = ?", customerID).First(&projects)
	if result.Error != nil {
		return &projects, result.Error
	}
	return &projects, nil

}

type Projects struct {
	Project []Project
}
type Project struct {
	ID string
}
