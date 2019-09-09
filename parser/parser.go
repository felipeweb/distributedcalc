package parser

import (
	"context"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strconv"
	"strings"
	"text/scanner"
	"unicode"

	"github.com/felipeweb/distributedcalc/proto"
	"google.golang.org/grpc"
)

type Input struct {
	Expression string             `json:"expression,omitempty"`
	Variables  map[string]float64 `json:"variables,omitempty"`
}

func replaceVars(_ context.Context, input Input) (str string, err error) {
	runes := make([]rune, 0)
	var s scanner.Scanner
	s.Init(strings.NewReader(input.Expression))
	r := s.Next()
	for r != scanner.EOF {
		next := s.Peek()
		if unicode.IsNumber(r) && (unicode.IsLetter(next) || next == '(') {
			runes = append(runes, r, '*')
			r = s.Next()
			continue
		}
		if r == ')' && (unicode.IsLetter(next) || unicode.IsNumber(next)) {
			runes = append(runes, r, '*')
			r = s.Next()
			continue
		}
		if unicode.IsLetter(r) {
			v, ok := input.Variables[string(r)]
			if !ok {
				return "", fmt.Errorf("undeclared variable %s on position %v", string(r), s.Pos().Offset)
			}
			runes = append(runes, []rune(fmt.Sprint(v))...)
			if unicode.IsNumber(next) || next == '(' {
				runes = append(runes, '*')
			}
			r = s.Next()
			continue
		}
		runes = append(runes, r)
		r = s.Next()
	}
	return string(runes), nil
}

func Eval(ctx context.Context, input Input) (float64, error) {
	str, err := replaceVars(ctx, input)
	if err != nil {
		return 0, err
	}
	exp, err := parser.ParseExpr(str)
	if err != nil {
		return 0, err
	}
	return eval(ctx, exp)
}

func eval(ctx context.Context, exp ast.Expr) (float64, error) {
	switch exp := exp.(type) {
	case *ast.BinaryExpr:
		return evalBinaryExpr(ctx, exp)
	case *ast.BasicLit:
		switch exp.Kind {
		case token.INT:
			return strconv.ParseFloat(exp.Value, 64)
		case token.FLOAT:
			return strconv.ParseFloat(exp.Value, 64)
		}
	case *ast.ParenExpr:
		return eval(ctx, exp.X)
	default:
		return 0, fmt.Errorf("unsupported expression in position %v", exp.Pos())
	}
	return 0, nil
}

func evalBinaryExpr(ctx context.Context, exp *ast.BinaryExpr) (float64, error) {
	left, err := eval(ctx, exp.X)
	if err != nil {
		return 0, err
	}
	right, err := eval(ctx, exp.Y)
	if err != nil {
		return 0, err
	}
	switch exp.Op {
	case token.ADD:
		conn, err := grpc.Dial(os.Getenv("ADD_SERVICE_ADDR"), grpc.WithInsecure())
		if err != nil {
			return 0, err
		}
		defer conn.Close()
		c := proto.NewAddClient(conn)
		resp, err := c.Add(ctx, &proto.OpRequest{
			Left:  left,
			Right: right,
		})
		if err != nil {
			return 0, err
		}
		return resp.GetResult(), nil
	case token.SUB:
		conn, err := grpc.Dial(os.Getenv("SUB_SERVICE_ADDR"), grpc.WithInsecure())
		if err != nil {
			return 0, err
		}
		defer conn.Close()
		c := proto.NewSubClient(conn)
		resp, err := c.Sub(ctx, &proto.OpRequest{
			Left:  left,
			Right: right,
		})
		if err != nil {
			return 0, err
		}
		return resp.GetResult(), nil
	case token.MUL:
		conn, err := grpc.Dial(os.Getenv("MUL_SERVICE_ADDR"), grpc.WithInsecure())
		if err != nil {
			return 0, err
		}
		defer conn.Close()
		c := proto.NewMulClient(conn)
		resp, err := c.Mul(ctx, &proto.OpRequest{
			Left:  left,
			Right: right,
		})
		if err != nil {
			return 0, err
		}
		return resp.GetResult(), nil
	case token.QUO:
		conn, err := grpc.Dial(os.Getenv("QUO_SERVICE_ADDR"), grpc.WithInsecure())
		if err != nil {
			return 0, err
		}
		defer conn.Close()
		c := proto.NewQuoClient(conn)
		resp, err := c.Quo(ctx, &proto.OpRequest{
			Left:  left,
			Right: right,
		})
		if err != nil {
			return 0, err
		}
		return resp.GetResult(), nil
	default:
		return 0, fmt.Errorf("unsupported operator %v in position %v", exp.Op, exp.OpPos)
	}
}
