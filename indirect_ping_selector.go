package memberlist

// IndirectPingNodeSelector returns an ordered list of indirect-ping
// witness candidates for the given target. candidates is the pre-
// filtered set of currently-Alive peers, with self and the target
// already removed. Memberlist consumes the returned slice from index 0
// and uses the first IndirectChecks (K) entries; the caller controls
// preference order but does not need to truncate.
//
// The selector exists so callers (in porx, libopenstorage/gossip) can
// apply failure-domain preferences — interleaving witnesses across
// distinct AZs / zones / racks reduces the chance that all K witnesses
// share a partition with the prober and produce a unanimous-but-wrong
// DOWN verdict.
//
// Leaving Config.IndirectPingNodeSelector nil preserves upstream
// behavior exactly (kRandomNodes is used). The selector is invoked on
// the goroutine that observed the probe failure while
// m.nodeLock.RLock() is held, so implementations MUST NOT block and
// MUST NOT call back into memberlist methods that take nodeLock for
// write. Decoding Node.Meta and consulting a precomputed map is fine;
// network calls and gRPC fan-out are not.
type IndirectPingNodeSelector func(target *Node, candidates []*Node) []Node
