package game

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/bitmapfont/v3"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

var (
	bgColor    = color.RGBA{R: 0x1b, G: 0x10, B: 0x3a, A: 0xff}
	titleColor = color.RGBA{R: 0xff, G: 0xd7, B: 0x4a, A: 0xff}
	textColor  = color.RGBA{R: 0xf5, G: 0xf1, B: 0xe3, A: 0xff}
	hintColor  = color.RGBA{R: 0xa0, G: 0x94, B: 0xc8, A: 0xff}
)

type Game struct {
	dice *Dice
	face text.Face
}

func New() *Game {
	return &Game{
		dice: NewDice(),
		face: text.NewGoXFace(bitmapfont.FaceEA),
	}
}

func (g *Game) Update() error {
	switch g.dice.State() {
	case stateIdle:
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.dice.Roll()
		}
	case stateResult:
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.dice.Roll()
			return nil
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyR) {
			g.dice.Reset()
			return nil
		}
		g.dice.Update()
	default:
		g.dice.Update()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(bgColor)

	g.drawText(screen, "ジャッジメントダイス", screenWidth/2, 36, titleColor)
	g.dice.Draw(screen, screenWidth/2, 170)

	switch g.dice.State() {
	case stateResult:
		g.drawText(screen, fmt.Sprintf("出目：%d", g.dice.Final()), screenWidth/2, 290, titleColor)
		g.drawText(screen, Effects[g.dice.Final()], screenWidth/2, 330, textColor)
	case stateRolling:
		g.drawText(screen, "ダイスを振っています……", screenWidth/2, 310, textColor)
	default:
		g.drawText(screen, "SPACE でダイスを振る", screenWidth/2, 310, textColor)
	}

	g.drawText(screen, "SPACE: Roll    R: Reset", screenWidth/2, screenHeight-24, hintColor)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) drawText(screen *ebiten.Image, s string, x, y float64, clr color.Color) {
	op := &text.DrawOptions{}
	op.GeoM.Translate(x, y)
	op.ColorScale.ScaleWithColor(clr)
	op.LayoutOptions.PrimaryAlign = text.AlignCenter
	op.LayoutOptions.SecondaryAlign = text.AlignCenter
	text.Draw(screen, s, g.face, op)
}
