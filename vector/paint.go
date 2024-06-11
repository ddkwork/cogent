// Copyright (c) 2021, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vector

import (
	"fmt"

	"cogentcore.org/core/base/reflectx"
	"cogentcore.org/core/colors"
	"cogentcore.org/core/core"
	"cogentcore.org/core/events"
	"cogentcore.org/core/styles"
	"cogentcore.org/core/styles/units"
	"cogentcore.org/core/svg"
	"cogentcore.org/core/views"
)

// PaintView provides editing of basic Stroke and Fill painting parameters
// for selected items
type PaintView struct {
	core.Frame

	// paint type for stroke
	StrokeType PaintTypes

	// name of gradient with stops
	StrokeStops string

	// paint type for fill
	FillType PaintTypes

	// name of gradient with stops
	FillStops string

	// the parent vectorview
	VectorView *VectorView `copier:"-" json:"-" xml:"-" view:"-"`
}

func (pv *PaintView) Init() {
	pv.Frame.Init()
	pv.StrokeType = PaintSolid
	pv.FillType = PaintSolid

	DashIconsInit()
	MarkerIconsInit()

	pv.Styler(func(s *styles.Style) {
		s.Direction = styles.Column
	})

	sty := &Settings.ShapeStyle

	core.AddChildAt(pv, "stroke-lab", func(w *core.Frame) {
		w.Styler(func(s *styles.Style) {
			s.Direction = styles.Row
		})
		core.AddChild(w, func(w *core.Text) {
			w.SetText("<b>Stroke Paint:  </b>")
		})
		core.AddChild(w, func(w *core.Switches) {
			w.SetEnum(pv.StrokeType)
		})
	})

	core.AddChildAt(pv, "stroke-width", func(w *core.Frame) {
		w.Styler(func(s *styles.Style) {
			s.Direction = styles.Row
		})
		core.AddChild(w, func(w *core.Text) {
			w.SetText("Width:  ").Styler(func(s *styles.Style) {
				s.Align.Items = styles.Center
			})
		})
		core.AddChild(w, func(w *core.Spinner) {
			w.SetMin(0).SetStep(0.05).
				SetValue(sty.StrokeStyle.Width.Value).OnChange(func(e events.Event) {
				if pv.IsStrokeOn() {
					pv.VectorView.SetStrokeWidth(pv.StrokeWidthProp(), false)
				}
			})
		})

		// uncb.SetCurrentIndex(int(Settings.Size.Units))
		core.AddChild(w, func(w *core.Chooser) {
			w.SetEnum(sty.StrokeStyle.Width.Unit).
				OnChange(func(e events.Event) {
					if pv.IsStrokeOn() {
						pv.VectorView.SetStrokeWidth(pv.StrokeWidthProp(), false)
					}
				})
		})

		// core.NewSpace(wr, "sp1").Styler(func(s *styles.Style) {
		// 	s.Min.X.Ch(5)
		// })

		core.AddChild(w, func(w *core.Chooser) {
			// dshcb.ItemsFromIconList(AllDashIcons, true, 0)
			// dshcb.SetProp("width", units.NewCh(15))
			w.OnChange(func(e events.Event) {
				if pv.IsStrokeOn() {
					pv.VectorView.SetDashProperties(pv.StrokeDashProp())
				}
			})
		})

		core.AddChildAt(w, "stroke-markers", func(w *core.Frame) {
			w.Styler(func(s *styles.Style) {
				s.Direction = styles.Row
			})
			core.AddChild(w, func(w *core.Chooser) { // start
				// mscb.SetProp("width", units.NewCh(20))
				// mscb.ItemsFromIconList(AllMarkerIcons, true, 0)
				w.OnChange(func(e events.Event) {
					if pv.IsStrokeOn() {
						pv.VectorView.SetMarkerProperties(pv.MarkerProperties())
					}
				})
			})
			core.AddChild(w, func(w *core.Chooser) { // start-color
				w.SetEnum(MarkerColorsN)
				// mscc.SetProp("width", units.NewCh(5))
				w.OnChange(func(e events.Event) {
					if pv.IsStrokeOn() {
						pv.VectorView.SetMarkerProperties(pv.MarkerProperties())
					}
				})
			})

			core.AddChild(w, func(w *core.Separator) {})

			core.AddChild(w, func(w *core.Chooser) { // mid
				// mscb.SetProp("width", units.NewCh(20))
				// mscb.ItemsFromIconList(AllMarkerIcons, true, 0)
				w.OnChange(func(e events.Event) {
					if pv.IsStrokeOn() {
						pv.VectorView.SetMarkerProperties(pv.MarkerProperties())
					}
				})
			})
			core.AddChild(w, func(w *core.Chooser) { // mid-color
				w.SetEnum(MarkerColorsN)
				// mmcc.SetProp("width", units.NewCh(5))
				w.OnChange(func(e events.Event) {
					if pv.IsStrokeOn() {
						pv.VectorView.SetMarkerProperties(pv.MarkerProperties())
					}
				})
			})

			core.AddChild(w, func(w *core.Separator) {})

			core.AddChild(w, func(w *core.Chooser) { // end
				// mscb.SetProp("width", units.NewCh(20))
				// mscb.ItemsFromIconList(AllMarkerIcons, true, 0)
				w.OnChange(func(e events.Event) {
					if pv.IsStrokeOn() {
						pv.VectorView.SetMarkerProperties(pv.MarkerProperties())
					}
				})
			})
			core.AddChild(w, func(w *core.Chooser) { // end-color
				w.SetEnum(MarkerColorsN)
				// mscc.SetProp("width", units.NewCh(5))
				w.OnChange(func(e events.Event) {
					if pv.IsStrokeOn() {
						pv.VectorView.SetMarkerProperties(pv.MarkerProperties())
					}
				})
			})
		})

		////////////////////////////////
		// stroke stack

		core.AddChildAt(w, "stroke-stack", func(w *core.Frame) {
			w.Styles.Display = styles.Stacked
			w.StackTop = 1
			// ss.StackTopOnly = true

			core.AddChild(w, func(w *core.Frame) {}) // "stroke-blank"

			// spt.ButtonSig.Connect(pv.This, func(recv, send tree.Node, sig int64, data any) {
			// 	pvv := recv.Embed(KiT_PaintView).(*PaintView)
			// 	sss := pvv.StrokeStack()
			// 	prev := pv.StrokeType
			// 	pvv.StrokeType = PaintTypes(sig)
			// 	update := pvv.UpdateStart()
			// 	pvv.SetFullReRender()
			// 	sp := pvv.StrokeProp()
			// 	switch pvv.StrokeType {
			// 	case PaintOff, PaintInherit:
			// 		sss.StackTop = 0
			// 	case PaintSolid:
			// 		sss.StackTop = 1
			// 	case PaintLinear, PaintRadial:
			// 		if pvv.StrokeStops == "" {
			// 			pvv.StrokeStops = pvv.VectorView.DefaultGradient()
			// 		}
			// 		sp = pvv.StrokeStops
			// 		sss.StackTop = 2
			// 		pvv.SelectStrokeGrad()
			// 	}
			// 	pvv.UpdateEnd(update)
			// 	pvv.VectorView.SetStroke(prev, pvv.StrokeType, sp)
			// })

			core.AddChild(w, func(w *views.ColorView) { // "stroke-clr")
				// sc.SetProp("vertical-align", styles.AlignTop)
				// sc.SetColor(sty.StrokeStyle.Color)
				// sc.ViewSig.Connect(pv.This, func(recv, send tree.Node, sig int64, data any) {
				// 	if pv.StrokeType == PaintSolid {
				// 		pv.VectorView.SetStrokeColor(pv.StrokeProp(), false) // not manip
				// 	}
				// })
				// sc.ManipSig.Connect(pv.This, func(recv, send tree.Node, sig int64, data any) {
				// 	if pv.StrokeType == PaintSolid {
				// 		pv.VectorView.SetStrokeColor(pv.StrokeProp(), true) // manip
				// 	}
				// })
			})

			core.AddChild(w, func(w *views.TableView) { // "stroke-grad"
				// sg.SetProp("index", true)
				// sg.SetProp("toolbar", true)
				// sg.SelectedIndex = -1
				w.SetSlice(&pv.VectorView.EditState.Gradients)
				// sg.WidgetSig.Connect(pv.This, func(recv, send tree.Node, sig int64, data any) {
				// 	if sig == int64(core.WidgetSelected) {
				// 		svv, _ := send.(*views.TableView)
				// 		if svv.SelectedIndex >= 0 {
				// 			pv.StrokeStops = pv.VectorView.EditState.Gradients[svv.SelectedIndex].Name
				// 			pv.VectorView.SetStroke(pv.StrokeType, pv.StrokeType, pv.StrokeStops) // handles full updating
				// 		}
				// 	}
				// })
			})
		})

		core.AddChild(w, func(w *core.Separator) {})

		core.AddChildAt(pv, "fill-lab", func(w *core.Frame) {
			w.Styler(func(s *styles.Style) {
				s.Direction = styles.Row
			})
			core.AddChild(w, func(w *core.Text) {
				w.SetText("<b>Fill Paint:  </b>")
			})
			core.AddChild(w, func(w *core.Switches) {
				w.SetEnum(pv.FillType)
			})
		})

		core.AddChildAt(w, "fill-stack", func(w *core.Frame) {
			w.Styles.Display = styles.Stacked
			w.StackTop = 1
			// fs.StackTopOnly = true

			core.AddChild(w, func(w *core.Frame) {}) // "fill-blank"

			// fpt.ButtonSig.Connect(pv.This, func(recv, send tree.Node, sig int64, data any) {
			// 	pvv := recv.Embed(KiT_PaintView).(*PaintView)
			// 	fss := pvv.FillStack()
			// 	prev := pvv.FillType
			// 	pvv.FillType = PaintTypes(sig)
			// 	update := fss.UpdateStart()
			// 	fss.SetFullReRender()
			// 	fp := pvv.FillProp()
			// 	switch pvv.FillType {
			// 	case PaintOff, PaintInherit:
			// 		fss.StackTop = 0
			// 	case PaintSolid:
			// 		fss.StackTop = 1
			// 	case PaintLinear, PaintRadial:
			// 		if pvv.FillStops == "" {
			// 			pvv.FillStops = pvv.VectorView.DefaultGradient()
			// 		}
			// 		fp = pvv.FillStops
			// 		fss.StackTop = 2
			// 		pvv.SelectFillGrad()
			// 	}
			// 	pvv.UpdateEnd(update)
			// 	pvv.VectorView.SetFill(prev, pvv.FillType, fp)
			// })

			core.AddChild(w, func(w *views.ColorView) { // "fill-clr")
				w.SetColor(colors.Scheme.Primary.Base)
				// fc.SetProp("vertical-align", styles.AlignTop)
				// fc.Config()
				// fc.SetColor(sty.FillStyle.Color.Color)
				// fc.ViewSig.Connect(pv.This, func(recv, send tree.Node, sig int64, data any) {
				// 	if pv.FillType == PaintSolid {
				// 		pv.VectorView.SetFillColor(pv.FillProp(), false)
				// 	}
				// })
				// fc.ManipSig.Connect(pv.This, func(recv, send tree.Node, sig int64, data any) {
				// 	if pv.FillType == PaintSolid {
				// 		pv.VectorView.SetFillColor(pv.FillProp(), true) // manip
				// 	}
				// })
			})

			core.AddChild(w, func(w *views.TableView) { // "fill-grad"
				// sg.SetProp("index", true)
				// sg.SetProp("toolbar", true)
				// sg.SelectedIndex = -1
				w.SetSlice(&pv.VectorView.EditState.Gradients)
				// fg.WidgetSig.Connect(pv.This, func(recv, send tree.Node, sig int64, data any) {
				// 	if sig == int64(core.WidgetSelected) {
				// 		svv, _ := send.(*views.TableView)
				// 		if svv.SelectedIndex >= 0 {
				// 			pv.FillStops = pv.VectorView.EditState.Gradients[svv.SelectedIndex].Name
				// 			pv.VectorView.SetFill(pv.FillType, pv.FillType, pv.FillStops) // this handles updating gradients etc to use stops
				// 		}
				// 	}
				// })
				// fg.SliceViewSig.Connect(pv.This, func(recv, send tree.Node, sig int64, data any) {
				// 	// fmt.Printf("svs: %v   %v\n", sig, data)
				// 	// svv, _ := send.(*views.TableView)
				// 	if sig == int64(views.SliceViewDeleted) { // not clear what we can do here
				// 	} else {
				// 		pv.VectorView.UpdateGradients()
				// 	}
				// })
				// fg.ViewSig.Connect(pv.This, func(recv, send tree.Node, sig int64, data any) {
				// 	// fmt.Printf("vs: %v   %v\n", sig, data)
				// 	// svv, _ := send.(*views.TableView)
				// 	pv.VectorView.UpdateGradients()
				// })
			})

			core.AddChild(w, func(w *core.Stretch) {})
		})
	})
}

/////////////////////////////////////////////////////////////////////////
//  Actions

// ManipAction manages all the updating etc associated with performing an action
// that includes an ongoing manipulation with a final non-manip update.
// runs given function to actually do the update.
func (vv *VectorView) ManipAction(act, data string, manip bool, fun func(sii svg.Node)) {
	es := &vv.EditState
	sv := vv.SVG()
	// update := false
	actStart := false
	finalAct := false
	if !manip && es.InAction() {
		finalAct = true
	}
	if manip && !es.InAction() {
		manip = false
		actStart = true
		es.ActStart(act, data)
		es.ActUnlock()
	}
	if !manip {
		if !finalAct {
			sv.UndoSave(act, data)
		}
		// update = sv.UpdateStart()
	}
	for itm := range es.Selected {
		vv.ManipActionFun(itm, fun)
	}
	if !manip {
		// sv.UpdateEnd(update)
		if !actStart {
			es.ActDone()
			vv.ChangeMade()
		}
	} else {
		sv.ManipUpdate()
	}
}

func (vv *VectorView) ManipActionFun(sii svg.Node, fun func(itm svg.Node)) {
	if gp, isgp := sii.(*svg.Group); isgp {
		for _, kid := range gp.Children {
			vv.ManipActionFun(kid.(svg.Node), fun)
		}
		return
	}
	fun(sii)
}

// SetColorNode sets the color properties of Node
// based on previous and current PaintType
func (vv *VectorView) SetColorNode(sii svg.Node, prop string, prev, pt PaintTypes, sp string) {
	if gp, isgp := sii.(*svg.Group); isgp {
		for _, kid := range gp.Children {
			vv.SetColorNode(kid.(svg.Node), prop, prev, pt, sp)
		}
		return
	}
	switch pt {
	// case PaintLinear:
	// 	svg.UpdateNodeGradientProp(sii, prop, false, sp)
	// case PaintRadial:
	// 	svg.UpdateNodeGradientProp(sii, prop, true, sp)
	default:
		if prev == PaintLinear || prev == PaintRadial {
			pstr := reflectx.ToString(sii.AsTree().Properties[prop])
			_ = pstr
			// svg.DeleteNodeGradient(sii, pstr)
		}
		sii.AsNodeBase().SetColorProperties(prop, sp)
	}
	vv.UpdateMarkerColors(sii)
}

// SetStroke sets the stroke properties of selected items
// based on previous and current PaintType
func (vv *VectorView) SetStroke(prev, pt PaintTypes, sp string) {
	es := &vv.EditState
	sv := vv.SVG()
	sv.UndoSave("SetStroke", sp)
	// update := sv.UpdateStart()
	for itm := range es.Selected {
		vv.SetColorNode(itm, "stroke", prev, pt, sp)
	}
	// sv.UpdateEnd(update)
	vv.ChangeMade()
}

// SetStrokeWidthNode sets the stroke width of Node
func (vv *VectorView) SetStrokeWidthNode(sii svg.Node, wp string) {
	if gp, isgp := sii.(*svg.Group); isgp {
		for _, kid := range gp.Children {
			vv.SetStrokeWidthNode(kid.(svg.Node), wp)
		}
		return
	}
	g := sii.AsNodeBase()
	if g.Paint.StrokeStyle.Color != nil {
		g.SetProperty("stroke-width", wp)
	}
}

// SetStrokeWidth sets the stroke width property for selected items
// manip means currently being manipulated -- don't save undo.
func (vv *VectorView) SetStrokeWidth(wp string, manip bool) {
	es := &vv.EditState
	sv := vv.SVG()
	// update := false
	if !manip {
		sv.UndoSave("SetStrokeWidth", wp)
		// update = sv.UpdateStart()
		// sv.SetFullReRender()
	}
	for itm := range es.Selected {
		vv.SetStrokeWidthNode(itm, wp)
	}
	if !manip {
		// sv.UpdateEnd(update)
		vv.ChangeMade()
	} else {
		sv.ManipUpdate()
	}
}

// SetStrokeColor sets the stroke color for selected items.
// manip means currently being manipulated -- don't save undo.
func (vv *VectorView) SetStrokeColor(sp string, manip bool) {
	vv.ManipAction("SetStrokeColor", sp, manip,
		func(itm svg.Node) {
			p := itm.AsTree().Properties["stroke"]
			if p != nil {
				itm.AsNodeBase().SetColorProperties("stroke", sp)
				vv.UpdateMarkerColors(itm)
			}
		})
}

// SetMarkerNode sets the marker properties of Node.
func (vv *VectorView) SetMarkerNode(sii svg.Node, start, mid, end string, sc, mc, ec MarkerColors) {
	if gp, isgp := sii.(*svg.Group); isgp {
		for _, kid := range gp.Children {
			vv.SetMarkerNode(kid.(svg.Node), start, mid, end, sc, mc, ec)
		}
		return
	}
	sv := vv.SVG()
	MarkerSetProp(sv.SSVG(), sii, "marker-start", start, sc)
	MarkerSetProp(sv.SSVG(), sii, "marker-mid", mid, mc)
	MarkerSetProp(sv.SSVG(), sii, "marker-end", end, ec)
}

// SetMarkerProperties sets the marker properties
func (vv *VectorView) SetMarkerProperties(start, mid, end string, sc, mc, ec MarkerColors) {
	es := &vv.EditState
	sv := vv.SVG()
	sv.UndoSave("SetMarkerProperties", start+" "+mid+" "+end)
	// update := sv.UpdateStart()
	// sv.SetFullReRender()
	for itm := range es.Selected {
		vv.SetMarkerNode(itm, start, mid, end, sc, mc, ec)
	}
	// sv.UpdateEnd(update)
	vv.ChangeMade()
}

// UpdateMarkerColors updates the marker colors, when setting fill or stroke
func (vv *VectorView) UpdateMarkerColors(sii svg.Node) {
	if sii == nil {
		return
	}
	sv := vv.SVG()
	MarkerUpdateColorProp(sv.SSVG(), sii, "marker-start")
	MarkerUpdateColorProp(sv.SSVG(), sii, "marker-mid")
	MarkerUpdateColorProp(sv.SSVG(), sii, "marker-end")
}

// SetDashNode sets the stroke-dasharray property of Node.
// multiplies dash values by the line width in dots.
func (vv *VectorView) SetDashNode(sii svg.Node, dary []float64) {
	if gp, isgp := sii.(*svg.Group); isgp {
		for _, kid := range gp.Children {
			vv.SetDashNode(kid.(svg.Node), dary)
		}
		return
	}
	if len(dary) == 0 {
		delete(sii.AsTree().Properties, "stroke-dasharray")
		return
	}
	g := sii.AsNodeBase()
	mary := DashMulWidth(float64(g.Paint.StrokeStyle.Width.Dots), dary)
	ds := DashString(mary)
	sii.AsTree().Properties["stroke-dasharray"] = ds
}

// SetDashProperties sets the dash properties
func (vv *VectorView) SetDashProperties(dary []float64) {
	es := &vv.EditState
	sv := vv.SVG()
	sv.UndoSave("SetDashProperties", "")
	// update := sv.UpdateStart()
	// sv.SetFullReRender()
	for itm := range es.Selected {
		vv.SetDashNode(itm, dary)
	}
	// sv.UpdateEnd(update)
	vv.ChangeMade()
}

// SetFill sets the fill properties of selected items
// based on previous and current PaintType
func (vv *VectorView) SetFill(prev, pt PaintTypes, fp string) {
	es := &vv.EditState
	sv := vv.SVG()
	sv.UndoSave("SetFill", fp)
	// update := sv.UpdateStart()
	// sv.SetFullReRender()
	for itm := range es.Selected {
		vv.SetColorNode(itm, "fill", prev, pt, fp)
	}
	// sv.UpdateEnd(update)
	vv.ChangeMade()
}

// SetFillColor sets the fill color for selected items
// manip means currently being manipulated -- don't save undo.
func (vv *VectorView) SetFillColor(fp string, manip bool) {
	vv.ManipAction("SetFillColor", fp, manip,
		func(itm svg.Node) {
			p := itm.AsTree().Properties["fill"]
			if p != nil {
				itm.AsNodeBase().SetColorProperties("fill", fp)
				vv.UpdateMarkerColors(itm)
			}
		})
}

// DefaultGradient returns the default gradient to use for setting stops
func (vv *VectorView) DefaultGradient() string {
	es := &vv.EditState
	sv := vv.SVG()
	if len(vv.EditState.Gradients) == 0 {
		es.ConfigDefaultGradient()
		sv.UpdateGradients(es.Gradients)
	}
	return es.Gradients[0].Name
}

// UpdateGradients updates gradients from EditState
func (vv *VectorView) UpdateGradients() {
	es := &vv.EditState
	sv := vv.SVG()
	// update := sv.UpdateStart()
	sv.UpdateGradients(es.Gradients)
	// sv.UpdateEnd(update)
}

///////////////////////////////////////////////////////////////
//  PaintView

// Update updates the current settings based on the values in the given Paint and
// properties from node (node can be nil)
/*
func (pv *PaintView) Update(pc *paint.Paint, kn tree.Node) {
	update := pv.UpdateStart()
	defer pv.UpdateEnd(update)

	pv.StrokeType, pv.StrokeStops = pv.DecodeType(kn, &pc.StrokeStyle.Color, "stroke")
	pv.FillType, pv.FillStops = pv.DecodeType(kn, &pc.FillStyle.Color, "fill")

	es := &pv.VectorView.EditState
	grl := &es.Gradients

	spt := pv.ChildByName("stroke-lab", 0).ChildByName("stroke-type", 1).(*core.ButtonBox)
	spt.SelectItem(int(pv.StrokeType))

	ss := pv.StrokeStack()

	switch pv.StrokeType {
	case PaintSolid:
		if ss.StackTop != 1 {
			ss.SetFullReRender()
		}
		ss.StackTop = 1
		sc := ss.ChildByName("stroke-clr", 1).(*views.ColorView)
		sc.SetColor(pc.StrokeStyle.Color.Color)
	case PaintLinear, PaintRadial:
		if ss.StackTop != 2 {
			ss.SetFullReRender()
		}
		ss.StackTop = 2
		sg := ss.ChildByName("stroke-grad", 1).(*views.TableView)
		sg.SetSlice(grl)
		pv.SelectStrokeGrad()
	default:
		if ss.StackTop != 0 {
			ss.SetFullReRender()
		}
		ss.StackTop = 0
	}

	wr := pv.ChildByName("stroke-width", 2)
	wsb := wr.ChildByName("width", 1).(*core.Spinner)
	wsb.SetValue(pc.StrokeStyle.Width.Val)
	uncb := wr.ChildByName("width-units", 2).(*core.Chooser)
	uncb.SetCurrentIndex(int(pc.StrokeStyle.Width.Un))

	dshcb := wr.ChildByName("dashes", 3).(*core.Chooser)
	nwdsh, dnm := DashMatchArray(float64(pc.StrokeStyle.Width.Dots), pc.StrokeStyle.Dashes)
	if nwdsh {
		dshcb.ItemsFromIconList(AllDashIcons, false, 0)
	}
	dshcb.SetCurVal(icons.Icon(dnm))

	mkr := pv.ChildByName("stroke-markers", 3)

	ms, _, mc := MarkerFromNodeProp(kn, "marker-start")
	mscb := mkr.ChildByName("marker-start", 0).(*core.Chooser)
	mscc := mkr.ChildByName("marker-start-color", 1).(*core.Chooser)
	if ms != "" {
		mscb.SetCurVal(MarkerNameToIcon(ms))
		mscc.SetCurrentIndex(int(mc))
	} else {
		mscb.SetCurrentIndex(0)
		mscc.SetCurrentIndex(0)
	}
	ms, _, mc = MarkerFromNodeProp(kn, "marker-mid")
	mmcb := mkr.ChildByName("marker-mid", 2).(*core.Chooser)
	mmcc := mkr.ChildByName("marker-mid-color", 3).(*core.Chooser)
	if ms != "" {
		mmcb.SetCurVal(MarkerNameToIcon(ms))
		mmcc.SetCurrentIndex(int(mc))
	} else {
		mmcb.SetCurrentIndex(0)
		mmcc.SetCurrentIndex(0)
	}
	ms, _, mc = MarkerFromNodeProp(kn, "marker-end")
	mecb := mkr.ChildByName("marker-end", 4).(*core.Chooser)
	mecc := mkr.ChildByName("marker-end-color", 5).(*core.Chooser)
	if ms != "" {
		mecb.SetCurVal(MarkerNameToIcon(ms))
		mecc.SetCurrentIndex(int(mc))
	} else {
		mecb.SetCurrentIndex(0)
		mecc.SetCurrentIndex(0)
	}

	fpt := pv.ChildByName("fill-lab", 0).ChildByName("fill-type", 1).(*core.ButtonBox)
	fpt.SelectItem(int(pv.FillType))

	fs := pv.FillStack()

	switch pv.FillType {
	case PaintSolid:
		if fs.StackTop != 1 {
			fs.SetFullReRender()
		}
		fs.StackTop = 1
		fc := fs.ChildByName("fill-clr", 1).(*views.ColorView)
		fc.SetColor(pc.FillStyle.Color.Color)
	case PaintLinear, PaintRadial:
		if fs.StackTop != 2 {
			fs.SetFullReRender()
		}
		fs.StackTop = 2
		fg := fs.ChildByName("fill-grad", 1).(*views.TableView)
		if fg.Slice != grl {
			pv.SetFullReRender()
		}
		fg.SetSlice(grl)
		pv.SelectFillGrad()
	default:
		if fs.StackTop != 0 {
			fs.SetFullReRender()
		}
		fs.StackTop = 0
	}
}
*/

/*
// GradStopsName returns the stopsname for gradient from url
func (pv *PaintView) GradStopsName(gii core.Node2D, url string) string {
	gr := svg.GradientByName(gii, url)
	if gr == nil {
		return ""
	}
	if gr.StopsName != "" {
		return gr.StopsName
	}
	return gr.Nm
}
*/

/*
// DecodeType decodes the paint type from paint and properties
// also returns the name of the gradient if using one.
func (pv *PaintView) DecodeType(kn tree.Node, cs *style.ColorSpec, prop string) (PaintTypes, string) {
	pstr := ""
	var gii core.Node2D
	if kn != nil {
		pstr = reflectx.ToString(kn.Prop(prop))
		gii = kn.(core.Node2D)
	}
	ptyp := PaintSolid
	grnm := ""
	switch {
	case pstr == "inherit":
		ptyp = PaintInherit
	case pstr == "none" || cs.IsNil():
		ptyp = PaintOff
	case strings.HasPrefix(pstr, "url(#linear") || (cs.Gradient != nil && !cs.Gradient.IsRadial):
		ptyp = PaintLinear
		if gii != nil {
			grnm = pv.GradStopsName(gii, pstr)
		}
	case strings.HasPrefix(pstr, "url(#radial") || (cs.Gradient != nil && cs.Gradient.IsRadial):
		ptyp = PaintRadial
		if gii != nil {
			grnm = pv.GradStopsName(gii, pstr)
		}
	default:
		ptyp = PaintSolid
	}
	if grnm == "" {
		if prop == "fill" {
			grnm = pv.FillStops
		} else {
			grnm = pv.StrokeStops
		}
	}
	return ptyp, grnm
}
*/

func (pv *PaintView) SelectStrokeGrad() {
	es := &pv.VectorView.EditState
	grl := &es.Gradients
	ss := pv.StrokeStack()
	sg := ss.ChildByName("stroke-grad", 1).(*views.TableView)
	sg.UnselectAllIndexes()
	for i, g := range *grl {
		if g.Name == pv.StrokeStops {
			sg.SelectIndex(i)
			break
		}
	}
}

func (pv *PaintView) SelectFillGrad() {
	es := &pv.VectorView.EditState
	grl := &es.Gradients
	fs := pv.FillStack()
	fg := fs.ChildByName("fill-grad", 1).(*views.TableView)
	fg.UnselectAllIndexes()
	for i, g := range *grl {
		if g.Name == pv.FillStops {
			fg.SelectIndex(i)
			break
		}
	}
}

// StrokeStack returns the stroke stack frame
func (pv *PaintView) StrokeStack() *core.Frame {
	return pv.ChildByName("stroke-stack", 1).(*core.Frame)
}

// FillStack returns the fill stack frame
func (pv *PaintView) FillStack() *core.Frame {
	return pv.ChildByName("fill-stack", 4).(*core.Frame)
}

// StrokeProp returns the stroke property string according to current settings
func (pv *PaintView) StrokeProp() string {
	// ss := pv.StrokeStack()
	switch pv.StrokeType {
	case PaintOff:
		return "none"
	case PaintSolid:
		// sc := ss.ChildByName("stroke-clr", 1).(*views.ColorView)
		// return sc.Color.HexString()
	case PaintLinear:
		return pv.StrokeStops
	case PaintRadial:
		return pv.StrokeStops
	case PaintInherit:
		return "inherit"
	}
	return "none"
}

// MarkerProp returns the marker property string according to current settings
// along with color type to set.
func (pv *PaintView) MarkerProperties() (start, mid, end string, sc, mc, ec MarkerColors) {
	// mkr := pv.ChildByName("stroke-markers", 3)
	//
	// mscb := mkr.ChildByName("marker-start", 0).(*core.Chooser)
	// mscc := mkr.ChildByName("marker-start-color", 1).(*core.Chooser)
	// start = IconToMarkerName(mscb.CurVal)
	// sc = MarkerColors(mscc.CurrentIndex)
	//
	// mmcb := mkr.ChildByName("marker-mid", 2).(*core.Chooser)
	// mmcc := mkr.ChildByName("marker-mid-color", 3).(*core.Chooser)
	// mid = IconToMarkerName(mmcb.CurVal)
	// mc = MarkerColors(mmcc.CurrentIndex)
	//
	// mecb := mkr.ChildByName("marker-end", 4).(*core.Chooser)
	// mecc := mkr.ChildByName("marker-end-color", 5).(*core.Chooser)
	// end = IconToMarkerName(mecb.CurVal)
	// ec = MarkerColors(mecc.CurrentIndex)

	return
}

// IsStrokeOn returns true if stroke is active
func (pv *PaintView) IsStrokeOn() bool {
	return pv.StrokeType >= PaintSolid && pv.StrokeType < PaintInherit
}

// StrokeWidthProp returns stroke-width property
func (pv *PaintView) StrokeWidthProp() string {
	wr := pv.ChildByName("stroke-width", 2)
	wsb := wr.AsTree().ChildByName("width", 1).(*core.Spinner)
	uncb := wr.AsTree().ChildByName("width-units", 2).(*core.Chooser)
	unnm := "px"
	if uncb.CurrentIndex > 0 {
		unvl := units.Units(uncb.CurrentIndex)
		unnm = unvl.String()
	}
	return fmt.Sprintf("%g%s", wsb.Value, unnm)
}

// StrokeDashProp returns stroke-dasharray property as an array (nil = none)
// these values need to be multiplied by line widths for each item.
func (pv *PaintView) StrokeDashProp() []float64 {
	wr := pv.ChildByName("stroke-width", 2)
	dshcb := wr.AsTree().ChildByName("dashes", 3).(*core.Chooser)
	if dshcb.CurrentIndex == 0 {
		return nil
	}
	dnm := reflectx.ToString(dshcb.CurrentItem.Value)
	if dnm == "" {
		return nil
	}
	dary, ok := AllDashesMap[dnm]
	if !ok {
		return nil
	}
	return dary
}

// IsFillOn returns true if Fill is active
func (pv *PaintView) IsFillOn() bool {
	return pv.FillType >= PaintSolid && pv.FillType < PaintInherit
}

// FillProp returns the fill property string according to current settings
func (pv *PaintView) FillProp() string {
	fs := pv.FillStack()
	switch pv.FillType {
	case PaintOff:
		return "none"
	case PaintSolid:
		sc := fs.ChildByName("fill-clr", 1).(*views.ColorView)
		return colors.AsHex(sc.Color)
	case PaintLinear:
		return pv.FillStops
	case PaintRadial:
		return pv.FillStops
	case PaintInherit:
		return "inherit"
	}
	return "none"
}

// SetProperties sets the properties for given node according to current settings
func (pv *PaintView) SetProperties(sii svg.Node) {
	pv.VectorView.SetColorNode(sii, "stroke", pv.StrokeType, pv.StrokeType, pv.StrokeProp())
	if pv.IsStrokeOn() {
		sii.AsTree().Properties["stroke-width"] = pv.StrokeWidthProp()
		start, mid, end, sc, mc, ec := pv.MarkerProperties()
		pv.VectorView.SetMarkerNode(sii, start, mid, end, sc, mc, ec)
	}
	pv.VectorView.SetColorNode(sii, "fill", pv.FillType, pv.FillType, pv.FillProp())
}

type PaintTypes int32 //enums:enum -trim-prefix Paint

const (
	PaintOff PaintTypes = iota
	PaintSolid
	PaintLinear
	PaintRadial
	PaintInherit
)
