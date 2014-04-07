package user

import (
	"errors"

	"github.com/jinzhu/gorm"
)

// Union of all the User* tables
type User struct {
	Id     int64
	UserId int64

	Basic        UserBasic
	Detail       UserDetail
	Addresses    []UserAddress
	Phones       []UserPhone
	ProfileElems []UserProfileElement
}

func GetUserIdFromUserName(db *gorm.DB, username string) (int64, error) {
	u, err := getUserBasicByName(db, username)
	if err != nil {
		return -1, err
	}
	return u.UserId, nil
}

func GetUserById(db *gorm.DB, uId int64) (*User, error) {
	// var u User
	// u.UserId = uId
	return nil, errors.New("GetUserById not implemented")
}

var MAX_ID_RETRY int = 10

func GenerateNewVisitorId(db *gorm.DB) (int64, error) {
	var unique bool
	var tries int
	for !unique {

		// TODO  some better conflict resolution, we just want them random and not sequential
		if tries > MAX_ID_RETRY {
			return -1, errors.New("too many attempts at generating a unqiue id")
		}

		// generate id
		id := generateRandId()

		// check for existance
		u, err := getVisitorByUserId(db, id)
		if u == nil {
			return id, nil
		} else if err != nil {
			// probably can't get here
			return -1, err
		} else {
			tries++
		}
	}
	panic("shouldn't get here")
	return -1, nil
}

func GenerateNewUserId(db *gorm.DB) (int64, error) {
	var unique bool
	var tries int
	for !unique {

		// TODO  some better conflict resolution, we just want them random and not sequential
		if tries > MAX_ID_RETRY {
			return -1, errors.New("too many attempts at generating a unqiue id")
		}

		// generate id
		id := generateRandId()

		// check for existance
		u, err := getUserBasicByUserId(db, id)
		if u == nil {
			return id, nil
		} else if err != nil {
			// probably can't get here
			return -1, err
		} else {
			tries++
		}
	}
	panic("shouldn't get here")
	return -1, nil
}

func AddVisitor(db *gorm.DB, uId int64, ip string) error {
	u := &Visitor{
		VisitorId: uId,
		VisitorIp: ip,
	}
	return addVisitor(db, u)
}
func UpdateVisitor(db *gorm.DB, v *Visitor) error {
	return addVisitor(db, v)
}

func DeleteVisitor(db *gorm.DB, uId int64) error {
	return deleteVisitor(db, uId)
}

func GetVisitorByVisitorId(db *gorm.DB, uId int64) (*Visitor, error) {
	return getVisitorByVisitorId(db, uId)
}

func GetVisitorByIp(db *gorm.DB, ip string) (*Visitor, error) {
	return getVisitorByVisitorIp(db, ip)
}

func AddUserBasic(db *gorm.DB, uId int64, username string) error {
	u := &UserBasic{
		UserId:   uId,
		UserName: username,
	}
	return addUserBasic(db, u)
}

func DeleteUserBasic(db *gorm.DB, uId int64) error {
	return deleteUserBasic(db, uId)
}

func GetUserBasicById(db *gorm.DB, uId int64) (*UserBasic, error) {
	return getUserBasicByUserId(db, uId)
}

func GetUserBasicByName(db *gorm.DB, username string) (*UserBasic, error) {
	return getUserBasicByName(db, username)
}

func AddUserAddress(db *gorm.DB, u *UserAddress) error {
	return addUserAddress(db, u)
}

func DeleteUserAddress(db *gorm.DB, uId int64) error {
	return deleteUserAddress(db, uId)
}

func GetUserAddressesById(db *gorm.DB, uId int64) ([]UserAddress, error) {
	return getUserAddressesByUserId(db, uId)
}

func GetUserAddressByIdType(db *gorm.DB, uId int64, a_type string) (*UserAddress, error) {
	return getUserAddressByUserIdType(db, uId, a_type)
}

func AddUserPhone(db *gorm.DB, uId int64, p_type, number string) error {
	return addUserPhone(db, uId, p_type, number)
}

func DeleteUserPhone(db *gorm.DB, uId int64) error {
	return deleteUserPhone(db, uId)
}

func GetUserPhonesById(db *gorm.DB, uId int64) ([]UserPhone, error) {
	return getUserPhonesByUserId(db, uId)
}

func GetUserPhoneByIdType(db *gorm.DB, uId int64, a_type string) (*UserPhone, error) {
	return getUserPhoneByUserIdType(db, uId, a_type)
}

func AddUserDetail(db *gorm.DB, uId int64, u *UserDetail) error {
	return addUserDetail(db, u)
}

func DeleteUserDetail(db *gorm.DB, uId int64) error {
	return deleteUserDetail(db, uId)
}

func GetUserDetailById(db *gorm.DB, uId int64) (*UserDetail, error) {
	return getUserDetailByUserId(db, uId)
}

func AddUserProfileElement(db *gorm.DB, uId int64, field, content string) error {
	return addUserProfileElement(db, uId, field, content)
}

func DeleteAllUserProfileElements(db *gorm.DB, uId int64) error {
	return deleteAllUserProfileElements(db, uId)
}

func DeleteUserProfileElement(db *gorm.DB, uId int64, field string) error {
	return deleteUserProfileElement(db, uId, field)
}

func GetUserProfileElementsById(db *gorm.DB, uId int64) ([]UserProfileElement, error) {
	return getUserProfileElementsByUserId(db, uId)
}

func GetUserProfileElementByIdField(db *gorm.DB, uId int64, field string) (*UserProfileElement, error) {
	return getUserProfileElementByUserIdField(db, uId, field)
}
