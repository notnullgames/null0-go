package null0

// This file is auto-generated from api/*.yml - DO NOT EDIT

import "unsafe"

// Type Definitions

// An RGBA color.
type Color struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

// The 2D size of something (width/height.)
type Dimensions struct {
	Width int32
	Height int32
}

// The 2D position + size of something (x/y/w/h.)
type Rectangle struct {
	X int32
	Y int32
	Width int32
	Height int32
}

// Sfx parameters.
type SfxParams struct {
	DeltaSlide float32
	ChangeAmount float32
	HpfCutoffSweep float32
	RandSeed uint32
	WaveType int32
	SustainPunch float32
	StartFrequency float32
	DutySweep float32
	RepeatSpeed float32
	LpfCutoff float32
	LpfCutoffSweep float32
	DecayTime float32
	VibratoDepth float32
	VibratoSpeed float32
	ChangeSpeed float32
	SquareDuty float32
	PhaserOffset float32
	LpfResonance float32
	HpfCutoff float32
	MinFrequency float32
	PhaserSweep float32
	AttackTime float32
	SustainTime float32
	Slide float32
}

// The 2D position of something (x/y.)
type Vector struct {
	X int32
	Y int32
}

// Represents a gamepad button.
type GamepadButton uint32

const (
	GAMEPAD_BUTTON_A GamepadButton = 7
	GAMEPAD_BUTTON_B GamepadButton = 6
	GAMEPAD_BUTTON_DOWN GamepadButton = 3
	GAMEPAD_BUTTON_LEFT GamepadButton = 4
	GAMEPAD_BUTTON_LEFT_SHOULDER GamepadButton = 9
	GAMEPAD_BUTTON_LEFT_THUMB GamepadButton = 16
	GAMEPAD_BUTTON_LEFT_TRIGGER GamepadButton = 10
	GAMEPAD_BUTTON_MENU GamepadButton = 14
	GAMEPAD_BUTTON_RIGHT GamepadButton = 2
	GAMEPAD_BUTTON_RIGHT_SHOULDER GamepadButton = 11
	GAMEPAD_BUTTON_RIGHT_THUMB GamepadButton = 17
	GAMEPAD_BUTTON_RIGHT_TRIGGER GamepadButton = 12
	GAMEPAD_BUTTON_SELECT GamepadButton = 13
	GAMEPAD_BUTTON_START GamepadButton = 15
	GAMEPAD_BUTTON_UNKNOWN GamepadButton = 0
	GAMEPAD_BUTTON_UP GamepadButton = 1
	GAMEPAD_BUTTON_X GamepadButton = 8
	GAMEPAD_BUTTON_Y GamepadButton = 5
)

// Potential image-filtering techniques for scale/etc.
type ImageFilter uint32

const (
	FILTER_BILINEAR ImageFilter = 1
	FILTER_NEARESTNEIGHBOR ImageFilter = 0
	FILTER_SMOOTH ImageFilter = 2
)

// Represents a keyboard key.
type Key uint32

const (
	KEY_0 Key = 48
	KEY_1 Key = 49
	KEY_2 Key = 50
	KEY_3 Key = 51
	KEY_4 Key = 52
	KEY_5 Key = 53
	KEY_6 Key = 54
	KEY_7 Key = 55
	KEY_8 Key = 56
	KEY_9 Key = 57
	KEY_A Key = 65
	KEY_APOSTROPHE Key = 39
	KEY_B Key = 66
	KEY_BACKSLASH Key = 92
	KEY_BACKSPACE Key = 259
	KEY_C Key = 67
	KEY_CAPS_LOCK Key = 280
	KEY_COMMA Key = 44
	KEY_D Key = 68
	KEY_DELETE Key = 261
	KEY_DOWN Key = 264
	KEY_E Key = 69
	KEY_END Key = 269
	KEY_ENTER Key = 257
	KEY_EQUAL Key = 61
	KEY_ESCAPE Key = 256
	KEY_F Key = 70
	KEY_F1 Key = 290
	KEY_F10 Key = 299
	KEY_F11 Key = 300
	KEY_F12 Key = 301
	KEY_F13 Key = 302
	KEY_F14 Key = 303
	KEY_F15 Key = 304
	KEY_F16 Key = 305
	KEY_F17 Key = 306
	KEY_F18 Key = 307
	KEY_F19 Key = 308
	KEY_F2 Key = 291
	KEY_F20 Key = 309
	KEY_F21 Key = 310
	KEY_F22 Key = 311
	KEY_F23 Key = 312
	KEY_F24 Key = 313
	KEY_F25 Key = 314
	KEY_F3 Key = 292
	KEY_F4 Key = 293
	KEY_F5 Key = 294
	KEY_F6 Key = 295
	KEY_F7 Key = 296
	KEY_F8 Key = 297
	KEY_F9 Key = 298
	KEY_G Key = 71
	KEY_GRAVE_ACCENT Key = 96
	KEY_H Key = 72
	KEY_HOME Key = 268
	KEY_I Key = 73
	KEY_INSERT Key = 260
	KEY_INVALID Key = 0
	KEY_J Key = 74
	KEY_K Key = 75
	KEY_KP_0 Key = 320
	KEY_KP_1 Key = 321
	KEY_KP_2 Key = 322
	KEY_KP_3 Key = 323
	KEY_KP_4 Key = 324
	KEY_KP_5 Key = 325
	KEY_KP_6 Key = 326
	KEY_KP_7 Key = 327
	KEY_KP_8 Key = 328
	KEY_KP_9 Key = 329
	KEY_KP_ADD Key = 334
	KEY_KP_DECIMAL Key = 330
	KEY_KP_DIVIDE Key = 331
	KEY_KP_ENTER Key = 335
	KEY_KP_EQUAL Key = 336
	KEY_KP_MULTIPLY Key = 332
	KEY_KP_SUBTRACT Key = 333
	KEY_L Key = 76
	KEY_LEFT Key = 263
	KEY_LEFT_ALT Key = 342
	KEY_LEFT_BRACKET Key = 91
	KEY_LEFT_CONTROL Key = 341
	KEY_LEFT_SHIFT Key = 340
	KEY_LEFT_SUPER Key = 343
	KEY_M Key = 77
	KEY_MENU Key = 348
	KEY_MINUS Key = 45
	KEY_N Key = 78
	KEY_NUM_LOCK Key = 282
	KEY_O Key = 79
	KEY_P Key = 80
	KEY_PAGE_DOWN Key = 267
	KEY_PAGE_UP Key = 266
	KEY_PAUSE Key = 284
	KEY_PERIOD Key = 46
	KEY_PRINT_SCREEN Key = 283
	KEY_Q Key = 81
	KEY_R Key = 82
	KEY_RIGHT Key = 262
	KEY_RIGHT_ALT Key = 346
	KEY_RIGHT_BRACKET Key = 93
	KEY_RIGHT_CONTROL Key = 345
	KEY_RIGHT_SHIFT Key = 344
	KEY_RIGHT_SUPER Key = 347
	KEY_S Key = 83
	KEY_SCROLL_LOCK Key = 281
	KEY_SEMICOLON Key = 59
	KEY_SLASH Key = 47
	KEY_SPACE Key = 32
	KEY_T Key = 84
	KEY_TAB Key = 258
	KEY_U Key = 85
	KEY_UP Key = 265
	KEY_V Key = 86
	KEY_W Key = 87
	KEY_WORLD_1 Key = 161
	KEY_WORLD_2 Key = 162
	KEY_X Key = 88
	KEY_Y Key = 89
	KEY_Z Key = 90
)

// Represents a mouse button.
type MouseButton uint32

const (
	MOUSE_BUTTON_LEFT MouseButton = 1
	MOUSE_BUTTON_MIDDLE MouseButton = 3
	MOUSE_BUTTON_RIGHT MouseButton = 2
	MOUSE_BUTTON_UNKNOWN MouseButton = 0
)

// Represents a Sfx preset type.
type SfxPresetType uint32

const (
	SFX_COIN SfxPresetType = 0
	SFX_EXPLOSION SfxPresetType = 2
	SFX_HURT SfxPresetType = 4
	SFX_JUMP SfxPresetType = 5
	SFX_LASER SfxPresetType = 1
	SFX_POWERUP SfxPresetType = 3
	SFX_SELECT SfxPresetType = 6
	SFX_SYNTH SfxPresetType = 7
)

// Low-level WASM imports (internal)

//go:wasmimport null0 clear
func clear(color uint32)

//go:wasmimport null0 clear_image
func clearImage(color uint32, destination uint32)

//go:wasmimport null0 current_time
func currentTime() uint64

//go:wasmimport null0 delta_time
func deltaTime() float32

//go:wasmimport null0 draw_arc
func drawArc(centerX int32, centerY int32, color uint32, endAngle float32, radius float32, segments int32, startAngle float32)

//go:wasmimport null0 draw_arc_outline
func drawArcOutline(centerX int32, centerY int32, color uint32, endAngle float32, radius float32, segments int32, startAngle float32, thickness int32)

//go:wasmimport null0 draw_circle
func drawCircle(centerX int32, centerY int32, color uint32, radius int32)

//go:wasmimport null0 draw_circle_on_image
func drawCircleOnImage(centerX int32, centerY int32, color uint32, destination uint32, radius int32)

//go:wasmimport null0 draw_circle_outline
func drawCircleOutline(centerX int32, centerY int32, color uint32, radius int32, thickness int32)

//go:wasmimport null0 draw_circle_outline_on_image
func drawCircleOutlineOnImage(centerX int32, centerY int32, color uint32, destination uint32, radius int32, thickness int32)

//go:wasmimport null0 draw_ellipse
func drawEllipse(centerX int32, centerY int32, color uint32, radiusX int32, radiusY int32)

//go:wasmimport null0 draw_ellipse_on_image
func drawEllipseOnImage(centerX int32, centerY int32, color uint32, destination uint32, radiusX int32, radiusY int32)

//go:wasmimport null0 draw_ellipse_outline
func drawEllipseOutline(centerX int32, centerY int32, color uint32, radiusX int32, radiusY int32, thickness int32)

//go:wasmimport null0 draw_ellipse_outline_on_image
func drawEllipseOutlineOnImage(centerX int32, centerY int32, color uint32, destination uint32, radiusX int32, radiusY int32, thickness int32)

//go:wasmimport null0 draw_image
func drawImage(posX int32, posY int32, src uint32)

//go:wasmimport null0 draw_image_flipped
func drawImageFlipped(flipDiagonal bool, flipHorizontal bool, flipVertical bool, posX int32, posY int32, src uint32)

//go:wasmimport null0 draw_image_flipped_on_image
func drawImageFlippedOnImage(destination uint32, flipDiagonal bool, flipHorizontal bool, flipVertical bool, posX int32, posY int32, src uint32)

//go:wasmimport null0 draw_image_on_image
func drawImageOnImage(destination uint32, posX int32, posY int32, src uint32)

//go:wasmimport null0 draw_image_rotated
func drawImageRotated(degrees float32, filter uint32, offsetX float32, offsetY float32, posX int32, posY int32, src uint32)

//go:wasmimport null0 draw_image_rotated_on_image
func drawImageRotatedOnImage(degrees float32, destination uint32, filter uint32, offsetX float32, offsetY float32, posX int32, posY int32, src uint32)

//go:wasmimport null0 draw_image_scaled
func drawImageScaled(filter uint32, offsetX float32, offsetY float32, posX int32, posY int32, scaleX float32, scaleY float32, src uint32)

//go:wasmimport null0 draw_image_scaled_on_image
func drawImageScaledOnImage(destination uint32, filter uint32, offsetX float32, offsetY float32, posX int32, posY int32, scaleX float32, scaleY float32, src uint32)

//go:wasmimport null0 draw_image_tint
func drawImageTint(posX int32, posY int32, src uint32, tint uint32)

//go:wasmimport null0 draw_image_tint_on_image
func drawImageTintOnImage(destination uint32, posX int32, posY int32, src uint32, tint uint32)

//go:wasmimport null0 draw_line
func drawLine(color uint32, endPosX int32, endPosY int32, startPosX int32, startPosY int32)

//go:wasmimport null0 draw_line_on_image
func drawLineOnImage(color uint32, destination uint32, endPosX int32, endPosY int32, startPosX int32, startPosY int32)

//go:wasmimport null0 draw_point
func drawPoint(color uint32, x int32, y int32)

//go:wasmimport null0 draw_point_on_image
func drawPointOnImage(color uint32, destination uint32, x int32, y int32)

//go:wasmimport null0 draw_polygon
func drawPolygon(color uint32, numPoints int32, points uint32)

//go:wasmimport null0 draw_polygon_on_image
func drawPolygonOnImage(color uint32, destination uint32, numPoints int32, points uint32)

//go:wasmimport null0 draw_polygon_outline
func drawPolygonOutline(color uint32, numPoints int32, points uint32, thickness int32)

//go:wasmimport null0 draw_polygon_outline_on_image
func drawPolygonOutlineOnImage(color uint32, destination uint32, numPoints int32, points uint32, thickness int32)

//go:wasmimport null0 draw_rectangle
func drawRectangle(color uint32, height int32, posX int32, posY int32, width int32)

//go:wasmimport null0 draw_rectangle_on_image
func drawRectangleOnImage(color uint32, destination uint32, height int32, posX int32, posY int32, width int32)

//go:wasmimport null0 draw_rectangle_outline
func drawRectangleOutline(color uint32, height int32, posX int32, posY int32, thickness int32, width int32)

//go:wasmimport null0 draw_rectangle_outline_on_image
func drawRectangleOutlineOnImage(color uint32, destination uint32, height int32, posX int32, posY int32, thickness int32, width int32)

//go:wasmimport null0 draw_rectangle_rounded
func drawRectangleRounded(color uint32, cornerRadius int32, height int32, width int32, x int32, y int32)

//go:wasmimport null0 draw_rectangle_rounded_on_image
func drawRectangleRoundedOnImage(color uint32, cornerRadius int32, destination uint32, height int32, width int32, x int32, y int32)

//go:wasmimport null0 draw_rectangle_rounded_outline
func drawRectangleRoundedOutline(color uint32, cornerRadius int32, height int32, thickness int32, width int32, x int32, y int32)

//go:wasmimport null0 draw_rectangle_rounded_outline_on_image
func drawRectangleRoundedOutlineOnImage(color uint32, cornerRadius int32, destination uint32, height int32, thickness int32, width int32, x int32, y int32)

//go:wasmimport null0 draw_text
func drawText(color uint32, font uint32, posX int32, posY int32, text uint32)

//go:wasmimport null0 draw_text_on_image
func drawTextOnImage(color uint32, destination uint32, font uint32, posX int32, posY int32, text uint32)

//go:wasmimport null0 draw_triangle
func drawTriangle(color uint32, x1 int32, x2 int32, x3 int32, y1 int32, y2 int32, y3 int32)

//go:wasmimport null0 draw_triangle_on_image
func drawTriangleOnImage(color uint32, destination uint32, x1 int32, x2 int32, x3 int32, y1 int32, y2 int32, y3 int32)

//go:wasmimport null0 draw_triangle_outline
func drawTriangleOutline(color uint32, thickness int32, x1 int32, x2 int32, x3 int32, y1 int32, y2 int32, y3 int32)

//go:wasmimport null0 draw_triangle_outline_on_image
func drawTriangleOutlineOnImage(color uint32, destination uint32, thickness int32, x1 int32, x2 int32, x3 int32, y1 int32, y2 int32, y3 int32)

//go:wasmimport null0 font_copy
func fontCopy(font uint32) uint32

//go:wasmimport null0 font_scale
func fontScale(filter uint32, font uint32, scaleX float32, scaleY float32) uint32

//go:wasmimport null0 gamepad_button_down
func gamepadButtonDown(button uint32, gamepad int32) bool

//go:wasmimport null0 gamepad_button_pressed
func gamepadButtonPressed(button uint32, gamepad int32) bool

//go:wasmimport null0 gamepad_button_released
func gamepadButtonReleased(button uint32, gamepad int32) bool

//go:wasmimport null0 image_alpha_border
func imageAlphaBorder(image uint32, threshold float32) uint32

//go:wasmimport null0 image_alpha_crop
func imageAlphaCrop(image uint32, threshold float32)

//go:wasmimport null0 image_alpha_mask
func imageAlphaMask(alphaMask uint32, image uint32, posX int32, posY int32)

//go:wasmimport null0 image_color_brightness
func imageColorBrightness(factor float32, image uint32)

//go:wasmimport null0 image_color_contrast
func imageColorContrast(contrast float32, image uint32)

//go:wasmimport null0 image_color_fade
func imageColorFade(alpha float32, image uint32)

//go:wasmimport null0 image_color_invert
func imageColorInvert(image uint32)

//go:wasmimport null0 image_color_replace
func imageColorReplace(color uint32, image uint32, replace uint32)

//go:wasmimport null0 image_color_tint
func imageColorTint(color uint32, image uint32)

//go:wasmimport null0 image_copy
func imageCopy(image uint32) uint32

//go:wasmimport null0 image_crop
func imageCrop(height int32, image uint32, width int32, x int32, y int32)

//go:wasmimport null0 image_flip
func imageFlip(horizontal bool, image uint32, vertical bool)

//go:wasmimport null0 image_gradient
func imageGradient(bottomLeft uint32, bottomRight uint32, height int32, topLeft uint32, topRight uint32, width int32) uint32

//go:wasmimport null0 image_resize
func imageResize(filter uint32, image uint32, newHeight int32, newWidth int32) uint32

//go:wasmimport null0 image_rotate
func imageRotate(degrees float32, filter uint32, image uint32) uint32

//go:wasmimport null0 image_scale
func imageScale(filter uint32, image uint32, scaleX float32, scaleY float32) uint32

//go:wasmimport null0 image_subimage
func imageSubimage(height int32, image uint32, width int32, x int32, y int32) uint32

//go:wasmimport null0 key_down
func keyDown(key uint32) bool

//go:wasmimport null0 key_pressed
func keyPressed(key uint32) bool

//go:wasmimport null0 key_released
func keyReleased(key uint32) bool

//go:wasmimport null0 key_up
func keyUp(key uint32) bool

//go:wasmimport null0 load_font_bmf
func loadFontBmf(characters uint32, filename uint32) uint32

//go:wasmimport null0 load_font_bmf_from_image
func loadFontBmfFromImage(characters uint32, image uint32) uint32

//go:wasmimport null0 load_font_ttf
func loadFontTtf(filename uint32, fontSize int32) uint32

//go:wasmimport null0 load_font_tty
func loadFontTty(characters uint32, filename uint32, glyphHeight int32, glyphWidth int32) uint32

//go:wasmimport null0 load_font_tty_from_image
func loadFontTtyFromImage(characters uint32, glyphHeight int32, glyphWidth int32, image uint32) uint32

//go:wasmimport null0 load_image
func loadImage(filename uint32) uint32

//go:wasmimport null0 load_sound
func loadSound(filename uint32) uint32

//go:wasmimport null0 measure_image
func measureImage(image uint32) uint32

//go:wasmimport null0 measure_text
func measureText(font uint32, text uint32, textLength int32) uint32

//go:wasmimport null0 mouse_button_down
func mouseButtonDown(button uint32) bool

//go:wasmimport null0 mouse_button_pressed
func mouseButtonPressed(button uint32) bool

//go:wasmimport null0 mouse_button_released
func mouseButtonReleased(button uint32) bool

//go:wasmimport null0 mouse_button_up
func mouseButtonUp(button uint32) bool

//go:wasmimport null0 mouse_position
func mousePosition() uint32

//go:wasmimport null0 new_image
func newImage(color uint32, height int32, width int32) uint32

//go:wasmimport null0 play_sound
func playSound(loop bool, sound uint32)

//go:wasmimport null0 random_int
func randomInt(max int32, min int32) int32

//go:wasmimport null0 random_seed_get
func randomSeedGet() uint64

//go:wasmimport null0 random_seed_set
func randomSeedSet(seed uint64)

//go:wasmimport null0 save_image
func saveImage(filename uint32, image uint32)

//go:wasmimport null0 sfx_generate
func sfxGenerate(type_ uint32) uint32

//go:wasmimport null0 sfx_sound
func sfxSound(params uint32) uint32

//go:wasmimport null0 stop_sound
func stopSound(sound uint32)

//go:wasmimport null0 tts_sound
func ttsSound(mouth int32, phonetic bool, pitch int32, sing bool, speed int32, text uint32, throat int32) uint32

//go:wasmimport null0 unload_font
func unloadFont(font uint32)

//go:wasmimport null0 unload_image
func unloadImage(image uint32)

//go:wasmimport null0 unload_sound
func unloadSound(sound uint32)

// Public API Functions

// Clear the screen.
func Clear(color Color) {
	clear(uint32(uintptr(unsafe.Pointer(&color))))
}

// Clear an image.
func ClearImage(color Color, destination uint32) {
	clearImage(uint32(uintptr(unsafe.Pointer(&color))), destination)
}

// Get system-time (ms) since unix epoch.
func CurrentTime() uint64 {
	return currentTime()
}

// Get the change in time (seconds) since the last update run.
func DeltaTime() float32 {
	return deltaTime()
}

// Draw a filled arc on the screen.
func DrawArc(centerX int32, centerY int32, color Color, endAngle float32, radius float32, segments int32, startAngle float32) {
	drawArc(centerX, centerY, uint32(uintptr(unsafe.Pointer(&color))), endAngle, radius, segments, startAngle)
}

// Draw a outlined (with thickness) arc on the screen.
func DrawArcOutline(centerX int32, centerY int32, color Color, endAngle float32, radius float32, segments int32, startAngle float32, thickness int32) {
	drawArcOutline(centerX, centerY, uint32(uintptr(unsafe.Pointer(&color))), endAngle, radius, segments, startAngle, thickness)
}

// Draw a filled circle on the screen.
func DrawCircle(centerX int32, centerY int32, color Color, radius int32) {
	drawCircle(centerX, centerY, uint32(uintptr(unsafe.Pointer(&color))), radius)
}

// Draw a circle on an image.
func DrawCircleOnImage(centerX int32, centerY int32, color Color, destination uint32, radius int32) {
	drawCircleOnImage(centerX, centerY, uint32(uintptr(unsafe.Pointer(&color))), destination, radius)
}

// Draw a outlined (with thickness) circle on the screen.
func DrawCircleOutline(centerX int32, centerY int32, color Color, radius int32, thickness int32) {
	drawCircleOutline(centerX, centerY, uint32(uintptr(unsafe.Pointer(&color))), radius, thickness)
}

// Draw a outlined (with thickness) circle on an image.
func DrawCircleOutlineOnImage(centerX int32, centerY int32, color Color, destination uint32, radius int32, thickness int32) {
	drawCircleOutlineOnImage(centerX, centerY, uint32(uintptr(unsafe.Pointer(&color))), destination, radius, thickness)
}

// Draw a filled ellipse on the screen.
func DrawEllipse(centerX int32, centerY int32, color Color, radiusX int32, radiusY int32) {
	drawEllipse(centerX, centerY, uint32(uintptr(unsafe.Pointer(&color))), radiusX, radiusY)
}

// Draw a filled ellipse on an image.
func DrawEllipseOnImage(centerX int32, centerY int32, color Color, destination uint32, radiusX int32, radiusY int32) {
	drawEllipseOnImage(centerX, centerY, uint32(uintptr(unsafe.Pointer(&color))), destination, radiusX, radiusY)
}

// Draw a outlined (with thickness) ellipse on the screen.
func DrawEllipseOutline(centerX int32, centerY int32, color Color, radiusX int32, radiusY int32, thickness int32) {
	drawEllipseOutline(centerX, centerY, uint32(uintptr(unsafe.Pointer(&color))), radiusX, radiusY, thickness)
}

// Draw a outlined (with thickness) ellipse on an image.
func DrawEllipseOutlineOnImage(centerX int32, centerY int32, color Color, destination uint32, radiusX int32, radiusY int32, thickness int32) {
	drawEllipseOutlineOnImage(centerX, centerY, uint32(uintptr(unsafe.Pointer(&color))), destination, radiusX, radiusY, thickness)
}

// Draw an image on the screen.
func DrawImage(posX int32, posY int32, src uint32) {
	drawImage(posX, posY, src)
}

// Draw an image, flipped, on the screen.
func DrawImageFlipped(flipDiagonal bool, flipHorizontal bool, flipVertical bool, posX int32, posY int32, src uint32) {
	drawImageFlipped(flipDiagonal, flipHorizontal, flipVertical, posX, posY, src)
}

// Draw an image, flipped, on an image.
func DrawImageFlippedOnImage(destination uint32, flipDiagonal bool, flipHorizontal bool, flipVertical bool, posX int32, posY int32, src uint32) {
	drawImageFlippedOnImage(destination, flipDiagonal, flipHorizontal, flipVertical, posX, posY, src)
}

// Draw an image on an image.
func DrawImageOnImage(destination uint32, posX int32, posY int32, src uint32) {
	drawImageOnImage(destination, posX, posY, src)
}

// Draw an image, rotated, on the screen.
func DrawImageRotated(degrees float32, filter ImageFilter, offsetX float32, offsetY float32, posX int32, posY int32, src uint32) {
	drawImageRotated(degrees, uint32(filter), offsetX, offsetY, posX, posY, src)
}

// Draw an image, rotated, on an image.
func DrawImageRotatedOnImage(degrees float32, destination uint32, filter ImageFilter, offsetX float32, offsetY float32, posX int32, posY int32, src uint32) {
	drawImageRotatedOnImage(degrees, destination, uint32(filter), offsetX, offsetY, posX, posY, src)
}

// Draw an image, scaled, on the screen.
func DrawImageScaled(filter ImageFilter, offsetX float32, offsetY float32, posX int32, posY int32, scaleX float32, scaleY float32, src uint32) {
	drawImageScaled(uint32(filter), offsetX, offsetY, posX, posY, scaleX, scaleY, src)
}

// Draw an image, scaled, on an image.
func DrawImageScaledOnImage(destination uint32, filter ImageFilter, offsetX float32, offsetY float32, posX int32, posY int32, scaleX float32, scaleY float32, src uint32) {
	drawImageScaledOnImage(destination, uint32(filter), offsetX, offsetY, posX, posY, scaleX, scaleY, src)
}

// Draw a tinted image on the screen.
func DrawImageTint(posX int32, posY int32, src uint32, tint Color) {
	drawImageTint(posX, posY, src, uint32(uintptr(unsafe.Pointer(&tint))))
}

// Draw a tinted image on an image.
func DrawImageTintOnImage(destination uint32, posX int32, posY int32, src uint32, tint Color) {
	drawImageTintOnImage(destination, posX, posY, src, uint32(uintptr(unsafe.Pointer(&tint))))
}

// Draw a line on the screen.
func DrawLine(color Color, endPosX int32, endPosY int32, startPosX int32, startPosY int32) {
	drawLine(uint32(uintptr(unsafe.Pointer(&color))), endPosX, endPosY, startPosX, startPosY)
}

// Draw a line on an image.
func DrawLineOnImage(color Color, destination uint32, endPosX int32, endPosY int32, startPosX int32, startPosY int32) {
	drawLineOnImage(uint32(uintptr(unsafe.Pointer(&color))), destination, endPosX, endPosY, startPosX, startPosY)
}

// Draw a single pixel on the screen.
func DrawPoint(color Color, x int32, y int32) {
	drawPoint(uint32(uintptr(unsafe.Pointer(&color))), x, y)
}

// Draw a single pixel on an image.
func DrawPointOnImage(color Color, destination uint32, x int32, y int32) {
	drawPointOnImage(uint32(uintptr(unsafe.Pointer(&color))), destination, x, y)
}

// Draw a filled polygon on the screen.
func DrawPolygon(color Color, numPoints int32, points []Vector) {
	drawPolygon(uint32(uintptr(unsafe.Pointer(&color))), numPoints, uint32(uintptr(unsafe.Pointer(&points))))
}

// Draw a filled polygon on an image.
func DrawPolygonOnImage(color Color, destination uint32, numPoints int32, points []Vector) {
	drawPolygonOnImage(uint32(uintptr(unsafe.Pointer(&color))), destination, numPoints, uint32(uintptr(unsafe.Pointer(&points))))
}

// Draw a outlined (with thickness) polygon on the screen.
func DrawPolygonOutline(color Color, numPoints int32, points []Vector, thickness int32) {
	drawPolygonOutline(uint32(uintptr(unsafe.Pointer(&color))), numPoints, uint32(uintptr(unsafe.Pointer(&points))), thickness)
}

// Draw a outlined (with thickness) polygon on an image.
func DrawPolygonOutlineOnImage(color Color, destination uint32, numPoints int32, points []Vector, thickness int32) {
	drawPolygonOutlineOnImage(uint32(uintptr(unsafe.Pointer(&color))), destination, numPoints, uint32(uintptr(unsafe.Pointer(&points))), thickness)
}

// Draw a filled rectangle on the screen.
func DrawRectangle(color Color, height int32, posX int32, posY int32, width int32) {
	drawRectangle(uint32(uintptr(unsafe.Pointer(&color))), height, posX, posY, width)
}

// Draw a filled rectangle on an image.
func DrawRectangleOnImage(color Color, destination uint32, height int32, posX int32, posY int32, width int32) {
	drawRectangleOnImage(uint32(uintptr(unsafe.Pointer(&color))), destination, height, posX, posY, width)
}

// Draw a outlined (with thickness) rectangle on the screen.
func DrawRectangleOutline(color Color, height int32, posX int32, posY int32, thickness int32, width int32) {
	drawRectangleOutline(uint32(uintptr(unsafe.Pointer(&color))), height, posX, posY, thickness, width)
}

// Draw a outlined (with thickness) rectangle on an image.
func DrawRectangleOutlineOnImage(color Color, destination uint32, height int32, posX int32, posY int32, thickness int32, width int32) {
	drawRectangleOutlineOnImage(uint32(uintptr(unsafe.Pointer(&color))), destination, height, posX, posY, thickness, width)
}

// Draw a filled round-rectangle on the screen.
func DrawRectangleRounded(color Color, cornerRadius int32, height int32, width int32, x int32, y int32) {
	drawRectangleRounded(uint32(uintptr(unsafe.Pointer(&color))), cornerRadius, height, width, x, y)
}

// Draw a filled round-rectangle on an image.
func DrawRectangleRoundedOnImage(color Color, cornerRadius int32, destination uint32, height int32, width int32, x int32, y int32) {
	drawRectangleRoundedOnImage(uint32(uintptr(unsafe.Pointer(&color))), cornerRadius, destination, height, width, x, y)
}

// Draw a outlined (with thickness) round-rectangle on the screen.
func DrawRectangleRoundedOutline(color Color, cornerRadius int32, height int32, thickness int32, width int32, x int32, y int32) {
	drawRectangleRoundedOutline(uint32(uintptr(unsafe.Pointer(&color))), cornerRadius, height, thickness, width, x, y)
}

// Draw a outlined (with thickness) round-rectangle on an image.
func DrawRectangleRoundedOutlineOnImage(color Color, cornerRadius int32, destination uint32, height int32, thickness int32, width int32, x int32, y int32) {
	drawRectangleRoundedOutlineOnImage(uint32(uintptr(unsafe.Pointer(&color))), cornerRadius, destination, height, thickness, width, x, y)
}

// Draw some text on the screen.
func DrawText(color Color, font uint32, posX int32, posY int32, text string) {
	drawText(uint32(uintptr(unsafe.Pointer(&color))), font, posX, posY, uint32(uintptr(unsafe.Pointer(unsafe.StringData(text)))))
}

// Draw some text on an image.
func DrawTextOnImage(color Color, destination uint32, font uint32, posX int32, posY int32, text string) {
	drawTextOnImage(uint32(uintptr(unsafe.Pointer(&color))), destination, font, posX, posY, uint32(uintptr(unsafe.Pointer(unsafe.StringData(text)))))
}

// Draw a filled triangle on the screen.
func DrawTriangle(color Color, x1 int32, x2 int32, x3 int32, y1 int32, y2 int32, y3 int32) {
	drawTriangle(uint32(uintptr(unsafe.Pointer(&color))), x1, x2, x3, y1, y2, y3)
}

// Draw a filled triangle on an image.
func DrawTriangleOnImage(color Color, destination uint32, x1 int32, x2 int32, x3 int32, y1 int32, y2 int32, y3 int32) {
	drawTriangleOnImage(uint32(uintptr(unsafe.Pointer(&color))), destination, x1, x2, x3, y1, y2, y3)
}

// Draw a outlined (with thickness) triangle on the screen.
func DrawTriangleOutline(color Color, thickness int32, x1 int32, x2 int32, x3 int32, y1 int32, y2 int32, y3 int32) {
	drawTriangleOutline(uint32(uintptr(unsafe.Pointer(&color))), thickness, x1, x2, x3, y1, y2, y3)
}

// Draw a outlined (with thickness) triangle on an image.
func DrawTriangleOutlineOnImage(color Color, destination uint32, thickness int32, x1 int32, x2 int32, x3 int32, y1 int32, y2 int32, y3 int32) {
	drawTriangleOutlineOnImage(uint32(uintptr(unsafe.Pointer(&color))), destination, thickness, x1, x2, x3, y1, y2, y3)
}

// Copy a font to a new font.
func FontCopy(font uint32) uint32 {
	return fontCopy(font)
}

// Scale a font, return a new font.
func FontScale(filter ImageFilter, font uint32, scaleX float32, scaleY float32) uint32 {
	return fontScale(uint32(filter), font, scaleX, scaleY)
}

// Is the button currently down?
func GamepadButtonDown(button GamepadButton, gamepad int32) bool {
	return gamepadButtonDown(uint32(button), gamepad)
}

// Has the button been pressed? (tracks unpress/read correctly.)
func GamepadButtonPressed(button GamepadButton, gamepad int32) bool {
	return gamepadButtonPressed(uint32(button), gamepad)
}

// Has the button been released? (tracks press/read correctly.)
func GamepadButtonReleased(button GamepadButton, gamepad int32) bool {
	return gamepadButtonReleased(uint32(button), gamepad)
}

// Calculate a rectangle representing the available alpha border in an image.
func ImageAlphaBorder(image uint32, threshold float32) Rectangle {
	var result Rectangle
	resultPtr := imageAlphaBorder(image, threshold)
	if resultPtr != 0 {
		result = *(*Rectangle)(unsafe.Pointer(uintptr(resultPtr)))
	}
	return result
}

// Crop an image based on the alpha border, in-place.
func ImageAlphaCrop(image uint32, threshold float32) {
	imageAlphaCrop(image, threshold)
}

// Use an image as an alpha-mask on another image.
func ImageAlphaMask(alphaMask uint32, image uint32, posX int32, posY int32) {
	imageAlphaMask(alphaMask, image, posX, posY)
}

// Adjust the brightness of an image, in-place.
func ImageColorBrightness(factor float32, image uint32) {
	imageColorBrightness(factor, image)
}

// Change the contrast of an image, in-place.
func ImageColorContrast(contrast float32, image uint32) {
	imageColorContrast(contrast, image)
}

// Fade a color in an image, in-place.
func ImageColorFade(alpha float32, image uint32) {
	imageColorFade(alpha, image)
}

// Invert the colors in an image, in-place.
func ImageColorInvert(image uint32) {
	imageColorInvert(image)
}

// Replace a color in an image, in-place.
func ImageColorReplace(color Color, image uint32, replace Color) {
	imageColorReplace(uint32(uintptr(unsafe.Pointer(&color))), image, uint32(uintptr(unsafe.Pointer(&replace))))
}

// Tint a color in an image, in-place.
func ImageColorTint(color Color, image uint32) {
	imageColorTint(uint32(uintptr(unsafe.Pointer(&color))), image)
}

// Copy an image to a new image.
func ImageCopy(image uint32) uint32 {
	return imageCopy(image)
}

// Crop an image, in-place.
func ImageCrop(height int32, image uint32, width int32, x int32, y int32) {
	imageCrop(height, image, width, x, y)
}

// Flip an image, in-place.
func ImageFlip(horizontal bool, image uint32, vertical bool) {
	imageFlip(horizontal, image, vertical)
}

// Create a new image of a gradient.
func ImageGradient(bottomLeft Color, bottomRight Color, height int32, topLeft Color, topRight Color, width int32) uint32 {
	return imageGradient(uint32(uintptr(unsafe.Pointer(&bottomLeft))), uint32(uintptr(unsafe.Pointer(&bottomRight))), height, uint32(uintptr(unsafe.Pointer(&topLeft))), uint32(uintptr(unsafe.Pointer(&topRight))), width)
}

// Resize an image, return copy.
func ImageResize(filter ImageFilter, image uint32, newHeight int32, newWidth int32) uint32 {
	return imageResize(uint32(filter), image, newHeight, newWidth)
}

// Create a new image, rotating another image.
func ImageRotate(degrees float32, filter ImageFilter, image uint32) uint32 {
	return imageRotate(degrees, uint32(filter), image)
}

// Scale an image, return copy.
func ImageScale(filter ImageFilter, image uint32, scaleX float32, scaleY float32) uint32 {
	return imageScale(uint32(filter), image, scaleX, scaleY)
}

// Create an image from a region of another image.
func ImageSubimage(height int32, image uint32, width int32, x int32, y int32) uint32 {
	return imageSubimage(height, image, width, x, y)
}

// Is the key currently down?
func KeyDown(key Key) bool {
	return keyDown(uint32(key))
}

// Has the key been pressed? (tracks unpress/read correctly.)
func KeyPressed(key Key) bool {
	return keyPressed(uint32(key))
}

// Has the key been released? (tracks press/read correctly.)
func KeyReleased(key Key) bool {
	return keyReleased(uint32(key))
}

// Is the key currently up?
func KeyUp(key Key) bool {
	return keyUp(uint32(key))
}

// Load a BMF font from a file in cart.
func LoadFontBmf(characters string, filename string) uint32 {
	return loadFontBmf(uint32(uintptr(unsafe.Pointer(unsafe.StringData(characters)))), uint32(uintptr(unsafe.Pointer(unsafe.StringData(filename)))))
}

// Load a BMF font from an image.
func LoadFontBmfFromImage(characters string, image uint32) uint32 {
	return loadFontBmfFromImage(uint32(uintptr(unsafe.Pointer(unsafe.StringData(characters)))), image)
}

// Load a TTF font from a file in cart.
func LoadFontTtf(filename string, fontSize int32) uint32 {
	return loadFontTtf(uint32(uintptr(unsafe.Pointer(unsafe.StringData(filename)))), fontSize)
}

// Load a TTY font from a file in cart.
func LoadFontTty(characters string, filename string, glyphHeight int32, glyphWidth int32) uint32 {
	return loadFontTty(uint32(uintptr(unsafe.Pointer(unsafe.StringData(characters)))), uint32(uintptr(unsafe.Pointer(unsafe.StringData(filename)))), glyphHeight, glyphWidth)
}

// Load a TTY font from an image.
func LoadFontTtyFromImage(characters string, glyphHeight int32, glyphWidth int32, image uint32) uint32 {
	return loadFontTtyFromImage(uint32(uintptr(unsafe.Pointer(unsafe.StringData(characters)))), glyphHeight, glyphWidth, image)
}

// Load an image from a file in cart.
func LoadImage(filename string) uint32 {
	return loadImage(uint32(uintptr(unsafe.Pointer(unsafe.StringData(filename)))))
}

// Load a sound from a file in cart.
func LoadSound(filename string) uint32 {
	return loadSound(uint32(uintptr(unsafe.Pointer(unsafe.StringData(filename)))))
}

// Meaure an image (use 0 for screen).
func MeasureImage(image uint32) Dimensions {
	var result Dimensions
	resultPtr := measureImage(image)
	if resultPtr != 0 {
		result = *(*Dimensions)(unsafe.Pointer(uintptr(resultPtr)))
	}
	return result
}

// Measure the size of some text.
func MeasureText(font uint32, text string, textLength int32) Dimensions {
	var result Dimensions
	resultPtr := measureText(font, uint32(uintptr(unsafe.Pointer(unsafe.StringData(text)))), textLength)
	if resultPtr != 0 {
		result = *(*Dimensions)(unsafe.Pointer(uintptr(resultPtr)))
	}
	return result
}

// Is the button currently down?
func MouseButtonDown(button MouseButton) bool {
	return mouseButtonDown(uint32(button))
}

// Has the button been pressed? (tracks unpress/read correctly.)
func MouseButtonPressed(button MouseButton) bool {
	return mouseButtonPressed(uint32(button))
}

// Has the button been released? (tracks press/read correctly.)
func MouseButtonReleased(button MouseButton) bool {
	return mouseButtonReleased(uint32(button))
}

// Is the button currently up?
func MouseButtonUp(button MouseButton) bool {
	return mouseButtonUp(uint32(button))
}

// Get current position of mouse.
func MousePosition() Vector {
	var result Vector
	resultPtr := mousePosition()
	if resultPtr != 0 {
		result = *(*Vector)(unsafe.Pointer(uintptr(resultPtr)))
	}
	return result
}

// Create a new blank image.
func NewImage(color Color, height int32, width int32) uint32 {
	return newImage(uint32(uintptr(unsafe.Pointer(&color))), height, width)
}

// Play a sound.
func PlaySound(loop bool, sound uint32) {
	playSound(loop, sound)
}

// Get a random integer between 2 numbers.
func RandomInt(max int32, min int32) int32 {
	return randomInt(max, min)
}

// Get the random-seed.
func RandomSeedGet() uint64 {
	return randomSeedGet()
}

// Set the random-seed.
func RandomSeedSet(seed uint64) {
	randomSeedSet(seed)
}

// Save an image to persistant storage.
func SaveImage(filename string, image uint32) {
	saveImage(uint32(uintptr(unsafe.Pointer(unsafe.StringData(filename)))), image)
}

// Create Sfx parameters.
func SfxGenerate(type_ SfxPresetType) SfxParams {
	var result SfxParams
	resultPtr := sfxGenerate(uint32(type_))
	if resultPtr != 0 {
		result = *(*SfxParams)(unsafe.Pointer(uintptr(resultPtr)))
	}
	return result
}

// Create Sfx sound.
func SfxSound(params SfxParams) uint32 {
	return sfxSound(uint32(uintptr(unsafe.Pointer(&params))))
}

// Stop a sound.
func StopSound(sound uint32) {
	stopSound(sound)
}

// Speak some text and return a sound. Set things to 0 for defaults.
func TtsSound(mouth int32, phonetic bool, pitch int32, sing bool, speed int32, text string, throat int32) uint32 {
	return ttsSound(mouth, phonetic, pitch, sing, speed, uint32(uintptr(unsafe.Pointer(unsafe.StringData(text)))), throat)
}

// Unload a font.
func UnloadFont(font uint32) {
	unloadFont(font)
}

// Unload an image.
func UnloadImage(image uint32) {
	unloadImage(image)
}

// Unload a sound.
func UnloadSound(sound uint32) {
	unloadSound(sound)
}

