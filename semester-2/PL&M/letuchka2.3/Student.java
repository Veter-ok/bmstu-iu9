public class Student {
    private String name;
    private String surname;
    private String name2;
    private int groupNumber;
    private int rank;

    public Student(String name, String surname, String name2, int groupNumber, int rank) {
        this.name = name;
        this.surname = surname;
        this.name2 = name2;
        this.groupNumber = groupNumber;
        this.rank = rank;
    }

    public String fulInfo(){
        String x = this.surname + " " + this.name + " " + this.name2 + " группа: " + this.groupNumber + " " + this.rank;
        return x;
    }
}
