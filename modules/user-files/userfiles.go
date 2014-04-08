package userfiles

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/revel/revel"
)

// a table to store agglomeration of Id's of users and data sets, meta info, etc
type UserFileInfo struct {
	// gorm fields
	Id        int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time

	// keys to the data table
	UserId int64 `sql:"not null"`
	FileId int64 `sql:"not null;unique"`

	Name   string
	Path   string
	Type   string
	Folder bool
	Size   int
}

type UserFile struct {
	// gorm fields
	Id        int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time

	// keys to the data table
	UserId int64 `sql:"not null"`
	FileId int64 `sql:"not null;unique"`

	Content []byte
}

type UserDataPermissions struct {
	// gorm fields
	Id        int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time

	// keys to the data table
	UserId int64 `sql:"not null"`
	FileId int64 `sql:"not null;unique"`

	// need to use RBAC
	// or
	// maybe Postgres has RBAC at the table level
	// if so, each user could have own table(s) DB(s)
	Permissions int64
}

type UserFileWire struct {
	Name    string
	Path    string
	Type    string
	Folder  bool
	Size    int
	Content string
}

func AddTables(db *gorm.DB) error {
	err := db.AutoMigrate(UserFileInfo{}).Error
	if err != nil {
		return err
	}
	return db.AutoMigrate(UserFile{}).Error
}

func DropTables(db *gorm.DB) error {
	err := db.DropTable(UserFileInfo{}).Error
	if err != nil {
		return err
	}
	return db.DropTable(UserFile{}).Error
}

func FillTables(db *gorm.DB) error {
	return errors.New("TODO")
}
func TestTables(db *gorm.DB) error {
	return errors.New("TODO")
}

func AddUserFile(db *gorm.DB, udt *UserFileInfo, content []byte) error {

	lastId, err := getLastFileIdByUserId(db, udt.UserId)
	if err != nil {
		revel.TRACE.Println(err)
		return err
	}
	nextId := lastId + 1
	// revel.INFO.Println("Next DS id: ", nextId)
	udt.FileId = nextId

	err = db.Save(udt).Error
	if err != nil {
		revel.TRACE.Println(err)
		return err
	}

	ds := &UserFile{
		UserId:  udt.UserId,
		FileId:  udt.FileId,
		Content: content,
	}

	err = db.Save(ds).Error
	if err != nil {
		revel.TRACE.Println(err)
		return err
	}

	return err
}

func getLastFileIdByUserId(db *gorm.DB, uId int64) (int64, error) {
	var udt UserFileInfo
	err := db.Where(&UserFileInfo{UserId: uId}).Order("file_id desc").First(&udt).Error
	if err == gorm.RecordNotFound {
		return 0, nil
	}
	if err != nil {
		revel.TRACE.Println(err)
		return 0, err
	}
	return udt.FileId, nil
}

func GetUserFileInfos(db *gorm.DB, uId int64) ([]UserFileInfo, error) {
	var err error
	var infos []UserFileInfo
	err = db.Where(&UserFileInfo{UserId: uId}).Order("file_id").Find(&infos).Error
	if err != nil {
		revel.ERROR.Println(err)
		return nil, err
	}
	return infos, nil
}

func GetUserFileById(db *gorm.DB, uId, dsId int64) (ds *UserFile, err error) {
	err = db.Where(&UserFileInfo{UserId: uId, FileId: dsId}).First(ds).Error
	if err != nil {
		revel.TRACE.Println(err)
		return nil, err
	}
	return
}

func GetUserFileInfoById(db *gorm.DB, uId, dsId int64) (info *UserFileInfo, err error) {
	err = db.Where(&UserFileInfo{UserId: uId, FileId: dsId}).First(info).Error
	if err != nil {
		revel.TRACE.Println(err)
		return nil, err
	}
	return
}

func UpdateUserFileInfo(db *gorm.DB, info *UserFileInfo) error {
	err := db.Save(info).Error
	if err != nil {
		revel.TRACE.Println(err)
		return err
	}
	return nil
}

func DeleteUserFileById(db *gorm.DB, uId, dsId int64) (bool, error) {
	// delete file info
	err := db.Where(&UserFileInfo{UserId: uId, FileId: dsId}).Delete(&UserFileInfo{}).Error
	if err != nil {
		revel.TRACE.Println(err)
		return false, err
	}

	// delete file contents
	err = db.Where(&UserFile{UserId: uId, FileId: dsId}).Delete(&UserFile{}).Error
	if err != nil {
		revel.TRACE.Println(err)
		return false, err
	}

	return true, nil
}
