public with sharing class Foo extends Bar implements Baz {
    public static String action() {
        Integer i = 0;
        for (Integer i = 0; i < 3; i++) {
            System.debug(i);
        }
        while(i < 5) {
            System.debug(i*10);
            i++;
            continue;
            System.debug(true);
        }
        if (true) {
            System.debug(true);
        } else {
            System.debug(false);
        }
        if (false) {
            System.debug(true);
        } else {
            System.debug(false);
        }
    }
}