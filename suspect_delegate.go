package memberlist

// SuspectDelegate is an optional hook fired whenever a node transitions
// to StateSuspect. It lets callers react earlier than the natural
// SuspicionMult × log(N+1) × ProbeInterval timeout would allow.
//
// The callback runs inline on the goroutine that observed the
// transition, AFTER m.nodeLock has been released (the call is deferred
// in suspectNode for exactly that reason). Implementations may call
// back into other Memberlist methods (NodeState, Members, …) without
// risk of deadlock, but should still avoid blocking for long because
// the calling goroutine is the same one that drives further suspect-
// message broadcasts. Dispatch any non-trivial work asynchronously.
type SuspectDelegate interface {
	// NotifySuspect is invoked when a node transitions to StateSuspect.
	// The Node argument must not be modified. from is the name of the
	// peer that originally reported the suspicion (the From field of
	// the suspect message); it is empty when the local probe layer
	// raised the suspicion directly.
	NotifySuspect(node *Node, from string)
}
