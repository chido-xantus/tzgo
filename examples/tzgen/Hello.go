// Code generated by tzgen - DO NOT EDIT.
// This file is a binding generated from Hello smart contract at address KT1FsGDhS7PUWn1Yff8r5nXE4MArZwD2XhEi.
// Any manual changes will be lost.

package main

import (
	"context"
	"encoding/json"
	"math/big"
	"time"

	"github.com/pkg/errors"
	"github.com/tzgo/contract"
	"github.com/tzgo/contract/bind"
	"github.com/tzgo/micheline"
	"github.com/tzgo/rpc"
	"github.com/tzgo/tezos"
)

// Hello is a generated binding to a Tezos smart contract.
type Hello struct {
	bind.Contract
	builder HelloBuilder
	rpc     bind.RPC
	script  *micheline.Script
}

// HelloSession is a generated binding to a Tezos smart contract, that will
// use Opts for every call.
type HelloSession struct {
	*Hello
	Opts *rpc.CallOptions
}

// HelloBuilder is a generated struct that builds micheline.Parameters from
// go types.
type HelloBuilder struct{}

// NewHello creates a new Hello handle, bound to the provided address
// with the given rpc.
//
// Returns an error if the contract was not found at the given address.
func NewHello(ctx context.Context, address tezos.Address, client *rpc.Client) (*Hello, error) {
	script, err := client.GetContractScript(ctx, address)
	if err != nil {
		return nil, err
	}

	return &Hello{
		Contract: contract.NewContract(address, client),
		rpc:      client,
		script:   script,
	}, nil
}

// Session returns a new HelloSession with the configured rpc.CallOptions.
func (_h *Hello) Session(opts *rpc.CallOptions) *HelloSession {
	return &HelloSession{Hello: _h, Opts: opts}
}

// Builder returns the builder struct for this contract.
func (_h *Hello) Builder() HelloBuilder {
	return _h.builder
}

// Storage queries the current storage of the contract.
func (_h *Hello) Storage(ctx context.Context) (string, error) {
	return _h.StorageAt(ctx, rpc.Head)
}

// StorageAt queries the contract's storage at the given block.
func (_h *Hello) StorageAt(ctx context.Context, block rpc.BlockID) (string, error) {
	var storage string
	prim, err := _h.rpc.GetContractStorage(ctx, _h.Contract.Address(), block)
	if err != nil {
		return storage, errors.Wrap(err, "failed to get storage")
	}

	err = bind.UnmarshalPrim(prim, &storage)
	if err != nil {
		return storage, errors.Wrap(err, "failed to unmarshal storage")
	}
	return storage, nil
}

// DeployHello deploys a Hello contract by using client and opts, and HelloMicheline.
//
// Returns the receipt and a handle to the Hello deployed contract.
func DeployHello(ctx context.Context, opts *rpc.CallOptions, client *rpc.Client, storage string) (*rpc.Receipt, *Hello, error) {
	var script *micheline.Script
	err := json.Unmarshal([]byte(HelloMicheline), &script)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to unmarshall contract's script")
	}

	prim, err := bind.MarshalPrim(storage, false)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to marshal storage")
	}
	script.Storage = prim

	c := contract.NewEmptyContract(client).WithScript(script)
	receipt, err := c.Deploy(ctx, opts)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to deploy contract")
	}
	return receipt, &Hello{Contract: c, rpc: client}, nil
}

// region Entrypoints

// Greet makes a call to the `greet` contract entry.
//
// greet(string0 string)
func (_h *Hello) Greet(ctx context.Context, opts *rpc.CallOptions, string0 string) (*rpc.Receipt, error) {
	params, err := _h.builder.Greet(string0)
	if err != nil {
		return nil, err
	}
	return _h.Contract.Call(ctx, &contract.TxArgs{Params: params}, opts)
}

// Greet makes a call to the `greet` contract entry.
//
// greet(string0 string)
func (_h *HelloSession) Greet(ctx context.Context, string0 string) (*rpc.Receipt, error) {
	return _h.Hello.Greet(ctx, _h.Opts, string0)
}

// Greet builds `greet` contract entry's parameters.
//
// greet(string0 string)
func (HelloBuilder) Greet(string0 string) (micheline.Parameters, error) {
	prim, err := bind.MarshalParams(false, string0)
	if err != nil {
		return micheline.Parameters{}, errors.Wrap(err, "failed to marshal params")
	}
	return micheline.Parameters{Entrypoint: HelloGreetEntry, Value: prim}, nil
}

// endregion

// Hello entry names
const (
	HelloGreetEntry = "greet"
)

const HelloMicheline = `{"code":[{"prim":"parameter","args":[{"prim":"string","annots":["%greet"]}]},{"prim":"storage","args":[{"prim":"string"}]},{"prim":"code","args":[[{"prim":"UNPAIR"},{"prim":"PUSH","args":[{"prim":"nat"},{"int":"0"}]},{"prim":"DUP","args":[{"int":"2"}]},{"prim":"SIZE"},{"prim":"COMPARE"},{"prim":"GT"},{"prim":"NOT"},{"prim":"IF","args":[[{"prim":"PUSH","args":[{"prim":"string"},{"string":"r1"}]},{"prim":"PUSH","args":[{"prim":"string"},{"string":"INVALID_CONDITION"}]},{"prim":"PAIR"},{"prim":"FAILWITH"}],[]]},{"prim":"DUP"},{"prim":"PUSH","args":[{"prim":"string"},{"string":"Hello "}]},{"prim":"CONCAT"},{"prim":"DIP","args":[[{"prim":"DIG","args":[{"int":"1"}]},{"prim":"DROP"}]]},{"prim":"DUG","args":[{"int":"1"}]},{"prim":"DROP"},{"prim":"NIL","args":[{"prim":"operation"}]},{"prim":"PAIR"}]]}],"storage":{}}`

var (
	_ = big.NewInt
	_ = micheline.NewPrim
	_ = bind.MarshalParams
	_ = time.Now
)
