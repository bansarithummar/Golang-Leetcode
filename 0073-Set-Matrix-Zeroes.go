func setZeroes(matrix [][]int)  {
    m, n := len(matrix), len(matrix[0])
    firstRowZero := false
    firstColZero := false

    // Check if the first row should be zero
    for j := 0; j < n; j++ {
        if matrix[0][j] == 0 {
            firstRowZero = true
            break
        }
    }

    // Check if the first column should be zero
    for i := 0; i < m; i++ {
        if matrix[i][0] == 0 {
            firstColZero = true
            break
        }
    }

    // Mark zeroes in the first row and column
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            if matrix[i][j] == 0 {
                matrix[i][0] = 0
                matrix[0][j] = 0
            }
        }
    }

    // Set rows to zero based on markers
    for i := 1; i < m; i++ {
        if matrix[i][0] == 0 {
            for j := 1; j < n; j++ {
                matrix[i][j] = 0
            }
        }
    }

    // Set columns to zero based on markers
    for j := 1; j < n; j++ {
        if matrix[0][j] == 0 {
            for i := 1; i < m; i++ {
                matrix[i][j] = 0
            }
        }
    }

    if firstRowZero {
        for j := 0; j < n; j++ {
            matrix[0][j] = 0
        }
    }

    if firstColZero {
        for i := 0; i < m; i++ {
            matrix[i][0] = 0
        }
    }   
}
