def reverseWords(s):
    ss = s.strip().split()
    s = ""
    flag = False
    for i in xrange(len(ss)-1, -1, -1):
        tmp = ss[i].strip()
        if flag and len(tmp) > 0:
            s += " " + tmp
        elif len(tmp) > 0:
            s += tmp
            flag = True
    return s


s = " ab    cd"
print reverseWords(s)
