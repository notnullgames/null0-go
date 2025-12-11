package main

// This file is auto-generated from api/*.yml - DO NOT EDIT

import (
	"context"
	"fmt"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
)

// registerNull0API registers all null0 API functions with the host module builder.
func registerNull0API(builder wazero.HostModuleBuilder) {
	// clear: Clear the screen.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] clear called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32},
			nil,
		).
		Export("clear")

	// clear_image: Clear an image.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] clear_image called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32},
			nil,
		).
		Export("clear_image")

	// current_time: Get system-time (ms) since unix epoch.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] current_time called (stub)")
			}),
			nil,
			[]api.ValueType{api.ValueTypeI64},
		).
		Export("current_time")

	// delta_time: Get the change in time (seconds) since the last update run.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] delta_time called (stub)")
			}),
			nil,
			[]api.ValueType{api.ValueTypeF32},
		).
		Export("delta_time")

	// draw_arc: Draw a filled arc on the screen.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] draw_arc called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeF32, api.ValueTypeF32, api.ValueTypeI32, api.ValueTypeF32},
			nil,
		).
		Export("draw_arc")

	// draw_arc_outline: Draw a outlined (with thickness) arc on the screen.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] draw_arc_outline called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeF32, api.ValueTypeF32, api.ValueTypeI32, api.ValueTypeF32, api.ValueTypeI32},
			nil,
		).
		Export("draw_arc_outline")

	// draw_circle: Draw a filled circle on the screen.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] draw_circle called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32},
			nil,
		).
		Export("draw_circle")

	// draw_circle_on_image: Draw a circle on an image.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] draw_circle_on_image called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32},
			nil,
		).
		Export("draw_circle_on_image")

	// draw_circle_outline: Draw a outlined (with thickness) circle on the screen.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] draw_circle_outline called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32},
			nil,
		).
		Export("draw_circle_outline")

	// draw_circle_outline_on_image: Draw a outlined (with thickness) circle on an image.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] draw_circle_outline_on_image called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32},
			nil,
		).
		Export("draw_circle_outline_on_image")

	// draw_ellipse: Draw a filled ellipse on the screen.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] draw_ellipse called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32},
			nil,
		).
		Export("draw_ellipse")

	// draw_ellipse_on_image: Draw a filled ellipse on an image.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] draw_ellipse_on_image called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32},
			nil,
		).
		Export("draw_ellipse_on_image")

	// draw_ellipse_outline: Draw a outlined (with thickness) ellipse on the screen.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] draw_ellipse_outline called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32},
			nil,
		).
		Export("draw_ellipse_outline")

	// draw_ellipse_outline_on_image: Draw a outlined (with thickness) ellipse on an image.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] draw_ellipse_outline_on_image called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32},
			nil,
		).
		Export("draw_ellipse_outline_on_image")

	// draw_image: Draw an image on the screen.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] draw_image called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32},
			nil,
		).
		Export("draw_image")

	// draw_image_flipped: Draw an image, flipped, on the screen.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] draw_image_flipped called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32},
			nil,
		).
		Export("draw_image_flipped")

	// draw_image_flipped_on_image: Draw an image, flipped, on an image.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] draw_image_flipped_on_image called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32},
			nil,
		).
		Export("draw_image_flipped_on_image")

	// draw_image_on_image: Draw an image on an image.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] draw_image_on_image called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32},
			nil,
		).
		Export("draw_image_on_image")

	// draw_image_rotated: Draw an image, rotated, on the screen.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] draw_image_rotated called (stub)")
			}),
			[]api.ValueType{api.ValueTypeF32, api.ValueTypeI32, api.ValueTypeF32, api.ValueTypeF32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32},
			nil,
		).
		Export("draw_image_rotated")

	// draw_image_rotated_on_image: Draw an image, rotated, on an image.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] draw_image_rotated_on_image called (stub)")
			}),
			[]api.ValueType{api.ValueTypeF32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeF32, api.ValueTypeF32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32},
			nil,
		).
		Export("draw_image_rotated_on_image")

	// draw_image_scaled: Draw an image, scaled, on the screen.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] draw_image_scaled called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeF32, api.ValueTypeF32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeF32, api.ValueTypeF32, api.ValueTypeI32},
			nil,
		).
		Export("draw_image_scaled")

	// draw_image_scaled_on_image: Draw an image, scaled, on an image.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] draw_image_scaled_on_image called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeF32, api.ValueTypeF32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeF32, api.ValueTypeF32, api.ValueTypeI32},
			nil,
		).
		Export("draw_image_scaled_on_image")

	// draw_image_tint: Draw a tinted image on the screen.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] draw_image_tint called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32},
			nil,
		).
		Export("draw_image_tint")

	// draw_image_tint_on_image: Draw a tinted image on an image.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] draw_image_tint_on_image called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32},
			nil,
		).
		Export("draw_image_tint_on_image")

	// draw_line: Draw a line on the screen.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] draw_line called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32},
			nil,
		).
		Export("draw_line")

	// draw_line_on_image: Draw a line on an image.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] draw_line_on_image called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32},
			nil,
		).
		Export("draw_line_on_image")

	// draw_point: Draw a single pixel on the screen.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] draw_point called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32},
			nil,
		).
		Export("draw_point")

	// draw_point_on_image: Draw a single pixel on an image.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] draw_point_on_image called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32},
			nil,
		).
		Export("draw_point_on_image")

	// draw_polygon: Draw a filled polygon on the screen.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] draw_polygon called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32},
			nil,
		).
		Export("draw_polygon")

	// draw_polygon_on_image: Draw a filled polygon on an image.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] draw_polygon_on_image called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32},
			nil,
		).
		Export("draw_polygon_on_image")

	// draw_polygon_outline: Draw a outlined (with thickness) polygon on the screen.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] draw_polygon_outline called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32},
			nil,
		).
		Export("draw_polygon_outline")

	// draw_polygon_outline_on_image: Draw a outlined (with thickness) polygon on an image.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] draw_polygon_outline_on_image called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32},
			nil,
		).
		Export("draw_polygon_outline_on_image")

	// draw_rectangle: Draw a filled rectangle on the screen.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] draw_rectangle called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32},
			nil,
		).
		Export("draw_rectangle")

	// draw_rectangle_on_image: Draw a filled rectangle on an image.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] draw_rectangle_on_image called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32},
			nil,
		).
		Export("draw_rectangle_on_image")

	// draw_rectangle_outline: Draw a outlined (with thickness) rectangle on the screen.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] draw_rectangle_outline called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32},
			nil,
		).
		Export("draw_rectangle_outline")

	// draw_rectangle_outline_on_image: Draw a outlined (with thickness) rectangle on an image.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] draw_rectangle_outline_on_image called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32},
			nil,
		).
		Export("draw_rectangle_outline_on_image")

	// draw_rectangle_rounded: Draw a filled round-rectangle on the screen.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] draw_rectangle_rounded called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32},
			nil,
		).
		Export("draw_rectangle_rounded")

	// draw_rectangle_rounded_on_image: Draw a filled round-rectangle on an image.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] draw_rectangle_rounded_on_image called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32},
			nil,
		).
		Export("draw_rectangle_rounded_on_image")

	// draw_rectangle_rounded_outline: Draw a outlined (with thickness) round-rectangle on the screen.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] draw_rectangle_rounded_outline called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32},
			nil,
		).
		Export("draw_rectangle_rounded_outline")

	// draw_rectangle_rounded_outline_on_image: Draw a outlined (with thickness) round-rectangle on an image.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] draw_rectangle_rounded_outline_on_image called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32},
			nil,
		).
		Export("draw_rectangle_rounded_outline_on_image")

	// draw_text: Draw some text on the screen.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] draw_text called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32},
			nil,
		).
		Export("draw_text")

	// draw_text_on_image: Draw some text on an image.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] draw_text_on_image called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32},
			nil,
		).
		Export("draw_text_on_image")

	// draw_triangle: Draw a filled triangle on the screen.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] draw_triangle called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32},
			nil,
		).
		Export("draw_triangle")

	// draw_triangle_on_image: Draw a filled triangle on an image.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] draw_triangle_on_image called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32},
			nil,
		).
		Export("draw_triangle_on_image")

	// draw_triangle_outline: Draw a outlined (with thickness) triangle on the screen.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] draw_triangle_outline called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32},
			nil,
		).
		Export("draw_triangle_outline")

	// draw_triangle_outline_on_image: Draw a outlined (with thickness) triangle on an image.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] draw_triangle_outline_on_image called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32},
			nil,
		).
		Export("draw_triangle_outline_on_image")

	// font_copy: Copy a font to a new font.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] font_copy called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32},
			[]api.ValueType{api.ValueTypeI32},
		).
		Export("font_copy")

	// font_scale: Scale a font, return a new font.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] font_scale called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeF32, api.ValueTypeF32},
			[]api.ValueType{api.ValueTypeI32},
		).
		Export("font_scale")

	// gamepad_button_down: Is the button currently down?
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] gamepad_button_down called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32},
			[]api.ValueType{api.ValueTypeI32},
		).
		Export("gamepad_button_down")

	// gamepad_button_pressed: Has the button been pressed? (tracks unpress/read correctly.)
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] gamepad_button_pressed called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32},
			[]api.ValueType{api.ValueTypeI32},
		).
		Export("gamepad_button_pressed")

	// gamepad_button_released: Has the button been released? (tracks press/read correctly.)
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] gamepad_button_released called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32},
			[]api.ValueType{api.ValueTypeI32},
		).
		Export("gamepad_button_released")

	// image_alpha_border: Calculate a rectangle representing the available alpha border in an image.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] image_alpha_border called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeF32},
			[]api.ValueType{api.ValueTypeI32},
		).
		Export("image_alpha_border")

	// image_alpha_crop: Crop an image based on the alpha border, in-place.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] image_alpha_crop called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeF32},
			nil,
		).
		Export("image_alpha_crop")

	// image_alpha_mask: Use an image as an alpha-mask on another image.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] image_alpha_mask called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32},
			nil,
		).
		Export("image_alpha_mask")

	// image_color_brightness: Adjust the brightness of an image, in-place.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] image_color_brightness called (stub)")
			}),
			[]api.ValueType{api.ValueTypeF32, api.ValueTypeI32},
			nil,
		).
		Export("image_color_brightness")

	// image_color_contrast: Change the contrast of an image, in-place.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] image_color_contrast called (stub)")
			}),
			[]api.ValueType{api.ValueTypeF32, api.ValueTypeI32},
			nil,
		).
		Export("image_color_contrast")

	// image_color_fade: Fade a color in an image, in-place.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] image_color_fade called (stub)")
			}),
			[]api.ValueType{api.ValueTypeF32, api.ValueTypeI32},
			nil,
		).
		Export("image_color_fade")

	// image_color_invert: Invert the colors in an image, in-place.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] image_color_invert called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32},
			nil,
		).
		Export("image_color_invert")

	// image_color_replace: Replace a color in an image, in-place.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] image_color_replace called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32},
			nil,
		).
		Export("image_color_replace")

	// image_color_tint: Tint a color in an image, in-place.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] image_color_tint called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32},
			nil,
		).
		Export("image_color_tint")

	// image_copy: Copy an image to a new image.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] image_copy called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32},
			[]api.ValueType{api.ValueTypeI32},
		).
		Export("image_copy")

	// image_crop: Crop an image, in-place.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] image_crop called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32},
			nil,
		).
		Export("image_crop")

	// image_flip: Flip an image, in-place.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] image_flip called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32},
			nil,
		).
		Export("image_flip")

	// image_gradient: Create a new image of a gradient.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] image_gradient called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32},
			[]api.ValueType{api.ValueTypeI32},
		).
		Export("image_gradient")

	// image_resize: Resize an image, return copy.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] image_resize called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32},
			[]api.ValueType{api.ValueTypeI32},
		).
		Export("image_resize")

	// image_rotate: Create a new image, rotating another image.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] image_rotate called (stub)")
			}),
			[]api.ValueType{api.ValueTypeF32, api.ValueTypeI32, api.ValueTypeI32},
			[]api.ValueType{api.ValueTypeI32},
		).
		Export("image_rotate")

	// image_scale: Scale an image, return copy.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] image_scale called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeF32, api.ValueTypeF32},
			[]api.ValueType{api.ValueTypeI32},
		).
		Export("image_scale")

	// image_subimage: Create an image from a region of another image.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] image_subimage called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32},
			[]api.ValueType{api.ValueTypeI32},
		).
		Export("image_subimage")

	// key_down: Is the key currently down?
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] key_down called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32},
			[]api.ValueType{api.ValueTypeI32},
		).
		Export("key_down")

	// key_pressed: Has the key been pressed? (tracks unpress/read correctly.)
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] key_pressed called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32},
			[]api.ValueType{api.ValueTypeI32},
		).
		Export("key_pressed")

	// key_released: Has the key been released? (tracks press/read correctly.)
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] key_released called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32},
			[]api.ValueType{api.ValueTypeI32},
		).
		Export("key_released")

	// key_up: Is the key currently up?
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] key_up called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32},
			[]api.ValueType{api.ValueTypeI32},
		).
		Export("key_up")

	// load_font_bmf: Load a BMF font from a file in cart.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] load_font_bmf called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32},
			[]api.ValueType{api.ValueTypeI32},
		).
		Export("load_font_bmf")

	// load_font_bmf_from_image: Load a BMF font from an image.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] load_font_bmf_from_image called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32},
			[]api.ValueType{api.ValueTypeI32},
		).
		Export("load_font_bmf_from_image")

	// load_font_ttf: Load a TTF font from a file in cart.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] load_font_ttf called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32},
			[]api.ValueType{api.ValueTypeI32},
		).
		Export("load_font_ttf")

	// load_font_tty: Load a TTY font from a file in cart.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] load_font_tty called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32},
			[]api.ValueType{api.ValueTypeI32},
		).
		Export("load_font_tty")

	// load_font_tty_from_image: Load a TTY font from an image.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] load_font_tty_from_image called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32},
			[]api.ValueType{api.ValueTypeI32},
		).
		Export("load_font_tty_from_image")

	// load_image: Load an image from a file in cart.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] load_image called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32},
			[]api.ValueType{api.ValueTypeI32},
		).
		Export("load_image")

	// load_sound: Load a sound from a file in cart.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] load_sound called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32},
			[]api.ValueType{api.ValueTypeI32},
		).
		Export("load_sound")

	// measure_image: Meaure an image (use 0 for screen).
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] measure_image called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32},
			[]api.ValueType{api.ValueTypeI32},
		).
		Export("measure_image")

	// measure_text: Measure the size of some text.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] measure_text called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32},
			[]api.ValueType{api.ValueTypeI32},
		).
		Export("measure_text")

	// mouse_button_down: Is the button currently down?
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] mouse_button_down called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32},
			[]api.ValueType{api.ValueTypeI32},
		).
		Export("mouse_button_down")

	// mouse_button_pressed: Has the button been pressed? (tracks unpress/read correctly.)
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] mouse_button_pressed called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32},
			[]api.ValueType{api.ValueTypeI32},
		).
		Export("mouse_button_pressed")

	// mouse_button_released: Has the button been released? (tracks press/read correctly.)
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] mouse_button_released called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32},
			[]api.ValueType{api.ValueTypeI32},
		).
		Export("mouse_button_released")

	// mouse_button_up: Is the button currently up?
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] mouse_button_up called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32},
			[]api.ValueType{api.ValueTypeI32},
		).
		Export("mouse_button_up")

	// mouse_position: Get current position of mouse.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] mouse_position called (stub)")
			}),
			nil,
			[]api.ValueType{api.ValueTypeI32},
		).
		Export("mouse_position")

	// new_image: Create a new blank image.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] new_image called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32},
			[]api.ValueType{api.ValueTypeI32},
		).
		Export("new_image")

	// play_sound: Play a sound.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] play_sound called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32},
			nil,
		).
		Export("play_sound")

	// random_int: Get a random integer between 2 numbers.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] random_int called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32},
			[]api.ValueType{api.ValueTypeI32},
		).
		Export("random_int")

	// random_seed_get: Get the random-seed.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] random_seed_get called (stub)")
			}),
			nil,
			[]api.ValueType{api.ValueTypeI64},
		).
		Export("random_seed_get")

	// random_seed_set: Set the random-seed.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] random_seed_set called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI64},
			nil,
		).
		Export("random_seed_set")

	// save_image: Save an image to persistant storage.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] save_image called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32},
			nil,
		).
		Export("save_image")

	// sfx_generate: Create Sfx parameters.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] sfx_generate called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32},
			[]api.ValueType{api.ValueTypeI32},
		).
		Export("sfx_generate")

	// sfx_sound: Create Sfx sound.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] sfx_sound called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32},
			[]api.ValueType{api.ValueTypeI32},
		).
		Export("sfx_sound")

	// stop_sound: Stop a sound.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] stop_sound called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32},
			nil,
		).
		Export("stop_sound")

	// tts_sound: Speak some text and return a sound. Set things to 0 for defaults.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] tts_sound called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32, api.ValueTypeI32},
			[]api.ValueType{api.ValueTypeI32},
		).
		Export("tts_sound")

	// unload_font: Unload a font.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] unload_font called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32},
			nil,
		).
		Export("unload_font")

	// unload_image: Unload an image.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] unload_image called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32},
			nil,
		).
		Export("unload_image")

	// unload_sound: Unload a sound.
	builder.NewFunctionBuilder().
		WithGoModuleFunction(
			api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
				fmt.Println("[null0] unload_sound called (stub)")
			}),
			[]api.ValueType{api.ValueTypeI32},
			nil,
		).
		Export("unload_sound")

}
