package events

import (
	"fmt"
	"time"

	"github.com/fossoreslp/go-uuid-v4"
)

// mock for tests
type orderIDGenerator interface {
	Generate() int64
}

type idempotentKeyGenerator interface {
	Generate() string
}

type operationMomentGenerator interface {
	Generate() time.Time
}

type Event struct {
	ID              int64     `json:"id"`
	EventType       string    `json:"event"`
	OperationMoment time.Time `json:"moment"`
	IdempotentKey   string    `json:"idempotent_key"`
}

type EventFactory struct {
	OrderIDGenerator orderIDGenerator
	IdempotentKey    idempotentKeyGenerator
	OperationMoment  operationMomentGenerator
}

func (f *EventFactory) Create(event string) Event {
	return Event{
		ID:              f.OrderIDGenerator.Generate(),
		EventType:       event,
		IdempotentKey:   f.IdempotentKey.Generate(),
		OperationMoment: f.OperationMoment.Generate(),
	}
}

func NewDefaultFactory(start int) *EventFactory {
	return New(
		NewSeqGen(start),
		NewUUIDv4Generator(),
		&Clock{},
	)
}

func New(
	idGen orderIDGenerator,
	idemKeyGen idempotentKeyGenerator,
	momentGen operationMomentGenerator,
) *EventFactory {
	return &EventFactory{
		idGen,
		idemKeyGen,
		momentGen,
	}
}

type SeqGen struct {
	cur int
}

func NewSeqGen(start int) *SeqGen {
	return &SeqGen{
		cur: start,
	}
}

type Clock struct{}

func (c *Clock) Generate() time.Time {
	return time.Now()
}

func (g *SeqGen) Generate() int64 {
	val := g.cur
	g.cur++

	return int64(val)
}

type UUIDv4Generator struct {
	gen uuid.UUID
}

func NewUUIDv4Generator() *UUIDv4Generator {
	gen, err := uuid.New()
	if err != nil {
		panic(fmt.Sprintf("NewUUIDv4Generator failed: %s", err.Error()))
	}

	return &UUIDv4Generator{
		gen: gen,
	}
}

func (g *UUIDv4Generator) Generate() string {
	return g.gen.String()
}
