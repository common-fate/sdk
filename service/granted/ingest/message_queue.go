package ingest

type messageQueue struct {
	pending      []Event
	bytes        int
	maxBatchSize int
}

func (q *messageQueue) push(m Event) (b []Event) {
	if q.pending == nil {
		q.pending = make([]Event, 0, q.maxBatchSize)
	}

	q.pending = append(q.pending, m)

	if b == nil && len(q.pending) == q.maxBatchSize {
		b = q.flush()
	}

	return
}

func (q *messageQueue) flush() (msgs []Event) {
	msgs, q.pending, q.bytes = q.pending, nil, 0
	return
}
