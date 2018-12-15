package main

import (
	"log"
	"runtime"
	"time"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

const fps = 60

var window *glfw.Window

func init() {
	runtime.LockOSThread()
}

var cube Object
var prog flatShaderProgram

// Create the engine
func create() {
	var err error

	// Init window
	checkPanic(glfw.Init())

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.Samples, 4)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	// Create a Window
	window, err = glfw.CreateWindow(1000, 1000, "3D Automata", nil, nil)
	checkPanic(err)

	window.MakeContextCurrent()

	// Init GL context
	checkPanic(gl.Init())

	// Display version
	log.Println("OpenGL version", gl.GoStr(gl.GetString(gl.VERSION)))

	// Shader
	prog.id = createShaderProgram(flatVertexShader, flatFragmentShader)

	gl.UseProgram(prog.id)
	prog.projUniform = gl.GetUniformLocation(prog.id, gl.Str("proj\x00"))
	prog.viewUniform = gl.GetUniformLocation(prog.id, gl.Str("view\x00"))
	prog.modelUniform = gl.GetUniformLocation(prog.id, gl.Str("model\x00"))
	prog.colorUniform = gl.GetUniformLocation(prog.id, gl.Str("color\x00"))

	cube.Create3D(cubeVertices, cubeNormals, nil, nil, gl.TRIANGLES)
}

// Loop through the game logic
func loop() {
	var cl Clock
	cl.Init()

	// Configure global settings
	gl.Enable(gl.DEPTH_TEST)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	gl.ClearColor(0, 0, 0, 1)

	// Matrices
	projection := mgl32.Perspective(mgl32.DegToRad(45.0), 1.0, 0.1, 100.0)
	model := mgl32.Ident4()

	gl.UseProgram(prog.id)
	gl.UniformMatrix4fv(prog.projUniform, 1, false, &projection[0])
	gl.UniformMatrix4fv(prog.modelUniform, 1, false, &model[0])
	gl.Uniform4f(prog.colorUniform, 0.6, 0.6, 0.6, 1)

	eye := mgl32.Vec3{5, 5, 5}
	center := mgl32.Vec3{0, 0, 0}
	up := mgl32.Vec3{0, 1, 0}

	for !window.ShouldClose() {
		// Get start of frame time
		cl.Tic()
		// print("\rFPS:", int(1.0/cl.GetElapsed()), "         ")

		// Scene rendering
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		// Rotate scene
		rotate := mgl32.Rotate3DY(0.01)
		eye = rotate.Mul3x1(eye)
		view := mgl32.LookAtV(eye, center, up)
		gl.UniformMatrix4fv(prog.viewUniform, 1, false, &view[0])

		// Draw scene
		cube.Draw()

		window.SwapBuffers()

		// Detect inputs
		glfw.PollEvents()

		// Get end of frame time & wait for PFS cap
		time.Sleep(time.Microsecond * time.Duration((1.0/fps-cl.Toc())*1e6))
	}
}

// Stop the engine's execution
func stop() {
	glfw.Terminate()
}
