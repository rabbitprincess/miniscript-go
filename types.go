package miniscript

import "strings"

const (
	// Length of a pubkey inside P2WSH, which are 33 byte compressed pubkeys.
	pubKeyLen = 33
	// Length of a pubkey data push in P2WSH, which is 1+33 (1 byte for the VarInt encoding of 33).
	pubKeyDataPushLen = 34

	// The maximum size in bytes of a standard witnessScript
	maxStandardP2WSHScriptSize = 3600
	// Maximum number of non-push operations per script
	maxOpsPerScript = 201

	// Maximum number of keys in a multisig.
	multisigMaxKeys = 20
)

// This type encapsulates the miniscript type system properties.
// Every miniscript expression is one of 4 basic types, and additionally has
// a number of boolean type properties.
type basicType string

// The basic types are:
const (
	// - "B" Base:
	//   - Takes its inputs from the top of the stack.
	//   - When satisfied, pushes a nonzero value of up to 4 bytes onto the stack.
	//   - When dissatisfied, pushes a 0 onto the stack.
	//   - This is used for most expressions, and required for the top level one.
	//   - For example: older(n) = <n> OP_CHECKSEQUENCEVERIFY.
	typeB basicType = "B"

	// - "V" Verify:
	//   - Takes its inputs from the top of the stack.
	//   - When satisfied, pushes nothing.
	//   - Cannot be dissatisfied.
	//   - This can be obtained by adding an OP_VERIFY to a B, modifying the last opcode
	//     of a B to its -VERIFY version (only for OP_CHECKSIG, OP_CHECKSIGVERIFY,
	//     OP_NUMEQUAL and OP_EQUAL), or by combining a V fragment under some conditions.
	//   - For example vc:pk_k(key) = <key> OP_CHECKSIGVERIFY
	typeV basicType = "V"

	// - "K" Key:
	//   - Takes its inputs from the top of the stack.
	//   - Becomes a B when followed by OP_CHECKSIG.
	//   - Always pushes a public key onto the stack, for which a signature is to be
	//     provided to satisfy the expression.
	//   - For example pk_h(key) = OP_DUP OP_HASH160 <Hash160(key)> OP_EQUALVERIFY
	typeK basicType = "K"

	// - "W" Wrapped:
	//   - Takes its input from one below the top of the stack.
	//   - When satisfied, pushes a nonzero value (like B) on top of the stack, or one below.
	//   - When dissatisfied, pushes 0 op top of the stack or one below.
	//   - Is always "OP_SWAP [B]" or "OP_TOALTSTACK [B] OP_FROMALTSTACK".
	//   - For example sc:pk_k(key) = OP_SWAP <key> OP_CHECKSIG
	typeW basicType = "W"
)

type properties struct {
	// Basic type properties
	z, o, n, d, u bool
	// Malleability properties.
	// If `m`, a non-malleable satisfaction is guaranteed to exist.
	// The purpose of s/f/e is only to compute `m` and can be disregarded afterwards.
	m, s, f, e bool
	// Check if the rightmost script byte produced by this node is OP_EQUAL, OP_CHECKSIG or
	// OP_CHECKMULTISIG.
	//
	// If so, it can be be converted into the VERIFY version if an ancestor is the verify wrapper
	// `v`, i.e. OP_EQUALVERIFY, OP_CHECKSIGVERIFY and OP_CHECKMULTISIGVERIFY instead of using two
	// opcodes, e.g. `OP_EQUAL OP_VERIFY`.
	canCollapseVerify bool
}

func (p properties) String() string {
	s := strings.Builder{}
	if p.z {
		s.WriteRune('z')
	}
	if p.o {
		s.WriteRune('o')
	}
	if p.n {
		s.WriteRune('n')
	}
	if p.d {
		s.WriteRune('d')
	}
	if p.u {
		s.WriteRune('u')
	}
	if p.m {
		s.WriteRune('m')
	}
	if p.s {
		s.WriteRune('s')
	}
	if p.f {
		s.WriteRune('f')
	}
	if p.e {
		s.WriteRune('e')
	}
	return s.String()
}

const (
	// All fragment identifiers.
	f_0         = "0"         // 0
	f_1         = "1"         // 1
	f_pk_k      = "pk_k"      // pk_k(key)
	f_pk_h      = "pk_h"      // pk_h(key)
	f_pk        = "pk"        // pk(key) = c:pk_k(key)
	f_pkh       = "pkh"       // pkh(key) = c:pk_h(key)
	f_sha256    = "sha256"    // sha256(h)
	f_ripemd160 = "ripemd160" // ripemd160(h)
	f_hash256   = "hash256"   // hash256(h)
	f_hash160   = "hash160"   // hash160(h)
	f_older     = "older"     // older(n)
	f_after     = "after"     // after(n)
	f_andor     = "andor"     // andor(X,Y,Z)
	f_and_v     = "and_v"     // and_v(X,Y)
	f_and_b     = "and_b"     // and_b(X,Y)
	f_and_n     = "and_n"     // and_n(X,Y) = andor(X,Y,0)
	f_or_b      = "or_b"      // or_b(X,Z)
	f_or_c      = "or_c"      // or_c(X,Z)
	f_or_d      = "or_d"      // or_d(X,Z)
	f_or_i      = "or_i"      // or_i(X,Z)
	f_thresh    = "thresh"    // thresh(k,X1,...,Xn)
	f_multi     = "multi"     // multi(k,key1,...,keyn)
	f_wrap_a    = "a"         // a:X
	f_wrap_s    = "s"         // s:X
	f_wrap_c    = "c"         // c:X
	f_wrap_d    = "d"         // d:X
	f_wrap_v    = "v"         // v:X
	f_wrap_j    = "j"         // j:X
	f_wrap_n    = "n"         // n:X
	f_wrap_t    = "t"         // t:X = and_v(X,1)
	f_wrap_l    = "l"         // l:X = or_i(0,X)
	f_wrap_u    = "u"         // u:X = or_i(X,0)
)
