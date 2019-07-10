#!/bin/bash
# go-objdump colorizes and reformats output of `go tool objdump`
# - it inserts an empty line after unconditional control-flow modifying instructions (JMP, RET, UD2)
# - it colors calls/returns in green
# - it colors traps (UD2) in red
# - it colors jumps (both conditional and unconditional) in blue
# - it colors padding/nops in violet
# - it colors the function name in yellow
# - it unindent the function body

function go-objdump() {
         go tool objdump "$@" |
                gsed -E "
                        s/^  ([^\t]+)(.*)/\1  \2/
                        s,^(TEXT )([^ ]+)(.*),$(tput setaf 3)\\1$(tput bold)\\2$(tput sgr0)$(tput setaf 3)\\3$(tput sgr0),
                        s/((JMP|RET|UD2).*)\$/\1\n/
                        s,.*(CALL |RET).*,$(tput setaf 2)&$(tput sgr0),
                        s,.*UD2.*,$(tput setaf 1)&$(tput sgr0),
                        s,.*J[A-Z]+.*,$(tput setaf 4)&$(tput sgr0),
                        s,.*(INT \\\$0x3|NOP).*,$(tput setaf 5)&$(tput sgr0),
                        "
}