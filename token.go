package token

const (
    // Single character tokens
    leftParen = iota
    rightParen
    leftBrace
    rightBrace
    comma
    dot
    minus
    plus
    semicolon
    slash
    star
    
    // One or two character tokens
    bang
    bangEqual
    equal
    equalEqual
    greater
    greateEqual
    less
    lessEqual

    // Literals
    identifier
    string
    number

    // Keywords
    and
    class
    else
    false
    fun
    for
    if
    null
    or
    print
    return
    super
    this
    true
    var
    while

    eof
)


