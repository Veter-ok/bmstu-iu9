import java.util.Arrays;

public class Stack<T> {
    private int count = 0;
    private Object[] buf = new Object[16];

    public boolean empty() {
        return count == 0; 
    }

    public void push(T x) {
        if (count == buf.length) {
            buf = Arrays.copyOf(buf, buf.length*2);
        }
        buf[count++] = x;
    }

    public T pop() {
        if (empty()){
            throw new RuntimeException("underflow");
        }
        return (T)buf[--count];
    }

    public static void main(String[] args) {
        Stack<Student> stackStud = new Stack<Student>();
        stackStud.push(new Student("rrr", "ttt", "yyy", 100, 200));
        stackStud.push(new Student("Rodion", "Lavrov", "Den", 1, 1));
        stackStud.push(new Student("Vova", "Gorin", "ytoot", 2, 1));
        while (!stackStud.empty()) {
            System.out.println(stackStud.pop().fulInfo());
        }
    }
}