package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

// Object ...
type Object struct {
	vao       VAO
	count     int32
	primitive uint32
	dim       int
}

// Create builds a VAO storing 3D geometry
func (o *Object) Create(vertices, normals, colors []float32, indices []uint32, primitive uint32) {
	// Create the Array Object based on the mesh
	o.vao = makeVAO(vertices, normals, colors, indices)
	if indices != nil {
		o.count = int32(len(indices))
	} else {
		o.count = int32(len(vertices) / 3)
	}
	o.primitive = primitive
	o.dim = 3
}

// ReplaceVertexBuffer recreates a vertex buffer
// assumed that the vao has no element indexing
func (o *Object) ReplaceVertexBuffer(newVertices []float32) {
	gl.BindBuffer(gl.ARRAY_BUFFER, o.vao.VertexBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, len(newVertices)*4, gl.Ptr(newVertices), gl.STATIC_DRAW)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	o.count = int32(len(newVertices) / o.dim)
}

// ReplaceColorBuffer recreates a color buffer
// assumed that the vao has no element indexing
func (o *Object) ReplaceColorBuffer(newColors []float32) {
	gl.BindBuffer(gl.ARRAY_BUFFER, o.vao.ColorsBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, len(newColors)*4, gl.Ptr(newColors), gl.STATIC_DRAW)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
}

// UpdateColorBuffer updates the content of the colors vbo
func (o *Object) UpdateColorBuffer(newColors []float32) {
	gl.BindBuffer(gl.ARRAY_BUFFER, o.vao.ColorsBuffer)
	gl.BufferSubData(gl.ARRAY_BUFFER, 0, len(newColors)*4, gl.Ptr(newColors))
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
}

// Draw renders the object on the screen
func (o Object) Draw() {
	gl.BindVertexArray(o.vao.Ptr)
	if o.vao.hasElements {
		gl.DrawElements(o.primitive, o.count, gl.UNSIGNED_INT, gl.PtrOffset(0))
	} else {
		gl.DrawArrays(o.primitive, 0, o.count)
	}
}
