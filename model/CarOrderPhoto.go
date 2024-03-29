package model

type CarOrderPhoto struct {
	Id           int32  `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	OrderNo      string `gorm:"column:order_no;NOT NULL;comment:'订单编号'"`
	Uid          int32  `gorm:"column:uid;NOT NULL;comment:'用户编号'"`
	MId          int32  `gorm:"column:m_id;NOT NULL;comment:'主订单id'"`
	Orderid      string `gorm:"column:orderid;NOT NULL;comment:'主订单编号'"`
	CouponId     int32  `gorm:"column:coupon_id;default:0;NOT NULL;comment:'优惠券ID'"`
	BatchNum     string `gorm:"column:batch_num;default:;NOT NULL;comment:'优惠券批次号'"`
	CouponAmount string `gorm:"column:coupon_amount;default:0.00;NOT NULL;comment:'优惠券金额'"`
	ProId        string `gorm:"column:pro_id;NOT NULL;comment:'产品id'"`
	ProName      string `gorm:"column:pro_name;default:;NOT NULL;comment:'产品名称'"`
	ProDetail    string `gorm:"column:pro_detail;default:;NOT NULL;comment:'产品规格介绍'"`
	ProImg       string `gorm:"column:pro_img;default:;NOT NULL;comment:'产品图片'"`
	OrderAmount  string `gorm:"column:order_amount;default:0.00;NOT NULL;comment:'订单金额'"`
	PayAmount    string `gorm:"column:pay_amount;default:0.00;NOT NULL;comment:'支付金额'"`
	PayNo        string `gorm:"column:pay_no;default:;NOT NULL;comment:'支付单号'"`
	PayAt        int32  `gorm:"column:pay_at;default:0;NOT NULL;comment:'支付时间'"`
	Style        string `gorm:"column:style;default:;NOT NULL;comment:'边框风格'"`
	Index        int8   `gorm:"column:index;default:0;NOT NULL;comment:'边框样式'"`
	Type         int8   `gorm:"column:type;default:50;NOT NULL;comment:'卡券类型：50-照片摆台，51-摆台拓客版'"`
	Contact      string `gorm:"column:contact;default:;NOT NULL;comment:'收货人'"`
	Mobile       string `gorm:"column:mobile;default:;NOT NULL;comment:'收货手机'"`
	Province     string `gorm:"column:province;default:;NOT NULL;comment:'省'"`
	City         string `gorm:"column:city;default:;NOT NULL;comment:'市'"`
	Area         string `gorm:"column:area;default:;NOT NULL;comment:'区'"`
	Address      string `gorm:"column:address;default:;NOT NULL;comment:'详细地址'"`
	CustomerInfo string `gorm:"column:customer_info;default:;NOT NULL;comment:'客户信息'"`
	Organ        string `gorm:"column:organ;default:;NOT NULL;comment:'客户机构'"`
	WorkNum      string `gorm:"column:work_num;default:;NOT NULL;comment:'客户工号'"`
	ThirdOrderId string `gorm:"column:third_order_id;default:;NOT NULL;comment:'第三方订单id'"`
	ThirdOrderNo string `gorm:"column:third_order_no;default:;NOT NULL;comment:'第三方订单编号'"`
	Company      int8   `gorm:"column:company;default:1;NOT NULL;comment:'照片摆台供应商：1-美印，2-影印'"`
	ShipNo       string `gorm:"column:ship_no;default:;NOT NULL;comment:'快递单号'"`
	ShipName     string `gorm:"column:ship_name;default:;NOT NULL;comment:'快递公司'"`
	ShipCom      string `gorm:"column:ship_com;default:;NOT NULL;comment:'快递公司代号'"`
	ShipTime     int32  `gorm:"column:ship_time;default:0;NOT NULL;comment:'发货时间'"`
	SendTime     int32  `gorm:"column:send_time;default:0;NOT NULL;comment:'推送时间'"`
	Status       int8   `gorm:"column:status;default:0;NOT NULL;comment:'订单状态：-1-已取消，0-已下单，1-已同步，2-已发货，3-已完成'"`
	CTime        int32  `gorm:"column:c_time;default:0;NOT NULL;comment:'创建时间'"`
	UTime        int32  `gorm:"column:u_time;default:0;NOT NULL;comment:'更新时间'"`
}

type PhotoCy struct {
	OrderNo      string `json:"order_no" tag:"订单编号"`      // 订单编号
	BatchNum     string `json:"batch_num" tag:"优惠券批次号"`   // 优惠券批次号
	ProName      string `json:"pro_name" tag:"产品名称"`      // 产品名称
	Contact      string `json:"contact" tag:"收货人"`        // 收货人
	Mobile       string `json:"mobile" tag:"收货手机"`        // 收货手机
	Province     string `json:"province" tag:"省"`         // 省
	City         string `json:"city" tag:"市"`             // 市
	Area         string `json:"area" tag:"区"`             // 区
	Address      string `json:"address" tag:"详细地址"`       // 详细地址
	CustomerInfo string `json:"customer_info" tag:"客户信息"` // 客户信息
	WorkNum      string `json:"work_num" tag:"客户工号"`      // 客户工号
	Company      int8   `json:"company" tag:"照片摆台供应商"`    // 照片摆台供应商：1-美印，2-影印
	ShipNo       string `json:"ship_no" tag:"快递单号"`       // 快递单号
	ShipName     string `json:"ship_name" tag:"快递公司"`     // 快递公司
	CTime        int32  `json:"c_time" tag:"创建时间"`        // 创建时间
}

func (c *CarOrderPhoto) TableName() string {
	return "car_order_photo"
}
