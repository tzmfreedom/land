/*
 [The "BSD licence"]
 Copyright (c) 2013 Terence Parr, Sam Harwell
 All rights reserved.

 Redistribution and use in source and binary forms, with or without
 modification, are permitted provided that the following conditions
 are met:
 1. Redistributions of source code must retain the above copyright
    notice, this list of conditions and the following disclaimer.
 2. Redistributions in binary form must reproduce the above copyright
    notice, this list of conditions and the following disclaimer in the
    documentation and/or other materials provided with the distribution.
 3. The name of the author may not be used to endorse or promote products
    derived from this software without specific prior written permission.

 THIS SOFTWARE IS PROVIDED BY THE AUTHOR ``AS IS'' AND ANY EXPRESS OR
 IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES
 OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.
 IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY DIRECT, INDIRECT,
 INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT
 NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
 DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
 THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF
 THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

/** 
 *  An Apexcode grammar derived from Java 1.7 grammar for ANTLR v4.
 *  Uses ANTLR v4's left-recursive expression notation.
 *  
 *  @maintainer: Andrey Gavrikov
 *
 *  You can test with
 *
 *  $ antlr4 Apexcode.g4
 *  $ javac *.java
 *  $ grun Apexcode compilationUnit *.cls
 */
grammar apex;

// starting point for parsing a apexcode file
compilationUnit
    :   typeDeclaration EOF
    ;

typeDeclaration
    :   classOrInterfaceModifier* classDeclaration
    |   classOrInterfaceModifier* enumDeclaration
    |   classOrInterfaceModifier* interfaceDeclaration
    |   triggerDeclaration
    |   ';'
    ;

triggerDeclaration
    :   TRIGGER apexIdentifier ON apexIdentifier '(' triggerTimings ')' block
    ;

triggerTimings
    :   triggerTiming (',' triggerTiming)*
    ;

triggerTiming
    :   timing=(BEFORE | AFTER) dml=(INSERT | UPDATE | UPSERT | DELETE | UNDELETE)
    ;

modifier
    :   classOrInterfaceModifier
    |   (
        TRANSIENT
        )
    ;

classOrInterfaceModifier
    :   annotation       // class or interface
    |   (   PUBLIC     // class or interface
        |   PROTECTED  // class or interface
        |   PRIVATE    // class or interface
        |   STATIC     // class or interface
        |   ABSTRACT   // class or interface
        |   FINAL      // class only -- does not apply to interfaces
        |   GLOBAL     // class or interface
        |   WEBSERVICE // class only -- does not apply to interfaces
        |   OVERRIDE   // method only
        |   VIRTUAL    // method only
        |   TESTMETHOD    // method only
		|	APEX_WITH_SHARING // class only
		|	APEX_WITHOUT_SHARING //class only
        )
    ;

variableModifier
    :   FINAL
    |   annotation
    ;

classDeclaration
    :   CLASS apexIdentifier
        (EXTENDS apexType)?
        (IMPLEMENTS typeList)?
        classBody
    ;

enumDeclaration
    :   ENUM apexIdentifier (IMPLEMENTS typeList)?
        '{' enumConstants? ','? enumBodyDeclarations? '}'
    ;

enumConstants
    :   enumConstant (',' enumConstant)*
    ;

enumConstant
    :   annotation* apexIdentifier arguments? classBody?
    ;

enumBodyDeclarations
    :   ';' classBodyDeclaration*
    ;

interfaceDeclaration
    :   INTERFACE apexIdentifier interfaceBody
    ;

typeList
    :   apexType (',' apexType)*
    ;

classBody
    :   '{' classBodyDeclaration* '}'
    ;

interfaceBody
    :   '{' interfaceBodyDeclaration* '}'
    ;

classBodyDeclaration
    :   ';'
    |   STATIC? block
    |   modifier* memberDeclaration
    ;

memberDeclaration
    :   methodDeclaration
    |   fieldDeclaration
    |   constructorDeclaration
    |   interfaceDeclaration
    |   classDeclaration
    |   enumDeclaration
    |   propertyDeclaration
    ;

/* We use rule this even for void methods which cannot have [] after parameters.
   This simplifies grammar and we can consider void to be a type, which
   renders the [] matching as a context-sensitive issue or a semantic check
   for invalid return type after parsing.
 */
methodDeclaration
    :   OVERRIDE? (apexType|VOID) apexIdentifier formalParameters ('[' ']')*
        (THROWS qualifiedNameList)?
        (   methodBody
        |   ';'
        )
    ;

constructorDeclaration
    :   apexIdentifier formalParameters (THROWS qualifiedNameList)?
        constructorBody
    ;

fieldDeclaration
    :   apexType variableDeclarators ';'
    ;

propertyDeclaration
    :   apexType variableDeclaratorId propertyBodyDeclaration
    ;

propertyBodyDeclaration
    :   '{' propertyBlock propertyBlock? '}'
    ;

interfaceBodyDeclaration
    :   modifier* interfaceMemberDeclaration
    |   ';'
    ;

interfaceMemberDeclaration
    :   constDeclaration
    |   interfaceMethodDeclaration
    |   interfaceDeclaration
    |   classDeclaration
    |   enumDeclaration
    ;

constDeclaration
    :   apexType constantDeclarator (',' constantDeclarator)* ';'
    ;

constantDeclarator
    :   apexIdentifier ('[' ']')* '=' variableInitializer
    ;

// see matching of [] comment in methodDeclaratorRest
interfaceMethodDeclaration
    :   (apexType|VOID) apexIdentifier formalParameters ('[' ']')*
        (THROWS qualifiedNameList)?
        ';'
    ;

variableDeclarators
    :   variableDeclarator (',' variableDeclarator)*
    ;

variableDeclarator
    :   variableDeclaratorId ('=' variableInitializer)?
    ;

variableDeclaratorId
    :   apexIdentifier ('[' ']')*
    ;

variableInitializer
    :   arrayInitializer
    |   expression
    ;

arrayInitializer
    :   '{' (variableInitializer (',' variableInitializer)* (',')? )? '}'
    ;

enumConstantName
    :   apexIdentifier
    ;

apexType
    :   classOrInterfaceType typedArray*
    |   primitiveType typedArray*
    ;

typedArray
    : '[' ']'
    ;
classOrInterfaceType
    :   typeIdentifier typeArguments? ('.' typeIdentifier typeArguments? )*
    |   SET typeArguments // 'set <' has to be defined explisitly, otherwise it clashes with SET of property setter
    ;

primitiveType
    :   BOOLEAN
    |   STRING
    |   INTEGER
    |   LONG
    |   FLOAT
    |   DOUBLE
    ;

typeArguments
    :   '<' typeArgument (',' typeArgument)* '>'
    ;

typeArgument
    :   apexType
    |   '?' ((EXTENDS | SUPER) apexType)?
    ;

qualifiedNameList
    :   qualifiedName (',' qualifiedName)*
    ;

formalParameters
    :   '(' formalParameterList? ')'
    ;

formalParameterList
    :   formalParameter (',' formalParameter)* (',' lastFormalParameter)?
    |   lastFormalParameter
    ;

formalParameter
    :   variableModifier* apexType variableDeclaratorId
    ;

lastFormalParameter
    :   variableModifier* apexType '...' variableDeclaratorId
    ;

methodBody
    :   block
    ;

constructorBody
    :   block
    ;

qualifiedName
    :   apexIdentifier ('.' apexIdentifier)*
    ;

literal
    :   IntegerLiteral
    |   FloatingPointLiteral
    |   StringLiteral
    |   BooleanLiteral
    |   NullLiteral
    ;

// ANNOTATIONS

annotation
    :   '@' annotationName ( '(' ( elementValuePairs | elementValue )? ')' )?
    ;

annotationName : qualifiedName ;

elementValuePairs
    :   elementValuePair elementValuePair*
    ;

elementValuePair
    :   apexIdentifier '=' elementValue
    ;

elementValue
    :   expression
    |   annotation
    |   elementValueArrayInitializer
    ;

elementValueArrayInitializer
    :   '{' (elementValue (',' elementValue)*)? (',')? '}'
    ;


// STATEMENTS / BLOCKS

block
    :   '{' blockStatement* '}'
    ;

blockStatement
    :   localVariableDeclarationStatement
    |   statement
    |   typeDeclaration
    ;

localVariableDeclarationStatement
    :    localVariableDeclaration ';'
    ;

localVariableDeclaration
    :   variableModifier* apexType variableDeclarators
    ;

statement
    :   block
    |   IF parExpression statement (ELSE statement)?
    |   SWITCH ON expression '{' whenStatements  (WHEN ELSE block)? '}'
    |   FOR '(' forControl ')' statement
    |   WHILE parExpression statement
    |   DO statement WHILE parExpression
    |   TRY block (catchClause+ finallyBlock? | finallyBlock)
    |   RETURN expression? ';'
    |   THROW expression ';'
    |   BREAK apexIdentifier? ';'
    |   CONTINUE apexIdentifier? ';'
    |   ';'
    |   statementExpression ';'
    |   apexDbExpression ';'
    |   SYSTEM '.' RUNAS '(' expression ')' block
    ;

propertyBlock
	:	modifier* (getter | setter)
	;

getter
 : GET (';' | methodBody)
 ;

setter
 : SET (';' | methodBody)
 ;


catchClause
    :   CATCH '(' variableModifier* catchType apexIdentifier ')' block
    ;

catchType
    :   qualifiedName ('|' qualifiedName)*
    ;

finallyBlock
    :   FINALLY block
    ;

whenStatements
    :   whenStatement whenStatement*
    ;

whenStatement
    :   WHEN whenExpression block
    ;

whenExpression
    :   literal (',' literal)*
    |   apexType apexIdentifier
    ;

forControl
    :   enhancedForControl
    |   forInit? ';' expression? ';' forUpdate?
    ;

forInit
    :   localVariableDeclaration
    |   expressionList
    ;

enhancedForControl
    :   variableModifier* apexType variableDeclaratorId ':' expression
    ;

forUpdate
    :   expressionList
    ;

// EXPRESSIONS

parExpression
    :   '(' expression ')'
    ;

expressionList
    :   expression (',' expression)*
    ;

statementExpression
    :   expression
    ;

constantExpression
    :   expression
    ;

apexDbExpressionShort
    :   dml=(INSERT | UPSERT | UPDATE | DELETE | UNDELETE) expression
    |   UPSERT expression apexIdentifier
    ;


apexDbExpression
	: apexDbExpressionShort
	;
	
expression
    :   primary                                                 # PrimaryExpression
    |   expression '.' apexIdentifier                           # FieldAccess
    |   expression '.' explicitGenericInvocation                # OpExpression
    |   expression '[' expression ']'                           # ArrayAccess
    |   expression '(' expressionList? ')'                      # MethodInvocation
    |   NEW creator                                             # NewObjectExpression
    |   '(' apexType ')' expression                                 # CastExpression
    |   expression op=('++' | '--')                             # PostUnaryExpression
    |   op=('++'|'--') expression                               # PreUnaryExpression
    |   op=('+'|'-'|'!') expression                             # UnaryExpression
    |   expression op=('*'|'/'|'%') expression                  # OpExpression
    |   expression op=('+'|'-') expression                      # OpExpression
    |   expression (op+='<' op+='<' | op+='>' op+='>' op+='>' | op+='>' op+='>') expression # ShiftExpression
    |   expression op=('<=' | '>=' | '>' | '<') expression      # OpExpression
    |   expression op=INSTANCEOF apexType                       # InstanceofExpression
    |   expression op=('===' | '==' | '!=') expression          # OpExpression
    |   expression op='&' expression                            # OpExpression
    |   expression op='^' expression                            # OpExpression
    |   expression op='|' expression                            # OpExpression
    |   expression op='&&' expression                           # OpExpression
    |   expression op='||' expression                           # OpExpression
    |   expression op='?' expression ':' expression             # TernalyExpression
    |   <assoc=right> expression
        op=(   '='
        |   '+='
        |   '-='
        |   '*='
        |   '/='
        |   '&='
        |   '|='
        |   '^='
        |   '>>='
        |   '>>>='
        |   '<<='
        |   '%='
        )
        expression                                             # OpExpression
    ;

primary
    :   '(' expression ')'
    |   THIS
    |   SUPER
    |   literal
    |   apexIdentifier
    |   apexType '.' CLASS
    |   VOID '.' CLASS
    |   nonWildcardTypeArguments (explicitGenericInvocationSuffix | THIS arguments)
    |   soqlLiteral
    |   soslLiteral
    |   primitiveType
    ;

creator
    :   nonWildcardTypeArguments createdName classCreatorRest
    |   createdName (arrayCreatorRest | classCreatorRest | mapCreatorRest | setCreatorRest)
    ;

createdName
    :   apexIdentifier typeArgumentsOrDiamond? ('.' apexIdentifier typeArgumentsOrDiamond?)*
    |   primitiveType
    |   SET typeArgumentsOrDiamond // 'set <' has to be defined explisitly, otherwise it clashes with SET of property setter
    ;

innerCreator
    :   apexIdentifier nonWildcardTypeArgumentsOrDiamond? classCreatorRest
    ;

arrayCreatorRest
    :   typedArray typedArray* arrayInitializer
    |   '[' expression ']' ('[' expression ']')* typedArray*
    ;

mapCreatorRest
    :   '{' '}'
    |   '{' mapKey '=>' mapValue (',' mapKey '=>' mapValue )* '}'
    ;

mapKey
    : apexIdentifier
    | expression
    ;

mapValue
    : literal
    | expression
    ;

setCreatorRest
	: '{' setValue (',' setValue)* '}'
	;

setValue
    : literal
    | expression
    ;

classCreatorRest
    :   arguments classBody?
    ;

explicitGenericInvocation
    :   nonWildcardTypeArguments explicitGenericInvocationSuffix
    ;

nonWildcardTypeArguments
    :   '<' typeList '>'
    ;

typeArgumentsOrDiamond
    :   '<' '>'
    |   typeArguments
    ;

nonWildcardTypeArgumentsOrDiamond
    :   '<' '>'
    |   nonWildcardTypeArguments
    ;

superSuffix
    :   arguments
    |   '.' apexIdentifier arguments?
    ;

explicitGenericInvocationSuffix
    :   SUPER superSuffix
    |   apexIdentifier arguments
    ;

arguments
    :   '(' expressionList? ')'
    ;

// Apex - SOQL literal

soqlLiteral
    : '[' query ']'
	;

query
    : selectClause
      fromClause
      whereClause?
      withClause?
      groupClause?
      orderClause?
      limitClause?
      offsetClause?
      viewClause?
    ;

selectClause
    : SELECT fieldList
    ;

fieldList
    : selectField (COMMA selectField)*
    ;

selectField
    : soqlField
    | subquery
    | TYPEOF soqlField
      (WHEN apexIdentifier THEN fieldList)+
      ELSE fieldList
      END
    ;

fromClause
    : FROM apexIdentifier (USING SCOPE filterScope)?
    ;

filterScope
    :
    ;

soqlField
    : (apexIdentifier DOT)* apexIdentifier  # SoqlFieldReference
    | apexIdentifier LPAREN (soqlField (COMMA soqlField)*)? RPAREN # SoqlFunctionCall
    ;

subquery
    : query
    ;

whereClause
    : WHERE whereFields
    ;

whereFields
    : whereField
    | whereFields and_or=(SOQL_AND|SOQL_OR) whereFields
    ;

whereField
    :
       SOQL_NOT?
       soqlField
       op=(
         '='
         | '<'
         | '>'
         | '<='
         | '>='
         | '!='
         | '<>'
         | LIKE
         | IN
       )
       soqlValue
    |  '(' whereFields ')'
    ;

limitClause
    :  LIMIT (IntegerLiteral | bindVariable)
    ;

orderClause
    :  ORDER BY soqlField (',' soqlField)* asc_desc=(ASC | DESC)? (NULLS nulls=(LAST | FIRST))?
    ;

bindVariable
    :  COLON expression
    ;

soqlValue
    :  literal
    |  bindVariable
    |  apexIdentifier COLON literal
    ;

withClause
    :  WITH DATA CATEGORY soqlFilteringExpression
    ;

soqlFilteringExpression
    :
    ;

groupClause
    :  GROUP BY soqlField (',' soqlField)* (HAVING havingConditionExpression)?
    ;

havingConditionExpression
    : whereFields
    ;

offsetClause
    :  OFFSET (IntegerLiteral | bindVariable)
    ;

viewClause
    : FOR (VIEW | REFERENCE) (UPDATE (TRACKING | VIEWSTAT))?
    ;

// Apex - SOSL literal

soslLiteral
    : '[' soslQuery ']'
	;

soslQuery
    : FIND literal IN ALL FIELDS RETURNING soslReturningObject (',' soslReturningObject)*
    ;

soslReturningObject
    : Identifier ('(' Identifier (',' Identifier)* ')')?
    ;
apexIdentifier
    :  Identifier
    |  GET
    |  SET
    |  DATA
    |  GROUP
    |  DELETE
    |  INSERT
    |  UPDATE
    |  UNDELETE
    |  UPSERT
    |  SCOPE
    |  CATEGORY
    |  REFERENCE
    |  OFFSET
    |  THEN
    |  FIND
    |  RETURNING
    |  ALL
    |  FIELDS
    |  RUNAS
    |  SYSTEM
    |  primitiveType
    ;

typeIdentifier
    :  Identifier
    |  GET
    |  SET
    |  DATA
    |  GROUP
    |  SCOPE
    |  CATEGORY
    |  REFERENCE
    |  OFFSET
    |  THEN
    |  FIND
    |  RETURNING
    |  ALL
    |  FIELDS
    |  SYSTEM
    ;
// LEXER

// ?3.9 Keywords

OVERRIDE      : O V E R R I D E;
VIRTUAL       : V I R T U A L;
SET           : S E T;
GET           : G E T;
ABSTRACT      : A B S T R A C T;
BOOLEAN       : B O O L E A N;
BREAK         : B R E A K;
CATCH         : C A T C H;
CLASS         : C L A S S;
CONST         : C O N S T;
CONTINUE      : C O N T I N U E;
DEFAULT       : D E F A U L T;
DO            : D O;
DOUBLE        : D O U B L E;
ELSE          : E L S E;
ENUM          : E N U M;
EXTENDS       : E X T E N D S;
FINAL         : F I N A L;
FINALLY       : F I N A L L Y;
FLOAT         : F L O A T;
FOR           : F O R;
IF            : I F;
GOTO          : G O T O;
IMPLEMENTS    : I M P L E M E N T S;
IMPORT        : I M P O R T;
INSTANCEOF    : I N S T A N C E O F;
INTEGER       : I N T E G E R;
STRING        : S T R I N G;
INTERFACE     : I N T E R F A C E;
LONG          : L O N G;
NATIVE        : N A T I V E;
NEW           : N E W;
PACKAGE       : P A C K A G E;
PRIVATE       : P R I V A T E;
PROTECTED     : P R O T E C T E D;
PUBLIC        : P U B L I C;
RETURN        : R E T U R N;
STATIC        : S T A T I C;

SUPER         : S U P E R;
SYNCHRONIZED  : S Y N C H R O N I Z E D;
THIS          : T H I S;
THROW         : T H R O W;
THROWS        : T H R O W S;
TRANSIENT     : T R A N S I E N T;
TRY           : T R Y;
VOID          : V O I D;
VOLATILE      : V O L A T I L E;
WHILE         : W H I L E;
SWITCH        : S W I T C H;
WHEN          : W H E N;

// Apexcode specific
GLOBAL	      : G L O B A L;
WEBSERVICE    : W E B S E R V I C E;
APEX_WITH_SHARING :    W I T H SPACE S H A R I N G;
APEX_WITHOUT_SHARING : W I T H O U T SPACE S H A R I N G;
SELECT        : S E L E C T;
FROM          : F R O M;
WHERE         : W H E R E;
LIMIT         : L I M I T;
ORDER         : O R D E R;
BY            : B Y;
ASC           : A S C;
DESC          : D E S C;
WITH          : W I T H;
TYPEOF        : T Y P E O F;
REFERENCE     : R E F E R E N C E;
VIEW          : V I E W;
VIEWSTAT      : V I E W S T A T;
TRACKING      : T R A C K I N G;
OFFSET        : O F F S E T;
IN            : I N;
END           : E N D;
USING         : U S I N G;
DATA          : D A T A;
CATEGORY      : C A T E G O R Y;
GROUP         : G R O U P;
HAVING        : H A V I N G;
NULLS         : N U L L S;
FIRST         : F I R S T;
LAST          : L A S T;
SCOPE         : S C O P E;
ROLLUP        : R O L L U P;
CUBE          : C U B E;
LIKE          : L I K E;
THEN          : T H E N;
INSERT     : I N S E R T;
UPSERT     : U P S E R T;
UPDATE     : U P D A T E;
DELETE     : D E L E T E;
UNDELETE   : U N D E L E T E;
SOQL_AND   : A N D;
SOQL_OR    : O R;
SOQL_NOT   : N O T;
FIND       : F I N D;
FIELDS     : F I E L D S;
RETURNING  : R E T U R N I N G;
ALL        : A L L;
TESTMETHOD   : T E S T M E T H O D;
TRIGGER       : T R I G G E R;
ON            : O N;
BEFORE        : B E F O R E;
AFTER         : A F T E R;
RUNAS         : R U N A S;
SYSTEM        : S Y S T E M;


// ?3.10.1 Integer Literals

IntegerLiteral
    :   DecimalIntegerLiteral
    |   HexIntegerLiteral
    |   OctalIntegerLiteral
    |   BinaryIntegerLiteral
    ;

fragment
DecimalIntegerLiteral
    :   DecimalNumeral IntegerTypeSuffix?
    ;

fragment
HexIntegerLiteral
    :   HexNumeral IntegerTypeSuffix?
    ;

fragment
OctalIntegerLiteral
    :   OctalNumeral IntegerTypeSuffix?
    ;

fragment
BinaryIntegerLiteral
    :   BinaryNumeral IntegerTypeSuffix?
    ;

fragment
IntegerTypeSuffix
    :   [lL]
    ;

fragment
DecimalNumeral
    :   Digits
    ;

fragment
Digits
    :   Digit (DigitOrUnderscore* Digit)?
    ;

fragment
Digit
    :   '0'
    |   NonZeroDigit
    ;

fragment
NonZeroDigit
    :   [1-9]
    ;

fragment
DigitOrUnderscore
    :   Digit
    |   '_'
    ;

fragment
Underscores
    :   '_'+
    ;

fragment
HexNumeral
    :   '0' [xX] HexDigits
    ;

fragment
HexDigits
    :   HexDigit (HexDigitOrUnderscore* HexDigit)?
    ;

fragment
HexDigit
    :   [0-9a-fA-F]
    ;

fragment
HexDigitOrUnderscore
    :   HexDigit
    |   '_'
    ;

fragment
OctalNumeral
    :   '0' Underscores? OctalDigits
    ;

fragment
OctalDigits
    :   OctalDigit (OctalDigitOrUnderscore* OctalDigit)?
    ;

fragment
OctalDigit
    :   [0-7]
    ;

fragment
OctalDigitOrUnderscore
    :   OctalDigit
    |   '_'
    ;

fragment
BinaryNumeral
    :   '0' [bB] BinaryDigits
    ;

fragment
BinaryDigits
    :   BinaryDigit (BinaryDigitOrUnderscore* BinaryDigit)?
    ;

fragment
BinaryDigit
    :   [01]
    ;

fragment
BinaryDigitOrUnderscore
    :   BinaryDigit
    |   '_'
    ;

// ?3.10.2 Floating-Point Literals

FloatingPointLiteral
    :   DecimalFloatingPointLiteral
    |   HexadecimalFloatingPointLiteral
    ;

fragment
DecimalFloatingPointLiteral
    :   Digits '.' Digits? ExponentPart? FloatTypeSuffix?
    |   '.' Digits ExponentPart? FloatTypeSuffix?
    |   Digits ExponentPart FloatTypeSuffix?
    |   Digits FloatTypeSuffix
    ;

fragment
ExponentPart
    :   ExponentIndicator SignedInteger
    ;

fragment
ExponentIndicator
    :   [eE]
    ;

fragment
SignedInteger
    :   Sign? Digits
    ;

fragment
Sign
    :   [+-]
    ;

fragment
FloatTypeSuffix
    :   [fFdD]
    ;

fragment
HexadecimalFloatingPointLiteral
    :   HexSignificand BinaryExponent FloatTypeSuffix?
    ;

fragment
HexSignificand
    :   HexNumeral '.'?
    |   '0' [xX] HexDigits? '.' HexDigits
    ;

fragment
BinaryExponent
    :   BinaryExponentIndicator SignedInteger
    ;

fragment
BinaryExponentIndicator
    :   [pP]
    ;

// ?3.10.3 Boolean Literals

BooleanLiteral
    :   T R U E
    |   F A L S E
    ;

// ?3.10.5 String Literals

StringLiteral
    :   QUOTE StringCharacters? QUOTE
    ;

fragment
StringCharacters
    :   StringCharacter+
    ;

fragment
StringCharacter
    :   ~['\\]
    |   EscapeSequence
    ;

// ?3.10.6 Escape Sequences for Character and String Literals

fragment
EscapeSequence
    :   '\\' [btnfr"'\\]
    |   OctalEscape
    |   UnicodeEscape
    ;

fragment
OctalEscape
    :   '\\' OctalDigit
    |   '\\' OctalDigit OctalDigit
    |   '\\' ZeroToThree OctalDigit OctalDigit
    ;

fragment
UnicodeEscape
    :   '\\' 'u' HexDigit HexDigit HexDigit HexDigit
    ;

fragment
ZeroToThree
    :   [0-3]
    ;

// ?3.10.7 The Null Literal

NullLiteral :   N U L L;


// ?3.11 Separators

LPAREN          : '(';
RPAREN          : ')';
LBRACE          : '{';
RBRACE          : '}';
LBRACK          : '[';
RBRACK          : ']';
SEMI            : ';';
COMMA           : ',';
DOT             : '.';

// ?3.12 Operators

ASSIGN          : '=';
GT              : '>';
LT              : '<';
BANG            : '!';
TILDE           : '~';
QUESTION        : '?';
COLON           : ':';
EQUAL           : '==';
T_EQUAL         : '===';
LE              : '<=';
GE              : '>=';
NOTEQUAL        : '!=';
AND             : '&&';
OR              : '||';
INC             : '++';
DEC             : '--';
ADD             : '+';
SUB             : '-';
MUL             : '*';
DIV             : '/';
BITAND          : '&';
BITOR           : '|';
CARET           : '^';
MOD             : '%';

ADD_ASSIGN      : '+=';
SUB_ASSIGN      : '-=';
MUL_ASSIGN      : '*=';
DIV_ASSIGN      : '/=';
AND_ASSIGN      : '&=';
OR_ASSIGN       : '|=';
XOR_ASSIGN      : '^=';
MOD_ASSIGN      : '%=';
LSHIFT_ASSIGN   : '<<=';
RSHIFT_ASSIGN   : '>>=';
URSHIFT_ASSIGN  : '>>>=';
LAMBDA_LIKE     : '=>';


// ?3.8 Identifiers (must appear after all keywords in the grammar)

Identifier
    :   JavaLetter JavaLetterOrDigit*
    ;

fragment
JavaLetter
    :   [a-zA-Z$_] // these are the "java letters" below 0xFF
    ;

fragment
JavaLetterOrDigit
    :   [a-zA-Z0-9$_] // these are the "java letters or digits" below 0xFF
    ;

//
// Additional symbols not defined in the lexical specification
//

AT : '@';
ELLIPSIS : '...';

//
// Whitespace and comments
//

WS  :  [ \t\r\n\u000C]+ -> skip
    ;

APEXDOC_COMMENT
    :   '/**' [\r\n] .*? '*/' -> skip
    ;

APEXDOC_COMMENT_START
    :   '/**' -> skip
    ;

COMMENT
    :   '/*' .*? '*/' -> skip
    ;

COMMENT_START
    :   '/*' -> skip
    ;

LINE_COMMENT
    :   '//' (~[\r\n])* -> skip
    ;

//
// Unexpected token for non recognized elements
//

QUOTE	:	'\'' -> skip;

// characters

fragment A : [aA];
fragment B : [bB];
fragment C : [cC];
fragment D : [dD];
fragment E : [eE];
fragment F : [fF];
fragment G : [gG];
fragment H : [hH];
fragment I : [iI];
fragment J : [jJ];
fragment K : [kK];
fragment L : [lL];
fragment M : [mM];
fragment N : [nN];
fragment O : [oO];
fragment P : [pP];
fragment Q : [qQ];
fragment R : [rR];
fragment S : [sS];
fragment T : [tT];
fragment U : [uU];
fragment V : [vV];
fragment W : [wW];
fragment X : [xX];
fragment Y : [yY];
fragment Z : [zZ];
fragment SPACE : ' ';
