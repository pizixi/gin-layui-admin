package router

import (
	"embed"
	"gin-layui-admin/controllers"
	"gin-layui-admin/web"
	"html/template"
	"io/fs"
	"net/http"
	"strings"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func loadTemplatesFromEmbedFS(embedFS embed.FS, filenames ...string) (*template.Template, error) {
	tmpl := template.New("")
	for _, filename := range filenames {
		file, err := embedFS.ReadFile(filename)
		if err != nil {
			return nil, err
		}
		tmpl, err = tmpl.Parse(string(file))
		if err != nil {
			return nil, err
		}
	}
	return tmpl, nil
}

func createMultiTpl(r *gin.Engine) {
	render := multitemplate.NewRenderer()

	// 从embed.FS加载模板
	templates := map[string][]string{
		"login": {"views/login/login.html"},
		"home":  {"views/public/main.html"},
		"start": {"views/public/layout.html", "views/home/start.html"},
		// 角色管理
		"roleIndex": {"views/public/layout.html", "views/role/list.html"},
		"roleEdit":  {"views/public/layout.html", "views/role/edit.html"},
		"roleAdd":   {"views/public/layout.html", "views/role/add.html"},
		// 用户管理
		"userIndex": {"views/public/layout.html", "views/user/list.html"},
		"userAdd":   {"views/public/layout.html", "views/user/add.html"},
		"userEdit":  {"views/public/layout.html", "views/user/edit.html"},
		// 菜单管理
		"menuIndex": {"views/public/layout.html", "views/menu/list.html"},
		// 个人中心
		"personal": {"views/public/layout.html", "views/personal/edit.html"},
	}
	for name, files := range templates {
		tmpl, err := loadTemplatesFromEmbedFS(web.ViewsFS, files...)
		if err != nil {
			panic(err)
		}
		render.Add(name, tmpl)
	}

	// // 从文件系统加载模板（原）
	// render.AddFromFiles("login", "web/views/login/login.html")
	// render.AddFromFiles("home", "web/views/public/main.html")
	// render.AddFromFiles("start", "web/views/public/layout.html", "web/views/home/start.html")

	// // 角色管理
	// render.AddFromFiles("roleIndex", "web/views/public/layout.html", "web/views/role/list.html")
	// render.AddFromFiles("roleEdit", "web/views/public/layout.html", "web/views/role/edit.html")
	// render.AddFromFiles("roleAdd", "web/views/public/layout.html", "web/views/role/add.html")

	// // 用户管理
	// render.AddFromFiles("userIndex", "web/views/public/layout.html", "web/views/user/list.html")
	// render.AddFromFiles("userAdd", "web/views/public/layout.html", "web/views/user/add.html")
	// render.AddFromFiles("userEdit", "web/views/public/layout.html", "web/views/user/edit.html")

	// // 菜单管理
	// render.AddFromFiles("menuIndex", "web/views/public/layout.html", "web/views/menu/list.html")

	// // 个人中心
	// render.AddFromFiles("personal", "web/views/public/layout.html", "web/views/personal/edit.html")

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

	// // 静态资源
	// r.Static("/static", "./web/static")

	// 使用embed打包静态资源
	staticRootFS, _ := fs.Sub(web.StaticFS, "static")
	r.StaticFS("/static", http.FS(staticRootFS))

	// login
	r.GET("/", controllers.LoginIndex)
	r.POST("/login", controllers.Login)
	r.GET("/logout", controllers.Logout)

	// private
	private := r.Group("/home")
	private.Use(AuthRequired())
	{
		private.GET("/", controllers.HomeIndex)
		private.GET("/start", controllers.HomeStartIndex)

		// 角色管理
		private.GET("/role", controllers.RoleIndex)
		private.GET("/role/table", controllers.RoleTable)
		private.GET("/role/add", controllers.RoleAddIndex)
		private.GET("/role/edit", controllers.RoleEditIndex)
		private.POST("/role/ajaxsave", controllers.RoleAjaxSave)
		private.POST("/role/ajaxdel", controllers.RoleAjaxDel)

		// 用户管理
		private.GET("/user", controllers.UserIndex)
		private.GET("/user/table", controllers.UserTable)
		private.GET("/user/add", controllers.UserAddIndex)
		private.GET("/user/edit", controllers.UserEditIndex)
		private.POST("/user/ajaxsave", controllers.UserAjaxSave)
		private.POST("/user/ajaxdel", controllers.UserAjaxDel)

		// 菜单管理
		private.GET("/menu", controllers.MenuIndex)
		private.POST("/menu/getnode", controllers.GetNode)
		private.POST("/menu/getnodes", controllers.GetNodes)
		private.POST("/menu/ajaxsave", controllers.MenuAjaxSave)
		private.POST("/menu/ajaxdel", controllers.MenuAjaxDel)

		// 个人中心
		private.GET("/personal", controllers.PersonalIndex)
		private.POST("/personal/ajaxsave", controllers.PersonalAjaxSave)

	}

}
