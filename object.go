package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

// Object ...
type Object struct {
	vao       VAO
	count     int32
	primitive uint32
}

// Create builds a VAO storing 3D geometry
func (o *Object) Create(vertices, normals []float32) {
	o.vao = makeVAO(vertices, normals)
	o.count = int32(len(vertices) / 3)
}

// ReplaceVertexBuffer recreates a vertex buffer
// assumed that the vao has no element indexing
func (o *Object) ReplaceVertexBuffer(newVertices []float32) {
	gl.BindBuffer(gl.ARRAY_BUFFER, o.vao.VertexBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, len(newVertices)*4, gl.Ptr(newVertices), gl.STATIC_DRAW)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	o.count = int32(len(newVertices) / 3)
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
	gl.DrawArrays(gl.TRIANGLES, 0, o.count)
}
