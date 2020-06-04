grammar cond;

conditionExpr: singleExpr;

singleExpr
   : prefix='!' expr=singleExpr 							#prefix
   | left=singleExpr op=('<' | '>' | '<=' | '>=' | '==' | '!=' | 'in' | 'not in') right=singleExpr	#test
   | left=singleExpr op=('&&' | '||') right=singleExpr					#logic
   | NUMBER 			#number
   | BOOL				#bool
   | NULL 				#nil
   | ID			        #identifier
   | STRING 			#string
   | '(' conditionExpr')' 		#paren
   ;

BOOL
   : 'true'
   | 'false'
   ;

NULL
   : 'null'
   ;

NUMBER
   : '-'? INT
   ;

ID
   : [a-z][.a-zA-Z0-9_-]*
   ;

STRING
   : '"' (ESC | SAFECODEPOINT)* '"'
   ;


WS
   : [ \r\n\t] -> skip
   ;

fragment ESC
   : '\\' (["\\/bfnrt] | UNICODE)
   ;
fragment UNICODE
   : 'u' HEX HEX HEX HEX
   ;
fragment HEX
   : [0-9a-fA-F]
   ;
fragment SAFECODEPOINT
   : ~ ["\\\u0000-\u001F]
   ;

fragment INT
   : '0' | [1-9] [0-9]*
   ;