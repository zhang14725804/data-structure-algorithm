var calculate = function(s) {
    var stack = []
    var n = ''
    var sign = '+'
    var a = typeof s === 'string' ? Array.from(s) : s
    while(a.length || n) {
        var ch = a.shift()
        if (ch === ' ') continue
        if (ch === '(')  {
            n = calculate(a)
        } else if (/\D/.test(ch)) {
            switch (sign) {
                case '+':
                    stack.push(n)
                break;
                case '-':
                    stack.push(-n)
                break;
                case '*':
                    stack.push(stack.pop() * n)
                break;
                case '/':
                    stack.push(stack.pop() / n | 0)
            }
            if (ch === ')') break
            sign = ch
            n = ''
        } else {
            n += ch
        }
    }
    return stack.reduce((ch, v) => ch + (v | 0), 0)
};