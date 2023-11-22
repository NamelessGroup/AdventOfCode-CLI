import System.IO
import System.Environment
import Task

main = do
    args <- getArgs
    let task = head args
    let test = head (tail args)

    let fileToRead = if test == "test" then "test.in" else "solve.in"

    fileContents <- readFile fileToRead
    let fileLines = lines fileContents

    let result = if task == "1" then taskOne fileLines else taskTwo fileLines

    putStrLn result