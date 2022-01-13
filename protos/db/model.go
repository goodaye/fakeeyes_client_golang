package db

import "time"

// ID          int       `json:"id" xorm:"not null pk autoincr INT"`
// Name        string    `json:"name" xorm:"not null unique VARCHAR(255) comment('产品名称')"`
// BizModules  string    `json:"biz_modules" xorm:"not null  VARCHAR(1024) comment('产品关联模块')"`
// Admin       string    `json:"admin" xorm:"VARCHAR(255)  comment('产品admin')"`
// GmtModified time.Time `json:"gmt_modified" xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP updated" `
// GmtCreated  time.Time `json:"gmt_created" xorm:"not null default '1970-01-01 08:00:01' TIMESTAMP created" `

// User user
type User struct {
	ID          int64     `json:"id" xorm:"not null  pk autoincr INT"`
	Name        string    `json:"name" xorm:"not null unique VARCHAR(255) comment('用户名')"`
	LastLogin   time.Time `json:"last_login" xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP" `
	GmtCreated  time.Time `json:"gmt_created" xorm:"not null default '1970-01-01 08:00:01' TIMESTAMP created" `
	GmtModified time.Time `json:"gmt_modified" xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP updated" `
}

type UserSession struct {
	ID          int64     `json:"id" xorm:"not null pk autoincr INT"`
	UserID      int64     `json:"user_id" xorm:"not null BIGINT unique comment('用户ID')"`
	Token       string    `json:"token" xorm:"not null  VARCHAR(255) unique comment('用户token')"`
	ExpireTime  time.Time `json:"expire_time" xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP" `
	GmtCreated  time.Time `json:"gmt_created" xorm:"not null default '1970-01-01 08:00:01' TIMESTAMP created" `
	GmtModified time.Time `json:"gmt_modified" xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP updated" `
}

type DeviceSession struct {
}

// Device
type Device struct {
	ID           int64     `json:"id" xorm:"not null pk autoincr INT"`
	SN           string    `json:"sn" xorm:"not null unique VARCHAR(255) comment('设备SN')"`
	Name         string    `json:"name" xorm:"not null  VARCHAR(255) comment('设备名')"`
	Type         string    `json:"type" xorm:" null  VARCHAR(255) comment('设备类型,比如4WD')"`
	Class        string    `json:"class" xorm:" null  VARCHAR(255) comment('设备类型,比如,树莓派小车')"`
	Manufacturer string    `json:"manufacturer" xorm:" null  VARCHAR(255) comment('设备制造商,比如,亚博')"`
	Status       int       `json:"status" xorm:" null  VARCHAR(255) comment('设备制状态，比如，在线、离线')"`
	LastLogin    time.Time `json:"last_login" xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP comment('上次登陆时间')" `
	GmtCreated   time.Time `json:"gmt_created" xorm:"not null default '1970-01-01 08:00:01' TIMESTAMP created" `
	GmtModified  time.Time `json:"gmt_modified" xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP updated" `
}

// Room 房间
type Room struct {
	ID          int64     `json:"id" xorm:"not null pk autoincr INT"`
	UUID        string    `json:"token" xorm:"not null unique  VARCHAR(255) unique comment('房间UUID')"`
	Name        string    `json:"name" xorm:"not null VARCHAR(255) comment('房间名')"`
	Status      int       `json:"status" xorm:" null  VARCHAR(255) comment('设备制状态')"`
	GmtCreated  time.Time `json:"gmt_created" xorm:"not null default '1970-01-01 08:00:01' TIMESTAMP created" `
	GmtModified time.Time `json:"gmt_modified" xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP updated" `
}

//RoomUser 房间中的人
type RoomUser struct {
	ID          int64 `json:"id" xorm:"not null pk autoincr INT"`
	RoomID      int64
	UserID      int64
	Role        int
	Status      int
	GmtCreated  time.Time `json:"gmt_created" xorm:"not null default '1970-01-01 08:00:01' TIMESTAMP created" `
	GmtModified time.Time `json:"gmt_modified" xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP updated" `
}

//房间中的设备
type RoomDevice struct {
	ID          int64 `json:"id" xorm:"not null pk autoincr INT"`
	RoomID      int64
	DeviceID    int64
	GmtCreated  time.Time `json:"gmt_created" xorm:"not null default '1970-01-01 08:00:01' TIMESTAMP created" `
	GmtModified time.Time `json:"gmt_modified" xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP updated" `
}
