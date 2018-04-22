// Template generated by reactGen

package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"

	"mvdan.cc/sh/syntax"

	"honnef.co/go/js/dom"
	"myitcv.io/react"
)

type AppDef struct {
	react.ComponentDef
}

//go:generate immutableGen

type lang int

const (
	langGo    lang = iota // Go
	langShell             // Shell
)

type _Imm_langState struct {
	Code  string
	Ast   string
	Error bool
}

type AppState struct {
	Go     *langState
	Shell  *langState
	Choice lang
}

func (a AppState) currLangState() *langState {
	switch a.Choice {
	case langGo:
		return a.Go
	case langShell:
		return a.Shell
	default:
		panic(fmt.Errorf("unable to handle language %v", a.Choice))
	}
}

func (a AppState) setCurrLangState(ls *langState) AppState {
	switch a.Choice {
	case langGo:
		a.Go = ls
	case langShell:
		a.Shell = ls
	default:
		panic(fmt.Errorf("unable to handle language %v", a.Choice))
	}

	return a
}

func App() *AppElem {
	return buildAppElem()
}

func (a AppDef) GetInitialState() AppState {
	return AppState{
		Go:     new(langState),
		Shell:  new(langState),
		Choice: langGo,
	}
}

func (a AppDef) Render() react.Element {
	s := a.State()
	curr := s.currLangState()

	outputClass := "ast"
	if curr.Error() {
		outputClass += " asterror"
	}

	buildLi := func(l lang) *react.LiElem {
		return react.Li(nil,
			react.A(
				&react.AProps{
					Href:    "#",
					OnClick: languageChange(a, l),
				},
				react.S(l.String()),
			),
		)
	}

	return react.Div(
		&react.DivProps{ClassName: "grid-container"},
		react.Div(
			&react.DivProps{ClassName: "header"},
			react.S("Syntax Viewer"),
			react.Div(
				&react.DivProps{ClassName: "dropdown", Style: &react.CSS{Float: "right"}},
				react.Button(
					&react.ButtonProps{
						ClassName:    "btn btn-default dropdown-toggle",
						Type:         "button",
						ID:           "dropdownMenu1",
						DataSet:      react.DataSet{"toggle": "dropdown"},
						AriaHasPopup: true,
						AriaExpanded: true,
					},
					react.Sprintf("%v ", s.Choice),
					react.Span(&react.SpanProps{ClassName: "caret"}),
				),
				react.Ul(
					&react.UlProps{
						ClassName:      "dropdown-menu dropdown-menu-right",
						AriaLabelledBy: "dropdownMenu1",
					},
					buildLi(langGo),
					buildLi(langShell),
				),
			),
		),
		react.Div(
			&react.DivProps{ClassName: "left"},
			react.TextArea(
				&react.TextAreaProps{
					ClassName:   "codeinput",
					Placeholder: "Your code here...",
					Value:       curr.Code(),
					OnChange:    inputChange(a),
				},
			),
		),
		react.Div(
			&react.DivProps{ClassName: "right"},
			react.Pre(
				&react.PreProps{ClassName: outputClass},
				react.S(curr.Ast()),
			),
		),
	)
}

func (a AppDef) handleEvent() {
	st := a.State().currLangState().AsMutable()
	defer func() {
		st.AsImmutable(nil)
		a.SetState(a.State().setCurrLangState(st))
	}()

	st.SetError(true)
	st.SetAst("")

	if st.Code() == "" {
		return
	}

	b := new(bytes.Buffer)

	switch a.State().Choice {
	case langGo:
		fset := token.NewFileSet()
		f, err := parser.ParseFile(fset, "", st.Code(), parser.ParseComments)
		if err != nil {
			st.SetAst(err.Error())
			return
		}

		if err := ast.Fprint(b, fset, f, nil); err != nil {
			st.SetAst(err.Error())
			return
		}

	case langShell:
		in := strings.NewReader(st.Code())
		f, err := syntax.NewParser().Parse(in, "stdin")
		if err != nil {
			st.SetAst(err.Error())
			return
		}

		if err := syntax.DebugPrint(b, f); err != nil {
			st.SetAst(err.Error())
			return
		}

	default:
		panic(fmt.Errorf("don't know how to handleEvent for %v", a.State().Choice))
	}

	st.SetAst(b.String())
	st.SetError(false)
}

type changeEvent struct {
	a AppDef
}

func languageChange(a AppDef, l lang) languageChangeEvent {
	return languageChangeEvent{
		changeEvent: changeEvent{
			a: a,
		},
		l: l,
	}
}

type languageChangeEvent struct {
	changeEvent
	l lang
}

func (l languageChangeEvent) OnClick(se *react.SyntheticMouseEvent) {
	se.PreventDefault()

	st := l.a.State()
	st.Choice = l.l
	l.a.SetState(st)

	l.a.handleEvent()
}

func inputChange(a AppDef) inputChangeEvent {
	return inputChangeEvent{
		changeEvent: changeEvent{a: a},
	}
}

type inputChangeEvent struct {
	changeEvent
}

func (i inputChangeEvent) OnChange(se *react.SyntheticEvent) {
	target := se.Target().(*dom.HTMLTextAreaElement)

	st := i.a.State()
	st = st.setCurrLangState(st.currLangState().SetCode(target.Value))
	i.a.SetState(st)

	i.a.handleEvent()
}