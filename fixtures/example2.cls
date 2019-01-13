public with sharing class Foo {
    public Integer i = 0;
    public Boolean b = false;
    public String s = 'foo';
    public Double d = 1.23;

    public Foo(Integer i) {
        this.i = i;
    }

    public String foo() {
        String i = 'foo';
        return i;
    }

    public static void action() {
        Foo f = new Foo(100);
        System.debug(f);
        System.debug(new Foo(200).i);
        System.debug(f.foo());
        System.debug(new Foo(123).foo());
        System.debug(f.i);
        System.debug(f.b);
        System.debug(f.s);
        System.debug(f.d);
    }
}