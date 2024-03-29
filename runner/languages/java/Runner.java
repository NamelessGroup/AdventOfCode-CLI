import java.io.File;
import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;

public class Runner {

    public static void main(String[] args) {
        if (args.length < 1) {
            System.out.println("Error with the run command");
            return;
        }

        // Load input file
        File file = new File((args.length >= 2 && args[1].equals("test") ? "test.in" : "solve.in"));
        if (!file.exists()) {
            System.out.println("File not found");
            return;
        }

        // Parse input file
        List<String> lines = new ArrayList<>();
        try (Scanner scanner = new Scanner(file)) {
            while (scanner.hasNextLine()) {
                lines.add(scanner.nextLine());
            }
        } catch (Exception e) {
            System.out.println(e.getMessage());
            return;
        }

        // Start task
        Tasks t = new Tasks();
        switch (args.length > 0 ? Integer.parseInt(args[0]) : 1) {
            case 1 -> t.taskOne(lines.toArray(new String[0]));
            case 2 -> t.taskTwo(lines.toArray(new String[0]));
            default -> System.out.println("Task not found");
        }
    }
}
