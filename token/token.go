package token

// TokenType이라는 사용자 정의 타입을 선언
type TokenType string

// Token 구조체 정의
type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// 식별자 + 리터럴
	IDENT = "IDENT"
	INT   = "INT"

	// 연산자
	ASSIGN = "ASSIGN"
	PLUS   = "PLUS"

	// 구분자
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// 예약어
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
