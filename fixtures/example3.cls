public with sharing class Foo {
    public static void action() {
        Integer i = 0;
        for (Integer j = 0; j < 3; j++) {
            System.debug(j);
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