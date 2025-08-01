package minesweeper

import "strconv"

// Annotate returns an annotated board
func Annotate(board []string) []string {
    h := len(board)

    if h == 0 {
        return board
    }
    
    w := len(board[0])

    if w == 0 {
        return board
    }

	res := make([]string, h)
    
    dir := [8][2]int{
        {-1,-1},
    	{-1,0},
        {-1,1},
        {0,-1},
        {0,1},
        {1,-1},
        {1,0},
        {1,1},
    }
           
	for r, row := range board {
        s := ""
        for c, char := range row {
			if char == '*' {
                s += string(char)
                continue
            }
            
            count := 0
            for _, d := range dir {
                if isWithin(w, h, c+d[1], r+d[0]) {
                    if board[r+d[0]][c+d[1]] == '*' {
                        count ++
                    }
                }
            }
            
            if count > 0 {
            	s += strconv.Itoa(count)
            } else {
                s += string(char)
            }
        }
        res[r] = s
    }
    
    return res
}

func isWithin(width int, height int, x int, y int) bool {
    return x >= 0 && y >= 0 && x < width && y < height
}
    