public with sharing class Foo {
    public static String action() {
        System.debug(1 + 2 + 3);
        System.debug(1 + 2 * 3);
        System.debug((1 + 2) * 3);
        System.debug('hoge');
        System.debug('foo' + '/' + 'bar');
        System.debug(1.2 * 1.3);
        return 'hello';
    }
}
