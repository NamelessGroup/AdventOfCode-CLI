if (Test-Path -Path './.venv/' -PathType container) {
    & ./.venv/Scripts/Activate.ps1
} else {
    echo "The usage of a virtual environment is recommended."
    echo "This script searches for an enviroment in: '.venv'"
}

$allArgs = $PsBoundParameters.Values + $args
python ./.aoc/aoc.py $allArgs