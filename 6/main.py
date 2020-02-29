def convert(s: str, numRows: int) -> str:
    if numRows == 1:
        return s
    d = {}
    pos = 0
    isDown = True
    res = ""
    for i, v in enumerate(s):
        if i < numRows:
            d[pos] = v
        else:
            d[pos] = d[pos] + v
        if pos == 0:
            isDown = True
        elif pos == numRows-1:
            isDown = False
        pos += 1 if isDown else -1
    for _, v in d.items():
        res += v
    return res


convert("PAYPALISHIRING", 1)
