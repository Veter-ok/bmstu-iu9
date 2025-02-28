public class Student {
    private int undoneWorks;
    
    public Student(int undoneWorks) {
        this.undoneWorks = undoneWorks;
        System.out.println("Конструктор: undoneWorks="+this.undoneWorks);
    }

    public int getUndoneWorks() {
        return this.undoneWorks;
    }

    public void setUndoneWorks(int newVal) {
        this.undoneWorks = newVal;
    }

    public static void main(String[] args) {
        Student a = new Student(0);
        System.out.println("Дефолтное значение поля "+ a.getUndoneWorks());
        a.setUndoneWorks(100);
        System.out.println("Новое значение поля "+ a.getUndoneWorks());
    }
}
