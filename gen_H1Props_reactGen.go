// Code generated by reactGen. DO NOT EDIT.

package react

// H1Props defines the properties for the <h1> element
type H1Props struct {
	ID                      string
	Key                     string
	ClassName               string
	Role                    string
	Style                   *CSS
	OnChange                func(e *SyntheticEvent)
	OnClick                 func(e *SyntheticMouseEvent)
	DangerouslySetInnerHTML *DangerousInnerHTMLDef
}

func (h *H1Props) assign(v *_H1Props) {

	if h.ID != "" {
		v.ID = h.ID
	}

	if h.Key != "" {
		v.Key = h.Key
	}

	v.ClassName = h.ClassName

	v.Role = h.Role

	// TODO: until we have a resolution on
	// https://github.com/gopherjs/gopherjs/issues/236
	v.Style = h.Style.hack()

	v.OnChange = h.OnChange

	v.OnClick = h.OnClick

	v.DangerouslySetInnerHTML = h.DangerouslySetInnerHTML

}
