package main

import (
	"fmt"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
)

const vertexShader = `
#version 330

uniform mat4 proj;
uniform mat4 view;
uniform ivec3 trans;
uniform vec4 color;

vec3 light_dir = vec3(0.6, 0.2, -0.5);

layout(location = 0) in vec3 vert;
layout(location = 1) in vec3 normal;

flat out vec4 fcolor;

void main() {
	// Set vertex position
	gl_Position = proj * view * vec4(vert + trans, 1);
	float intensity = 0.45 * max(-0.5, dot(normal, normalize(light_dir)));
	fcolor = clamp(color + vec4(intensity, intensity, intensity, 1.0), 0.0, 1.0);
	fcolor.w = color.w;
}
` + "\x00"

const fragmentShader = `
#version 330

flat in vec4 fcolor;
out vec4 frag_colour;

void main() {
	frag_colour = fcolor;
}
` + "\x00"

type shaderProgram struct {
	id uint32

	projUniform  int32
	viewUniform  int32
	transUniform int32
	colorUniform int32
}

func (prog *shaderProgram) Create() {
	prog.id = createShaderProgram(vertexShader, fragmentShader)

	gl.UseProgram(prog.id)
	prog.projUniform = gl.GetUniformLocation(prog.id, gl.Str("proj\x00"))
	prog.viewUniform = gl.GetUniformLocation(prog.id, gl.Str("view\x00"))
	prog.transUniform = gl.GetUniformLocation(prog.id, gl.Str("trans\x00"))
	prog.colorUniform = gl.GetUniformLocation(prog.id, gl.Str("color\x00"))
}

// Compile the shader
func compileShader(source string, shaderType uint32) (uint32, error) {
	glSrcs, freeFn := gl.Strs(source)
	defer freeFn()

	shader := gl.CreateShader(shaderType)
	gl.ShaderSource(shader, 1, glSrcs, nil)
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}

// Create load the shader files and create a OpenGL program
func createShaderProgram(vertexSource, fragmentSource string) uint32 {
	vertexShader, err := compileShader(vertexSource, gl.VERTEX_SHADER)
	checkPanic(err)

	fragmentShader, err := compileShader(fragmentSource, gl.FRAGMENT_SHADER)
	checkPanic(err)

	program := gl.CreateProgram()

	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)

	gl.LinkProgram(program)

	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(log))

		panic(fmt.Errorf("failed to link program: %v", log))
	}

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	return program
}
