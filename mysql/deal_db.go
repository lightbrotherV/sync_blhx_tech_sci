package mysql

import (
	"database/sql"
	"errors"
)

type AzurLane struct {
	ID                      int64  `json:"id" db:"id"`
	CreatedById             int64  `json:"createdById,omitempty" db:"created_by_id"`
	UpdatedById             int64  `json:"updatedById,omitempty" db:"updated_by_id"`
	Code                    string `json:"code" db:"code"`
	Name                    string `json:"name" db:"name"`
	Camp                    string `json:"camp" db:"camp"`
	ShipType                string `json:"ship_type" db:"ship_type"`
	TechPointGet            int64  `json:"tech_point_get" db:"tech_point_get"`
	TechPointStar           int64  `json:"tech_point_star" db:"tech_point_star"`
	TechPointLv120          int64  `json:"tech_point_lv120" db:"tech_point_lv120"`
	TechPointTotal          int64  `json:"tech_point_total,omitempty" db:"tech_point_total"`
	AttributeGetApplyShip   string `json:"attribute_get_apply_ship" db:"attribute_get_apply_ship"`
	AttributeLv120ApplyShip string `json:"attribute_lv120_apply_ship" db:"attribute_lv120_apply_ship"`
	AttributeNameGet        string `json:"attribute_name_get" db:"attribute_name_get"`
	AttributeNameLv120      string `json:"attribute_name_lv120" db:"attribute_name_lv120"`
	AttributeGet            int64  `json:"attribute_get" db:"attribute_get"`
	AttributeLv120          int64  `json:"attribute_lv120" db:"attribute_lv120"`
	IsGetTech               string `json:"is_get_tech" db:"is_get_tech"`
	CreatedAt               string `json:"createdAt" db:"created_at"`
	UpdatedAt               string `json:"updatedAt" db:"updated_at"`
}

const (
	GetRecordByCodeSql = "SELECT id, updatedAt, createdAt, createdById, updatedById, code, name, camp, ship_type, tech_point_get, tech_point_star, tech_point_lv120, tech_point_total, attribute_get_apply_ship, attribute_lv120_apply_ship, attribute_name_get, attribute_name_lv120, attribute_get, attribute_lv120, is_get_tech FROM `azur_lane` WHERE code = ?"

	UpdateRecordByCodeSql = "UPDATE azur_lane SET updatedAt=?, tech_point_get=?, tech_point_star=?, tech_point_lv120=?, tech_point_total=?, attribute_get_apply_ship=?, attribute_lv120_apply_ship=?, attribute_name_get=?, attribute_name_lv120=?, attribute_get=?, attribute_lv120=?, is_get_tech=? WHERE code = ?"

	InsertRecordByCodeSql = "INSERT INTO azur_lane(updatedAt, createdAt, createdById, updatedById, code, name, camp, ship_type, tech_point_get, tech_point_star, tech_point_lv120, tech_point_total, attribute_get_apply_ship, attribute_lv120_apply_ship, attribute_name_get, attribute_name_lv120, attribute_get, attribute_lv120, is_get_tech) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
)

func GetRecordByCode(code string) (*AzurLane, error) {
	res := &AzurLane{}
	row := MysqlDb.QueryRow(GetRecordByCodeSql, code)
	err := row.Scan(&res.ID, &res.UpdatedAt, &res.CreatedAt, &res.CreatedById, &res.UpdatedById, &res.Code, &res.Name, &res.Camp, &res.ShipType, &res.TechPointGet, &res.TechPointStar, &res.TechPointLv120, &res.TechPointTotal, &res.AttributeGetApplyShip, &res.AttributeLv120ApplyShip, &res.AttributeNameGet, &res.AttributeNameLv120, &res.AttributeGet, &res.AttributeLv120, &res.IsGetTech)
	if errors.Is(err, sql.ErrNoRows) {
		return res, nil
	}
	return res, err
}

func UpdateRecordByCode(record *AzurLane) error {
	_, err := MysqlDb.Exec(UpdateRecordByCodeSql, record.UpdatedAt, record.TechPointGet, record.TechPointStar, record.TechPointLv120, record.TechPointTotal, record.AttributeGetApplyShip, record.AttributeLv120ApplyShip, record.AttributeNameGet, record.AttributeNameLv120, record.AttributeGet, record.AttributeLv120, record.IsGetTech, record.Code)
	return err
}

func InsertRecordByCode(record *AzurLane) error {
	_, err := MysqlDb.Exec(InsertRecordByCodeSql, record.UpdatedAt, record.CreatedAt, record.CreatedById, record.UpdatedById, record.Code, record.Name, record.Camp, record.ShipType, record.TechPointGet, record.TechPointStar, record.TechPointLv120, record.TechPointTotal, record.AttributeGetApplyShip, record.AttributeLv120ApplyShip, record.AttributeNameGet, record.AttributeNameLv120, record.AttributeGet, record.AttributeLv120, record.IsGetTech)
	return err
}
