package frps

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/xxl6097/go-frp-panel/pkg/utils"
	utils2 "github.com/xxl6097/go-service/gservice/utils"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type User struct {
	ID         string   `json:"id,omitempty"`
	User       string   `json:"user,omitempty"`
	Token      string   `json:"token,omitempty"`
	Comment    string   `json:"comment,omitempty"`
	Ports      []any    `json:"ports,omitempty"`
	Domains    []string `json:"domains,omitempty"`
	Subdomains []string `json:"subdomains,omitempty"`
	Enable     bool     `json:"enable,omitempty"`
}

func (u *User) CreateUserByID() error {
	if u.User == "" {
		return errors.New("用户名空")
	}
	if u.Token == "" {
		return errors.New("凭证空")
	}
	if u.ID == "" {
		return errors.New("ID空")
	}
	userFilePath := GetJsonPath(u.ID)
	if utils2.FileExists(userFilePath) {
		return errors.New("user already exists")
	}

	err := os.MkdirAll(filepath.Dir(userFilePath), 0755)
	if err != nil {
		log.Fatal(err)
		return err
	}
	// 序列化为 JSON（带缩进格式化）
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return err
	}
	return utils.Write(userFilePath, jsonData)
}

func (u *User) UpdateUser() error {
	if u.Token == "" {
		return errors.New("token is empty")
	}
	if u.ID == "" {
		return errors.New("id is empty")
	}
	userFilePath := GetJsonPath(u.ID)
	if utils2.FileExists(userFilePath) {
		err := os.Remove(userFilePath)
		if err != nil {
			return err
		}
	}

	err := os.MkdirAll(filepath.Dir(userFilePath), 0755)
	if err != nil {
		log.Fatal(err)
		return err
	}
	// 序列化为 JSON（带缩进格式化）
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return err
	}
	return utils.Write(userFilePath, jsonData)
}

func GetUserDir() string {
	binpath, err := os.Executable()
	if err != nil {
		return ""
	}
	return filepath.Join(filepath.Dir(binpath), "user")
}
func GetJsonPath(fileName string) string {
	return filepath.Join(GetUserDir(), fmt.Sprintf("%s.json", fileName))
}

func Read(filePath string) (*User, error) {
	content, err := utils.Read(filePath)
	if err != nil {
		return nil, err
	}
	var u User
	err = json.Unmarshal(content, &u)
	return &u, err
}

func DeleteUser(id string) error {
	if id == "" {
		return errors.New("id is empty")
	}
	return os.Remove(GetJsonPath(id))
}

func GetUserAll() ([]User, error) {
	binpath, err := os.Executable()
	if err != nil {
		return nil, err
	}
	files, err := filepath.Glob(filepath.Join(filepath.Dir(binpath), "user", "*.json"))
	if err != nil {
		return nil, err
	}
	var users []User
	for _, file := range files {
		user, err := Read(file)
		if err == nil {
			users = append(users, *user)
		}
	}
	return users, nil
}

//	type UserGorm struct {
//		gorm.Model
//		Name string `gorm:"column:name;unique;not null;default:'';comment:'用户名'"`
//		User datatypes.JSONType[User]
//	}
//
//	type UserRepository struct {
//		db *gorm.DB
//	}
//
//	func (u UserGorm) TableName() string {
//		return "frps_user_table"
//	}
//
//	func NewUserRepository(db *gorm.DB) *UserRepository {
//		err := db.Debug().AutoMigrate(&UserGorm{})
//		if err != nil {
//			fmt.Println("user table created failed", err)
//		} else {
//			fmt.Println("user table created")
//		}
//		return &UserRepository{
//			db: db,
//		}
//	}
//
//	func (this *UserRepository) Create(obj *UserGorm) error {
//		return this.db.Create(obj).Error
//	}
//
//	func (this *UserRepository) Delete(name string) error {
//		//return this.db.Where("id = ?", Key.User).Unscoped().Delete(&Key).Error
//		return this.db.Where("name = ?", name).Delete(&UserGorm{}).Error
//	}
//
//	func (this *UserRepository) Update(u *UserGorm) error {
//		//fmt.Printf("update token %+v\n", obj)
//		//return this.db.Model(&Token{}).Omit("enable").Where("id = ?", obj.ID).Updates(obj).Error
//		//return this.db.Model(&Token{}).Where("id = ?", obj.ID).Save(obj).Error
//		return this.db.Model(&UserGorm{}).Where("name = ?", u.Name).Updates(u).Error
//	}
//
//	func (this *UserRepository) Find(name string) (User, error) {
//		var u UserGorm
//		err := this.db.Where("name = ?", name).First(&u).Error
//		return u.User.Data(), err
//	}
//
//	func (this *UserRepository) FindAll() ([]UserGorm, error) {
//		//this.db.Where(datatypes.JSONQuery("conf").Equals("dark", "theme")).Find(&result)
//		var us []UserGorm
//		err := this.db.Find(&us).Error
//		return us, err
//	}

func ToPorts(ports []any) []any {
	ps := make([]interface{}, len(ports))
	for _, port := range ports {
		if str, ok := port.(string); ok {
			ps = append(ps, str)
		} else if ints, ok := port.(int); ok {
			ps = append(ps, ints)
		}
	}
	return ps
}

func JudgeToken(id, token string) (bool, error) {
	u, err := Read(GetJsonPath(id))
	if err != nil {
		return false, fmt.Errorf("[token校验]frps服务器不存在该用户ID:%v,err: %v", id, err)
	}
	if u.Token != token {
		return false, fmt.Errorf("[token校验]ID【%s】的token【%s】校验错误❌", id, token)
	}
	if !u.Enable {
		return false, fmt.Errorf("[token校验]用户【%s】被禁用", id)
	}
	return true, nil
}

func JudgePort(id, proxyType string, userPort int, userDomains []string, userSubdomain string) (bool, error) {
	u, err := Read(GetJsonPath(id))
	//u, err := this.Find(user)
	if err != nil {
		return false, fmt.Errorf("[port校验]frps服务器不存在该用户:%v", err)
	}
	var portErr error
	portAllowed := true
	var reject = false
	if proxyType == "tcp" || proxyType == "udp" {
		portAllowed = false
		for _, port := range u.Ports {
			if str, ok := port.(string); ok {
				if strings.Contains(str, "-") {
					allowedRanges := strings.Split(str, "-")
					if len(allowedRanges) != 2 {
						portErr = fmt.Errorf("id [%v] port range [%v] format error", id, port)
						break
					}
					start, err := strconv.Atoi(strings.TrimSpace(allowedRanges[0]))
					if err != nil {
						portErr = fmt.Errorf("id [%v] port rang [%v] start port [%v] is not a number", id, port, allowedRanges[0])
						break
					}
					end, err := strconv.Atoi(strings.TrimSpace(allowedRanges[1]))
					if err != nil {
						portErr = fmt.Errorf("id [%v] port rang [%v] end port [%v] is not a number", id, port, allowedRanges[0])
						break
					}
					if max(int64(userPort), int64(start)) == int64(userPort) && min(int64(userPort), int64(end)) == int64(userPort) {
						portAllowed = true
						break
					}
				} else {
					if str == "" {
						portAllowed = true
						break
					}
					allowed, err := strconv.Atoi(str)
					if err != nil {
						portErr = fmt.Errorf("id [%v] allowed port [%v] is not a number", id, port)
					}
					if int64(allowed) == int64(userPort) {
						portAllowed = true
						break
					}
				}
			} else {
				num, okk := port.(float64)
				if okk && num == float64(userPort) {
					portAllowed = true
					break
				}
				//allowed := int64(port)
				//if int64(allowed) == int64(userPort) {
				//	portAllowed = true
				//	break
				//}
			}

		}
	}
	//判断port是否合法
	if !portAllowed {
		if portErr == nil {
			portErr = fmt.Errorf("id [%v] port [%v] is not allowed", id, userPort)
		}
		reject = true
	}

	domainAllowed := true
	if proxyType == "http" || proxyType == "https" || proxyType == "tcpmux" {
		if portAllowed {
			if utils.StringContains("", u.Domains) {
				domainAllowed = true
			} else {
				for _, userDomain := range userDomains {
					if !utils.StringContains(userDomain, u.Domains) {
						domainAllowed = false
						break
					}
				}
			}
			if !domainAllowed {
				portErr = fmt.Errorf("id [%v] domain [%v] is not allowed", id, strings.Join(userDomains, ","))
				reject = true
			}
		}
	}

	subdomainAllowed := true
	if proxyType == "http" || proxyType == "https" {
		subdomainAllowed = false
		if portAllowed && domainAllowed {
			arr := u.Subdomains
			if utils.StringContains("", arr) {
				subdomainAllowed = true
			} else {
				for _, subdomain := range arr {
					if subdomain == userSubdomain {
						subdomainAllowed = true
						break
					}
				}
			}
			if !subdomainAllowed {
				portErr = fmt.Errorf("id [%v] subdomain [%v] is not allowed", id, userSubdomain)
				reject = true
			}
		}
	}
	return reject, portErr
}
