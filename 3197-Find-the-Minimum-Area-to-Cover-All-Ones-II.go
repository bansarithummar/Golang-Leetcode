func min(a, b int) int { if a < b { return a }; return b }
func max(a, b int) int { if a > b { return a }; return b }
func minimumSum(grid [][]int) int {
    const INF = int(1 << 29)
	H, W := len(grid), len(grid[0])

	nextRight := make([][]int, H)
	prevLeft := make([][]int, H)
	for r := 0; r < H; r++ {
		nextRight[r] = make([]int, W)
		prevLeft[r] = make([]int, W)
		nxt := W
		for c := W - 1; c >= 0; c-- {
			if grid[r][c] == 1 { nxt = c }
			nextRight[r][c] = nxt
		}
		prv := -1
		for c := 0; c < W; c++ {
			if grid[r][c] == 1 { prv = c }
			prevLeft[r][c] = prv
		}
	}

	nextDown := make([][]int, W)
	prevUp := make([][]int, W)
	for c := 0; c < W; c++ {
		nextDown[c] = make([]int, H)
		prevUp[c] = make([]int, H)
		nxt := H
		for r := H - 1; r >= 0; r-- {
			if grid[r][c] == 1 { nxt = r }
			nextDown[c][r] = nxt
		}
		prv := -1
		for r := 0; r < H; r++ {
			if grid[r][c] == 1 { prv = r }
			prevUp[c][r] = prv
		}
	}

	costCol := make([][]int, W)
	{
		colHas := make([]bool, W)
		colMin := make([]int, W)
		colMax := make([]int, W)
		for c := 0; c < W; c++ { colMin[c], colMax[c] = H, -1 }
		for c := 0; c < W; c++ {
			for r := 0; r < H; r++ {
				if grid[r][c] == 1 {
					if r < colMin[c] { colMin[c] = r }
					if r > colMax[c] { colMax[c] = r }
				}
			}
			if colMin[c] <= colMax[c] { colHas[c] = true }
		}
		for L := 0; L < W; L++ {
			costCol[L] = make([]int, W)
			top, bot, left, right := H, -1, -1, -1
			any := false
			for R := L; R < W; R++ {
				if colHas[R] {
					if !any { left = R }
					right = R; any = true
					if colMin[R] < top { top = colMin[R] }
					if colMax[R] > bot { bot = colMax[R] }
				}
				if any { costCol[L][R] = (right-left+1)*(bot-top+1) } else { costCol[L][R] = 1 }
			}
		}
	}

	costRow := make([][]int, H)
	{
		rowHas := make([]bool, H)
		rowMin := make([]int, H)
		rowMax := make([]int, H)
		for r := 0; r < H; r++ { rowMin[r], rowMax[r] = W, -1 }
		for r := 0; r < H; r++ {
			for c := 0; c < W; c++ {
				if grid[r][c] == 1 {
					if c < rowMin[r] { rowMin[r] = c }
					if c > rowMax[r] { rowMax[r] = c }
				}
			}
			if rowMin[r] <= rowMax[r] { rowHas[r] = true }
		}
		for T := 0; T < H; T++ {
			costRow[T] = make([]int, H)
			left, right, top, bot := W, -1, -1, -1
			any := false
			for B := T; B < H; B++ {
				if rowHas[B] {
					if !any { top = B }
					bot = B; any = true
					if rowMin[B] < left { left = rowMin[B] }
					if rowMax[B] > right { right = rowMax[B] }
				}
				if any { costRow[T][B] = (right-left+1)*(bot-top+1) } else { costRow[T][B] = 1 }
			}
		}
	}

	best2H := make([][]int, W)
	for L := 0; L < W; L++ {
		best2H[L] = make([]int, W)
		for R := L; R < W; R++ {
			pref := make([]int, H)
			suf := make([]int, H)
			any := false
			top, bot, left, right := H, -1, W, -1
			for r := 0; r < H; r++ {
				lp := nextRight[r][L]
				if lp <= R {
					any = true
					if top == H { top = r }
					bot = r
					if lp < left { left = lp }
					if prevLeft[r][R] > right { right = prevLeft[r][R] }
				}
				if any { pref[r] = (bot-top+1)*(right-left+1) } else { pref[r] = 1 }
			}
			any = false; top, bot, left, right = H, -1, W, -1
			for r := H - 1; r >= 0; r-- {
				lp := nextRight[r][L]
				if lp <= R {
					any = true
					if bot == -1 { bot = r }
					top = r
					if lp < left { left = lp }
					if prevLeft[r][R] > right { right = prevLeft[r][R] }
				}
				if any { suf[r] = (bot-top+1)*(right-left+1) } else { suf[r] = 1 }
			}
			best := INF
			for r := 0; r+1 < H; r++ {
				if v := pref[r] + suf[r+1]; v < best { best = v }
			}
			best2H[L][R] = best
		}
	}

	best2V := make([][]int, H)
	for T := 0; T < H; T++ {
		best2V[T] = make([]int, H)
		for B := T; B < H; B++ {
			pref := make([]int, W)
			suf := make([]int, W)
			any := false
			l, rgt, top, bot := -1, -1, H, -1
			for c := 0; c < W; c++ {
				fr := nextDown[c][T]
				if fr <= B {
					any = true
					if l == -1 { l = c }
					rgt = c
					if fr < top { top = fr }
					if prevUp[c][B] > bot { bot = prevUp[c][B] }
				}
				if any { pref[c] = (rgt-l+1)*(bot-top+1) } else { pref[c] = 1 }
			}
			any = false; l, rgt, top, bot = -1, -1, H, -1
			for c := W - 1; c >= 0; c-- {
				fr := nextDown[c][T]
				if fr <= B {
					any = true
					if rgt == -1 { rgt = c }
					l = c
					if fr < top { top = fr }
					if prevUp[c][B] > bot { bot = prevUp[c][B] }
				}
				if any { suf[c] = (rgt-l+1)*(bot-top+1) } else { suf[c] = 1 }
			}
			best := INF
			for c := 0; c+1 < W; c++ {
				if v := pref[c] + suf[c+1]; v < best { best = v }
			}
			best2V[T][B] = best
		}
	}

	three := func(cost [][]int, N int) int {
		dp1 := make([]int, N)
		for i := 0; i < N; i++ { dp1[i] = cost[0][i] }
		dp2 := make([]int, N)
		for i := range dp2 { dp2[i] = INF }
		for i := 0; i < N; i++ {
			for j := 0; j < i; j++ {
				if v := dp1[j] + cost[j+1][i]; v < dp2[i] { dp2[i] = v }
			}
		}
		dp3 := make([]int, N)
		for i := range dp3 { dp3[i] = INF }
		for i := 0; i < N; i++ {
			for j := 0; j < i; j++ {
				if dp2[j] < INF {
					if v := dp2[j] + cost[j+1][i]; v < dp3[i] { dp3[i] = v }
				}
			}
		}
		return dp3[N-1]
	}

	ans := three(costCol, W)
	if v := three(costRow, H); v < ans { ans = v }

	for c := 0; c+1 < W; c++ {
		if v := costCol[0][c] + best2H[c+1][W-1]; v < ans { ans = v }
		if v := best2H[0][c] + costCol[c+1][W-1]; v < ans { ans = v }
	}
	for r := 0; r+1 < H; r++ {
		if v := costRow[0][r] + best2V[r+1][H-1]; v < ans { ans = v }
		if v := best2V[0][r] + costRow[r+1][H-1]; v < ans { ans = v }
	}
	return ans
}
