public with sharing class Foo extends Bar implements Baz {
    public static String action() {
        List<Account> accounts = [
            SELECT
                id,
                Name
            FROM
                Account
            WHERE
                Name != ''
            ORDER BY
                id,
                Name ASC
            LIMIT
                10
            OFFSET
                1000
        ];
        System.debug(accounts);
    }
}