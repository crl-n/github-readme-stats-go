package svg

const (
	bgColor       = "gray"
	progressColor = "yellow"
)

type ProgressBar struct {
	progressPercentage float32
	x                  int
	y                  int
	height             int
	width              int
}

func NewProgressBar(progressPercentage float32, x, y, height, width int) *ProgressBar {
	return &ProgressBar{
		progressPercentage: progressPercentage,
		x:                  x,
		y:                  y,
		height:             height,
		width:              width,
	}
}

func (pb *ProgressBar) ToElementGroup() *Group {
	group := NewGroup()

	bgRect := NewRect(RectParams{
		Width:  pb.width,
		Height: pb.height,
		Fill:   bgColor,
		X:      pb.x,
		Y:      pb.y,
	})
	group.AppendElement(bgRect)

	progressRect := NewRect(RectParams{
		Width:  int(pb.progressPercentage / 100.0 * float32(pb.width)),
		Height: pb.height,
		Fill:   progressColor,
		X:      pb.x,
		Y:      pb.y,
	})
	group.AppendElement(progressRect)

	return group
}
