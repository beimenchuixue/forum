package user

import (
	"forum/utils/names"
	"github.com/beego/beego/v2/core/validation"
)

// 对数据进行校验
// ValidData 对用户登录提交的数据进行校验
func ValidData(sf *User) (ok bool, err error, errMap map[string]interface{}) {
	//2. 参数校验
	valid := validation.Validation{}
	ok, err = valid.Valid(sf)
	// 处理验证错误
	if err != nil {
		return false, err, nil
	}

	// 验证失败返回错误信息
	if !ok {
		// e.field e.tmpl
		errors := valid.Errors
		errMap = make(map[string]interface{}, len(errors))
		for _, e := range errors {
			errMap[names.UnMarshal(e.Field)] = e.Tmpl
		}
		return ok, nil, errMap
		//for _, e := range valid.Errors {
		//	fmt.Println("name:", e.Name)
		//	fmt.Println("key:", e.Key)
		//	fmt.Println("value:", e.Value)
		//	fmt.Println("field:", e.Field)
		//	fmt.Println("msg:", e.Message)
		//	fmt.Println("tmpl:", e.Tmpl)
		//
	}
	return ok, nil, nil
}
