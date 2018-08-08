package resources

import (
	"fmt"
	"github.com/beevik/etree"
	"strconv"
)

// Group struct
type Group struct {
	XMLData *etree.Element
}

// CreateGroup constructs Group
func CreateGroup(id int) *Group {
	doc := etree.NewDocument()
	doc.CreateProcInst("xml", `version="1.0"`)

	el := doc.CreateElement("GROUP").CreateElement("ID")
	el.SetText(fmt.Sprintf("%d", id))

	return &Group{doc.Root()}
}

// Attribute method
func (g Group) Attribute(path string) string {
	elements := g.XMLData.FindElements(path)
	if elements == nil {
		return ""
	}
	return elements[0].Text()
}

// ID method
func (g Group) ID() int {
	i, err := strconv.Atoi(g.Attribute("ID"))
	if err != nil {
		return -1
	}
	return i
}

// Name method
func (g Group) Name() string {
	return g.Attribute("NAME")
}

// Users method
func (g Group) Users() []User {
	elements := g.XMLData.FindElements("USERS/ID")
	users := make([]User, len(elements))
	for i, e := range elements {
		id, err := strconv.Atoi(e.Text())
		if err != nil {
			users[i] = *CreateUser(id)
		}
	}
	return users
}
