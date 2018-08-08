package resources

import (
	"fmt"
	"github.com/beevik/etree"
	"strconv"
)

// Host struct
type Host struct {
	XMLData *etree.Element
}

// CreateHost constructs Host
func CreateHost(id int) *Host {
	doc := etree.NewDocument()
	doc.CreateProcInst("xml", `version="1.0"`)

	el := doc.CreateElement("HOST").CreateElement("ID")
	el.SetText(fmt.Sprintf("%d", id))

	return &Host{doc.Root()}
}

// Attribute method
func (h Host) Attribute(path string) string {
	elements := h.XMLData.FindElements(path)
	if elements == nil {
		return ""
	}
	return elements[0].Text()
}

// ID method
func (h Host) ID() int {
	i, err := strconv.Atoi(h.Attribute("ID"))
	if err != nil {
		return -1
	}
	return i
}

// Name method
func (h Host) Name() string {
	return h.Attribute("NAME")
}

// ImMad method
func (h Host) ImMad() string {
	return h.Attribute("IM_MAD")
}

// VMMad method
func (h Host) VMMad() string {
	return h.Attribute("VM_MAD")
}

// Cluster method
func (h Host) Cluster() []Cluster {
	elements := h.XMLData.FindElements("CLUSTERS")
	clusters := make([]Cluster, len(elements))
	for i, e := range elements {
		clusters[i] = Cluster{XMLData: e}
	}
	return clusters
}
