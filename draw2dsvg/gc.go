// Copyright 2015 The draw2d Authors. All rights reserved.
// created: 16/12/2017 by Drahoslav Bednář

package draw2dsvg

import (
	"bytes"
	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dbase"
	"image"
	"strings"
)

const ()

var ()

type drawType int

const (
	filled drawType = 1 << iota
	stroked
)

type SVG bytes.Buffer

func NewSvg() *Svg {
	return &Svg{
		Xmlns:      "http://www.w3.org/2000/svg",
		FillStroke: FillStroke{Fill: "none", Stroke: "none"},
	}
}

// GraphicContext implements the draw2d.GraphicContext interface
// It provides draw2d with a svg backend
type GraphicContext struct {
	*draw2dbase.StackGraphicContext
	svg *Svg
}

func NewGraphicContext(svg *Svg) *GraphicContext {
	gc := &GraphicContext{draw2dbase.NewStackGraphicContext(), svg}
	return gc
}

// Clear fills the current canvas with a default transparent color
func (gc *GraphicContext) Clear() {
	gc.svg.Groups = nil
	gc.svg.Groups = append(gc.svg.Groups, Group{
	// TODO add background color?
	})
}

// Stroke strokes the paths with the color specified by SetStrokeColor
func (gc *GraphicContext) Stroke(paths ...*draw2d.Path) {
	gc.drawPaths(stroked, paths...)
	gc.Current.Path.Clear()
}

// Fill fills the paths with the color specified by SetFillColor
func (gc *GraphicContext) Fill(paths ...*draw2d.Path) {
	gc.drawPaths(filled, paths...)
	gc.Current.Path.Clear()
}

// FillStroke first fills the paths and than strokes them
func (gc *GraphicContext) FillStroke(paths ...*draw2d.Path) {
	gc.drawPaths(filled|stroked, paths...)
	gc.Current.Path.Clear()
}

func (gc *GraphicContext) drawPaths(drawType drawType, paths ...*draw2d.Path) {
	paths = append(paths, gc.Current.Path)

	svgPath := Path{}
	group := Group{}

	svgPathsDesc := make([]string, len(paths))

	// multiple pathes has to be joined to single svg path description
	// because fill-rule wont work for whole group
	for i, path := range paths {
		svgPathsDesc[i] = toSvgPathDesc(path)
	}
	svgPath.Desc = strings.Join(svgPathsDesc, " ")

	if drawType&stroked == stroked {
		group.Stroke = toSvgRGBA(gc.Current.StrokeColor)
		group.StrokeWidth = toSvgLength(gc.Current.LineWidth)
		group.StrokeLinecap = gc.Current.Cap.String()
		group.StrokeLinejoin = gc.Current.Join.String()
		if len(gc.Current.Dash) > 0 {
			group.StrokeDasharray = toSvgArray(gc.Current.Dash)
			group.StrokeDashoffset = toSvgLength(gc.Current.DashOffset)
		}
	}

	if drawType&filled == filled {
		group.Fill = toSvgRGBA(gc.Current.FillColor)
		group.FillRule = toSvgFillRule(gc.Current.FillRule)
	}

	group.Paths = []Path{svgPath}

	gc.svg.Groups = append(gc.svg.Groups, group)
}

///////////////////////////////////////
// TODO implement following methods (or remove if not neccesary)

// SetFontData sets the current FontData
func (gc *GraphicContext) SetFontData(fontData draw2d.FontData) {

}

// GetFontData gets the current FontData
func (gc *GraphicContext) GetFontData() draw2d.FontData {
	return draw2d.FontData{}
}

// GetFontName gets the current FontData as a string
func (gc *GraphicContext) GetFontName() string {
	return ""
}

// DrawImage draws the raster image in the current canvas
func (gc *GraphicContext) DrawImage(image image.Image) {

}

// Save the context and push it to the context stack
func (gc *GraphicContext) Save() {

}

// Restore remove the current context and restore the last one
func (gc *GraphicContext) Restore() {

}

// ClearRect fills the specified rectangle with a default transparent color
func (gc *GraphicContext) ClearRect(x1, y1, x2, y2 int) {

}

// SetDPI sets the current DPI
func (gc *GraphicContext) SetDPI(dpi int) {

}

// GetDPI gets the current DPI
func (gc *GraphicContext) GetDPI() int {
	return 0
}

// GetStringBounds gets pixel bounds(dimensions) of given string
func (gc *GraphicContext) GetStringBounds(s string) (left, top, right, bottom float64) {
	return 0, 0, 0, 0
}

// CreateStringPath creates a path from the string s at x, y
func (gc *GraphicContext) CreateStringPath(text string, x, y float64) (cursor float64) {
	return 0
}

// FillString draws the text at point (0, 0)
func (gc *GraphicContext) FillString(text string) (cursor float64) {
	return 0
}

// FillStringAt draws the text at the specified point (x, y)
func (gc *GraphicContext) FillStringAt(text string, x, y float64) (cursor float64) {
	return 0
}

// StrokeString draws the contour of the text at point (0, 0)
func (gc *GraphicContext) StrokeString(text string) (cursor float64) {
	return 0
}

// StrokeStringAt draws the contour of the text at point (x, y)
func (gc *GraphicContext) StrokeStringAt(text string, x, y float64) (cursor float64) {
	return 0
}
