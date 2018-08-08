package resources

import (
	"fmt"
	"github.com/beevik/etree"
	"strconv"
)

// User struct
type User struct {
	XMLData *etree.Element
}

// CreateUser constructs User
func CreateUser(id int) *User {
	doc := etree.NewDocument()
	doc.CreateProcInst("xml", `version="1.0"`)

	el := doc.CreateElement("USER").CreateElement("ID")
	el.SetText(fmt.Sprintf("%d", id))

	return &User{doc.Root()}
}

// Attribute method
func (u User) Attribute(path string) string {
	elements := u.XMLData.FindElements(path)
	if elements == nil {
		return ""
	}
	return elements[0].Text()
}

// ID method
func (u User) ID() int {
	i, err := strconv.Atoi(u.Attribute("ID"))
	if err != nil {
		return -1
	}
	return i
}

// Name method
func (u User) Name() string {
	return u.Attribute("NAME")
}

// Password method
func (u User) Password() string {
	return u.Attribute("PASSWORD")
}

// AuthDriver method
func (u User) AuthDriver() string {
	return u.Attribute("AUTH_DRIVER")
}

// Groups method
func (u User) Groups() []Group {
	elements := u.XMLData.FindElements("GROUPS/ID")
	groups := make([]Group, len(elements))
	for i, e := range elements {
		id, err := strconv.Atoi(e.Text())
		if err != nil {
			groups[i] = *CreateGroup(id)
		}
	}
	return groups
}
