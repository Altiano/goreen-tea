package visitorCounter

import "context"

func (d domain) IncreaseVisits(ctx context.Context) error {
	_, span := d.Tracer.Start(ctx, "VisitorCounterDomain.IncreaseVisits")
	defer span.End()

	// Increase the live tracking number
	_, err := d.Memcache.Increment("visitorCounter")

	// Return
	return err
}
