import java.util.Stack;

class A {
    public String valA;
    public A(String vA) {valA = vA;}
    public String toString() {return String.format("%s", this.valA);}
}

class B extends A {
    public String valB;
    public B(String vA, String vB) {super(vA); valB = vB;}
    public String toString() {return String.format("%s, %s", this.valA, this.valB);}
}

class C extends A {
    public String valC;
    public C(String vA, String vC) {super(vA); valC = vC;}
    public String toString() {return String.format("%s, %s", this.valA, this.valC);}
}

class D extends B {
    public String valD;
    public D(String vA, String vB, String vD) {super(vA, vB); valD = vD;}
    public String toString() {return String.format("%s, %s, %s", this.valA, this.valB, this.valD);}
}

class E extends B {
    public String valE;
    public E(String vA, String vB, String vE) {super(vA, vB); valE = vE;}
    public String toString() {return String.format("%s, %s, %s", this.valA, this.valB, this.valE);}
}

class F extends C {
    public String valF;
    public F(String vA, String vC, String vF) {super(vA, vC); valF = vF;}
    public String toString() {return String.format("%s, %s, %s", this.valA, this.valC, this.valF);}
}

class G extends D {
    public String valG;
    public G(String vA, String vB, String vD, String vG) {super(vA, vB, vD); valG = vG;}
    public String toString() {return String.format("%s, %s, %s, %s", this.valA, this.valB, this.valD, this.valG);}
}

class J extends D {
    public String valJ;
    public J(String vA, String vB, String vD, String vJ) {super(vA, vB, vD); valJ = vJ;}
    public String toString() {return String.format("%s, %s, %s, %s", this.valA, this.valB, this.valD, this.valJ);}
}

class K extends D {
    public String valK;
    public K(String vA, String vB, String vD, String vK) {super(vA, vB, vD); valK = vK;}
    public String toString() {return String.format("%s, %s, %s, %s", this.valA, this.valB, this.valD, this.valK);}
}

public class Task {
    public static void pushReversed(
            Stack<? super D> dest,
            Stack<? extends D> src)
    {
        while(!src.empty()){
            dest.push(src.pop());
        }

        while (!dest.empty()) {
            System.out.println(dest.pop());
        }
    }

    public static void main(String[] args) {
        Stack<A> dest = new Stack<>();
        Stack<D> src = new Stack<>();
        src.push(new D("A1", "B1", "D1"));
        src.push(new D("A2", "B2", "D2"));
        src.push(new D("A3", "B3", "D3"));

        // Stack<A> dest = new Stack<>();
        // Stack<J> src = new Stack<>();
        // src.push(new J("A1", "B1", "D1", "J1"));
        // src.push(new J("A2", "B2", "D2", "J2"));
        // src.push(new J("A3", "B3", "D3", "J3"));

        pushReversed(dest, src);
    }
}