package main

import (
	"cmd/go/internal/base"
	"cmd/go/internal/bug"
	"cmd/go/internal/clean"
	"cmd/go/internal/doc"
	"cmd/go/internal/envcmd"
	"cmd/go/internal/fix"
	"cmd/go/internal/fmtcmd"
	"cmd/go/internal/generate"
	"cmd/go/internal/get"
	"cmd/go/internal/help"
	"cmd/go/internal/list"
	"cmd/go/internal/modcmd"
	"cmd/go/internal/modfetch"
	"cmd/go/internal/modget"
	"cmd/go/internal/modload"
	"cmd/go/internal/run"
	"cmd/go/internal/test"
	"cmd/go/internal/tool"
	"cmd/go/internal/version"
	"cmd/go/internal/vet"
	"cmd/go/internal/work"
)

func init() {
	base.Go.Commands = []*base.Command{
		bug.CmdBug,
		work.CmdBuild,
		clean.CmdClean,
		doc.CmdDoc,
		envcmd.CmdEnv,
		fix.CmdFix,
		fmtcmd.CmdFmt,
		generate.CmdGenerate,
		get.CmdGet,
		work.CmdInstall,
		list.CmdList,
		modcmd.CmdMod,
		run.CmdRun,
		test.CmdTest,
		tool.CmdTool,
		version.CmdVersion,
		vet.CmdVet,

		help.HelpBuildmode,
		help.HelpC,
		help.HelpCache,
		help.HelpEnvironment,
		help.HelpFileType,
		modload.HelpGoMod,
		help.HelpGopath,
		get.HelpGopathGet,
		modfetch.HelpGoproxy,
		help.HelpImportPath,
		modload.HelpModules,
		modget.HelpModuleGet,
		help.HelpPackages,
		test.HelpTestflag,
		test.HelpTestfunc,
	}
}
