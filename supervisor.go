package ergonode

import (
	"github.com/halturin/ergonode/etf"
	"github.com/halturin/ergonode/lib"
)

type SupervisorStrategy = string
type SupervisorChildRestart = string
type SupervisorChild = string

const (
	// Restart strategies:

	// SupervisorStrategyOneForOne If one child process terminates and is to be restarted, only
	// that child process is affected. This is the default restart strategy.
	SupervisorStrategyOneForOne = "one_for_one"

	// SupervisorStrategyOneForAll If one child process terminates and is to be restarted, all other
	// child processes are terminated and then all child processes are restarted.
	SupervisorStrategyOneForAll = "one_for_all"

	// SupervisorStrategyRestForOne If one child process terminates and is to be restarted,
	// the 'rest' of the child processes (that is, the child
	// processes after the terminated child process in the start order)
	// are terminated. Then the terminated child process and all
	// child processes after it are restarted
	SupervisorStrategyRestForOne = "rest_for_one"

	// SupervisorStrategySimpleOneForOne A simplified one_for_one supervisor, where all
	// child processes are dynamically added instances
	// of the same process type, that is, running the same code.
	SupervisorStrategySimpleOneForOne = "simple_one_for_one"

	// Restart types:

	// SupervisorChildRestartPermanent child process is always restarted
	SupervisorChildRestartPermanent = "permanent"

	// SupervisorChildRestartTemporary child process is never restarted
	// (not even when the supervisor restart strategy is rest_for_one
	// or one_for_all and a sibling death causes the temporary process
	// to be terminated)
	SupervisorChildRestartTemporary = "temporary"

	// SupervisorChildRestartTransient child process is restarted only if
	// it terminates abnormally, that is, with an exit reason other
	// than normal, shutdown, or {shutdown,Term}.
	SupervisorChildRestartTransient = "transient"
)

// SupervisorBehavior interface
type SupervisorBehavior interface {
	Init(process *Process, args ...interface{}) SupervisorSpec
	StartChild()
	StartLink()
	// RestartChild()
	// DeleteChild()
	// TerminateChild()
	// WhichChildren()
	// CountChildren()
}

type SupervisorSpec struct {
	children []SupervisorChildSpec
	strategy SupervisorStrategy
}

type SupervisorChildSpec struct {
	child   ProcessBehaviour
	restart SupervisorChildRestart
}

// Supervisor is implementation of ProcessBehavior interface
type Supervisor struct {
	process Process
}

func (sv *Supervisor) loop(p *Process, object interface{}, args ...interface{}) {
	spec := object.(SupervisorBehavior).Init(p, args...)
	lib.Log("Supervisor spec %#v\n", spec)
	p.ready <- true
	// var stop chan string
	// stop = make(chan string)
	for {
		// var message etf.Term
		var fromPid etf.Pid
		select {
		// case reason := <-stop:
		// 	object.(SupervisorBehavior).Terminate(reason, state)
		// 	return
		// case messageLocal := <-p.local:
		// 	message = messageLocal
		// case messageRemote := <-p.remote:
		// 	message = messageRemote[1]
		// 	fromPid = messageRemote[0].(etf.Pid)
		case <-p.context.Done():
			// object.(GenServerBehavior).Terminate("immediate", p.state)
			return
		}

		lib.Log("[%#v]. Message from %#v\n", p.self, fromPid)
		// switch m := message.(type) {
		// case etf.Tuple:
		// 	switch mtag := m[0].(type) {
		// 	case etf.Atom:
		// 		switch mtag {
		// 		case etf.Atom("$gen_call"):
		// 			fromTuple := m[1].(etf.Tuple)
		// 			code, reply, state1 := object.(GenServerBehavior).HandleCall(&fromTuple, &m[2], p.state)

		// 			p.state = state1
		// 			if code == "stop" {
		// 				stop <- code
		// 				continue
		// 			}

		// 			if reply != nil && code == "reply" {
		// 				// pid := fromTuple[0].(etf.Pid)
		// 				// ref := fromTuple[1]
		// 				// rep := etf.Term(etf.Tuple{ref, *reply})
		// 				// gs.Send(pid, &rep)
		// 			}

		// 		case etf.Atom("$gen_cast"):
		// 			code, state1 := object.(GenServerBehavior).HandleCast(&m[1], p.state)
		// 			p.state = state1
		// 			if code == "stop" {
		// 				stop <- code
		// 				continue
		// 			}
		// 		default:
		// 			code, state1 := object.(GenServerBehavior).HandleInfo(&message, p.state)
		// 			p.state = state1
		// 			if code == "stop" {
		// 				stop <- code
		// 				return
		// 			}
		// 		}
		// 	case etf.Ref:
		// 		lib.Log("got reply: %#v\n%#v", mtag, message)
		// 		// gs.chreply <- m
		// 	default:
		// 		lib.Log("mtag: %#v", mtag)
		// 		go func() {
		// 			code, state1 := object.(GenServerBehavior).HandleInfo(&message, p.state)
		// 			p.state = state1
		// 			if code == "stop" {
		// 				stop <- code
		// 				return
		// 			}
		// 		}()
		// 	}
		// default:
		// 	lib.Log("m: %#v", m)
		// 	go func() {
		// 		code, state1 := object.(GenServerBehavior).HandleInfo(&message, p.state)
		// 		p.state = state1
		// 		if code == "stop" {
		// 			stop <- code
		// 			return
		// 		}
		// 	}()
		// }
	}
}