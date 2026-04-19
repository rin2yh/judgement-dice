package game

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type duelPhase int

const (
	duelIdle duelPhase = iota
	duelRolling
	duelJudgement
)

type Duel struct {
	phase    duelPhase
	cpu      *Dice
	player   *Dice
	cpuFinal int
	plyFinal int
}

func NewDuel() *Duel {
	return &Duel{
		phase:  duelIdle,
		cpu:    NewDice(),
		player: NewDice(),
	}
}

func (d *Duel) Reset() {
	d.phase = duelIdle
	d.cpu.Reset()
	d.player.Reset()
	d.cpuFinal = 0
	d.plyFinal = 0
}

func (d *Duel) Update() {
	switch d.phase {
	case duelIdle:
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			d.cpu.Roll()
			d.player.Roll()
			d.phase = duelRolling
		}
	case duelRolling:
		d.cpu.Update()
		d.player.Update()
		if d.cpu.State() == stateResult && d.player.State() == stateResult {
			d.cpuFinal = d.cpu.Final()
			d.plyFinal = d.player.Final()
			d.phase = duelJudgement
		}
	case duelJudgement:
		d.cpu.Update()
		d.player.Update()
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			d.Reset()
		}
	}
}

func (d *Duel) Draw(screen *ebiten.Image, face text.Face) {
	const cpuX, playerX float32 = 170, 470
	const diceY float32 = 200

	drawText(screen, face, "タンクトップ小隊", screenWidth/2, 36, titleColor)
	drawText(screen, face, "CPU", float64(cpuX), 110, textColor)
	drawText(screen, face, "YOU", float64(playerX), 110, textColor)

	d.cpu.Draw(screen, cpuX, diceY)
	d.player.Draw(screen, playerX, diceY)

	switch d.phase {
	case duelIdle:
		drawText(screen, face, "偽遊戯：ルールは単純！　二人同時にサイコロを振る", screenWidth/2, 330, textColor)
		drawText(screen, face, "オレより小さい目を出せばあんたの勝ち！　同じか大きいなら負けだ", screenWidth/2, 360, textColor)
		drawText(screen, face, "SPACE でスタート", screenWidth/2, 390, textColor)
	case duelRolling:
		drawText(screen, face, "サイコロが回っている……", screenWidth/2, 330, textColor)
	case duelJudgement:
		drawText(screen, face, fmt.Sprintf("CPU：%d  /  YOU：%d", d.cpuFinal, d.plyFinal), screenWidth/2, 330, titleColor)
		drawText(screen, face, d.Result(), screenWidth/2, 370, titleColor)
		drawText(screen, face, "SPACE でもう一勝負", screenWidth/2, 400, textColor)
	}
}

func (d *Duel) Result() string {
	if d.plyFinal < d.cpuFinal {
		return "あんたの勝ち！"
	}
	return "あんたの負け…"
}
