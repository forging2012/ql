// CAUTION: Generated file - DO NOT EDIT.

// Copyright (c) 2014 The ql Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ql

import (
	"fmt"
	"strconv"
)

// Implements yyLexer.
func (l *lexer) Lex(lval *yySymType) (r int) {
	const (
		INITIAL = iota
		S1
		S2
	)
	if n := l.inj; n != 0 {
		l.inj = 0
		return n
	}

	defer func() {
		pos := l.file.Position(l.First.Pos())
		l.line = pos.Line
		l.col = pos.Column
		lval.line, lval.col = l.line, l.col
	}()
	c := l.Enter()

yystate0:
	yyrule := -1
	_ = yyrule
	c = l.Rule0()

	switch yyt := l.sc; yyt {
	default:
		panic(fmt.Errorf(`invalid start condition %d`, yyt))
	case 0: // start condition: INITIAL
		goto yystart1
	case 1: // start condition: S1
		goto yystart314
	case 2: // start condition: S2
		goto yystart320
	}

	goto yystate0 // silence unused label error
	goto yyAction // silence unused label error
yyAction:
	switch yyrule {
	case 1:
		goto yyrule1
	case 2:
		goto yyrule2
	case 3:
		goto yyrule3
	case 4:
		goto yyrule4
	case 5:
		goto yyrule5
	case 6:
		goto yyrule6
	case 7:
		goto yyrule7
	case 8:
		goto yyrule8
	case 9:
		goto yyrule9
	case 10:
		goto yyrule10
	case 11:
		goto yyrule11
	case 12:
		goto yyrule12
	case 13:
		goto yyrule13
	case 14:
		goto yyrule14
	case 15:
		goto yyrule15
	case 16:
		goto yyrule16
	case 17:
		goto yyrule17
	case 18:
		goto yyrule18
	case 19:
		goto yyrule19
	case 20:
		goto yyrule20
	case 21:
		goto yyrule21
	case 22:
		goto yyrule22
	case 23:
		goto yyrule23
	case 24:
		goto yyrule24
	case 25:
		goto yyrule25
	case 26:
		goto yyrule26
	case 27:
		goto yyrule27
	case 28:
		goto yyrule28
	case 29:
		goto yyrule29
	case 30:
		goto yyrule30
	case 31:
		goto yyrule31
	case 32:
		goto yyrule32
	case 33:
		goto yyrule33
	case 34:
		goto yyrule34
	case 35:
		goto yyrule35
	case 36:
		goto yyrule36
	case 37:
		goto yyrule37
	case 38:
		goto yyrule38
	case 39:
		goto yyrule39
	case 40:
		goto yyrule40
	case 41:
		goto yyrule41
	case 42:
		goto yyrule42
	case 43:
		goto yyrule43
	case 44:
		goto yyrule44
	case 45:
		goto yyrule45
	case 46:
		goto yyrule46
	case 47:
		goto yyrule47
	case 48:
		goto yyrule48
	case 49:
		goto yyrule49
	case 50:
		goto yyrule50
	case 51:
		goto yyrule51
	case 52:
		goto yyrule52
	case 53:
		goto yyrule53
	case 54:
		goto yyrule54
	case 55:
		goto yyrule55
	case 56:
		goto yyrule56
	case 57:
		goto yyrule57
	case 58:
		goto yyrule58
	case 59:
		goto yyrule59
	case 60:
		goto yyrule60
	case 61:
		goto yyrule61
	case 62:
		goto yyrule62
	case 63:
		goto yyrule63
	case 64:
		goto yyrule64
	case 65:
		goto yyrule65
	case 66:
		goto yyrule66
	case 67:
		goto yyrule67
	case 68:
		goto yyrule68
	case 69:
		goto yyrule69
	case 70:
		goto yyrule70
	case 71:
		goto yyrule71
	case 72:
		goto yyrule72
	case 73:
		goto yyrule73
	case 74:
		goto yyrule74
	case 75:
		goto yyrule75
	case 76:
		goto yyrule76
	case 77:
		goto yyrule77
	case 78:
		goto yyrule78
	case 79:
		goto yyrule79
	case 80:
		goto yyrule80
	case 81:
		goto yyrule81
	case 82:
		goto yyrule82
	case 83:
		goto yyrule83
	case 84:
		goto yyrule84
	case 85:
		goto yyrule85
	case 86:
		goto yyrule86
	case 87:
		goto yyrule87
	case 88:
		goto yyrule88
	case 89:
		goto yyrule89
	case 90:
		goto yyrule90
	case 91:
		goto yyrule91
	case 92:
		goto yyrule92
	case 93:
		goto yyrule93
	case 94:
		goto yyrule94
	case 95:
		goto yyrule95
	case 96:
		goto yyrule96
	case 97:
		goto yyrule97
	case 98:
		goto yyrule98
	case 99:
		goto yyrule99
	case 100:
		goto yyrule100
	}
	goto yystate1 // silence unused label error
yystate1:
	c = l.Next()
yystart1:
	switch {
	default:
		goto yyabort
	case c == '!':
		goto yystate3
	case c == '"':
		goto yystate5
	case c == '$' || c == '?':
		goto yystate6
	case c == '&':
		goto yystate8
	case c == '-':
		goto yystate15
	case c == '.':
		goto yystate17
	case c == '/':
		goto yystate23
	case c == '0':
		goto yystate28
	case c == '<':
		goto yystate35
	case c == '=':
		goto yystate38
	case c == '>':
		goto yystate40
	case c == 'A' || c == 'a':
		goto yystate43
	case c == 'B' || c == 'b':
		goto yystate55
	case c == 'C' || c == 'c':
		goto yystate82
	case c == 'D' || c == 'd':
		goto yystate106
	case c == 'E' || c == 'e':
		goto yystate136
	case c == 'F' || c == 'f':
		goto yystate147
	case c == 'G' || c == 'g':
		goto yystate166
	case c == 'H' || c == 'K' || c == 'M' || c == 'P' || c == 'Q' || c >= 'X' && c <= 'Z' || c == '_' || c == 'h' || c == 'k' || c == 'm' || c == 'p' || c == 'q' || c >= 'x' && c <= 'z' || c == '\u0081':
		goto yystate44
	case c == 'I' || c == 'i':
		goto yystate171
	case c == 'J' || c == 'j':
		goto yystate191
	case c == 'L' || c == 'l':
		goto yystate195
	case c == 'N' || c == 'n':
		goto yystate205
	case c == 'O' || c == 'o':
		goto yystate211
	case c == 'R' || c == 'r':
		goto yystate226
	case c == 'S' || c == 's':
		goto yystate241
	case c == 'T' || c == 't':
		goto yystate253
	case c == 'U' || c == 'u':
		goto yystate278
	case c == 'V' || c == 'v':
		goto yystate299
	case c == 'W' || c == 'w':
		goto yystate305
	case c == '\'':
		goto yystate11
	case c == '\t' || c == '\n' || c == '\r' || c == ' ':
		goto yystate2
	case c == '\u0080':
		goto yystate313
	case c == '`':
		goto yystate310
	case c == '|':
		goto yystate311
	case c >= '1' && c <= '9':
		goto yystate34
	}

yystate2:
	c = l.Next()
	yyrule = 2
	l.Mark()
	switch {
	default:
		goto yyrule2
	case c == '\t' || c == '\n' || c == '\r' || c == ' ':
		goto yystate2
	}

yystate3:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '=':
		goto yystate4
	}

yystate4:
	c = l.Next()
	yyrule = 15
	l.Mark()
	goto yyrule15

yystate5:
	c = l.Next()
	yyrule = 10
	l.Mark()
	goto yyrule10

yystate6:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9':
		goto yystate7
	}

yystate7:
	c = l.Next()
	yyrule = 100
	l.Mark()
	switch {
	default:
		goto yyrule100
	case c >= '0' && c <= '9':
		goto yystate7
	}

yystate8:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '&':
		goto yystate9
	case c == '^':
		goto yystate10
	}

yystate9:
	c = l.Next()
	yyrule = 16
	l.Mark()
	goto yyrule16

yystate10:
	c = l.Next()
	yyrule = 17
	l.Mark()
	goto yyrule17

yystate11:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '\'':
		goto yystate12
	case c == '\\':
		goto yystate13
	case c >= '\x01' && c <= '&' || c >= '(' && c <= '[' || c >= ']' && c <= 'ÿ':
		goto yystate11
	}

yystate12:
	c = l.Next()
	yyrule = 12
	l.Mark()
	goto yyrule12

yystate13:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '\'':
		goto yystate14
	case c == '\\':
		goto yystate13
	case c >= '\x01' && c <= '&' || c >= '(' && c <= '[' || c >= ']' && c <= 'ÿ':
		goto yystate11
	}

yystate14:
	c = l.Next()
	yyrule = 12
	l.Mark()
	switch {
	default:
		goto yyrule12
	case c == '\'':
		goto yystate12
	case c == '\\':
		goto yystate13
	case c >= '\x01' && c <= '&' || c >= '(' && c <= '[' || c >= ']' && c <= 'ÿ':
		goto yystate11
	}

yystate15:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '-':
		goto yystate16
	}

yystate16:
	c = l.Next()
	yyrule = 3
	l.Mark()
	switch {
	default:
		goto yyrule3
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= 'ÿ':
		goto yystate16
	}

yystate17:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9':
		goto yystate18
	}

yystate18:
	c = l.Next()
	yyrule = 9
	l.Mark()
	switch {
	default:
		goto yyrule9
	case c == 'E' || c == 'e':
		goto yystate19
	case c == 'I' || c == 'i':
		goto yystate22
	case c >= '0' && c <= '9':
		goto yystate18
	}

yystate19:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '+' || c == '-':
		goto yystate20
	case c >= '0' && c <= '9':
		goto yystate21
	}

yystate20:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9':
		goto yystate21
	}

yystate21:
	c = l.Next()
	yyrule = 9
	l.Mark()
	switch {
	default:
		goto yyrule9
	case c == 'I' || c == 'i':
		goto yystate22
	case c >= '0' && c <= '9':
		goto yystate21
	}

yystate22:
	c = l.Next()
	yyrule = 7
	l.Mark()
	goto yyrule7

yystate23:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '*':
		goto yystate24
	case c == '/':
		goto yystate27
	}

yystate24:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '*':
		goto yystate25
	case c >= '\x01' && c <= ')' || c >= '+' && c <= 'ÿ':
		goto yystate24
	}

yystate25:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '*':
		goto yystate25
	case c == '/':
		goto yystate26
	case c >= '\x01' && c <= ')' || c >= '+' && c <= '.' || c >= '0' && c <= 'ÿ':
		goto yystate24
	}

yystate26:
	c = l.Next()
	yyrule = 5
	l.Mark()
	goto yyrule5

yystate27:
	c = l.Next()
	yyrule = 4
	l.Mark()
	switch {
	default:
		goto yyrule4
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= 'ÿ':
		goto yystate27
	}

yystate28:
	c = l.Next()
	yyrule = 8
	l.Mark()
	switch {
	default:
		goto yyrule8
	case c == '.':
		goto yystate18
	case c == '8' || c == '9':
		goto yystate30
	case c == 'E' || c == 'e':
		goto yystate19
	case c == 'I' || c == 'i':
		goto yystate31
	case c == 'X' || c == 'x':
		goto yystate32
	case c >= '0' && c <= '7':
		goto yystate29
	}

yystate29:
	c = l.Next()
	yyrule = 8
	l.Mark()
	switch {
	default:
		goto yyrule8
	case c == '.':
		goto yystate18
	case c == '8' || c == '9':
		goto yystate30
	case c == 'E' || c == 'e':
		goto yystate19
	case c == 'I' || c == 'i':
		goto yystate31
	case c >= '0' && c <= '7':
		goto yystate29
	}

yystate30:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '.':
		goto yystate18
	case c == 'E' || c == 'e':
		goto yystate19
	case c == 'I' || c == 'i':
		goto yystate31
	case c >= '0' && c <= '9':
		goto yystate30
	}

yystate31:
	c = l.Next()
	yyrule = 6
	l.Mark()
	goto yyrule6

yystate32:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate33
	}

yystate33:
	c = l.Next()
	yyrule = 8
	l.Mark()
	switch {
	default:
		goto yyrule8
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate33
	}

yystate34:
	c = l.Next()
	yyrule = 8
	l.Mark()
	switch {
	default:
		goto yyrule8
	case c == '.':
		goto yystate18
	case c == 'E' || c == 'e':
		goto yystate19
	case c == 'I' || c == 'i':
		goto yystate31
	case c >= '0' && c <= '9':
		goto yystate34
	}

yystate35:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '<':
		goto yystate36
	case c == '=':
		goto yystate37
	}

yystate36:
	c = l.Next()
	yyrule = 18
	l.Mark()
	goto yyrule18

yystate37:
	c = l.Next()
	yyrule = 19
	l.Mark()
	goto yyrule19

yystate38:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '=':
		goto yystate39
	}

yystate39:
	c = l.Next()
	yyrule = 20
	l.Mark()
	goto yyrule20

yystate40:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '=':
		goto yystate41
	case c == '>':
		goto yystate42
	}

yystate41:
	c = l.Next()
	yyrule = 21
	l.Mark()
	goto yyrule21

yystate42:
	c = l.Next()
	yyrule = 22
	l.Mark()
	goto yyrule22

yystate43:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'D' || c == 'd':
		goto yystate45
	case c == 'L' || c == 'l':
		goto yystate47
	case c == 'N' || c == 'n':
		goto yystate51
	case c == 'S' || c == 's':
		goto yystate53
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'C' || c >= 'E' && c <= 'K' || c == 'M' || c >= 'O' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'k' || c == 'm' || c >= 'o' && c <= 'r' || c >= 't' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate44:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate45:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'D' || c == 'd':
		goto yystate46
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'C' || c >= 'E' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate46:
	c = l.Next()
	yyrule = 24
	l.Mark()
	switch {
	default:
		goto yyrule24
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate47:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'T' || c == 't':
		goto yystate48
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate48:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'E' || c == 'e':
		goto yystate49
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate49:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'R' || c == 'r':
		goto yystate50
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate50:
	c = l.Next()
	yyrule = 25
	l.Mark()
	switch {
	default:
		goto yyrule25
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate51:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'D' || c == 'd':
		goto yystate52
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'C' || c >= 'E' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate52:
	c = l.Next()
	yyrule = 26
	l.Mark()
	switch {
	default:
		goto yyrule26
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate53:
	c = l.Next()
	yyrule = 28
	l.Mark()
	switch {
	default:
		goto yyrule28
	case c == 'C' || c == 'c':
		goto yystate54
	case c >= '0' && c <= '9' || c == 'A' || c == 'B' || c >= 'D' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate54:
	c = l.Next()
	yyrule = 27
	l.Mark()
	switch {
	default:
		goto yyrule27
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate55:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'E' || c == 'e':
		goto yystate56
	case c == 'I' || c == 'i':
		goto yystate65
	case c == 'L' || c == 'l':
		goto yystate73
	case c == 'O' || c == 'o':
		goto yystate76
	case c == 'Y' || c == 'y':
		goto yystate79
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'H' || c == 'J' || c == 'K' || c == 'M' || c == 'N' || c >= 'P' && c <= 'X' || c == 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'h' || c == 'j' || c == 'k' || c == 'm' || c == 'n' || c >= 'p' && c <= 'x' || c == 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate56:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'G' || c == 'g':
		goto yystate57
	case c == 'T' || c == 't':
		goto yystate60
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'H' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'f' || c >= 'h' && c <= 's' || c >= 'u' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate57:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'I' || c == 'i':
		goto yystate58
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate58:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'N' || c == 'n':
		goto yystate59
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate59:
	c = l.Next()
	yyrule = 29
	l.Mark()
	switch {
	default:
		goto yyrule29
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate60:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'W' || c == 'w':
		goto yystate61
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'V' || c >= 'X' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'v' || c >= 'x' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate61:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'E' || c == 'e':
		goto yystate62
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate62:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'E' || c == 'e':
		goto yystate63
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate63:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'N' || c == 'n':
		goto yystate64
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate64:
	c = l.Next()
	yyrule = 30
	l.Mark()
	switch {
	default:
		goto yyrule30
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate65:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'G' || c == 'g':
		goto yystate66
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'H' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'f' || c >= 'h' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate66:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'I' || c == 'i':
		goto yystate67
	case c == 'R' || c == 'r':
		goto yystate70
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'q' || c >= 's' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate67:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'N' || c == 'n':
		goto yystate68
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate68:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'T' || c == 't':
		goto yystate69
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate69:
	c = l.Next()
	yyrule = 75
	l.Mark()
	switch {
	default:
		goto yyrule75
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate70:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'A' || c == 'a':
		goto yystate71
	case c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate71:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'T' || c == 't':
		goto yystate72
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate72:
	c = l.Next()
	yyrule = 76
	l.Mark()
	switch {
	default:
		goto yyrule76
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate73:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'O' || c == 'o':
		goto yystate74
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate74:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'B' || c == 'b':
		goto yystate75
	case c >= '0' && c <= '9' || c == 'A' || c >= 'C' && c <= 'Z' || c == '_' || c == 'a' || c >= 'c' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate75:
	c = l.Next()
	yyrule = 77
	l.Mark()
	switch {
	default:
		goto yyrule77
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate76:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'O' || c == 'o':
		goto yystate77
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate77:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'L' || c == 'l':
		goto yystate78
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate78:
	c = l.Next()
	yyrule = 78
	l.Mark()
	switch {
	default:
		goto yyrule78
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate79:
	c = l.Next()
	yyrule = 31
	l.Mark()
	switch {
	default:
		goto yyrule31
	case c == 'T' || c == 't':
		goto yystate80
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate80:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'E' || c == 'e':
		goto yystate81
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate81:
	c = l.Next()
	yyrule = 79
	l.Mark()
	switch {
	default:
		goto yyrule79
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate82:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'O' || c == 'o':
		goto yystate83
	case c == 'R' || c == 'r':
		goto yystate101
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c == 'P' || c == 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c == 'p' || c == 'q' || c >= 's' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate83:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'L' || c == 'l':
		goto yystate84
	case c == 'M' || c == 'm':
		goto yystate88
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'N' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'n' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate84:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'U' || c == 'u':
		goto yystate85
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'T' || c >= 'V' && c <= 'Z' || c == '_' || c >= 'a' && c <= 't' || c >= 'v' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate85:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'M' || c == 'm':
		goto yystate86
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'L' || c >= 'N' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'l' || c >= 'n' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate86:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'N' || c == 'n':
		goto yystate87
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate87:
	c = l.Next()
	yyrule = 32
	l.Mark()
	switch {
	default:
		goto yyrule32
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate88:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'M' || c == 'm':
		goto yystate89
	case c == 'P' || c == 'p':
		goto yystate92
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'L' || c == 'N' || c == 'O' || c >= 'Q' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'l' || c == 'n' || c == 'o' || c >= 'q' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate89:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'I' || c == 'i':
		goto yystate90
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate90:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'T' || c == 't':
		goto yystate91
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate91:
	c = l.Next()
	yyrule = 33
	l.Mark()
	switch {
	default:
		goto yyrule33
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate92:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'L' || c == 'l':
		goto yystate93
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate93:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'E' || c == 'e':
		goto yystate94
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate94:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'X' || c == 'x':
		goto yystate95
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'W' || c == 'Y' || c == 'Z' || c == '_' || c >= 'a' && c <= 'w' || c == 'y' || c == 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate95:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == '0' || c >= '2' && c <= '5' || c >= '7' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	case c == '1':
		goto yystate96
	case c == '6':
		goto yystate99
	}

yystate96:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == '0' || c == '1' || c >= '3' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	case c == '2':
		goto yystate97
	}

yystate97:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == '8':
		goto yystate98
	case c >= '0' && c <= '7' || c == '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate98:
	c = l.Next()
	yyrule = 80
	l.Mark()
	switch {
	default:
		goto yyrule80
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate99:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == '4':
		goto yystate100
	case c >= '0' && c <= '3' || c >= '5' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate100:
	c = l.Next()
	yyrule = 81
	l.Mark()
	switch {
	default:
		goto yyrule81
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate101:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'E' || c == 'e':
		goto yystate102
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate102:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'A' || c == 'a':
		goto yystate103
	case c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate103:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'T' || c == 't':
		goto yystate104
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate104:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'E' || c == 'e':
		goto yystate105
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate105:
	c = l.Next()
	yyrule = 34
	l.Mark()
	switch {
	default:
		goto yyrule34
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate106:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'E' || c == 'e':
		goto yystate107
	case c == 'I' || c == 'i':
		goto yystate119
	case c == 'R' || c == 'r':
		goto yystate126
	case c == 'U' || c == 'u':
		goto yystate129
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'H' || c >= 'J' && c <= 'Q' || c == 'S' || c == 'T' || c >= 'V' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'h' || c >= 'j' && c <= 'q' || c == 's' || c == 't' || c >= 'v' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate107:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'F' || c == 'f':
		goto yystate108
	case c == 'L' || c == 'l':
		goto yystate113
	case c == 'S' || c == 's':
		goto yystate117
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'E' || c >= 'G' && c <= 'K' || c >= 'M' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'e' || c >= 'g' && c <= 'k' || c >= 'm' && c <= 'r' || c >= 't' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate108:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'A' || c == 'a':
		goto yystate109
	case c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate109:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'U' || c == 'u':
		goto yystate110
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'T' || c >= 'V' && c <= 'Z' || c == '_' || c >= 'a' && c <= 't' || c >= 'v' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate110:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'L' || c == 'l':
		goto yystate111
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate111:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'T' || c == 't':
		goto yystate112
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate112:
	c = l.Next()
	yyrule = 35
	l.Mark()
	switch {
	default:
		goto yyrule35
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate113:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'E' || c == 'e':
		goto yystate114
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate114:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'T' || c == 't':
		goto yystate115
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate115:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'E' || c == 'e':
		goto yystate116
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate116:
	c = l.Next()
	yyrule = 36
	l.Mark()
	switch {
	default:
		goto yyrule36
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate117:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'C' || c == 'c':
		goto yystate118
	case c >= '0' && c <= '9' || c == 'A' || c == 'B' || c >= 'D' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate118:
	c = l.Next()
	yyrule = 37
	l.Mark()
	switch {
	default:
		goto yyrule37
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate119:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'S' || c == 's':
		goto yystate120
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate120:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'T' || c == 't':
		goto yystate121
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate121:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'I' || c == 'i':
		goto yystate122
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate122:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'N' || c == 'n':
		goto yystate123
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate123:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'C' || c == 'c':
		goto yystate124
	case c >= '0' && c <= '9' || c == 'A' || c == 'B' || c >= 'D' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate124:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'T' || c == 't':
		goto yystate125
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate125:
	c = l.Next()
	yyrule = 38
	l.Mark()
	switch {
	default:
		goto yyrule38
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate126:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'O' || c == 'o':
		goto yystate127
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate127:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'P' || c == 'p':
		goto yystate128
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'O' || c >= 'Q' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'o' || c >= 'q' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate128:
	c = l.Next()
	yyrule = 39
	l.Mark()
	switch {
	default:
		goto yyrule39
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate129:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'R' || c == 'r':
		goto yystate130
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate130:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'A' || c == 'a':
		goto yystate131
	case c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate131:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'T' || c == 't':
		goto yystate132
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate132:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'I' || c == 'i':
		goto yystate133
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate133:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'O' || c == 'o':
		goto yystate134
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate134:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'N' || c == 'n':
		goto yystate135
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate135:
	c = l.Next()
	yyrule = 82
	l.Mark()
	switch {
	default:
		goto yyrule82
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate136:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'X' || c == 'x':
		goto yystate137
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'W' || c == 'Y' || c == 'Z' || c == '_' || c >= 'a' && c <= 'w' || c == 'y' || c == 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate137:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'I' || c == 'i':
		goto yystate138
	case c == 'P' || c == 'p':
		goto yystate142
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'O' || c >= 'Q' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'o' || c >= 'q' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate138:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'S' || c == 's':
		goto yystate139
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate139:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'T' || c == 't':
		goto yystate140
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate140:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'S' || c == 's':
		goto yystate141
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate141:
	c = l.Next()
	yyrule = 40
	l.Mark()
	switch {
	default:
		goto yyrule40
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate142:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'L' || c == 'l':
		goto yystate143
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate143:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'A' || c == 'a':
		goto yystate144
	case c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate144:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'I' || c == 'i':
		goto yystate145
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate145:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'N' || c == 'n':
		goto yystate146
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate146:
	c = l.Next()
	yyrule = 41
	l.Mark()
	switch {
	default:
		goto yyrule41
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate147:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'A' || c == 'a':
		goto yystate148
	case c == 'L' || c == 'l':
		goto yystate152
	case c == 'R' || c == 'r':
		goto yystate160
	case c == 'U' || c == 'u':
		goto yystate163
	case c >= '0' && c <= '9' || c >= 'B' && c <= 'K' || c >= 'M' && c <= 'Q' || c == 'S' || c == 'T' || c >= 'V' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'k' || c >= 'm' && c <= 'q' || c == 's' || c == 't' || c >= 'v' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate148:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'L' || c == 'l':
		goto yystate149
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate149:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'S' || c == 's':
		goto yystate150
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate150:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'E' || c == 'e':
		goto yystate151
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate151:
	c = l.Next()
	yyrule = 73
	l.Mark()
	switch {
	default:
		goto yyrule73
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate152:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'O' || c == 'o':
		goto yystate153
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate153:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'A' || c == 'a':
		goto yystate154
	case c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate154:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'T' || c == 't':
		goto yystate155
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate155:
	c = l.Next()
	yyrule = 83
	l.Mark()
	switch {
	default:
		goto yyrule83
	case c == '3':
		goto yystate156
	case c == '6':
		goto yystate158
	case c >= '0' && c <= '2' || c == '4' || c == '5' || c >= '7' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate156:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == '0' || c == '1' || c >= '3' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	case c == '2':
		goto yystate157
	}

yystate157:
	c = l.Next()
	yyrule = 84
	l.Mark()
	switch {
	default:
		goto yyrule84
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate158:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == '4':
		goto yystate159
	case c >= '0' && c <= '3' || c >= '5' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate159:
	c = l.Next()
	yyrule = 85
	l.Mark()
	switch {
	default:
		goto yyrule85
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate160:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'O' || c == 'o':
		goto yystate161
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate161:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'M' || c == 'm':
		goto yystate162
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'L' || c >= 'N' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'l' || c >= 'n' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate162:
	c = l.Next()
	yyrule = 42
	l.Mark()
	switch {
	default:
		goto yyrule42
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate163:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'L' || c == 'l':
		goto yystate164
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate164:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'L' || c == 'l':
		goto yystate165
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate165:
	c = l.Next()
	yyrule = 43
	l.Mark()
	switch {
	default:
		goto yyrule43
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate166:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'R' || c == 'r':
		goto yystate167
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate167:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'O' || c == 'o':
		goto yystate168
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate168:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'U' || c == 'u':
		goto yystate169
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'T' || c >= 'V' && c <= 'Z' || c == '_' || c >= 'a' && c <= 't' || c >= 'v' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate169:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'P' || c == 'p':
		goto yystate170
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'O' || c >= 'Q' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'o' || c >= 'q' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate170:
	c = l.Next()
	yyrule = 44
	l.Mark()
	switch {
	default:
		goto yyrule44
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate171:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'F' || c == 'f':
		goto yystate172
	case c == 'N' || c == 'n':
		goto yystate173
	case c == 'S' || c == 's':
		goto yystate190
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'E' || c >= 'G' && c <= 'M' || c >= 'O' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'e' || c >= 'g' && c <= 'm' || c >= 'o' && c <= 'r' || c >= 't' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate172:
	c = l.Next()
	yyrule = 45
	l.Mark()
	switch {
	default:
		goto yyrule45
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate173:
	c = l.Next()
	yyrule = 49
	l.Mark()
	switch {
	default:
		goto yyrule49
	case c == 'D' || c == 'd':
		goto yystate174
	case c == 'S' || c == 's':
		goto yystate177
	case c == 'T' || c == 't':
		goto yystate181
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'C' || c >= 'E' && c <= 'R' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'r' || c >= 'u' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate174:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'E' || c == 'e':
		goto yystate175
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate175:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'X' || c == 'x':
		goto yystate176
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'W' || c == 'Y' || c == 'Z' || c == '_' || c >= 'a' && c <= 'w' || c == 'y' || c == 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate176:
	c = l.Next()
	yyrule = 46
	l.Mark()
	switch {
	default:
		goto yyrule46
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate177:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'E' || c == 'e':
		goto yystate178
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate178:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'R' || c == 'r':
		goto yystate179
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate179:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'T' || c == 't':
		goto yystate180
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate180:
	c = l.Next()
	yyrule = 47
	l.Mark()
	switch {
	default:
		goto yyrule47
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate181:
	c = l.Next()
	yyrule = 86
	l.Mark()
	switch {
	default:
		goto yyrule86
	case c == '0' || c == '2' || c == '4' || c == '5' || c == '7' || c == '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	case c == '1':
		goto yystate182
	case c == '3':
		goto yystate184
	case c == '6':
		goto yystate186
	case c == '8':
		goto yystate188
	case c == 'O' || c == 'o':
		goto yystate189
	}

yystate182:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == '6':
		goto yystate183
	case c >= '0' && c <= '5' || c >= '7' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate183:
	c = l.Next()
	yyrule = 87
	l.Mark()
	switch {
	default:
		goto yyrule87
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate184:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == '0' || c == '1' || c >= '3' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	case c == '2':
		goto yystate185
	}

yystate185:
	c = l.Next()
	yyrule = 88
	l.Mark()
	switch {
	default:
		goto yyrule88
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate186:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == '4':
		goto yystate187
	case c >= '0' && c <= '3' || c >= '5' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate187:
	c = l.Next()
	yyrule = 89
	l.Mark()
	switch {
	default:
		goto yyrule89
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate188:
	c = l.Next()
	yyrule = 90
	l.Mark()
	switch {
	default:
		goto yyrule90
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate189:
	c = l.Next()
	yyrule = 48
	l.Mark()
	switch {
	default:
		goto yyrule48
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate190:
	c = l.Next()
	yyrule = 50
	l.Mark()
	switch {
	default:
		goto yyrule50
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate191:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'O' || c == 'o':
		goto yystate192
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate192:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'I' || c == 'i':
		goto yystate193
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate193:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'N' || c == 'n':
		goto yystate194
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate194:
	c = l.Next()
	yyrule = 51
	l.Mark()
	switch {
	default:
		goto yyrule51
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate195:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'E' || c == 'e':
		goto yystate196
	case c == 'I' || c == 'i':
		goto yystate199
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'h' || c >= 'j' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate196:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'F' || c == 'f':
		goto yystate197
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'E' || c >= 'G' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'e' || c >= 'g' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate197:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'T' || c == 't':
		goto yystate198
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate198:
	c = l.Next()
	yyrule = 52
	l.Mark()
	switch {
	default:
		goto yyrule52
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate199:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'K' || c == 'k':
		goto yystate200
	case c == 'M' || c == 'm':
		goto yystate202
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'J' || c == 'L' || c >= 'N' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'j' || c == 'l' || c >= 'n' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate200:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'E' || c == 'e':
		goto yystate201
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate201:
	c = l.Next()
	yyrule = 53
	l.Mark()
	switch {
	default:
		goto yyrule53
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate202:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'I' || c == 'i':
		goto yystate203
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate203:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'T' || c == 't':
		goto yystate204
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate204:
	c = l.Next()
	yyrule = 54
	l.Mark()
	switch {
	default:
		goto yyrule54
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate205:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'O' || c == 'o':
		goto yystate206
	case c == 'U' || c == 'u':
		goto yystate208
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'T' || c >= 'V' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 't' || c >= 'v' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate206:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'T' || c == 't':
		goto yystate207
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate207:
	c = l.Next()
	yyrule = 55
	l.Mark()
	switch {
	default:
		goto yyrule55
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate208:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'L' || c == 'l':
		goto yystate209
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate209:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'L' || c == 'l':
		goto yystate210
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate210:
	c = l.Next()
	yyrule = 72
	l.Mark()
	switch {
	default:
		goto yyrule72
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate211:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'F' || c == 'f':
		goto yystate212
	case c == 'N' || c == 'n':
		goto yystate217
	case c == 'R' || c == 'r':
		goto yystate218
	case c == 'U' || c == 'u':
		goto yystate222
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'E' || c >= 'G' && c <= 'M' || c >= 'O' && c <= 'Q' || c == 'S' || c == 'T' || c >= 'V' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'e' || c >= 'g' && c <= 'm' || c >= 'o' && c <= 'q' || c == 's' || c == 't' || c >= 'v' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate212:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'F' || c == 'f':
		goto yystate213
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'E' || c >= 'G' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'e' || c >= 'g' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate213:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'S' || c == 's':
		goto yystate214
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate214:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'E' || c == 'e':
		goto yystate215
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate215:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'T' || c == 't':
		goto yystate216
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate216:
	c = l.Next()
	yyrule = 56
	l.Mark()
	switch {
	default:
		goto yyrule56
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate217:
	c = l.Next()
	yyrule = 57
	l.Mark()
	switch {
	default:
		goto yyrule57
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate218:
	c = l.Next()
	yyrule = 59
	l.Mark()
	switch {
	default:
		goto yyrule59
	case c == 'D' || c == 'd':
		goto yystate219
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'C' || c >= 'E' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate219:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'E' || c == 'e':
		goto yystate220
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate220:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'R' || c == 'r':
		goto yystate221
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate221:
	c = l.Next()
	yyrule = 58
	l.Mark()
	switch {
	default:
		goto yyrule58
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate222:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'T' || c == 't':
		goto yystate223
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate223:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'E' || c == 'e':
		goto yystate224
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate224:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'R' || c == 'r':
		goto yystate225
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate225:
	c = l.Next()
	yyrule = 60
	l.Mark()
	switch {
	default:
		goto yyrule60
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate226:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'I' || c == 'i':
		goto yystate227
	case c == 'O' || c == 'o':
		goto yystate231
	case c == 'U' || c == 'u':
		goto yystate238
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'N' || c >= 'P' && c <= 'T' || c >= 'V' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'n' || c >= 'p' && c <= 't' || c >= 'v' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate227:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'G' || c == 'g':
		goto yystate228
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'H' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'f' || c >= 'h' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate228:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'H' || c == 'h':
		goto yystate229
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'G' || c >= 'I' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'g' || c >= 'i' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate229:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'T' || c == 't':
		goto yystate230
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate230:
	c = l.Next()
	yyrule = 61
	l.Mark()
	switch {
	default:
		goto yyrule61
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate231:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'L' || c == 'l':
		goto yystate232
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate232:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'L' || c == 'l':
		goto yystate233
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate233:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'B' || c == 'b':
		goto yystate234
	case c >= '0' && c <= '9' || c == 'A' || c >= 'C' && c <= 'Z' || c == '_' || c == 'a' || c >= 'c' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate234:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'A' || c == 'a':
		goto yystate235
	case c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate235:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'C' || c == 'c':
		goto yystate236
	case c >= '0' && c <= '9' || c == 'A' || c == 'B' || c >= 'D' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate236:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'K' || c == 'k':
		goto yystate237
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'J' || c >= 'L' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'j' || c >= 'l' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate237:
	c = l.Next()
	yyrule = 62
	l.Mark()
	switch {
	default:
		goto yyrule62
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate238:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'N' || c == 'n':
		goto yystate239
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate239:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'E' || c == 'e':
		goto yystate240
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate240:
	c = l.Next()
	yyrule = 91
	l.Mark()
	switch {
	default:
		goto yyrule91
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate241:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'E' || c == 'e':
		goto yystate242
	case c == 'T' || c == 't':
		goto yystate248
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 's' || c >= 'u' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate242:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'L' || c == 'l':
		goto yystate243
	case c == 'T' || c == 't':
		goto yystate247
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 's' || c >= 'u' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate243:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'E' || c == 'e':
		goto yystate244
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate244:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'C' || c == 'c':
		goto yystate245
	case c >= '0' && c <= '9' || c == 'A' || c == 'B' || c >= 'D' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate245:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'T' || c == 't':
		goto yystate246
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate246:
	c = l.Next()
	yyrule = 63
	l.Mark()
	switch {
	default:
		goto yyrule63
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate247:
	c = l.Next()
	yyrule = 64
	l.Mark()
	switch {
	default:
		goto yyrule64
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate248:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'R' || c == 'r':
		goto yystate249
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate249:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'I' || c == 'i':
		goto yystate250
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate250:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'N' || c == 'n':
		goto yystate251
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate251:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'G' || c == 'g':
		goto yystate252
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'H' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'f' || c >= 'h' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate252:
	c = l.Next()
	yyrule = 92
	l.Mark()
	switch {
	default:
		goto yyrule92
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate253:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'A' || c == 'a':
		goto yystate254
	case c == 'I' || c == 'i':
		goto yystate258
	case c == 'R' || c == 'r':
		goto yystate261
	case c >= '0' && c <= '9' || c >= 'B' && c <= 'H' || c >= 'J' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'h' || c >= 'j' && c <= 'q' || c >= 's' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate254:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'B' || c == 'b':
		goto yystate255
	case c >= '0' && c <= '9' || c == 'A' || c >= 'C' && c <= 'Z' || c == '_' || c == 'a' || c >= 'c' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate255:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'L' || c == 'l':
		goto yystate256
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate256:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'E' || c == 'e':
		goto yystate257
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate257:
	c = l.Next()
	yyrule = 65
	l.Mark()
	switch {
	default:
		goto yyrule65
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate258:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'M' || c == 'm':
		goto yystate259
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'L' || c >= 'N' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'l' || c >= 'n' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate259:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'E' || c == 'e':
		goto yystate260
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate260:
	c = l.Next()
	yyrule = 93
	l.Mark()
	switch {
	default:
		goto yyrule93
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate261:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'A' || c == 'a':
		goto yystate262
	case c == 'U' || c == 'u':
		goto yystate271
	case c >= '0' && c <= '9' || c >= 'B' && c <= 'T' || c >= 'V' && c <= 'Z' || c == '_' || c >= 'b' && c <= 't' || c >= 'v' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate262:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'N' || c == 'n':
		goto yystate263
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate263:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'S' || c == 's':
		goto yystate264
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate264:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'A' || c == 'a':
		goto yystate265
	case c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate265:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'C' || c == 'c':
		goto yystate266
	case c >= '0' && c <= '9' || c == 'A' || c == 'B' || c >= 'D' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate266:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'T' || c == 't':
		goto yystate267
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate267:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'I' || c == 'i':
		goto yystate268
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate268:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'O' || c == 'o':
		goto yystate269
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate269:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'N' || c == 'n':
		goto yystate270
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate270:
	c = l.Next()
	yyrule = 66
	l.Mark()
	switch {
	default:
		goto yyrule66
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate271:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'E' || c == 'e':
		goto yystate272
	case c == 'N' || c == 'n':
		goto yystate273
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'm' || c >= 'o' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate272:
	c = l.Next()
	yyrule = 74
	l.Mark()
	switch {
	default:
		goto yyrule74
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate273:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'C' || c == 'c':
		goto yystate274
	case c >= '0' && c <= '9' || c == 'A' || c == 'B' || c >= 'D' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate274:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'A' || c == 'a':
		goto yystate275
	case c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate275:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'T' || c == 't':
		goto yystate276
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate276:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'E' || c == 'e':
		goto yystate277
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate277:
	c = l.Next()
	yyrule = 67
	l.Mark()
	switch {
	default:
		goto yyrule67
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate278:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'I' || c == 'i':
		goto yystate279
	case c == 'N' || c == 'n':
		goto yystate289
	case c == 'P' || c == 'p':
		goto yystate294
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'M' || c == 'O' || c >= 'Q' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'm' || c == 'o' || c >= 'q' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate279:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'N' || c == 'n':
		goto yystate280
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate280:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'T' || c == 't':
		goto yystate281
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate281:
	c = l.Next()
	yyrule = 94
	l.Mark()
	switch {
	default:
		goto yyrule94
	case c == '0' || c == '2' || c == '4' || c == '5' || c == '7' || c == '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	case c == '1':
		goto yystate282
	case c == '3':
		goto yystate284
	case c == '6':
		goto yystate286
	case c == '8':
		goto yystate288
	}

yystate282:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == '6':
		goto yystate283
	case c >= '0' && c <= '5' || c >= '7' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate283:
	c = l.Next()
	yyrule = 95
	l.Mark()
	switch {
	default:
		goto yyrule95
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate284:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == '0' || c == '1' || c >= '3' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	case c == '2':
		goto yystate285
	}

yystate285:
	c = l.Next()
	yyrule = 96
	l.Mark()
	switch {
	default:
		goto yyrule96
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate286:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == '4':
		goto yystate287
	case c >= '0' && c <= '3' || c >= '5' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate287:
	c = l.Next()
	yyrule = 97
	l.Mark()
	switch {
	default:
		goto yyrule97
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate288:
	c = l.Next()
	yyrule = 98
	l.Mark()
	switch {
	default:
		goto yyrule98
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate289:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'I' || c == 'i':
		goto yystate290
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate290:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'Q' || c == 'q':
		goto yystate291
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'P' || c >= 'R' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'p' || c >= 'r' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate291:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'U' || c == 'u':
		goto yystate292
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'T' || c >= 'V' && c <= 'Z' || c == '_' || c >= 'a' && c <= 't' || c >= 'v' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate292:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'E' || c == 'e':
		goto yystate293
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate293:
	c = l.Next()
	yyrule = 69
	l.Mark()
	switch {
	default:
		goto yyrule69
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate294:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'D' || c == 'd':
		goto yystate295
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'C' || c >= 'E' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate295:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'A' || c == 'a':
		goto yystate296
	case c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate296:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'T' || c == 't':
		goto yystate297
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate297:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'E' || c == 'e':
		goto yystate298
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate298:
	c = l.Next()
	yyrule = 68
	l.Mark()
	switch {
	default:
		goto yyrule68
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate299:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'A' || c == 'a':
		goto yystate300
	case c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate300:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'L' || c == 'l':
		goto yystate301
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate301:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'U' || c == 'u':
		goto yystate302
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'T' || c >= 'V' && c <= 'Z' || c == '_' || c >= 'a' && c <= 't' || c >= 'v' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate302:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'E' || c == 'e':
		goto yystate303
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate303:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'S' || c == 's':
		goto yystate304
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate304:
	c = l.Next()
	yyrule = 70
	l.Mark()
	switch {
	default:
		goto yyrule70
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate305:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'H' || c == 'h':
		goto yystate306
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'G' || c >= 'I' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'g' || c >= 'i' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate306:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'E' || c == 'e':
		goto yystate307
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate307:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'R' || c == 'r':
		goto yystate308
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate308:
	c = l.Next()
	yyrule = 99
	l.Mark()
	switch {
	default:
		goto yyrule99
	case c == 'E' || c == 'e':
		goto yystate309
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate309:
	c = l.Next()
	yyrule = 71
	l.Mark()
	switch {
	default:
		goto yyrule71
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0081' || c == '\u0082':
		goto yystate44
	}

yystate310:
	c = l.Next()
	yyrule = 11
	l.Mark()
	goto yyrule11

yystate311:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '|':
		goto yystate312
	}

yystate312:
	c = l.Next()
	yyrule = 23
	l.Mark()
	goto yyrule23

yystate313:
	c = l.Next()
	yyrule = 1
	l.Mark()
	goto yyrule1

	goto yystate314 // silence unused label error
yystate314:
	c = l.Next()
yystart314:
	switch {
	default:
		goto yyabort
	case c == '"':
		goto yystate316
	case c == '\\':
		goto yystate317
	case c == '\u0080':
		goto yystate319
	case c >= '\x01' && c <= '!' || c >= '#' && c <= '[' || c >= ']' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate315
	}

yystate315:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '"':
		goto yystate316
	case c == '\\':
		goto yystate317
	case c >= '\x01' && c <= '!' || c >= '#' && c <= '[' || c >= ']' && c <= 'ÿ':
		goto yystate315
	}

yystate316:
	c = l.Next()
	yyrule = 13
	l.Mark()
	goto yyrule13

yystate317:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '"':
		goto yystate318
	case c == '\\':
		goto yystate317
	case c >= '\x01' && c <= '!' || c >= '#' && c <= '[' || c >= ']' && c <= 'ÿ':
		goto yystate315
	}

yystate318:
	c = l.Next()
	yyrule = 13
	l.Mark()
	switch {
	default:
		goto yyrule13
	case c == '"':
		goto yystate316
	case c == '\\':
		goto yystate317
	case c >= '\x01' && c <= '!' || c >= '#' && c <= '[' || c >= ']' && c <= 'ÿ':
		goto yystate315
	}

yystate319:
	c = l.Next()
	yyrule = 1
	l.Mark()
	switch {
	default:
		goto yyrule1
	case c == '"':
		goto yystate316
	case c == '\\':
		goto yystate317
	case c >= '\x01' && c <= '!' || c >= '#' && c <= '[' || c >= ']' && c <= 'ÿ':
		goto yystate315
	}

	goto yystate320 // silence unused label error
yystate320:
	c = l.Next()
yystart320:
	switch {
	default:
		goto yyabort
	case c == '\u0080':
		goto yystate323
	case c == '`':
		goto yystate322
	case c >= '\x01' && c <= '_' || c >= 'a' && c <= '\u007f' || c >= '\u0081' && c <= 'ÿ':
		goto yystate321
	}

yystate321:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '`':
		goto yystate322
	case c >= '\x01' && c <= '_' || c >= 'a' && c <= 'ÿ':
		goto yystate321
	}

yystate322:
	c = l.Next()
	yyrule = 14
	l.Mark()
	goto yyrule14

yystate323:
	c = l.Next()
	yyrule = 1
	l.Mark()
	switch {
	default:
		goto yyrule1
	case c == '`':
		goto yystate322
	case c >= '\x01' && c <= '_' || c >= 'a' && c <= 'ÿ':
		goto yystate321
	}

yyrule1: // \x80
	{
		return 0
	}
yyrule2: // [ \t\n\r]+

	goto yystate0
yyrule3: // --.*

	goto yystate0
yyrule4: // \/\/.*

	goto yystate0
yyrule5: // \/\*([^*]|\*+[^*/])*\*+\/

	goto yystate0
yyrule6: // {imaginary_ilit}
	{
		return l.int(lval, true)
	}
yyrule7: // {imaginary_lit}
	{
		return l.float(lval, true)
	}
yyrule8: // {int_lit}
	{
		return l.int(lval, false)
	}
yyrule9: // {float_lit}
	{
		return l.float(lval, false)
	}
yyrule10: // \"
	{
		l.sc = S1
		goto yystate0
	}
yyrule11: // `
	{
		l.sc = S2
		goto yystate0
	}
yyrule12: // '(\\.|[^'])*'
	{
		if ret := l.str(lval, ""); ret != stringLit {
			return ret
		}
		lval.item = idealRune(lval.item.(string)[0])
		return intLit
	}
yyrule13: // (\\.|[^\"])*\"
	{
		return l.str(lval, "\"")
	}
yyrule14: // ([^`]|\n)*`
	{
		return l.str(lval, "`")
	}
yyrule15: // "!="
	{
		return neq
	}
yyrule16: // "&&"
	{
		return andand
	}
yyrule17: // "&^"
	{
		return andnot
	}
yyrule18: // "<<"
	{
		return lsh
	}
yyrule19: // "<="
	{
		return le
	}
yyrule20: // "=="
	{
		return eq
	}
yyrule21: // ">="
	{
		return ge
	}
yyrule22: // ">>"
	{
		return rsh
	}
yyrule23: // "||"
	{
		return oror
	}
yyrule24: // {add}
	{
		return add
	}
yyrule25: // {alter}
	{
		return alter
	}
yyrule26: // {and}
	{
		return and
	}
yyrule27: // {asc}
	{
		return asc
	}
yyrule28: // {as}
	{
		return as
	}
yyrule29: // {begin}
	{
		return begin
	}
yyrule30: // {between}
	{
		return between
	}
yyrule31: // {by}
	{
		return by
	}
yyrule32: // {column}
	{
		return column
	}
yyrule33: // {commit}
	{
		return commit
	}
yyrule34: // {create}
	{
		return create
	}
yyrule35: // {default}
	{
		return defaultKwd
	}
yyrule36: // {delete}
	{
		return deleteKwd
	}
yyrule37: // {desc}
	{
		return desc
	}
yyrule38: // {distinct}
	{
		return distinct
	}
yyrule39: // {drop}
	{
		return drop
	}
yyrule40: // {exists}
	{
		return exists
	}
yyrule41: // {explain}
	{
		return explain
	}
yyrule42: // {from}
	{
		return from
	}
yyrule43: // {full}
	{
		return full
	}
yyrule44: // {group}
	{
		return group
	}
yyrule45: // {if}
	{
		return ifKwd
	}
yyrule46: // {index}
	{
		return index
	}
yyrule47: // {insert}
	{
		return insert
	}
yyrule48: // {into}
	{
		return into
	}
yyrule49: // {in}
	{
		return in
	}
yyrule50: // {is}
	{
		return is
	}
yyrule51: // {join}
	{
		return join
	}
yyrule52: // {left}
	{
		return left
	}
yyrule53: // {like}
	{
		return like
	}
yyrule54: // {limit}
	{
		return limit
	}
yyrule55: // {not}
	{
		return not
	}
yyrule56: // {offset}
	{
		return offset
	}
yyrule57: // {on}
	{
		return on
	}
yyrule58: // {order}
	{
		return order
	}
yyrule59: // {or}
	{
		return or
	}
yyrule60: // {outer}
	{
		return outer
	}
yyrule61: // {right}
	{
		return right
	}
yyrule62: // {rollback}
	{
		return rollback
	}
yyrule63: // {select}
	{
		l.agg = append(l.agg, false)
		return selectKwd
	}
yyrule64: // {set}
	{
		return set
	}
yyrule65: // {table}
	{
		return tableKwd
	}
yyrule66: // {transaction}
	{
		return transaction
	}
yyrule67: // {truncate}
	{
		return truncate
	}
yyrule68: // {update}
	{
		return update
	}
yyrule69: // {unique}
	{
		return unique
	}
yyrule70: // {values}
	{
		return values
	}
yyrule71: // {where}
	{
		return where
	}
yyrule72: // {null}
	{
		lval.item = nil
		return null
	}
yyrule73: // {false}
	{
		lval.item = false
		return falseKwd
	}
yyrule74: // {true}
	{
		lval.item = true
		return trueKwd
	}
yyrule75: // {bigint}
	{
		lval.item = qBigInt
		return bigIntType
	}
yyrule76: // {bigrat}
	{
		lval.item = qBigRat
		return bigRatType
	}
yyrule77: // {blob}
	{
		lval.item = qBlob
		return blobType
	}
yyrule78: // {bool}
	{
		lval.item = qBool
		return boolType
	}
yyrule79: // {byte}
	{
		lval.item = qUint8
		return byteType
	}
yyrule80: // {complex}128
	{
		lval.item = qComplex128
		return complex128Type
	}
yyrule81: // {complex}64
	{
		lval.item = qComplex64
		return complex64Type
	}
yyrule82: // {duration}
	{
		lval.item = qDuration
		return durationType
	}
yyrule83: // {float}
	{
		lval.item = qFloat64
		return floatType
	}
yyrule84: // {float}32
	{
		lval.item = qFloat32
		return float32Type
	}
yyrule85: // {float}64
	{
		lval.item = qFloat64
		return float64Type
	}
yyrule86: // {int}
	{
		lval.item = qInt64
		return intType
	}
yyrule87: // {int}16
	{
		lval.item = qInt16
		return int16Type
	}
yyrule88: // {int}32
	{
		lval.item = qInt32
		return int32Type
	}
yyrule89: // {int}64
	{
		lval.item = qInt64
		return int64Type
	}
yyrule90: // {int}8
	{
		lval.item = qInt8
		return int8Type
	}
yyrule91: // {rune}
	{
		lval.item = qInt32
		return runeType
	}
yyrule92: // {string}
	{
		lval.item = qString
		return stringType
	}
yyrule93: // {time}
	{
		lval.item = qTime
		return timeType
	}
yyrule94: // {uint}
	{
		lval.item = qUint64
		return uintType
	}
yyrule95: // {uint}16
	{
		lval.item = qUint16
		return uint16Type
	}
yyrule96: // {uint}32
	{
		lval.item = qUint32
		return uint32Type
	}
yyrule97: // {uint}64
	{
		lval.item = qUint64
		return uint64Type
	}
yyrule98: // {uint}8
	{
		lval.item = qUint8
		return uint8Type
	}
yyrule99: // {ident}
	{
		lval.item = string(l.TokenBytes(nil))
		return identifier
	}
yyrule100: // ($|\?){D}
	{
		lval.item, _ = strconv.Atoi(string(l.TokenBytes(nil)[1:]))
		return qlParam
	}
	panic("unreachable")

	goto yyabort // silence unused label error

yyabort: // no lexem recognized
	if c, ok := l.Abort(); ok {
		return c
	}

	goto yyAction
}
