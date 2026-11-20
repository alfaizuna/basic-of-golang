//go:build ignore

package main

// example.go — hanya untuk referensi, jangan diubah

type Row struct {
	ID       string  `yaml:"id"`
	Item     string  `yaml:"item"`
	Category string  `yaml:"category"`
	Price    float64 `yaml:"price"`
	Quantity int     `yaml:"quantity"`
	Total    float64 `yaml:"total"`
}

type Stage func([]Row) []Row

type Pipeline struct {
	stages []Stage
}

func NewPipeline(stages ...Stage) *Pipeline {
	return &Pipeline{stages: stages}
}

func (p *Pipeline) AddStage(stage Stage) {
	p.stages = append(p.stages, stage)
}

func (p *Pipeline) Execute(input []Row) []Row {
	current := input
	for _, stage := range p.stages {
		current = stage(current)
	}
	return current
}
