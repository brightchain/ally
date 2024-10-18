package models

import "time"

type CarBrand struct {
	Id        int32     `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Name      string    `gorm:"column:name;default:;NOT NULL;comment:'名称'"`
	Brand     string    `gorm:"column:brand;default:;NOT NULL;comment:'品牌名称'"`
	FullName  string    `gorm:"column:full_name;default:;NOT NULL;comment:'全称'"`
	Initial   string    `gorm:"column:initial;default:A;NOT NULL;comment:'首字母'"`
	Logo      string    `gorm:"column:logo;default:;NOT NULL;comment:'logo'"`
	Salestate string    `gorm:"column:salestate;default:;NOT NULL;comment:'销售状态'"`
	Depth     int8      `gorm:"column:depth;default:1;NOT NULL;comment:'层级'"`
	CreatedAt time.Time `gorm:"column:created_at;comment:'创建时间'"`
	UpdatedAt time.Time `gorm:"column:updated_at;comment:'更新时间'"`
}

func (c *CarBrand) TableName() string {
	return "car_brand"
}
