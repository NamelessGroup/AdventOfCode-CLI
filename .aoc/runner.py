import subprocess
import time

def runDay(day, console, lang, task):
    if lang.hasIndividualTaskRunCommands():
        if task != -1:
            prepareAndRun(day, task, console, lang)
        else:
            prepareAndRun(day, 1, console, lang)
            prepareAndRun(day, 2, console, lang)
    else:
        prepareAndRun(day, -1, console, lang)

def testDay(day, console, lang, task):
    if lang.hasIndividualTaskRunCommands():
        if task != -1:
            prepareAndTest(day, task, console, lang)
        else:
            prepareAndTest(day, 1, console, lang)
            prepareAndTest(day, 2, console, lang)
    else:
        prepareAndTest(day, -1, console, lang)

def prepareAndRun(day, task, console, lang):
    prepareTask(day, task, console, lang)
    runTask(day, task, console, lang)

def prepareAndTest(day, task, console, lang):
    prepareTask(day, task, console, lang)
    runTask(day, task, console, lang, True)

def prepareTask(day, task, console, lang):
    taskStr = f"Preparing Task {task}"
    if task < 0:
        taskStr = f"Preparing all tasks"
    preTasks = lang.getPreRunCommand(day, task, f"./src/day{str(day).rjust(2, '0')}")
    if len(preTasks) <= 0:
        return
    console.rule(f"[yellow]{taskStr}")
    for preTask in preTasks:
        if type(preTask) is tuple:
            runPreRunCommand(console, preTask[1], preTask[0])
        else:
            runPreRunCommand(console, preTask)

def runTask(day, task, console, lang, test=False):
    taskStr = f"Task {task}"
    if task < 0:
        taskStr = "All tasks"
    console.rule(f"[yellow]{taskStr}")
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
        console.log(f"[green]{taskStr} finished execution successfully in {timeDiff}s.")
    else:
        console.log(f"[bold red]{taskStr} failed execution in {timeDiff}s!")

def runPreRunCommand(console, command, name=""):
    p = subprocess.Popen(command, shell=True, stdout=subprocess.PIPE, stderr=subprocess.STDOUT)
    p.wait()

    taskName = name or command

    if p.returncode == 0:
        console.log(f"[green]{taskName}... :heavy_check_mark:")
    else:
        console.log(f"[red bold]{taskName}... :cross_mark:")
        console.log(f"Task failed with code {p.returncode}")
        exit(1)