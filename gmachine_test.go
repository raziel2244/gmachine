package gmachine_test

import (
	"gmachine"
	"testing"
)

func TestNew(t *testing.T) {
	t.Parallel()

	g := gmachine.New()

	var wantA, gotA uint64 = 0, g.A
	if wantA != gotA {
		t.Errorf("want initial A value %d, got %d", wantA, gotA)
	}

	var wantP, gotP uint64 = 0, g.P
	if wantP != gotP {
		t.Errorf("want initial P value %d, got %d", wantP, gotP)
	}

	wantMemSize := gmachine.DefaultMemSize
	gotMemSize := len(g.Memory)
	if wantMemSize != gotMemSize {
		t.Errorf("want %d words of memory, got %d", wantMemSize, gotMemSize)
	}

	var wantMemValue uint64 = 0
	gotMemValue := g.Memory[gmachine.DefaultMemSize-1]
	if wantMemValue != gotMemValue {
		t.Errorf("want last memory location to contain %d, got %d", wantMemValue, gotMemValue)
	}
}

func TestHalt(t *testing.T) {
	t.Parallel()

	g := gmachine.New()
	g.Run()

	var wantP, gotP uint64 = 1, g.P
	if wantP != gotP {
		t.Errorf("want P == %v, got %v", wantP, gotP)
	}
}

func TestNOOP(t *testing.T) {
	t.Parallel()

	program := gmachine.Program{
		gmachine.OpNOOP,
		gmachine.OpHALT,
	}

	g := gmachine.New()
	g.RunProgram(program)

	var wantP, gotP uint64 = 2, g.P
	if wantP != gotP {
		t.Errorf("want P == %v, got %v", wantP, gotP)
	}
}

func TestINCA(t *testing.T) {
	t.Parallel()

	program := gmachine.Program{
		gmachine.OpINCA,
		gmachine.OpHALT,
	}

	g := gmachine.New()
	g.RunProgram(program)

	var wantA, gotA uint64 = 1, g.A
	if wantA != gotA {
		t.Errorf("want A == %v, got %v", wantA, gotA)
	}
}

func TestDECA(t *testing.T) {
	t.Parallel()

	program := gmachine.Program{
		gmachine.OpDECA,
		gmachine.OpHALT,
	}

	g := gmachine.New()
	g.A = 2
	g.RunProgram(program)

	var wantA, gotA uint64 = 1, g.A
	if wantA != gotA {
		t.Errorf("want A == %v, got %v", wantA, gotA)
	}
}

func TestSETA(t *testing.T) {
	t.Parallel()

	program := gmachine.Program{
		gmachine.OpSETA, 1,
		gmachine.OpHALT,
	}

	g := gmachine.New()
	g.RunProgram(program)

	var wantA, gotA uint64 = 1, g.A
	if wantA != gotA {
		t.Errorf("want A == %v, got %v", wantA, gotA)
	}

	var wantP, gotP uint64 = 3, g.P
	if wantP != gotP {
		t.Errorf("want P == %v, got %v", wantP, gotP)
	}
}

func TestSub(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		input, want uint64
	}{
		{input: 3, want: 1},
		{input: 5, want: 3},
		{input: 7, want: 5},
	}

	for _, tc := range testCases {
		program := gmachine.Program{
			gmachine.OpSETA, tc.input,
			gmachine.OpDECA,
			gmachine.OpDECA,
			gmachine.OpHALT,
		}

		g := gmachine.New()
		g.RunProgram(program)

		var wantA, gotA uint64 = tc.want, g.A
		if wantA != gotA {
			t.Errorf("want A == %v, got %v", wantA, gotA)
		}
	}
}
