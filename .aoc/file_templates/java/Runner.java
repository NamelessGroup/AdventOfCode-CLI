import java.io.File;
import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;

public class Runner {

    public static void main(String[] args) {
        if (args.length < 2) {
            System.out.println("Error with the run command");
            return;
        }

        // Load input file
        File file = new File(args[1] + (args.length >= 3 && args[1].equals("test") ? "test" : "input"));
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
        switch (args.length > 0 ? Integer.parseInt(args[0]) : 1) {
            case 1 -> Task1.main(lines.toArray(new String[0]));
            case 2 -> Task2.main(lines.toArray(new String[0]));
            default -> System.out.println("Task not found");
        }
    }
}
