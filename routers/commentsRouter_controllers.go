package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["rbacAdmin/controllers:CourseController"] = append(beego.GlobalControllerRouter["rbacAdmin/controllers:CourseController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rbacAdmin/controllers:CourseController"] = append(beego.GlobalControllerRouter["rbacAdmin/controllers:CourseController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rbacAdmin/controllers:CourseController"] = append(beego.GlobalControllerRouter["rbacAdmin/controllers:CourseController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rbacAdmin/controllers:CourseController"] = append(beego.GlobalControllerRouter["rbacAdmin/controllers:CourseController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rbacAdmin/controllers:CourseController"] = append(beego.GlobalControllerRouter["rbacAdmin/controllers:CourseController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rbacAdmin/controllers:LoginController"] = append(beego.GlobalControllerRouter["rbacAdmin/controllers:LoginController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rbacAdmin/controllers:LoginController"] = append(beego.GlobalControllerRouter["rbacAdmin/controllers:LoginController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rbacAdmin/controllers:LoginController"] = append(beego.GlobalControllerRouter["rbacAdmin/controllers:LoginController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rbacAdmin/controllers:LoginController"] = append(beego.GlobalControllerRouter["rbacAdmin/controllers:LoginController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rbacAdmin/controllers:LoginController"] = append(beego.GlobalControllerRouter["rbacAdmin/controllers:LoginController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"get","post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rbacAdmin/controllers:LoginController"] = append(beego.GlobalControllerRouter["rbacAdmin/controllers:LoginController"],
        beego.ControllerComments{
            Method: "LoginOut",
            Router: `login_out`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rbacAdmin/controllers:MainController"] = append(beego.GlobalControllerRouter["rbacAdmin/controllers:MainController"],
        beego.ControllerComments{
            Method: "Welcome",
            Router: `/welcome`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rbacAdmin/controllers:RbacController"] = append(beego.GlobalControllerRouter["rbacAdmin/controllers:RbacController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rbacAdmin/controllers:RbacController"] = append(beego.GlobalControllerRouter["rbacAdmin/controllers:RbacController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rbacAdmin/controllers:RbacController"] = append(beego.GlobalControllerRouter["rbacAdmin/controllers:RbacController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rbacAdmin/controllers:RbacController"] = append(beego.GlobalControllerRouter["rbacAdmin/controllers:RbacController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rbacAdmin/controllers:RbacController"] = append(beego.GlobalControllerRouter["rbacAdmin/controllers:RbacController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rbacAdmin/controllers:RbacController"] = append(beego.GlobalControllerRouter["rbacAdmin/controllers:RbacController"],
        beego.ControllerComments{
            Method: "AdminAdd",
            Router: `/admin_add`,
            AllowHTTPMethods: []string{"get","post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rbacAdmin/controllers:RbacController"] = append(beego.GlobalControllerRouter["rbacAdmin/controllers:RbacController"],
        beego.ControllerComments{
            Method: "AdminDel",
            Router: `/admin_del`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rbacAdmin/controllers:RbacController"] = append(beego.GlobalControllerRouter["rbacAdmin/controllers:RbacController"],
        beego.ControllerComments{
            Method: "AdminEdit",
            Router: `/admin_edit`,
            AllowHTTPMethods: []string{"get","post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rbacAdmin/controllers:RbacController"] = append(beego.GlobalControllerRouter["rbacAdmin/controllers:RbacController"],
        beego.ControllerComments{
            Method: "AdminList",
            Router: `/admin_list`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rbacAdmin/controllers:RbacController"] = append(beego.GlobalControllerRouter["rbacAdmin/controllers:RbacController"],
        beego.ControllerComments{
            Method: "CategoryAdd",
            Router: `/category_add`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rbacAdmin/controllers:RbacController"] = append(beego.GlobalControllerRouter["rbacAdmin/controllers:RbacController"],
        beego.ControllerComments{
            Method: "CategoryDel",
            Router: `/category_del`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rbacAdmin/controllers:RbacController"] = append(beego.GlobalControllerRouter["rbacAdmin/controllers:RbacController"],
        beego.ControllerComments{
            Method: "CategoryEdit",
            Router: `/category_edit`,
            AllowHTTPMethods: []string{"get","post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rbacAdmin/controllers:RbacController"] = append(beego.GlobalControllerRouter["rbacAdmin/controllers:RbacController"],
        beego.ControllerComments{
            Method: "CategoryList",
            Router: `/category_list`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rbacAdmin/controllers:RbacController"] = append(beego.GlobalControllerRouter["rbacAdmin/controllers:RbacController"],
        beego.ControllerComments{
            Method: "PermissionAdd",
            Router: `/permission_add`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rbacAdmin/controllers:RbacController"] = append(beego.GlobalControllerRouter["rbacAdmin/controllers:RbacController"],
        beego.ControllerComments{
            Method: "PermissionDel",
            Router: `/permission_del`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rbacAdmin/controllers:RbacController"] = append(beego.GlobalControllerRouter["rbacAdmin/controllers:RbacController"],
        beego.ControllerComments{
            Method: "PermissionEdit",
            Router: `/permission_edit`,
            AllowHTTPMethods: []string{"get","post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rbacAdmin/controllers:RbacController"] = append(beego.GlobalControllerRouter["rbacAdmin/controllers:RbacController"],
        beego.ControllerComments{
            Method: "PermissionList",
            Router: `/permission_list`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rbacAdmin/controllers:RbacController"] = append(beego.GlobalControllerRouter["rbacAdmin/controllers:RbacController"],
        beego.ControllerComments{
            Method: "RoleAdd",
            Router: `/role_add`,
            AllowHTTPMethods: []string{"get","post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rbacAdmin/controllers:RbacController"] = append(beego.GlobalControllerRouter["rbacAdmin/controllers:RbacController"],
        beego.ControllerComments{
            Method: "RoleDel",
            Router: `/role_del`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rbacAdmin/controllers:RbacController"] = append(beego.GlobalControllerRouter["rbacAdmin/controllers:RbacController"],
        beego.ControllerComments{
            Method: "RoleEdit",
            Router: `/role_edit`,
            AllowHTTPMethods: []string{"get","post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rbacAdmin/controllers:RbacController"] = append(beego.GlobalControllerRouter["rbacAdmin/controllers:RbacController"],
        beego.ControllerComments{
            Method: "RoleList",
            Router: `/role_list`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rbacAdmin/controllers:TeacherController"] = append(beego.GlobalControllerRouter["rbacAdmin/controllers:TeacherController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rbacAdmin/controllers:TeacherController"] = append(beego.GlobalControllerRouter["rbacAdmin/controllers:TeacherController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rbacAdmin/controllers:TeacherController"] = append(beego.GlobalControllerRouter["rbacAdmin/controllers:TeacherController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rbacAdmin/controllers:TeacherController"] = append(beego.GlobalControllerRouter["rbacAdmin/controllers:TeacherController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rbacAdmin/controllers:TeacherController"] = append(beego.GlobalControllerRouter["rbacAdmin/controllers:TeacherController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rbacAdmin/controllers:UserController"] = append(beego.GlobalControllerRouter["rbacAdmin/controllers:UserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rbacAdmin/controllers:UserController"] = append(beego.GlobalControllerRouter["rbacAdmin/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rbacAdmin/controllers:UserController"] = append(beego.GlobalControllerRouter["rbacAdmin/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rbacAdmin/controllers:UserController"] = append(beego.GlobalControllerRouter["rbacAdmin/controllers:UserController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rbacAdmin/controllers:UserController"] = append(beego.GlobalControllerRouter["rbacAdmin/controllers:UserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rbacAdmin/controllers:VideoController"] = append(beego.GlobalControllerRouter["rbacAdmin/controllers:VideoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rbacAdmin/controllers:VideoController"] = append(beego.GlobalControllerRouter["rbacAdmin/controllers:VideoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rbacAdmin/controllers:VideoController"] = append(beego.GlobalControllerRouter["rbacAdmin/controllers:VideoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rbacAdmin/controllers:VideoController"] = append(beego.GlobalControllerRouter["rbacAdmin/controllers:VideoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rbacAdmin/controllers:VideoController"] = append(beego.GlobalControllerRouter["rbacAdmin/controllers:VideoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
