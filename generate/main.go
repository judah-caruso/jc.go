package main

import (
	"fmt"
	"strings"

	"github.com/dave/jennifer/jen"
)

func main() {
	// mg
	must(generateVectorType(2).Save("v2.go"))
	must(generateVectorType(3).Save("v3.go"))
	must(generateVectorType(4).Save("v4.go"))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func generateVectorType(dimensions int) *jen.File {
	f := jen.NewFile("mg")

	const Float = "Float"

	type_name := fmt.Sprintf("V%d", dimensions)
	f.Commentf("%s is a %d-dimensional vector type", type_name, dimensions)
	f.Type().Id(type_name).Index(jen.Lit(dimensions)).Id(Float)

	fields_by_dimension := [][]string{
		nil, nil,
		// Vec2
		{"X", "Y", "U", "V", "Min", "Max", "Width", "Height"},
		// Vec3
		{"X", "Y", "Z", "R", "G", "B", "Width", "Height", "Depth"},
		// Vec4
		{"X", "Y", "Z", "W", "R", "G", "B", "A", "", "", "Width", "Height"},
	}

	operator_templates := []struct {
		name     string
		operator string
	}{
		{"Add", "+"},
		{"Sub", "-"},
		{"Mul", "*"},
		{"Div", "/"},
	}

	// func (v *VX|VX) func_name
	startMethod := func(func_name string, ptr bool) *jen.Statement {
		var receiver *jen.Statement
		if ptr {
			receiver = jen.Id("v").Op("*").Id(type_name)
		} else {
			receiver = jen.Id("v").Id(type_name)
		}

		return f.Func().Params(receiver).Id(func_name)
	}

	switch dimensions {
	case 3: // Vec3
		// AsV2 func
		asv2 := startMethod("AsV2", false).Params().Id("V2")
		asv2.Op("{").Return(jen.Id("V2").ValuesFunc(func(g *jen.Group) {
			for i := range dimensions - 1 {
				idx := jen.Lit(i)
				g.Id("v").Index(idx)
			}
		}).Op("}"))

		// AsV4 func
		asv4 := startMethod("AsV4", false).Params().Id("V4")
		asv4.Op("{").Return(jen.Id("V4").ValuesFunc(func(g *jen.Group) {
			for i := range dimensions {
				idx := jen.Lit(i)
				g.Id("v").Index(idx)
			}

			g.Lit(1.0)
		}).Op("}"))

		// AsV4w func
		asv4w := startMethod("AsV4w", false).Params(jen.Id("w").Id(Float)).Id("V4")
		asv4w.Op("{").Return(jen.Id("V4").ValuesFunc(func(g *jen.Group) {
			for i := range dimensions {
				idx := jen.Lit(i)
				g.Id("v").Index(idx)
			}

			g.Id("w")
		}).Op("}"))

		f.Line()
	}

	// Generate field accessors
	fields := fields_by_dimension[dimensions]
	for len(fields) != 0 {
		group := fields[:dimensions]
		fields = fields[dimensions:]

		for i, g := range group {
			if len(g) == 0 {
				continue
			}

			fn := startMethod(g, false).Params().Id(Float)
			fn.Op("{").Return(jen.Id("v").Index(jen.Lit(i))).Op("}")
		}
	}

	f.Line()

	// Generate pointer field accessors
	fields = fields_by_dimension[dimensions]
	for len(fields) != 0 {
		group := fields[:dimensions]
		fields = fields[dimensions:]

		for i, g := range group {
			if len(g) == 0 {
				continue
			}

			fn := startMethod(g+"p", true).Params().Op("*").Id(Float)
			fn.Op("{").Return(jen.Op("&").Id("v").Index(jen.Lit(i))).Op("}")
		}
	}

	f.Line()

	// Generate access funcs
	fields = fields_by_dimension[dimensions]
	for len(fields) != 0 {
		group := fields[:dimensions]
		fields = fields[dimensions:]

		for i, g := range group {
			if len(g) == 0 {
				continue
			}

			fn := startMethod("Set"+g, true).Params(jen.Id("f").Id(Float))
			fn.Op("{").Id("v").Index(jen.Lit(i)).Op("=").Id("f").Op("}")
		}
	}

	f.Line()

	// Basic component-wise operator funcs
	for _, operator := range operator_templates {
		args := jen.Id("o").Id(type_name)

		basic := startMethod(operator.name, false).Params(args).Id(type_name)
		basic.BlockFunc(func(g *jen.Group) {
			g.Return(jen.Id(type_name).ValuesFunc(func(g *jen.Group) {
				for i := range dimensions {
					idx := jen.Lit(i)
					g.Id("v").Index(idx).Op(operator.operator).Id("o").Index(idx)
				}
			}))
		}).Line()

		scalar := startMethod(fmt.Sprintf("%sf", operator.name), false).Params(jen.Id("f").Id(Float)).Id(type_name)
		scalar.BlockFunc(func(g *jen.Group) {
			g.Return(jen.Id(type_name).ValuesFunc(func(g *jen.Group) {
				for i := range dimensions {
					g.Id("v").Index(jen.Lit(i)).Op(operator.operator).Id("f")
				}
			}))
		}).Line()

		assign := startMethod(operator.name+"Mut", true).Params(args)
		assign.BlockFunc(func(g *jen.Group) {
			op := operator.operator + "="
			for i := range dimensions {
				idx := jen.Lit(i)
				g.Id("v").Index(idx).Op(op).Id("o").Index(idx)
			}
		}).Line()

		scalar_assign := startMethod(operator.name+"Mutf", true).Params(jen.Id("f").Id(Float))
		scalar_assign.BlockFunc(func(g *jen.Group) {
			op := operator.operator + "="
			for i := range dimensions {
				g.Id("v").Index(jen.Lit(i)).Op(op).Id("f")
			}
		}).Line()
	}

	// Eq func
	startMethod("Eq", false).Params(jen.Id("o").Id(type_name)).Bool().BlockFunc(func(g *jen.Group) {
		g.ReturnFunc(func(g *jen.Group) {
			expr := jen.Empty()
			for i := range dimensions {
				idx := jen.Lit(i)
				expr.Id("Abs").Call(jen.Id("v").Index(idx).Op("-").Id("o").Index(idx)).Op("<").Id("FloatMin")
				if i < dimensions-1 {
					expr.Op("&&")
				}
			}
			g.Add(expr)
		})
	}).Line()

	// CloseEnough func
	startMethod("CloseEnough", false).Params(jen.Id("o").Id(type_name)).Bool().BlockFunc(func(g *jen.Group) {
		g.ReturnFunc(func(g *jen.Group) {
			expr := jen.Empty()
			for i := range dimensions {
				idx := jen.Lit(i)
				expr.Id("CloseEnough").Call(jen.Id("v").Index(idx), jen.Id("o").Index(idx))
				if i < dimensions-1 {
					expr.Op("&&")
				}
			}
			g.Add(expr)
		})
	}).Line()

	// Dot func
	startMethod("Dot", false).Params(jen.Id("o").Id(type_name)).Id(Float).BlockFunc(func(g *jen.Group) {
		g.ReturnFunc(func(g *jen.Group) {
			expr := jen.Empty()
			for i := range dimensions {
				idx := jen.Lit(i)
				expr.Id("v").Index(idx).Op("*").Id("o").Index(idx)
				if i < dimensions-1 {
					expr.Op("+")
				}
			}
			g.Add(expr)
		})
	}).Line()

	// Mag func
	startMethod("Mag", false).Params().Id(Float).BlockFunc(func(g *jen.Group) {
		g.Return(jen.Id("Sqrt").Call(jen.Id("v").Op(".").Id("Dot").Call(jen.Id("v"))))
	}).Line()

	// Distance func
	startMethod("Distance", false).Params(jen.Id("to").Id(type_name)).Id(Float).BlockFunc(func(g *jen.Group) {
		g.Return(jen.Id("v").Op(".").Id("Sub").Call(jen.Id("to")).Op(".").Id("Mag").Call())
	}).Line()

	// Angle func
	startMethod("Angle", false).Params(jen.Id("to").Id(type_name)).Id("Angle").BlockFunc(func(g *jen.Group) {
		vmag := jen.Id("v").Dot("Mag").Call()
		tmag := jen.Id("to").Dot("Mag").Call()
		magmul := jen.Parens(vmag.Op("*").Add(tmag))
		dot := jen.Id("v").Dot("Dot").Call(jen.Id("to"))
		g.Return(jen.Id("FromRad").Call(jen.Id("Acos").Call(dot.Op("/").Add(magmul))))
	}).Line()

	// Abs func
	startMethod("Abs", false).Params().Id(type_name).BlockFunc(func(g *jen.Group) {
		g.Return(jen.Id(type_name).ValuesFunc(func(g *jen.Group) {
			for i := range dimensions {
				g.Id("Abs").Call(jen.Id("v").Index(jen.Lit(i)))
			}
		}))
	}).Line()

	// @todo(judah): Normalize should return zero vec when mag is zero?

	// Normal func
	startMethod("Normalize", false).Params().Id(type_name).BlockFunc(func(g *jen.Group) {
		g.Return(jen.Id("v").Op(".").Id("Mulf").Call(jen.Lit(1.0).Op("/").Id("v").Op(".").Id("Mag").Call()))
	}).Line()

	// NormalizeMut func
	startMethod("NormalizeMut", true).Params().BlockFunc(func(g *jen.Group) {
		g.Id("v").Op(".").Id("MulMutf").Call(jen.Lit(1.0).Op("/").Id("v").Op(".").Id("Mag").Call())
	}).Line()

	// Invert func
	startMethod("Invert", false).Params().Id(type_name).BlockFunc(func(g *jen.Group) {
		g.Return(jen.Id(type_name).ValuesFunc(func(g *jen.Group) {
			for i := range dimensions {
				g.Op("-").Id("v").Index(jen.Lit(i))
			}
		}))
	}).Line()

	// InvertMut func
	startMethod("InvertMut", true).Params().BlockFunc(func(g *jen.Group) {
		for i := range dimensions {
			idx := jen.Lit(i)
			g.Id("v").Index(idx).Op("=").Op("-").Id("v").Index(idx)
		}
	}).Line()

	// Clamp func
	startMethod("Clamp", false).Params(jen.Id("min").Id(Float), jen.Id("max").Id(Float)).Id(type_name).BlockFunc(func(g *jen.Group) {
		g.Return(jen.Id(type_name).ValuesFunc(func(g *jen.Group) {
			for i := range dimensions {
				g.Id("Clamp").Call(jen.Id("v").Index(jen.Lit(i)), jen.Id("min"), jen.Id("max"))
			}
		}))
	}).Line()

	// ClampMut func
	startMethod("ClampMut", true).Params(jen.Id("min").Id(Float), jen.Id("max").Id(Float)).BlockFunc(func(g *jen.Group) {
		for i := range dimensions {
			idx := jen.Lit(i)
			g.Id("v").Index(idx).Op("=").Id("Clamp").Call(jen.Id("v").Index(idx), jen.Id("min"), jen.Id("max"))
		}
	}).Line()

	// Lerp func
	lerp := startMethod("Lerp", false).Params(jen.Id("to").Id(type_name), jen.Id("t").Id(Float)).Id(type_name)
	lerp.BlockFunc(func(g *jen.Group) {
		g.Return(jen.Id(type_name).ValuesFunc(func(g *jen.Group) {
			for i := range dimensions {
				idx := jen.Lit(i)
				g.Id("Lerp").Call(jen.Id("v").Index(idx), jen.Id("to").Index(idx), jen.Id("t"))
			}
		}))
	}).Line()

	// LerpMut func
	lerpmut := startMethod("LerpMut", true).Params(jen.Id("to").Id(type_name), jen.Id("t").Id(Float))
	lerpmut.BlockFunc(func(g *jen.Group) {
		for i := range dimensions {
			idx := jen.Lit(i)
			g.Id("v").Index(idx).Op("=").Id("Lerp").Call(jen.Id("v").Index(idx), jen.Id("to").Index(idx), jen.Id("t"))
		}
	}).Line()

	// Generate extra funcs (specific type # dimensions)
	switch dimensions {
	case 2: // Vec2
		f.Op(`
			func (v V2) Rotate(angle Angle) V2 {
				cos := Float(angle.Cos())
				sin := Float(angle.Sin())
				return V2{
					v[0] * cos - v[1] * sin,
					v[0] * sin + v[1] * cos,
				}
			}

			func (v *V2) RotateMut(angle Angle){
				cos := Float(angle.Cos())
				sin := Float(angle.Sin())
				v[0] = v[0] * cos - v[1] * sin
				v[1] = v[0] * sin + v[1] * cos
			}

			func (v V2) Reflect(normal V2) V2 {
				dot := v.Dot(normal)
				return V2{
					v[0] - (2 * normal[0]) * dot,
					v[1] - (2 * normal[1]) * dot,
				}
			}

			func (v *V2) ReflectMut(normal V2) {
				dot := v.Dot(normal)
				v[0] = v[0] - (2 * normal[0]) * dot
				v[1] = v[1] - (2 * normal[1]) * dot
			}
		`).Line()
	case 3: // Vec3
		f.Op(`
			func (v V3) Cross(u V3) V3 {
				return V3{
					v[1]*u[2] - v[2]*u[1],
					v[2]*u[0] - v[0]*u[2],
					v[0]*u[1] - v[1]*u[0],
				}
			}

			// Implements the [color.Color] interface
			func (v V3) RGBA() (r, g, b, a uint32) {
				r = uint32(v[0] * 0xFFFF)
				g = uint32(v[1] * 0xFFFF)
				b = uint32(v[2] * 0xFFFF)
				a = 0xFFFF
				return
			}
		`).Line()
	case 4: // Vec4
		f.Op(`
			// Implements the [color.Color] interface
			func (v V4) RGBA() (r, g, b, a uint32) {
				r = uint32(v[0] * 0xFFFF)
				g = uint32(v[1] * 0xFFFF)
				b = uint32(v[2] * 0xFFFF)
				a = uint32(v[3] * 0xFFFF)
				return
			}
		`).Line()
	}

	// String func
	f.Comment("Implements the [fmt.Stringer] interface")
	startMethod("String", false).Params().String().BlockFunc(func(g *jen.Group) {
		format_str := strings.Repeat("%f, ", dimensions)
		format_str = format_str[:len(format_str)-2]
		args := []jen.Code{
			jen.Lit("(" + format_str + ")"),
		}
		for i := range dimensions {
			args = append(args, jen.Id("v").Index(jen.Lit(i)))
		}
		g.Return(jen.Qual("fmt", "Sprintf").Call(args...))
	}).Line()

	return f
}
