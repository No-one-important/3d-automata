package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

// VAO stores the VAO and its buffers
type VAO struct {
	Ptr uint32

	VertexBuffer  uint32
	NormalsBuffer uint32
	ColorsBuffer  uint32
	IndexBuffer   uint32

	hasElements bool
	count       int32
}

// returns a vertex array from the vertices provided
func makeVAO(vertices, normals, colors []float32, indices []uint32) VAO {
	var vao VAO

	// Create VAO buffer
	gl.GenVertexArrays(1, &vao.Ptr)
	gl.BindVertexArray(vao.Ptr)

	// Create Vertices buffer
	if vertices != nil {
		gl.GenBuffers(1, &vao.VertexBuffer)
		gl.BindBuffer(gl.ARRAY_BUFFER, vao.VertexBuffer)
		gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

		gl.EnableVertexAttribArray(0)
		gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)
	} else {
		gl.DisableVertexAttribArray(0)
	}

	// Create Normals buffer
	if normals != nil {
		gl.GenBuffers(1, &vao.NormalsBuffer)
		gl.BindBuffer(gl.ARRAY_BUFFER, vao.NormalsBuffer)
		gl.BufferData(gl.ARRAY_BUFFER, len(normals)*4, gl.Ptr(normals), gl.STATIC_DRAW)

		gl.EnableVertexAttribArray(1)
		gl.VertexAttribPointer(1, 3, gl.FLOAT, false, 0, nil)
	} else {
		gl.DisableVertexAttribArray(1)
	}

	// Create Color buffer
	if colors != nil {
		gl.GenBuffers(1, &vao.ColorsBuffer)
		gl.BindBuffer(gl.ARRAY_BUFFER, vao.ColorsBuffer)
		gl.BufferData(gl.ARRAY_BUFFER, len(colors)*4, gl.Ptr(colors), gl.STATIC_DRAW)

		gl.EnableVertexAttribArray(2)
		gl.VertexAttribPointer(2, 4, gl.FLOAT, false, 0, nil)
	} else {
		gl.DisableVertexAttribArray(2)
	}

	if indices != nil {
		gl.GenBuffers(1, &vao.IndexBuffer)
		gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, vao.IndexBuffer)
		gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*4, gl.Ptr(indices), gl.STATIC_DRAW)

		vao.hasElements = true
	}

	// Disable VAO
	gl.BindVertexArray(0)

	return vao
}
