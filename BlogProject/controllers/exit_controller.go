package controllers

type ExitController struct {
	BaseController
}

func (c *ExitController) Get()  {
	// 清楚该用户登录状态的数据
	c.DelSession("loginuser")
	c.Redirect("/", 302)
}
