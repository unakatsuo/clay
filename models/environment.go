package models

type Environment struct {
	ID                    int    `json:"id" form:"id" gorm:"primary_key;AUTO_INCREMENT"`
	TemplateID            int    `json:"template_id" gorm:"index" sql:"type:integer references templates(id) on delete set null"`
	GitRepositoryURI      string `json:"git_repository_uri"`
	DesignFileName        string `json:"design_file_name"`
	TemplateFileName      string `json:"template_file_name"`
	TestCaseDirectoryName string `json:"test_case_directory_name"`
}

var EnvironmentModel = &Environment{}
