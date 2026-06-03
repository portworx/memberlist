package memberlist

// IndirectPingNodeSelector picks up to k indirect-ping witness nodes
// for the given target. candidates is the pre-filtered set of currently-
// Alive peers, with self and the target already removed. The returned
// slice should contain at most k entries; anything beyond k is ignored.
//
// The selector exists so callers (in porx, libopenstorage/gossip) can
// apply failure-domain preferences — picking witnesses across distinct
// AZs / zones / racks reduces the chance that all K witnesses share a
// partition with the prober and produce a unanimous-but-wrong DOWN
// verdict.
//
// Leaving Config.IndirectPingNodeSelector nil preserves upstream
// behavior exactly (kRandomNodes is used). The selector is invoked on
// the goroutine that observed the probe failure, while
// m.nodeLock.RLock() is held, so implementations MUST NOT block and
// MUST NOT call back into memberlist methods that take nodeLock for
// write. Decoding Node.Meta and consulting a precomputed map is fine;
// network calls and gRPC fan-out are not.
type IndirectPingNodeSelector func(k int, target *Node, candidates []*Node) []Node
