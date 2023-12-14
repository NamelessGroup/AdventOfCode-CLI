import { readFile } from "fs/promises";
import { taskOne, taskTwo } from "./task";

async function main() {
    let lastArgument = process.argv.pop() as string;
    let taskNumber = 1;
    let isTest = false;

    if (lastArgument === "test") {
        isTest = true;
        taskNumber = parseInt(process.argv.pop() as string);
    } else {
        taskNumber = parseInt(lastArgument)
    }

    const fileToLoad = isTest ? "test.in" : "solve.in";
    const fileContents = await readFile(fileToLoad, "utf-8")
    
    const lines = fileContents.trimEnd().split("\n");

    if (taskNumber === 1) {
        await taskOne(lines);
    }
    if (taskNumber === 2) {
        await taskTwo(lines);
    }
}

void main();