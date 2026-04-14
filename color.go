package colorschemes

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Color struct {
	r, g, b int
}

func (c Color) String() string {
	return fmt.Sprintf("#%02x%02x%02x", c.r, c.g, c.b)
}

func (c Color) MarshalJSON() ([]byte, error) {
	return []byte("\"" + c.String() + "\""), nil
}

func (c Color) Lerp(other Color, t float64) Color {
	return Color{
		r: lerp(c.r, other.r, t),
		g: lerp(c.g, other.g, t),
		b: lerp(c.b, other.b, t),
	}
}

func (c Color) Lighten(amount float64) Color {
	h, s, l := c.toHSL()
	return ColorFromHSL(h, s, min(l+amount, 1.0))
}

func (c Color) Darken(amount float64) Color {
	h, s, l := c.toHSL()
	return ColorFromHSL(h, s, max(l-amount, 0.0))
}

func (c Color) toHSL() (h, s, l float64) {
	r := float64(c.r) / 255
	g := float64(c.g) / 255
	b := float64(c.b) / 255

	cMax := max(r, g, b)
	cMin := min(r, g, b)
	delta := cMax - cMin

	l = (cMax + cMin) / 2

	if delta == 0 {
		return 0, 0, l
	}

	s = delta / (1 - math.Abs(2*l-1))

	switch cMax {
	case r:
		h = math.Mod((g-b)/delta, 6)
	case g:
		h = (b-r)/delta + 2
	case b:
		h = (r-g)/delta + 4
	}
	h /= 6
	if h < 0 {
		h += 1
	}

	return h, s, l
}

func ColorMustFromHex(hex string) Color {
	c, err := ColorFromHex(hex)
	if err != nil {
		panic(err)
	}
	return c
}

func ColorFromHex(hex string) (Color, error) {
	switch len(hex) {
	case 3:
		// hex = "fff"
		r, err := strconv.ParseInt(strings.Repeat(string(hex[0]), 2), 16, 64)
		if err != nil {
			return Color{}, err
		}
		g, err := strconv.ParseInt(strings.Repeat(string(hex[1]), 2), 16, 64)
		if err != nil {
			return Color{}, err
		}
		b, err := strconv.ParseInt(strings.Repeat(string(hex[2]), 2), 16, 64)
		if err != nil {
			return Color{}, err
		}
		return ColorFromRGB(int(r), int(g), int(b)), nil
	case 4:
		// hex = "#fff"
		return ColorFromHex(hex[1:])
	case 6:
		// hex = "ffffff"
		r, err := strconv.ParseInt(hex[:2], 16, 64)
		if err != nil {
			return Color{}, err
		}
		g, err := strconv.ParseInt(hex[2:4], 16, 64)
		if err != nil {
			return Color{}, err
		}
		b, err := strconv.ParseInt(hex[4:], 16, 64)
		if err != nil {
			return Color{}, err
		}
		return ColorFromRGB(int(r), int(g), int(b)), nil
	case 7:
		// hex = "#ffffff"
		return ColorFromHex(hex[1:])
	default:
		return Color{}, fmt.Errorf("invalid hex length")
	}
}

func ColorFromRGB(r, g, b int) Color {
	return Color{
		r: r,
		g: g,
		b: b,
	}
}

func ColorFromHSL(h, s, l float64) Color {
	c := (1 - math.Abs(2*l-1)) * s
	x := c * (1 - math.Abs(math.Mod(h*6, 2)-1))
	m := l - c/2

	var r, g, b float64
	switch {
	case h < 1.0/6:
		r, g, b = c, x, 0
	case h < 2.0/6:
		r, g, b = x, c, 0
	case h < 3.0/6:
		r, g, b = 0, c, x
	case h < 4.0/6:
		r, g, b = 0, x, c
	case h < 5.0/6:
		r, g, b = x, 0, c
	default:
		r, g, b = c, 0, x
	}

	return Color{
		r: int(math.Round((r + m) * 255)),
		g: int(math.Round((g + m) * 255)),
		b: int(math.Round((b + m) * 255)),
	}
}

func lerp(start, end int, t float64) int {
	return int(math.Round(float64(start) + float64(end-start)*t))
}
