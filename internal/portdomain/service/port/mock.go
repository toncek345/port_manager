package port

import "context"

type Mock struct {
	UpsertPortFn func(context.Context, *Port) error
	GetPortFn    func(context.Context, int64) (*Port, error)
}

func (m *Mock) UpsertPort(ctx context.Context, port *Port) error {
	return m.UpsertPortFn(ctx, port)
}

func (m *Mock) GetPort(ctx context.Context, id int64) (*Port, error) {
	return m.GetPortFn(ctx, id)
}
