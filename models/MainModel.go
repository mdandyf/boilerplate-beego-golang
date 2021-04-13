package models

type MainDataModel struct {
	ApplicationName    string `db:"application_name" json:"application_name"`
	ApplicationVersion string `db:"version" json:"version"`
}
