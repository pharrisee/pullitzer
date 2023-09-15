package hx

func (c *Ctx) HXIsHTMX() bool {
	return c.Request().Header.Get("HX-Request") == "true"
}

func (c *Ctx) HXIsBoosted() bool {
	return c.Request().Header.Get("HX-Boosted") == "true"
}

func (c *Ctx) HXCurrentURL() string {
	return c.Request().Header.Get("HX-Current-URL")
}

func (c *Ctx) HXHistoryRestore() bool {
	return c.Request().Header.Get("HX-History-Restore") == "true"
}

func (c *Ctx) HXPrompt() string {
	return c.Request().Header.Get("HX-Prompt")
}

func (c *Ctx) HXTarget() string {
	return c.Request().Header.Get("HX-Target")
}

func (c *Ctx) HXTrigger() string {
	return c.Request().Header.Get("HX-Trigger")
}

func (c *Ctx) HXTriggerName() string {
	return c.Request().Header.Get("HX-Trigger-Name")
}
