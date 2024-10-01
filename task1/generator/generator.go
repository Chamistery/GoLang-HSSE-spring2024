package generator

type IDGenerator struct {
	generateFunc func(int) int
	currentID    int
}

func (g *IDGenerator) GenerateID() int {
	g.currentID = g.generateFunc(g.currentID)
	return g.currentID
}

func NewIDGenerator(generateFunc func(int) int) *IDGenerator {
	return &IDGenerator{
		generateFunc: generateFunc,
		currentID:    0,
	}
}
