def lengthOfLongestSubstring(s):
    d, maxlen, start, = {}, 0, 0
    for i, v in enumerate(s):
        if v in d:
            maxlen = max(maxlen, i - start)
            # 文字が連続した時の対処 (abba)
            start = max(start, d[v]+1)
        d[v] = i
    return max(maxlen, len(s)-start)


print(lengthOfLongestSubstring('abbba'))
