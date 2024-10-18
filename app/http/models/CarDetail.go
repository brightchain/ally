package models

import "time"

type CarDetail struct {
	Id        int32     `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Name string `gorm:"column:name;default:;NOT NULL;comment:'名称'"`
    Brand string `gorm:"column:brand;default:;NOT NULL;comment:'品牌名称'"`
    Pid int32 `gorm:"column:pid;default:0;NOT NULL;comment:'上级id1'"`
    Fullname string `gorm:"column:fullname;default:;NOT NULL;comment:'全称'"`
    Initial string `gorm:"column:initial;default:A;NOT NULL;comment:'首字母'"`
    Logo string `gorm:"column:logo;default:;NOT NULL;comment:'logo'"`
    Price string `gorm:"column:price;default:;NOT NULL;comment:'售价'"`
    Yeartype string `gorm:"column:yeartype;default:;NOT NULL;comment:'年款'"`
    Salestate string `gorm:"column:salestate;default:;NOT NULL;comment:'销售状态'"`
    Sizetype string `gorm:"column:sizetype;default:;NOT NULL;comment:'尺寸类型'"`
    Productionstate string `gorm:"column:productionstate;default:;NOT NULL;comment:'生产状态'"`
    MachineoilVolume string `gorm:"column:machineoil_volume;default:;NOT NULL;comment:'机油参考用量'"`
    MachineoilViscosity string `gorm:"column:machineoil_viscosity;default:;NOT NULL;comment:'机油粘稠度'"`
    MachineoilGrade string `gorm:"column:machineoil_grade;default:;NOT NULL;comment:'机油分类'"`
    MachineoilLevel string `gorm:"column:machineoil_level;default:;NOT NULL;comment:'机油质量等级'"`
    EnginePosition string `gorm:"column:engine_position;default:;NOT NULL;comment:'发动机位置'"`
    EngineModel string `gorm:"column:engine_model;default:;NOT NULL;comment:'发动机型号'"`
    EngineDisplacement string `gorm:"column:engine_displacement;default:;NOT NULL;comment:'发动机排量'"`
    EngineIntakeform string `gorm:"column:engine_intakeform;default:;NOT NULL;comment:'发动机进气形式'"`
    EngineCylindernum string `gorm:"column:engine_cylindernum;default:;NOT NULL;comment:'发动机气缸数(个)'"`
    EngineMaxpowerspeed string `gorm:"column:engine_maxpowerspeed;default:;NOT NULL;comment:'发动机最大功率转速'"`
    Gearbox string `gorm:"column:gearbox;default:;NOT NULL;comment:'变速箱型号'"`
	CreatedAt time.Time `gorm:"column:created_at;comment:'创建时间'"`
	UpdatedAt time.Time `gorm:"column:updated_at;comment:'更新时间'"`
}

func (c *CarDetail) TableName() string {
	return "car_detail"
}
