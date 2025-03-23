import java.util.Scanner;

public class Test {
    public static void main(String[] args) {
        Scanner reader = new Scanner(System.in);

        System.out.println("Введите кол-во частиц для 1ой вселенной: ");
        int x = reader.nextInt();
        Universe universe1 = new Universe(x);

        System.out.println("Введите кол-во частиц для 2ой вселенной: ");
        x = reader.nextInt();
        Universe universe2 = new Universe(x);

        System.out.printf("Расстояние между вселенными: %f\n", universe1.distantToUniversal(universe2));
        
        reader.close();
    }
}
