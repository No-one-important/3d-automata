package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

// Object ...
type Object struct {
	VertexCount   int32
	InstanceCount int32

	Ptr uint32

	VertexBuffer   uint32
	NormalsBuffer  uint32
	InstanceBuffer uint32
}

// Create builds a VAO storing 3D geometry
func (o *Object) Create(vertices, normals []float32) {
	o.VertexCount = int32(len(vertices) / 3)

	// Create VAO buffer
	gl.GenVertexArrays(1, &o.Ptr)
	gl.BindVertexArray(o.Ptr)

	// Create Vertices buffer
	if vertices != nil {
		gl.GenBuffers(1, &o.VertexBuffer)
		gl.BindBuffer(gl.ARRAY_BUFFER, o.VertexBuffer)
		gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

		gl.EnableVertexAttribArray(0)
		gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)
	} else {
		gl.DisableVertexAttribArray(0)
	}

	// Create Normals buffer
	if normals != nil {
		gl.GenBuffers(1, &o.NormalsBuffer)
		gl.BindBuffer(gl.ARRAY_BUFFER, o.NormalsBuffer)
		gl.BufferData(gl.ARRAY_BUFFER, len(normals)*4, gl.Ptr(normals), gl.STATIC_DRAW)

		gl.EnableVertexAttribArray(1)
		gl.VertexAttribPointer(1, 3, gl.FLOAT, false, 0, nil)
	} else {
		gl.DisableVertexAttribArray(1)
	}

	// Create Instance buffer
	gl.GenBuffers(1, &o.InstanceBuffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, o.InstanceBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, 0, nil, gl.STATIC_DRAW)

	gl.EnableVertexAttribArray(2)
	gl.VertexAttribPointer(2, 3, gl.FLOAT, false, 0, nil)
	gl.VertexAttribDivisor(2, 1)

	// Disable VAO
	gl.BindVertexArray(0)
}

// ReplaceVertexBuffer recreates a vertex buffer
// assumed that the vao has no element indexing
func (o *Object) ReplaceVertexBuffer(newVertices []float32) {
	gl.BindBuffer(gl.ARRAY_BUFFER, o.VertexBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, len(newVertices)*4, gl.Ptr(newVertices), gl.STATIC_DRAW)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	o.VertexCount = int32(len(newVertices) / 3)
}

// UpdateInstances ...
func (o *Object) UpdateInstances(newInstances []int32) {
	gl.BindBuffer(gl.ARRAY_BUFFER, o.InstanceBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, len(newInstances)*4, gl.Ptr(newInstances), gl.STATIC_DRAW)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	o.InstanceCount = int32(len(newInstances) / 3)
}

// Draw ...
func (o *Object) Draw() {
	gl.BindVertexArray(o.Ptr)
	gl.DrawArraysInstanced(gl.TRIANGLES, 0, o.VertexCount, o.InstanceCount)
}
