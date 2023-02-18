package repository
import (
	"github.com/jinzhu/gorm"
	"github.com/rayallen20/user/domain/model"
)
type IUserRepository interface{
    InitTable() error
    FindUserByID(int64) (*model.User, error)
	CreateUser(*model.User) (int64, error)
	DeleteUserByID(int64) error
	UpdateUser(*model.User) error
	FindAll()([]model.User,error)

}
//创建userRepository
func NewUserRepository(db *gorm.DB) IUserRepository  {
	return &UserRepository{mysqlDb:db}
}

type UserRepository struct {
	mysqlDb *gorm.DB
}

//初始化表
func (u *UserRepository)InitTable() error  {
	return u.mysqlDb.CreateTable(&model.User{}).Error
}

//根据ID查找User信息
func (u *UserRepository)FindUserByID(userID int64) (user *model.User,err error) {
	user = &model.User{}
	return user, u.mysqlDb.First(user,userID).Error
}

//创建User信息
func (u *UserRepository) CreateUser(user *model.User) (int64, error) {
	return user.ID, u.mysqlDb.Create(user).Error
}

//根据ID删除User信息
func (u *UserRepository) DeleteUserByID(userID int64) error {
	return u.mysqlDb.Where("id = ?",userID).Delete(&model.User{}).Error
}

//更新User信息
func (u *UserRepository) UpdateUser(user *model.User) error {
	return u.mysqlDb.Model(user).Update(user).Error
}

//获取结果集
func (u *UserRepository) FindAll()(userAll []model.User,err error) {
	return userAll, u.mysqlDb.Find(&userAll).Error
}

