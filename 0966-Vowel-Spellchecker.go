func spellchecker(wordlist []string, queries []string) []string {
    vowels := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}

    lower := func(s string) string {
        b := []byte(s)
        for i := 0; i < len(b); i++ {
            if b[i] >= 'A' && b[i] <= 'Z' {
                b[i] = b[i] - 'A' + 'a'
            }
        }
        return string(b)
    }

    devowel := func(s string) string {
        b := []byte(lower(s))
        for i := 0; i < len(b); i++ {
            if vowels[b[i]] {
                b[i] = '*'
            }
        }
        return string(b)
    }

    exact := make(map[string]string)
    capMap := make(map[string]string)
    vowMap := make(map[string]string)

    for _, w := range wordlist {
        exact[w] = w
        lw := lower(w)
        if _, ok := capMap[lw]; !ok {
            capMap[lw] = w
        }
        dv := devowel(w)
        if _, ok := vowMap[dv]; !ok {
            vowMap[dv] = w
        }
    }

    ans := make([]string, len(queries))
    for i, q := range queries {
        if _, ok := exact[q]; ok {
            ans[i] = q
            continue
        }
        lq := lower(q)
        if w, ok := capMap[lq]; ok {
            ans[i] = w
            continue
        }
        dq := devowel(q)
        if w, ok := vowMap[dq]; ok {
            ans[i] = w
            continue
        }
        ans[i] = ""
    }

    return ans
}
