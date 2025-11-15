// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package flagdoc contains flag definitions for 'go doc cmd/link'.
// It duplicates flags from internal/ld/main.go to show full help.
package flagdoc

import (
	"cmd/internal/quoted"
	"flag"
)

var (
	// === LINKER FLAGS FOR go doc cmd/link ===
	// Duplicated from cmd/link/internal/ld/main.go
	Buildid          = flag.String("buildid", "", "record `id` as Go toolchain build id")
	BindNow          = flag.Bool("bindnow", false, "mark a dynamically linked ELF object for immediate function binding")
	Outfile          = flag.String("o", "", "write output to `file`")
	PluginPath       = flag.String("pluginpath", "", "full path name for plugin")
	Fipso            = flag.String("fipso", "", "write fips module to `file`")
	InstallSuffix    = flag.String("installsuffix", "", "set package directory `suffix`")
	DumpDep          = flag.Bool("dumpdep", false, "dump symbol dependency graph")
	Race             = flag.Bool("race", false, "enable race detector")
	Msan             = flag.Bool("msan", false, "enable MSan interface")
	Asan             = flag.Bool("asan", false, "enable ASan interface")
	Aslr             = flag.Bool("aslr", true, "enable ASLR for buildmode=c-shared on windows")
	FieldTrack       = flag.String("k", "", "set field tracking `symbol`")
	LibGCC           = flag.String("libgcc", "", "compiler support lib for internal linking; use \"none\" to disable")
	Tmpdir           = flag.String("tmpdir", "", "use `directory` for temporary files")
	Extld            = new(quoted.Flag)
	Extldflags       = new(quoted.Flag)
	Extar            = flag.String("extar", "", "archive program for buildmode=c-archive")
	CaptureHostObjs  = flag.String("capturehostobjs", "", "capture host object files loaded during internal linking to specified dir")
	Aflag            = flag.Bool("a", false, "no-op (deprecated)")
	Cflag            = flag.Bool("c", false, "dump call graph")
	Dflag            = flag.Bool("d", false, "disable dynamic executable")
	Fflag            = flag.Bool("f", false, "ignore version mismatch")
	Gflag            = flag.Bool("g", false, "disable go package data checks")
	Hflag            = flag.Bool("h", false, "halt on error")
	Nflag            = flag.Bool("n", false, "no-op (deprecated)")
	Sflag            = flag.Bool("s", false, "disable symbol table")
	HostBuildid      = flag.String("B", "", "set ELF NT_GNU_BUILD_ID `note` or Mach-O UUID; use \"gobuildid\" to generate it from the Go build ID; \"none\" to disable")
	Interpreter      = flag.String("I", "", "use `linker` as ELF dynamic linker")
	CheckLinkname    = flag.Bool("checklinkname", true, "check linkname symbol references")
	DebugTramp       = flag.Int("debugtramp", 0, "debug trampolines")
	DebugTextSize    = flag.Int("debugtextsize", 0, "debug text section max size")
	DebugNosplit     = flag.Bool("debugnosplit", false, "dump nosplit call graph")
	StrictDups       = flag.Int("strictdups", 0, "sanity check duplicate symbol contents during object file reading (1=warn 2=err).")
	Round            = flag.Int64("R", -1, "set address rounding `quantum`")
	TextAddr         = flag.Int64("T", -1, "set the start address of text symbols")
	DataAddr         = flag.Int64("D", -1, "set the start address of data symbols")
	FuncAlign        = flag.Int("funcalign", 0, "set function align to `N` bytes")
	EntrySymbol      = flag.String("E", "", "set `entry` symbol name")
	PruneWeakMap     = flag.Bool("pruneweakmap", true, "prune weak mapinit refs")
	RandLayout       = flag.Int64("randlayout", 0, "randomize function layout")
	AllErrors        = flag.Bool("e", false, "no limit on number of errors reported")
	CpuProfile       = flag.String("cpuprofile", "", "write cpu profile to `file`")
	MemProfile       = flag.String("memprofile", "", "write memory profile to `file`")
	MemProfileRate   = flag.Int64("memprofilerate", 0, "set runtime.MemProfileRate to `rate`")
	Benchmark        = flag.String("benchmark", "", "set to 'mem' or 'cpu' to enable phase benchmarking")
	BenchmarkProfile = flag.String("benchmarkprofile", "", "emit phase profiles to `base`_phase.{cpu,mem}prof")
	Wflag            = new(bool)
)
