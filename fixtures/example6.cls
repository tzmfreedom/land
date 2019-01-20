public without sharing class Builtin {
    public static void blob() {
        System.debug(EncodingUtil.base64Encode(Blob.valueOf('あいうえお')));
        System.debug(String.valueOf(EncodingUtil.base64Decode('44GC44GE44GG44GI44GK')));
        System.debug(EncodingUtil.urlEncode('あいうえお', 'UTF-8'));
    }

    public static void string() {
        System.assertEquals('foofuga', 'hogefuga'.replace('hoge', 'foo'));
        System.assertEquals('hogafuga', 'hogefuge'.replaceAll('ge', 'ga'));
        System.assertEquals(3, 'hogefuge'.indexOf('ef'));
        System.assertEquals(true, ''.isBlank());
        System.assertEquals(true, '  '.isBlank());
        System.assertEquals(false, 'a'.isBlank());
        System.assertEquals(false, ''.isNotBlank());
        System.assertEquals(false, '  '.isNotBlank());
        System.assertEquals(true, 'a'.isNotBlank());
        System.assertEquals(true, ''.isEmpty());
        System.assertEquals(false, '  '.isEmpty());
        System.assertEquals(false, 'a'.isEmpty());
        System.assertEquals(false, ''.isNotEmpty());
        System.assertEquals(true, '  '.isNotEmpty());
        System.assertEquals(true, 'a'.isNotEmpty());
        System.assertEquals(true, 'a'.equals('a'));
        System.assertEquals(false, 'a'.equals('b'));
        System.assertEquals(false, 'hoge'.contains('a'));
        System.assertEquals(true, 'hoge'.contains('g'));
        System.assertEquals(4, 'hoge'.length());
        System.debug('foo,bar,baz'.split(','));
        System.assertEquals('burger', 'hamburger'.substring(3));
        System.assertEquals('urge', 'hamburger'.substring(4, 8));
        System.assertEquals('hamburger', 'HamBurger'.toLowerCase());
        System.assertEquals('HAMBURGER', 'HamBurger'.toUpperCase());
        System.assertEquals('a,bb,ccc', String.join(new List<String>{'a', 'bb', 'ccc'}, ','));
    }

    public static void http() {
        HttpRequest req = new HttpRequest();
        req.setHeader('Foo', 'Bar');
        req.setEndpoint('https://httpbin.org/post');
        req.setMethod('post');
        req.setBody('this is body');
        Http http = new Http();
        HTTPResponse res = http.send(req);
        System.debug(res.getBody());
    }
}
