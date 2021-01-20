package shell

// Bash is the bash shell integration.
var Bash = Shell(`# Put the line below in ~/.bashrc or ~/bash_profile:
#
#   eval "$(jump shell bash)"
#
# The following lines are autogenerated:

__jump_prompt_command() {
  local status=$?
  jump chdir && return $status
}

__jump_hint() {
  local term="${COMP_LINE/#{{.Bind}} /}"
  echo \'$(jump hint "$term")\'
}

{{.Bind}}() {
  local dir="$(jump cd "$@")"
  test -d "$dir"  && cd "$dir"
}

[[ "$PROMPT_COMMAND" =~ __jump_prompt_command ]] || {
  PROMPT_COMMAND="__jump_prompt_command;$PROMPT_COMMAND"
}

complete -o dirnames -C '__jump_hint' {{.Bind}}
`)
