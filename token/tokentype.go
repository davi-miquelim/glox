package token

const (
    // Single character tokens
    LeftParen = iota
    RightParen
    LeftBrace
    RightBrace
    Comma
    Dot
    Minus
    Plus
    SemiColon
    Slash
    Star
    
    // One or two character tokens
    Bang
    BangEqual
    Equal
    EqualEqual
    Greater
    GreaterEqual
    Less
    LessEqual

    // Literals
    Identifier
    String
    Number

    // Keywords
    And
    Class
    Else
    False
    Fun
    For
    If
    Null
    Or
    Print
    Return
    Super
    This
    True
    Var
    While

    Eof
)
