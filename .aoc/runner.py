import subprocess
import time

def runDay(day, console, lang):
    if lang.hasIndividualTaskRunCommands():
        runTask(day, 1, console, lang)
        runTask(day, 2, console, lang)
    else:
        runTask(day, -1, console, lang)

def testDay(day, console, lang):
    if lang.hasIndividualTaskRunCommands():
        runTask(day, 1, console, lang, True)
        runTask(day, 2, console, lang, True)
    else:
        runTask(day, -1, console, lang, True)

def runTask(day, task, console, lang, test=False):
    if task >= 0:
        console.rule(f"[yellow]Task {task}")
    else:
        console.rule("[yellow]All tasks")
    startTime = time.time()
    if test:
        cmd = lang.getTestCommand(day, task, f"./src/day{str(day).rjust(2, '0')}")
    else:
        cmd = lang.getRunCommand(day, task, f"./src/day{str(day).rjust(2, '0')}")
    p = subprocess.Popen(cmd, shell=True, stdout=subprocess.PIPE, stderr=subprocess.STDOUT)

    buf = b""
    while p.poll() is None:
        out = p.stdout.read(1)
        if out != b'':
            if out == b'\n':
                console.log(buf.decode("utf8"))
                buf = b""
            else:
                buf += out
    out = p.stdout.read()
    if out != b'':
        buf += out
        console.log(buf.decode("utf8"))
    endTime = time.time()
    timeDiff = round(endTime - startTime, 3)
    if p.returncode == 0:
        console.log(f"[green]Task {task} finished execution successfully in {timeDiff}s.")
    else:
        console.log(f"[bold red]Task {task} failed execution in {timeDiff}s!")