package hx

func (c *Ctx) HXLocation(href string) {
	c.SetHeader("HX-Location", href)
}

func (c *Ctx) HXPushURL(href string) {
	c.SetHeader("HX-Push-Url", href)
}

func (c *Ctx) HXRedirect(href string) {
	c.SetHeader("HX-Redirect", href)
}

func (c *Ctx) HXRefresh() {
	c.SetHeader("HX-Refresh", "true")
}

func (c *Ctx) HXReplaceUrl(href string) {
	c.SetHeader("HX-Replace-Url", href)
}

func (c *Ctx) HXReswap(swap string) {
	c.SetHeader("HX-Reswap", swap)
}

func (c *Ctx) HXReselect(swap string) {
	c.SetHeader("HX-Reselect", swap)
}

func (c *Ctx) HXSetTrigger(swap string) {
	c.SetHeader("HX-Trigger", swap)
}

func (c *Ctx) HXSetTriggerAfterSettle(swap string) {
	c.SetHeader("HX-Trigger-After-Settle", swap)
}

func (c *Ctx) HXSetTriggerAfterSwap(swap string) {
	c.SetHeader("HX-Trigger-After-Swap", swap)
}

func (c *Ctx) SetHeader(key, value string) {
	if c.HasHeader(key) {
		c.Response().Header().Set(key, value)
		return
	}
	c.Response().Header().Add(key, value)
}

func (c *Ctx) GetHeader(key string) string {
	return c.Request().Header.Get(key)
}

func (c *Ctx) HasHeader(key string) bool {
	return c.GetHeader(key) != ""
}
