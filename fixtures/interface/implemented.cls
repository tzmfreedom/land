public class Implemented implements InterfaceClass {
    public Double discount() {
        return 1.2;
    }

    public static void main() {
        InterfaceClass h = new Implemented();
        System.debug(h.discount());
        h = new Implemented2();
        System.debug(h.discount());
    }
}
