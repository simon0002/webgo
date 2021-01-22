package models

import (
    //"fmt"
	"github.com/astaxie/beego/orm"
   // _ "github.com/go-sql-driver/mysql"
	
)

// Model Struct
type User struct {
    BaseModel

    Id          int
    Username    string 
    Password    string 
    Email       string 
    Avatar      string 
}

func init() {
    // register model
    orm.RegisterModel(new(User))
}

// func main() {
//     // o := orm.NewOrm()

//     // user := User{Name: "slene"}

//     // // insert
//     // id, err := o.Insert(&user)
//     // fmt.Printf("ID: %d, ERR: %v\n", id, err)

//     // // update
//     // user.Name = "astaxie"
//     // num, err := o.Update(&user)
//     // fmt.Printf("NUM: %d, ERR: %v\n", num, err)

//     // // read one
//     // u := User{Id: user.Id}
//     // err = o.Read(&u)
//     // fmt.Printf("ERR: %v\n", err)

//     // // delete
//     // num, err = o.Delete(&u)
//     // fmt.Printf("NUM: %d, ERR: %v\n", num, err)
// }

func UserList(offset int, limit int) []*User{
    o := orm.NewOrm()
    var user User
    var users []*User
    //users := o.Raw("SELECT * FROM user")
    o.QueryTable(user).Limit(limit, offset).All(&users)
    //num, err := o.QueryTable(user).All(&users)
    //fmt.Printf("Returned Rows Num: %s, %s", num, err)
    return users
}

func UserSave(user *User) int64{
    //user := User{Username: username, Password: password, Email: email}

    o := orm.NewOrm()
    id, _ := o.Insert(user)
    return id
}

func UserGetByName(username string) (User,error){
    o := orm.NewOrm()
    var user User
    err := o.QueryTable(user).Filter("Username", username).One(&user)

    return user, err
}

func UserCount() int64{
    o := orm.NewOrm()
    var user User
    count,_ := o.QueryTable(user).Count()
    //fmt.Println(count)
    return count
}

func UserUpdate(user *User) int64{
    o := orm.NewOrm()
    id, _ := o.Insert(user)
    return id
}