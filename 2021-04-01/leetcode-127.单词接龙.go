package _021_04_01

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2021/4/1 6:17 下午
 * @description：
 * @modified By：
 * @version    ：$
 */

// 参考题解： https://leetcode-cn.com/problems/word-ladder/solution/127-dan-ci-jie-long-wei-shi-yao-yao-yong-yan-sou-x/

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordSet := map[string]bool{}
	for _, v := range wordList {
		wordSet[v] = true
	}
	if _, ok := wordSet[endWord]; !ok {
		return 0
	}
	queue := make([]string, 1)
	queue[0] = beginWord
	visit := map[string]int{}
	visit[beginWord] = 1
	for len(queue) > 0 {
		// 队列第一个单词出队
		word := queue[0]
		queue = queue[1:]
		// 获取这个单词的路径长度
		path := visit[word]
		for i := 0; i < len(word); i++ {
			newWord := word
			for j := 0; j < 26; j++ {
				if i < len(word)-1 {
					newWord = newWord[:i] + string(j+'a') + newWord[i+1:]
				} else {
					newWord = newWord[:i] + string(j+'a')
				}
				// 因为开头我们判断了endWord如果不是目标直接返回0了，这里endWord一定存在
				if newWord == endWord {
					return path + 1
				}
				// wordSet中可以查询到，并且没有被访问过则加入到队列中
				if _, ok := wordSet[newWord]; ok && visit[newWord] == 0 {
					visit[newWord] = path + 1
					queue = append(queue, newWord)
				}
			}
		}
	}
	return 0
}
