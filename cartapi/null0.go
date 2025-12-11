package null0

// This file is auto-generated from api/*.yml - DO NOT EDIT

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
	WaveType int32
	MinFrequency float32
	VibratoSpeed float32
	ChangeSpeed float32
	PhaserOffset float32
	LpfResonance float32
	HpfCutoff float32
	AttackTime float32
	StartFrequency float32
	VibratoDepth float32
	ChangeAmount float32
	LpfCutoff float32
	SustainTime float32
	SustainPunch float32
	Slide float32
	DeltaSlide float32
	SquareDuty float32
	DutySweep float32
	PhaserSweep float32
	HpfCutoffSweep float32
	RandSeed uint32
	DecayTime float32
	RepeatSpeed float32
	LpfCutoffSweep float32
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

// API Functions

// Clear the screen.
//go:wasmimport null0 clear
func Clear(color uint32)

// Clear an image.
//go:wasmimport null0 clear_image
func ClearImage(color uint32, destination uint32)

// Get system-time (ms) since unix epoch.
//go:wasmimport null0 current_time
func CurrentTime() uint64

// Get the change in time (seconds) since the last update run.
//go:wasmimport null0 delta_time
func DeltaTime() float32

// Draw a filled arc on the screen.
//go:wasmimport null0 draw_arc
func DrawArc(centerX int32, centerY int32, color uint32, endAngle float32, radius float32, segments int32, startAngle float32)

// Draw a outlined (with thickness) arc on the screen.
//go:wasmimport null0 draw_arc_outline
func DrawArcOutline(centerX int32, centerY int32, color uint32, endAngle float32, radius float32, segments int32, startAngle float32, thickness int32)

// Draw a filled circle on the screen.
//go:wasmimport null0 draw_circle
func DrawCircle(centerX int32, centerY int32, color uint32, radius int32)

// Draw a circle on an image.
//go:wasmimport null0 draw_circle_on_image
func DrawCircleOnImage(centerX int32, centerY int32, color uint32, destination uint32, radius int32)

// Draw a outlined (with thickness) circle on the screen.
//go:wasmimport null0 draw_circle_outline
func DrawCircleOutline(centerX int32, centerY int32, color uint32, radius int32, thickness int32)

// Draw a outlined (with thickness) circle on an image.
//go:wasmimport null0 draw_circle_outline_on_image
func DrawCircleOutlineOnImage(centerX int32, centerY int32, color uint32, destination uint32, radius int32, thickness int32)

// Draw a filled ellipse on the screen.
//go:wasmimport null0 draw_ellipse
func DrawEllipse(centerX int32, centerY int32, color uint32, radiusX int32, radiusY int32)

// Draw a filled ellipse on an image.
//go:wasmimport null0 draw_ellipse_on_image
func DrawEllipseOnImage(centerX int32, centerY int32, color uint32, destination uint32, radiusX int32, radiusY int32)

// Draw a outlined (with thickness) ellipse on the screen.
//go:wasmimport null0 draw_ellipse_outline
func DrawEllipseOutline(centerX int32, centerY int32, color uint32, radiusX int32, radiusY int32, thickness int32)

// Draw a outlined (with thickness) ellipse on an image.
//go:wasmimport null0 draw_ellipse_outline_on_image
func DrawEllipseOutlineOnImage(centerX int32, centerY int32, color uint32, destination uint32, radiusX int32, radiusY int32, thickness int32)

// Draw an image on the screen.
//go:wasmimport null0 draw_image
func DrawImage(posX int32, posY int32, src uint32)

// Draw an image, flipped, on the screen.
//go:wasmimport null0 draw_image_flipped
func DrawImageFlipped(flipDiagonal bool, flipHorizontal bool, flipVertical bool, posX int32, posY int32, src uint32)

// Draw an image, flipped, on an image.
//go:wasmimport null0 draw_image_flipped_on_image
func DrawImageFlippedOnImage(destination uint32, flipDiagonal bool, flipHorizontal bool, flipVertical bool, posX int32, posY int32, src uint32)

// Draw an image on an image.
//go:wasmimport null0 draw_image_on_image
func DrawImageOnImage(destination uint32, posX int32, posY int32, src uint32)

// Draw an image, rotated, on the screen.
//go:wasmimport null0 draw_image_rotated
func DrawImageRotated(degrees float32, filter uint32, offsetX float32, offsetY float32, posX int32, posY int32, src uint32)

// Draw an image, rotated, on an image.
//go:wasmimport null0 draw_image_rotated_on_image
func DrawImageRotatedOnImage(degrees float32, destination uint32, filter uint32, offsetX float32, offsetY float32, posX int32, posY int32, src uint32)

// Draw an image, scaled, on the screen.
//go:wasmimport null0 draw_image_scaled
func DrawImageScaled(filter uint32, offsetX float32, offsetY float32, posX int32, posY int32, scaleX float32, scaleY float32, src uint32)

// Draw an image, scaled, on an image.
//go:wasmimport null0 draw_image_scaled_on_image
func DrawImageScaledOnImage(destination uint32, filter uint32, offsetX float32, offsetY float32, posX int32, posY int32, scaleX float32, scaleY float32, src uint32)

// Draw a tinted image on the screen.
//go:wasmimport null0 draw_image_tint
func DrawImageTint(posX int32, posY int32, src uint32, tint uint32)

// Draw a tinted image on an image.
//go:wasmimport null0 draw_image_tint_on_image
func DrawImageTintOnImage(destination uint32, posX int32, posY int32, src uint32, tint uint32)

// Draw a line on the screen.
//go:wasmimport null0 draw_line
func DrawLine(color uint32, endPosX int32, endPosY int32, startPosX int32, startPosY int32)

// Draw a line on an image.
//go:wasmimport null0 draw_line_on_image
func DrawLineOnImage(color uint32, destination uint32, endPosX int32, endPosY int32, startPosX int32, startPosY int32)

// Draw a single pixel on the screen.
//go:wasmimport null0 draw_point
func DrawPoint(color uint32, x int32, y int32)

// Draw a single pixel on an image.
//go:wasmimport null0 draw_point_on_image
func DrawPointOnImage(color uint32, destination uint32, x int32, y int32)

// Draw a filled polygon on the screen.
//go:wasmimport null0 draw_polygon
func DrawPolygon(color uint32, numPoints int32, points uint32)

// Draw a filled polygon on an image.
//go:wasmimport null0 draw_polygon_on_image
func DrawPolygonOnImage(color uint32, destination uint32, numPoints int32, points uint32)

// Draw a outlined (with thickness) polygon on the screen.
//go:wasmimport null0 draw_polygon_outline
func DrawPolygonOutline(color uint32, numPoints int32, points uint32, thickness int32)

// Draw a outlined (with thickness) polygon on an image.
//go:wasmimport null0 draw_polygon_outline_on_image
func DrawPolygonOutlineOnImage(color uint32, destination uint32, numPoints int32, points uint32, thickness int32)

// Draw a filled rectangle on the screen.
//go:wasmimport null0 draw_rectangle
func DrawRectangle(color uint32, height int32, posX int32, posY int32, width int32)

// Draw a filled rectangle on an image.
//go:wasmimport null0 draw_rectangle_on_image
func DrawRectangleOnImage(color uint32, destination uint32, height int32, posX int32, posY int32, width int32)

// Draw a outlined (with thickness) rectangle on the screen.
//go:wasmimport null0 draw_rectangle_outline
func DrawRectangleOutline(color uint32, height int32, posX int32, posY int32, thickness int32, width int32)

// Draw a outlined (with thickness) rectangle on an image.
//go:wasmimport null0 draw_rectangle_outline_on_image
func DrawRectangleOutlineOnImage(color uint32, destination uint32, height int32, posX int32, posY int32, thickness int32, width int32)

// Draw a filled round-rectangle on the screen.
//go:wasmimport null0 draw_rectangle_rounded
func DrawRectangleRounded(color uint32, cornerRadius int32, height int32, width int32, x int32, y int32)

// Draw a filled round-rectangle on an image.
//go:wasmimport null0 draw_rectangle_rounded_on_image
func DrawRectangleRoundedOnImage(color uint32, cornerRadius int32, destination uint32, height int32, width int32, x int32, y int32)

// Draw a outlined (with thickness) round-rectangle on the screen.
//go:wasmimport null0 draw_rectangle_rounded_outline
func DrawRectangleRoundedOutline(color uint32, cornerRadius int32, height int32, thickness int32, width int32, x int32, y int32)

// Draw a outlined (with thickness) round-rectangle on an image.
//go:wasmimport null0 draw_rectangle_rounded_outline_on_image
func DrawRectangleRoundedOutlineOnImage(color uint32, cornerRadius int32, destination uint32, height int32, thickness int32, width int32, x int32, y int32)

// Draw some text on the screen.
//go:wasmimport null0 draw_text
func DrawText(color uint32, font uint32, posX int32, posY int32, text uint32)

// Draw some text on an image.
//go:wasmimport null0 draw_text_on_image
func DrawTextOnImage(color uint32, destination uint32, font uint32, posX int32, posY int32, text uint32)

// Draw a filled triangle on the screen.
//go:wasmimport null0 draw_triangle
func DrawTriangle(color uint32, x1 int32, x2 int32, x3 int32, y1 int32, y2 int32, y3 int32)

// Draw a filled triangle on an image.
//go:wasmimport null0 draw_triangle_on_image
func DrawTriangleOnImage(color uint32, destination uint32, x1 int32, x2 int32, x3 int32, y1 int32, y2 int32, y3 int32)

// Draw a outlined (with thickness) triangle on the screen.
//go:wasmimport null0 draw_triangle_outline
func DrawTriangleOutline(color uint32, thickness int32, x1 int32, x2 int32, x3 int32, y1 int32, y2 int32, y3 int32)

// Draw a outlined (with thickness) triangle on an image.
//go:wasmimport null0 draw_triangle_outline_on_image
func DrawTriangleOutlineOnImage(color uint32, destination uint32, thickness int32, x1 int32, x2 int32, x3 int32, y1 int32, y2 int32, y3 int32)

// Copy a font to a new font.
//go:wasmimport null0 font_copy
func FontCopy(font uint32) uint32

// Scale a font, return a new font.
//go:wasmimport null0 font_scale
func FontScale(filter uint32, font uint32, scaleX float32, scaleY float32) uint32

// Is the button currently down?
//go:wasmimport null0 gamepad_button_down
func GamepadButtonDown(button uint32, gamepad int32) bool

// Has the button been pressed? (tracks unpress/read correctly.)
//go:wasmimport null0 gamepad_button_pressed
func GamepadButtonPressed(button uint32, gamepad int32) bool

// Has the button been released? (tracks press/read correctly.)
//go:wasmimport null0 gamepad_button_released
func GamepadButtonReleased(button uint32, gamepad int32) bool

// Calculate a rectangle representing the available alpha border in an image.
//go:wasmimport null0 image_alpha_border
func ImageAlphaBorder(image uint32, threshold float32) uint32

// Crop an image based on the alpha border, in-place.
//go:wasmimport null0 image_alpha_crop
func ImageAlphaCrop(image uint32, threshold float32)

// Use an image as an alpha-mask on another image.
//go:wasmimport null0 image_alpha_mask
func ImageAlphaMask(alphaMask uint32, image uint32, posX int32, posY int32)

// Adjust the brightness of an image, in-place.
//go:wasmimport null0 image_color_brightness
func ImageColorBrightness(factor float32, image uint32)

// Change the contrast of an image, in-place.
//go:wasmimport null0 image_color_contrast
func ImageColorContrast(contrast float32, image uint32)

// Fade a color in an image, in-place.
//go:wasmimport null0 image_color_fade
func ImageColorFade(alpha float32, image uint32)

// Invert the colors in an image, in-place.
//go:wasmimport null0 image_color_invert
func ImageColorInvert(image uint32)

// Replace a color in an image, in-place.
//go:wasmimport null0 image_color_replace
func ImageColorReplace(color uint32, image uint32, replace uint32)

// Tint a color in an image, in-place.
//go:wasmimport null0 image_color_tint
func ImageColorTint(color uint32, image uint32)

// Copy an image to a new image.
//go:wasmimport null0 image_copy
func ImageCopy(image uint32) uint32

// Crop an image, in-place.
//go:wasmimport null0 image_crop
func ImageCrop(height int32, image uint32, width int32, x int32, y int32)

// Flip an image, in-place.
//go:wasmimport null0 image_flip
func ImageFlip(horizontal bool, image uint32, vertical bool)

// Create a new image of a gradient.
//go:wasmimport null0 image_gradient
func ImageGradient(bottomLeft uint32, bottomRight uint32, height int32, topLeft uint32, topRight uint32, width int32) uint32

// Resize an image, return copy.
//go:wasmimport null0 image_resize
func ImageResize(filter uint32, image uint32, newHeight int32, newWidth int32) uint32

// Create a new image, rotating another image.
//go:wasmimport null0 image_rotate
func ImageRotate(degrees float32, filter uint32, image uint32) uint32

// Scale an image, return copy.
//go:wasmimport null0 image_scale
func ImageScale(filter uint32, image uint32, scaleX float32, scaleY float32) uint32

// Create an image from a region of another image.
//go:wasmimport null0 image_subimage
func ImageSubimage(height int32, image uint32, width int32, x int32, y int32) uint32

// Is the key currently down?
//go:wasmimport null0 key_down
func KeyDown(key uint32) bool

// Has the key been pressed? (tracks unpress/read correctly.)
//go:wasmimport null0 key_pressed
func KeyPressed(key uint32) bool

// Has the key been released? (tracks press/read correctly.)
//go:wasmimport null0 key_released
func KeyReleased(key uint32) bool

// Is the key currently up?
//go:wasmimport null0 key_up
func KeyUp(key uint32) bool

// Load a BMF font from a file in cart.
//go:wasmimport null0 load_font_bmf
func LoadFontBmf(characters uint32, filename uint32) uint32

// Load a BMF font from an image.
//go:wasmimport null0 load_font_bmf_from_image
func LoadFontBmfFromImage(characters uint32, image uint32) uint32

// Load a TTF font from a file in cart.
//go:wasmimport null0 load_font_ttf
func LoadFontTtf(filename uint32, fontSize int32) uint32

// Load a TTY font from a file in cart.
//go:wasmimport null0 load_font_tty
func LoadFontTty(characters uint32, filename uint32, glyphHeight int32, glyphWidth int32) uint32

// Load a TTY font from an image.
//go:wasmimport null0 load_font_tty_from_image
func LoadFontTtyFromImage(characters uint32, glyphHeight int32, glyphWidth int32, image uint32) uint32

// Load an image from a file in cart.
//go:wasmimport null0 load_image
func LoadImage(filename uint32) uint32

// Load a sound from a file in cart.
//go:wasmimport null0 load_sound
func LoadSound(filename uint32) uint32

// Meaure an image (use 0 for screen).
//go:wasmimport null0 measure_image
func MeasureImage(image uint32) uint32

// Measure the size of some text.
//go:wasmimport null0 measure_text
func MeasureText(font uint32, text uint32, textLength int32) uint32

// Is the button currently down?
//go:wasmimport null0 mouse_button_down
func MouseButtonDown(button uint32) bool

// Has the button been pressed? (tracks unpress/read correctly.)
//go:wasmimport null0 mouse_button_pressed
func MouseButtonPressed(button uint32) bool

// Has the button been released? (tracks press/read correctly.)
//go:wasmimport null0 mouse_button_released
func MouseButtonReleased(button uint32) bool

// Is the button currently up?
//go:wasmimport null0 mouse_button_up
func MouseButtonUp(button uint32) bool

// Get current position of mouse.
//go:wasmimport null0 mouse_position
func MousePosition() uint32

// Create a new blank image.
//go:wasmimport null0 new_image
func NewImage(color uint32, height int32, width int32) uint32

// Play a sound.
//go:wasmimport null0 play_sound
func PlaySound(loop bool, sound uint32)

// Get a random integer between 2 numbers.
//go:wasmimport null0 random_int
func RandomInt(max int32, min int32) int32

// Get the random-seed.
//go:wasmimport null0 random_seed_get
func RandomSeedGet() uint64

// Set the random-seed.
//go:wasmimport null0 random_seed_set
func RandomSeedSet(seed uint64)

// Save an image to persistant storage.
//go:wasmimport null0 save_image
func SaveImage(filename uint32, image uint32)

// Create Sfx parameters.
//go:wasmimport null0 sfx_generate
func SfxGenerate(type_ uint32) uint32

// Create Sfx sound.
//go:wasmimport null0 sfx_sound
func SfxSound(params uint32) uint32

// Stop a sound.
//go:wasmimport null0 stop_sound
func StopSound(sound uint32)

// Speak some text and return a sound. Set things to 0 for defaults.
//go:wasmimport null0 tts_sound
func TtsSound(mouth int32, phonetic bool, pitch int32, sing bool, speed int32, text uint32, throat int32) uint32

// Unload a font.
//go:wasmimport null0 unload_font
func UnloadFont(font uint32)

// Unload an image.
//go:wasmimport null0 unload_image
func UnloadImage(image uint32)

// Unload a sound.
//go:wasmimport null0 unload_sound
func UnloadSound(sound uint32)

