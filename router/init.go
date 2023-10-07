package router

import (
	"net/http"
	"strings"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func createMultiTpl(r *gin.Engine) {
	render := multitemplate.NewRenderer()
	render.AddFromFiles("login", "views/login/login.html")
	render.AddFromFiles("home", "views/public/main.html")
	render.AddFromFiles("start", "views/public/layout.html", "views/home/start.html")

	// 角色管理
	render.AddFromFiles("roleIndex", "views/public/layout.html", "views/role/list.html")
	render.AddFromFiles("roleEdit", "views/public/layout.html", "views/role/edit.html")
	render.AddFromFiles("roleAdd", "views/public/layout.html", "views/role/add.html")

	// 用户管理
	render.AddFromFiles("userIndex", "views/public/layout.html", "views/user/list.html")
	render.AddFromFiles("userAdd", "views/public/layout.html", "views/user/add.html")
	render.AddFromFiles("userEdit", "views/public/layout.html", "views/user/edit.html")

	// 菜单管理
	render.AddFromFiles("menuIndex", "views/public/layout.html", "views/menu/list.html")

	// 个人中心
	render.AddFromFiles("personal", "views/public/layout.html", "views/personal/edit.html")

	r.HTMLRender = render
}

// AuthRequired AuthRequired
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("auth")
		if user == nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"flashError": "没有权限"})
		} else {
			if cookie, ok := user.(string); ok {
				s := strings.Split(cookie, "|")
				if len(s) == 3 {
					c.Set("userId", s[0])
					c.Set("userName", s[1])
					c.Next()
					return
				}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"flashError": "没有权限"})
		}
	}
}

// RegRouter RegRouter
func RegRouter(r *gin.Engine) {
	createMultiTpl(r)

	// session中间件
	store := sessions.NewCookieStore([]byte("secret"))
	store.Options(sessions.Options{
		MaxAge: 7 * 86400,
		// Path:   "/",
	})
	r.Use(sessions.Sessions("kgogame", store))

	// 静态资源
	r.Static("/static", "./static")

	// login
	r.GET("/", LoginIndex)
	r.POST("/login", Login)
	r.GET("/logout", Logout)

	// private
	private := r.Group("/home")
	private.Use(AuthRequired())
	{
		private.GET("/", HomeIndex)
		private.GET("/start", HomeStartIndex)

		// 角色管理
		private.GET("/role", RoleIndex)
		private.GET("/role/table", RoleTable)
		private.GET("/role/add", RoleAddIndex)
		private.GET("/role/edit", RoleEditIndex)
		private.POST("/role/ajaxsave", RoleAjaxSave)

		// 用户管理
		private.GET("/user", UserIndex)
		private.GET("/user/table", UserTable)
		private.GET("/user/add", UserAddIndex)
		private.GET("/user/edit", UserEditIndex)
		private.POST("/user/ajaxsave", UserAjaxSave)
		private.POST("/user/ajaxdel", UserAjaxDel)

		// 菜单管理
		private.GET("/menu", MenuIndex)
		private.POST("/menu/getnode", GetNode)
		private.POST("/menu/getnodes", GetNodes)
		private.POST("/menu/ajaxsave", MenuAjaxSave)
		private.POST("/menu/ajaxdel", MenuAjaxDel)

		// 个人中心
		private.GET("/personal", PersonalIndex)
		private.POST("/personal/ajaxsave", PersonalAjaxSave)

	}

}
