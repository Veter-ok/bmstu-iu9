public class Test {
    public static void main(String[] args) {
        Number a = new Number(Double.parseDouble(args[0]));
        Number b = new Number(Double.parseDouble(args[1]));
        if (args[2].equals("mul")) {
            Number res = a.multiply(b);
            System.out.println(res);
        }else if (args[2].equals("add")) {
            Number res = a.add(b);
            System.out.println(res);
        }
    }
}
