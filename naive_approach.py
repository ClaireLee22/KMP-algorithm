def stringMatching(string, pattern):
    for i in range(len(string)):
        j = 0
        while j < len(pattern):
            if string[i+j] != pattern[j]:
                break
            j += 1
        if j == len(pattern):
            return True
    return False


if __name__ == "__main__":
    string = "kmppkmpkmpkmdkmpkmpk"
    pattern = "kmpkmdkmpkmpk"
    print("does match:", stringMatching(string, pattern))

"""
output:
('does match:', True)
"""