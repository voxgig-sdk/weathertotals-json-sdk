package entity

import (
	"encoding/json"
	"fmt"

	"github.com/voxgig-sdk/weathertotals-json-sdk/go/core"

	vs "github.com/voxgig-sdk/weathertotals-json-sdk/go/utility/struct"
)

type SnowpackEntity struct {
	name    string
	client  *core.WeathertotalsJsonSDK
	utility *core.Utility
	entopts map[string]any
	data    any
	match   map[string]any
	entctx  *core.Context
	deleted bool
}

func NewSnowpackEntity(client *core.WeathertotalsJsonSDK, entopts map[string]any) *SnowpackEntity {
	if entopts == nil {
		entopts = map[string]any{}
	}
	if _, ok := entopts["active"]; !ok {
		entopts["active"] = true
	} else if entopts["active"] == false {
		// keep false
	} else {
		entopts["active"] = true
	}

	e := &SnowpackEntity{
		name:    "snowpack",
		client:  client,
		utility: client.GetUtility(),
		entopts: entopts,
		data:    map[string]any{},
		match:   map[string]any{},
	}

	e.entctx = e.utility.MakeContext(map[string]any{
		"entity":  e,
		"entopts": entopts,
	}, client.GetRootCtx())

	e.utility.FeatureHook(e.entctx, "PostConstructEntity")

	return e
}

func (e *SnowpackEntity) GetName() string { return e.name }

// An entity prints and serialises as its data, as ts's toString and toJSON
// do: the client it holds carries the options.
func (e *SnowpackEntity) String() string {
	return "Snowpack " + vs.Jsonify(e.data, map[string]any{"indent": 0})
}

func (e *SnowpackEntity) GoString() string {
	return e.String()
}

func (e *SnowpackEntity) MarshalJSON() ([]byte, error) {
	data, ok := e.data.(map[string]any)
	if !ok {
		return json.Marshal(e.data)
	}
	out := map[string]any{}
	for k, v := range data {
		out[k] = v
	}
	out["voxgig$entity"] = "Snowpack"
	return json.Marshal(out)
}

func (e *SnowpackEntity) MarkDeleted() {
	e.deleted = true
}


// Deleted reports whether a successful Remove has resolved on this instance.
func (e *SnowpackEntity) Deleted() bool {
	return e.deleted
}


func (e *SnowpackEntity) Make() core.Entity {
	opts := map[string]any{}
	for k, v := range e.entopts {
		opts[k] = v
	}
	return NewSnowpackEntity(e.client, opts)
}

func (e *SnowpackEntity) Data(args ...any) any {
	if len(args) > 0 && args[0] != nil {
		e.data = vs.Clone(args[0])
		e.utility.FeatureHook(e.entctx, "SetData")
	}

	e.utility.FeatureHook(e.entctx, "GetData")
	out := vs.Clone(e.data)
	return out
}

func (e *SnowpackEntity) Match(args ...any) any {
	if len(args) > 0 && args[0] != nil {
		e.match = core.ToMapAny(vs.Clone(args[0]))
		if e.match == nil {
			e.match = map[string]any{}
		}
		e.utility.FeatureHook(e.entctx, "SetMatch")
	}

	e.utility.FeatureHook(e.entctx, "GetMatch")
	out := vs.Clone(e.match)
	return out
}

// DataTyped is the statically-typed accessor for this entity's data. With no
// argument it returns the current data as an Snowpack; with an argument it
// sets the data and returns the stored value. It delegates to the untyped Data
// (identical runtime) and converts at the typed boundary.
func (e *SnowpackEntity) DataTyped(data ...Snowpack) Snowpack {
	if len(data) > 0 {
		return typedFrom[Snowpack](e.Data(asMap(data[0])))
	}
	return typedFrom[Snowpack](e.Data())
}

// MatchTyped mirrors DataTyped for the entity's match filter. The match is a
// partial of the entity, so it round-trips through Snowpack (all fields
// optional at the wire level).
func (e *SnowpackEntity) MatchTyped(match ...Snowpack) Snowpack {
	if len(match) > 0 {
		return typedFrom[Snowpack](e.Match(asMap(match[0])))
	}
	return typedFrom[Snowpack](e.Match())
}

func (e *SnowpackEntity) Stream(action string, args map[string]any, callopts map[string]any) <-chan core.StreamItem {
	out := make(chan core.StreamItem)

	if callopts == nil {
		callopts = map[string]any{}
	}

	var signal <-chan struct{}
	switch s := callopts["signal"].(type) {
	case <-chan struct{}:
		signal = s
	case chan struct{}:
		signal = s
	}

	ctrl := map[string]any{}
	if c := core.ToMapAny(callopts["ctrl"]); c != nil {
		for k, v := range c {
			ctrl[k] = v
		}
	}

	ctxmap := map[string]any{
		"opname": action,
		"ctrl":   ctrl,
		"match":  e.match,
		"data":   e.data,
	}
	for k, v := range args {
		ctxmap[k] = v
	}

	utility := e.utility
	ctx := utility.MakeContext(ctxmap, e.entctx)
	ctx.Meta["stream"] = callopts

	// Outbound: expose the caller's payload so the request builder / transport
	// can stream it as the request body.
	if body := callopts["body"]; body != nil {
		ctx.Reqdata["body$"] = body
		ctx.Meta["stream_out"] = body
	}

	send := func(item core.StreamItem) bool {
		select {
		case <-signal:
			return false
		case out <- item:
			return true
		}
	}

	// What MakeError or Done hands back: the error, as the last value, or
	// under `throw: false` the data there is.
	sendData := func(data any, err error) {
		if err != nil {
			send(core.StreamItem{Err: err})
			return
		}
		switch d := data.(type) {
		case []any:
			for _, item := range d {
				if !send(core.StreamItem{Item: item}) {
					return
				}
			}
		case nil:
			// nothing to yield
		default:
			send(core.StreamItem{Item: d})
		}
	}

	go func() {
		defer close(out)

		// A panicking hook or stream function leaves through MakeError, as
		// runOp's does. A goroutine the stream function starts is out of reach.
		defer func() {
			if r := recover(); r != nil {
				sendData(e.recovered(ctx, r))
			}
		}()

		// A failed step leaves through MakeError, as an operation's does.
		if err := e.streamSteps(ctx); err != nil {
			sendData(utility.MakeError(ctx, err))
			return
		}

		// Inbound: prefer the streaming feature's incremental iterator; else
		// fall back to the materialised items so Stream always yields.
		if ctx.Result != nil && ctx.Result.Stream != nil {
			// Done does not run on this path, so its record is cleaned here.
			utility.CleanExplain(ctx)
			for item := range ctx.Result.Stream() {
				if !send(core.StreamItem{Item: item}) {
					return
				}
			}
			return
		}

		sendData(utility.Done(ctx))
	}()

	return out
}

// The steps an operation runs, with their hooks; the first that fails hands
// back its error.
func (e *SnowpackEntity) streamSteps(ctx *core.Context) error {
	utility := e.utility

	utility.FeatureHook(ctx, "PrePoint")
	point, err := utility.MakePoint(ctx)
	ctx.Out["point"] = point
	if err != nil {
		return err
	}

	utility.FeatureHook(ctx, "PreSpec")
	spec, err := utility.MakeSpec(ctx)
	ctx.Out["spec"] = spec
	if err != nil {
		return err
	}

	utility.FeatureHook(ctx, "PreRequest")
	req, err := utility.MakeRequest(ctx)
	ctx.Out["request"] = req
	if err != nil {
		return err
	}

	utility.FeatureHook(ctx, "PreResponse")
	resp, err := utility.MakeResponse(ctx)
	ctx.Out["response"] = resp
	if err != nil {
		return err
	}

	utility.FeatureHook(ctx, "PreResult")
	result, err := utility.MakeResult(ctx)
	ctx.Out["result"] = result
	if err != nil {
		return err
	}

	utility.FeatureHook(ctx, "PreDone")
	return nil
}


func (e *SnowpackEntity) Load(reqmatch map[string]any, ctrl map[string]any) (any, error) {
	utility := e.utility
	ctx := utility.MakeContext(map[string]any{
		"opname":   "load",
		"ctrl":     ctrl,
		"match":    e.match,
		"data":     e.data,
		"reqmatch": reqmatch,
	}, e.entctx)

	return e.runOp(ctx, func() {
		if ctx.Result != nil {
			if ctx.Result.Resmatch != nil {
				e.match = ctx.Result.Resmatch
			}
			if ctx.Result.Resdata != nil {
				e.data = map[string]any{}
				if data := core.ToMapAny(vs.Clone(ctx.Result.Resdata)); data != nil {
					e.data = data
				}
			}
		}
	})
}

// LoadTyped is the statically-typed variant of Load: it takes an
// SnowpackLoadMatch and returns an Snowpack. It delegates to the untyped
// Load (identical runtime) and converts at the typed boundary.
func (e *SnowpackEntity) LoadTyped(reqmatch SnowpackLoadMatch, ctrl map[string]any) (Snowpack, error) {
	res, err := e.Load(asMap(reqmatch), ctrl)
	if err != nil {
		return Snowpack{}, err
	}
	return typedFrom[Snowpack](res), nil
}



func (e *SnowpackEntity) List(_ map[string]any, _ map[string]any) (any, error) {
	return core.UnsupportedOp("list", e.name)
}


func (e *SnowpackEntity) Create(_ map[string]any, _ map[string]any) (any, error) {
	return core.UnsupportedOp("create", e.name)
}


func (e *SnowpackEntity) Update(_ map[string]any, _ map[string]any) (any, error) {
	return core.UnsupportedOp("update", e.name)
}


func (e *SnowpackEntity) Patch(_ map[string]any, _ map[string]any) (any, error) {
	return core.UnsupportedOp("patch", e.name)
}


func (e *SnowpackEntity) Remove(_ map[string]any, _ map[string]any) (any, error) {
	return core.UnsupportedOp("remove", e.name)
}


func (e *SnowpackEntity) runOp(ctx *core.Context, postDone func()) (out any, err error) {
	utility := e.utility

	defer func() {
		if r := recover(); r != nil {
			out, err = e.recovered(ctx, r)
		}
	}()

	utility.FeatureHook(ctx, "PrePoint")
	point, err := utility.MakePoint(ctx)
	ctx.Out["point"] = point
	if err != nil {
		return utility.MakeError(ctx, err)
	}

	utility.FeatureHook(ctx, "PreSpec")
	spec, err := utility.MakeSpec(ctx)
	ctx.Out["spec"] = spec
	if err != nil {
		return utility.MakeError(ctx, err)
	}

	utility.FeatureHook(ctx, "PreRequest")
	resp, err := utility.MakeRequest(ctx)
	ctx.Out["request"] = resp
	if err != nil {
		return utility.MakeError(ctx, err)
	}

	utility.FeatureHook(ctx, "PreResponse")
	resp2, err := utility.MakeResponse(ctx)
	ctx.Out["response"] = resp2
	if err != nil {
		return utility.MakeError(ctx, err)
	}

	utility.FeatureHook(ctx, "PreResult")
	result, err := utility.MakeResult(ctx)
	ctx.Out["result"] = result
	if err != nil {
		return utility.MakeError(ctx, err)
	}

	utility.FeatureHook(ctx, "PreDone")
	postDone()

	out, err = utility.Done(ctx)
	if err != nil {
		return out, err
	}

	opname := ""
	if ctx.Op != nil {
		opname = ctx.Op.Name
	}

	if ctx.Result != nil && ctx.Result.Ok && opname != "list" {
		if opname == "remove" {
			e.MarkDeleted()
		}
		return e, nil
	}

	return out, nil
}

// A hook, fetcher or parser that panics never reached MakeError, and its
// message can quote the request.
func (e *SnowpackEntity) recovered(ctx *core.Context, r any) (any, error) {
	perr, ok := r.(error)
	if !ok {
		perr = fmt.Errorf("%v", r)
	}
	return e.utility.MakeError(ctx, perr)
}
