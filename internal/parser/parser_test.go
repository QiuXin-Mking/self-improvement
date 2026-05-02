package parser

import (
	"testing"
)

func TestParseContent_Q_A_Markers(t *testing.T) {
	qp := &QuestionParser{}

	content := `# q
什么是闭包？
# a
闭包是指有权访问另一个函数作用域中变量的函数。

# q
什么是原型链？
# a
原型链是 JavaScript 中实现继承的机制。`

	questions := qp.ParseContent(content, "test.md")

	if len(questions) != 2 {
		t.Fatalf("expected 2 questions, got %d", len(questions))
	}
	if questions[0].QuestionText != "什么是闭包？" {
		t.Errorf("q1 mismatch: %s", questions[0].QuestionText)
	}
	if questions[0].AnswerText != "闭包是指有权访问另一个函数作用域中变量的函数。" {
		t.Errorf("a1 mismatch: %s", questions[0].AnswerText)
	}
	if questions[1].QuestionText != "什么是原型链？" {
		t.Errorf("q2 mismatch: %s", questions[1].QuestionText)
	}
}

func TestParseContent_Question_Answer_Markers(t *testing.T) {
	qp := &QuestionParser{}

	content := `# question
什么是装饰器？
# answer
装饰器本质上是一个接受函数作为参数并返回函数的高阶函数。

# question
Python 中的 GIL 是什么？
# answer
GIL 是 Python 解释器中的互斥锁，确保同一时刻只有一个线程执行 Python 字节码。`

	questions := qp.ParseContent(content, "test.md")

	if len(questions) != 2 {
		t.Fatalf("expected 2 questions, got %d", len(questions))
	}
	if questions[0].QuestionText != "什么是装饰器？" {
		t.Errorf("q1 mismatch: %s", questions[0].QuestionText)
	}
	if questions[1].QuestionText != "Python 中的 GIL 是什么？" {
		t.Errorf("q2 mismatch: %s", questions[1].QuestionText)
	}
}

func TestParseContent_Mixed_Markers(t *testing.T) {
	qp := &QuestionParser{}

	content := `# q
Q1 内容
# answer
A1 内容

# question
Q2 内容
# a
A2 内容`

	questions := qp.ParseContent(content, "test.md")

	if len(questions) != 2 {
		t.Fatalf("expected 2 questions, got %d", len(questions))
	}
	if questions[0].QuestionText != "Q1 内容" {
		t.Errorf("q1 mismatch: %s", questions[0].QuestionText)
	}
	if questions[0].AnswerText != "A1 内容" {
		t.Errorf("a1 mismatch: %s", questions[0].AnswerText)
	}
	if questions[1].QuestionText != "Q2 内容" {
		t.Errorf("q2 mismatch: %s", questions[1].QuestionText)
	}
	if questions[1].AnswerText != "A2 内容" {
		t.Errorf("a2 mismatch: %s", questions[1].AnswerText)
	}
}

func TestParseContent_WithSpacesInMarker(t *testing.T) {
	qp := &QuestionParser{}

	content := `# question 这是一道基础题
什么是闭包？
# answer 核心概念
闭包是指有权访问另一个函数作用域中变量的函数。`

	questions := qp.ParseContent(content, "test.md")

	if len(questions) != 1 {
		t.Fatalf("expected 1 question, got %d", len(questions))
	}
	if questions[0].QuestionText != "什么是闭包？" {
		t.Errorf("q mismatch: %s", questions[0].QuestionText)
	}
}

func TestParseContent_EmptyContent(t *testing.T) {
	qp := &QuestionParser{}

	questions := qp.ParseContent("", "test.md")
	if len(questions) != 0 {
		t.Errorf("expected 0 questions, got %d", len(questions))
	}

	questions = qp.ParseContent("# q\n# a\n", "test.md")
	if len(questions) != 0 {
		t.Errorf("expected 0 questions for empty q/a, got %d", len(questions))
	}
}

func TestParseContent_NoMarkers(t *testing.T) {
	qp := &QuestionParser{}

	questions := qp.ParseContent("这是一段没有任何标记的文字。", "test.md")
	if len(questions) != 0 {
		t.Errorf("expected 0 questions, got %d", len(questions))
	}
}
