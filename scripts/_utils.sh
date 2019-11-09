## Outputs Front-Mater formatted failures for functions not returning 0
## Use the following line after sourcing this file to set failure trap
##    trap 'failure "LINENO" "BASH_LINENO" "${BASH_COMMAND}" "${?}"' ERR
failure(){
    local -n _lineno="${1:-LINENO}"
    local -n _bash_lineno="${2:-BASH_LINENO}"
    local _last_command="${3:-${BASH_COMMAND}}"
    local _code="${4:-0}"

    ## Workaround for read EOF combo tripping traps
    if ! ((_code)); then
        return "${_code}"
    fi

    local _last_command_height="$(wc -l <<<"${_last_command}")"

    local -a _output_array=()
    _output_array+=(
        '---'
        "lines_history: [${_lineno} ${_bash_lineno[*]}]"
        "function_trace: [${FUNCNAME[*]}]"
        "exit_code: ${_code}"
    )

    if [[ "${#BASH_SOURCE[@]}" -gt '1' ]]; then
        _output_array+=('source_trace:')
        for _item in "${BASH_SOURCE[@]}"; do
            _output_array+=("  - ${_item}")
        done
    else
        _output_array+=("source_trace: [${BASH_SOURCE[*]}]")
    fi

    if [[ "${_last_command_height}" -gt '1' ]]; then
        _output_array+=(
            'last_command: ->'
            "${_last_command}"
        )
    else
        _output_array+=("last_command: ${_last_command}")
    fi

    _output_array+=('---')
    printf '%s\n' "${_output_array[@]}" >&2
    exit ${_code}
}
