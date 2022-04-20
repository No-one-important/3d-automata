package main

import (
	"runtime"
	"time"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

// Engine ...
type Engine struct {
	window   *glfw.Window
	prog     shaderProgram
	automata Automata
}

func init() {
	runtime.LockOSThread()
}

// Init the engine
func (e *Engine) Init() {
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
	e.window, err = glfw.CreateWindow(1000, 1000, "3D Automata", nil, nil)
	checkPanic(err)

	// Init GL context
	e.window.MakeContextCurrent()
	checkPanic(gl.Init())

	// Display version
	log.Println("OpenGL version", gl.GoStr(gl.GetString(gl.VERSION)))

	e.automata.Init(70)

	e.prog.Create()
}

// Run through the game logic
func (e *Engine) Run() {
	const fps = 120

	var cl Clock
	cl.Init()

	// Configure global settings
	gl.Enable(gl.DEPTH_TEST)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	gl.ClearColor(0, 0, 0, 1)

	// Matrices
	projection := mgl32.Perspective(mgl32.DegToRad(45.0), 1.0, 0.1, 1000.0)

	gl.UseProgram(e.prog.id)
	gl.UniformMatrix4fv(e.prog.projUniform, 1, false, &projection[0])
	gl.Uniform4f(e.prog.colorUniform, 0.6, 0.6, 0.6, 1)

	eye := mgl32.Vec3{270, 270, 270}
	center := mgl32.Vec3{0, 0, 0}
	up := mgl32.Vec3{0, 1, 0}

	t := 0.0
	for !e.window.ShouldClose() {
		cl.Tic()
		print("\rFPS:", int(1.0/cl.GetElapsed()), "         ")

		// Simulate automata
		if t > 0.2 {
			e.automata.Simulate()
			t = 0.0
		}

		// Scene rendering
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		// Rotate scene
		rotate := mgl32.Rotate3DY(0.005)
		eye = rotate.Mul3x1(eye)
		view := mgl32.LookAtV(eye, center, up)
		gl.UniformMatrix4fv(e.prog.viewUniform, 1, false, &view[0])

		e.automata.Draw(&e.prog)

		e.window.SwapBuffers()

		// Detect inputs
		glfw.PollEvents()

		// Get end of frame time & wait for PFS cap
		toc := cl.Toc()
		t += toc
		time.Sleep(time.Duration(1.0/fps - toc))
	}
}

// Stop the engine's execution
func (e *Engine) Stop() {
	glfw.Terminate()
}
