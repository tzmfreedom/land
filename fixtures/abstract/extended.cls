public class Extended extends AbstractClass {
    public override String hello() {
        return 'hello';
    }

    public static void main() {
        AbstractClass klass = new Extended();
        System.debug(klass.hello());
        klass = new Extended2();
        System.debug(klass.hello());
    }
}
