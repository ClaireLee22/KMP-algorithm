def kmpAlgorithm(string, pattern):
    LPSArr = buildLPSArr(pattern)
    return stringMatching(string, pattern, LPSArr)

def buildLPSArr(pattern):
    LPSArr = [-1]*len(pattern)
    print("Init LPS array", LPSArr)

    i, j = 1, 0
    while i < len(pattern):
        print("j="+ str(j), "i="+ str(i), "LPS array", LPSArr)
        if pattern[i] == pattern[j]:
            LPSArr[i] = j
            i += 1
            j += 1
        elif j > 0:
            j = LPSArr[j-1] + 1
        else:
            i += 1
    return LPSArr

def stringMatching(string, pattern, LPSArr):
    i, j = 0, 0
    while i + len(pattern) - j <= len(string):
        if string[i] == pattern[j]:
            if j == len(pattern) - 1:
                return True
            i += 1
            j += 1
        elif j > 0:
            j = LPSArr[j-1] + 1
        else:
            i += 1
    return False

if __name__ == "__main__":
    string = "kmppkmpkmpkmdkmpkmpk"
    pattern = "kmpkmdkmpkmpk"
    print("does match:", kmpAlgorithm(string, pattern))

"""
output:
('Init LPS array', [-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1])
('j=0', 'i=1', 'LPS array', [-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1])
('j=0', 'i=2', 'LPS array', [-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1])
('j=0', 'i=3', 'LPS array', [-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1])
('j=1', 'i=4', 'LPS array', [-1, -1, -1, 0, -1, -1, -1, -1, -1, -1, -1, -1, -1])
('j=2', 'i=5', 'LPS array', [-1, -1, -1, 0, 1, -1, -1, -1, -1, -1, -1, -1, -1])
('j=0', 'i=5', 'LPS array', [-1, -1, -1, 0, 1, -1, -1, -1, -1, -1, -1, -1, -1])
('j=0', 'i=6', 'LPS array', [-1, -1, -1, 0, 1, -1, -1, -1, -1, -1, -1, -1, -1])
('j=1', 'i=7', 'LPS array', [-1, -1, -1, 0, 1, -1, 0, -1, -1, -1, -1, -1, -1])
('j=2', 'i=8', 'LPS array', [-1, -1, -1, 0, 1, -1, 0, 1, -1, -1, -1, -1, -1])
('j=3', 'i=9', 'LPS array', [-1, -1, -1, 0, 1, -1, 0, 1, 2, -1, -1, -1, -1])
('j=4', 'i=10', 'LPS array', [-1, -1, -1, 0, 1, -1, 0, 1, 2, 3, -1, -1, -1])
('j=5', 'i=11', 'LPS array', [-1, -1, -1, 0, 1, -1, 0, 1, 2, 3, 4, -1, -1])
('j=2', 'i=11', 'LPS array', [-1, -1, -1, 0, 1, -1, 0, 1, 2, 3, 4, -1, -1])
('j=3', 'i=12', 'LPS array', [-1, -1, -1, 0, 1, -1, 0, 1, 2, 3, 4, 2, -1])
('does match:', True)
"""