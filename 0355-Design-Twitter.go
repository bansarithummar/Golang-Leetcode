type Tweet struct {
    id        int
    timestamp int
}

type Twitter struct {
    userTweets  map[int][]Tweet
    userFollows map[int]map[int]struct{}
    timestamp   int    
}

func Constructor() Twitter {
    return Twitter{
        userTweets:  make(map[int][]Tweet),
        userFollows: make(map[int]map[int]struct{}),
        timestamp:   0,
    }   
}

func (this *Twitter) PostTweet(userId int, tweetId int)  {
    this.timestamp++
    tweet := Tweet{id: tweetId, timestamp: this.timestamp}
    this.userTweets[userId] = append(this.userTweets[userId], tweet)    
}

func (this *Twitter) GetNewsFeed(userId int) []int {
    pq := &TweetHeap{}
    heap.Init(pq)
    if tweets, exists := this.userTweets[userId]; exists {
        for _, tweet := range tweets {
            heap.Push(pq, tweet)
            if pq.Len() > 10 {
                heap.Pop(pq)
            }
        }
    }
    if follows, exists := this.userFollows[userId]; exists {
        for followeeId := range follows {
            if tweets, exists := this.userTweets[followeeId]; exists {
                for _, tweet := range tweets {
                    heap.Push(pq, tweet)
                    if pq.Len() > 10 {
                        heap.Pop(pq)
                    }
                }
            }
        }
    }
    result := make([]int, pq.Len())
    for i := len(result) - 1; i >= 0; i-- {
        result[i] = heap.Pop(pq).(Tweet).id
    }
    return result    
}


func (this *Twitter) Follow(followerId int, followeeId int)  {
    if this.userFollows[followerId] == nil {
        this.userFollows[followerId] = make(map[int]struct{})
    }
    this.userFollows[followerId][followeeId] = struct{}{}    
}

func (this *Twitter) Unfollow(followerId int, followeeId int)  {
    if follows, exists := this.userFollows[followerId]; exists {
        delete(follows, followeeId)
    }    
}

type TweetHeap []Tweet

func (h TweetHeap) Len() int           { return len(h) }
func (h TweetHeap) Less(i, j int) bool { return h[i].timestamp < h[j].timestamp }
func (h TweetHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *TweetHeap) Push(x interface{}) {
    *h = append(*h, x.(Tweet))
}

func (h *TweetHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
