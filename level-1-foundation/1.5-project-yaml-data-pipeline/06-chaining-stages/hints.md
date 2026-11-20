## Hint 1 (arah umum)
Definisikan `type Stage func([]Row) []Row`.
Definisikan `type Pipeline struct { stages []Stage }`.

---

## Hint 2 (sintaks yang relevan)
```go
func NewPipeline(stages ...Stage) *Pipeline {
	return &Pipeline{stages: stages}
}

func (p *Pipeline) AddStage(stage Stage) {
	p.stages = append(p.stages, stage)
}
```

---

## Hint 3 (hampir jawaban)
```go
func (p *Pipeline) Execute(input []Row) []Row {
	current := input
	for _, stage := range p.stages {
		current = stage(current)
	}
	return current
}
```
